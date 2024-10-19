// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.19.4
// source: order.proto

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
	OrderService_GetActivePlanListByUserId_FullMethodName = "/order.OrderService/GetActivePlanListByUserId"
	OrderService_CommitNewOrder_FullMethodName            = "/order.OrderService/CommitNewOrder"
	OrderService_GetAllMyOrders_FullMethodName            = "/order.OrderService/GetAllMyOrders"
)

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 定义 Order 服务
type OrderServiceClient interface {
	// GetActivePlanListByUserId 獲取該用戶所有有效的訂單
	GetActivePlanListByUserId(ctx context.Context, in *GetActivePlanListByUserIdRequest, opts ...grpc.CallOption) (*GetActivePlanListByUserIdResponse, error)
	// CommitNewOrder 提交新訂單時
	CommitNewOrder(ctx context.Context, in *CommitNewOrderRequest, opts ...grpc.CallOption) (*CommitNewOrderResponse, error)
	// GetAllMyOrders 獲取該用戶所有的訂單 不管是有效還是無效
	GetAllMyOrders(ctx context.Context, in *GetAllMyOrdersRequest, opts ...grpc.CallOption) (*GetAllMyOrdersResponse, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) GetActivePlanListByUserId(ctx context.Context, in *GetActivePlanListByUserIdRequest, opts ...grpc.CallOption) (*GetActivePlanListByUserIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetActivePlanListByUserIdResponse)
	err := c.cc.Invoke(ctx, OrderService_GetActivePlanListByUserId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) CommitNewOrder(ctx context.Context, in *CommitNewOrderRequest, opts ...grpc.CallOption) (*CommitNewOrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommitNewOrderResponse)
	err := c.cc.Invoke(ctx, OrderService_CommitNewOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetAllMyOrders(ctx context.Context, in *GetAllMyOrdersRequest, opts ...grpc.CallOption) (*GetAllMyOrdersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllMyOrdersResponse)
	err := c.cc.Invoke(ctx, OrderService_GetAllMyOrders_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
// All implementations must embed UnimplementedOrderServiceServer
// for forward compatibility.
//
// 定义 Order 服务
type OrderServiceServer interface {
	// GetActivePlanListByUserId 獲取該用戶所有有效的訂單
	GetActivePlanListByUserId(context.Context, *GetActivePlanListByUserIdRequest) (*GetActivePlanListByUserIdResponse, error)
	// CommitNewOrder 提交新訂單時
	CommitNewOrder(context.Context, *CommitNewOrderRequest) (*CommitNewOrderResponse, error)
	// GetAllMyOrders 獲取該用戶所有的訂單 不管是有效還是無效
	GetAllMyOrders(context.Context, *GetAllMyOrdersRequest) (*GetAllMyOrdersResponse, error)
	mustEmbedUnimplementedOrderServiceServer()
}

// UnimplementedOrderServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedOrderServiceServer struct{}

func (UnimplementedOrderServiceServer) GetActivePlanListByUserId(context.Context, *GetActivePlanListByUserIdRequest) (*GetActivePlanListByUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActivePlanListByUserId not implemented")
}
func (UnimplementedOrderServiceServer) CommitNewOrder(context.Context, *CommitNewOrderRequest) (*CommitNewOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommitNewOrder not implemented")
}
func (UnimplementedOrderServiceServer) GetAllMyOrders(context.Context, *GetAllMyOrdersRequest) (*GetAllMyOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllMyOrders not implemented")
}
func (UnimplementedOrderServiceServer) mustEmbedUnimplementedOrderServiceServer() {}
func (UnimplementedOrderServiceServer) testEmbeddedByValue()                      {}

// UnsafeOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServiceServer will
// result in compilation errors.
type UnsafeOrderServiceServer interface {
	mustEmbedUnimplementedOrderServiceServer()
}

func RegisterOrderServiceServer(s grpc.ServiceRegistrar, srv OrderServiceServer) {
	// If the following call pancis, it indicates UnimplementedOrderServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&OrderService_ServiceDesc, srv)
}

func _OrderService_GetActivePlanListByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActivePlanListByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetActivePlanListByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_GetActivePlanListByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetActivePlanListByUserId(ctx, req.(*GetActivePlanListByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_CommitNewOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommitNewOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CommitNewOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_CommitNewOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CommitNewOrder(ctx, req.(*CommitNewOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetAllMyOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllMyOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetAllMyOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_GetAllMyOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetAllMyOrders(ctx, req.(*GetAllMyOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderService_ServiceDesc is the grpc.ServiceDesc for OrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetActivePlanListByUserId",
			Handler:    _OrderService_GetActivePlanListByUserId_Handler,
		},
		{
			MethodName: "CommitNewOrder",
			Handler:    _OrderService_CommitNewOrder_Handler,
		},
		{
			MethodName: "GetAllMyOrders",
			Handler:    _OrderService_GetAllMyOrders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order.proto",
}
