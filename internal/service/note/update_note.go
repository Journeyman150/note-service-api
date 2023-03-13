package note

import (
	"context"

	desc "github.com/Journeyman150/note-service-api/pkg/note_v1"
)

func (s *Service) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) error {
	_, err := s.noteRepository.UpdateNote(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
