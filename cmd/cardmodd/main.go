package main

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/iamnande/cardmod/internal/config"
	"github.com/iamnande/cardmod/internal/database"
	"github.com/iamnande/cardmod/internal/logger"
	"github.com/iamnande/cardmod/internal/repositories"
	"github.com/iamnande/cardmod/internal/server/grpc"
	"github.com/iamnande/cardmod/internal/server/rest"
)

var (
	// ServiceName is the name of the service.
	ServiceName = "cardmod"

	// ServiceVersion is the version of the service being deployed.
	// note: this should be overwritten by the linker, using ldflags, during the compilation process.
	ServiceVersion = "v1.0.0-dev"
)

func main() {

	// api: initialize root context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// api: initialize application config
	cfg := config.MustLoad()

	// api: initialize logger
	log := logger.NewLogger(ServiceName, ServiceVersion)

	// api: initialize database client
	dbConn, err := sql.Open("pgx", cfg.DatabaseEndpoint)
	if err != nil {
		log.Error(err, "failed to connect to the database")
		os.Exit(1)
	}
	driver := entsql.OpenDB("postgres", dbConn)
	dbClient := database.NewClient(database.Driver(driver))

	// api: initialize repositories
	cardRepository := repositories.NewCardRepository(dbClient)
	magicRepository := repositories.NewMagicRepository(dbClient)

	// api: initialize the gRPC server
	grpcServer := grpc.NewServer(&grpc.ServerConfig{
		Logger:  log,
		Version: "v1.0.0",
		Port:    cfg.GRPCPort,

		// Repositories
		CardRepository:  cardRepository,
		MagicRepository: magicRepository,
	})

	// api: initialize REST server
	restServer, err := rest.NewServer(&rest.ServerConfig{
		Context:      ctx,
		Logger:       log,
		GRPCEndpoint: cfg.GRPCPort,
		RESTEndpoint: cfg.RESTPort,
	})
	if err != nil {
		log.Error(err, "failed to initialize REST server instance")
		os.Exit(1)
	}

	// api: start gRPC listener
	go func() {
		log.Info("starting gRPC server")
		if err = grpcServer.Serve(); err != nil {
			log.Error(err, "failed to start gRPC server")
			os.Exit(1)
		}
	}()

	// api: start REST listener
	go func() {
		log.Info("starting REST server")
		if err = restServer.Serve(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				log.Error(err, "failed to start REST server")
				os.Exit(1)
			}
		}
	}()

	// api: support SIG* signals to handle gracefully stopping gRPC and REST servers
	doneChan := make(chan os.Signal, 1)
	signal.Notify(doneChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP)
	<-doneChan
	log.Info("shutdown signal received")
	cancel()

	// api: setup graceful gRPC server stop
	grpcStop := stopServer(func(ctx context.Context) {
		log.Info("stopping gRPC server")
		grpcServer.Stop()
	})

	// api: setup graceful REST server stop
	restStop := stopServer(func(ctx context.Context) {
		log.Info("stopping REST server")
		if err = restServer.Stop(ctx); err != nil {
			log.Error(err, "failed to stop REST server")
		}
	})

	// api: gracefully stop all the things
	<-grpcStop.Done()
	<-restStop.Done()

	// api: we've completely closed everything, goodbye
	log.Info("goodbye, neo.")

}

func stopServer(stopHandler func(context.Context)) context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	go func() {
		stopHandler(ctx)
		cancel()
	}()
	return ctx
}
