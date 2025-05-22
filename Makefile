LOCAL_BIN:=$(CURDIR)/bin
MIGRATIONS_DIR:=migrations

export GOBIN:=${LOCAL_BIN}

# -------------------------------- Local -------------------------------- #

include env/local.env
export

run:
	PG_DSN=${LOCAL_PG_DSN} go run ./cmd/sputnik

# -------------------------------- Docker Env -------------------------------- #

docker-env-up:
	docker-compose -f ./env/docker-compose.local.yaml up -d --wait && \
	sleep 2 && \
	make migration-apply

docker-env-stop:
	docker-compose -f ./env/docker-compose.local.yaml stop

docker-env-down:
	docker-compose -f ./env/docker-compose.local.yaml down --remove-orphans

docker-psql:
	psql ${LOCAL_PG_DSN}

# -------------------------------- Migrations -------------------------------- #

install-goose:
	go install github.com/pressly/goose/v3/cmd/goose@latest

migration-create: install-goose
	@[ "${name}" ] || (echo 'usage: make name="<migration name>" migration-create'; exit 1)
	${GOBIN}/goose -dir ${MIGRATIONS_DIR} create ${name} sql

migration-apply:
	${GOBIN}/goose -dir ${MIGRATIONS_DIR} postgres "${LOCAL_PG_DSN}" up
