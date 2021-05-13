package backend

import (
	"context"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	limiter "github.com/ivanstanev/rate-limiter/limiter"
)

type RedisBackend struct {
	limiter.Backend
	redisClient *redis.Client
}

func NewRedisBackend(redisClient *redis.Client) *RedisBackend {
	return &RedisBackend{
		redisClient: redisClient,
	}
}

func (rb *RedisBackend) Set(key string, value int, expiration time.Duration) error {
	setCtx, setCancel := context.WithDeadline(context.Background(), time.Now().Add(time.Minute))
	defer setCancel()

	return rb.redisClient.Set(setCtx, key, strconv.Itoa(value), expiration).Err()
}

func (rb *RedisBackend) Get(key string) (int, error) {
	getCtx, getCancel := context.WithDeadline(context.Background(), time.Now().Add(time.Minute))
	defer getCancel()

	token, err := rb.redisClient.Get(getCtx, key).Result()
	if err == redis.Nil {
		token = "0"
	} else if err != nil {
		return 0, err
	}

	return strconv.Atoi(token)
}
