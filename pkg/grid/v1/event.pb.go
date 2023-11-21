// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: grid/v1/event.proto

//  <!-- markdownlint-disable -->
//Messages to support event detection in the platform.
//The event detected may be originated from different sources: devices (e.g. a PMU, RTU), services (e.g. state estimator), or an external service (e.g. SCADA).

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

type EventStatus int32

const (
	EventStatus_EVENT_STATUS_UNSPECIFIED EventStatus = 0 //No status defined
	EventStatus_EVENT_STATUS_ACTIVE      EventStatus = 1 //Event is still active
	EventStatus_EVENT_STATUS_COMPLETE    EventStatus = 2 //Event is completed
	EventStatus_EVENT_STATUS_UNKNOWN     EventStatus = 3 //Information available don't allow us to know if the even is active or complete
)

// Enum value maps for EventStatus.
var (
	EventStatus_name = map[int32]string{
		0: "EVENT_STATUS_UNSPECIFIED",
		1: "EVENT_STATUS_ACTIVE",
		2: "EVENT_STATUS_COMPLETE",
		3: "EVENT_STATUS_UNKNOWN",
	}
	EventStatus_value = map[string]int32{
		"EVENT_STATUS_UNSPECIFIED": 0,
		"EVENT_STATUS_ACTIVE":      1,
		"EVENT_STATUS_COMPLETE":    2,
		"EVENT_STATUS_UNKNOWN":     3,
	}
)

func (x EventStatus) Enum() *EventStatus {
	p := new(EventStatus)
	*p = x
	return p
}

func (x EventStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_grid_v1_event_proto_enumTypes[0].Descriptor()
}

func (EventStatus) Type() protoreflect.EnumType {
	return &file_grid_v1_event_proto_enumTypes[0]
}

