package backoff

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Intervaler is an interface that defines method to be used when implementing interval
// of two subsequent time.
type Intervaler interface {
	// Interval defines the n-th interval.
	// Please, note that the valid order (n value) is a positive integer.
	// If non-positive integer (including zero) is given, the method should return 0,
	// which indicates that there is no interval.
	Interval(order int) time.Duration
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
