# Copyright 2024 Zaphiro Technologies
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

from google.protobuf import any_pb2 as _any_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class TaskType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    TASK_TYPE_UNSPECIFIED: _ClassVar[TaskType]
    TASK_TYPE_COLLECTION: _ClassVar[TaskType]
    TASK_TYPE_TOPOLOGY: _ClassVar[TaskType]
    TASK_TYPE_STATE: _ClassVar[TaskType]
    TASK_TYPE_FAULT: _ClassVar[TaskType]

class NotificationType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    NOTIFICATION_TYPE_UNSPECIFIED: _ClassVar[NotificationType]
    NOTIFICATION_TYPE_DATA_COMPLETE: _ClassVar[NotificationType]
    NOTIFICATION_TYPE_DATA_TIMEOUT_SHORT: _ClassVar[NotificationType]
    NOTIFICATION_TYPE_DATA_TIMEOUT_MEDIUM: _ClassVar[NotificationType]
    NOTIFICATION_TYPE_TRIGGER: _ClassVar[NotificationType]
    NOTIFICATION_TYPE_TOPOLOGY_COMPUTED: _ClassVar[NotificationType]
TASK_TYPE_UNSPECIFIED: TaskType
TASK_TYPE_COLLECTION: TaskType
TASK_TYPE_TOPOLOGY: TaskType
TASK_TYPE_STATE: TaskType
TASK_TYPE_FAULT: TaskType
NOTIFICATION_TYPE_UNSPECIFIED: NotificationType
NOTIFICATION_TYPE_DATA_COMPLETE: NotificationType
NOTIFICATION_TYPE_DATA_TIMEOUT_SHORT: NotificationType
NOTIFICATION_TYPE_DATA_TIMEOUT_MEDIUM: NotificationType
NOTIFICATION_TYPE_TRIGGER: NotificationType
NOTIFICATION_TYPE_TOPOLOGY_COMPUTED: NotificationType

class Parameter(_message.Message):
    __slots__ = ("name", "value")
    NAME_FIELD_NUMBER: _ClassVar[int]
    VALUE_FIELD_NUMBER: _ClassVar[int]
    name: str
    value: _any_pb2.Any
    def __init__(self, name: _Optional[str] = ..., value: _Optional[_Union[_any_pb2.Any, _Mapping]] = ...) -> None: ...

class Task(_message.Message):
    __slots__ = ("taskType", "createdAt", "parameters")
    TASKTYPE_FIELD_NUMBER: _ClassVar[int]
    CREATEDAT_FIELD_NUMBER: _ClassVar[int]
    PARAMETERS_FIELD_NUMBER: _ClassVar[int]
    taskType: TaskType
    createdAt: int
    parameters: _containers.RepeatedCompositeFieldContainer[Parameter]
    def __init__(self, taskType: _Optional[_Union[TaskType, str]] = ..., createdAt: _Optional[int] = ..., parameters: _Optional[_Iterable[_Union[Parameter, _Mapping]]] = ...) -> None: ...

class Notification(_message.Message):
    __slots__ = ("notificationType", "createdAt", "message", "parameters")
    NOTIFICATIONTYPE_FIELD_NUMBER: _ClassVar[int]
    CREATEDAT_FIELD_NUMBER: _ClassVar[int]
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    PARAMETERS_FIELD_NUMBER: _ClassVar[int]
    notificationType: NotificationType
    createdAt: int
    message: str
    parameters: _containers.RepeatedCompositeFieldContainer[Parameter]
    def __init__(self, notificationType: _Optional[_Union[NotificationType, str]] = ..., createdAt: _Optional[int] = ..., message: _Optional[str] = ..., parameters: _Optional[_Iterable[_Union[Parameter, _Mapping]]] = ...) -> None: ...
