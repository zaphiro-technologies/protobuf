// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: grid/v1/fault.proto

// <!-- markdownlint-disable -->
//Messages describing faults.

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PhaseConnectedFaultKind int32

const (
	PhaseConnectedFaultKind_PHASE_CONNECTED_FAULT_KIND_UNSPECIFIED PhaseConnectedFaultKind = 0
	// The fault connects the indicated phases to ground. The line to line fault impedance is not used and assumed infinite. The full ground impedance is connected between each phase specified in the fault and ground, but not between the phases.
	PhaseConnectedFaultKind_PHASE_CONNECTED_FAULT_KIND_LINE_TO_GROUND PhaseConnectedFaultKind = 1
	// The fault connects the specified phases together without a connection to ground. The ground impedance of this fault is ignored. The line to line impedance is connected between each of the phases specified in the fault. For example three times for a three phase fault, one time for a two phase fault. A single phase fault should not be specified.
	PhaseConnectedFaultKind_PHASE_CONNECTED_FAULT_KIND_LINE_TO_LINE PhaseConnectedFaultKind = 2
	// The fault connects the indicated phases to ground and to each other. The line to line impedance is connected between each of the phases specified in the fault in a full mesh. For example three times for a three phase fault, one time for a two phase fault. A single phase fault should not be specified. The full ground impedance is connected between each phase specified in the fault and ground.
	PhaseConnectedFaultKind_PHASE_CONNECTED_FAULT_KIND_LINE_TO_LINE_TO_GROUND PhaseConnectedFaultKind = 3
	// The fault is when the conductor path is broken between two terminals. Additional coexisting faults may be required if the broken conductor also causes connections to grounds or other lines or phases.
	PhaseConnectedFaultKind_PHASE_CONNECTED_FAULT_KIND_LINE_OPEN PhaseConnectedFaultKind = 4
)

// Enum value maps for PhaseConnectedFaultKind.
var (
	PhaseConnectedFaultKind_name = map[int32]string{
		0: "PHASE_CONNECTED_FAULT_KIND_UNSPECIFIED",
		1: "PHASE_CONNECTED_FAULT_KIND_LINE_TO_GROUND",
		2: "PHASE_CONNECTED_FAULT_KIND_LINE_TO_LINE",
		3: "PHASE_CONNECTED_FAULT_KIND_LINE_TO_LINE_TO_GROUND",
		4: "PHASE_CONNECTED_FAULT_KIND_LINE_OPEN",
	}
	PhaseConnectedFaultKind_value = map[string]int32{
		"PHASE_CONNECTED_FAULT_KIND_UNSPECIFIED":            0,
		"PHASE_CONNECTED_FAULT_KIND_LINE_TO_GROUND":         1,
		"PHASE_CONNECTED_FAULT_KIND_LINE_TO_LINE":           2,
		"PHASE_CONNECTED_FAULT_KIND_LINE_TO_LINE_TO_GROUND": 3,
		"PHASE_CONNECTED_FAULT_KIND_LINE_OPEN":              4,
	}
)

func (x PhaseConnectedFaultKind) Enum() *PhaseConnectedFaultKind {
	p := new(PhaseConnectedFaultKind)
	*p = x
	return p
}

func (x PhaseConnectedFaultKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PhaseConnectedFaultKind) Descriptor() protoreflect.EnumDescriptor {
	return file_grid_v1_fault_proto_enumTypes[0].Descriptor()
}

func (PhaseConnectedFaultKind) Type() protoreflect.EnumType {
	return &file_grid_v1_fault_proto_enumTypes[0]
}

