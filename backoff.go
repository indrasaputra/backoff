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

// ConstantBackoff implements Backoff using constant interval.
type ConstantBackoff struct {
	// BackoffInterval defines how long the next interval will be, compared to the previous one.
	BackoffInterval time.Duration
	// JitterInterval defines the additional value for interval.
	// The additional value will be in range [0,JitterInterval).
	JitterInterval time.Duration
}

// NextInterval returns next interval.
func (c *ConstantBackoff) NextInterval() time.Duration {
	if c.JitterInterval <= 0 {
		return c.BackoffInterval
	}

	jitter := rand.Int63n(int64(c.JitterInterval))
	return time.Duration(c.BackoffInterval + time.Duration(jitter))
}

// Reset resets Constant Backoff.
// Actually, it does nothing
// since constant backoff will always constant all the time.
func (c *ConstantBackoff) Reset() {
}

// ExponentialBackoff implements Backoff using exponential interval.
type ExponentialBackoff struct {
	// BackoffInterval defines how long the next interval will be, compared to the previous one.
	BackoffInterval time.Duration
	// JitterInterval defines the additional value for interval.
	// The additional value will be in range [0,JitterInterval).
	JitterInterval time.Duration
	// MaxInterval defines the maximum interval allowed.
	// If this field is let empty,
	// it means that there is no maximum value for interval.
	// Please, keep in mind that if this field is empty,
	// the interval can be a very long time.
	MaxInterval time.Duration
	// Multipler defines the multipler for the next interval.
	// Default value for Multiplier is 1.
	// Using `Multiplier = 1` means ExponentialBackoff can behave
	// like ConstantBackoff.
	Multiplier int
	// state defines the state of the backoff.
	// basically, it is just an integer indicating
	// how many times NextInterval is called.
	state int
}
