package port

import (
	"context"

	"github.com/senago/sputnik/internal/domain"
)

type GetSatellites func(ctx context.Context) ([]domain.Satellite, error)

type GetSatellitesByNameLike func(ctx context.Context, nameLike string) ([]domain.Satellite, error)

type UpdateSatellite func(ctx context.Context, satellite domain.Satellite) error

type InsertSatellite func(ctx context.Context, satellite domain.Satellite) error

type DeleteSatellites func(ctx context.Context, satelliteIDs []domain.SatelliteID) error
