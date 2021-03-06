// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/oauth/oauth.proto

package io_github_entere_srv_oauth

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

type LoginRequest struct {
	LoginName            string   `protobuf:"bytes,1,opt,name=loginName,proto3" json:"loginName,omitempty"`
	LoginPwd             string   `protobuf:"bytes,2,opt,name=loginPwd,proto3" json:"loginPwd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfe27e33eeec8e22, []int{0}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetLoginName() string {
	if m != nil {
		return m.LoginName
	}
	return ""
}

func (m *LoginRequest) GetLoginPwd() string {
	if m != nil {
		return m.LoginPwd
	}
	return ""
}

type LoginResponse struct {
	Error                *Error     `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Data                 *LoginData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfe27e33eeec8e22, []int{1}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *LoginResponse) GetData() *LoginData {
	if m != nil {
		return m.Data
	}
	return nil
}

type LoginData struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	LoginName            string   `protobuf:"bytes,2,opt,name=loginName,proto3" json:"loginName,omitempty"`
	LoginPwd             string   `protobuf:"bytes,3,opt,name=loginPwd,proto3" json:"loginPwd,omitempty"`
	Token                string   `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginData) Reset()         { *m = LoginData{} }
func (m *LoginData) String() string { return proto.CompactTextString(m) }
func (*LoginData) ProtoMessage()    {}
func (*LoginData) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfe27e33eeec8e22, []int{2}
}

func (m *LoginData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginData.Unmarshal(m, b)
}
func (m *LoginData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginData.Marshal(b, m, deterministic)
}
func (m *LoginData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginData.Merge(m, src)
}
func (m *LoginData) XXX_Size() int {
	return xxx_messageInfo_LoginData.Size(m)
}
func (m *LoginData) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginData.DiscardUnknown(m)
}

var xxx_messageInfo_LoginData proto.InternalMessageInfo

func (m *LoginData) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *LoginData) GetLoginName() string {
	if m != nil {
		return m.LoginName
	}
	return ""
}

func (m *LoginData) GetLoginPwd() string {
	if m != nil {
		return m.LoginPwd
	}
	return ""
}

func (m *LoginData) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfe27e33eeec8e22, []int{3}
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

func (m *Error) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "io.github.entere.srv.oauth.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "io.github.entere.srv.oauth.LoginResponse")
	proto.RegisterType((*LoginData)(nil), "io.github.entere.srv.oauth.LoginData")
	proto.RegisterType((*Error)(nil), "io.github.entere.srv.oauth.Error")
}

func init() { proto.RegisterFile("proto/oauth/oauth.proto", fileDescriptor_cfe27e33eeec8e22) }

var fileDescriptor_cfe27e33eeec8e22 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x35, 0x6d, 0x52, 0xcc, 0x54, 0x51, 0x06, 0xd1, 0x10, 0x3c, 0x68, 0x40, 0xa8, 0x07, 0x57,
	0x88, 0x07, 0xf1, 0x2a, 0x15, 0x14, 0x44, 0x4b, 0xfe, 0x60, 0x6b, 0x96, 0x24, 0x68, 0x33, 0xed,
	0xee, 0xc6, 0xe2, 0xd9, 0x1f, 0x97, 0x4c, 0x96, 0x4a, 0x0f, 0xd6, 0x5e, 0xc2, 0xbc, 0x99, 0x37,
	0x33, 0xef, 0x65, 0x16, 0x4e, 0xe6, 0x9a, 0x2c, 0x5d, 0x93, 0x6c, 0x6c, 0xd9, 0x7d, 0x05, 0x67,
	0x30, 0xae, 0x48, 0x14, 0x95, 0x2d, 0x9b, 0xa9, 0x50, 0xb5, 0x55, 0x5a, 0x09, 0xa3, 0x3f, 0x05,
	0x33, 0x92, 0x47, 0xd8, 0x7b, 0xa6, 0xa2, 0xaa, 0x33, 0xb5, 0x68, 0x94, 0xb1, 0x78, 0x0a, 0xe1,
	0x47, 0x8b, 0x5f, 0xe4, 0x4c, 0x45, 0xde, 0x99, 0x37, 0x0a, 0xb3, 0xdf, 0x04, 0xc6, 0xb0, 0xcb,
	0x60, 0xb2, 0xcc, 0xa3, 0x1e, 0x17, 0x57, 0x38, 0xf9, 0xf6, 0x60, 0xdf, 0x8d, 0x32, 0x73, 0xaa,
	0x8d, 0xc2, 0x5b, 0x08, 0x94, 0xd6, 0xa4, 0x79, 0xce, 0x30, 0x3d, 0x17, 0x7f, 0xeb, 0x10, 0x0f,
	0x2d, 0x31, 0xeb, 0xf8, 0x78, 0x07, 0x7e, 0x2e, 0xad, 0xe4, 0x15, 0xc3, 0xf4, 0x62, 0x53, 0x1f,
	0x6f, 0x1c, 0x4b, 0x2b, 0x33, 0x6e, 0x49, 0x0c, 0x84, 0xab, 0x14, 0x1e, 0xc3, 0xa0, 0x31, 0x4a,
	0x3f, 0x8d, 0x9d, 0x13, 0x87, 0xd6, 0x4d, 0xf6, 0x36, 0x99, 0xec, 0xaf, 0x9b, 0xc4, 0x23, 0x08,
	0x2c, 0xbd, 0xab, 0x3a, 0xf2, 0xb9, 0xd0, 0x81, 0xe4, 0x0a, 0x02, 0xd6, 0x8f, 0x08, 0xfe, 0x1b,
	0xe5, 0xdd, 0x8f, 0x0b, 0x32, 0x8e, 0xf1, 0x10, 0xfa, 0x33, 0x53, 0xb8, 0x35, 0x6d, 0x98, 0x2e,
	0x20, 0x78, 0x6d, 0xc5, 0x63, 0x09, 0x07, 0x2c, 0xf6, 0xfe, 0x6b, 0x22, 0x8d, 0x59, 0x92, 0xce,
	0x71, 0xf4, 0xaf, 0x59, 0x77, 0xa9, 0xf8, 0x72, 0x0b, 0x66, 0x77, 0x88, 0x64, 0x67, 0x3a, 0xe0,
	0x97, 0x70, 0xf3, 0x13, 0x00, 0x00, 0xff, 0xff, 0xe1, 0x00, 0x05, 0x06, 0x24, 0x02, 0x00, 0x00,
}
