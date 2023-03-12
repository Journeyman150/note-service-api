package note

import (
	"context"
	"fmt"

	desc "github.com/Journeyman150/note-service-api/pkg/note_v1"
)

func (s *Service) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (*desc.Empty, error) {
	result, err := s.noteRepository.UpdateNote(ctx, req)
	fmt.Println("Update note")
	fmt.Println("Note Id:", req.GetId())
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	fmt.Println("Rows affected: ", rowsAffected)
	fmt.Println()
	return &desc.Empty{}, err
}
