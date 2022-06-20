// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: internal/module/gitops/rpc/rpc.proto

package rpc

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/tool/grpctool/automata"
	agentcfg "gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/pkg/agentcfg"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ObjectsToSynchronizeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId string             `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	CommitId  string             `protobuf:"bytes,2,opt,name=commit_id,json=commitId,proto3" json:"commit_id,omitempty"`
	Paths     []*agentcfg.PathCF `protobuf:"bytes,3,rep,name=paths,proto3" json:"paths,omitempty"`
}

func (x *ObjectsToSynchronizeRequest) Reset() {
	*x = ObjectsToSynchronizeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_module_gitops_rpc_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectsToSynchronizeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectsToSynchronizeRequest) ProtoMessage() {}

func (x *ObjectsToSynchronizeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_module_gitops_rpc_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectsToSynchronizeRequest.ProtoReflect.Descriptor instead.
func (*ObjectsToSynchronizeRequest) Descriptor() ([]byte, []int) {
	return file_internal_module_gitops_rpc_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *ObjectsToSynchronizeRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *ObjectsToSynchronizeRequest) GetCommitId() string {
	if x != nil {
		return x.CommitId
	}
	return ""
}

func (x *ObjectsToSynchronizeRequest) GetPaths() []*agentcfg.PathCF {
	if x != nil {
		return x.Paths
	}
	return nil
}

type ObjectsToSynchronizeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//	*ObjectsToSynchronizeResponse_Header_
	//	*ObjectsToSynchronizeResponse_Object_
	//	*ObjectsToSynchronizeResponse_Trailer_
	Message isObjectsToSynchronizeResponse_Message `protobuf_oneof:"message"`
}

func (x *ObjectsToSynchronizeResponse) Reset() {
	*x = ObjectsToSynchronizeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_module_gitops_rpc_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectsToSynchronizeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectsToSynchronizeResponse) ProtoMessage() {}

func (x *ObjectsToSynchronizeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_module_gitops_rpc_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectsToSynchronizeResponse.ProtoReflect.Descriptor instead.
func (*ObjectsToSynchronizeResponse) Descriptor() ([]byte, []int) {
	return file_internal_module_gitops_rpc_rpc_proto_rawDescGZIP(), []int{1}
}

func (m *ObjectsToSynchronizeResponse) GetMessage() isObjectsToSynchronizeResponse_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *ObjectsToSynchronizeResponse) GetHeader() *ObjectsToSynchronizeResponse_Header {
	if x, ok := x.GetMessage().(*ObjectsToSynchronizeResponse_Header_); ok {
		return x.Header
	}
	return nil
}

func (x *ObjectsToSynchronizeResponse) GetObject() *ObjectsToSynchronizeResponse_Object {
	if x, ok := x.GetMessage().(*ObjectsToSynchronizeResponse_Object_); ok {
		return x.Object
	}
	return nil
}

func (x *ObjectsToSynchronizeResponse) GetTrailer() *ObjectsToSynchronizeResponse_Trailer {
	if x, ok := x.GetMessage().(*ObjectsToSynchronizeResponse_Trailer_); ok {
		return x.Trailer
	}
	return nil
}

type isObjectsToSynchronizeResponse_Message interface {
	isObjectsToSynchronizeResponse_Message()
}

type ObjectsToSynchronizeResponse_Header_ struct {
	Header *ObjectsToSynchronizeResponse_Header `protobuf:"bytes,1,opt,name=header,proto3,oneof"`
}

type ObjectsToSynchronizeResponse_Object_ struct {
	Object *ObjectsToSynchronizeResponse_Object `protobuf:"bytes,2,opt,name=object,proto3,oneof"`
}

type ObjectsToSynchronizeResponse_Trailer_ struct {
	Trailer *ObjectsToSynchronizeResponse_Trailer `protobuf:"bytes,3,opt,name=trailer,proto3,oneof"`
}

func (*ObjectsToSynchronizeResponse_Header_) isObjectsToSynchronizeResponse_Message() {}

func (*ObjectsToSynchronizeResponse_Object_) isObjectsToSynchronizeResponse_Message() {}

func (*ObjectsToSynchronizeResponse_Trailer_) isObjectsToSynchronizeResponse_Message() {}

type ObjectsToSynchronizeResponse_Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommitId  string `protobuf:"bytes,1,opt,name=commit_id,json=commitId,proto3" json:"commit_id,omitempty"`
	ProjectId int64  `protobuf:"varint,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *ObjectsToSynchronizeResponse_Header) Reset() {
	*x = ObjectsToSynchronizeResponse_Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_module_gitops_rpc_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectsToSynchronizeResponse_Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectsToSynchronizeResponse_Header) ProtoMessage() {}

