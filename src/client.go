package src

import (
	"net/http"
)

type Client struct {
	token      string
	HttpClient *http.Client
}

func NewClient(token string, client ...*http.Client) *Client {
	c := &Client{token: token}
	if len(client) != 0 {
		c.HttpClient = client[0]
	} else {
		c.HttpClient = http.DefaultClient
	}
	return c
}