func (x PhaseConnectedFaultKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PhaseConnectedFaultKind.Descriptor instead.
func (PhaseConnectedFaultKind) EnumDescriptor() ([]byte, []int) {
	return file_grid_v1_fault_proto_rawDescGZIP(), []int{0}
}

type PhaseCode int32

const (
	PhaseCode_PHASE_CODE_UNSPECIFIED PhaseCode = 0  //No phases specified.
	PhaseCode_PHASE_CODE_A           PhaseCode = 1  //Phase A.
	PhaseCode_PHASE_CODE_B           PhaseCode = 2  //Phase B.
	PhaseCode_PHASE_CODE_C           PhaseCode = 3  //Phase C.
	PhaseCode_PHASE_CODE_N           PhaseCode = 4  //Neutral phase.
	PhaseCode_PHASE_CODE_AB          PhaseCode = 5  //Phases A and B.
	PhaseCode_PHASE_CODE_AC          PhaseCode = 6  //Phases A and C.
	PhaseCode_PHASE_CODE_BC          PhaseCode = 7  //Phases B and C.
	PhaseCode_PHASE_CODE_AN          PhaseCode = 8  //Phases A and neutral.
	PhaseCode_PHASE_CODE_BN          PhaseCode = 9  //Phases B and neutral.
	PhaseCode_PHASE_CODE_CN          PhaseCode = 10 //Phases C and neutral.
	PhaseCode_PHASE_CODE_ABC         PhaseCode = 11 //Phases A, B, and C.
	PhaseCode_PHASE_CODE_ABN         PhaseCode = 12 //Phases A, B, and neutral.
	PhaseCode_PHASE_CODE_ACN         PhaseCode = 13 //Phases A, C and neutral.
	PhaseCode_PHASE_CODE_BCN         PhaseCode = 14 //Phases B, C, and neutral.
	PhaseCode_PHASE_CODE_ABCN        PhaseCode = 15 //Phases A, B, C, and N.
	PhaseCode_PHASE_CODE_S1N         PhaseCode = 16 //Secondary phase 1 and neutral.
	PhaseCode_PHASE_CODE_S2N         PhaseCode = 17 //Secondary phase 2 and neutral.
	PhaseCode_PHASE_CODE_S12N        PhaseCode = 18 //Secondary phases 1, 2, and neutral.
	PhaseCode_PHASE_CODE_S1          PhaseCode = 19 //Secondary phase 1.
	PhaseCode_PHASE_CODE_S2          PhaseCode = 20 //Secondary phase 2.
	PhaseCode_PHASE_CODE_S12         PhaseCode = 21 //Secondary phase 1 and 2.
	PhaseCode_PHASE_CODE_X           PhaseCode = 22 //Unknown non-neutral phase.
	PhaseCode_PHASE_CODE_XY          PhaseCode = 23 //Two unknown non-neutral phases.
	PhaseCode_PHASE_CODE_XN          PhaseCode = 24 //Unknown non-neutral phase plus neutral.
	PhaseCode_PHASE_CODE_XYN         PhaseCode = 25 //Two unknown non-neutral phases plus neutral.
)

// Enum value maps for PhaseCode.
var (
	PhaseCode_name = map[int32]string{
		0:  "PHASE_CODE_UNSPECIFIED",
		1:  "PHASE_CODE_A",
		2:  "PHASE_CODE_B",
		3:  "PHASE_CODE_C",
		4:  "PHASE_CODE_N",
		5:  "PHASE_CODE_AB",
		6:  "PHASE_CODE_AC",
		7:  "PHASE_CODE_BC",
		8:  "PHASE_CODE_AN",
		9:  "PHASE_CODE_BN",
		10: "PHASE_CODE_CN",
		11: "PHASE_CODE_ABC",
		12: "PHASE_CODE_ABN",
		13: "PHASE_CODE_ACN",
		14: "PHASE_CODE_BCN",
		15: "PHASE_CODE_ABCN",
		16: "PHASE_CODE_S1N",
		17: "PHASE_CODE_S2N",
		18: "PHASE_CODE_S12N",
		19: "PHASE_CODE_S1",
		20: "PHASE_CODE_S2",
		21: "PHASE_CODE_S12",
		22: "PHASE_CODE_X",
		23: "PHASE_CODE_XY",
		24: "PHASE_CODE_XN",
		25: "PHASE_CODE_XYN",
	}
	PhaseCode_value = map[string]int32{
		"PHASE_CODE_UNSPECIFIED": 0,
		"PHASE_CODE_A":           1,
		"PHASE_CODE_B":           2,
		"PHASE_CODE_C":           3,
		"PHASE_CODE_N":           4,
		"PHASE_CODE_AB":          5,
		"PHASE_CODE_AC":          6,
		"PHASE_CODE_BC":          7,
		"PHASE_CODE_AN":          8,
		"PHASE_CODE_BN":          9,
		"PHASE_CODE_CN":          10,
		"PHASE_CODE_ABC":         11,
		"PHASE_CODE_ABN":         12,
		"PHASE_CODE_ACN":         13,
		"PHASE_CODE_BCN":         14,
		"PHASE_CODE_ABCN":        15,
		"PHASE_CODE_S1N":         16,
		"PHASE_CODE_S2N":         17,
		"PHASE_CODE_S12N":        18,
		"PHASE_CODE_S1":          19,
		"PHASE_CODE_S2":          20,
		"PHASE_CODE_S12":         21,
		"PHASE_CODE_X":           22,
		"PHASE_CODE_XY":          23,
		"PHASE_CODE_XN":          24,
		"PHASE_CODE_XYN":         25,
	}
)

func (x PhaseCode) Enum() *PhaseCode {
	p := new(PhaseCode)
	*p = x
	return p
}

func (x PhaseCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PhaseCode) Descriptor() protoreflect.EnumDescriptor {
	return file_grid_v1_fault_proto_enumTypes[1].Descriptor()
}

func (PhaseCode) Type() protoreflect.EnumType {
	return &file_grid_v1_fault_proto_enumTypes[1]
}

func (x PhaseCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PhaseCode.Descriptor instead.
func (PhaseCode) EnumDescriptor() ([]byte, []int) {
	return file_grid_v1_fault_proto_rawDescGZIP(), []int{1}
}

type Fault struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   string                  `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`                                           //The textual id of the fault.
	Description          *string                 `protobuf:"bytes,2,opt,name=description,proto3,oneof" json:"description,omitempty"`                   //The textual description of the fault.
	Kind                 PhaseConnectedFaultKind `protobuf:"varint,3,opt,name=kind,proto3,enum=grid.v1.PhaseConnectedFaultKind" json:"kind,omitempty"` //The kind of phase fault.
	Phases               PhaseCode               `protobuf:"varint,4,opt,name=phases,proto3,enum=grid.v1.PhaseCode" json:"phases,omitempty"`           //The phases participating in the fault. The fault connections into these phases are further specified by the type of fault.
	OccurredAt           int64                   `protobuf:"varint,5,opt,name=occurredAt,proto3" json:"occurredAt,omitempty"`                          //The date and time at which the fault occurred (Unix msec timestamp).
	FaultyEquipmentId    *string                 `protobuf:"bytes,6,opt,name=faultyEquipmentId,proto3,oneof" json:"faultyEquipmentId,omitempty"`       //The equipment with the fault.
	LocatedAt            *int64                  `protobuf:"varint,7,opt,name=locatedAt,proto3,oneof" json:"locatedAt,omitempty"`                      //The time when the fault was located.
	FaultCurrent         *float32                `protobuf:"fixed32,8,opt,name=faultCurrent,proto3,oneof" json:"faultCurrent,omitempty"`               //The current associated to the fault.
	Located              *bool                   `protobuf:"varint,9,opt,name=located,proto3,oneof" json:"located,omitempty"`                          //Was the fault located.
	ImpactedEquipmentIds []string                `protobuf:"bytes,10,rep,name=impactedEquipmentIds,proto3" json:"impactedEquipmentIds,omitempty"`      //The set of IDs of equipments impacted by the fault.
}

func (x *Fault) Reset() {
	*x = Fault{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grid_v1_fault_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fault) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fault) ProtoMessage() {}

func (x *Fault) ProtoReflect() protoreflect.Message {
	mi := &file_grid_v1_fault_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fault.ProtoReflect.Descriptor instead.
func (*Fault) Descriptor() ([]byte, []int) {
	return file_grid_v1_fault_proto_rawDescGZIP(), []int{0}
}

func (x *Fault) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Fault) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *Fault) GetKind() PhaseConnectedFaultKind {
	if x != nil {
		return x.Kind
	}
	return PhaseConnectedFaultKind_PHASE_CONNECTED_FAULT_KIND_UNSPECIFIED
}

