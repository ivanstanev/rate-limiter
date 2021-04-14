package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	limiter "github.com/ivanstanev/rate-limiter"
	rateLimiter "github.com/ivanstanev/rate-limiter/limiter"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	rlConfig := &rateLimiter.RateLimiterConfiguration{
		Window: rateLimiter.Window{
			Tokens:      0,
			RefreshRate: time.Minute,
		},
	}
	rateLimiter := limiter.NewRedisRateLimiter(redisClient, rlConfig)

	// cancellationCtx := context.WithDeadline(context.TODO(), time.Now().Add(time.Second*5))

	if rateLimiter.ShouldLimit("foo") {
		panic("Got rate limited!")
	} else {
		fmt.Println("Bye!")
	}
}
