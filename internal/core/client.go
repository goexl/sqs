package core

import (
	"context"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/goexl/sqs/internal/internal"
	"github.com/goexl/sqs/internal/param"
)

type Client struct {
	sqs        *sqs.Client
	urls       *sync.Map
	attributes *sync.Map
	params     *param.Client
}

func NewClient(params *param.Client) (client *Client) {
	client = new(Client)
	client.params = params
	client.urls = new(sync.Map)
	client.attributes = new(sync.Map)

	options := sqs.Options{}
	options.Credentials = aws.NewCredentialsCache(params.Credentials)
	options.Region = params.Region
	options.HTTPClient = internal.NewClient(params.Http)
	client.sqs = sqs.New(options)

	return
}

func (c *Client) Stop(_ context.Context) (err error) {
	if nil == c || nil == c.params {
		return
	}

	c.params.Exit = true
	if nil != c.params.Cancel {
		c.params.Cancel()
	}

	return
}
