package proxyserver

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/master-bogdan/reverse-proxy/balancers"
	"github.com/master-bogdan/reverse-proxy/config"
)

type ProxyServer struct {
	backends []*balancers.Backend
	balancer balancers.Balancer
	config   config.Config
}

func NewProxyServer(cfg config.Config) *ProxyServer {
	var backends []*balancers.Backend

	for _, rawURL := range cfg.Backends {
		backendUrl, err := url.Parse(rawURL)
		if err != nil {
			panic(err)
		}

		backends = append(backends, &balancers.Backend{URL: backendUrl, IsHealthy: true})
	}

	var balancer balancers.Balancer
	switch cfg.Balancer {
	case "least_conn":
		balancer = balancers.NewLeastConnBalancer(backends)
	default:
		balancer = balancers.NewRoundRobinBalancer(backends)
	}

	return &ProxyServer{backends: backends, balancer: balancer, config: cfg}
}

func (p *ProxyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
