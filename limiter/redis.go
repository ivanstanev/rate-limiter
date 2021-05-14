package limiter

import (
	"fmt"
	"time"
)

type RedisRateLimiter struct {
	algorithm Algorithm
	backend   Backend
}

func NewRedisRateLimiter(algorithm Algorithm, backend Backend) RateLimiter {
	return &RedisRateLimiter{
		algorithm: algorithm,
		backend:   backend,
	}
}

func (rl *RedisRateLimiter) ShouldLimit(key string) (bool, error) {
	evaluation, err := rl.Evaluate(key)
	if err != nil {
		return false, err
	}

	return evaluation.ShouldLimit, nil
}

func (rl *RedisRateLimiter) Evaluate(key string) (Calculation, error) {
	now := time.Now()
	currentWindow := now.Truncate(rl.algorithm.GetRefreshRate())
	nextWindow := currentWindow.Add(rl.algorithm.GetRefreshRate())

	keyPrefix := currentWindow.Format(time.Stamp)
	redisKey := fmt.Sprintf("%s %s", keyPrefix, key)

	tokens, err := rl.backend.Get(redisKey)
	if err != nil {
		return Calculation{}, err
	}
	defer func(newTokens int) {
		rl.backend.Set(redisKey, newTokens, rl.algorithm.GetRefreshRate())
	}(tokens + 1)

	shouldLimit := tokens >= rl.algorithm.GetTokens()
	retryAfter := nextWindow.Sub(now)

	return Calculation{
		ShouldLimit: shouldLimit,
		RetryAfter:  retryAfter,
	}, nil
}
