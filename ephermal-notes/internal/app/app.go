package app

import (
	"net/http"

	"github.com/master-bogdan/ephermal-notes/internal/notes"
	_ "github.com/master-bogdan/ephermal-notes/pkg/docs"
	"github.com/redis/go-redis/v9"
	httpSwagger "github.com/swaggo/http-swagger"
)

type App struct {
	Router *http.ServeMux
	Client *redis.Client
}

func Init(app App) {
	app.Router.Handle("/swagger/", httpSwagger.WrapHandler)

	notes.RouterNew(app.Router, app.Client)
}
