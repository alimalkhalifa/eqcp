// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/login_server.proto

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

type Server struct {
	Localip              string   `protobuf:"bytes,1,opt,name=localip,proto3" json:"localip,omitempty"`
	Playersonline        int64    `protobuf:"varint,2,opt,name=playersonline,proto3" json:"playersonline,omitempty"`
	Remoteip             string   `protobuf:"bytes,3,opt,name=remoteip,proto3" json:"remoteip,omitempty"`
	Serverlistid         int64    `protobuf:"varint,4,opt,name=serverlistid,proto3" json:"serverlistid,omitempty"`
	Serverlongname       string   `protobuf:"bytes,5,opt,name=serverlongname,proto3" json:"serverlongname,omitempty"`
	Servershortname      string   `protobuf:"bytes,6,opt,name=servershortname,proto3" json:"servershortname,omitempty"`
	Serverstatus         int64    `protobuf:"varint,7,opt,name=serverstatus,proto3" json:"serverstatus,omitempty"`
	Zonesbooted          int64    `protobuf:"varint,8,opt,name=zonesbooted,proto3" json:"zonesbooted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Server) Reset()         { *m = Server{} }
func (m *Server) String() string { return proto.CompactTextString(m) }
func (*Server) ProtoMessage()    {}
func (*Server) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc26b5533683afe0, []int{0}
}

func (m *Server) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Server.Unmarshal(m, b)
}
func (m *Server) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Server.Marshal(b, m, deterministic)
}
func (m *Server) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Server.Merge(m, src)
}
func (m *Server) XXX_Size() int {
	return xxx_messageInfo_Server.Size(m)
}
func (m *Server) XXX_DiscardUnknown() {
	xxx_messageInfo_Server.DiscardUnknown(m)
}

var xxx_messageInfo_Server proto.InternalMessageInfo

func (m *Server) GetLocalip() string {
	if m != nil {
		return m.Localip
	}
	return ""
}

func (m *Server) GetPlayersonline() int64 {
	if m != nil {
		return m.Playersonline
	}
	return 0
}

func (m *Server) GetRemoteip() string {
	if m != nil {
		return m.Remoteip
	}
	return ""
}

func (m *Server) GetServerlistid() int64 {
	if m != nil {
		return m.Serverlistid
	}
	return 0
}

func (m *Server) GetServerlongname() string {
	if m != nil {
		return m.Serverlongname
	}
	return ""
}

func (m *Server) GetServershortname() string {
	if m != nil {
		return m.Servershortname
	}
	return ""
}

func (m *Server) GetServerstatus() int64 {
	if m != nil {
		return m.Serverstatus
	}
	return 0
}

func (m *Server) GetZonesbooted() int64 {
	if m != nil {
		return m.Zonesbooted
	}
	return 0
}

func init() {
	proto.RegisterType((*Server)(nil), "pb.Server")
}

func init() { proto.RegisterFile("proto/login_server.proto", fileDescriptor_fc26b5533683afe0) }

var fileDescriptor_fc26b5533683afe0 = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xdd, 0x4a, 0xc3, 0x40,
	0x10, 0x85, 0x49, 0xaa, 0x69, 0x1d, 0xff, 0x60, 0xaf, 0x86, 0xe2, 0x45, 0x28, 0x22, 0xb9, 0x32,
	0x17, 0xbe, 0x49, 0x7d, 0x00, 0xd9, 0xd8, 0x21, 0x2e, 0x6c, 0x67, 0x96, 0xdd, 0x51, 0xd0, 0xa7,
	0xf1, 0x51, 0x85, 0x89, 0xda, 0x36, 0x97, 0xe7, 0xe3, 0xdb, 0x73, 0x96, 0x01, 0x4c, 0x59, 0x54,
	0xfa, 0x28, 0x63, 0xe0, 0x97, 0x42, 0xf9, 0x83, 0xf2, 0xa3, 0x21, 0x57, 0xa7, 0x61, 0x7d, 0x37,
	0x8a, 0x8c, 0x91, 0x7a, 0x9f, 0x42, 0xef, 0x99, 0x45, 0xbd, 0x06, 0xe1, 0x32, 0x19, 0x9b, 0xef,
	0x1a, 0x9a, 0x67, 0x7b, 0xe2, 0x10, 0x96, 0x51, 0x5e, 0x7d, 0x0c, 0x09, 0xab, 0xb6, 0xea, 0x2e,
	0xb6, 0x7f, 0xd1, 0xdd, 0xc3, 0x75, 0x8a, 0xfe, 0x93, 0x72, 0x11, 0x8e, 0x81, 0x09, 0xeb, 0xb6,
	0xea, 0x16, 0xdb, 0x53, 0xe8, 0xd6, 0xb0, 0xca, 0xb4, 0x17, 0xa5, 0x90, 0x70, 0x61, 0x05, 0xff,
	0xd9, 0x6d, 0xe0, 0x6a, 0xfa, 0x58, 0x0c, 0x45, 0xc3, 0x0e, 0xcf, 0xac, 0xe0, 0x84, 0xb9, 0x07,
	0xb8, 0xf9, 0xcd, 0xc2, 0x23, 0xfb, 0x3d, 0xe1, 0xb9, 0xb5, 0xcc, 0xa8, 0xeb, 0xe0, 0x76, 0x22,
	0xe5, 0x4d, 0xb2, 0x9a, 0xd8, 0x98, 0x38, 0xc7, 0x87, 0xd5, 0xa2, 0x5e, 0xdf, 0x0b, 0x2e, 0x8f,
	0x57, 0x27, 0xe6, 0x5a, 0xb8, 0xfc, 0x12, 0xa6, 0x32, 0x88, 0x28, 0xed, 0x70, 0x65, 0xca, 0x31,
	0x1a, 0x1a, 0xbb, 0xd4, 0xd3, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6d, 0xba, 0x48, 0x8a, 0x67,
	0x01, 0x00, 0x00,
}
