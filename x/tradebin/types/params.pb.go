// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tradebin/params.proto

package types

import (
	fmt "fmt"
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
type Params struct {
	CreateMarketFee     string `protobuf:"bytes,1,opt,name=createMarketFee,proto3" json:"createMarketFee,omitempty" yaml:"create_market_fee"`
	MarketMakerFee      string `protobuf:"bytes,2,opt,name=marketMakerFee,proto3" json:"marketMakerFee,omitempty" yaml:"market_maker_fee"`
	MarketTakerFee      string `protobuf:"bytes,3,opt,name=marketTakerFee,proto3" json:"marketTakerFee,omitempty" yaml:"market_taker_fee"`
	MakerFeeDestination string `protobuf:"bytes,4,opt,name=makerFeeDestination,proto3" json:"makerFeeDestination,omitempty" yaml:"maker_fee_destination"`
	TakerFeeDestination string `protobuf:"bytes,5,opt,name=takerFeeDestination,proto3" json:"takerFeeDestination,omitempty" yaml:"taker_fee_destination"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_db27696fb9409af7, []int{0}
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

func (m *Params) GetCreateMarketFee() string {
	if m != nil {
		return m.CreateMarketFee
	}
	return ""
}

func (m *Params) GetMarketMakerFee() string {
	if m != nil {
		return m.MarketMakerFee
	}
	return ""
}

func (m *Params) GetMarketTakerFee() string {
	if m != nil {
		return m.MarketTakerFee
	}
	return ""
}

func (m *Params) GetMakerFeeDestination() string {
	if m != nil {
		return m.MakerFeeDestination
	}
	return ""
}

func (m *Params) GetTakerFeeDestination() string {
	if m != nil {
		return m.TakerFeeDestination
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "bze.tradebin.v1.Params")
}

func init() { proto.RegisterFile("tradebin/params.proto", fileDescriptor_db27696fb9409af7) }

var fileDescriptor_db27696fb9409af7 = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x29, 0x4a, 0x4c,
	0x49, 0x4d, 0xca, 0xcc, 0xd3, 0x2f, 0x48, 0x2c, 0x4a, 0xcc, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0x4f, 0xaa, 0x4a, 0xd5, 0x83, 0x49, 0xe9, 0x95, 0x19, 0x4a, 0x89, 0xa4, 0xe7,
	0xa7, 0xe7, 0x83, 0xe5, 0xf4, 0x41, 0x2c, 0x88, 0x32, 0xa5, 0x0e, 0x66, 0x2e, 0xb6, 0x00, 0xb0,
	0x3e, 0x21, 0x37, 0x2e, 0xfe, 0xe4, 0xa2, 0xd4, 0xc4, 0x92, 0x54, 0xdf, 0xc4, 0xa2, 0xec, 0xd4,
	0x12, 0xb7, 0xd4, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x4e, 0x27, 0x99, 0x4f, 0xf7, 0xe4, 0x25,
	0x2a, 0x13, 0x73, 0x73, 0xac, 0x94, 0x20, 0x0a, 0xe2, 0x73, 0xc1, 0x2a, 0xe2, 0xd3, 0x52, 0x53,
	0x95, 0x82, 0xd0, 0x35, 0x09, 0x39, 0x73, 0xf1, 0x41, 0xe4, 0x7d, 0x13, 0xb3, 0x53, 0x8b, 0x40,
	0xc6, 0x30, 0x81, 0x8d, 0x91, 0xfe, 0x74, 0x4f, 0x5e, 0x1c, 0x62, 0x0c, 0x54, 0x7f, 0x2e, 0x48,
	0x01, 0xc4, 0x14, 0x34, 0x2d, 0x08, 0x43, 0x42, 0x60, 0x86, 0x30, 0xe3, 0x30, 0xa4, 0x04, 0xc3,
	0x10, 0x98, 0x16, 0xa1, 0x20, 0x2e, 0xe1, 0x5c, 0x28, 0xdb, 0x25, 0xb5, 0xb8, 0x24, 0x33, 0x2f,
	0xb1, 0x24, 0x33, 0x3f, 0x4f, 0x82, 0x05, 0x6c, 0x92, 0xc2, 0xa7, 0x7b, 0xf2, 0x32, 0x30, 0x93,
	0xa0, 0x46, 0xc4, 0xa7, 0x20, 0x94, 0x29, 0x05, 0x61, 0xd3, 0x0c, 0x32, 0xb3, 0x04, 0x8b, 0x99,
	0xac, 0xe8, 0x66, 0x96, 0xe0, 0x30, 0x13, 0x8b, 0x66, 0x2b, 0x96, 0x19, 0x0b, 0xe4, 0x19, 0x9c,
	0xdc, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f,
	0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x37, 0x3d, 0xb3, 0x24,
	0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x3f, 0xa9, 0x2a, 0x55, 0x37, 0x31, 0xa7, 0x20, 0x23,
	0xb1, 0x24, 0x35, 0x11, 0xcc, 0xd3, 0xaf, 0xd0, 0x87, 0xa7, 0x80, 0x92, 0xca, 0x82, 0xd4, 0xe2,
	0x24, 0x36, 0x70, 0xd4, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x75, 0xbf, 0x0c, 0xdb, 0x1a,
	0x02, 0x00, 0x00,
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
	if len(m.TakerFeeDestination) > 0 {
		i -= len(m.TakerFeeDestination)
		copy(dAtA[i:], m.TakerFeeDestination)
		i = encodeVarintParams(dAtA, i, uint64(len(m.TakerFeeDestination)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.MakerFeeDestination) > 0 {
		i -= len(m.MakerFeeDestination)
		copy(dAtA[i:], m.MakerFeeDestination)
		i = encodeVarintParams(dAtA, i, uint64(len(m.MakerFeeDestination)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.MarketTakerFee) > 0 {
		i -= len(m.MarketTakerFee)
		copy(dAtA[i:], m.MarketTakerFee)
		i = encodeVarintParams(dAtA, i, uint64(len(m.MarketTakerFee)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.MarketMakerFee) > 0 {
		i -= len(m.MarketMakerFee)
		copy(dAtA[i:], m.MarketMakerFee)
		i = encodeVarintParams(dAtA, i, uint64(len(m.MarketMakerFee)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.CreateMarketFee) > 0 {
		i -= len(m.CreateMarketFee)
		copy(dAtA[i:], m.CreateMarketFee)
		i = encodeVarintParams(dAtA, i, uint64(len(m.CreateMarketFee)))
		i--
		dAtA[i] = 0xa
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
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CreateMarketFee)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.MarketMakerFee)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.MarketTakerFee)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.MakerFeeDestination)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.TakerFeeDestination)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateMarketFee", wireType)
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
			m.CreateMarketFee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarketMakerFee", wireType)
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
			m.MarketMakerFee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarketTakerFee", wireType)
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
			m.MarketTakerFee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MakerFeeDestination", wireType)
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
			m.MakerFeeDestination = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TakerFeeDestination", wireType)
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
			m.TakerFeeDestination = string(dAtA[iNdEx:postIndex])
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
