// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orderer/configuration.proto

package orderer // import "github.com/hyperledger/udo/protos/orderer"

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

type ConsensusType struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	// Opaque metadata, dependent on the consensus type.
	Metadata             []byte   `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConsensusType) Reset()         { *m = ConsensusType{} }
func (m *ConsensusType) String() string { return proto.CompactTextString(m) }
func (*ConsensusType) ProtoMessage()    {}
func (*ConsensusType) Descriptor() ([]byte, []int) {
	return fileDescriptor_configuration_3289caae61166f4d, []int{0}
}
func (m *ConsensusType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsensusType.Unmarshal(m, b)
}
func (m *ConsensusType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsensusType.Marshal(b, m, deterministic)
}
func (dst *ConsensusType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsensusType.Merge(dst, src)
}
func (m *ConsensusType) XXX_Size() int {
	return xxx_messageInfo_ConsensusType.Size(m)
}
func (m *ConsensusType) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsensusType.DiscardUnknown(m)
}

var xxx_messageInfo_ConsensusType proto.InternalMessageInfo

func (m *ConsensusType) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ConsensusType) GetMetadata() []byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type BatchSize struct {
	// Simply specified as number of messages for now, in the future
	// we may want to allow this to be specified by size in bytes
	MaxMessageCount uint32 `protobuf:"varint,1,opt,name=max_message_count,json=maxMessageCount" json:"max_message_count,omitempty"`
	// The byte count of the serialized messages in a batch cannot
	// exceed this value.
	AbsoluteMaxBytes uint32 `protobuf:"varint,2,opt,name=absolute_max_bytes,json=absoluteMaxBytes" json:"absolute_max_bytes,omitempty"`
	// The byte count of the serialized messages in a batch should not
	// exceed this value.
	PreferredMaxBytes    uint32   `protobuf:"varint,3,opt,name=preferred_max_bytes,json=preferredMaxBytes" json:"preferred_max_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BatchSize) Reset()         { *m = BatchSize{} }
func (m *BatchSize) String() string { return proto.CompactTextString(m) }
func (*BatchSize) ProtoMessage()    {}
func (*BatchSize) Descriptor() ([]byte, []int) {
	return fileDescriptor_configuration_3289caae61166f4d, []int{1}
}
func (m *BatchSize) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchSize.Unmarshal(m, b)
}
func (m *BatchSize) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchSize.Marshal(b, m, deterministic)
}
func (dst *BatchSize) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchSize.Merge(dst, src)
}
func (m *BatchSize) XXX_Size() int {
	return xxx_messageInfo_BatchSize.Size(m)
}
func (m *BatchSize) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchSize.DiscardUnknown(m)
}

var xxx_messageInfo_BatchSize proto.InternalMessageInfo

func (m *BatchSize) GetMaxMessageCount() uint32 {
	if m != nil {
		return m.MaxMessageCount
	}
	return 0
}

func (m *BatchSize) GetAbsoluteMaxBytes() uint32 {
	if m != nil {
		return m.AbsoluteMaxBytes
	}
	return 0
}

func (m *BatchSize) GetPreferredMaxBytes() uint32 {
	if m != nil {
		return m.PreferredMaxBytes
	}
	return 0
}

