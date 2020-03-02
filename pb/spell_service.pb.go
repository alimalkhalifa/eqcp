// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/spell_service.proto

package pb

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SpellSearchRequest struct {
	Values               map[string]string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Limit                int64             `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               int64             `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Orderby              string            `protobuf:"bytes,4,opt,name=orderby,proto3" json:"orderby,omitempty"`
	Orderdesc            bool              `protobuf:"varint,5,opt,name=orderdesc,proto3" json:"orderdesc,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SpellSearchRequest) Reset()         { *m = SpellSearchRequest{} }
func (m *SpellSearchRequest) String() string { return proto.CompactTextString(m) }
func (*SpellSearchRequest) ProtoMessage()    {}
func (*SpellSearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d20b828b8f248bc, []int{0}
}

func (m *SpellSearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpellSearchRequest.Unmarshal(m, b)
}
func (m *SpellSearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpellSearchRequest.Marshal(b, m, deterministic)
}
func (m *SpellSearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpellSearchRequest.Merge(m, src)
}
func (m *SpellSearchRequest) XXX_Size() int {
	return xxx_messageInfo_SpellSearchRequest.Size(m)
}
func (m *SpellSearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SpellSearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SpellSearchRequest proto.InternalMessageInfo

func (m *SpellSearchRequest) GetValues() map[string]string {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *SpellSearchRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *SpellSearchRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *SpellSearchRequest) GetOrderby() string {
	if m != nil {
		return m.Orderby
	}
	return ""
}

func (m *SpellSearchRequest) GetOrderdesc() bool {
	if m != nil {
		return m.Orderdesc
	}
	return false
}

type SpellSearchResponse struct {
	Spells               []*Spell `protobuf:"bytes,1,rep,name=Spells,proto3" json:"Spells,omitempty"`
	Total                int64    `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpellSearchResponse) Reset()         { *m = SpellSearchResponse{} }
func (m *SpellSearchResponse) String() string { return proto.CompactTextString(m) }
func (*SpellSearchResponse) ProtoMessage()    {}
func (*SpellSearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d20b828b8f248bc, []int{1}
}

