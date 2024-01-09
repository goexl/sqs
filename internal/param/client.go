package param

import (
	"context"
	"time"

	"github.com/goexl/http"
	"github.com/goexl/log"
)

type Client struct {
	*Provider

	Region string
	Wait   time.Duration
	Queues map[string]*string
	Http   *http.Client
	Logger log.Logger
	Exit   bool
	Cancel context.CancelFunc
}

func NewClient() *Client {
	return &Client{
		Provider: NewProvider(),

		Wait:   20 * time.Second,
		Queues: make(map[string]*string),
		Http:   http.New().Build(),
		Logger: log.New().Apply(),
	}
}
