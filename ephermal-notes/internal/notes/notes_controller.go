package notes

import (
	"encoding/json"
	"net/http"

	memory_db "github.com/master-bogdan/ephermal-notes/internal/infra/db/redis"
)

type NotesController interface {
	GetNote(w http.ResponseWriter, r *http.Request)
	CreateNote(w http.ResponseWriter, r *http.Request)
	DeleteNote(w http.ResponseWriter, r *http.Request)
}

type notesController struct {
	service NotesService
}

func NewNotesController(service NotesService) NotesController {
	return &notesController{
		service: service,
	}
}

func (c *notesController) GetNote(w http.ResponseWriter, r *http.Request) {
	noteID := r.PathValue("id")
	if noteID == "" {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	note, err := c.service.GetNote(noteID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(note)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (c *notesController) CreateNote(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var input struct {
		Message string `json:"message"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// change to dto later
	note := &memory_db.NotesModel{
		Message: input.Message,
	}

	createdNote, err := c.service.CreateNote(note)
	if err != nil {
		http.Error(w, "Failed to create note", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(createdNote)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (c *notesController) DeleteNote(w http.ResponseWriter, r *http.Request) {
	noteID := r.PathValue("id")
	if noteID == "" {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	err := c.service.DeleteNote(noteID)
	if err != nil {
		http.Error(w, "Failed to delete note", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
