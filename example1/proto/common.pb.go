// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example1/proto/common.proto

package model

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RespType int32

const (
	RespType_NONE    RespType = 0
	RespType_ASCEND  RespType = 1
	RespType_DESCEND RespType = 2
)

var RespType_name = map[int32]string{
	0: "NONE",
	1: "ASCEND",
	2: "DESCEND",
}
var RespType_value = map[string]int32{
	"NONE":    0,
	"ASCEND":  1,
	"DESCEND": 2,
}

func (x RespType) String() string {
	return proto.EnumName(RespType_name, int32(x))
}
func (RespType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_common_8ae4e6f4c85c4fd8, []int{0}
}

type SayParam struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SayParam) Reset()         { *m = SayParam{} }
func (m *SayParam) String() string { return proto.CompactTextString(m) }
func (*SayParam) ProtoMessage()    {}
func (*SayParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_8ae4e6f4c85c4fd8, []int{0}
}
func (m *SayParam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SayParam.Unmarshal(m, b)
}
func (m *SayParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SayParam.Marshal(b, m, deterministic)
}
func (dst *SayParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SayParam.Merge(dst, src)
}
func (m *SayParam) XXX_Size() int {
	return xxx_messageInfo_SayParam.Size(m)
}
func (m *SayParam) XXX_DiscardUnknown() {
	xxx_messageInfo_SayParam.DiscardUnknown(m)
}

var xxx_messageInfo_SayParam proto.InternalMessageInfo

