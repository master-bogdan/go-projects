package notes

import "net/http"

func RouterNew(m *http.ServeMux) {
	controller := NewNotesController()

	m.HandleFunc("GET /notes/{id}", controller.GetNote)
	m.HandleFunc("POST /notes", controller.CreateNote)
	m.HandleFunc("DELETE /notes/{id}", controller.DeleteNote)
}
