package ioc

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/senago/sputnik/internal/details/queries"
	"github.com/senago/sputnik/internal/domain"
	"github.com/senago/sputnik/internal/port"
)

func (c *Container) PortInsertSatellite() port.InsertSatellite {
	return func(ctx context.Context, satellite domain.Satellite) error {
		ctx = c.db.MasterIntoContext(ctx)

		// Have you heard of transactions?

		if err := queries.InsertSatellites(ctx, []domain.Satellite{satellite}); err != nil {
			return errors.Wrap(err, "queries.InsertSatellites")
		}

		if err := queries.InsertOrUpdateSatellitePositions(ctx, []domain.Satellite{satellite}); err != nil {
			return errors.Wrap(err, "queries.InsertOrUpdateSatellitePositions")
		}

		return nil
	}
}

func (c *Container) PortUpdateSatellite() port.UpdateSatellites {
	return func(ctx context.Context, satellites []domain.Satellite) error {
		ctx = c.db.MasterIntoContext(ctx)

		// Have you heard of transactions?

		if err := queries.UpdateSatellites(ctx, satellites); err != nil {
			return errors.Wrap(err, "queries.UpdateSatellites")
		}

		if err := queries.InsertOrUpdateSatellitePositions(ctx, satellites); err != nil {
			return errors.Wrap(err, "InsertOrUpdateSatellitePositions")
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
