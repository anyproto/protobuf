// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: imp/imp2/imp2-enum.proto

package imp2 // import "github.com/gogo/protobuf/protoc-gen-gogo/testdata/imp/imp2"

import proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type EnumPB_Type int32

const (
	EnumPB_TYPE_1 EnumPB_Type = 0
)

var EnumPB_Type_name = map[int32]string{
	0: "TYPE_1",
}
var EnumPB_Type_value = map[string]int32{
	"TYPE_1": 0,
}

func (x EnumPB_Type) Enum() *EnumPB_Type {
	p := new(EnumPB_Type)
	*p = x
	return p
}
func (x EnumPB_Type) String() string {
	return proto.EnumName(EnumPB_Type_name, int32(x))
}
func (x *EnumPB_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EnumPB_Type_value, data, "EnumPB_Type")
	if err != nil {
		return err
	}
	*x = EnumPB_Type(value)
	return nil
}
func (EnumPB_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_imp2_enum_55af38dfa913421d, []int{0, 0}
}

type EnumPB struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnumPB) Reset()         { *m = EnumPB{} }
func (m *EnumPB) String() string { return proto.CompactTextString(m) }
func (*EnumPB) ProtoMessage()    {}
func (*EnumPB) Descriptor() ([]byte, []int) {
	return fileDescriptor_imp2_enum_55af38dfa913421d, []int{0}
}
func (m *EnumPB) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnumPB.Unmarshal(m, b)
}
func (m *EnumPB) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnumPB.Marshal(b, m, deterministic)
}
func (dst *EnumPB) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnumPB.Merge(dst, src)
}
func (m *EnumPB) XXX_Size() int {
	return xxx_messageInfo_EnumPB.Size(m)
}
func (m *EnumPB) XXX_DiscardUnknown() {
	xxx_messageInfo_EnumPB.DiscardUnknown(m)
}

var xxx_messageInfo_EnumPB proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EnumPB)(nil), "imp.EnumPB")
	proto.RegisterEnum("imp.EnumPB_Type", EnumPB_Type_name, EnumPB_Type_value)
}

func init() { proto.RegisterFile("imp/imp2/imp2-enum.proto", fileDescriptor_imp2_enum_55af38dfa913421d) }

var fileDescriptor_imp2_enum_55af38dfa913421d = []byte{
	// 127 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xc8, 0xcc, 0x2d, 0xd0,
	0xcf, 0xcc, 0x2d, 0x30, 0x02, 0x13, 0xba, 0xa9, 0x79, 0xa5, 0xb9, 0x7a, 0x05, 0x45, 0xf9, 0x25,
	0xf9, 0x42, 0xcc, 0x99, 0xb9, 0x05, 0x4a, 0x32, 0x5c, 0x6c, 0xae, 0x79, 0xa5, 0xb9, 0x01, 0x4e,
	0x4a, 0x42, 0x5c, 0x2c, 0x21, 0x95, 0x05, 0xa9, 0x42, 0x5c, 0x5c, 0x6c, 0x21, 0x91, 0x01, 0xae,
	0xf1, 0x86, 0x02, 0x0c, 0x4e, 0x36, 0x51, 0x56, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9,
	0xf9, 0xb9, 0xfa, 0xe9, 0xf9, 0xe9, 0xf9, 0xfa, 0x60, 0xbd, 0x49, 0xa5, 0x69, 0x10, 0x46, 0xb2,
	0x6e, 0x7a, 0x6a, 0x9e, 0x2e, 0x58, 0xa2, 0x24, 0xb5, 0xb8, 0x24, 0x25, 0xb1, 0x24, 0x51, 0x1f,
	0x66, 0x21, 0x20, 0x00, 0x00, 0xff, 0xff, 0xb1, 0xd5, 0xeb, 0xd2, 0x7b, 0x00, 0x00, 0x00,
}