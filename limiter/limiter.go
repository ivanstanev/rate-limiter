package limiter

import (
	"time"
)

type RateLimiter interface {
	ShouldLimit(key string) (bool, error)
}

type Window struct {
	Tokens      uint          `default:"10"`
	RefreshRate time.Duration `default:"1m"`
}

type Configuration struct {
	Window Window
}
