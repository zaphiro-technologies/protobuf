# ProtoBuf Release Notes

## 0.0.7-dev - 2024-06-18

### Dependencies

- Bump google.golang.org/protobuf from 1.34.1 to 1.34.2 (PR #60 by
  @dependabot[bot])

### Refactoring

- Refactor `Fault` messages to support measurement and measurement timestamp
  reference (PR #62 by @chicco785)

## 0.0.6 - 2024-06-04

### Features

- Add support for Measurement Types used in SIL (PR #57 by @chicco785)
- Update topology protobuf to include branches (PR #53 by @hiimjako)
- Support `NormalizedResidual` and `ObjectiveFunction` data types (PR #49 by
  @chicco785)

### Continuous Integration

- Configure Static Code & License Analysis (PR #55 by @chicco785)
- Replace custom proto-gen-md-diagrams with official one (PR #50 by @chicco785)

### Dependencies

- Bump google.golang.org/protobuf from 1.33.0 to 1.34.1 (PR #52 by
  @dependabot[bot])

### Refactoring

- Adopt `FaultStatus` to manage fault evolution over time (PR #59 by @chicco785)

## 0.0.5 - 2024-03-22

### Features

- Support python generation (PR #47 by @chicco785)

### Documentation

- Add `aligned` header to Data and DataSet (PR #42 by @hiimjako)

### Dependencies

- Bump github.com/stretchr/testify from 1.8.4 to 1.9.0 (PR #43 by
  @dependabot[bot])
- Bump google.golang.org/protobuf from 1.32.0 to 1.33.0 (PR #44 by
  @dependabot[bot])

## 0.0.4 - 2024-02-21

### Features

- Add `Stat` protobuf (PR #40 by @hiimjako)
- Support `ComputedTopology` protocol buffer (PR #39 by @hiimjako)

### Continuous Integration

- Update `releasenotes.yaml` to use `pull_request_review` (PR #35 by @chicco785)

### Dependencies

- Bump github.com/google/uuid from 1.5.0 to 1.6.0 (PR #36 by @dependabot[bot])

### Refactoring

- Remove `Event` from the name of `DeviceEvent` subclasses (PR #37 by
  @chicco785)

## 0.0.3 - 2024-02-09

### Features

- Add `NOTIFICATION_TYPE_TOPOLOGY_COMPUTED` (PR #34 by @chicco785)
- Support `DeviceEvent` messages (PR #31 by @chicco785)
- Add protocol buffer for grid events (PR #19 by @chicco785)

### Documentation

- Document headers/application properties data types for every Protobuf. (PR #27
  by @hiimjako)
- [Documentation] Change fault id data type (PR #26 by @tejo)

### Dependencies

- Bump google.golang.org/protobuf from 1.31.0 to 1.32.0 (PR #28 by
  @dependabot[bot])
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
