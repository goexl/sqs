package callback

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type ChangeMessageVisibility func(ctx context.Context, params *sqs.ChangeMessageVisibilityInput, optFns ...func(*sqs.Options)) (*sqs.ChangeMessageVisibilityOutput, error)
