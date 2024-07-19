package client

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/spf13/viper"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

func NewClient(host string) *Client {
	return &Client{
		host:     host,
		basePath: viper.GetString("api_basepath"),
		client:   http.Client{},
	}
}

func (c *Client) FormatBaseUrl(rout string) url.URL {
	return url.URL{
		Scheme: viper.GetString("api_sheme"),
		Host:   c.host,
		Path:   path.Join(c.basePath, rout),
	}
}

func (c *Client) DoRequest(req *http.Request) ([]byte, error) {
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
