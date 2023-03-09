package note_v1

import (
	"context"
	"fmt"
	desc "github.com/Journeyman150/note-service-api/pkg/note_v1"
)

func (n *Note) Update(ctx context.Context, req *desc.UpdateNoteRequest) (*desc.UpdateNoteResponse, error) {
	fmt.Println("Update note")
	fmt.Println("Note Id:", req.GetId())
	fmt.Println("New values:")
	if len(req.GetTitle()) != 0 {
	fmt.Println("title:", req.GetTitle())
	}
	if len(req.GetText()) != 0 {
		fmt.Println("text:", req.GetText())
	}
	if len(req.GetAuthor()) != 0 {
		fmt.Println("author:", req.GetAuthor())
	}
	fmt.Println()
	return &desc.UpdateNoteResponse {
		Id: req.GetId(),
	}, nil
}