type BatchTimeout struct {
	// Any duration string parseable by ParseDuration():
	// https://golang.org/pkg/time/#ParseDuration
	Timeout              string   `protobuf:"bytes,1,opt,name=timeout" json:"timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BatchTimeout) Reset()         { *m = BatchTimeout{} }
func (m *BatchTimeout) String() string { return proto.CompactTextString(m) }
func (*BatchTimeout) ProtoMessage()    {}
func (*BatchTimeout) Descriptor() ([]byte, []int) {
	return fileDescriptor_configuration_3289caae61166f4d, []int{2}
}
func (m *BatchTimeout) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchTimeout.Unmarshal(m, b)
}
func (m *BatchTimeout) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchTimeout.Marshal(b, m, deterministic)
}
func (dst *BatchTimeout) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchTimeout.Merge(dst, src)
}
func (m *BatchTimeout) XXX_Size() int {
	return xxx_messageInfo_BatchTimeout.Size(m)
}
func (m *BatchTimeout) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchTimeout.DiscardUnknown(m)
}

var xxx_messageInfo_BatchTimeout proto.InternalMessageInfo

func (m *BatchTimeout) GetTimeout() string {
	if m != nil {
		return m.Timeout
	}
	return ""
}

// Carries a list of bootstrap brokers, i.e. this is not the exclusive set of
// brokers an ordering service
type KafkaBrokers struct {
	// Each broker here should be identified using the (IP|host):port notation,
	// e.g. 127.0.0.1:7050, or localhost:7050 are valid entries
	Brokers              []string `protobuf:"bytes,1,rep,name=brokers" json:"brokers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KafkaBrokers) Reset()         { *m = KafkaBrokers{} }
func (m *KafkaBrokers) String() string { return proto.CompactTextString(m) }
func (*KafkaBrokers) ProtoMessage()    {}
func (*KafkaBrokers) Descriptor() ([]byte, []int) {
	return fileDescriptor_configuration_3289caae61166f4d, []int{3}
}
func (m *KafkaBrokers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KafkaBrokers.Unmarshal(m, b)
}
func (m *KafkaBrokers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KafkaBrokers.Marshal(b, m, deterministic)
}
func (dst *KafkaBrokers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KafkaBrokers.Merge(dst, src)
}
func (m *KafkaBrokers) XXX_Size() int {
	return xxx_messageInfo_KafkaBrokers.Size(m)
}
func (m *KafkaBrokers) XXX_DiscardUnknown() {
	xxx_messageInfo_KafkaBrokers.DiscardUnknown(m)
}

var xxx_messageInfo_KafkaBrokers proto.InternalMessageInfo

func (m *KafkaBrokers) GetBrokers() []string {
	if m != nil {
		return m.Brokers
	}
	return nil
}

// ChannelRestrictions is the mssage which conveys restrictions on channel creation for an orderer
type ChannelRestrictions struct {
	MaxCount             uint64   `protobuf:"varint,1,opt,name=max_count,json=maxCount" json:"max_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChannelRestrictions) Reset()         { *m = ChannelRestrictions{} }
func (m *ChannelRestrictions) String() string { return proto.CompactTextString(m) }
func (*ChannelRestrictions) ProtoMessage()    {}
func (*ChannelRestrictions) Descriptor() ([]byte, []int) {
	return fileDescriptor_configuration_3289caae61166f4d, []int{4}
}
func (m *ChannelRestrictions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChannelRestrictions.Unmarshal(m, b)
}
func (m *ChannelRestrictions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChannelRestrictions.Marshal(b, m, deterministic)
}
func (dst *ChannelRestrictions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelRestrictions.Merge(dst, src)
}
func (m *ChannelRestrictions) XXX_Size() int {
	return xxx_messageInfo_ChannelRestrictions.Size(m)
}
func (m *ChannelRestrictions) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelRestrictions.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelRestrictions proto.InternalMessageInfo

func (m *ChannelRestrictions) GetMaxCount() uint64 {
	if m != nil {
		return m.MaxCount
	}
	return 0
}

func init() {
	proto.RegisterType((*ConsensusType)(nil), "orderer.ConsensusType")
	proto.RegisterType((*BatchSize)(nil), "orderer.BatchSize")
	proto.RegisterType((*BatchTimeout)(nil), "orderer.BatchTimeout")
	proto.RegisterType((*KafkaBrokers)(nil), "orderer.KafkaBrokers")
	proto.RegisterType((*ChannelRestrictions)(nil), "orderer.ChannelRestrictions")
}

func init() {
	proto.RegisterFile("orderer/configuration.proto", fileDescriptor_configuration_3289caae61166f4d)
}

var fileDescriptor_configuration_3289caae61166f4d = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0x4f, 0x6b, 0xf2, 0x40,
	0x10, 0xc6, 0xc9, 0xab, 0xbc, 0xea, 0xa2, 0xbc, 0xaf, 0xeb, 0x25, 0xd4, 0x8b, 0x04, 0x0a, 0x52,
	0x24, 0x81, 0xf6, 0x03, 0x14, 0xe2, 0xb1, 0x78, 0x49, 0xed, 0xa5, 0x17, 0x99, 0x24, 0x93, 0x3f,
	0x68, 0x76, 0xc3, 0xec, 0x06, 0x92, 0x7e, 0x8f, 0x7e, 0xdf, 0xb2, 0x9b, 0x68, 0xbd, 0xcd, 0x33,
	0xcf, 0x6f, 0x87, 0x79, 0x76, 0xd8, 0x5a, 0x52, 0x8a, 0x84, 0x14, 0x24, 0x52, 0x64, 0x65, 0xde,
	0x10, 0xe8, 0x52, 0x0a, 0xbf, 0x26, 0xa9, 0x25, 0x9f, 0x0c, 0xa6, 0xf7, 0xca, 0x16, 0x7b, 0x29,
	0x14, 0x0a, 0xd5, 0xa8, 0x63, 0x57, 0x23, 0xe7, 0x6c, 0xac, 0xbb, 0x1a, 0x5d, 0x67, 0xe3, 0x6c,
	0x67, 0x91, 0xad, 0xf9, 0x03, 0x9b, 0x56, 0xa8, 0x21, 0x05, 0x0d, 0xee, 0x9f, 0x8d, 0xb3, 0x9d,
	0x47, 0x37, 0xed, 0x7d, 0x3b, 0x6c, 0x16, 0x82, 0x4e, 0x8a, 0xf7, 0xf2, 0x0b, 0xf9, 0x13, 0x5b,
	0x56, 0xd0, 0x9e, 0x2a, 0x54, 0x0a, 0x72, 0x3c, 0x25, 0xb2, 0x11, 0xda, 0x8e, 0x5a, 0x44, 0xff,
	0x2a, 0x68, 0x0f, 0x7d, 0x7f, 0x6f, 0xda, 0x7c, 0xc7, 0x38, 0xc4, 0x4a, 0x5e, 0x1a, 0x8d, 0x27,
	0xf3, 0x28, 0xee, 0x34, 0x2a, 0x3b, 0x7f, 0x11, 0xfd, 0xbf, 0x3a, 0x07, 0x68, 0x43, 0xd3, 0xe7,
	0x3e, 0x5b, 0xd5, 0x84, 0x19, 0x12, 0x61, 0x7a, 0x87, 0x8f, 0x2c, 0xbe, 0xbc, 0x59, 0x57, 0xde,
	0xdb, 0xb2, 0xb9, 0x5d, 0xeb, 0x58, 0x56, 0x28, 0x1b, 0xcd, 0x5d, 0x36, 0xd1, 0x7d, 0x39, 0x44,
	0xbb, 0x4a, 0x43, 0xbe, 0x41, 0x76, 0x86, 0x90, 0xe4, 0x19, 0x49, 0x19, 0x32, 0xee, 0x4b, 0xd7,
	0xd9, 0x8c, 0x0c, 0x39, 0x48, 0xef, 0x99, 0xad, 0xf6, 0x05, 0x08, 0x81, 0x97, 0x08, 0x95, 0xa6,
	0x32, 0x31, 0x3f, 0xaa, 0xf8, 0x9a, 0xcd, 0xcc, 0x42, 0xbf, 0x61, 0xc7, 0xd1, 0xb4, 0x82, 0xd6,
	0xa6, 0x0c, 0x3f, 0xd8, 0xa3, 0xa4, 0xdc, 0x2f, 0xba, 0x1a, 0xe9, 0x82, 0x69, 0x8e, 0xe4, 0x67,
	0x10, 0x53, 0x99, 0xf4, 0x97, 0x50, 0xfe, 0x70, 0x89, 0xcf, 0x5d, 0x5e, 0xea, 0xa2, 0x89, 0xfd,
	0x44, 0x56, 0xc1, 0x1d, 0x1d, 0xf4, 0x74, 0xd0, 0xd3, 0xc1, 0x40, 0xc7, 0x7f, 0xad, 0x7e, 0xf9,
	0x09, 0x00, 0x00, 0xff, 0xff, 0xb5, 0x9c, 0xb6, 0xa5, 0xe6, 0x01, 0x00, 0x00,
}
