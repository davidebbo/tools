// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protoc-gen-crds/testdata/basic.proto

package testdata

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "istio.io/tools/kubernetes/resource"
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

// BasicMessage is a message that I use for testing.
type BasicMessage struct {
	FieldInt32           int32           `protobuf:"varint,1,opt,name=field_int32,json=fieldInt32,proto3" json:"field_int32,omitempty"`
	FieldString          string          `protobuf:"bytes,2,opt,name=field_string,json=fieldString,proto3" json:"field_string,omitempty"`
	FieldRepeatedString  []string        `protobuf:"bytes,3,rep,name=field_repeated_string,json=fieldRepeatedString,proto3" json:"field_repeated_string,omitempty"`
	FieldMessage         *InnerMessage   `protobuf:"bytes,4,opt,name=field_message,json=fieldMessage,proto3" json:"field_message,omitempty"`
	FieldRepeatedMessage []*InnerMessage `protobuf:"bytes,5,rep,name=field_repeated_message,json=fieldRepeatedMessage,proto3" json:"field_repeated_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *BasicMessage) Reset()         { *m = BasicMessage{} }
func (m *BasicMessage) String() string { return proto.CompactTextString(m) }
func (*BasicMessage) ProtoMessage()    {}
func (*BasicMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d69e4565f225795, []int{0}
}

func (m *BasicMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BasicMessage.Unmarshal(m, b)
}
func (m *BasicMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BasicMessage.Marshal(b, m, deterministic)
}
func (m *BasicMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BasicMessage.Merge(m, src)
}
func (m *BasicMessage) XXX_Size() int {
	return xxx_messageInfo_BasicMessage.Size(m)
}
func (m *BasicMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_BasicMessage.DiscardUnknown(m)
}

var xxx_messageInfo_BasicMessage proto.InternalMessageInfo

func (m *BasicMessage) GetFieldInt32() int32 {
	if m != nil {
		return m.FieldInt32
	}
	return 0
}

func (m *BasicMessage) GetFieldString() string {
	if m != nil {
		return m.FieldString
	}
	return ""
}

func (m *BasicMessage) GetFieldRepeatedString() []string {
	if m != nil {
		return m.FieldRepeatedString
	}
	return nil
}

func (m *BasicMessage) GetFieldMessage() *InnerMessage {
	if m != nil {
		return m.FieldMessage
	}
	return nil
}

func (m *BasicMessage) GetFieldRepeatedMessage() []*InnerMessage {
	if m != nil {
		return m.FieldRepeatedMessage
	}
	return nil
}

type InnerMessage struct {
	FieldBool            bool     `protobuf:"varint,1,opt,name=field_bool,json=fieldBool,proto3" json:"field_bool,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InnerMessage) Reset()         { *m = InnerMessage{} }
func (m *InnerMessage) String() string { return proto.CompactTextString(m) }
func (*InnerMessage) ProtoMessage()    {}
func (*InnerMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d69e4565f225795, []int{1}
}

func (m *InnerMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InnerMessage.Unmarshal(m, b)
}
func (m *InnerMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InnerMessage.Marshal(b, m, deterministic)
}
func (m *InnerMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InnerMessage.Merge(m, src)
}
func (m *InnerMessage) XXX_Size() int {
	return xxx_messageInfo_InnerMessage.Size(m)
}
func (m *InnerMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_InnerMessage.DiscardUnknown(m)
}

var xxx_messageInfo_InnerMessage proto.InternalMessageInfo

func (m *InnerMessage) GetFieldBool() bool {
	if m != nil {
		return m.FieldBool
	}
	return false
}

func init() {
	proto.RegisterType((*BasicMessage)(nil), "testpkg.BasicMessage")
	proto.RegisterType((*InnerMessage)(nil), "testpkg.InnerMessage")
}

func init() {
	proto.RegisterFile("protoc-gen-crds/testdata/basic.proto", fileDescriptor_7d69e4565f225795)
}

var fileDescriptor_7d69e4565f225795 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x3f, 0x4f, 0xc2, 0x40,
	0x18, 0xc6, 0x2d, 0x88, 0xc2, 0x81, 0x31, 0xa9, 0x62, 0x2e, 0x26, 0xc6, 0x42, 0x4c, 0xac, 0x03,
	0xad, 0xc0, 0xc6, 0xc8, 0x46, 0x8c, 0xcb, 0xb9, 0xb9, 0x90, 0x6b, 0x79, 0xa9, 0x17, 0xea, 0xbd,
	0xcd, 0xdd, 0xe1, 0xec, 0xe8, 0xec, 0xe8, 0xe8, 0x77, 0xe8, 0x07, 0xf0, 0x9b, 0x19, 0xee, 0x8a,
	0x51, 0x13, 0xd6, 0xdf, 0xf3, 0xa7, 0x4f, 0xdf, 0x23, 0x57, 0x85, 0x42, 0x83, 0xe9, 0x20, 0x03,
	0x39, 0x48, 0xd5, 0x42, 0xc7, 0x06, 0xb4, 0x59, 0x70, 0xc3, 0xe3, 0x84, 0x6b, 0x91, 0x46, 0x56,
	0xf6, 0x0f, 0x37, 0xb4, 0x58, 0x65, 0xe7, 0xbd, 0xd5, 0x3a, 0x01, 0x25, 0xc1, 0x80, 0x8e, 0x15,
	0x68, 0x5c, 0xab, 0x14, 0x62, 0x2c, 0x8c, 0x40, 0xa9, 0x9d, 0xb7, 0xff, 0x55, 0x23, 0x9d, 0xe9,
	0x26, 0x7b, 0x0f, 0x5a, 0xf3, 0x0c, 0xfc, 0x4b, 0xd2, 0x5e, 0x0a, 0xc8, 0x17, 0x73, 0x21, 0xcd,
	0x78, 0x44, 0xbd, 0xc0, 0x0b, 0x1b, 0x8c, 0x58, 0x34, 0xdb, 0x10, 0xbf, 0x47, 0x3a, 0xce, 0xa0,
	0x8d, 0x12, 0x32, 0xa3, 0xb5, 0xc0, 0x0b, 0x5b, 0xcc, 0x85, 0x1e, 0x2c, 0xf2, 0x47, 0xa4, 0xeb,
	0x2c, 0x0a, 0x0a, 0xe0, 0x06, 0x7e, 0xbc, 0xf5, 0xa0, 0x1e, 0xb6, 0xd8, 0x89, 0x15, 0x59, 0xa5,
	0x55, 0x99, 0x09, 0x39, 0x72, 0x99, 0x67, 0x37, 0x84, 0xee, 0x07, 0x5e, 0xd8, 0x1e, 0x75, 0xa3,
	0xea, 0x67, 0xa2, 0x99, 0x94, 0xa0, 0xaa, 0x95, 0xcc, 0x4d, 0xd8, 0x6e, 0xbe, 0x23, 0x67, 0xff,
	0xbe, 0xb7, 0x2d, 0x69, 0x04, 0xf5, 0xdd, 0x25, 0xa7, 0x7f, 0x76, 0x54, 0x74, 0x72, 0xfb, 0x5a,
	0x52, 0xef, 0xad, 0xa4, 0x7b, 0xef, 0x25, 0x3d, 0x4e, 0x51, 0x2e, 0x45, 0x16, 0x09, 0x6d, 0x04,
	0x46, 0x02, 0x3f, 0x4a, 0xda, 0x7c, 0x19, 0xf2, 0xbc, 0x78, 0xe2, 0xc3, 0xcf, 0x92, 0x36, 0xec,
	0xd5, 0xfb, 0x03, 0xd2, 0xf9, 0xdd, 0xeb, 0x5f, 0x10, 0x77, 0xaf, 0x79, 0x82, 0x98, 0xdb, 0x0b,
	0x36, 0x59, 0xcb, 0x92, 0x29, 0x62, 0x3e, 0xbd, 0x79, 0xbc, 0xde, 0xf6, 0xc5, 0x06, 0x31, 0xd7,
	0xf1, 0xae, 0x57, 0x4d, 0x0e, 0xac, 0x32, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x37, 0xe8, 0x14,
	0xc2, 0xf8, 0x01, 0x00, 0x00,
}
