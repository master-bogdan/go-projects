package ratelimiter

import (
	"net/http"
	"sync"
	"time"
)

// internal token bucket
type bucket struct {
	mu       sync.Mutex
	tokens   int
	max      int
	interval time.Duration
	last     time.Time
}

func newBucket(max int, refillInterval time.Duration) *bucket {
	return &bucket{
		tokens:   max,
		max:      max,
		interval: refillInterval,
		last:     time.Now(),
	}
}

func (b *bucket) allow() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(b.last)

	if elapsed >= b.interval {
		b.tokens = b.max
		b.last = now
	}

	if b.tokens > 0 {
		b.tokens--
		return true
	}
	return false
}

// exported RateLimiter wrapper
type RateLimiter struct {
	b *bucket
}

func NewRateLimiter(max int, interval time.Duration) *RateLimiter {
	return &RateLimiter{
		b: newBucket(max, interval),
	}
}

// Middleware wraps an http.Handler
func (rl *RateLimiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !rl.b.allow() {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
