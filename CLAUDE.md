# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

A minimal Go library that wraps `sirupsen/logrus` with a custom `PlainFormatter` for consistent plain-text log output. Single-file project (`logrus.go`).

## Development Commands

| Task         | Command          |
|--------------|------------------|
| Build        | `go build ./...` |
| Test         | `go test ./...`  |
| Tidy modules | `go mod tidy`    |

## Notes

- No linter config — uses standard `gofmt` formatting
- Dependency updates are managed by Dependabot (weekly)
- go.mod version must stay aligned with upstream logrus release tag (currently go 1.17)

## Project Memory System

This project maintains institutional knowledge in `docs/project_notes/` for consistency across sessions.

### Memory Files

- **bugs.md** - Bug log with dates, solutions, and prevention notes
- **decisions.md** - Architectural Decision Records (ADRs) with context and trade-offs
- **key_facts.md** - Project configuration, important constants, references
- **issues.md** - Work log with descriptions and dates

### Memory-Aware Protocols

**Before proposing architectural changes:** Check `docs/project_notes/decisions.md` for existing decisions. If conflicting, acknowledge and explain why a change is warranted.

**When encountering errors:** Search `docs/project_notes/bugs.md` for similar issues. Document new bugs and solutions when resolved.

**When looking up project config:** Check `docs/project_notes/key_facts.md` for module info, log format details, dependency versions.

**When completing work:** Log in `docs/project_notes/issues.md` with date and brief description.
