// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: user.proto

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
	UserService_Login_FullMethodName                 = "/proto.UserService/Login"
	UserService_Register_FullMethodName              = "/proto.UserService/Register"
	UserService_CheckPreviousPassword_FullMethodName = "/proto.UserService/CheckPreviousPassword"
	UserService_ApplyNewPassword_FullMethodName      = "/proto.UserService/ApplyNewPassword"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	// 定义用户登录 RPC 方法
	Login(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error)
	// 定义用户注册 RPC 方法
	Register(ctx context.Context, in *UserRegisterRequest, opts ...grpc.CallOption) (*UserRegisterResponse, error)
	// 验证旧密码
	CheckPreviousPassword(ctx context.Context, in *CheckPreviousPasswordRequest, opts ...grpc.CallOption) (*CheckPreviousPasswordResponse, error)
	// 更新用户密码
	ApplyNewPassword(ctx context.Context, in *ApplyNewPasswordRequest, opts ...grpc.CallOption) (*ApplyNewPasswordResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Login(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserLoginResponse)
	err := c.cc.Invoke(ctx, UserService_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Register(ctx context.Context, in *UserRegisterRequest, opts ...grpc.CallOption) (*UserRegisterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserRegisterResponse)
	err := c.cc.Invoke(ctx, UserService_Register_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CheckPreviousPassword(ctx context.Context, in *CheckPreviousPasswordRequest, opts ...grpc.CallOption) (*CheckPreviousPasswordResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckPreviousPasswordResponse)
	err := c.cc.Invoke(ctx, UserService_CheckPreviousPassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ApplyNewPassword(ctx context.Context, in *ApplyNewPasswordRequest, opts ...grpc.CallOption) (*ApplyNewPasswordResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ApplyNewPasswordResponse)
	err := c.cc.Invoke(ctx, UserService_ApplyNewPassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility.
type UserServiceServer interface {
	// 定义用户登录 RPC 方法
	Login(context.Context, *UserLoginRequest) (*UserLoginResponse, error)
	// 定义用户注册 RPC 方法
	Register(context.Context, *UserRegisterRequest) (*UserRegisterResponse, error)
	// 验证旧密码
	CheckPreviousPassword(context.Context, *CheckPreviousPasswordRequest) (*CheckPreviousPasswordResponse, error)
	// 更新用户密码
	ApplyNewPassword(context.Context, *ApplyNewPasswordRequest) (*ApplyNewPasswordResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserServiceServer struct{}

func (UnimplementedUserServiceServer) Login(context.Context, *UserLoginRequest) (*UserLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServiceServer) Register(context.Context, *UserRegisterRequest) (*UserRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserServiceServer) CheckPreviousPassword(context.Context, *CheckPreviousPasswordRequest) (*CheckPreviousPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPreviousPassword not implemented")
}
func (UnimplementedUserServiceServer) ApplyNewPassword(context.Context, *ApplyNewPasswordRequest) (*ApplyNewPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyNewPassword not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}
func (UnimplementedUserServiceServer) testEmbeddedByValue()                     {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	// If the following call pancis, it indicates UnimplementedUserServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Login(ctx, req.(*UserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Register(ctx, req.(*UserRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CheckPreviousPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckPreviousPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CheckPreviousPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_CheckPreviousPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CheckPreviousPassword(ctx, req.(*CheckPreviousPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ApplyNewPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyNewPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ApplyNewPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_ApplyNewPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ApplyNewPassword(ctx, req.(*ApplyNewPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _UserService_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _UserService_Register_Handler,
		},
		{
			MethodName: "CheckPreviousPassword",
			Handler:    _UserService_CheckPreviousPassword_Handler,
		},
		{
			MethodName: "ApplyNewPassword",
			Handler:    _UserService_ApplyNewPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
