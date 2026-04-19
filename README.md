# Workerbee API

Workerbee is the main API for the Beehive and Queenbee applications. It is built with Go and the Gin framework, uses PostgreSQL as its primary database, and serves both public and protected endpoints. API documentation is available at `/api/v2/docs` once the application is running.

The API also integrates with object storage for file handling, which is used for images and other media uploads.

## Beehive Database

This repository includes the database setup used by Login for Workerbee. The `db` directory contains:

- `init.up.sql` for the base database structure
- `dummydata.sql` for the same structure with seeded test data

## Running the project

Start the full stack with:

```sh
docker compose up --build
```

This starts both the PostgreSQL database and the Workerbee API.

For local development, use the seeded SQL file if you want test data. Otherwise use the base schema.

## Environment

Workerbee reads configuration from a `.env` file in the project root. Start by getting a working `.env` file from 1Password.

The main settings cover:

- database connection
- application host and port
- object storage credentials
- protected endpoint rate limiting

## Project structure

- `api/` contains the Go application
- `api/handlers/` contains the HTTP handlers
- `api/services/` contains business logic
- `api/repositories/` and `api/db/` contain data access code
- `api/routes_internal/` registers routes under `/api/v2`
- `api/docs/` contains the generated API documentation
- `db/` contains database initialization files
- `docker-compose.yml` starts the API and database together
- `Dockerfile` builds the application container

## Getting started

1. Get a valid `.env` file.
2. Run `docker compose up --build`.
3. Open `/api/v2/docs` to explore the API.
4. Check `db/` if you need seeded or empty local database initialization.
