package note

import (
	"context"
	"fmt"

	desc "github.com/Journeyman150/note-service-api/pkg/note_v1"
)

func (s *Service) CreateNote(ctx context.Context, req *desc.CreateNoteRequest) (*desc.CreateNoteResponse, error) {
	id, err := s.noteRepository.CreateNote(ctx, req)
	if err != nil {
		return nil, err
	}

	fmt.Println("Create Note")
	fmt.Println("id:", id)
	fmt.Println("title:", req.GetTitle())
	fmt.Println("text:", req.GetText())
	fmt.Println("author:", req.GetAuthor())
	fmt.Println("email:", req.GetEmail())
	fmt.Println()

	return &desc.CreateNoteResponse{
		Id: id,
	}, nil
}
