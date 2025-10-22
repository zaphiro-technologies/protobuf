from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class DataType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    DATA_TYPE_UNSPECIFIED: _ClassVar[DataType]
    DATA_TYPE_ACTIVE_POWER: _ClassVar[DataType]
    DATA_TYPE_ANGLE: _ClassVar[DataType]
    DATA_TYPE_AUTOMATIC: _ClassVar[DataType]
    DATA_TYPE_BOOLEAN: _ClassVar[DataType]
    DATA_TYPE_CONFIG_CHANGE: _ClassVar[DataType]
    DATA_TYPE_CURRENT_SEQUENCE_NEGATIVE: _ClassVar[DataType]
    DATA_TYPE_CURRENT_SEQUENCE_POSITIVE: _ClassVar[DataType]
    DATA_TYPE_CURRENT_SEQUENCE_ZERO: _ClassVar[DataType]
    DATA_TYPE_DATA_ERROR: _ClassVar[DataType]
    DATA_TYPE_DATA_MODIFIED: _ClassVar[DataType]
    DATA_TYPE_DATA_SORTING: _ClassVar[DataType]
    DATA_TYPE_ENERGY: _ClassVar[DataType]
    DATA_TYPE_FREQUENCY: _ClassVar[DataType]
    DATA_TYPE_LINE_CURRENT: _ClassVar[DataType]
    DATA_TYPE_LINE_TO_LINE_VOLTAGE: _ClassVar[DataType]
    DATA_TYPE_LOCAL_OPERATION: _ClassVar[DataType]
    DATA_TYPE_OPERATION_COUNT: _ClassVar[DataType]
    DATA_TYPE_PACKET_LATENCY: _ClassVar[DataType]
    DATA_TYPE_PHASE_VOLTAGE: _ClassVar[DataType]
    DATA_TYPE_PHASOR_CURRENT: _ClassVar[DataType]
    DATA_TYPE_PHASOR_VOLTAGE: _ClassVar[DataType]
    DATA_TYPE_PMU_SYNC: _ClassVar[DataType]
    DATA_TYPE_PMU_TIME_QUALITY: _ClassVar[DataType]
    DATA_TYPE_PMU_TRIGGER_DETECTED: _ClassVar[DataType]
    DATA_TYPE_POWER: _ClassVar[DataType]
    DATA_TYPE_POWER_FACTOR: _ClassVar[DataType]
    DATA_TYPE_PRESSURE: _ClassVar[DataType]
    DATA_TYPE_RATE_OF_CHANGE_OF_FREQUENCY: _ClassVar[DataType]
    DATA_TYPE_REACTIVE_POWER: _ClassVar[DataType]
    DATA_TYPE_SWITCH_POSITION: _ClassVar[DataType]
    DATA_TYPE_TAP_POSITION: _ClassVar[DataType]
    DATA_TYPE_TEMPERATURE: _ClassVar[DataType]
    DATA_TYPE_THREE_PHASE_ACTIVE_POWER: _ClassVar[DataType]
    DATA_TYPE_THREE_PHASE_CURRENT: _ClassVar[DataType]
    DATA_TYPE_THREE_PHASE_POWER: _ClassVar[DataType]
    DATA_TYPE_THREE_PHASE_POWER_FACTOR: _ClassVar[DataType]
    DATA_TYPE_THREE_PHASE_REACTIVE_POWER: _ClassVar[DataType]
    DATA_TYPE_TRIGGER_REASON: _ClassVar[DataType]
    DATA_TYPE_UNLOCKED_TIME: _ClassVar[DataType]
    DATA_TYPE_VOLTAGE_SEQUENCE_NEGATIVE: _ClassVar[DataType]
    DATA_TYPE_VOLTAGE_SEQUENCE_POSITIVE: _ClassVar[DataType]
    DATA_TYPE_VOLTAGE_SEQUENCE_ZERO: _ClassVar[DataType]
    DATA_TYPE_DISCRETE: _ClassVar[DataType]
    DATA_TYPE_ANALOG: _ClassVar[DataType]
    DATA_TYPE_NORMALIZED_RESIDUAL: _ClassVar[DataType]
    DATA_TYPE_OBJECTIVE_FUNCTION: _ClassVar[DataType]
    DATA_TYPE_SWITCH_POSITION_CHECK: _ClassVar[DataType]
    DATA_TYPE_TEMPERATURE_ALARM: _ClassVar[DataType]
    DATA_TYPE_PRESSURE_ALARM: _ClassVar[DataType]
    DATA_TYPE_DOOR_ALARM: _ClassVar[DataType]
    DATA_TYPE_ABSOLUTE_RESIDUAL: _ClassVar[DataType]
