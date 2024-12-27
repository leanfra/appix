// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: api/appix/v1/applications.proto

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
	Applications_CreateApplications_FullMethodName = "/api.appix.v1.Applications/CreateApplications"
	Applications_UpdateApplications_FullMethodName = "/api.appix.v1.Applications/UpdateApplications"
	Applications_DeleteApplications_FullMethodName = "/api.appix.v1.Applications/DeleteApplications"
	Applications_GetApplications_FullMethodName    = "/api.appix.v1.Applications/GetApplications"
	Applications_ListApplications_FullMethodName   = "/api.appix.v1.Applications/ListApplications"
)

// ApplicationsClient is the client API for Applications service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApplicationsClient interface {
	CreateApplications(ctx context.Context, in *CreateApplicationsRequest, opts ...grpc.CallOption) (*CreateApplicationsReply, error)
	UpdateApplications(ctx context.Context, in *UpdateApplicationsRequest, opts ...grpc.CallOption) (*UpdateApplicationsReply, error)
	DeleteApplications(ctx context.Context, in *DeleteApplicationsRequest, opts ...grpc.CallOption) (*DeleteApplicationsReply, error)
	GetApplications(ctx context.Context, in *GetApplicationsRequest, opts ...grpc.CallOption) (*GetApplicationsReply, error)
	ListApplications(ctx context.Context, in *ListApplicationsRequest, opts ...grpc.CallOption) (*ListApplicationsReply, error)
}

type applicationsClient struct {
	cc grpc.ClientConnInterface
}

func NewApplicationsClient(cc grpc.ClientConnInterface) ApplicationsClient {
	return &applicationsClient{cc}
}

func (c *applicationsClient) CreateApplications(ctx context.Context, in *CreateApplicationsRequest, opts ...grpc.CallOption) (*CreateApplicationsReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateApplicationsReply)
	err := c.cc.Invoke(ctx, Applications_CreateApplications_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationsClient) UpdateApplications(ctx context.Context, in *UpdateApplicationsRequest, opts ...grpc.CallOption) (*UpdateApplicationsReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateApplicationsReply)
	err := c.cc.Invoke(ctx, Applications_UpdateApplications_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationsClient) DeleteApplications(ctx context.Context, in *DeleteApplicationsRequest, opts ...grpc.CallOption) (*DeleteApplicationsReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteApplicationsReply)
	err := c.cc.Invoke(ctx, Applications_DeleteApplications_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationsClient) GetApplications(ctx context.Context, in *GetApplicationsRequest, opts ...grpc.CallOption) (*GetApplicationsReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetApplicationsReply)
	err := c.cc.Invoke(ctx, Applications_GetApplications_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationsClient) ListApplications(ctx context.Context, in *ListApplicationsRequest, opts ...grpc.CallOption) (*ListApplicationsReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListApplicationsReply)
	err := c.cc.Invoke(ctx, Applications_ListApplications_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApplicationsServer is the server API for Applications service.
// All implementations must embed UnimplementedApplicationsServer
// for forward compatibility.
type ApplicationsServer interface {
	CreateApplications(context.Context, *CreateApplicationsRequest) (*CreateApplicationsReply, error)
	UpdateApplications(context.Context, *UpdateApplicationsRequest) (*UpdateApplicationsReply, error)
	DeleteApplications(context.Context, *DeleteApplicationsRequest) (*DeleteApplicationsReply, error)
	GetApplications(context.Context, *GetApplicationsRequest) (*GetApplicationsReply, error)
	ListApplications(context.Context, *ListApplicationsRequest) (*ListApplicationsReply, error)
	mustEmbedUnimplementedApplicationsServer()
}

// UnimplementedApplicationsServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedApplicationsServer struct{}

func (UnimplementedApplicationsServer) CreateApplications(context.Context, *CreateApplicationsRequest) (*CreateApplicationsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateApplications not implemented")
}
func (UnimplementedApplicationsServer) UpdateApplications(context.Context, *UpdateApplicationsRequest) (*UpdateApplicationsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateApplications not implemented")
}
func (UnimplementedApplicationsServer) DeleteApplications(context.Context, *DeleteApplicationsRequest) (*DeleteApplicationsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteApplications not implemented")
}
func (UnimplementedApplicationsServer) GetApplications(context.Context, *GetApplicationsRequest) (*GetApplicationsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApplications not implemented")
}
func (UnimplementedApplicationsServer) ListApplications(context.Context, *ListApplicationsRequest) (*ListApplicationsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListApplications not implemented")
}
func (UnimplementedApplicationsServer) mustEmbedUnimplementedApplicationsServer() {}
func (UnimplementedApplicationsServer) testEmbeddedByValue()                      {}

// UnsafeApplicationsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApplicationsServer will
// result in compilation errors.
type UnsafeApplicationsServer interface {
	mustEmbedUnimplementedApplicationsServer()
}

func RegisterApplicationsServer(s grpc.ServiceRegistrar, srv ApplicationsServer) {
	// If the following call pancis, it indicates UnimplementedApplicationsServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Applications_ServiceDesc, srv)
}

func _Applications_CreateApplications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateApplicationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationsServer).CreateApplications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Applications_CreateApplications_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationsServer).CreateApplications(ctx, req.(*CreateApplicationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Applications_UpdateApplications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateApplicationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationsServer).UpdateApplications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Applications_UpdateApplications_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationsServer).UpdateApplications(ctx, req.(*UpdateApplicationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Applications_DeleteApplications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteApplicationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationsServer).DeleteApplications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Applications_DeleteApplications_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationsServer).DeleteApplications(ctx, req.(*DeleteApplicationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Applications_GetApplications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApplicationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationsServer).GetApplications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Applications_GetApplications_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationsServer).GetApplications(ctx, req.(*GetApplicationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Applications_ListApplications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListApplicationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationsServer).ListApplications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Applications_ListApplications_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationsServer).ListApplications(ctx, req.(*ListApplicationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Applications_ServiceDesc is the grpc.ServiceDesc for Applications service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Applications_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.appix.v1.Applications",
	HandlerType: (*ApplicationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateApplications",
			Handler:    _Applications_CreateApplications_Handler,
		},
		{
			MethodName: "UpdateApplications",
			Handler:    _Applications_UpdateApplications_Handler,
		},
		{
			MethodName: "DeleteApplications",
			Handler:    _Applications_DeleteApplications_Handler,
		},
		{
			MethodName: "GetApplications",
			Handler:    _Applications_GetApplications_Handler,
		},
		{
			MethodName: "ListApplications",
			Handler:    _Applications_ListApplications_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/appix/v1/applications.proto",
}
