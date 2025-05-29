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

# -------------------------------- Load tests -------------------------------- #

run-api:
	PG_DSN=${LOCAL_PG_DSN} go run ./cmd/api

run-k6:
	(cd tests/k6 && k6 run script.js)

# -------------------------------- Building -------------------------------- #

build-windows:
	GOOS=windows CGO_ENABLED=1 go build -ldflags="-s -w -H=windowsgui" ./cmd/sputnik

compress-windows:
	upx -5 sputnik.exe
