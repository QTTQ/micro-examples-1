// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/auth/auth.proto

package io_github_entere_srv_auth

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

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Detail               string   `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{0}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

type Request struct {
	UserId               uint64   `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Request) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *Request) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type Response struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Response) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Response) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*Error)(nil), "io.github.entere.srv.auth.Error")
	proto.RegisterType((*Request)(nil), "io.github.entere.srv.auth.Request")
	proto.RegisterType((*Response)(nil), "io.github.entere.srv.auth.Response")
}

func init() { proto.RegisterFile("proto/auth/auth.proto", fileDescriptor_82b5829f48cfb8e5) }

var fileDescriptor_82b5829f48cfb8e5 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x91, 0x31, 0x4f, 0xfb, 0x30,
	0x10, 0xc5, 0xff, 0xf9, 0xd3, 0x34, 0xe1, 0x18, 0x90, 0x4e, 0x80, 0x42, 0xa7, 0xc8, 0x2c, 0x9d,
	0x8c, 0xd4, 0x4a, 0xec, 0x48, 0x30, 0x30, 0xc0, 0xe0, 0xc2, 0x86, 0x84, 0xd2, 0xe4, 0x44, 0xa3,
	0x96, 0xb8, 0x9c, 0xed, 0x7e, 0x46, 0x3e, 0x16, 0xf2, 0x25, 0x30, 0x41, 0x37, 0x16, 0xeb, 0x7e,
	0xba, 0x77, 0xf7, 0xfc, 0x6c, 0x38, 0xdd, 0xb2, 0xf5, 0xf6, 0xb2, 0x0a, 0x7e, 0x25, 0x87, 0x16,
	0xc6, 0xf3, 0xd6, 0xea, 0xd7, 0xd6, 0xaf, 0xc2, 0x52, 0x53, 0xe7, 0x89, 0x49, 0x3b, 0xde, 0xe9,
	0x28, 0x50, 0x73, 0x48, 0x6f, 0x99, 0x2d, 0x23, 0xc2, 0xa8, 0xb6, 0x0d, 0x15, 0x49, 0x99, 0x4c,
	0x53, 0x23, 0x35, 0x9e, 0xc1, 0xb8, 0x21, 0x5f, 0xb5, 0x9b, 0xe2, 0x7f, 0x99, 0x4c, 0x0f, 0xcd,
	0x40, 0x6a, 0x01, 0x99, 0xa1, 0xf7, 0x40, 0xce, 0x47, 0x49, 0x70, 0xc4, 0x77, 0x8d, 0x0c, 0x8e,
	0xcc, 0x40, 0x38, 0x81, 0x3c, 0x56, 0x0f, 0xd5, 0x1b, 0x0d, 0xc3, 0xdf, 0x8c, 0x27, 0x90, 0x7a,
	0xbb, 0xa6, 0xae, 0x38, 0x90, 0x46, 0x0f, 0x8a, 0x21, 0x37, 0xe4, 0xb6, 0xb6, 0x73, 0x84, 0x05,
	0x64, 0x2e, 0xd4, 0x35, 0x39, 0x27, 0x6b, 0x73, 0xf3, 0x85, 0x78, 0x05, 0x29, 0xc5, 0xfb, 0xca,
	0xd2, 0xa3, 0x59, 0xa9, 0x7f, 0x8d, 0xa6, 0x25, 0x97, 0xe9, 0xe5, 0x3f, 0x7b, 0xce, 0x3e, 0x12,
	0xc8, 0x16, 0xc4, 0xbb, 0xb6, 0x26, 0x7c, 0x86, 0xe3, 0xfb, 0x6a, 0x4d, 0xd7, 0xe2, 0xf3, 0x18,
	0xdb, 0xa8, 0xf6, 0x6c, 0x1f, 0x1e, 0x60, 0x72, 0xb1, 0x57, 0xd3, 0xe7, 0x51, 0xff, 0xf0, 0x05,
	0xf0, 0x86, 0x36, 0x4f, 0x8e, 0xf8, 0x6f, 0x0c, 0x96, 0x63, 0xf9, 0xea, 0xf9, 0x67, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xd1, 0x18, 0x25, 0x70, 0x03, 0x02, 0x00, 0x00,
}
