# ProtoBuf Release Notes

## 0.0.20-dev - 2026-01-14

### Dependencies

- Bump google.golang.org/protobuf from 1.36.10 to 1.36.11 (PR #127 by
  @dependabot[bot])
- Bump actions/checkout from 5 to 6 (PR #125 by @dependabot[bot])
- Bump golangci/golangci-lint-action from 8 to 9 (PR #123 by @dependabot[bot])
- Bump github.com/ccoveille/go-safecast from 1.7.0 to 1.8.2 (PR #124 by
  @dependabot[bot])

## 0.0.19 - 2025-10-24

### Features

- Change source type to string (PR #121 by @hiimjako)

## 0.0.18 - 2025-10-23

### Features

- Add casting methods for constants (PR #120 by @hiimjako)

## 0.0.17 - 2025-10-23

- no changes

## 0.0.17 - 2025-10-23

- no changes

## 0.0.17 - 2025-10-23

### Features

- Add shared constants to model message headers (PR #119 by @hiimjako)

## 0.0.16 - 2025-10-21

### Features

- data proto: add absolute residual data type (PR #116 by @hiimjako)

### Continuous Integration

- Install locally proto-gen-md-diagrams (PR #117 by @chicco785)

### Documentation

- Add blacklist headers to identify the message (PR #113 by @hiimjako)

### Dependencies

- Bump github.com/ccoveille/go-safecast from 1.6.1 to 1.7.0 (PR #115 by
  @dependabot[bot])
- Bump google.golang.org/protobuf from 1.36.9 to 1.36.10 (PR #114 by
  @dependabot[bot])
- Bump google.golang.org/protobuf from 1.36.7 to 1.36.9 (PR #111 by
  @dependabot[bot])

### Refactoring

- Deprecate `message` field in `Event` (PR #112 by @chicco785)

## 0.0.15 - 2025-09-15

### Features

- stats and dataset proto: add latency header (PR #110 by @chicco785)

### Dependencies

- Bump github.com/stretchr/testify from 1.10.0 to 1.11.0 (PR #107 by
  @dependabot[bot])
- Bump actions/setup-go from 5 to 6 (PR #109 by @dependabot[bot])
- Bump actions/checkout from 4 to 5 (PR #106 by @dependabot[bot])
- Bump google.golang.org/protobuf from 1.36.6 to 1.36.7 (PR #105 by
  @dependabot[bot])

## 0.0.14 - 2025-06-19

### Dependencies

- Bump protobuf from 5.26.0 to 5.29.5 (PR #102 by @dependabot[bot])

### Refactoring

- Refactor blacklist proto to support blacklisting specific for service (PR #104
  by @chicco785)

## 0.0.13 - 2025-05-12

### Dependencies

- Bump golangci/golangci-lint-action from 7 to 8 (PR #95 by @dependabot[bot])

### Refactoring

- `Event` message: deprecate `sourceId` and `sourceType` and move to headers (PR
  #100 by @chicco785)

## 0.0.12 - 2025-05-08

### Features

- Add blacklist message (PR #96 by @chicco785)

## 0.0.11 - 2025-05-08

### Features

- Add event test service event source type (PR #98 by @tejo)

### Continuous Integration

- Add auto-approve & auto-merge Dependabot PRs workflow (PR #94 by @chicco785)

## 0.0.10 - 2025-04-30

### Features

- Add config 3 compatibility in config proto (PR #93 by @tejo)

### Dependencies

- Bump google.golang.org/protobuf from 1.36.5 to 1.36.6 (PR #91 by
  @dependabot[bot])
- Bump github.com/ccoveille/go-safecast from 1.6.0 to 1.6.1 (PR #89 by
  @dependabot[bot])
- Bump golangci/golangci-lint-action from 6 to 7 (PR #90 by @dependabot[bot])
- Bump github.com/ccoveille/go-safecast from 1.5.0 to 1.6.0 (PR #88 by
  @dependabot[bot])
- Bump google.golang.org/protobuf from 1.36.4 to 1.36.5 (PR #87 by
  @dependabot[bot])
- Bump google.golang.org/protobuf from 1.36.3 to 1.36.4 (PR #86 by
  @dependabot[bot])

## 0.0.9 - 2025-01-21

### Features

- Add Updated to `FaultEventType` (PR #85 by @chicco785)
- Add `FrequencyVariation` Event (PR #80 by @chicco785)

### Continuous Integration

- Add trivy cache workflow (PR #81 by @chicco785)

### Dependencies

- Bump google.golang.org/protobuf from 1.36.2 to 1.36.3 (PR #84 by
  @dependabot[bot])
- Bump google.golang.org/protobuf from 1.36.1 to 1.36.2 (PR #83 by
  @dependabot[bot])
- Bump github.com/ccoveille/go-safecast from 1.2.0 to 1.5.0 (PR #78 by
  @dependabot[bot])
- Bump google.golang.org/protobuf from 1.36.0 to 1.36.1 (PR #79 by
  @dependabot[bot])
- Bump github.com/stretchr/testify from 1.9.0 to 1.10.0 (PR #72 by
  @dependabot[bot])
- Bump google.golang.org/protobuf from 1.34.2 to 1.36.0 (PR #77 by
  @dependabot[bot])

## 0.0.8 - 2024-12-17

### Features

- Add `samplingPeriod` property (PR #73 by @chicco785)

### Continuous Integration

- Add `new-release` workflow (PR #76 by @chicco785)
- Pass `-coverpkg=./...` to `go test` to compute coverage across all packages
  (PR #61 by @chicco785)

### Documentation

- Fix documentation of headers for faults (PR #68 by @chicco785)
- Update documentation for open source release (PR #66 by @chicco785)

## 0.0.7 - 2024-06-24

### Documentation

- Update stat producerId documentation (PR #64 by @hiimjako)

### Dependencies

- Bump google.golang.org/protobuf from 1.34.1 to 1.34.2 (PR #60 by
  @dependabot[bot])

### Refactoring

- Fault messages as set of events, rather than single event with multiple status
  (PR #63 by @chicco785)
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
