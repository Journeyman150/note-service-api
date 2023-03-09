package note_v1

import (
	"context"
	"fmt"

	desc "github.com/Journeyman150/note-service-api/pkg/note_v1"
)

func (n *Note) GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponse, error) {
	fmt.Println("Get note")
	fmt.Println("Note Id:", req.GetId())
	fmt.Println()
	return &desc.GetNoteResponse{
		Title:  "Title received",
		Text:   "Text received",
		Author: "Author received",
	}, nil
}
