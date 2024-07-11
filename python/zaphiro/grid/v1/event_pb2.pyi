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

from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class EventStatus(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    EVENT_STATUS_UNSPECIFIED: _ClassVar[EventStatus]
    EVENT_STATUS_STARTED: _ClassVar[EventStatus]
    EVENT_STATUS_IN_PROGRESS: _ClassVar[EventStatus]
    EVENT_STATUS_ENDED: _ClassVar[EventStatus]
    EVENT_STATUS_UNKNOWN: _ClassVar[EventStatus]

class EventSourceType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    EVENT_SOURCE_UNSPECIFIED: _ClassVar[EventSourceType]
    EVENT_SOURCE_DEVICE: _ClassVar[EventSourceType]
    EVENT_SOURCE_SERVICE: _ClassVar[EventSourceType]
    EVENT_SOURCE_EXTERNAL_SERVICE: _ClassVar[EventSourceType]
EVENT_STATUS_UNSPECIFIED: EventStatus
EVENT_STATUS_STARTED: EventStatus
EVENT_STATUS_IN_PROGRESS: EventStatus
EVENT_STATUS_ENDED: EventStatus
EVENT_STATUS_UNKNOWN: EventStatus
EVENT_SOURCE_UNSPECIFIED: EventSourceType
EVENT_SOURCE_DEVICE: EventSourceType
EVENT_SOURCE_SERVICE: EventSourceType
EVENT_SOURCE_EXTERNAL_SERVICE: EventSourceType

class Event(_message.Message):
    __slots__ = ("Id", "sourceId", "sourceType", "occurredAt", "detectedAt", "message", "status")
    ID_FIELD_NUMBER: _ClassVar[int]
    SOURCEID_FIELD_NUMBER: _ClassVar[int]
    SOURCETYPE_FIELD_NUMBER: _ClassVar[int]
    OCCURREDAT_FIELD_NUMBER: _ClassVar[int]
    DETECTEDAT_FIELD_NUMBER: _ClassVar[int]
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    Id: str
    sourceId: str
    sourceType: EventSourceType
    occurredAt: int
    detectedAt: int
    message: str
    status: EventStatus
    def __init__(self, Id: _Optional[str] = ..., sourceId: _Optional[str] = ..., sourceType: _Optional[_Union[EventSourceType, str]] = ..., occurredAt: _Optional[int] = ..., detectedAt: _Optional[int] = ..., message: _Optional[str] = ..., status: _Optional[_Union[EventStatus, str]] = ...) -> None: ...
