package callback

import (
	"context"
)

type Url func(ctx context.Context, label string) (url *string, err error)
