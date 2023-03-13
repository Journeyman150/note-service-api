package note

import (
	"context"
	"database/sql"
	"time"

	desc "github.com/Journeyman150/note-service-api/pkg/note_v1"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	tableName = "note"
)

type Repository interface {
	CreateNote(ctx context.Context, req *desc.CreateNoteRequest) (int64, error)
	GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponse, error)
	GetListNote(ctx context.Context, req *desc.GetListNoteRequest) (*desc.GetListNoteResponse, error)
	UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (sql.Result, error)
	DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) (sql.Result, error)
}

type repository struct {
	db *sqlx.DB
}

func NewNoteRepository(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r repository) CreateNote(ctx context.Context, req *desc.CreateNoteRequest) (int64, error) {
	builder := sq.Insert(tableName).
		Columns("title, text, author, email").
		Values(req.GetTitle(), req.GetText(), req.GetAuthor(), req.GetEmail()).
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
	builder := sq.Select("title, text, author, email, created_at, updated_at").
		From(tableName).
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

	var title, text, author, email string
	var createdAt time.Time
	var nullableUpdatedAt sql.NullTime
	var updatedAt time.Time

	if row.Next() {
		err = row.Scan(&title, &text, &author, &email, &createdAt, &nullableUpdatedAt)
		if err != nil {
			return nil, err
		}
		if nullableUpdatedAt.Valid {
			updatedAt = nullableUpdatedAt.Time
		}
	} else {
		return &desc.GetNoteResponse{}, nil
	}

	return &desc.GetNoteResponse{
		Id:        req.GetId(),
		Title:     title,
		Text:      text,
		Author:    author,
		Email:     email,
		CreatedAt: timestamppb.New(createdAt),
		UpdatedAt: timestamppb.New(updatedAt),
	}, nil
}

func (r repository) GetListNote(ctx context.Context, req *desc.GetListNoteRequest) (*desc.GetListNoteResponse, error) {
	builder := sq.Select("id, title, text, author, email, created_at, updated_at").
		From(tableName).
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
	var title, text, author, email string
	var createdAt time.Time
	var nullableUpdatedAt sql.NullTime
	noteList := make([]*desc.GetNoteResponse, 0, 10)
	for row.Next() {
		err = row.Scan(&id, &title, &text, &author, &email, &createdAt, &nullableUpdatedAt)
		var updatedAt time.Time
		if nullableUpdatedAt.Valid {
			updatedAt = nullableUpdatedAt.Time
		}
		noteList = append(noteList, &desc.GetNoteResponse{
			Id:        id,
			Title:     title,
			Text:      text,
			Author:    author,
			Email:     email,
			CreatedAt: timestamppb.New(createdAt),
			UpdatedAt: timestamppb.New(updatedAt),
		})
		if err != nil {
			return nil, err
		}
	}

	return &desc.GetListNoteResponse{
		Notes: noteList,
	}, nil
}

func (r repository) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (sql.Result, error) {
	builder := sq.Update(tableName)

	if len(req.GetTitle()) != 0 {
		builder = builder.Set("title", req.GetTitle())
	}
	if len(req.GetText()) != 0 {
		builder = builder.Set("text", req.GetText())
	}
	if len(req.GetAuthor()) != 0 {
		builder = builder.Set("author", req.GetAuthor())
	}
	if len(req.GetEmail()) != 0 {
		builder = builder.Set("email", req.GetEmail())
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

func (r repository) DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) (sql.Result, error) {
	builder := sq.Delete(tableName).
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
