from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Topology(_message.Message):
    __slots__ = ("createdAt", "tp")
    CREATEDAT_FIELD_NUMBER: _ClassVar[int]
    TP_FIELD_NUMBER: _ClassVar[int]
    createdAt: int
    tp: bytes
    def __init__(self, createdAt: _Optional[int] = ..., tp: _Optional[bytes] = ...) -> None: ...

class TopologicalNode(_message.Message):
    __slots__ = ("TerminalIds", "ConnectivityNodeIds", "PowerTransferEndIds", "ConnectivityNodeContainerId", "BaseVoltageId", "BaseVoltage")
    TERMINALIDS_FIELD_NUMBER: _ClassVar[int]
    CONNECTIVITYNODEIDS_FIELD_NUMBER: _ClassVar[int]
    POWERTRANSFERENDIDS_FIELD_NUMBER: _ClassVar[int]
    CONNECTIVITYNODECONTAINERID_FIELD_NUMBER: _ClassVar[int]
    BASEVOLTAGEID_FIELD_NUMBER: _ClassVar[int]
    BASEVOLTAGE_FIELD_NUMBER: _ClassVar[int]
    TerminalIds: _containers.RepeatedScalarFieldContainer[str]
    ConnectivityNodeIds: _containers.RepeatedScalarFieldContainer[str]
    PowerTransferEndIds: _containers.RepeatedScalarFieldContainer[str]
    ConnectivityNodeContainerId: str
    BaseVoltageId: str
    BaseVoltage: float
    def __init__(self, TerminalIds: _Optional[_Iterable[str]] = ..., ConnectivityNodeIds: _Optional[_Iterable[str]] = ..., PowerTransferEndIds: _Optional[_Iterable[str]] = ..., ConnectivityNodeContainerId: _Optional[str] = ..., BaseVoltageId: _Optional[str] = ..., BaseVoltage: _Optional[float] = ...) -> None: ...

class TopologicalIsland(_message.Message):
    __slots__ = ("TopologicalNodeIds",)
    TOPOLOGICALNODEIDS_FIELD_NUMBER: _ClassVar[int]
    TopologicalNodeIds: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, TopologicalNodeIds: _Optional[_Iterable[str]] = ...) -> None: ...

class ComputedTopology(_message.Message):
    __slots__ = ("eqId", "topologicalNodes", "topologicalIslands")
    class TopologicalNodesEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: TopologicalNode
        def __init__(self, key: _Optional[str] = ..., value: _Optional[_Union[TopologicalNode, _Mapping]] = ...) -> None: ...
    class TopologicalIslandsEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: TopologicalIsland
        def __init__(self, key: _Optional[str] = ..., value: _Optional[_Union[TopologicalIsland, _Mapping]] = ...) -> None: ...
    EQID_FIELD_NUMBER: _ClassVar[int]
    TOPOLOGICALNODES_FIELD_NUMBER: _ClassVar[int]
    TOPOLOGICALISLANDS_FIELD_NUMBER: _ClassVar[int]
    eqId: str
    topologicalNodes: _containers.MessageMap[str, TopologicalNode]
    topologicalIslands: _containers.MessageMap[str, TopologicalIsland]
    def __init__(self, eqId: _Optional[str] = ..., topologicalNodes: _Optional[_Mapping[str, TopologicalNode]] = ..., topologicalIslands: _Optional[_Mapping[str, TopologicalIsland]] = ...) -> None: ...
