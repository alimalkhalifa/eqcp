// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/bug.proto

package pb

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

type Bug struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Zone                 string   `protobuf:"bytes,2,opt,name=zone,proto3" json:"zone,omitempty"`
	Clientversionid      int64    `protobuf:"varint,3,opt,name=clientversionid,proto3" json:"clientversionid,omitempty"`
	Clientversionname    string   `protobuf:"bytes,4,opt,name=clientversionname,proto3" json:"clientversionname,omitempty"`
	Accountid            int64    `protobuf:"varint,5,opt,name=accountid,proto3" json:"accountid,omitempty"`
	Characterid          int64    `protobuf:"varint,6,opt,name=characterid,proto3" json:"characterid,omitempty"`
	Charactername        string   `protobuf:"bytes,7,opt,name=charactername,proto3" json:"charactername,omitempty"`
	Reporterspoof        int64    `protobuf:"varint,8,opt,name=reporterspoof,proto3" json:"reporterspoof,omitempty"`
	Categoryid           int64    `protobuf:"varint,9,opt,name=categoryid,proto3" json:"categoryid,omitempty"`
	Categoryname         string   `protobuf:"bytes,10,opt,name=categoryname,proto3" json:"categoryname,omitempty"`
	Reportername         string   `protobuf:"bytes,11,opt,name=reportername,proto3" json:"reportername,omitempty"`
	Uipath               string   `protobuf:"bytes,12,opt,name=uipath,proto3" json:"uipath,omitempty"`
	Posx                 int64    `protobuf:"varint,13,opt,name=posx,proto3" json:"posx,omitempty"`
	Posy                 int64    `protobuf:"varint,14,opt,name=posy,proto3" json:"posy,omitempty"`
	Posz                 int64    `protobuf:"varint,15,opt,name=posz,proto3" json:"posz,omitempty"`
	Heading              int64    `protobuf:"varint,16,opt,name=heading,proto3" json:"heading,omitempty"`
	Timeplayed           int64    `protobuf:"varint,17,opt,name=timeplayed,proto3" json:"timeplayed,omitempty"`
	Targetid             int64    `protobuf:"varint,18,opt,name=targetid,proto3" json:"targetid,omitempty"`
	Targetname           string   `protobuf:"bytes,19,opt,name=targetname,proto3" json:"targetname,omitempty"`
	Optionalinfomask     int64    `protobuf:"varint,20,opt,name=optionalinfomask,proto3" json:"optionalinfomask,omitempty"`
	Canduplicate         int64    `protobuf:"varint,21,opt,name=canduplicate,proto3" json:"canduplicate,omitempty"`
	Crashbug             int64    `protobuf:"varint,22,opt,name=crashbug,proto3" json:"crashbug,omitempty"`
	Targetinfo           int64    `protobuf:"varint,23,opt,name=targetinfo,proto3" json:"targetinfo,omitempty"`
	Characterflags       int64    `protobuf:"varint,24,opt,name=characterflags,proto3" json:"characterflags,omitempty"`
	Unknownvalue         int64    `protobuf:"varint,25,opt,name=unknownvalue,proto3" json:"unknownvalue,omitempty"`
	Bugreport            string   `protobuf:"bytes,26,opt,name=bugreport,proto3" json:"bugreport,omitempty"`
	Systeminfo           string   `protobuf:"bytes,27,opt,name=systeminfo,proto3" json:"systeminfo,omitempty"`
	Reportdatetime       int64    `protobuf:"varint,28,opt,name=reportdatetime,proto3" json:"reportdatetime,omitempty"`
	Bugstatus            int64    `protobuf:"varint,29,opt,name=bugstatus,proto3" json:"bugstatus,omitempty"`
	Lastreview           int64    `protobuf:"varint,30,opt,name=lastreview,proto3" json:"lastreview,omitempty"`
	Lastreviewer         string   `protobuf:"bytes,31,opt,name=lastreviewer,proto3" json:"lastreviewer,omitempty"`
	Reviewernotes        string   `protobuf:"bytes,32,opt,name=reviewernotes,proto3" json:"reviewernotes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bug) Reset()         { *m = Bug{} }
func (m *Bug) String() string { return proto.CompactTextString(m) }
func (*Bug) ProtoMessage()    {}
func (*Bug) Descriptor() ([]byte, []int) {
	return fileDescriptor_e891962b823e3f16, []int{0}
}

func (m *Bug) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bug.Unmarshal(m, b)
}
func (m *Bug) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bug.Marshal(b, m, deterministic)
}
func (m *Bug) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bug.Merge(m, src)
}
func (m *Bug) XXX_Size() int {
	return xxx_messageInfo_Bug.Size(m)
}
func (m *Bug) XXX_DiscardUnknown() {
	xxx_messageInfo_Bug.DiscardUnknown(m)
}

var xxx_messageInfo_Bug proto.InternalMessageInfo

func (m *Bug) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Bug) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func (m *Bug) GetClientversionid() int64 {
	if m != nil {
		return m.Clientversionid
	}
	return 0
}

func (m *Bug) GetClientversionname() string {
	if m != nil {
		return m.Clientversionname
	}
	return ""
}

func (m *Bug) GetAccountid() int64 {
	if m != nil {
		return m.Accountid
	}
	return 0
}

func (m *Bug) GetCharacterid() int64 {
	if m != nil {
		return m.Characterid
	}
	return 0
}

func (m *Bug) GetCharactername() string {
	if m != nil {
		return m.Charactername
	}
	return ""
}

func (m *Bug) GetReporterspoof() int64 {
	if m != nil {
		return m.Reporterspoof
	}
	return 0
}

func (m *Bug) GetCategoryid() int64 {
	if m != nil {
		return m.Categoryid
	}
	return 0
}

func (m *Bug) GetCategoryname() string {
	if m != nil {
		return m.Categoryname
	}
	return ""
}

func (m *Bug) GetReportername() string {
	if m != nil {
		return m.Reportername
	}
	return ""
}

func (m *Bug) GetUipath() string {
	if m != nil {
		return m.Uipath
	}
	return ""
}

func (m *Bug) GetPosx() int64 {
	if m != nil {
		return m.Posx
	}
	return 0
}

func (m *Bug) GetPosy() int64 {
	if m != nil {
		return m.Posy
	}
	return 0
}

func (m *Bug) GetPosz() int64 {
	if m != nil {
		return m.Posz
	}
	return 0
}

func (m *Bug) GetHeading() int64 {
	if m != nil {
		return m.Heading
	}
	return 0
}

func (m *Bug) GetTimeplayed() int64 {
	if m != nil {
		return m.Timeplayed
	}
	return 0
}

func (m *Bug) GetTargetid() int64 {
	if m != nil {
		return m.Targetid
	}
	return 0
}

func (m *Bug) GetTargetname() string {
	if m != nil {
		return m.Targetname
	}
	return ""
}

func (m *Bug) GetOptionalinfomask() int64 {
	if m != nil {
		return m.Optionalinfomask
	}
	return 0
}

func (m *Bug) GetCanduplicate() int64 {
	if m != nil {
		return m.Canduplicate
	}
	return 0
}

func (m *Bug) GetCrashbug() int64 {
	if m != nil {
		return m.Crashbug
	}
	return 0
}

func (m *Bug) GetTargetinfo() int64 {
	if m != nil {
		return m.Targetinfo
	}
	return 0
}

func (m *Bug) GetCharacterflags() int64 {
	if m != nil {
		return m.Characterflags
	}
	return 0
}

func (m *Bug) GetUnknownvalue() int64 {
	if m != nil {
		return m.Unknownvalue
	}
	return 0
}

func (m *Bug) GetBugreport() string {
	if m != nil {
		return m.Bugreport
	}
	return ""
}

func (m *Bug) GetSysteminfo() string {
	if m != nil {
		return m.Systeminfo
	}
	return ""
}

func (m *Bug) GetReportdatetime() int64 {
	if m != nil {
		return m.Reportdatetime
	}
	return 0
}

func (m *Bug) GetBugstatus() int64 {
	if m != nil {
		return m.Bugstatus
	}
	return 0
}

func (m *Bug) GetLastreview() int64 {
	if m != nil {
		return m.Lastreview
	}
	return 0
}

func (m *Bug) GetLastreviewer() string {
	if m != nil {
		return m.Lastreviewer
	}
	return ""
}

func (m *Bug) GetReviewernotes() string {
	if m != nil {
		return m.Reviewernotes
	}
	return ""
}

func init() {
	proto.RegisterType((*Bug)(nil), "pb.Bug")
}

func init() {
	proto.RegisterFile("proto/bug.proto", fileDescriptor_e891962b823e3f16)
}

var fileDescriptor_e891962b823e3f16 = []byte{
	// 489 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x94, 0xdb, 0x6e, 0x13, 0x31,
	0x10, 0x86, 0x95, 0xa4, 0xa4, 0x89, 0xdb, 0x26, 0xad, 0x81, 0x32, 0x94, 0x52, 0xa2, 0x0a, 0xa1,
	0x08, 0x21, 0xb8, 0xe0, 0x0d, 0x78, 0x84, 0xbe, 0x81, 0x77, 0x3d, 0xd9, 0x58, 0xdd, 0xd8, 0x2b,
	0x1f, 0x52, 0x36, 0xf7, 0xbc, 0x37, 0xf2, 0x78, 0xb3, 0x87, 0xe6, 0x6e, 0xe6, 0xf3, 0xbf, 0x33,
	0xff, 0x8c, 0x9d, 0xb0, 0x65, 0x65, 0x8d, 0x37, 0xbf, 0xb2, 0x50, 0xfc, 0xa4, 0x88, 0x8f, 0xab,
	0xec, 0xf1, 0xdf, 0x8c, 0x4d, 0xfe, 0x84, 0x82, 0x2f, 0xd8, 0x58, 0x49, 0x18, 0xad, 0x46, 0xeb,
	0xc9, 0xd3, 0x58, 0x49, 0xce, 0xd9, 0xd9, 0xc1, 0x68, 0x84, 0xf1, 0x6a, 0xb4, 0x9e, 0x3f, 0x51,
	0xcc, 0xd7, 0x6c, 0x99, 0x97, 0x0a, 0xb5, 0xdf, 0xa3, 0x75, 0xca, 0x68, 0x25, 0x61, 0x42, 0x1f,
	0xbc, 0xc6, 0xfc, 0x07, 0xbb, 0x19, 0x20, 0x2d, 0x76, 0x08, 0x67, 0x54, 0xea, 0xf4, 0x80, 0xdf,
	0xb3, 0xb9, 0xc8, 0x73, 0x13, 0xb4, 0x57, 0x12, 0xde, 0x50, 0xc5, 0x0e, 0xf0, 0x15, 0xbb, 0xc8,
	0xb7, 0xc2, 0x8a, 0xdc, 0xa3, 0x55, 0x12, 0xa6, 0x74, 0xde, 0x47, 0xfc, 0x2b, 0xbb, 0x6a, 0x53,
	0xea, 0x74, 0x4e, 0x9d, 0x86, 0x30, 0xaa, 0x2c, 0x56, 0xc6, 0x7a, 0xb4, 0xae, 0x32, 0x66, 0x03,
	0x33, 0xaa, 0x34, 0x84, 0xfc, 0x81, 0xb1, 0x5c, 0x78, 0x2c, 0x8c, 0xad, 0x95, 0x84, 0x39, 0x49,
	0x7a, 0x84, 0x3f, 0xb2, 0xcb, 0x63, 0x46, 0xad, 0x18, 0xb5, 0x1a, 0xb0, 0xa8, 0x39, 0x16, 0x25,
	0xcd, 0x45, 0xd2, 0xf4, 0x19, 0xbf, 0x65, 0xd3, 0xa0, 0x2a, 0xe1, 0xb7, 0x70, 0x49, 0xa7, 0x4d,
	0x16, 0xf7, 0x5e, 0x19, 0xf7, 0x17, 0xae, 0xa8, 0x33, 0xc5, 0x0d, 0xab, 0x61, 0xd1, 0xb2, 0xba,
	0x61, 0x07, 0x58, 0xb6, 0xec, 0xc0, 0x81, 0x9d, 0x6f, 0x51, 0x48, 0xa5, 0x0b, 0xb8, 0x26, 0x7c,
	0x4c, 0xe3, 0x54, 0x5e, 0xed, 0xb0, 0x2a, 0x45, 0x8d, 0x12, 0x6e, 0xd2, 0x54, 0x1d, 0xe1, 0x77,
	0x6c, 0xe6, 0x85, 0x2d, 0x30, 0x5e, 0x00, 0xa7, 0xd3, 0x36, 0xa7, 0x6f, 0x29, 0xa6, 0x59, 0xde,
	0x92, 0xdb, 0x1e, 0xe1, 0xdf, 0xd9, 0xb5, 0xa9, 0xbc, 0x32, 0x5a, 0x94, 0x4a, 0x6f, 0xcc, 0x4e,
	0xb8, 0x67, 0x78, 0x47, 0x35, 0x4e, 0x78, 0xda, 0x9e, 0x96, 0xa1, 0x2a, 0x55, 0xdc, 0x18, 0xbc,
	0x27, 0xdd, 0x80, 0x45, 0x2f, 0xb9, 0x15, 0x6e, 0x9b, 0x85, 0x02, 0x6e, 0x93, 0x97, 0x63, 0xde,
	0x79, 0x89, 0x15, 0xe1, 0x43, 0x33, 0x47, 0x4b, 0xf8, 0x37, 0xb6, 0x68, 0x2f, 0x7d, 0x53, 0x8a,
	0xc2, 0x01, 0x90, 0xe6, 0x15, 0x8d, 0x3e, 0x82, 0x7e, 0xd6, 0xe6, 0x45, 0xef, 0x45, 0x19, 0x10,
	0x3e, 0x26, 0x1f, 0x7d, 0x16, 0x5f, 0x65, 0x16, 0x8a, 0x74, 0x69, 0x70, 0x47, 0x63, 0x77, 0x20,
	0x3a, 0x71, 0xb5, 0xf3, 0xb8, 0x23, 0x27, 0x9f, 0xd2, 0x56, 0x3a, 0x12, 0x9d, 0x24, 0xa5, 0x14,
	0x1e, 0xe3, 0xa6, 0xe1, 0x3e, 0x39, 0x19, 0xd2, 0xa6, 0x8b, 0xf3, 0xc2, 0x07, 0x07, 0x9f, 0xd3,
	0xdb, 0x6f, 0x41, 0xec, 0x52, 0x0a, 0xe7, 0x2d, 0xee, 0x15, 0xbe, 0xc0, 0x43, 0x9a, 0xb7, 0x23,
	0x71, 0x8e, 0x2e, 0x43, 0x0b, 0x5f, 0xd2, 0x4b, 0xeb, 0xb3, 0xf4, 0xee, 0x53, 0xac, 0x8d, 0x47,
	0x07, 0xab, 0xf4, 0xeb, 0x18, 0xc0, 0x6c, 0x4a, 0x7f, 0x09, 0xbf, 0xff, 0x07, 0x00, 0x00, 0xff,
	0xff, 0x56, 0x08, 0xbf, 0xef, 0x25, 0x04, 0x00, 0x00,
}
