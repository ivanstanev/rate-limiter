package limiter_test

import (
	"testing"
	"time"

	miniredis "github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	limiter "github.com/ivanstanev/rate-limiter/limiter"
	redisRateLimiter "github.com/ivanstanev/rate-limiter/redis"
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
		config := &limiter.Configuration{
			Window: limiter.Window{
				Tokens:      1,
				RefreshRate: time.Minute,
			},
		}

		rl := redisRateLimiter.NewRateLimiter(client, config)
		got, err := rl.ShouldLimit("Boo")
		want := false

		assert.Equal(t, want, got, "Rate limiting should not be applied")
		assert.Nil(t, err, "Rate limiting should not have errors")
	})

	t.Run("when exceeding limits", func(t *testing.T) {
		config := &limiter.Configuration{
			Window: limiter.Window{
				Tokens:      1,
				RefreshRate: time.Minute,
			},
		}

		rl := redisRateLimiter.NewRateLimiter(client, config)
		got, err := rl.ShouldLimit("Far")
		want := false
		assert.Equal(t, want, got, "Rate limiting should not be applied")
		assert.Nil(t, err, "Rate limiting should not have errors")

		got, err = rl.ShouldLimit("Far")
		want = true
		assert.Equal(t, want, got, "Rate limiting should be applied")
		assert.Nil(t, err, "Rate limiting should not have errors")
	})
}

var result bool

func BenchmarkRedisRateLimiterShouldLimit(t *testing.B) {
	miniredis, err := miniredis.Run()
	require.Nil(t, err, "miniredis could not start")
	defer miniredis.Close()

	client := redis.NewClient(&redis.Options{
		Addr: miniredis.Addr(),
	})

	var res bool
	t.Run("performance", func(t *testing.B) {
		config := &limiter.Configuration{
			Window: limiter.Window{
				Tokens:      uint(t.N),
				RefreshRate: time.Minute,
			},
		}
		rl := redisRateLimiter.NewRateLimiter(client, config)

		for i := 0; i < t.N; i++ {
			res, _ = rl.ShouldLimit("")
		}

		result = res
	})
}
