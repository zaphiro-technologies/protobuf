# Package: zaphiro.grid.v1

<!-- markdownlint-disable -->
Messages to support topology data exchange in the platform.



## Imports

| Import | Description |
|--------|-------------|



## Options

| Name       | Value     | Description |
|------------|-----------|-------------|
| go_package | ./grid/v1 |             |




### Topology Diagram

```mermaid
classDiagram
direction LR

%% A topology computed information.
%% Headers used in rabbitMQ:
%% * `id` (string): id of the `Topology`
%% * `type` (string): always `Topology`
%% * `producerId` (string): the id of the producer (e.g. a PMU) linked to the dataset.
%% * `timestampId` (int64): related measurement Unix msec timestamp (if any)
%% * `subnetworkId` (string): the sub network id for which the topology was computed
%% 

class Topology {
  + int64 createdAt
  + bytes tp
}

```
### TopologicalNode Diagram

```mermaid
classDiagram
direction LR

%% A topology Node information.
%% * To be used in ComputedTopology message.
%% 

class TopologicalNode {
  + List~string~ TerminalIds
  + List~string~ ConnectivityNodeIds
  + List~string~ PowerTransferEndIds
  + string ConnectivityNodeContainerId
  + string BaseVoltageId
  + double BaseVoltage
}

```
### TopologicalIsland Diagram

```mermaid
classDiagram
direction LR

%% A topology Island information.
%% * To be used in ComputedTopology message.
%% 

class TopologicalIsland {
  + List~string~ TopologicalNodeIds
}

```
### ComputedTopology Diagram

```mermaid
classDiagram
direction LR

%% A processed topology information.
%% Headers used in rabbitMQ:
%% * `id` (string): id of the `Topology`
%% * `type` (string): always `ComputedTopology`
%% * `producerId` (string): the id of the producer (e.g. a PMU) linked to the dataset.
%% * `timestampId` (int64): related measurement Unix msec timestamp (if any)
%% * `subnetworkId` (string): the sub network id for which the topology was computed
%% 

class ComputedTopology {
  + string eqId
  + Map~string,  TopologicalNode~ topologicalNodes
  + Map~string,  TopologicalIsland~ topologicalIslands
}
ComputedTopology .. ` TopologicalNode`
ComputedTopology .. ` TopologicalIsland`

```

## Message: Topology

**FQN**: zaphiro.grid.v1.Topology

A topology computed information.
Headers used in rabbitMQ:
* `id` (string): id of the `Topology`
* `type` (string): always `Topology`
* `producerId` (string): the id of the producer (e.g. a PMU) linked to the dataset.
* `timestampId` (int64): related measurement Unix msec timestamp (if any)
* `subnetworkId` (string): the sub network id for which the topology was computed



| Field       | Ordinal | Type    | Label | Description                                                       |
|-------------|---------|---------|-------|-------------------------------------------------------------------|
| `createdAt` | 1       | `int64` |       | The time of creation of the topology data (Unix msec timestamp).  |
| `tp`        | 2       | `bytes` |       | The TP profile file serialized as bytes.                          |




## Message: TopologicalNode

**FQN**: zaphiro.grid.v1.TopologicalNode

A topology Node information.
* To be used in ComputedTopology message.



| Field                         | Ordinal | Type     | Label    | Description                                                      |
|-------------------------------|---------|----------|----------|------------------------------------------------------------------|
| `TerminalIds`                 | 1       | `string` | Repeated | The list of Terminal ids in the TopologicalNode.                 |
| `ConnectivityNodeIds`         | 2       | `string` | Repeated | The list of ConnectivityNode ids in the TopologicalNode.         |
| `PowerTransferEndIds`         | 3       | `string` | Repeated | The list of PowerTransferEnd ids in the TopologicalNode.         |
| `ConnectivityNodeContainerId` | 4       | `string` |          | The id of the ConnectivityNodeContainer in the TopologicalNode.  |
| `BaseVoltageId`               | 5       | `string` |          | The id of the BaseVoltage in the TopologicalNode.                |
| `BaseVoltage`                 | 6       | `double` |          | The BaseVoltage in the TopologicalNode.                          |




## Message: TopologicalIsland

**FQN**: zaphiro.grid.v1.TopologicalIsland

A topology Island information.
* To be used in ComputedTopology message.



| Field                | Ordinal | Type     | Label    | Description                                                |
|----------------------|---------|----------|----------|------------------------------------------------------------|
| `TopologicalNodeIds` | 1       | `string` | Repeated | The list of TopologicalNode ids in the TopologicalIsland.  |




## Message: ComputedTopology

**FQN**: zaphiro.grid.v1.ComputedTopology

A processed topology information.
Headers used in rabbitMQ:
* `id` (string): id of the `Topology`
* `type` (string): always `ComputedTopology`
* `producerId` (string): the id of the producer (e.g. a PMU) linked to the dataset.
* `timestampId` (int64): related measurement Unix msec timestamp (if any)
* `subnetworkId` (string): the sub network id for which the topology was computed



| Field                | Ordinal | Type                        | Label | Description                                     |
|----------------------|---------|-----------------------------|-------|-------------------------------------------------|
| `eqId`               | 1       | `string`                    |       | The id of the EQ file used.                     |
| `topologicalNodes`   | 2       | `string, TopologicalNode`   | Map   | The map of TopologicalNodes in the Topology.    |
| `topologicalIslands` | 3       | `string, TopologicalIsland` | Map   | The map of TopologicalIslands in the Topology.  |






<!-- Created by: Proto Diagram Tool -->
<!-- https://github.com/GoogleCloudPlatform/proto-gen-md-diagrams -->
