package note

import (
	"context"

	"github.com/Journeyman150/note-service-api/internal/model"
)

func (s *Service) GetListNote(ctx context.Context) ([]*model.Note, error) {
	notes, err := s.noteRepository.GetListNote(ctx)
	if err != nil {
		return nil, err
	}

	return notes, nil
}
