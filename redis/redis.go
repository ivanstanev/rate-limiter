package limiter

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	limiter "github.com/ivanstanev/rate-limiter/limiter"
)

type RedisRateLimiter struct {
	limiter.RateLimiter
	RedisClient   *redis.Client
	Configuration *limiter.Configuration
}

func NewRateLimiter(redisClient *redis.Client, config *limiter.Configuration) limiter.RateLimiter {
	return &RedisRateLimiter{
		RedisClient:   redisClient,
		Configuration: config,
	}
}

func (rl *RedisRateLimiter) ShouldLimit(key string) (bool, error) {
	currentWindow := getCurrentWindow(rl.Configuration)

	keyPrefix := currentWindow.Format(time.Stamp)
	redisKey := fmt.Sprintf("%s %s", keyPrefix, key)

	getCtx, getCancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	defer getCancel()

	tokens, err := rl.RedisClient.Get(getCtx, redisKey).Result()
	if err == redis.Nil {
		tokens = "0"
	} else if err != nil {
		return false, err
	}

	i, err := strconv.Atoi(tokens)
	if err != nil {
		return false, err
	}

	setCtx, setCancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	defer setCancel()

	newTokens := i + 1
	err = rl.RedisClient.Set(setCtx, redisKey, strconv.Itoa(newTokens), rl.Configuration.Window.RefreshRate).Err()
	if err != nil {
		return false, err
	}

	shouldLimit := int64(newTokens) > int64(rl.Configuration.Window.Tokens)
	return shouldLimit, nil
}

func getCurrentWindow(configuration *limiter.Configuration) time.Time {
	return time.Now().Truncate(configuration.Window.RefreshRate)
}
