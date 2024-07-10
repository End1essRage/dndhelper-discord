package client

//универсальный клиент

import (
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

func NewClient(host string) *Client {
	return &Client{
		host:     host,
		basePath: "/api",
		client:   http.Client{},
	}
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
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
