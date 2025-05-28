package queries

import (
	"context"
	"fmt"
)

const queryMigration001 = `-- Migration001
create table if not exists satellite (
    id uuid not null,
    orbit_id uuid not null,
    name text not null,
    description text not null,
    type text not null,
    
    primary key (id)
);

create unique index if not exists satellite_name_idx on satellite (name);

create table if not exists orbit (
    id uuid not null,
    name text not null,
    height_km bigint not null,

    primary key (id)
);

create unique index if not exists orbit_name_idx on orbit (name);

create table if not exists satellite_position (
    id uuid not null,
    x real not null,
    y real not null,

    primary key (id)
) with (fillfactor=90);
`

func ApplyMigrations(ctx context.Context) error {
	_, err := resolveConn(ctx).Exec(
		ctx,
		queryMigration001,
	)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	return nil
}
