package note

import (
	"context"
	"fmt"

	desc "github.com/Journeyman150/note-service-api/pkg/note_v1"
)

func (s *Service) GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponse, error) {
	res, err := s.noteRepository.GetNote(ctx, req)
	if err != nil {
		return nil, err
	}

	fmt.Println("Get note")
	fmt.Println("Note Id:", res.GetId())
	fmt.Println()

	return res, nil
}
