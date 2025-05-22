# Sputnik

Домашнее задание по курсу СУБД

## Usage
### Запуск без docker
Зависимости:
- postgres@13
- golang@1.24

В [env/local.env](env/local.env) указать:
- LOCAL_PG_PORT: порт, на котором развернут постгрес
- LOCAL_PG_DSN: dsn до инстанса постгреса

На постгресе необходимо применить миграции:
```sh
make migration-apply
```

После этого запуск:
```sh
make run
```

### Запуск с docker
```sh
make docker-env-up
make run
```
