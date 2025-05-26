package queries

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/senago/sputnik/internal/domain"
)

type sattelitePositionModel struct {
	ID       domain.SatelliteID
	Position domain.Position
}

const queryInsertOrUpdateSatellitePositions = `-- InsertOrUpdateSatellitePositions
insert into satellite_position (
	id,
	x,
	y
)
select * from unnest (
	$1::uuid[],
	$2::real[],
	$3::real[]
)
on conflict (id) do update
set
	x = excluded.x,
	y = excluded.y
`

func InsertOrUpdateSatellitePositions(ctx context.Context, satellites []domain.Satellite) error {
	args := []any{
		nest(satellites, func(s domain.Satellite) string { return s.ID.String() }),
		nest(satellites, func(s domain.Satellite) float32 { return s.Position.X }),
		nest(satellites, func(s domain.Satellite) float32 { return s.Position.Y }),
	}

	_, err := resolveConn(ctx).Exec(
		ctx,
		queryInsertOrUpdateSatellitePositions,
		args...,
	)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	return nil
}

const queryGetSatellitePositions = `-- GetSatellitePositions
select
	id,
	x,
	y
from satellite_position
where id = any($1)
`

func getSatellitePositions(ctx context.Context, ids []domain.SatelliteID) ([]sattelitePositionModel, error) {
	satelliteRows, err := resolveConn(ctx).Query(
		ctx,
		queryGetSatellitePositions,
		nest(ids, func(sid domain.SatelliteID) string { return sid.String() }),
	)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	return scanRows(
		satelliteRows,
		func(row pgx.Row, item *sattelitePositionModel) error {
			return row.Scan(
				&item.ID,
				&item.Position.X,
				&item.Position.Y,
			)
		},
	)
}
