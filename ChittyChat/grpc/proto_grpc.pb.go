// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: grpc/proto.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ChittychatDB_GetConnectionLog_FullMethodName = "/ChittychatDB/GetConnectionLog"
	ChittychatDB_PublishPost_FullMethodName      = "/ChittychatDB/PublishPost"
	ChittychatDB_Connect_FullMethodName          = "/ChittychatDB/Connect"
	ChittychatDB_Disconnect_FullMethodName       = "/ChittychatDB/Disconnect"
)

// ChittychatDBClient is the client API for ChittychatDB service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChittychatDBClient interface {
	GetConnectionLog(ctx context.Context, in *ClientLT, opts ...grpc.CallOption) (*ConnectionsLog, error)
	PublishPost(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Posted, error)
	Connect(ctx context.Context, in *ClientInfo, opts ...grpc.CallOption) (*Empty, error)
	Disconnect(ctx context.Context, in *ClientInfo, opts ...grpc.CallOption) (*Empty, error)
}

type chittychatDBClient struct {
	cc grpc.ClientConnInterface
}

func NewChittychatDBClient(cc grpc.ClientConnInterface) ChittychatDBClient {
	return &chittychatDBClient{cc}
}

func (c *chittychatDBClient) GetConnectionLog(ctx context.Context, in *ClientLT, opts ...grpc.CallOption) (*ConnectionsLog, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ConnectionsLog)
	err := c.cc.Invoke(ctx, ChittychatDB_GetConnectionLog_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chittychatDBClient) PublishPost(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Posted, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Posted)
	err := c.cc.Invoke(ctx, ChittychatDB_PublishPost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chittychatDBClient) Connect(ctx context.Context, in *ClientInfo, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, ChittychatDB_Connect_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chittychatDBClient) Disconnect(ctx context.Context, in *ClientInfo, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, ChittychatDB_Disconnect_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChittychatDBServer is the server API for ChittychatDB service.
// All implementations must embed UnimplementedChittychatDBServer
// for forward compatibility.
type ChittychatDBServer interface {
	GetConnectionLog(context.Context, *ClientLT) (*ConnectionsLog, error)
	PublishPost(context.Context, *Post) (*Posted, error)
	Connect(context.Context, *ClientInfo) (*Empty, error)
	Disconnect(context.Context, *ClientInfo) (*Empty, error)
	mustEmbedUnimplementedChittychatDBServer()
}

// UnimplementedChittychatDBServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedChittychatDBServer struct{}

func (UnimplementedChittychatDBServer) GetConnectionLog(context.Context, *ClientLT) (*ConnectionsLog, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConnectionLog not implemented")
}
func (UnimplementedChittychatDBServer) PublishPost(context.Context, *Post) (*Posted, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishPost not implemented")
}
func (UnimplementedChittychatDBServer) Connect(context.Context, *ClientInfo) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedChittychatDBServer) Disconnect(context.Context, *ClientInfo) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disconnect not implemented")
}
func (UnimplementedChittychatDBServer) mustEmbedUnimplementedChittychatDBServer() {}
func (UnimplementedChittychatDBServer) testEmbeddedByValue()                      {}

// UnsafeChittychatDBServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChittychatDBServer will
// result in compilation errors.
type UnsafeChittychatDBServer interface {
	mustEmbedUnimplementedChittychatDBServer()
}

func RegisterChittychatDBServer(s grpc.ServiceRegistrar, srv ChittychatDBServer) {
	// If the following call pancis, it indicates UnimplementedChittychatDBServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ChittychatDB_ServiceDesc, srv)
}

func _ChittychatDB_GetConnectionLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientLT)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChittychatDBServer).GetConnectionLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChittychatDB_GetConnectionLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChittychatDBServer).GetConnectionLog(ctx, req.(*ClientLT))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChittychatDB_PublishPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Post)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChittychatDBServer).PublishPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChittychatDB_PublishPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChittychatDBServer).PublishPost(ctx, req.(*Post))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChittychatDB_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChittychatDBServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChittychatDB_Connect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChittychatDBServer).Connect(ctx, req.(*ClientInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChittychatDB_Disconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChittychatDBServer).Disconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChittychatDB_Disconnect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChittychatDBServer).Disconnect(ctx, req.(*ClientInfo))
	}
	return interceptor(ctx, in, info, handler)
}

// ChittychatDB_ServiceDesc is the grpc.ServiceDesc for ChittychatDB service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChittychatDB_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ChittychatDB",
	HandlerType: (*ChittychatDBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConnectionLog",
			Handler:    _ChittychatDB_GetConnectionLog_Handler,
		},
		{
			MethodName: "PublishPost",
			Handler:    _ChittychatDB_PublishPost_Handler,
		},
		{
			MethodName: "Connect",
			Handler:    _ChittychatDB_Connect_Handler,
		},
		{
			MethodName: "Disconnect",
			Handler:    _ChittychatDB_Disconnect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/proto.proto",
}
