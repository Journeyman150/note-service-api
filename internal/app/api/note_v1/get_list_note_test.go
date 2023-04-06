package note_v1

import (
	"context"
	"testing"

	"github.com/Journeyman150/note-service-api/internal/model"
	noteMocks "github.com/Journeyman150/note-service-api/internal/repository/note/mocks"
	"github.com/Journeyman150/note-service-api/internal/service/note"
	desc "github.com/Journeyman150/note-service-api/pkg/note_v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Test_GetListNote(t *testing.T) {
	var (
		ctx      = context.Background()
		mockCtrl = gomock.NewController(t)

		repoErr     = gofakeit.Error()
		repoErrText = repoErr.Error()

		validRepoNote *model.Note
		validDescNote *desc.Note
	)
	validRepoNotes := make([]*model.Note, 0, 10)
	validDescNotes := make([]*desc.Note, 0, 10)

	gofakeit.Slice(&validRepoNotes)
	for i := range validRepoNotes {
		validRepoNote = validRepoNotes[i]

		validRepoNote.UpdatedAt.Valid = true

		validDescNote = &desc.Note{
			Id: validRepoNote.Id,
			NoteInfo: &desc.NoteInfo{
				Title:  validRepoNote.Info.Title,
				Text:   validRepoNote.Info.Text,
				Author: validRepoNote.Info.Author,
				Email:  validRepoNote.Info.Email,
			},
			CreatedAt: timestamppb.New(validRepoNote.CreatedAt),
			UpdatedAt: timestamppb.New(validRepoNote.UpdatedAt.Time),
		}
		validDescNotes = append(validDescNotes, validDescNote)
	}
	validRes := &desc.GetListNoteResponse{Notes: validDescNotes}

	noteRepoMock := noteMocks.NewMockRepository(mockCtrl)
	api := NewMockNoteV1(Note{
		noteService: note.NewMockNoteService(noteRepoMock),
	})

	t.Run("success GetListNote case", func(t *testing.T) {
		noteRepoMock.EXPECT().GetList(ctx).Return(validRepoNotes, nil)

		res, err := api.GetListNote(ctx, &emptypb.Empty{})

		require.Equal(t, validRes, res)
		require.Nil(t, err)
	})

	t.Run("note repo returning error to GetListNote", func(t *testing.T) {
		noteRepoMock.EXPECT().GetList(ctx).Return(nil, repoErr)

		res, err := api.GetListNote(ctx, &emptypb.Empty{})

		require.Nil(t, res)
		require.NotNil(t, err)
		require.Equal(t, repoErrText, err.Error())
	})

}
