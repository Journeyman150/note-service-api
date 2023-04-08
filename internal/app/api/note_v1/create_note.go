package note_v1

import (
	"context"

	"github.com/Journeyman150/note-service-api/internal/converter"
	desc "github.com/Journeyman150/note-service-api/pkg/note_v1"
)

func (n *Note) CreateNote(ctx context.Context, req *desc.CreateNoteRequest) (*desc.CreateNoteResponse, error) {
	id, err := n.noteService.Create(ctx, converter.ToNoteInfo(req.GetNoteInfo()))
	if err != nil {
		return nil, err
	}

	return &desc.CreateNoteResponse{Id: id}, nil
}
