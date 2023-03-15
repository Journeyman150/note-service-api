package note

import (
	"context"

	"github.com/Journeyman150/note-service-api/internal/model"
)

func (s *Service) UpdateNote(ctx context.Context, req *model.UpdateNoteRequest) error {
	_, err := s.noteRepository.UpdateNote(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
