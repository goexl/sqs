package internal

import (
	"github.com/goexl/http"
)

type Client struct {
	http *http.Client
}

func NewClient(http *http.Client) *Client {
	return &Client{
		http: http,
	}
}
