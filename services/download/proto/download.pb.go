// Code generated by protoc-gen-go. DO NOT EDIT.
// source: download.proto

package proto

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

type DownloadEndpointReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownloadEndpointReq) Reset()         { *m = DownloadEndpointReq{} }
func (m *DownloadEndpointReq) String() string { return proto.CompactTextString(m) }
func (*DownloadEndpointReq) ProtoMessage()    {}
func (*DownloadEndpointReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7ce52b48c9eea83, []int{0}
}

func (m *DownloadEndpointReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadEndpointReq.Unmarshal(m, b)
}
func (m *DownloadEndpointReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadEndpointReq.Marshal(b, m, deterministic)
}
func (m *DownloadEndpointReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadEndpointReq.Merge(m, src)
}
func (m *DownloadEndpointReq) XXX_Size() int {
	return xxx_messageInfo_DownloadEndpointReq.Size(m)
}
func (m *DownloadEndpointReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadEndpointReq.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadEndpointReq proto.InternalMessageInfo

type DownloadEndpointResp struct {
	Code                 int64                      `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string                     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data                 *DownloadEndpointResp_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *DownloadEndpointResp) Reset()         { *m = DownloadEndpointResp{} }
func (m *DownloadEndpointResp) String() string { return proto.CompactTextString(m) }
func (*DownloadEndpointResp) ProtoMessage()    {}
func (*DownloadEndpointResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7ce52b48c9eea83, []int{1}
}

func (m *DownloadEndpointResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadEndpointResp.Unmarshal(m, b)
}
func (m *DownloadEndpointResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadEndpointResp.Marshal(b, m, deterministic)
}
func (m *DownloadEndpointResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadEndpointResp.Merge(m, src)
}
func (m *DownloadEndpointResp) XXX_Size() int {
	return xxx_messageInfo_DownloadEndpointResp.Size(m)
}
func (m *DownloadEndpointResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadEndpointResp.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadEndpointResp proto.InternalMessageInfo

func (m *DownloadEndpointResp) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *DownloadEndpointResp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *DownloadEndpointResp) GetData() *DownloadEndpointResp_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type DownloadEndpointResp_Data struct {
	Endpoint             string   `protobuf:"bytes,3,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownloadEndpointResp_Data) Reset()         { *m = DownloadEndpointResp_Data{} }
func (m *DownloadEndpointResp_Data) String() string { return proto.CompactTextString(m) }
func (*DownloadEndpointResp_Data) ProtoMessage()    {}
func (*DownloadEndpointResp_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7ce52b48c9eea83, []int{1, 0}
}

func (m *DownloadEndpointResp_Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadEndpointResp_Data.Unmarshal(m, b)
}
func (m *DownloadEndpointResp_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadEndpointResp_Data.Marshal(b, m, deterministic)
}
func (m *DownloadEndpointResp_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadEndpointResp_Data.Merge(m, src)
}
func (m *DownloadEndpointResp_Data) XXX_Size() int {
	return xxx_messageInfo_DownloadEndpointResp_Data.Size(m)
}
func (m *DownloadEndpointResp_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadEndpointResp_Data.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadEndpointResp_Data proto.InternalMessageInfo

func (m *DownloadEndpointResp_Data) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func init() {
	proto.RegisterType((*DownloadEndpointReq)(nil), "proto.DownloadEndpointReq")
	proto.RegisterType((*DownloadEndpointResp)(nil), "proto.DownloadEndpointResp")
	proto.RegisterType((*DownloadEndpointResp_Data)(nil), "proto.DownloadEndpointResp.Data")
}

func init() { proto.RegisterFile("download.proto", fileDescriptor_f7ce52b48c9eea83) }

var fileDescriptor_f7ce52b48c9eea83 = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0xc9, 0x2f, 0xcf,
	0xcb, 0xc9, 0x4f, 0x4c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xa2,
	0x5c, 0xc2, 0x2e, 0x50, 0x09, 0xd7, 0xbc, 0x94, 0x82, 0xfc, 0xcc, 0xbc, 0x92, 0xa0, 0xd4, 0x42,
	0xa5, 0x79, 0x8c, 0x5c, 0x22, 0x98, 0xe2, 0xc5, 0x05, 0x42, 0x42, 0x5c, 0x2c, 0xc9, 0xf9, 0x29,
	0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x60, 0xb6, 0x90, 0x04, 0x17, 0x7b, 0x6e, 0x6a,
	0x71, 0x71, 0x62, 0x7a, 0xaa, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x8c, 0x2b, 0x64, 0xc2,
	0xc5, 0x92, 0x92, 0x58, 0x92, 0x28, 0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0x6d, 0xa4, 0x00, 0xb1, 0x5a,
	0x0f, 0x9b, 0xc1, 0x7a, 0x2e, 0x89, 0x25, 0x89, 0x41, 0x60, 0xd5, 0x52, 0x4a, 0x5c, 0x2c, 0x20,
	0x9e, 0x90, 0x14, 0x17, 0x47, 0x2a, 0x54, 0x09, 0xd8, 0x04, 0xce, 0x20, 0x38, 0xdf, 0x28, 0x8e,
	0x8b, 0x1f, 0x66, 0x4c, 0x70, 0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x90, 0x37, 0x97, 0x00, 0x92,
	0xc9, 0x01, 0x20, 0x65, 0x42, 0x52, 0x38, 0xad, 0x2c, 0x94, 0x92, 0xc6, 0xe3, 0x9c, 0x24, 0x36,
	0xb0, 0x9c, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x2e, 0x37, 0x65, 0xe2, 0x37, 0x01, 0x00, 0x00,
}