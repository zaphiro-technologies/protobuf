# Package: zaphiro.grid.v1

Copyright 2024 Zaphiro Technologies Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0 Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License. <!-- markdownlint-disable -->
Messages to support event detection in the platform.
The event detected may be originated from different sources: devices (e.g. a
PMU, RTU), services (e.g. state estimator), or an external service (e.g. SCADA).



## Imports

| Import | Description |
|--------|-------------|



## Options

| Name       | Value     | Description |
|------------|-----------|-------------|
| go_package | ./grid/v1 |             |



## Enum: EventStatus

**FQN**: zaphiro.grid.v1.EventStatus

The collection of Event Status defined so far.


| Name                       | Ordinal | Description                                                                     |
|----------------------------|---------|---------------------------------------------------------------------------------|
| `EVENT_STATUS_UNSPECIFIED` | 0       | No status defined                                                               |
| `EVENT_STATUS_STARTED`     | 1       | Event started                                                                   |
| `EVENT_STATUS_IN_PROGRESS` | 2       | Event is still active                                                           |
| `EVENT_STATUS_ENDED`       | 3       | Event ended                                                                     |
| `EVENT_STATUS_UNKNOWN`     | 4       | Information available don't allow us to know if the even is active or complete  |


## Enum: EventSourceType

**FQN**: zaphiro.grid.v1.EventSourceType




| Name                            | Ordinal | Description                                                                           |
|---------------------------------|---------|---------------------------------------------------------------------------------------|
| `EVENT_SOURCE_UNSPECIFIED`      | 0       | No source type defined                                                                |
| `EVENT_SOURCE_DEVICE`           | 1       | The source of the event was a device (e.g. PMU)                                       |
| `EVENT_SOURCE_SERVICE`          | 2       | The source of the event was a service (e.g. state estimator)                          |
| `EVENT_SOURCE_EXTERNAL_SERVICE` | 3       | The source of the event was a service external to SynchroGuard platform (e.g. SCADA)  |
| `EVENT_SOURCE_TEST_SERVICE`     | 4       | The source of the event was a service in test mode.                                   |



### EventStatus Diagram

```mermaid
classDiagram
direction LR
%% The collection of Event Status defined so far.

class EventStatus{
  <<enumeration>>
  EVENT_STATUS_UNSPECIFIED
  EVENT_STATUS_STARTED
  EVENT_STATUS_IN_PROGRESS
  EVENT_STATUS_ENDED
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
  EVENT_SOURCE_TEST_SERVICE
}
```
### Event Diagram

```mermaid
classDiagram
direction LR

%% A generic event.
%% Headers used in rabbitMQ (only if not sent as part of `DataSet`):
%% * `id` (string): id of the `Event`
%% * `type` (string): always `Event` - used for routing.
%% * `eventType` (string): the specific type of `Event`, this is required in
%% addition to `type` for de-serialization of the messages.
%% * `sourceId` (string): DEPRECATED: use: producerId. the id of the source (e.g. a PMU) that generated the event
%% * `producerId` (string): the id of the producer (e.g. a PMU) that generated the event
%% * `sourceType` (string): the Event source type
%% event.
%% * `timestampId` (int64): related measurement Unix msec timestamp (if any)
%% 

class Event {
  + string Id
  + string sourceId
  + EventSourceType sourceType
  + int64 occurredAt
  + Optional~int64~ detectedAt
  + string message
  + Optional~EventStatus~ status
}
Event --> `EventSourceType`
Event --> `EventStatus`

```

## Message: Event

**FQN**: zaphiro.grid.v1.Event

A generic event.
Headers used in rabbitMQ (only if not sent as part of `DataSet`):
* `id` (string): id of the `Event`
* `type` (string): always `Event` - used for routing.
* `eventType` (string): the specific type of `Event`, this is required in
addition to `type` for de-serialization of the messages.
* `sourceId` (string): DEPRECATED: use: producerId. the id of the source (e.g. a PMU) that generated the event
* `producerId` (string): the id of the producer (e.g. a PMU) that generated the event
* `sourceType` (string): the Event source type
event.
* `timestampId` (int64): related measurement Unix msec timestamp (if any)



| Field        | Ordinal | Type              | Label    | Description                                                                                         |
|--------------|---------|-------------------|----------|-----------------------------------------------------------------------------------------------------|
| `Id`         | 1       | `string`          |          | The uuid of the event.                                                                              |
| `sourceId`   | 2       | `string`          |          | The id of the source (e.g. a PMU) that generated the event.                                         |
| `sourceType` | 3       | `EventSourceType` |          | The type of data see `DataType` enum.                                                               |
| `occurredAt` | 4       | `int64`           |          | The time of occurency of the event (Unix msec timestamp) usually is the same value as timestampId.  |
| `detectedAt` | 5       | `int64`           | Optional | The time of detection of the event (Unix msec timestamp).                                           |
| `message`    | 6       | `string`          |          | Event message.                                                                                      |
| `status`     | 7       | `EventStatus`     | Optional | The status of the event.                                                                            |






<!-- Created by: Proto Diagram Tool -->
<!-- https://github.com/GoogleCloudPlatform/proto-gen-md-diagrams -->
