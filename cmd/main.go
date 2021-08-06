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
	rl := limiter.New(algorithm, backend)

	shouldLimit, err := rl.ShouldLimit("foo")
	if err != nil {
		panic(err)
	}
	if shouldLimit {
		panic("should not have limited")
	}

	shouldLimit, err = rl.ShouldLimit("foo")
	if err != nil {
		panic(err)
	}
	if !shouldLimit {
		panic("should have limited")
	}

	fmt.Println("all good, bye!")
}
