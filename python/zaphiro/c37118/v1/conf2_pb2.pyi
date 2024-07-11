from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Conf2Frame(_message.Message):
    __slots__ = ("header", "configs", "DATA_RATE")
    HEADER_FIELD_NUMBER: _ClassVar[int]
    CONFIGS_FIELD_NUMBER: _ClassVar[int]
    DATA_RATE_FIELD_NUMBER: _ClassVar[int]
    header: Conf2Header
    configs: _containers.RepeatedCompositeFieldContainer[Config]
    DATA_RATE: int
    def __init__(self, header: _Optional[_Union[Conf2Header, _Mapping]] = ..., configs: _Optional[_Iterable[_Union[Config, _Mapping]]] = ..., DATA_RATE: _Optional[int] = ...) -> None: ...

class Conf2Header(_message.Message):
    __slots__ = ("SYNC", "FRAMESIZE", "IDCODE", "SOC", "FRACSEC", "TIME_BASE", "NUM_PMU")
    SYNC_FIELD_NUMBER: _ClassVar[int]
    FRAMESIZE_FIELD_NUMBER: _ClassVar[int]
    IDCODE_FIELD_NUMBER: _ClassVar[int]
    SOC_FIELD_NUMBER: _ClassVar[int]
    FRACSEC_FIELD_NUMBER: _ClassVar[int]
    TIME_BASE_FIELD_NUMBER: _ClassVar[int]
    NUM_PMU_FIELD_NUMBER: _ClassVar[int]
    SYNC: int
    FRAMESIZE: int
    IDCODE: int
    SOC: int
    FRACSEC: int
    TIME_BASE: int
    NUM_PMU: int
    def __init__(self, SYNC: _Optional[int] = ..., FRAMESIZE: _Optional[int] = ..., IDCODE: _Optional[int] = ..., SOC: _Optional[int] = ..., FRACSEC: _Optional[int] = ..., TIME_BASE: _Optional[int] = ..., NUM_PMU: _Optional[int] = ...) -> None: ...

class Config(_message.Message):
    __slots__ = ("STN", "IDCODE", "FORMAT", "PHNMR", "ANNMR", "DGNMR", "CHNAM", "PHUNIT", "ANUNIT", "DIGUNIT", "FNOM", "CFGCNT")
    STN_FIELD_NUMBER: _ClassVar[int]
    IDCODE_FIELD_NUMBER: _ClassVar[int]
    FORMAT_FIELD_NUMBER: _ClassVar[int]
    PHNMR_FIELD_NUMBER: _ClassVar[int]
    ANNMR_FIELD_NUMBER: _ClassVar[int]
    DGNMR_FIELD_NUMBER: _ClassVar[int]
    CHNAM_FIELD_NUMBER: _ClassVar[int]
    PHUNIT_FIELD_NUMBER: _ClassVar[int]
    ANUNIT_FIELD_NUMBER: _ClassVar[int]
    DIGUNIT_FIELD_NUMBER: _ClassVar[int]
    FNOM_FIELD_NUMBER: _ClassVar[int]
    CFGCNT_FIELD_NUMBER: _ClassVar[int]
    STN: str
    IDCODE: int
    FORMAT: int
    PHNMR: int
    ANNMR: int
    DGNMR: int
    CHNAM: str
    PHUNIT: _containers.RepeatedScalarFieldContainer[int]
    ANUNIT: _containers.RepeatedScalarFieldContainer[int]
    DIGUNIT: _containers.RepeatedScalarFieldContainer[int]
    FNOM: int
    CFGCNT: int
    def __init__(self, STN: _Optional[str] = ..., IDCODE: _Optional[int] = ..., FORMAT: _Optional[int] = ..., PHNMR: _Optional[int] = ..., ANNMR: _Optional[int] = ..., DGNMR: _Optional[int] = ..., CHNAM: _Optional[str] = ..., PHUNIT: _Optional[_Iterable[int]] = ..., ANUNIT: _Optional[_Iterable[int]] = ..., DIGUNIT: _Optional[_Iterable[int]] = ..., FNOM: _Optional[int] = ..., CFGCNT: _Optional[int] = ...) -> None: ...
