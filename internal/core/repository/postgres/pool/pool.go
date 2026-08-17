package core_postgres_pool

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ConnectionPool struct {
	*pgxpool.Pool
}

func NewConnnectionPool(
	ctx context.Context,
	config Config,
) (*ConnectionPool, error) {
	connectionString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sllmode=disabled",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)

	pgxconfig, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return nil, fmt.Errorf("parse pgxconfig %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, pgxconfig)

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("pgxpool ping %w", err)
	}

	return &ConnectionPool{
		Pool: pool,
	}, nil
}
