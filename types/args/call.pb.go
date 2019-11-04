// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types/args/call.proto

package args

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	abi "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/types/abi"
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

// Call contains information for a contract call
type Call struct {
	// Contract to call method on
	Contract *abi.Contract `protobuf:"bytes,1,opt,name=contract,proto3" json:"contract,omitempty"`
	// Method to call
	Method *abi.Method `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	// Arguments to feed on transaction call
	Args                 []string `protobuf:"bytes,3,rep,name=args,proto3" json:"args,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Call) Reset()         { *m = Call{} }
func (m *Call) String() string { return proto.CompactTextString(m) }
func (*Call) ProtoMessage()    {}
func (*Call) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9df5d94cbe4727e, []int{0}
}

func (m *Call) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Call.Unmarshal(m, b)
}
func (m *Call) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Call.Marshal(b, m, deterministic)
}
func (m *Call) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Call.Merge(m, src)
}
func (m *Call) XXX_Size() int {
	return xxx_messageInfo_Call.Size(m)
}
func (m *Call) XXX_DiscardUnknown() {
	xxx_messageInfo_Call.DiscardUnknown(m)
}

var xxx_messageInfo_Call proto.InternalMessageInfo

func (m *Call) GetContract() *abi.Contract {
	if m != nil {
		return m.Contract
	}
	return nil
}

func (m *Call) GetMethod() *abi.Method {
	if m != nil {
		return m.Method
	}
	return nil
}

func (m *Call) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func init() {
	proto.RegisterType((*Call)(nil), "args.Call")
}

func init() { proto.RegisterFile("types/args/call.proto", fileDescriptor_a9df5d94cbe4727e) }

var fileDescriptor_a9df5d94cbe4727e = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8f, 0x31, 0xab, 0xc2, 0x30,
	0x10, 0xc7, 0xe9, 0x6b, 0x29, 0xef, 0xa5, 0xbc, 0x25, 0x22, 0x14, 0xa7, 0xa2, 0x4b, 0x1d, 0x4c,
	0x40, 0xbf, 0x41, 0xeb, 0xea, 0x52, 0x37, 0xb7, 0x34, 0xc6, 0x1a, 0x4d, 0x93, 0x92, 0xdc, 0xd2,
	0x6f, 0x2f, 0x3d, 0x8b, 0x0e, 0x07, 0x7f, 0xfe, 0xbf, 0x1f, 0xdc, 0x1d, 0x59, 0xc2, 0x38, 0xa8,
	0xc0, 0x85, 0xef, 0x02, 0x97, 0xc2, 0x18, 0x36, 0x78, 0x07, 0x8e, 0x26, 0x53, 0xb1, 0x5a, 0xcc,
	0xb0, 0xd5, 0xd3, 0xbc, 0xd1, 0xfa, 0x41, 0x92, 0x5a, 0x18, 0x43, 0xb7, 0xe4, 0x57, 0x3a, 0x0b,
	0x5e, 0x48, 0xc8, 0xa3, 0x22, 0x2a, 0xb3, 0xfd, 0x3f, 0x9b, 0xac, 0x7a, 0x2e, 0x9b, 0x0f, 0xa6,
	0x1b, 0x92, 0xf6, 0x0a, 0xee, 0xee, 0x9a, 0xff, 0xa0, 0x98, 0xa1, 0x78, 0xc2, 0xaa, 0x99, 0x11,
	0xa5, 0x04, 0x97, 0xe6, 0x71, 0x11, 0x97, 0x7f, 0x0d, 0xe6, 0xea, 0x78, 0xa9, 0x3a, 0x0d, 0x46,
	0xb4, 0x4c, 0xba, 0x9e, 0xd7, 0xce, 0x06, 0x65, 0xcf, 0x63, 0xe0, 0xd2, 0x68, 0x65, 0x81, 0xdf,
	0x3c, 0x97, 0xce, 0xab, 0x5d, 0x00, 0x21, 0x9f, 0x18, 0x31, 0xb1, 0x4e, 0x03, 0xff, 0xfe, 0xd5,
	0xa6, 0x78, 0xf8, 0xe1, 0x15, 0x00, 0x00, 0xff, 0xff, 0xe1, 0xcc, 0xd4, 0xe4, 0xec, 0x00, 0x00,
	0x00,
}