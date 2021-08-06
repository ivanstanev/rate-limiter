package algorithm

import (
	"time"

	limiter "github.com/ivanstanev/rate-limiter/limiter"
)

type FixedWindowCounter struct {
	limiter.Algorithm

	tokens      int
	refreshRate time.Duration
}

func NewFixedWindowCounter(tokens int, refreshRate time.Duration) limiter.Algorithm {
	return &FixedWindowCounter{
		tokens:      tokens,
		refreshRate: refreshRate,
	}
}

func (alg *FixedWindowCounter) GetTokens() int {
	return alg.tokens
}

func (alg *FixedWindowCounter) GetRefreshRate() time.Duration {
	return alg.refreshRate
}