func (x *Fault) GetPhases() PhaseCode {
	if x != nil {
		return x.Phases
	}
	return PhaseCode_PHASE_CODE_UNSPECIFIED
}

func (x *Fault) GetOccurredAt() int64 {
	if x != nil {
		return x.OccurredAt
	}
	return 0
}

func (x *Fault) GetFaultyEquipmentId() string {
	if x != nil && x.FaultyEquipmentId != nil {
		return *x.FaultyEquipmentId
	}
	return ""
}

func (x *Fault) GetLocatedAt() int64 {
	if x != nil && x.LocatedAt != nil {
		return *x.LocatedAt
	}
	return 0
}

func (x *Fault) GetFaultCurrent() float32 {
	if x != nil && x.FaultCurrent != nil {
		return *x.FaultCurrent
	}
	return 0
}

func (x *Fault) GetLocated() bool {
	if x != nil && x.Located != nil {
		return *x.Located
	}
	return false
}

func (x *Fault) GetImpactedEquipmentIds() []string {
	if x != nil {
		return x.ImpactedEquipmentIds
	}
	return nil
}

type LineFault struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fault               *Fault   `protobuf:"bytes,1,opt,name=fault,proto3" json:"fault,omitempty"`                                     //The base fault message.
	LengthFromTerminal1 *float32 `protobuf:"fixed32,2,opt,name=lengthFromTerminal1,proto3,oneof" json:"lengthFromTerminal1,omitempty"` //The length to the place where the fault is located starting from terminal with sequence number 1 of the faulted line segment.
	AcLineSegmentID     *string  `protobuf:"bytes,3,opt,name=acLineSegmentID,proto3,oneof" json:"acLineSegmentID,omitempty"`           //The line segment of this line fault.
}

