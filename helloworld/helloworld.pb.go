// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloworld/helloworld.proto

package helloworld

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// The request message that contains id
type CreateIDRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateIDRequest) Reset()         { *m = CreateIDRequest{} }
func (m *CreateIDRequest) String() string { return proto.CompactTextString(m) }
func (*CreateIDRequest) ProtoMessage()    {}
func (*CreateIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_73149fedf49f4319, []int{0}
}

func (m *CreateIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateIDRequest.Unmarshal(m, b)
}
func (m *CreateIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateIDRequest.Marshal(b, m, deterministic)
}
func (m *CreateIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateIDRequest.Merge(m, src)
}
func (m *CreateIDRequest) XXX_Size() int {
	return xxx_messageInfo_CreateIDRequest.Size(m)
}
func (m *CreateIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateIDRequest proto.InternalMessageInfo

func (m *CreateIDRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CreateIDReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateIDReply) Reset()         { *m = CreateIDReply{} }
func (m *CreateIDReply) String() string { return proto.CompactTextString(m) }
func (*CreateIDReply) ProtoMessage()    {}
func (*CreateIDReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_73149fedf49f4319, []int{1}
}

func (m *CreateIDReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateIDReply.Unmarshal(m, b)
}
func (m *CreateIDReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateIDReply.Marshal(b, m, deterministic)
}
func (m *CreateIDReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateIDReply.Merge(m, src)
}
func (m *CreateIDReply) XXX_Size() int {
	return xxx_messageInfo_CreateIDReply.Size(m)
}
func (m *CreateIDReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateIDReply.DiscardUnknown(m)
}

var xxx_messageInfo_CreateIDReply proto.InternalMessageInfo

func (m *CreateIDReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// The request message containing the user's name.
type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_73149fedf49f4319, []int{2}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_73149fedf49f4319, []int{3}
}

func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (m *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(m, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateIDRequest)(nil), "helloworld.CreateIDRequest")
	proto.RegisterType((*CreateIDReply)(nil), "helloworld.CreateIDReply")
	proto.RegisterType((*HelloRequest)(nil), "helloworld.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "helloworld.HelloReply")
}

func init() {
	proto.RegisterFile("helloworld/helloworld.proto", fileDescriptor_73149fedf49f4319)
}

var fileDescriptor_73149fedf49f4319 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0xed, 0x22, 0xee, 0x3a, 0xa8, 0x0b, 0x39, 0x48, 0xdd, 0x5e, 0x34, 0x07, 0xd1, 0x4b,
	0x0a, 0x8a, 0x57, 0x0f, 0xeb, 0x82, 0x7a, 0x5b, 0xd6, 0x83, 0xe0, 0x2d, 0xda, 0x21, 0x16, 0xa6,
	0x3b, 0x31, 0xad, 0x68, 0x7f, 0x85, 0x7f, 0x59, 0x1a, 0x1a, 0x1a, 0xa4, 0x78, 0x7b, 0xed, 0x7c,
	0xef, 0x0d, 0x2f, 0x03, 0xd9, 0x3b, 0x12, 0xf1, 0x17, 0x3b, 0x2a, 0xf2, 0x41, 0x2a, 0xeb, 0xb8,
	0x61, 0x01, 0xc3, 0x1f, 0x79, 0x06, 0xf3, 0x3b, 0x87, 0xba, 0xc1, 0xc7, 0xd5, 0x06, 0x3f, 0x3e,
	0xb1, 0x6e, 0xc4, 0x11, 0x4c, 0xca, 0x22, 0x4d, 0x4e, 0x93, 0x8b, 0xfd, 0xcd, 0xa4, 0x2c, 0xe4,
	0x25, 0x1c, 0x0e, 0x88, 0xa5, 0x56, 0xa4, 0x30, 0xad, 0xb0, 0xae, 0xb5, 0xc1, 0x9e, 0x0a, 0x9f,
	0x52, 0xc2, 0xc1, 0x43, 0x97, 0x1d, 0xa2, 0x04, 0xec, 0x6e, 0x75, 0x15, 0x30, 0xaf, 0xe5, 0x39,
	0x40, 0xcf, 0xfc, 0x9b, 0x75, 0xf5, 0x93, 0xc0, 0xf4, 0xde, 0x21, 0x36, 0xe8, 0xc4, 0x2d, 0xcc,
	0x9e, 0x74, 0xeb, 0x6d, 0x22, 0x55, 0x51, 0xa1, 0x78, 0xdb, 0xe2, 0x78, 0x64, 0x62, 0xa9, 0x95,
	0x3b, 0x62, 0x05, 0xb3, 0x50, 0x41, 0x64, 0x31, 0xf5, 0xa7, 0xfb, 0xe2, 0x64, 0x7c, 0xe8, 0x53,
	0x96, 0x06, 0xb2, 0x92, 0x95, 0x71, 0xf6, 0x4d, 0xe1, 0xb7, 0xae, 0x2c, 0x61, 0x1d, 0xe1, 0xcb,
	0xb9, 0x5f, 0xf9, 0xdc, 0xe9, 0x75, 0xf7, 0xce, 0xeb, 0xe4, 0xe5, 0xc6, 0x30, 0x1b, 0x42, 0x65,
	0x98, 0xf4, 0xd6, 0x28, 0x76, 0x26, 0xef, 0xec, 0x79, 0xb0, 0xe7, 0xa3, 0x67, 0x7a, 0xdd, 0xf3,
	0x77, 0xba, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x3f, 0xe4, 0x46, 0x1c, 0xc6, 0x01, 0x00, 0x00,
}
