package limiter

import limiter "github.com/ivanstanev/rate-limiter/limiter"

type InMemoryRateLimiter struct {
	limiter.RateLimiter
	Configuration *limiter.RateLimiterConfiguration
}

func (rl *InMemoryRateLimiter) ShouldLimit(key string) bool {
	return false
}
