package port

import (
	"context"

	"github.com/senago/sputnik/internal/domain"
)

type InsertOrbit func(ctx context.Context, orbit domain.Orbit) error

type GetOrbits func(ctx context.Context) ([]domain.Orbit, error)
