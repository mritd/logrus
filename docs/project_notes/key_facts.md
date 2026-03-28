# Key Facts

Project configuration, important constants, and reference information.

> **Security**: Never store secrets, passwords, or API keys here. Use env vars or secret managers.

## Project

- **Module**: github.com/mritd/logrus
- **Type**: Go library (not an application)
- **Go version**: 1.17 (aligned with upstream logrus v1.9.4 tag)
- **Primary dependency**: github.com/sirupsen/logrus v1.9.4

## Log Format

- **Pattern**: `{timestamp} {level} {message} [{key=value}...]`
- **Timestamp**: `2006-01-02 15:04:05`
- **Levels**: PANC, FATL, ERRO, WARN, INFO, DEBG, TRAC (4-char abbreviations)
- **Unknown level fallback**: `UNKN`

## Dependency Management

- Dependabot enabled (weekly schedule) for Go modules
