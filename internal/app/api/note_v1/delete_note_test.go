package note_v1 //nolint:revive

import (
	"context"
	"testing"

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
	"google.golang.org/protobuf/types/known/emptypb"
)

func Test_DeleteNote(t *testing.T) {
	var (
		ctx      = context.Background()
		mockCtrl = gomock.NewController(t)

		noteId   = gofakeit.Int64()
		validReq = &desc.DeleteNoteRequest{Id: noteId}
		validRes = &emptypb.Empty{}

		repoErr   = gofakeit.Error()
		txErr     = errors.Wrap(repoErr, "failed executing code inside transaction")
		txErrText = txErr.Error()
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

	t.Run("success DeleteNote case", func(t *testing.T) {
		dbMock.EXPECT().BeginTx(ctx, pgxV4.TxOptions{IsoLevel: pgxV4.ReadCommitted}).Return(txMock, nil)
		noteRepoMock.EXPECT().Delete(txCtx, noteId).Return(nil)
		logRepoMock.EXPECT().Delete(txCtx, noteId).Return(nil)
		txMock.EXPECT().Commit(txCtx).Return(nil)

		res, err := api.DeleteNote(ctx, validReq)

		require.Nil(t, err)
		require.Equal(t, validRes, res)
	})

	t.Run("note repository returning error to Delete", func(t *testing.T) {
		dbMock.EXPECT().BeginTx(ctx, pgxV4.TxOptions{IsoLevel: pgxV4.ReadCommitted}).Return(txMock, nil)
		logRepoMock.EXPECT().Delete(txCtx, noteId).Return(nil)
		noteRepoMock.EXPECT().Delete(txCtx, noteId).Return(repoErr)
		txMock.EXPECT().Rollback(txCtx).Return(nil)

		_, err := api.DeleteNote(ctx, validReq)

		require.NotNil(t, err)
		require.Equal(t, txErrText, err.Error())
	})

	t.Run("log repository returning error to Delete", func(t *testing.T) {
		dbMock.EXPECT().BeginTx(ctx, pgxV4.TxOptions{IsoLevel: pgxV4.ReadCommitted}).Return(txMock, nil)
		logRepoMock.EXPECT().Delete(txCtx, noteId).Return(repoErr)
		txMock.EXPECT().Rollback(txCtx).Return(nil)

		_, err := api.DeleteNote(ctx, validReq)

		require.NotNil(t, err)
		require.Equal(t, txErrText, err.Error())
	})
}
