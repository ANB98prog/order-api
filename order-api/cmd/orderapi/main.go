package main

import (
	"context"
	"github.com/ANB98prog/order-api/internal/app"
	"github.com/ANB98prog/order-api/internal/config"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := config.Load()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	// для Gracefully shutdown
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	application, err := app.New(ctx, logger, cfg)
	if err != nil {
		logger.Error("failed to initialize application", slog.Any("error", err))
		os.Exit(1)
	}
	defer func() {
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := application.Shutdown(shutdownCtx); err != nil {
			logger.Error("failed to shutdown application", slog.Any("error", err))
		}
	}()

	if err := application.Run(ctx); err != nil {
		logger.Error("application stopped with error", slog.Any("error", err))
		os.Exit(1)
	}
}
