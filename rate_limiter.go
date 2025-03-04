package rate_limiter

import (
	"context"
	"time"
)

type RateLimiter struct {
	ticker *time.Ticker
}

func NewRateLimiter(duration time.Duration) *RateLimiter {
	return &RateLimiter{
		ticker: time.NewTicker(duration),
	}
}

func (rl *RateLimiter) Wait(ctx context.Context) error {
	select {
	case <-rl.ticker.C:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (rl *RateLimiter) Stop() {
	rl.ticker.Stop()
}
