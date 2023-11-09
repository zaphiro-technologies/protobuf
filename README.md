# Proto Buf

This repository includes the [Protocol Buffer](https://protobuf.dev/) used in
Zaphiro's platform.

## Requirements

1. [Golang](https://go.dev/doc/install)
1. [Protocol Buffer Compiler](https://grpc.io/docs/protoc-installation/)
1. [Go plugins for the protocol compiler](https://grpc.io/docs/languages/go/quickstart/#prerequisites)
1. [Buf](https://buf.build/docs/installation)
1. [proto-gen-md-diagrams](https://github.com/GoogleCloudPlatform/proto-gen-md-diagrams)

## Available Protocol Buffer

At the time being this repository includes the following Protocol Buffer
packages:

- _Grid_ package collects grid related messages, defined in
  [`proto/grid/v1`](./proto/grid/v1/), which currently include:
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
- _Platform_ package collects platform related messages, defined in
  [`proto/platform/v1`](./proto/platform/v1/), which currently include:
  - `Task`: a task to be performed by an service in the platform.
  - `Notification`: a notification produced by a service in the platform.
- _Conf_ package collects messages related to configuration in the platform,
  defined in [`proto/conf/v1`](./proto/conf/v1/), which currently include:
  - `Conf2Frame`: a Protocol Buffer used to store PMU configuration frames.

## Protocol Buffers Management

All protocol buffers are documented in [docs](./docs).

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
