package internal

import (
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type Base struct {
	Label string
	fns   []func(*sqs.Options)
}
