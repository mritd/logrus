# Architectural Decision Records

Document key decisions with context, alternatives, and consequences.

### ADR-001: Use PlainFormatter instead of logrus built-in formatters (2025-01-01)

**Context:**
- Need consistent, human-readable log output across projects
- logrus built-in TextFormatter includes field prefixes and quoting that clutters output

**Decision:**
- Custom PlainFormatter with format: `{timestamp} {level} {message} {fields}`
- 4-char level abbreviations: PANC, FATL, ERRO, WARN, INFO, DEBG, TRAC

**Consequences:**
- Simpler log output, easier to read in terminal
- Fields auto-sorted by key for stable output
- Values with special chars (spaces, equals, quotes) are automatically quoted

### ADR-002: Keep go.mod aligned with upstream logrus (2026-03-28)

**Context:**
- This library wraps sirupsen/logrus and should be usable by the same Go versions
- golang.org/x/sys version must also align to avoid pulling in higher go directive

**Decision:**
- Match `go` directive and `golang.org/x/sys` version to upstream logrus release tag (not main branch)

**Consequences:**
- Broader compatibility with consumer projects
- Must check upstream when bumping dependencies
