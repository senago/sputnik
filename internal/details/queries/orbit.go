package queries

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/senago/sputnik/internal/domain"
)

const queryInsertOrbits = `-- InsertOrbits
insert into orbit (
	id,
	name,
	height_km
)
select * from unnest (
	$1::uuid[],
	$2::text[],
	$3::bigint[]
)
`

func InsertOrbits(ctx context.Context, orbits []domain.Orbit) error {
	args := []any{
		nest(orbits, func(o domain.Orbit) string { return o.ID.String() }),
		nest(orbits, func(o domain.Orbit) string { return o.Name }),
		nest(orbits, func(o domain.Orbit) int64 { return o.HeightKm }),
	}

	_, err := resolveConn(ctx).Exec(
		ctx,
		queryInsertOrbits,
		args...,
	)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	return nil
}

const queryGetOrbits = `-- GetOrbits
select
	id,
	name,
	height_km
from orbit
`

func GetOrbits(ctx context.Context) ([]domain.Orbit, error) {
	rows, err := resolveConn(ctx).Query(
		ctx,
		queryGetOrbits,
	)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	return scanRows(
		rows,
		func(row pgx.Row, item *domain.Orbit) error {
			return row.Scan(
				&item.ID,
				&item.Name,
				&item.HeightKm,
			)
		},
	)
}

const queryGetOrbitsByID = `-- GetOrbitsByID
select
	id,
	name,
	height_km
from orbit
where id = any($1)
`

func GetOrbitsByID(ctx context.Context, orbitIDs []domain.OrbitID) ([]domain.Orbit, error) {
	rows, err := resolveConn(ctx).Query(
		ctx,
		queryGetOrbitsByID,
		nest(orbitIDs, func(oid domain.OrbitID) string { return oid.String() }),
	)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	return scanRows(
		rows,
		func(row pgx.Row, item *domain.Orbit) error {
			return row.Scan(
				&item.ID,
				&item.Name,
				&item.HeightKm,
			)
		},
	)
}
