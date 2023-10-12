// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fusionchain/treasury/wallet.proto

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

// WalletType specifies the Layer 1 blockchain that this wallet will be used
// for.
type WalletType int32

const (
	// The wallet type is missing
	WalletType_WALLET_TYPE_UNSPECIFIED WalletType = 0
	// The wallet type for native Fusion chain cosmos accounts (not ERC-20 QRDO
	// tokens)
	WalletType_WALLET_TYPE_QRDO WalletType = 1
	// The wallet type for mainnet ETH and its ERC-20 tokens (including non-native
	// QRDO)
	WalletType_WALLET_TYPE_ETH WalletType = 2
	// The wallet type for Sepolia testnet ETH and its ERC-20 tokens
	WalletType_WALLET_TYPE_ETH_SEPOLIA WalletType = 3
)

var WalletType_name = map[int32]string{
	0: "WALLET_TYPE_UNSPECIFIED",
	1: "WALLET_TYPE_QRDO",
	2: "WALLET_TYPE_ETH",
	3: "WALLET_TYPE_ETH_SEPOLIA",
}

var WalletType_value = map[string]int32{
	"WALLET_TYPE_UNSPECIFIED": 0,
	"WALLET_TYPE_QRDO":        1,
	"WALLET_TYPE_ETH":         2,
	"WALLET_TYPE_ETH_SEPOLIA": 3,
}

func (x WalletType) String() string {
	return proto.EnumName(WalletType_name, int32(x))
}

func (WalletType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_51fb94234f9ffc53, []int{0}
}

// WalletRequestType used at the request level for query_keys
type WalletRequestType int32

const (
	// The wallet type is missing
	WalletRequestType_WALLET_REQUEST_TYPE_UNSPECIFIED WalletRequestType = 0
	// The wallet type for all wallets to be derived
	WalletRequestType_WALLET_REQUEST_TYPE_ALL WalletRequestType = 1
	// The wallet type for native Fusion chain cosmos accounts (not ERC-20 QRDO
	// tokens)
	WalletRequestType_WALLET_REQUEST_TYPE_QRDO WalletRequestType = 2
	// The wallet type for mainnet ETH and its ERC-20 tokens (including non-native
	// QRDO)
	WalletRequestType_WALLET_REQUEST_TYPE_ETH WalletRequestType = 3
	// The wallet type for Sepolia testnet ETH and its ERC-20 tokens
	WalletRequestType_WALLET_REQUEST_TYPE_ETH_SEPOLIA WalletRequestType = 4
)

var WalletRequestType_name = map[int32]string{
	0: "WALLET_REQUEST_TYPE_UNSPECIFIED",
	1: "WALLET_REQUEST_TYPE_ALL",
	2: "WALLET_REQUEST_TYPE_QRDO",
	3: "WALLET_REQUEST_TYPE_ETH",
	4: "WALLET_REQUEST_TYPE_ETH_SEPOLIA",
}

var WalletRequestType_value = map[string]int32{
	"WALLET_REQUEST_TYPE_UNSPECIFIED": 0,
	"WALLET_REQUEST_TYPE_ALL":         1,
	"WALLET_REQUEST_TYPE_QRDO":        2,
	"WALLET_REQUEST_TYPE_ETH":         3,
	"WALLET_REQUEST_TYPE_ETH_SEPOLIA": 4,
}

func (x WalletRequestType) String() string {
	return proto.EnumName(WalletRequestType_name, int32(x))
}

func (WalletRequestType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_51fb94234f9ffc53, []int{1}
}