func (x *ObjectsToSynchronizeResponse_Header) ProtoReflect() protoreflect.Message {
	mi := &file_internal_module_gitops_rpc_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectsToSynchronizeResponse_Header.ProtoReflect.Descriptor instead.
func (*ObjectsToSynchronizeResponse_Header) Descriptor() ([]byte, []int) {
	return file_internal_module_gitops_rpc_rpc_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ObjectsToSynchronizeResponse_Header) GetCommitId() string {
	if x != nil {
		return x.CommitId
	}
	return ""
}

func (x *ObjectsToSynchronizeResponse_Header) GetProjectId() int64 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

type ObjectsToSynchronizeResponse_Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Data   []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ObjectsToSynchronizeResponse_Object) Reset() {
	*x = ObjectsToSynchronizeResponse_Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_module_gitops_rpc_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectsToSynchronizeResponse_Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectsToSynchronizeResponse_Object) ProtoMessage() {}

func (x *ObjectsToSynchronizeResponse_Object) ProtoReflect() protoreflect.Message {
	mi := &file_internal_module_gitops_rpc_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectsToSynchronizeResponse_Object.ProtoReflect.Descriptor instead.
func (*ObjectsToSynchronizeResponse_Object) Descriptor() ([]byte, []int) {
	return file_internal_module_gitops_rpc_rpc_proto_rawDescGZIP(), []int{1, 1}
}

func (x *ObjectsToSynchronizeResponse_Object) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *ObjectsToSynchronizeResponse_Object) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type ObjectsToSynchronizeResponse_Trailer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ObjectsToSynchronizeResponse_Trailer) Reset() {
	*x = ObjectsToSynchronizeResponse_Trailer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_module_gitops_rpc_rpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectsToSynchronizeResponse_Trailer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectsToSynchronizeResponse_Trailer) ProtoMessage() {}

func (x *ObjectsToSynchronizeResponse_Trailer) ProtoReflect() protoreflect.Message {
	mi := &file_internal_module_gitops_rpc_rpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectsToSynchronizeResponse_Trailer.ProtoReflect.Descriptor instead.
func (*ObjectsToSynchronizeResponse_Trailer) Descriptor() ([]byte, []int) {
	return file_internal_module_gitops_rpc_rpc_proto_rawDescGZIP(), []int{1, 2}
}

var File_internal_module_gitops_rpc_rpc_proto protoreflect.FileDescriptor

var file_internal_module_gitops_rpc_rpc_proto_rawDesc = []byte{
	0x0a, 0x24, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x2f, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x70, 0x63,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x1a,
	0x1b, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x63, 0x66, 0x67, 0x2f, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x63, 0x66, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x6f, 0x6f, 0x6c, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x74, 0x6f, 0x6f, 0x6c, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x61, 0x2f, 0x61, 0x75,
	0x74, 0x6f, 0x6d, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x01, 0x0a, 0x1b, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x54, 0x6f, 0x53, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x20, 0x01, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x05, 0x70, 0x61,
	0x74, 0x68, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x69, 0x74, 0x6c,
	0x61, 0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x63, 0x66,
	0x67, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x43, 0x46, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02,
	0x08, 0x01, 0x52, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73, 0x22, 0x86, 0x04, 0x0a, 0x1c, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x54, 0x6f, 0x53, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69,
	0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x60, 0x0a, 0x06, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x67, 0x69, 0x74,
	0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x54, 0x6f, 0x53, 0x79,
	0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x08, 0x80, 0xf6, 0x2c, 0x02, 0x80, 0xf6,
	0x2c, 0x03, 0x48, 0x00, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x60, 0x0a, 0x06,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x67,
	0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x69, 0x74, 0x6f,
	0x70, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x54, 0x6f,
	0x53, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x08, 0x80, 0xf6, 0x2c, 0x02,
	0x80, 0xf6, 0x2c, 0x03, 0x48, 0x00, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x68,
	0x0a, 0x07, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x3d, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x67,
	0x69, 0x74, 0x6f, 0x70, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x54, 0x6f, 0x53, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x42, 0x0d,
	0x80, 0xf6, 0x2c, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x48, 0x00, 0x52,
	0x07, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x1a, 0x4d, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x12, 0x24, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x08,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x1a, 0x3d, 0x0a, 0x06, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x1f, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x09, 0x0a, 0x07, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x65,
	0x72, 0x42, 0x1f, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x88, 0xf6,
	0x2c, 0x01, 0x88, 0xf6, 0x2c, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0xf8,
	0x42, 0x01, 0x32, 0x95, 0x01, 0x0a, 0x06, 0x47, 0x69, 0x74, 0x6f, 0x70, 0x73, 0x12, 0x8a, 0x01,
	0x0a, 0x17, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x54, 0x6f, 0x53, 0x79,
	0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x65, 0x12, 0x34, 0x2e, 0x67, 0x69, 0x74, 0x6c,
	0x61, 0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x54, 0x6f, 0x53, 0x79, 0x6e,
	0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x35, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x67,
	0x69, 0x74, 0x6f, 0x70, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x54, 0x6f, 0x53, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x57, 0x5a, 0x55, 0x67, 0x69,
	0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2d,
	0x6f, 0x72, 0x67, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2d, 0x69, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2d, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73, 0x2f,
	0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_module_gitops_rpc_rpc_proto_rawDescOnce sync.Once
	file_internal_module_gitops_rpc_rpc_proto_rawDescData = file_internal_module_gitops_rpc_rpc_proto_rawDesc
)

func file_internal_module_gitops_rpc_rpc_proto_rawDescGZIP() []byte {
	file_internal_module_gitops_rpc_rpc_proto_rawDescOnce.Do(func() {
		file_internal_module_gitops_rpc_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_module_gitops_rpc_rpc_proto_rawDescData)
	})
	return file_internal_module_gitops_rpc_rpc_proto_rawDescData
}

