package db

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	pgxV4 "github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type key string

const (
	TxKey key = "tx"
)

type Query struct {
	Name     string
	QueryRaw string
}

type Transactor interface {
	BeginTx(ctx context.Context, txOptions pgxV4.TxOptions) (pgxV4.Tx, error)
}

type SQLExecer interface {
	NamedExecer
	QueryExecer
}

type NamedExecer interface {
	GetContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
}

type QueryExecer interface {
	ExecContext(ctx context.Context, q Query, args ...interface{}) (pgconn.CommandTag, error)
	QueryContext(ctx context.Context, q Query, args ...interface{}) (pgxV4.Rows, error)
	QueryRowContext(ctx context.Context, q Query, args ...interface{}) pgxV4.Row
}

type Pinger interface {
	Ping(ctx context.Context) error
}

type DB interface {
	SQLExecer
	Transactor
	Pinger
	Close()
}

type db struct {
	pool *pgxpool.Pool
}

func (d *db) GetContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error {
	return pgxscan.Get(ctx, d.pool, dest, q.QueryRaw, args...)
}

func (d *db) SelectContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error {
	return pgxscan.Select(ctx, d.pool, dest, q.QueryRaw, args...)
}

func (d *db) ExecContext(ctx context.Context, q Query, args ...interface{}) (pgconn.CommandTag, error) {
	tx, ok := ctx.Value(TxKey).(pgxV4.Tx)
	if ok {
		return tx.Exec(ctx, q.QueryRaw, args...)
	}

	return d.pool.Exec(ctx, q.QueryRaw, args...)
}

func (d *db) QueryContext(ctx context.Context, q Query, args ...interface{}) (pgxV4.Rows, error) {
	tx, ok := ctx.Value(TxKey).(pgxV4.Tx)
	if ok {
		return tx.Query(ctx, q.QueryRaw, args...)
	}

	return d.pool.Query(ctx, q.QueryRaw, args...)
}

func (d *db) QueryRowContext(ctx context.Context, q Query, args ...interface{}) pgxV4.Row {
	tx, ok := ctx.Value(TxKey).(pgxV4.Tx)
	if ok {
		return tx.QueryRow(ctx, q.QueryRaw, args...)
	}

	return d.pool.QueryRow(ctx, q.QueryRaw, args...)
}

func (d *db) BeginTx(ctx context.Context, txOptions pgxV4.TxOptions) (pgxV4.Tx, error) {
	return d.pool.BeginTx(ctx, txOptions)
}

func (d *db) Ping(ctx context.Context) error {
	return d.pool.Ping(ctx)
}

func (d *db) Close() {
	d.pool.Close()
}

func GetContextTx(ctx context.Context, tx pgxV4.Tx) context.Context {
	return context.WithValue(ctx, TxKey, tx)
}
