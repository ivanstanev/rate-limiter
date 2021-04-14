package limiter

import (
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	limiter "github.com/ivanstanev/rate-limiter/limiter"
	"github.com/stretchr/testify/assert"
)

func TestRedisRateLimiterShouldLimit(t *testing.T) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	rlConfig := &limiter.RateLimiterConfiguration{
		Window: limiter.Window{
			Tokens:      0,
			RefreshRate: time.Minute,
		},
	}
	rl := &RedisRateLimiter{
		RedisClient:   redisClient,
		Configuration: rlConfig,
	}
	got := rl.ShouldLimit("Boo")
	want := false

	assert.Equal(t, want, got, "Rate limiting should not be applied")
}

func TestGetCurrentWindow(t *testing.T) {
	getCurrentWindow(&limiter.RateLimiterConfiguration{})
}
