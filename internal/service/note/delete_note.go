package note

import (
	"context"
)

func (s *Service) DeleteNote(ctx context.Context, noteId int64) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {

		errTx := s.logRepository.Delete(ctx, noteId)
		if errTx != nil {
			return errTx
		}

		errTx = s.noteRepository.Delete(ctx, noteId)
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
