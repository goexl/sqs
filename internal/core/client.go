package core

import (
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/goexl/sqs/internal/param"
)

type Client struct {
	client *sqs.Client
	urls   sync.Map
	param  *param.Client
}

func NewClient(param *param.Client) (client *Client) {
	client = new(Client)
	client.param = param

	options := sqs.Options{}
	options.Credentials = aws.NewCredentialsCache(param.Provider)
	options.Region = param.Region
	client.client = sqs.New(options)

	return
}
