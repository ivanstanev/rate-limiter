package limiter

import "github.com/go-redis/redis/v8"

type RedisRateLimiter struct {
	RedisClient *redis.Client
}

func (rl *RedisRateLimiter) ShouldLimit(key string) bool {
	return false
}
