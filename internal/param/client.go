package param

import (
	"time"
)

type Client struct {
	*Provider

	Region string
	Wait   time.Duration
	Queues map[string]*string
}

func NewClient() *Client {
	return &Client{
		Provider: NewProvider(),

		Wait:   20 * time.Second,
		Queues: make(map[string]*string),
	}
}
