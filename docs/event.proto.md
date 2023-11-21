# Package: grid.v1

<div class="comment"><span><!-- markdownlint-disable --></span><br/><span>Messages to support event detection in the platform.</span><br/><span>The event detected may be originated from different sources: devices (e.g. a PMU, RTU), services (e.g. state estimator), or an external service (e.g. SCADA).</span><br/><span></span><br/></div>

## Imports

| Import | Description |
|--------|-------------|



## Options

| Name       | Value     | Description |
|------------|-----------|-------------|
| go_package | ./grid/v1 |             |



## Enum: EventStatus
<div style="font-size: 12px; margin-top: -10px;" class="fqn">FQN: grid.v1.EventStatus</div>

<div class="comment"><span>The collection of Event Status defined so far.</span><br/></div>

| Name                     | Ordinal | Description                                                                     |
|--------------------------|---------|---------------------------------------------------------------------------------|
| EVENT_STATUS_UNSPECIFIED | 0       | No status defined                                                               |
| EVENT_STATUS_ACTIVE      | 1       | Event is still active                                                           |
| EVENT_STATUS_COMPLETE    | 2       | Event is completed                                                              |
| EVENT_STATUS_UNKNOWN     | 3       | Information available don't allow us to know if the even is active or complete  |


## Enum: EventSourceType
<div style="font-size: 12px; margin-top: -10px;" class="fqn">FQN: grid.v1.EventSourceType</div>

<div class="comment"><span></span><br/></div>

| Name                          | Ordinal | Description                                                                           |
|-------------------------------|---------|---------------------------------------------------------------------------------------|
| EVENT_SOURCE_UNSPECIFIED      | 0       | No source type defined                                                                |
| EVENT_SOURCE_DEVICE           | 1       | The source of the event was a device (e.g. PMU)                                       |
| EVENT_SOURCE_SERVICE          | 2       | The source of the event was a service (e.g. state estimator)                          |
| EVENT_SOURCE_EXTERNAL_SERVICE | 3       | The source of the event was a service external to SynchroGuard platform (e.g. SCADA)  |



### EventStatus Diagram

```mermaid
classDiagram
direction LR
%% The collection of Event Status defined so far.

class EventStatus{
  <<enumeration>>
  EVENT_STATUS_UNSPECIFIED
  EVENT_STATUS_ACTIVE
  EVENT_STATUS_COMPLETE
  EVENT_STATUS_UNKNOWN
}
```
### EventSourceType Diagram

```mermaid
classDiagram
direction LR
%% 

class EventSourceType{
  <<enumeration>>
  EVENT_SOURCE_UNSPECIFIED
  EVENT_SOURCE_DEVICE
  EVENT_SOURCE_SERVICE
  EVENT_SOURCE_EXTERNAL_SERVICE
}
```
### Event Diagram

```mermaid
classDiagram
direction LR

%% A generic event.
%% Headers used in rabbitMQ (only if not sent as part of `DataSet`):
%% * `id`: id of the `Event`
%% * `type`: always `Event`
%% * `sourceId`: the id of the source (e.g. a PMU) that generated the event.
%% * `timestampId`: related measurement Unix msec timestamp (if any)
%% 

class Event {
  + Optional~int64~ detectedAt
  + string message
  + Optional~int64~ occurredAt
  + EventSourceType sourceType
  + Optional~EventStatus~ status
  + Optional~uint64~ value
}
Event --> `EventSourceType`
Event --> `EventStatus`

```

## Message: Event
<div style="font-size: 12px; margin-top: -10px;" class="fqn">FQN: grid.v1.Event</div>

<div class="comment"><span>A generic event.</span><br/><span>Headers used in rabbitMQ (only if not sent as part of `DataSet`):</span><br/><span>* `id`: id of the `Event`</span><br/><span>* `type`: always `Event`</span><br/><span>* `sourceId`: the id of the source (e.g. a PMU) that generated the event.</span><br/><span>* `timestampId`: related measurement Unix msec timestamp (if any)</span><br/><span></span><br/></div>

| Field      | Ordinal | Type            | Label    | Description                                                |
|------------|---------|-----------------|----------|------------------------------------------------------------|
| detectedAt | 3       | int64           | Optional | The time of detection of the event (Unix msec timestamp).  |
| message    | 5       | string          |          | Event message.                                             |
| occurredAt | 2       | int64           | Optional | The time of occurency of the event (Unix msec timestamp).  |
| sourceType | 1       | EventSourceType |          | The type of data see `DataType` enum.                      |
| status     | 6       | EventStatus     | Optional | The status of the event.                                   |
| value      | 4       | uint64          | Optional | The data value casted to uint64.                           |






<!-- Created by: Proto Diagram Tool -->
<!-- https://github.com/GoogleCloudPlatform/proto-gen-md-diagrams -->
