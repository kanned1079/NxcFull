// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: key.proto

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
	KeyService_GetAllMyKeysByUserIdAsc_FullMethodName     = "/key.KeyService/GetAllMyKeysByUserIdAsc"
	KeyService_GetKeyInfoById_FullMethodName              = "/key.KeyService/GetKeyInfoById"
	KeyService_RegisterClientWithKey_FullMethodName       = "/key.KeyService/RegisterClientWithKey"
	KeyService_CancelRegisterClientWithKey_FullMethodName = "/key.KeyService/CancelRegisterClientWithKey"
	KeyService_BindOrActiveMyKey2App_FullMethodName       = "/key.KeyService/BindOrActiveMyKey2App"
	KeyService_UnbindKey_FullMethodName                   = "/key.KeyService/UnbindKey"
	KeyService_AlterKeyBindRemark_FullMethodName          = "/key.KeyService/AlterKeyBindRemark"
	KeyService_CheckAccountAndKeyUsability_FullMethodName = "/key.KeyService/CheckAccountAndKeyUsability"
	KeyService_GetActivateLogByUserId_FullMethodName      = "/key.KeyService/GetActivateLogByUserId"
	KeyService_GetActivateLogByAdmin_FullMethodName       = "/key.KeyService/GetActivateLogByAdmin"
	KeyService_AlterKeyValidInfoByAdmin_FullMethodName    = "/key.KeyService/AlterKeyValidInfoByAdmin"
	KeyService_GetAllKeysByAdminDesc_FullMethodName       = "/key.KeyService/GetAllKeysByAdminDesc"
)

// KeyServiceClient is the client API for KeyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 定义 Key 服务
type KeyServiceClient interface {
	// 获取用户所有的key 包含有效和无效的 按照时间和是否有效升序排序
	GetAllMyKeysByUserIdAsc(ctx context.Context, in *GetAllMyKeysByUserIdAscRequest, opts ...grpc.CallOption) (*GetAllMyKeysByUserIdAscResponse, error)
	// 獲取單個密鑰的信息
	GetKeyInfoById(ctx context.Context, in *GetKeyInfoByIdRequest, opts ...grpc.CallOption) (*GetKeyInfoByIdResponse, error)
	// 第三方API 激活密钥
	RegisterClientWithKey(ctx context.Context, in *RegisterClientWithKeyRequest, opts ...grpc.CallOption) (*RegisterClientWithKeyResponse, error)
	// 第三方API 取消激活密钥
	CancelRegisterClientWithKey(ctx context.Context, in *CancelRegisterClientWithKeyRequest, opts ...grpc.CallOption) (*CancelRegisterClientWithKeyResponse, error)
	// 密钥使用记录
	BindOrActiveMyKey2App(ctx context.Context, in *BindOrActiveMyKey2AppRequest, opts ...grpc.CallOption) (*BindOrActiveMyKey2AppResponse, error)
	UnbindKey(ctx context.Context, in *UnbindKeyRequest, opts ...grpc.CallOption) (*UnbindKeyResponse, error)
	AlterKeyBindRemark(ctx context.Context, in *AlterKeyBindRemarkRequest, opts ...grpc.CallOption) (*AlterKeyBindRemarkResponse, error)
	CheckAccountAndKeyUsability(ctx context.Context, in *CheckAccountAndKeyUsabilityRequest, opts ...grpc.CallOption) (*CheckAccountAndKeyUsabilityResponse, error)
	GetActivateLogByUserId(ctx context.Context, in *GetActivateLogByUserIdRequest, opts ...grpc.CallOption) (*GetActivateLogByUserIdResponse, error)
	GetActivateLogByAdmin(ctx context.Context, in *GetActivateLogByAdminRequest, opts ...grpc.CallOption) (*GetActivateLogByAdminResponse, error)
	AlterKeyValidInfoByAdmin(ctx context.Context, in *AlterKeyValidInfoByAdminRequest, opts ...grpc.CallOption) (*AlterKeyValidInfoByAdminResponse, error)
	// Admin
	GetAllKeysByAdminDesc(ctx context.Context, in *GetAllKeysByAdminDescRequest, opts ...grpc.CallOption) (*GetAllKeysByAdminDescResponse, error)
}

type keyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeyServiceClient(cc grpc.ClientConnInterface) KeyServiceClient {
	return &keyServiceClient{cc}
}

