package helpers

//Обработка флага форматирования сообщений

import (
	"encoding/json"
	"strings"

	t "github.com/end1essrage/dndhelper-discord/pkg/types"
)

type Formatter interface {
	FormatSpellInfo(t.Spell) string
}

// --------------------------------------------------------------------------
type SimpleFormatter struct {
}

func NewSimpleFormatter() *SimpleFormatter {
	return &SimpleFormatter{}
}

func (f *SimpleFormatter) FormatSpellInfo(spell t.Spell) string {
	var sb strings.Builder

	sb.WriteString("Spell Name : " + spell.Name + "\n")

	sb.WriteString("Description : \n")
	for i := 0; i < len(spell.Desc); i++ {
		sb.WriteString(spell.Desc[i] + "\n")
	}

	sb.WriteString("Damage type is " + spell.Damage.DamageType.Name + "\n")
	for key, value := range spell.Damage.DamageAtSlotLevel {
		sb.WriteString("slot " + key + " - damage " + value + " \n")
	}

	sb.WriteString("Range : " + spell.Range)
	sb.WriteString("School : " + spell.School.Name)

	return sb.String()
}

// --------------------------------------------------------------------------
type MinimalisticFormatter struct {
}

func NewMinimalisticFormatter() *MinimalisticFormatter {
	return &MinimalisticFormatter{}
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

// --------------------------------------------------------------------------
type DefaultFormatter struct {
}

func NewDefaultFormatter() *DefaultFormatter {
	return &DefaultFormatter{}
}

// --------------------------------------------------------------------------
type DevFormatter struct {
}

func NewDevFormatter() *DevFormatter {
	return &DevFormatter{}
}

func (f *DevFormatter) FormatSpellInfo(spell t.Spell) string {
	sSpell, _ := json.Marshal(&spell)

	return string(sSpell)
}
