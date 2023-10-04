# go-template Release Notes

## 0.0.1-dev - 2023-09-27

### Features

- Manage environment variables via `.env` files (PR #11 by @hiimjako)
- set-up base template, including log lib, testing lib, lint wf, test wf, and
  benchmark wf (PR #2 by @chicco785)

### Continuous Integration

- Benchmark job: add fail fast and skip tests (PR #16 by @hiimjako)
- Remove annoying lint rules (PR #10 by @hiimjako)
- Add default go folders `cmd/` and `vendor/` to ci paths (PR #12 by @hiimjako)
- Add write content permission to golang ci workflow (PR #7 by @chicco785)

### Dependencies

- updated to new version of logger version (PR #14 by @chicco785)
- Bump actions/cache from 1 to 3 (PR #4 by @dependabot[bot])
- Bump actions/checkout from 3 to 4 (PR #5 by @dependabot[bot])

### Refactoring

- Use zaphiro logging wrapper (PR #9 by @chicco785)
