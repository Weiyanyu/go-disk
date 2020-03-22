// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package userinterface

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

type RegisterReq struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterReq) Reset()         { *m = RegisterReq{} }
func (m *RegisterReq) String() string { return proto.CompactTextString(m) }
func (*RegisterReq) ProtoMessage()    {}
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *RegisterReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterReq.Unmarshal(m, b)
}
func (m *RegisterReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterReq.Marshal(b, m, deterministic)
}
func (m *RegisterReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterReq.Merge(m, src)
}
func (m *RegisterReq) XXX_Size() int {
	return xxx_messageInfo_RegisterReq.Size(m)
}
func (m *RegisterReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterReq.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterReq proto.InternalMessageInfo

func (m *RegisterReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *RegisterReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RegisterResp struct {
	Code                 int64    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResp) Reset()         { *m = RegisterResp{} }
func (m *RegisterResp) String() string { return proto.CompactTextString(m) }
func (*RegisterResp) ProtoMessage()    {}
func (*RegisterResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *RegisterResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResp.Unmarshal(m, b)
}
func (m *RegisterResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResp.Marshal(b, m, deterministic)
}
func (m *RegisterResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResp.Merge(m, src)
}
func (m *RegisterResp) XXX_Size() int {
	return xxx_messageInfo_RegisterResp.Size(m)
}
func (m *RegisterResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResp.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResp proto.InternalMessageInfo

func (m *RegisterResp) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RegisterResp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type QueryUserInfoReq struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryUserInfoReq) Reset()         { *m = QueryUserInfoReq{} }
func (m *QueryUserInfoReq) String() string { return proto.CompactTextString(m) }
func (*QueryUserInfoReq) ProtoMessage()    {}
func (*QueryUserInfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *QueryUserInfoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryUserInfoReq.Unmarshal(m, b)
}
func (m *QueryUserInfoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryUserInfoReq.Marshal(b, m, deterministic)
}
func (m *QueryUserInfoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUserInfoReq.Merge(m, src)
}
func (m *QueryUserInfoReq) XXX_Size() int {
	return xxx_messageInfo_QueryUserInfoReq.Size(m)
}
func (m *QueryUserInfoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUserInfoReq.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUserInfoReq proto.InternalMessageInfo

func (m *QueryUserInfoReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type QueryUserInfoResp struct {
	Code                 int64                   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string                  `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data                 *QueryUserInfoResp_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *QueryUserInfoResp) Reset()         { *m = QueryUserInfoResp{} }
func (m *QueryUserInfoResp) String() string { return proto.CompactTextString(m) }
func (*QueryUserInfoResp) ProtoMessage()    {}
func (*QueryUserInfoResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *QueryUserInfoResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryUserInfoResp.Unmarshal(m, b)
}
func (m *QueryUserInfoResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryUserInfoResp.Marshal(b, m, deterministic)
}
func (m *QueryUserInfoResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUserInfoResp.Merge(m, src)
}
func (m *QueryUserInfoResp) XXX_Size() int {
	return xxx_messageInfo_QueryUserInfoResp.Size(m)
}
func (m *QueryUserInfoResp) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUserInfoResp.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUserInfoResp proto.InternalMessageInfo

func (m *QueryUserInfoResp) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *QueryUserInfoResp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *QueryUserInfoResp) GetData() *QueryUserInfoResp_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type QueryUserInfoResp_Data struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string   `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Profile              string   `protobuf:"bytes,4,opt,name=profile,proto3" json:"profile,omitempty"`
	LastActive           string   `protobuf:"bytes,5,opt,name=lastActive,proto3" json:"lastActive,omitempty"`
	SignupAt             string   `protobuf:"bytes,6,opt,name=signupAt,proto3" json:"signupAt,omitempty"`
	EmailValidated       bool     `protobuf:"varint,7,opt,name=emailValidated,proto3" json:"emailValidated,omitempty"`
	PhoneValidated       bool     `protobuf:"varint,8,opt,name=phoneValidated,proto3" json:"phoneValidated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryUserInfoResp_Data) Reset()         { *m = QueryUserInfoResp_Data{} }
func (m *QueryUserInfoResp_Data) String() string { return proto.CompactTextString(m) }
func (*QueryUserInfoResp_Data) ProtoMessage()    {}
func (*QueryUserInfoResp_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3, 0}
}

func (m *QueryUserInfoResp_Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryUserInfoResp_Data.Unmarshal(m, b)
}
func (m *QueryUserInfoResp_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryUserInfoResp_Data.Marshal(b, m, deterministic)
}
func (m *QueryUserInfoResp_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUserInfoResp_Data.Merge(m, src)
}
func (m *QueryUserInfoResp_Data) XXX_Size() int {
	return xxx_messageInfo_QueryUserInfoResp_Data.Size(m)
}
func (m *QueryUserInfoResp_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUserInfoResp_Data.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUserInfoResp_Data proto.InternalMessageInfo

func (m *QueryUserInfoResp_Data) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *QueryUserInfoResp_Data) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *QueryUserInfoResp_Data) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *QueryUserInfoResp_Data) GetProfile() string {
	if m != nil {
		return m.Profile
	}
	return ""
}

func (m *QueryUserInfoResp_Data) GetLastActive() string {
	if m != nil {
		return m.LastActive
	}
	return ""
}

func (m *QueryUserInfoResp_Data) GetSignupAt() string {
	if m != nil {
		return m.SignupAt
	}
	return ""
}

func (m *QueryUserInfoResp_Data) GetEmailValidated() bool {
	if m != nil {
		return m.EmailValidated
	}
	return false
}

func (m *QueryUserInfoResp_Data) GetPhoneValidated() bool {
	if m != nil {
		return m.PhoneValidated
	}
	return false
}

func init() {
	proto.RegisterType((*RegisterReq)(nil), "userinterface.RegisterReq")
	proto.RegisterType((*RegisterResp)(nil), "userinterface.RegisterResp")
	proto.RegisterType((*QueryUserInfoReq)(nil), "userinterface.QueryUserInfoReq")
	proto.RegisterType((*QueryUserInfoResp)(nil), "userinterface.QueryUserInfoResp")
	proto.RegisterType((*QueryUserInfoResp_Data)(nil), "userinterface.QueryUserInfoResp.Data")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4d, 0x4b, 0xc3, 0x30,
	0x18, 0xc7, 0xe9, 0xd6, 0xbd, 0x3d, 0xdb, 0x44, 0x83, 0x87, 0x50, 0x41, 0x4b, 0x41, 0xd9, 0xa9,
	0x87, 0x79, 0x12, 0xbc, 0x0c, 0x14, 0xf1, 0x68, 0x44, 0xef, 0xb1, 0x7d, 0x36, 0x03, 0x5d, 0xdb,
	0x25, 0xd9, 0xc4, 0xaf, 0xe4, 0xd7, 0xd3, 0x0f, 0x20, 0x49, 0x56, 0xf7, 0x22, 0x6e, 0x78, 0xcb,
	0xef, 0x79, 0xfd, 0xf3, 0x7f, 0x02, 0x30, 0x57, 0x28, 0xe3, 0x52, 0x16, 0xba, 0x20, 0x7d, 0xf3,
	0x16, 0xb9, 0x46, 0x39, 0xe6, 0x09, 0x46, 0xb7, 0xd0, 0x65, 0x38, 0x11, 0x4a, 0xa3, 0x64, 0x38,
	0x23, 0x01, 0xb4, 0x4d, 0x3e, 0xe7, 0x53, 0xa4, 0x5e, 0xe8, 0x0d, 0x3a, 0xec, 0x87, 0x4d, 0xae,
	0xe4, 0x4a, 0xbd, 0x15, 0x32, 0xa5, 0x35, 0x97, 0xab, 0x38, 0xba, 0x86, 0xde, 0x6a, 0x8c, 0x2a,
	0x09, 0x01, 0x3f, 0x29, 0x52, 0x37, 0xa3, 0xce, 0xec, 0x9b, 0x50, 0x68, 0x4d, 0x51, 0x29, 0x3e,
	0xc1, 0x65, 0x7b, 0x85, 0x51, 0x0c, 0x87, 0x0f, 0x73, 0x94, 0xef, 0x4f, 0x0a, 0xe5, 0x7d, 0x3e,
	0x2e, 0xf6, 0x28, 0x89, 0x3e, 0x6b, 0x70, 0xb4, 0xd5, 0xf0, 0xdf, 0x9d, 0xe4, 0x0a, 0xfc, 0x94,
	0x6b, 0x4e, 0xeb, 0xa1, 0x37, 0xe8, 0x0e, 0xcf, 0xe3, 0x0d, 0x5b, 0xe2, 0x5f, 0xd3, 0xe3, 0x1b,
	0xae, 0x39, 0xb3, 0x2d, 0xc1, 0x97, 0x07, 0xbe, 0xc1, 0x9d, 0x6e, 0x1d, 0x43, 0x03, 0xa7, 0x5c,
	0x64, 0xcb, 0xbd, 0x0e, 0x4c, 0xb4, 0x7c, 0x2d, 0x72, 0xb4, 0x6b, 0x3b, 0xcc, 0x81, 0x51, 0x59,
	0xca, 0x62, 0x2c, 0x32, 0xa4, 0xbe, 0x53, 0xb9, 0x44, 0x72, 0x0a, 0x90, 0x71, 0xa5, 0x47, 0x89,
	0x16, 0x0b, 0xa4, 0x0d, 0x9b, 0x5c, 0x8b, 0x18, 0x05, 0x4a, 0x4c, 0xf2, 0x79, 0x39, 0xd2, 0xb4,
	0xe9, 0x14, 0x54, 0x4c, 0x2e, 0xe0, 0xc0, 0x2e, 0x7d, 0xe6, 0x99, 0x48, 0xb9, 0xc6, 0x94, 0xb6,
	0x42, 0x6f, 0xd0, 0x66, 0x5b, 0x51, 0x53, 0x67, 0x65, 0xac, 0xea, 0xda, 0xae, 0x6e, 0x33, 0x3a,
	0xfc, 0xf0, 0xa0, 0x6b, 0x2c, 0x79, 0x44, 0xb9, 0x10, 0x09, 0x92, 0x3b, 0xe8, 0x19, 0xac, 0xee,
	0x4e, 0x82, 0x2d, 0x0f, 0xd7, 0xfe, 0x55, 0x70, 0xf2, 0x67, 0x4e, 0x95, 0x84, 0x41, 0x7f, 0xc3,
	0x6f, 0x72, 0xb6, 0xfb, 0x1a, 0xb3, 0x20, 0xdc, 0x77, 0xae, 0x97, 0xa6, 0xfd, 0xed, 0x97, 0xdf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xec, 0xa6, 0xe0, 0xbd, 0xfb, 0x02, 0x00, 0x00,
}
