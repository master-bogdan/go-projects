package notes

import (
	"github.com/master-bogdan/ephermal-notes/internal/infra/db/redis"
	"github.com/redis/go-redis/v9"
	"net/http"
)

func RouterNew(m *http.ServeMux, client *redis.Client) {
	repo := memory_db.NewNotesRepository(client)
	service := NewNotesService(repo)
	controller := NewNotesController(service)

	m.HandleFunc("GET /notes/{id}", controller.GetNote)
	m.HandleFunc("POST /notes", controller.CreateNote)
	m.HandleFunc("DELETE /notes/{id}", controller.DeleteNote)
}
