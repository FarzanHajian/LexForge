# LexPilot

LexPilot is a multi-user foreign-language vocabulary learning application. It combines vocabulary management, topic-based organization, daily spaced-practice sessions, learning-history tracking, and read-only topic sharing between users.

This is also a portfolio project: code quality, testability, and observability are treated as first-class goals, alongside a deliberate effort to avoid over-engineering features ahead of need.

## Status

Early scaffolding stage. Backend and frontend project skeletons exist; core domain features (topics, words, daily practice, spaced repetition, sharing) are not yet implemented. See [`docs/LexPilot_Coding_Agent_Handoff.md`](docs/LexPilot_Coding_Agent_Handoff.md) for the full set of architectural and product decisions driving the build-out.

## Tech stack

**Backend**
- Go
- [Echo](https://echo.labstack.com/) — HTTP/API framework
- [GORM](https://gorm.io/) — ORM
- MySQL — production database
- Auth0 (official Go SDK) — authentication/identity

**Frontend**
- React + TypeScript
- Vite
- React Router
- Tailwind CSS

The frontend builds to a static SPA. In production, a single Go binary serves both the REST API (`/api/*`) and the built React app (`/*`, with SPA fallback to `index.html`) — there is no separate Node.js server, SSR, or dedicated frontend host at runtime. Node/npm are development and build-time tools only.

## Repository layout

```
LexForge/
├── backend/            Go module (Echo, GORM, MySQL, Auth0 adapter)
│   ├── cmd/server/      application entrypoint
│   └── internal/
│       ├── domain/      core business entities and rules
│       ├── service/      application services orchestrating domain logic
│       ├── repository/   repository interfaces + GORM/MySQL implementations
│       ├── http/         Echo handlers, routing, middleware
│       ├── auth/         auth abstraction + Auth0 adapter
│       ├── scheduler/     spaced-repetition scheduling abstraction
│       └── config/       configuration / environment / system-wide settings
├── frontend/           React + TypeScript + Vite + Tailwind SPA
└── docs/               specs, plans, ADRs, and the handoff document
```

## Architecture

Business logic is isolated from infrastructure via interfaces, so Echo, GORM, MySQL, and Auth0 can be swapped or unit-tested around without touching domain code:

```
HTTP Handler -> Application Service -> Repository/Infra Interface -> GORM Repository -> MySQL
Application Core -> Authentication Interface -> Auth0 Adapter -> Auth0 Go SDK
```

The spaced-repetition scheduler follows the same pattern: it starts as a simple due/overdue-first algorithm and is expected to evolve toward SM-2, and potentially FSRS, without changes to the rest of the application.

## Development

Backend and frontend commands will be documented here as the projects are scaffolded out (build, lint, test, run). For now:

```bash
# Backend
cd backend
go build ./...

# Frontend
cd frontend
npm install
npm run dev
```

## Environment variables

### Backend (`backend/`)

| Variable | Required | Default | Description |
|---|---|---|---|
| `LEXFORGE_SERVER_ADDRESS` | No | `http://localhost:5656` | Address the Echo server binds to |
| `LEXFORGE_MAX_SESSIONS_PER_DAY` | No | `1` | Max practice sessions per user/notebook/day |
| `LEXFORGE_MIGRATE_DATABASE` | No | `false` | Whether to run GORM auto-migration on startup |
| `LEXFORGE_DATABASE_DSN` | Yes | — | MySQL connection string (DSN) |
| `LEXFORGE_AUTH0_DOMAIN` | Yes | — | Auth0 tenant domain; used as the OIDC issuer when validating access tokens |
| `LEXFORGE_AUTH0_AUDIENCE` | Yes | — | Auth0 API identifier this backend accepts tokens for |

### Frontend (`frontend/`)

| Variable | Required | Default | Description |
|---|---|---|---|
| `VITE_LEXFORGE_SERVER_ADDRESS` | No | `''` (relative paths) | Base URL of the backend API the SPA calls |
| `VITE_AUTH0_DOMAIN` | Yes | — | Auth0 tenant domain (same tenant as the backend's `LEXFORGE_AUTH0_DOMAIN`) |
| `VITE_AUTH0_CLIENT_ID` | Yes | — | Auth0 SPA application client ID |
| `VITE_AUTH0_AUDIENCE` | Yes | — | Must match the backend's `LEXFORGE_AUTH0_AUDIENCE` — identifies the API the requested access token is issued for |

This repo has no `.env` files (frontend `.env*` is gitignored). Locally, backend variables are set in the "Debug BE" configuration in `.vscode/launch.json`, and frontend variables in the "Start FE dev server" task in `.vscode/tasks.json`.

## Testing

- **Unit tests** (Go `testing` + Testify): colocated as `*_test.go` files next to the code they test in each `internal/` package. Business rules are tested against repository/service interfaces with mocks/fakes (see `internal/testutil`) — no database involved.
- **Integration tests** (Testcontainers + real MySQL): live under `backend/test/integration`, gated behind the `integration` build tag so a plain `go test ./...` never needs Docker. Repository behavior, migrations, and DB-level constraints (e.g. one daily session per user/topic/day) are tested against an actual MySQL instance rather than SQLite.

```bash
cd backend

# unit tests only (no Docker required)
go test ./...

# integration tests (requires Docker running locally)
go test -tags=integration ./test/integration/...
```

## Observability

Planned stack: OpenTelemetry → OpenTelemetry Collector → Prometheus (metrics) / Tempo (traces) / Loki (logs) → Grafana, introduced incrementally alongside features rather than upfront.

## Contributing / guidance for coding agents

See [`CLAUDE.md`](CLAUDE.md) for architectural constraints, domain rules, and implementation principles to follow when working in this repository.
