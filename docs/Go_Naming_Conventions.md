# Go Naming Conventions

Reference notes on idiomatic Go naming, for consistency across the LexForge backend.

## Packages

- Short, lowercase, single word — no underscores, no mixedCaps (`config`, `domain`, not `http_api`/`httpAPI`).
- Named for what the package *provides*, not what it contains — avoid generic names like `util`/`common`/`helpers`.
- Package name is the last element of its import path (`internal/config` → package `config`).

## Files

- Lowercase, underscore-separated for multi-word names (`sysconfig.go`, `user_repository.go`), matching the file's primary type/concern.
- Test files: `<name>_test.go`, same package as the code under test (or `<pkg>_test` package for black-box tests).
- OS/arch build-constraint suffixes are recognized by the toolchain (`_windows.go`, `_amd64.go`).
- `doc.go`: a file with only a package clause and a package-level doc comment, used when there's no single file to hang that comment on. Every `internal/*` package in this repo has one — read it before adding code to a package to pick up any conventions it records.

## Types

- Exported: `PascalCase` (`User`, `SysConfig`, `PracticeSession`).
- Unexported: `camelCase` (`userModel`, `sysConfig`).
- Single-method interfaces are conventionally named `<Verb>er` (`Reader`, `Scheduler`, `TokenValidator`).
- **No `I` prefix on interfaces** (not `IUserRepository`) and **no `Impl` suffix on implementations** — Go doesn't use the C#/Java conventions here (compare `io.Reader`, not `IReader`; `bytes.Buffer`, not `BufferImpl`).
  - The interface gets the clean, meaningful name: `UserRepository`.
  - The implementation is named after the concrete detail it adds (the technology, not "being an implementation"): `GormUserRepository`. If there were a second implementation, it'd be distinguished the same way (`InMemoryUserRepository`), never by a generic suffix.

## Functions / methods

- Exported: `PascalCase` (`NewUser`, `NewTopic`, `Load`).
- Unexported: `camelCase` (`parseEnv`, `newTopicRow`).
- Constructors: `New<Type>`, returning the type (or `(*T, error)` if construction can fail) — e.g. `NewUser`, `NewTopic`, `NewGormUserRepository`.
- No `Get` prefix on getters (`user.Name()`, not `user.GetName()`); setters do use `Set` (`user.SetName(...)`).
- Receiver names: short (1–2 letters, an abbreviation of the type), consistent across all methods of that type, never `this`/`self` (`func (u *User) ...`).

## Variables / constants

- `camelCase` for locals and unexported package vars; `PascalCase` for exported.
- Acronyms stay fully cased, not mixed: `ID`/`userID`/`httpURL`, not `Id`/`userId`. (Note: this repo currently uses `Id`/`UserId`/`ExternalId` in the domain structs — a deliberate deviation from strict idiom, not yet reconciled.)
- Constants grouped in `const (...)` blocks, often with `iota` for enums; no `SCREAMING_SNAKE_CASE` (a C convention, not Go's).

## Modules

- Module path is a URL-like string, typically the repo location, lowercase, no underscores (`github.com/FarzanHajian/lexforge/backend`).

## Errors

- Sentinel errors are `camelCase`/`PascalCase` values named `Err<Thing>` (`ErrNotFound`, `ErrUserNotFound`), not stringly-typed or exception-style.
