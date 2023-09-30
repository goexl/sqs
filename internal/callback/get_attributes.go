package callback

import (
	"context"

	"github.com/goexl/sqs/internal/internal"
)

type GetAttributes func(ctx context.Context, url *string) (*internal.Attributes, error)
