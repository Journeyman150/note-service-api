package repository

import (
	"context"

	"github.com/Journeyman150/note-service-api/internal/repository/table"
	desc "github.com/Journeyman150/note-service-api/pkg/note_v1"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type NoteRepository interface {
	CreateNote(ctx context.Context, req *desc.CreateNoteRequest) (int64, error)
}

type repository struct {
	db *sqlx.DB
}

func (r repository) CreateNote(ctx context.Context, req *desc.CreateNoteRequest) (int64, error) {
	builder := sq.Insert(table.Note).
		Columns("title, text, author").
		Values(req.GetTitle(), req.GetText(), req.GetAuthor()).
		Suffix("returning id").
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	row, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	defer row.Close()

	row.Next()
	var id int64
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func NewNoteRepository(db *sqlx.DB) NoteRepository {
	return &repository{
		db: db,
	}
}
