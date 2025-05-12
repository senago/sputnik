package main

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/senago/sputnik/cmd/sputnik/closer"
	"github.com/senago/sputnik/internal/details/db"
)

type App struct {
	config Config
	closer *closer.Closer

	dbPool *pgxpool.Pool
	db     *db.DB
}

func NewApp(ctx context.Context, config Config) (*App, error) {
	lifoCloser := closer.New()

	dbPool, err := pgxpool.New(ctx, config.DSN)
	if err != nil {
		return nil, errors.Wrap(err, "pgxpool.New")
	}

	lifoCloser.Add(func() error {
		dbPool.Close()
		return nil
	})

	return &App{
		config: config,
		closer: lifoCloser,
		dbPool: dbPool,
		db:     db.New(dbPool),
	}, nil
}

func (a *App) Close() error {
	return a.closer.Close()
}
