package commands

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/erickir/LandOfLifeDiscordBot/internal/bot"
	"github.com/erickir/LandOfLifeDiscordBot/internal/core"
	"github.com/erickir/LandOfLifeDiscordBot/internal/domain"
	"github.com/google/uuid"
)

type LandOfLifeCommand struct {
	service *core.RegistrationService
}

func NewLandOfLifeCommand(service *core.RegistrationService) *LandOfLifeCommand {
	return &LandOfLifeCommand{service: service}
}

func (c *LandOfLifeCommand) Definition() bot.CommandDefinition {
	return bot.CommandDefinition{
		Name:        "lol",
		Description: "Retrieve Land of Life information and manage your registrations",
		Options: []bot.OptionDefinition{
			{
				Name:        "slot",
				Description: "View, register and unregister Land of Life slots",
				Type:        bot.OptionTypeSubCommandGroup,
				Options: []bot.OptionDefinition{
					{
						Name:        "register",
						Description: "Register for a Land of Life slot",
						Type:        bot.OptionTypeSubCommand,
						Options: []bot.OptionDefinition{
							{Name: "username", Description: "Your in-game username", Type: bot.OptionTypeString, Required: true},
							{Name: "hour", Description: "Slot hour (01:00, 03:00, ... or 1,3,5)", Type: bot.OptionTypeString, Required: true},
							{Name: "channel", Description: "Channel number (1-7)", Type: bot.OptionTypeInteger, Required: true},
							{Name: "class", Description: "Character class", Type: bot.OptionTypeString, Required: true},
							{Name: "level", Description: "Player level", Type: bot.OptionTypeString, Required: true},
							{Name: "pet", Description: "Pet name", Type: bot.OptionTypeString, Required: true},
						},
					},
					{
						Name:        "unregister",
						Description: "Remove your registration from a Land of Life slot",
						Type:        bot.OptionTypeSubCommand,
						Options: []bot.OptionDefinition{
							{Name: "username", Description: "Your in-game username", Type: bot.OptionTypeString, Required: true},
							{Name: "hour", Description: "Slot hour to remove (01:00, 03:00, ... or 1,3,5)", Type: bot.OptionTypeString, Required: true},
							{Name: "channel", Description: "Channel number (1-7)", Type: bot.OptionTypeInteger, Required: true},
						},
					},
					{
						Name:        "status",
						Description: "List Land of Life's today slots",
						Type:        bot.OptionTypeSubCommand,
						Options: []bot.OptionDefinition{
							{Name: "hour", Description: "Slot hour (01:00, 03:00, ... or 1,3,5)", Type: bot.OptionTypeString, Required: false},
						},
					},
				},
			},
		},
	}
}

func (c *LandOfLifeCommand) Handle(ctx context.Context, in bot.Invocation) bot.Response {
	if in.SubCommandGroup != "slot" {
		return bot.Response{Content: "Unknown command group.", Ephemeral: true}
	}

	switch in.SubCommand {
	case "register":
		return c.handleRegister(ctx, in)
	case "unregister":
		return c.handleUnregister(ctx, in)
	case "status":
		return c.handleStatus(ctx, in)
	default:
		return bot.Response{Content: "Unknown slot action.", Ephemeral: true}
	}
}

func (c *LandOfLifeCommand) handleRegister(ctx context.Context, in bot.Invocation) bot.Response {
	username := parseStringOption(in.Options, "username")
	if username == "" {
		return bot.Response{Content: "Username is required.", Ephemeral: true}
	}

	hour, err := parseHourOption(in.Options, "hour")
	if err != nil {
		return bot.Response{Content: err.Error(), Ephemeral: true}
	}

	channel, err := parseChannelOption(in.Options, "channel")
	if err != nil {
		return bot.Response{Content: err.Error(), Ephemeral: true}
	}

	class := parseStringOption(in.Options, "class")
	if class == "" {
		return bot.Response{Content: "Class is required.", Ephemeral: true}
	}

	level := parseStringOption(in.Options, "level")
	if level == "" {
		return bot.Response{Content: "Level is required.", Ephemeral: true}
	}

	pet := parseStringOption(in.Options, "pet")
	if pet == "" {
		return bot.Response{Content: "Pet is required.", Ephemeral: true}
	}

	registration := domain.Registration{
		ID:            uuid.NewString(),
		EventHour:     hour,
		Channel:       channel,
		DiscordUserID: in.UserID,
		Username:      username,
		Class:         class,
		Level:         level,
		Pet:           pet,
	}

	if err := c.service.Register(ctx, registration); err != nil {
		switch {
		case errors.Is(err, domain.ErrInvalidHour):
			return bot.Response{Content: "Invalid hour. Use 01:00, 03:00, ... or 1,3,5.", Ephemeral: true}
		case errors.Is(err, domain.ErrInvalidChannel):
			return bot.Response{Content: "Invalid channel. Enter a number between 1 and 7.", Ephemeral: true}
		case errors.Is(err, domain.ErrSlotFull):
			return bot.Response{Content: "This slot is full. Please choose another channel or hour.", Ephemeral: true}
		case errors.Is(err, domain.ErrAlreadyRegistered):
			return bot.Response{Content: "You are already registered for this hour.", Ephemeral: true}
		default:
			return bot.Response{Content: fmt.Sprintf("Could not register slot: %v", err), Ephemeral: true}
		}
	}

	return bot.Response{Content: fmt.Sprintf("Registration complete: %s | %s | %s", hour.String(), channel.String(), username), Ephemeral: false}
}

