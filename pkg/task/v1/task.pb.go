// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: proto/task/v1/task.proto

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

type TaskType int32

const (
	TaskType_TASK_TYPE_UNSPECIFIED TaskType = 0
	TaskType_TASK_TYPE_COLLECTION  TaskType = 1
	TaskType_TASK_TYPE_TOPOLOGY    TaskType = 2
	TaskType_TASK_TYPE_STATE       TaskType = 3
	TaskType_TASK_TYPE_FAULT       TaskType = 4
)

// Enum value maps for TaskType.
var (
	TaskType_name = map[int32]string{
		0: "TASK_TYPE_UNSPECIFIED",
		1: "TASK_TYPE_COLLECTION",
		2: "TASK_TYPE_TOPOLOGY",
		3: "TASK_TYPE_STATE",
		4: "TASK_TYPE_FAULT",
	}
	TaskType_value = map[string]int32{
		"TASK_TYPE_UNSPECIFIED": 0,
		"TASK_TYPE_COLLECTION":  1,
		"TASK_TYPE_TOPOLOGY":    2,
		"TASK_TYPE_STATE":       3,
		"TASK_TYPE_FAULT":       4,
	}
)

func (x TaskType) Enum() *TaskType {
	p := new(TaskType)
	*p = x
	return p
}

func (x TaskType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_task_v1_task_proto_enumTypes[0].Descriptor()
}

func (TaskType) Type() protoreflect.EnumType {
	return &file_proto_task_v1_task_proto_enumTypes[0]
}

func (x TaskType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskType.Descriptor instead.
func (TaskType) EnumDescriptor() ([]byte, []int) {
	return file_proto_task_v1_task_proto_rawDescGZIP(), []int{0}
}

type NotificationType int32

const (
	NotificationType_NOTIFICATION_TYPE_UNSPECIFIED    NotificationType = 0
	NotificationType_NOTIFICATION_TYPE_DATA_COMPLETE  NotificationType = 1
	NotificationType_NOTIFICATION_TYPE_DATA_TIMEOUT_1 NotificationType = 2
	NotificationType_NOTIFICATION_TYPE_DATA_TIMEOUT_2 NotificationType = 3
)

// Enum value maps for NotificationType.
var (
	NotificationType_name = map[int32]string{
		0: "NOTIFICATION_TYPE_UNSPECIFIED",
		1: "NOTIFICATION_TYPE_DATA_COMPLETE",
		2: "NOTIFICATION_TYPE_DATA_TIMEOUT_1",
		3: "NOTIFICATION_TYPE_DATA_TIMEOUT_2",
	}
	NotificationType_value = map[string]int32{
		"NOTIFICATION_TYPE_UNSPECIFIED":    0,
		"NOTIFICATION_TYPE_DATA_COMPLETE":  1,
		"NOTIFICATION_TYPE_DATA_TIMEOUT_1": 2,
		"NOTIFICATION_TYPE_DATA_TIMEOUT_2": 3,
	}
)

func (x NotificationType) Enum() *NotificationType {
	p := new(NotificationType)
	*p = x
	return p
}

func (x NotificationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NotificationType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_task_v1_task_proto_enumTypes[1].Descriptor()
}

func (NotificationType) Type() protoreflect.EnumType {
	return &file_proto_task_v1_task_proto_enumTypes[1]
}

func (x NotificationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NotificationType.Descriptor instead.
func (NotificationType) EnumDescriptor() ([]byte, []int) {
	return file_proto_task_v1_task_proto_rawDescGZIP(), []int{1}
}

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TaskType      TaskType `protobuf:"varint,2,opt,name=taskType,proto3,enum=task.v1.TaskType" json:"taskType,omitempty"`
	CreatedAt     int64    `protobuf:"varint,3,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	MeasurementID *string  `protobuf:"bytes,4,opt,name=measurementID,proto3,oneof" json:"measurementID,omitempty"`
	TimestampID   *int64   `protobuf:"varint,5,opt,name=timestampID,proto3,oneof" json:"timestampID,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_task_v1_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_v1_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_proto_task_v1_task_proto_rawDescGZIP(), []int{0}
}

func (x *Task) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Task) GetTaskType() TaskType {
	if x != nil {
		return x.TaskType
	}
	return TaskType_TASK_TYPE_UNSPECIFIED
}

func (x *Task) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Task) GetMeasurementID() string {
	if x != nil && x.MeasurementID != nil {
		return *x.MeasurementID
	}
	return ""
}

func (x *Task) GetTimestampID() int64 {
	if x != nil && x.TimestampID != nil {
		return *x.TimestampID
	}
	return 0
}

type Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	NotificationType NotificationType `protobuf:"varint,2,opt,name=notificationType,proto3,enum=task.v1.NotificationType" json:"notificationType,omitempty"`
	CreatedAt        int64            `protobuf:"varint,3,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	Message          string           `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	MeasurementID    *string          `protobuf:"bytes,5,opt,name=measurementID,proto3,oneof" json:"measurementID,omitempty"`
	TimestampID      *int64           `protobuf:"varint,6,opt,name=timestampID,proto3,oneof" json:"timestampID,omitempty"`
}

