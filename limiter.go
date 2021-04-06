package limiter

import (
	"github.com/go-redis/redis/v8"
	limiter "github.com/ivanstanev/rate-limiter/internal"
)

type RateLimiter interface {
	ShouldLimit(key string) bool
}

func NewRedisRateLimiter(redisClient *redis.Client) RateLimiter {
	return &limiter.RedisRateLimiter{
		RedisClient: redisClient,
	}
}

func NewInMemoryRateLimiter() RateLimiter {
	return &limiter.InMemoryRateLimiter{}
}
