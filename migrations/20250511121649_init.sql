-- +goose Up
-- +goose StatementBegin
create table satellite (
    id uuid not null,
    orbit_id uuid not null,
    name text not null,
    description text not null,
    type text not null,
    
    primary key (id)
);

create unique index satellite_name_idx on satellite (name);

create table orbit (
    id uuid not null,
    name text not null,
    height_km bigint not null,

    primary key (id)
);

create unique index orbit_name_idx on orbit (name);

create table satellite_position (
    id uuid not null,
    x real not null,
    y real not null,

    primary key (id)
) with (fillfactor=90);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table orbit;
drop table satellite;
-- +goose StatementEnd
