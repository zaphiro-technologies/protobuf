# ProtoBuffer documentation

- _Grid_ package collects grid related messages, which currently include:
  - [`Data`](./grid/v1/data.proto.md) dedicated to measurements related messages.
  - [`Fault`](./grid/v1/fault.proto.md) dedicated to faults related messages.
  - [`Topology`](./grid/v1/topology.proto.md) dedicated to topology related messages.
  - [`Event`](./grid/v1/event.proto.md) dedicated to general event messages,
    inherited by:
    - [`DeviceEvent`](./grid/v1/device_event.proto.md): events related to devices.
    - [`GridEvent`](./grid/v1/grid_event.proto.md): events related to grid.
- _Platform_ package collects platform related messages, which currently include:
  - [`Task`](./platform/v1/task.proto.md) dedicated to platform tasks and notifications.
  - [`Blacklist`](./platform/v1/blacklist.proto.md) dedicated to management of
    measurements blacklist.
- _C37118_ package collects messages related to the IEEE c37.118 standard,
  which currently include:
  - [`Conf2Frame`](./c37118/v1/conf.proto.md) dedicated to PMU configuration frames.
  - [`Stat`](./c37118/v1/stat.proto.md) dedicated to PMU statistics.
