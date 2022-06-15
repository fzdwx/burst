// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: protocol/burst.proto

package protocol

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BurstType int32

const (
	// 添加代理信息
	BurstType_ADD_PROXY_INFO BurstType = 0
	// 关闭
	BurstType_CLOSE BurstType = 1
	// 客户连接
	BurstType_USER_CONNECT BurstType = 2
	// 转发数据
	BurstType_FORWARD_DATA BurstType = 3
	// 删除代理信息
	BurstType_REMOVE_PROXY_INFO BurstType = 4
)

// Enum value maps for BurstType.
var (
	BurstType_name = map[int32]string{
		0: "ADD_PROXY_INFO",
		1: "CLOSE",
		2: "USER_CONNECT",
		3: "FORWARD_DATA",
		4: "REMOVE_PROXY_INFO",
	}
	BurstType_value = map[string]int32{
		"ADD_PROXY_INFO":    0,
		"CLOSE":             1,
		"USER_CONNECT":      2,
		"FORWARD_DATA":      3,
		"REMOVE_PROXY_INFO": 4,
	}
)

func (x BurstType) Enum() *BurstType {
	p := new(BurstType)
	*p = x
	return p
}

func (x BurstType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BurstType) Descriptor() protoreflect.EnumDescriptor {
	return file_protocol_burst_proto_enumTypes[0].Descriptor()
}

func (BurstType) Type() protoreflect.EnumType {
	return &file_protocol_burst_proto_enumTypes[0]
}

func (x BurstType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BurstType.Descriptor instead.
func (BurstType) EnumDescriptor() ([]byte, []int) {
	return file_protocol_burst_proto_rawDescGZIP(), []int{0}
}

type Headers int32

const (
	// 错误信息 string
	Headers_ERROR Headers = 0
	// port mapping Map<int,Proxy>
	Headers_PORTS Headers = 1
	// 服务端暴露的端口 int32
	Headers_SERVER_EXPORT_PORT Headers = 2
	// 外网客户的id string
	Headers_USER_CONNECT_ID Headers = 3
	// register token string
	Headers_TOKEN Headers = 4
)

// Enum value maps for Headers.
var (
	Headers_name = map[int32]string{
		0: "ERROR",
		1: "PORTS",
		2: "SERVER_EXPORT_PORT",
		3: "USER_CONNECT_ID",
		4: "TOKEN",
	}
	Headers_value = map[string]int32{
		"ERROR":              0,
		"PORTS":              1,
		"SERVER_EXPORT_PORT": 2,
		"USER_CONNECT_ID":    3,
		"TOKEN":              4,
	}
)

func (x Headers) Enum() *Headers {
	p := new(Headers)
	*p = x
	return p
}

func (x Headers) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Headers) Descriptor() protoreflect.EnumDescriptor {
	return file_protocol_burst_proto_enumTypes[1].Descriptor()
}

func (Headers) Type() protoreflect.EnumType {
	return &file_protocol_burst_proto_enumTypes[1]
}

func (x Headers) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Headers.Descriptor instead.
func (Headers) EnumDescriptor() ([]byte, []int) {
	return file_protocol_burst_proto_rawDescGZIP(), []int{1}
}

type BurstMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   BurstType            `protobuf:"varint,1,opt,name=type,proto3,enum=protocol.BurstType" json:"type,omitempty"`
	Header map[int32]*anypb.Any `protobuf:"bytes,2,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Data   []byte               `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	// 当type = REMOVE_PROXY_INFO 时有效
	ServerPort []int32 `protobuf:"varint,4,rep,packed,name=serverPort,proto3" json:"serverPort,omitempty"`
}

func (x *BurstMessage) Reset() {
	*x = BurstMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_burst_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BurstMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BurstMessage) ProtoMessage() {}

func (x *BurstMessage) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_burst_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BurstMessage.ProtoReflect.Descriptor instead.
func (*BurstMessage) Descriptor() ([]byte, []int) {
	return file_protocol_burst_proto_rawDescGZIP(), []int{0}
}

func (x *BurstMessage) GetType() BurstType {
	if x != nil {
		return x.Type
	}
	return BurstType_ADD_PROXY_INFO
}

func (x *BurstMessage) GetHeader() map[int32]*anypb.Any {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *BurstMessage) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *BurstMessage) GetServerPort() []int32 {
	if x != nil {
		return x.ServerPort
	}
	return nil
}

type Proxy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip   string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Port int32  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *Proxy) Reset() {
	*x = Proxy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_burst_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Proxy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Proxy) ProtoMessage() {}

func (x *Proxy) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_burst_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Proxy.ProtoReflect.Descriptor instead.
func (*Proxy) Descriptor() ([]byte, []int) {
	return file_protocol_burst_proto_rawDescGZIP(), []int{1}
}

