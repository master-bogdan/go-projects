package balancers

import "sync"

type LeastConnBalancer struct {
	backends []*Backend
	mu       sync.Mutex
}

func NewLeastConnBalancer(backends []*Backend) *LeastConnBalancer {
	return &LeastConnBalancer{backends: backends}
}
