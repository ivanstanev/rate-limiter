package main

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	limiter "github.com/ivanstanev/rate-limiter"
)

func main() {
	opt, err := redis.ParseURL("redis://localhost:6379/0")
	if err != nil {
		panic(err)
	}

	redisClient := redis.NewClient(opt)
	rateLimiter := limiter.NewRedisRateLimiter(redisClient)

	// cancellationCtx := context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))

	if rateLimiter.ShouldLimit("foo") {
		panic("Got rate limited!")
	} else {
		fmt.Println("Bye!")
	}
}
