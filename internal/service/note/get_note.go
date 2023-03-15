package note

import (
	"context"

	"github.com/Journeyman150/note-service-api/internal/model"
)

func (s *Service) GetNote(ctx context.Context, id int64) (*model.GetNoteResponse, error) {
	res, err := s.noteRepository.GetNote(ctx, id)
	if err != nil {
		return nil, err
	}

	return res, nil
}
