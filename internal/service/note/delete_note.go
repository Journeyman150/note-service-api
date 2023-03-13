package note

import (
	"context"

	desc "github.com/Journeyman150/note-service-api/pkg/note_v1"
)

func (s *Service) DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) error {
	_, err := s.noteRepository.DeleteNote(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
