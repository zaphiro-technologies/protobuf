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

At the time being this repository includes the following Protocol Buffer:

- _Data_ used to manage measurements injected in the platform, defined in
  [`proto/data`](./proto/data), which currently include:
  - `Data`: a single data defined by a `value`, a type of data (`dataType`) -
    which can be `DIGITAL` or `PHASOR` -, and the time when the data was
    measured (`measuredAt`).
  - `DataSet`: a collection of data defined by a producer(`producerId`) and a
    map of `Data` (where the key is the `uuid` of that `Data`).
- _Task_ used to manage tasks in the platform, defined in
  [`proto/task`](./proto/task), which currently include:
  - `Task`: a task identified by an `id`, a `taskType`, a time of creation
    (`createdAt`) and `timestampID` and/or a `measurementID`.
  - `Notification`: a notification identified by an `id`, a `notificationType`,
    a time of creation (`createdAt`), a `message` and `timestampID` and/or a
    `measurementID`.
- _Conf_ used to manage configurations in the platform, defined in
  [`proto/conf`](./proto/conf), which currently include:
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
