package backend

import (
	"context"
	"time"

	limiter "github.com/ivanstanev/rate-limiter/limiter"
)

type storageValue struct {
	context    context.Context
	cancelFunc context.CancelFunc
	value      int
}

type MemoryBackend struct {
	limiter.Backend

	storage map[string]storageValue
}

func NewMemoryBackend() *MemoryBackend {
	return &MemoryBackend{
		storage: make(map[string]storageValue),
	}
}

func (mb *MemoryBackend) Set(key string, value int, expiration time.Duration) error {
	mb.storage[key].cancelFunc()

	ctx, cancelFunc := context.WithTimeout(context.Background(), expiration)
	mb.storage[key] = storageValue{
		context:    ctx,
		cancelFunc: cancelFunc,
		value:      value,
	}
	return nil
}

func (mb *MemoryBackend) Get(key string) (int, error) {
	return mb.storage[key].value, nil
}
