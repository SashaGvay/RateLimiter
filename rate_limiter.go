package main

import (
	"context"
	"time"
)

type RateLimiter struct {
	ticker *time.Ticker
}

func NewRateLimiter(rps int) *RateLimiter {
	if rps <= 0 {
		panic("rps must be greater than 0")
	}

	return &RateLimiter{
		ticker: time.NewTicker(time.Second / time.Duration(rps)),
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
