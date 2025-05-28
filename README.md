# Sputnik

Домашнее задание по курсу СУБД

## Usage
### Запуск без docker
Зависимости:
- postgres@13
- golang@1.24

После этого запуск:
```sh
go run ./cmd/sputnik --pg_dsn=<dsn>
```

### Запуск с docker
```sh
make docker-env-up
make run
```

## Load tests
```sh
make run-api
```

```sh
make rrun-k6
```
