package note

//go:generate mockgen --build_flags=--mod=mod -destination=mocks/mock_repository.go -package=mocks . Repository

import (
	"context"
	"time"

	"github.com/Journeyman150/note-service-api/internal/model"
	"github.com/Journeyman150/note-service-api/internal/pkg/db"
	sq "github.com/Masterminds/squirrel"
)

const (
	tableName = "note"
)

type Repository interface {
	Create(ctx context.Context, noteInfo *model.NoteInfo) (int64, error)
	Get(ctx context.Context, id int64) (*model.Note, error)
	GetList(ctx context.Context) ([]*model.Note, error)
	Update(ctx context.Context, id int64, req *model.UpdateNoteInfo) error
	Delete(ctx context.Context, id int64) error
}

type repository struct {
	client db.Client
}

func NewNoteRepository(client db.Client) Repository {
	return &repository{
		client: client,
	}
}

func (r *repository) Create(ctx context.Context, noteInfo *model.NoteInfo) (int64, error) {
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

func (r *repository) Get(ctx context.Context, id int64) (*model.Note, error) {
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

	getNoteResponse := new(model.Note)
	err = r.client.DB().GetContext(ctx, getNoteResponse, q, args...)
	if err != nil {
		return nil, err
	}

	return getNoteResponse, nil
}

func (r *repository) GetList(ctx context.Context) ([]*model.Note, error) {
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

	var notes []*model.Note
	err = r.client.DB().SelectContext(ctx, &notes, q, args...)
	if err != nil {
		return nil, err
	}

	return notes, nil
}

func (r *repository) Update(ctx context.Context, id int64, noteInfo *model.UpdateNoteInfo) error {
	builder := sq.Update(tableName)

	if noteInfo.Title.Valid {
		builder = builder.Set("title", noteInfo.Title)
	}
	if noteInfo.Text.Valid {
		builder = builder.Set("text", noteInfo.Text)
	}
	if noteInfo.Author.Valid {
		builder = builder.Set("author", noteInfo.Author)
	}
	if noteInfo.Email.Valid {
		builder = builder.Set("email", noteInfo.Email)
	}

	builder = builder.Set("updated_at", time.Now().UTC().Format(time.RFC3339)).
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "UpdateNote",
		QueryRaw: query,
	}

	_, err = r.client.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, id int64) error {
	builder := sq.Delete(tableName).
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "DeleteNote",
		QueryRaw: query,
	}

	_, err = r.client.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
