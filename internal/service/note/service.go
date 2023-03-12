package note

import "github.com/Journeyman150/note-service-api/internal/repository"

type Service struct {
	noteRepository repository.NoteRepository
}

func NewService(noteRepository repository.NoteRepository) *Service {
	return &Service{
		noteRepository: noteRepository,
	}
}
