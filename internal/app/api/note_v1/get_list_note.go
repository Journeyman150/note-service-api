package note_v1

import (
	"context"
	"fmt"

	desc "github.com/Journeyman150/note-service-api/pkg/note_v1"
)

func (n *Note) GetListNote(ctx context.Context, req *desc.GetListNoteRequest) (*desc.GetListNoteResponse, error) {
	res, err := n.noteService.GetListNote(ctx, req)
	if err != nil {
		return nil, err
	}

	fmt.Println("Get list note")
	fmt.Println()

	return res, nil
}
