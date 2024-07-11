# Copyright 2024 Zaphiro Technologies
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

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
    __slots__ = ("TerminalIds", "ConnectivityNodeIds", "BranchIds", "ConnectivityNodeContainerId", "BaseVoltageId", "BaseVoltage")
    TERMINALIDS_FIELD_NUMBER: _ClassVar[int]
    CONNECTIVITYNODEIDS_FIELD_NUMBER: _ClassVar[int]
    BRANCHIDS_FIELD_NUMBER: _ClassVar[int]
    CONNECTIVITYNODECONTAINERID_FIELD_NUMBER: _ClassVar[int]
    BASEVOLTAGEID_FIELD_NUMBER: _ClassVar[int]
    BASEVOLTAGE_FIELD_NUMBER: _ClassVar[int]
    TerminalIds: _containers.RepeatedScalarFieldContainer[str]
    ConnectivityNodeIds: _containers.RepeatedScalarFieldContainer[str]
    BranchIds: _containers.RepeatedScalarFieldContainer[str]
    ConnectivityNodeContainerId: str
    BaseVoltageId: str
    BaseVoltage: float
    def __init__(self, TerminalIds: _Optional[_Iterable[str]] = ..., ConnectivityNodeIds: _Optional[_Iterable[str]] = ..., BranchIds: _Optional[_Iterable[str]] = ..., ConnectivityNodeContainerId: _Optional[str] = ..., BaseVoltageId: _Optional[str] = ..., BaseVoltage: _Optional[float] = ...) -> None: ...

class TopologicalIsland(_message.Message):
    __slots__ = ("Id", "TopologicalNodes")
    class TopologicalNodesEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: TopologicalNode
        def __init__(self, key: _Optional[str] = ..., value: _Optional[_Union[TopologicalNode, _Mapping]] = ...) -> None: ...
    ID_FIELD_NUMBER: _ClassVar[int]
    TOPOLOGICALNODES_FIELD_NUMBER: _ClassVar[int]
    Id: str
    TopologicalNodes: _containers.MessageMap[str, TopologicalNode]
    def __init__(self, Id: _Optional[str] = ..., TopologicalNodes: _Optional[_Mapping[str, TopologicalNode]] = ...) -> None: ...

class BranchConnection(_message.Message):
    __slots__ = ("TopologicalNodeIds",)
    TOPOLOGICALNODEIDS_FIELD_NUMBER: _ClassVar[int]
    TopologicalNodeIds: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, TopologicalNodeIds: _Optional[_Iterable[str]] = ...) -> None: ...

class ComputedTopology(_message.Message):
    __slots__ = ("eqId", "topologicalIslands", "BranchConnections")
    class TopologicalIslandsEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: TopologicalIsland
        def __init__(self, key: _Optional[str] = ..., value: _Optional[_Union[TopologicalIsland, _Mapping]] = ...) -> None: ...
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
    topologicalIslands: _containers.MessageMap[str, TopologicalIsland]
    BranchConnections: _containers.MessageMap[str, BranchConnection]
    def __init__(self, eqId: _Optional[str] = ..., topologicalIslands: _Optional[_Mapping[str, TopologicalIsland]] = ..., BranchConnections: _Optional[_Mapping[str, BranchConnection]] = ...) -> None: ...
