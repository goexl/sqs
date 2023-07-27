package callback

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type ReceiveMessage func(ctx context.Context, params *sqs.ReceiveMessageInput, optFns ...func(*sqs.Options)) (*sqs.ReceiveMessageOutput, error)
