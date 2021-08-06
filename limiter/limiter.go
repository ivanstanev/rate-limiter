package limiter

import (
	"fmt"
	"time"
)

type Limiter struct {
	RateLimiter
	algorithm Algorithm
	backend   Backend
}

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

func New(algorithm Algorithm, backend Backend) RateLimiter {
	return &Limiter{
		algorithm: algorithm,
		backend:   backend,
	}
}

func (rl Limiter) ShouldLimit(key string) (bool, error) {
	evaluation, err := rl.Evaluate(key)
	if err != nil {
		return false, err
	}

	return evaluation.ShouldLimit, nil
}

func (rl Limiter) Evaluate(key string) (Calculation, error) {
	now := time.Now()
	currentWindow := now.Truncate(rl.algorithm.GetRefreshRate())
	nextWindow := currentWindow.Add(rl.algorithm.GetRefreshRate())

	keyPrefix := currentWindow.Format(time.Stamp)
	redisKey := fmt.Sprintf("%s %s", keyPrefix, key)

	tokens, err := rl.backend.Get(redisKey)
	if err != nil {
		panic(err)
	}
	defer func(newTokens int) {
		err := rl.backend.Set(redisKey, newTokens, rl.algorithm.GetRefreshRate())
		if err != nil {
			panic(err)
		}
	}(tokens + 1)

	shouldLimit := tokens >= rl.algorithm.GetTokens()
	retryAfter := nextWindow.Sub(now)

	return Calculation{
		ShouldLimit: shouldLimit,
		RetryAfter:  retryAfter,
	}, nil
}
