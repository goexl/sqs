package param

import (
	"time"

	"github.com/goexl/log"
)

type Client struct {
	*Provider

	Region string
	Wait   time.Duration
	Queues map[string]*string
	Logger log.Logger
	Exit   bool
}

func NewClient() *Client {
	return &Client{
		Provider: NewProvider(),

		Wait:   20 * time.Second,
		Queues: make(map[string]*string),
		Logger: log.New().Apply(),
	}
}
