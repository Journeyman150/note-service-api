package note_v1

import (
	"context"
	"database/sql"
	"testing"

	"github.com/Journeyman150/note-service-api/internal/model"
	noteMocks "github.com/Journeyman150/note-service-api/internal/repository/note/mocks"
	"github.com/Journeyman150/note-service-api/internal/service/note"
	desc "github.com/Journeyman150/note-service-api/pkg/note_v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func Test_UpdateNote(t *testing.T) {
	var (
		ctx      = context.Background()
		mockCtrl = gomock.NewController(t)

		id    = gofakeit.Int64()
		title = gofakeit.BeerName()
		text  = gofakeit.Sentence(4)
		//author = gofakeit.Name()
		email = gofakeit.Email()

		repoErr     = gofakeit.Error()
		repoErrText = repoErr.Error()

		validReq = &desc.UpdateNoteRequest{
			Id: id,
			UpdateNoteInfo: &desc.UpdateNoteInfo{
				Title:  &wrapperspb.StringValue{Value: title},
				Text:   &wrapperspb.StringValue{Value: text},
				Author: nil,
				Email:  &wrapperspb.StringValue{Value: email},
			},
		}
		validUpdateNoteInfoModel = &model.UpdateNoteInfo{
			Title:  sql.NullString{String: title, Valid: true},
			Text:   sql.NullString{String: text, Valid: true},
			Author: sql.NullString{Valid: false},
			Email:  sql.NullString{String: email, Valid: true},
		}
	)

	noteRepoMock := noteMocks.NewMockRepository(mockCtrl)
	api := NewMockNoteV1(Note{
		noteService: note.NewMockNoteService(noteRepoMock),
	})

	t.Run("success UpdateNote case", func(t *testing.T) {
		noteRepoMock.EXPECT().UpdateNote(ctx, validReq.GetId(), validUpdateNoteInfoModel).Return(nil)

		empty, err := api.Update(ctx, validReq)

		require.Equal(t, &emptypb.Empty{}, empty)
		require.Nil(t, err)
	})

	t.Run("note repo returning error to UpdateNote", func(t *testing.T) {
		noteRepoMock.EXPECT().UpdateNote(ctx, validReq.GetId(), validUpdateNoteInfoModel).Return(repoErr)

		empty, err := api.Update(ctx, validReq)

		require.Nil(t, empty)
		require.NotNil(t, err)
		require.Equal(t, repoErrText, err.Error())
	})
}
