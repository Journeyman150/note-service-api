package note_v1

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	desc "github.com/Journeyman150/note-service-api/pkg/note_v1"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"google.golang.org/protobuf/types/known/timestamppb"
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
		From(noteTable).
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

	var title, text, author string
	var createdAt time.Time
	var updatedAt sql.NullTime
	noteList := make([]*desc.GetNoteResponse, 0, 10)
	for row.Next() {
		err = row.Scan(&title, &text, &author, &createdAt, &updatedAt)
		var checkedUpdatedAt *timestamppb.Timestamp
		if updatedAt.Valid {
			checkedUpdatedAt = timestamppb.New(updatedAt.Time)
		}
		noteList = append(noteList, &desc.GetNoteResponse{
			Title:     title,
			Text:      text,
			Author:    author,
			CreatedAt: timestamppb.New(createdAt),
			UpdatedAt: checkedUpdatedAt,
		})
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
