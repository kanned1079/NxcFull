// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: coupon.proto

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
	CouponService_AddNewCoupon_FullMethodName   = "/coupon.CouponService/AddNewCoupon"
	CouponService_VerifyCoupon_FullMethodName   = "/coupon.CouponService/VerifyCoupon"
	CouponService_GetAllCoupons_FullMethodName  = "/coupon.CouponService/GetAllCoupons"
	CouponService_ActivateCoupon_FullMethodName = "/coupon.CouponService/ActivateCoupon"
	CouponService_DeleteCoupon_FullMethodName   = "/coupon.CouponService/DeleteCoupon"
	CouponService_UpdateCoupon_FullMethodName   = "/coupon.CouponService/UpdateCoupon"
)

// CouponServiceClient is the client API for CouponService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 定义 gRPC 服务
type CouponServiceClient interface {
	// 新增优惠券
	AddNewCoupon(ctx context.Context, in *AddNewCouponRequest, opts ...grpc.CallOption) (*AddNewCouponResponse, error)
	// 验证优惠券
	VerifyCoupon(ctx context.Context, in *VerifyCouponRequest, opts ...grpc.CallOption) (*VerifyCouponResponse, error)
	// 获取所有优惠券
	GetAllCoupons(ctx context.Context, in *GetAllCouponsRequest, opts ...grpc.CallOption) (*GetAllCouponsResponse, error)
	// 启用/禁用优惠券
	ActivateCoupon(ctx context.Context, in *ActivateCouponRequest, opts ...grpc.CallOption) (*ActivateCouponResponse, error)
	// 删除优惠券
	DeleteCoupon(ctx context.Context, in *DeleteCouponRequest, opts ...grpc.CallOption) (*DeleteCouponResponse, error)
	UpdateCoupon(ctx context.Context, in *UpdateCouponRequest, opts ...grpc.CallOption) (*UpdateCouponResponse, error)
}

type couponServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCouponServiceClient(cc grpc.ClientConnInterface) CouponServiceClient {
	return &couponServiceClient{cc}
}

