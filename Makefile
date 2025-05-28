# -------------------------------- Local -------------------------------- #

include env/local.env
export

run:
	PG_DSN=${LOCAL_PG_DSN} go run ./cmd/sputnik

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
