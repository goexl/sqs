package callback

import (
	"context"

	"github.com/goexl/sqs/internal/core"
)

type GetAttributes func(ctx context.Context, url *string) (*core.Attributes, error)
