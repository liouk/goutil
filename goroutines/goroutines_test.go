package goroutines

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimedOut(t *testing.T) {
	timedOut := WaitTimeout(2*time.Millisecond, func() {
		time.Sleep(3 * time.Millisecond)
	})

	assert.True(t, timedOut)
}

func TestNoTimeout(t *testing.T) {
	timedOut := WaitTimeout(2*time.Millisecond, func() {
		time.Sleep(1 * time.Millisecond)
	})

	assert.False(t, timedOut)
}
