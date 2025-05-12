package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *DB {
	return &DB{
		pool: pool,
	}
}

func (db *DB) MasterIntoContext(ctx context.Context) context.Context {
	return WithConnContext(ctx, ConnContext{
		Resolver: db.connectMaster,
	})
}

func (db *DB) connectMaster(ctx context.Context) Executor {
	return db.pool
}
