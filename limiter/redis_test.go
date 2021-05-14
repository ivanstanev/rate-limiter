package limiter_test

import (
	"testing"
	"time"

	miniredis "github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/ivanstanev/rate-limiter/algorithm"
	"github.com/ivanstanev/rate-limiter/backend"
	"github.com/ivanstanev/rate-limiter/limiter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRedisRateLimiterShouldLimit(t *testing.T) {
	miniredis, err := miniredis.Run()
	require.Nil(t, err, "miniredis could not start")
	defer miniredis.Close()

	client := redis.NewClient(&redis.Options{
		Addr: miniredis.Addr(),
	})

	t.Run("when within limits", func(t *testing.T) {
		algorithm := algorithm.NewFixedWindowCounter(1, time.Minute)
		backend := backend.NewRedisBackend(client)
		rl := limiter.NewRedisRateLimiter(algorithm, backend)

		calc, err := rl.Evaluate("Boo")
		got := calc.ShouldLimit
		want := false
		assert.Equal(t, want, got, "Rate limiting should not be applied")
		assert.Nil(t, err, "Rate limiting should not have errors")
	})

	t.Run("when exceeding limits", func(t *testing.T) {
		algorithm := algorithm.NewFixedWindowCounter(1, time.Minute)
		backend := backend.NewRedisBackend(client)
		rl := limiter.NewRedisRateLimiter(algorithm, backend)

		calc, err := rl.Evaluate("Far")
		got := calc.ShouldLimit
		want := false
		assert.Equal(t, want, got, "Rate limiting should not be applied")
		assert.Nil(t, err, "Rate limiting should not have errors")

		calc, err = rl.Evaluate("Far")
		got = calc.ShouldLimit
		want = true
		assert.Equal(t, want, got, "Rate limiting should be applied")
		assert.Nil(t, err, "Rate limiting should not have errors")
	})
}

var calculation limiter.Calculation

func BenchmarkRedisRateLimiterShouldLimit(t *testing.B) {
	miniredis, err := miniredis.Run()
	require.Nil(t, err, "miniredis could not start")
	defer miniredis.Close()

	client := redis.NewClient(&redis.Options{
		Addr: miniredis.Addr(),
	})

	var calc limiter.Calculation
	t.Run("performance", func(t *testing.B) {
		algorithm := algorithm.NewFixedWindowCounter(t.N, time.Minute)
		backend := backend.NewRedisBackend(client)
		rl := limiter.NewRedisRateLimiter(algorithm, backend)

		for i := 0; i < t.N; i++ {
			calc, _ = rl.Evaluate("")
		}

		calculation = calc
	})
}
