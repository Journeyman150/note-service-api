package note

import (
	"context"

	"github.com/Journeyman150/note-service-api/internal/model"
)

func (s *Service) UpdateNote(ctx context.Context, id int64, noteInfo *model.UpdateNoteInfo) error {
	_, err := s.noteRepository.UpdateNote(ctx, id, noteInfo)
	if err != nil {
		return err
	}

	return nil
}
