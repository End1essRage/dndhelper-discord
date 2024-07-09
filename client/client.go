package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

func NewClient(host string) *Client {
	return &Client{
		host:     host,
		basePath: "/api/spell-book",
		client:   http.Client{},
	}
}

func (c *Client) GetSpellInfo(spellName string) (*Spell, error) {
	resp, err := c.sendRequest(spellName)

	var spell Spell
	err = json.Unmarshal(resp, &spell)

	return &spell, err
}

func (c *Client) sendRequest(spellname string) ([]byte, error) {
	u := url.URL{
		Scheme: "http",
		Host:   c.host,
		Path:   c.basePath,
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	q := req.URL.Query()
	q.Add("name", spellname)

	req.URL.RawQuery = q.Encode()

	if err != nil {
		return nil, fmt.Errorf("can't do request: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("can't do request: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("can't read response: %w", err)
	}

	return body, nil
}
