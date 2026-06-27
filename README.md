<div align="center">

<img src="https://s3.login.no/beehive/img/logo/logo-white-small.svg" alt="Login logo" width="80" height="80" />

<h1>Workerbee</h1>

<p>
  <img src="https://img.shields.io/badge/Go-fd8738?style=flat-square&logo=go&logoColor=white" alt="Go" />
  <img src="https://img.shields.io/badge/Gin-fd8738?style=flat-square&logo=go&logoColor=white" alt="Gin" />
  <img src="https://img.shields.io/badge/PostgreSQL-fd8738?style=flat-square&logo=postgresql&logoColor=white" alt="PostgreSQL" />
  <img src="https://img.shields.io/badge/S3-fd8738?style=flat-square&logo=amazons3&logoColor=white" alt="S3" />
  <img src="https://img.shields.io/badge/Varnish-fd8738?style=flat-square&logo=varnish&logoColor=white" alt="Varnish" />
  <img src="https://img.shields.io/badge/Docker-fd8738?style=flat-square&logo=docker&logoColor=white" alt="Docker" />
</p>

</div>

---

The main API for the Beehive and Queenbee applications, built for [Login](https://login.no).

## Features

- **Bearer token authentication** via Authentik, admin endpoints require the `QueenBee` group
- **Public and protected endpoints** under `/api/v2`
- **Swagger docs** available at `/api/v2/docs`
- **Object storage** for image and media uploads via S3-compatible API
- **Image processing** with WebP conversion
- **Rate limiting** on protected endpoints
- **Varnish cache** in front of the API

## Getting Started

1. **Configure environment**

   Create a `.env` file in the repo root. See [Configuration](#configuration) below or grab the values from 1Password.

2. **Start**

   ```bash
   docker compose up --build
   ```

   | Service | URL                               |
   |---------|-----------------------------------|
   | API     | http://localhost:8500             |
   | Docs    | http://localhost:8500/api/v2/docs |

   Port 8500 is the Varnish cache layer. The Go app listens on `PORT` (default `8080`) inside the container.

## Configuration

All variables go in the root `.env` file.

| Name                         | Default     | Notes                                               |
|------------------------------|-------------|-----------------------------------------------------|
| `HOST`                       | `0.0.0.0`   | API bind address                                    |
| `PORT`                       | `8080`      | Container port; must match the docker-compose mapping (`8500:8080`) |
| `DB`                         | `workerbee` | Postgres database name                              |
| `DB_HOST`                    | `localhost` | Postgres host                                       |
| `DB_PORT`                    | `5432`      | Postgres port                                       |
| `DB_USER`                    | `workerbee` | Postgres username                                   |
| `DB_PASSWORD`                |             | Postgres password                                   |
| `S3_URL`                     |             | S3-compatible storage endpoint                      |
| `S3_ACCESS_KEY_ID`           |             | Storage access key                                  |
| `S3_SECRET_ACCESS_KEY`       |             | Storage secret key                                  |
| `S3_REGION`                  | `us-east-1` | Storage region                                      |
| `ALLOWED_PROTECTED_REQUESTS` | `25`        | Max requests per minute on protected endpoints      |
| `LOAD_DUMMY_DATA`            | `false`     | Set to `true` to seed the database with test data   |

## Project Structure

- `api/` - Go application
- `api/handlers/` - HTTP handlers
- `api/services/` - business logic
- `api/repositories/` and `api/db/` - data access
- `api/routes_internal/` - route registration under `/api/v2`
- `api/docs/` - generated Swagger documentation
- `db/init.up.sql` - base database schema
- `db/dummydata.sql` - schema with seeded test data
