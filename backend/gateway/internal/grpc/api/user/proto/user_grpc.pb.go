// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
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
	UserService_Setup2FA_FullMethodName              = "/proto.UserService/Setup2FA"
	UserService_Test2FA_FullMethodName               = "/proto.UserService/Test2FA"
	UserService_CancelSetup2FA_FullMethodName        = "/proto.UserService/CancelSetup2FA"
	UserService_Get2FAStatus_FullMethodName          = "/proto.UserService/Get2FAStatus"
	UserService_Disable2FA_FullMethodName            = "/proto.UserService/Disable2FA"
	UserService_UpdateUserInfo_FullMethodName        = "/proto.UserService/UpdateUserInfo"
	UserService_CreateNewTicket_FullMethodName       = "/proto.UserService/CreateNewTicket"
	UserService_SendChatMessage_FullMethodName       = "/proto.UserService/SendChatMessage"
	UserService_DeleteMyAccount_FullMethodName       = "/proto.UserService/DeleteMyAccount"
	UserService_GetAllUsers_FullMethodName           = "/proto.UserService/GetAllUsers"
	UserService_AddUserManually_FullMethodName       = "/proto.UserService/AddUserManually"
	UserService_UpdateUserInfoAdmin_FullMethodName   = "/proto.UserService/UpdateUserInfoAdmin"
	UserService_Block2UnblockUserById_FullMethodName = "/proto.UserService/Block2UnblockUserById"
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
	Setup2FA(ctx context.Context, in *Setup2FARequest, opts ...grpc.CallOption) (*Setup2FAResponse, error)
	Test2FA(ctx context.Context, in *Test2FARequest, opts ...grpc.CallOption) (*Test2FAResponse, error)
	CancelSetup2FA(ctx context.Context, in *CancelSetup2FARequest, opts ...grpc.CallOption) (*CancelSetup2FAResponse, error)
	Get2FAStatus(ctx context.Context, in *Get2FAStatusRequest, opts ...grpc.CallOption) (*Get2FAStatusResponse, error)
	Disable2FA(ctx context.Context, in *Disable2FARequest, opts ...grpc.CallOption) (*Disable2FAResponse, error)
	UpdateUserInfo(ctx context.Context, in *UpdateUserInfoRequest, opts ...grpc.CallOption) (*UpdateUserInfoResponse, error)
	CreateNewTicket(ctx context.Context, in *CreateNewTicketRequest, opts ...grpc.CallOption) (*CreateNewTicketResponse, error)
	SendChatMessage(ctx context.Context, in *SendChatMessageRequest, opts ...grpc.CallOption) (*SendChatMessageResponse, error)
	DeleteMyAccount(ctx context.Context, in *DeleteMyAccountRequest, opts ...grpc.CallOption) (*DeleteMyAccountResponse, error)
	GetAllUsers(ctx context.Context, in *GetAllUsersRequest, opts ...grpc.CallOption) (*GetAllUsersResponse, error)
	AddUserManually(ctx context.Context, in *AddUserManuallyRequest, opts ...grpc.CallOption) (*AddUserManuallyResponse, error)
	UpdateUserInfoAdmin(ctx context.Context, in *UpdateUserInfoAdminRequest, opts ...grpc.CallOption) (*UpdateUserInfoAdminResponse, error)
	Block2UnblockUserById(ctx context.Context, in *Block2UnblockUserByIdRequest, opts ...grpc.CallOption) (*Block2UnblockUserByIdResponse, error)
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

