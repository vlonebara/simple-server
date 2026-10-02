# simple-server

Универсальный минимальный шаблон HTTP-сервера на Go без внешних зависимостей. Подходит как стартовая точка для API и сервисов.

## Что включено

- `net/http` и новый маршрутизатор стандартной библиотеки Go;
- JSON-ответы;
- middleware логирования запросов и recovery после panic;
- конфигурация из переменных окружения;
- корректное завершение по `SIGINT` и `SIGTERM`;
- разделение HTTP-слоя, бизнес-логики и инфраструктурных интерфейсов.

## Быстрый старт

Требуется Go 1.24+.

```bash
go run ./cmd/server
```

Сервер будет доступен на `http://localhost:8080`.

```bash
curl http://localhost:8080/healthz
curl 'http://localhost:8080/api/v1/hello?name=Кирилл'
```

## Конфигурация

| Переменная | Значение по умолчанию |
|---|---|
| `HTTP_ADDR` | `:8080` |
| `HTTP_READ_TIMEOUT` | `5s` |
| `HTTP_WRITE_TIMEOUT` | `10s` |
| `HTTP_IDLE_TIMEOUT` | `60s` |
| `SHUTDOWN_TIMEOUT` | `10s` |
| `LOG_LEVEL` | `info` |

## Структура

```text
cmd/server/          точка входа
internal/config/     конфигурация
internal/httpapi/    маршруты, handlers и middleware
internal/service/    бизнес-логика и интерфейсы зависимостей
```

Развивайте проект по цепочке: `HTTP handler → service → repository interface → concrete adapter`. Благодаря этому HTTP, бизнес-логика и конкретная база данных остаются независимыми друг от друга.
