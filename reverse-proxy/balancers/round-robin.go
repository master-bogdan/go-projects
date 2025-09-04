package balancers

import (
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

type RoundRobinBalancer struct {
	backends []*Backend
	mu       sync.Mutex
	idx      int
}

func NewRoundRobinBalancer(backends []*Backend) *RoundRobinBalancer {
	return &RoundRobinBalancer{backends: backends}
}

func (b *RoundRobinBalancer) Next(r *http.Request) (*Backend, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	n := len(b.backends)
	if n == 0 {
		return nil, fmt.Errorf("no backends available")
	}

	// Try up to n times to find a healthy backend
	for i := 0; i < n; i++ {
		backend := b.backends[b.idx]
		b.idx = (b.idx + 1) % n
		if backend.IsHealthy {
			return backend, nil
		}
	}

	return nil, fmt.Errorf("no healthy backends")
}

func (b *RoundRobinBalancer) OnStart(backend *Backend) {
	atomic.AddInt64(&backend.ConnCount, 1)
}

func (b *RoundRobinBalancer) OnFinish(backend *Backend, isSuccess bool, duration time.Duration) {
	atomic.AddInt64(&backend.ConnCount, -1)
	b.mu.Lock()
	defer b.mu.Unlock()
	if !isSuccess {
		backend.Failures++
		backend.LastFailure = time.Now()
	} else {
		backend.Failures = 0
	}
}
