package main

import (
	"log"
	"net/http"
	"time"

	"github.com/master-bogdan/ephermal-notes/internal/app"
	"github.com/master-bogdan/ephermal-notes/internal/infra/db/redis"
	"github.com/master-bogdan/ephermal-notes/pkg/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	client, err := memory_db.Connect(cfg)
	if err != nil {
		log.Fatalf("failed connect to redis: %v", err)
	}

	mux := http.NewServeMux()

	App := &app.App{
		Router: mux,
		Client: client,
	}

	app.Init(*App)

	addr := cfg.Server.Host + ":" + cfg.Server.Port
	server := http.Server{
		Addr:        addr,
		ReadTimeout: 3 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
