package limiter

import (
	"testing"

	limiter "github.com/ivanstanev/rate-limiter/limiter"
	"github.com/stretchr/testify/assert"
)

func TestInMemoryRateLimiterShouldLimit(t *testing.T) {
	rl := RedisRateLimiter{}
	got := rl.ShouldLimit("Boo")
	want := false

	assert.Equal(t, want, got, "Rate limiting should not be applied")
}

func TestGetCurrentWindow(t *testing.T) {
	getCurrentWindow(&limiter.RateLimiterConfiguration{})
}
