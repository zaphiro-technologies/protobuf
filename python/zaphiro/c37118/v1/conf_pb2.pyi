from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class FrameType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    FRAME_TYPE_UNSPECIFIED: _ClassVar[FrameType]
    FRAME_TYPE_DATA: _ClassVar[FrameType]
    FRAME_TYPE_HEADER: _ClassVar[FrameType]
    FRAME_TYPE_CONFIG_1: _ClassVar[FrameType]
    FRAME_TYPE_CONFIG_2: _ClassVar[FrameType]
    FRAME_TYPE_COMMAND: _ClassVar[FrameType]
    FRAME_TYPE_CONFIG_3: _ClassVar[FrameType]
FRAME_TYPE_UNSPECIFIED: FrameType
FRAME_TYPE_DATA: FrameType
FRAME_TYPE_HEADER: FrameType
FRAME_TYPE_CONFIG_1: FrameType
FRAME_TYPE_CONFIG_2: FrameType
FRAME_TYPE_COMMAND: FrameType
FRAME_TYPE_CONFIG_3: FrameType

class ConfFrame(_message.Message):
    __slots__ = ("header", "configs", "DATA_RATE")
    HEADER_FIELD_NUMBER: _ClassVar[int]
    CONFIGS_FIELD_NUMBER: _ClassVar[int]
    DATA_RATE_FIELD_NUMBER: _ClassVar[int]
    header: ConfHeader
    configs: _containers.RepeatedCompositeFieldContainer[Config]
    DATA_RATE: int
    def __init__(self, header: _Optional[_Union[ConfHeader, _Mapping]] = ..., configs: _Optional[_Iterable[_Union[Config, _Mapping]]] = ..., DATA_RATE: _Optional[int] = ...) -> None: ...

class ConfHeader(_message.Message):
    __slots__ = ("SYNC", "FRAMESIZE", "IDCODE", "SOC", "FRACSEC", "TIME_BASE", "NUM_PMU", "version")
    SYNC_FIELD_NUMBER: _ClassVar[int]
    FRAMESIZE_FIELD_NUMBER: _ClassVar[int]
    IDCODE_FIELD_NUMBER: _ClassVar[int]
    SOC_FIELD_NUMBER: _ClassVar[int]
    FRACSEC_FIELD_NUMBER: _ClassVar[int]
    TIME_BASE_FIELD_NUMBER: _ClassVar[int]
    NUM_PMU_FIELD_NUMBER: _ClassVar[int]
    VERSION_FIELD_NUMBER: _ClassVar[int]
    SYNC: int
    FRAMESIZE: int
    IDCODE: int
    SOC: int
    FRACSEC: int
    TIME_BASE: int
    NUM_PMU: int
    version: int
    def __init__(self, SYNC: _Optional[int] = ..., FRAMESIZE: _Optional[int] = ..., IDCODE: _Optional[int] = ..., SOC: _Optional[int] = ..., FRACSEC: _Optional[int] = ..., TIME_BASE: _Optional[int] = ..., NUM_PMU: _Optional[int] = ..., version: _Optional[int] = ...) -> None: ...

class Config(_message.Message):
    __slots__ = ("STN", "IDCODE", "FORMAT", "PHNMR", "ANNMR", "DGNMR", "CHNAM", "PHUNIT", "ANUNIT", "DIGUNIT", "FNOM", "CFGCNT", "G_PMU_ID", "PHSCALE", "ANSCALE", "PMU_LAT", "PMU_LON", "PMU_ELEV", "SVC_CLASS", "WINDOW", "GRP_DLY")
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
    G_PMU_ID_FIELD_NUMBER: _ClassVar[int]
    PHSCALE_FIELD_NUMBER: _ClassVar[int]
    ANSCALE_FIELD_NUMBER: _ClassVar[int]
    PMU_LAT_FIELD_NUMBER: _ClassVar[int]
    PMU_LON_FIELD_NUMBER: _ClassVar[int]
    PMU_ELEV_FIELD_NUMBER: _ClassVar[int]
    SVC_CLASS_FIELD_NUMBER: _ClassVar[int]
    WINDOW_FIELD_NUMBER: _ClassVar[int]
    GRP_DLY_FIELD_NUMBER: _ClassVar[int]
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
    G_PMU_ID: bytes
    PHSCALE: _containers.RepeatedCompositeFieldContainer[PhasorScaling]
    ANSCALE: _containers.RepeatedCompositeFieldContainer[AnalogScaling]
    PMU_LAT: float
    PMU_LON: float
    PMU_ELEV: float
    SVC_CLASS: str
    WINDOW: int
    GRP_DLY: int
    def __init__(self, STN: _Optional[str] = ..., IDCODE: _Optional[int] = ..., FORMAT: _Optional[int] = ..., PHNMR: _Optional[int] = ..., ANNMR: _Optional[int] = ..., DGNMR: _Optional[int] = ..., CHNAM: _Optional[str] = ..., PHUNIT: _Optional[_Iterable[int]] = ..., ANUNIT: _Optional[_Iterable[int]] = ..., DIGUNIT: _Optional[_Iterable[int]] = ..., FNOM: _Optional[int] = ..., CFGCNT: _Optional[int] = ..., G_PMU_ID: _Optional[bytes] = ..., PHSCALE: _Optional[_Iterable[_Union[PhasorScaling, _Mapping]]] = ..., ANSCALE: _Optional[_Iterable[_Union[AnalogScaling, _Mapping]]] = ..., PMU_LAT: _Optional[float] = ..., PMU_LON: _Optional[float] = ..., PMU_ELEV: _Optional[float] = ..., SVC_CLASS: _Optional[str] = ..., WINDOW: _Optional[int] = ..., GRP_DLY: _Optional[int] = ...) -> None: ...

class PhasorScaling(_message.Message):
    __slots__ = ("flags", "scale_factor", "angle_offset")
    FLAGS_FIELD_NUMBER: _ClassVar[int]
    SCALE_FACTOR_FIELD_NUMBER: _ClassVar[int]
    ANGLE_OFFSET_FIELD_NUMBER: _ClassVar[int]
    flags: int
    scale_factor: float
    angle_offset: float
    def __init__(self, flags: _Optional[int] = ..., scale_factor: _Optional[float] = ..., angle_offset: _Optional[float] = ...) -> None: ...

class AnalogScaling(_message.Message):
    __slots__ = ("scale_factor", "offset")
    SCALE_FACTOR_FIELD_NUMBER: _ClassVar[int]
    OFFSET_FIELD_NUMBER: _ClassVar[int]
    scale_factor: float
    offset: float
    def __init__(self, scale_factor: _Optional[float] = ..., offset: _Optional[float] = ...) -> None: ...
