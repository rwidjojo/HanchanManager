package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// DB wraps a pgxpool.Pool and implements all store interfaces.
// Each domain area is a method receiver on DB, keeping the pool
// in one place while allowing clean separation by file if needed.
type DB struct {
	pool *pgxpool.Pool
}

func NewDB(pool *pgxpool.Pool) *DB {
	return &DB{pool: pool}
}

func (d *DB) Pool() *pgxpool.Pool {
	return d.pool
}

// Connect opens a pgxpool connection and verifies it with a ping.
func Connect(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("pgxpool.New: %w", err)
	}
	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("pool.Ping: %w", err)
	}
	return pool, nil
}
