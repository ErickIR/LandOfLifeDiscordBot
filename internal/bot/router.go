package bot

import (
	"context"
	"fmt"
)

type Router struct {
	commands map[string]Command
}

func NewRouter(cmds ...Command) (*Router, error) {
	r := &Router{
		commands: make(map[string]Command),
	}

	for _, cmd := range cmds {
		def := cmd.Definition()
		if def.Name == "" {
			return nil, fmt.Errorf("command name cannot be empty")
		}
		if _, exists := r.commands[def.Name]; exists {
			return nil, fmt.Errorf("duplicate command: %s", def.Name)
		}
		r.commands[def.Name] = cmd
	}

	return r, nil
}

func (r *Router) Definitions() []CommandDefinition {
	out := make([]CommandDefinition, 0, len(r.commands))
	for _, cmd := range r.commands {
		out = append(out, cmd.Definition())
	}
	return out
}

func (r *Router) Dispatch(ctx context.Context, in Invocation) Response {
	cmd, ok := r.commands[in.CommandName]
	if !ok {
		return Response{
			Content:   "Unknown command.",
			Ephemeral: true,
		}
	}
	return cmd.Handle(ctx, in)
}
