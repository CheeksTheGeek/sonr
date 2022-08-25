// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: common/v1/ipns.proto

package common

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

type IpnsEntry_ValidityType int32

const (
	// setting an EOL says "this record is valid until..."
	IpnsEntry_EOL IpnsEntry_ValidityType = 0
)

var IpnsEntry_ValidityType_name = map[int32]string{
	0: "EOL",
}

var IpnsEntry_ValidityType_value = map[string]int32{
	"EOL": 0,
}

func (x IpnsEntry_ValidityType) String() string {
	return proto.EnumName(IpnsEntry_ValidityType_name, int32(x))
}

func (IpnsEntry_ValidityType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5738f9ecddc69498, []int{0, 0}
}

type IpnsEntry struct {
	Value        []byte                 `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Signature    []byte                 `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	ValidityType IpnsEntry_ValidityType `protobuf:"varint,3,opt,name=validityType,proto3,enum=sonrio.common.v1.IpnsEntry_ValidityType" json:"validityType,omitempty"`
	Validity     []byte                 `protobuf:"bytes,4,opt,name=validity,proto3" json:"validity,omitempty"`
	Sequence     uint64                 `protobuf:"varint,5,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Ttl          uint64                 `protobuf:"varint,6,opt,name=ttl,proto3" json:"ttl,omitempty"`
	PubKey       []byte                 `protobuf:"bytes,7,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
}

func (m *IpnsEntry) Reset()         { *m = IpnsEntry{} }
func (m *IpnsEntry) String() string { return proto.CompactTextString(m) }
func (*IpnsEntry) ProtoMessage()    {}
func (*IpnsEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_5738f9ecddc69498, []int{0}
}
func (m *IpnsEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IpnsEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IpnsEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IpnsEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpnsEntry.Merge(m, src)
}
func (m *IpnsEntry) XXX_Size() int {
	return m.Size()
}
func (m *IpnsEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_IpnsEntry.DiscardUnknown(m)
}

var xxx_messageInfo_IpnsEntry proto.InternalMessageInfo

func (m *IpnsEntry) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *IpnsEntry) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *IpnsEntry) GetValidityType() IpnsEntry_ValidityType {
	if m != nil {
		return m.ValidityType
	}
	return IpnsEntry_EOL
}

func (m *IpnsEntry) GetValidity() []byte {
	if m != nil {
		return m.Validity
	}
	return nil
}

func (m *IpnsEntry) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *IpnsEntry) GetTtl() uint64 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

func (m *IpnsEntry) GetPubKey() []byte {
	if m != nil {
		return m.PubKey
	}
	return nil
}

func init() {
	proto.RegisterEnum("sonrio.common.v1.IpnsEntry_ValidityType", IpnsEntry_ValidityType_name, IpnsEntry_ValidityType_value)
	proto.RegisterType((*IpnsEntry)(nil), "sonrio.common.v1.IpnsEntry")
}

func init() { proto.RegisterFile("common/v1/ipns.proto", fileDescriptor_5738f9ecddc69498) }

