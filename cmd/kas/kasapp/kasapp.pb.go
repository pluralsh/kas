// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.7
// source: cmd/kas/kasapp/kasapp.proto

package kasapp

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/tool/grpctool/automata"
	prototool "gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/tool/prototool"
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

type StartStreaming struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StartStreaming) Reset() {
	*x = StartStreaming{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_kas_kasapp_kasapp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartStreaming) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartStreaming) ProtoMessage() {}

func (x *StartStreaming) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_kas_kasapp_kasapp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartStreaming.ProtoReflect.Descriptor instead.
func (*StartStreaming) Descriptor() ([]byte, []int) {
	return file_cmd_kas_kasapp_kasapp_proto_rawDescGZIP(), []int{0}
}

type GatewayKasResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Msg:
	//
	//	*GatewayKasResponse_TunnelReady_
	//	*GatewayKasResponse_Header_
	//	*GatewayKasResponse_Message_
	//	*GatewayKasResponse_Trailer_
	//	*GatewayKasResponse_Error_
	Msg isGatewayKasResponse_Msg `protobuf_oneof:"msg"`
}

func (x *GatewayKasResponse) Reset() {
	*x = GatewayKasResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_kas_kasapp_kasapp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewayKasResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayKasResponse) ProtoMessage() {}

func (x *GatewayKasResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_kas_kasapp_kasapp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayKasResponse.ProtoReflect.Descriptor instead.
func (*GatewayKasResponse) Descriptor() ([]byte, []int) {
	return file_cmd_kas_kasapp_kasapp_proto_rawDescGZIP(), []int{1}
}

func (m *GatewayKasResponse) GetMsg() isGatewayKasResponse_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (x *GatewayKasResponse) GetTunnelReady() *GatewayKasResponse_TunnelReady {
	if x, ok := x.GetMsg().(*GatewayKasResponse_TunnelReady_); ok {
		return x.TunnelReady
	}
	return nil
}

func (x *GatewayKasResponse) GetHeader() *GatewayKasResponse_Header {
	if x, ok := x.GetMsg().(*GatewayKasResponse_Header_); ok {
		return x.Header
	}
	return nil
}

func (x *GatewayKasResponse) GetMessage() *GatewayKasResponse_Message {
	if x, ok := x.GetMsg().(*GatewayKasResponse_Message_); ok {
		return x.Message
	}
	return nil
}

func (x *GatewayKasResponse) GetTrailer() *GatewayKasResponse_Trailer {
	if x, ok := x.GetMsg().(*GatewayKasResponse_Trailer_); ok {
		return x.Trailer
	}
	return nil
}

func (x *GatewayKasResponse) GetError() *GatewayKasResponse_Error {
	if x, ok := x.GetMsg().(*GatewayKasResponse_Error_); ok {
		return x.Error
	}
	return nil
}

type isGatewayKasResponse_Msg interface {
	isGatewayKasResponse_Msg()
}

type GatewayKasResponse_TunnelReady_ struct {
	TunnelReady *GatewayKasResponse_TunnelReady `protobuf:"bytes,1,opt,name=tunnel_ready,json=tunnelReady,proto3,oneof"`
}

type GatewayKasResponse_Header_ struct {
	Header *GatewayKasResponse_Header `protobuf:"bytes,2,opt,name=header,proto3,oneof"`
}

type GatewayKasResponse_Message_ struct {
	Message *GatewayKasResponse_Message `protobuf:"bytes,3,opt,name=message,proto3,oneof"`
}

type GatewayKasResponse_Trailer_ struct {
	Trailer *GatewayKasResponse_Trailer `protobuf:"bytes,4,opt,name=trailer,proto3,oneof"`
}

type GatewayKasResponse_Error_ struct {
	Error *GatewayKasResponse_Error `protobuf:"bytes,5,opt,name=error,proto3,oneof"`
}

func (*GatewayKasResponse_TunnelReady_) isGatewayKasResponse_Msg() {}

func (*GatewayKasResponse_Header_) isGatewayKasResponse_Msg() {}

func (*GatewayKasResponse_Message_) isGatewayKasResponse_Msg() {}

func (*GatewayKasResponse_Trailer_) isGatewayKasResponse_Msg() {}

func (*GatewayKasResponse_Error_) isGatewayKasResponse_Msg() {}

type GatewayKasResponse_TunnelReady struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GatewayKasResponse_TunnelReady) Reset() {
	*x = GatewayKasResponse_TunnelReady{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_kas_kasapp_kasapp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewayKasResponse_TunnelReady) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayKasResponse_TunnelReady) ProtoMessage() {}

