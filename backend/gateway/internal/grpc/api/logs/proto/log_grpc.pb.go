// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: log.proto

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
	LogService_GetApiAccessLog_FullMethodName     = "/proto.LogService/GetApiAccessLog"
	LogService_SaveApiAccessLog2Db_FullMethodName = "/proto.LogService/SaveApiAccessLog2Db"
	LogService_GetServerLiveStatus_FullMethodName = "/proto.LogService/GetServerLiveStatus"
	LogService_ClearPreviousApiLog_FullMethodName = "/proto.LogService/ClearPreviousApiLog"
)

// LogServiceClient is the client API for LogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogServiceClient interface {
	GetApiAccessLog(ctx context.Context, in *GetApiAccessLogRequest, opts ...grpc.CallOption) (*GetApiAccessLogResponse, error)
	SaveApiAccessLog2Db(ctx context.Context, in *SaveApiAccessLog2DbRequest, opts ...grpc.CallOption) (*SaveApiAccessLog2DbResponse, error)
	GetServerLiveStatus(ctx context.Context, in *GetServerLiveStatusRequest, opts ...grpc.CallOption) (*GetServerLiveStatusResponse, error)
	ClearPreviousApiLog(ctx context.Context, in *ClearPreviousApiLogRequest, opts ...grpc.CallOption) (*ClearPreviousApiLogResponse, error)
}

type logServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLogServiceClient(cc grpc.ClientConnInterface) LogServiceClient {
	return &logServiceClient{cc}
}

func (c *logServiceClient) GetApiAccessLog(ctx context.Context, in *GetApiAccessLogRequest, opts ...grpc.CallOption) (*GetApiAccessLogResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetApiAccessLogResponse)
	err := c.cc.Invoke(ctx, LogService_GetApiAccessLog_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logServiceClient) SaveApiAccessLog2Db(ctx context.Context, in *SaveApiAccessLog2DbRequest, opts ...grpc.CallOption) (*SaveApiAccessLog2DbResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SaveApiAccessLog2DbResponse)
	err := c.cc.Invoke(ctx, LogService_SaveApiAccessLog2Db_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logServiceClient) GetServerLiveStatus(ctx context.Context, in *GetServerLiveStatusRequest, opts ...grpc.CallOption) (*GetServerLiveStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetServerLiveStatusResponse)
	err := c.cc.Invoke(ctx, LogService_GetServerLiveStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logServiceClient) ClearPreviousApiLog(ctx context.Context, in *ClearPreviousApiLogRequest, opts ...grpc.CallOption) (*ClearPreviousApiLogResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ClearPreviousApiLogResponse)
	err := c.cc.Invoke(ctx, LogService_ClearPreviousApiLog_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogServiceServer is the server API for LogService service.
// All implementations must embed UnimplementedLogServiceServer
// for forward compatibility.
type LogServiceServer interface {
	GetApiAccessLog(context.Context, *GetApiAccessLogRequest) (*GetApiAccessLogResponse, error)
	SaveApiAccessLog2Db(context.Context, *SaveApiAccessLog2DbRequest) (*SaveApiAccessLog2DbResponse, error)
	GetServerLiveStatus(context.Context, *GetServerLiveStatusRequest) (*GetServerLiveStatusResponse, error)
	ClearPreviousApiLog(context.Context, *ClearPreviousApiLogRequest) (*ClearPreviousApiLogResponse, error)
	mustEmbedUnimplementedLogServiceServer()
}

// UnimplementedLogServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLogServiceServer struct{}

func (UnimplementedLogServiceServer) GetApiAccessLog(context.Context, *GetApiAccessLogRequest) (*GetApiAccessLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApiAccessLog not implemented")
}
func (UnimplementedLogServiceServer) SaveApiAccessLog2Db(context.Context, *SaveApiAccessLog2DbRequest) (*SaveApiAccessLog2DbResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveApiAccessLog2Db not implemented")
}
func (UnimplementedLogServiceServer) GetServerLiveStatus(context.Context, *GetServerLiveStatusRequest) (*GetServerLiveStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerLiveStatus not implemented")
}
func (UnimplementedLogServiceServer) ClearPreviousApiLog(context.Context, *ClearPreviousApiLogRequest) (*ClearPreviousApiLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearPreviousApiLog not implemented")
}
func (UnimplementedLogServiceServer) mustEmbedUnimplementedLogServiceServer() {}
func (UnimplementedLogServiceServer) testEmbeddedByValue()                    {}

// UnsafeLogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogServiceServer will
// result in compilation errors.
type UnsafeLogServiceServer interface {
	mustEmbedUnimplementedLogServiceServer()
}

func RegisterLogServiceServer(s grpc.ServiceRegistrar, srv LogServiceServer) {
	// If the following call pancis, it indicates UnimplementedLogServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LogService_ServiceDesc, srv)
}

func _LogService_GetApiAccessLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApiAccessLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServiceServer).GetApiAccessLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogService_GetApiAccessLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServiceServer).GetApiAccessLog(ctx, req.(*GetApiAccessLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogService_SaveApiAccessLog2Db_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveApiAccessLog2DbRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServiceServer).SaveApiAccessLog2Db(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogService_SaveApiAccessLog2Db_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServiceServer).SaveApiAccessLog2Db(ctx, req.(*SaveApiAccessLog2DbRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogService_GetServerLiveStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServerLiveStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServiceServer).GetServerLiveStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogService_GetServerLiveStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServiceServer).GetServerLiveStatus(ctx, req.(*GetServerLiveStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogService_ClearPreviousApiLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearPreviousApiLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServiceServer).ClearPreviousApiLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogService_ClearPreviousApiLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServiceServer).ClearPreviousApiLog(ctx, req.(*ClearPreviousApiLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LogService_ServiceDesc is the grpc.ServiceDesc for LogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.LogService",
	HandlerType: (*LogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetApiAccessLog",
			Handler:    _LogService_GetApiAccessLog_Handler,
		},
		{
			MethodName: "SaveApiAccessLog2Db",
			Handler:    _LogService_SaveApiAccessLog2Db_Handler,
		},
		{
			MethodName: "GetServerLiveStatus",
			Handler:    _LogService_GetServerLiveStatus_Handler,
		},
		{
			MethodName: "ClearPreviousApiLog",
			Handler:    _LogService_ClearPreviousApiLog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "log.proto",
}
