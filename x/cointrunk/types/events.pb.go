// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bze/cointrunk/v1/events.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ArticleAddedEvent struct {
	Publisher string `protobuf:"bytes,1,opt,name=publisher,proto3" json:"publisher,omitempty"`
	ArticleId uint64 `protobuf:"varint,2,opt,name=article_id,json=articleId,proto3" json:"article_id,omitempty"`
	Paid      bool   `protobuf:"varint,3,opt,name=paid,proto3" json:"paid,omitempty"`
}

func (m *ArticleAddedEvent) Reset()         { *m = ArticleAddedEvent{} }
func (m *ArticleAddedEvent) String() string { return proto.CompactTextString(m) }
func (*ArticleAddedEvent) ProtoMessage()    {}
func (*ArticleAddedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_673ee938d5c50d20, []int{0}
}
func (m *ArticleAddedEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ArticleAddedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ArticleAddedEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ArticleAddedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArticleAddedEvent.Merge(m, src)
}
func (m *ArticleAddedEvent) XXX_Size() int {
	return m.Size()
}
func (m *ArticleAddedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ArticleAddedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ArticleAddedEvent proto.InternalMessageInfo

func (m *ArticleAddedEvent) GetPublisher() string {
	if m != nil {
		return m.Publisher
	}
	return ""
}

func (m *ArticleAddedEvent) GetArticleId() uint64 {
	if m != nil {
		return m.ArticleId
	}
	return 0
}

func (m *ArticleAddedEvent) GetPaid() bool {
	if m != nil {
		return m.Paid
	}
	return false
}

type PublisherAddedEvent struct {
	Publisher *Publisher `protobuf:"bytes,1,opt,name=publisher,proto3" json:"publisher,omitempty"`
}

func (m *PublisherAddedEvent) Reset()         { *m = PublisherAddedEvent{} }
func (m *PublisherAddedEvent) String() string { return proto.CompactTextString(m) }
func (*PublisherAddedEvent) ProtoMessage()    {}
func (*PublisherAddedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_673ee938d5c50d20, []int{1}
}
func (m *PublisherAddedEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PublisherAddedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PublisherAddedEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PublisherAddedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublisherAddedEvent.Merge(m, src)
}
func (m *PublisherAddedEvent) XXX_Size() int {
	return m.Size()
}
func (m *PublisherAddedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_PublisherAddedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_PublisherAddedEvent proto.InternalMessageInfo

func (m *PublisherAddedEvent) GetPublisher() *Publisher {
	if m != nil {
		return m.Publisher
	}
	return nil
}

type PublisherUpdatedEvent struct {
	Publisher *Publisher `protobuf:"bytes,1,opt,name=publisher,proto3" json:"publisher,omitempty"`
}

func (m *PublisherUpdatedEvent) Reset()         { *m = PublisherUpdatedEvent{} }
func (m *PublisherUpdatedEvent) String() string { return proto.CompactTextString(m) }
func (*PublisherUpdatedEvent) ProtoMessage()    {}
func (*PublisherUpdatedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_673ee938d5c50d20, []int{2}
}
func (m *PublisherUpdatedEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PublisherUpdatedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PublisherUpdatedEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PublisherUpdatedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublisherUpdatedEvent.Merge(m, src)
}
func (m *PublisherUpdatedEvent) XXX_Size() int {
	return m.Size()
}
func (m *PublisherUpdatedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_PublisherUpdatedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_PublisherUpdatedEvent proto.InternalMessageInfo

func (m *PublisherUpdatedEvent) GetPublisher() *Publisher {
	if m != nil {
		return m.Publisher
	}
	return nil
}

type AcceptedDomainAddedEvent struct {
	AcceptedDomain *AcceptedDomain `protobuf:"bytes,1,opt,name=accepted_domain,json=acceptedDomain,proto3" json:"accepted_domain,omitempty"`
}

func (m *AcceptedDomainAddedEvent) Reset()         { *m = AcceptedDomainAddedEvent{} }
func (m *AcceptedDomainAddedEvent) String() string { return proto.CompactTextString(m) }
func (*AcceptedDomainAddedEvent) ProtoMessage()    {}
func (*AcceptedDomainAddedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_673ee938d5c50d20, []int{3}
}
func (m *AcceptedDomainAddedEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AcceptedDomainAddedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AcceptedDomainAddedEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AcceptedDomainAddedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AcceptedDomainAddedEvent.Merge(m, src)
}
func (m *AcceptedDomainAddedEvent) XXX_Size() int {
	return m.Size()
}
func (m *AcceptedDomainAddedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_AcceptedDomainAddedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_AcceptedDomainAddedEvent proto.InternalMessageInfo

func (m *AcceptedDomainAddedEvent) GetAcceptedDomain() *AcceptedDomain {
	if m != nil {
		return m.AcceptedDomain
	}
	return nil
}

type AcceptedDomainUpdatedEvent struct {
	AcceptedDomain *AcceptedDomain `protobuf:"bytes,1,opt,name=accepted_domain,json=acceptedDomain,proto3" json:"accepted_domain,omitempty"`
}

func (m *AcceptedDomainUpdatedEvent) Reset()         { *m = AcceptedDomainUpdatedEvent{} }
func (m *AcceptedDomainUpdatedEvent) String() string { return proto.CompactTextString(m) }
func (*AcceptedDomainUpdatedEvent) ProtoMessage()    {}
func (*AcceptedDomainUpdatedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_673ee938d5c50d20, []int{4}
}
func (m *AcceptedDomainUpdatedEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AcceptedDomainUpdatedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AcceptedDomainUpdatedEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AcceptedDomainUpdatedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AcceptedDomainUpdatedEvent.Merge(m, src)
}
func (m *AcceptedDomainUpdatedEvent) XXX_Size() int {
	return m.Size()
}
func (m *AcceptedDomainUpdatedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_AcceptedDomainUpdatedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_AcceptedDomainUpdatedEvent proto.InternalMessageInfo

func (m *AcceptedDomainUpdatedEvent) GetAcceptedDomain() *AcceptedDomain {
	if m != nil {
		return m.AcceptedDomain
	}
	return nil
}

type PublisherRespectPaidEvent struct {
	RespectPaid        uint64 `protobuf:"varint,1,opt,name=respect_paid,json=respectPaid,proto3" json:"respect_paid,omitempty"`
	PublisherReward    uint64 `protobuf:"varint,2,opt,name=publisher_reward,json=publisherReward,proto3" json:"publisher_reward,omitempty"`
	CommunityPoolFunds uint64 `protobuf:"varint,3,opt,name=community_pool_funds,json=communityPoolFunds,proto3" json:"community_pool_funds,omitempty"`
	Publisher          string `protobuf:"bytes,4,opt,name=publisher,proto3" json:"publisher,omitempty"`
}

func (m *PublisherRespectPaidEvent) Reset()         { *m = PublisherRespectPaidEvent{} }
func (m *PublisherRespectPaidEvent) String() string { return proto.CompactTextString(m) }
func (*PublisherRespectPaidEvent) ProtoMessage()    {}
func (*PublisherRespectPaidEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_673ee938d5c50d20, []int{5}
}
func (m *PublisherRespectPaidEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PublisherRespectPaidEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PublisherRespectPaidEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PublisherRespectPaidEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublisherRespectPaidEvent.Merge(m, src)
}
func (m *PublisherRespectPaidEvent) XXX_Size() int {
	return m.Size()
}
func (m *PublisherRespectPaidEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_PublisherRespectPaidEvent.DiscardUnknown(m)
}

var xxx_messageInfo_PublisherRespectPaidEvent proto.InternalMessageInfo

func (m *PublisherRespectPaidEvent) GetRespectPaid() uint64 {
	if m != nil {
		return m.RespectPaid
	}
	return 0
}

func (m *PublisherRespectPaidEvent) GetPublisherReward() uint64 {
	if m != nil {
		return m.PublisherReward
	}
	return 0
}

func (m *PublisherRespectPaidEvent) GetCommunityPoolFunds() uint64 {
	if m != nil {
		return m.CommunityPoolFunds
	}
	return 0
}

func (m *PublisherRespectPaidEvent) GetPublisher() string {
	if m != nil {
		return m.Publisher
	}
	return ""
}

func init() {
	proto.RegisterType((*ArticleAddedEvent)(nil), "bze.cointrunk.v1.ArticleAddedEvent")
	proto.RegisterType((*PublisherAddedEvent)(nil), "bze.cointrunk.v1.PublisherAddedEvent")
	proto.RegisterType((*PublisherUpdatedEvent)(nil), "bze.cointrunk.v1.PublisherUpdatedEvent")
	proto.RegisterType((*AcceptedDomainAddedEvent)(nil), "bze.cointrunk.v1.AcceptedDomainAddedEvent")
	proto.RegisterType((*AcceptedDomainUpdatedEvent)(nil), "bze.cointrunk.v1.AcceptedDomainUpdatedEvent")
	proto.RegisterType((*PublisherRespectPaidEvent)(nil), "bze.cointrunk.v1.PublisherRespectPaidEvent")
}

func init() { proto.RegisterFile("bze/cointrunk/v1/events.proto", fileDescriptor_673ee938d5c50d20) }

var fileDescriptor_673ee938d5c50d20 = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x4d, 0xaf, 0x93, 0x40,
	0x14, 0xed, 0x28, 0x31, 0x76, 0x9e, 0xf1, 0x3d, 0x47, 0x4d, 0xb0, 0xfa, 0x08, 0xb2, 0x30, 0x75,
	0x21, 0xf8, 0x74, 0xe5, 0xb2, 0x46, 0x8d, 0xdd, 0x11, 0x12, 0x37, 0x6e, 0xc8, 0xc0, 0x5c, 0xdb,
	0x89, 0xc0, 0x4c, 0x86, 0xa1, 0xda, 0xfe, 0x0a, 0xff, 0x8e, 0xff, 0xc0, 0x65, 0x97, 0x2e, 0x4d,
	0xfb, 0x47, 0x4c, 0x81, 0xf2, 0xd5, 0x6d, 0x77, 0x70, 0xee, 0x3d, 0xe7, 0xdc, 0xb9, 0xf7, 0xe0,
	0xeb, 0x68, 0x03, 0x5e, 0x2c, 0x78, 0xa6, 0x55, 0x91, 0x7d, 0xf7, 0x56, 0x37, 0x1e, 0xac, 0x20,
	0xd3, 0xb9, 0x2b, 0x95, 0xd0, 0x82, 0x5c, 0x45, 0x1b, 0x70, 0x9b, 0xb2, 0xbb, 0xba, 0x99, 0xd8,
	0x27, 0x04, 0x59, 0x44, 0x09, 0xcf, 0x97, 0xa0, 0x2a, 0xce, 0xe4, 0xc5, 0x49, 0x07, 0x8d, 0x63,
	0x90, 0x1a, 0x58, 0xc8, 0x44, 0x4a, 0x79, 0x56, 0xf5, 0x39, 0x0c, 0x3f, 0x98, 0x29, 0xcd, 0xe3,
	0x04, 0x66, 0x8c, 0x01, 0xfb, 0x78, 0xf0, 0x25, 0xcf, 0xf0, 0xb8, 0xd1, 0x33, 0x91, 0x8d, 0xa6,
	0xe3, 0xa0, 0x05, 0xc8, 0x35, 0xc6, 0xb4, 0xa2, 0x84, 0x9c, 0x99, 0xb7, 0x6c, 0x34, 0x35, 0x82,
	0x71, 0x8d, 0xcc, 0x19, 0x21, 0xd8, 0x90, 0x94, 0x33, 0xf3, 0xb6, 0x8d, 0xa6, 0x77, 0x83, 0xf2,
	0xdb, 0xf1, 0xf1, 0x43, 0xff, 0xc8, 0xef, 0xf8, 0xbc, 0x1b, 0xfa, 0x5c, 0xbc, 0x79, 0xea, 0x0e,
	0x1f, 0xeb, 0x36, 0xcc, 0xce, 0x10, 0x4e, 0x80, 0x1f, 0x37, 0xf8, 0x17, 0xc9, 0xa8, 0x3e, 0x83,
	0x26, 0x60, 0x73, 0x56, 0x2f, 0xe9, 0x43, 0xb9, 0xa3, 0xce, 0xa8, 0x73, 0x7c, 0x39, 0x58, 0x60,
	0x2d, 0x6e, 0x9f, 0x8a, 0xf7, 0x45, 0x82, 0xfb, 0xb4, 0xf7, 0xef, 0x2c, 0xf0, 0xa4, 0xdf, 0xd1,
	0x9b, 0xff, 0x8c, 0x46, 0xbf, 0x11, 0x7e, 0xd2, 0x3e, 0x14, 0x72, 0x09, 0xb1, 0xf6, 0x29, 0xaf,
	0x8d, 0x9e, 0xe3, 0x7b, 0xaa, 0xc2, 0xc2, 0xf2, 0x5e, 0xa8, 0x3c, 0xe4, 0x85, 0x6a, 0xfb, 0xc8,
	0x4b, 0x7c, 0xd5, 0x6c, 0x27, 0x54, 0xf0, 0x83, 0xaa, 0xe3, 0xbd, 0x2f, 0x65, 0xab, 0x7b, 0x80,
	0xc9, 0x6b, 0xfc, 0x28, 0x16, 0x69, 0x5a, 0x64, 0x5c, 0xaf, 0x43, 0x29, 0x44, 0x12, 0x7e, 0x2b,
	0x32, 0x96, 0x97, 0x29, 0x30, 0x02, 0xd2, 0xd4, 0x7c, 0x21, 0x92, 0x4f, 0x87, 0x4a, 0x3f, 0x64,
	0xc6, 0x20, 0x64, 0xef, 0x3f, 0xff, 0xd9, 0x59, 0x68, 0xbb, 0xb3, 0xd0, 0xbf, 0x9d, 0x85, 0x7e,
	0xed, 0xad, 0xd1, 0x76, 0x6f, 0x8d, 0xfe, 0xee, 0xad, 0xd1, 0x57, 0x77, 0xc1, 0xf5, 0xb2, 0x88,
	0xdc, 0x58, 0xa4, 0x5e, 0xb4, 0x81, 0x57, 0x34, 0x91, 0x4b, 0xaa, 0x81, 0x96, 0x7f, 0xde, 0xcf,
	0x4e, 0xe8, 0xf5, 0x5a, 0x42, 0x1e, 0xdd, 0x29, 0x83, 0xfe, 0xf6, 0x7f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xae, 0x03, 0x61, 0xca, 0x65, 0x03, 0x00, 0x00,
}

func (m *ArticleAddedEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ArticleAddedEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ArticleAddedEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Paid {
		i--
		if m.Paid {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.ArticleId != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.ArticleId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Publisher) > 0 {
		i -= len(m.Publisher)
		copy(dAtA[i:], m.Publisher)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Publisher)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PublisherAddedEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PublisherAddedEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PublisherAddedEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Publisher != nil {
		{
			size, err := m.Publisher.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvents(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PublisherUpdatedEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PublisherUpdatedEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PublisherUpdatedEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Publisher != nil {
		{
			size, err := m.Publisher.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvents(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AcceptedDomainAddedEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AcceptedDomainAddedEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AcceptedDomainAddedEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AcceptedDomain != nil {
		{
			size, err := m.AcceptedDomain.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvents(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AcceptedDomainUpdatedEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AcceptedDomainUpdatedEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AcceptedDomainUpdatedEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AcceptedDomain != nil {
		{
			size, err := m.AcceptedDomain.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvents(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PublisherRespectPaidEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PublisherRespectPaidEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PublisherRespectPaidEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Publisher) > 0 {
		i -= len(m.Publisher)
		copy(dAtA[i:], m.Publisher)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Publisher)))
		i--
		dAtA[i] = 0x22
	}
	if m.CommunityPoolFunds != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.CommunityPoolFunds))
		i--
		dAtA[i] = 0x18
	}
	if m.PublisherReward != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.PublisherReward))
		i--
		dAtA[i] = 0x10
	}
	if m.RespectPaid != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.RespectPaid))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ArticleAddedEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Publisher)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.ArticleId != 0 {
		n += 1 + sovEvents(uint64(m.ArticleId))
	}
	if m.Paid {
		n += 2
	}
	return n
}

func (m *PublisherAddedEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Publisher != nil {
		l = m.Publisher.Size()
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *PublisherUpdatedEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Publisher != nil {
		l = m.Publisher.Size()
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *AcceptedDomainAddedEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AcceptedDomain != nil {
		l = m.AcceptedDomain.Size()
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *AcceptedDomainUpdatedEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AcceptedDomain != nil {
		l = m.AcceptedDomain.Size()
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *PublisherRespectPaidEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RespectPaid != 0 {
		n += 1 + sovEvents(uint64(m.RespectPaid))
	}
	if m.PublisherReward != 0 {
		n += 1 + sovEvents(uint64(m.PublisherReward))
	}
	if m.CommunityPoolFunds != 0 {
		n += 1 + sovEvents(uint64(m.CommunityPoolFunds))
	}
	l = len(m.Publisher)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ArticleAddedEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ArticleAddedEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ArticleAddedEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Publisher", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Publisher = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ArticleId", wireType)
			}
			m.ArticleId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ArticleId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Paid", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Paid = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PublisherAddedEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PublisherAddedEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PublisherAddedEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Publisher", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Publisher == nil {
				m.Publisher = &Publisher{}
			}
			if err := m.Publisher.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PublisherUpdatedEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PublisherUpdatedEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PublisherUpdatedEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Publisher", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Publisher == nil {
				m.Publisher = &Publisher{}
			}
			if err := m.Publisher.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AcceptedDomainAddedEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AcceptedDomainAddedEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AcceptedDomainAddedEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AcceptedDomain", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AcceptedDomain == nil {
				m.AcceptedDomain = &AcceptedDomain{}
			}
			if err := m.AcceptedDomain.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AcceptedDomainUpdatedEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AcceptedDomainUpdatedEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AcceptedDomainUpdatedEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AcceptedDomain", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AcceptedDomain == nil {
				m.AcceptedDomain = &AcceptedDomain{}
			}
			if err := m.AcceptedDomain.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PublisherRespectPaidEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PublisherRespectPaidEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PublisherRespectPaidEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RespectPaid", wireType)
			}
			m.RespectPaid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RespectPaid |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublisherReward", wireType)
			}
			m.PublisherReward = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PublisherReward |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommunityPoolFunds", wireType)
			}
			m.CommunityPoolFunds = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CommunityPoolFunds |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Publisher", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Publisher = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
