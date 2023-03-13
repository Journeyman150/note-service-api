package note

import (
	"github.com/Journeyman150/note-service-api/internal/repository/note"
)

type Service struct {
	noteRepository note.NoteRepository
}

func NewService(noteRepository note.NoteRepository) *Service {
	return &Service{
		noteRepository: noteRepository,
	}
}
