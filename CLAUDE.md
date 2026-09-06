# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project status

Backend scaffolding has started: `backend/go.mod` (module `github.com/FarzanHajian/lexforge/backend`) exists, along with early domain types under `backend/internal/domain` (`User`, `Notebook`). There is still no build/test tooling, migrations, HTTP layer, or GORM repository implementation beyond these domain types. Before writing code, check whether more has changed — if `/docs/specs`, `/docs/plans`, `/docs/adr`, `/issues`, `CHANGELOG.md`, or a real `README.md` exist, treat them as authoritative over this file and over both handoff docs, and flag any conflicts rather than silently overriding them.

The authoritative handoff doc is `docs/LexForge_Coding_Agent_Handoff.md`. An earlier `docs/LexPilot_Coding_Agent_Handoff.md` still exists but is superseded where the two disagree (e.g. product name, `Word` → `StudyItem`) — treat the LexForge doc as newer/authoritative per its own conflict-resolution rule.

Decisions recorded in this file (and in the handoff docs) are a starting point, not a binding contract. The user may deviate from them as real implementation choices get made along the way — defer to those in-the-moment decisions rather than citing this file or a handoff doc to push back, and update this file to match reality when a recorded decision changes.

Once build/test tooling exists, update this file with the actual commands (Go build/test/lint, frontend dev/build/lint commands, how to run a single test, etc.) — do not invent them before they exist.

## What LexForge is

A multi-user foreign-language vocabulary learning app (formerly referred to as "LexPilot" in older docs — the product/repo name is now **LexForge**): vocabulary management, notebook-based organization, daily spaced-practice sessions, learning-history tracking, and read-only notebook sharing between users. It also doubles as a portfolio project, so code quality, testability, and observability are deliberate goals — but avoid building sophistication (scheduling algorithms, analytics, infra) ahead of actual need.

## Technology decisions (fixed, do not silently change)

- **Backend:** Go, Echo (HTTP), GORM (ORM), MySQL (production DB), JSON REST API. Entity IDs are UUIDs (see Core domain model).
- **Auth:** Auth0 via the official Auth0 Go SDK. Auth0 types must never leak into application/domain code — isolate behind an application-level auth interface with an Auth0 adapter implementation. Internal users have an `auth0_sub`-equivalent field linking them to Auth0 (in code, `User.ExternalId` — named generically rather than `Auth0Sub` to keep the domain type provider-agnostic).
- **Frontend:** React + TypeScript + Vite + React Router + Tailwind CSS, built as a static SPA. No Next.js, no SSR, no server-generated HTML/HTMX/Alpine, no Node.js runtime in production.
- **Deployment shape:** one Go binary serves both `/api/*` (Echo) and `/*` (static SPA build, with fallback to `index.html` for client-side routes). Node/npm are dev/build-time only. A future iteration may `go:embed` the Vite `dist` output into the Go binary.

## Architecture / dependency direction

Layered, interface-isolated, no DI framework:

```
HTTP Handler -> Application Service -> Repository/Infra Interface -> GORM Repository -> MySQL
Application Core -> Authentication Interface -> Auth0 Adapter -> Auth0 Go SDK
```

Business logic must stay independent of Echo, GORM, MySQL, and Auth0 where practical — it should depend only on small, consumer-defined interfaces (Go interfaces are implicit, so prefer narrow interfaces over large "repository" ones). This is about behavior, not struct tags: GORM tags on application structs are an **intentional, accepted pragmatic choice** for this project (e.g. `domain.User`, `domain.Notebook` carry `gorm:"primaryKey"`, `type:char(36)`, `uniqueIndex`, etc. directly). Do not create separate domain/persistence models solely to eliminate GORM tags — tags are inert reflected metadata with no Go import on `gorm`, and a mapping layer between domain and persistence models isn't worth the boilerplate at this project's scale. Revisit only if a persistence detail genuinely diverges from the domain shape.

The spaced-repetition scheduler is likewise a replaceable component behind an interface: start with a simple due/overdue-first scheduler, evolve toward SM-2-style, potentially FSRS later — never hard-code a single scheduling implementation into callers.

## Core domain model

- **IDs**: primary keys are UUIDs, not auto-increment integers — Go field type `string`, generated via `github.com/google/uuid` (`uuid.NewString()`) at construction time, mapped to `CHAR(36)` in MySQL (`gorm:"type:char(36);primaryKey"`). Applies across all domain entities.
- **User**: internal record — currently `Id`, `Name`, `ExternalId` in code (`backend/internal/domain/user.go`). `email`/`created_at` from the original handoff design aren't in the struct yet. The app owns its own authorization rules independent of Auth0.
- **Notebook** (formerly `Topic` — renamed to better match the product's "physical notebook" metaphor): owned by a user, has exactly one template, can be shared read-only with other users via `OWNER`/`VIEWER` roles. Viewers can read a notebook and its study items but never modify them. In code (`backend/internal/domain/notebook.go`), `Template` is a plain raw string (Markdown-formatted), capped at `varchar(2048)`, stored directly on `Notebook` — **not** a structured/JSON schema and **not** a separate `NotebookTemplate` entity/table (don't introduce one unless templates need an independent lifecycle later).
- **StudyItem** (not `Word` — renamed per `docs/LexForge_Coding_Agent_Handoff.md`): belongs to exactly one notebook. Can be a word, phrase, expression, idiom, sentence, or other language-learning item — content is **template-driven**, not a fixed schema; never hard-code the vocabulary model around fields like `translation`/`example`/`notes`. Supports Markdown/rich text.
- **Learning progress/state**: keyed by `(user, study_item)`, not by study item alone — a shared study item has independent learning state per user. This is the central multi-tenancy rule of the domain: never let per-item data leak across users. Keep current aggregate state (`UserStudyItemState`: e.g. `user_id`, `study_item_id`, `status`, `last_reviewed_at`, `next_review_at`, `correct_count`, `wrong_count`) separate from historical review events (`StudyItemReview`).
- **PracticeSession**: one session per `(user_id, notebook_id, date)`, enforced at the DB level with a `UNIQUE(user_id, notebook_id, date)` constraint. The user picks the session size (5/10/15/20) at creation time; that size and the session's item set/order are a **snapshot** fixed at creation — never dynamically recomputed as the user progresses (`PracticeSessionItem` rows capture that snapshot). If today's session already exists, resume it instead of creating another.
  - Statuses: `IN_PROGRESS`, `COMPLETED`, `EXPIRED`.
  - **Expiration is lazy** — do not run a midnight background job. When the user attempts to start a new session, the backend checks for an existing `IN_PROGRESS` session from a previous calendar day and transitions it to `EXPIRED` before creating the new day's session. The backend must also reject answers submitted against an expired/previous-day session.
  - Session statistics only include `COMPLETED` sessions (`IN_PROGRESS`/`EXPIRED` excluded), but reviews actually submitted before a session was abandoned remain valid learning-history records — i.e. session stats come from completed sessions, learning stats come from actual recorded reviews.
