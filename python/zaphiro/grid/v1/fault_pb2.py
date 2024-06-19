# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: zaphiro/grid/v1/fault.proto
# Protobuf Python Version: 5.26.0
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1bzaphiro/grid/v1/fault.proto\x12\x0fzaphiro.grid.v1\"\xc2\x05\n\x05\x46\x61ult\x12\x0e\n\x02Id\x18\x01 \x01(\tR\x02Id\x12%\n\x0b\x64\x65scription\x18\x02 \x01(\tH\x00R\x0b\x64\x65scription\x88\x01\x01\x12<\n\x04kind\x18\x03 \x01(\x0e\x32(.zaphiro.grid.v1.PhaseConnectedFaultKindR\x04kind\x12\x32\n\x06phases\x18\x04 \x01(\x0e\x32\x1a.zaphiro.grid.v1.PhaseCodeR\x06phases\x12\x1c\n\tupdatedAt\x18\x05 \x01(\x03R\tupdatedAt\x12\x37\n\x06status\x18\x06 \x01(\x0e\x32\x1f.zaphiro.grid.v1.FaultEventTypeR\x06status\x12\x31\n\x11\x66\x61ultyEquipmentId\x18\x07 \x01(\tH\x01R\x11\x66\x61ultyEquipmentId\x88\x01\x01\x12\'\n\x0c\x66\x61ultCurrent\x18\x08 \x01(\x02H\x02R\x0c\x66\x61ultCurrent\x88\x01\x01\x12\x32\n\x14impactedEquipmentIds\x18\t \x03(\tR\x14impactedEquipmentIds\x12Q\n\x12usedMeasurementIds\x18\n \x03(\x0b\x32!.zaphiro.grid.v1.FaultMeasurementR\x12usedMeasurementIds\x12\x37\n\x14measurementTimestamp\x18\x0b \x01(\x03H\x03R\x14measurementTimestamp\x88\x01\x01\x12\x35\n\x13locationProbability\x18\x0c \x01(\x02H\x04R\x13locationProbability\x88\x01\x01\x42\x0e\n\x0c_descriptionB\x14\n\x12_faultyEquipmentIdB\x0f\n\r_faultCurrentB\x17\n\x15_measurementTimestampB\x16\n\x14_locationProbability\"\xd1\x01\n\tLineFault\x12,\n\x05\x66\x61ult\x18\x01 \x01(\x0b\x32\x16.zaphiro.grid.v1.FaultR\x05\x66\x61ult\x12\x35\n\x13lengthFromTerminal1\x18\x02 \x01(\x02H\x00R\x13lengthFromTerminal1\x88\x01\x01\x12\x31\n\x11lengthUncertainty\x18\x03 \x01(\x02H\x01R\x11lengthUncertainty\x88\x01\x01\x42\x16\n\x14_lengthFromTerminal1B\x14\n\x12_lengthUncertainty\"r\n\x0e\x45quipmentFault\x12,\n\x05\x66\x61ult\x18\x01 \x01(\x0b\x32\x16.zaphiro.grid.v1.FaultR\x05\x66\x61ult\x12#\n\nterminalID\x18\x02 \x01(\tH\x00R\nterminalID\x88\x01\x01\x42\r\n\x0b_terminalID\"\\\n\x10\x46\x61ultMeasurement\x12\"\n\x0cpositiveSign\x18\x01 \x01(\x08R\x0cpositiveSign\x12$\n\rmeasurementID\x18\x02 \x01(\tR\rmeasurementID*\x82\x02\n\x17PhaseConnectedFaultKind\x12*\n&PHASE_CONNECTED_FAULT_KIND_UNSPECIFIED\x10\x00\x12-\n)PHASE_CONNECTED_FAULT_KIND_LINE_TO_GROUND\x10\x01\x12+\n\'PHASE_CONNECTED_FAULT_KIND_LINE_TO_LINE\x10\x02\x12\x35\n1PHASE_CONNECTED_FAULT_KIND_LINE_TO_LINE_TO_GROUND\x10\x03\x12(\n$PHASE_CONNECTED_FAULT_KIND_LINE_OPEN\x10\x04*\x89\x04\n\tPhaseCode\x12\x1a\n\x16PHASE_CODE_UNSPECIFIED\x10\x00\x12\x10\n\x0cPHASE_CODE_A\x10\x01\x12\x10\n\x0cPHASE_CODE_B\x10\x02\x12\x10\n\x0cPHASE_CODE_C\x10\x03\x12\x10\n\x0cPHASE_CODE_N\x10\x04\x12\x11\n\rPHASE_CODE_AB\x10\x05\x12\x11\n\rPHASE_CODE_AC\x10\x06\x12\x11\n\rPHASE_CODE_BC\x10\x07\x12\x11\n\rPHASE_CODE_AN\x10\x08\x12\x11\n\rPHASE_CODE_BN\x10\t\x12\x11\n\rPHASE_CODE_CN\x10\n\x12\x12\n\x0ePHASE_CODE_ABC\x10\x0b\x12\x12\n\x0ePHASE_CODE_ABN\x10\x0c\x12\x12\n\x0ePHASE_CODE_ACN\x10\r\x12\x12\n\x0ePHASE_CODE_BCN\x10\x0e\x12\x13\n\x0fPHASE_CODE_ABCN\x10\x0f\x12\x12\n\x0ePHASE_CODE_S1N\x10\x10\x12\x12\n\x0ePHASE_CODE_S2N\x10\x11\x12\x13\n\x0fPHASE_CODE_S12N\x10\x12\x12\x11\n\rPHASE_CODE_S1\x10\x13\x12\x11\n\rPHASE_CODE_S2\x10\x14\x12\x12\n\x0ePHASE_CODE_S12\x10\x15\x12\x10\n\x0cPHASE_CODE_X\x10\x16\x12\x11\n\rPHASE_CODE_XY\x10\x17\x12\x11\n\rPHASE_CODE_XN\x10\x18\x12\x12\n\x0ePHASE_CODE_XYN\x10\x19*\xa8\x01\n\x0e\x46\x61ultEventType\x12 \n\x1c\x46\x41ULT_EVENT_TYPE_UNSPECIFIED\x10\x00\x12\x1c\n\x18\x46\x41ULT_EVENT_TYPE_STARTED\x10\x01\x12\x1c\n\x18\x46\x41ULT_EVENT_TYPE_LOCATED\x10\x02\x12\x1a\n\x16\x46\x41ULT_EVENT_TYPE_ENDED\x10\x03\x12\x1c\n\x18\x46\x41ULT_EVENT_TYPE_UNKNOWN\x10\x04\x42\x0bZ\t./grid/v1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'zaphiro.grid.v1.fault_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\t./grid/v1'
  _globals['_PHASECONNECTEDFAULTKIND']._serialized_start=1180
  _globals['_PHASECONNECTEDFAULTKIND']._serialized_end=1438
  _globals['_PHASECODE']._serialized_start=1441
  _globals['_PHASECODE']._serialized_end=1962
  _globals['_FAULTEVENTTYPE']._serialized_start=1965
  _globals['_FAULTEVENTTYPE']._serialized_end=2133
  _globals['_FAULT']._serialized_start=49
  _globals['_FAULT']._serialized_end=755
  _globals['_LINEFAULT']._serialized_start=758
  _globals['_LINEFAULT']._serialized_end=967
  _globals['_EQUIPMENTFAULT']._serialized_start=969
  _globals['_EQUIPMENTFAULT']._serialized_end=1083
  _globals['_FAULTMEASUREMENT']._serialized_start=1085
  _globals['_FAULTMEASUREMENT']._serialized_end=1177
# @@protoc_insertion_point(module_scope)
