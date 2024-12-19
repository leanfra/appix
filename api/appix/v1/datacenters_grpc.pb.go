// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: appix/v1/datacenters.proto

package v1

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
	Datacenters_CreateDatacenters_FullMethodName = "/api.appix.v1.Datacenters/CreateDatacenters"
	Datacenters_UpdateDatacenters_FullMethodName = "/api.appix.v1.Datacenters/UpdateDatacenters"
	Datacenters_DeleteDatacenters_FullMethodName = "/api.appix.v1.Datacenters/DeleteDatacenters"
	Datacenters_GetDatacenters_FullMethodName    = "/api.appix.v1.Datacenters/GetDatacenters"
	Datacenters_ListDatacenters_FullMethodName   = "/api.appix.v1.Datacenters/ListDatacenters"
)

// DatacentersClient is the client API for Datacenters service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DatacentersClient interface {
	CreateDatacenters(ctx context.Context, in *CreateDatacentersRequest, opts ...grpc.CallOption) (*CreateDatacentersReply, error)
	UpdateDatacenters(ctx context.Context, in *UpdateDatacentersRequest, opts ...grpc.CallOption) (*UpdateDatacentersReply, error)
	DeleteDatacenters(ctx context.Context, in *DeleteDatacentersRequest, opts ...grpc.CallOption) (*DeleteDatacentersReply, error)
	GetDatacenters(ctx context.Context, in *GetDatacentersRequest, opts ...grpc.CallOption) (*GetDatacentersReply, error)
	ListDatacenters(ctx context.Context, in *ListDatacentersRequest, opts ...grpc.CallOption) (*ListDatacentersReply, error)
}

type datacentersClient struct {
	cc grpc.ClientConnInterface
}

func NewDatacentersClient(cc grpc.ClientConnInterface) DatacentersClient {
	return &datacentersClient{cc}
}

func (c *datacentersClient) CreateDatacenters(ctx context.Context, in *CreateDatacentersRequest, opts ...grpc.CallOption) (*CreateDatacentersReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateDatacentersReply)
	err := c.cc.Invoke(ctx, Datacenters_CreateDatacenters_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datacentersClient) UpdateDatacenters(ctx context.Context, in *UpdateDatacentersRequest, opts ...grpc.CallOption) (*UpdateDatacentersReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateDatacentersReply)
	err := c.cc.Invoke(ctx, Datacenters_UpdateDatacenters_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datacentersClient) DeleteDatacenters(ctx context.Context, in *DeleteDatacentersRequest, opts ...grpc.CallOption) (*DeleteDatacentersReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteDatacentersReply)
	err := c.cc.Invoke(ctx, Datacenters_DeleteDatacenters_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datacentersClient) GetDatacenters(ctx context.Context, in *GetDatacentersRequest, opts ...grpc.CallOption) (*GetDatacentersReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDatacentersReply)
	err := c.cc.Invoke(ctx, Datacenters_GetDatacenters_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datacentersClient) ListDatacenters(ctx context.Context, in *ListDatacentersRequest, opts ...grpc.CallOption) (*ListDatacentersReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListDatacentersReply)
	err := c.cc.Invoke(ctx, Datacenters_ListDatacenters_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatacentersServer is the server API for Datacenters service.
// All implementations must embed UnimplementedDatacentersServer
// for forward compatibility.
type DatacentersServer interface {
	CreateDatacenters(context.Context, *CreateDatacentersRequest) (*CreateDatacentersReply, error)
	UpdateDatacenters(context.Context, *UpdateDatacentersRequest) (*UpdateDatacentersReply, error)
	DeleteDatacenters(context.Context, *DeleteDatacentersRequest) (*DeleteDatacentersReply, error)
	GetDatacenters(context.Context, *GetDatacentersRequest) (*GetDatacentersReply, error)
	ListDatacenters(context.Context, *ListDatacentersRequest) (*ListDatacentersReply, error)
	mustEmbedUnimplementedDatacentersServer()
}

// UnimplementedDatacentersServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDatacentersServer struct{}

func (UnimplementedDatacentersServer) CreateDatacenters(context.Context, *CreateDatacentersRequest) (*CreateDatacentersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDatacenters not implemented")
}
func (UnimplementedDatacentersServer) UpdateDatacenters(context.Context, *UpdateDatacentersRequest) (*UpdateDatacentersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDatacenters not implemented")
}
func (UnimplementedDatacentersServer) DeleteDatacenters(context.Context, *DeleteDatacentersRequest) (*DeleteDatacentersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDatacenters not implemented")
}
func (UnimplementedDatacentersServer) GetDatacenters(context.Context, *GetDatacentersRequest) (*GetDatacentersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDatacenters not implemented")
}
func (UnimplementedDatacentersServer) ListDatacenters(context.Context, *ListDatacentersRequest) (*ListDatacentersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDatacenters not implemented")
}
func (UnimplementedDatacentersServer) mustEmbedUnimplementedDatacentersServer() {}
func (UnimplementedDatacentersServer) testEmbeddedByValue()                     {}

// UnsafeDatacentersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DatacentersServer will
// result in compilation errors.
type UnsafeDatacentersServer interface {
	mustEmbedUnimplementedDatacentersServer()
}

func RegisterDatacentersServer(s grpc.ServiceRegistrar, srv DatacentersServer) {
	// If the following call pancis, it indicates UnimplementedDatacentersServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Datacenters_ServiceDesc, srv)
}

func _Datacenters_CreateDatacenters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDatacentersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatacentersServer).CreateDatacenters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Datacenters_CreateDatacenters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatacentersServer).CreateDatacenters(ctx, req.(*CreateDatacentersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Datacenters_UpdateDatacenters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDatacentersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatacentersServer).UpdateDatacenters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Datacenters_UpdateDatacenters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatacentersServer).UpdateDatacenters(ctx, req.(*UpdateDatacentersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Datacenters_DeleteDatacenters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDatacentersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatacentersServer).DeleteDatacenters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Datacenters_DeleteDatacenters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatacentersServer).DeleteDatacenters(ctx, req.(*DeleteDatacentersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Datacenters_GetDatacenters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDatacentersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatacentersServer).GetDatacenters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Datacenters_GetDatacenters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatacentersServer).GetDatacenters(ctx, req.(*GetDatacentersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Datacenters_ListDatacenters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDatacentersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatacentersServer).ListDatacenters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Datacenters_ListDatacenters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatacentersServer).ListDatacenters(ctx, req.(*ListDatacentersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Datacenters_ServiceDesc is the grpc.ServiceDesc for Datacenters service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Datacenters_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.appix.v1.Datacenters",
	HandlerType: (*DatacentersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDatacenters",
			Handler:    _Datacenters_CreateDatacenters_Handler,
		},
		{
			MethodName: "UpdateDatacenters",
			Handler:    _Datacenters_UpdateDatacenters_Handler,
		},
		{
			MethodName: "DeleteDatacenters",
			Handler:    _Datacenters_DeleteDatacenters_Handler,
		},
		{
			MethodName: "GetDatacenters",
			Handler:    _Datacenters_GetDatacenters_Handler,
		},
		{
			MethodName: "ListDatacenters",
			Handler:    _Datacenters_ListDatacenters_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "appix/v1/datacenters.proto",
}
