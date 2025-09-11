package notes

import (
	"net/http"
)

type NotesController interface {
	GetNote(w http.ResponseWriter, r *http.Request)
	CreateNote(w http.ResponseWriter, r *http.Request)
	DeleteNote(w http.ResponseWriter, r *http.Request)
}

type notesController struct {
	service
	repository
}

func NewNotesController() NotesController {
	return &notesController{}
}

func (c *notesController) GetNote(w http.ResponseWriter, r *http.Request) {
	noteID := r.PathValue("id")
	if noteID == "" {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

}

func (c *notesController) CreateNote(w http.ResponseWriter, r *http.Request) {

}

func (c *notesController) DeleteNote(w http.ResponseWriter, r *http.Request) {

}
