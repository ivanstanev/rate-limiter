package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/ivanstanev/rate-limiter/limiter"
	redisRateLimiter "github.com/ivanstanev/rate-limiter/redis"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	rlConfig := &limiter.Configuration{
		Window: limiter.Window{
			Tokens:      1,
			RefreshRate: time.Minute,
		},
	}
	rl := redisRateLimiter.NewRateLimiter(redisClient, rlConfig)

	if shouldLimit, err := rl.ShouldLimit("foo"); shouldLimit && err != nil {
		panic("Got rate limited!")
	} else {
		fmt.Println("Bye!")
	}
}
