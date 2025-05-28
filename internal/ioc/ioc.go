package ioc

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/senago/sputnik/cmd/sputnik/closer"
	"github.com/senago/sputnik/internal/details/db"
	"github.com/senago/sputnik/internal/details/queries"
	"github.com/senago/sputnik/internal/domain"
	"github.com/senago/sputnik/internal/port"
)

type Container struct {
	closer *closer.Closer

	db *db.DB
}

func New(ctx context.Context) (*Container, error) {
	lifoCloser := closer.New()

	return &Container{
		closer: lifoCloser,
	}, nil
}

func (c *Container) ConnectDB(ctx context.Context, dsn string) error {
	dbPool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return fmt.Errorf("pgxpool.New: %w", err)
	}

	c.closer.Add(func() error {
		dbPool.Close()
		return nil
	})

	if err := dbPool.Ping(ctx); err != nil {
		return fmt.Errorf("ping: %w", err)
	}

	c.db = db.New(dbPool)

	if err := c.applyMigrations(ctx); err != nil {
		return fmt.Errorf("applyMigrations: %w", err)
	}

	return nil
}

func (c *Container) Close() error {
	return c.closer.Close()
}

func (c *Container) PortInsertSatellite() port.InsertSatellite {
	return func(ctx context.Context, satellite domain.Satellite) error {
		ctx = c.db.MasterIntoContext(ctx)

		// Have you heard of transactions?

		if err := queries.InsertSatellites(ctx, []domain.Satellite{satellite}); err != nil {
			return fmt.Errorf("queries.InsertSatellites: %w", err)
		}

		if err := queries.InsertOrUpdateSatellitePositions(ctx, []domain.Satellite{satellite}); err != nil {
			return fmt.Errorf("queries.InsertOrUpdateSatellitePositions: %w", err)
		}

		return nil
	}
}

func (c *Container) PortUpdateSatellite() port.UpdateSatellites {
	return func(ctx context.Context, satellites []domain.Satellite) error {
		ctx = c.db.MasterIntoContext(ctx)

		// Have you heard of transactions?

		if err := queries.UpdateSatellites(ctx, satellites); err != nil {
			return fmt.Errorf("queries.UpdateSatellites: %w", err)
		}

		if err := queries.InsertOrUpdateSatellitePositions(ctx, satellites); err != nil {
			return fmt.Errorf("InsertOrUpdateSatellitePositions: %w", err)
		}

		return nil
	}
}

func (c *Container) PortGetSatellites() port.GetSatellites {
	return func(ctx context.Context) ([]domain.Satellite, error) {
		ctx = c.db.MasterIntoContext(ctx)

		return queries.GetSatellites(ctx, queries.GetSatellitesParams{})
	}
}

func (c *Container) PortGetSatellitesByNameLike() port.GetSatellitesByNameLike {
	return func(ctx context.Context, nameLike string) ([]domain.Satellite, error) {
		ctx = c.db.MasterIntoContext(ctx)

		return queries.GetSatellites(
			ctx,
			queries.GetSatellitesParams{
				NameLike: nameLike,
			},
		)
	}
}

func (c *Container) PortDeleteSatellites() port.DeleteSatellites {
	return func(ctx context.Context, satelliteIDs []domain.SatelliteID) error {
		ctx = c.db.MasterIntoContext(ctx)

		return queries.DeleteSatellites(
			ctx,
			satelliteIDs,
		)
	}
}

func (c *Container) PortInsertOrbit() port.InsertOrbit {
	return func(ctx context.Context, orbit domain.Orbit) error {
		ctx = c.db.MasterIntoContext(ctx)

		return queries.InsertOrbits(ctx, []domain.Orbit{orbit})
	}
}

func (c *Container) PortGetOrbits() port.GetOrbits {
	return func(ctx context.Context) ([]domain.Orbit, error) {
		ctx = c.db.MasterIntoContext(ctx)

		return queries.GetOrbits(ctx)
	}
}

func (c *Container) applyMigrations(ctx context.Context) error {
	ctx = c.db.MasterIntoContext(ctx)

	return queries.ApplyMigrations(ctx)
}

func (c *Container) IsConnectedToDB() bool {
	return c.db != nil
}
