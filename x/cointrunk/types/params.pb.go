// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cointrunk/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// Params defines the parameters for the module.
type PublisherRespectParams struct {
	Tax   github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=tax,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"tax" yaml:"tax"`
	Denom string                                 `protobuf:"bytes,5,opt,name=denom,proto3" json:"denom,omitempty" yaml:"denom"`
}

func (m *PublisherRespectParams) Reset()         { *m = PublisherRespectParams{} }
func (m *PublisherRespectParams) String() string { return proto.CompactTextString(m) }
func (*PublisherRespectParams) ProtoMessage()    {}
func (*PublisherRespectParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_f29b1ecdb66bf452, []int{0}
}
func (m *PublisherRespectParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PublisherRespectParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PublisherRespectParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PublisherRespectParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublisherRespectParams.Merge(m, src)
}
func (m *PublisherRespectParams) XXX_Size() int {
	return m.Size()
}
func (m *PublisherRespectParams) XXX_DiscardUnknown() {
	xxx_messageInfo_PublisherRespectParams.DiscardUnknown(m)
}

var xxx_messageInfo_PublisherRespectParams proto.InternalMessageInfo

func (m *PublisherRespectParams) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

type Params struct {
	AnonArticleLimit       uint64                 `protobuf:"varint,1,opt,name=anonArticleLimit,proto3" json:"anonArticleLimit,omitempty" yaml:"anon_article_limit"`
	AnonArticleCost        types.Coin             `protobuf:"bytes,2,opt,name=anonArticleCost,proto3" json:"anon_article_cost,omitempty" yaml:"anon_article_cost"`
	PublisherRespectParams PublisherRespectParams `protobuf:"bytes,3,opt,name=publisherRespectParams,proto3" json:"publisher_respect,omitempty" yaml:"publisher_respect"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_f29b1ecdb66bf452, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetAnonArticleLimit() uint64 {
	if m != nil {
		return m.AnonArticleLimit
	}
	return 0
}

func (m *Params) GetAnonArticleCost() types.Coin {
	if m != nil {
		return m.AnonArticleCost
	}
	return types.Coin{}
}

func (m *Params) GetPublisherRespectParams() PublisherRespectParams {
	if m != nil {
		return m.PublisherRespectParams
	}
	return PublisherRespectParams{}
}

func init() {
	proto.RegisterType((*PublisherRespectParams)(nil), "bze.cointrunk.v1.PublisherRespectParams")
	proto.RegisterType((*Params)(nil), "bze.cointrunk.v1.Params")
}

func init() { proto.RegisterFile("cointrunk/params.proto", fileDescriptor_f29b1ecdb66bf452) }

var fileDescriptor_f29b1ecdb66bf452 = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x31, 0x6b, 0xdc, 0x30,
	0x14, 0xc7, 0xad, 0x5e, 0x12, 0xa8, 0x5a, 0xe8, 0x61, 0xca, 0xe1, 0xa4, 0xd4, 0x0a, 0x1a, 0xc2,
	0x0d, 0x8d, 0xc4, 0xb5, 0x5b, 0x28, 0x94, 0x3a, 0x1d, 0x5a, 0x28, 0x25, 0x78, 0xec, 0x72, 0xc8,
	0x8e, 0xb8, 0x13, 0xb1, 0x2c, 0x63, 0xe9, 0xc2, 0x5d, 0xe6, 0x7c, 0x80, 0x8c, 0x1d, 0xdb, 0x6f,
	0x93, 0x31, 0x63, 0xe9, 0x60, 0xca, 0xdd, 0xd6, 0xd1, 0x9f, 0xa0, 0x48, 0x3a, 0x92, 0x34, 0xe7,
	0x4e, 0xb6, 0xf4, 0x7e, 0xef, 0xbd, 0xbf, 0xde, 0xfb, 0xc3, 0x41, 0xae, 0x44, 0x69, 0xea, 0x59,
	0x79, 0x46, 0x2b, 0x56, 0x33, 0xa9, 0x49, 0x55, 0x2b, 0xa3, 0xc2, 0x7e, 0x76, 0xc1, 0xc9, 0x6d,
	0x8c, 0x9c, 0x8f, 0xf6, 0x9e, 0x4f, 0xd4, 0x44, 0xb9, 0x20, 0xb5, 0x7f, 0x9e, 0xdb, 0x8b, 0x73,
	0xa5, 0xa5, 0xd2, 0x34, 0x63, 0x9a, 0xd3, 0xf3, 0x51, 0xc6, 0x0d, 0x1b, 0x51, 0x9b, 0xe7, 0xe3,
	0xf8, 0x0a, 0xc0, 0xc1, 0xc9, 0x2c, 0x2b, 0x84, 0x9e, 0xf2, 0x3a, 0xe5, 0xba, 0xe2, 0xb9, 0x39,
	0x71, 0x8d, 0xc2, 0x2f, 0xb0, 0x67, 0xd8, 0x3c, 0x02, 0xfb, 0x60, 0xf8, 0x38, 0x79, 0x7b, 0xdd,
	0xa0, 0xe0, 0x57, 0x83, 0x0e, 0x26, 0xc2, 0x4c, 0x67, 0x19, 0xc9, 0x95, 0xa4, 0xeb, 0xd2, 0xfe,
	0x73, 0xa8, 0x4f, 0xcf, 0xa8, 0x59, 0x54, 0x5c, 0x93, 0x0f, 0x3c, 0x6f, 0x1b, 0x04, 0x17, 0x4c,
	0x16, 0x47, 0xd8, 0xb0, 0x39, 0x4e, 0x6d, 0xa1, 0xf0, 0x00, 0x6e, 0x9f, 0xf2, 0x52, 0xc9, 0x68,
	0xdb, 0x55, 0xec, 0xb7, 0x0d, 0x7a, 0xea, 0x19, 0x77, 0x8d, 0x53, 0x1f, 0xc6, 0x97, 0x3d, 0xb8,
	0xb3, 0x96, 0xf0, 0x09, 0xf6, 0x59, 0xa9, 0xca, 0xf7, 0xb5, 0x11, 0x79, 0xc1, 0x3f, 0x0b, 0x29,
	0x8c, 0xd3, 0xb3, 0x95, 0xbc, 0x6c, 0x1b, 0xb4, 0xeb, 0xb3, 0x2d, 0x31, 0x66, 0x1e, 0x19, 0x17,
	0x96, 0xc1, 0xe9, 0x46, 0x5a, 0x78, 0x09, 0xe0, 0xb3, 0x7b, 0x97, 0xc7, 0x4a, 0x9b, 0xe8, 0xd1,
	0x3e, 0x18, 0x3e, 0x79, 0xbd, 0x4b, 0xfc, 0x0b, 0x88, 0x9d, 0x11, 0x59, 0xcf, 0x88, 0x1c, 0x2b,
	0x51, 0x26, 0xef, 0xec, 0xab, 0xff, 0x34, 0xe8, 0xc5, 0x3f, 0x3d, 0x72, 0xa5, 0xcd, 0x2b, 0x25,
	0x85, 0xe1, 0xb2, 0x32, 0x8b, 0xb6, 0x41, 0x51, 0x87, 0x10, 0x0b, 0xe1, 0xf4, 0x61, 0xcb, 0xf0,
	0x07, 0x80, 0x83, 0xaa, 0x73, 0xde, 0x51, 0xcf, 0xa9, 0x19, 0x92, 0x87, 0x9b, 0x25, 0xdd, 0xfb,
	0xb9, 0x13, 0x77, 0x5b, 0x6f, 0x5c, 0x7b, 0xa0, 0x4b, 0xdc, 0x06, 0x84, 0xd3, 0xff, 0x08, 0x39,
	0xda, 0xfa, 0xf6, 0x1d, 0x05, 0xc9, 0xc7, 0xeb, 0x65, 0x0c, 0x6e, 0x96, 0x31, 0xf8, 0xbd, 0x8c,
	0xc1, 0xd5, 0x2a, 0x0e, 0x6e, 0x56, 0x71, 0xf0, 0x73, 0x15, 0x07, 0x5f, 0xc9, 0x3d, 0x0f, 0x64,
	0x17, 0xfc, 0x90, 0x15, 0xd5, 0x94, 0x19, 0xce, 0xdc, 0x89, 0xce, 0xe9, 0x9d, 0x65, 0x9d, 0x1f,
	0xb2, 0x1d, 0x67, 0xb5, 0x37, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x1e, 0x82, 0x0e, 0x47, 0xcc,
	0x02, 0x00, 0x00,
}

func (m *PublisherRespectParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PublisherRespectParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PublisherRespectParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x2a
	}
	{
		size := m.Tax.Size()
		i -= size
		if _, err := m.Tax.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.PublisherRespectParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.AnonArticleCost.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.AnonArticleLimit != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.AnonArticleLimit))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PublisherRespectParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Tax.Size()
	n += 1 + l + sovParams(uint64(l))
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AnonArticleLimit != 0 {
		n += 1 + sovParams(uint64(m.AnonArticleLimit))
	}
	l = m.AnonArticleCost.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.PublisherRespectParams.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PublisherRespectParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: PublisherRespectParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PublisherRespectParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tax", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Tax.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AnonArticleLimit", wireType)
			}
			m.AnonArticleLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AnonArticleLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AnonArticleCost", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AnonArticleCost.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublisherRespectParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PublisherRespectParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
