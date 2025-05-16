package queries

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/samber/lo"
	"github.com/senago/sputnik/internal/domain"
)

type GetSatellitesParams struct {
	NameLike string
}

func GetSatellites(ctx context.Context, params GetSatellitesParams) ([]domain.Satellite, error) {
	satelliteModels, err := getSatellitesModels(ctx, params)
	if err != nil {
		return nil, errors.Wrap(err, "getSatellitesModels")
	}

	if len(satelliteModels) == 0 {
		return []domain.Satellite{}, nil
	}

	orbitIDs := lo.Uniq(lo.Map(
		satelliteModels,
		func(sm satelliteModel, _ int) domain.OrbitID {
			return sm.OrbitID
		},
	))

	orbits, err := GetOrbitsByID(ctx, orbitIDs)
	if err != nil {
		return nil, errors.Wrap(err, "GetOrbitsByID")
	}

	orbitsByID := lo.SliceToMap(
		orbits,
		func(orbit domain.Orbit) (domain.OrbitID, domain.Orbit) {
			return orbit.ID, orbit
		},
	)

	satelliteIDs := lo.Map(
		satelliteModels,
		func(sm satelliteModel, _ int) domain.SatelliteID {
			return sm.ID
		},
	)

	positions, err := getSatellitePositions(ctx, satelliteIDs)
	if err != nil {
		return nil, errors.Wrap(err, "getSatellitePositions")
	}

	positionsByID := lo.SliceToMap(
		positions,
		func(spm sattelitePositionModel) (domain.SatelliteID, domain.Position) {
			return spm.ID, spm.Position
		},
	)

	satellites := lo.Map(
		satelliteModels,
		func(sm satelliteModel, _ int) domain.Satellite {
			return domain.Satellite{
				ID:          sm.ID,
				Orbit:       orbitsByID[sm.OrbitID],
				Position:    positionsByID[sm.ID],
				Name:        sm.Name,
				Description: sm.Description,
				Type:        sm.Type,
			}
		},
	)

	return satellites, nil
}
