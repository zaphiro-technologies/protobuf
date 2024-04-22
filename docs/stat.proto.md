# Package: zaphiro.c37118.v1

<!-- markdownlint-disable -->
Messages to support stat data injection exchange in the platform.



## Imports

| Import | Description |
|--------|-------------|



## Options

| Name       | Value       | Description |
|------------|-------------|-------------|
| go_package | ./c37118/v1 |             |




### Stat Diagram

```mermaid
classDiagram
direction LR

%% A Stat of PMU data, check IEEE C37.118 standard for more details.
%% Headers used in rabbitMQ:
%% * `id` (string): id of the `Stat` message.
%% * `producerId` (string): the id of the PMU linked to the Stat.
%% * `timestampId` (int64): related measurement Unix msec timestamp (if any)
%% 

class Stat {
  + int64 measuredAt
  + uint32 error
  + bool sync
  + bool sorting
  + bool trigger
  + bool configChange
  + bool dataModified
  + uint32 timeQuality
  + uint32 unlockedTime
  + uint32 triggerReason
}

```

## Message: Stat

**FQN**: zaphiro.c37118.v1.Stat

A Stat of PMU data, check IEEE C37.118 standard for more details.
Headers used in rabbitMQ:
* `id` (string): id of the `Stat` message.
* `producerId` (string): the id of the PMU linked to the Stat.
* `timestampId` (int64): related measurement Unix msec timestamp (if any)



| Field           | Ordinal | Type     | Label | Description                                                   |
|-----------------|---------|----------|-------|---------------------------------------------------------------|
| `measuredAt`    | 1       | `int64`  |       | The time of creation of the stat data (Unix msec timestamp).  |
| `error`         | 2       | `uint32` |       | Error code uint8                                              |
| `sync`          | 3       | `bool`   |       | Synchronization status                                        |
| `sorting`       | 4       | `bool`   |       | Sorting status                                                |
| `trigger`       | 5       | `bool`   |       | Trigger status                                                |
| `configChange`  | 6       | `bool`   |       | Configuration change status                                   |
| `dataModified`  | 7       | `bool`   |       | Data modification status                                      |
| `timeQuality`   | 8       | `uint32` |       | Time quality uint8                                            |
| `unlockedTime`  | 9       | `uint32` |       | Unlocked time uint8                                           |
| `triggerReason` | 10      | `uint32` |       | Trigger reason uint8                                          |






<!-- Created by: Proto Diagram Tool -->
<!-- https://github.com/GoogleCloudPlatform/proto-gen-md-diagrams -->
