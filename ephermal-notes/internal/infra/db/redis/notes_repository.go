package redis

import "github.com/redis/go-redis/v9"

type NotesModel struct {
	ID      string `redis:"id"`
	Key     string `redis:"key"`
	Message string `redis:"message"`
}

type NotesRepository interface {
	Get()
	Create()
	Delete()
}

type notesRepository struct {
	client *redis.Client
}

func NewNotesRepository(client *redis.Client) NotesRepository {
	return &notesRepository{
		client: client,
	}
}

func (r *notesRepository) Get() {}

func (r *notesRepository) Create() {}

func (r *notesRepository) Delete() {}
