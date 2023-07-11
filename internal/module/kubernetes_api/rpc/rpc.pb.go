// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.7
// source: internal/module/kubernetes_api/rpc/rpc.proto

package rpc

import (
	reflect "reflect"
	sync "sync"

	grpctool "gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/tool/grpctool"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HeaderExtra struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImpConfig *ImpersonationConfig `protobuf:"bytes,1,opt,name=imp_config,json=impConfig,proto3" json:"imp_config,omitempty"`
}

func (x *HeaderExtra) Reset() {
	*x = HeaderExtra{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_module_kubernetes_api_rpc_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeaderExtra) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeaderExtra) ProtoMessage() {}

func (x *HeaderExtra) ProtoReflect() protoreflect.Message {
	mi := &file_internal_module_kubernetes_api_rpc_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeaderExtra.ProtoReflect.Descriptor instead.
func (*HeaderExtra) Descriptor() ([]byte, []int) {
	return file_internal_module_kubernetes_api_rpc_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *HeaderExtra) GetImpConfig() *ImpersonationConfig {
	if x != nil {
		return x.ImpConfig
	}
	return nil
}

type ImpersonationConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string         `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Groups   []string       `protobuf:"bytes,2,rep,name=groups,proto3" json:"groups,omitempty"`
	Uid      string         `protobuf:"bytes,3,opt,name=uid,proto3" json:"uid,omitempty"`
	Extra    []*ExtraKeyVal `protobuf:"bytes,4,rep,name=extra,proto3" json:"extra,omitempty"`
}

func (x *ImpersonationConfig) Reset() {
	*x = ImpersonationConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_module_kubernetes_api_rpc_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImpersonationConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImpersonationConfig) ProtoMessage() {}

func (x *ImpersonationConfig) ProtoReflect() protoreflect.Message {
	mi := &file_internal_module_kubernetes_api_rpc_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImpersonationConfig.ProtoReflect.Descriptor instead.
func (*ImpersonationConfig) Descriptor() ([]byte, []int) {
	return file_internal_module_kubernetes_api_rpc_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *ImpersonationConfig) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ImpersonationConfig) GetGroups() []string {
	if x != nil {
		return x.Groups
	}
	return nil
}

func (x *ImpersonationConfig) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *ImpersonationConfig) GetExtra() []*ExtraKeyVal {
	if x != nil {
		return x.Extra
	}
	return nil
}

type ExtraKeyVal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Val []string `protobuf:"bytes,2,rep,name=val,proto3" json:"val,omitempty"`
}

func (x *ExtraKeyVal) Reset() {
	*x = ExtraKeyVal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_module_kubernetes_api_rpc_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtraKeyVal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtraKeyVal) ProtoMessage() {}

func (x *ExtraKeyVal) ProtoReflect() protoreflect.Message {
	mi := &file_internal_module_kubernetes_api_rpc_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtraKeyVal.ProtoReflect.Descriptor instead.
func (*ExtraKeyVal) Descriptor() ([]byte, []int) {
	return file_internal_module_kubernetes_api_rpc_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *ExtraKeyVal) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ExtraKeyVal) GetVal() []string {
	if x != nil {
		return x.Val
	}
	return nil
}

var File_internal_module_kubernetes_api_rpc_rpc_proto protoreflect.FileDescriptor

var file_internal_module_kubernetes_api_rpc_rpc_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f,
	0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x6b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x1a,
	0x25, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x6f, 0x6f, 0x6c, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x74, 0x6f, 0x6f, 0x6c, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x74, 0x6f, 0x6f, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x62, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x45, 0x78, 0x74, 0x72, 0x61, 0x12, 0x53, 0x0a, 0x0a, 0x69, 0x6d, 0x70, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x67, 0x69, 0x74, 0x6c,
	0x61, 0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x6d, 0x70, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x09, 0x69, 0x6d, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x9f, 0x01, 0x0a, 0x13, 0x49,
	0x6d, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x42, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72,
	0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62,
	0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x4b,
	0x65, 0x79, 0x56, 0x61, 0x6c, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x22, 0x31, 0x0a, 0x0b,
	0x45, 0x78, 0x74, 0x72, 0x61, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x32,
	0x6d, 0x0a, 0x0d, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x41, 0x70, 0x69,
	0x12, 0x5c, 0x0a, 0x0b, 0x4d, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x22, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x74, 0x6f, 0x6f, 0x6c, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x74, 0x6f, 0x6f, 0x6c, 0x2e, 0x48, 0x74, 0x74, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x5f,
	0x5a, 0x5d, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x69, 0x74,
	0x6c, 0x61, 0x62, 0x2d, 0x6f, 0x72, 0x67, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2d,
	0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x69, 0x74, 0x6c,
	0x61, 0x62, 0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x6b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x70, 0x63, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_module_kubernetes_api_rpc_rpc_proto_rawDescOnce sync.Once
	file_internal_module_kubernetes_api_rpc_rpc_proto_rawDescData = file_internal_module_kubernetes_api_rpc_rpc_proto_rawDesc
)

func file_internal_module_kubernetes_api_rpc_rpc_proto_rawDescGZIP() []byte {
	file_internal_module_kubernetes_api_rpc_rpc_proto_rawDescOnce.Do(func() {
		file_internal_module_kubernetes_api_rpc_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_module_kubernetes_api_rpc_rpc_proto_rawDescData)
	})
	return file_internal_module_kubernetes_api_rpc_rpc_proto_rawDescData
}

var file_internal_module_kubernetes_api_rpc_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_internal_module_kubernetes_api_rpc_rpc_proto_goTypes = []interface{}{
	(*HeaderExtra)(nil),           // 0: gitlab.agent.kubernetes_api.rpc.HeaderExtra
	(*ImpersonationConfig)(nil),   // 1: gitlab.agent.kubernetes_api.rpc.ImpersonationConfig
	(*ExtraKeyVal)(nil),           // 2: gitlab.agent.kubernetes_api.rpc.ExtraKeyVal
	(*grpctool.HttpRequest)(nil),  // 3: gitlab.agent.grpctool.HttpRequest
	(*grpctool.HttpResponse)(nil), // 4: gitlab.agent.grpctool.HttpResponse
}
var file_internal_module_kubernetes_api_rpc_rpc_proto_depIdxs = []int32{
	1, // 0: gitlab.agent.kubernetes_api.rpc.HeaderExtra.imp_config:type_name -> gitlab.agent.kubernetes_api.rpc.ImpersonationConfig
	2, // 1: gitlab.agent.kubernetes_api.rpc.ImpersonationConfig.extra:type_name -> gitlab.agent.kubernetes_api.rpc.ExtraKeyVal
	3, // 2: gitlab.agent.kubernetes_api.rpc.KubernetesApi.MakeRequest:input_type -> gitlab.agent.grpctool.HttpRequest
	4, // 3: gitlab.agent.kubernetes_api.rpc.KubernetesApi.MakeRequest:output_type -> gitlab.agent.grpctool.HttpResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_internal_module_kubernetes_api_rpc_rpc_proto_init() }
func file_internal_module_kubernetes_api_rpc_rpc_proto_init() {
	if File_internal_module_kubernetes_api_rpc_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_module_kubernetes_api_rpc_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeaderExtra); i {
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
		file_internal_module_kubernetes_api_rpc_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImpersonationConfig); i {
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
		file_internal_module_kubernetes_api_rpc_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtraKeyVal); i {
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
			RawDescriptor: file_internal_module_kubernetes_api_rpc_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_module_kubernetes_api_rpc_rpc_proto_goTypes,
		DependencyIndexes: file_internal_module_kubernetes_api_rpc_rpc_proto_depIdxs,
		MessageInfos:      file_internal_module_kubernetes_api_rpc_rpc_proto_msgTypes,
	}.Build()
	File_internal_module_kubernetes_api_rpc_rpc_proto = out.File
	file_internal_module_kubernetes_api_rpc_rpc_proto_rawDesc = nil
	file_internal_module_kubernetes_api_rpc_rpc_proto_goTypes = nil
	file_internal_module_kubernetes_api_rpc_rpc_proto_depIdxs = nil
}
