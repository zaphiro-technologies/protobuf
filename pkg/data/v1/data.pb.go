// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: proto/data/v1/data.proto

//  <!-- markdownlint-disable -->
//Messages to support data injection in the platform.
//The data injected may be originated from different sources (e.g. a PMU, RTU, an external service).
//
//Data are grouped into sets, where each id identifies a specific measurement. The id does not identify the instance of measurement, but the class of measurement. Measurement ID can be used to retrieve additional medata about the measurement, from example, in the CIM OP profile associated to the monitored grid.

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

type DataType int32

const (
	DataType_DATA_TYPE_UNSPECIFIED                 DataType = 0  //No type define
	DataType_DATA_TYPE_ACTIVE_POWER                DataType = 1  //Active Power. It maps to an Analog.
	DataType_DATA_TYPE_ANGLE                       DataType = 2  //Angle. It maps to an Analog.
	DataType_DATA_TYPE_AUTOMATIC                   DataType = 3  //Auomatic. It maps to Discrete (True or False).
	DataType_DATA_TYPE_BOOLEAN                     DataType = 4  //Booleam. It maps to Discrete (True or False).
	DataType_DATA_TYPE_CONFIG_CHANGE               DataType = 5  // Configuration change.  It maps to Discrete (True or False).
	DataType_DATA_TYPE_CURRENT_SEQUENCE_NEGATIVE   DataType = 6  //PHASOR
	DataType_DATA_TYPE_CURRENT_SEQUENCE_POSITIVE   DataType = 7  //PHASOR
	DataType_DATA_TYPE_CURRENT_SEQUENCE_ZERO       DataType = 8  //PHASOR
	DataType_DATA_TYPE_DATA_ERROR                  DataType = 9  //DISCRETE (True or False)
	DataType_DATA_TYPE_DATA_MODIFIED               DataType = 10 //DISCRETE (True or False)
	DataType_DATA_TYPE_DATA_SORTING                DataType = 11 //DISCRETE
	DataType_DATA_TYPE_ENERGY                      DataType = 12 //ANALOG
	DataType_DATA_TYPE_FREQUENCY                   DataType = 13 //ANALOG
	DataType_DATA_TYPE_LINE_CURRENT                DataType = 14 //ANALOG
	DataType_DATA_TYPE_LINE_TO_LINE_VOLTAGE        DataType = 15 //ANALOG
	DataType_DATA_TYPE_LOCAL_OPERATION             DataType = 16 //DISCRETE
	DataType_DATA_TYPE_OPERATION_COUNT             DataType = 17 //DISCRETE
	DataType_DATA_TYPE_PACKET_LATENCY              DataType = 18 //ANALOG
	DataType_DATA_TYPE_PHASE_VOLTAGE               DataType = 19 //ANALOG
	DataType_DATA_TYPE_PHASOR_CURRENT              DataType = 20 //PHASOR
	DataType_DATA_TYPE_PHASOR_VOLTAGE              DataType = 21 //PHASOR
	DataType_DATA_TYPE_PMU_SYNC                    DataType = 22 //DISCRETE (True or False)
	DataType_DATA_TYPE_PMU_TIME_QUALITY            DataType = 23 //DISCRETE
	DataType_DATA_TYPE_PMU_TRIGGER_DETECTED        DataType = 24 //DISCRETE (True or False)
	DataType_DATA_TYPE_POWER                       DataType = 25 //ANALOG
	DataType_DATA_TYPE_POWER_FACTOR                DataType = 26 //ANALOG
	DataType_DATA_TYPE_PRESSURE                    DataType = 27 //ANALOG
	DataType_DATA_TYPE_RATE_OF_CHANGE_OF_FREQUENCY DataType = 28 //ANALOG
	DataType_DATA_TYPE_REACTIVE_POWER              DataType = 29 //ANALOG
	DataType_DATA_TYPE_SWITCH_POSITION             DataType = 30 //DISCRETE (True or False)
	DataType_DATA_TYPE_TAP_POSITION                DataType = 31 //DISCRETE
	DataType_DATA_TYPE_TEMPERATURE                 DataType = 32 //ANALOG
	DataType_DATA_TYPE_THREE_PHASE_ACTIVE_POWER    DataType = 33 //ANALOG
	DataType_DATA_TYPE_THREE_PHASE_CURRENT         DataType = 34 //ANALOG
	DataType_DATA_TYPE_THREE_PHASE_POWER           DataType = 35 //ANALOG
	DataType_DATA_TYPE_THREE_PHASE_POWER_FACTOR    DataType = 36 //ANALOG
	DataType_DATA_TYPE_THREE_PHASE_REACTIVE_POWER  DataType = 37 //ANALOG
	DataType_DATA_TYPE_TRIGGER_REASON              DataType = 38 //DISCRETE
	DataType_DATA_TYPE_UNLOCKED_TIME               DataType = 39 //DISCRETE
	DataType_DATA_TYPE_VOLTAGE_SEQUENCE_NEGATIVE   DataType = 40 //PHASOR
	DataType_DATA_TYPE_VOLTAGE_SEQUENCE_POSITIVE   DataType = 41 //PHASOR
	DataType_DATA_TYPE_VOLTAGE_SEQUENCE_ZERO       DataType = 42 //PHASOR
	DataType_DATA_TYPE_DISCRETE                    DataType = 43 //DISCRETE
	DataType_DATA_TYPE_ANALOG                      DataType = 44 //ANALOG
)

