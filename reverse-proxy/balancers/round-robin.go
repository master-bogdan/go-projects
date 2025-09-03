package balancers

import "sync"

type RoundRobinBalancer struct {
	backends []*Backend
	mu       sync.Mutex
	idx      int
}

func NewRoundRobinBalancer(backends []*Backend) *RoundRobinBalancer {
	return &RoundRobinBalancer{backends: backends}
}
