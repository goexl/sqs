package exception

import (
	"fmt"
	"time"
)

type Delay struct {
	delay *time.Duration
}

func NewDelay(delay *time.Duration) *Delay {
	return &Delay{
		delay: delay,
	}
}

func (d *Delay) Error() string {
	return fmt.Sprintf(`{"delay": %s}`, d.delay.String())
}

func (d *Delay) Duration() time.Duration {
	return *d.delay
}
