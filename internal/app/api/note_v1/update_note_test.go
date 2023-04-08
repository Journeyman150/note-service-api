package note_v1

import (
	"context"
	"database/sql"
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
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func Test_UpdateNote(t *testing.T) {
	var (
		ctx      = context.Background()
		mockCtrl = gomock.NewController(t)

		noteId = gofakeit.Int64()
		title  = gofakeit.BeerName()
		text   = gofakeit.Sentence(4)
		author = gofakeit.Name()
		email  = gofakeit.Email()

		repoErr   = gofakeit.Error()
		txErr     = errors.Wrap(repoErr, "failed executing code inside transaction")
		txErrText = txErr.Error()

		validReq = &desc.UpdateNoteRequest{
			Id: noteId,
			UpdateNoteInfo: &desc.UpdateNoteInfo{
				Title:  &wrapperspb.StringValue{Value: title},
				Text:   &wrapperspb.StringValue{Value: text},
				Author: &wrapperspb.StringValue{Value: author},
				Email:  &wrapperspb.StringValue{Value: email},
			},
		}
		validRes = &emptypb.Empty{}

		validUpdateNoteInfoModel = &model.UpdateNoteInfo{
			Title:  sql.NullString{String: title, Valid: true},
			Text:   sql.NullString{String: text, Valid: true},
			Author: sql.NullString{String: author, Valid: true},
			Email:  sql.NullString{String: email, Valid: true},
		}

		validLog = &model.Log{
			NoteId: noteId,
			Msg:    fmt.Sprintf("note with id %d was updated", noteId),
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

	t.Run("success UpdateNote case", func(t *testing.T) {
		dbMock.EXPECT().BeginTx(ctx, pgxV4.TxOptions{IsoLevel: pgxV4.ReadCommitted}).Return(txMock, nil)
		noteRepoMock.EXPECT().Update(txCtx, noteId, validUpdateNoteInfoModel).Return(nil)
		logRepoMock.EXPECT().Create(txCtx, validLog).Return(nil)
		txMock.EXPECT().Commit(txCtx).Return(nil)

		res, err := api.Update(ctx, validReq)

		require.Nil(t, err)
		require.Equal(t, validRes, res)
	})

	t.Run("note repository returning error to Update", func(t *testing.T) {
		dbMock.EXPECT().BeginTx(ctx, pgxV4.TxOptions{IsoLevel: pgxV4.ReadCommitted}).Return(txMock, nil)
		noteRepoMock.EXPECT().Update(txCtx, noteId, validUpdateNoteInfoModel).Return(repoErr)
		txMock.EXPECT().Rollback(txCtx).Return(nil)

		_, err := api.Update(ctx, validReq)

		require.NotNil(t, err)
		require.Equal(t, txErrText, err.Error())
	})

	t.Run("log repository returning error to Create", func(t *testing.T) {
		dbMock.EXPECT().BeginTx(ctx, pgxV4.TxOptions{IsoLevel: pgxV4.ReadCommitted}).Return(txMock, nil)
		noteRepoMock.EXPECT().Update(txCtx, noteId, validUpdateNoteInfoModel).Return(nil)
		logRepoMock.EXPECT().Create(txCtx, validLog).Return(repoErr)
		txMock.EXPECT().Rollback(txCtx).Return(nil)

		_, err := api.Update(ctx, validReq)

		require.NotNil(t, err)
		require.Equal(t, txErrText, err.Error())
	})
}
