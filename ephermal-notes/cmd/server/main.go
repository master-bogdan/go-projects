package main

import (
	"log"
	"net/http"
	"time"

	"github.com/master-bogdan/ephermal-notes/config"
	"github.com/master-bogdan/ephermal-notes/internal/infra/db/redis"
	"github.com/master-bogdan/ephermal-notes/internal/notes"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	_, err = redis.Connect(cfg)
	if err != nil {
		log.Fatalf("failed connect to redis: %v", err)
	}

	mux := http.NewServeMux()

	notes.RouterNew(mux)

	addr := cfg.Server.Host + ":" + cfg.Server.Port
	server := http.Server{
		Addr:        addr,
		ReadTimeout: 3 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
