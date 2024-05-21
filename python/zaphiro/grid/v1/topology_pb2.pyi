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
    __slots__ = ("Id", "TerminalIds", "ConnectivityNodeIds", "BranchIds", "ConnectivityNodeContainerId", "BaseVoltageId", "BaseVoltage")
    ID_FIELD_NUMBER: _ClassVar[int]
    TERMINALIDS_FIELD_NUMBER: _ClassVar[int]
    CONNECTIVITYNODEIDS_FIELD_NUMBER: _ClassVar[int]
    BRANCHIDS_FIELD_NUMBER: _ClassVar[int]
    CONNECTIVITYNODECONTAINERID_FIELD_NUMBER: _ClassVar[int]
    BASEVOLTAGEID_FIELD_NUMBER: _ClassVar[int]
    BASEVOLTAGE_FIELD_NUMBER: _ClassVar[int]
    Id: str
    TerminalIds: _containers.RepeatedScalarFieldContainer[str]
    ConnectivityNodeIds: _containers.RepeatedScalarFieldContainer[str]
    BranchIds: _containers.RepeatedScalarFieldContainer[str]
    ConnectivityNodeContainerId: str
    BaseVoltageId: str
    BaseVoltage: float
    def __init__(self, Id: _Optional[str] = ..., TerminalIds: _Optional[_Iterable[str]] = ..., ConnectivityNodeIds: _Optional[_Iterable[str]] = ..., BranchIds: _Optional[_Iterable[str]] = ..., ConnectivityNodeContainerId: _Optional[str] = ..., BaseVoltageId: _Optional[str] = ..., BaseVoltage: _Optional[float] = ...) -> None: ...

class BranchConnection(_message.Message):
    __slots__ = ("TopologicalNodeIds",)
    TOPOLOGICALNODEIDS_FIELD_NUMBER: _ClassVar[int]
    TopologicalNodeIds: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, TopologicalNodeIds: _Optional[_Iterable[str]] = ...) -> None: ...

class TopologicalIsland(_message.Message):
    __slots__ = ("Id", "TopologicalNodes")
    ID_FIELD_NUMBER: _ClassVar[int]
    TOPOLOGICALNODES_FIELD_NUMBER: _ClassVar[int]
    Id: str
    TopologicalNodes: _containers.RepeatedCompositeFieldContainer[TopologicalNode]
    def __init__(self, Id: _Optional[str] = ..., TopologicalNodes: _Optional[_Iterable[_Union[TopologicalNode, _Mapping]]] = ...) -> None: ...

class ComputedTopology(_message.Message):
    __slots__ = ("eqId", "topologicalIslands", "BranchConnections")
    class BranchConnectionsEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: BranchConnection
        def __init__(self, key: _Optional[str] = ..., value: _Optional[_Union[BranchConnection, _Mapping]] = ...) -> None: ...
    EQID_FIELD_NUMBER: _ClassVar[int]
    TOPOLOGICALISLANDS_FIELD_NUMBER: _ClassVar[int]
    BRANCHCONNECTIONS_FIELD_NUMBER: _ClassVar[int]
    eqId: str
    topologicalIslands: _containers.RepeatedCompositeFieldContainer[TopologicalIsland]
    BranchConnections: _containers.MessageMap[str, BranchConnection]
    def __init__(self, eqId: _Optional[str] = ..., topologicalIslands: _Optional[_Iterable[_Union[TopologicalIsland, _Mapping]]] = ..., BranchConnections: _Optional[_Mapping[str, BranchConnection]] = ...) -> None: ...
