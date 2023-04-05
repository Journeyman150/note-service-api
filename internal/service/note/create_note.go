package note

import (
	"context"
	"fmt"

	"github.com/Journeyman150/note-service-api/internal/model"
)

func (s *Service) CreateNote(ctx context.Context, noteInfo *model.NoteInfo) (int64, error) {
	var noteId int64
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error

		noteId, errTx = s.noteRepository.Create(ctx, noteInfo)
		if errTx != nil {
			return errTx
		}

		log := &model.Log{
			NoteId: noteId,
			Msg:    fmt.Sprintf("note with id %d was created", noteId),
		}

		errTx = s.logRepository.Create(ctx, log)
		if errTx != nil {
			return errTx
		}

		return nil
	})
	if err != nil {
		return 0, err
	}

	return noteId, nil
}
