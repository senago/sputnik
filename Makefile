# -------------------------------- Local -------------------------------- #

include env/local.env
export

run:
	go run ./cmd/sputnik --pg_dsn=${LOCAL_PG_DSN}

# -------------------------------- Docker Env -------------------------------- #

docker-env-up:
	docker-compose -f ./env/docker-compose.local.yaml up -d --wait

docker-env-stop:
	docker-compose -f ./env/docker-compose.local.yaml stop

docker-env-down:
	docker-compose -f ./env/docker-compose.local.yaml down --remove-orphans

docker-psql:
	psql ${LOCAL_PG_DSN}

docker-pg-dump:
	pg_dump --data-only --column-inserts ${LOCAL_PG_DSN} > dump/dump.sql

# -------------------------------- Load tests -------------------------------- #

run-api:
	PG_DSN=${LOCAL_PG_DSN} go run ./cmd/api

run-k6:
	(cd tests/k6 && k6 run script.js)

# -------------------------------- Building -------------------------------- #

compress-windows:
	upx -5 ./bin/sputnik.exe

build-windows:
	docker build -t sputnik-windows .
	docker create --name tmp-sputnik-windows sputnik-windows
	docker cp tmp-sputnik-windows:/app/sputnik.exe ./bin/sputnik.exe
	docker rm tmp-sputnik-windows
