package note

import (
	"context"

	"github.com/Journeyman150/note-service-api/internal/model"
)

func (s *Service) CreateNote(ctx context.Context, noteInfo *model.NoteInfo) (int64, error) {
	id, err := s.noteRepository.CreateNote(ctx, noteInfo)
	if err != nil {
		return 0, err
	}

	return id, nil
}
