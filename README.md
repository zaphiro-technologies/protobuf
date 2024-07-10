# SynchroGuard's Protocol Buffers

In SynchroGuard, real-time data exchange relies on protocol buffer messages
exchanged over [RabbitMQ](https://www.rabbitmq.com/). Thus, consuming and/or
publishing protocol buffer messages on RabbitMQ is the preferred way to
integrate third parties with SynchroGuard for real-time data exchange.

This repository includes the [Protocol Buffer](https://protobuf.dev/) used in
Zaphiro's platform and the generated code for GO and Python.

## Available Protocol Buffer

At the time being this repository includes the following Protocol Buffer
packages:

- _Grid_ package collects grid related messages, defined in
  [`zaphiro/grid/v1`](./zaphiro/grid/v1/), which currently include:
  - `Data` protocol buffer package, that includes the following messages
    - `Data`: a single measurement data
    - `DataSet`: a collection of measurement data.
  - `Fault` protocol buffer package, that includes the following messages:
    - `Fault`: an abnormal condition causing current flow through conducting
      equipment, such as caused by equipment failure or short circuits from
      objects not typically modelled.
    - `Line Fault`: a fault that occurs on an AC line segment at some point
      along the length.
    - `Equipment Fault`: a fault applied at the terminal, external to the
      equipment. This class is not used to specify faults internal to the
      equipment.
  - `Topology` protocol buffer package, that includes the following messages:
    - `Topology`: an message represented an topology computed by the system.
  - `Event` protocol buffer package that define general event messages,
    inherited by:
    - `GridEvent`: a message that represent grid events.
- _Platform_ package collects platform related messages, defined in
  [`zaphiro/platform/v1`](./zaphiro/platform/v1/), which currently include:
  - `Task`: a task to be performed by an service in the platform.
  - `Notification`: a notification produced by a service in the platform.
- _C37118_ package collects messages related to the IEEE c37.118 standard,
  defined in [`zaphiro/c37118/v1`](./zaphiro/c37118/v1/), which currently
  include:
  - `Conf2Frame`: a Protocol Buffer used to store PMU configuration frames.
  - `Stat`: a Protocol Buffer used to store PMU measurement Stat.

All protocol buffers are documented in [docs](./docs).

## Installation

Currently code is generated for:

- [go](go)
- [python](python)

### Python

Installation was tested with poetry:

```bash
poetry add git+https://github.com/zaphiro-technologies/protobuf.git#v0.0.5
```

### Go

```bash
go install github.com/zaphiro-technologies/protobuf/go@v0.0.5
```

## Contributing to Protocol Buffers

Should you need to modify the protocol buffer message or add new ones, you will
need to set-up the dependencies listed in [Requirements](#requirements).

Protocol buffers are versioned (current version is v1), and should be developed
following best practices, as implemented by [Buf](https://buf.build) and defined
in [Protobuf programming
guides](https://protobuf.dev/programming-guides/dos-donts/).
In particular, it is important - even more within the same version - to preserve
compatibility, to avoid services breaking up.

## Requirements

1. [Golang](https://go.dev/doc/install)
1. [Protocol Buffer Compiler](https://grpc.io/docs/protoc-installation/)
1. [Go plugins for the protocol compiler](https://grpc.io/docs/languages/go/quickstart/#prerequisites)
1. [Buf](https://buf.build/docs/installation)
1. [proto-gen-md-diagrams](https://github.com/GoogleCloudPlatform/proto-gen-md-diagrams)

### Generate Code from Protocol Buffers

```bash
make generate
```

### Generate Documentation from Protocol Buffers

```bash
make docs
```

### Lint Protocol Buffers

```bash
make proto-lint
```

## References

1. [10 Protobuf Versioning Best Practices](https://climbtheladder.com/10-protobuf-versioning-best-practices/)
1. [Protos Best Practices](https://protobuf.dev/programming-guides/dos-donts/)
