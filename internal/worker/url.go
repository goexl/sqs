package worker

import (
	"context"

	"github.com/goexl/sqs/internal/callback"
	"github.com/goexl/sqs/internal/internal"
)

type Url struct {
	param *internal.Base
	url   callback.Url
}

func NewUrl(param *internal.Base, url callback.Url) *Url {
	return &Url{
		param: param,
		url:   url,
	}
}

func (u *Url) Do(ctx context.Context) (*string, error) {
	return u.url(ctx, u.param)
}
