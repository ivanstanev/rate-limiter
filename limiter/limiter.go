package limiter

import (
	"time"
)

type RateLimiter interface {
	Evaluate(key string) (Calculation, error)
	ShouldLimit(key string) (bool, error)
}

type Calculation struct {
	ShouldLimit bool
	RetryAfter  time.Duration
}

type Algorithm interface {
	Configuration
}

type Backend interface {
	Set(key string, value int, expiration time.Duration) error
	Get(key string) (int, error)
}

type Configuration interface {
	GetTokens() int
	GetRefreshRate() time.Duration
}
