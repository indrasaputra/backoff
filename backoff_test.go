package backoff_test

import (
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

func TestConstantBackoff_Interval(t *testing.T) {
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
		b := e.NextInterval()
		assert.Equal(t, backoffInterval, b)

		e.Reset()
		b = e.NextInterval()
		assert.Equal(t, backoffInterval, b)
	})
}
