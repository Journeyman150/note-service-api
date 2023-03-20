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

	//fmt.Println("Test_GetNote variables:\n",
	//	"id:", validId,
	//	"title:", validTitle,
	//	"text:", validText,
	//	"author:", validAuthor,
	//	"email:", validEmail,
	//	"createdAt:", timestamppb.New(validCreatedAt),
	//	"updatedAt:", timestamppb.New(validUpdatedAt.Time),
	//	"\nRequest id:", validReq.GetId(),
	//	"\ntest error text:", repoErrText,
	//)

	// fmt.Println("Test_GetNote expect returning variables:\n",
	// 	"id:", validRes.GetNote().GetId(),
	// 	"title:", validRes.GetNote().GetNoteInfo().GetTitle(),
	// 	"text:", validRes.GetNote().GetNoteInfo().GetText(),
	// 	"author:", validRes.GetNote().GetNoteInfo().GetAuthor(),
	// 	"email:", validRes.GetNote().GetNoteInfo().GetEmail(),
	// 	"createdAt:", validRes.GetNote().GetCreatedAt(),
	// 	"updatedAt:", validRes.GetNote().GetUpdatedAt(),
	// )

	t.Run("success GetNote case", func(t *testing.T) {
		noteRepoMock.EXPECT().GetNote(ctx, validReq.GetId()).Return(validRepoNote, nil)

		res, err := api.GetNote(ctx, validReq)

		//fmt.Println("Test_GetNote returning variables:\n",
		//	"id:", res.GetNote().GetId(),
		//	"title:", res.GetNote().GetNoteInfo().GetTitle(),
		//	"text:", res.GetNote().GetNoteInfo().GetText(),
		//	"author:", res.GetNote().GetNoteInfo().GetAuthor(),
		//	"email:", res.GetNote().GetNoteInfo().GetEmail(),
		//	"createdAt:", res.GetNote().GetCreatedAt(),
		//	"updatedAt:", res.GetNote().GetUpdatedAt(),
		//)

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
