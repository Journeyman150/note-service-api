package note

import (
	"github.com/Journeyman150/note-service-api/internal/repository/note"
)

type Service struct {
	noteRepository note.Repository
}

func NewService(noteRepository note.Repository) *Service {
	return &Service{
		noteRepository: noteRepository,
	}
}

func NewMockNoteService(deps ...interface{}) *Service {
	service := Service{}

	for _, v := range deps {
		switch s := v.(type) {
		case note.Repository:
			service.noteRepository = s
		}
	}

	return &service
}
