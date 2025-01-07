from zaphiro.grid.v1 import event_pb2 as _event_pb2
from zaphiro.grid.v1 import fault_pb2 as _fault_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class GridEvent(_message.Message):
    __slots__ = ("event", "componentID", "substationID", "value", "referenceLimit", "probability")
    EVENT_FIELD_NUMBER: _ClassVar[int]
    COMPONENTID_FIELD_NUMBER: _ClassVar[int]
    SUBSTATIONID_FIELD_NUMBER: _ClassVar[int]
    VALUE_FIELD_NUMBER: _ClassVar[int]
    REFERENCELIMIT_FIELD_NUMBER: _ClassVar[int]
    PROBABILITY_FIELD_NUMBER: _ClassVar[int]
    event: _event_pb2.Event
    componentID: str
    substationID: str
    value: float
    referenceLimit: float
    probability: float
    def __init__(self, event: _Optional[_Union[_event_pb2.Event, _Mapping]] = ..., componentID: _Optional[str] = ..., substationID: _Optional[str] = ..., value: _Optional[float] = ..., referenceLimit: _Optional[float] = ..., probability: _Optional[float] = ...) -> None: ...

class VoltageEvent(_message.Message):
    __slots__ = ("event", "phaseCode")
    EVENT_FIELD_NUMBER: _ClassVar[int]
    PHASECODE_FIELD_NUMBER: _ClassVar[int]
    event: GridEvent
    phaseCode: _fault_pb2.PhaseCode
    def __init__(self, event: _Optional[_Union[GridEvent, _Mapping]] = ..., phaseCode: _Optional[_Union[_fault_pb2.PhaseCode, str]] = ...) -> None: ...

class CurrentEvent(_message.Message):
    __slots__ = ("event",)
    EVENT_FIELD_NUMBER: _ClassVar[int]
    event: GridEvent
    def __init__(self, event: _Optional[_Union[GridEvent, _Mapping]] = ...) -> None: ...

class PhaseEvent(_message.Message):
    __slots__ = ("event",)
    EVENT_FIELD_NUMBER: _ClassVar[int]
    event: GridEvent
    def __init__(self, event: _Optional[_Union[GridEvent, _Mapping]] = ...) -> None: ...

class FrequencyEvent(_message.Message):
    __slots__ = ("event",)
    EVENT_FIELD_NUMBER: _ClassVar[int]
    event: GridEvent
    def __init__(self, event: _Optional[_Union[GridEvent, _Mapping]] = ...) -> None: ...

class LineCongestion(_message.Message):
    __slots__ = ("event",)
    EVENT_FIELD_NUMBER: _ClassVar[int]
    event: CurrentEvent
    def __init__(self, event: _Optional[_Union[CurrentEvent, _Mapping]] = ...) -> None: ...

class TransformerCongestion(_message.Message):
    __slots__ = ("event",)
    EVENT_FIELD_NUMBER: _ClassVar[int]
    event: CurrentEvent
    def __init__(self, event: _Optional[_Union[CurrentEvent, _Mapping]] = ...) -> None: ...

class VoltageUnbalance(_message.Message):
    __slots__ = ("event",)
    EVENT_FIELD_NUMBER: _ClassVar[int]
    event: VoltageEvent
    def __init__(self, event: _Optional[_Union[VoltageEvent, _Mapping]] = ...) -> None: ...

class VoltageDip(_message.Message):
    __slots__ = ("event",)
    EVENT_FIELD_NUMBER: _ClassVar[int]
    event: VoltageEvent
    def __init__(self, event: _Optional[_Union[VoltageEvent, _Mapping]] = ...) -> None: ...

class VoltageInterruption(_message.Message):
    __slots__ = ("event",)
    EVENT_FIELD_NUMBER: _ClassVar[int]
    event: VoltageEvent
    def __init__(self, event: _Optional[_Union[VoltageEvent, _Mapping]] = ...) -> None: ...

class VoltageSwell(_message.Message):
    __slots__ = ("event",)
    EVENT_FIELD_NUMBER: _ClassVar[int]
    event: VoltageEvent
    def __init__(self, event: _Optional[_Union[VoltageEvent, _Mapping]] = ...) -> None: ...

class VoltageLimit(_message.Message):
    __slots__ = ("event",)
    EVENT_FIELD_NUMBER: _ClassVar[int]
    event: VoltageEvent
    def __init__(self, event: _Optional[_Union[VoltageEvent, _Mapping]] = ...) -> None: ...

class VoltageRapidChange(_message.Message):
    __slots__ = ("event",)
    EVENT_FIELD_NUMBER: _ClassVar[int]
    event: VoltageEvent
    def __init__(self, event: _Optional[_Union[VoltageEvent, _Mapping]] = ...) -> None: ...

class OverFrequency(_message.Message):
    __slots__ = ("event",)
    EVENT_FIELD_NUMBER: _ClassVar[int]
    event: FrequencyEvent
    def __init__(self, event: _Optional[_Union[FrequencyEvent, _Mapping]] = ...) -> None: ...

class UnderFrequency(_message.Message):
    __slots__ = ("event",)
    EVENT_FIELD_NUMBER: _ClassVar[int]
    event: FrequencyEvent
    def __init__(self, event: _Optional[_Union[FrequencyEvent, _Mapping]] = ...) -> None: ...

class FrequencyVariation(_message.Message):
    __slots__ = ("event",)
    EVENT_FIELD_NUMBER: _ClassVar[int]
    event: FrequencyEvent
    def __init__(self, event: _Optional[_Union[FrequencyEvent, _Mapping]] = ...) -> None: ...

class SteadyOscillation(_message.Message):
    __slots__ = ("event",)
    EVENT_FIELD_NUMBER: _ClassVar[int]
    event: PhaseEvent
    def __init__(self, event: _Optional[_Union[PhaseEvent, _Mapping]] = ...) -> None: ...

class TransientOscillation(_message.Message):
    __slots__ = ("event",)
    EVENT_FIELD_NUMBER: _ClassVar[int]
    event: PhaseEvent
    def __init__(self, event: _Optional[_Union[PhaseEvent, _Mapping]] = ...) -> None: ...
