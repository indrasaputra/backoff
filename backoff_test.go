package backoff_test

import (
	"math"
	"strconv"
	"testing"
	"time"

	"github.com/indrasaputra/backoff"
	"github.com/stretchr/testify/assert"
)

// tests for constant backoff

func TestConstantBackoff_(t *testing.T) {
	c := &backoff.ConstantBackoff{}
	assert.Implements(t, (*backoff.Backoff)(nil), c)
}

func TestConstantBackoff_NextInterval(t *testing.T) {
	backoffInterval := 500 * time.Millisecond
	jitterInterval := 100 * time.Millisecond
	c1 := backoff.ConstantBackoff{
		BackoffInterval: backoffInterval,
		JitterInterval:  jitterInterval,
	}

	// test without reset
	for i := 0; i < 50; i++ {
		t.Run("constant backoff with jitter :"+strconv.Itoa(i), func(t *testing.T) {
			b := c1.NextInterval()
			assert.True(t, b >= backoffInterval)
			assert.True(t, b < backoffInterval+jitterInterval)
		})
	}

	// test with reset
	for i := 0; i < 50; i++ {
		t.Run("constant backoff with jitter :"+strconv.Itoa(i), func(t *testing.T) {
			b := c1.NextInterval()
			assert.True(t, b >= backoffInterval)
			assert.True(t, b < backoffInterval+jitterInterval)

			c1.Reset()
			b = c1.NextInterval()
			assert.True(t, b >= backoffInterval)
			assert.True(t, b < backoffInterval+jitterInterval)
		})
	}

	backoffInterval = 100 * time.Millisecond
	c2 := backoff.ConstantBackoff{
		BackoffInterval: backoffInterval,
	}

	// test without reset
	for i := 0; i < 50; i++ {
		t.Run("constant backoff without jitter :"+strconv.Itoa(i), func(t *testing.T) {
			b := c2.NextInterval()
			assert.Equal(t, backoffInterval, b)
		})
	}

	// test with reset
	for i := 0; i < 50; i++ {
		t.Run("constant backoff without jitter :"+strconv.Itoa(i), func(t *testing.T) {
			b := c2.NextInterval()
			assert.Equal(t, backoffInterval, b)

			c2.Reset()
			b = c2.NextInterval()
			assert.Equal(t, backoffInterval, b)
		})
	}
}

func TestConstantBackoff_Reset(t *testing.T) {
	c := backoff.ConstantBackoff{}
	assert.NotPanics(t, func() { c.Reset() })
}

// tests for exponential backoff

func TestExponentialBackoff_(t *testing.T) {
	b := &backoff.ExponentialBackoff{}
	assert.Implements(t, (*backoff.Backoff)(nil), b)
}

func TestExponentialBackoff_Reset(t *testing.T) {
	backoffInterval := 100 * time.Millisecond
	e := backoff.ExponentialBackoff{
		BackoffInterval: backoffInterval,
	}

	t.Run("exponential backoff reset", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			e.NextInterval()
		}

		e.Reset()
		b := e.NextInterval()
		assert.Equal(t, backoffInterval, b)
	})
}

func TestExponentialBackoff_NextInterval(t *testing.T) {
	backoffInterval := 100 * time.Millisecond
	jitterInterval := 50 * time.Millisecond
	multiplier := 2
	maxInterval := 10 * time.Second

	t.Run("exponential backoff only backoff interval", func(t *testing.T) {
		e := backoff.ExponentialBackoff{
			BackoffInterval: backoffInterval,
		}

		for i := 0; i < 100; i++ {
			b := e.NextInterval()
			assert.Equal(t, backoffInterval, b)
		}
	})

	t.Run("use backoff interval and jitter", func(t *testing.T) {
		e := backoff.ExponentialBackoff{
			BackoffInterval: backoffInterval,
			JitterInterval:  jitterInterval,
		}

		for i := 0; i < 100; i++ {
			b := e.NextInterval()
			assert.True(t, b >= backoffInterval)
			assert.True(t, b < backoffInterval+jitterInterval)
		}
	})

	t.Run("use backoff, jitter, and multiplier", func(t *testing.T) {
		e := backoff.ExponentialBackoff{
			BackoffInterval: backoffInterval,
			JitterInterval:  jitterInterval,
			Multiplier:      multiplier,
		}

		// iterate 30 times should be enough
		// more than 30 can be overflow
		for i := 0; i < 30; i++ {
			m := math.Pow(float64(multiplier), float64(i))

			b := e.NextInterval()
			assert.True(t, b >= backoffInterval*time.Duration(m))
			assert.True(t, b < backoffInterval*time.Duration(m)+jitterInterval)
		}
	})

	t.Run("use backoff, jitter, multiplier, and max interval", func(t *testing.T) {
		e := backoff.ExponentialBackoff{
			BackoffInterval: backoffInterval,
			JitterInterval:  jitterInterval,
			Multiplier:      multiplier,
			MaxInterval:     maxInterval,
		}

		// iterate 30 times should be enough
		// more than 30 can be overflow
		for i := 0; i < 30; i++ {
			b := e.NextInterval()
			assert.True(t, b <= maxInterval+jitterInterval)
		}
	})
}
