package discord

import (
	"context"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/erickir/LandOfLifeDiscordBot/internal/bot"
	"github.com/erickir/LandOfLifeDiscordBot/internal/config"
)

type Client struct {
	session *discordgo.Session
	guildID string
	router  *bot.Router
}

func NewClient(configuration *config.Config, router *bot.Router) (*Client, error) {
	s, err := discordgo.New("Bot " + configuration.TokenID)
	if err != nil {
		return nil, fmt.Errorf("create discord session error: %w", err)
	}

	c := &Client{
		session: s,
		guildID: configuration.GuildID,
		router:  router,
	}

	c.session.AddHandler(c.onInteractionCreate)

	return c, nil
}

func (c *Client) onInteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	data := i.ApplicationCommandData()
	username := ""
	if i.Member != nil && i.Member.User != nil {
		username = i.Member.User.Username
	} else if i.User != nil {
		username = i.User.Username
	}
	log.Printf("Received interaction: command=%q user=%q guild=%q channel=%q\n", data.Name, username, i.GuildID, i.ChannelID)

	in := fromDiscordInteraction(i)

	resp := c.router.Dispatch(context.Background(), in)
	log.Printf("Dispatched command=%q response=%q ephemeral=%t\n", in.CommandName, resp.Content, resp.Ephemeral)

	flags := discordgo.MessageFlags(0)
	if resp.Ephemeral {
		flags = discordgo.MessageFlagsEphemeral
	}

	if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: resp.Content,
			Flags:   flags,
		},
	}); err != nil {
		log.Printf("Failed to send interaction response: %v", err)
	}
}

func fromDiscordInteraction(i *discordgo.InteractionCreate) bot.Invocation {
	data := i.ApplicationCommandData()

	opts := make(map[string]any)
	subCommand := ""
	subCommandGroup := ""
	for _, opt := range data.Options {
		extractDiscordOptions(opt, opts, &subCommand, &subCommandGroup)
	}

	username := ""
	userID := ""

	if i.Member != nil && i.Member.User != nil {
		username = i.Member.User.Username
		userID = i.Member.User.ID
	} else if i.User != nil {
		username = i.User.Username
		userID = i.User.ID
	}

	return bot.Invocation{
		CommandName:     data.Name,
		GuildID:         i.GuildID,
		ChannelID:       i.ChannelID,
		UserID:          userID,
		Username:        username,
		Options:         opts,
		SubCommand:      subCommand,
		SubCommandGroup: subCommandGroup,
	}
}

func extractDiscordOptions(option *discordgo.ApplicationCommandInteractionDataOption, opts map[string]any, subCommand, subCommandGroup *string) {
	switch option.Type {
	case discordgo.ApplicationCommandOptionSubCommandGroup:
		*subCommandGroup = option.Name
		for _, child := range option.Options {
			extractDiscordOptions(child, opts, subCommand, subCommandGroup)
		}
	case discordgo.ApplicationCommandOptionSubCommand:
		*subCommand = option.Name
		for _, child := range option.Options {
			extractDiscordOptions(child, opts, subCommand, subCommandGroup)
		}
	default:
		if len(option.Options) == 0 {
			opts[option.Name] = option.Value
			return
		}
		for _, child := range option.Options {
			extractDiscordOptions(child, opts, subCommand, subCommandGroup)
		}
	}
}

func (c *Client) Open(ctx context.Context) error {
	if err := c.session.Open(); err != nil {
		return err
	}

	for _, def := range c.router.Definitions() {
		cmd := toDiscordCommand(def)
		log.Printf("Registering command: name=%q description=%q guild=%q\n", def.Name, def.Description, c.guildID)

		registered, err := c.session.ApplicationCommandCreate(
			c.session.State.User.ID,
			c.guildID, // use "" for global, guild ID for dev
			cmd,
		)
		if err != nil {
			return fmt.Errorf("register command %s: %w", def.Name, err)
		}
		log.Printf("Registered command %q with ID %q\n", def.Name, registered.ID)
	}

	log.Println("Bot is now running. Press CTRL-C to exit.")
	<-ctx.Done()
	return nil
}

func toDiscordCommand(def bot.CommandDefinition) *discordgo.ApplicationCommand {
	options := make([]*discordgo.ApplicationCommandOption, 0, len(def.Options))

	for _, opt := range def.Options {
		options = append(options, toDiscordOption(opt))
	}

	return &discordgo.ApplicationCommand{
		Name:        def.Name,
		Description: def.Description,
		Options:     options,
	}
}

func toDiscordOption(opt bot.OptionDefinition) *discordgo.ApplicationCommandOption {
	newOpt := &discordgo.ApplicationCommandOption{
		Type:        toDiscordOptionType(opt.Type),
		Name:        opt.Name,
		Description: opt.Description,
		Required:    opt.Required,
	}

	if len(opt.Options) > 0 {
		newOpt.Options = make([]*discordgo.ApplicationCommandOption, 0, len(opt.Options))
		for _, child := range opt.Options {
			newOpt.Options = append(newOpt.Options, toDiscordOption(child))
		}
	}

	return newOpt
}

func toDiscordOptionType(t bot.OptionType) discordgo.ApplicationCommandOptionType {
	switch t {
	case bot.OptionTypeString:
		return discordgo.ApplicationCommandOptionString
	case bot.OptionTypeInteger:
		return discordgo.ApplicationCommandOptionInteger
	case bot.OptionTypeBoolean:
		return discordgo.ApplicationCommandOptionBoolean
	case bot.OptionTypeSubCommand:
		return discordgo.ApplicationCommandOptionSubCommand
	case bot.OptionTypeSubCommandGroup:
		return discordgo.ApplicationCommandOptionSubCommandGroup
	default:
		return discordgo.ApplicationCommandOptionString
	}
}

func (c *Client) Close() error {
	return c.session.Close()
}