func (m *SayParam) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type Pair struct {
	Key                  int32    `protobuf:"varint,1,opt,name=key" json:"key,omitempty"`
	Values               string   `protobuf:"bytes,2,opt,name=values" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pair) Reset()         { *m = Pair{} }
func (m *Pair) String() string { return proto.CompactTextString(m) }
func (*Pair) ProtoMessage()    {}
func (*Pair) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_8ae4e6f4c85c4fd8, []int{1}
}
func (m *Pair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pair.Unmarshal(m, b)
}
func (m *Pair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pair.Marshal(b, m, deterministic)
}
func (dst *Pair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pair.Merge(dst, src)
}
func (m *Pair) XXX_Size() int {
	return xxx_messageInfo_Pair.Size(m)
}
func (m *Pair) XXX_DiscardUnknown() {
	xxx_messageInfo_Pair.DiscardUnknown(m)
}

var xxx_messageInfo_Pair proto.InternalMessageInfo

func (m *Pair) GetKey() int32 {
	if m != nil {
		return m.Key
	}
	return 0
}

func (m *Pair) GetValues() string {
	if m != nil {
		return m.Values
	}
	return ""
}

type SayResponse struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
	// 数组
	Values []string `protobuf:"bytes,2,rep,name=values" json:"values,omitempty"`
	// map
	Header               map[string]*Pair `protobuf:"bytes,3,rep,name=header" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Type                 RespType         `protobuf:"varint,4,opt,name=type,enum=model.RespType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SayResponse) Reset()         { *m = SayResponse{} }
func (m *SayResponse) String() string { return proto.CompactTextString(m) }
func (*SayResponse) ProtoMessage()    {}
func (*SayResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_8ae4e6f4c85c4fd8, []int{2}
}
func (m *SayResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SayResponse.Unmarshal(m, b)
}
func (m *SayResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SayResponse.Marshal(b, m, deterministic)
}
func (dst *SayResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SayResponse.Merge(dst, src)
}
func (m *SayResponse) XXX_Size() int {
	return xxx_messageInfo_SayResponse.Size(m)
}
func (m *SayResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SayResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SayResponse proto.InternalMessageInfo

func (m *SayResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *SayResponse) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *SayResponse) GetHeader() map[string]*Pair {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SayResponse) GetType() RespType {
	if m != nil {
		return m.Type
	}
	return RespType_NONE
}

func init() {
	proto.RegisterType((*SayParam)(nil), "model.SayParam")
	proto.RegisterType((*Pair)(nil), "model.Pair")
	proto.RegisterType((*SayResponse)(nil), "model.SayResponse")
	proto.RegisterMapType((map[string]*Pair)(nil), "model.SayResponse.HeaderEntry")
	proto.RegisterEnum("model.RespType", RespType_name, RespType_value)
}

func init() { proto.RegisterFile("example1/proto/common.proto", fileDescriptor_common_8ae4e6f4c85c4fd8) }

var fileDescriptor_common_8ae4e6f4c85c4fd8 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x5d, 0x4b, 0xf3, 0x30,
	0x14, 0xc7, 0x97, 0xf5, 0xe5, 0xd9, 0x4e, 0xe0, 0xd9, 0xc8, 0x85, 0x94, 0x29, 0x32, 0xeb, 0x4d,
	0x11, 0xec, 0xb4, 0xa2, 0x88, 0x77, 0xe2, 0x2a, 0xbb, 0xaa, 0x23, 0xf5, 0x0b, 0x44, 0x77, 0x50,
	0xb1, 0x69, 0x4a, 0x5b, 0xc5, 0x7c, 0x57, 0x3f, 0x8c, 0x34, 0xad, 0xac, 0xb2, 0xbb, 0x73, 0xf8,
	0xbf, 0x70, 0x7e, 0x09, 0xec, 0xe3, 0x97, 0x90, 0x45, 0x86, 0xe7, 0x8b, 0xa2, 0x54, 0xb5, 0x5a,
	0x3c, 0x2b, 0x29, 0x55, 0x1e, 0x9a, 0x85, 0x39, 0x52, 0x6d, 0x30, 0xf3, 0x0f, 0x60, 0x94, 0x0a,
	0xbd, 0x16, 0xa5, 0x90, 0x6c, 0x0a, 0x96, 0xac, 0x5e, 0x3c, 0x32, 0x27, 0xc1, 0x98, 0x37, 0xa3,
	0x7f, 0x06, 0xf6, 0x5a, 0xbc, 0x95, 0x8d, 0xf2, 0x8e, 0xda, 0x28, 0x0e, 0x6f, 0x46, 0xb6, 0x07,
	0xee, 0xa7, 0xc8, 0x3e, 0xb0, 0xf2, 0x86, 0xc6, 0xde, 0x6d, 0xfe, 0x37, 0x01, 0x9a, 0x0a, 0xcd,
	0xb1, 0x2a, 0x54, 0x5e, 0xe1, 0x6e, 0xe7, 0x9f, 0xa4, 0xb5, 0x4d, 0xb2, 0x2b, 0x70, 0x5f, 0x51,
	0x6c, 0xb0, 0xf4, 0xac, 0xb9, 0x15, 0xd0, 0xe8, 0x30, 0x34, 0x17, 0x86, 0xbd, 0xb6, 0x70, 0x65,
	0x0c, 0x71, 0x5e, 0x97, 0x9a, 0x77, 0x6e, 0x76, 0x0c, 0x76, 0xad, 0x0b, 0xf4, 0xec, 0x39, 0x09,
	0xfe, 0x47, 0x93, 0x2e, 0xd5, 0x44, 0x1e, 0x75, 0x81, 0xdc, 0x88, 0xb3, 0x7b, 0xa0, 0xbd, 0x6c,
	0x9f, 0x67, 0xdc, 0xf2, 0x1c, 0x81, 0x63, 0xee, 0x30, 0x38, 0x34, 0xa2, 0x5d, 0x4d, 0x43, 0xcf,
	0x5b, 0xe5, 0x66, 0x78, 0x4d, 0x4e, 0x4e, 0x61, 0xf4, 0xdb, 0xcc, 0x46, 0x60, 0x27, 0x0f, 0x49,
	0x3c, 0x1d, 0x30, 0x00, 0xf7, 0x36, 0xbd, 0x8b, 0x93, 0xe5, 0x94, 0x30, 0x0a, 0xff, 0x96, 0x71,
	0xbb, 0x0c, 0xa3, 0x4b, 0xb0, 0x52, 0xa1, 0x59, 0x08, 0xce, 0x0a, 0xb3, 0x4c, 0xb1, 0xc9, 0x96,
	0xc9, 0x3c, 0xf9, 0x8c, 0xed, 0x42, 0xfa, 0x83, 0x27, 0xd7, 0x7c, 0xd1, 0xc5, 0x4f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x6a, 0x28, 0x0e, 0x6e, 0xc1, 0x01, 0x00, 0x00,
}
