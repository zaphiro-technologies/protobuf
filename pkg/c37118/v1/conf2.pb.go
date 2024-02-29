// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: c37118/v1/conf2.proto

// <!-- markdownlint-disable -->
//Messages describing PMU C37.118 Configurations.
//See [C37.118](https://www.typhoon-hil.com/documentation/typhoon-hil-software-manual/References/c37_118_protocol.html)
//protocol.

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

type Conf2Frame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header    *Conf2Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`                      //Configuration frame 2 header
	Configs   []*Config    `protobuf:"bytes,2,rep,name=configs,proto3" json:"configs,omitempty"`                    //Set of PMU configurations
	DATA_RATE uint32       `protobuf:"varint,3,opt,name=DATA_RATE,json=DATARATE,proto3" json:"DATA_RATE,omitempty"` //Rate of data transmission
}

func (x *Conf2Frame) Reset() {
	*x = Conf2Frame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_c37118_v1_conf2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conf2Frame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conf2Frame) ProtoMessage() {}

func (x *Conf2Frame) ProtoReflect() protoreflect.Message {
	mi := &file_c37118_v1_conf2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Conf2Frame.ProtoReflect.Descriptor instead.
func (*Conf2Frame) Descriptor() ([]byte, []int) {
	return file_c37118_v1_conf2_proto_rawDescGZIP(), []int{0}
}

func (x *Conf2Frame) GetHeader() *Conf2Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Conf2Frame) GetConfigs() []*Config {
	if x != nil {
		return x.Configs
	}
	return nil
}

func (x *Conf2Frame) GetDATA_RATE() uint32 {
	if x != nil {
		return x.DATA_RATE
	}
	return 0
}

type Conf2Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SYNC      uint32 `protobuf:"varint,1,opt,name=SYNC,proto3" json:"SYNC,omitempty"`                         //Sync byte followed by frame type and version number
	FRAMESIZE uint32 `protobuf:"varint,2,opt,name=FRAMESIZE,proto3" json:"FRAMESIZE,omitempty"`               //Number of bytes in the frame
	IDCODE    uint32 `protobuf:"varint,3,opt,name=IDCODE,proto3" json:"IDCODE,omitempty"`                     //Stream source ID number
	SOC       uint32 `protobuf:"varint,4,opt,name=SOC,proto3" json:"SOC,omitempty"`                           //SOC time stamp
	FRACSEC   uint32 `protobuf:"varint,5,opt,name=FRACSEC,proto3" json:"FRACSEC,omitempty"`                   //Fraction of Second and Message Time Quality
	TIME_BASE uint32 `protobuf:"varint,6,opt,name=TIME_BASE,json=TIMEBASE,proto3" json:"TIME_BASE,omitempty"` //Resolution of FRACSEC time stamp
	NUM_PMU   uint32 `protobuf:"varint,7,opt,name=NUM_PMU,json=NUMPMU,proto3" json:"NUM_PMU,omitempty"`       //The number of PMUs included in the data frame
}

func (x *Conf2Header) Reset() {
	*x = Conf2Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_c37118_v1_conf2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conf2Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conf2Header) ProtoMessage() {}

func (x *Conf2Header) ProtoReflect() protoreflect.Message {
	mi := &file_c37118_v1_conf2_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Conf2Header.ProtoReflect.Descriptor instead.
func (*Conf2Header) Descriptor() ([]byte, []int) {
	return file_c37118_v1_conf2_proto_rawDescGZIP(), []int{1}
}

func (x *Conf2Header) GetSYNC() uint32 {
	if x != nil {
		return x.SYNC
	}
	return 0
}

func (x *Conf2Header) GetFRAMESIZE() uint32 {
	if x != nil {
		return x.FRAMESIZE
	}
	return 0
}

func (x *Conf2Header) GetIDCODE() uint32 {
	if x != nil {
		return x.IDCODE
	}
	return 0
}

func (x *Conf2Header) GetSOC() uint32 {
	if x != nil {
		return x.SOC
	}
	return 0
}

func (x *Conf2Header) GetFRACSEC() uint32 {
	if x != nil {
		return x.FRACSEC
	}
	return 0
}

func (x *Conf2Header) GetTIME_BASE() uint32 {
	if x != nil {
		return x.TIME_BASE
	}
	return 0
}

func (x *Conf2Header) GetNUM_PMU() uint32 {
	if x != nil {
		return x.NUM_PMU
	}
	return 0
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	STN     string   `protobuf:"bytes,1,opt,name=STN,proto3" json:"STN,omitempty"`                  //Station name
	IDCODE  uint32   `protobuf:"varint,2,opt,name=IDCODE,proto3" json:"IDCODE,omitempty"`           //Data source ID number
	FORMAT  uint32   `protobuf:"varint,3,opt,name=FORMAT,proto3" json:"FORMAT,omitempty"`           //Data format within data frame
	PHNMR   uint32   `protobuf:"varint,4,opt,name=PHNMR,proto3" json:"PHNMR,omitempty"`             //Number of phasors
	ANNMR   uint32   `protobuf:"varint,5,opt,name=ANNMR,proto3" json:"ANNMR,omitempty"`             //Number of analog values
	DGNMR   uint32   `protobuf:"varint,6,opt,name=DGNMR,proto3" json:"DGNMR,omitempty"`             //Number of digital status words
	CHNAM   string   `protobuf:"bytes,7,opt,name=CHNAM,proto3" json:"CHNAM,omitempty"`              //Phasor and channel names
	PHUNIT  []uint32 `protobuf:"varint,8,rep,packed,name=PHUNIT,proto3" json:"PHUNIT,omitempty"`    //Conversion factor for phasor channels
	ANUNIT  []uint32 `protobuf:"varint,9,rep,packed,name=ANUNIT,proto3" json:"ANUNIT,omitempty"`    //Conversion factor for analog channels
	DIGUNIT []uint32 `protobuf:"varint,10,rep,packed,name=DIGUNIT,proto3" json:"DIGUNIT,omitempty"` //Mask words for digital status words
	FNOM    uint32   `protobuf:"varint,11,opt,name=FNOM,proto3" json:"FNOM,omitempty"`              //Nominal line frequency code and flags
	CFGCNT  uint32   `protobuf:"varint,12,opt,name=CFGCNT,proto3" json:"CFGCNT,omitempty"`          //Configuration change count
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_c37118_v1_conf2_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_c37118_v1_conf2_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_c37118_v1_conf2_proto_rawDescGZIP(), []int{2}
}

func (x *Config) GetSTN() string {
	if x != nil {
		return x.STN
	}
	return ""
}

func (x *Config) GetIDCODE() uint32 {
	if x != nil {
		return x.IDCODE
	}
	return 0
}

func (x *Config) GetFORMAT() uint32 {
	if x != nil {
		return x.FORMAT
	}
	return 0
}

func (x *Config) GetPHNMR() uint32 {
	if x != nil {
		return x.PHNMR
	}
	return 0
}

func (x *Config) GetANNMR() uint32 {
	if x != nil {
		return x.ANNMR
	}
	return 0
}

func (x *Config) GetDGNMR() uint32 {
	if x != nil {
		return x.DGNMR
	}
	return 0
}

func (x *Config) GetCHNAM() string {
	if x != nil {
		return x.CHNAM
	}
	return ""
}

func (x *Config) GetPHUNIT() []uint32 {
	if x != nil {
		return x.PHUNIT
	}
	return nil
}

func (x *Config) GetANUNIT() []uint32 {
	if x != nil {
		return x.ANUNIT
	}
	return nil
}

func (x *Config) GetDIGUNIT() []uint32 {
	if x != nil {
		return x.DIGUNIT
	}
	return nil
}

func (x *Config) GetFNOM() uint32 {
	if x != nil {
		return x.FNOM
	}
	return 0
}

func (x *Config) GetCFGCNT() uint32 {
	if x != nil {
		return x.CFGCNT
	}
	return 0
}

var File_c37118_v1_conf2_proto protoreflect.FileDescriptor

var file_c37118_v1_conf2_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x33, 0x37, 0x31, 0x31, 0x38, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x33, 0x37, 0x31, 0x31, 0x38, 0x2e,
	0x76, 0x31, 0x22, 0x86, 0x01, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x66, 0x32, 0x46, 0x72, 0x61, 0x6d,
	0x65, 0x12, 0x2e, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x63, 0x33, 0x37, 0x31, 0x31, 0x38, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x32, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x12, 0x2b, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x33, 0x37, 0x31, 0x31, 0x38, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x1b,
	0x0a, 0x09, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x52, 0x41, 0x54, 0x45, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x44, 0x41, 0x54, 0x41, 0x52, 0x41, 0x54, 0x45, 0x22, 0xb9, 0x01, 0x0a, 0x0b,
	0x43, 0x6f, 0x6e, 0x66, 0x32, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x53,
	0x59, 0x4e, 0x43, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x53, 0x59, 0x4e, 0x43, 0x12,
	0x1c, 0x0a, 0x09, 0x46, 0x52, 0x41, 0x4d, 0x45, 0x53, 0x49, 0x5a, 0x45, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x46, 0x52, 0x41, 0x4d, 0x45, 0x53, 0x49, 0x5a, 0x45, 0x12, 0x16, 0x0a,
	0x06, 0x49, 0x44, 0x43, 0x4f, 0x44, 0x45, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x49,
	0x44, 0x43, 0x4f, 0x44, 0x45, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x4f, 0x43, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x03, 0x53, 0x4f, 0x43, 0x12, 0x18, 0x0a, 0x07, 0x46, 0x52, 0x41, 0x43, 0x53,
	0x45, 0x43, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x46, 0x52, 0x41, 0x43, 0x53, 0x45,
	0x43, 0x12, 0x1b, 0x0a, 0x09, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x42, 0x41, 0x53, 0x45, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x54, 0x49, 0x4d, 0x45, 0x42, 0x41, 0x53, 0x45, 0x12, 0x17,
	0x0a, 0x07, 0x4e, 0x55, 0x4d, 0x5f, 0x50, 0x4d, 0x55, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x4e, 0x55, 0x4d, 0x50, 0x4d, 0x55, 0x22, 0x98, 0x02, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x54, 0x4e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x53, 0x54, 0x4e, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x44, 0x43, 0x4f, 0x44, 0x45, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x49, 0x44, 0x43, 0x4f, 0x44, 0x45, 0x12, 0x16, 0x0a, 0x06,
	0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x46, 0x4f,
	0x52, 0x4d, 0x41, 0x54, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x48, 0x4e, 0x4d, 0x52, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x50, 0x48, 0x4e, 0x4d, 0x52, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x4e,
	0x4e, 0x4d, 0x52, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x41, 0x4e, 0x4e, 0x4d, 0x52,
	0x12, 0x14, 0x0a, 0x05, 0x44, 0x47, 0x4e, 0x4d, 0x52, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x44, 0x47, 0x4e, 0x4d, 0x52, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x48, 0x4e, 0x41, 0x4d, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x43, 0x48, 0x4e, 0x41, 0x4d, 0x12, 0x16, 0x0a, 0x06,
	0x50, 0x48, 0x55, 0x4e, 0x49, 0x54, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x06, 0x50, 0x48,
	0x55, 0x4e, 0x49, 0x54, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x4e, 0x55, 0x4e, 0x49, 0x54, 0x18, 0x09,
	0x20, 0x03, 0x28, 0x0d, 0x52, 0x06, 0x41, 0x4e, 0x55, 0x4e, 0x49, 0x54, 0x12, 0x18, 0x0a, 0x07,
	0x44, 0x49, 0x47, 0x55, 0x4e, 0x49, 0x54, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x07, 0x44,
	0x49, 0x47, 0x55, 0x4e, 0x49, 0x54, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x4e, 0x4f, 0x4d, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x46, 0x4e, 0x4f, 0x4d, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x46,
	0x47, 0x43, 0x4e, 0x54, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x43, 0x46, 0x47, 0x43,
	0x4e, 0x54, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x63, 0x33, 0x37, 0x31, 0x31, 0x38, 0x2f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_c37118_v1_conf2_proto_rawDescOnce sync.Once
	file_c37118_v1_conf2_proto_rawDescData = file_c37118_v1_conf2_proto_rawDesc
)

func file_c37118_v1_conf2_proto_rawDescGZIP() []byte {
	file_c37118_v1_conf2_proto_rawDescOnce.Do(func() {
		file_c37118_v1_conf2_proto_rawDescData = protoimpl.X.CompressGZIP(file_c37118_v1_conf2_proto_rawDescData)
	})
	return file_c37118_v1_conf2_proto_rawDescData
}

var file_c37118_v1_conf2_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_c37118_v1_conf2_proto_goTypes = []interface{}{
	(*Conf2Frame)(nil),  // 0: c37118.v1.Conf2Frame
	(*Conf2Header)(nil), // 1: c37118.v1.Conf2Header
	(*Config)(nil),      // 2: c37118.v1.Config
}
var file_c37118_v1_conf2_proto_depIdxs = []int32{
	1, // 0: c37118.v1.Conf2Frame.header:type_name -> c37118.v1.Conf2Header
	2, // 1: c37118.v1.Conf2Frame.configs:type_name -> c37118.v1.Config
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_c37118_v1_conf2_proto_init() }
func file_c37118_v1_conf2_proto_init() {
	if File_c37118_v1_conf2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_c37118_v1_conf2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Conf2Frame); i {
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
		file_c37118_v1_conf2_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Conf2Header); i {
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
		file_c37118_v1_conf2_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_c37118_v1_conf2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_c37118_v1_conf2_proto_goTypes,
		DependencyIndexes: file_c37118_v1_conf2_proto_depIdxs,
		MessageInfos:      file_c37118_v1_conf2_proto_msgTypes,
	}.Build()
	File_c37118_v1_conf2_proto = out.File
	file_c37118_v1_conf2_proto_rawDesc = nil
	file_c37118_v1_conf2_proto_goTypes = nil
	file_c37118_v1_conf2_proto_depIdxs = nil
}