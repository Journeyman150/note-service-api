package note_v1

import (
	"context"

	desc "github.com/Journeyman150/note-service-api/pkg/note_v1"
)

func (n *Note) DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) (*desc.Empty, error) {
	res, err := n.noteService.DeleteNote(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
