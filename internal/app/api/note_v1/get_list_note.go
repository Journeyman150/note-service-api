package note_v1

import (
	"context"

	"github.com/Journeyman150/note-service-api/internal/converter"
	desc "github.com/Journeyman150/note-service-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (n *Note) GetListNote(ctx context.Context, _ *emptypb.Empty) (*desc.GetListNoteResponse, error) {
	res, err := n.noteService.GetListNote(ctx)
	if err != nil {
		return nil, err
	}

	return converter.ToDescGetListNoteResponse(res), nil
}
