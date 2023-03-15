package note_v1

import (
	"context"

	"github.com/Journeyman150/note-service-api/internal/converter"
	desc "github.com/Journeyman150/note-service-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (n *Note) Update(ctx context.Context, req *desc.UpdateNoteRequest) (*emptypb.Empty, error) {
	err := n.noteService.UpdateNote(ctx, converter.ToUpdateNoteRequest(req))
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
