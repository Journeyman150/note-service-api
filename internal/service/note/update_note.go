package note

import (
	"context"
	"fmt"

	"github.com/Journeyman150/note-service-api/internal/model"
)

func (s *Service) Update(ctx context.Context, noteId int64, noteInfo *model.UpdateNoteInfo) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {

		errTx := s.noteRepository.Update(ctx, noteId, noteInfo)
		if errTx != nil {
			return errTx
		}

		log := &model.Log{
			NoteId: noteId,
			Msg:    fmt.Sprintf("note with id %d was updated", noteId),
		}

		errTx = s.logRepository.Create(ctx, log)
		if errTx != nil {
			return errTx
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
