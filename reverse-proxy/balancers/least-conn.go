package balancers

import (
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

type LeastConnBalancer struct {
	backends []*Backend
	mu       sync.Mutex
}

func NewLeastConnBalancer(backends []*Backend) *LeastConnBalancer {
	return &LeastConnBalancer{backends: backends}
}

func (b *LeastConnBalancer) Next(r *http.Request) (*Backend, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	var chosen *Backend
	minConn := int64(1<<63 - 1)

	for _, backend := range b.backends {
		if !backend.IsHealthy {
			continue
		}

		if backend.ConnCount < minConn {
			chosen = backend
			minConn = backend.ConnCount
		}
	}

	if chosen == nil {
		return nil, fmt.Errorf("no healthy backends")
	}

	return chosen, nil
}

func (b *LeastConnBalancer) OnStart(backend *Backend) {
	atomic.AddInt64(&backend.ConnCount, 1)
}

func (b *LeastConnBalancer) OnFinish(backend *Backend, isSuccess bool, duration time.Duration) {
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
