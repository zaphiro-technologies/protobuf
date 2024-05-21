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
  + string Id
  + List~string~ TerminalIds
  + List~string~ ConnectivityNodeIds
  + List~string~ BranchIds
  + string ConnectivityNodeContainerId
  + string BaseVoltageId
  + double BaseVoltage
}

```
### BranchConnection Diagram

```mermaid
classDiagram
direction LR

%% A wrapper in order to use an array in the map in TopologicalIsland.
%% * It represents the connections between the TopologicalNodes in the TopologicalIsland.
%% 

class BranchConnection {
  + List~string~ TopologicalNodeIds
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
  + string Id
  + List~TopologicalNode~ TopologicalNodes
}
TopologicalIsland --> `TopologicalNode`

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
  + List~TopologicalIsland~ topologicalIslands
  + Map~string,  BranchConnection~ BranchConnections
}
ComputedTopology --> `TopologicalIsland`
ComputedTopology .. ` BranchConnection`

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
| `Id`                          | 1       | `string` |          | The id of the TopologicalNode.                                   |
| `TerminalIds`                 | 2       | `string` | Repeated | The list of Terminal ids in the TopologicalNode.                 |
| `ConnectivityNodeIds`         | 3       | `string` | Repeated | The list of ConnectivityNode ids in the TopologicalNode.         |
| `BranchIds`                   | 4       | `string` | Repeated | The list of Branch ids in the TopologicalNode.                   |
| `ConnectivityNodeContainerId` | 5       | `string` |          | The id of the ConnectivityNodeContainer in the TopologicalNode.  |
| `BaseVoltageId`               | 6       | `string` |          | The id of the BaseVoltage in the TopologicalNode.                |
| `BaseVoltage`                 | 7       | `double` |          | The BaseVoltage in the TopologicalNode.                          |




## Message: BranchConnection

**FQN**: zaphiro.grid.v1.BranchConnection

A wrapper in order to use an array in the map in TopologicalIsland.
* It represents the connections between the TopologicalNodes in the TopologicalIsland.



| Field                | Ordinal | Type     | Label    | Description                                               |
|----------------------|---------|----------|----------|-----------------------------------------------------------|
| `TopologicalNodeIds` | 1       | `string` | Repeated | The list of TopologicalNode ids connected to the Branch.  |




## Message: TopologicalIsland

**FQN**: zaphiro.grid.v1.TopologicalIsland

A topology Island information.
* To be used in ComputedTopology message.



| Field              | Ordinal | Type              | Label    | Description                                            |
|--------------------|---------|-------------------|----------|--------------------------------------------------------|
| `Id`               | 1       | `string`          |          | The id of the TopologicalIsland.                       |
| `TopologicalNodes` | 2       | `TopologicalNode` | Repeated | The list of TopologicalNode in the TopologicalIsland.  |




## Message: ComputedTopology

**FQN**: zaphiro.grid.v1.ComputedTopology

A processed topology information.
Headers used in rabbitMQ:
* `id` (string): id of the `Topology`
* `type` (string): always `ComputedTopology`
* `producerId` (string): the id of the producer (e.g. a PMU) linked to the dataset.
* `timestampId` (int64): related measurement Unix msec timestamp (if any)
* `subnetworkId` (string): the sub network id for which the topology was computed



| Field                | Ordinal | Type                       | Label    | Description                                                                                         |
|----------------------|---------|----------------------------|----------|-----------------------------------------------------------------------------------------------------|
| `eqId`               | 1       | `string`                   |          | The id of the EQ file used.                                                                         |
| `topologicalIslands` | 2       | `TopologicalIsland`        | Repeated | The map of TopologicalIslands in the Topology.                                                      |
| `BranchConnections`  | 3       | `string, BranchConnection` | Map      | All the connections between the TopologicalNodes in the TopologicalIsland. TN id -> TN connection.  |






<!-- Created by: Proto Diagram Tool -->
<!-- https://github.com/GoogleCloudPlatform/proto-gen-md-diagrams -->
