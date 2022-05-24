// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protocol/burst.proto

package protocol

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type BurstType int32

const (
	// 初始化
	BurstType_INIT BurstType = 0
	// 关闭
	BurstType_CLOSE BurstType = 1
	// 客户连接
	BurstType_USER_CONNECT BurstType = 2
	// 转发数据
	BurstType_FORWARD_DATA BurstType = 3
)

var BurstType_name = map[int32]string{
	0: "INIT",
	1: "CLOSE",
	2: "USER_CONNECT",
	3: "FORWARD_DATA",
}

var BurstType_value = map[string]int32{
	"INIT":         0,
	"CLOSE":        1,
	"USER_CONNECT": 2,
	"FORWARD_DATA": 3,
}

func (x BurstType) String() string {
	return proto.EnumName(BurstType_name, int32(x))
}

func (BurstType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_187b98f11cb75ff0, []int{0}
}

type Headers int32

const (
	// 错误信息 string
	Headers_ERROR Headers = 0
	// port mapping Map<int,int>
	Headers_PORTS Headers = 1
	// 服务端暴露的端口 int32
	Headers_SERVER_EXPORT_PORT Headers = 2
	// 外网客户的id string
	Headers_USER_CONNECT_ID Headers = 3
)

var Headers_name = map[int32]string{
	0: "ERROR",
	1: "PORTS",
	2: "SERVER_EXPORT_PORT",
	3: "USER_CONNECT_ID",
}

var Headers_value = map[string]int32{
	"ERROR":              0,
	"PORTS":              1,
	"SERVER_EXPORT_PORT": 2,
	"USER_CONNECT_ID":    3,
}

func (x Headers) String() string {
	return proto.EnumName(Headers_name, int32(x))
}

func (Headers) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_187b98f11cb75ff0, []int{1}
}

type BurstMessage struct {
	Type                 BurstType          `protobuf:"varint,1,opt,name=type,proto3,enum=protocol.BurstType" json:"type,omitempty"`
	Header               map[int32]*any.Any `protobuf:"bytes,2,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Data                 []byte             `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *BurstMessage) Reset()         { *m = BurstMessage{} }
func (m *BurstMessage) String() string { return proto.CompactTextString(m) }
func (*BurstMessage) ProtoMessage()    {}
func (*BurstMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_187b98f11cb75ff0, []int{0}
}

func (m *BurstMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BurstMessage.Unmarshal(m, b)
}
func (m *BurstMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BurstMessage.Marshal(b, m, deterministic)
}
func (m *BurstMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BurstMessage.Merge(m, src)
}
func (m *BurstMessage) XXX_Size() int {
	return xxx_messageInfo_BurstMessage.Size(m)
}
func (m *BurstMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_BurstMessage.DiscardUnknown(m)
}

var xxx_messageInfo_BurstMessage proto.InternalMessageInfo

func (m *BurstMessage) GetType() BurstType {
	if m != nil {
		return m.Type
	}
	return BurstType_INIT
}

