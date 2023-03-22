package note_v1

import (
	"context"
	"testing"

	noteMocks "github.com/Journeyman150/note-service-api/internal/repository/note/mocks"
	"github.com/Journeyman150/note-service-api/internal/service/note"
	desc "github.com/Journeyman150/note-service-api/pkg/note_v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
)

func Test_DeleteNote(t *testing.T) {
	var (
		ctx      = context.Background()
		mockCtrl = gomock.NewController(t)

		id       = gofakeit.Int64()
		validReq = &desc.DeleteNoteRequest{Id: id}

		repoErr     = gofakeit.Error()
		repoErrText = repoErr.Error()
	)

	noteRepoMock := noteMocks.NewMockRepository(mockCtrl)
	api := NewMockNoteV1(Note{
		noteService: note.NewMockNoteService(noteRepoMock),
	})

	t.Run("success DeleteNote case", func(t *testing.T) {
		noteRepoMock.EXPECT().DeleteNote(ctx, validReq.GetId()).Return(nil)

		empty, err := api.DeleteNote(ctx, validReq)

		require.Equal(t, &emptypb.Empty{}, empty)
		require.Nil(t, err)
	})

	t.Run("note repo returning error to DeleteNote", func(t *testing.T) {
		noteRepoMock.EXPECT().DeleteNote(ctx, validReq.GetId()).Return(repoErr)

		empty, err := api.DeleteNote(ctx, validReq)

		require.Nil(t, empty)
		require.NotNil(t, err)
		require.Equal(t, repoErrText, err.Error())
	})
}
