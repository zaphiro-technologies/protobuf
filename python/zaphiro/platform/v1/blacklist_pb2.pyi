from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class Blacklist(_message.Message):
    __slots__ = ("common", "fault_locator", "event_handler", "state_estimator", "topology_processor", "storer")
    COMMON_FIELD_NUMBER: _ClassVar[int]
    FAULT_LOCATOR_FIELD_NUMBER: _ClassVar[int]
    EVENT_HANDLER_FIELD_NUMBER: _ClassVar[int]
    STATE_ESTIMATOR_FIELD_NUMBER: _ClassVar[int]
    TOPOLOGY_PROCESSOR_FIELD_NUMBER: _ClassVar[int]
    STORER_FIELD_NUMBER: _ClassVar[int]
    common: _containers.RepeatedScalarFieldContainer[str]
    fault_locator: _containers.RepeatedScalarFieldContainer[str]
    event_handler: _containers.RepeatedScalarFieldContainer[str]
    state_estimator: _containers.RepeatedScalarFieldContainer[str]
    topology_processor: _containers.RepeatedScalarFieldContainer[str]
    storer: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, common: _Optional[_Iterable[str]] = ..., fault_locator: _Optional[_Iterable[str]] = ..., event_handler: _Optional[_Iterable[str]] = ..., state_estimator: _Optional[_Iterable[str]] = ..., topology_processor: _Optional[_Iterable[str]] = ..., storer: _Optional[_Iterable[str]] = ...) -> None: ...
