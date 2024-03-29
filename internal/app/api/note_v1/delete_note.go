package note_v1 //nolint:revive

import (
	"context"

	desc "github.com/Journeyman150/note-service-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (n *Note) DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) (*emptypb.Empty, error) {
	err := n.noteService.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