func (x *LineFault) Reset() {
	*x = LineFault{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grid_v1_fault_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LineFault) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LineFault) ProtoMessage() {}

func (x *LineFault) ProtoReflect() protoreflect.Message {
	mi := &file_grid_v1_fault_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LineFault.ProtoReflect.Descriptor instead.
func (*LineFault) Descriptor() ([]byte, []int) {
	return file_grid_v1_fault_proto_rawDescGZIP(), []int{1}
}

func (x *LineFault) GetFault() *Fault {
	if x != nil {
		return x.Fault
	}
	return nil
}

func (x *LineFault) GetLengthFromTerminal1() float32 {
	if x != nil && x.LengthFromTerminal1 != nil {
		return *x.LengthFromTerminal1
	}
	return 0
}

func (x *LineFault) GetAcLineSegmentID() string {
	if x != nil && x.AcLineSegmentID != nil {
		return *x.AcLineSegmentID
	}
	return ""
}

type EquipmentFault struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fault      *Fault  `protobuf:"bytes,1,opt,name=fault,proto3" json:"fault,omitempty"`                 //The base fault message.
	TerminalID *string `protobuf:"bytes,2,opt,name=terminalID,proto3,oneof" json:"terminalID,omitempty"` //The terminal connecting to the bus to which the fault is applied.
}

func (x *EquipmentFault) Reset() {
	*x = EquipmentFault{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grid_v1_fault_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EquipmentFault) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EquipmentFault) ProtoMessage() {}

