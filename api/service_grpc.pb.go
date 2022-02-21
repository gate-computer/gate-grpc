// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: api/service.proto

package api

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RootClient is the client API for Root service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RootClient interface {
	Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*InitResponse, error)
}

type rootClient struct {
	cc grpc.ClientConnInterface
}

func NewRootClient(cc grpc.ClientConnInterface) RootClient {
	return &rootClient{cc}
}

func (c *rootClient) Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*InitResponse, error) {
	out := new(InitResponse)
	err := c.cc.Invoke(ctx, "/gate.service.Root/Init", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RootServer is the server API for Root service.
// All implementations must embed UnimplementedRootServer
// for forward compatibility
type RootServer interface {
	Init(context.Context, *InitRequest) (*InitResponse, error)
	mustEmbedUnimplementedRootServer()
}

// UnimplementedRootServer must be embedded to have forward compatible implementations.
type UnimplementedRootServer struct {
}

func (UnimplementedRootServer) Init(context.Context, *InitRequest) (*InitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (UnimplementedRootServer) mustEmbedUnimplementedRootServer() {}

// UnsafeRootServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RootServer will
// result in compilation errors.
type UnsafeRootServer interface {
	mustEmbedUnimplementedRootServer()
}

func RegisterRootServer(s grpc.ServiceRegistrar, srv RootServer) {
	s.RegisterService(&Root_ServiceDesc, srv)
}

func _Root_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RootServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gate.service.Root/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RootServer).Init(ctx, req.(*InitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Root_ServiceDesc is the grpc.ServiceDesc for Root service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Root_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gate.service.Root",
	HandlerType: (*RootServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Init",
			Handler:    _Root_Init_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/service.proto",
}

// InstanceClient is the client API for Instance service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InstanceClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Receive(ctx context.Context, in *ReceiveRequest, opts ...grpc.CallOption) (Instance_ReceiveClient, error)
	Handle(ctx context.Context, in *HandleRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Suspend(ctx context.Context, in *SuspendRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Snapshot(ctx context.Context, in *SnapshotRequest, opts ...grpc.CallOption) (*wrappers.BytesValue, error)
}

type instanceClient struct {
	cc grpc.ClientConnInterface
}

func NewInstanceClient(cc grpc.ClientConnInterface) InstanceClient {
	return &instanceClient{cc}
}

func (c *instanceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/gate.service.Instance/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceClient) Receive(ctx context.Context, in *ReceiveRequest, opts ...grpc.CallOption) (Instance_ReceiveClient, error) {
	stream, err := c.cc.NewStream(ctx, &Instance_ServiceDesc.Streams[0], "/gate.service.Instance/Receive", opts...)
	if err != nil {
		return nil, err
	}
	x := &instanceReceiveClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Instance_ReceiveClient interface {
	Recv() (*wrappers.BytesValue, error)
	grpc.ClientStream
}

type instanceReceiveClient struct {
	grpc.ClientStream
}

func (x *instanceReceiveClient) Recv() (*wrappers.BytesValue, error) {
	m := new(wrappers.BytesValue)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *instanceClient) Handle(ctx context.Context, in *HandleRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/gate.service.Instance/Handle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceClient) Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/gate.service.Instance/Shutdown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceClient) Suspend(ctx context.Context, in *SuspendRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/gate.service.Instance/Suspend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceClient) Snapshot(ctx context.Context, in *SnapshotRequest, opts ...grpc.CallOption) (*wrappers.BytesValue, error) {
	out := new(wrappers.BytesValue)
	err := c.cc.Invoke(ctx, "/gate.service.Instance/Snapshot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InstanceServer is the server API for Instance service.
// All implementations must embed UnimplementedInstanceServer
// for forward compatibility
type InstanceServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Receive(*ReceiveRequest, Instance_ReceiveServer) error
	Handle(context.Context, *HandleRequest) (*empty.Empty, error)
	Shutdown(context.Context, *ShutdownRequest) (*empty.Empty, error)
	Suspend(context.Context, *SuspendRequest) (*empty.Empty, error)
	Snapshot(context.Context, *SnapshotRequest) (*wrappers.BytesValue, error)
	mustEmbedUnimplementedInstanceServer()
}

// UnimplementedInstanceServer must be embedded to have forward compatible implementations.
type UnimplementedInstanceServer struct {
}

func (UnimplementedInstanceServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedInstanceServer) Receive(*ReceiveRequest, Instance_ReceiveServer) error {
	return status.Errorf(codes.Unimplemented, "method Receive not implemented")
}
func (UnimplementedInstanceServer) Handle(context.Context, *HandleRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handle not implemented")
}
func (UnimplementedInstanceServer) Shutdown(context.Context, *ShutdownRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Shutdown not implemented")
}
func (UnimplementedInstanceServer) Suspend(context.Context, *SuspendRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Suspend not implemented")
}
func (UnimplementedInstanceServer) Snapshot(context.Context, *SnapshotRequest) (*wrappers.BytesValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Snapshot not implemented")
}
func (UnimplementedInstanceServer) mustEmbedUnimplementedInstanceServer() {}

// UnsafeInstanceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InstanceServer will
// result in compilation errors.
type UnsafeInstanceServer interface {
	mustEmbedUnimplementedInstanceServer()
}

func RegisterInstanceServer(s grpc.ServiceRegistrar, srv InstanceServer) {
	s.RegisterService(&Instance_ServiceDesc, srv)
}

func _Instance_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gate.service.Instance/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Instance_Receive_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReceiveRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(InstanceServer).Receive(m, &instanceReceiveServer{stream})
}

type Instance_ReceiveServer interface {
	Send(*wrappers.BytesValue) error
	grpc.ServerStream
}

type instanceReceiveServer struct {
	grpc.ServerStream
}

func (x *instanceReceiveServer) Send(m *wrappers.BytesValue) error {
	return x.ServerStream.SendMsg(m)
}

func _Instance_Handle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServer).Handle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gate.service.Instance/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServer).Handle(ctx, req.(*HandleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Instance_Shutdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShutdownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServer).Shutdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gate.service.Instance/Shutdown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServer).Shutdown(ctx, req.(*ShutdownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Instance_Suspend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuspendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServer).Suspend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gate.service.Instance/Suspend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServer).Suspend(ctx, req.(*SuspendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Instance_Snapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServer).Snapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gate.service.Instance/Snapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServer).Snapshot(ctx, req.(*SnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Instance_ServiceDesc is the grpc.ServiceDesc for Instance service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Instance_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gate.service.Instance",
	HandlerType: (*InstanceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Instance_Create_Handler,
		},
		{
			MethodName: "Handle",
			Handler:    _Instance_Handle_Handler,
		},
		{
			MethodName: "Shutdown",
			Handler:    _Instance_Shutdown_Handler,
		},
		{
			MethodName: "Suspend",
			Handler:    _Instance_Suspend_Handler,
		},
		{
			MethodName: "Snapshot",
			Handler:    _Instance_Snapshot_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Receive",
			Handler:       _Instance_Receive_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/service.proto",
}