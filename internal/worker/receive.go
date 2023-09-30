package worker

import (
	"context"

	"github.com/goexl/sqs/internal/output"
	"github.com/goexl/sqs/internal/param"
)

type Receive struct {
	param *param.Receive
}

func NewReceive(param *param.Receive) *Receive {
	return &Receive{
		param: param,
	}
}

func (r *Receive) Do(ctx context.Context) (out *output.Receive, err error) {
	if url, ue := r.param.Url(ctx, r.param.Base); nil != ue {
		err = ue
	} else {
		out, err = r.do(ctx, url)
	}

	return
}

func (r *Receive) do(ctx context.Context, url *string) (out *output.Receive, err error) {
	return r.param.Do(ctx, url)
}
