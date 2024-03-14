// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rewards/genesis.proto

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

// GenesisState defines the rewards module's genesis state.
type GenesisState struct {
	Params                Params          `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	StakingRewardList     []StakingReward `protobuf:"bytes,2,rep,name=staking_reward_list,json=stakingRewardList,proto3" json:"staking_reward_list"`
	StakingRewardsCounter uint64          `protobuf:"varint,3,opt,name=staking_rewards_counter,json=stakingRewardsCounter,proto3" json:"staking_rewards_counter,omitempty"`
	TradingRewardsCounter uint64          `protobuf:"varint,4,opt,name=trading_rewards_counter,json=tradingRewardsCounter,proto3" json:"trading_rewards_counter,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_04e261dd19d07881, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetStakingRewardList() []StakingReward {
	if m != nil {
		return m.StakingRewardList
	}
	return nil
}

func (m *GenesisState) GetStakingRewardsCounter() uint64 {
	if m != nil {
		return m.StakingRewardsCounter
	}
	return 0
}

func (m *GenesisState) GetTradingRewardsCounter() uint64 {
	if m != nil {
		return m.TradingRewardsCounter
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "bze.v1.rewards.GenesisState")
}

func init() { proto.RegisterFile("rewards/genesis.proto", fileDescriptor_04e261dd19d07881) }

var fileDescriptor_04e261dd19d07881 = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x4a, 0x2d, 0x4f,
	0x2c, 0x4a, 0x29, 0xd6, 0x4f, 0x4f, 0xcd, 0x4b, 0x2d, 0xce, 0x2c, 0xd6, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0x4b, 0xaa, 0x4a, 0xd5, 0x2b, 0x33, 0xd4, 0x83, 0xca, 0x4a, 0x89, 0xa4, 0xe7,
	0xa7, 0xe7, 0x83, 0xa5, 0xf4, 0x41, 0x2c, 0x88, 0x2a, 0x29, 0x11, 0x98, 0xe6, 0x82, 0xc4, 0xa2,
	0xc4, 0x5c, 0xa8, 0x5e, 0x29, 0x19, 0x98, 0x68, 0x71, 0x49, 0x62, 0x76, 0x66, 0x5e, 0x7a, 0x3c,
	0x84, 0x0f, 0x91, 0x55, 0xea, 0x64, 0xe2, 0xe2, 0x71, 0x87, 0xd8, 0x15, 0x5c, 0x92, 0x58, 0x92,
	0x2a, 0x64, 0xc2, 0xc5, 0x06, 0xd1, 0x2e, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x6d, 0x24, 0xa6, 0x87,
	0x6a, 0xb7, 0x5e, 0x00, 0x58, 0xd6, 0x89, 0xe5, 0xc4, 0x3d, 0x79, 0x86, 0x20, 0xa8, 0x5a, 0xa1,
	0x60, 0x2e, 0x61, 0x54, 0xe3, 0xe3, 0x73, 0x32, 0x8b, 0x4b, 0x24, 0x98, 0x14, 0x98, 0x35, 0xb8,
	0x8d, 0x64, 0xd1, 0x8d, 0x08, 0x86, 0x28, 0x0d, 0x02, 0x73, 0xa1, 0x26, 0x09, 0x16, 0x23, 0x0b,
	0xfa, 0x64, 0x16, 0x97, 0x08, 0x99, 0x71, 0x89, 0xa3, 0x1a, 0x5a, 0x1c, 0x9f, 0x9c, 0x5f, 0x9a,
	0x57, 0x92, 0x5a, 0x24, 0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0x12, 0x24, 0x8a, 0xa2, 0xa7, 0xd8, 0x19,
	0x22, 0x09, 0xd2, 0x57, 0x52, 0x94, 0x98, 0x82, 0x4d, 0x1f, 0x0b, 0x44, 0x1f, 0x54, 0x1a, 0x55,
	0x9f, 0x93, 0xdb, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38,
	0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0xe9, 0xa4, 0x67,
	0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x27, 0x55, 0xa5, 0xea, 0x26, 0xe6, 0x14,
	0x64, 0x24, 0x96, 0xa4, 0x26, 0x82, 0x79, 0xfa, 0x15, 0xfa, 0xb0, 0x20, 0x2e, 0xa9, 0x2c, 0x48,
	0x2d, 0x4e, 0x62, 0x03, 0x07, 0xad, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xc3, 0xb1, 0x6e, 0x4a,
	0xcd, 0x01, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TradingRewardsCounter != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.TradingRewardsCounter))
		i--
		dAtA[i] = 0x20
	}
	if m.StakingRewardsCounter != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.StakingRewardsCounter))
		i--
		dAtA[i] = 0x18
	}
	if len(m.StakingRewardList) > 0 {
		for iNdEx := len(m.StakingRewardList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.StakingRewardList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.StakingRewardList) > 0 {
		for _, e := range m.StakingRewardList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.StakingRewardsCounter != 0 {
		n += 1 + sovGenesis(uint64(m.StakingRewardsCounter))
	}
	if m.TradingRewardsCounter != 0 {
		n += 1 + sovGenesis(uint64(m.TradingRewardsCounter))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakingRewardList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StakingRewardList = append(m.StakingRewardList, StakingReward{})
			if err := m.StakingRewardList[len(m.StakingRewardList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakingRewardsCounter", wireType)
			}
			m.StakingRewardsCounter = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StakingRewardsCounter |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradingRewardsCounter", wireType)
			}
			m.TradingRewardsCounter = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TradingRewardsCounter |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
