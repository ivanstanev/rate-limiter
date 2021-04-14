package limiter

import (
	"github.com/go-redis/redis/v8"
	memoryRateLimiter "github.com/ivanstanev/rate-limiter/internal/memory"
	redisRateLimiter "github.com/ivanstanev/rate-limiter/internal/redis"
	rateLimiter "github.com/ivanstanev/rate-limiter/limiter"
)

func NewRedisRateLimiter(redisClient *redis.Client, config *rateLimiter.RateLimiterConfiguration) rateLimiter.RateLimiter {
	return &redisRateLimiter.RedisRateLimiter{
		RedisClient:   redisClient,
		Configuration: config,
	}
}

func NewInMemoryRateLimiter() rateLimiter.RateLimiter {
	return &memoryRateLimiter.InMemoryRateLimiter{}
}
