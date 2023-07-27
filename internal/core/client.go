package core

import (
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/goexl/sqs/internal/param"
)

type Client struct {
	sqs    *sqs.Client
	urls   *sync.Map
	queues map[string]*string
	param  *param.Client
}

func NewClient(param *param.Client) (client *Client) {
	client = new(Client)
	client.param = param
	client.urls = new(sync.Map)
	client.queues = make(map[string]*string)

	options := sqs.Options{}
	options.Credentials = aws.NewCredentialsCache(param.Provider)
	options.Region = param.Region
	client.sqs = sqs.New(options)

	return
}