func (x *Notification) Reset() {
	*x = Notification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_task_v1_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification) ProtoMessage() {}

func (x *Notification) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_v1_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notification.ProtoReflect.Descriptor instead.
func (*Notification) Descriptor() ([]byte, []int) {
	return file_proto_task_v1_task_proto_rawDescGZIP(), []int{1}
}

func (x *Notification) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Notification) GetNotificationType() NotificationType {
	if x != nil {
		return x.NotificationType
	}
	return NotificationType_NOTIFICATION_TYPE_UNSPECIFIED
}

func (x *Notification) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Notification) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Notification) GetMeasurementID() string {
	if x != nil && x.MeasurementID != nil {
		return *x.MeasurementID
	}
	return ""
}

func (x *Notification) GetTimestampID() int64 {
	if x != nil && x.TimestampID != nil {
		return *x.TimestampID
	}
	return 0
}

var File_proto_task_v1_task_proto protoreflect.FileDescriptor

var file_proto_task_v1_task_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x76, 0x31, 0x2f,
	0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x74, 0x61, 0x73, 0x6b,
	0x2e, 0x76, 0x31, 0x22, 0xd7, 0x01, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2d, 0x0a, 0x08,
	0x74, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11,
	0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x29, 0x0a, 0x0d, 0x6d, 0x65, 0x61,
	0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x0d, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x44, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x0b, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x49, 0x44, 0x88, 0x01, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x5f,
	0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x42, 0x0e, 0x0a,
	0x0c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x49, 0x44, 0x22, 0x91, 0x02,
	0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x45,
	0x0a, 0x10, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e,
	0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x10, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x29, 0x0a,
	0x0d, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0d, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52,
	0x0b, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x49, 0x44, 0x88, 0x01, 0x01, 0x42,
	0x10, 0x0a, 0x0e, 0x5f, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x44, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x49,
	0x44, 0x2a, 0x81, 0x01, 0x0a, 0x08, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19,
	0x0a, 0x15, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x54, 0x41, 0x53,
	0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x4c, 0x4c, 0x45, 0x43, 0x54, 0x49, 0x4f,
	0x4e, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x54, 0x4f, 0x50, 0x4f, 0x4c, 0x4f, 0x47, 0x59, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x54,
	0x41, 0x53, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x10, 0x03,
	0x12, 0x13, 0x0a, 0x0f, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x41,
	0x55, 0x4c, 0x54, 0x10, 0x04, 0x2a, 0xa6, 0x01, 0x0a, 0x10, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x1d, 0x4e, 0x4f,
	0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x23, 0x0a,
	0x1f, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45,
	0x10, 0x01, 0x12, 0x24, 0x0a, 0x20, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x49, 0x4d,
	0x45, 0x4f, 0x55, 0x54, 0x5f, 0x31, 0x10, 0x02, 0x12, 0x24, 0x0a, 0x20, 0x4e, 0x4f, 0x54, 0x49,
	0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x41,
	0x54, 0x41, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x5f, 0x32, 0x10, 0x03, 0x42, 0x0b,
	0x5a, 0x09, 0x2e, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_task_v1_task_proto_rawDescOnce sync.Once
	file_proto_task_v1_task_proto_rawDescData = file_proto_task_v1_task_proto_rawDesc
)

func file_proto_task_v1_task_proto_rawDescGZIP() []byte {
	file_proto_task_v1_task_proto_rawDescOnce.Do(func() {
		file_proto_task_v1_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_task_v1_task_proto_rawDescData)
	})
	return file_proto_task_v1_task_proto_rawDescData
}

var file_proto_task_v1_task_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_task_v1_task_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_task_v1_task_proto_goTypes = []interface{}{
	(TaskType)(0),         // 0: task.v1.TaskType
	(NotificationType)(0), // 1: task.v1.NotificationType
	(*Task)(nil),          // 2: task.v1.Task
	(*Notification)(nil),  // 3: task.v1.Notification
}
var file_proto_task_v1_task_proto_depIdxs = []int32{
	0, // 0: task.v1.Task.taskType:type_name -> task.v1.TaskType
	1, // 1: task.v1.Notification.notificationType:type_name -> task.v1.NotificationType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_task_v1_task_proto_init() }
func file_proto_task_v1_task_proto_init() {
	if File_proto_task_v1_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_task_v1_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_proto_task_v1_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notification); i {
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
	file_proto_task_v1_task_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_proto_task_v1_task_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_task_v1_task_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_task_v1_task_proto_goTypes,
		DependencyIndexes: file_proto_task_v1_task_proto_depIdxs,
		EnumInfos:         file_proto_task_v1_task_proto_enumTypes,
		MessageInfos:      file_proto_task_v1_task_proto_msgTypes,
	}.Build()
	File_proto_task_v1_task_proto = out.File
	file_proto_task_v1_task_proto_rawDesc = nil
	file_proto_task_v1_task_proto_goTypes = nil
	file_proto_task_v1_task_proto_depIdxs = nil
}
