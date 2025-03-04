package rate_limiter

import (
	"context"
	"testing"
	"time"
)

func TestRateLimiter_Wait(t *testing.T) {
	ctx := context.Background()
	rps := 5
	rl := NewRateLimiter(rps)
	defer rl.Stop()

	start := time.Now()
	err := rl.Wait(ctx)
	duration := time.Since(start)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedMin := time.Second / time.Duration(rps)
	if duration < expectedMin {
		t.Errorf("Expected at least %v, but got %v", expectedMin, duration)
	}
}

func TestRateLimiter_WaitWithContextCancel(t *testing.T) {
	rl := NewRateLimiter(1)
	defer rl.Stop()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	start := time.Now()
	err := rl.Wait(ctx)
	duration := time.Since(start)

	if err == nil {
		t.Fatal("Expected context cancellation error, got nil")
	}

	if duration > 50*time.Millisecond {
		t.Errorf("Wait took too long: expected ~10ms, got %v", duration)
	}
}

func TestRateLimiter_Stop(t *testing.T) {
	rl := NewRateLimiter(2) // 2 запроса в секунду
	rl.Stop()

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	err := rl.Wait(ctx)

	if err == nil {
		t.Fatal("Expected error due to stopped ticker, got nil")
	}
}
