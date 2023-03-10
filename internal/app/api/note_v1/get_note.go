package note_v1

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	desc "github.com/Journeyman150/note-service-api/pkg/note_v1"
	sq "github.com/Masterminds/squirrel"
)

func (n *Note) GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponse, error) {
	dbDsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, dbUser, dbPassword, dbName, sslMode,
	)
	db, err := sqlx.Open("pgx", dbDsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	builder := sq.Select("title, text, author, created_at, updated_at").
		From(noteTable).
		Where("id = ?", req.GetId()).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	row, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	row.Next()
	var title, text, author, createdAt string
	var updatedAt sql.NullString
	err = row.Scan(&title, &text, &author, &createdAt, &updatedAt)
	if err != nil {
		return nil, err
	}

	fmt.Println("Get note")
	fmt.Println("Note Id:", req.GetId())
	fmt.Println()

	return &desc.GetNoteResponse{
		Title:     title,
		Text:      text,
		Author:    author,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt.String,
	}, nil
}
