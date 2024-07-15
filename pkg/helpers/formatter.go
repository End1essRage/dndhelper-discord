package helpers

//Обработка флага форматирования сообщений

import (
	"strings"

	t "github.com/end1essrage/dndhelper-discord/pkg/types"
)

type Formatter interface {
	FormatSpellInfo(t.Spell) string
}

type SimpleFormatter struct {
}

type MinimalisticFormatter struct {
}

func NewSimpleFormatter() *SimpleFormatter {
	return &SimpleFormatter{}
}

func NewMinimalisticFormatter() *MinimalisticFormatter {
	return &MinimalisticFormatter{}
}

func (f *SimpleFormatter) FormatSpellInfo(spell t.Spell) string {
	var sb strings.Builder

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

	return sb.String()
}

func (f *MinimalisticFormatter) FormatSpellInfo(spell t.Spell) string {
	var sb strings.Builder

	sb.WriteString("Spell Name : " + spell.Name)
	sb.WriteString("\n")

	sb.WriteString("Damage is " + spell.Damage.DamageType.Name + "\n")
	for key, value := range spell.Damage.DamageAtSlotLevel {
		sb.WriteString(key + " - " + value + "\n")
	}

	return sb.String()
}
