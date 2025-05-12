LOCAL_BIN:=$(CURDIR)/bin
MIGRATIONS_DIR:=migrations

export GOBIN:=${LOCAL_BIN}

# -------------------------------- Local -------------------------------- #

include env/local.env
export

local:
	make local-env-up && trap "make local-env-stop" SIGINT && make local-run

local-run:
	PG_DSN=${LOCAL_PG_DSN} go run ./cmd/sputnik

# -------------------------------- Local Env -------------------------------- #

local-env-up:
	make local-db-up

local-env-stop:
	make local-db-stop

local-env-down:
	make local-db-down

# -------------------------------- Local DB -------------------------------- #

local-db-up:
	docker-compose -f ./env/docker-compose.local.yaml up -d --wait && \
	sleep 2 && \
	${GOBIN}/goose -dir ${MIGRATIONS_DIR} postgres "${LOCAL_PG_DSN}" up

local-db-stop:
	docker-compose -f ./env/docker-compose.local.yaml stop

local-db-down:
	docker-compose -f ./env/docker-compose.local.yaml down --remove-orphans

local-psql:
	psql ${LOCAL_PG_DSN}

# -------------------------------- Migrations -------------------------------- #

install-goose:
	go install github.com/pressly/goose/v3/cmd/goose@latest

migration-create: install-goose
	@[ "${name}" ] || (echo 'usage: make name="<migration name>" migration-create'; exit 1)
	${GOBIN}/goose -dir ${MIGRATIONS_DIR} create ${name} sql
