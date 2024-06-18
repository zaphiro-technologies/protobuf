from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class PhaseConnectedFaultKind(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    PHASE_CONNECTED_FAULT_KIND_UNSPECIFIED: _ClassVar[PhaseConnectedFaultKind]
    PHASE_CONNECTED_FAULT_KIND_LINE_TO_GROUND: _ClassVar[PhaseConnectedFaultKind]
    PHASE_CONNECTED_FAULT_KIND_LINE_TO_LINE: _ClassVar[PhaseConnectedFaultKind]
    PHASE_CONNECTED_FAULT_KIND_LINE_TO_LINE_TO_GROUND: _ClassVar[PhaseConnectedFaultKind]
    PHASE_CONNECTED_FAULT_KIND_LINE_OPEN: _ClassVar[PhaseConnectedFaultKind]

class PhaseCode(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    PHASE_CODE_UNSPECIFIED: _ClassVar[PhaseCode]
    PHASE_CODE_A: _ClassVar[PhaseCode]
    PHASE_CODE_B: _ClassVar[PhaseCode]
    PHASE_CODE_C: _ClassVar[PhaseCode]
    PHASE_CODE_N: _ClassVar[PhaseCode]
    PHASE_CODE_AB: _ClassVar[PhaseCode]
    PHASE_CODE_AC: _ClassVar[PhaseCode]
    PHASE_CODE_BC: _ClassVar[PhaseCode]
    PHASE_CODE_AN: _ClassVar[PhaseCode]
    PHASE_CODE_BN: _ClassVar[PhaseCode]
    PHASE_CODE_CN: _ClassVar[PhaseCode]
    PHASE_CODE_ABC: _ClassVar[PhaseCode]
    PHASE_CODE_ABN: _ClassVar[PhaseCode]
    PHASE_CODE_ACN: _ClassVar[PhaseCode]
    PHASE_CODE_BCN: _ClassVar[PhaseCode]
    PHASE_CODE_ABCN: _ClassVar[PhaseCode]
    PHASE_CODE_S1N: _ClassVar[PhaseCode]
    PHASE_CODE_S2N: _ClassVar[PhaseCode]
    PHASE_CODE_S12N: _ClassVar[PhaseCode]
    PHASE_CODE_S1: _ClassVar[PhaseCode]
    PHASE_CODE_S2: _ClassVar[PhaseCode]
    PHASE_CODE_S12: _ClassVar[PhaseCode]
    PHASE_CODE_X: _ClassVar[PhaseCode]
    PHASE_CODE_XY: _ClassVar[PhaseCode]
    PHASE_CODE_XN: _ClassVar[PhaseCode]
    PHASE_CODE_XYN: _ClassVar[PhaseCode]

class FaultStatus(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    FAULT_STATUS_UNSPECIFIED: _ClassVar[FaultStatus]
    FAULT_STATUS_STARTED: _ClassVar[FaultStatus]
    FAULT_STATUS_ENDED_AND_LOCATED: _ClassVar[FaultStatus]
    FAULT_STATUS_ENDED: _ClassVar[FaultStatus]
    FAULT_STATUS_UNKNOWN: _ClassVar[FaultStatus]
PHASE_CONNECTED_FAULT_KIND_UNSPECIFIED: PhaseConnectedFaultKind
PHASE_CONNECTED_FAULT_KIND_LINE_TO_GROUND: PhaseConnectedFaultKind
PHASE_CONNECTED_FAULT_KIND_LINE_TO_LINE: PhaseConnectedFaultKind
PHASE_CONNECTED_FAULT_KIND_LINE_TO_LINE_TO_GROUND: PhaseConnectedFaultKind
PHASE_CONNECTED_FAULT_KIND_LINE_OPEN: PhaseConnectedFaultKind
PHASE_CODE_UNSPECIFIED: PhaseCode
PHASE_CODE_A: PhaseCode
PHASE_CODE_B: PhaseCode
PHASE_CODE_C: PhaseCode
PHASE_CODE_N: PhaseCode
PHASE_CODE_AB: PhaseCode
PHASE_CODE_AC: PhaseCode
PHASE_CODE_BC: PhaseCode
PHASE_CODE_AN: PhaseCode
PHASE_CODE_BN: PhaseCode
PHASE_CODE_CN: PhaseCode
PHASE_CODE_ABC: PhaseCode
PHASE_CODE_ABN: PhaseCode
PHASE_CODE_ACN: PhaseCode
PHASE_CODE_BCN: PhaseCode
PHASE_CODE_ABCN: PhaseCode
PHASE_CODE_S1N: PhaseCode
PHASE_CODE_S2N: PhaseCode
PHASE_CODE_S12N: PhaseCode
PHASE_CODE_S1: PhaseCode
PHASE_CODE_S2: PhaseCode
PHASE_CODE_S12: PhaseCode
PHASE_CODE_X: PhaseCode
PHASE_CODE_XY: PhaseCode
PHASE_CODE_XN: PhaseCode
PHASE_CODE_XYN: PhaseCode
FAULT_STATUS_UNSPECIFIED: FaultStatus
FAULT_STATUS_STARTED: FaultStatus
FAULT_STATUS_ENDED_AND_LOCATED: FaultStatus
FAULT_STATUS_ENDED: FaultStatus
FAULT_STATUS_UNKNOWN: FaultStatus

class Fault(_message.Message):
    __slots__ = ("Id", "description", "kind", "phases", "updatedAt", "status", "faultyEquipmentId", "faultCurrent", "impactedEquipmentIds", "usedMeasurementIds", "measurementTimestamp")
    ID_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    KIND_FIELD_NUMBER: _ClassVar[int]
    PHASES_FIELD_NUMBER: _ClassVar[int]
    UPDATEDAT_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    FAULTYEQUIPMENTID_FIELD_NUMBER: _ClassVar[int]
    FAULTCURRENT_FIELD_NUMBER: _ClassVar[int]
    IMPACTEDEQUIPMENTIDS_FIELD_NUMBER: _ClassVar[int]
    USEDMEASUREMENTIDS_FIELD_NUMBER: _ClassVar[int]
    MEASUREMENTTIMESTAMP_FIELD_NUMBER: _ClassVar[int]
    Id: str
    description: str
    kind: PhaseConnectedFaultKind
    phases: PhaseCode
    updatedAt: int
    status: FaultStatus
    faultyEquipmentId: str
    faultCurrent: float
    impactedEquipmentIds: _containers.RepeatedScalarFieldContainer[str]
    usedMeasurementIds: _containers.RepeatedCompositeFieldContainer[FaultMeasurement]
    measurementTimestamp: int
    def __init__(self, Id: _Optional[str] = ..., description: _Optional[str] = ..., kind: _Optional[_Union[PhaseConnectedFaultKind, str]] = ..., phases: _Optional[_Union[PhaseCode, str]] = ..., updatedAt: _Optional[int] = ..., status: _Optional[_Union[FaultStatus, str]] = ..., faultyEquipmentId: _Optional[str] = ..., faultCurrent: _Optional[float] = ..., impactedEquipmentIds: _Optional[_Iterable[str]] = ..., usedMeasurementIds: _Optional[_Iterable[_Union[FaultMeasurement, _Mapping]]] = ..., measurementTimestamp: _Optional[int] = ...) -> None: ...

class LineFault(_message.Message):
    __slots__ = ("fault", "lengthFromTerminal1")
    FAULT_FIELD_NUMBER: _ClassVar[int]
    LENGTHFROMTERMINAL1_FIELD_NUMBER: _ClassVar[int]
    fault: Fault
    lengthFromTerminal1: float
    def __init__(self, fault: _Optional[_Union[Fault, _Mapping]] = ..., lengthFromTerminal1: _Optional[float] = ...) -> None: ...

class EquipmentFault(_message.Message):
    __slots__ = ("fault", "terminalID")
    FAULT_FIELD_NUMBER: _ClassVar[int]
    TERMINALID_FIELD_NUMBER: _ClassVar[int]
    fault: Fault
    terminalID: str
    def __init__(self, fault: _Optional[_Union[Fault, _Mapping]] = ..., terminalID: _Optional[str] = ...) -> None: ...

class FaultMeasurement(_message.Message):
    __slots__ = ("positiveSign", "measurementID")
    POSITIVESIGN_FIELD_NUMBER: _ClassVar[int]
    MEASUREMENTID_FIELD_NUMBER: _ClassVar[int]
    positiveSign: bool
    measurementID: str
    def __init__(self, positiveSign: bool = ..., measurementID: _Optional[str] = ...) -> None: ...
