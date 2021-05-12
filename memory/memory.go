package limiter

import "github.com/ivanstanev/rate-limiter/limiter"

type InMemoryRateLimiter struct {
	limiter.RateLimiter
	Configuration *limiter.Configuration
}

func NewRateLimiter() limiter.RateLimiter {
	return &InMemoryRateLimiter{}
}

func (rl *InMemoryRateLimiter) ShouldLimit(key string) (bool, error) {
	return false, nil
}
