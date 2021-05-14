package limiter_test

import (
	"testing"
	"time"

	"github.com/ivanstanev/rate-limiter/algorithm"
	"github.com/ivanstanev/rate-limiter/backend"
	"github.com/ivanstanev/rate-limiter/limiter"
	"github.com/stretchr/testify/assert"
)

func TestInMemoryRateLimiterShouldLimit(t *testing.T) {
	algorithm := algorithm.NewFixedWindowCounter(1, time.Minute)
	backend := backend.NewMemoryBackend()
	rl := limiter.NewInMemoryRateLimiter(algorithm, backend)
	calc, err := rl.Evaluate("Boo")
	got := calc.ShouldLimit
	want := false

	assert.Equal(t, want, got, "Rate limiting should not be applied")
	assert.Nil(t, err, "Rate limiting should not have errors")
}
