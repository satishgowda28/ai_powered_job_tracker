# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

AI-powered job application tracker with a Go backend (Fiber framework) and planned React + TypeScript frontend. Currently in active development with authentication implemented.

## Common Commands

All commands should be run from the `Backend/` directory:

```bash
# Development with live reload
air

# Manual build and run
go build -o ./tmp/main .
./tmp/main

# Run database migrations (goose)
goose -dir db/migration postgres "$DATABASE_URL" up
goose -dir db/migration postgres "$DATABASE_URL" down

# Regenerate SQL code after modifying queries
sqlc generate

# Run tests
go test ./...
```

## Environment Variables

Required in `Backend/.env`:
- `DATABASE_URL` - PostgreSQL connection string
- `JWT_SECRET` - Secret for JWT token signing
- `ISSUER` - JWT issuer (required, no default)
- `PORT` - Server port (default: 8080)
- `ENV` - Environment (default: development)

## Architecture

### Layered Architecture with Dependency Injection

```
cmd/server/main.go          → Entry point, wires everything together
internal/container/         → DI container that instantiates all dependencies
internal/handlers/          → HTTP handlers (receive requests, call services)
internal/services/          → Business logic layer
internal/respositories/     → Data access layer (uses sqlc-generated queries)
internal/routes/            → Route registration
internal/config/            → Singleton config loaded from environment
internal/database/          → Global pgxpool connection
internal/auth/              → JWT and password utilities (argon2id)
```

### Database Layer (sqlc)

- **Schema**: `db/migration/` - Goose migration files define the schema
- **Queries**: `db/queries/` - SQL files with sqlc annotations
- **Generated**: `db/generated/` - Type-safe Go code (regenerate with `sqlc generate`)

The repository layer uses the generated `Queries` interface from sqlc with `database.DB` pool.

### Adding a New Feature

1. Add migration in `db/migration/` (use goose naming: `YYYYMMDDHHMMSS_description.sql`)
2. Add queries in `db/queries/` with sqlc annotations
3. Run `sqlc generate`
4. Create repository in `internal/respositories/`
5. Create service in `internal/services/`
6. Create handler in `internal/handlers/`
7. Register routes in `internal/routes/`
8. Wire dependencies in `internal/container/container.go`

### Current API Endpoints

- `GET /health` - Health check
- `POST /auth/register` - User registration
- `POST /auth/login` - User login (returns JWT + refresh token)
