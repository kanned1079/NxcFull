// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: order_handle.proto

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
	OrderHandleService_GetOrderStatus_FullMethodName           = "/orderHandle.OrderHandleService/GetOrderStatus"
	OrderHandleService_CancelOrder_FullMethodName              = "/orderHandle.OrderHandleService/CancelOrder"
	OrderHandleService_PlaceOrder_FullMethodName               = "/orderHandle.OrderHandleService/PlaceOrder"
	OrderHandleService_ManualPassOrderPayment_FullMethodName   = "/orderHandle.OrderHandleService/ManualPassOrderPayment"
	OrderHandleService_ManualCancelOrderPayment_FullMethodName = "/orderHandle.OrderHandleService/ManualCancelOrderPayment"
)

// OrderHandleServiceClient is the client API for OrderHandleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 定义 Order 服务
type OrderHandleServiceClient interface {
	// 查询订单是否成功 轮询请求
	GetOrderStatus(ctx context.Context, in *GetOrderStatusRequest, opts ...grpc.CallOption) (*GetOrderStatusResponse, error)
	// 取消订单
	CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*CancelOrderResponse, error)
	// 确认订单
	PlaceOrder(ctx context.Context, in *PlaceOrderRequest, opts ...grpc.CallOption) (*PlaceOrderResponse, error)
	ManualPassOrderPayment(ctx context.Context, in *ManualPassOrderPaymentRequest, opts ...grpc.CallOption) (*ManualPassOrderPaymentResponse, error)
	ManualCancelOrderPayment(ctx context.Context, in *ManualCancelOrderPaymentRequest, opts ...grpc.CallOption) (*ManualCancelOrderPaymentResponse, error)
}

type orderHandleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderHandleServiceClient(cc grpc.ClientConnInterface) OrderHandleServiceClient {
	return &orderHandleServiceClient{cc}
}

func (c *orderHandleServiceClient) GetOrderStatus(ctx context.Context, in *GetOrderStatusRequest, opts ...grpc.CallOption) (*GetOrderStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetOrderStatusResponse)
	err := c.cc.Invoke(ctx, OrderHandleService_GetOrderStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderHandleServiceClient) CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*CancelOrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CancelOrderResponse)
	err := c.cc.Invoke(ctx, OrderHandleService_CancelOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderHandleServiceClient) PlaceOrder(ctx context.Context, in *PlaceOrderRequest, opts ...grpc.CallOption) (*PlaceOrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PlaceOrderResponse)
	err := c.cc.Invoke(ctx, OrderHandleService_PlaceOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderHandleServiceClient) ManualPassOrderPayment(ctx context.Context, in *ManualPassOrderPaymentRequest, opts ...grpc.CallOption) (*ManualPassOrderPaymentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ManualPassOrderPaymentResponse)
	err := c.cc.Invoke(ctx, OrderHandleService_ManualPassOrderPayment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderHandleServiceClient) ManualCancelOrderPayment(ctx context.Context, in *ManualCancelOrderPaymentRequest, opts ...grpc.CallOption) (*ManualCancelOrderPaymentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ManualCancelOrderPaymentResponse)
	err := c.cc.Invoke(ctx, OrderHandleService_ManualCancelOrderPayment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderHandleServiceServer is the server API for OrderHandleService service.
// All implementations must embed UnimplementedOrderHandleServiceServer
// for forward compatibility.
//
// 定义 Order 服务
type OrderHandleServiceServer interface {
	// 查询订单是否成功 轮询请求
	GetOrderStatus(context.Context, *GetOrderStatusRequest) (*GetOrderStatusResponse, error)
	// 取消订单
	CancelOrder(context.Context, *CancelOrderRequest) (*CancelOrderResponse, error)
	// 确认订单
	PlaceOrder(context.Context, *PlaceOrderRequest) (*PlaceOrderResponse, error)
	ManualPassOrderPayment(context.Context, *ManualPassOrderPaymentRequest) (*ManualPassOrderPaymentResponse, error)
	ManualCancelOrderPayment(context.Context, *ManualCancelOrderPaymentRequest) (*ManualCancelOrderPaymentResponse, error)
	mustEmbedUnimplementedOrderHandleServiceServer()
}

// UnimplementedOrderHandleServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedOrderHandleServiceServer struct{}

func (UnimplementedOrderHandleServiceServer) GetOrderStatus(context.Context, *GetOrderStatusRequest) (*GetOrderStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderStatus not implemented")
}
func (UnimplementedOrderHandleServiceServer) CancelOrder(context.Context, *CancelOrderRequest) (*CancelOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOrder not implemented")
}
func (UnimplementedOrderHandleServiceServer) PlaceOrder(context.Context, *PlaceOrderRequest) (*PlaceOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaceOrder not implemented")
}
func (UnimplementedOrderHandleServiceServer) ManualPassOrderPayment(context.Context, *ManualPassOrderPaymentRequest) (*ManualPassOrderPaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ManualPassOrderPayment not implemented")
}
func (UnimplementedOrderHandleServiceServer) ManualCancelOrderPayment(context.Context, *ManualCancelOrderPaymentRequest) (*ManualCancelOrderPaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ManualCancelOrderPayment not implemented")
}
func (UnimplementedOrderHandleServiceServer) mustEmbedUnimplementedOrderHandleServiceServer() {}
func (UnimplementedOrderHandleServiceServer) testEmbeddedByValue()                            {}

// UnsafeOrderHandleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderHandleServiceServer will
// result in compilation errors.
type UnsafeOrderHandleServiceServer interface {
	mustEmbedUnimplementedOrderHandleServiceServer()
}

func RegisterOrderHandleServiceServer(s grpc.ServiceRegistrar, srv OrderHandleServiceServer) {
	// If the following call pancis, it indicates UnimplementedOrderHandleServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&OrderHandleService_ServiceDesc, srv)
}

func _OrderHandleService_GetOrderStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderHandleServiceServer).GetOrderStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderHandleService_GetOrderStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderHandleServiceServer).GetOrderStatus(ctx, req.(*GetOrderStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderHandleService_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderHandleServiceServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderHandleService_CancelOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderHandleServiceServer).CancelOrder(ctx, req.(*CancelOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderHandleService_PlaceOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaceOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderHandleServiceServer).PlaceOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderHandleService_PlaceOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderHandleServiceServer).PlaceOrder(ctx, req.(*PlaceOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderHandleService_ManualPassOrderPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ManualPassOrderPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderHandleServiceServer).ManualPassOrderPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderHandleService_ManualPassOrderPayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderHandleServiceServer).ManualPassOrderPayment(ctx, req.(*ManualPassOrderPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderHandleService_ManualCancelOrderPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ManualCancelOrderPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderHandleServiceServer).ManualCancelOrderPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderHandleService_ManualCancelOrderPayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderHandleServiceServer).ManualCancelOrderPayment(ctx, req.(*ManualCancelOrderPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderHandleService_ServiceDesc is the grpc.ServiceDesc for OrderHandleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderHandleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "orderHandle.OrderHandleService",
	HandlerType: (*OrderHandleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrderStatus",
			Handler:    _OrderHandleService_GetOrderStatus_Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _OrderHandleService_CancelOrder_Handler,
		},
		{
			MethodName: "PlaceOrder",
			Handler:    _OrderHandleService_PlaceOrder_Handler,
		},
		{
			MethodName: "ManualPassOrderPayment",
			Handler:    _OrderHandleService_ManualPassOrderPayment_Handler,
		},
		{
			MethodName: "ManualCancelOrderPayment",
			Handler:    _OrderHandleService_ManualCancelOrderPayment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order_handle.proto",
}
