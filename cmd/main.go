package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/ivanstanev/rate-limiter/algorithm"
	"github.com/ivanstanev/rate-limiter/backend"
	"github.com/ivanstanev/rate-limiter/limiter"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	algorithm := algorithm.NewFixedWindowCounter(1, time.Minute)
	backend := backend.NewRedisBackend(redisClient)
	rl := limiter.NewRedisRateLimiter(algorithm, backend)

	if shouldLimit, err := rl.ShouldLimit("foo"); shouldLimit && err != nil {
		panic("Got rate limited!")
	} else {
		fmt.Println("Bye!")
	}
}
