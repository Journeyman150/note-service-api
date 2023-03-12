package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/jmoiron/sqlx"

	"github.com/Journeyman150/note-service-api/internal/repository"

	"github.com/Journeyman150/note-service-api/internal/service/note"

	"github.com/Journeyman150/note-service-api/internal/app/api/note_v1"
	desc "github.com/Journeyman150/note-service-api/pkg/note_v1"
	grpcValidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	hostGrpc = "localhost:50051"
	hostHttp = "localhost:8090"
)

const (
	host       = "localhost"
	port       = "54321"
	dbUser     = "note-service-user"
	dbPassword = "note-service-password"
	dbName     = "note-service"
	sslMode    = "disable"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		startGRPC()
	}()
	go func() {
		defer wg.Done()
		startHttp()
	}()

	wg.Wait()
}

func startGRPC() error {
	list, err := net.Listen("tcp", hostGrpc)
	if err != nil {
		log.Fatalf("failed to mapping port: %s", err.Error())
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpcValidator.UnaryServerInterceptor()),
	)

	dbDsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, dbUser, dbPassword, dbName, sslMode,
	)
	db, err := sqlx.Open("pgx", dbDsn)
	if err != nil {
		return err
	}
	defer db.Close()

	noteRepository := repository.NewNoteRepository(db)
	noteService := note.NewService(noteRepository)

	desc.RegisterNoteV1Server(s, note_v1.NewNote(noteService))
	fmt.Printf("grpc server is running on port%s\n", hostGrpc)

	if err = s.Serve(list); err != nil {
		log.Fatalf("failed to serve: %s", err.Error())
		return err
	}

	return nil
}

func startHttp() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err := desc.RegisterNoteV1HandlerFromEndpoint(ctx, mux, hostGrpc, opts)
	if err != nil {
		return err
	}
	fmt.Printf("http server is running on port%s\n", hostHttp)
	if err = http.ListenAndServe(hostHttp, mux); err != nil {
		log.Fatalf("failed to serve: %s", err.Error())
		return err
	}
	return nil
}
