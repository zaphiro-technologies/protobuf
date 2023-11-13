# ProtoBuf Release Notes

## 0.0.2-dev - 2023-11-13

### Bug Fixes

- Bug: Use int64 for topology's createdAt. (PR #16 by @tejo)

## 0.0.1 - 2023-10-27

### Features

- New protocol buffer for topology data (PR #12 by @chicco785)
- Notification message: support Trigger requirements (PR #6 by @chicco785)
- Add protocol buffer to describe faults (PR #3 by @chicco785)

### Bug Fixes

- fix typo in `Fault` field `faultCurrent` (PR #7 by @tejo)

### Continuous Integration

- Use new shared workflow for Golang (PR #11 by @chicco785)

### Refactoring

- Refactor protobuffer to align with architecture discussion (PR #9 by
  @chicco785)
- Migrate existing protobuf from
  `https://github.com/zaphiro-technologies/rabbitmq-poc` (PR #2 by @chicco785)
