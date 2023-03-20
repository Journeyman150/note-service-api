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
)

func Test_CreateNote(t *testing.T) {
	var (
		ctx      = context.Background()
		mockCtrl = gomock.NewController(t)

		validTitle  = gofakeit.BeerName()
		validText   = gofakeit.LetterN(10)
		validAuthor = gofakeit.Name()
		validEmail  = gofakeit.Email()

		validResId = gofakeit.Int64()

		repoErr     = gofakeit.Error()
		repoErrText = repoErr.Error()

		validReq = &desc.CreateNoteRequest{NoteInfo: &desc.NoteInfo{
			Title:  validTitle,
			Text:   validText,
			Author: validAuthor,
			Email:  validEmail,
		}}

		validNoteInfoModel = &model.NoteInfo{
			Title:  validTitle,
			Text:   validText,
			Author: validAuthor,
			Email:  validEmail,
		}

		validRes = &desc.CreateNoteResponse{Id: validResId}
	)

	noteRepoMock := noteMocks.NewMockRepository(mockCtrl)
	api := NewMockNoteV1(Note{
		noteService: note.NewMockNoteService(noteRepoMock),
	})

	//fmt.Println("Test_CreateNote variables:\n",
	//	"title:", validReq.NoteInfo.GetTitle(),
	//	"text:", validReq.NoteInfo.GetText(),
	//	"author:", validReq.NoteInfo.GetAuthor(),
	//	"email:", validReq.NoteInfo.GetEmail(),
	//	"\nid:", validRes.GetId(),
	//	"\ntest error text:", repoErrText)

	t.Run("success CreateNote case", func(t *testing.T) {
		noteRepoMock.EXPECT().CreateNote(ctx, validNoteInfoModel).Return(validResId, nil)

		res, err := api.CreateNote(ctx, validReq)

		require.Nil(t, err)
		require.Equal(t, validRes, res)
	})

	t.Run("note repository returning error to CreateNote", func(t *testing.T) {
		noteRepoMock.EXPECT().CreateNote(ctx, validNoteInfoModel).Return(int64(0), repoErr)

		_, err := api.CreateNote(ctx, validReq)

		require.NotNil(t, err)
		require.Equal(t, repoErrText, err.Error())
	})
}
