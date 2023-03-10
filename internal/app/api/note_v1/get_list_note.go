package note_v1

import (
	"context"
	"database/sql"
	"fmt"

	desc "github.com/Journeyman150/note-service-api/pkg/note_v1"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

func (n *Note) GetListNote(ctx context.Context, req *desc.GetListNoteRequest) (*desc.GetListNoteResponse, error) {
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
		From(noteTable)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	row, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var title, text, author, createdAt string
	var updatedAt sql.NullString
	noteList := make([]*desc.GetNoteResponse, 0, 10)
	for row.Next() {
		err = row.Scan(&title, &text, &author, &createdAt, &updatedAt)
		note := desc.GetNoteResponse{
			Title:     title,
			Text:      text,
			Author:    author,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt.String,
		}
		noteList = append(noteList, &note)
		if err != nil {
			return nil, err
		}
	}

	fmt.Println("Get list note")
	fmt.Println()
	return &desc.GetListNoteResponse{
		Notes: noteList,
	}, nil
}
