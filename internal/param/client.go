package param

import (
	"time"

	"github.com/goexl/sqs/internal/core"
)

type Client struct {
	Region   string
	Provider core.CredentialProvider
	Wait     time.Duration
}

func NewClient() *Client {
	return &Client{
		Wait: 20 * time.Second,
	}
}
