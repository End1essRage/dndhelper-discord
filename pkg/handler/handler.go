package handler

import (
	"encoding/json"
	"fmt"

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

func (h *Handler) Handle(command string, value string, params map[string]string) (string, error) {
	switch command {
	case "spellinfo":
		return h.getSpellInfo(value, params)
	case "help":
		return h.getHelpMessage(), nil
	default:
		return h.getHelpMessage(), nil
	}
}

func (h *Handler) getHelpMessage() string {
	return "help"
}

// Реализовать автодокуиентирование доступных команд
func (h *Handler) getSpellInfo(spellName string, params map[string]string) (string, error) {

	command := factory.NewSpellInfoCommand(h.client, spellName, params["lang"])
	//create request
	req, err := command.Execute()

	if err != nil {
		return "", fmt.Errorf("can't do request: %w", err)
	}
	//send request
	resp, err := h.client.DoRequest(req)

	if err != nil {
		return "", fmt.Errorf("can't do request: %w", err)
	}

	var spell t.Spell
	err = json.Unmarshal(resp, &spell)

	if err != nil {
		return "", fmt.Errorf("can't unmarshall: %w", err)
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

	return format.FormatSpellInfo(spell), nil
}
