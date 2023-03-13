package note_v1

import (
	"context"

	desc "github.com/Journeyman150/note-service-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (n *Note) Update(ctx context.Context, req *desc.UpdateNoteRequest) (*emptypb.Empty, error) {
	err := n.noteService.UpdateNote(ctx, req)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}
