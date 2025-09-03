package main

import (
	"log"
	"net/http"
	"time"

	"github.com/master-bogdan/reverse-proxy/config"
	proxyserver "github.com/master-bogdan/reverse-proxy/proxy-server"
)

func main() {
	cfg := config.New()
	proxyServer := proxyserver.NewProxyServer(*cfg)

	srv := &http.Server{
		Addr:         cfg.Listen,
		Handler:      proxyServer,
		ReadTimeout:  time.Duration(cfg.Timeout.Read) * time.Second,
		WriteTimeout: time.Duration(cfg.Timeout.Write) * time.Second,
		IdleTimeout:  time.Duration(cfg.Timeout.Idle) * time.Second,
	}

	log.Printf("starting proxy on %s with %s balancer", cfg.Listen, cfg.Balancer)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
