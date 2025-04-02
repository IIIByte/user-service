package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/IIIByte/user-service/config"
	"github.com/IIIByte/user-service/pkg/logger"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	logger  logger.Logger
	Pool    *pgxpool.Pool
	Builder squirrel.StatementBuilderType
}

// NewDB - инициализация подключения к PostgreSQL
func NewDB(ctx context.Context, logger logger.Logger, cfg *config.PostgresConfig) (*DB, error) {
	if cfg == nil {
		return nil, errors.New("Postgres (NewDB): postgres config is nil")
	}

	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	poolConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %v", err)
	}
	poolConfig.MaxConns = int32(cfg.MaxConnections)

	// Создаем пул соединений
	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, err
	}

	return &DB{
		logger:  logger,
		Pool:    pool,
		Builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}, nil
}
