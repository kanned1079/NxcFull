// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.3
// source: ticketHandle.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetUserTicketsByUserIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetUserTicketsByUserIdRequest) Reset() {
	*x = GetUserTicketsByUserIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticketHandle_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserTicketsByUserIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserTicketsByUserIdRequest) ProtoMessage() {}

func (x *GetUserTicketsByUserIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ticketHandle_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserTicketsByUserIdRequest.ProtoReflect.Descriptor instead.
func (*GetUserTicketsByUserIdRequest) Descriptor() ([]byte, []int) {
	return file_ticketHandle_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserTicketsByUserIdRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetUserTicketsByUserIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg     string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Tickets []byte `protobuf:"bytes,3,opt,name=tickets,proto3" json:"tickets,omitempty"`
}

func (x *GetUserTicketsByUserIdResponse) Reset() {
	*x = GetUserTicketsByUserIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticketHandle_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserTicketsByUserIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserTicketsByUserIdResponse) ProtoMessage() {}

func (x *GetUserTicketsByUserIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ticketHandle_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserTicketsByUserIdResponse.ProtoReflect.Descriptor instead.
func (*GetUserTicketsByUserIdResponse) Descriptor() ([]byte, []int) {
	return file_ticketHandle_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserTicketsByUserIdResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetUserTicketsByUserIdResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetUserTicketsByUserIdResponse) GetTickets() []byte {
	if x != nil {
		return x.Tickets
	}
	return nil
}

type CloseTicketByTicketIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TicketId int64 `protobuf:"varint,2,opt,name=ticket_id,json=ticketId,proto3" json:"ticket_id,omitempty"`
}

func (x *CloseTicketByTicketIdRequest) Reset() {
	*x = CloseTicketByTicketIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticketHandle_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseTicketByTicketIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseTicketByTicketIdRequest) ProtoMessage() {}

func (x *CloseTicketByTicketIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ticketHandle_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseTicketByTicketIdRequest.ProtoReflect.Descriptor instead.
func (*CloseTicketByTicketIdRequest) Descriptor() ([]byte, []int) {
	return file_ticketHandle_proto_rawDescGZIP(), []int{2}
}

func (x *CloseTicketByTicketIdRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CloseTicketByTicketIdRequest) GetTicketId() int64 {
	if x != nil {
		return x.TicketId
	}
	return 0
}

type CloseTicketByTicketIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code   int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Closed bool   `protobuf:"varint,2,opt,name=closed,proto3" json:"closed,omitempty"`
	Msg    string `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *CloseTicketByTicketIdResponse) Reset() {
	*x = CloseTicketByTicketIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticketHandle_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseTicketByTicketIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseTicketByTicketIdResponse) ProtoMessage() {}

func (x *CloseTicketByTicketIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ticketHandle_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseTicketByTicketIdResponse.ProtoReflect.Descriptor instead.
func (*CloseTicketByTicketIdResponse) Descriptor() ([]byte, []int) {
	return file_ticketHandle_proto_rawDescGZIP(), []int{3}
}

func (x *CloseTicketByTicketIdResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CloseTicketByTicketIdResponse) GetClosed() bool {
	if x != nil {
		return x.Closed
	}
	return false
}

func (x *CloseTicketByTicketIdResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type GetChatContentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TicketId int64 `protobuf:"varint,2,opt,name=ticket_id,json=ticketId,proto3" json:"ticket_id,omitempty"`
}

func (x *GetChatContentRequest) Reset() {
	*x = GetChatContentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticketHandle_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChatContentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChatContentRequest) ProtoMessage() {}

func (x *GetChatContentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ticketHandle_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChatContentRequest.ProtoReflect.Descriptor instead.
func (*GetChatContentRequest) Descriptor() ([]byte, []int) {
	return file_ticketHandle_proto_rawDescGZIP(), []int{4}
}

func (x *GetChatContentRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetChatContentRequest) GetTicketId() int64 {
	if x != nil {
		return x.TicketId
	}
	return 0
}

type GetChatContentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg     string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Content []byte `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *GetChatContentResponse) Reset() {
	*x = GetChatContentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticketHandle_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChatContentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChatContentResponse) ProtoMessage() {}

func (x *GetChatContentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ticketHandle_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChatContentResponse.ProtoReflect.Descriptor instead.
func (*GetChatContentResponse) Descriptor() ([]byte, []int) {
	return file_ticketHandle_proto_rawDescGZIP(), []int{5}
}

func (x *GetChatContentResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetChatContentResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetChatContentResponse) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type GetAllTicketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PendingPage  int64 `protobuf:"varint,1,opt,name=pending_page,json=pendingPage,proto3" json:"pending_page,omitempty"`
	PendingSize  int64 `protobuf:"varint,2,opt,name=pending_size,json=pendingSize,proto3" json:"pending_size,omitempty"`
	FinishedPage int64 `protobuf:"varint,3,opt,name=finished_page,json=finishedPage,proto3" json:"finished_page,omitempty"`
	FinishedSize int64 `protobuf:"varint,4,opt,name=finished_size,json=finishedSize,proto3" json:"finished_size,omitempty"`
}

func (x *GetAllTicketRequest) Reset() {
	*x = GetAllTicketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticketHandle_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllTicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTicketRequest) ProtoMessage() {}

func (x *GetAllTicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ticketHandle_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTicketRequest.ProtoReflect.Descriptor instead.
func (*GetAllTicketRequest) Descriptor() ([]byte, []int) {
	return file_ticketHandle_proto_rawDescGZIP(), []int{6}
}

func (x *GetAllTicketRequest) GetPendingPage() int64 {
	if x != nil {
		return x.PendingPage
	}
	return 0
}

func (x *GetAllTicketRequest) GetPendingSize() int64 {
	if x != nil {
		return x.PendingSize
	}
	return 0
}

func (x *GetAllTicketRequest) GetFinishedPage() int64 {
	if x != nil {
		return x.FinishedPage
	}
	return 0
}

func (x *GetAllTicketRequest) GetFinishedSize() int64 {
	if x != nil {
		return x.FinishedSize
	}
	return 0
}

type GetAllTicketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code              int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg               string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Tickets           []byte `protobuf:"bytes,3,opt,name=tickets,proto3" json:"tickets,omitempty"`
	PendingTickets    []byte `protobuf:"bytes,4,opt,name=pending_tickets,json=pendingTickets,proto3" json:"pending_tickets,omitempty"`
	PendingPageCount  int64  `protobuf:"varint,5,opt,name=pending_page_count,json=pendingPageCount,proto3" json:"pending_page_count,omitempty"`
	FinishedPageCount int64  `protobuf:"varint,6,opt,name=finished_page_count,json=finishedPageCount,proto3" json:"finished_page_count,omitempty"`
}

func (x *GetAllTicketResponse) Reset() {
	*x = GetAllTicketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticketHandle_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllTicketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTicketResponse) ProtoMessage() {}

func (x *GetAllTicketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ticketHandle_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTicketResponse.ProtoReflect.Descriptor instead.
func (*GetAllTicketResponse) Descriptor() ([]byte, []int) {
	return file_ticketHandle_proto_rawDescGZIP(), []int{7}
}

func (x *GetAllTicketResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetAllTicketResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetAllTicketResponse) GetTickets() []byte {
	if x != nil {
		return x.Tickets
	}
	return nil
}

func (x *GetAllTicketResponse) GetPendingTickets() []byte {
	if x != nil {
		return x.PendingTickets
	}
	return nil
}

func (x *GetAllTicketResponse) GetPendingPageCount() int64 {
	if x != nil {
		return x.PendingPageCount
	}
	return 0
}

func (x *GetAllTicketResponse) GetFinishedPageCount() int64 {
	if x != nil {
		return x.FinishedPageCount
	}
	return 0
}

var File_ticketHandle_proto protoreflect.FileDescriptor

var file_ticketHandle_proto_rawDesc = []byte{
	0x0a, 0x12, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x48, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x22, 0x38, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x60, 0x0a, 0x1e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x42, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x22, 0x54,
	0x0a, 0x1c, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x42, 0x79, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x49, 0x64, 0x22, 0x5d, 0x0a, 0x1d, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x42, 0x79, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6c, 0x6f,
	0x73, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x63, 0x6c, 0x6f, 0x73, 0x65,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x22, 0x4d, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x49, 0x64, 0x22, 0x58, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0xa5, 0x01, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x70, 0x65, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x50, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x65, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x70,
	0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x50, 0x61, 0x67, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x53, 0x69, 0x7a, 0x65, 0x22, 0xdd, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x27, 0x0a,
	0x0f, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x10, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x67, 0x65, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x11, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x50, 0x61, 0x67, 0x65, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x32, 0xb0, 0x03, 0x0a, 0x13, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x48,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x73, 0x0a, 0x16,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x42, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2b, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x48,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x48, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x70, 0x0a, 0x15, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x42, 0x79, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x12, 0x2a, 0x2e, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x42, 0x79, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x48,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x42, 0x79, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x48, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61,
	0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x55, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x12, 0x21, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x48, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ticketHandle_proto_rawDescOnce sync.Once
	file_ticketHandle_proto_rawDescData = file_ticketHandle_proto_rawDesc
)

func file_ticketHandle_proto_rawDescGZIP() []byte {
	file_ticketHandle_proto_rawDescOnce.Do(func() {
		file_ticketHandle_proto_rawDescData = protoimpl.X.CompressGZIP(file_ticketHandle_proto_rawDescData)
	})
	return file_ticketHandle_proto_rawDescData
}

var file_ticketHandle_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_ticketHandle_proto_goTypes = []any{
	(*GetUserTicketsByUserIdRequest)(nil),  // 0: ticketHandle.GetUserTicketsByUserIdRequest
	(*GetUserTicketsByUserIdResponse)(nil), // 1: ticketHandle.GetUserTicketsByUserIdResponse
	(*CloseTicketByTicketIdRequest)(nil),   // 2: ticketHandle.CloseTicketByTicketIdRequest
	(*CloseTicketByTicketIdResponse)(nil),  // 3: ticketHandle.CloseTicketByTicketIdResponse
	(*GetChatContentRequest)(nil),          // 4: ticketHandle.GetChatContentRequest
	(*GetChatContentResponse)(nil),         // 5: ticketHandle.GetChatContentResponse
	(*GetAllTicketRequest)(nil),            // 6: ticketHandle.GetAllTicketRequest
	(*GetAllTicketResponse)(nil),           // 7: ticketHandle.GetAllTicketResponse
}
var file_ticketHandle_proto_depIdxs = []int32{
	0, // 0: ticketHandle.TicketHandleService.GetUserTicketsByUserId:input_type -> ticketHandle.GetUserTicketsByUserIdRequest
	2, // 1: ticketHandle.TicketHandleService.CloseTicketByTicketId:input_type -> ticketHandle.CloseTicketByTicketIdRequest
	4, // 2: ticketHandle.TicketHandleService.GetChatContent:input_type -> ticketHandle.GetChatContentRequest
	6, // 3: ticketHandle.TicketHandleService.GetAllTicket:input_type -> ticketHandle.GetAllTicketRequest
	1, // 4: ticketHandle.TicketHandleService.GetUserTicketsByUserId:output_type -> ticketHandle.GetUserTicketsByUserIdResponse
	3, // 5: ticketHandle.TicketHandleService.CloseTicketByTicketId:output_type -> ticketHandle.CloseTicketByTicketIdResponse
	5, // 6: ticketHandle.TicketHandleService.GetChatContent:output_type -> ticketHandle.GetChatContentResponse
	7, // 7: ticketHandle.TicketHandleService.GetAllTicket:output_type -> ticketHandle.GetAllTicketResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ticketHandle_proto_init() }
func file_ticketHandle_proto_init() {
	if File_ticketHandle_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ticketHandle_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetUserTicketsByUserIdRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ticketHandle_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetUserTicketsByUserIdResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ticketHandle_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CloseTicketByTicketIdRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ticketHandle_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CloseTicketByTicketIdResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ticketHandle_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetChatContentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ticketHandle_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetChatContentResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ticketHandle_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllTicketRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ticketHandle_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllTicketResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ticketHandle_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ticketHandle_proto_goTypes,
		DependencyIndexes: file_ticketHandle_proto_depIdxs,
		MessageInfos:      file_ticketHandle_proto_msgTypes,
	}.Build()
	File_ticketHandle_proto = out.File
	file_ticketHandle_proto_rawDesc = nil
	file_ticketHandle_proto_goTypes = nil
	file_ticketHandle_proto_depIdxs = nil
}
