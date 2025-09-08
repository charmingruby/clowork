package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/charmingruby/clowork/config"
	"github.com/charmingruby/clowork/internal/chat"
	"github.com/charmingruby/clowork/internal/platform"
	"github.com/charmingruby/clowork/pkg/database/postgres"
	"github.com/charmingruby/clowork/pkg/delivery/http/grpc"
	"github.com/charmingruby/clowork/pkg/delivery/http/rest"
	"github.com/charmingruby/clowork/pkg/telemetry/logger"

	"github.com/joho/godotenv"
)

func main() {
	log := logger.New()

	if err := godotenv.Load(); err != nil {
		log.Warn("failed to find .env file", "error", err)
	}

	log.Info("loading environment variables...")

	cfg, err := config.New()
	if err != nil {
		log.Error("failed to loading environment variables", "error", err)
		failAndExit(log, nil, nil, nil)
	}

	log.Info("environment variables loaded")

	logLevel := logger.ChangeLevel(cfg.LogLevel)

	log.Info("log level configured", "level", logLevel)

	log.Info("connecting to Postgres...")

	db, err := postgres.New(log, cfg.PostgresURL)
	if err != nil {
		log.Error("failed connect to Postgres", "error", err)
		failAndExit(log, nil, nil, nil)
	}

	log.Info("connected to Postgres successfully")

	restSrv, r := rest.New(cfg.RestServerPort)

	grpcAddr := fmt.Sprintf("%s:%s", cfg.GRPCServerHost, cfg.GRPCServerPort)
	grpcSrv := grpc.New(grpcAddr)

	platform.New(r, db)

	if err := chat.New(log, db.Conn, grpcSrv.Conn); err != nil {
		log.Error("failed create Chat module", "error", err)
		failAndExit(log, restSrv, &grpcSrv, db)
	}

	go func() {
		log.Info("REST server is running...", "port", cfg.RestServerPort)

		if err := restSrv.Start(); err != nil {
			log.Error("failed starting REST server", "error", err)
			failAndExit(log, restSrv, nil, db)
		}
	}()

	go func() {
		log.Info("gRPC server is running...", "port", cfg.GRPCServerPort)

		if err := grpcSrv.Start(); err != nil {
			log.Error("failed starting gRPC server", "error", err)
			failAndExit(log, restSrv, &grpcSrv, db)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Info("received an interrupt signal")

	log.Info("starting graceful shutdown...")

	signal := gracefulShutdown(log, restSrv, &grpcSrv, db)

	log.Info(fmt.Sprintf("gracefully shutdown, with code %d", signal))

	os.Exit(signal)
}

func failAndExit(log *logger.Logger, restSrv *rest.Server, grpcSrv *grpc.Server, db *postgres.Client) {
	gracefulShutdown(log, restSrv, grpcSrv, db)
	os.Exit(1)
}

func gracefulShutdown(log *logger.Logger, restSrv *rest.Server, grpcSrv *grpc.Server, db *postgres.Client) int {
	parentCtx := context.Background()

	var hasError bool

	if restSrv != nil {
		ctx, cancel := context.WithTimeout(parentCtx, 15*time.Second)
		defer cancel()

		if err := restSrv.Close(ctx); err != nil {
			log.Error("error closing REST server", "error", err)
			hasError = true
		}
	}

	if grpcSrv != nil {
		grpcSrv.Close()
	}

	if db != nil {
		if err := db.Close(); err != nil {
			log.Error("error closing Postgres connection", "error", err)
			hasError = true
		}
	}

	if hasError {
		return 1
	}

	return 0
}
