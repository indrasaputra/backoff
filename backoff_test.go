package backoff_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/indrasaputra/backoff"
	"github.com/stretchr/testify/assert"
)

func TestConstantBackoff_Interval(t *testing.T) {
	backoffInterval := 500 * time.Millisecond
	jitterInterval := 200 * time.Millisecond
	c := backoff.ConstantBackoff{
		BackoffInterval: backoffInterval,
		JitterInterval:  jitterInterval,
	}

	for i := -100; i < 1; i++ {
		t.Run("test:"+strconv.Itoa(i), func(t *testing.T) {
			b := c.Interval(i)
			assert.Equal(t, time.Duration(0), b)
		})
	}

	for i := 1; i < 100; i++ {
		t.Run("test:"+strconv.Itoa(i), func(t *testing.T) {
			b := c.Interval(i)
			assert.True(t, b >= backoffInterval)
			assert.True(t, b < backoffInterval+jitterInterval)
		})
	}
}
