// Code generated by protoc-gen-go. DO NOT EDIT.
// source: unwrap_request_evm.proto

package events

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

type UnwrapRequestEvmProto struct {
	NetworkType          uint32   `protobuf:"varint,1,opt,name=networkType,proto3" json:"networkType,omitempty"`
	ChainId              uint32   `protobuf:"varint,2,opt,name=chainId,proto3" json:"chainId,omitempty"`
	BlockNumber          uint64   `protobuf:"varint,3,opt,name=blockNumber,proto3" json:"blockNumber,omitempty"`
	BlockHash            []byte   `protobuf:"bytes,4,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
	TransactionHash      []byte   `protobuf:"bytes,5,opt,name=transactionHash,proto3" json:"transactionHash,omitempty"`
	LogIndex             uint32   `protobuf:"varint,6,opt,name=logIndex,proto3" json:"logIndex,omitempty"`
	From                 []byte   `protobuf:"bytes,7,opt,name=from,proto3" json:"from,omitempty"`
	To                   string   `protobuf:"bytes,8,opt,name=to,proto3" json:"to,omitempty"`
	Token                []byte   `protobuf:"bytes,9,opt,name=token,proto3" json:"token,omitempty"`
	Amount               []byte   `protobuf:"bytes,10,opt,name=amount,proto3" json:"amount,omitempty"`
	Signature            string   `protobuf:"bytes,11,opt,name=signature,proto3" json:"signature,omitempty"`
	RedeemStatus         uint32   `protobuf:"varint,12,opt,name=redeemStatus,proto3" json:"redeemStatus,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnwrapRequestEvmProto) Reset()         { *m = UnwrapRequestEvmProto{} }
func (m *UnwrapRequestEvmProto) String() string { return proto.CompactTextString(m) }
func (*UnwrapRequestEvmProto) ProtoMessage()    {}
func (*UnwrapRequestEvmProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a195c65809385cd, []int{0}
}

func (m *UnwrapRequestEvmProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnwrapRequestEvmProto.Unmarshal(m, b)
}
func (m *UnwrapRequestEvmProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnwrapRequestEvmProto.Marshal(b, m, deterministic)
}
func (m *UnwrapRequestEvmProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnwrapRequestEvmProto.Merge(m, src)
}
func (m *UnwrapRequestEvmProto) XXX_Size() int {
	return xxx_messageInfo_UnwrapRequestEvmProto.Size(m)
}
func (m *UnwrapRequestEvmProto) XXX_DiscardUnknown() {
	xxx_messageInfo_UnwrapRequestEvmProto.DiscardUnknown(m)
}

var xxx_messageInfo_UnwrapRequestEvmProto proto.InternalMessageInfo

func (m *UnwrapRequestEvmProto) GetNetworkType() uint32 {
	if m != nil {
		return m.NetworkType
	}
	return 0
}

func (m *UnwrapRequestEvmProto) GetChainId() uint32 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *UnwrapRequestEvmProto) GetBlockNumber() uint64 {
	if m != nil {
		return m.BlockNumber
	}
	return 0
}

func (m *UnwrapRequestEvmProto) GetBlockHash() []byte {
	if m != nil {
		return m.BlockHash
	}
	return nil
}

func (m *UnwrapRequestEvmProto) GetTransactionHash() []byte {
	if m != nil {
		return m.TransactionHash
	}
	return nil
}

func (m *UnwrapRequestEvmProto) GetLogIndex() uint32 {
	if m != nil {
		return m.LogIndex
	}
	return 0
}

func (m *UnwrapRequestEvmProto) GetFrom() []byte {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *UnwrapRequestEvmProto) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *UnwrapRequestEvmProto) GetToken() []byte {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *UnwrapRequestEvmProto) GetAmount() []byte {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *UnwrapRequestEvmProto) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *UnwrapRequestEvmProto) GetRedeemStatus() uint32 {
	if m != nil {
		return m.RedeemStatus
	}
	return 0
}

func init() {
	proto.RegisterType((*UnwrapRequestEvmProto)(nil), "events.UnwrapRequestEvmProto")
}

