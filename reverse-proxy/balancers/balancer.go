package balancers

import (
	"net/http"
	"net/url"
	"time"
)

type Backend struct {
	URL         *url.URL
	ConnCount   int64
	IsHealthy   bool
	Failures    int
	LastFailure time.Time
}

type Balancer interface {
	Next(r *http.Request) (*Backend, error)
	OnStart(b *Backend)
	OnFinish(b *Backend, isSuccess bool, duration time.Duration)
}
