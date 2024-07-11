# Package: zaphiro.grid.v1

<!-- markdownlint-disable -->
Messages to support device event detection in the platform. Device events are sub classes of Events.



## Imports

| Import                      | Description |
|-----------------------------|-------------|
| zaphiro/grid/v1/event.proto |             |



## Options

| Name       | Value     | Description |
|------------|-----------|-------------|
| go_package | ./grid/v1 |             |




### DeviceEvent Diagram

```mermaid
classDiagram
direction LR

%% A device event.
%% Headers used in rabbitMQ:
%% * `id`: id of the `Event`
%% * `type`: always `Event` - used for routing.
%% * `eventType`: the specific type of `DeviceEvent`, this is required in addition 
%%  to `type` for de-serialization of the messages.
%% * `sourceId`: the id of the source (e.g. a PMU) that generated the event.
%% * `timestampId`: related measurement Unix msec timestamp (if any)
%% 

class DeviceEvent {
  + Event event
  + string deviceID
  + string substationID
  + Optional~double~ value
  + Optional~double~ referenceLimit
  + Optional~string~ code
}
DeviceEvent --> `Event`

```
### CommunicationError Diagram

```mermaid
classDiagram
direction LR

%% 

class CommunicationError {
  + DeviceEvent event
}
CommunicationError --> `DeviceEvent`

```
### TimeQuality Diagram

```mermaid
classDiagram
direction LR

%% 

class TimeQuality {
  + DeviceEvent event
}
TimeQuality --> `DeviceEvent`

```
### SyncStatus Diagram

```mermaid
classDiagram
direction LR

%% 

class SyncStatus {
  + DeviceEvent event
}
SyncStatus --> `DeviceEvent`

```
### Power Diagram

```mermaid
classDiagram
direction LR

%% 

class Power {
  + DeviceEvent event
}
Power --> `DeviceEvent`

```
### Config Diagram

```mermaid
classDiagram
direction LR

%% 

class Config {
  + DeviceEvent event
}
Config --> `DeviceEvent`

```
### Trigger Diagram

```mermaid
classDiagram
direction LR

%% 

class Trigger {
  + DeviceEvent event
}
Trigger --> `DeviceEvent`

```
### DataError Diagram

```mermaid
classDiagram
direction LR

%% 

class DataError {
  + DeviceEvent event
}
DataError --> `DeviceEvent`

```

## Message: DeviceEvent

**FQN**: zaphiro.grid.v1.DeviceEvent

A device event.
Headers used in rabbitMQ:
* `id`: id of the `Event`
* `type`: always `Event` - used for routing.
* `eventType`: the specific type of `DeviceEvent`, this is required in addition 
 to `type` for de-serialization of the messages.
* `sourceId`: the id of the source (e.g. a PMU) that generated the event.
* `timestampId`: related measurement Unix msec timestamp (if any)



| Field            | Ordinal | Type     | Label    | Description                                               |
|------------------|---------|----------|----------|-----------------------------------------------------------|
| `event`          | 1       | `Event`  |          | The base event message                                    |
| `deviceID`       | 2       | `string` |          | The ID of the device where the event occurred.            |
| `substationID`   | 3       | `string` |          | The ID of the substation where the event occurred.        |
| `value`          | 4       | `double` | Optional | The measured / estimated value in relation to the event.  |
| `referenceLimit` | 5       | `double` | Optional | The reference limit or expected value.                    |
| `code`           | 6       | `string` | Optional | The device event code (or the mapped string)              |




## Message: CommunicationError

**FQN**: zaphiro.grid.v1.CommunicationError




| Field   | Ordinal | Type          | Label | Description                    |
|---------|---------|---------------|-------|--------------------------------|
| `event` | 1       | `DeviceEvent` |       | The base device event message  |




## Message: TimeQuality

**FQN**: zaphiro.grid.v1.TimeQuality




| Field   | Ordinal | Type          | Label | Description                    |
|---------|---------|---------------|-------|--------------------------------|
| `event` | 1       | `DeviceEvent` |       | The base device event message  |




## Message: SyncStatus

**FQN**: zaphiro.grid.v1.SyncStatus




| Field   | Ordinal | Type          | Label | Description                    |
|---------|---------|---------------|-------|--------------------------------|
| `event` | 1       | `DeviceEvent` |       | The base device event message  |




## Message: Power

**FQN**: zaphiro.grid.v1.Power




| Field   | Ordinal | Type          | Label | Description                    |
|---------|---------|---------------|-------|--------------------------------|
| `event` | 1       | `DeviceEvent` |       | The base device event message  |




## Message: Config

**FQN**: zaphiro.grid.v1.Config




| Field   | Ordinal | Type          | Label | Description                    |
|---------|---------|---------------|-------|--------------------------------|
| `event` | 1       | `DeviceEvent` |       | The base device event message  |




## Message: Trigger

**FQN**: zaphiro.grid.v1.Trigger




| Field   | Ordinal | Type          | Label | Description                    |
|---------|---------|---------------|-------|--------------------------------|
| `event` | 1       | `DeviceEvent` |       | The base device event message  |




## Message: DataError

**FQN**: zaphiro.grid.v1.DataError




| Field   | Ordinal | Type          | Label | Description                    |
|---------|---------|---------------|-------|--------------------------------|
| `event` | 1       | `DeviceEvent` |       | The base device event message  |






<!-- Created by: Proto Diagram Tool -->
<!-- https://github.com/GoogleCloudPlatform/proto-gen-md-diagrams -->
