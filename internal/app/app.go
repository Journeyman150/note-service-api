package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/Journeyman150/note-service-api/internal/app/api/note_v1"
	desc "github.com/Journeyman150/note-service-api/pkg/note_v1"
	grpcValidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	serviceProvider *serviceProvider
	pathConfig      string
	noteImpl        *note_v1.Note
	mux             *runtime.ServeMux
	grpcServer      *grpc.Server
	grpcAddress     string
	httpAddress     string
}

func NewApp(ctx context.Context, pathConfig string) (*App, error) {
	a := &App{
		pathConfig: pathConfig,
	}
	err := a.initDeps(ctx)

	return a, err
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initServiceProvider,
		a.initServer,
		a.initGRPCServer,
		a.initHTTPServer,
	}

	for _, fn := range inits {
		err := fn(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) Run() error {
	defer func() {
		a.serviceProvider.db.Close()
	}()
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		err := a.startGRPC()
		if err != nil {
			log.Fatalf("error in startGrpc(): %s", err)
		}
	}()
	go func() {
		defer wg.Done()
		err := a.startHttp()
		if err != nil {
			log.Fatalf("error in startHttp(): %s", err)
		}
	}()

	wg.Wait()
	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider(a.pathConfig)

	return nil
}

func (a *App) initServer(ctx context.Context) error {
	a.noteImpl = note_v1.NewNote(a.serviceProvider.GetNoteService(ctx))

	return nil
}

func (a *App) initGRPCServer(_ context.Context) error {
	a.grpcServer = grpc.NewServer(
		grpc.UnaryInterceptor(grpcValidator.UnaryServerInterceptor()),
	)
	desc.RegisterNoteV1Server(a.grpcServer, a.noteImpl)

	a.grpcAddress = net.JoinHostPort(
		a.serviceProvider.GetConfig().GRPC.Host,
		a.serviceProvider.GetConfig().GRPC.Port,
	)

	return nil
}

func (a *App) initHTTPServer(ctx context.Context) error {
	a.mux = runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err := desc.RegisterNoteV1HandlerFromEndpoint(ctx, a.mux, a.grpcAddress, opts)
	if err != nil {
		return err
	}

	a.httpAddress = net.JoinHostPort(
		a.serviceProvider.GetConfig().HTTP.Host,
		a.serviceProvider.GetConfig().HTTP.Port,
	)

	return nil
}

func (a *App) startGRPC() error {
	list, err := net.Listen("tcp", a.grpcAddress)
	if err != nil {
		log.Fatalf("failed to mapping port: %s", err.Error())
	}

	fmt.Printf("grpc server is running on host%s\n", a.grpcAddress)

	if err = a.grpcServer.Serve(list); err != nil {
		log.Fatalf("failed to serve: %s", err.Error())
		return err
	}

	return nil
}

func (a *App) startHttp() error {
	fmt.Printf("http server is running on host%s\n", a.httpAddress)
	if err := http.ListenAndServe(a.httpAddress, a.mux); err != nil {
		log.Fatalf("failed to serve: %s", err.Error())
		return err
	}

	return nil
}
