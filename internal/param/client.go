package param

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
)

type Client struct {
	Region   string
	Provider aws.CredentialsProvider
	Wait     time.Duration
	Queues   map[string]*string
}

func NewClient() *Client {
	return &Client{
		Wait:   20 * time.Second,
		Queues: make(map[string]*string),
	}
}
