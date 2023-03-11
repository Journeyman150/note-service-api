package note_v1

import (
	"context"
	"fmt"

	desc "github.com/Journeyman150/note-service-api/pkg/note_v1"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

func (n *Note) DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) (*desc.Empty, error) {
	dbDsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, dbUser, dbPassword, dbName, sslMode,
	)
	db, err := sqlx.Open("pgx", dbDsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	builder := sq.Delete(noteTable).
		Where(sq.Eq{"id": req.GetId()}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	result, err := db.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	fmt.Println("Delete note")
	fmt.Println("Note Id:", req.GetId())
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	fmt.Println("Rows affected: ", rowsAffected)
	fmt.Println()
	return &desc.Empty{}, nil
}
