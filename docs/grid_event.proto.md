# Package: grid.v1

<div class="comment"><span><!-- markdownlint-disable --></span><br/><span>Messages to support grid event detection in the platform. Grid events are sub classes of Events.</span><br/><span></span><br/></div>

## Imports

| Import              | Description |
|---------------------|-------------|
| grid/v1/event.proto |             |
| grid/v1/fault.proto |             |



## Options

| Name       | Value     | Description |
|------------|-----------|-------------|
| go_package | ./grid/v1 |             |




### GridEvent Diagram

```mermaid
classDiagram
direction LR

%% A grid event.
%% Headers used in rabbitMQ (only if not sent as part of `DataSet`):
%% * `id`: id of the `Event`
%% * `type`: always `Event`
%% * `sourceId`: the id of the source (e.g. a PMU) that generated the event.
%% * `timestampId`: related measurement Unix msec timestamp (if any)
%% 

class GridEvent {
  + string componentID
  + int64 endedAt
  + Event event
  + float nominalValue
  + float percentage
  + float probability
  + int64 startedAt
  + Optional~string~ substationID
}
GridEvent --> `Event`

```
### VoltageEvent Diagram

```mermaid
classDiagram
direction LR

%% 

class VoltageEvent {
  + GridEvent event
  + Optional~PhaseCode~ phaseCode
}
VoltageEvent --> `GridEvent`
VoltageEvent --> `PhaseCode`

```
### CurrentEvent Diagram

```mermaid
classDiagram
direction LR

%% 

class CurrentEvent {
  + GridEvent event
}
CurrentEvent --> `GridEvent`

```
### PhaseEvent Diagram

```mermaid
classDiagram
direction LR

%% 

class PhaseEvent {
  + GridEvent event
}
PhaseEvent --> `GridEvent`

```
### FrequencyEvent Diagram

```mermaid
classDiagram
direction LR

%% 

class FrequencyEvent {
  + GridEvent event
}
FrequencyEvent --> `GridEvent`

```
### LineCongestion Diagram

```mermaid
classDiagram
direction LR

%% 

class LineCongestion {
  + CurrentEvent event
}
LineCongestion --> `CurrentEvent`

```
### TransformerCongestion Diagram

```mermaid
classDiagram
direction LR

%% 

class TransformerCongestion {
  + CurrentEvent event
}
TransformerCongestion --> `CurrentEvent`

```
### VoltageUnbalance Diagram

```mermaid
classDiagram
direction LR

%% 

class VoltageUnbalance {
  + VoltageEvent event
}
VoltageUnbalance --> `VoltageEvent`

```
### VoltageDip Diagram

```mermaid
classDiagram
direction LR

%% 

class VoltageDip {
  + VoltageEvent event
}
VoltageDip --> `VoltageEvent`

```
### VoltageInterruption Diagram

```mermaid
classDiagram
direction LR

%% 

class VoltageInterruption {
  + VoltageEvent event
}
VoltageInterruption --> `VoltageEvent`

```
### VoltageSwell Diagram

```mermaid
classDiagram
direction LR

%% 

class VoltageSwell {
  + VoltageEvent event
}
VoltageSwell --> `VoltageEvent`

```
### VoltageLimit Diagram

```mermaid
classDiagram
direction LR

%% 

class VoltageLimit {
  + VoltageEvent event
}
VoltageLimit --> `VoltageEvent`

```
### VoltageRapidChange Diagram

```mermaid
classDiagram
direction LR

%% 

class VoltageRapidChange {
  + VoltageEvent event
}
VoltageRapidChange --> `VoltageEvent`

```
### OverFrequency Diagram

```mermaid
classDiagram
direction LR

%% 

class OverFrequency {
  + FrequencyEvent event
}
OverFrequency --> `FrequencyEvent`

```
### UnderFrequency Diagram

```mermaid
classDiagram
direction LR

%% 

class UnderFrequency {
  + FrequencyEvent event
}
UnderFrequency --> `FrequencyEvent`

```
### SteadyOscillation Diagram

