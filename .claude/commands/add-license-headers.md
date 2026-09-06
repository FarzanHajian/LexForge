---
description: Add the BSD 3-Clause license header to backend and frontend source files that are missing it
---

Add the project's standard BSD 3-Clause copyright header to every source file under `backend/` and `frontend/src/` (plus `frontend/index.html` and `frontend/vite.config.ts`) that doesn't already have one. Optional argument: `$ARGUMENTS` — if given, a path (or glob) to restrict the scan to instead of the whole repo.

## 1. Determine the exact header text

Read the "License" section of `CLAUDE.md` at the repo root and use the header template given there verbatim (copyright holder + year come from that section, which itself should match `LICENSE`). Do not hardcode the year/holder in this command — always source it fresh from `CLAUDE.md` so the command stays correct if the license/year ever changes. If `CLAUDE.md` has no License section or no template, stop and ask the user for the exact text before doing anything else.

The template is 4 lines of plain text (no comment markers yet), e.g.:
```
Copyright (c) <year>, <holder>
All rights reserved.

This source code is licensed under the BSD 3-Clause license found in the
LICENSE file in the root directory of this source tree.
```

## 2. Find candidate files

Scope (or `$ARGUMENTS` if provided instead):
- `backend/**/*.go` (excluding `go.sum`, generated files, and anything under `vendor/`)
- `frontend/src/**/*.ts` and `frontend/src/**/*.tsx` (include hand-authored `.d.ts` files like `vite-env.d.ts`; skip only files that are clearly auto-generated/vendored)
- `frontend/index.html`
- `frontend/vite.config.ts`

Never touch: `.json`, `.yaml`/`.yml`, `.md`, `.env*`, lockfiles, `node_modules/`, `dist/`, `build/`, or any other build output/dependency directory. These file types intentionally don't get a header per `CLAUDE.md`.

## 3. Skip files that already have the header

For each candidate file, check the first ~10 lines for an existing `Copyright (c)` line. If present, leave the file untouched — never duplicate or re-insert a header.

## 4. Insert the header using the right comment syntax, at the very top of the file (before package/import/doctype/anything else)

- `.go`, `.ts`, `.tsx` → `//` line comments, one per line of the template, e.g.:
  ```go
  // Copyright (c) 2026, Farzan Hajian
  // All rights reserved.
  //
  // This source code is licensed under the BSD 3-Clause license found in the
  // LICENSE file in the root directory of this source tree.
  ```
- `.html` → a single HTML block comment above `<!doctype html>`/`<html>`, e.g.:
  ```html
  <!--
  Copyright (c) 2026, Farzan Hajian
  All rights reserved.

  This source code is licensed under the BSD 3-Clause license found in the
  LICENSE file in the root directory of this source tree.
  -->
  ```
- Any `.css` file in scope → `/* */` block comment in the same shape as the HTML one.

Leave exactly one blank line between the header and the first real line of code (existing files in this repo already follow this convention — match it).

## 5. Report

After processing, print a concise summary: files updated (list), files skipped because they already had a header (list), and any files skipped because they didn't match an expected extension/pattern. Do not run `git add`/`git commit` — just leave the changes in the working tree for the user to review.
