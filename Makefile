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

build-windows:
	make build-run && make build-extract && make build-compress

build-run:
	docker-buildx build -t sputnik-windows .
	docker run --rm -v ${PWD}:/build sputnik-windows

build-extract:
	docker create --name tmp-sputnik-windows sputnik-windows
	docker cp tmp-sputnik-windows:/app/sputnik.exe ./build/sputnik.exe
	docker rm tmp-sputnik-windows

build-compress:
	upx -5 ./build/sputnik.exe
