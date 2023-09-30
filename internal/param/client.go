package param

import (
	"time"

	"github.com/goexl/simaqian"
)

type Client struct {
	*Provider

	Region string
	Wait   time.Duration
	Queues map[string]*string
	Logger simaqian.Logger
}

func NewClient() *Client {
	return &Client{
		Provider: NewProvider(),

		Wait:   20 * time.Second,
		Queues: make(map[string]*string),
		Logger: simaqian.Default(),
	}
}