func (x *GatewayKasResponse_TunnelReady) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_kas_kasapp_kasapp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayKasResponse_TunnelReady.ProtoReflect.Descriptor instead.
func (*GatewayKasResponse_TunnelReady) Descriptor() ([]byte, []int) {
	return file_cmd_kas_kasapp_kasapp_proto_rawDescGZIP(), []int{1, 0}
}

type GatewayKasResponse_Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta map[string]*prototool.Values `protobuf:"bytes,1,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GatewayKasResponse_Header) Reset() {
	*x = GatewayKasResponse_Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_kas_kasapp_kasapp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewayKasResponse_Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayKasResponse_Header) ProtoMessage() {}

func (x *GatewayKasResponse_Header) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_kas_kasapp_kasapp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayKasResponse_Header.ProtoReflect.Descriptor instead.
func (*GatewayKasResponse_Header) Descriptor() ([]byte, []int) {
	return file_cmd_kas_kasapp_kasapp_proto_rawDescGZIP(), []int{1, 1}
}

func (x *GatewayKasResponse_Header) GetMeta() map[string]*prototool.Values {
	if x != nil {
		return x.Meta
	}
	return nil
}

type GatewayKasResponse_Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GatewayKasResponse_Message) Reset() {
	*x = GatewayKasResponse_Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_kas_kasapp_kasapp_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewayKasResponse_Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayKasResponse_Message) ProtoMessage() {}

func (x *GatewayKasResponse_Message) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_kas_kasapp_kasapp_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayKasResponse_Message.ProtoReflect.Descriptor instead.
func (*GatewayKasResponse_Message) Descriptor() ([]byte, []int) {
	return file_cmd_kas_kasapp_kasapp_proto_rawDescGZIP(), []int{1, 2}
}

func (x *GatewayKasResponse_Message) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type GatewayKasResponse_Trailer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta map[string]*prototool.Values `protobuf:"bytes,1,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GatewayKasResponse_Trailer) Reset() {
	*x = GatewayKasResponse_Trailer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_kas_kasapp_kasapp_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewayKasResponse_Trailer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayKasResponse_Trailer) ProtoMessage() {}

func (x *GatewayKasResponse_Trailer) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_kas_kasapp_kasapp_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayKasResponse_Trailer.ProtoReflect.Descriptor instead.
func (*GatewayKasResponse_Trailer) Descriptor() ([]byte, []int) {
	return file_cmd_kas_kasapp_kasapp_proto_rawDescGZIP(), []int{1, 3}
}

func (x *GatewayKasResponse_Trailer) GetMeta() map[string]*prototool.Values {
	if x != nil {
		return x.Meta
	}
	return nil
}

type GatewayKasResponse_Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *status.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GatewayKasResponse_Error) Reset() {
	*x = GatewayKasResponse_Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_kas_kasapp_kasapp_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewayKasResponse_Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayKasResponse_Error) ProtoMessage() {}

func (x *GatewayKasResponse_Error) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_kas_kasapp_kasapp_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayKasResponse_Error.ProtoReflect.Descriptor instead.
func (*GatewayKasResponse_Error) Descriptor() ([]byte, []int) {
	return file_cmd_kas_kasapp_kasapp_proto_rawDescGZIP(), []int{1, 4}
}

func (x *GatewayKasResponse_Error) GetStatus() *status.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

var File_cmd_kas_kasapp_kasapp_proto protoreflect.FileDescriptor