func (x EventStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventStatus.Descriptor instead.
func (EventStatus) EnumDescriptor() ([]byte, []int) {
	return file_grid_v1_event_proto_rawDescGZIP(), []int{0}
}

type EventSourceType int32

const (
	EventSourceType_EVENT_SOURCE_UNSPECIFIED      EventSourceType = 0 //No source type defined
	EventSourceType_EVENT_SOURCE_DEVICE           EventSourceType = 1 //The source of the event was a device (e.g. PMU)
	EventSourceType_EVENT_SOURCE_SERVICE          EventSourceType = 2 //The source of the event was a service (e.g. state estimator)
	EventSourceType_EVENT_SOURCE_EXTERNAL_SERVICE EventSourceType = 3 //The source of the event was a service external to SynchroGuard platform (e.g. SCADA)
)

// Enum value maps for EventSourceType.
var (
	EventSourceType_name = map[int32]string{
		0: "EVENT_SOURCE_UNSPECIFIED",
		1: "EVENT_SOURCE_DEVICE",
		2: "EVENT_SOURCE_SERVICE",
		3: "EVENT_SOURCE_EXTERNAL_SERVICE",
	}
	EventSourceType_value = map[string]int32{
		"EVENT_SOURCE_UNSPECIFIED":      0,
		"EVENT_SOURCE_DEVICE":           1,
		"EVENT_SOURCE_SERVICE":          2,
		"EVENT_SOURCE_EXTERNAL_SERVICE": 3,
	}
)

func (x EventSourceType) Enum() *EventSourceType {
	p := new(EventSourceType)
	*p = x
	return p
}

func (x EventSourceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventSourceType) Descriptor() protoreflect.EnumDescriptor {
	return file_grid_v1_event_proto_enumTypes[1].Descriptor()
}

func (EventSourceType) Type() protoreflect.EnumType {
	return &file_grid_v1_event_proto_enumTypes[1]
}

func (x EventSourceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventSourceType.Descriptor instead.
func (EventSourceType) EnumDescriptor() ([]byte, []int) {
	return file_grid_v1_event_proto_rawDescGZIP(), []int{1}
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceType EventSourceType `protobuf:"varint,1,opt,name=sourceType,proto3,enum=grid.v1.EventSourceType" json:"sourceType,omitempty"` //The type of data see `DataType` enum.
	OccurredAt *int64          `protobuf:"varint,2,opt,name=occurredAt,proto3,oneof" json:"occurredAt,omitempty"`                        //The time of occurency of the event (Unix msec timestamp).
	DetectedAt *int64          `protobuf:"varint,3,opt,name=detectedAt,proto3,oneof" json:"detectedAt,omitempty"`                        //The time of detection of the event (Unix msec timestamp).
	Value      *uint64         `protobuf:"varint,4,opt,name=value,proto3,oneof" json:"value,omitempty"`                                  //The data value casted to uint64.
	Message    string          `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`                                     //Event message.
	Status     *EventStatus    `protobuf:"varint,6,opt,name=status,proto3,enum=grid.v1.EventStatus,oneof" json:"status,omitempty"`       //The status of the event.
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grid_v1_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_grid_v1_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_grid_v1_event_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetSourceType() EventSourceType {
	if x != nil {
		return x.SourceType
	}
	return EventSourceType_EVENT_SOURCE_UNSPECIFIED
}

func (x *Event) GetOccurredAt() int64 {
	if x != nil && x.OccurredAt != nil {
		return *x.OccurredAt
	}
	return 0
}

func (x *Event) GetDetectedAt() int64 {
	if x != nil && x.DetectedAt != nil {
		return *x.DetectedAt
	}
	return 0
}

func (x *Event) GetValue() uint64 {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return 0
}

func (x *Event) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Event) GetStatus() EventStatus {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return EventStatus_EVENT_STATUS_UNSPECIFIED
}

var File_grid_v1_event_proto protoreflect.FileDescriptor

var file_grid_v1_event_proto_rawDesc = []byte{
	0x0a, 0x13, 0x67, 0x72, 0x69, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x67, 0x72, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x22, 0xa6,
	0x02, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x67,
	0x72, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x23, 0x0a, 0x0a, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x0a, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x64, 0x65, 0x74, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x0a, 0x64,
	0x65, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x48, 0x02, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x88, 0x01, 0x01, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x31, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x14, 0x2e, 0x67, 0x72, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x64, 0x41, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x79, 0x0a, 0x0b, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x18, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x19, 0x0a,
	0x15, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f,
	0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x56, 0x45, 0x4e,
	0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x03, 0x2a, 0x85, 0x01, 0x0a, 0x0f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f,
	0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x4f,
	0x55, 0x52, 0x43, 0x45, 0x5f, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x10, 0x01, 0x12, 0x18, 0x0a,
	0x14, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x53, 0x45,
	0x52, 0x56, 0x49, 0x43, 0x45, 0x10, 0x02, 0x12, 0x21, 0x0a, 0x1d, 0x45, 0x56, 0x45, 0x4e, 0x54,
	0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x45, 0x58, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c,
	0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x10, 0x03, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f,
	0x67, 0x72, 0x69, 0x64, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grid_v1_event_proto_rawDescOnce sync.Once
	file_grid_v1_event_proto_rawDescData = file_grid_v1_event_proto_rawDesc
)

func file_grid_v1_event_proto_rawDescGZIP() []byte {
	file_grid_v1_event_proto_rawDescOnce.Do(func() {
		file_grid_v1_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_grid_v1_event_proto_rawDescData)
	})
	return file_grid_v1_event_proto_rawDescData
}

var file_grid_v1_event_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_grid_v1_event_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_grid_v1_event_proto_goTypes = []interface{}{
	(EventStatus)(0),     // 0: grid.v1.EventStatus
	(EventSourceType)(0), // 1: grid.v1.EventSourceType
	(*Event)(nil),        // 2: grid.v1.Event
}
var file_grid_v1_event_proto_depIdxs = []int32{
	1, // 0: grid.v1.Event.sourceType:type_name -> grid.v1.EventSourceType
	0, // 1: grid.v1.Event.status:type_name -> grid.v1.EventStatus
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_grid_v1_event_proto_init() }
func file_grid_v1_event_proto_init() {
	if File_grid_v1_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grid_v1_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
	file_grid_v1_event_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grid_v1_event_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_grid_v1_event_proto_goTypes,
		DependencyIndexes: file_grid_v1_event_proto_depIdxs,
		EnumInfos:         file_grid_v1_event_proto_enumTypes,
		MessageInfos:      file_grid_v1_event_proto_msgTypes,
	}.Build()
	File_grid_v1_event_proto = out.File
	file_grid_v1_event_proto_rawDesc = nil
	file_grid_v1_event_proto_goTypes = nil
	file_grid_v1_event_proto_depIdxs = nil
}
