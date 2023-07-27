package core

import (
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/goexl/sqs/internal/param"
)

type Client struct {
	sqs   *sqs.Client
	urls  *sync.Map
	param *param.Client
}

func NewClient(param *param.Client) (client *Client) {
	client = new(Client)
	client.param = param
	client.urls = new(sync.Map)

	options := sqs.Options{}
	options.Credentials = aws.NewCredentialsCache(param.Credentials)
	options.Region = param.Region
	client.sqs = sqs.New(options)

	return
}
