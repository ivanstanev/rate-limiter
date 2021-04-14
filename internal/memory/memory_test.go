package limiter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInMemoryRateLimiterShouldLimit(t *testing.T) {
	rl := InMemoryRateLimiter{}
	got := rl.ShouldLimit("Boo")
	want := false

	assert.Equal(t, want, got, "Rate limiting should not be applied")
}
