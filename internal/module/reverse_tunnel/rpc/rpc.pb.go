// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: internal/module/reverse_tunnel/rpc/rpc.proto

package rpc

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	info "gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/module/reverse_tunnel/info"
	_ "gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/tool/grpctool/automata"
	prototool "gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/tool/prototool"
	status "google.golang.org/genproto/googleapis/rpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Descriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentDescriptor *info.AgentDescriptor `protobuf:"bytes,1,opt,name=agent_descriptor,json=agentDescriptor,proto3" json:"agent_descriptor,omitempty"`
}

func (x *Descriptor) Reset() {
	*x = Descriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_module_reverse_tunnel_rpc_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Descriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Descriptor) ProtoMessage() {}

func (x *Descriptor) ProtoReflect() protoreflect.Message {
	mi := &file_internal_module_reverse_tunnel_rpc_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Descriptor.ProtoReflect.Descriptor instead.
func (*Descriptor) Descriptor() ([]byte, []int) {
	return file_internal_module_reverse_tunnel_rpc_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *Descriptor) GetAgentDescriptor() *info.AgentDescriptor {
	if x != nil {
		return x.AgentDescriptor
	}
	return nil
}

type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta map[string]*prototool.Values `protobuf:"bytes,1,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_module_reverse_tunnel_rpc_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_internal_module_reverse_tunnel_rpc_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_internal_module_reverse_tunnel_rpc_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *Header) GetMeta() map[string]*prototool.Values {
	if x != nil {
		return x.Meta
	}
	return nil
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_module_reverse_tunnel_rpc_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_internal_module_reverse_tunnel_rpc_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_internal_module_reverse_tunnel_rpc_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *Message) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type Trailer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta map[string]*prototool.Values `protobuf:"bytes,1,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Trailer) Reset() {
	*x = Trailer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_module_reverse_tunnel_rpc_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trailer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trailer) ProtoMessage() {}

func (x *Trailer) ProtoReflect() protoreflect.Message {
	mi := &file_internal_module_reverse_tunnel_rpc_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trailer.ProtoReflect.Descriptor instead.
func (*Trailer) Descriptor() ([]byte, []int) {
	return file_internal_module_reverse_tunnel_rpc_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *Trailer) GetMeta() map[string]*prototool.Values {
	if x != nil {
		return x.Meta
	}
	return nil
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *status.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_module_reverse_tunnel_rpc_rpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_internal_module_reverse_tunnel_rpc_rpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_internal_module_reverse_tunnel_rpc_rpc_proto_rawDescGZIP(), []int{4}
}

func (x *Error) GetStatus() *status.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

type ConnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Msg:
	//	*ConnectRequest_Descriptor_
	//	*ConnectRequest_Header
	//	*ConnectRequest_Message
	//	*ConnectRequest_Trailer
	//	*ConnectRequest_Error
	Msg isConnectRequest_Msg `protobuf_oneof:"msg"`
}

func (x *ConnectRequest) Reset() {
	*x = ConnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_module_reverse_tunnel_rpc_rpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectRequest) ProtoMessage() {}

func (x *ConnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_module_reverse_tunnel_rpc_rpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectRequest.ProtoReflect.Descriptor instead.
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return file_internal_module_reverse_tunnel_rpc_rpc_proto_rawDescGZIP(), []int{5}
}

func (m *ConnectRequest) GetMsg() isConnectRequest_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (x *ConnectRequest) GetDescriptor_() *Descriptor {
	if x, ok := x.GetMsg().(*ConnectRequest_Descriptor_); ok {
		return x.Descriptor_
	}
	return nil
}

func (x *ConnectRequest) GetHeader() *Header {
	if x, ok := x.GetMsg().(*ConnectRequest_Header); ok {
		return x.Header
	}
	return nil
}

func (x *ConnectRequest) GetMessage() *Message {
	if x, ok := x.GetMsg().(*ConnectRequest_Message); ok {
		return x.Message
	}
	return nil
}

func (x *ConnectRequest) GetTrailer() *Trailer {
	if x, ok := x.GetMsg().(*ConnectRequest_Trailer); ok {
		return x.Trailer
	}
	return nil
}

func (x *ConnectRequest) GetError() *Error {
	if x, ok := x.GetMsg().(*ConnectRequest_Error); ok {
		return x.Error
	}
	return nil
}

type isConnectRequest_Msg interface {
	isConnectRequest_Msg()
}

type ConnectRequest_Descriptor_ struct {
	Descriptor_ *Descriptor `protobuf:"bytes,1,opt,name=descriptor,proto3,oneof"`
}

