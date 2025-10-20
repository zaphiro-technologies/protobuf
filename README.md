# SynchroGuard's Protocol Buffer Messages

In SynchroGuard, real-time data exchange relies on protocol buffer messages
exchanged over [RabbitMQ](https://www.rabbitmq.com/). Thus, consuming and/or
publishing protocol buffer messages on RabbitMQ is the preferred way to
integrate third parties with SynchroGuard for real-time data exchange.

This repository includes the [Protocol Buffer](https://protobuf.dev/) used in
Zaphiro's platform and the generated code for GO and Python.

## Available Protocol Buffer Messages

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
    - `Topology`: a message represented an topology computed by the system.
  - `Event` protocol buffer package that define general event messages,
    inherited by:
    - `GridEvent`: messages that represent grid events.
    - `DeviceEvent`: messages that represent events related to devices.
- _Platform_ package collects platform related messages, defined in
  [`zaphiro/platform/v1`](./zaphiro/platform/v1/), which currently include:
  - `Task`: a task to be performed by an service in the platform.
  - `Notification`: a notification produced by a service in the platform.
  - `Blacklist`: a message that describes for each service which measurement to
    not use.
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
poetry add git+https://github.com/zaphiro-technologies/protobuf.git#v0.0.7
```

### Go

```bash
go get -v github.com/zaphiro-technologies/protobuf/go@v0.0.7
```

## Examples

For your convenience, in the [examples](examples) folder we provide Go code to:

- Produce and consume measurements (uses RabbitMQ streams).
- Produce and consume faults (uses RabbitMQ exchanges).

### Requirements

- Golang 1.21
- [Docker](https://docs.docker.com/get-docker/)

### Produce and consume measurements

#### Go

1. Set-up the infrastructure using `make docker-start`.
1. Launch the example `make example-measurements-go`
1. Press any key to stop.

```bash
Getting started with Streaming client for RabbitMQ
Connecting to RabbitMQ streaming ...
Sent 10 messages
Press any key to stop

consumer name: my_consumer, measurement_id: Dev0000-0005, measurement_time 1720698360980, measurement_type 20, measurement_value 4592455024224327647
consumer name: my_consumer, measurement_id: Dev0000-0006, measurement_time 1720698360980, measurement_type 20, measurement_value 4604241342922663796
consumer name: my_consumer, measurement_id: Dev0000-0009, measurement_time 1720698360980, measurement_type 44, measurement_value 4603643564277470555
consumer name: my_consumer, measurement_id: Dev0000-0001, measurement_time 1720698360980, measurement_type 28, measurement_value 4598295346455196891
consumer name: my_consumer, measurement_id: Dev0000-0004, measurement_time 1720698360980, measurement_type 20, measurement_value 4605375802349382247
consumer name: my_consumer, measurement_id: Dev0000-0007, measurement_time 1720698360980, measurement_type 20, measurement_value 4605095969928423361
consumer name: my_consumer, measurement_id: Dev0000-0002, measurement_time 1720698360980, measurement_type 21, measurement_value 4600833284970982103
consumer name: my_consumer, measurement_id: Dev0000-0008, measurement_time 1720698360980, measurement_type 20, measurement_value 4597854724057569946
consumer name: my_consumer, measurement_id: Dev0000-0000, measurement_time 1720698360980, measurement_type 13, measurement_value 4603013690508174157
consumer name: my_consumer, measurement_id: Dev0000-0003, measurement_time 1720698360980, measurement_type 21, measurement_value 4598093240485266093
consumer name: my_consumer, measurement_id: Dev0000-0008, measurement_time 1720698360980, measurement_type 20, measurement_value 4600766170106400936
consumer name: my_consumer, measurement_id: Dev0000-0007, measurement_time 1720698360980, measurement_type 20, measurement_value 4606916483465028581
```

### Produce and consume faults

#### Go

1. Set-up the infrastructure using `make docker-start`.
1. Launch the example `make example-faults-go`
1. Press any key to stop.

```bash
Getting started with AMPQ client for RabbitMQ
Connecting to RabbitMQ ...
Received a fault event message: 1948cd36-f835-4dc5-83c2-ba5a7612362a, event type: FAULT_EVENT_TYPE_STARTED
Press any key to stop

Received a fault event message: 1948cd36-f835-4dc5-83c2-ba5a7612362a, event type: FAULT_EVENT_TYPE_ENDED
Received a line fault event message: 1948cd36-f835-4dc5-83c2-ba5a7612362a, event type: FAULT_EVENT_TYPE_LOCATED, faulty line: EQ-1, probability: 0.330000, length from t1: 0.526424
Received a line fault event message: 1948cd36-f835-4dc5-83c2-ba5a7612362a, event type: FAULT_EVENT_TYPE_LOCATED, faulty line: EQ-2, probability: 0.330000, length from t1: 0.484818
Received a line fault event message: 1948cd36-f835-4dc5-83c2-ba5a7612362a, event type: FAULT_EVENT_TYPE_LOCATED, faulty line: EQ-3, probability: 0.330000, length from t1: 0.315550
```

## Contributing to Protocol Buffers

Should you need to modify the protocol buffer message or add new ones, you will
need to set-up the dependencies listed in [Requirements](#requirements).

Protocol buffers are versioned (current version is v1), and should be developed
following best practices, as implemented by [Buf](https://buf.build) and defined
in
[Protobuf programming guides](https://protobuf.dev/best-practices/dos-donts/).
In particular, it is important - even more within the same version - to preserve
compatibility, to avoid services breaking up.

You find contributing guidelines [here](CONTRIBUTING.md).

### Requirements

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
make install-proto-gen-md
make docs
```

### Lint Protocol Buffers

```bash
make proto-lint
```

## References

1. [10 Protobuf Versioning Best Practices](https://climbtheladder.com/10-protobuf-versioning-best-practices/)
1. [Protos Best Practices](https://protobuf.dev/best-practices/dos-donts/)