func (m *SpellSearchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpellSearchResponse.Unmarshal(m, b)
}
func (m *SpellSearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpellSearchResponse.Marshal(b, m, deterministic)
}
func (m *SpellSearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpellSearchResponse.Merge(m, src)
}
func (m *SpellSearchResponse) XXX_Size() int {
	return xxx_messageInfo_SpellSearchResponse.Size(m)
}
func (m *SpellSearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SpellSearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SpellSearchResponse proto.InternalMessageInfo

func (m *SpellSearchResponse) GetSpells() []*Spell {
	if m != nil {
		return m.Spells
	}
	return nil
}

func (m *SpellSearchResponse) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

type SpellCreateRequest struct {
	Values               map[string]string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SpellCreateRequest) Reset()         { *m = SpellCreateRequest{} }
func (m *SpellCreateRequest) String() string { return proto.CompactTextString(m) }
func (*SpellCreateRequest) ProtoMessage()    {}
func (*SpellCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d20b828b8f248bc, []int{2}
}

func (m *SpellCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpellCreateRequest.Unmarshal(m, b)
}
func (m *SpellCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpellCreateRequest.Marshal(b, m, deterministic)
}
func (m *SpellCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpellCreateRequest.Merge(m, src)
}
func (m *SpellCreateRequest) XXX_Size() int {
	return xxx_messageInfo_SpellCreateRequest.Size(m)
}
func (m *SpellCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SpellCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SpellCreateRequest proto.InternalMessageInfo

func (m *SpellCreateRequest) GetValues() map[string]string {
	if m != nil {
		return m.Values
	}
	return nil
}

type SpellCreateResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpellCreateResponse) Reset()         { *m = SpellCreateResponse{} }
func (m *SpellCreateResponse) String() string { return proto.CompactTextString(m) }
func (*SpellCreateResponse) ProtoMessage()    {}
func (*SpellCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d20b828b8f248bc, []int{3}
}

func (m *SpellCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpellCreateResponse.Unmarshal(m, b)
}
func (m *SpellCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpellCreateResponse.Marshal(b, m, deterministic)
}
func (m *SpellCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpellCreateResponse.Merge(m, src)
}
func (m *SpellCreateResponse) XXX_Size() int {
	return xxx_messageInfo_SpellCreateResponse.Size(m)
}
func (m *SpellCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SpellCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SpellCreateResponse proto.InternalMessageInfo

func (m *SpellCreateResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type SpellReadRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpellReadRequest) Reset()         { *m = SpellReadRequest{} }
func (m *SpellReadRequest) String() string { return proto.CompactTextString(m) }
func (*SpellReadRequest) ProtoMessage()    {}
func (*SpellReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d20b828b8f248bc, []int{4}
}

func (m *SpellReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpellReadRequest.Unmarshal(m, b)
}
func (m *SpellReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpellReadRequest.Marshal(b, m, deterministic)
}
func (m *SpellReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpellReadRequest.Merge(m, src)
}
func (m *SpellReadRequest) XXX_Size() int {
	return xxx_messageInfo_SpellReadRequest.Size(m)
}
func (m *SpellReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SpellReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SpellReadRequest proto.InternalMessageInfo

func (m *SpellReadRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type SpellReadResponse struct {
	Spell                *Spell   `protobuf:"bytes,1,opt,name=spell,proto3" json:"spell,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpellReadResponse) Reset()         { *m = SpellReadResponse{} }
func (m *SpellReadResponse) String() string { return proto.CompactTextString(m) }
func (*SpellReadResponse) ProtoMessage()    {}
func (*SpellReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d20b828b8f248bc, []int{5}
}

func (m *SpellReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpellReadResponse.Unmarshal(m, b)
}
func (m *SpellReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpellReadResponse.Marshal(b, m, deterministic)
}
func (m *SpellReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpellReadResponse.Merge(m, src)
}
func (m *SpellReadResponse) XXX_Size() int {
	return xxx_messageInfo_SpellReadResponse.Size(m)
}
func (m *SpellReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SpellReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SpellReadResponse proto.InternalMessageInfo

func (m *SpellReadResponse) GetSpell() *Spell {
	if m != nil {
		return m.Spell
	}
	return nil
}

type SpellUpdateRequest struct {
	Id                   int64             `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Values               map[string]string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SpellUpdateRequest) Reset()         { *m = SpellUpdateRequest{} }
func (m *SpellUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*SpellUpdateRequest) ProtoMessage()    {}
func (*SpellUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d20b828b8f248bc, []int{6}
}

func (m *SpellUpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpellUpdateRequest.Unmarshal(m, b)
}
func (m *SpellUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpellUpdateRequest.Marshal(b, m, deterministic)
}
func (m *SpellUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpellUpdateRequest.Merge(m, src)
}
func (m *SpellUpdateRequest) XXX_Size() int {
	return xxx_messageInfo_SpellUpdateRequest.Size(m)
}
func (m *SpellUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SpellUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SpellUpdateRequest proto.InternalMessageInfo

func (m *SpellUpdateRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SpellUpdateRequest) GetValues() map[string]string {
	if m != nil {
		return m.Values
	}
	return nil
}

type SpellUpdateResponse struct {
	Rowsaffected         int64    `protobuf:"varint,1,opt,name=rowsaffected,proto3" json:"rowsaffected,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpellUpdateResponse) Reset()         { *m = SpellUpdateResponse{} }
func (m *SpellUpdateResponse) String() string { return proto.CompactTextString(m) }
func (*SpellUpdateResponse) ProtoMessage()    {}
func (*SpellUpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d20b828b8f248bc, []int{7}
}

func (m *SpellUpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpellUpdateResponse.Unmarshal(m, b)
}
func (m *SpellUpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpellUpdateResponse.Marshal(b, m, deterministic)
}
func (m *SpellUpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpellUpdateResponse.Merge(m, src)
}
func (m *SpellUpdateResponse) XXX_Size() int {
	return xxx_messageInfo_SpellUpdateResponse.Size(m)
}
func (m *SpellUpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SpellUpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SpellUpdateResponse proto.InternalMessageInfo

func (m *SpellUpdateResponse) GetRowsaffected() int64 {
	if m != nil {
		return m.Rowsaffected
	}
	return 0
}

type SpellDeleteRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpellDeleteRequest) Reset()         { *m = SpellDeleteRequest{} }
func (m *SpellDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*SpellDeleteRequest) ProtoMessage()    {}
func (*SpellDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d20b828b8f248bc, []int{8}
}

func (m *SpellDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpellDeleteRequest.Unmarshal(m, b)
}
func (m *SpellDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpellDeleteRequest.Marshal(b, m, deterministic)
}
func (m *SpellDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpellDeleteRequest.Merge(m, src)
}
func (m *SpellDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_SpellDeleteRequest.Size(m)
}
func (m *SpellDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SpellDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SpellDeleteRequest proto.InternalMessageInfo

func (m *SpellDeleteRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type SpellDeleteResponse struct {
	Rowsaffected         int64    `protobuf:"varint,1,opt,name=rowsaffected,proto3" json:"rowsaffected,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpellDeleteResponse) Reset()         { *m = SpellDeleteResponse{} }
func (m *SpellDeleteResponse) String() string { return proto.CompactTextString(m) }
func (*SpellDeleteResponse) ProtoMessage()    {}
func (*SpellDeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d20b828b8f248bc, []int{9}
}

func (m *SpellDeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpellDeleteResponse.Unmarshal(m, b)
}
func (m *SpellDeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpellDeleteResponse.Marshal(b, m, deterministic)
}
func (m *SpellDeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpellDeleteResponse.Merge(m, src)
}
func (m *SpellDeleteResponse) XXX_Size() int {
	return xxx_messageInfo_SpellDeleteResponse.Size(m)
}
func (m *SpellDeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SpellDeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SpellDeleteResponse proto.InternalMessageInfo

func (m *SpellDeleteResponse) GetRowsaffected() int64 {
	if m != nil {
		return m.Rowsaffected
	}
	return 0
}

type SpellPatchRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpellPatchRequest) Reset()         { *m = SpellPatchRequest{} }
func (m *SpellPatchRequest) String() string { return proto.CompactTextString(m) }
func (*SpellPatchRequest) ProtoMessage()    {}
func (*SpellPatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d20b828b8f248bc, []int{10}
}

func (m *SpellPatchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpellPatchRequest.Unmarshal(m, b)
}
func (m *SpellPatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpellPatchRequest.Marshal(b, m, deterministic)
}
func (m *SpellPatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpellPatchRequest.Merge(m, src)
}
func (m *SpellPatchRequest) XXX_Size() int {
	return xxx_messageInfo_SpellPatchRequest.Size(m)
}
func (m *SpellPatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SpellPatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SpellPatchRequest proto.InternalMessageInfo

func (m *SpellPatchRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SpellPatchRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SpellPatchRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type SpellPatchResponse struct {
	Rowsaffected         int64    `protobuf:"varint,1,opt,name=rowsaffected,proto3" json:"rowsaffected,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpellPatchResponse) Reset()         { *m = SpellPatchResponse{} }
func (m *SpellPatchResponse) String() string { return proto.CompactTextString(m) }
func (*SpellPatchResponse) ProtoMessage()    {}
func (*SpellPatchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d20b828b8f248bc, []int{11}
}

func (m *SpellPatchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpellPatchResponse.Unmarshal(m, b)
}
func (m *SpellPatchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpellPatchResponse.Marshal(b, m, deterministic)
}
func (m *SpellPatchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpellPatchResponse.Merge(m, src)
}
func (m *SpellPatchResponse) XXX_Size() int {
	return xxx_messageInfo_SpellPatchResponse.Size(m)
}
func (m *SpellPatchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SpellPatchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SpellPatchResponse proto.InternalMessageInfo

func (m *SpellPatchResponse) GetRowsaffected() int64 {
	if m != nil {
		return m.Rowsaffected
	}
	return 0
}

func init() {
	proto.RegisterType((*SpellSearchRequest)(nil), "pb.SpellSearchRequest")
	proto.RegisterMapType((map[string]string)(nil), "pb.SpellSearchRequest.ValuesEntry")
	proto.RegisterType((*SpellSearchResponse)(nil), "pb.SpellSearchResponse")
	proto.RegisterType((*SpellCreateRequest)(nil), "pb.SpellCreateRequest")
	proto.RegisterMapType((map[string]string)(nil), "pb.SpellCreateRequest.ValuesEntry")
	proto.RegisterType((*SpellCreateResponse)(nil), "pb.SpellCreateResponse")
	proto.RegisterType((*SpellReadRequest)(nil), "pb.SpellReadRequest")
	proto.RegisterType((*SpellReadResponse)(nil), "pb.SpellReadResponse")
	proto.RegisterType((*SpellUpdateRequest)(nil), "pb.SpellUpdateRequest")
	proto.RegisterMapType((map[string]string)(nil), "pb.SpellUpdateRequest.ValuesEntry")
	proto.RegisterType((*SpellUpdateResponse)(nil), "pb.SpellUpdateResponse")
	proto.RegisterType((*SpellDeleteRequest)(nil), "pb.SpellDeleteRequest")
	proto.RegisterType((*SpellDeleteResponse)(nil), "pb.SpellDeleteResponse")
	proto.RegisterType((*SpellPatchRequest)(nil), "pb.SpellPatchRequest")
	proto.RegisterType((*SpellPatchResponse)(nil), "pb.SpellPatchResponse")
}

func init() {
	proto.RegisterFile("proto/spell_service.proto", fileDescriptor_4d20b828b8f248bc)
}

var fileDescriptor_4d20b828b8f248bc = []byte{
	// 1531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x98, 0xcd, 0x6e, 0xdb, 0xc6,
	0x16, 0xc7, 0x43, 0xf9, 0xda, 0xf7, 0x7a, 0x1c, 0x3b, 0x0e, 0xed, 0xeb, 0xab, 0xab, 0x1b, 0xe0,
	0x32, 0xe3, 0xe4, 0x5e, 0x37, 0x70, 0x6c, 0x91, 0x92, 0xfc, 0xc1, 0xa2, 0x49, 0x84, 0xc6, 0x9b,
	0xa0, 0x8b, 0x40, 0xf9, 0x00, 0xba, 0x0a, 0x46, 0xe4, 0x48, 0x19, 0x84, 0x9c, 0x61, 0x38, 0x94,
	0x1d, 0xd5, 0x70, 0x17, 0x45, 0x9f, 0x40, 0x5d, 0xb6, 0x8b, 0xa0, 0x8b, 0x3e, 0x42, 0xd1, 0x4d,
	0xbb, 0xe9, 0x23, 0x64, 0x57, 0xa0, 0x40, 0x17, 0x79, 0x81, 0xae, 0xba, 0xe8, 0xa2, 0xc5, 0x7c,
	0x50, 0xfc, 0x52, 0x02, 0x04, 0x4d, 0x56, 0xe1, 0xfc, 0x74, 0x38, 0x73, 0xfe, 0xe7, 0x9c, 0x39,
	0x3c, 0x0e, 0xf8, 0x77, 0x14, 0xb3, 0x84, 0xed, 0xf2, 0x08, 0x07, 0xc1, 0x23, 0x8e, 0xe3, 0x63,
	0xe2, 0xe1, 0x1d, 0xc9, 0xcc, 0x5a, 0xd4, 0x6f, 0x5c, 0x1a, 0x32, 0x36, 0x0c, 0xf0, 0x2e, 0x8a,
	0xc8, 0x2e, 0xa2, 0x94, 0x25, 0x28, 0x21, 0x8c, 0x72, 0x65, 0xd1, 0xd8, 0x96, 0xff, 0x78, 0xd7,
	0x87, 0x98, 0x5e, 0xe7, 0x27, 0x68, 0x38, 0xc4, 0xf1, 0x2e, 0x8b, 0xa4, 0xc5, 0x0c, 0xeb, 0x8b,
	0xb9, 0xa3, 0x14, 0x82, 0xbf, 0x1a, 0xc0, 0xbc, 0x27, 0xd6, 0xf7, 0x30, 0x8a, 0xbd, 0xc7, 0x3d,
	0xfc, 0x74, 0x84, 0x79, 0x62, 0xba, 0x60, 0xe1, 0x18, 0x05, 0x23, 0xcc, 0xeb, 0x86, 0x35, 0xb7,
	0xb5, 0xe4, 0xc0, 0x9d, 0xa8, 0xbf, 0x53, 0xb5, 0xdb, 0x79, 0x28, 0x8d, 0x8e, 0x68, 0x12, 0x8f,
	0x7b, 0xfa, 0x0d, 0x73, 0x1d, 0xcc, 0x07, 0x24, 0x24, 0x49, 0xbd, 0x66, 0x19, 0x5b, 0x73, 0x3d,
	0xb5, 0x30, 0x37, 0xc0, 0x02, 0x1b, 0x0c, 0x38, 0x4e, 0xea, 0x73, 0x12, 0xeb, 0x95, 0x59, 0x07,
	0x7f, 0x67, 0xb1, 0x8f, 0xe3, 0xfe, 0xb8, 0xfe, 0x37, 0xcb, 0xd8, 0x5a, 0xec, 0xa5, 0x4b, 0xf3,
	0x12, 0x58, 0x94, 0x8f, 0x3e, 0xe6, 0x5e, 0x7d, 0xde, 0x32, 0xb6, 0xfe, 0xd1, 0xcb, 0x40, 0xe3,
	0x10, 0x2c, 0xe5, 0x0e, 0x37, 0x57, 0xc1, 0xdc, 0x13, 0x3c, 0xae, 0x1b, 0x72, 0x0b, 0xf1, 0x28,
	0xdc, 0x90, 0x0e, 0x49, 0x37, 0x16, 0x7b, 0x6a, 0xe1, 0xd6, 0x0e, 0x0c, 0xf8, 0x00, 0xac, 0x15,
	0xa4, 0xf0, 0x88, 0x51, 0x8e, 0xcd, 0xcb, 0x60, 0x41, 0xe2, 0x54, 0xf3, 0xe2, 0x54, 0x73, 0x4f,
	0xff, 0x20, 0xf6, 0x4c, 0x58, 0x82, 0x82, 0x54, 0x9a, 0x5c, 0xb8, 0xf3, 0x93, 0x6e, 0xcd, 0x39,
	0x07, 0x27, 0x69, 0x28, 0x3f, 0x8c, 0x31, 0x4a, 0x70, 0x35, 0x94, 0xb5, 0x52, 0x28, 0x0b, 0x76,
	0xb3, 0x42, 0xf9, 0x17, 0x44, 0xa6, 0x4e, 0x6d, 0x6b, 0xad, 0xe9, 0x59, 0x5a, 0xeb, 0x0a, 0xa8,
	0x11, 0x5f, 0x6e, 0x34, 0xd7, 0xab, 0x11, 0x3f, 0xb5, 0x86, 0x60, 0x55, 0x09, 0xc6, 0xc8, 0x4f,
	0xfd, 0x2f, 0x99, 0xc2, 0x2f, 0x57, 0xc0, 0xc5, 0x9c, 0x91, 0xde, 0xf0, 0xbf, 0x60, 0x5e, 0x96,
	0x95, 0x34, 0x2c, 0xc4, 0x4e, 0x71, 0xf7, 0xe5, 0xf2, 0xa4, 0xfb, 0xcb, 0xb2, 0xf3, 0xf3, 0xb2,
	0xf9, 0xd3, 0xf2, 0x29, 0x94, 0x0c, 0xba, 0xa7, 0x90, 0xf8, 0xd0, 0x85, 0x76, 0xb3, 0x09, 0xb7,
	0x21, 0x45, 0x21, 0x86, 0x2e, 0xbc, 0x37, 0x0a, 0x43, 0x46, 0xad, 0xfb, 0x8f, 0x63, 0x76, 0x42,
	0xe8, 0xd0, 0xba, 0x2d, 0xab, 0x1a, 0x6e, 0xc3, 0x28, 0x40, 0x63, 0x1c, 0xdb, 0xd0, 0x85, 0x77,
	0x3f, 0xea, 0x7e, 0x7c, 0xd4, 0x7b, 0x64, 0xc3, 0x6d, 0xe8, 0x21, 0x9e, 0x24, 0x44, 0xbe, 0xd8,
	0x6e, 0xca, 0x6d, 0x62, 0xec, 0xb1, 0x63, 0x1c, 0x8f, 0x35, 0xb5, 0x3b, 0x29, 0xcd, 0x2c, 0x35,
	0x0b, 0x11, 0x45, 0xd0, 0x85, 0x8e, 0x78, 0xc6, 0x83, 0x01, 0xf6, 0x92, 0x3e, 0xe2, 0x58, 0x86,
	0x50, 0x9c, 0x72, 0xd0, 0xb2, 0xf7, 0xa5, 0xd5, 0x33, 0x3b, 0xb5, 0x22, 0x1e, 0xa3, 0xe2, 0xb9,
	0x63, 0x3b, 0xe2, 0x17, 0x1c, 0xa6, 0xc0, 0x6e, 0xee, 0x09, 0x67, 0x58, 0x18, 0x31, 0x8a, 0x69,
	0xc2, 0xc5, 0x1b, 0xd7, 0xed, 0x02, 0x72, 0xaa, 0xa8, 0x55, 0x45, 0xed, 0x0a, 0xf2, 0xd8, 0x48,
	0x6f, 0x38, 0x03, 0x3b, 0xb3, 0x71, 0x6b, 0x36, 0x6e, 0x6b, 0x4c, 0x19, 0x7e, 0x16, 0x61, 0xea,
	0xc7, 0x18, 0x0d, 0x31, 0x4d, 0xa6, 0xce, 0x96, 0xb8, 0xf3, 0x0a, 0xde, 0x7a, 0x05, 0x9f, 0xfa,
	0x3e, 0x60, 0x71, 0x38, 0x0a, 0x90, 0xf4, 0xb9, 0x79, 0x98, 0x01, 0x67, 0x9a, 0x6b, 0x0d, 0x5a,
	0x65, 0xd0, 0x2e, 0x83, 0x4e, 0x19, 0xec, 0x95, 0xc1, 0x7e, 0x19, 0x1c, 0x94, 0xc1, 0x61, 0x19,
	0xd8, 0xcd, 0x0a, 0xb1, 0x2b, 0x24, 0x73, 0x76, 0xc8, 0x98, 0xaf, 0xea, 0x43, 0xc7, 0x4f, 0x2d,
	0x88, 0x2f, 0x5e, 0x6a, 0x39, 0x39, 0xe0, 0xc8, 0xe2, 0x68, 0xe7, 0x48, 0xab, 0x42, 0xda, 0x15,
	0xd2, 0xa9, 0x90, 0xbd, 0x0a, 0xd9, 0xaf, 0x90, 0x83, 0x0a, 0x39, 0xac, 0x10, 0x29, 0xb4, 0x84,
	0xec, 0x2a, 0xca, 0xdc, 0x4e, 0x50, 0x3c, 0xc4, 0x49, 0x32, 0x8e, 0xc4, 0x35, 0x11, 0x25, 0x2d,
	0x6e, 0x84, 0x4f, 0x06, 0x03, 0x69, 0x01, 0xb7, 0x21, 0x7f, 0x42, 0xc4, 0xc5, 0x85, 0xb6, 0xb0,
	0xfe, 0x84, 0x51, 0xac, 0x6d, 0x55, 0xcd, 0x06, 0x88, 0x73, 0xcc, 0xd5, 0x09, 0x9d, 0x0c, 0x38,
	0x65, 0xd0, 0x2a, 0x83, 0x76, 0x19, 0x74, 0xca, 0x60, 0xaf, 0x0c, 0xf6, 0xcb, 0xe0, 0xa0, 0x0c,
	0x0e, 0xcb, 0x40, 0x87, 0x23, 0x4f, 0x2a, 0xbe, 0xda, 0x15, 0x67, 0x6d, 0x59, 0xa5, 0x79, 0x50,
	0xf1, 0xd6, 0xae, 0xb8, 0x6b, 0xe7, 0xfc, 0x45, 0x3c, 0x21, 0x74, 0x88, 0x28, 0x09, 0x45, 0x97,
	0x6a, 0x89, 0x28, 0x8a, 0xf6, 0x87, 0x54, 0xf8, 0xa9, 0x8f, 0x9f, 0x09, 0x2e, 0x4a, 0x8f, 0xe0,
	0xc0, 0xb7, 0xdb, 0x59, 0xe5, 0x51, 0x7c, 0x92, 0x76, 0x98, 0xe9, 0x5b, 0x6a, 0x9b, 0x56, 0x53,
	0x6c, 0xed, 0xe3, 0x00, 0x27, 0x18, 0xf5, 0x03, 0xac, 0x2b, 0x53, 0x7c, 0x1a, 0xe9, 0x28, 0x9c,
	0x6e, 0x20, 0xb2, 0x93, 0x63, 0x07, 0xd3, 0xac, 0xe7, 0xa0, 0xdd, 0x9c, 0x1e, 0xbd, 0x97, 0x5d,
	0xc8, 0xe8, 0x38, 0x8a, 0x31, 0x27, 0x5c, 0xf6, 0x44, 0xe8, 0xc2, 0x4e, 0x9e, 0x79, 0x28, 0xf0,
	0xa4, 0xa5, 0x5d, 0xa4, 0x51, 0x7a, 0x88, 0x74, 0xd5, 0x43, 0x09, 0x1e, 0xb2, 0x78, 0x2c, 0x1b,
	0xa4, 0x68, 0xa5, 0x34, 0xf2, 0x46, 0x1c, 0x0f, 0x46, 0x01, 0xc5, 0x9c, 0x8b, 0x3d, 0xc5, 0x41,
	0x7c, 0xc4, 0x45, 0x0f, 0xc9, 0x89, 0x90, 0xbe, 0x38, 0xcd, 0xbd, 0x69, 0x3b, 0x91, 0x6b, 0x75,
	0x69, 0x45, 0x13, 0xa6, 0x1e, 0x0b, 0xfb, 0x28, 0xbd, 0x8c, 0x6c, 0x94, 0xb0, 0x41, 0x81, 0x68,
	0xfb, 0x4e, 0x71, 0xb9, 0x5f, 0xdc, 0xce, 0x11, 0x99, 0xdf, 0xcf, 0x96, 0x22, 0xe8, 0xad, 0x4e,
	0xb6, 0x6e, 0x95, 0xcc, 0xd3, 0xdd, 0x42, 0x42, 0x7d, 0xc2, 0x93, 0x90, 0xf9, 0xd0, 0xb5, 0xe5,
	0xd7, 0x21, 0x5b, 0x9e, 0x9d, 0xc1, 0xaf, 0xd3, 0x21, 0xe0, 0x41, 0xe4, 0xe7, 0x86, 0x80, 0xd2,
	0x47, 0xf4, 0x35, 0x43, 0x41, 0xe1, 0xbd, 0x77, 0x34, 0x14, 0x3c, 0xd4, 0x43, 0x41, 0x7a, 0x96,
	0xfe, 0x86, 0x43, 0x70, 0x3e, 0x66, 0x27, 0x5c, 0xd5, 0x25, 0x4e, 0xdd, 0x2d, 0x30, 0xf7, 0x3f,
	0x93, 0x6e, 0xdd, 0xd9, 0x30, 0xd7, 0x4f, 0x61, 0x1e, 0xcb, 0xd8, 0x9c, 0xc1, 0x2b, 0x5a, 0xfb,
	0x6d, 0x59, 0x96, 0xaf, 0x1a, 0x20, 0xd2, 0xd3, 0x53, 0xab, 0xb7, 0x75, 0x7a, 0x4f, 0xcf, 0x25,
	0x77, 0x51, 0x92, 0x0d, 0xb2, 0xe5, 0xc0, 0xeb, 0x68, 0xd5, 0x66, 0x44, 0x6b, 0x2e, 0x1f, 0x2d,
	0x1d, 0xa9, 0x07, 0x5a, 0x91, 0xde, 0xf3, 0x2d, 0xb9, 0xea, 0x3c, 0x5f, 0x02, 0xe7, 0xf5, 0x08,
	0x2a, 0xe7, 0x7d, 0xf3, 0x85, 0x01, 0x96, 0x72, 0x33, 0xa9, 0xb9, 0x31, 0x7b, 0xde, 0x6e, 0xfc,
	0xab, 0xc2, 0x95, 0x4b, 0x70, 0x62, 0x4c, 0xba, 0x27, 0xe6, 0xb2, 0xa2, 0x96, 0xbc, 0x7e, 0xbc,
	0x71, 0xa5, 0xb0, 0xb4, 0x4e, 0x48, 0xf2, 0xd8, 0x8a, 0x62, 0x76, 0x4c, 0x7c, 0xec, 0x5b, 0x4f,
	0x47, 0x38, 0x26, 0x98, 0xef, 0xdc, 0xb9, 0x05, 0xe6, 0x9c, 0x66, 0xd3, 0x3c, 0x04, 0x2b, 0xda,
	0xce, 0xc7, 0x09, 0x22, 0x01, 0x37, 0xff, 0x0f, 0xae, 0x36, 0x36, 0x37, 0x77, 0x7d, 0x3c, 0x20,
	0x94, 0xa8, 0xbf, 0x29, 0xa2, 0xfe, 0x8c, 0xc3, 0x3f, 0x7b, 0xf1, 0xf2, 0x8b, 0x9a, 0x69, 0xae,
	0xee, 0x1e, 0xdb, 0xea, 0x6f, 0x8b, 0x5d, 0xae, 0x54, 0xfc, 0x91, 0xaa, 0x52, 0xd3, 0x67, 0x4e,
	0x55, 0x61, 0xf4, 0xcd, 0xa9, 0x2a, 0x8e, 0xa9, 0xf0, 0x07, 0x63, 0xd2, 0xfd, 0xc6, 0x30, 0x57,
	0x15, 0xb6, 0x28, 0x3e, 0x51, 0x5a, 0x1a, 0x37, 0x34, 0x41, 0x19, 0xdb, 0x01, 0x77, 0x71, 0x1c,
	0x12, 0xce, 0x09, 0xa3, 0x96, 0xfc, 0x6c, 0x59, 0x8c, 0xba, 0xd6, 0x66, 0x38, 0xe6, 0x4f, 0x03,
	0x6b, 0xcb, 0xc7, 0x03, 0x34, 0x0a, 0x12, 0xd7, 0x72, 0x3a, 0x9d, 0xf7, 0xee, 0x1c, 0x29, 0xcd,
	0x37, 0xc0, 0x1a, 0xf1, 0x2d, 0x36, 0x50, 0x3b, 0x58, 0x9e, 0xdc, 0xd4, 0x7f, 0x8d, 0xf0, 0xa2,
	0x7f, 0xfd, 0x25, 0xb0, 0x08, 0xe6, 0xef, 0xb3, 0x27, 0x98, 0x9a, 0xe7, 0x64, 0x14, 0x56, 0xe0,
	0xe2, 0x34, 0x0a, 0xae, 0x71, 0xcd, 0xfc, 0xce, 0x00, 0x8b, 0xd3, 0x61, 0xd9, 0x5c, 0xcf, 0xa6,
	0xe2, 0x6c, 0xc0, 0x6e, 0xfc, 0xb3, 0x44, 0xb5, 0xf6, 0x4f, 0x27, 0x5d, 0x6c, 0x5e, 0x10, 0x48,
	0x7b, 0x47, 0xe8, 0x80, 0x35, 0xd6, 0x72, 0x40, 0xe7, 0x69, 0xe7, 0xce, 0x07, 0x4a, 0xcd, 0x1e,
	0x58, 0x2e, 0xfc, 0x60, 0x5e, 0x05, 0x9b, 0x8d, 0xcb, 0x33, 0x75, 0xe4, 0x4f, 0x92, 0x8e, 0xaf,
	0x9a, 0x2b, 0x59, 0xfa, 0x4e, 0x89, 0x7f, 0x66, 0x7e, 0x5e, 0xd3, 0xc9, 0x53, 0x5d, 0x22, 0x97,
	0xbc, 0x42, 0x8b, 0xca, 0x25, 0xaf, 0xd8, 0x4e, 0xe0, 0x8f, 0x2a, 0x79, 0xeb, 0x0a, 0x17, 0x7d,
	0x6e, 0xdc, 0x9a, 0x45, 0xdf, 0x28, 0x85, 0x37, 0x95, 0xe8, 0x03, 0xb0, 0x2c, 0x2e, 0x98, 0x95,
	0xde, 0xb0, 0xd7, 0x24, 0xaf, 0xe8, 0x5f, 0x35, 0x79, 0x6b, 0x8d, 0x52, 0x0c, 0x44, 0x06, 0x7f,
	0x4b, 0x6b, 0x58, 0xb5, 0xab, 0x5c, 0x18, 0x0a, 0x5d, 0x2e, 0x17, 0x86, 0x62, 0x5f, 0x83, 0xdf,
	0x1a, 0x93, 0xee, 0x57, 0x86, 0x79, 0x41, 0x61, 0x0b, 0x51, 0x5d, 0xc2, 0xef, 0x97, 0xc0, 0x3b,
	0x16, 0x5f, 0xf4, 0xaa, 0x2a, 0x7e, 0xf5, 0x5a, 0xb9, 0x00, 0x7e, 0x37, 0x00, 0xc8, 0x9a, 0x9f,
	0x99, 0x95, 0x69, 0xbe, 0xc1, 0x36, 0x36, 0xca, 0x58, 0xcb, 0xfe, 0xde, 0x98, 0x74, 0x9f, 0x1b,
	0xe6, 0xea, 0x91, 0x4f, 0x12, 0x9d, 0x65, 0xf9, 0x25, 0x6d, 0xdc, 0x94, 0x24, 0x55, 0xad, 0xe0,
	0x1b, 0x69, 0xbf, 0xa1, 0xb4, 0xef, 0x97, 0xb5, 0xff, 0x0f, 0x5c, 0x69, 0xc0, 0x99, 0xda, 0x0b,
	0x9e, 0xcd, 0xc8, 0xbb, 0x53, 0xcd, 0x7b, 0x7f, 0x41, 0xfe, 0xff, 0x48, 0xeb, 0xcf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x44, 0x14, 0xa6, 0x86, 0x9f, 0x11, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SpellServiceClient is the client API for SpellService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SpellServiceClient interface {
	SpellSearch(ctx context.Context, in *SpellSearchRequest, opts ...grpc.CallOption) (*SpellSearchResponse, error)
	SpellCreate(ctx context.Context, in *SpellCreateRequest, opts ...grpc.CallOption) (*SpellCreateResponse, error)
	SpellRead(ctx context.Context, in *SpellReadRequest, opts ...grpc.CallOption) (*SpellReadResponse, error)
	SpellUpdate(ctx context.Context, in *SpellUpdateRequest, opts ...grpc.CallOption) (*SpellUpdateResponse, error)
	SpellDelete(ctx context.Context, in *SpellDeleteRequest, opts ...grpc.CallOption) (*SpellDeleteResponse, error)
	SpellPatch(ctx context.Context, in *SpellPatchRequest, opts ...grpc.CallOption) (*SpellPatchResponse, error)
}

type spellServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSpellServiceClient(cc grpc.ClientConnInterface) SpellServiceClient {
	return &spellServiceClient{cc}
}

func (c *spellServiceClient) SpellSearch(ctx context.Context, in *SpellSearchRequest, opts ...grpc.CallOption) (*SpellSearchResponse, error) {
	out := new(SpellSearchResponse)
	err := c.cc.Invoke(ctx, "/pb.SpellService/SpellSearch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spellServiceClient) SpellCreate(ctx context.Context, in *SpellCreateRequest, opts ...grpc.CallOption) (*SpellCreateResponse, error) {
	out := new(SpellCreateResponse)
	err := c.cc.Invoke(ctx, "/pb.SpellService/SpellCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spellServiceClient) SpellRead(ctx context.Context, in *SpellReadRequest, opts ...grpc.CallOption) (*SpellReadResponse, error) {
	out := new(SpellReadResponse)
	err := c.cc.Invoke(ctx, "/pb.SpellService/SpellRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spellServiceClient) SpellUpdate(ctx context.Context, in *SpellUpdateRequest, opts ...grpc.CallOption) (*SpellUpdateResponse, error) {
	out := new(SpellUpdateResponse)
	err := c.cc.Invoke(ctx, "/pb.SpellService/SpellUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spellServiceClient) SpellDelete(ctx context.Context, in *SpellDeleteRequest, opts ...grpc.CallOption) (*SpellDeleteResponse, error) {
	out := new(SpellDeleteResponse)
	err := c.cc.Invoke(ctx, "/pb.SpellService/SpellDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spellServiceClient) SpellPatch(ctx context.Context, in *SpellPatchRequest, opts ...grpc.CallOption) (*SpellPatchResponse, error) {
	out := new(SpellPatchResponse)
	err := c.cc.Invoke(ctx, "/pb.SpellService/SpellPatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpellServiceServer is the server API for SpellService service.
type SpellServiceServer interface {
	SpellSearch(context.Context, *SpellSearchRequest) (*SpellSearchResponse, error)
	SpellCreate(context.Context, *SpellCreateRequest) (*SpellCreateResponse, error)
	SpellRead(context.Context, *SpellReadRequest) (*SpellReadResponse, error)
	SpellUpdate(context.Context, *SpellUpdateRequest) (*SpellUpdateResponse, error)
	SpellDelete(context.Context, *SpellDeleteRequest) (*SpellDeleteResponse, error)
	SpellPatch(context.Context, *SpellPatchRequest) (*SpellPatchResponse, error)
}

// UnimplementedSpellServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSpellServiceServer struct {
}

func (*UnimplementedSpellServiceServer) SpellSearch(ctx context.Context, req *SpellSearchRequest) (*SpellSearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SpellSearch not implemented")
}
func (*UnimplementedSpellServiceServer) SpellCreate(ctx context.Context, req *SpellCreateRequest) (*SpellCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SpellCreate not implemented")
}
func (*UnimplementedSpellServiceServer) SpellRead(ctx context.Context, req *SpellReadRequest) (*SpellReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SpellRead not implemented")
}
func (*UnimplementedSpellServiceServer) SpellUpdate(ctx context.Context, req *SpellUpdateRequest) (*SpellUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SpellUpdate not implemented")
}
func (*UnimplementedSpellServiceServer) SpellDelete(ctx context.Context, req *SpellDeleteRequest) (*SpellDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SpellDelete not implemented")
}
func (*UnimplementedSpellServiceServer) SpellPatch(ctx context.Context, req *SpellPatchRequest) (*SpellPatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SpellPatch not implemented")
}

func RegisterSpellServiceServer(s *grpc.Server, srv SpellServiceServer) {
	s.RegisterService(&_SpellService_serviceDesc, srv)
}

func _SpellService_SpellSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpellSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpellServiceServer).SpellSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SpellService/SpellSearch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpellServiceServer).SpellSearch(ctx, req.(*SpellSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpellService_SpellCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpellCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpellServiceServer).SpellCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SpellService/SpellCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpellServiceServer).SpellCreate(ctx, req.(*SpellCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpellService_SpellRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpellReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpellServiceServer).SpellRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SpellService/SpellRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpellServiceServer).SpellRead(ctx, req.(*SpellReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpellService_SpellUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpellUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpellServiceServer).SpellUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SpellService/SpellUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpellServiceServer).SpellUpdate(ctx, req.(*SpellUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpellService_SpellDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpellDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpellServiceServer).SpellDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SpellService/SpellDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpellServiceServer).SpellDelete(ctx, req.(*SpellDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpellService_SpellPatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpellPatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpellServiceServer).SpellPatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SpellService/SpellPatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpellServiceServer).SpellPatch(ctx, req.(*SpellPatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SpellService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.SpellService",
	HandlerType: (*SpellServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SpellSearch",
			Handler:    _SpellService_SpellSearch_Handler,
		},
		{
			MethodName: "SpellCreate",
			Handler:    _SpellService_SpellCreate_Handler,
		},
		{
			MethodName: "SpellRead",
			Handler:    _SpellService_SpellRead_Handler,
		},
		{
			MethodName: "SpellUpdate",
			Handler:    _SpellService_SpellUpdate_Handler,
		},
		{
			MethodName: "SpellDelete",
			Handler:    _SpellService_SpellDelete_Handler,
		},
		{
			MethodName: "SpellPatch",
			Handler:    _SpellService_SpellPatch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/spell_service.proto",
}
