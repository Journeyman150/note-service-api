package note

import (
	"context"
	"fmt"

	desc "github.com/Journeyman150/note-service-api/pkg/note_v1"
)

func (s *Service) DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) (*desc.Empty, error) {
	result, err := s.noteRepository.DeleteNote(ctx, req)
	if err != nil {
		return nil, err
	}
	fmt.Println("Delete note")
	fmt.Println("Note Id:", req.GetId())
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	fmt.Println("Rows affected: ", rowsAffected)
	fmt.Println()
	return &desc.Empty{}, nil
}
