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
	Interval(order int) time.Duration
}
