// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: gfs.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	MasterService_HeartBeat_FullMethodName = "/MasterService/HeartBeat"
)

// MasterServiceClient is the client API for MasterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MasterServiceClient interface {
	HeartBeat(ctx context.Context, in *HeartBeatRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type masterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMasterServiceClient(cc grpc.ClientConnInterface) MasterServiceClient {
	return &masterServiceClient{cc}
}

func (c *masterServiceClient) HeartBeat(ctx context.Context, in *HeartBeatRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MasterService_HeartBeat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MasterServiceServer is the server API for MasterService service.
// All implementations must embed UnimplementedMasterServiceServer
// for forward compatibility.
type MasterServiceServer interface {
	HeartBeat(context.Context, *HeartBeatRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedMasterServiceServer()
}

// UnimplementedMasterServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMasterServiceServer struct{}

func (UnimplementedMasterServiceServer) HeartBeat(context.Context, *HeartBeatRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HeartBeat not implemented")
}
func (UnimplementedMasterServiceServer) mustEmbedUnimplementedMasterServiceServer() {}
func (UnimplementedMasterServiceServer) testEmbeddedByValue()                       {}

// UnsafeMasterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MasterServiceServer will
// result in compilation errors.
type UnsafeMasterServiceServer interface {
	mustEmbedUnimplementedMasterServiceServer()
}

func RegisterMasterServiceServer(s grpc.ServiceRegistrar, srv MasterServiceServer) {
	// If the following call pancis, it indicates UnimplementedMasterServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MasterService_ServiceDesc, srv)
}

func _MasterService_HeartBeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartBeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).HeartBeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MasterService_HeartBeat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).HeartBeat(ctx, req.(*HeartBeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MasterService_ServiceDesc is the grpc.ServiceDesc for MasterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MasterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MasterService",
	HandlerType: (*MasterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HeartBeat",
			Handler:    _MasterService_HeartBeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gfs.proto",
}

const (
	ChunkServerService_StoreChunk_FullMethodName = "/ChunkServerService/StoreChunk"
	ChunkServerService_CopyChunk_FullMethodName  = "/ChunkServerService/CopyChunk"
)

// ChunkServerServiceClient is the client API for ChunkServerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChunkServerServiceClient interface {
	StoreChunk(ctx context.Context, in *ChunkRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CopyChunk(ctx context.Context, in *CopyChunkRequest, opts ...grpc.CallOption) (*CopyChunkResponse, error)
}

type chunkServerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChunkServerServiceClient(cc grpc.ClientConnInterface) ChunkServerServiceClient {
	return &chunkServerServiceClient{cc}
}

func (c *chunkServerServiceClient) StoreChunk(ctx context.Context, in *ChunkRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ChunkServerService_StoreChunk_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chunkServerServiceClient) CopyChunk(ctx context.Context, in *CopyChunkRequest, opts ...grpc.CallOption) (*CopyChunkResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CopyChunkResponse)
	err := c.cc.Invoke(ctx, ChunkServerService_CopyChunk_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChunkServerServiceServer is the server API for ChunkServerService service.
// All implementations must embed UnimplementedChunkServerServiceServer
// for forward compatibility.
type ChunkServerServiceServer interface {
	StoreChunk(context.Context, *ChunkRequest) (*emptypb.Empty, error)
	CopyChunk(context.Context, *CopyChunkRequest) (*CopyChunkResponse, error)
	mustEmbedUnimplementedChunkServerServiceServer()
}

// UnimplementedChunkServerServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedChunkServerServiceServer struct{}

func (UnimplementedChunkServerServiceServer) StoreChunk(context.Context, *ChunkRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreChunk not implemented")
}
func (UnimplementedChunkServerServiceServer) CopyChunk(context.Context, *CopyChunkRequest) (*CopyChunkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CopyChunk not implemented")
}
func (UnimplementedChunkServerServiceServer) mustEmbedUnimplementedChunkServerServiceServer() {}
func (UnimplementedChunkServerServiceServer) testEmbeddedByValue()                            {}

// UnsafeChunkServerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChunkServerServiceServer will
// result in compilation errors.
type UnsafeChunkServerServiceServer interface {
	mustEmbedUnimplementedChunkServerServiceServer()
}

func RegisterChunkServerServiceServer(s grpc.ServiceRegistrar, srv ChunkServerServiceServer) {
	// If the following call pancis, it indicates UnimplementedChunkServerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ChunkServerService_ServiceDesc, srv)
}

func _ChunkServerService_StoreChunk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChunkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChunkServerServiceServer).StoreChunk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChunkServerService_StoreChunk_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChunkServerServiceServer).StoreChunk(ctx, req.(*ChunkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChunkServerService_CopyChunk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CopyChunkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChunkServerServiceServer).CopyChunk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChunkServerService_CopyChunk_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChunkServerServiceServer).CopyChunk(ctx, req.(*CopyChunkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChunkServerService_ServiceDesc is the grpc.ServiceDesc for ChunkServerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChunkServerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ChunkServerService",
	HandlerType: (*ChunkServerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StoreChunk",
			Handler:    _ChunkServerService_StoreChunk_Handler,
		},
		{
			MethodName: "CopyChunk",
			Handler:    _ChunkServerService_CopyChunk_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gfs.proto",
}
