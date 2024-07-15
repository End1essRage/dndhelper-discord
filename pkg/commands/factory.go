package commands

import (
	"net/http"

	formatter "github.com/end1essrage/dndhelper-discord/pkg/helpers"
)

//Переделать на билдер возможно?

func NewSpellInfoCommand(spellName, localization string, format formatter.Formatter) *GetSpellInfoCommand {
	cmd := GetSpellInfoCommand{}
	cmd.apiRout = "spell-book"
	cmd.format = format
	cmd.method = http.MethodGet
	cmd.spellName = spellName
	if localization != "" {
		cmd.localization = localization
	} else {
		cmd.localization = "EN"
	}

	return &cmd
}
