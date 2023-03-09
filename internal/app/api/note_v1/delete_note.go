package note_v1

import (
	"context"
	"fmt"

	desc "github.com/Journeyman150/note-service-api/pkg/note_v1"
)

func (n *Note) DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) (*desc.Empty, error) {
	fmt.Println("Delete note")
	fmt.Println("Note Id:", req.GetId())
	fmt.Println()
	return &desc.Empty{}, nil
}
