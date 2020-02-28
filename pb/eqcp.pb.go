// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/eqcp.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

func init() { proto.RegisterFile("proto/eqcp.proto", fileDescriptor_01c2dd79b7d9ec64) }

var fileDescriptor_01c2dd79b7d9ec64 = []byte{
	// 413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xdf, 0x8a, 0xd3, 0x40,
	0x14, 0xc6, 0x71, 0xfd, 0x53, 0x0d, 0x8b, 0xe8, 0xd8, 0x25, 0x10, 0x14, 0x43, 0x2f, 0x8b, 0x9b,
	0xd9, 0xad, 0x77, 0x7b, 0xb7, 0xd6, 0x82, 0x82, 0x84, 0x74, 0x65, 0xf1, 0x52, 0x26, 0x93, 0x93,
	0x64, 0x30, 0x9d, 0x39, 0x3b, 0x33, 0x59, 0x2d, 0xa5, 0x37, 0x3e, 0x82, 0x3e, 0x9a, 0xaf, 0xe0,
	0x95, 0x4f, 0x21, 0x9d, 0x24, 0xa5, 0x09, 0x7b, 0x95, 0x7c, 0xdf, 0x9c, 0xfc, 0xe6, 0x17, 0x38,
	0xde, 0x33, 0xd4, 0xca, 0x2a, 0x0a, 0x37, 0x1c, 0x23, 0xf7, 0x4a, 0x8e, 0x30, 0x0d, 0x5e, 0x16,
	0x4a, 0x15, 0x15, 0x50, 0x86, 0x82, 0x32, 0x29, 0x95, 0x65, 0x56, 0x28, 0x69, 0x9a, 0x89, 0xe0,
	0x8d, 0x7b, 0xf0, 0xd3, 0x02, 0xe4, 0xa9, 0xf9, 0xce, 0x8a, 0x02, 0x34, 0x55, 0xe8, 0x26, 0xee,
	0x98, 0xf6, 0x9b, 0x1b, 0x24, 0xf2, 0xaf, 0x06, 0xf4, 0xad, 0xe0, 0xd0, 0x1c, 0xcc, 0xfe, 0xdd,
	0xf7, 0x1e, 0x2c, 0x96, 0xf3, 0x84, 0x7c, 0xf1, 0x9e, 0xc4, 0xc8, 0x3f, 0x03, 0xd3, 0xbc, 0x24,
	0xe3, 0x08, 0xd3, 0x68, 0x1f, 0xaf, 0xe0, 0xa6, 0x06, 0x63, 0x83, 0x93, 0x41, 0x6b, 0x50, 0x49,
	0x03, 0x93, 0x57, 0x3f, 0xff, 0xfc, 0xfd, 0x7d, 0xe4, 0x93, 0x13, 0x7a, 0x7b, 0xbe, 0xbb, 0x81,
	0x1a, 0x77, 0x4e, 0x37, 0x92, 0xad, 0x60, 0x4b, 0x3e, 0x39, 0xf0, 0x5c, 0x03, 0xb3, 0xb0, 0x07,
	0x37, 0x71, 0x08, 0xee, 0xda, 0x16, 0x4c, 0x1c, 0xf8, 0x78, 0x32, 0x6a, 0xc1, 0x17, 0xf7, 0xa6,
	0xe4, 0x83, 0x37, 0x8a, 0x91, 0x5f, 0x01, 0xcb, 0x08, 0x69, 0xbf, 0xda, 0x85, 0x8e, 0xf4, 0xa2,
	0xd7, 0xb5, 0x9c, 0xb1, 0xe3, 0x3c, 0x25, 0xc7, 0x9d, 0xe0, 0x46, 0x64, 0x5b, 0xb2, 0x74, 0x5e,
	0xd7, 0x98, 0x1d, 0x7a, 0x35, 0x71, 0xe8, 0xd5, 0xb5, 0x2d, 0xcf, 0x77, 0xbc, 0xe7, 0x41, 0x8f,
	0xb7, 0x93, 0x8b, 0x1d, 0xf2, 0x3d, 0x54, 0x70, 0x80, 0x6c, 0xe2, 0x10, 0xd9, 0xb5, 0x7d, 0xc5,
	0x69, 0x5f, 0x31, 0xf6, 0x1e, 0xc7, 0xc8, 0x13, 0x66, 0x79, 0x49, 0xba, 0x3f, 0x73, 0xa9, 0xa3,
	0x8d, 0xfb, 0x65, 0xdf, 0x6f, 0x36, 0xf4, 0x7b, 0x67, 0x7e, 0x5d, 0x22, 0x91, 0x9e, 0xbf, 0x58,
	0x2e, 0x56, 0x75, 0x38, 0x57, 0xd2, 0x6a, 0x55, 0x85, 0x09, 0x93, 0x50, 0x85, 0x97, 0xc9, 0x47,
	0x12, 0x5e, 0x1b, 0xc8, 0xc2, 0x5c, 0xe9, 0x50, 0x64, 0x20, 0xad, 0xc8, 0xd7, 0x42, 0x16, 0x61,
	0x33, 0x6d, 0x6c, 0x9d, 0xe7, 0x93, 0xa9, 0x37, 0x32, 0x35, 0xa2, 0xd2, 0x96, 0xbc, 0x2e, 0xad,
	0x45, 0x73, 0x41, 0x69, 0x21, 0x6c, 0x59, 0xa7, 0x11, 0x57, 0x2b, 0xfa, 0x83, 0xf1, 0x6f, 0xa0,
	0xd7, 0x6e, 0x93, 0x67, 0x0f, 0xcf, 0xa2, 0xb3, 0xe8, 0x3c, 0x7d, 0xe4, 0x16, 0xed, 0xed, 0xff,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x98, 0x18, 0x5e, 0xe5, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EQCPClient is the client API for EQCP service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EQCPClient interface {
	NpcSearch(ctx context.Context, in *NpcSearchRequest, opts ...grpc.CallOption) (*NpcSearchResponse, error)
	NpcCreate(ctx context.Context, in *NpcCreateRequest, opts ...grpc.CallOption) (*NpcCreateResponse, error)
	NpcRead(ctx context.Context, in *NpcReadRequest, opts ...grpc.CallOption) (*NpcReadResponse, error)
	NpcUpdate(ctx context.Context, in *NpcUpdateRequest, opts ...grpc.CallOption) (*NpcUpdateResponse, error)
	NpcDelete(ctx context.Context, in *NpcDeleteRequest, opts ...grpc.CallOption) (*NpcDeleteResponse, error)
	NpcPatch(ctx context.Context, in *NpcPatchRequest, opts ...grpc.CallOption) (*NpcPatchResponse, error)
}

type eQCPClient struct {
	cc grpc.ClientConnInterface
}

func NewEQCPClient(cc grpc.ClientConnInterface) EQCPClient {
	return &eQCPClient{cc}
}

func (c *eQCPClient) NpcSearch(ctx context.Context, in *NpcSearchRequest, opts ...grpc.CallOption) (*NpcSearchResponse, error) {
	out := new(NpcSearchResponse)
	err := c.cc.Invoke(ctx, "/pb.EQCP/NpcSearch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eQCPClient) NpcCreate(ctx context.Context, in *NpcCreateRequest, opts ...grpc.CallOption) (*NpcCreateResponse, error) {
	out := new(NpcCreateResponse)
	err := c.cc.Invoke(ctx, "/pb.EQCP/NpcCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eQCPClient) NpcRead(ctx context.Context, in *NpcReadRequest, opts ...grpc.CallOption) (*NpcReadResponse, error) {
	out := new(NpcReadResponse)
	err := c.cc.Invoke(ctx, "/pb.EQCP/NpcRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eQCPClient) NpcUpdate(ctx context.Context, in *NpcUpdateRequest, opts ...grpc.CallOption) (*NpcUpdateResponse, error) {
	out := new(NpcUpdateResponse)
	err := c.cc.Invoke(ctx, "/pb.EQCP/NpcUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eQCPClient) NpcDelete(ctx context.Context, in *NpcDeleteRequest, opts ...grpc.CallOption) (*NpcDeleteResponse, error) {
	out := new(NpcDeleteResponse)
	err := c.cc.Invoke(ctx, "/pb.EQCP/NpcDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eQCPClient) NpcPatch(ctx context.Context, in *NpcPatchRequest, opts ...grpc.CallOption) (*NpcPatchResponse, error) {
	out := new(NpcPatchResponse)
	err := c.cc.Invoke(ctx, "/pb.EQCP/NpcPatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EQCPServer is the server API for EQCP service.
type EQCPServer interface {
	NpcSearch(context.Context, *NpcSearchRequest) (*NpcSearchResponse, error)
	NpcCreate(context.Context, *NpcCreateRequest) (*NpcCreateResponse, error)
	NpcRead(context.Context, *NpcReadRequest) (*NpcReadResponse, error)
	NpcUpdate(context.Context, *NpcUpdateRequest) (*NpcUpdateResponse, error)
	NpcDelete(context.Context, *NpcDeleteRequest) (*NpcDeleteResponse, error)
	NpcPatch(context.Context, *NpcPatchRequest) (*NpcPatchResponse, error)
}

// UnimplementedEQCPServer can be embedded to have forward compatible implementations.
type UnimplementedEQCPServer struct {
}

func (*UnimplementedEQCPServer) NpcSearch(ctx context.Context, req *NpcSearchRequest) (*NpcSearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NpcSearch not implemented")
}
func (*UnimplementedEQCPServer) NpcCreate(ctx context.Context, req *NpcCreateRequest) (*NpcCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NpcCreate not implemented")
}
func (*UnimplementedEQCPServer) NpcRead(ctx context.Context, req *NpcReadRequest) (*NpcReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NpcRead not implemented")
}
func (*UnimplementedEQCPServer) NpcUpdate(ctx context.Context, req *NpcUpdateRequest) (*NpcUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NpcUpdate not implemented")
}
func (*UnimplementedEQCPServer) NpcDelete(ctx context.Context, req *NpcDeleteRequest) (*NpcDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NpcDelete not implemented")
}
func (*UnimplementedEQCPServer) NpcPatch(ctx context.Context, req *NpcPatchRequest) (*NpcPatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NpcPatch not implemented")
}

func RegisterEQCPServer(s *grpc.Server, srv EQCPServer) {
	s.RegisterService(&_EQCP_serviceDesc, srv)
}

func _EQCP_NpcSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NpcSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EQCPServer).NpcSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.EQCP/NpcSearch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EQCPServer).NpcSearch(ctx, req.(*NpcSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EQCP_NpcCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NpcCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EQCPServer).NpcCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.EQCP/NpcCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EQCPServer).NpcCreate(ctx, req.(*NpcCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EQCP_NpcRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NpcReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EQCPServer).NpcRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.EQCP/NpcRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EQCPServer).NpcRead(ctx, req.(*NpcReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EQCP_NpcUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NpcUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EQCPServer).NpcUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.EQCP/NpcUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EQCPServer).NpcUpdate(ctx, req.(*NpcUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EQCP_NpcDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NpcDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EQCPServer).NpcDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.EQCP/NpcDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EQCPServer).NpcDelete(ctx, req.(*NpcDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EQCP_NpcPatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NpcPatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EQCPServer).NpcPatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.EQCP/NpcPatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EQCPServer).NpcPatch(ctx, req.(*NpcPatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EQCP_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.EQCP",
	HandlerType: (*EQCPServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NpcSearch",
			Handler:    _EQCP_NpcSearch_Handler,
		},
		{
			MethodName: "NpcCreate",
			Handler:    _EQCP_NpcCreate_Handler,
		},
		{
			MethodName: "NpcRead",
			Handler:    _EQCP_NpcRead_Handler,
		},
		{
			MethodName: "NpcUpdate",
			Handler:    _EQCP_NpcUpdate_Handler,
		},
		{
			MethodName: "NpcDelete",
			Handler:    _EQCP_NpcDelete_Handler,
		},
		{
			MethodName: "NpcPatch",
			Handler:    _EQCP_NpcPatch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/eqcp.proto",
}