package commands

import (
	"net/http"
	"strings"

	client "github.com/end1essrage/dndhelper-discord/pkg/client"
	formatter "github.com/end1essrage/dndhelper-discord/pkg/helpers"
	"github.com/spf13/viper"
)

type Command interface {
	Execute(client *client.Client) (string, error)
}

type CommandBase struct {
	apiRout string
	method  string
}

// Spell Info Command
// TODO Add Summary
/*//////Summary////////////
literal: spellinfo
params:
	lang : EN|RU|...
	display : min|Max
*/
const apiRoute = "spell-book"

type GetSpellInfoCommand struct {
	CommandBase
	spellName    string
	localization string
	format       formatter.Formatter
}

func HelpSpellInfo() string {
	sb := strings.Builder{}
	sb.WriteString("params: \n")
	sb.WriteString("	lang : EN|RU|... \n")
	sb.WriteString("	display : min|Max \n")
	return sb.String()
}

func NewSpellInfoCommand(spellName, localization string, format formatter.Formatter) *GetSpellInfoCommand {
	cmd := GetSpellInfoCommand{}
	cmd.apiRout = apiRoute
	cmd.format = format
	cmd.method = http.MethodGet
	cmd.spellName = spellName
	if localization != "" {
		cmd.localization = localization
	} else {
		cmd.localization = viper.GetString("localization_default")
	}

	return &cmd
}
