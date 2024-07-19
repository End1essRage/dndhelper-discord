package commands

import (
	"encoding/json"
	"fmt"
	"net/http"

	client "github.com/end1essrage/dndhelper-discord/pkg/client"
	t "github.com/end1essrage/dndhelper-discord/pkg/types"
	"github.com/sirupsen/logrus"
)

func (c *GetSpellInfoCommand) Execute(client *client.Client) (string, error) {
	u := client.FormatBaseUrl(c.apiRout)
	req, err := http.NewRequest(c.method, u.String(), nil)

	if err != nil {
		logrus.Error("Error creating request")
	}

	req.URL.RawQuery = c.encodeQuery(req)

	logrus.Info("raw query is : " + req.URL.RawQuery)

	resp, err := client.DoRequest(req)
	if err != nil {
		return "", fmt.Errorf("can't do request: %w", err)
	}

	var spell t.Spell
	err = json.Unmarshal(resp, &spell)
	if err != nil {
		return "", fmt.Errorf("can't unmarshall response: %w", err)
	}

	return c.format.FormatSpellInfo(spell), nil
}

func (c *GetSpellInfoCommand) encodeQuery(req *http.Request) string {
	q := req.URL.Query()
	q.Add("name", c.spellName)
	q.Add("localization", c.localization)
	return q.Encode()
}
