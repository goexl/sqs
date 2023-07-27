package callback

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type SendMessage func(ctx context.Context, params *sqs.SendMessageInput, optFns ...func(*sqs.Options)) (*sqs.SendMessageOutput, error)