func (c *keyServiceClient) GetAllMyKeysByUserIdAsc(ctx context.Context, in *GetAllMyKeysByUserIdAscRequest, opts ...grpc.CallOption) (*GetAllMyKeysByUserIdAscResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllMyKeysByUserIdAscResponse)
	err := c.cc.Invoke(ctx, KeyService_GetAllMyKeysByUserIdAsc_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyServiceClient) GetKeyInfoById(ctx context.Context, in *GetKeyInfoByIdRequest, opts ...grpc.CallOption) (*GetKeyInfoByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetKeyInfoByIdResponse)
	err := c.cc.Invoke(ctx, KeyService_GetKeyInfoById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyServiceClient) RegisterClientWithKey(ctx context.Context, in *RegisterClientWithKeyRequest, opts ...grpc.CallOption) (*RegisterClientWithKeyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterClientWithKeyResponse)
	err := c.cc.Invoke(ctx, KeyService_RegisterClientWithKey_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyServiceClient) CancelRegisterClientWithKey(ctx context.Context, in *CancelRegisterClientWithKeyRequest, opts ...grpc.CallOption) (*CancelRegisterClientWithKeyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CancelRegisterClientWithKeyResponse)
	err := c.cc.Invoke(ctx, KeyService_CancelRegisterClientWithKey_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyServiceClient) BindOrActiveMyKey2App(ctx context.Context, in *BindOrActiveMyKey2AppRequest, opts ...grpc.CallOption) (*BindOrActiveMyKey2AppResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BindOrActiveMyKey2AppResponse)
	err := c.cc.Invoke(ctx, KeyService_BindOrActiveMyKey2App_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyServiceClient) UnbindKey(ctx context.Context, in *UnbindKeyRequest, opts ...grpc.CallOption) (*UnbindKeyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UnbindKeyResponse)
	err := c.cc.Invoke(ctx, KeyService_UnbindKey_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyServiceClient) AlterKeyBindRemark(ctx context.Context, in *AlterKeyBindRemarkRequest, opts ...grpc.CallOption) (*AlterKeyBindRemarkResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AlterKeyBindRemarkResponse)
	err := c.cc.Invoke(ctx, KeyService_AlterKeyBindRemark_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyServiceClient) CheckAccountAndKeyUsability(ctx context.Context, in *CheckAccountAndKeyUsabilityRequest, opts ...grpc.CallOption) (*CheckAccountAndKeyUsabilityResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckAccountAndKeyUsabilityResponse)
	err := c.cc.Invoke(ctx, KeyService_CheckAccountAndKeyUsability_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyServiceClient) GetActivateLogByUserId(ctx context.Context, in *GetActivateLogByUserIdRequest, opts ...grpc.CallOption) (*GetActivateLogByUserIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetActivateLogByUserIdResponse)
	err := c.cc.Invoke(ctx, KeyService_GetActivateLogByUserId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyServiceClient) GetActivateLogByAdmin(ctx context.Context, in *GetActivateLogByAdminRequest, opts ...grpc.CallOption) (*GetActivateLogByAdminResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetActivateLogByAdminResponse)
	err := c.cc.Invoke(ctx, KeyService_GetActivateLogByAdmin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyServiceClient) AlterKeyValidInfoByAdmin(ctx context.Context, in *AlterKeyValidInfoByAdminRequest, opts ...grpc.CallOption) (*AlterKeyValidInfoByAdminResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AlterKeyValidInfoByAdminResponse)
	err := c.cc.Invoke(ctx, KeyService_AlterKeyValidInfoByAdmin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyServiceClient) GetAllKeysByAdminDesc(ctx context.Context, in *GetAllKeysByAdminDescRequest, opts ...grpc.CallOption) (*GetAllKeysByAdminDescResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllKeysByAdminDescResponse)
	err := c.cc.Invoke(ctx, KeyService_GetAllKeysByAdminDesc_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeyServiceServer is the server API for KeyService service.
// All implementations must embed UnimplementedKeyServiceServer
// for forward compatibility.
//
// 定义 Key 服务
type KeyServiceServer interface {
	// 获取用户所有的key 包含有效和无效的 按照时间和是否有效升序排序
	GetAllMyKeysByUserIdAsc(context.Context, *GetAllMyKeysByUserIdAscRequest) (*GetAllMyKeysByUserIdAscResponse, error)
	// 獲取單個密鑰的信息
	GetKeyInfoById(context.Context, *GetKeyInfoByIdRequest) (*GetKeyInfoByIdResponse, error)
	// 第三方API 激活密钥
	RegisterClientWithKey(context.Context, *RegisterClientWithKeyRequest) (*RegisterClientWithKeyResponse, error)
	// 第三方API 取消激活密钥
	CancelRegisterClientWithKey(context.Context, *CancelRegisterClientWithKeyRequest) (*CancelRegisterClientWithKeyResponse, error)
	// 密钥使用记录
	BindOrActiveMyKey2App(context.Context, *BindOrActiveMyKey2AppRequest) (*BindOrActiveMyKey2AppResponse, error)
	UnbindKey(context.Context, *UnbindKeyRequest) (*UnbindKeyResponse, error)
	AlterKeyBindRemark(context.Context, *AlterKeyBindRemarkRequest) (*AlterKeyBindRemarkResponse, error)
	CheckAccountAndKeyUsability(context.Context, *CheckAccountAndKeyUsabilityRequest) (*CheckAccountAndKeyUsabilityResponse, error)
	GetActivateLogByUserId(context.Context, *GetActivateLogByUserIdRequest) (*GetActivateLogByUserIdResponse, error)
	GetActivateLogByAdmin(context.Context, *GetActivateLogByAdminRequest) (*GetActivateLogByAdminResponse, error)
	AlterKeyValidInfoByAdmin(context.Context, *AlterKeyValidInfoByAdminRequest) (*AlterKeyValidInfoByAdminResponse, error)
	// Admin
	GetAllKeysByAdminDesc(context.Context, *GetAllKeysByAdminDescRequest) (*GetAllKeysByAdminDescResponse, error)
	mustEmbedUnimplementedKeyServiceServer()
}

// UnimplementedKeyServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedKeyServiceServer struct{}

func (UnimplementedKeyServiceServer) GetAllMyKeysByUserIdAsc(context.Context, *GetAllMyKeysByUserIdAscRequest) (*GetAllMyKeysByUserIdAscResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllMyKeysByUserIdAsc not implemented")
}
func (UnimplementedKeyServiceServer) GetKeyInfoById(context.Context, *GetKeyInfoByIdRequest) (*GetKeyInfoByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKeyInfoById not implemented")
}
func (UnimplementedKeyServiceServer) RegisterClientWithKey(context.Context, *RegisterClientWithKeyRequest) (*RegisterClientWithKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterClientWithKey not implemented")
}
func (UnimplementedKeyServiceServer) CancelRegisterClientWithKey(context.Context, *CancelRegisterClientWithKeyRequest) (*CancelRegisterClientWithKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelRegisterClientWithKey not implemented")
}
func (UnimplementedKeyServiceServer) BindOrActiveMyKey2App(context.Context, *BindOrActiveMyKey2AppRequest) (*BindOrActiveMyKey2AppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BindOrActiveMyKey2App not implemented")
}
func (UnimplementedKeyServiceServer) UnbindKey(context.Context, *UnbindKeyRequest) (*UnbindKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnbindKey not implemented")
}
func (UnimplementedKeyServiceServer) AlterKeyBindRemark(context.Context, *AlterKeyBindRemarkRequest) (*AlterKeyBindRemarkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlterKeyBindRemark not implemented")
}
func (UnimplementedKeyServiceServer) CheckAccountAndKeyUsability(context.Context, *CheckAccountAndKeyUsabilityRequest) (*CheckAccountAndKeyUsabilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAccountAndKeyUsability not implemented")
}
func (UnimplementedKeyServiceServer) GetActivateLogByUserId(context.Context, *GetActivateLogByUserIdRequest) (*GetActivateLogByUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActivateLogByUserId not implemented")
}
func (UnimplementedKeyServiceServer) GetActivateLogByAdmin(context.Context, *GetActivateLogByAdminRequest) (*GetActivateLogByAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActivateLogByAdmin not implemented")
}
func (UnimplementedKeyServiceServer) AlterKeyValidInfoByAdmin(context.Context, *AlterKeyValidInfoByAdminRequest) (*AlterKeyValidInfoByAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlterKeyValidInfoByAdmin not implemented")
}
func (UnimplementedKeyServiceServer) GetAllKeysByAdminDesc(context.Context, *GetAllKeysByAdminDescRequest) (*GetAllKeysByAdminDescResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllKeysByAdminDesc not implemented")
}
func (UnimplementedKeyServiceServer) mustEmbedUnimplementedKeyServiceServer() {}
func (UnimplementedKeyServiceServer) testEmbeddedByValue()                    {}

// UnsafeKeyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KeyServiceServer will
// result in compilation errors.
type UnsafeKeyServiceServer interface {
	mustEmbedUnimplementedKeyServiceServer()
}

func RegisterKeyServiceServer(s grpc.ServiceRegistrar, srv KeyServiceServer) {
	// If the following call pancis, it indicates UnimplementedKeyServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&KeyService_ServiceDesc, srv)
}

func _KeyService_GetAllMyKeysByUserIdAsc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllMyKeysByUserIdAscRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyServiceServer).GetAllMyKeysByUserIdAsc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeyService_GetAllMyKeysByUserIdAsc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyServiceServer).GetAllMyKeysByUserIdAsc(ctx, req.(*GetAllMyKeysByUserIdAscRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyService_GetKeyInfoById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKeyInfoByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyServiceServer).GetKeyInfoById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeyService_GetKeyInfoById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyServiceServer).GetKeyInfoById(ctx, req.(*GetKeyInfoByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyService_RegisterClientWithKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterClientWithKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyServiceServer).RegisterClientWithKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeyService_RegisterClientWithKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyServiceServer).RegisterClientWithKey(ctx, req.(*RegisterClientWithKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyService_CancelRegisterClientWithKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelRegisterClientWithKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyServiceServer).CancelRegisterClientWithKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeyService_CancelRegisterClientWithKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyServiceServer).CancelRegisterClientWithKey(ctx, req.(*CancelRegisterClientWithKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyService_BindOrActiveMyKey2App_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BindOrActiveMyKey2AppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyServiceServer).BindOrActiveMyKey2App(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeyService_BindOrActiveMyKey2App_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyServiceServer).BindOrActiveMyKey2App(ctx, req.(*BindOrActiveMyKey2AppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyService_UnbindKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnbindKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyServiceServer).UnbindKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeyService_UnbindKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyServiceServer).UnbindKey(ctx, req.(*UnbindKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyService_AlterKeyBindRemark_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlterKeyBindRemarkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyServiceServer).AlterKeyBindRemark(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeyService_AlterKeyBindRemark_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyServiceServer).AlterKeyBindRemark(ctx, req.(*AlterKeyBindRemarkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyService_CheckAccountAndKeyUsability_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckAccountAndKeyUsabilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyServiceServer).CheckAccountAndKeyUsability(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeyService_CheckAccountAndKeyUsability_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyServiceServer).CheckAccountAndKeyUsability(ctx, req.(*CheckAccountAndKeyUsabilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyService_GetActivateLogByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActivateLogByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyServiceServer).GetActivateLogByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeyService_GetActivateLogByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyServiceServer).GetActivateLogByUserId(ctx, req.(*GetActivateLogByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyService_GetActivateLogByAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActivateLogByAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyServiceServer).GetActivateLogByAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeyService_GetActivateLogByAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyServiceServer).GetActivateLogByAdmin(ctx, req.(*GetActivateLogByAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyService_AlterKeyValidInfoByAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlterKeyValidInfoByAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyServiceServer).AlterKeyValidInfoByAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeyService_AlterKeyValidInfoByAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyServiceServer).AlterKeyValidInfoByAdmin(ctx, req.(*AlterKeyValidInfoByAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyService_GetAllKeysByAdminDesc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllKeysByAdminDescRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyServiceServer).GetAllKeysByAdminDesc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeyService_GetAllKeysByAdminDesc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyServiceServer).GetAllKeysByAdminDesc(ctx, req.(*GetAllKeysByAdminDescRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KeyService_ServiceDesc is the grpc.ServiceDesc for KeyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KeyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "key.KeyService",
	HandlerType: (*KeyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllMyKeysByUserIdAsc",
			Handler:    _KeyService_GetAllMyKeysByUserIdAsc_Handler,
		},
		{
			MethodName: "GetKeyInfoById",
			Handler:    _KeyService_GetKeyInfoById_Handler,
		},
		{
			MethodName: "RegisterClientWithKey",
			Handler:    _KeyService_RegisterClientWithKey_Handler,
		},
		{
			MethodName: "CancelRegisterClientWithKey",
			Handler:    _KeyService_CancelRegisterClientWithKey_Handler,
		},
		{
			MethodName: "BindOrActiveMyKey2App",
			Handler:    _KeyService_BindOrActiveMyKey2App_Handler,
		},
		{
			MethodName: "UnbindKey",
			Handler:    _KeyService_UnbindKey_Handler,
		},
		{
			MethodName: "AlterKeyBindRemark",
			Handler:    _KeyService_AlterKeyBindRemark_Handler,
		},
		{
			MethodName: "CheckAccountAndKeyUsability",
			Handler:    _KeyService_CheckAccountAndKeyUsability_Handler,
		},
		{
			MethodName: "GetActivateLogByUserId",
			Handler:    _KeyService_GetActivateLogByUserId_Handler,
		},
		{
			MethodName: "GetActivateLogByAdmin",
			Handler:    _KeyService_GetActivateLogByAdmin_Handler,
		},
		{
			MethodName: "AlterKeyValidInfoByAdmin",
			Handler:    _KeyService_AlterKeyValidInfoByAdmin_Handler,
		},
		{
			MethodName: "GetAllKeysByAdminDesc",
			Handler:    _KeyService_GetAllKeysByAdminDesc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "key.proto",
}
