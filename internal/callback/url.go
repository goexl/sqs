package callback

import (
	"context"

	"github.com/goexl/sqs/internal/internal"
)

type Url func(ctx context.Context, base *internal.Base) (url *string, err error)
