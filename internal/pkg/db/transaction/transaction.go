package transaction

import (
	"context"

	"github.com/Journeyman150/note-service-api/internal/pkg/db"
	pgxV4 "github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

type Manager interface {
	ReadCommitted(ctx context.Context, f Handler) error
	RepeatableRead(ctx context.Context, f Handler) error
	Serializable(ctx context.Context, f Handler) error
}

type Handler func(ctx context.Context) error

type manager struct {
	db db.Transactor
}

func NewTransactionManager(db db.Transactor) *manager {
	return &manager{
		db: db,
	}
}

func (s *manager) transaction(ctx context.Context, txOptions pgxV4.TxOptions, fn Handler) (err error) {
	tx, ok := ctx.Value(db.TxKey).(pgxV4.Tx)
	if ok {
		return fn(ctx)
	}

	tx, err = s.db.BeginTx(ctx, txOptions)
	if err != nil {
		return errors.Wrap(err, "can't begin transaction")
	}

	ctx = db.GetContextTx(ctx, tx)

	defer func() {
		if r := recover(); r != nil {
			err = errors.Errorf("panic recovered: %v", r)
		}

		if err == nil {
			err = tx.Commit(ctx)
			if err != nil {
				err = errors.Wrap(err, "tx commit failed")
			}
		}

		if err != nil {
			if errRollback := tx.Rollback(ctx); errRollback != nil {
				err = errors.Wrapf(err, "errRollback: %v", errRollback)
			}
		}
	}()

	if err = fn(ctx); err != nil {
		err = errors.Wrap(err, "failed executing code inside transaction")
	}

	return err
}

func (s *manager) ReadCommitted(ctx context.Context, f Handler) error {
	txOptions := pgxV4.TxOptions{IsoLevel: pgxV4.ReadCommitted}
	return s.transaction(ctx, txOptions, f)
}

func (s *manager) RepeatableRead(ctx context.Context, f Handler) error {
	txOptions := pgxV4.TxOptions{IsoLevel: pgxV4.RepeatableRead}
	return s.transaction(ctx, txOptions, f)
}

func (s *manager) Serializable(ctx context.Context, f Handler) error {
	txOptions := pgxV4.TxOptions{IsoLevel: pgxV4.Serializable}
	return s.transaction(ctx, txOptions, f)
}
