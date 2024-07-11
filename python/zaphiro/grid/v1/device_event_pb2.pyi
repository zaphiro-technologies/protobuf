from zaphiro.grid.v1 import event_pb2 as _event_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class DeviceEvent(_message.Message):
    __slots__ = ("event", "deviceID", "substationID", "value", "referenceLimit", "code")
    EVENT_FIELD_NUMBER: _ClassVar[int]
    DEVICEID_FIELD_NUMBER: _ClassVar[int]
    SUBSTATIONID_FIELD_NUMBER: _ClassVar[int]
    VALUE_FIELD_NUMBER: _ClassVar[int]
    REFERENCELIMIT_FIELD_NUMBER: _ClassVar[int]
    CODE_FIELD_NUMBER: _ClassVar[int]
    event: _event_pb2.Event
    deviceID: str
    substationID: str
    value: float
    referenceLimit: float
    code: str
    def __init__(self, event: _Optional[_Union[_event_pb2.Event, _Mapping]] = ..., deviceID: _Optional[str] = ..., substationID: _Optional[str] = ..., value: _Optional[float] = ..., referenceLimit: _Optional[float] = ..., code: _Optional[str] = ...) -> None: ...

class CommunicationError(_message.Message):
    __slots__ = ("event",)
    EVENT_FIELD_NUMBER: _ClassVar[int]
    event: DeviceEvent
    def __init__(self, event: _Optional[_Union[DeviceEvent, _Mapping]] = ...) -> None: ...

class TimeQuality(_message.Message):
    __slots__ = ("event",)
    EVENT_FIELD_NUMBER: _ClassVar[int]
    event: DeviceEvent
    def __init__(self, event: _Optional[_Union[DeviceEvent, _Mapping]] = ...) -> None: ...

class SyncStatus(_message.Message):
    __slots__ = ("event",)
    EVENT_FIELD_NUMBER: _ClassVar[int]
    event: DeviceEvent
    def __init__(self, event: _Optional[_Union[DeviceEvent, _Mapping]] = ...) -> None: ...

class Power(_message.Message):
    __slots__ = ("event",)
    EVENT_FIELD_NUMBER: _ClassVar[int]
    event: DeviceEvent
    def __init__(self, event: _Optional[_Union[DeviceEvent, _Mapping]] = ...) -> None: ...

class Config(_message.Message):
    __slots__ = ("event",)
    EVENT_FIELD_NUMBER: _ClassVar[int]
    event: DeviceEvent
    def __init__(self, event: _Optional[_Union[DeviceEvent, _Mapping]] = ...) -> None: ...

class Trigger(_message.Message):
    __slots__ = ("event",)
    EVENT_FIELD_NUMBER: _ClassVar[int]
    event: DeviceEvent
    def __init__(self, event: _Optional[_Union[DeviceEvent, _Mapping]] = ...) -> None: ...

class DataError(_message.Message):
    __slots__ = ("event",)
    EVENT_FIELD_NUMBER: _ClassVar[int]
    event: DeviceEvent
    def __init__(self, event: _Optional[_Union[DeviceEvent, _Mapping]] = ...) -> None: ...
