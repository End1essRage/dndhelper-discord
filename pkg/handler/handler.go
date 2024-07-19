package handler

import (
	"strings"

	client "github.com/end1essrage/dndhelper-discord/pkg/client"
	commands "github.com/end1essrage/dndhelper-discord/pkg/commands"
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
	sb := strings.Builder{}
	sb.WriteString("spellinfo - позволяет получить информацию о заклинании (для подробной информации =spellinfo -help)")
	return sb.String()
}

func (h *Handler) getSpellInfo(spellName string, params map[string]string) (string, error) {

	_, hExists := params["h"]
	_, exists := params["help"]
	if exists || hExists {
		return commands.HelpSpellInfo(), nil
	}

	var format formatter.Formatter

	switch params["display"] {
	case "min":
		format = formatter.NewMinimalisticFormatter()
	case "max":
		format = formatter.NewSimpleFormatter()
	default:
		format = formatter.NewSimpleFormatter()
	}

	command := commands.NewSpellInfoCommand(spellName, params["lang"], format)

	return command.Execute(h.client)
}