func (x *EquipmentFault) ProtoReflect() protoreflect.Message {
	mi := &file_grid_v1_fault_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EquipmentFault.ProtoReflect.Descriptor instead.
func (*EquipmentFault) Descriptor() ([]byte, []int) {
	return file_grid_v1_fault_proto_rawDescGZIP(), []int{2}
}

func (x *EquipmentFault) GetFault() *Fault {
	if x != nil {
		return x.Fault
	}
	return nil
}

func (x *EquipmentFault) GetTerminalID() string {
	if x != nil && x.TerminalID != nil {
		return *x.TerminalID
	}
	return ""
}

var File_grid_v1_fault_proto protoreflect.FileDescriptor

var file_grid_v1_fault_proto_rawDesc = []byte{
	0x0a, 0x13, 0x67, 0x72, 0x69, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x67, 0x72, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x22, 0xe3,
	0x03, 0x0a, 0x05, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12,
	0x34, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e,
	0x67, 0x72, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x68, 0x61, 0x73, 0x65, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x4b, 0x69, 0x6e, 0x64, 0x52,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x2a, 0x0a, 0x06, 0x70, 0x68, 0x61, 0x73, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x67, 0x72, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x68, 0x61, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x70, 0x68, 0x61, 0x73, 0x65,
	0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x31, 0x0a, 0x11, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x79, 0x45, 0x71, 0x75, 0x69, 0x70,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x11,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x79, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x48, 0x02, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0c, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x48, 0x03, 0x52,
	0x0c, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01,
	0x12, 0x1d, 0x0a, 0x07, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x08, 0x48, 0x04, 0x52, 0x07, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x32, 0x0a, 0x14, 0x69, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x65, 0x64, 0x45, 0x71, 0x75, 0x69, 0x70,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x14, 0x69,
	0x6d, 0x70, 0x61, 0x63, 0x74, 0x65, 0x64, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x73, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x79, 0x45, 0x71,
	0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x65, 0x64, 0x22, 0xc3, 0x01, 0x0a, 0x09, 0x4c, 0x69, 0x6e, 0x65, 0x46, 0x61, 0x75,
	0x6c, 0x74, 0x12, 0x24, 0x0a, 0x05, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x67, 0x72, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x61, 0x75, 0x6c,
	0x74, 0x52, 0x05, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x35, 0x0a, 0x13, 0x6c, 0x65, 0x6e, 0x67,
	0x74, 0x68, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x31, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x02, 0x48, 0x00, 0x52, 0x13, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x46,
	0x72, 0x6f, 0x6d, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x31, 0x88, 0x01, 0x01, 0x12,
	0x2d, 0x0a, 0x0f, 0x61, 0x63, 0x4c, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0f, 0x61, 0x63, 0x4c, 0x69,
	0x6e, 0x65, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x88, 0x01, 0x01, 0x42, 0x16,
	0x0a, 0x14, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x31, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x61, 0x63, 0x4c, 0x69, 0x6e,
	0x65, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x22, 0x6a, 0x0a, 0x0e, 0x45, 0x71,
	0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x24, 0x0a, 0x05,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x67, 0x72,
	0x69, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x05, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x12, 0x23, 0x0a, 0x0a, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e,
	0x61, 0x6c, 0x49, 0x44, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x74, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x61, 0x6c, 0x49, 0x44, 0x2a, 0x82, 0x02, 0x0a, 0x17, 0x50, 0x68, 0x61, 0x73, 0x65,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x4b, 0x69,
	0x6e, 0x64, 0x12, 0x2a, 0x0a, 0x26, 0x50, 0x48, 0x41, 0x53, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x4e,
	0x45, 0x43, 0x54, 0x45, 0x44, 0x5f, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x5f, 0x4b, 0x49, 0x4e, 0x44,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x2d,
	0x0a, 0x29, 0x50, 0x48, 0x41, 0x53, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x45,
	0x44, 0x5f, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x4c, 0x49, 0x4e,
	0x45, 0x5f, 0x54, 0x4f, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x01, 0x12, 0x2b, 0x0a,
	0x27, 0x50, 0x48, 0x41, 0x53, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x45, 0x44,
	0x5f, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x4c, 0x49, 0x4e, 0x45,
	0x5f, 0x54, 0x4f, 0x5f, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x02, 0x12, 0x35, 0x0a, 0x31, 0x50, 0x48,
	0x41, 0x53, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x45, 0x44, 0x5f, 0x46, 0x41,
	0x55, 0x4c, 0x54, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x54, 0x4f,
	0x5f, 0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x54, 0x4f, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x4e, 0x44, 0x10,
	0x03, 0x12, 0x28, 0x0a, 0x24, 0x50, 0x48, 0x41, 0x53, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45,
	0x43, 0x54, 0x45, 0x44, 0x5f, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f,
	0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x10, 0x04, 0x2a, 0x89, 0x04, 0x0a, 0x09,
	0x50, 0x68, 0x61, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x48, 0x41,
	0x53, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x48, 0x41, 0x53, 0x45, 0x5f, 0x43,
	0x4f, 0x44, 0x45, 0x5f, 0x41, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x48, 0x41, 0x53, 0x45,
	0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x42, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x48, 0x41,
	0x53, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x43, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x50,
	0x48, 0x41, 0x53, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4e, 0x10, 0x04, 0x12, 0x11, 0x0a,
	0x0d, 0x50, 0x48, 0x41, 0x53, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x41, 0x42, 0x10, 0x05,
	0x12, 0x11, 0x0a, 0x0d, 0x50, 0x48, 0x41, 0x53, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x41,
	0x43, 0x10, 0x06, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x48, 0x41, 0x53, 0x45, 0x5f, 0x43, 0x4f, 0x44,
	0x45, 0x5f, 0x42, 0x43, 0x10, 0x07, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x48, 0x41, 0x53, 0x45, 0x5f,
	0x43, 0x4f, 0x44, 0x45, 0x5f, 0x41, 0x4e, 0x10, 0x08, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x48, 0x41,
	0x53, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x42, 0x4e, 0x10, 0x09, 0x12, 0x11, 0x0a, 0x0d,
	0x50, 0x48, 0x41, 0x53, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x43, 0x4e, 0x10, 0x0a, 0x12,
	0x12, 0x0a, 0x0e, 0x50, 0x48, 0x41, 0x53, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x41, 0x42,
	0x43, 0x10, 0x0b, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x48, 0x41, 0x53, 0x45, 0x5f, 0x43, 0x4f, 0x44,
	0x45, 0x5f, 0x41, 0x42, 0x4e, 0x10, 0x0c, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x48, 0x41, 0x53, 0x45,
	0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x41, 0x43, 0x4e, 0x10, 0x0d, 0x12, 0x12, 0x0a, 0x0e, 0x50,
	0x48, 0x41, 0x53, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x42, 0x43, 0x4e, 0x10, 0x0e, 0x12,
	0x13, 0x0a, 0x0f, 0x50, 0x48, 0x41, 0x53, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x41, 0x42,
	0x43, 0x4e, 0x10, 0x0f, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x48, 0x41, 0x53, 0x45, 0x5f, 0x43, 0x4f,
	0x44, 0x45, 0x5f, 0x53, 0x31, 0x4e, 0x10, 0x10, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x48, 0x41, 0x53,
	0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x53, 0x32, 0x4e, 0x10, 0x11, 0x12, 0x13, 0x0a, 0x0f,
	0x50, 0x48, 0x41, 0x53, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x53, 0x31, 0x32, 0x4e, 0x10,
	0x12, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x48, 0x41, 0x53, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f,
	0x53, 0x31, 0x10, 0x13, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x48, 0x41, 0x53, 0x45, 0x5f, 0x43, 0x4f,
	0x44, 0x45, 0x5f, 0x53, 0x32, 0x10, 0x14, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x48, 0x41, 0x53, 0x45,
	0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x53, 0x31, 0x32, 0x10, 0x15, 0x12, 0x10, 0x0a, 0x0c, 0x50,
	0x48, 0x41, 0x53, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x58, 0x10, 0x16, 0x12, 0x11, 0x0a,
	0x0d, 0x50, 0x48, 0x41, 0x53, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x58, 0x59, 0x10, 0x17,
	0x12, 0x11, 0x0a, 0x0d, 0x50, 0x48, 0x41, 0x53, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x58,
	0x4e, 0x10, 0x18, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x48, 0x41, 0x53, 0x45, 0x5f, 0x43, 0x4f, 0x44,
	0x45, 0x5f, 0x58, 0x59, 0x4e, 0x10, 0x19, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x67, 0x72, 0x69,
	0x64, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grid_v1_fault_proto_rawDescOnce sync.Once
	file_grid_v1_fault_proto_rawDescData = file_grid_v1_fault_proto_rawDesc
)

func file_grid_v1_fault_proto_rawDescGZIP() []byte {
	file_grid_v1_fault_proto_rawDescOnce.Do(func() {
		file_grid_v1_fault_proto_rawDescData = protoimpl.X.CompressGZIP(file_grid_v1_fault_proto_rawDescData)
	})
	return file_grid_v1_fault_proto_rawDescData
}

var file_grid_v1_fault_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_grid_v1_fault_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_grid_v1_fault_proto_goTypes = []interface{}{
	(PhaseConnectedFaultKind)(0), // 0: grid.v1.PhaseConnectedFaultKind
	(PhaseCode)(0),               // 1: grid.v1.PhaseCode
	(*Fault)(nil),                // 2: grid.v1.Fault
	(*LineFault)(nil),            // 3: grid.v1.LineFault
	(*EquipmentFault)(nil),       // 4: grid.v1.EquipmentFault
}
var file_grid_v1_fault_proto_depIdxs = []int32{
	0, // 0: grid.v1.Fault.kind:type_name -> grid.v1.PhaseConnectedFaultKind
	1, // 1: grid.v1.Fault.phases:type_name -> grid.v1.PhaseCode
	2, // 2: grid.v1.LineFault.fault:type_name -> grid.v1.Fault
	2, // 3: grid.v1.EquipmentFault.fault:type_name -> grid.v1.Fault
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_grid_v1_fault_proto_init() }
func file_grid_v1_fault_proto_init() {
	if File_grid_v1_fault_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grid_v1_fault_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fault); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grid_v1_fault_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LineFault); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grid_v1_fault_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EquipmentFault); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_grid_v1_fault_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_grid_v1_fault_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_grid_v1_fault_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grid_v1_fault_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_grid_v1_fault_proto_goTypes,
		DependencyIndexes: file_grid_v1_fault_proto_depIdxs,
		EnumInfos:         file_grid_v1_fault_proto_enumTypes,
		MessageInfos:      file_grid_v1_fault_proto_msgTypes,
	}.Build()
	File_grid_v1_fault_proto = out.File
	file_grid_v1_fault_proto_rawDesc = nil
	file_grid_v1_fault_proto_goTypes = nil
	file_grid_v1_fault_proto_depIdxs = nil
}
