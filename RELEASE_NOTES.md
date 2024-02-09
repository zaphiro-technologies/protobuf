# ProtoBuf Release Notes

## 0.0.3-dev - 2024-02-09

### Features

- Add `NOTIFICATION_TYPE_TOPOLOGY_COMPUTED` (PR #34 by @chicco785)
- Support `DeviceEvent` messages (PR #31 by @chicco785)
- Add protocol buffer for grid events (PR #19 by @chicco785)

### Documentation

- Document headers/application properties data types for every Protobuf. (PR #27
  by @hiimjako)
- [Documentation] Change fault id data type (PR #26 by @tejo)

### Dependencies

- Bump github.com/google/uuid from 1.3.1 to 1.5.0 (PR #29 by @dependabot[bot])
- Bump actions/setup-go from 4 to 5 (PR #25 by @dependabot[bot])

### Refactoring

- Event: set `occurredAt` as required, correct Event Status to `Started`,
  `Ended`, and add `eventType` in the header (PR #23 by @chicco785)
- Fault: uniform datetime field naming (PR #21 by @chicco785)

## 0.0.2 - 2023-11-14

### Bug Fixes

- Bug: Use int64 for topology's createdAt. (PR #16 by @tejo)

### Documentation

- Update documentation to reflect protocol buffer refactoring (PR #15 by
  @chicco785)

### Refactoring

- Update NotificationType names (PR #17 by @hiimjako)

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
