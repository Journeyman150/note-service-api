package note_v1

import (
	"context"
	"fmt"
	"testing"

	"github.com/Journeyman150/note-service-api/internal/model"
	"github.com/Journeyman150/note-service-api/internal/pkg/db"
	dbMocks "github.com/Journeyman150/note-service-api/internal/pkg/db/mocksDB"
	txMocks "github.com/Journeyman150/note-service-api/internal/pkg/db/mocksTx"
	"github.com/Journeyman150/note-service-api/internal/pkg/db/transaction"
	logMocks "github.com/Journeyman150/note-service-api/internal/repository/log/mocks"
	noteMocks "github.com/Journeyman150/note-service-api/internal/repository/note/mocks"
	"github.com/Journeyman150/note-service-api/internal/service/note"
	desc "github.com/Journeyman150/note-service-api/pkg/note_v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	pgxV4 "github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
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

		repoErr   = gofakeit.Error()
		txErr     = errors.Wrap(repoErr, "failed executing code inside transaction")
		txErrText = txErr.Error()

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

		validLog = &model.Log{
			NoteId: validResId,
			Msg:    fmt.Sprintf("note with id %d was created", validResId),
		}
	)

	noteRepoMock := noteMocks.NewMockRepository(mockCtrl)
	logRepoMock := logMocks.NewMockRepository(mockCtrl)

	dbMock := dbMocks.NewMockDB(mockCtrl)
	txMock := txMocks.NewMockTx(mockCtrl)
	trManagerMock := transaction.NewMockTransactionManager(dbMock)
	txCtx := db.GetContextTx(ctx, txMock)

	api := NewMockNoteV1(Note{
		noteService: note.NewMockNoteService(noteRepoMock, logRepoMock, trManagerMock),
	})

	t.Run("success CreateNote case", func(t *testing.T) {
		dbMock.EXPECT().BeginTx(ctx, pgxV4.TxOptions{IsoLevel: pgxV4.ReadCommitted}).Return(txMock, nil)
		noteRepoMock.EXPECT().Create(txCtx, validNoteInfoModel).Return(validResId, nil)
		logRepoMock.EXPECT().Create(txCtx, validLog).Return(nil)
		txMock.EXPECT().Commit(txCtx).Return(nil)

		res, err := api.CreateNote(ctx, validReq)

		require.Nil(t, err)
		require.Equal(t, validRes, res)
	})

	t.Run("note repository returning error to Create", func(t *testing.T) {
		dbMock.EXPECT().BeginTx(ctx, pgxV4.TxOptions{IsoLevel: pgxV4.ReadCommitted}).Return(txMock, nil)
		noteRepoMock.EXPECT().Create(txCtx, validNoteInfoModel).Return(int64(0), repoErr)
		txMock.EXPECT().Rollback(txCtx).Return(nil)

		_, err := api.CreateNote(ctx, validReq)

		require.NotNil(t, err)
		require.Equal(t, txErrText, err.Error())
	})

	t.Run("log repository returning error to Create", func(t *testing.T) {
		dbMock.EXPECT().BeginTx(ctx, pgxV4.TxOptions{IsoLevel: pgxV4.ReadCommitted}).Return(txMock, nil)
		noteRepoMock.EXPECT().Create(txCtx, validNoteInfoModel).Return(validResId, nil)
		logRepoMock.EXPECT().Create(txCtx, validLog).Return(repoErr)
		txMock.EXPECT().Rollback(txCtx).Return(nil)

		_, err := api.CreateNote(ctx, validReq)

		require.NotNil(t, err)
		require.Equal(t, txErrText, err.Error())
	})
}