```mermaid
classDiagram
direction LR

%% 

class SteadyOscillation {
  + PhaseEvent event
}
SteadyOscillation --> `PhaseEvent`

```
### TransientOscillation Diagram

```mermaid
classDiagram
direction LR

%% 

class TransientOscillation {
  + PhaseEvent event
}
TransientOscillation --> `PhaseEvent`

```

## Message: GridEvent
<div style="font-size: 12px; margin-top: -10px;" class="fqn">FQN: grid.v1.GridEvent</div>

<div class="comment"><span>A grid event.</span><br/><span>Headers used in rabbitMQ (only if not sent as part of `DataSet`):</span><br/><span>* `id`: id of the `Event`</span><br/><span>* `type`: always `Event`</span><br/><span>* `sourceId`: the id of the source (e.g. a PMU) that generated the event.</span><br/><span>* `timestampId`: related measurement Unix msec timestamp (if any)</span><br/><span></span><br/></div>

| Field        | Ordinal | Type   | Label    | Description                                                                                                                                                           |
|--------------|---------|--------|----------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| componentID  | 2       | string |          | The ID of the component where the event occurred.                                                                                                                     |
| endedAt      | 5       | int64  |          | The timestamp where the event ended.                                                                                                                                  |
| event        | 1       | Event  |          | The base event message                                                                                                                                                |
| nominalValue | 6       | float  |          | The measured / estimated value in relation to the event (e.g. in the case of a `VoltageEvent` is the voltage, in the case of a `CurrentEvent` is the current event).  |
| percentage   | 7       | float  |          | The percentage reached by the value compared to the reference limit or expected value.                                                                                |
| probability  | 8       | float  |          | The probability that the event actually occurred.                                                                                                                     |
| startedAt    | 4       | int64  |          | The timestamp where the event started (should be equal to timestampId in header).                                                                                     |
| substationID | 3       | string | Optional | The ID of the substation where the event occurred.                                                                                                                    |




## Message: VoltageEvent
<div style="font-size: 12px; margin-top: -10px;" class="fqn">FQN: grid.v1.VoltageEvent</div>

<div class="comment"><span></span><br/></div>

| Field     | Ordinal | Type      | Label    | Description |
|-----------|---------|-----------|----------|-------------|
| event     | 1       | GridEvent |          |             |
| phaseCode | 2       | PhaseCode | Optional |             |




## Message: CurrentEvent
<div style="font-size: 12px; margin-top: -10px;" class="fqn">FQN: grid.v1.CurrentEvent</div>

<div class="comment"><span></span><br/></div>

| Field | Ordinal | Type      | Label | Description |
|-------|---------|-----------|-------|-------------|
| event | 1       | GridEvent |       |             |




## Message: PhaseEvent
<div style="font-size: 12px; margin-top: -10px;" class="fqn">FQN: grid.v1.PhaseEvent</div>

<div class="comment"><span></span><br/></div>

| Field | Ordinal | Type      | Label | Description |
|-------|---------|-----------|-------|-------------|
| event | 1       | GridEvent |       |             |




## Message: FrequencyEvent
<div style="font-size: 12px; margin-top: -10px;" class="fqn">FQN: grid.v1.FrequencyEvent</div>

<div class="comment"><span></span><br/></div>

| Field | Ordinal | Type      | Label | Description |
|-------|---------|-----------|-------|-------------|
| event | 1       | GridEvent |       |             |




## Message: LineCongestion
<div style="font-size: 12px; margin-top: -10px;" class="fqn">FQN: grid.v1.LineCongestion</div>

<div class="comment"><span></span><br/></div>

| Field | Ordinal | Type         | Label | Description |
|-------|---------|--------------|-------|-------------|
| event | 1       | CurrentEvent |       |             |




## Message: TransformerCongestion
<div style="font-size: 12px; margin-top: -10px;" class="fqn">FQN: grid.v1.TransformerCongestion</div>

