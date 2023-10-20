# Package: grid.v1

<div class="comment"><span><!-- markdownlint-disable --></span><br/><span>Messages to support topology data exchange in the platform.</span><br/><span></span><br/></div>

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
%% * `id`: id of the `Topology`
%% * `type`: always `Topology`
%% * `producerId`: the id of the producer (e.g. a PMU) linked to the dataset.
%% * `timestampId`: related measurement Unix msec timestamp (if any)
%% * `subnetworkId`: the sub network id for which the topology was computed
%% 

class Topology {
  + uint64 createdAt
  + bytes ssh
  + bytes tp
}

```

## Message: Topology
<div style="font-size: 12px; margin-top: -10px;" class="fqn">FQN: grid.v1.Topology</div>

<div class="comment"><span>A topology computed information.</span><br/><span>Headers used in rabbitMQ:</span><br/><span>* `id`: id of the `Topology`</span><br/><span>* `type`: always `Topology`</span><br/><span>* `producerId`: the id of the producer (e.g. a PMU) linked to the dataset.</span><br/><span>* `timestampId`: related measurement Unix msec timestamp (if any)</span><br/><span>* `subnetworkId`: the sub network id for which the topology was computed</span><br/><span></span><br/></div>

| Field     | Ordinal | Type   | Label | Description                                                       |
|-----------|---------|--------|-------|-------------------------------------------------------------------|
| createdAt | 1       | uint64 |       | The time of creation of the topology data (Unix msec timestamp).  |
| ssh       | 3       | bytes  |       | The SSH profile file serialized as bytes.                         |
| tp        | 2       | bytes  |       | The TP profile file serialized as bytes.                          |




<!-- Created by: Proto Diagram Tool -->
<!-- https://github.com/GoogleCloudPlatform/proto-gen-md-diagrams -->
