// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: appix/v1/admin.proto

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
	Admin_CreateUsers_FullMethodName = "/api.appix.v1.Admin/CreateUsers"
	Admin_UpdateUsers_FullMethodName = "/api.appix.v1.Admin/UpdateUsers"
	Admin_DeleteUsers_FullMethodName = "/api.appix.v1.Admin/DeleteUsers"
	Admin_GetUsers_FullMethodName    = "/api.appix.v1.Admin/GetUsers"
	Admin_ListUsers_FullMethodName   = "/api.appix.v1.Admin/ListUsers"
	Admin_Login_FullMethodName       = "/api.appix.v1.Admin/Login"
	Admin_Logout_FullMethodName      = "/api.appix.v1.Admin/Logout"
)

// AdminClient is the client API for Admin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminClient interface {
	CreateUsers(ctx context.Context, in *CreateUsersRequest, opts ...grpc.CallOption) (*CreateUsersReply, error)
	UpdateUsers(ctx context.Context, in *UpdateUsersRequest, opts ...grpc.CallOption) (*UpdateUsersReply, error)
	DeleteUsers(ctx context.Context, in *DeleteUsersRequest, opts ...grpc.CallOption) (*DeleteUsersReply, error)
	GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersReply, error)
	ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersReply, error)
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginReply, error)
	Logout(ctx context.Context, in *LogoutReq, opts ...grpc.CallOption) (*LogoutReply, error)
}

type adminClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminClient(cc grpc.ClientConnInterface) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) CreateUsers(ctx context.Context, in *CreateUsersRequest, opts ...grpc.CallOption) (*CreateUsersReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateUsersReply)
	err := c.cc.Invoke(ctx, Admin_CreateUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) UpdateUsers(ctx context.Context, in *UpdateUsersRequest, opts ...grpc.CallOption) (*UpdateUsersReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateUsersReply)
	err := c.cc.Invoke(ctx, Admin_UpdateUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) DeleteUsers(ctx context.Context, in *DeleteUsersRequest, opts ...grpc.CallOption) (*DeleteUsersReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteUsersReply)
	err := c.cc.Invoke(ctx, Admin_DeleteUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUsersReply)
	err := c.cc.Invoke(ctx, Admin_GetUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListUsersReply)
	err := c.cc.Invoke(ctx, Admin_ListUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, Admin_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) Logout(ctx context.Context, in *LogoutReq, opts ...grpc.CallOption) (*LogoutReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LogoutReply)
	err := c.cc.Invoke(ctx, Admin_Logout_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServer is the server API for Admin service.
// All implementations must embed UnimplementedAdminServer
// for forward compatibility.
type AdminServer interface {
	CreateUsers(context.Context, *CreateUsersRequest) (*CreateUsersReply, error)
	UpdateUsers(context.Context, *UpdateUsersRequest) (*UpdateUsersReply, error)
	DeleteUsers(context.Context, *DeleteUsersRequest) (*DeleteUsersReply, error)
	GetUsers(context.Context, *GetUsersRequest) (*GetUsersReply, error)
	ListUsers(context.Context, *ListUsersRequest) (*ListUsersReply, error)
	Login(context.Context, *LoginReq) (*LoginReply, error)
	Logout(context.Context, *LogoutReq) (*LogoutReply, error)
	mustEmbedUnimplementedAdminServer()
}

// UnimplementedAdminServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAdminServer struct{}

func (UnimplementedAdminServer) CreateUsers(context.Context, *CreateUsersRequest) (*CreateUsersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUsers not implemented")
}
func (UnimplementedAdminServer) UpdateUsers(context.Context, *UpdateUsersRequest) (*UpdateUsersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUsers not implemented")
}
func (UnimplementedAdminServer) DeleteUsers(context.Context, *DeleteUsersRequest) (*DeleteUsersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUsers not implemented")
}
func (UnimplementedAdminServer) GetUsers(context.Context, *GetUsersRequest) (*GetUsersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedAdminServer) ListUsers(context.Context, *ListUsersRequest) (*ListUsersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (UnimplementedAdminServer) Login(context.Context, *LoginReq) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAdminServer) Logout(context.Context, *LogoutReq) (*LogoutReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedAdminServer) mustEmbedUnimplementedAdminServer() {}
func (UnimplementedAdminServer) testEmbeddedByValue()               {}

// UnsafeAdminServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServer will
// result in compilation errors.
type UnsafeAdminServer interface {
	mustEmbedUnimplementedAdminServer()
}

func RegisterAdminServer(s grpc.ServiceRegistrar, srv AdminServer) {
	// If the following call pancis, it indicates UnimplementedAdminServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Admin_ServiceDesc, srv)
}

func _Admin_CreateUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).CreateUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_CreateUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).CreateUsers(ctx, req.(*CreateUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_UpdateUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).UpdateUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_UpdateUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).UpdateUsers(ctx, req.(*UpdateUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_DeleteUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).DeleteUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_DeleteUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).DeleteUsers(ctx, req.(*DeleteUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_GetUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).GetUsers(ctx, req.(*GetUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_ListUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).ListUsers(ctx, req.(*ListUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_Logout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).Logout(ctx, req.(*LogoutReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Admin_ServiceDesc is the grpc.ServiceDesc for Admin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Admin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.appix.v1.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUsers",
			Handler:    _Admin_CreateUsers_Handler,
		},
		{
			MethodName: "UpdateUsers",
			Handler:    _Admin_UpdateUsers_Handler,
		},
		{
			MethodName: "DeleteUsers",
			Handler:    _Admin_DeleteUsers_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _Admin_GetUsers_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _Admin_ListUsers_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Admin_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _Admin_Logout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "appix/v1/admin.proto",
}
