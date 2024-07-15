package commands

import (
	"net/http"

	client "github.com/end1essrage/dndhelper-discord/pkg/client"
	"github.com/sirupsen/logrus"
)

type Command interface {
	Execute() (string, error)
}

type CommandBase struct {
	client  *client.Client
	apiRout string
	method  string
}

type GetSpellInfoCommand struct {
	CommandBase
	spellName    string
	localization string
}

func (c *GetSpellInfoCommand) Execute() (*http.Request, error) {
	u := c.client.FormatBaseUrl(c.apiRout)

	req, err := http.NewRequest(c.method, u.String(), nil)

	if err != nil {
		logrus.Error("Error creating request")
	}

	q := req.URL.Query()
	q.Add("name", c.spellName)

	req.URL.RawQuery = q.Encode()

	return req, nil
}
