package app

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc/reflection"
	"user-api/database/postgres"
	"user-api/internal/config"
	"user-api/internal/delivery/srv"
	"user-api/internal/repo"
	"user-api/internal/service"
	"user-api/pkg/gensqlc"
	"user-api/pkg/proto/gengrpc"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	grpcServer *grpc.Server
}

func New() *App {
	return &App{
		grpcServer: grpc.NewServer(),
	}
}

func (a *App) Run(ctx context.Context) error {
	cfg := config.MustLoad()

	serverCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	a.grpcServer = grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.ChainUnaryInterceptor(
			recovery.UnaryServerInterceptor(),
		),
	)
	reflection.Register(a.grpcServer)

	pgPool, err := postgres.NewPool(cfg, serverCtx)

	if err != nil {
		return err
	}

	queries := gensqlc.New(pgPool)

	defer pgPool.Close()

	// repositories
	userRepo := repo.NewUserRepo(pgPool, queries)

	// services
	tokenService := service.NewTokenService()
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(userService, tokenService, cfg)

	grpcServer := srv.NewGRPCServer(userService, authService)

	// gRPC services
	gengrpc.RegisterUserServiceServer(a.grpcServer, grpcServer)
	gengrpc.RegisterAuthServiceServer(a.grpcServer, grpcServer)

	listener, err := net.Listen("tcp", cfg.App.Host+":"+cfg.App.Port)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", cfg.App.Host+":"+cfg.App.Port, err)
		return err
	}

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("Starting gRPC server on %s:%s", cfg.App.Host, cfg.App.Port)
		if err := a.grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	<-stopChan
	log.Println("Shutting down gRPC server gracefully...")

	_, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	a.grpcServer.GracefulStop()
	log.Println("gRPC server stopped successfully")

	return nil
}
