package handler

import (
	"strings"

	c "github.com/end1essrage/dndhelper-discord/pkg"
	client "github.com/end1essrage/dndhelper-discord/pkg/client"
	commands "github.com/end1essrage/dndhelper-discord/pkg/commands"
	formatter "github.com/end1essrage/dndhelper-discord/pkg/helpers"
	t "github.com/end1essrage/dndhelper-discord/pkg/types"
)

type Handler struct {
	client *client.Client
	env    string
}

func NewHandler(client *client.Client, env string) *Handler {
	return &Handler{client: client, env: env}
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

	if h.env == c.ENV_LOCAL {
		format = formatter.NewDevFormatter()
	}

	if h.env == c.ENV_DEV {
		switch params["display"] {
		case "min":
			format = formatter.NewSimpleFormatter()
		case "max":
			format = formatter.NewSimpleFormatter()
		default:
			format = formatter.NewSimpleFormatter()
		}
	}

	command := commands.NewSpellInfoCommand(spellName, params["lang"], format)

	return command.Execute(h.client)
}
