package handler

import (
	"strings"

	client "github.com/end1essrage/dndhelper-discord/pkg/client"
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

	var sb strings.Builder

	//считываем параметр
	sb.WriteString("lang is " + params["lang"])

	spell, err := h.client.GetSpellInfo(spellName)
	if err != nil {
		sb.WriteString("ERROR OCCURED: ")
		sb.WriteString(err.Error())
		return sb.String(), err
	}

	sb.WriteString("Spell Name : " + spell.Name)
	sb.WriteString("\n")

	sb.WriteString("Description: ")
	for i := 0; i < len(spell.Desc); i++ {
		sb.WriteString(spell.Desc[i])
		sb.WriteString("\n")
	}

	sb.WriteString("Damage is " + spell.Damage.DamageType.Name + "\n")
	for key, value := range spell.Damage.DamageAtSlotLevel {
		sb.WriteString(key + " - " + value + "\n")
	}

	return sb.String(), nil
}
