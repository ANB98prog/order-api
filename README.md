# Order API

Простое API для создания заказов

## Стек 

| Name     | Version |
|----------|---------|
| Go       | 1.25.0  |
| Postgres | 16.4    |
| Redis    | 7.2     |


СПО развернуто в Docker контейнерах

## Запуск приложения

1. Необходимо установить зависимости проекта
```shell
go mod tidy
```
либо сразу собрать приложение

```shell
go build ./order-api/cmd
```


2. Необходимо собрать и запустить контейнеры с Redis и Postgres
Выполняем команду из корня проекта:

```shell
docker compose -f ./order-api/infrastructure/docker-compose.yml up -d
```

3. Далее необходимо сделать миграции находясь в директории migrations

```shell
goose postgres "host=localhost user=postgres password=my_pass dbname=market port=5433 sslmode=disable" up
```

Для упрощения в проекте оставлен .env файл
