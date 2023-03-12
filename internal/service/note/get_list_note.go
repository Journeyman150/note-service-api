package note

import (
	"context"
	"fmt"

	desc "github.com/Journeyman150/note-service-api/pkg/note_v1"
)

func (s *Service) GetListNote(ctx context.Context, req *desc.GetListNoteRequest) (*desc.GetListNoteResponse, error) {
	res, err := s.noteRepository.GetListNote(ctx, req)
	if err != nil {
		return nil, err
	}

	fmt.Println("Get list note")
	fmt.Println()

	return res, nil
}
