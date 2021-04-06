package limiter

type InMemoryRateLimiter struct{}

func (rl *InMemoryRateLimiter) ShouldLimit(key string) bool {
	return false
}
