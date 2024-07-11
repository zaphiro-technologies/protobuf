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

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class Stat(_message.Message):
    __slots__ = ("measuredAt", "error", "sync", "sorting", "trigger", "configChange", "dataModified", "timeQuality", "unlockedTime", "triggerReason")
    MEASUREDAT_FIELD_NUMBER: _ClassVar[int]
    ERROR_FIELD_NUMBER: _ClassVar[int]
    SYNC_FIELD_NUMBER: _ClassVar[int]
    SORTING_FIELD_NUMBER: _ClassVar[int]
    TRIGGER_FIELD_NUMBER: _ClassVar[int]
    CONFIGCHANGE_FIELD_NUMBER: _ClassVar[int]
    DATAMODIFIED_FIELD_NUMBER: _ClassVar[int]
    TIMEQUALITY_FIELD_NUMBER: _ClassVar[int]
    UNLOCKEDTIME_FIELD_NUMBER: _ClassVar[int]
    TRIGGERREASON_FIELD_NUMBER: _ClassVar[int]
    measuredAt: int
    error: int
    sync: bool
    sorting: bool
    trigger: bool
    configChange: bool
    dataModified: bool
    timeQuality: int
    unlockedTime: int
    triggerReason: int
    def __init__(self, measuredAt: _Optional[int] = ..., error: _Optional[int] = ..., sync: bool = ..., sorting: bool = ..., trigger: bool = ..., configChange: bool = ..., dataModified: bool = ..., timeQuality: _Optional[int] = ..., unlockedTime: _Optional[int] = ..., triggerReason: _Optional[int] = ...) -> None: ...