func (x *Proxy) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *Proxy) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type Ports struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ports map[int32]*Proxy `protobuf:"bytes,1,rep,name=ports,proto3" json:"ports,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Ports) Reset() {
	*x = Ports{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_burst_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ports) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ports) ProtoMessage() {}

func (x *Ports) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_burst_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ports.ProtoReflect.Descriptor instead.
func (*Ports) Descriptor() ([]byte, []int) {
	return file_protocol_burst_proto_rawDescGZIP(), []int{2}
}

func (x *Ports) GetPorts() map[int32]*Proxy {
	if x != nil {
		return x.Ports
	}
	return nil
}

var File_protocol_burst_proto protoreflect.FileDescriptor

var file_protocol_burst_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x62, 0x75, 0x72, 0x73, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf8, 0x01, 0x0a, 0x0c,
	0x42, 0x75, 0x72, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x27, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x42, 0x75, 0x72, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x42, 0x75, 0x72, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50,
	0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x50, 0x6f, 0x72, 0x74, 0x1a, 0x4f, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x2b, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70,
	0x6f, 0x72, 0x74, 0x22, 0x84, 0x01, 0x0a, 0x05, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x30, 0x0a,
	0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x2e, 0x50, 0x6f,
	0x72, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x1a,
	0x49, 0x0a, 0x0a, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x25, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x65, 0x0a, 0x09, 0x42, 0x75,
	0x72, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x44, 0x44, 0x5f, 0x50,
	0x52, 0x4f, 0x58, 0x59, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x43,
	0x4c, 0x4f, 0x53, 0x45, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x43,
	0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x46, 0x4f, 0x52, 0x57,
	0x41, 0x52, 0x44, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x52, 0x45,
	0x4d, 0x4f, 0x56, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x58, 0x59, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x10,
	0x04, 0x2a, 0x57, 0x0a, 0x07, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x09, 0x0a, 0x05,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x4f, 0x52, 0x54, 0x53,
	0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x45, 0x58, 0x50,
	0x4f, 0x52, 0x54, 0x5f, 0x50, 0x4f, 0x52, 0x54, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x55, 0x53,
	0x45, 0x52, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x5f, 0x49, 0x44, 0x10, 0x03, 0x12,
	0x09, 0x0a, 0x05, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0x04, 0x42, 0x29, 0x0a, 0x0e, 0x62, 0x75,
	0x72, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x42, 0x05, 0x42, 0x75,
	0x72, 0x73, 0x74, 0x50, 0x01, 0x5a, 0x0e, 0x62, 0x75, 0x72, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocol_burst_proto_rawDescOnce sync.Once
	file_protocol_burst_proto_rawDescData = file_protocol_burst_proto_rawDesc
)

func file_protocol_burst_proto_rawDescGZIP() []byte {
	file_protocol_burst_proto_rawDescOnce.Do(func() {
		file_protocol_burst_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_burst_proto_rawDescData)
	})
	return file_protocol_burst_proto_rawDescData
}

var file_protocol_burst_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_protocol_burst_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protocol_burst_proto_goTypes = []interface{}{
	(BurstType)(0),       // 0: protocol.BurstType
	(Headers)(0),         // 1: protocol.Headers
	(*BurstMessage)(nil), // 2: protocol.BurstMessage
	(*Proxy)(nil),        // 3: protocol.Proxy
	(*Ports)(nil),        // 4: protocol.Ports
	nil,                  // 5: protocol.BurstMessage.HeaderEntry
	nil,                  // 6: protocol.Ports.PortsEntry
	(*anypb.Any)(nil),    // 7: google.protobuf.Any
}
var file_protocol_burst_proto_depIdxs = []int32{
	0, // 0: protocol.BurstMessage.type:type_name -> protocol.BurstType
	5, // 1: protocol.BurstMessage.header:type_name -> protocol.BurstMessage.HeaderEntry
	6, // 2: protocol.Ports.ports:type_name -> protocol.Ports.PortsEntry
	7, // 3: protocol.BurstMessage.HeaderEntry.value:type_name -> google.protobuf.Any
	3, // 4: protocol.Ports.PortsEntry.value:type_name -> protocol.Proxy
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_protocol_burst_proto_init() }
func file_protocol_burst_proto_init() {
	if File_protocol_burst_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocol_burst_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BurstMessage); i {
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
		file_protocol_burst_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Proxy); i {
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
		file_protocol_burst_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ports); i {
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
			RawDescriptor: file_protocol_burst_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protocol_burst_proto_goTypes,
		DependencyIndexes: file_protocol_burst_proto_depIdxs,
		EnumInfos:         file_protocol_burst_proto_enumTypes,
		MessageInfos:      file_protocol_burst_proto_msgTypes,
	}.Build()
	File_protocol_burst_proto = out.File
	file_protocol_burst_proto_rawDesc = nil
	file_protocol_burst_proto_goTypes = nil
	file_protocol_burst_proto_depIdxs = nil
}