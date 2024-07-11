package client

import (
	"encoding/json"
	"fmt"

	"net/http"
	"net/url"
	"path"

	t "github.com/end1essrage/dndhelper-discord/pkg/types"
)

// Через интерфейс и наследование вынести как отдельный клиент
const rout = "spell-book"

// Добавить ДТО для параметризации запроса
func (c *Client) GetSpellInfo(spellName string) (*t.Spell, error) {
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

	var spell t.Spell
	err = json.Unmarshal(resp, &spell)

	return &spell, err
}
