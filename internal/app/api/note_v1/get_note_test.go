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
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Test_GetNote(t *testing.T) {
	var (
		ctx      = context.Background()
		mockCtrl = gomock.NewController(t)

		validId        = gofakeit.Int64()
		validTitle     = gofakeit.BeerName()
		validText      = gofakeit.LetterN(10)
		validAuthor    = gofakeit.Name()
		validEmail     = gofakeit.Email()
		validCreatedAt = gofakeit.Date()
		validUpdatedAt = sql.NullTime{
			Time:  gofakeit.Date(),
			Valid: true,
		}
		repoErr     = gofakeit.Error()
		repoErrText = repoErr.Error()

		validRepoNoteInfo = &model.NoteInfo{
			Title:  validTitle,
			Text:   validText,
			Author: validAuthor,
			Email:  validEmail,
		}
		validRepoNote = &model.Note{
			Id:        validId,
			Info:      validRepoNoteInfo,
			CreatedAt: validCreatedAt,
			UpdatedAt: validUpdatedAt,
		}

		validDescNoteInfo = &desc.NoteInfo{
			Title:  validTitle,
			Text:   validText,
			Author: validAuthor,
			Email:  validEmail,
		}
		validDescNote = &desc.Note{
			Id:        validId,
			NoteInfo:  validDescNoteInfo,
			CreatedAt: timestamppb.New(validCreatedAt),
			UpdatedAt: timestamppb.New(validUpdatedAt.Time),
		}

		validRes = &desc.GetNoteResponse{Note: validDescNote}
		validReq = &desc.GetNoteRequest{Id: validId}
	)

	noteRepoMock := noteMocks.NewMockRepository(mockCtrl)
	api := NewMockNoteV1(Note{
		noteService: note.NewMockNoteService(noteRepoMock),
	})

	t.Run("success GetNote case", func(t *testing.T) {
		noteRepoMock.EXPECT().GetNote(ctx, validReq.GetId()).Return(validRepoNote, nil)

		res, err := api.GetNote(ctx, validReq)

		require.Equal(t, validRes, res)
		require.Nil(t, err)
	})

	t.Run("note repo returning error to GetNote", func(t *testing.T) {
		noteRepoMock.EXPECT().GetNote(ctx, validReq.GetId()).Return(nil, repoErr)

		res, err := api.GetNote(ctx, validReq)

		require.Nil(t, res)
		require.NotNil(t, err)
		require.Equal(t, repoErrText, err.Error())
	})
}
