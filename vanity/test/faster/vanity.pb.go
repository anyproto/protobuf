// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: vanity.proto

package vanity

import (
	fmt "fmt"
	github_com_gogo_protobuf_proto "github.com/anyproto/protobuf/proto"
	proto "github.com/anyproto/protobuf/proto"
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

type A struct {
	Strings string `protobuf:"bytes,1,opt,name=Strings" json:"Strings"`
	Int     int64  `protobuf:"varint,2,req,name=Int" json:"Int"`
}

func (m *A) Reset()         { *m = A{} }
func (m *A) String() string { return proto.CompactTextString(m) }
func (*A) ProtoMessage()    {}
func (*A) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4f40d14cd1329d6, []int{0}
}
func (m *A) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *A) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_A.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *A) XXX_Merge(src proto.Message) {
	xxx_messageInfo_A.Merge(m, src)
}
func (m *A) XXX_Size() int {
	return m.Size()
}
func (m *A) XXX_DiscardUnknown() {
	xxx_messageInfo_A.DiscardUnknown(m)
}

var xxx_messageInfo_A proto.InternalMessageInfo

func (m *A) GetStrings() string {
	if m != nil {
		return m.Strings
	}
	return ""
}

func (m *A) GetInt() int64 {
	if m != nil {
		return m.Int
	}
	return 0
}

func init() {
	proto.RegisterType((*A)(nil), "vanity.A")
}

func init() { proto.RegisterFile("vanity.proto", fileDescriptor_d4f40d14cd1329d6) }

var fileDescriptor_d4f40d14cd1329d6 = []byte{
	// 117 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4b, 0xcc, 0xcb,
	0x2c, 0xa9, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0xac, 0xb9, 0x18,
	0x1d, 0x85, 0xe4, 0xb8, 0xd8, 0x83, 0x4b, 0x8a, 0x32, 0xf3, 0xd2, 0x8b, 0x25, 0x18, 0x15, 0x18,
	0x35, 0x38, 0x9d, 0x58, 0x4e, 0xdc, 0x93, 0x67, 0x08, 0x82, 0x09, 0x0a, 0x89, 0x71, 0x31, 0x7b,
	0xe6, 0x95, 0x48, 0x30, 0x29, 0x30, 0x69, 0x30, 0x43, 0xe5, 0x40, 0x02, 0x4e, 0x12, 0x27, 0x1e,
	0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17,
	0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x00, 0x08, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x2e, 0x6a,
	0x0c, 0x6d, 0x00, 0x00, 0x00,
}

func (m *A) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *A) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *A) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i = encodeVarintVanity(dAtA, i, uint64(m.Int))
	i--
	dAtA[i] = 0x10
	i -= len(m.Strings)
	copy(dAtA[i:], m.Strings)
	i = encodeVarintVanity(dAtA, i, uint64(len(m.Strings)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintVanity(dAtA []byte, offset int, v uint64) int {
	offset -= sovVanity(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *A) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Strings)
	n += 1 + l + sovVanity(uint64(l))
	n += 1 + sovVanity(uint64(m.Int))
	return n
}

func sovVanity(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVanity(x uint64) (n int) {
	return sovVanity(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *A) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVanity
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
			return fmt.Errorf("proto: A: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: A: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Strings", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVanity
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
				return ErrInvalidLengthVanity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVanity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Strings = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Int", wireType)
			}
			m.Int = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVanity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Int |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000001)
		default:
			iNdEx = preIndex
			skippy, err := skipVanity(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVanity
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("Int")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipVanity(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVanity
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
					return 0, ErrIntOverflowVanity
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
					return 0, ErrIntOverflowVanity
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
				return 0, ErrInvalidLengthVanity
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVanity
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVanity
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVanity        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVanity          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVanity = fmt.Errorf("proto: unexpected end of group")
)
