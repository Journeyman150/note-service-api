package note_v1

import (
	"context"
	"fmt"
	desc "github.com/Journeyman150/note-service-api/pkg/note_v1"
)

func (n *Note) GetListNote(ctx context.Context, req *desc.GetListNoteRequest) (*desc.GetListNoteResponse, error) {
	fmt.Println("Get list note")
	fmt.Println()
	return &desc.GetListNoteResponse {
		Notes: []*desc.GetNoteResponse {
			{
				Title: "Title 1",
				Text: "Text 1",
				Author: "Author 1",
			},
			{
				Title: "Title 2",
				Text: "Text 2",
				Author: "Author 2",
			},
			{
				Title:  "Title 3",
				Text:   "Text 3",
				Author: "Author 3",
			},
		},
	}, nil
}
