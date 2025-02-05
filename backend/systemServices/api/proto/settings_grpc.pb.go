// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: settings.proto

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
	SettingsService_UpdateSingleOption_FullMethodName               = "/settings.SettingsService/UpdateSingleOption"
	SettingsService_GetStartTheme_FullMethodName                    = "/settings.SettingsService/GetStartTheme"
	SettingsService_GetSystemSettings_FullMethodName                = "/settings.SettingsService/GetSystemSettings"
	SettingsService_AddPaymentMethod_FullMethodName                 = "/settings.SettingsService/AddPaymentMethod"
	SettingsService_GetAllPaymentSettingsKv_FullMethodName          = "/settings.SettingsService/GetAllPaymentSettingsKv"
	SettingsService_GetPaymentMethodDetailsByName_FullMethodName    = "/settings.SettingsService/GetPaymentMethodDetailsByName"
	SettingsService_EditPaymentSettingsBySystemName_FullMethodName  = "/settings.SettingsService/EditPaymentSettingsBySystemName"
	SettingsService_EnablePaymentSettingBySystemName_FullMethodName = "/settings.SettingsService/EnablePaymentSettingBySystemName"
	SettingsService_DeletePaymentSettingBySystemName_FullMethodName = "/settings.SettingsService/DeletePaymentSettingBySystemName"
	SettingsService_GetInviteUserMsg_FullMethodName                 = "/settings.SettingsService/GetInviteUserMsg"
	SettingsService_GetAdminDashboardData_FullMethodName            = "/settings.SettingsService/GetAdminDashboardData"
	SettingsService_GetBasicRuntimeEnvConfig_FullMethodName         = "/settings.SettingsService/GetBasicRuntimeEnvConfig"
	SettingsService_GetRegisterEnvConfig_FullMethodName             = "/settings.SettingsService/GetRegisterEnvConfig"
	SettingsService_GetWelcomePageConfig_FullMethodName             = "/settings.SettingsService/GetWelcomePageConfig"
	SettingsService_GetAppDownloadLink_FullMethodName               = "/settings.SettingsService/GetAppDownloadLink"
	SettingsService_GetServerLiveStatus_FullMethodName              = "/settings.SettingsService/GetServerLiveStatus"
)

// SettingsServiceClient is the client API for SettingsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 定义 SettingsService 服务
type SettingsServiceClient interface {
	// 更新单个设置项
	UpdateSingleOption(ctx context.Context, in *UpdateSingleOptionRequest, opts ...grpc.CallOption) (*UpdateSingleOptionResponse, error)
	// 获取启动页面主题设置
	GetStartTheme(ctx context.Context, in *GetStartThemeRequest, opts ...grpc.CallOption) (*GetStartThemeResponse, error)
	// 获取所有的设置
	GetSystemSettings(ctx context.Context, in *GetSystemSettingsRequest, opts ...grpc.CallOption) (*GetSystemSettingsResponse, error)
	// pass 添加新的付款方式
	AddPaymentMethod(ctx context.Context, in *AddPaymentMethodRequest, opts ...grpc.CallOption) (*AddPaymentMethodResponse, error)
	// 获取所有支付方式的键值对 名称 是否被启用
	GetAllPaymentSettingsKv(ctx context.Context, in *GetAllPaymentSettingsKvRequest, opts ...grpc.CallOption) (*GetAllPaymentSettingsKvResponse, error)
	// 根据传入的支付方式名称来获取其详细信息
	GetPaymentMethodDetailsByName(ctx context.Context, in *GetPaymentMethodDetailsByNameRequest, opts ...grpc.CallOption) (*GetPaymentMethodDetailsByNameResponse, error)
	// 根据名称来新增或修改支付方式的信息
	EditPaymentSettingsBySystemName(ctx context.Context, in *EditPaymentSettingsBySystemNameRequest, opts ...grpc.CallOption) (*EditPaymentSettingsBySystemNameResponse, error)
	// 是否启用该支付防止 这个调用后取反enabled字段 之后需要刷新Redis缓存和页面
	EnablePaymentSettingBySystemName(ctx context.Context, in *EnablePaymentSettingBySystemNameRequest, opts ...grpc.CallOption) (*EnablePaymentSettingBySystemNameResponse, error)
	// pass 删除付款方式
	DeletePaymentSettingBySystemName(ctx context.Context, in *DeletePaymentSettingBySystemNameRequest, opts ...grpc.CallOption) (*DeletePaymentSettingBySystemNameResponse, error)
	// 杂项
	GetInviteUserMsg(ctx context.Context, in *GetInviteUserMsgRequest, opts ...grpc.CallOption) (*GetInviteUserMsgResponse, error)
	// 管理员首页的图表
	GetAdminDashboardData(ctx context.Context, in *GetAdminDashboardDataRequest, opts ...grpc.CallOption) (*GetAdminDashboardDataResponse, error)
	// 获取整个页面的配置信息
	GetBasicRuntimeEnvConfig(ctx context.Context, in *GetBasicRuntimeEnvConfigRequest, opts ...grpc.CallOption) (*GetBasicRuntimeEnvConfigResponse, error)
	GetRegisterEnvConfig(ctx context.Context, in *GetRegisterEnvConfigRequest, opts ...grpc.CallOption) (*GetRegisterEnvConfigResponse, error)
	GetWelcomePageConfig(ctx context.Context, in *GetWelcomePageConfigRequest, opts ...grpc.CallOption) (*GetWelcomePageConfigResponse, error)
	GetAppDownloadLink(ctx context.Context, in *GetAppDownloadLinkRequest, opts ...grpc.CallOption) (*GetAppDownloadLinkResponse, error)
	GetServerLiveStatus(ctx context.Context, in *GetServerLiveStatusRequest, opts ...grpc.CallOption) (*GetServerLiveStatusResponse, error)
}

type settingsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSettingsServiceClient(cc grpc.ClientConnInterface) SettingsServiceClient {
	return &settingsServiceClient{cc}
}

func (c *settingsServiceClient) UpdateSingleOption(ctx context.Context, in *UpdateSingleOptionRequest, opts ...grpc.CallOption) (*UpdateSingleOptionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateSingleOptionResponse)
	err := c.cc.Invoke(ctx, SettingsService_UpdateSingleOption_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsServiceClient) GetStartTheme(ctx context.Context, in *GetStartThemeRequest, opts ...grpc.CallOption) (*GetStartThemeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetStartThemeResponse)
	err := c.cc.Invoke(ctx, SettingsService_GetStartTheme_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsServiceClient) GetSystemSettings(ctx context.Context, in *GetSystemSettingsRequest, opts ...grpc.CallOption) (*GetSystemSettingsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSystemSettingsResponse)
	err := c.cc.Invoke(ctx, SettingsService_GetSystemSettings_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsServiceClient) AddPaymentMethod(ctx context.Context, in *AddPaymentMethodRequest, opts ...grpc.CallOption) (*AddPaymentMethodResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddPaymentMethodResponse)
	err := c.cc.Invoke(ctx, SettingsService_AddPaymentMethod_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsServiceClient) GetAllPaymentSettingsKv(ctx context.Context, in *GetAllPaymentSettingsKvRequest, opts ...grpc.CallOption) (*GetAllPaymentSettingsKvResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllPaymentSettingsKvResponse)
	err := c.cc.Invoke(ctx, SettingsService_GetAllPaymentSettingsKv_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsServiceClient) GetPaymentMethodDetailsByName(ctx context.Context, in *GetPaymentMethodDetailsByNameRequest, opts ...grpc.CallOption) (*GetPaymentMethodDetailsByNameResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPaymentMethodDetailsByNameResponse)
	err := c.cc.Invoke(ctx, SettingsService_GetPaymentMethodDetailsByName_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsServiceClient) EditPaymentSettingsBySystemName(ctx context.Context, in *EditPaymentSettingsBySystemNameRequest, opts ...grpc.CallOption) (*EditPaymentSettingsBySystemNameResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EditPaymentSettingsBySystemNameResponse)
	err := c.cc.Invoke(ctx, SettingsService_EditPaymentSettingsBySystemName_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsServiceClient) EnablePaymentSettingBySystemName(ctx context.Context, in *EnablePaymentSettingBySystemNameRequest, opts ...grpc.CallOption) (*EnablePaymentSettingBySystemNameResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EnablePaymentSettingBySystemNameResponse)
	err := c.cc.Invoke(ctx, SettingsService_EnablePaymentSettingBySystemName_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsServiceClient) DeletePaymentSettingBySystemName(ctx context.Context, in *DeletePaymentSettingBySystemNameRequest, opts ...grpc.CallOption) (*DeletePaymentSettingBySystemNameResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeletePaymentSettingBySystemNameResponse)
	err := c.cc.Invoke(ctx, SettingsService_DeletePaymentSettingBySystemName_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsServiceClient) GetInviteUserMsg(ctx context.Context, in *GetInviteUserMsgRequest, opts ...grpc.CallOption) (*GetInviteUserMsgResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetInviteUserMsgResponse)
	err := c.cc.Invoke(ctx, SettingsService_GetInviteUserMsg_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsServiceClient) GetAdminDashboardData(ctx context.Context, in *GetAdminDashboardDataRequest, opts ...grpc.CallOption) (*GetAdminDashboardDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAdminDashboardDataResponse)
	err := c.cc.Invoke(ctx, SettingsService_GetAdminDashboardData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsServiceClient) GetBasicRuntimeEnvConfig(ctx context.Context, in *GetBasicRuntimeEnvConfigRequest, opts ...grpc.CallOption) (*GetBasicRuntimeEnvConfigResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBasicRuntimeEnvConfigResponse)
	err := c.cc.Invoke(ctx, SettingsService_GetBasicRuntimeEnvConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsServiceClient) GetRegisterEnvConfig(ctx context.Context, in *GetRegisterEnvConfigRequest, opts ...grpc.CallOption) (*GetRegisterEnvConfigResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRegisterEnvConfigResponse)
	err := c.cc.Invoke(ctx, SettingsService_GetRegisterEnvConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsServiceClient) GetWelcomePageConfig(ctx context.Context, in *GetWelcomePageConfigRequest, opts ...grpc.CallOption) (*GetWelcomePageConfigResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetWelcomePageConfigResponse)
	err := c.cc.Invoke(ctx, SettingsService_GetWelcomePageConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsServiceClient) GetAppDownloadLink(ctx context.Context, in *GetAppDownloadLinkRequest, opts ...grpc.CallOption) (*GetAppDownloadLinkResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAppDownloadLinkResponse)
	err := c.cc.Invoke(ctx, SettingsService_GetAppDownloadLink_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsServiceClient) GetServerLiveStatus(ctx context.Context, in *GetServerLiveStatusRequest, opts ...grpc.CallOption) (*GetServerLiveStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetServerLiveStatusResponse)
	err := c.cc.Invoke(ctx, SettingsService_GetServerLiveStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SettingsServiceServer is the server API for SettingsService service.
// All implementations must embed UnimplementedSettingsServiceServer
// for forward compatibility.
//
// 定义 SettingsService 服务
type SettingsServiceServer interface {
	// 更新单个设置项
	UpdateSingleOption(context.Context, *UpdateSingleOptionRequest) (*UpdateSingleOptionResponse, error)
	// 获取启动页面主题设置
	GetStartTheme(context.Context, *GetStartThemeRequest) (*GetStartThemeResponse, error)
	// 获取所有的设置
	GetSystemSettings(context.Context, *GetSystemSettingsRequest) (*GetSystemSettingsResponse, error)
	// pass 添加新的付款方式
	AddPaymentMethod(context.Context, *AddPaymentMethodRequest) (*AddPaymentMethodResponse, error)
	// 获取所有支付方式的键值对 名称 是否被启用
	GetAllPaymentSettingsKv(context.Context, *GetAllPaymentSettingsKvRequest) (*GetAllPaymentSettingsKvResponse, error)
	// 根据传入的支付方式名称来获取其详细信息
	GetPaymentMethodDetailsByName(context.Context, *GetPaymentMethodDetailsByNameRequest) (*GetPaymentMethodDetailsByNameResponse, error)
	// 根据名称来新增或修改支付方式的信息
	EditPaymentSettingsBySystemName(context.Context, *EditPaymentSettingsBySystemNameRequest) (*EditPaymentSettingsBySystemNameResponse, error)
	// 是否启用该支付防止 这个调用后取反enabled字段 之后需要刷新Redis缓存和页面
	EnablePaymentSettingBySystemName(context.Context, *EnablePaymentSettingBySystemNameRequest) (*EnablePaymentSettingBySystemNameResponse, error)
	// pass 删除付款方式
	DeletePaymentSettingBySystemName(context.Context, *DeletePaymentSettingBySystemNameRequest) (*DeletePaymentSettingBySystemNameResponse, error)
	// 杂项
	GetInviteUserMsg(context.Context, *GetInviteUserMsgRequest) (*GetInviteUserMsgResponse, error)
	// 管理员首页的图表
	GetAdminDashboardData(context.Context, *GetAdminDashboardDataRequest) (*GetAdminDashboardDataResponse, error)
	// 获取整个页面的配置信息
	GetBasicRuntimeEnvConfig(context.Context, *GetBasicRuntimeEnvConfigRequest) (*GetBasicRuntimeEnvConfigResponse, error)
	GetRegisterEnvConfig(context.Context, *GetRegisterEnvConfigRequest) (*GetRegisterEnvConfigResponse, error)
	GetWelcomePageConfig(context.Context, *GetWelcomePageConfigRequest) (*GetWelcomePageConfigResponse, error)
	GetAppDownloadLink(context.Context, *GetAppDownloadLinkRequest) (*GetAppDownloadLinkResponse, error)
	GetServerLiveStatus(context.Context, *GetServerLiveStatusRequest) (*GetServerLiveStatusResponse, error)
	mustEmbedUnimplementedSettingsServiceServer()
}

// UnimplementedSettingsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSettingsServiceServer struct{}

func (UnimplementedSettingsServiceServer) UpdateSingleOption(context.Context, *UpdateSingleOptionRequest) (*UpdateSingleOptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSingleOption not implemented")
}
func (UnimplementedSettingsServiceServer) GetStartTheme(context.Context, *GetStartThemeRequest) (*GetStartThemeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStartTheme not implemented")
}
func (UnimplementedSettingsServiceServer) GetSystemSettings(context.Context, *GetSystemSettingsRequest) (*GetSystemSettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSystemSettings not implemented")
}
func (UnimplementedSettingsServiceServer) AddPaymentMethod(context.Context, *AddPaymentMethodRequest) (*AddPaymentMethodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPaymentMethod not implemented")
}
func (UnimplementedSettingsServiceServer) GetAllPaymentSettingsKv(context.Context, *GetAllPaymentSettingsKvRequest) (*GetAllPaymentSettingsKvResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPaymentSettingsKv not implemented")
}
func (UnimplementedSettingsServiceServer) GetPaymentMethodDetailsByName(context.Context, *GetPaymentMethodDetailsByNameRequest) (*GetPaymentMethodDetailsByNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPaymentMethodDetailsByName not implemented")
}
func (UnimplementedSettingsServiceServer) EditPaymentSettingsBySystemName(context.Context, *EditPaymentSettingsBySystemNameRequest) (*EditPaymentSettingsBySystemNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditPaymentSettingsBySystemName not implemented")
}
func (UnimplementedSettingsServiceServer) EnablePaymentSettingBySystemName(context.Context, *EnablePaymentSettingBySystemNameRequest) (*EnablePaymentSettingBySystemNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnablePaymentSettingBySystemName not implemented")
}
func (UnimplementedSettingsServiceServer) DeletePaymentSettingBySystemName(context.Context, *DeletePaymentSettingBySystemNameRequest) (*DeletePaymentSettingBySystemNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePaymentSettingBySystemName not implemented")
}
func (UnimplementedSettingsServiceServer) GetInviteUserMsg(context.Context, *GetInviteUserMsgRequest) (*GetInviteUserMsgResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInviteUserMsg not implemented")
}
func (UnimplementedSettingsServiceServer) GetAdminDashboardData(context.Context, *GetAdminDashboardDataRequest) (*GetAdminDashboardDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdminDashboardData not implemented")
}
func (UnimplementedSettingsServiceServer) GetBasicRuntimeEnvConfig(context.Context, *GetBasicRuntimeEnvConfigRequest) (*GetBasicRuntimeEnvConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBasicRuntimeEnvConfig not implemented")
}
func (UnimplementedSettingsServiceServer) GetRegisterEnvConfig(context.Context, *GetRegisterEnvConfigRequest) (*GetRegisterEnvConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRegisterEnvConfig not implemented")
}
func (UnimplementedSettingsServiceServer) GetWelcomePageConfig(context.Context, *GetWelcomePageConfigRequest) (*GetWelcomePageConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWelcomePageConfig not implemented")
}
func (UnimplementedSettingsServiceServer) GetAppDownloadLink(context.Context, *GetAppDownloadLinkRequest) (*GetAppDownloadLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppDownloadLink not implemented")
}
func (UnimplementedSettingsServiceServer) GetServerLiveStatus(context.Context, *GetServerLiveStatusRequest) (*GetServerLiveStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerLiveStatus not implemented")
}
func (UnimplementedSettingsServiceServer) mustEmbedUnimplementedSettingsServiceServer() {}
func (UnimplementedSettingsServiceServer) testEmbeddedByValue()                         {}

// UnsafeSettingsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SettingsServiceServer will
// result in compilation errors.
type UnsafeSettingsServiceServer interface {
	mustEmbedUnimplementedSettingsServiceServer()
}

func RegisterSettingsServiceServer(s grpc.ServiceRegistrar, srv SettingsServiceServer) {
	// If the following call pancis, it indicates UnimplementedSettingsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SettingsService_ServiceDesc, srv)
}

func _SettingsService_UpdateSingleOption_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSingleOptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServiceServer).UpdateSingleOption(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SettingsService_UpdateSingleOption_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServiceServer).UpdateSingleOption(ctx, req.(*UpdateSingleOptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SettingsService_GetStartTheme_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStartThemeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServiceServer).GetStartTheme(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SettingsService_GetStartTheme_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServiceServer).GetStartTheme(ctx, req.(*GetStartThemeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SettingsService_GetSystemSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSystemSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServiceServer).GetSystemSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SettingsService_GetSystemSettings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServiceServer).GetSystemSettings(ctx, req.(*GetSystemSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SettingsService_AddPaymentMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPaymentMethodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServiceServer).AddPaymentMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SettingsService_AddPaymentMethod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServiceServer).AddPaymentMethod(ctx, req.(*AddPaymentMethodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SettingsService_GetAllPaymentSettingsKv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllPaymentSettingsKvRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServiceServer).GetAllPaymentSettingsKv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SettingsService_GetAllPaymentSettingsKv_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServiceServer).GetAllPaymentSettingsKv(ctx, req.(*GetAllPaymentSettingsKvRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SettingsService_GetPaymentMethodDetailsByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPaymentMethodDetailsByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServiceServer).GetPaymentMethodDetailsByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SettingsService_GetPaymentMethodDetailsByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServiceServer).GetPaymentMethodDetailsByName(ctx, req.(*GetPaymentMethodDetailsByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SettingsService_EditPaymentSettingsBySystemName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditPaymentSettingsBySystemNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServiceServer).EditPaymentSettingsBySystemName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SettingsService_EditPaymentSettingsBySystemName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServiceServer).EditPaymentSettingsBySystemName(ctx, req.(*EditPaymentSettingsBySystemNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SettingsService_EnablePaymentSettingBySystemName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnablePaymentSettingBySystemNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServiceServer).EnablePaymentSettingBySystemName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SettingsService_EnablePaymentSettingBySystemName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServiceServer).EnablePaymentSettingBySystemName(ctx, req.(*EnablePaymentSettingBySystemNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SettingsService_DeletePaymentSettingBySystemName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePaymentSettingBySystemNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServiceServer).DeletePaymentSettingBySystemName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SettingsService_DeletePaymentSettingBySystemName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServiceServer).DeletePaymentSettingBySystemName(ctx, req.(*DeletePaymentSettingBySystemNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SettingsService_GetInviteUserMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInviteUserMsgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServiceServer).GetInviteUserMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SettingsService_GetInviteUserMsg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServiceServer).GetInviteUserMsg(ctx, req.(*GetInviteUserMsgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SettingsService_GetAdminDashboardData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdminDashboardDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServiceServer).GetAdminDashboardData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SettingsService_GetAdminDashboardData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServiceServer).GetAdminDashboardData(ctx, req.(*GetAdminDashboardDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SettingsService_GetBasicRuntimeEnvConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBasicRuntimeEnvConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServiceServer).GetBasicRuntimeEnvConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SettingsService_GetBasicRuntimeEnvConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServiceServer).GetBasicRuntimeEnvConfig(ctx, req.(*GetBasicRuntimeEnvConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SettingsService_GetRegisterEnvConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRegisterEnvConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServiceServer).GetRegisterEnvConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SettingsService_GetRegisterEnvConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServiceServer).GetRegisterEnvConfig(ctx, req.(*GetRegisterEnvConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SettingsService_GetWelcomePageConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWelcomePageConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServiceServer).GetWelcomePageConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SettingsService_GetWelcomePageConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServiceServer).GetWelcomePageConfig(ctx, req.(*GetWelcomePageConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SettingsService_GetAppDownloadLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppDownloadLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServiceServer).GetAppDownloadLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SettingsService_GetAppDownloadLink_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServiceServer).GetAppDownloadLink(ctx, req.(*GetAppDownloadLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SettingsService_GetServerLiveStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServerLiveStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServiceServer).GetServerLiveStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SettingsService_GetServerLiveStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServiceServer).GetServerLiveStatus(ctx, req.(*GetServerLiveStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SettingsService_ServiceDesc is the grpc.ServiceDesc for SettingsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SettingsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "settings.SettingsService",
	HandlerType: (*SettingsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateSingleOption",
			Handler:    _SettingsService_UpdateSingleOption_Handler,
		},
		{
			MethodName: "GetStartTheme",
			Handler:    _SettingsService_GetStartTheme_Handler,
		},
		{
			MethodName: "GetSystemSettings",
			Handler:    _SettingsService_GetSystemSettings_Handler,
		},
		{
			MethodName: "AddPaymentMethod",
			Handler:    _SettingsService_AddPaymentMethod_Handler,
		},
		{
			MethodName: "GetAllPaymentSettingsKv",
			Handler:    _SettingsService_GetAllPaymentSettingsKv_Handler,
		},
		{
			MethodName: "GetPaymentMethodDetailsByName",
			Handler:    _SettingsService_GetPaymentMethodDetailsByName_Handler,
		},
		{
			MethodName: "EditPaymentSettingsBySystemName",
			Handler:    _SettingsService_EditPaymentSettingsBySystemName_Handler,
		},
		{
			MethodName: "EnablePaymentSettingBySystemName",
			Handler:    _SettingsService_EnablePaymentSettingBySystemName_Handler,
		},
		{
			MethodName: "DeletePaymentSettingBySystemName",
			Handler:    _SettingsService_DeletePaymentSettingBySystemName_Handler,
		},
		{
			MethodName: "GetInviteUserMsg",
			Handler:    _SettingsService_GetInviteUserMsg_Handler,
		},
		{
			MethodName: "GetAdminDashboardData",
			Handler:    _SettingsService_GetAdminDashboardData_Handler,
		},
		{
			MethodName: "GetBasicRuntimeEnvConfig",
			Handler:    _SettingsService_GetBasicRuntimeEnvConfig_Handler,
		},
		{
			MethodName: "GetRegisterEnvConfig",
			Handler:    _SettingsService_GetRegisterEnvConfig_Handler,
		},
		{
			MethodName: "GetWelcomePageConfig",
			Handler:    _SettingsService_GetWelcomePageConfig_Handler,
		},
		{
			MethodName: "GetAppDownloadLink",
			Handler:    _SettingsService_GetAppDownloadLink_Handler,
		},
		{
			MethodName: "GetServerLiveStatus",
			Handler:    _SettingsService_GetServerLiveStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "settings.proto",
}
