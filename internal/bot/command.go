package bot

import "context"

type Command interface {
	Definition() CommandDefinition
	Handle(ctx context.Context, in Invocation) Response
}

type CommandDefinition struct {
	Name        string
	Description string
	Options     []OptionDefinition
}

type OptionDefinition struct {
	Name        string
	Description string
	Type        OptionType
	Required    bool
	Options     []OptionDefinition // For subcommands and groups
}

type OptionType int

const (
	OptionTypeUnknown OptionType = iota
	OptionTypeSubCommand
	OptionTypeSubCommandGroup
	OptionTypeString
	OptionTypeInteger
	OptionTypeBoolean
	OptionTypeUser
	OptionTypeChannel
	OptionTypeRole
	OptionTypeMentionable
	OptionTypeNumber
	OptionTypeAttachment
)

type Invocation struct {
	CommandName    string
	GuildID        string
	ChannelID      string
	UserID         string
	Username       string
	Options        map[string]any
	SubCommand     string
	SubCommandGroup string
}

type Response struct {
	Content   string
	Ephemeral bool
}