func (m *BurstMessage) GetHeader() map[int32]*any.Any {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *BurstMessage) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Ports struct {
	Ports                map[int32]int32 `protobuf:"bytes,1,rep,name=ports,proto3" json:"ports,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Ports) Reset()         { *m = Ports{} }
func (m *Ports) String() string { return proto.CompactTextString(m) }
func (*Ports) ProtoMessage()    {}
func (*Ports) Descriptor() ([]byte, []int) {
	return fileDescriptor_187b98f11cb75ff0, []int{1}
}

func (m *Ports) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ports.Unmarshal(m, b)
}
func (m *Ports) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ports.Marshal(b, m, deterministic)
}
func (m *Ports) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ports.Merge(m, src)
}
func (m *Ports) XXX_Size() int {
	return xxx_messageInfo_Ports.Size(m)
}
func (m *Ports) XXX_DiscardUnknown() {
	xxx_messageInfo_Ports.DiscardUnknown(m)
}

var xxx_messageInfo_Ports proto.InternalMessageInfo

func (m *Ports) GetPorts() map[int32]int32 {
	if m != nil {
		return m.Ports
	}
	return nil
}

func init() {
	proto.RegisterEnum("protocol.BurstType", BurstType_name, BurstType_value)
	proto.RegisterEnum("protocol.Headers", Headers_name, Headers_value)
	proto.RegisterType((*BurstMessage)(nil), "protocol.BurstMessage")
	proto.RegisterMapType((map[int32]*any.Any)(nil), "protocol.BurstMessage.HeaderEntry")
	proto.RegisterType((*Ports)(nil), "protocol.Ports")
	proto.RegisterMapType((map[int32]int32)(nil), "protocol.Ports.PortsEntry")
}

func init() { proto.RegisterFile("protocol/burst.proto", fileDescriptor_187b98f11cb75ff0) }

var fileDescriptor_187b98f11cb75ff0 = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4f, 0x8f, 0xa2, 0x40,
	0x10, 0xc5, 0xb7, 0x41, 0x5c, 0x2d, 0x8d, 0xdb, 0x69, 0xcd, 0x06, 0x39, 0x11, 0x2f, 0x4b, 0x3c,
	0xe0, 0xc6, 0xbd, 0x18, 0x6f, 0x28, 0xbd, 0x59, 0x13, 0x57, 0x4c, 0xc3, 0xfc, 0xb9, 0x11, 0xd4,
	0x1e, 0x27, 0x19, 0x22, 0x04, 0x70, 0x12, 0xbe, 0xe9, 0x7c, 0x9c, 0x09, 0xdd, 0x3a, 0x92, 0xc9,
	0x5c, 0x2a, 0xc5, 0xe3, 0x57, 0x8f, 0x57, 0x15, 0x60, 0x90, 0x66, 0x49, 0x91, 0xec, 0x93, 0x78,
	0xb2, 0x3b, 0x67, 0x79, 0x61, 0x8b, 0x47, 0xd2, 0xba, 0xaa, 0xc6, 0xf0, 0x98, 0x24, 0xc7, 0x98,
	0x4f, 0x84, 0xb0, 0x3b, 0x3f, 0x4d, 0xa2, 0x53, 0x29, 0xa1, 0xd1, 0x1b, 0x82, 0xee, 0xa2, 0x1a,
	0xfa, 0xcf, 0xf3, 0x3c, 0x3a, 0x72, 0xf2, 0x0b, 0x1a, 0x45, 0x99, 0x72, 0x1d, 0x99, 0xc8, 0xea,
	0x4d, 0xfb, 0xf6, 0xd5, 0xc4, 0x16, 0x54, 0x50, 0xa6, 0x9c, 0x09, 0x80, 0xcc, 0xa1, 0xf9, 0xcc,
	0xa3, 0x03, 0xcf, 0x74, 0xc5, 0x54, 0xad, 0xce, 0x74, 0xf4, 0x09, 0xbd, 0x18, 0xda, 0xff, 0x04,
	0x44, 0x4f, 0x45, 0x56, 0xb2, 0xcb, 0x04, 0x21, 0xd0, 0x38, 0x44, 0x45, 0xa4, 0xab, 0x26, 0xb2,
	0xba, 0x4c, 0xf4, 0x86, 0x07, 0x9d, 0x1a, 0x4a, 0x30, 0xa8, 0x2f, 0xbc, 0x14, 0x31, 0x34, 0x56,
	0xb5, 0x64, 0x0c, 0xda, 0x6b, 0x14, 0x9f, 0xb9, 0xae, 0x98, 0xc8, 0xea, 0x4c, 0x07, 0xb6, 0xdc,
	0xca, 0xbe, 0x6e, 0x65, 0x3b, 0xa7, 0x92, 0x49, 0x64, 0xae, 0xcc, 0xd0, 0x28, 0x07, 0x6d, 0x9b,
	0x64, 0x45, 0x4e, 0x7e, 0x83, 0x96, 0x56, 0x8d, 0x8e, 0x44, 0x50, 0xe3, 0x16, 0x54, 0xbc, 0x97,
	0x55, 0x06, 0x94, 0xa0, 0x31, 0x03, 0xb8, 0x89, 0x5f, 0x44, 0x19, 0xd4, 0xa3, 0x68, 0xb5, 0x8f,
	0x8e, 0x5d, 0x68, 0x7f, 0x1c, 0x8a, 0xb4, 0xa0, 0xb1, 0xda, 0xac, 0x02, 0xfc, 0x8d, 0xb4, 0x41,
	0x5b, 0xae, 0x3d, 0x9f, 0x62, 0x44, 0x30, 0x74, 0xef, 0x7c, 0xca, 0xc2, 0xa5, 0xb7, 0xd9, 0xd0,
	0x65, 0x80, 0x95, 0x4a, 0xf9, 0xeb, 0xb1, 0x07, 0x87, 0xb9, 0xa1, 0xeb, 0x04, 0x0e, 0x56, 0xc7,
	0x6b, 0xf8, 0x2e, 0x6f, 0x91, 0x57, 0x93, 0x94, 0x31, 0x8f, 0x49, 0x93, 0xad, 0xc7, 0x02, 0x1f,
	0x23, 0xf2, 0x13, 0x88, 0x4f, 0xd9, 0x3d, 0x65, 0x21, 0x7d, 0xac, 0xb4, 0xb0, 0x2a, 0x58, 0x21,
	0x7d, 0xf8, 0x51, 0x37, 0x0f, 0x57, 0x2e, 0x56, 0x17, 0x43, 0xe8, 0xd5, 0xfe, 0x8b, 0x7d, 0x12,
	0x2f, 0x34, 0x91, 0x71, 0x8b, 0x76, 0x4d, 0x21, 0xfd, 0x79, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x94,
	0x81, 0x24, 0xc7, 0x42, 0x02, 0x00, 0x00,
}