DATA_TYPE_UNSPECIFIED: DataType
DATA_TYPE_ACTIVE_POWER: DataType
DATA_TYPE_ANGLE: DataType
DATA_TYPE_AUTOMATIC: DataType
DATA_TYPE_BOOLEAN: DataType
DATA_TYPE_CONFIG_CHANGE: DataType
DATA_TYPE_CURRENT_SEQUENCE_NEGATIVE: DataType
DATA_TYPE_CURRENT_SEQUENCE_POSITIVE: DataType
DATA_TYPE_CURRENT_SEQUENCE_ZERO: DataType
DATA_TYPE_DATA_ERROR: DataType
DATA_TYPE_DATA_MODIFIED: DataType
DATA_TYPE_DATA_SORTING: DataType
DATA_TYPE_ENERGY: DataType
DATA_TYPE_FREQUENCY: DataType
DATA_TYPE_LINE_CURRENT: DataType
DATA_TYPE_LINE_TO_LINE_VOLTAGE: DataType
DATA_TYPE_LOCAL_OPERATION: DataType
DATA_TYPE_OPERATION_COUNT: DataType
DATA_TYPE_PACKET_LATENCY: DataType
DATA_TYPE_PHASE_VOLTAGE: DataType
DATA_TYPE_PHASOR_CURRENT: DataType
DATA_TYPE_PHASOR_VOLTAGE: DataType
DATA_TYPE_PMU_SYNC: DataType
DATA_TYPE_PMU_TIME_QUALITY: DataType
DATA_TYPE_PMU_TRIGGER_DETECTED: DataType
DATA_TYPE_POWER: DataType
DATA_TYPE_POWER_FACTOR: DataType
DATA_TYPE_PRESSURE: DataType
DATA_TYPE_RATE_OF_CHANGE_OF_FREQUENCY: DataType
DATA_TYPE_REACTIVE_POWER: DataType
DATA_TYPE_SWITCH_POSITION: DataType
DATA_TYPE_TAP_POSITION: DataType
DATA_TYPE_TEMPERATURE: DataType
DATA_TYPE_THREE_PHASE_ACTIVE_POWER: DataType
DATA_TYPE_THREE_PHASE_CURRENT: DataType
DATA_TYPE_THREE_PHASE_POWER: DataType
DATA_TYPE_THREE_PHASE_POWER_FACTOR: DataType
DATA_TYPE_THREE_PHASE_REACTIVE_POWER: DataType
DATA_TYPE_TRIGGER_REASON: DataType
DATA_TYPE_UNLOCKED_TIME: DataType
DATA_TYPE_VOLTAGE_SEQUENCE_NEGATIVE: DataType
DATA_TYPE_VOLTAGE_SEQUENCE_POSITIVE: DataType
DATA_TYPE_VOLTAGE_SEQUENCE_ZERO: DataType
DATA_TYPE_DISCRETE: DataType
DATA_TYPE_ANALOG: DataType
DATA_TYPE_NORMALIZED_RESIDUAL: DataType
DATA_TYPE_OBJECTIVE_FUNCTION: DataType
DATA_TYPE_SWITCH_POSITION_CHECK: DataType
DATA_TYPE_TEMPERATURE_ALARM: DataType
DATA_TYPE_PRESSURE_ALARM: DataType
DATA_TYPE_DOOR_ALARM: DataType
DATA_TYPE_ABSOLUTE_RESIDUAL: DataType

class Data(_message.Message):
    __slots__ = ("dataType", "measuredAt", "value")
    DATATYPE_FIELD_NUMBER: _ClassVar[int]
    MEASUREDAT_FIELD_NUMBER: _ClassVar[int]
    VALUE_FIELD_NUMBER: _ClassVar[int]
    dataType: DataType
    measuredAt: int
    value: int
    def __init__(self, dataType: _Optional[_Union[DataType, str]] = ..., measuredAt: _Optional[int] = ..., value: _Optional[int] = ...) -> None: ...

class DataSet(_message.Message):
    __slots__ = ("producerId", "data")
    class DataEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: Data
        def __init__(self, key: _Optional[str] = ..., value: _Optional[_Union[Data, _Mapping]] = ...) -> None: ...
    PRODUCERID_FIELD_NUMBER: _ClassVar[int]
    DATA_FIELD_NUMBER: _ClassVar[int]
    producerId: str
    data: _containers.MessageMap[str, Data]
    def __init__(self, producerId: _Optional[str] = ..., data: _Optional[_Mapping[str, Data]] = ...) -> None: ...
