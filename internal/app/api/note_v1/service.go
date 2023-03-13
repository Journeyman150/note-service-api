package note_v1

import (
	"github.com/Journeyman150/note-service-api/internal/service/note"
	desc "github.com/Journeyman150/note-service-api/pkg/note_v1"
)

type Note struct {
	desc.UnimplementedNoteV1Server
	noteService *note.Service
}

func NewNote(noteService *note.Service) *Note {
	return &Note{
		noteService: noteService,
	}
}
