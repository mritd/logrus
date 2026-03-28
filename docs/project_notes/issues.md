# Work Log

Track completed work with dates, descriptions, and references.

### 2026-03-28 - Add WithFields support and performance optimization
- **Status**: Completed
- **Description**: PlainFormatter now outputs entry.Data fields. Refactored to use entry.Buffer, type-specific writeValue, and direct buffer writes for reduced allocations.
- **Notes**: Also added boundary checks for LevelDesc/TimestampFormat zero values, and test suite with 91.3% coverage.
