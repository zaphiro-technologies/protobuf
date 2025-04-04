# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: zaphiro/grid/v1/grid_event.proto
# Protobuf Python Version: 5.26.0
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from zaphiro.grid.v1 import event_pb2 as zaphiro_dot_grid_dot_v1_dot_event__pb2
from zaphiro.grid.v1 import fault_pb2 as zaphiro_dot_grid_dot_v1_dot_fault__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n zaphiro/grid/v1/grid_event.proto\x12\x0fzaphiro.grid.v1\x1a\x1bzaphiro/grid/v1/event.proto\x1a\x1bzaphiro/grid/v1/fault.proto\"\x8a\x02\n\tGridEvent\x12,\n\x05\x65vent\x18\x01 \x01(\x0b\x32\x16.zaphiro.grid.v1.EventR\x05\x65vent\x12 \n\x0b\x63omponentID\x18\x02 \x01(\tR\x0b\x63omponentID\x12\'\n\x0csubstationID\x18\x03 \x01(\tH\x00R\x0csubstationID\x88\x01\x01\x12\x14\n\x05value\x18\x04 \x01(\x01R\x05value\x12&\n\x0ereferenceLimit\x18\x05 \x01(\x01R\x0ereferenceLimit\x12%\n\x0bprobability\x18\x06 \x01(\x01H\x01R\x0bprobability\x88\x01\x01\x42\x0f\n\r_substationIDB\x0e\n\x0c_probability\"\x8d\x01\n\x0cVoltageEvent\x12\x30\n\x05\x65vent\x18\x01 \x01(\x0b\x32\x1a.zaphiro.grid.v1.GridEventR\x05\x65vent\x12=\n\tphaseCode\x18\x02 \x01(\x0e\x32\x1a.zaphiro.grid.v1.PhaseCodeH\x00R\tphaseCode\x88\x01\x01\x42\x0c\n\n_phaseCode\"@\n\x0c\x43urrentEvent\x12\x30\n\x05\x65vent\x18\x01 \x01(\x0b\x32\x1a.zaphiro.grid.v1.GridEventR\x05\x65vent\">\n\nPhaseEvent\x12\x30\n\x05\x65vent\x18\x01 \x01(\x0b\x32\x1a.zaphiro.grid.v1.GridEventR\x05\x65vent\"B\n\x0e\x46requencyEvent\x12\x30\n\x05\x65vent\x18\x01 \x01(\x0b\x32\x1a.zaphiro.grid.v1.GridEventR\x05\x65vent\"E\n\x0eLineCongestion\x12\x33\n\x05\x65vent\x18\x01 \x01(\x0b\x32\x1d.zaphiro.grid.v1.CurrentEventR\x05\x65vent\"L\n\x15TransformerCongestion\x12\x33\n\x05\x65vent\x18\x01 \x01(\x0b\x32\x1d.zaphiro.grid.v1.CurrentEventR\x05\x65vent\"G\n\x10VoltageUnbalance\x12\x33\n\x05\x65vent\x18\x01 \x01(\x0b\x32\x1d.zaphiro.grid.v1.VoltageEventR\x05\x65vent\"A\n\nVoltageDip\x12\x33\n\x05\x65vent\x18\x01 \x01(\x0b\x32\x1d.zaphiro.grid.v1.VoltageEventR\x05\x65vent\"J\n\x13VoltageInterruption\x12\x33\n\x05\x65vent\x18\x01 \x01(\x0b\x32\x1d.zaphiro.grid.v1.VoltageEventR\x05\x65vent\"C\n\x0cVoltageSwell\x12\x33\n\x05\x65vent\x18\x01 \x01(\x0b\x32\x1d.zaphiro.grid.v1.VoltageEventR\x05\x65vent\"C\n\x0cVoltageLimit\x12\x33\n\x05\x65vent\x18\x01 \x01(\x0b\x32\x1d.zaphiro.grid.v1.VoltageEventR\x05\x65vent\"I\n\x12VoltageRapidChange\x12\x33\n\x05\x65vent\x18\x01 \x01(\x0b\x32\x1d.zaphiro.grid.v1.VoltageEventR\x05\x65vent\"F\n\rOverFrequency\x12\x35\n\x05\x65vent\x18\x01 \x01(\x0b\x32\x1f.zaphiro.grid.v1.FrequencyEventR\x05\x65vent\"G\n\x0eUnderFrequency\x12\x35\n\x05\x65vent\x18\x01 \x01(\x0b\x32\x1f.zaphiro.grid.v1.FrequencyEventR\x05\x65vent\"K\n\x12\x46requencyVariation\x12\x35\n\x05\x65vent\x18\x01 \x01(\x0b\x32\x1f.zaphiro.grid.v1.FrequencyEventR\x05\x65vent\"F\n\x11SteadyOscillation\x12\x31\n\x05\x65vent\x18\x01 \x01(\x0b\x32\x1b.zaphiro.grid.v1.PhaseEventR\x05\x65vent\"I\n\x14TransientOscillation\x12\x31\n\x05\x65vent\x18\x01 \x01(\x0b\x32\x1b.zaphiro.grid.v1.PhaseEventR\x05\x65ventB\x0bZ\t./grid/v1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'zaphiro.grid.v1.grid_event_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\t./grid/v1'
  _globals['_GRIDEVENT']._serialized_start=112
  _globals['_GRIDEVENT']._serialized_end=378
  _globals['_VOLTAGEEVENT']._serialized_start=381
  _globals['_VOLTAGEEVENT']._serialized_end=522
  _globals['_CURRENTEVENT']._serialized_start=524
  _globals['_CURRENTEVENT']._serialized_end=588
  _globals['_PHASEEVENT']._serialized_start=590
  _globals['_PHASEEVENT']._serialized_end=652
  _globals['_FREQUENCYEVENT']._serialized_start=654
  _globals['_FREQUENCYEVENT']._serialized_end=720
  _globals['_LINECONGESTION']._serialized_start=722
  _globals['_LINECONGESTION']._serialized_end=791
  _globals['_TRANSFORMERCONGESTION']._serialized_start=793
  _globals['_TRANSFORMERCONGESTION']._serialized_end=869
  _globals['_VOLTAGEUNBALANCE']._serialized_start=871
  _globals['_VOLTAGEUNBALANCE']._serialized_end=942
  _globals['_VOLTAGEDIP']._serialized_start=944
  _globals['_VOLTAGEDIP']._serialized_end=1009
  _globals['_VOLTAGEINTERRUPTION']._serialized_start=1011
  _globals['_VOLTAGEINTERRUPTION']._serialized_end=1085
  _globals['_VOLTAGESWELL']._serialized_start=1087
  _globals['_VOLTAGESWELL']._serialized_end=1154
  _globals['_VOLTAGELIMIT']._serialized_start=1156
  _globals['_VOLTAGELIMIT']._serialized_end=1223
  _globals['_VOLTAGERAPIDCHANGE']._serialized_start=1225
  _globals['_VOLTAGERAPIDCHANGE']._serialized_end=1298
  _globals['_OVERFREQUENCY']._serialized_start=1300
  _globals['_OVERFREQUENCY']._serialized_end=1370
  _globals['_UNDERFREQUENCY']._serialized_start=1372
  _globals['_UNDERFREQUENCY']._serialized_end=1443
  _globals['_FREQUENCYVARIATION']._serialized_start=1445
  _globals['_FREQUENCYVARIATION']._serialized_end=1520
  _globals['_STEADYOSCILLATION']._serialized_start=1522
  _globals['_STEADYOSCILLATION']._serialized_end=1592
  _globals['_TRANSIENTOSCILLATION']._serialized_start=1594
  _globals['_TRANSIENTOSCILLATION']._serialized_end=1667
# @@protoc_insertion_point(module_scope)
