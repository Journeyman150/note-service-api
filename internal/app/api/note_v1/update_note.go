package note_v1 //nolint:revive

import (
	"context"

	"github.com/Journeyman150/note-service-api/internal/converter"
	desc "github.com/Journeyman150/note-service-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (n *Note) Update(ctx context.Context, req *desc.UpdateNoteRequest) (*emptypb.Empty, error) {
	err := n.noteService.Update(ctx, req.GetId(), converter.ToUpdateNoteInfo(req.UpdateNoteInfo))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