func (c *LandOfLifeCommand) handleUnregister(ctx context.Context, in bot.Invocation) bot.Response {
	username := parseStringOption(in.Options, "username")
	if username == "" {
		return bot.Response{Content: "Username is required.", Ephemeral: true}
	}

	hour, err := parseHourOption(in.Options, "hour")
	if err != nil {
		return bot.Response{Content: err.Error(), Ephemeral: true}
	}

	channel, err := parseChannelOption(in.Options, "channel")
	if err != nil {
		return bot.Response{Content: err.Error(), Ephemeral: true}
	}

	if err := c.service.Deregister(ctx, username, hour, channel); err != nil {
		switch {
		case errors.Is(err, domain.ErrInvalidHour):
			return bot.Response{Content: "Invalid hour. Use 01:00, 03:00, ... or 1,3,5.", Ephemeral: true}
		case errors.Is(err, domain.ErrInvalidChannel):
			return bot.Response{Content: "Invalid channel. Enter a number between 1 and 7.", Ephemeral: true}
		case errors.Is(err, domain.ErrRegistrationNotFound):
			return bot.Response{Content: fmt.Sprintf("No registration found for %s at %s channel %d.", username, hour.String(), int(channel)), Ephemeral: true}
		default:
			return bot.Response{Content: fmt.Sprintf("Could not unregister slot: %v", err), Ephemeral: true}
		}
	}

	return bot.Response{Content: fmt.Sprintf("Unregistered %s from %s | %s", username, hour.String(), channel.String()), Ephemeral: false}
}

func (c *LandOfLifeCommand) handleStatus(ctx context.Context, in bot.Invocation) bot.Response {
	hour := domain.UnknownHour
	hourRaw := parseStringOption(in.Options, "hour")
	if hourRaw != "" {
		parsedHour, err := parseHourOption(in.Options, "hour")
		if err != nil {
			return bot.Response{Content: err.Error(), Ephemeral: true}
		}
		hour = parsedHour
	}

	date := time.Now().Format("2006-01-02")
	slots, err := c.service.ListRegistrationForSlot(ctx, date)
	if err != nil {
		return bot.Response{Content: fmt.Sprintf("Could not fetch slots: %v", err), Ephemeral: true}
	}

	filteredSlots := make([]domain.Slot, 0, len(slots))
	for _, slot := range slots {
		if hour != domain.UnknownHour && slot.Hour != hour {
			continue
		}

		filteredSlots = append(filteredSlots, slot)
	}

	lines := []string{fmt.Sprintf("Slots for %s:", date)}
	counted := 0
	for _, slot := range filteredSlots {
		if hour == domain.UnknownHour && len(slot.Registrations) == 0 {
			continue
		}

		if len(slot.Registrations) == 0 {
			lines = append(lines, fmt.Sprintf("- %s @ %s | 0/%d", slot.Channel.String(), slot.Hour.String(), slot.Capacity))
			continue
		}

		counted++
		usernames := make([]string, 0, len(slot.Registrations))
		for _, reg := range slot.Registrations {
			usernames = append(usernames, fmt.Sprintf("%s (%s) - %s (%s)", reg.Username, reg.Level, reg.Class, reg.Pet))
		}

		lines = append(lines, fmt.Sprintf("- %s @ %s | %d/%d | users: %s", slot.Channel.String(), slot.Hour.String(), len(slot.Registrations), slot.Capacity, strings.Join(usernames, ", ")))
	}

	if counted == 0 {
		return bot.Response{Content: "No registrations, feel free to register in any LoL slot.", Ephemeral: false}
	}

	response := strings.Join(lines, "\n")

	if len(response) > 2000 {
		response = fmt.Sprintf("There are %d registrations for %s, but the message is too long to display. Try fetching status for specific hours instead.", counted, date)
	}

	return bot.Response{Content: response, Ephemeral: false}
}
