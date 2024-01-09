package internal

import (
	"net/http"
)

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.http.Do(req)
}
