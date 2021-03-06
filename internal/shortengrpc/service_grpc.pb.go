// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: internal/shortengrpc/service.proto

package shortengrpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ShortenGrpcClient is the client API for ShortenGrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShortenGrpcClient interface {
	GetShort(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*OriginResponse, error)
	GetURLs(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*ListResponse, error)
	PostBatch(ctx context.Context, in *BatchRequest, opts ...grpc.CallOption) (*BatchResponse, error)
	PostURL(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*ShortResponse, error)
	SoftDelete(ctx context.Context, in *DelRequest, opts ...grpc.CallOption) (*DelResponse, error)
}

type shortenGrpcClient struct {
	cc grpc.ClientConnInterface
}

func NewShortenGrpcClient(cc grpc.ClientConnInterface) ShortenGrpcClient {
	return &shortenGrpcClient{cc}
}

func (c *shortenGrpcClient) GetShort(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*OriginResponse, error) {
	out := new(OriginResponse)
	err := c.cc.Invoke(ctx, "/shortengrpc.ShortenGrpc/GetShort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortenGrpcClient) GetURLs(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/shortengrpc.ShortenGrpc/GetURLs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortenGrpcClient) PostBatch(ctx context.Context, in *BatchRequest, opts ...grpc.CallOption) (*BatchResponse, error) {
	out := new(BatchResponse)
	err := c.cc.Invoke(ctx, "/shortengrpc.ShortenGrpc/PostBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortenGrpcClient) PostURL(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*ShortResponse, error) {
	out := new(ShortResponse)
	err := c.cc.Invoke(ctx, "/shortengrpc.ShortenGrpc/PostURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortenGrpcClient) SoftDelete(ctx context.Context, in *DelRequest, opts ...grpc.CallOption) (*DelResponse, error) {
	out := new(DelResponse)
	err := c.cc.Invoke(ctx, "/shortengrpc.ShortenGrpc/SoftDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShortenGrpcServer is the grpcserver API for ShortenGrpc service.
// All implementations must embed UnimplementedShortenGrpcServer
// for forward compatibility
type ShortenGrpcServer interface {
	GetShort(context.Context, *GetRequest) (*OriginResponse, error)
	GetURLs(context.Context, *UUID) (*ListResponse, error)
	PostBatch(context.Context, *BatchRequest) (*BatchResponse, error)
	PostURL(context.Context, *PostRequest) (*ShortResponse, error)
	SoftDelete(context.Context, *DelRequest) (*DelResponse, error)
	mustEmbedUnimplementedShortenGrpcServer()
}

// UnimplementedShortenGrpcServer must be embedded to have forward compatible implementations.
type UnimplementedShortenGrpcServer struct {
}

func (UnimplementedShortenGrpcServer) GetShort(context.Context, *GetRequest) (*OriginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShort not implemented")
}
func (UnimplementedShortenGrpcServer) GetURLs(context.Context, *UUID) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetURLs not implemented")
}
func (UnimplementedShortenGrpcServer) PostBatch(context.Context, *BatchRequest) (*BatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostBatch not implemented")
}
func (UnimplementedShortenGrpcServer) PostURL(context.Context, *PostRequest) (*ShortResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostURL not implemented")
}
func (UnimplementedShortenGrpcServer) SoftDelete(context.Context, *DelRequest) (*DelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SoftDelete not implemented")
}
func (UnimplementedShortenGrpcServer) mustEmbedUnimplementedShortenGrpcServer() {}

// UnsafeShortenGrpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShortenGrpcServer will
// result in compilation errors.
type UnsafeShortenGrpcServer interface {
	mustEmbedUnimplementedShortenGrpcServer()
}

func RegisterShortenGrpcServer(s grpc.ServiceRegistrar, srv ShortenGrpcServer) {
	s.RegisterService(&ShortenGrpc_ServiceDesc, srv)
}

func _ShortenGrpc_GetShort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenGrpcServer).GetShort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortengrpc.ShortenGrpc/GetShort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenGrpcServer).GetShort(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShortenGrpc_GetURLs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenGrpcServer).GetURLs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortengrpc.ShortenGrpc/GetURLs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenGrpcServer).GetURLs(ctx, req.(*UUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShortenGrpc_PostBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenGrpcServer).PostBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortengrpc.ShortenGrpc/PostBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenGrpcServer).PostBatch(ctx, req.(*BatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShortenGrpc_PostURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenGrpcServer).PostURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortengrpc.ShortenGrpc/PostURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenGrpcServer).PostURL(ctx, req.(*PostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShortenGrpc_SoftDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenGrpcServer).SoftDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortengrpc.ShortenGrpc/SoftDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenGrpcServer).SoftDelete(ctx, req.(*DelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ShortenGrpc_ServiceDesc is the grpc.ServiceDesc for ShortenGrpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShortenGrpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shortengrpc.ShortenGrpc",
	HandlerType: (*ShortenGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetShort",
			Handler:    _ShortenGrpc_GetShort_Handler,
		},
		{
			MethodName: "GetURLs",
			Handler:    _ShortenGrpc_GetURLs_Handler,
		},
		{
			MethodName: "PostBatch",
			Handler:    _ShortenGrpc_PostBatch_Handler,
		},
		{
			MethodName: "PostURL",
			Handler:    _ShortenGrpc_PostURL_Handler,
		},
		{
			MethodName: "SoftDelete",
			Handler:    _ShortenGrpc_SoftDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/shortengrpc/service.proto",
}