// Enum value maps for DataType.
var (
	DataType_name = map[int32]string{
		0:  "DATA_TYPE_UNSPECIFIED",
		1:  "DATA_TYPE_ACTIVE_POWER",
		2:  "DATA_TYPE_ANGLE",
		3:  "DATA_TYPE_AUTOMATIC",
		4:  "DATA_TYPE_BOOLEAN",
		5:  "DATA_TYPE_CONFIG_CHANGE",
		6:  "DATA_TYPE_CURRENT_SEQUENCE_NEGATIVE",
		7:  "DATA_TYPE_CURRENT_SEQUENCE_POSITIVE",
		8:  "DATA_TYPE_CURRENT_SEQUENCE_ZERO",
		9:  "DATA_TYPE_DATA_ERROR",
		10: "DATA_TYPE_DATA_MODIFIED",
		11: "DATA_TYPE_DATA_SORTING",
		12: "DATA_TYPE_ENERGY",
		13: "DATA_TYPE_FREQUENCY",
		14: "DATA_TYPE_LINE_CURRENT",
		15: "DATA_TYPE_LINE_TO_LINE_VOLTAGE",
		16: "DATA_TYPE_LOCAL_OPERATION",
		17: "DATA_TYPE_OPERATION_COUNT",
		18: "DATA_TYPE_PACKET_LATENCY",
		19: "DATA_TYPE_PHASE_VOLTAGE",
		20: "DATA_TYPE_PHASOR_CURRENT",
		21: "DATA_TYPE_PHASOR_VOLTAGE",
		22: "DATA_TYPE_PMU_SYNC",
		23: "DATA_TYPE_PMU_TIME_QUALITY",
		24: "DATA_TYPE_PMU_TRIGGER_DETECTED",
		25: "DATA_TYPE_POWER",
		26: "DATA_TYPE_POWER_FACTOR",
		27: "DATA_TYPE_PRESSURE",
		28: "DATA_TYPE_RATE_OF_CHANGE_OF_FREQUENCY",
		29: "DATA_TYPE_REACTIVE_POWER",
		30: "DATA_TYPE_SWITCH_POSITION",
		31: "DATA_TYPE_TAP_POSITION",
		32: "DATA_TYPE_TEMPERATURE",
		33: "DATA_TYPE_THREE_PHASE_ACTIVE_POWER",
		34: "DATA_TYPE_THREE_PHASE_CURRENT",
		35: "DATA_TYPE_THREE_PHASE_POWER",
		36: "DATA_TYPE_THREE_PHASE_POWER_FACTOR",
		37: "DATA_TYPE_THREE_PHASE_REACTIVE_POWER",
		38: "DATA_TYPE_TRIGGER_REASON",
		39: "DATA_TYPE_UNLOCKED_TIME",
		40: "DATA_TYPE_VOLTAGE_SEQUENCE_NEGATIVE",
		41: "DATA_TYPE_VOLTAGE_SEQUENCE_POSITIVE",
		42: "DATA_TYPE_VOLTAGE_SEQUENCE_ZERO",
		43: "DATA_TYPE_DISCRETE",
		44: "DATA_TYPE_ANALOG",
	}
	DataType_value = map[string]int32{
		"DATA_TYPE_UNSPECIFIED":                 0,
		"DATA_TYPE_ACTIVE_POWER":                1,
		"DATA_TYPE_ANGLE":                       2,
		"DATA_TYPE_AUTOMATIC":                   3,
		"DATA_TYPE_BOOLEAN":                     4,
		"DATA_TYPE_CONFIG_CHANGE":               5,
		"DATA_TYPE_CURRENT_SEQUENCE_NEGATIVE":   6,
		"DATA_TYPE_CURRENT_SEQUENCE_POSITIVE":   7,
		"DATA_TYPE_CURRENT_SEQUENCE_ZERO":       8,
		"DATA_TYPE_DATA_ERROR":                  9,
		"DATA_TYPE_DATA_MODIFIED":               10,
		"DATA_TYPE_DATA_SORTING":                11,
		"DATA_TYPE_ENERGY":                      12,
		"DATA_TYPE_FREQUENCY":                   13,
		"DATA_TYPE_LINE_CURRENT":                14,
		"DATA_TYPE_LINE_TO_LINE_VOLTAGE":        15,
		"DATA_TYPE_LOCAL_OPERATION":             16,
		"DATA_TYPE_OPERATION_COUNT":             17,
		"DATA_TYPE_PACKET_LATENCY":              18,
		"DATA_TYPE_PHASE_VOLTAGE":               19,
		"DATA_TYPE_PHASOR_CURRENT":              20,
		"DATA_TYPE_PHASOR_VOLTAGE":              21,
		"DATA_TYPE_PMU_SYNC":                    22,
		"DATA_TYPE_PMU_TIME_QUALITY":            23,
		"DATA_TYPE_PMU_TRIGGER_DETECTED":        24,
		"DATA_TYPE_POWER":                       25,
		"DATA_TYPE_POWER_FACTOR":                26,
		"DATA_TYPE_PRESSURE":                    27,
		"DATA_TYPE_RATE_OF_CHANGE_OF_FREQUENCY": 28,
		"DATA_TYPE_REACTIVE_POWER":              29,
		"DATA_TYPE_SWITCH_POSITION":             30,
		"DATA_TYPE_TAP_POSITION":                31,
		"DATA_TYPE_TEMPERATURE":                 32,
		"DATA_TYPE_THREE_PHASE_ACTIVE_POWER":    33,
		"DATA_TYPE_THREE_PHASE_CURRENT":         34,
		"DATA_TYPE_THREE_PHASE_POWER":           35,
		"DATA_TYPE_THREE_PHASE_POWER_FACTOR":    36,
		"DATA_TYPE_THREE_PHASE_REACTIVE_POWER":  37,
		"DATA_TYPE_TRIGGER_REASON":              38,
		"DATA_TYPE_UNLOCKED_TIME":               39,
		"DATA_TYPE_VOLTAGE_SEQUENCE_NEGATIVE":   40,
		"DATA_TYPE_VOLTAGE_SEQUENCE_POSITIVE":   41,
		"DATA_TYPE_VOLTAGE_SEQUENCE_ZERO":       42,
		"DATA_TYPE_DISCRETE":                    43,
		"DATA_TYPE_ANALOG":                      44,
	}
)

