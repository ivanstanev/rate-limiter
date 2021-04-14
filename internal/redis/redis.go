package limiter

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	limiter "github.com/ivanstanev/rate-limiter/limiter"
)

type RedisRateLimiter struct {
	limiter.RateLimiter
	RedisClient   *redis.Client
	Configuration *limiter.RateLimiterConfiguration
}

func (rl *RedisRateLimiter) ShouldLimit(key string) bool {
	currentWindow := getCurrentWindow(rl.Configuration)

	keyPrefix := currentWindow.Format(time.Stamp)
	redisKey := fmt.Sprintf("%s %s", keyPrefix, key)

	getCtx, getCancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	defer getCancel()
	val, err := rl.RedisClient.Get(getCtx, redisKey).Result()

	if err == redis.Nil {
		setCtx, setCancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
		defer setCancel()
		err = rl.RedisClient.Set(setCtx, redisKey, "1", rl.Configuration.Window.RefreshRate).Err()
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	}

	fmt.Println("redis value:", val)
	if redisKey == "" {
		return true
	}
	return false
}

func getCurrentWindow(configuration *limiter.RateLimiterConfiguration) time.Time {
	return time.Now().Truncate(configuration.Window.RefreshRate)
}