type ConnectRequest_Header struct {
	Header *Header `protobuf:"bytes,2,opt,name=header,proto3,oneof"`
}

type ConnectRequest_Message struct {
	Message *Message `protobuf:"bytes,3,opt,name=message,proto3,oneof"`
}

type ConnectRequest_Trailer struct {
	Trailer *Trailer `protobuf:"bytes,4,opt,name=trailer,proto3,oneof"`
}

type ConnectRequest_Error struct {
	Error *Error `protobuf:"bytes,5,opt,name=error,proto3,oneof"`
}

func (*ConnectRequest_Descriptor_) isConnectRequest_Msg() {}

func (*ConnectRequest_Header) isConnectRequest_Msg() {}

func (*ConnectRequest_Message) isConnectRequest_Msg() {}

func (*ConnectRequest_Trailer) isConnectRequest_Msg() {}

func (*ConnectRequest_Error) isConnectRequest_Msg() {}

type RequestInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MethodName string                       `protobuf:"bytes,1,opt,name=method_name,json=methodName,proto3" json:"method_name,omitempty"`
	Meta       map[string]*prototool.Values `protobuf:"bytes,2,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RequestInfo) Reset() {
	*x = RequestInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_module_reverse_tunnel_rpc_rpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestInfo) ProtoMessage() {}

func (x *RequestInfo) ProtoReflect() protoreflect.Message {
	mi := &file_internal_module_reverse_tunnel_rpc_rpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestInfo.ProtoReflect.Descriptor instead.
func (*RequestInfo) Descriptor() ([]byte, []int) {
	return file_internal_module_reverse_tunnel_rpc_rpc_proto_rawDescGZIP(), []int{6}
}

func (x *RequestInfo) GetMethodName() string {
	if x != nil {
		return x.MethodName
	}
	return ""
}

func (x *RequestInfo) GetMeta() map[string]*prototool.Values {
	if x != nil {
		return x.Meta
	}
	return nil
}

type CloseSend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CloseSend) Reset() {
	*x = CloseSend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_module_reverse_tunnel_rpc_rpc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseSend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseSend) ProtoMessage() {}

func (x *CloseSend) ProtoReflect() protoreflect.Message {
	mi := &file_internal_module_reverse_tunnel_rpc_rpc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseSend.ProtoReflect.Descriptor instead.
func (*CloseSend) Descriptor() ([]byte, []int) {
	return file_internal_module_reverse_tunnel_rpc_rpc_proto_rawDescGZIP(), []int{7}
}

type ConnectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Msg:
	//	*ConnectResponse_RequestInfo
	//	*ConnectResponse_Message
	//	*ConnectResponse_CloseSend
	Msg isConnectResponse_Msg `protobuf_oneof:"msg"`
}

func (x *ConnectResponse) Reset() {
	*x = ConnectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_module_reverse_tunnel_rpc_rpc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectResponse) ProtoMessage() {}

func (x *ConnectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_module_reverse_tunnel_rpc_rpc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectResponse.ProtoReflect.Descriptor instead.
func (*ConnectResponse) Descriptor() ([]byte, []int) {
	return file_internal_module_reverse_tunnel_rpc_rpc_proto_rawDescGZIP(), []int{8}
}

func (m *ConnectResponse) GetMsg() isConnectResponse_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (x *ConnectResponse) GetRequestInfo() *RequestInfo {
	if x, ok := x.GetMsg().(*ConnectResponse_RequestInfo); ok {
		return x.RequestInfo
	}
	return nil
}

func (x *ConnectResponse) GetMessage() *Message {
	if x, ok := x.GetMsg().(*ConnectResponse_Message); ok {
		return x.Message
	}
	return nil
}

func (x *ConnectResponse) GetCloseSend() *CloseSend {
	if x, ok := x.GetMsg().(*ConnectResponse_CloseSend); ok {
		return x.CloseSend
	}
	return nil
}

type isConnectResponse_Msg interface {
	isConnectResponse_Msg()
}

type ConnectResponse_RequestInfo struct {
	RequestInfo *RequestInfo `protobuf:"bytes,1,opt,name=request_info,json=requestInfo,proto3,oneof"`
}

type ConnectResponse_Message struct {
	Message *Message `protobuf:"bytes,2,opt,name=message,proto3,oneof"`
}

type ConnectResponse_CloseSend struct {
	CloseSend *CloseSend `protobuf:"bytes,3,opt,name=close_send,json=closeSend,proto3,oneof"`
}

func (*ConnectResponse_RequestInfo) isConnectResponse_Msg() {}

func (*ConnectResponse_Message) isConnectResponse_Msg() {}

func (*ConnectResponse_CloseSend) isConnectResponse_Msg() {}

var File_internal_module_reverse_tunnel_rpc_rpc_proto protoreflect.FileDescriptor

var file_internal_module_reverse_tunnel_rpc_rpc_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x2f, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x5f, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c,
	0x2f, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f,
	0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x72, 0x65, 0x76,
	0x65, 0x72, 0x73, 0x65, 0x5f, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x72, 0x70, 0x63, 0x1a,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x6f, 0x6f, 0x6c, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x74, 0x6f, 0x6f, 0x6c, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x61,
	0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x27, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x6f, 0x6f, 0x6c, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x74, 0x6f, 0x6f, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x74, 0x6f,
	0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73,
	0x65, 0x5f, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x69, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x74, 0x0a, 0x0a, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x66, 0x0a, 0x10, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x31, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x2e, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x5f, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c,
	0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52,
	0x0f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x22, 0xa8, 0x01, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x45, 0x0a, 0x04, 0x6d,
	0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x67, 0x69, 0x74, 0x6c,
	0x61, 0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65,
	0x5f, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6d, 0x65,
	0x74, 0x61, 0x1a, 0x57, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x34, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x74, 0x6f, 0x6f, 0x6c, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x1d, 0x0a, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xaa, 0x01, 0x0a, 0x07, 0x54,
	0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x12, 0x46, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x2e, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x5f, 0x74, 0x75, 0x6e, 0x6e,
	0x65, 0x6c, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x1a, 0x57,
	0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x34, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67,
	0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x74, 0x6f, 0x6f, 0x6c, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x3d, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x34, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xe6, 0x03, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x5f, 0x0a, 0x0a, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e,
	0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x72, 0x65, 0x76,
	0x65, 0x72, 0x73, 0x65, 0x5f, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x42, 0x10, 0x80, 0xf6, 0x2c, 0x02,
	0x80, 0xf6, 0x2c, 0x05, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x0a,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x53, 0x0a, 0x06, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x67, 0x69, 0x74,
	0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73,
	0x65, 0x5f, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x42, 0x10, 0x80, 0xf6, 0x2c, 0x03, 0x80, 0xf6, 0x2c, 0x04, 0xfa, 0x42, 0x05,
	0x8a, 0x01, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x56, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e,
	0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x5f, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x10, 0x80, 0xf6, 0x2c, 0x03,
	0x80, 0xf6, 0x2c, 0x04, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x5f, 0x0a, 0x07, 0x74, 0x72, 0x61, 0x69, 0x6c,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61,
	0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x5f,
	0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x6c,
	0x65, 0x72, 0x42, 0x19, 0x80, 0xf6, 0x2c, 0x05, 0x80, 0xf6, 0x2c, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0x01, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52,
	0x07, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x12, 0x55, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62,
	0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x5f, 0x74,
	0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x42,
	0x15, 0x80, 0xf6, 0x2c, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0xfa, 0x42,
	0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42,
	0x0e, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x07, 0x88, 0xf6, 0x2c, 0x01, 0xf8, 0x42, 0x01, 0x22,
	0xd3, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x4a, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36,
	0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x72, 0x65,
	0x76, 0x65, 0x72, 0x73, 0x65, 0x5f, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x4d, 0x65, 0x74,
	0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x1a, 0x57, 0x0a, 0x09,
	0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x34, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x69, 0x74,
	0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x74,
	0x6f, 0x6f, 0x6c, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x0b, 0x0a, 0x09, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x53, 0x65,
	0x6e, 0x64, 0x22, 0xe9, 0x02, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x70, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67,
	0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x72, 0x65, 0x76, 0x65,
	0x72, 0x73, 0x65, 0x5f, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x1d, 0x80, 0xf6, 0x2c, 0x02,
	0x80, 0xf6, 0x2c, 0x03, 0x80, 0xf6, 0x2c, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0x01, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x0b, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x63, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x69, 0x74, 0x6c,
	0x61, 0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65,
	0x5f, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x42, 0x1d, 0x80, 0xf6, 0x2c, 0x02, 0x80, 0xf6, 0x2c, 0x03, 0x80, 0xf6, 0x2c,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02,
	0x10, 0x01, 0x48, 0x00, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x62, 0x0a,
	0x0a, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x2e, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x5f, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x42, 0x15, 0x80,
	0xf6, 0x2c, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0xfa, 0x42, 0x05, 0x8a,
	0x01, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x09, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x53, 0x65, 0x6e,
	0x64, 0x42, 0x1b, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x14, 0x88, 0xf6, 0x2c, 0x01, 0x88, 0xf6,
	0x2c, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0xf8, 0x42, 0x01, 0x32, 0x83,
	0x01, 0x0a, 0x0d, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c,
	0x12, 0x72, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x2f, 0x2e, 0x67, 0x69,
	0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x72, 0x65, 0x76, 0x65, 0x72,
	0x73, 0x65, 0x5f, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x67,
	0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x72, 0x65, 0x76, 0x65,
	0x72, 0x73, 0x65, 0x5f, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x28, 0x01, 0x30, 0x01, 0x42, 0x5f, 0x5a, 0x5d, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2d, 0x6f, 0x72, 0x67, 0x2f, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x76,
	0x31, 0x35, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x2f, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x5f, 0x74, 0x75, 0x6e, 0x6e, 0x65,
	0x6c, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_module_reverse_tunnel_rpc_rpc_proto_rawDescOnce sync.Once
	file_internal_module_reverse_tunnel_rpc_rpc_proto_rawDescData = file_internal_module_reverse_tunnel_rpc_rpc_proto_rawDesc
)

func file_internal_module_reverse_tunnel_rpc_rpc_proto_rawDescGZIP() []byte {
	file_internal_module_reverse_tunnel_rpc_rpc_proto_rawDescOnce.Do(func() {
		file_internal_module_reverse_tunnel_rpc_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_module_reverse_tunnel_rpc_rpc_proto_rawDescData)
	})
	return file_internal_module_reverse_tunnel_rpc_rpc_proto_rawDescData
}

var file_internal_module_reverse_tunnel_rpc_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_internal_module_reverse_tunnel_rpc_rpc_proto_goTypes = []interface{}{
	(*Descriptor)(nil),           // 0: gitlab.agent.reverse_tunnel.rpc.Descriptor
	(*Header)(nil),               // 1: gitlab.agent.reverse_tunnel.rpc.Header
	(*Message)(nil),              // 2: gitlab.agent.reverse_tunnel.rpc.Message
	(*Trailer)(nil),              // 3: gitlab.agent.reverse_tunnel.rpc.Trailer
	(*Error)(nil),                // 4: gitlab.agent.reverse_tunnel.rpc.Error
	(*ConnectRequest)(nil),       // 5: gitlab.agent.reverse_tunnel.rpc.ConnectRequest
	(*RequestInfo)(nil),          // 6: gitlab.agent.reverse_tunnel.rpc.RequestInfo
	(*CloseSend)(nil),            // 7: gitlab.agent.reverse_tunnel.rpc.CloseSend
	(*ConnectResponse)(nil),      // 8: gitlab.agent.reverse_tunnel.rpc.ConnectResponse
	nil,                          // 9: gitlab.agent.reverse_tunnel.rpc.Header.MetaEntry
	nil,                          // 10: gitlab.agent.reverse_tunnel.rpc.Trailer.MetaEntry
	nil,                          // 11: gitlab.agent.reverse_tunnel.rpc.RequestInfo.MetaEntry
	(*info.AgentDescriptor)(nil), // 12: gitlab.agent.reverse_tunnel.info.AgentDescriptor
	(*status.Status)(nil),        // 13: google.rpc.Status
	(*prototool.Values)(nil),     // 14: gitlab.agent.prototool.Values
}
var file_internal_module_reverse_tunnel_rpc_rpc_proto_depIdxs = []int32{
	12, // 0: gitlab.agent.reverse_tunnel.rpc.Descriptor.agent_descriptor:type_name -> gitlab.agent.reverse_tunnel.info.AgentDescriptor
	9,  // 1: gitlab.agent.reverse_tunnel.rpc.Header.meta:type_name -> gitlab.agent.reverse_tunnel.rpc.Header.MetaEntry
	10, // 2: gitlab.agent.reverse_tunnel.rpc.Trailer.meta:type_name -> gitlab.agent.reverse_tunnel.rpc.Trailer.MetaEntry
	13, // 3: gitlab.agent.reverse_tunnel.rpc.Error.status:type_name -> google.rpc.Status
	0,  // 4: gitlab.agent.reverse_tunnel.rpc.ConnectRequest.descriptor:type_name -> gitlab.agent.reverse_tunnel.rpc.Descriptor
	1,  // 5: gitlab.agent.reverse_tunnel.rpc.ConnectRequest.header:type_name -> gitlab.agent.reverse_tunnel.rpc.Header
	2,  // 6: gitlab.agent.reverse_tunnel.rpc.ConnectRequest.message:type_name -> gitlab.agent.reverse_tunnel.rpc.Message
	3,  // 7: gitlab.agent.reverse_tunnel.rpc.ConnectRequest.trailer:type_name -> gitlab.agent.reverse_tunnel.rpc.Trailer
	4,  // 8: gitlab.agent.reverse_tunnel.rpc.ConnectRequest.error:type_name -> gitlab.agent.reverse_tunnel.rpc.Error
	11, // 9: gitlab.agent.reverse_tunnel.rpc.RequestInfo.meta:type_name -> gitlab.agent.reverse_tunnel.rpc.RequestInfo.MetaEntry
	6,  // 10: gitlab.agent.reverse_tunnel.rpc.ConnectResponse.request_info:type_name -> gitlab.agent.reverse_tunnel.rpc.RequestInfo
	2,  // 11: gitlab.agent.reverse_tunnel.rpc.ConnectResponse.message:type_name -> gitlab.agent.reverse_tunnel.rpc.Message
	7,  // 12: gitlab.agent.reverse_tunnel.rpc.ConnectResponse.close_send:type_name -> gitlab.agent.reverse_tunnel.rpc.CloseSend
	14, // 13: gitlab.agent.reverse_tunnel.rpc.Header.MetaEntry.value:type_name -> gitlab.agent.prototool.Values
	14, // 14: gitlab.agent.reverse_tunnel.rpc.Trailer.MetaEntry.value:type_name -> gitlab.agent.prototool.Values
	14, // 15: gitlab.agent.reverse_tunnel.rpc.RequestInfo.MetaEntry.value:type_name -> gitlab.agent.prototool.Values
	5,  // 16: gitlab.agent.reverse_tunnel.rpc.ReverseTunnel.Connect:input_type -> gitlab.agent.reverse_tunnel.rpc.ConnectRequest
	8,  // 17: gitlab.agent.reverse_tunnel.rpc.ReverseTunnel.Connect:output_type -> gitlab.agent.reverse_tunnel.rpc.ConnectResponse
	17, // [17:18] is the sub-list for method output_type
	16, // [16:17] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_internal_module_reverse_tunnel_rpc_rpc_proto_init() }
func file_internal_module_reverse_tunnel_rpc_rpc_proto_init() {
	if File_internal_module_reverse_tunnel_rpc_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_module_reverse_tunnel_rpc_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Descriptor); i {
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
		file_internal_module_reverse_tunnel_rpc_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Header); i {
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
		file_internal_module_reverse_tunnel_rpc_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_internal_module_reverse_tunnel_rpc_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Trailer); i {
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
		file_internal_module_reverse_tunnel_rpc_rpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_internal_module_reverse_tunnel_rpc_rpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectRequest); i {
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
		file_internal_module_reverse_tunnel_rpc_rpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestInfo); i {
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
		file_internal_module_reverse_tunnel_rpc_rpc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloseSend); i {
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
		file_internal_module_reverse_tunnel_rpc_rpc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectResponse); i {
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
	file_internal_module_reverse_tunnel_rpc_rpc_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*ConnectRequest_Descriptor_)(nil),
		(*ConnectRequest_Header)(nil),
		(*ConnectRequest_Message)(nil),
		(*ConnectRequest_Trailer)(nil),
		(*ConnectRequest_Error)(nil),
	}
	file_internal_module_reverse_tunnel_rpc_rpc_proto_msgTypes[8].OneofWrappers = []interface{}{
		(*ConnectResponse_RequestInfo)(nil),
		(*ConnectResponse_Message)(nil),
		(*ConnectResponse_CloseSend)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_module_reverse_tunnel_rpc_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_module_reverse_tunnel_rpc_rpc_proto_goTypes,
		DependencyIndexes: file_internal_module_reverse_tunnel_rpc_rpc_proto_depIdxs,
		MessageInfos:      file_internal_module_reverse_tunnel_rpc_rpc_proto_msgTypes,
	}.Build()
	File_internal_module_reverse_tunnel_rpc_rpc_proto = out.File
	file_internal_module_reverse_tunnel_rpc_rpc_proto_rawDesc = nil
	file_internal_module_reverse_tunnel_rpc_rpc_proto_goTypes = nil
	file_internal_module_reverse_tunnel_rpc_rpc_proto_depIdxs = nil
}