type Wallet struct {
	Id    uint64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type  WalletType `protobuf:"varint,2,opt,name=type,proto3,enum=fusionchain.treasury.WalletType" json:"type,omitempty"`
	KeyId uint64     `protobuf:"varint,3,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
}

func (m *Wallet) Reset()         { *m = Wallet{} }
func (m *Wallet) String() string { return proto.CompactTextString(m) }
func (*Wallet) ProtoMessage()    {}
func (*Wallet) Descriptor() ([]byte, []int) {
	return fileDescriptor_51fb94234f9ffc53, []int{0}
}
func (m *Wallet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Wallet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Wallet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Wallet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Wallet.Merge(m, src)
}
func (m *Wallet) XXX_Size() int {
	return m.Size()
}
func (m *Wallet) XXX_DiscardUnknown() {
	xxx_messageInfo_Wallet.DiscardUnknown(m)
}

var xxx_messageInfo_Wallet proto.InternalMessageInfo

func (m *Wallet) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Wallet) GetType() WalletType {
	if m != nil {
		return m.Type
	}
	return WalletType_WALLET_TYPE_UNSPECIFIED
}

func (m *Wallet) GetKeyId() uint64 {
	if m != nil {
		return m.KeyId
	}
	return 0
}

func init() {
	proto.RegisterEnum("fusionchain.treasury.WalletType", WalletType_name, WalletType_value)
	proto.RegisterEnum("fusionchain.treasury.WalletRequestType", WalletRequestType_name, WalletRequestType_value)
	proto.RegisterType((*Wallet)(nil), "fusionchain.treasury.Wallet")
}

func init() { proto.RegisterFile("fusionchain/treasury/wallet.proto", fileDescriptor_51fb94234f9ffc53) }

var fileDescriptor_51fb94234f9ffc53 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x4d, 0x4e, 0xc2, 0x40,
	0x14, 0xc7, 0x3b, 0x05, 0x59, 0xbc, 0x05, 0x8e, 0x23, 0xc6, 0x26, 0x9a, 0x11, 0x75, 0x43, 0x48,
	0x6c, 0x13, 0xf5, 0x02, 0x28, 0xa3, 0x34, 0x69, 0x04, 0x4a, 0x09, 0xd1, 0x4d, 0x03, 0x74, 0x94,
	0x09, 0x48, 0xa1, 0x1f, 0xd1, 0xde, 0xc2, 0xcb, 0x78, 0x07, 0x97, 0x2c, 0x5d, 0x1a, 0xb8, 0x88,
	0xa1, 0x68, 0xa8, 0xa6, 0x6c, 0xdf, 0xfb, 0xfd, 0x3f, 0x92, 0x3f, 0x1c, 0x3f, 0x86, 0xbe, 0x70,
	0xc7, 0xfd, 0x41, 0x57, 0x8c, 0xb5, 0xc0, 0xe3, 0x5d, 0x3f, 0xf4, 0x22, 0xed, 0xa5, 0x3b, 0x1a,
	0xf1, 0x40, 0x9d, 0x78, 0x6e, 0xe0, 0x92, 0x42, 0x02, 0x51, 0x7f, 0x91, 0x13, 0x0e, 0xb9, 0x4e,
	0x4c, 0x91, 0x3c, 0xc8, 0xc2, 0x51, 0x50, 0x11, 0x95, 0xb2, 0xa6, 0x2c, 0x1c, 0x72, 0x09, 0xd9,
	0x20, 0x9a, 0x70, 0x45, 0x2e, 0xa2, 0x52, 0xfe, 0xbc, 0xa8, 0xa6, 0xc9, 0xd5, 0x95, 0xd6, 0x8a,
	0x26, 0xdc, 0x8c, 0x69, 0xb2, 0x07, 0xb9, 0x21, 0x8f, 0x6c, 0xe1, 0x28, 0x99, 0xd8, 0x69, 0x6b,
	0xc8, 0x23, 0xdd, 0x29, 0x4f, 0x01, 0xd6, 0x28, 0x39, 0x80, 0xfd, 0x4e, 0xc5, 0x30, 0x98, 0x65,
	0x5b, 0xf7, 0x0d, 0x66, 0xb7, 0xef, 0x5a, 0x0d, 0x76, 0xad, 0xdf, 0xe8, 0xac, 0x8a, 0x25, 0x52,
	0x00, 0x9c, 0x7c, 0x36, 0xcd, 0x6a, 0x1d, 0x23, 0xb2, 0x0b, 0xdb, 0xc9, 0x2b, 0xb3, 0x6a, 0x58,
	0xfe, 0xef, 0xc3, 0xac, 0x9a, 0xdd, 0x62, 0x8d, 0xba, 0xa1, 0x57, 0x70, 0xa6, 0xfc, 0x8e, 0x60,
	0x67, 0x95, 0x69, 0xf2, 0x69, 0xc8, 0xfd, 0x55, 0xf4, 0x29, 0x1c, 0xfd, 0x48, 0x4c, 0xd6, 0x6c,
	0xb3, 0x56, 0x6a, 0x85, 0xb5, 0xef, 0x1f, 0xa8, 0x62, 0x18, 0x18, 0x91, 0x43, 0x50, 0xd2, 0x9e,
	0x71, 0x4f, 0x79, 0x93, 0x74, 0xd9, 0x37, 0xb3, 0x29, 0x3c, 0xd9, 0x3b, 0x7b, 0x75, 0xfb, 0x31,
	0xa7, 0x68, 0x36, 0xa7, 0xe8, 0x6b, 0x4e, 0xd1, 0xdb, 0x82, 0x4a, 0xb3, 0x05, 0x95, 0x3e, 0x17,
	0x54, 0x7a, 0x38, 0x7b, 0x12, 0xc1, 0x20, 0xec, 0xa9, 0x7d, 0xf7, 0x59, 0x9b, 0x7a, 0xdc, 0x71,
	0xb5, 0xe4, 0xea, 0xaf, 0xeb, 0xdd, 0x97, 0x4b, 0xf8, 0xbd, 0x5c, 0xbc, 0xfb, 0xc5, 0x77, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x61, 0xd2, 0x20, 0x7a, 0x1c, 0x02, 0x00, 0x00,
}

func (m *Wallet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Wallet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Wallet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.KeyId != 0 {
		i = encodeVarintWallet(dAtA, i, uint64(m.KeyId))
		i--
		dAtA[i] = 0x18
	}
	if m.Type != 0 {
		i = encodeVarintWallet(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintWallet(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintWallet(dAtA []byte, offset int, v uint64) int {
	offset -= sovWallet(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Wallet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovWallet(uint64(m.Id))
	}
	if m.Type != 0 {
		n += 1 + sovWallet(uint64(m.Type))
	}
	if m.KeyId != 0 {
		n += 1 + sovWallet(uint64(m.KeyId))
	}
	return n
}

func sovWallet(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWallet(x uint64) (n int) {
	return sovWallet(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Wallet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWallet
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
			return fmt.Errorf("proto: Wallet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Wallet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWallet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWallet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= WalletType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyId", wireType)
			}
			m.KeyId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWallet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.KeyId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipWallet(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWallet
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
func skipWallet(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWallet
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
					return 0, ErrIntOverflowWallet
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
					return 0, ErrIntOverflowWallet
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
				return 0, ErrInvalidLengthWallet
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWallet
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWallet
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWallet        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWallet          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWallet = fmt.Errorf("proto: unexpected end of group")
)
