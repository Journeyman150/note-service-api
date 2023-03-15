package note_v1

import (
	"context"

	"github.com/Journeyman150/note-service-api/internal/converter"
	desc "github.com/Journeyman150/note-service-api/pkg/note_v1"
)

func (n *Note) GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponse, error) {
	res, err := n.noteService.GetNote(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return converter.ToDescGetNoteResponse(res), nil
}
