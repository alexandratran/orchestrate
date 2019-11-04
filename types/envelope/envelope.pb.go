// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types/envelope/envelope.proto

package envelope

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	args "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/types/args"
	chain "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/types/chain"
	error1 "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/types/error"
	ethereum "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/types/ethereum"
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

// Metadata attached to an Envelope
type Metadata struct {
	// ID of the Envelope in UUID RFC 4122, ISO/IEC 9834-8:2005 format
	// e.g a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Extra information (optional)
	Extra                map[string]string `protobuf:"bytes,2,rep,name=extra,proto3" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_3335256a9be09a8d, []int{0}
}

func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Metadata) GetExtra() map[string]string {
	if m != nil {
		return m.Extra
	}
	return nil
}

// Args are provided when requesting an execution
type Args struct {
	// Arguments to craft transaction
	Call *args.Call `protobuf:"bytes,1,opt,name=call,proto3" json:"call,omitempty"`
	// Private arguments
	Private *args.Private `protobuf:"bytes,2,opt,name=private,proto3" json:"private,omitempty"`
	// Arbitrary data provided by user
	Data                 *ethereum.Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Args) Reset()         { *m = Args{} }
func (m *Args) String() string { return proto.CompactTextString(m) }
func (*Args) ProtoMessage()    {}
func (*Args) Descriptor() ([]byte, []int) {
	return fileDescriptor_3335256a9be09a8d, []int{1}
}

func (m *Args) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Args.Unmarshal(m, b)
}
func (m *Args) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Args.Marshal(b, m, deterministic)
}
func (m *Args) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Args.Merge(m, src)
}
func (m *Args) XXX_Size() int {
	return xxx_messageInfo_Args.Size(m)
}
func (m *Args) XXX_DiscardUnknown() {
	xxx_messageInfo_Args.DiscardUnknown(m)
}

var xxx_messageInfo_Args proto.InternalMessageInfo

func (m *Args) GetCall() *args.Call {
	if m != nil {
		return m.Call
	}
	return nil
}

func (m *Args) GetPrivate() *args.Private {
	if m != nil {
		return m.Private
	}
	return nil
}

func (m *Args) GetData() *ethereum.Data {
	if m != nil {
		return m.Data
	}
	return nil
}

// Envelope wraps all information contextual to the transaction orchestrated
type Envelope struct {
	// Chain the transaction is orchestrated for
	Chain *chain.Chain `protobuf:"bytes,1,opt,name=chain,proto3" json:"chain,omitempty"`
	// Protocol the transaction is orchestrated for
	Protocol *chain.Protocol `protobuf:"bytes,2,opt,name=protocol,proto3" json:"protocol,omitempty"`
	// Sender of the transaction
	// DATA (20 Bytes) - Ethereum Account Address
	// e.g 0xAf84242d70aE9D268E2bE3616ED497BA28A7b62C
	From *ethereum.Account `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	// Transaction
	Tx *ethereum.Transaction `protobuf:"bytes,4,opt,name=tx,proto3" json:"tx,omitempty"`
	// Receipt of the transaction
	Receipt *ethereum.Receipt `protobuf:"bytes,5,opt,name=receipt,proto3" json:"receipt,omitempty"`
	// Errors encountered while orchestrating the transaction
	Errors []*error1.Error `protobuf:"bytes,6,rep,name=errors,proto3" json:"errors,omitempty"`
	// Arguments provided by user
	Args *Args `protobuf:"bytes,7,opt,name=args,proto3" json:"args,omitempty"`
	// Metadata of the envelope
	Metadata             *Metadata `protobuf:"bytes,8,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Envelope) Reset()         { *m = Envelope{} }
func (m *Envelope) String() string { return proto.CompactTextString(m) }
func (*Envelope) ProtoMessage()    {}
func (*Envelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_3335256a9be09a8d, []int{2}
}

func (m *Envelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Envelope.Unmarshal(m, b)
}
func (m *Envelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Envelope.Marshal(b, m, deterministic)
}
func (m *Envelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Envelope.Merge(m, src)
}
func (m *Envelope) XXX_Size() int {
	return xxx_messageInfo_Envelope.Size(m)
}
func (m *Envelope) XXX_DiscardUnknown() {
	xxx_messageInfo_Envelope.DiscardUnknown(m)
}

var xxx_messageInfo_Envelope proto.InternalMessageInfo

func (m *Envelope) GetChain() *chain.Chain {
	if m != nil {
		return m.Chain
	}
	return nil
}

func (m *Envelope) GetProtocol() *chain.Protocol {
	if m != nil {
		return m.Protocol
	}
	return nil
}

func (m *Envelope) GetFrom() *ethereum.Account {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *Envelope) GetTx() *ethereum.Transaction {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *Envelope) GetReceipt() *ethereum.Receipt {
	if m != nil {
		return m.Receipt
	}
	return nil
}

func (m *Envelope) GetErrors() []*error1.Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *Envelope) GetArgs() *Args {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *Envelope) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*Metadata)(nil), "envelope.Metadata")
	proto.RegisterMapType((map[string]string)(nil), "envelope.Metadata.ExtraEntry")
	proto.RegisterType((*Args)(nil), "envelope.Args")
	proto.RegisterType((*Envelope)(nil), "envelope.Envelope")
}

func init() { proto.RegisterFile("types/envelope/envelope.proto", fileDescriptor_3335256a9be09a8d) }

var fileDescriptor_3335256a9be09a8d = []byte{
	// 494 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x93, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0xd5, 0x34, 0x6d, 0xc3, 0xe9, 0x18, 0x60, 0x31, 0x61, 0x22, 0x86, 0xaa, 0x88, 0x89,
	0x49, 0x13, 0x8e, 0xd4, 0xdd, 0x4c, 0xdc, 0x8d, 0x52, 0xc4, 0x0d, 0xd2, 0x64, 0xb8, 0xe2, 0xce,
	0x75, 0xbd, 0x2e, 0x5a, 0x1a, 0x47, 0xb6, 0x5b, 0xb5, 0x8f, 0xc0, 0xdb, 0xf1, 0x48, 0xc8, 0xff,
	0xda, 0xd2, 0xdd, 0x58, 0x27, 0xe7, 0xf7, 0xe5, 0xe4, 0x9c, 0xf3, 0x39, 0x70, 0x6e, 0xb6, 0xad,
	0xd0, 0xa5, 0x68, 0xd6, 0xa2, 0x96, 0xad, 0xd8, 0x05, 0xa4, 0x55, 0xd2, 0x48, 0x94, 0xc5, 0xe7,
	0xfc, 0x8d, 0x17, 0xf2, 0x07, 0x56, 0x35, 0xfe, 0xf4, 0x92, 0x3c, 0x3f, 0x04, 0x2e, 0xc5, 0x65,
	0x1d, 0xd8, 0xdb, 0x50, 0xdd, 0x3c, 0x08, 0x25, 0x56, 0xcb, 0x72, 0xc6, 0x74, 0xa8, 0x9c, 0x8f,
	0x8e, 0x90, 0x51, 0xac, 0xd1, 0x8c, 0x9b, 0x4a, 0xc6, 0xc2, 0xef, 0x8e, 0x14, 0x4a, 0x70, 0x51,
	0xb5, 0x26, 0x50, 0xec, 0x29, 0x53, 0x0b, 0x5d, 0xb6, 0xaa, 0x5a, 0x33, 0x13, 0x2b, 0x9f, 0x1d,
	0x10, 0xce, 0xea, 0xd8, 0x4b, 0x18, 0x40, 0x28, 0x25, 0x95, 0x3f, 0x3d, 0x28, 0xfe, 0x74, 0x20,
	0xfb, 0x21, 0x0c, 0x9b, 0x33, 0xc3, 0xd0, 0x29, 0x24, 0xd5, 0x1c, 0x77, 0x46, 0x9d, 0xcb, 0x67,
	0x34, 0xa9, 0xe6, 0xe8, 0x1a, 0x7a, 0x62, 0x63, 0x14, 0xc3, 0xc9, 0xa8, 0x7b, 0x39, 0x1c, 0x9f,
	0x93, 0xdd, 0x82, 0xe2, 0x2b, 0x64, 0x6a, 0xf9, 0xb4, 0x31, 0x6a, 0x4b, 0xbd, 0x36, 0xbf, 0x01,
	0xd8, 0x27, 0xd1, 0x4b, 0xe8, 0x3e, 0x8a, 0x6d, 0xa8, 0x69, 0x43, 0xf4, 0x1a, 0x7a, 0x6b, 0x56,
	0xaf, 0x04, 0x4e, 0x5c, 0xce, 0x3f, 0x7c, 0x4e, 0x6e, 0x3a, 0x85, 0x86, 0xf4, 0x56, 0x2d, 0x34,
	0x7a, 0x0f, 0xa9, 0x6d, 0xdd, 0xbd, 0x34, 0x1c, 0x03, 0xb1, 0xc3, 0x90, 0x09, 0xab, 0x6b, 0xea,
	0xf2, 0xe8, 0x23, 0x0c, 0xc2, 0xd0, 0xae, 0xc6, 0x70, 0xfc, 0xdc, 0x4b, 0xee, 0x7c, 0x92, 0x46,
	0x8a, 0x0a, 0x48, 0x6d, 0x93, 0xb8, 0xeb, 0x54, 0xa7, 0x24, 0x6e, 0x93, 0x7c, 0x65, 0x86, 0x51,
	0xc7, 0x8a, 0xbf, 0x09, 0x64, 0xd3, 0x30, 0x16, 0x2a, 0xa0, 0xe7, 0xac, 0x0c, 0x9f, 0x3e, 0x21,
	0xde, 0xeb, 0x89, 0x3d, 0xa9, 0x47, 0xe8, 0x0a, 0xb2, 0x68, 0x74, 0xf8, 0xfc, 0x8b, 0x20, 0xbb,
	0x0b, 0x69, 0xba, 0x13, 0xa0, 0x0b, 0x48, 0xef, 0x95, 0x5c, 0x86, 0x0e, 0x5e, 0xed, 0x3b, 0xb8,
	0xe5, 0x5c, 0xae, 0x1a, 0x43, 0x1d, 0x46, 0x17, 0x90, 0x98, 0x0d, 0x4e, 0x9d, 0xe8, 0x6c, 0x2f,
	0xfa, 0xb5, 0xbf, 0x16, 0x34, 0x31, 0x1b, 0x74, 0x05, 0x83, 0x70, 0x0f, 0x70, 0xef, 0xb8, 0x20,
	0xf5, 0x80, 0x46, 0x05, 0xfa, 0x00, 0x7d, 0x67, 0xb4, 0xc6, 0x7d, 0xe7, 0xde, 0x09, 0xf1, 0xbe,
	0x4f, 0xed, 0x49, 0x03, 0xb3, 0x2b, 0xb2, 0xbb, 0xc3, 0x83, 0xb8, 0xa2, 0xe8, 0xb0, 0x75, 0x82,
	0x3a, 0x86, 0x08, 0x64, 0xcb, 0xe0, 0x37, 0xce, 0x9c, 0x0e, 0x3d, 0xbd, 0x09, 0x74, 0xa7, 0xf9,
	0xf2, 0xfd, 0xf7, 0xb7, 0x45, 0x65, 0x6a, 0x36, 0x23, 0x5c, 0x2e, 0xcb, 0x89, 0x6c, 0xb4, 0x68,
	0x7e, 0x6e, 0x75, 0xc9, 0xeb, 0x4a, 0x34, 0xa6, 0xbc, 0x57, 0x25, 0x97, 0x4a, 0x7c, 0xd2, 0x86,
	0xf1, 0x47, 0x17, 0xba, 0x88, 0x2c, 0x2a, 0x53, 0xfe, 0xff, 0x43, 0xce, 0xfa, 0x6e, 0x91, 0xd7,
	0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x16, 0xcf, 0x91, 0xa9, 0x03, 0x00, 0x00,
}