package note_v1

import (
	"context"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	desc "github.com/Journeyman150/note-service-api/pkg/note_v1"
)

func (n *Note) Update(ctx context.Context, req *desc.UpdateNoteRequest) (*desc.Empty, error) {
	dbDsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, dbUser, dbPassword, dbName, sslMode,
	)
	db, err := sqlx.Open("pgx", dbDsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	builder := sq.Update(noteTable)

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
		Where("id = ?", req.GetId()).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	result, err := db.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	fmt.Println("Update note")
	fmt.Println("Note Id:", req.GetId())
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	fmt.Println("Rows affected: ", rowsAffected)
	fmt.Println()
	return &desc.Empty{}, nil
}