var file_internal_module_gitops_rpc_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_internal_module_gitops_rpc_rpc_proto_goTypes = []interface{}{
	(*ObjectsToSynchronizeRequest)(nil),          // 0: gitlab.agent.gitops.rpc.ObjectsToSynchronizeRequest
	(*ObjectsToSynchronizeResponse)(nil),         // 1: gitlab.agent.gitops.rpc.ObjectsToSynchronizeResponse
	(*ObjectsToSynchronizeResponse_Header)(nil),  // 2: gitlab.agent.gitops.rpc.ObjectsToSynchronizeResponse.Header
	(*ObjectsToSynchronizeResponse_Object)(nil),  // 3: gitlab.agent.gitops.rpc.ObjectsToSynchronizeResponse.Object
	(*ObjectsToSynchronizeResponse_Trailer)(nil), // 4: gitlab.agent.gitops.rpc.ObjectsToSynchronizeResponse.Trailer
	(*agentcfg.PathCF)(nil),                      // 5: gitlab.agent.agentcfg.PathCF
}
var file_internal_module_gitops_rpc_rpc_proto_depIdxs = []int32{
	5, // 0: gitlab.agent.gitops.rpc.ObjectsToSynchronizeRequest.paths:type_name -> gitlab.agent.agentcfg.PathCF
	2, // 1: gitlab.agent.gitops.rpc.ObjectsToSynchronizeResponse.header:type_name -> gitlab.agent.gitops.rpc.ObjectsToSynchronizeResponse.Header
	3, // 2: gitlab.agent.gitops.rpc.ObjectsToSynchronizeResponse.object:type_name -> gitlab.agent.gitops.rpc.ObjectsToSynchronizeResponse.Object
	4, // 3: gitlab.agent.gitops.rpc.ObjectsToSynchronizeResponse.trailer:type_name -> gitlab.agent.gitops.rpc.ObjectsToSynchronizeResponse.Trailer
	0, // 4: gitlab.agent.gitops.rpc.Gitops.GetObjectsToSynchronize:input_type -> gitlab.agent.gitops.rpc.ObjectsToSynchronizeRequest
	1, // 5: gitlab.agent.gitops.rpc.Gitops.GetObjectsToSynchronize:output_type -> gitlab.agent.gitops.rpc.ObjectsToSynchronizeResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_internal_module_gitops_rpc_rpc_proto_init() }
func file_internal_module_gitops_rpc_rpc_proto_init() {
	if File_internal_module_gitops_rpc_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_module_gitops_rpc_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectsToSynchronizeRequest); i {
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
		file_internal_module_gitops_rpc_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectsToSynchronizeResponse); i {
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
		file_internal_module_gitops_rpc_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectsToSynchronizeResponse_Header); i {
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
		file_internal_module_gitops_rpc_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectsToSynchronizeResponse_Object); i {
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
		file_internal_module_gitops_rpc_rpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectsToSynchronizeResponse_Trailer); i {
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
	file_internal_module_gitops_rpc_rpc_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ObjectsToSynchronizeResponse_Header_)(nil),
		(*ObjectsToSynchronizeResponse_Object_)(nil),
		(*ObjectsToSynchronizeResponse_Trailer_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_module_gitops_rpc_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_module_gitops_rpc_rpc_proto_goTypes,
		DependencyIndexes: file_internal_module_gitops_rpc_rpc_proto_depIdxs,
		MessageInfos:      file_internal_module_gitops_rpc_rpc_proto_msgTypes,
	}.Build()
	File_internal_module_gitops_rpc_rpc_proto = out.File
	file_internal_module_gitops_rpc_rpc_proto_rawDesc = nil
	file_internal_module_gitops_rpc_rpc_proto_goTypes = nil
	file_internal_module_gitops_rpc_rpc_proto_depIdxs = nil
}
