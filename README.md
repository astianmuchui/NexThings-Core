# NexThings Core

NexThings Core is the backend service for user authentication and device-registry foundations in the NexThings platform.

## Stack

- Go
- Fiber
- GORM
- SQLite (local) / PostgreSQL (production)
- Swagger UI (OpenAPI)

## Project Structure

- `cmd/main.go` application entrypoint
- `internal/routes` route registration
- `internal/handlers/api` API handlers
- `internal/models` database models
- `internal/schemas` request/response schema definitions
- `internal/templates` landing page templates
- `docs` generated OpenAPI spec files

## Run Locally

1. Install dependencies:

```bash
go mod tidy
```

2. Create a `.env` file in the project root:

```env
NEXTHINGS_PORT=8080
NEXTHINGS_ENVIRONMENT=Local
NEXTHINGS_POSTGRES_DSN=
NEXTHINGS_EMAIL_FROM=
NEXTHINGS_EMAIL_APP_PASSWORD=
```

3. Start the server:

```bash
go run ./cmd/main.go
```

Server starts on `http://localhost:8080` by default.

## API Documentation (Swagger UI)

- Swagger UI: `http://localhost:8080/docs/index.html`
- OpenAPI JSON: `http://localhost:8080/docs/doc.json`

If you change handler annotations, regenerate docs:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
$(go env GOPATH)/bin/swag init -g cmd/main.go -o docs
```

## Current Auth Endpoints

- `POST /api/v1/users/register`
- `POST /api/v1/users/login`
- `GET /api/v1/users/verify-account/:uid/:token`
- `PATCH /api/v1/users/reset-password/`

## Notes

- Local mode uses SQLite file `nexthings.sqlite`.
- Production mode uses PostgreSQL DSN from `NEXTHINGS_POSTGRES_DSN`.