var fileDescriptor_5738f9ecddc69498 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x4f, 0x4a, 0x03, 0x31,
	0x18, 0xc5, 0x27, 0xfd, 0x6b, 0x43, 0x91, 0x12, 0x8a, 0x06, 0x91, 0x50, 0xba, 0x9a, 0x8d, 0x89,
	0xd5, 0x1b, 0x08, 0x5d, 0x88, 0x05, 0xa5, 0x88, 0x0b, 0x37, 0x32, 0xd3, 0x86, 0x36, 0xd0, 0x26,
	0x31, 0xc9, 0x0c, 0xe4, 0x06, 0x2e, 0x3d, 0x96, 0xcb, 0x2e, 0x5d, 0x4a, 0x7b, 0x11, 0x99, 0x4c,
	0x1d, 0xab, 0xab, 0xe4, 0xf7, 0x5e, 0x5e, 0x1e, 0x1f, 0x1f, 0xec, 0xcf, 0xd4, 0x7a, 0xad, 0x24,
	0xcb, 0x47, 0x4c, 0x68, 0x69, 0xa9, 0x36, 0xca, 0x29, 0xd4, 0xb3, 0x4a, 0x1a, 0xa1, 0x68, 0x69,
	0xd2, 0x7c, 0x34, 0x7c, 0xab, 0xc1, 0xce, 0xad, 0x96, 0x76, 0x2c, 0x9d, 0xf1, 0xa8, 0x0f, 0x9b,
	0x79, 0xb2, 0xca, 0x38, 0x06, 0x03, 0x10, 0x77, 0xa7, 0x25, 0xa0, 0x73, 0xd8, 0xb1, 0x62, 0x21,
	0x13, 0x97, 0x19, 0x8e, 0x6b, 0xc1, 0xf9, 0x15, 0xd0, 0x04, 0x76, 0xf3, 0x64, 0x25, 0xe6, 0xc2,
	0xf9, 0x47, 0xaf, 0x39, 0xae, 0x0f, 0x40, 0x7c, 0x7c, 0x15, 0xd3, 0xff, 0x55, 0xb4, 0xaa, 0xa1,
	0x4f, 0x07, 0xef, 0xa7, 0x7f, 0xd2, 0xe8, 0x0c, 0x1e, 0xfd, 0x30, 0x6e, 0x84, 0xaa, 0x8a, 0x0b,
	0xcf, 0xf2, 0xd7, 0x8c, 0xcb, 0x19, 0xc7, 0xcd, 0x01, 0x88, 0x1b, 0xd3, 0x8a, 0x51, 0x0f, 0xd6,
	0x9d, 0x5b, 0xe1, 0x56, 0x90, 0x8b, 0x2b, 0x3a, 0x81, 0x2d, 0x9d, 0xa5, 0x77, 0xdc, 0xe3, 0x76,
	0xf8, 0x67, 0x4f, 0xc3, 0x53, 0xd8, 0x3d, 0xec, 0x47, 0x6d, 0x58, 0x1f, 0xdf, 0x4f, 0x7a, 0xd1,
	0x4d, 0xfa, 0xb1, 0x25, 0x60, 0xb3, 0x25, 0xe0, 0x6b, 0x4b, 0xc0, 0xfb, 0x8e, 0x44, 0x9b, 0x1d,
	0x89, 0x3e, 0x77, 0x24, 0x82, 0x7d, 0xa1, 0xc2, 0x38, 0xd4, 0x79, 0xcd, 0xed, 0x7e, 0xa4, 0x07,
	0xf0, 0x7c, 0xb9, 0x10, 0x6e, 0x99, 0xa5, 0x85, 0xc0, 0x0a, 0xff, 0x42, 0xa8, 0x70, 0x32, 0xb7,
	0x14, 0x66, 0xfe, 0xa2, 0x13, 0xe3, 0x3c, 0x0b, 0x19, 0x56, 0x66, 0xd2, 0x56, 0xd8, 0xc3, 0xf5,
	0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x28, 0x4a, 0xe1, 0x55, 0x9f, 0x01, 0x00, 0x00,
}

func (m *IpnsEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IpnsEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IpnsEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PubKey) > 0 {
		i -= len(m.PubKey)
		copy(dAtA[i:], m.PubKey)
		i = encodeVarintIpns(dAtA, i, uint64(len(m.PubKey)))
		i--
		dAtA[i] = 0x3a
	}
	if m.Ttl != 0 {
		i = encodeVarintIpns(dAtA, i, uint64(m.Ttl))
		i--
		dAtA[i] = 0x30
	}
	if m.Sequence != 0 {
		i = encodeVarintIpns(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Validity) > 0 {
		i -= len(m.Validity)
		copy(dAtA[i:], m.Validity)
		i = encodeVarintIpns(dAtA, i, uint64(len(m.Validity)))
		i--
		dAtA[i] = 0x22
	}
	if m.ValidityType != 0 {
		i = encodeVarintIpns(dAtA, i, uint64(m.ValidityType))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintIpns(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintIpns(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintIpns(dAtA []byte, offset int, v uint64) int {
	offset -= sovIpns(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IpnsEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovIpns(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovIpns(uint64(l))
	}
	if m.ValidityType != 0 {
		n += 1 + sovIpns(uint64(m.ValidityType))
	}
	l = len(m.Validity)
	if l > 0 {
		n += 1 + l + sovIpns(uint64(l))
	}
	if m.Sequence != 0 {
		n += 1 + sovIpns(uint64(m.Sequence))
	}
	if m.Ttl != 0 {
		n += 1 + sovIpns(uint64(m.Ttl))
	}
	l = len(m.PubKey)
	if l > 0 {
		n += 1 + l + sovIpns(uint64(l))
	}
	return n
}

func sovIpns(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIpns(x uint64) (n int) {
	return sovIpns(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IpnsEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIpns
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
			return fmt.Errorf("proto: IpnsEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IpnsEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIpns
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthIpns
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIpns
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIpns
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthIpns
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIpns
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidityType", wireType)
			}
			m.ValidityType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIpns
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ValidityType |= IpnsEntry_ValidityType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validity", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIpns
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthIpns
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIpns
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validity = append(m.Validity[:0], dAtA[iNdEx:postIndex]...)
			if m.Validity == nil {
				m.Validity = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIpns
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ttl", wireType)
			}
			m.Ttl = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIpns
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ttl |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIpns
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthIpns
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIpns
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKey = append(m.PubKey[:0], dAtA[iNdEx:postIndex]...)
			if m.PubKey == nil {
				m.PubKey = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIpns(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIpns
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
func skipIpns(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIpns
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
					return 0, ErrIntOverflowIpns
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
					return 0, ErrIntOverflowIpns
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
				return 0, ErrInvalidLengthIpns
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIpns
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIpns
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIpns        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIpns          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIpns = fmt.Errorf("proto: unexpected end of group")
)
