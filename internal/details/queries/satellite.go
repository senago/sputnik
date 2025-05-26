package queries

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/senago/sputnik/internal/domain"
)

type satelliteModel struct {
	ID          domain.SatelliteID
	OrbitID     domain.OrbitID
	Name        string
	Description string
	Type        string
}

const queryInsertSatellites = `-- InsertSatellites
insert into satellite (
	id,
	orbit_id,
	name,
	description,
	type
)
select * from unnest (
	$1::uuid[],
	$2::uuid[],
	$3::text[],
	$4::text[],
	$5::text[]
)
`

func InsertSatellites(ctx context.Context, satellites []domain.Satellite) error {
	args := []any{
		nest(satellites, func(s domain.Satellite) string { return s.ID.String() }),
		nest(satellites, func(s domain.Satellite) string { return s.Orbit.ID.String() }),
		nest(satellites, func(s domain.Satellite) string { return s.Name }),
		nest(satellites, func(s domain.Satellite) string { return s.Description }),
		nest(satellites, func(s domain.Satellite) string { return s.Type }),
	}

	_, err := resolveConn(ctx).Exec(
		ctx,
		queryInsertSatellites,
		args...,
	)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	return nil
}

const queryUpdateSatellites = `-- UpdateSatellites
update satellite s
set
	orbit_id = bulk.orbit_id,
	name = bulk.name,
	description = bulk.description,
	type = bulk.type
from (
	select * from unnest (
		$1::uuid[],
		$2::uuid[],
		$3::text[],
		$4::text[],
		$5::text[]
	) as t(id, orbit_id, name, description, type)
) as bulk
where s.id = bulk.id
`

func UpdateSatellites(ctx context.Context, satellites []domain.Satellite) error {
	args := []any{
		nest(satellites, func(s domain.Satellite) string { return s.ID.String() }),
		nest(satellites, func(s domain.Satellite) string { return s.Orbit.ID.String() }),
		nest(satellites, func(s domain.Satellite) string { return s.Name }),
		nest(satellites, func(s domain.Satellite) string { return s.Description }),
		nest(satellites, func(s domain.Satellite) string { return s.Type }),
	}

	_, err := resolveConn(ctx).Exec(
		ctx,
		queryUpdateSatellites,
		args...,
	)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	return nil
}

const queryDeleteSatellites = `-- DeleteSatellites
delete from satellite
where id = any($1)
`

func DeleteSatellites(ctx context.Context, satelliteIDs []domain.SatelliteID) error {
	_, err := resolveConn(ctx).Exec(
		ctx,
		queryDeleteSatellites,
		nest(satelliteIDs, func(sid domain.SatelliteID) string { return sid.String() }),
	)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	return nil
}

const queryGetSatellitesModels = `-- GetSatellitesModels
select
	id,
	orbit_id,
	name,
	description,
	type
from satellite
`

const queryGetSatellitesModelsNameLike = `-- GetSatellitesModelsNameLike
select
	id,
	orbit_id,
	name,
	description,
	type
from satellite
where name ilike $1
`

func getSatellitesModels(ctx context.Context, params GetSatellitesParams) ([]satelliteModel, error) {
	query := queryGetSatellitesModels
	if params.NameLike != "" {
		query = queryGetSatellitesModelsNameLike
	}

	var args []any
	if params.NameLike != "" {
		args = append(args, "%"+params.NameLike+"%")
	}

	satelliteRows, err := resolveConn(ctx).Query(
		ctx,
		query,
		args...,
	)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	return scanRows(
		satelliteRows,
		func(row pgx.Row, item *satelliteModel) error {
			return row.Scan(
				&item.ID,
				&item.OrbitID,
				&item.Name,
				&item.Description,
				&item.Type,
			)
		},
	)
}
