package limiter

type InMemoryRateLimiter struct {
	algorithm Algorithm
	backend   Backend
}

func NewInMemoryRateLimiter(algorithm Algorithm, backend Backend) RateLimiter {
	return &InMemoryRateLimiter{
		algorithm: algorithm,
		backend:   backend,
	}
}

func (rl *InMemoryRateLimiter) ShouldLimit(key string) (bool, error) {
	evaluation, err := rl.Evaluate(key)
	if err != nil {
		return false, err
	}

	return evaluation.ShouldLimit, nil
}

func (rl *InMemoryRateLimiter) Evaluate(key string) (Calculation, error) {
	return Calculation{}, nil
}
