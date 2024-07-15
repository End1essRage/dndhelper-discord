package handler

import (
	client "github.com/end1essrage/dndhelper-discord/pkg/client"
	factory "github.com/end1essrage/dndhelper-discord/pkg/commands"
	formatter "github.com/end1essrage/dndhelper-discord/pkg/helpers"
	t "github.com/end1essrage/dndhelper-discord/pkg/types"
)

type Handler struct {
	client *client.Client
}

func NewHandler(client *client.Client) *Handler {
	return &Handler{client: client}
}

func (h *Handler) Handle(command t.BotCommand) (string, error) {
	switch command.Command {
	case "spellinfo":
		return h.getSpellInfo(command.Value, command.Params)
	case "help":
		return h.getHelpMessage(), nil
	default:
		return h.getHelpMessage(), nil
	}
}

func (h *Handler) getHelpMessage() string {
	return "help"
}

func (h *Handler) getSpellInfo(spellName string, params map[string]string) (string, error) {
	var format formatter.Formatter

	switch params["display"] {
	case "min":
		format = formatter.NewMinimalisticFormatter()
	case "max":
		format = formatter.NewSimpleFormatter()
	default:
		format = formatter.NewSimpleFormatter()
	}

	command := factory.NewSpellInfoCommand(spellName, params["lang"], format)

	return command.Execute(h.client)
}
