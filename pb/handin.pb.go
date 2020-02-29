// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/handin.proto

package pb

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Handin struct {
	Handinid             int64    `protobuf:"varint,1,opt,name=handinid,proto3" json:"handinid,omitempty"`
	Time                 int64    `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	Questid              int64    `protobuf:"varint,3,opt,name=questid,proto3" json:"questid,omitempty"`
	Charid               int64    `protobuf:"varint,4,opt,name=charid,proto3" json:"charid,omitempty"`
	Charpp               int64    `protobuf:"varint,5,opt,name=charpp,proto3" json:"charpp,omitempty"`
	Chargp               int64    `protobuf:"varint,6,opt,name=chargp,proto3" json:"chargp,omitempty"`
	Charsp               int64    `protobuf:"varint,7,opt,name=charsp,proto3" json:"charsp,omitempty"`
	Charcp               int64    `protobuf:"varint,8,opt,name=charcp,proto3" json:"charcp,omitempty"`
	Charitems            int64    `protobuf:"varint,9,opt,name=charitems,proto3" json:"charitems,omitempty"`
	Npcid                int64    `protobuf:"varint,10,opt,name=npcid,proto3" json:"npcid,omitempty"`
	Npcpp                int64    `protobuf:"varint,11,opt,name=npcpp,proto3" json:"npcpp,omitempty"`
	Npcgp                int64    `protobuf:"varint,12,opt,name=npcgp,proto3" json:"npcgp,omitempty"`
	Npcsp                int64    `protobuf:"varint,13,opt,name=npcsp,proto3" json:"npcsp,omitempty"`
	Npccp                int64    `protobuf:"varint,14,opt,name=npccp,proto3" json:"npccp,omitempty"`
	Npcitems             int64    `protobuf:"varint,15,opt,name=npcitems,proto3" json:"npcitems,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Handin) Reset()         { *m = Handin{} }
func (m *Handin) String() string { return proto.CompactTextString(m) }
func (*Handin) ProtoMessage()    {}
func (*Handin) Descriptor() ([]byte, []int) {
	return fileDescriptor_06508675011c01fd, []int{0}
}

func (m *Handin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Handin.Unmarshal(m, b)
}
func (m *Handin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Handin.Marshal(b, m, deterministic)
}
func (m *Handin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Handin.Merge(m, src)
}
func (m *Handin) XXX_Size() int {
	return xxx_messageInfo_Handin.Size(m)
}
func (m *Handin) XXX_DiscardUnknown() {
	xxx_messageInfo_Handin.DiscardUnknown(m)
}

var xxx_messageInfo_Handin proto.InternalMessageInfo

func (m *Handin) GetHandinid() int64 {
	if m != nil {
		return m.Handinid
	}
	return 0
}

func (m *Handin) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Handin) GetQuestid() int64 {
	if m != nil {
		return m.Questid
	}
	return 0
}

func (m *Handin) GetCharid() int64 {
	if m != nil {
		return m.Charid
	}
	return 0
}

func (m *Handin) GetCharpp() int64 {
	if m != nil {
		return m.Charpp
	}
	return 0
}

func (m *Handin) GetChargp() int64 {
	if m != nil {
		return m.Chargp
	}
	return 0
}

func (m *Handin) GetCharsp() int64 {
	if m != nil {
		return m.Charsp
	}
	return 0
}

func (m *Handin) GetCharcp() int64 {
	if m != nil {
		return m.Charcp
	}
	return 0
}

func (m *Handin) GetCharitems() int64 {
	if m != nil {
		return m.Charitems
	}
	return 0
}

func (m *Handin) GetNpcid() int64 {
	if m != nil {
		return m.Npcid
	}
	return 0
}

func (m *Handin) GetNpcpp() int64 {
	if m != nil {
		return m.Npcpp
	}
	return 0
}

func (m *Handin) GetNpcgp() int64 {
	if m != nil {
		return m.Npcgp
	}
	return 0
}

func (m *Handin) GetNpcsp() int64 {
	if m != nil {
		return m.Npcsp
	}
	return 0
}

func (m *Handin) GetNpccp() int64 {
	if m != nil {
		return m.Npccp
	}
	return 0
}

func (m *Handin) GetNpcitems() int64 {
	if m != nil {
		return m.Npcitems
	}
	return 0
}

func init() {
	proto.RegisterType((*Handin)(nil), "pb.Handin")
}

func init() { proto.RegisterFile("proto/handin.proto", fileDescriptor_06508675011c01fd) }

var fileDescriptor_06508675011c01fd = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0xd0, 0x4d, 0x6e, 0xc2, 0x30,
	0x10, 0x05, 0x60, 0x11, 0x20, 0x80, 0xfb, 0x27, 0x8d, 0xaa, 0x6a, 0x84, 0x58, 0x54, 0x5d, 0x75,
	0xd5, 0x2c, 0x7a, 0x89, 0xae, 0x7b, 0x03, 0x63, 0x47, 0xc6, 0x52, 0xb1, 0x5f, 0xb1, 0x7b, 0xf1,
	0x9e, 0x00, 0x31, 0x89, 0xed, 0x5d, 0xde, 0x37, 0x19, 0x8d, 0xf5, 0x14, 0xe1, 0x12, 0x73, 0x1c,
	0x4e, 0x3a, 0x58, 0x1f, 0x3e, 0x24, 0x50, 0x87, 0xe3, 0xfe, 0xe0, 0x62, 0x74, 0x3f, 0xe3, 0xa0,
	0xe1, 0x07, 0x1d, 0x42, 0xcc, 0x3a, 0xfb, 0x18, 0xd2, 0xf4, 0xc7, 0xdb, 0x7f, 0xa7, 0xfa, 0x2f,
	0x59, 0xa1, 0xbd, 0xda, 0x4e, 0xcb, 0xde, 0xf2, 0xe2, 0x75, 0xf1, 0xbe, 0xfc, 0xae, 0x99, 0x48,
	0xad, 0xb2, 0x3f, 0x8f, 0xdc, 0x89, 0xcb, 0x37, 0xb1, 0xda, 0xfc, 0xfe, 0x8d, 0x29, 0x7b, 0xcb,
	0x4b, 0xe1, 0x12, 0xe9, 0x45, 0xf5, 0xe6, 0xa4, 0x2f, 0xde, 0xf2, 0x4a, 0x06, 0x73, 0x2a, 0x0e,
	0xf0, 0xba, 0x39, 0x50, 0xdc, 0x81, 0xfb, 0xe6, 0xae, 0x7a, 0x02, 0x6f, 0x9a, 0xa7, 0xea, 0x06,
	0xbc, 0x6d, 0x6e, 0x40, 0x07, 0xb5, 0x93, 0x4b, 0x79, 0x3c, 0x27, 0xde, 0xc9, 0xa8, 0x01, 0x3d,
	0xab, 0x75, 0x80, 0xf1, 0x96, 0x95, 0x4c, 0xa6, 0x30, 0x2b, 0xc0, 0x77, 0x55, 0x81, 0x59, 0x1d,
	0xf8, 0xbe, 0xaa, 0x2b, 0x9a, 0xc0, 0x0f, 0x55, 0x53, 0x51, 0x03, 0x7e, 0xac, 0x6a, 0x70, 0x6b,
	0xf3, 0x76, 0x40, 0x9e, 0xf2, 0x34, 0xb5, 0x59, 0xf2, 0xb1, 0x97, 0xee, 0x3f, 0xaf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xb9, 0x0c, 0x1e, 0x3a, 0xb3, 0x01, 0x00, 0x00,
}