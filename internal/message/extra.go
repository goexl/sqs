package message

import (
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

type Extra struct {
	Attributes map[string]string
	Messages   map[string]types.MessageAttributeValue
	Id         *string
	Handle     *string
}