<div class="comment"><span></span><br/></div>

| Field | Ordinal | Type         | Label | Description |
|-------|---------|--------------|-------|-------------|
| event | 1       | CurrentEvent |       |             |




## Message: VoltageUnbalance
<div style="font-size: 12px; margin-top: -10px;" class="fqn">FQN: grid.v1.VoltageUnbalance</div>

<div class="comment"><span></span><br/></div>

| Field | Ordinal | Type         | Label | Description |
|-------|---------|--------------|-------|-------------|
| event | 1       | VoltageEvent |       |             |




## Message: VoltageDip
<div style="font-size: 12px; margin-top: -10px;" class="fqn">FQN: grid.v1.VoltageDip</div>

<div class="comment"><span></span><br/></div>

| Field | Ordinal | Type         | Label | Description |
|-------|---------|--------------|-------|-------------|
| event | 1       | VoltageEvent |       |             |




## Message: VoltageInterruption
<div style="font-size: 12px; margin-top: -10px;" class="fqn">FQN: grid.v1.VoltageInterruption</div>

<div class="comment"><span></span><br/></div>

| Field | Ordinal | Type         | Label | Description |
|-------|---------|--------------|-------|-------------|
| event | 1       | VoltageEvent |       |             |




## Message: VoltageSwell
<div style="font-size: 12px; margin-top: -10px;" class="fqn">FQN: grid.v1.VoltageSwell</div>

<div class="comment"><span></span><br/></div>

| Field | Ordinal | Type         | Label | Description |
|-------|---------|--------------|-------|-------------|
| event | 1       | VoltageEvent |       |             |




## Message: VoltageLimit
<div style="font-size: 12px; margin-top: -10px;" class="fqn">FQN: grid.v1.VoltageLimit</div>

<div class="comment"><span></span><br/></div>

| Field | Ordinal | Type         | Label | Description |
|-------|---------|--------------|-------|-------------|
| event | 1       | VoltageEvent |       |             |




## Message: VoltageRapidChange
<div style="font-size: 12px; margin-top: -10px;" class="fqn">FQN: grid.v1.VoltageRapidChange</div>

<div class="comment"><span></span><br/></div>

| Field | Ordinal | Type         | Label | Description |
|-------|---------|--------------|-------|-------------|
| event | 1       | VoltageEvent |       |             |




## Message: OverFrequency
<div style="font-size: 12px; margin-top: -10px;" class="fqn">FQN: grid.v1.OverFrequency</div>

<div class="comment"><span></span><br/></div>

| Field | Ordinal | Type           | Label | Description |
|-------|---------|----------------|-------|-------------|
| event | 1       | FrequencyEvent |       |             |




## Message: UnderFrequency
<div style="font-size: 12px; margin-top: -10px;" class="fqn">FQN: grid.v1.UnderFrequency</div>

<div class="comment"><span></span><br/></div>

| Field | Ordinal | Type           | Label | Description |
|-------|---------|----------------|-------|-------------|
| event | 1       | FrequencyEvent |       |             |




## Message: SteadyOscillation
<div style="font-size: 12px; margin-top: -10px;" class="fqn">FQN: grid.v1.SteadyOscillation</div>

<div class="comment"><span></span><br/></div>

| Field | Ordinal | Type       | Label | Description |
|-------|---------|------------|-------|-------------|
| event | 1       | PhaseEvent |       |             |




## Message: TransientOscillation
<div style="font-size: 12px; margin-top: -10px;" class="fqn">FQN: grid.v1.TransientOscillation</div>

<div class="comment"><span></span><br/></div>

| Field | Ordinal | Type       | Label | Description |
|-------|---------|------------|-------|-------------|
| event | 1       | PhaseEvent |       |             |






<!-- Created by: Proto Diagram Tool -->
<!-- https://github.com/GoogleCloudPlatform/proto-gen-md-diagrams -->
