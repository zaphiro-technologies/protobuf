# Package: task.v1

<div class="comment"><span><!-- markdownlint-disable --></span><br/></div>

## Imports

| Import | Description |
|--------|-------------|



## Options

| Name       | Value     | Description |
|------------|-----------|-------------|
| go_package | ./task/v1 |             |



## Enum: TaskType
<div style="font-size: 12px; margin-top: -10px;" class="fqn">FQN: task.v1.TaskType</div>

<div class="comment"><span></span><br/></div>

| Name                  | Ordinal | Description |
|-----------------------|---------|-------------|
| TASK_TYPE_UNSPECIFIED | 0       |             |
| TASK_TYPE_COLLECTION  | 1       |             |
| TASK_TYPE_TOPOLOGY    | 2       |             |
| TASK_TYPE_STATE       | 3       |             |
| TASK_TYPE_FAULT       | 4       |             |


## Enum: NotificationType
<div style="font-size: 12px; margin-top: -10px;" class="fqn">FQN: task.v1.NotificationType</div>

<div class="comment"><span></span><br/></div>

| Name                             | Ordinal | Description |
|----------------------------------|---------|-------------|
| NOTIFICATION_TYPE_UNSPECIFIED    | 0       |             |
| NOTIFICATION_TYPE_DATA_COMPLETE  | 1       |             |
| NOTIFICATION_TYPE_DATA_TIMEOUT_1 | 2       |             |
| NOTIFICATION_TYPE_DATA_TIMEOUT_2 | 3       |             |



### TaskType Diagram

```mermaid
classDiagram
direction LR
%% 

class TaskType{
  <<enumeration>>
  TASK_TYPE_UNSPECIFIED
  TASK_TYPE_COLLECTION
  TASK_TYPE_TOPOLOGY
  TASK_TYPE_STATE
  TASK_TYPE_FAULT
}
```
### NotificationType Diagram

```mermaid
classDiagram
direction LR
%% 

class NotificationType{
  <<enumeration>>
  NOTIFICATION_TYPE_UNSPECIFIED
  NOTIFICATION_TYPE_DATA_COMPLETE
  NOTIFICATION_TYPE_DATA_TIMEOUT_1
  NOTIFICATION_TYPE_DATA_TIMEOUT_2
}
```
### Task Diagram

```mermaid
classDiagram
direction LR

%% 

class Task {
  + int64 createdAt
  + string id
  + Optional~string~ measurementID
  + TaskType taskType
  + Optional~int64~ timestampID
}
Task --> `TaskType`

```
### Notification Diagram

```mermaid
classDiagram
direction LR

%% 

class Notification {
  + int64 createdAt
  + string id
  + Optional~string~ measurementID
  + string message
  + NotificationType notificationType
  + Optional~int64~ timestampID
}
Notification --> `NotificationType`

```

## Message: Task
<div style="font-size: 12px; margin-top: -10px;" class="fqn">FQN: task.v1.Task</div>

<div class="comment"><span></span><br/></div>

| Field         | Ordinal | Type     | Label    | Description |
|---------------|---------|----------|----------|-------------|
| createdAt     | 3       | int64    |          |             |
| id            | 1       | string   |          |             |
| measurementID | 4       | string   | Optional |             |
| taskType      | 2       | TaskType |          |             |
| timestampID   | 5       | int64    | Optional |             |


## Message: Notification
<div style="font-size: 12px; margin-top: -10px;" class="fqn">FQN: task.v1.Notification</div>

<div class="comment"><span></span><br/></div>

| Field            | Ordinal | Type             | Label    | Description |
|------------------|---------|------------------|----------|-------------|
| createdAt        | 3       | int64            |          |             |
| id               | 1       | string           |          |             |
| measurementID    | 5       | string           | Optional |             |
| message          | 4       | string           |          |             |
| notificationType | 2       | NotificationType |          |             |
| timestampID      | 6       | int64            | Optional |             |




<!-- Created by: Proto Diagram Tool -->
<!-- https://github.com/GoogleCloudPlatform/proto-gen-md-diagrams -->
