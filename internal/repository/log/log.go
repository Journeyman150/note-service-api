package log

//go:generate mockgen --build_flags=--mod=mod -destination=mocks/mock_repository.go -package=mocks . Repository

import (
	"context"

	"github.com/Journeyman150/note-service-api/internal/model"
	"github.com/Journeyman150/note-service-api/internal/pkg/db"
	sq "github.com/Masterminds/squirrel"
)

const (
	tableName = "log"
)

type Repository interface {
	Create(ctx context.Context, log *model.Log) error
	Get(ctx context.Context, id int64) (*model.Log, error)
	GetList(ctx context.Context) ([]*model.Log, error)
	Delete(ctx context.Context, noteId int64) error
}

type repository struct {
	client db.Client
}

func NewLogRepository(client db.Client) Repository {
	return &repository{
		client: client,
	}
}

func (r *repository) Create(ctx context.Context, log *model.Log) error {
	builder := sq.Insert(tableName).
		Columns("note_id, msg").
		Values(log.NoteId, log.Msg).
		Suffix("returning id").
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "log.Create",
		QueryRaw: query,
	}

	_, err = r.client.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Get(ctx context.Context, id int64) (*model.Log, error) {
	builder := sq.Select("note_id, msg, created_at").
		From(tableName).
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "log.Get",
		QueryRaw: query,
	}

	log := &model.Log{}
	err = r.client.DB().GetContext(ctx, log, q, args...)
	if err != nil {
		return nil, err
	}

	return log, nil
}

func (r *repository) GetList(ctx context.Context) ([]*model.Log, error) {
	builder := sq.Select("id, msg, created_at").
		From(tableName).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "log.GetList",
		QueryRaw: query,
	}

	var logs []*model.Log
	err = r.client.DB().SelectContext(ctx, &logs, q, args...)
	if err != nil {
		return nil, err
	}

	return logs, nil
}

func (r *repository) Delete(ctx context.Context, noteId int64) error {
	builder := sq.Delete(tableName).
		Where(sq.Eq{"note_id": noteId}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "log.Delete",
		QueryRaw: query,
	}

	_, err = r.client.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
