package backoff

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Backoff is an interface for backoff-strategy.
type Backoff interface {
	// NextInterval returns an interval before the next process is executed.
	NextInterval() time.Duration
	// Reset resets backoff to its initial state.
	Reset()
}

// ConstantBackoff implements Intervaler using constant interval.
type ConstantBackoff struct {
	// BackoffInterval defines how long the next interval will be, compared to the previous one.
	BackoffInterval time.Duration
	// JitterInterval defines the additional value for interval.
	// The additional value will be in range [0,JitterInterval).
	JitterInterval time.Duration
}

// Interval returns n-th interval.
func (c *ConstantBackoff) Interval(order int) time.Duration {
	if order <= 0 {
		return 0
	}

	jitter := rand.Int63n(int64(c.JitterInterval))
	return time.Duration(c.BackoffInterval + time.Duration(jitter))
}
