package app

import (
	"net/http"

	"github.com/master-bogdan/ephermal-notes/internal/notes"
	"github.com/redis/go-redis/v9"
)

type App struct {
	Router *http.ServeMux
	Client *redis.Client
}

func Init(app App) {
	notes.RouterNew(app.Router, app.Client)
}