func (x DataType) Enum() *DataType {
	p := new(DataType)
	*p = x
	return p
}

func (x DataType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_data_v1_data_proto_enumTypes[0].Descriptor()
}

func (DataType) Type() protoreflect.EnumType {
	return &file_proto_data_v1_data_proto_enumTypes[0]
}

func (x DataType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataType.Descriptor instead.
func (DataType) EnumDescriptor() ([]byte, []int) {
	return file_proto_data_v1_data_proto_rawDescGZIP(), []int{0}
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataType   DataType `protobuf:"varint,1,opt,name=dataType,proto3,enum=data.v1.DataType" json:"dataType,omitempty"`
	MeasuredAt int64    `protobuf:"varint,2,opt,name=measuredAt,proto3" json:"measuredAt,omitempty"`
	Value      *uint64  `protobuf:"varint,3,opt,name=value,proto3,oneof" json:"value,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_data_v1_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_proto_data_v1_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_proto_data_v1_data_proto_rawDescGZIP(), []int{0}
}

func (x *Data) GetDataType() DataType {
	if x != nil {
		return x.DataType
	}
	return DataType_DATA_TYPE_UNSPECIFIED
}

func (x *Data) GetMeasuredAt() int64 {
	if x != nil {
		return x.MeasuredAt
	}
	return 0
}

func (x *Data) GetValue() uint64 {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return 0
}

type DataSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProducerId string           `protobuf:"bytes,1,opt,name=producerId,proto3" json:"producerId,omitempty"`
	Data       map[string]*Data `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *DataSet) Reset() {
	*x = DataSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_data_v1_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataSet) ProtoMessage() {}

func (x *DataSet) ProtoReflect() protoreflect.Message {
	mi := &file_proto_data_v1_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataSet.ProtoReflect.Descriptor instead.
func (*DataSet) Descriptor() ([]byte, []int) {
	return file_proto_data_v1_data_proto_rawDescGZIP(), []int{1}
}

func (x *DataSet) GetProducerId() string {
	if x != nil {
		return x.ProducerId
	}
	return ""
}

func (x *DataSet) GetData() map[string]*Data {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_proto_data_v1_data_proto protoreflect.FileDescriptor

var file_proto_data_v1_data_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x76, 0x31, 0x22, 0x7a, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2d, 0x0a, 0x08, 0x64,
	0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x65,
	0x61, 0x73, 0x75, 0x72, 0x65, 0x64, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x64, 0x41, 0x74, 0x12, 0x19, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0xa1, 0x01, 0x0a, 0x07, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x46, 0x0a, 0x09, 0x44,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x2a, 0xf3, 0x0a, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x19, 0x0a, 0x15, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x44,
	0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f,
	0x50, 0x4f, 0x57, 0x45, 0x52, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x44, 0x41, 0x54, 0x41, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x4e, 0x47, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13,
	0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x55, 0x54, 0x4f, 0x4d, 0x41,
	0x54, 0x49, 0x43, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x42, 0x4f, 0x4f, 0x4c, 0x45, 0x41, 0x4e, 0x10, 0x04, 0x12, 0x1b, 0x0a, 0x17,
	0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47,
	0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x05, 0x12, 0x27, 0x0a, 0x23, 0x44, 0x41, 0x54,
	0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x55, 0x52, 0x52, 0x45, 0x4e, 0x54, 0x5f, 0x53,
	0x45, 0x51, 0x55, 0x45, 0x4e, 0x43, 0x45, 0x5f, 0x4e, 0x45, 0x47, 0x41, 0x54, 0x49, 0x56, 0x45,
	0x10, 0x06, 0x12, 0x27, 0x0a, 0x23, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x43, 0x55, 0x52, 0x52, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x45, 0x51, 0x55, 0x45, 0x4e, 0x43, 0x45,
	0x5f, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x49, 0x56, 0x45, 0x10, 0x07, 0x12, 0x23, 0x0a, 0x1f, 0x44,
	0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x55, 0x52, 0x52, 0x45, 0x4e, 0x54,
	0x5f, 0x53, 0x45, 0x51, 0x55, 0x45, 0x4e, 0x43, 0x45, 0x5f, 0x5a, 0x45, 0x52, 0x4f, 0x10, 0x08,
	0x12, 0x18, 0x0a, 0x14, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x41,
	0x54, 0x41, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x09, 0x12, 0x1b, 0x0a, 0x17, 0x44, 0x41,
	0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x4d, 0x4f, 0x44,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x0a, 0x12, 0x1a, 0x0a, 0x16, 0x44, 0x41, 0x54, 0x41, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x53, 0x4f, 0x52, 0x54, 0x49, 0x4e,
	0x47, 0x10, 0x0b, 0x12, 0x14, 0x0a, 0x10, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x45, 0x4e, 0x45, 0x52, 0x47, 0x59, 0x10, 0x0c, 0x12, 0x17, 0x0a, 0x13, 0x44, 0x41, 0x54,
	0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x52, 0x45, 0x51, 0x55, 0x45, 0x4e, 0x43, 0x59,
	0x10, 0x0d, 0x12, 0x1a, 0x0a, 0x16, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x43, 0x55, 0x52, 0x52, 0x45, 0x4e, 0x54, 0x10, 0x0e, 0x12, 0x22,
	0x0a, 0x1e, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4c, 0x49, 0x4e, 0x45,
	0x5f, 0x54, 0x4f, 0x5f, 0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x56, 0x4f, 0x4c, 0x54, 0x41, 0x47, 0x45,
	0x10, 0x0f, 0x12, 0x1d, 0x0a, 0x19, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10,
	0x10, 0x12, 0x1d, 0x0a, 0x19, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f,
	0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x11,
	0x12, 0x1c, 0x0a, 0x18, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x41,
	0x43, 0x4b, 0x45, 0x54, 0x5f, 0x4c, 0x41, 0x54, 0x45, 0x4e, 0x43, 0x59, 0x10, 0x12, 0x12, 0x1b,
	0x0a, 0x17, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x48, 0x41, 0x53,
	0x45, 0x5f, 0x56, 0x4f, 0x4c, 0x54, 0x41, 0x47, 0x45, 0x10, 0x13, 0x12, 0x1c, 0x0a, 0x18, 0x44,
	0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x48, 0x41, 0x53, 0x4f, 0x52, 0x5f,
	0x43, 0x55, 0x52, 0x52, 0x45, 0x4e, 0x54, 0x10, 0x14, 0x12, 0x1c, 0x0a, 0x18, 0x44, 0x41, 0x54,
	0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x48, 0x41, 0x53, 0x4f, 0x52, 0x5f, 0x56, 0x4f,
	0x4c, 0x54, 0x41, 0x47, 0x45, 0x10, 0x15, 0x12, 0x16, 0x0a, 0x12, 0x44, 0x41, 0x54, 0x41, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x4d, 0x55, 0x5f, 0x53, 0x59, 0x4e, 0x43, 0x10, 0x16, 0x12,
	0x1e, 0x0a, 0x1a, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x4d, 0x55,
	0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x51, 0x55, 0x41, 0x4c, 0x49, 0x54, 0x59, 0x10, 0x17, 0x12,
	0x22, 0x0a, 0x1e, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x4d, 0x55,
	0x5f, 0x54, 0x52, 0x49, 0x47, 0x47, 0x45, 0x52, 0x5f, 0x44, 0x45, 0x54, 0x45, 0x43, 0x54, 0x45,
	0x44, 0x10, 0x18, 0x12, 0x13, 0x0a, 0x0f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x50, 0x4f, 0x57, 0x45, 0x52, 0x10, 0x19, 0x12, 0x1a, 0x0a, 0x16, 0x44, 0x41, 0x54, 0x41,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x4f, 0x57, 0x45, 0x52, 0x5f, 0x46, 0x41, 0x43, 0x54,
	0x4f, 0x52, 0x10, 0x1a, 0x12, 0x16, 0x0a, 0x12, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x50, 0x52, 0x45, 0x53, 0x53, 0x55, 0x52, 0x45, 0x10, 0x1b, 0x12, 0x29, 0x0a, 0x25,
	0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x41, 0x54, 0x45, 0x5f, 0x4f,
	0x46, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x4f, 0x46, 0x5f, 0x46, 0x52, 0x45, 0x51,
	0x55, 0x45, 0x4e, 0x43, 0x59, 0x10, 0x1c, 0x12, 0x1c, 0x0a, 0x18, 0x44, 0x41, 0x54, 0x41, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x50, 0x4f,
	0x57, 0x45, 0x52, 0x10, 0x1d, 0x12, 0x1d, 0x0a, 0x19, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x53, 0x57, 0x49, 0x54, 0x43, 0x48, 0x5f, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x49,
	0x4f, 0x4e, 0x10, 0x1e, 0x12, 0x1a, 0x0a, 0x16, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x54, 0x41, 0x50, 0x5f, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x1f,
	0x12, 0x19, 0x0a, 0x15, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x45,
	0x4d, 0x50, 0x45, 0x52, 0x41, 0x54, 0x55, 0x52, 0x45, 0x10, 0x20, 0x12, 0x26, 0x0a, 0x22, 0x44,
	0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x48, 0x52, 0x45, 0x45, 0x5f, 0x50,
	0x48, 0x41, 0x53, 0x45, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x50, 0x4f, 0x57, 0x45,
	0x52, 0x10, 0x21, 0x12, 0x21, 0x0a, 0x1d, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x54, 0x48, 0x52, 0x45, 0x45, 0x5f, 0x50, 0x48, 0x41, 0x53, 0x45, 0x5f, 0x43, 0x55, 0x52,
	0x52, 0x45, 0x4e, 0x54, 0x10, 0x22, 0x12, 0x1f, 0x0a, 0x1b, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x54, 0x48, 0x52, 0x45, 0x45, 0x5f, 0x50, 0x48, 0x41, 0x53, 0x45, 0x5f,
	0x50, 0x4f, 0x57, 0x45, 0x52, 0x10, 0x23, 0x12, 0x26, 0x0a, 0x22, 0x44, 0x41, 0x54, 0x41, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x48, 0x52, 0x45, 0x45, 0x5f, 0x50, 0x48, 0x41, 0x53, 0x45,
	0x5f, 0x50, 0x4f, 0x57, 0x45, 0x52, 0x5f, 0x46, 0x41, 0x43, 0x54, 0x4f, 0x52, 0x10, 0x24, 0x12,
	0x28, 0x0a, 0x24, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x48, 0x52,
	0x45, 0x45, 0x5f, 0x50, 0x48, 0x41, 0x53, 0x45, 0x5f, 0x52, 0x45, 0x41, 0x43, 0x54, 0x49, 0x56,
	0x45, 0x5f, 0x50, 0x4f, 0x57, 0x45, 0x52, 0x10, 0x25, 0x12, 0x1c, 0x0a, 0x18, 0x44, 0x41, 0x54,
	0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x52, 0x49, 0x47, 0x47, 0x45, 0x52, 0x5f, 0x52,
	0x45, 0x41, 0x53, 0x4f, 0x4e, 0x10, 0x26, 0x12, 0x1b, 0x0a, 0x17, 0x44, 0x41, 0x54, 0x41, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x5f, 0x54, 0x49,
	0x4d, 0x45, 0x10, 0x27, 0x12, 0x27, 0x0a, 0x23, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x56, 0x4f, 0x4c, 0x54, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x45, 0x51, 0x55, 0x45, 0x4e,
	0x43, 0x45, 0x5f, 0x4e, 0x45, 0x47, 0x41, 0x54, 0x49, 0x56, 0x45, 0x10, 0x28, 0x12, 0x27, 0x0a,
	0x23, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x56, 0x4f, 0x4c, 0x54, 0x41,
	0x47, 0x45, 0x5f, 0x53, 0x45, 0x51, 0x55, 0x45, 0x4e, 0x43, 0x45, 0x5f, 0x50, 0x4f, 0x53, 0x49,
	0x54, 0x49, 0x56, 0x45, 0x10, 0x29, 0x12, 0x23, 0x0a, 0x1f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x56, 0x4f, 0x4c, 0x54, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x45, 0x51, 0x55,
	0x45, 0x4e, 0x43, 0x45, 0x5f, 0x5a, 0x45, 0x52, 0x4f, 0x10, 0x2a, 0x12, 0x16, 0x0a, 0x12, 0x44,
	0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x49, 0x53, 0x43, 0x52, 0x45, 0x54,
	0x45, 0x10, 0x2b, 0x12, 0x14, 0x0a, 0x10, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x41, 0x4e, 0x41, 0x4c, 0x4f, 0x47, 0x10, 0x2c, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_data_v1_data_proto_rawDescOnce sync.Once
	file_proto_data_v1_data_proto_rawDescData = file_proto_data_v1_data_proto_rawDesc
)

func file_proto_data_v1_data_proto_rawDescGZIP() []byte {
	file_proto_data_v1_data_proto_rawDescOnce.Do(func() {
		file_proto_data_v1_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_data_v1_data_proto_rawDescData)
	})
	return file_proto_data_v1_data_proto_rawDescData
}

var file_proto_data_v1_data_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_data_v1_data_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_data_v1_data_proto_goTypes = []interface{}{
	(DataType)(0),   // 0: data.v1.DataType
	(*Data)(nil),    // 1: data.v1.Data
	(*DataSet)(nil), // 2: data.v1.DataSet
	nil,             // 3: data.v1.DataSet.DataEntry
}
var file_proto_data_v1_data_proto_depIdxs = []int32{
	0, // 0: data.v1.Data.dataType:type_name -> data.v1.DataType
	3, // 1: data.v1.DataSet.data:type_name -> data.v1.DataSet.DataEntry
	1, // 2: data.v1.DataSet.DataEntry.value:type_name -> data.v1.Data
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_data_v1_data_proto_init() }
func file_proto_data_v1_data_proto_init() {
	if File_proto_data_v1_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_data_v1_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_proto_data_v1_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataSet); i {
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
	file_proto_data_v1_data_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_data_v1_data_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_data_v1_data_proto_goTypes,
		DependencyIndexes: file_proto_data_v1_data_proto_depIdxs,
		EnumInfos:         file_proto_data_v1_data_proto_enumTypes,
		MessageInfos:      file_proto_data_v1_data_proto_msgTypes,
	}.Build()
	File_proto_data_v1_data_proto = out.File
	file_proto_data_v1_data_proto_rawDesc = nil
	file_proto_data_v1_data_proto_goTypes = nil
	file_proto_data_v1_data_proto_depIdxs = nil
}