- `max_sessions_per_day` is a system-wide config value, not a per-user preference — keep it configurable rather than a scattered magic number.
- **Review history (`StudyItemReview`)**: every review is preserved (never overwritten), attributed to `user_id`, `study_item_id`, `session_id`, `result`, `reviewed_at` (optional response time later). Initial UX uses a binary `KNOWN`/`UNKNOWN` result; richer `Again/Hard/Good/Easy` grading is a later evolution once a real scheduler needs it.

## Testing strategy

- Go `testing` + Testify (assertions) + Testify/mock (mocking).
- **Unit tests**: colocated `*_test.go` files alongside the code in each `internal/` package (not a separate folder) — pure business logic against repository/service interfaces with mocks/fakes, no database. Shared fakes/mocks/fixtures live in `internal/testutil`. Priority areas: practice-session rules, notebook/viewer permissions, study-item selection, learning-state calculations, scheduler behavior.
- **Integration tests**: live in `backend/test/integration`, using real MySQL via Testcontainers (`github.com/testcontainers/testcontainers-go/modules/mysql`) — not SQLite — because SQL semantics, constraints, and concurrency matter. Important cases: practice-session uniqueness, session expiration, migrations, notebook sharing, learning-state persistence, repository queries. Gated behind the `integration` build tag (`//go:build integration`) so plain `go test ./...` never requires Docker; run them explicitly with `go test -tags=integration ./test/integration/...`.

## Observability (portfolio-relevant, introduce incrementally)

Planned stack: OpenTelemetry -> OTel Collector -> Prometheus (metrics) / Tempo (traces) / Loki (logs) -> Grafana. Trace meaningful workflows (e.g. start-practice-session: auth -> authz -> find due study items -> select items -> create session -> DB ops). Useful metrics: HTTP latency/error rate, DB timing, sessions completed, study items reviewed, known/unknown results, practice duration. Use structured logging, never `Println`-style logs, and never log secrets/tokens/passwords/unnecessary personal data. Don't add Kubernetes, service meshes, or infra the app doesn't need yet.

## Implementation principles

1. Prefer simplicity; don't add technology because it's interesting.
2. Keep business logic independent of Echo/GORM/MySQL/Auth0 where practical.
3. Use interfaces at real boundaries (repositories, auth, scheduler) — not everywhere.
4. Business rules must be testable without a database.
5. Use real MySQL (not SQLite) for integration tests.
6. Preserve full review/learning history; never overwrite it.
7. Learning state is always per-user, even for shared notebooks/study items.
8. A practice session is an immutable snapshot once started.
9. Expire unfinished sessions lazily (on next session-start attempt); never a midnight job.
10. Exclude expired sessions from session statistics; keep their actually-submitted reviews in learning history.
11. The scheduler must stay swappable (simple -> SM-2 -> FSRS).
12. Don't over-engineer the first version of anything (scheduler, analytics, sharing permissions beyond OWNER/VIEWER).
13. Use `StudyItem`, not `Word`, throughout the domain.
14. Keep the notebook template on `Notebook` as a raw string, not a separate entity.
15. Use UUID identifiers consistently across all entities.
16. Treat GORM tags directly on domain structs as an intentional pragmatic choice, not a violation of layering.
17. No SSR, no Node production runtime.

## License

This project is distributed under the BSD 3-Clause license (full text in `LICENSE` at the repo root, copyright Farzan Hajian). **When creating any new source file** (`.go`, `.ts`, `.tsx`, etc. — not needed for Markdown/docs/config files like `.md`, `.json`, `.yaml`), add a short header comment at the very top, before the package/import declarations, using that file type's comment syntax:

```go
// Copyright (c) 2026, Farzan Hajian
// All rights reserved.
//
// This source code is licensed under the BSD 3-Clause license found in the
// LICENSE file in the root directory of this source tree.
```

Keep the copyright year current; don't add this header to files that already have one (e.g. when editing an existing file, leave its existing header as-is rather than duplicating).

## Known open issue to flag, not silently resolve

Current token model is access token in a secure cookie, refresh token in local storage. Local-storage refresh tokens raise XSS blast-radius concerns and should be reviewed before real production deployment — if you're implementing production auth, flag this rather than changing the decision unilaterally.
