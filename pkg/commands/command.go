package commands

import (
	"encoding/json"
	"fmt"
	"net/http"

	client "github.com/end1essrage/dndhelper-discord/pkg/client"
	formatter "github.com/end1essrage/dndhelper-discord/pkg/helpers"
	t "github.com/end1essrage/dndhelper-discord/pkg/types"
	"github.com/sirupsen/logrus"
)

type Command interface {
	Execute(client *client.Client) (string, error)
}

type CommandBase struct {
	apiRout string
	method  string
}

type GetSpellInfoCommand struct {
	CommandBase
	spellName    string
	localization string
	format       formatter.Formatter
}

func (c *GetSpellInfoCommand) Execute(client *client.Client) (string, error) {
	u := client.FormatBaseUrl(c.apiRout)
	req, err := http.NewRequest(c.method, u.String(), nil)

	if err != nil {
		logrus.Error("Error creating request")
	}

	q := req.URL.Query()
	q.Add("name", c.spellName)
	//q.Add("localization", c.localization)
	req.URL.RawQuery = q.Encode()

	resp, err := client.DoRequest(req)

	if err != nil {
		return "", fmt.Errorf("can't do request: %w", err)
	}

	var spell t.Spell
	err = json.Unmarshal(resp, &spell)

	if err != nil {
		return "", fmt.Errorf("can't unmarshall: %w", err)
	}

	return c.format.FormatSpellInfo(spell), nil
}
