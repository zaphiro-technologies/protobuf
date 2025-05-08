from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class Blacklist(_message.Message):
    __slots__ = ("measurementIds",)
    MEASUREMENTIDS_FIELD_NUMBER: _ClassVar[int]
    measurementIds: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, measurementIds: _Optional[_Iterable[str]] = ...) -> None: ...
