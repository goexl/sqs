package builder

import (
	"github.com/goexl/sqs/internal/callback"
	"github.com/goexl/sqs/internal/worker"
)

type Url struct {
	*Base

	url callback.Url
}

func NewUrl(url callback.Url) *Url {
	return &Url{
		Base: NewBase(),

		url: url,
	}
}

func (u *Url) Build() *worker.Url {
	return worker.NewUrl(u.param, u.url)
}
