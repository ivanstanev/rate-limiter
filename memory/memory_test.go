package limiter_test

import (
	"testing"

	limiter "github.com/ivanstanev/rate-limiter/memory"
	"github.com/stretchr/testify/assert"
)

func TestInMemoryRateLimiterShouldLimit(t *testing.T) {
	rl := limiter.NewRateLimiter()
	got, err := rl.ShouldLimit("Boo")
	want := false

	assert.Equal(t, want, got, "Rate limiting should not be applied")
	assert.Nil(t, err, "Rate limiting should not have errors")
}
