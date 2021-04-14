package limiter

import "time"

type RateLimiter interface {
	ShouldLimit(key string) bool
}

type Window struct {
	Tokens      uint          `default:"10"`
	RefreshRate time.Duration `default:"1m"`
}

type RateLimiterConfiguration struct {
	Window Window
}
