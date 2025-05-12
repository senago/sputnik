package main

import (
	"context"

	"github.com/senago/sputnik/internal/details/queries"
	"github.com/senago/sputnik/internal/domain"
	"github.com/senago/sputnik/internal/port"
)

func (a *App) PortInsertSatellite() port.InsertSatellite {
	return func(ctx context.Context, satellite domain.Satellite) error {
		ctx = a.db.MasterIntoContext(ctx)

		return queries.InsertSatellites(ctx, []domain.Satellite{satellite})
	}
}

func (a *App) PortUpdateSatellite() port.UpdateSatellite {
	return func(ctx context.Context, satellite domain.Satellite) error {
		ctx = a.db.MasterIntoContext(ctx)

		return queries.UpdateSatellite(ctx, satellite)
	}
}

func (a *App) PortGetSatellites() port.GetSatellites {
	return func(ctx context.Context) ([]domain.Satellite, error) {
		ctx = a.db.MasterIntoContext(ctx)

		return queries.GetSatellites(ctx, queries.GetSatellitesParams{})
	}
}

func (a *App) PortGetSatellitesByNameLike() port.GetSatellitesByNameLike {
	return func(ctx context.Context, nameLike string) ([]domain.Satellite, error) {
		ctx = a.db.MasterIntoContext(ctx)

		return queries.GetSatellites(
			ctx,
			queries.GetSatellitesParams{
				NameLike: nameLike,
			},
		)
	}
}

func (a *App) PortDeleteSatellites() port.DeleteSatellites {
	return func(ctx context.Context, satelliteIDs []domain.SatelliteID) error {
		ctx = a.db.MasterIntoContext(ctx)

		return queries.DeleteSatellites(
			ctx,
			satelliteIDs,
		)
	}
}

func (a *App) PortInsertOrbit() port.InsertOrbit {
	return func(ctx context.Context, orbit domain.Orbit) error {
		ctx = a.db.MasterIntoContext(ctx)

		return queries.InsertOrbits(ctx, []domain.Orbit{orbit})
	}
}

func (a *App) PortGetOrbits() port.GetOrbits {
	return func(ctx context.Context) ([]domain.Orbit, error) {
		ctx = a.db.MasterIntoContext(ctx)

		return queries.GetOrbits(ctx)
	}
}
