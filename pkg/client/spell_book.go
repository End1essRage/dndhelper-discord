package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
)

const rout = "spell-book"

func (c *Client) GetSpellInfo(spellName string) (*Spell, error) {
	u := url.URL{
		Scheme: "http",
		Host:   c.host,
		Path:   path.Join(c.basePath, rout),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	q := req.URL.Query()
	q.Add("name", spellName)

	req.URL.RawQuery = q.Encode()

	if err != nil {
		return nil, fmt.Errorf("can't do request: %w", err)
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, fmt.Errorf("can't do request: %w", err)
	}

	var spell Spell
	err = json.Unmarshal(resp, &spell)

	return &spell, err
}