func init() { proto.RegisterFile("unwrap_request_evm.proto", fileDescriptor_0a195c65809385cd) }

var fileDescriptor_0a195c65809385cd = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x41, 0x4f, 0x32, 0x31,
	0x10, 0x86, 0xb3, 0xfb, 0xc1, 0x02, 0x03, 0x9f, 0x26, 0x8d, 0x9a, 0x89, 0xf1, 0xb0, 0xe1, 0xb4,
	0x27, 0x3c, 0xf8, 0x0f, 0x4c, 0x4c, 0xe4, 0x62, 0xcc, 0xaa, 0x17, 0x2f, 0xa4, 0xc0, 0x08, 0x04,
	0xda, 0xae, 0xed, 0x74, 0xd1, 0x1f, 0xed, 0x7f, 0x30, 0xcc, 0xaa, 0xa0, 0xb7, 0xbe, 0x4f, 0x9f,
	0x69, 0xde, 0x74, 0x00, 0xa3, 0xdd, 0x7a, 0x5d, 0x4d, 0x3c, 0xbd, 0x46, 0x0a, 0x3c, 0xa1, 0xda,
	0x8c, 0x2a, 0xef, 0xd8, 0xa9, 0x8c, 0x6a, 0xb2, 0x1c, 0x86, 0x1f, 0x29, 0x9c, 0x3e, 0x89, 0x54,
	0x36, 0xce, 0x4d, 0x6d, 0xee, 0xc5, 0xc8, 0xa1, 0x6f, 0x89, 0xb7, 0xce, 0xaf, 0x1f, 0xdf, 0x2b,
	0xc2, 0x24, 0x4f, 0x8a, 0xff, 0xe5, 0x21, 0x52, 0x08, 0x9d, 0xd9, 0x52, 0xaf, 0xec, 0x78, 0x8e,
	0xa9, 0xdc, 0x7e, 0xc7, 0xdd, 0xec, 0x74, 0xe3, 0x66, 0xeb, 0xbb, 0x68, 0xa6, 0xe4, 0xf1, 0x5f,
	0x9e, 0x14, 0xad, 0xf2, 0x10, 0xa9, 0x0b, 0xe8, 0x49, 0xbc, 0xd5, 0x61, 0x89, 0xad, 0x3c, 0x29,
	0x06, 0xe5, 0x1e, 0xa8, 0x02, 0x8e, 0xd9, 0x6b, 0x1b, 0xf4, 0x8c, 0x57, 0xce, 0x8a, 0xd3, 0x16,
	0xe7, 0x2f, 0x56, 0xe7, 0xd0, 0xdd, 0xb8, 0xc5, 0xd8, 0xce, 0xe9, 0x0d, 0x33, 0x29, 0xf1, 0x93,
	0x95, 0x82, 0xd6, 0x8b, 0x77, 0x06, 0x3b, 0x32, 0x2a, 0x67, 0x75, 0x04, 0x29, 0x3b, 0xec, 0xe6,
	0x49, 0xd1, 0x2b, 0x53, 0x76, 0xea, 0x04, 0xda, 0xec, 0xd6, 0x64, 0xb1, 0x27, 0x52, 0x13, 0xd4,
	0x19, 0x64, 0xda, 0xb8, 0x68, 0x19, 0x41, 0xf0, 0x57, 0xda, 0xb5, 0x0e, 0xab, 0x85, 0xd5, 0x1c,
	0x3d, 0x61, 0x5f, 0x1e, 0xd9, 0x03, 0x35, 0x84, 0x81, 0xa7, 0x39, 0x91, 0x79, 0x60, 0xcd, 0x31,
	0xe0, 0x40, 0xfa, 0xfc, 0x62, 0xd7, 0xf0, 0xdc, 0x1d, 0x5d, 0x36, 0x7f, 0x3f, 0xcd, 0x64, 0x15,
	0x57, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9c, 0xda, 0x62, 0xf1, 0xa6, 0x01, 0x00, 0x00,
}