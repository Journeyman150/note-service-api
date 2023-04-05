package note

import (
	"github.com/Journeyman150/note-service-api/internal/pkg/db/transaction"
	noteLog "github.com/Journeyman150/note-service-api/internal/repository/log"
	"github.com/Journeyman150/note-service-api/internal/repository/note"
)

type Service struct {
	noteRepository note.Repository
	logRepository  noteLog.Repository
	txManager      transaction.Manager
}

func NewService(
	noteRepository note.Repository,
	logRepository noteLog.Repository,
	txManager transaction.Manager,
) *Service {
	return &Service{
		noteRepository: noteRepository,
		logRepository:  logRepository,
		txManager:      txManager,
	}
}

func NewMockNoteService(deps ...interface{}) *Service {
	service := Service{}

	for _, v := range deps {
		switch s := v.(type) {
		case note.Repository:
			service.noteRepository = s
		}
	}

	return &service
}
