//go:build !solution

package ratelimit

import (
	"context"
	"errors"
	"time"
)

// Limiter is precise rate limiter with context support.
type Limiter struct{}

var ErrStopped = errors.New("limiter stopped")

// NewLimiter returns limiter that throttles rate of successful Acquire() calls
// to maxCount events at any given interval.
func NewLimiter(maxCount int, interval time.Duration) *Limiter {
	_, _ = maxCount, interval
	return nil // TODO: create the limiter and start its manager.
}

func (l *Limiter) Acquire(ctx context.Context) error {
	_ = ctx
	return nil // TODO: wait for a permit, cancellation, or Stop.
}

func (l *Limiter) Stop() {
}
