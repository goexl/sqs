package callback

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type DeleteMessage func(ctx context.Context, params *sqs.DeleteMessageInput, optFns ...func(*sqs.Options)) (*sqs.DeleteMessageOutput, error)
