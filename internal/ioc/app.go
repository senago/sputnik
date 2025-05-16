package ioc

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/senago/sputnik/cmd/sputnik/closer"
	"github.com/senago/sputnik/internal/details/db"
)

type Container struct {
	config Config
	closer *closer.Closer

	dbPool *pgxpool.Pool
	db     *db.DB
}

func New(ctx context.Context, config Config) (*Container, error) {
	lifoCloser := closer.New()

	dbPool, err := pgxpool.New(ctx, config.DSN)
	if err != nil {
		return nil, errors.Wrap(err, "pgxpool.New")
	}

	lifoCloser.Add(func() error {
		dbPool.Close()
		return nil
	})

	return &Container{
		config: config,
		closer: lifoCloser,
		dbPool: dbPool,
		db:     db.New(dbPool),
	}, nil
}

func (a *Container) Close() error {
	return a.closer.Close()
}
