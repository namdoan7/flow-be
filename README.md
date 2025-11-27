# go-be

Simple example Go REST API using Gin with a clean-ish project layout.

## Getting Started

```bash
make run
```

The server listens on `:8080` by default. Override with `PORT`.

## API

| Method | Path            | Description      |
|--------|-----------------|------------------|
| POST   | /api/v1/users   | Create a user    |
| GET    | /api/v1/users   | List all users   |
| GET    | /api/v1/users/:id | Fetch a user  |
| PUT    | /api/v1/users/:id | Update a user |
| DELETE | /api/v1/users/:id | Delete a user |

Sample payload:

```json
{
  "name": "Ada Lovelace",
  "email": "ada@example.com"
}
```

## Docker

```bash
docker build -t go-be .
docker run -p 8080:8080 go-be
```

## Project Layout

```
cmd/server        # application entrypoint
internal/model    # domain models
internal/repository  # data access layer (in-memory for example)
internal/service  # business logic
internal/handler  # HTTP handlers
internal/router   # gin router wiring
pkg/response      # reusable response helpers
```

