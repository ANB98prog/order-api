package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/ANB98prog/order-api/internal/config"
	"github.com/ANB98prog/order-api/internal/server/orderapi"
	"github.com/ANB98prog/order-api/pkg/db"
	"github.com/redis/go-redis/v9"
	"log/slog"
)

type OrderApiApp struct {
	logger      *slog.Logger
	cfg         config.Config
	postgresDb  *db.Db
	redisClient *redis.Client
	server      *orderapi.Server
}

func New(ctx context.Context, logger *slog.Logger, cfg config.Config) (*OrderApiApp, error) {
	postgresDb, err := db.NewDb(cfg.Db)
	if err != nil {
		return nil, fmt.Errorf("coonect postgres: %w", err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Cache.Addr,
		Password: cfg.Cache.Password,
		DB:       cfg.Cache.Db,
	})

	if err = redisClient.Ping(ctx).Err(); err != nil {
		_ = redisClient.Close()
		return nil, fmt.Errorf("ping redis: %w", err)
	}

	deps := orderapi.Dependencies{
		Logger:      logger,
		Config:      cfg,
		RedisClient: redisClient,
		DB:          postgresDb,
	}
	orderApi := orderapi.New(deps)

	return &OrderApiApp{
		logger:      logger,
		cfg:         cfg,
		postgresDb:  postgresDb,
		redisClient: redisClient,
		server:      orderApi,
	}, nil
}

func (a *OrderApiApp) Run(ctx context.Context) error {
	a.logger.Info("starting OrderApi server")

	if err := a.server.Start(ctx); err != nil && !errors.Is(err, context.Canceled) {
		return fmt.Errorf("server start: %w", err)
	}

	return nil
}

func (a *OrderApiApp) Shutdown(ctx context.Context) error {
	var shutdownErr error

	a.logger.Info("shutting down OrderApi server")

	if a.server != nil {
		shutdownErr = errors.Join(shutdownErr, a.server.Stop(ctx))
	}

	if a.redisClient != nil {
		shutdownErr = errors.Join(shutdownErr, a.redisClient.Close())
	}

	return shutdownErr
}