func (c *userServiceClient) Setup2FA(ctx context.Context, in *Setup2FARequest, opts ...grpc.CallOption) (*Setup2FAResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Setup2FAResponse)
	err := c.cc.Invoke(ctx, UserService_Setup2FA_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Test2FA(ctx context.Context, in *Test2FARequest, opts ...grpc.CallOption) (*Test2FAResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Test2FAResponse)
	err := c.cc.Invoke(ctx, UserService_Test2FA_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CancelSetup2FA(ctx context.Context, in *CancelSetup2FARequest, opts ...grpc.CallOption) (*CancelSetup2FAResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CancelSetup2FAResponse)
	err := c.cc.Invoke(ctx, UserService_CancelSetup2FA_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Get2FAStatus(ctx context.Context, in *Get2FAStatusRequest, opts ...grpc.CallOption) (*Get2FAStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Get2FAStatusResponse)
	err := c.cc.Invoke(ctx, UserService_Get2FAStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Disable2FA(ctx context.Context, in *Disable2FARequest, opts ...grpc.CallOption) (*Disable2FAResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Disable2FAResponse)
	err := c.cc.Invoke(ctx, UserService_Disable2FA_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserInfo(ctx context.Context, in *UpdateUserInfoRequest, opts ...grpc.CallOption) (*UpdateUserInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateUserInfoResponse)
	err := c.cc.Invoke(ctx, UserService_UpdateUserInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateNewTicket(ctx context.Context, in *CreateNewTicketRequest, opts ...grpc.CallOption) (*CreateNewTicketResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateNewTicketResponse)
	err := c.cc.Invoke(ctx, UserService_CreateNewTicket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SendChatMessage(ctx context.Context, in *SendChatMessageRequest, opts ...grpc.CallOption) (*SendChatMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendChatMessageResponse)
	err := c.cc.Invoke(ctx, UserService_SendChatMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteMyAccount(ctx context.Context, in *DeleteMyAccountRequest, opts ...grpc.CallOption) (*DeleteMyAccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteMyAccountResponse)
	err := c.cc.Invoke(ctx, UserService_DeleteMyAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAllUsers(ctx context.Context, in *GetAllUsersRequest, opts ...grpc.CallOption) (*GetAllUsersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllUsersResponse)
	err := c.cc.Invoke(ctx, UserService_GetAllUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddUserManually(ctx context.Context, in *AddUserManuallyRequest, opts ...grpc.CallOption) (*AddUserManuallyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddUserManuallyResponse)
	err := c.cc.Invoke(ctx, UserService_AddUserManually_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserInfoAdmin(ctx context.Context, in *UpdateUserInfoAdminRequest, opts ...grpc.CallOption) (*UpdateUserInfoAdminResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateUserInfoAdminResponse)
	err := c.cc.Invoke(ctx, UserService_UpdateUserInfoAdmin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Block2UnblockUserById(ctx context.Context, in *Block2UnblockUserByIdRequest, opts ...grpc.CallOption) (*Block2UnblockUserByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Block2UnblockUserByIdResponse)
	err := c.cc.Invoke(ctx, UserService_Block2UnblockUserById_FullMethodName, in, out, cOpts...)
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
	Setup2FA(context.Context, *Setup2FARequest) (*Setup2FAResponse, error)
	Test2FA(context.Context, *Test2FARequest) (*Test2FAResponse, error)
	CancelSetup2FA(context.Context, *CancelSetup2FARequest) (*CancelSetup2FAResponse, error)
	Get2FAStatus(context.Context, *Get2FAStatusRequest) (*Get2FAStatusResponse, error)
	Disable2FA(context.Context, *Disable2FARequest) (*Disable2FAResponse, error)
	UpdateUserInfo(context.Context, *UpdateUserInfoRequest) (*UpdateUserInfoResponse, error)
	CreateNewTicket(context.Context, *CreateNewTicketRequest) (*CreateNewTicketResponse, error)
	SendChatMessage(context.Context, *SendChatMessageRequest) (*SendChatMessageResponse, error)
	DeleteMyAccount(context.Context, *DeleteMyAccountRequest) (*DeleteMyAccountResponse, error)
	GetAllUsers(context.Context, *GetAllUsersRequest) (*GetAllUsersResponse, error)
	AddUserManually(context.Context, *AddUserManuallyRequest) (*AddUserManuallyResponse, error)
	UpdateUserInfoAdmin(context.Context, *UpdateUserInfoAdminRequest) (*UpdateUserInfoAdminResponse, error)
	Block2UnblockUserById(context.Context, *Block2UnblockUserByIdRequest) (*Block2UnblockUserByIdResponse, error)
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
func (UnimplementedUserServiceServer) Setup2FA(context.Context, *Setup2FARequest) (*Setup2FAResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Setup2FA not implemented")
}
func (UnimplementedUserServiceServer) Test2FA(context.Context, *Test2FARequest) (*Test2FAResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Test2FA not implemented")
}
func (UnimplementedUserServiceServer) CancelSetup2FA(context.Context, *CancelSetup2FARequest) (*CancelSetup2FAResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelSetup2FA not implemented")
}
func (UnimplementedUserServiceServer) Get2FAStatus(context.Context, *Get2FAStatusRequest) (*Get2FAStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get2FAStatus not implemented")
}
func (UnimplementedUserServiceServer) Disable2FA(context.Context, *Disable2FARequest) (*Disable2FAResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disable2FA not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserInfo(context.Context, *UpdateUserInfoRequest) (*UpdateUserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserInfo not implemented")
}
func (UnimplementedUserServiceServer) CreateNewTicket(context.Context, *CreateNewTicketRequest) (*CreateNewTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewTicket not implemented")
}
func (UnimplementedUserServiceServer) SendChatMessage(context.Context, *SendChatMessageRequest) (*SendChatMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendChatMessage not implemented")
}
func (UnimplementedUserServiceServer) DeleteMyAccount(context.Context, *DeleteMyAccountRequest) (*DeleteMyAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMyAccount not implemented")
}
func (UnimplementedUserServiceServer) GetAllUsers(context.Context, *GetAllUsersRequest) (*GetAllUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUsers not implemented")
}
func (UnimplementedUserServiceServer) AddUserManually(context.Context, *AddUserManuallyRequest) (*AddUserManuallyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserManually not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserInfoAdmin(context.Context, *UpdateUserInfoAdminRequest) (*UpdateUserInfoAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserInfoAdmin not implemented")
}
func (UnimplementedUserServiceServer) Block2UnblockUserById(context.Context, *Block2UnblockUserByIdRequest) (*Block2UnblockUserByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Block2UnblockUserById not implemented")
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

func _UserService_Setup2FA_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Setup2FARequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Setup2FA(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Setup2FA_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Setup2FA(ctx, req.(*Setup2FARequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Test2FA_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Test2FARequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Test2FA(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Test2FA_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Test2FA(ctx, req.(*Test2FARequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CancelSetup2FA_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelSetup2FARequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CancelSetup2FA(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_CancelSetup2FA_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CancelSetup2FA(ctx, req.(*CancelSetup2FARequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Get2FAStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Get2FAStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Get2FAStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Get2FAStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Get2FAStatus(ctx, req.(*Get2FAStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Disable2FA_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Disable2FARequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Disable2FA(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Disable2FA_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Disable2FA(ctx, req.(*Disable2FARequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserInfo(ctx, req.(*UpdateUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateNewTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNewTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateNewTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_CreateNewTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateNewTicket(ctx, req.(*CreateNewTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SendChatMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendChatMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SendChatMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_SendChatMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SendChatMessage(ctx, req.(*SendChatMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteMyAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMyAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteMyAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_DeleteMyAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteMyAccount(ctx, req.(*DeleteMyAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetAllUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAllUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetAllUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAllUsers(ctx, req.(*GetAllUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddUserManually_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserManuallyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddUserManually(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_AddUserManually_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddUserManually(ctx, req.(*AddUserManuallyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserInfoAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserInfoAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserInfoAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateUserInfoAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserInfoAdmin(ctx, req.(*UpdateUserInfoAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Block2UnblockUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Block2UnblockUserByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Block2UnblockUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Block2UnblockUserById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Block2UnblockUserById(ctx, req.(*Block2UnblockUserByIdRequest))
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
		{
			MethodName: "Setup2FA",
			Handler:    _UserService_Setup2FA_Handler,
		},
		{
			MethodName: "Test2FA",
			Handler:    _UserService_Test2FA_Handler,
		},
		{
			MethodName: "CancelSetup2FA",
			Handler:    _UserService_CancelSetup2FA_Handler,
		},
		{
			MethodName: "Get2FAStatus",
			Handler:    _UserService_Get2FAStatus_Handler,
		},
		{
			MethodName: "Disable2FA",
			Handler:    _UserService_Disable2FA_Handler,
		},
		{
			MethodName: "UpdateUserInfo",
			Handler:    _UserService_UpdateUserInfo_Handler,
		},
		{
			MethodName: "CreateNewTicket",
			Handler:    _UserService_CreateNewTicket_Handler,
		},
		{
			MethodName: "SendChatMessage",
			Handler:    _UserService_SendChatMessage_Handler,
		},
		{
			MethodName: "DeleteMyAccount",
			Handler:    _UserService_DeleteMyAccount_Handler,
		},
		{
			MethodName: "GetAllUsers",
			Handler:    _UserService_GetAllUsers_Handler,
		},
		{
			MethodName: "AddUserManually",
			Handler:    _UserService_AddUserManually_Handler,
		},
		{
			MethodName: "UpdateUserInfoAdmin",
			Handler:    _UserService_UpdateUserInfoAdmin_Handler,
		},
		{
			MethodName: "Block2UnblockUserById",
			Handler:    _UserService_Block2UnblockUserById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
