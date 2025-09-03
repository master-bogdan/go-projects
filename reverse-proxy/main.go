package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"

	"gopkg.in/yaml.v3"
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

type Config struct {
	Listen   string
	Backends []string
	Balancer string
	Health   struct {
		Path                   string
		Interval               string
		Timeout                int
		PassiveFailuresForOpen int
		Coodown                int
	}
	Retry struct {
		Max     int
		Backoff int
	}
	Timeout struct {
		Read  int
		Write int
		Idle  int
	}
	Transport struct {
		DialTimeout         int
		TLSHandshakeTimeout int
		MaxIdlePerHost      int
	}
}

type ProxyServer struct {
	backends []*Backend
	balancer Balancer
	config   Config
}

func NewProxyServer(cfg Config) *ProxyServer {
	var backends []*Backend

	for _, rawURL := range cfg.Backends {
		backendUrl, err := url.Parse(rawURL)
		if err != nil {
			panic(err)
		}

		backends = append(backends, &Backend{URL: backendUrl, IsHealthy: true})
	}

	var balancer Balancer
	switch cfg.Balancer {
	case "least_conn":
		balancer = NewLeastConnBalancer(backends)
	default:
		balancer = NewRoundRobinBalancer(backends)
	}

	return &ProxyServer{backends: backends, balancer: balancer, config: cfg}
}

func (p *ProxyServer) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	backend, err := p.balancer.Next(r)
	if err != nil {
		http.Error(w, "No backend available", http.StatusServiceUnavailable)
		return
	}

	p.balancer.OnStart(backend)
	start := time.Now()

	proxy := httputil.NewSingleHostReverseProxy(backend.URL)

	proxy.ErrorHandler = func(rw http.ResponseWriter, req *http.Request, err error) {
		p.balancer.OnFinish(backend, false, time.Since(start))
		http.Error(rw, "Backend error: "+err.Error(), http.StatusBadGateway)
	}

	proxy.ModifyResponse = func(r *http.Response) error {
		p.balancer.OnFinish(backend, true, time.Since(start))
		return nil
	}

	proxy.ServeHTTP(w, r)
}

func main() {
	file, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	var config Config
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}

}
