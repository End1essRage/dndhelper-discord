package commands

import (
	"net/http"

	client "github.com/end1essrage/dndhelper-discord/pkg/client"
)

//пределать на метод который считывает параметры и рассовывает в поля команды

func NewSpellInfoCommand(client *client.Client, spellName, localization string) *GetSpellInfoCommand {
	cmd := GetSpellInfoCommand{}
	cmd.client = client
	cmd.apiRout = "spell-book"
	cmd.method = http.MethodGet
	cmd.spellName = spellName
	if localization != "" {
		cmd.localization = localization
	} else {
		cmd.localization = "EN"
	}

	return &cmd
}