func (c *couponServiceClient) AddNewCoupon(ctx context.Context, in *AddNewCouponRequest, opts ...grpc.CallOption) (*AddNewCouponResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddNewCouponResponse)
	err := c.cc.Invoke(ctx, CouponService_AddNewCoupon_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *couponServiceClient) VerifyCoupon(ctx context.Context, in *VerifyCouponRequest, opts ...grpc.CallOption) (*VerifyCouponResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VerifyCouponResponse)
	err := c.cc.Invoke(ctx, CouponService_VerifyCoupon_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *couponServiceClient) GetAllCoupons(ctx context.Context, in *GetAllCouponsRequest, opts ...grpc.CallOption) (*GetAllCouponsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllCouponsResponse)
	err := c.cc.Invoke(ctx, CouponService_GetAllCoupons_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *couponServiceClient) ActivateCoupon(ctx context.Context, in *ActivateCouponRequest, opts ...grpc.CallOption) (*ActivateCouponResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ActivateCouponResponse)
	err := c.cc.Invoke(ctx, CouponService_ActivateCoupon_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *couponServiceClient) DeleteCoupon(ctx context.Context, in *DeleteCouponRequest, opts ...grpc.CallOption) (*DeleteCouponResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteCouponResponse)
	err := c.cc.Invoke(ctx, CouponService_DeleteCoupon_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *couponServiceClient) UpdateCoupon(ctx context.Context, in *UpdateCouponRequest, opts ...grpc.CallOption) (*UpdateCouponResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateCouponResponse)
	err := c.cc.Invoke(ctx, CouponService_UpdateCoupon_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CouponServiceServer is the server API for CouponService service.
// All implementations must embed UnimplementedCouponServiceServer
// for forward compatibility.
//
// 定义 gRPC 服务
type CouponServiceServer interface {
	// 新增优惠券
	AddNewCoupon(context.Context, *AddNewCouponRequest) (*AddNewCouponResponse, error)
	// 验证优惠券
	VerifyCoupon(context.Context, *VerifyCouponRequest) (*VerifyCouponResponse, error)
	// 获取所有优惠券
	GetAllCoupons(context.Context, *GetAllCouponsRequest) (*GetAllCouponsResponse, error)
	// 启用/禁用优惠券
	ActivateCoupon(context.Context, *ActivateCouponRequest) (*ActivateCouponResponse, error)
	// 删除优惠券
	DeleteCoupon(context.Context, *DeleteCouponRequest) (*DeleteCouponResponse, error)
	UpdateCoupon(context.Context, *UpdateCouponRequest) (*UpdateCouponResponse, error)
	mustEmbedUnimplementedCouponServiceServer()
}

// UnimplementedCouponServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCouponServiceServer struct{}

func (UnimplementedCouponServiceServer) AddNewCoupon(context.Context, *AddNewCouponRequest) (*AddNewCouponResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNewCoupon not implemented")
}
func (UnimplementedCouponServiceServer) VerifyCoupon(context.Context, *VerifyCouponRequest) (*VerifyCouponResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyCoupon not implemented")
}
func (UnimplementedCouponServiceServer) GetAllCoupons(context.Context, *GetAllCouponsRequest) (*GetAllCouponsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCoupons not implemented")
}
func (UnimplementedCouponServiceServer) ActivateCoupon(context.Context, *ActivateCouponRequest) (*ActivateCouponResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivateCoupon not implemented")
}
func (UnimplementedCouponServiceServer) DeleteCoupon(context.Context, *DeleteCouponRequest) (*DeleteCouponResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCoupon not implemented")
}
func (UnimplementedCouponServiceServer) UpdateCoupon(context.Context, *UpdateCouponRequest) (*UpdateCouponResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCoupon not implemented")
}
func (UnimplementedCouponServiceServer) mustEmbedUnimplementedCouponServiceServer() {}
func (UnimplementedCouponServiceServer) testEmbeddedByValue()                       {}

// UnsafeCouponServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CouponServiceServer will
// result in compilation errors.
type UnsafeCouponServiceServer interface {
	mustEmbedUnimplementedCouponServiceServer()
}

func RegisterCouponServiceServer(s grpc.ServiceRegistrar, srv CouponServiceServer) {
	// If the following call pancis, it indicates UnimplementedCouponServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CouponService_ServiceDesc, srv)
}

func _CouponService_AddNewCoupon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddNewCouponRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CouponServiceServer).AddNewCoupon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CouponService_AddNewCoupon_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CouponServiceServer).AddNewCoupon(ctx, req.(*AddNewCouponRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CouponService_VerifyCoupon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyCouponRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CouponServiceServer).VerifyCoupon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CouponService_VerifyCoupon_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CouponServiceServer).VerifyCoupon(ctx, req.(*VerifyCouponRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CouponService_GetAllCoupons_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllCouponsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CouponServiceServer).GetAllCoupons(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CouponService_GetAllCoupons_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CouponServiceServer).GetAllCoupons(ctx, req.(*GetAllCouponsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CouponService_ActivateCoupon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivateCouponRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CouponServiceServer).ActivateCoupon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CouponService_ActivateCoupon_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CouponServiceServer).ActivateCoupon(ctx, req.(*ActivateCouponRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CouponService_DeleteCoupon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCouponRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CouponServiceServer).DeleteCoupon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CouponService_DeleteCoupon_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CouponServiceServer).DeleteCoupon(ctx, req.(*DeleteCouponRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CouponService_UpdateCoupon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCouponRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CouponServiceServer).UpdateCoupon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CouponService_UpdateCoupon_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CouponServiceServer).UpdateCoupon(ctx, req.(*UpdateCouponRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CouponService_ServiceDesc is the grpc.ServiceDesc for CouponService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CouponService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "coupon.CouponService",
	HandlerType: (*CouponServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddNewCoupon",
			Handler:    _CouponService_AddNewCoupon_Handler,
		},
		{
			MethodName: "VerifyCoupon",
			Handler:    _CouponService_VerifyCoupon_Handler,
		},
		{
			MethodName: "GetAllCoupons",
			Handler:    _CouponService_GetAllCoupons_Handler,
		},
		{
			MethodName: "ActivateCoupon",
			Handler:    _CouponService_ActivateCoupon_Handler,
		},
		{
			MethodName: "DeleteCoupon",
			Handler:    _CouponService_DeleteCoupon_Handler,
		},
		{
			MethodName: "UpdateCoupon",
			Handler:    _CouponService_UpdateCoupon_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "coupon.proto",
}
