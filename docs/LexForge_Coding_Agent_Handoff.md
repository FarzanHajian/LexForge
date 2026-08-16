# LexForge — Coding Agent Handoff

## Current decisions

### Stack

**Frontend**
- React
- TypeScript
- Vite
- Tailwind CSS
- React Router
- Static SPA; no Next.js and no SSR
- Vite build output is served directly by the Go/Echo backend
- No Node.js runtime, CDN, or separate frontend HTTP server in production

**Backend**
- Go
- Echo
- GORM
- MySQL
- REST API returning JSON
- UUIDs for entity IDs

**Authentication**
- Auth0 handles authentication/identity
- Use the official Auth0 Go SDK
- Isolate Auth0 SDK behind an application-level authentication abstraction
- Internal users have an `auth0_sub` linking them to Auth0
- Current token decision: access token in a secure cookie; refresh token in local storage
- Review refresh-token storage security before production; do not silently change the chosen architecture

### Domain terminology

The vocabulary entity is called **`StudyItem`**, not `Word`.

A study item can be a word, phrase, expression, idiom, sentence, or other language-learning item.

Use names such as:

```text
StudyItem
StudyItemReview
UserStudyItemState
PracticeSessionItem
```

### Topic template

A topic has one template, stored directly as a raw string column on `Topic`.

Do **not** introduce a separate `TopicTemplate` entity/table unless a future requirement gives templates an independent lifecycle.

### GORM tags

GORM tags are intentionally placed on application structs. This couples the structs to GORM, but is an accepted pragmatic choice for this project.

Do not create separate domain/persistence models solely to eliminate GORM tags.

Keep database access behind repository/infrastructure boundaries:

```text
Handler
  -> Application Service
    -> Repository interface
      -> GORM repository
        -> MySQL
```

### Practice sessions

A user can practice a topic at most once per calendar day.

Enforce this with a database constraint equivalent to:

```text
UNIQUE(user_id, topic_id, date)
```

Session statuses:

```text
IN_PROGRESS
COMPLETED
EXPIRED
```

The user chooses the session size before starting each day's session. Allowed values:

```text
5, 10, 15, 20
```

The selected value is stored on the session and must not change after it starts.

`max_sessions_per_day` is system-wide configuration.

A practice session is a snapshot: its selected study items/order should not change while it is being practiced.

### Session expiration

**Do not run a background job at midnight just to expire sessions.**

Expiration is lazy.

When the user attempts to start a new practice session, the backend checks the relevant existing session/date. If an unfinished session from a previous calendar day is encountered:

```text
IN_PROGRESS + old date -> EXPIRED
```

Then the new day's session can be created.

The backend must enforce expiration. Answers must not be accepted for an expired/previous-day session.

Expired sessions do **not** participate in session statistics:

```text
COMPLETED  -> included
IN_PROGRESS -> not included
EXPIRED     -> not included
```

However, reviews that were actually submitted before a session was abandoned remain valid learning-history records.

Therefore:

```text
Session statistics -> completed sessions
Learning statistics -> actual recorded reviews
```

### Learning state and reviews

Learning state is per `(user, study_item)`.

Keep current state separate from historical reviews:

```text
UserStudyItemState -> current aggregate state
StudyItemReview    -> historical review events
```

A conceptual state can include:

```text
user_id
study_item_id
status
last_reviewed_at
next_review_at
correct_count
wrong_count
```

A review records at least:

```text
user_id
study_item_id
session_id
result
reviewed_at
```

Initially the UI can use:

```text
KNOWN
UNKNOWN
```

### Spaced repetition

Keep the scheduler replaceable.

Intended progression:

```text
Simple scheduler
    -> SM-2-style scheduler
    -> potentially FSRS later
```

New study items must be supported.

A first implementation can prioritize due/overdue items, difficult/unknown items, new items, then other eligible items.

## Testing

Preferred tools:
- Go `testing`
- Testify
- Testify/mock
- Testcontainers for integration tests

Use mocks/fakes for pure unit tests.

Use **real MySQL**, not SQLite, for integration tests.

Important integration cases:
- daily-session uniqueness
- session expiration
- migrations
- topic sharing
- learning-state persistence
- repository queries

## Topic sharing

A user owns topics and can share them with another user in read-only mode.

Initial permissions:

```text
OWNER
VIEWER
```

A viewer can read the topic/study items but cannot modify them.

## Observability

Planned portfolio observability stack:

- OpenTelemetry
- OpenTelemetry Collector
- Prometheus
- Grafana
- Tempo
- Loki

Desired flow:

```text
Go/Echo
  -> OpenTelemetry Collector
      -> Prometheus (metrics)
      -> Tempo (traces)
      -> Loki (logs)
              -> Grafana
```

Useful traces include starting a practice session and its authentication, authorization, study-item selection, session creation, and database operations.

Useful metrics include HTTP latency/error rate, database timing, sessions completed, study items reviewed, known/unknown results, and practice duration.

Use structured logs and never log secrets, tokens, passwords, or unnecessary personal data.

Do not add Kubernetes/service-mesh complexity unless a concrete requirement emerges.

## Deployment

Production should be deliberately simple:

```text
Browser
   |
   v
Go/Echo application
   |-- REST API
   |-- static React/Vite files
   |
   v
MySQL
```

The Go application may eventually embed the Vite `dist` directory using `go:embed`.

## Suggested skills

Invoke only skills relevant to the current task.

### Core
- Go
- Echo
- GORM
- MySQL
- React
- TypeScript
- Vite
- Tailwind CSS
- React Router

### Authentication
- Auth0
- OAuth 2.0 / OpenID Connect
- JWT validation
- Auth0 Go SDK
- secure cookie authentication

### Testing
- Go testing
- Testify
- Testify/mock
- Testcontainers
- MySQL integration testing

### Observability
- OpenTelemetry
- OpenTelemetry Go
- OpenTelemetry Collector
- Prometheus
- Grafana
- Tempo
- Loki

### Engineering
- REST API design
- SQL/database design
- Go interfaces/dependency inversion
- pragmatic layered architecture
- secure web application development
- Docker

## Implementation principles

1. Prefer simplicity; do not introduce technology merely because it is interesting.
2. Keep business logic independent of Echo, GORM, MySQL, and Auth0 where practical.
3. Use interfaces at meaningful boundaries.
4. Test business rules independently of the database.
5. Use real MySQL for integration tests.
6. Preserve review history.
7. Keep learning state user-specific.
8. Treat a practice session as a snapshot.
9. Expire unfinished sessions lazily; do not use a midnight expiration job.
10. Exclude expired sessions from session statistics.
11. Keep actual reviews from expired sessions in learning history.
12. Keep the scheduler replaceable.
13. Do not over-engineer the initial version.
14. Use `StudyItem`, not `Word`, throughout the domain.
15. Keep the topic template on `Topic` as a raw string.
16. Use UUID identifiers consistently.
17. Treat GORM tags on structs as an intentional pragmatic choice.
18. Do not add SSR or a Node production runtime.

## Existing artifacts

Do not duplicate detailed specs, plans, ADRs, issues, commits, or diffs here. Reference them instead.

When present, treat these as authoritative:

```text
/docs/
/docs/specs/
/docs/plans/
/docs/adr/
/issues/
/CHANGELOG.md
/README.md
```

If a newer authoritative artifact conflicts with this handoff, follow the newer artifact and flag the conflict.

## Agent startup

Before coding, inspect:
1. repository structure
2. existing specs/plans/ADRs
3. current implementation status
4. tests
5. configuration/environment handling
6. frontend/backend development workflow

Then implement incrementally without silently changing the decisions above.
