package note

import (
	"context"
	"time"

	"github.com/Journeyman150/note-service-api/internal/model"
	"github.com/Journeyman150/note-service-api/internal/pkg/db"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgconn"
)

const (
	tableName = "note"
)

type Repository interface {
	CreateNote(ctx context.Context, noteInfo *model.NoteInfo) (int64, error)
	GetNote(ctx context.Context, id int64) (*model.GetNoteResponse, error)
	GetListNote(ctx context.Context) ([]*model.GetNoteResponse, error)
	UpdateNote(ctx context.Context, req *model.UpdateNoteRequest) (pgconn.CommandTag, error)
	DeleteNote(ctx context.Context, id int64) (pgconn.CommandTag, error)
}

type repository struct {
	client db.Client
}

func NewNoteRepository(client db.Client) Repository {
	return &repository{
		client: client,
	}
}

func (r repository) CreateNote(ctx context.Context, noteInfo *model.NoteInfo) (int64, error) {
	builder := sq.Insert(tableName).
		Columns("title, text, author, email").
		Values(noteInfo.Title, noteInfo.Text, noteInfo.Author, noteInfo.Email).
		Suffix("returning id").
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "CreateNote",
		QueryRaw: query,
	}

	row, err := r.client.DB().QueryContext(ctx, q, args...)
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

func (r repository) GetNote(ctx context.Context, id int64) (*model.GetNoteResponse, error) {
	builder := sq.Select("id, title, text, author, email, created_at, updated_at").
		From(tableName).
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "GetNote",
		QueryRaw: query,
	}

	getNoteResponse := new(model.GetNoteResponse)
	err = r.client.DB().GetContext(ctx, getNoteResponse, q, args...)
	if err != nil {
		return nil, err
	}

	return getNoteResponse, nil
}

func (r repository) GetListNote(ctx context.Context) ([]*model.GetNoteResponse, error) {
	builder := sq.Select("id, title, text, author, email, created_at, updated_at").
		From(tableName).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "GetListNote",
		QueryRaw: query,
	}

	var notes []*model.GetNoteResponse
	err = r.client.DB().SelectContext(ctx, &notes, q, args...)

	return notes, nil
}

func (r repository) UpdateNote(ctx context.Context, req *model.UpdateNoteRequest) (pgconn.CommandTag, error) {
	builder := sq.Update(tableName)

	if req.Title.Valid {
		builder = builder.Set("title", req.Title)
	}
	if req.Text.Valid {
		builder = builder.Set("text", req.Text)
	}
	if req.Author.Valid {
		builder = builder.Set("author", req.Author)
	}
	if req.Email.Valid {
		builder = builder.Set("email", req.Email)
	}

	builder = builder.Set("updated_at", time.Now().UTC().Format(time.RFC3339)).
		Where(sq.Eq{"id": req.Id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "UpdateNote",
		QueryRaw: query,
	}

	result, err := r.client.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r repository) DeleteNote(ctx context.Context, id int64) (pgconn.CommandTag, error) {
	builder := sq.Delete(tableName).
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "DeleteNote",
		QueryRaw: query,
	}

	result, err := r.client.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}

	return result, nil
}
