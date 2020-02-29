// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/trade_service.proto

package pb

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

type TradeSearchRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Limit                int64    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               int64    `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Orderby              string   `protobuf:"bytes,4,opt,name=orderby,proto3" json:"orderby,omitempty"`
	Orderdesc            bool     `protobuf:"varint,5,opt,name=orderdesc,proto3" json:"orderdesc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TradeSearchRequest) Reset()         { *m = TradeSearchRequest{} }
func (m *TradeSearchRequest) String() string { return proto.CompactTextString(m) }
func (*TradeSearchRequest) ProtoMessage()    {}
func (*TradeSearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3c9fc3a348d0932, []int{0}
}

func (m *TradeSearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TradeSearchRequest.Unmarshal(m, b)
}
func (m *TradeSearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TradeSearchRequest.Marshal(b, m, deterministic)
}
func (m *TradeSearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TradeSearchRequest.Merge(m, src)
}
func (m *TradeSearchRequest) XXX_Size() int {
	return xxx_messageInfo_TradeSearchRequest.Size(m)
}
func (m *TradeSearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TradeSearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TradeSearchRequest proto.InternalMessageInfo

func (m *TradeSearchRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TradeSearchRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *TradeSearchRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *TradeSearchRequest) GetOrderby() string {
	if m != nil {
		return m.Orderby
	}
	return ""
}

func (m *TradeSearchRequest) GetOrderdesc() bool {
	if m != nil {
		return m.Orderdesc
	}
	return false
}

type TradeSearchResponse struct {
	Trades               []*Trade `protobuf:"bytes,1,rep,name=Trades,proto3" json:"Trades,omitempty"`
	Total                int64    `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TradeSearchResponse) Reset()         { *m = TradeSearchResponse{} }
func (m *TradeSearchResponse) String() string { return proto.CompactTextString(m) }
func (*TradeSearchResponse) ProtoMessage()    {}
func (*TradeSearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3c9fc3a348d0932, []int{1}
}

func (m *TradeSearchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TradeSearchResponse.Unmarshal(m, b)
}
func (m *TradeSearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TradeSearchResponse.Marshal(b, m, deterministic)
}
func (m *TradeSearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TradeSearchResponse.Merge(m, src)
}
func (m *TradeSearchResponse) XXX_Size() int {
	return xxx_messageInfo_TradeSearchResponse.Size(m)
}
func (m *TradeSearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TradeSearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TradeSearchResponse proto.InternalMessageInfo

func (m *TradeSearchResponse) GetTrades() []*Trade {
	if m != nil {
		return m.Trades
	}
	return nil
}

func (m *TradeSearchResponse) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

type TradeCreateRequest struct {
	Values               map[string]string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TradeCreateRequest) Reset()         { *m = TradeCreateRequest{} }
func (m *TradeCreateRequest) String() string { return proto.CompactTextString(m) }
func (*TradeCreateRequest) ProtoMessage()    {}
func (*TradeCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3c9fc3a348d0932, []int{2}
}

func (m *TradeCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TradeCreateRequest.Unmarshal(m, b)
}
func (m *TradeCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TradeCreateRequest.Marshal(b, m, deterministic)
}
func (m *TradeCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TradeCreateRequest.Merge(m, src)
}
func (m *TradeCreateRequest) XXX_Size() int {
	return xxx_messageInfo_TradeCreateRequest.Size(m)
}
func (m *TradeCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TradeCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TradeCreateRequest proto.InternalMessageInfo

func (m *TradeCreateRequest) GetValues() map[string]string {
	if m != nil {
		return m.Values
	}
	return nil
}

type TradeCreateResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TradeCreateResponse) Reset()         { *m = TradeCreateResponse{} }
func (m *TradeCreateResponse) String() string { return proto.CompactTextString(m) }
func (*TradeCreateResponse) ProtoMessage()    {}
func (*TradeCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3c9fc3a348d0932, []int{3}
}

func (m *TradeCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TradeCreateResponse.Unmarshal(m, b)
}
func (m *TradeCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TradeCreateResponse.Marshal(b, m, deterministic)
}
func (m *TradeCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TradeCreateResponse.Merge(m, src)
}
func (m *TradeCreateResponse) XXX_Size() int {
	return xxx_messageInfo_TradeCreateResponse.Size(m)
}
func (m *TradeCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TradeCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TradeCreateResponse proto.InternalMessageInfo

func (m *TradeCreateResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type TradeReadRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TradeReadRequest) Reset()         { *m = TradeReadRequest{} }
func (m *TradeReadRequest) String() string { return proto.CompactTextString(m) }
func (*TradeReadRequest) ProtoMessage()    {}
func (*TradeReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3c9fc3a348d0932, []int{4}
}

func (m *TradeReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TradeReadRequest.Unmarshal(m, b)
}
func (m *TradeReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TradeReadRequest.Marshal(b, m, deterministic)
}
func (m *TradeReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TradeReadRequest.Merge(m, src)
}
func (m *TradeReadRequest) XXX_Size() int {
	return xxx_messageInfo_TradeReadRequest.Size(m)
}
func (m *TradeReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TradeReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TradeReadRequest proto.InternalMessageInfo

func (m *TradeReadRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type TradeReadResponse struct {
	Trade                *Trade   `protobuf:"bytes,1,opt,name=trade,proto3" json:"trade,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TradeReadResponse) Reset()         { *m = TradeReadResponse{} }
func (m *TradeReadResponse) String() string { return proto.CompactTextString(m) }
func (*TradeReadResponse) ProtoMessage()    {}
func (*TradeReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3c9fc3a348d0932, []int{5}
}

func (m *TradeReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TradeReadResponse.Unmarshal(m, b)
}
func (m *TradeReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TradeReadResponse.Marshal(b, m, deterministic)
}
func (m *TradeReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TradeReadResponse.Merge(m, src)
}
func (m *TradeReadResponse) XXX_Size() int {
	return xxx_messageInfo_TradeReadResponse.Size(m)
}
func (m *TradeReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TradeReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TradeReadResponse proto.InternalMessageInfo

func (m *TradeReadResponse) GetTrade() *Trade {
	if m != nil {
		return m.Trade
	}
	return nil
}

type TradeUpdateRequest struct {
	Id                   int64             `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Values               map[string]string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TradeUpdateRequest) Reset()         { *m = TradeUpdateRequest{} }
func (m *TradeUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*TradeUpdateRequest) ProtoMessage()    {}
func (*TradeUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3c9fc3a348d0932, []int{6}
}

func (m *TradeUpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TradeUpdateRequest.Unmarshal(m, b)
}
func (m *TradeUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TradeUpdateRequest.Marshal(b, m, deterministic)
}
func (m *TradeUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TradeUpdateRequest.Merge(m, src)
}
func (m *TradeUpdateRequest) XXX_Size() int {
	return xxx_messageInfo_TradeUpdateRequest.Size(m)
}
func (m *TradeUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TradeUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TradeUpdateRequest proto.InternalMessageInfo

func (m *TradeUpdateRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TradeUpdateRequest) GetValues() map[string]string {
	if m != nil {
		return m.Values
	}
	return nil
}

type TradeUpdateResponse struct {
	Rowsaffected         int64    `protobuf:"varint,1,opt,name=rowsaffected,proto3" json:"rowsaffected,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TradeUpdateResponse) Reset()         { *m = TradeUpdateResponse{} }
func (m *TradeUpdateResponse) String() string { return proto.CompactTextString(m) }
func (*TradeUpdateResponse) ProtoMessage()    {}
func (*TradeUpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3c9fc3a348d0932, []int{7}
}

func (m *TradeUpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TradeUpdateResponse.Unmarshal(m, b)
}
func (m *TradeUpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TradeUpdateResponse.Marshal(b, m, deterministic)
}
func (m *TradeUpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TradeUpdateResponse.Merge(m, src)
}
func (m *TradeUpdateResponse) XXX_Size() int {
	return xxx_messageInfo_TradeUpdateResponse.Size(m)
}
func (m *TradeUpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TradeUpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TradeUpdateResponse proto.InternalMessageInfo

func (m *TradeUpdateResponse) GetRowsaffected() int64 {
	if m != nil {
		return m.Rowsaffected
	}
	return 0
}

type TradeDeleteRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TradeDeleteRequest) Reset()         { *m = TradeDeleteRequest{} }
func (m *TradeDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*TradeDeleteRequest) ProtoMessage()    {}
func (*TradeDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3c9fc3a348d0932, []int{8}
}

func (m *TradeDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TradeDeleteRequest.Unmarshal(m, b)
}
func (m *TradeDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TradeDeleteRequest.Marshal(b, m, deterministic)
}
func (m *TradeDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TradeDeleteRequest.Merge(m, src)
}
func (m *TradeDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_TradeDeleteRequest.Size(m)
}
func (m *TradeDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TradeDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TradeDeleteRequest proto.InternalMessageInfo

func (m *TradeDeleteRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type TradeDeleteResponse struct {
	Rowsaffected         int64    `protobuf:"varint,1,opt,name=rowsaffected,proto3" json:"rowsaffected,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TradeDeleteResponse) Reset()         { *m = TradeDeleteResponse{} }
func (m *TradeDeleteResponse) String() string { return proto.CompactTextString(m) }
func (*TradeDeleteResponse) ProtoMessage()    {}
func (*TradeDeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3c9fc3a348d0932, []int{9}
}

func (m *TradeDeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TradeDeleteResponse.Unmarshal(m, b)
}
func (m *TradeDeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TradeDeleteResponse.Marshal(b, m, deterministic)
}
func (m *TradeDeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TradeDeleteResponse.Merge(m, src)
}
func (m *TradeDeleteResponse) XXX_Size() int {
	return xxx_messageInfo_TradeDeleteResponse.Size(m)
}
func (m *TradeDeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TradeDeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TradeDeleteResponse proto.InternalMessageInfo

func (m *TradeDeleteResponse) GetRowsaffected() int64 {
	if m != nil {
		return m.Rowsaffected
	}
	return 0
}

type TradePatchRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TradePatchRequest) Reset()         { *m = TradePatchRequest{} }
func (m *TradePatchRequest) String() string { return proto.CompactTextString(m) }
func (*TradePatchRequest) ProtoMessage()    {}
func (*TradePatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3c9fc3a348d0932, []int{10}
}

func (m *TradePatchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TradePatchRequest.Unmarshal(m, b)
}
func (m *TradePatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TradePatchRequest.Marshal(b, m, deterministic)
}
func (m *TradePatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TradePatchRequest.Merge(m, src)
}
func (m *TradePatchRequest) XXX_Size() int {
	return xxx_messageInfo_TradePatchRequest.Size(m)
}
func (m *TradePatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TradePatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TradePatchRequest proto.InternalMessageInfo

func (m *TradePatchRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TradePatchRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *TradePatchRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type TradePatchResponse struct {
	Rowsaffected         int64    `protobuf:"varint,1,opt,name=rowsaffected,proto3" json:"rowsaffected,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TradePatchResponse) Reset()         { *m = TradePatchResponse{} }
func (m *TradePatchResponse) String() string { return proto.CompactTextString(m) }
func (*TradePatchResponse) ProtoMessage()    {}
func (*TradePatchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3c9fc3a348d0932, []int{11}
}

func (m *TradePatchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TradePatchResponse.Unmarshal(m, b)
}
func (m *TradePatchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TradePatchResponse.Marshal(b, m, deterministic)
}
func (m *TradePatchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TradePatchResponse.Merge(m, src)
}
func (m *TradePatchResponse) XXX_Size() int {
	return xxx_messageInfo_TradePatchResponse.Size(m)
}
func (m *TradePatchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TradePatchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TradePatchResponse proto.InternalMessageInfo

func (m *TradePatchResponse) GetRowsaffected() int64 {
	if m != nil {
		return m.Rowsaffected
	}
	return 0
}

func init() {
	proto.RegisterType((*TradeSearchRequest)(nil), "pb.TradeSearchRequest")
	proto.RegisterType((*TradeSearchResponse)(nil), "pb.TradeSearchResponse")
	proto.RegisterType((*TradeCreateRequest)(nil), "pb.TradeCreateRequest")
	proto.RegisterMapType((map[string]string)(nil), "pb.TradeCreateRequest.ValuesEntry")
	proto.RegisterType((*TradeCreateResponse)(nil), "pb.TradeCreateResponse")
	proto.RegisterType((*TradeReadRequest)(nil), "pb.TradeReadRequest")
	proto.RegisterType((*TradeReadResponse)(nil), "pb.TradeReadResponse")
	proto.RegisterType((*TradeUpdateRequest)(nil), "pb.TradeUpdateRequest")
	proto.RegisterMapType((map[string]string)(nil), "pb.TradeUpdateRequest.ValuesEntry")
	proto.RegisterType((*TradeUpdateResponse)(nil), "pb.TradeUpdateResponse")
	proto.RegisterType((*TradeDeleteRequest)(nil), "pb.TradeDeleteRequest")
	proto.RegisterType((*TradeDeleteResponse)(nil), "pb.TradeDeleteResponse")
	proto.RegisterType((*TradePatchRequest)(nil), "pb.TradePatchRequest")
	proto.RegisterType((*TradePatchResponse)(nil), "pb.TradePatchResponse")
}

func init() { proto.RegisterFile("proto/trade_service.proto", fileDescriptor_b3c9fc3a348d0932) }

var fileDescriptor_b3c9fc3a348d0932 = []byte{
	// 606 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0xed, 0x24, 0xe0, 0x49, 0x15, 0xa5, 0xdb, 0x90, 0x3a, 0x56, 0xa5, 0x86, 0x15, 0x48,
	0x51, 0x0f, 0xb1, 0x08, 0x1c, 0xda, 0x5c, 0x81, 0x13, 0x12, 0xaa, 0x4c, 0x89, 0xc4, 0xa9, 0xda,
	0xc4, 0x9b, 0x62, 0x91, 0xda, 0xc6, 0xde, 0x06, 0x45, 0x55, 0x2f, 0x5c, 0x91, 0xb8, 0x70, 0xe4,
	0xc6, 0x5f, 0xe2, 0x2f, 0xf0, 0x43, 0x90, 0xf7, 0x23, 0xde, 0x75, 0x01, 0x81, 0xc4, 0xcd, 0xf3,
	0x32, 0x7e, 0xf3, 0xde, 0xe4, 0x8d, 0x61, 0x90, 0xe5, 0x29, 0x4b, 0x03, 0x96, 0x93, 0x88, 0x9e,
	0x17, 0x34, 0x5f, 0xc7, 0x0b, 0x3a, 0xe6, 0x18, 0xb2, 0xb3, 0xb9, 0x7f, 0x70, 0x91, 0xa6, 0x17,
	0x2b, 0x1a, 0x90, 0x2c, 0x0e, 0x48, 0x92, 0xa4, 0x8c, 0xb0, 0x38, 0x4d, 0x0a, 0xd1, 0xe1, 0xef,
	0x6a, 0x2f, 0x0b, 0x08, 0x7f, 0xb6, 0x00, 0x9d, 0x95, 0xf5, 0x2b, 0x4a, 0xf2, 0xc5, 0xdb, 0x90,
	0xbe, 0xbf, 0xa2, 0x05, 0x43, 0x08, 0x1a, 0x09, 0xb9, 0xa4, 0x9e, 0x35, 0xb4, 0x46, 0x6e, 0xc8,
	0x9f, 0x51, 0x0f, 0x9a, 0xab, 0xf8, 0x32, 0x66, 0x9e, 0x3d, 0xb4, 0x46, 0x4e, 0x28, 0x0a, 0xd4,
	0x87, 0x56, 0xba, 0x5c, 0x16, 0x94, 0x79, 0x0e, 0x87, 0x65, 0x85, 0x3c, 0xb8, 0x93, 0xe6, 0x11,
	0xcd, 0xe7, 0x1b, 0xaf, 0xc1, 0x49, 0x54, 0x89, 0x0e, 0xc0, 0xe5, 0x8f, 0x11, 0x2d, 0x16, 0x5e,
	0x73, 0x68, 0x8d, 0xee, 0x86, 0x15, 0x80, 0x5f, 0xc2, 0x9e, 0xa1, 0xa7, 0xc8, 0xd2, 0xa4, 0xa0,
	0xe8, 0x3e, 0xb4, 0x38, 0x5c, 0x78, 0xd6, 0xd0, 0x19, 0xb5, 0x27, 0xee, 0x38, 0x9b, 0x8f, 0x39,
	0x12, 0xca, 0x1f, 0x4a, 0x7d, 0x2c, 0x65, 0x64, 0xa5, 0xf4, 0xf1, 0x02, 0x7f, 0x52, 0x06, 0x9f,
	0xe6, 0x94, 0x30, 0xaa, 0x0c, 0x4e, 0xa1, 0xb5, 0x26, 0xab, 0x2b, 0x5a, 0x78, 0x36, 0xe7, 0xc3,
	0x5b, 0x3e, 0xa3, 0x6f, 0x3c, 0xe3, 0x4d, 0xcf, 0x13, 0x96, 0x6f, 0x42, 0xf9, 0x86, 0x7f, 0x02,
	0x6d, 0x0d, 0x46, 0x5d, 0x70, 0xde, 0xd1, 0x8d, 0x5c, 0x55, 0xf9, 0x58, 0x2a, 0xe1, 0xad, 0x5c,
	0x89, 0x1b, 0x8a, 0x62, 0x6a, 0x1f, 0x5b, 0xf8, 0xa1, 0x74, 0xa7, 0x86, 0x48, 0x77, 0x1d, 0xb0,
	0xe3, 0x88, 0x33, 0x38, 0xa1, 0x1d, 0x47, 0x18, 0x43, 0x57, 0x78, 0xa3, 0x24, 0x52, 0x8a, 0xeb,
	0x3d, 0x4f, 0x60, 0x57, 0xeb, 0x91, 0x44, 0x87, 0xd0, 0xe4, 0xff, 0x2e, 0xef, 0x33, 0xb6, 0x24,
	0x70, 0xfc, 0x55, 0xad, 0xe3, 0x75, 0x16, 0x69, 0xeb, 0xa8, 0x91, 0xff, 0x61, 0x3d, 0xc6, 0x7b,
	0xff, 0x7b, 0x3d, 0x27, 0x72, 0x3d, 0x6a, 0x88, 0x74, 0x85, 0x61, 0x27, 0x4f, 0x3f, 0x14, 0x64,
	0xb9, 0xa4, 0x0b, 0x46, 0x95, 0x4e, 0x03, 0xc3, 0x0f, 0xa4, 0xaf, 0x67, 0x74, 0x45, 0x7f, 0xeb,
	0x6b, 0x3b, 0x40, 0x75, 0xfd, 0xc3, 0x80, 0x17, 0x72, 0xdf, 0xa7, 0x84, 0x55, 0x77, 0x52, 0xdf,
	0x9b, 0x34, 0x6b, 0xff, 0xc2, 0xac, 0xa3, 0x99, 0xc5, 0xc7, 0x52, 0xad, 0x24, 0xfb, 0x7b, 0x19,
	0x93, 0x6f, 0x0d, 0xd8, 0x91, 0x07, 0xc2, 0x8f, 0x1f, 0x9d, 0x43, 0x5b, 0x3b, 0x18, 0xd4, 0xdf,
	0xfe, 0x53, 0xc6, 0x45, 0xfb, 0xfb, 0xb7, 0x70, 0x31, 0x14, 0x1f, 0x7e, 0xfc, 0xfe, 0xe3, 0x8b,
	0x3d, 0x40, 0xfb, 0xc1, 0xfa, 0x91, 0xf8, 0x34, 0x04, 0x05, 0xef, 0x08, 0xae, 0xcb, 0xb3, 0xbf,
	0x41, 0x67, 0x72, 0x80, 0xc8, 0xac, 0x36, 0xc0, 0xb8, 0x14, 0x6d, 0x80, 0x19, 0x6e, 0xdc, 0xe3,
	0x03, 0x3a, 0xd8, 0xdd, 0x0e, 0x98, 0x5a, 0x47, 0xe8, 0x14, 0xdc, 0x6d, 0x7c, 0x51, 0xaf, 0xca,
	0x69, 0x95, 0x78, 0xff, 0x5e, 0x0d, 0x95, 0x7c, 0x7d, 0xce, 0xd7, 0x45, 0x9d, 0x4a, 0xf0, 0x75,
	0x1c, 0xdd, 0xa0, 0x37, 0x52, 0xa7, 0x08, 0x8f, 0xa6, 0xd3, 0x88, 0xac, 0xa6, 0xd3, 0x4c, 0x19,
	0x1e, 0x70, 0xde, 0x3d, 0xbf, 0xc6, 0x5b, 0x8a, 0x9d, 0x49, 0x6a, 0x11, 0x1b, 0x8d, 0xda, 0x48,
	0x9b, 0x46, 0x6d, 0xe6, 0x4b, 0x49, 0x3e, 0xaa, 0x4b, 0x9e, 0x01, 0x54, 0x31, 0x40, 0x95, 0x5f,
	0x3d, 0x63, 0x7e, 0xbf, 0x0e, 0x9b, 0x7a, 0x27, 0xb7, 0xf5, 0xce, 0x5b, 0xfc, 0xe3, 0xfe, 0xf8,
	0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x59, 0x93, 0xe9, 0x2e, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TradeServiceClient is the client API for TradeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TradeServiceClient interface {
	TradeSearch(ctx context.Context, in *TradeSearchRequest, opts ...grpc.CallOption) (*TradeSearchResponse, error)
	TradeCreate(ctx context.Context, in *TradeCreateRequest, opts ...grpc.CallOption) (*TradeCreateResponse, error)
	TradeRead(ctx context.Context, in *TradeReadRequest, opts ...grpc.CallOption) (*TradeReadResponse, error)
	TradeUpdate(ctx context.Context, in *TradeUpdateRequest, opts ...grpc.CallOption) (*TradeUpdateResponse, error)
	TradeDelete(ctx context.Context, in *TradeDeleteRequest, opts ...grpc.CallOption) (*TradeDeleteResponse, error)
	TradePatch(ctx context.Context, in *TradePatchRequest, opts ...grpc.CallOption) (*TradePatchResponse, error)
}

type tradeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTradeServiceClient(cc grpc.ClientConnInterface) TradeServiceClient {
	return &tradeServiceClient{cc}
}

func (c *tradeServiceClient) TradeSearch(ctx context.Context, in *TradeSearchRequest, opts ...grpc.CallOption) (*TradeSearchResponse, error) {
	out := new(TradeSearchResponse)
	err := c.cc.Invoke(ctx, "/pb.TradeService/TradeSearch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeServiceClient) TradeCreate(ctx context.Context, in *TradeCreateRequest, opts ...grpc.CallOption) (*TradeCreateResponse, error) {
	out := new(TradeCreateResponse)
	err := c.cc.Invoke(ctx, "/pb.TradeService/TradeCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeServiceClient) TradeRead(ctx context.Context, in *TradeReadRequest, opts ...grpc.CallOption) (*TradeReadResponse, error) {
	out := new(TradeReadResponse)
	err := c.cc.Invoke(ctx, "/pb.TradeService/TradeRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeServiceClient) TradeUpdate(ctx context.Context, in *TradeUpdateRequest, opts ...grpc.CallOption) (*TradeUpdateResponse, error) {
	out := new(TradeUpdateResponse)
	err := c.cc.Invoke(ctx, "/pb.TradeService/TradeUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeServiceClient) TradeDelete(ctx context.Context, in *TradeDeleteRequest, opts ...grpc.CallOption) (*TradeDeleteResponse, error) {
	out := new(TradeDeleteResponse)
	err := c.cc.Invoke(ctx, "/pb.TradeService/TradeDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeServiceClient) TradePatch(ctx context.Context, in *TradePatchRequest, opts ...grpc.CallOption) (*TradePatchResponse, error) {
	out := new(TradePatchResponse)
	err := c.cc.Invoke(ctx, "/pb.TradeService/TradePatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TradeServiceServer is the server API for TradeService service.
type TradeServiceServer interface {
	TradeSearch(context.Context, *TradeSearchRequest) (*TradeSearchResponse, error)
	TradeCreate(context.Context, *TradeCreateRequest) (*TradeCreateResponse, error)
	TradeRead(context.Context, *TradeReadRequest) (*TradeReadResponse, error)
	TradeUpdate(context.Context, *TradeUpdateRequest) (*TradeUpdateResponse, error)
	TradeDelete(context.Context, *TradeDeleteRequest) (*TradeDeleteResponse, error)
	TradePatch(context.Context, *TradePatchRequest) (*TradePatchResponse, error)
}

// UnimplementedTradeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTradeServiceServer struct {
}

func (*UnimplementedTradeServiceServer) TradeSearch(ctx context.Context, req *TradeSearchRequest) (*TradeSearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TradeSearch not implemented")
}
func (*UnimplementedTradeServiceServer) TradeCreate(ctx context.Context, req *TradeCreateRequest) (*TradeCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TradeCreate not implemented")
}
func (*UnimplementedTradeServiceServer) TradeRead(ctx context.Context, req *TradeReadRequest) (*TradeReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TradeRead not implemented")
}
func (*UnimplementedTradeServiceServer) TradeUpdate(ctx context.Context, req *TradeUpdateRequest) (*TradeUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TradeUpdate not implemented")
}
func (*UnimplementedTradeServiceServer) TradeDelete(ctx context.Context, req *TradeDeleteRequest) (*TradeDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TradeDelete not implemented")
}
func (*UnimplementedTradeServiceServer) TradePatch(ctx context.Context, req *TradePatchRequest) (*TradePatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TradePatch not implemented")
}

func RegisterTradeServiceServer(s *grpc.Server, srv TradeServiceServer) {
	s.RegisterService(&_TradeService_serviceDesc, srv)
}

func _TradeService_TradeSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TradeSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeServiceServer).TradeSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TradeService/TradeSearch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeServiceServer).TradeSearch(ctx, req.(*TradeSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeService_TradeCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TradeCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeServiceServer).TradeCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TradeService/TradeCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeServiceServer).TradeCreate(ctx, req.(*TradeCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeService_TradeRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TradeReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeServiceServer).TradeRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TradeService/TradeRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeServiceServer).TradeRead(ctx, req.(*TradeReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeService_TradeUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TradeUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeServiceServer).TradeUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TradeService/TradeUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeServiceServer).TradeUpdate(ctx, req.(*TradeUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeService_TradeDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TradeDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeServiceServer).TradeDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TradeService/TradeDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeServiceServer).TradeDelete(ctx, req.(*TradeDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeService_TradePatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TradePatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeServiceServer).TradePatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TradeService/TradePatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeServiceServer).TradePatch(ctx, req.(*TradePatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TradeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.TradeService",
	HandlerType: (*TradeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TradeSearch",
			Handler:    _TradeService_TradeSearch_Handler,
		},
		{
			MethodName: "TradeCreate",
			Handler:    _TradeService_TradeCreate_Handler,
		},
		{
			MethodName: "TradeRead",
			Handler:    _TradeService_TradeRead_Handler,
		},
		{
			MethodName: "TradeUpdate",
			Handler:    _TradeService_TradeUpdate_Handler,
		},
		{
			MethodName: "TradeDelete",
			Handler:    _TradeService_TradeDelete_Handler,
		},
		{
			MethodName: "TradePatch",
			Handler:    _TradeService_TradePatch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/trade_service.proto",
}
