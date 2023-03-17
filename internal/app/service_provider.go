package app

import (
	"context"
	"log"

	"github.com/Journeyman150/note-service-api/internal/config"
	"github.com/Journeyman150/note-service-api/internal/pkg/db"
	noteRepo "github.com/Journeyman150/note-service-api/internal/repository/note"
	noteService "github.com/Journeyman150/note-service-api/internal/service/note"
)

type serviceProvider struct {
	db         db.Client
	configPath string
	config     *config.Config

	noteRepository noteRepo.Repository
	noteService    *noteService.Service
}

func newServiceProvider(configPath string) *serviceProvider {
	return &serviceProvider{
		configPath: configPath,
	}
}

func (s *serviceProvider) GetDB(ctx context.Context) db.Client {
	if s.db == nil {
		cfg, err := s.GetConfig().GetDBConfig()
		if err != nil {
			log.Fatalf("failed to get db config: %s", err.Error())
		}

		dbc, err := db.NewClient(ctx, cfg)
		if err != nil {
			log.Fatalf("can't connect to db err: %s", err.Error())
		}
		s.db = dbc
	}

	return s.db
}

func (s *serviceProvider) GetConfig() *config.Config {
	if s.config == nil {
		cfg, err := config.NewConfig(s.configPath)
		if err != nil {
			log.Fatalf("failed to get config: %s", err.Error())
		}

		s.config = cfg
	}

	return s.config
}

func (s *serviceProvider) GetNoteRepository(ctx context.Context) noteRepo.Repository {
	if s.noteRepository == nil {
		s.noteRepository = noteRepo.NewNoteRepository(s.GetDB(ctx))
	}

	return s.noteRepository
}

func (s *serviceProvider) GetNoteService(ctx context.Context) *noteService.Service {
	if s.noteService == nil {
		s.noteService = noteService.NewService(s.GetNoteRepository(ctx))
	}

	return s.noteService
}