var file_cmd_kas_kasapp_kasapp_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x6d, 0x64, 0x2f, 0x6b, 0x61, 0x73, 0x2f, 0x6b, 0x61, 0x73, 0x61, 0x70, 0x70,
	0x2f, 0x6b, 0x61, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x67,
	0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x6b, 0x61, 0x73, 0x1a,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x6f, 0x6f, 0x6c, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x74, 0x6f, 0x6f, 0x6c, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x61,
	0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x27, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x6f, 0x6f, 0x6c, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x74, 0x6f, 0x6f, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x74, 0x6f,
	0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x10, 0x0a, 0x0e, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x22, 0xcb, 0x07, 0x0a,
	0x12, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4b, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x63, 0x0a, 0x0c, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x72, 0x65,
	0x61, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x67, 0x69, 0x74, 0x6c,
	0x61, 0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x6b, 0x61, 0x73, 0x2e, 0x47, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x4b, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x61, 0x64, 0x79, 0x42, 0x0c, 0xfa, 0x42, 0x05,
	0x8a, 0x01, 0x02, 0x10, 0x01, 0x80, 0xf6, 0x2c, 0x02, 0x48, 0x00, 0x52, 0x0b, 0x74, 0x75, 0x6e,
	0x6e, 0x65, 0x6c, 0x52, 0x65, 0x61, 0x64, 0x79, 0x12, 0x57, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61,
	0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x6b, 0x61, 0x73, 0x2e, 0x47, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x4b, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x10, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x80,
	0xf6, 0x2c, 0x03, 0x80, 0xf6, 0x2c, 0x04, 0x48, 0x00, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x12, 0x5a, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x2e, 0x6b, 0x61, 0x73, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4b, 0x61, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x42, 0x10, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x80, 0xf6, 0x2c, 0x03, 0x80, 0xf6,
	0x2c, 0x04, 0x48, 0x00, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x63, 0x0a,
	0x07, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c,
	0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x6b, 0x61,
	0x73, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4b, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x42, 0x19, 0xfa, 0x42,
	0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x80, 0xf6, 0x2c, 0x05, 0x80, 0xf6, 0x2c, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x48, 0x00, 0x52, 0x07, 0x74, 0x72, 0x61, 0x69, 0x6c,
	0x65, 0x72, 0x12, 0x59, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x2e, 0x6b, 0x61, 0x73, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4b, 0x61, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x15, 0xfa,
	0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x80, 0xf6, 0x2c, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0x01, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0x0d, 0x0a,
	0x0b, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x61, 0x64, 0x79, 0x1a, 0xac, 0x01, 0x0a,
	0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x49, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x2e, 0x6b, 0x61, 0x73, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x4b, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6d, 0x65,
	0x74, 0x61, 0x1a, 0x57, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x34, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x74, 0x6f, 0x6f, 0x6c, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x1d, 0x0a, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0xae, 0x01, 0x0a, 0x07, 0x54,
	0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x12, 0x4a, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x2e, 0x6b, 0x61, 0x73, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4b,
	0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x6c,
	0x65, 0x72, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6d, 0x65,
	0x74, 0x61, 0x1a, 0x57, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x34, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x74, 0x6f, 0x6f, 0x6c, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3d, 0x0a, 0x05, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x34, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02,
	0x10, 0x01, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0e, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x12, 0x07, 0xf8, 0x42, 0x01, 0x88, 0xf6, 0x2c, 0x01, 0x42, 0x4b, 0x5a, 0x49, 0x67, 0x69,
	0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2d,
	0x6f, 0x72, 0x67, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2d, 0x69, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2d, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x6b, 0x61, 0x73,
	0x2f, 0x6b, 0x61, 0x73, 0x61, 0x70, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmd_kas_kasapp_kasapp_proto_rawDescOnce sync.Once
	file_cmd_kas_kasapp_kasapp_proto_rawDescData = file_cmd_kas_kasapp_kasapp_proto_rawDesc
)

func file_cmd_kas_kasapp_kasapp_proto_rawDescGZIP() []byte {
	file_cmd_kas_kasapp_kasapp_proto_rawDescOnce.Do(func() {
		file_cmd_kas_kasapp_kasapp_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_kas_kasapp_kasapp_proto_rawDescData)
	})
	return file_cmd_kas_kasapp_kasapp_proto_rawDescData
}

