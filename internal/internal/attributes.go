package internal

import (
	"time"

	"github.com/goexl/sqs/internal/constant"
)

type Attributes struct {
	values map[string]string
}

func NewAttributes(values map[string]string) *Attributes {
	return &Attributes{
		values: values,
	}
}

func (a *Attributes) Invalid(sent time.Time) bool {
	return time.Now().Sub(sent)+a.Visibility()*2 > a.Period()
}

func (a *Attributes) Period() time.Duration {
	return a.getDuration(constant.KeyPeriod, constant.Zero)
}

func (a *Attributes) Visibility() time.Duration {
	return a.getDuration(constant.KeyVisibility, constant.Zero)
}

func (a *Attributes) Delay() time.Duration {
	return a.getDuration(constant.KeyDelay, constant.Zero)
}

func (a *Attributes) getDuration(key string, def time.Duration) (duration time.Duration) {
	if value, ok := a.values[key]; ok {
		duration = a.parseDuration(value, def)
	} else {
		duration = def
	}

	return
}

func (a *Attributes) parseDuration(value string, def time.Duration) (duration time.Duration) {
	if parsed, pde := time.ParseDuration(value); nil != pde {
		duration = def
	} else {
		duration = parsed
	}

	return
}
