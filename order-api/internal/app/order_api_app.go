package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/ANB98prog/order-api/internal/config"
	"github.com/ANB98prog/order-api/internal/server/orderapi"
	"github.com/ANB98prog/order-api/pkg/db"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"log/slog"
)

type OrderApiApp struct {
	logger      *slog.Logger
	cfg         config.Config
	postgresDb  *pgxpool.Pool
	redisClient *redis.Client
	server      *orderapi.Server
}

func New(ctx context.Context, logger *slog.Logger, cfg config.Config) (*OrderApiApp, error) {
	postgresPool, err := pgxpool.New(ctx, cfg.Db.Dsn)
	if err != nil {
		return nil, fmt.Errorf("create postgres pool: %w", err)
	}

	if err = postgresPool.Ping(ctx); err != nil {
		postgresPool.Close()
		return nil, fmt.Errorf("ping postgres pool: %w", err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Cache.Addr,
		Password: cfg.Cache.Password,
		DB:       cfg.Cache.Db,
	})

	if err = redisClient.Ping(ctx).Err(); err != nil {
		postgresPool.Close()
		_ = redisClient.Close()
		return nil, fmt.Errorf("ping redis: %w", err)
	}

	pgdb := db.NewDb(cfg.Db)

	deps := orderapi.Dependencies{
		Logger:      logger,
		Config:      cfg,
		RedisClient: redisClient,
		DB:          pgdb,
	}
	orderApi := orderapi.New(deps)

	return &OrderApiApp{
		logger:      logger,
		cfg:         cfg,
		postgresDb:  postgresPool,
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

	if a.postgresDb != nil {
		a.postgresDb.Close()
	}

	return shutdownErr
}