var file_cmd_kas_kasapp_kasapp_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_cmd_kas_kasapp_kasapp_proto_goTypes = []interface{}{
	(*StartStreaming)(nil),                 // 0: gitlab.agent.kas.StartStreaming
	(*GatewayKasResponse)(nil),             // 1: gitlab.agent.kas.GatewayKasResponse
	(*GatewayKasResponse_TunnelReady)(nil), // 2: gitlab.agent.kas.GatewayKasResponse.TunnelReady
	(*GatewayKasResponse_Header)(nil),      // 3: gitlab.agent.kas.GatewayKasResponse.Header
	(*GatewayKasResponse_Message)(nil),     // 4: gitlab.agent.kas.GatewayKasResponse.Message
	(*GatewayKasResponse_Trailer)(nil),     // 5: gitlab.agent.kas.GatewayKasResponse.Trailer
	(*GatewayKasResponse_Error)(nil),       // 6: gitlab.agent.kas.GatewayKasResponse.Error
	nil,                                    // 7: gitlab.agent.kas.GatewayKasResponse.Header.MetaEntry
	nil,                                    // 8: gitlab.agent.kas.GatewayKasResponse.Trailer.MetaEntry
	(*status.Status)(nil),                  // 9: google.rpc.Status
	(*prototool.Values)(nil),               // 10: gitlab.agent.prototool.Values
}
var file_cmd_kas_kasapp_kasapp_proto_depIdxs = []int32{
	2,  // 0: gitlab.agent.kas.GatewayKasResponse.tunnel_ready:type_name -> gitlab.agent.kas.GatewayKasResponse.TunnelReady
	3,  // 1: gitlab.agent.kas.GatewayKasResponse.header:type_name -> gitlab.agent.kas.GatewayKasResponse.Header
	4,  // 2: gitlab.agent.kas.GatewayKasResponse.message:type_name -> gitlab.agent.kas.GatewayKasResponse.Message
	5,  // 3: gitlab.agent.kas.GatewayKasResponse.trailer:type_name -> gitlab.agent.kas.GatewayKasResponse.Trailer
	6,  // 4: gitlab.agent.kas.GatewayKasResponse.error:type_name -> gitlab.agent.kas.GatewayKasResponse.Error
	7,  // 5: gitlab.agent.kas.GatewayKasResponse.Header.meta:type_name -> gitlab.agent.kas.GatewayKasResponse.Header.MetaEntry
	8,  // 6: gitlab.agent.kas.GatewayKasResponse.Trailer.meta:type_name -> gitlab.agent.kas.GatewayKasResponse.Trailer.MetaEntry
	9,  // 7: gitlab.agent.kas.GatewayKasResponse.Error.status:type_name -> google.rpc.Status
	10, // 8: gitlab.agent.kas.GatewayKasResponse.Header.MetaEntry.value:type_name -> gitlab.agent.prototool.Values
	10, // 9: gitlab.agent.kas.GatewayKasResponse.Trailer.MetaEntry.value:type_name -> gitlab.agent.prototool.Values
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_cmd_kas_kasapp_kasapp_proto_init() }
func file_cmd_kas_kasapp_kasapp_proto_init() {
	if File_cmd_kas_kasapp_kasapp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmd_kas_kasapp_kasapp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartStreaming); i {
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
		file_cmd_kas_kasapp_kasapp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewayKasResponse); i {
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
		file_cmd_kas_kasapp_kasapp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewayKasResponse_TunnelReady); i {
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
		file_cmd_kas_kasapp_kasapp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewayKasResponse_Header); i {
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
		file_cmd_kas_kasapp_kasapp_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewayKasResponse_Message); i {
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
		file_cmd_kas_kasapp_kasapp_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewayKasResponse_Trailer); i {
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
		file_cmd_kas_kasapp_kasapp_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewayKasResponse_Error); i {
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
	file_cmd_kas_kasapp_kasapp_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*GatewayKasResponse_TunnelReady_)(nil),
		(*GatewayKasResponse_Header_)(nil),
		(*GatewayKasResponse_Message_)(nil),
		(*GatewayKasResponse_Trailer_)(nil),
		(*GatewayKasResponse_Error_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cmd_kas_kasapp_kasapp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_kas_kasapp_kasapp_proto_goTypes,
		DependencyIndexes: file_cmd_kas_kasapp_kasapp_proto_depIdxs,
		MessageInfos:      file_cmd_kas_kasapp_kasapp_proto_msgTypes,
	}.Build()
	File_cmd_kas_kasapp_kasapp_proto = out.File
	file_cmd_kas_kasapp_kasapp_proto_rawDesc = nil
	file_cmd_kas_kasapp_kasapp_proto_goTypes = nil
	file_cmd_kas_kasapp_kasapp_proto_depIdxs = nil
}
