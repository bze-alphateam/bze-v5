// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cointrunk/accepted_domain_proposal.proto

package types

import (
	fmt "fmt"
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

type AcceptedDomainProposal struct {
	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Domain      string `protobuf:"bytes,3,opt,name=domain,proto3" json:"domain,omitempty"`
	Active      bool   `protobuf:"varint,4,opt,name=active,proto3" json:"active,omitempty"`
}

func (m *AcceptedDomainProposal) Reset()         { *m = AcceptedDomainProposal{} }
func (m *AcceptedDomainProposal) String() string { return proto.CompactTextString(m) }
func (*AcceptedDomainProposal) ProtoMessage()    {}
func (*AcceptedDomainProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b0b7ddc6f58933e, []int{0}
}
func (m *AcceptedDomainProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AcceptedDomainProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AcceptedDomainProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AcceptedDomainProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AcceptedDomainProposal.Merge(m, src)
}
func (m *AcceptedDomainProposal) XXX_Size() int {
	return m.Size()
}
func (m *AcceptedDomainProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_AcceptedDomainProposal.DiscardUnknown(m)
}

var xxx_messageInfo_AcceptedDomainProposal proto.InternalMessageInfo

func (m *AcceptedDomainProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *AcceptedDomainProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *AcceptedDomainProposal) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *AcceptedDomainProposal) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func init() {
	proto.RegisterType((*AcceptedDomainProposal)(nil), "bze.cointrunk.v1.AcceptedDomainProposal")
}

func init() {
	proto.RegisterFile("cointrunk/accepted_domain_proposal.proto", fileDescriptor_5b0b7ddc6f58933e)
}

var fileDescriptor_5b0b7ddc6f58933e = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x48, 0xce, 0xcf, 0xcc,
	0x2b, 0x29, 0x2a, 0xcd, 0xcb, 0xd6, 0x4f, 0x4c, 0x4e, 0x4e, 0x2d, 0x28, 0x49, 0x4d, 0x89, 0x4f,
	0xc9, 0xcf, 0x4d, 0xcc, 0xcc, 0x8b, 0x2f, 0x28, 0xca, 0x2f, 0xc8, 0x2f, 0x4e, 0xcc, 0xd1, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x48, 0xaa, 0x4a, 0xd5, 0x83, 0xab, 0xd6, 0x2b, 0x33, 0x54,
	0x6a, 0x60, 0xe4, 0x12, 0x73, 0x84, 0x6a, 0x72, 0x01, 0xeb, 0x09, 0x80, 0x6a, 0x11, 0x12, 0xe1,
	0x62, 0x2d, 0xc9, 0x2c, 0xc9, 0x49, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x84,
	0x14, 0xb8, 0xb8, 0x53, 0x52, 0x8b, 0x93, 0x8b, 0x32, 0x0b, 0x4a, 0x32, 0xf3, 0xf3, 0x24, 0x98,
	0xc0, 0x72, 0xc8, 0x42, 0x42, 0x62, 0x5c, 0x6c, 0x10, 0xdb, 0x25, 0x98, 0xc1, 0x92, 0x50, 0x1e,
	0x48, 0x3c, 0x31, 0xb9, 0x24, 0xb3, 0x2c, 0x55, 0x82, 0x45, 0x81, 0x51, 0x83, 0x23, 0x08, 0xca,
	0x73, 0xf2, 0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27,
	0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xbd, 0xf4, 0xcc,
	0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0xa4, 0xaa, 0x54, 0xdd, 0xc4, 0x9c, 0x82,
	0x8c, 0xc4, 0x92, 0xd4, 0x44, 0x30, 0x4f, 0xbf, 0x42, 0x1f, 0xe1, 0xef, 0x92, 0xca, 0x82, 0xd4,
	0xe2, 0x24, 0x36, 0xb0, 0x2f, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb8, 0x7d, 0x99, 0x3d,
	0x11, 0x01, 0x00, 0x00,
}

func (m *AcceptedDomainProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AcceptedDomainProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AcceptedDomainProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Active {
		i--
		if m.Active {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.Domain) > 0 {
		i -= len(m.Domain)
		copy(dAtA[i:], m.Domain)
		i = encodeVarintAcceptedDomainProposal(dAtA, i, uint64(len(m.Domain)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintAcceptedDomainProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintAcceptedDomainProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAcceptedDomainProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovAcceptedDomainProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AcceptedDomainProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovAcceptedDomainProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovAcceptedDomainProposal(uint64(l))
	}
	l = len(m.Domain)
	if l > 0 {
		n += 1 + l + sovAcceptedDomainProposal(uint64(l))
	}
	if m.Active {
		n += 2
	}
	return n
}

func sovAcceptedDomainProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAcceptedDomainProposal(x uint64) (n int) {
	return sovAcceptedDomainProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AcceptedDomainProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAcceptedDomainProposal
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
			return fmt.Errorf("proto: AcceptedDomainProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AcceptedDomainProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcceptedDomainProposal
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
				return ErrInvalidLengthAcceptedDomainProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAcceptedDomainProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcceptedDomainProposal
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
				return ErrInvalidLengthAcceptedDomainProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAcceptedDomainProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Domain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcceptedDomainProposal
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
				return ErrInvalidLengthAcceptedDomainProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAcceptedDomainProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Domain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Active", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcceptedDomainProposal
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
			m.Active = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipAcceptedDomainProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAcceptedDomainProposal
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
func skipAcceptedDomainProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAcceptedDomainProposal
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
					return 0, ErrIntOverflowAcceptedDomainProposal
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
					return 0, ErrIntOverflowAcceptedDomainProposal
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
				return 0, ErrInvalidLengthAcceptedDomainProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAcceptedDomainProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAcceptedDomainProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAcceptedDomainProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAcceptedDomainProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAcceptedDomainProposal = fmt.Errorf("proto: unexpected end of group")
)
