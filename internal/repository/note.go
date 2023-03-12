package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/Journeyman150/note-service-api/internal/repository/table"
	desc "github.com/Journeyman150/note-service-api/pkg/note_v1"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type NoteRepository interface {
	CreateNote(ctx context.Context, req *desc.CreateNoteRequest) (int64, error)
	GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponse, error)
	GetListNote(ctx context.Context, req *desc.GetListNoteRequest) (*desc.GetListNoteResponse, error)
	UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (sql.Result, error)
}

type repository struct {
	db *sqlx.DB
}

func NewNoteRepository(db *sqlx.DB) NoteRepository {
	return &repository{
		db: db,
	}
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

func (r repository) GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponse, error) {
	builder := sq.Select("title, text, author, created_at, updated_at").
		From(table.Note).
		Where(sq.Eq{"id": req.GetId()}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	row, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	row.Next()
	var title, text, author string
	var createdAt time.Time
	var nullableUpdatedAt sql.NullTime
	err = row.Scan(&title, &text, &author, &createdAt, &nullableUpdatedAt)
	if err != nil {
		return nil, err
	}
	var updatedAt time.Time
	if nullableUpdatedAt.Valid {
		updatedAt = nullableUpdatedAt.Time
	}
	return &desc.GetNoteResponse{
		Id:        req.GetId(),
		Title:     title,
		Text:      text,
		Author:    author,
		CreatedAt: timestamppb.New(createdAt),
		UpdatedAt: timestamppb.New(updatedAt),
	}, nil
}

func (r repository) GetListNote(ctx context.Context, req *desc.GetListNoteRequest) (*desc.GetListNoteResponse, error) {
	builder := sq.Select("id, title, text, author, created_at, updated_at").
		From(table.Note).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	row, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var id int64
	var title, text, author string
	var createdAt time.Time
	var nullableUpdatedAt sql.NullTime
	noteList := make([]*desc.GetNoteResponse, 0, 10)
	for row.Next() {
		err = row.Scan(&id, &title, &text, &author, &createdAt, &nullableUpdatedAt)
		var updatedAt time.Time
		if nullableUpdatedAt.Valid {
			updatedAt = nullableUpdatedAt.Time
		}
		noteList = append(noteList, &desc.GetNoteResponse{
			Id:        id,
			Title:     title,
			Text:      text,
			Author:    author,
			CreatedAt: timestamppb.New(createdAt),
			UpdatedAt: timestamppb.New(updatedAt),
		})
		if err != nil {
			return nil, err
		}
	}
	return &desc.GetListNoteResponse{
		Notes: noteList,
	}, err
}

func (r repository) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (sql.Result, error) {
	builder := sq.Update(table.Note)

	if len(req.GetTitle()) != 0 {
		builder = builder.Set("title", req.GetTitle())
	}
	if len(req.GetText()) != 0 {
		builder = builder.Set("text", req.GetText())
	}
	if len(req.GetAuthor()) != 0 {
		builder = builder.Set("author", req.GetAuthor())
	}

	builder = builder.Set("updated_at", time.Now().UTC().Format(time.RFC3339)).
		Where(sq.Eq{"id": req.GetId()}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	result, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return result, nil
}
