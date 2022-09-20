// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: internal/module/configuration_project/rpc/rpc.proto

package rpc

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	modserver "gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/module/modserver"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListAgentConfigFilesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Repository    *modserver.Repository    `protobuf:"bytes,1,opt,name=repository,proto3" json:"repository,omitempty"`
	GitalyAddress *modserver.GitalyAddress `protobuf:"bytes,2,opt,name=gitaly_address,json=gitalyAddress,proto3" json:"gitaly_address,omitempty"`
	DefaultBranch string                   `protobuf:"bytes,3,opt,name=default_branch,json=defaultBranch,proto3" json:"default_branch,omitempty"`
}

func (x *ListAgentConfigFilesRequest) Reset() {
	*x = ListAgentConfigFilesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_module_configuration_project_rpc_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAgentConfigFilesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAgentConfigFilesRequest) ProtoMessage() {}

func (x *ListAgentConfigFilesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_module_configuration_project_rpc_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAgentConfigFilesRequest.ProtoReflect.Descriptor instead.
func (*ListAgentConfigFilesRequest) Descriptor() ([]byte, []int) {
	return file_internal_module_configuration_project_rpc_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *ListAgentConfigFilesRequest) GetRepository() *modserver.Repository {
	if x != nil {
		return x.Repository
	}
	return nil
}

func (x *ListAgentConfigFilesRequest) GetGitalyAddress() *modserver.GitalyAddress {
	if x != nil {
		return x.GitalyAddress
	}
	return nil
}

func (x *ListAgentConfigFilesRequest) GetDefaultBranch() string {
	if x != nil {
		return x.DefaultBranch
	}
	return ""
}

type ListAgentConfigFilesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConfigFiles []*AgentConfigFile `protobuf:"bytes,1,rep,name=config_files,proto3" json:"config_files,omitempty"`
}

func (x *ListAgentConfigFilesResponse) Reset() {
	*x = ListAgentConfigFilesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_module_configuration_project_rpc_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAgentConfigFilesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAgentConfigFilesResponse) ProtoMessage() {}

func (x *ListAgentConfigFilesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_module_configuration_project_rpc_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAgentConfigFilesResponse.ProtoReflect.Descriptor instead.
func (*ListAgentConfigFilesResponse) Descriptor() ([]byte, []int) {
	return file_internal_module_configuration_project_rpc_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *ListAgentConfigFilesResponse) GetConfigFiles() []*AgentConfigFile {
	if x != nil {
		return x.ConfigFiles
	}
	return nil
}

type AgentConfigFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	AgentName string `protobuf:"bytes,2,opt,name=agent_name,proto3" json:"agent_name,omitempty"`
}

func (x *AgentConfigFile) Reset() {
	*x = AgentConfigFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_module_configuration_project_rpc_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentConfigFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentConfigFile) ProtoMessage() {}

func (x *AgentConfigFile) ProtoReflect() protoreflect.Message {
	mi := &file_internal_module_configuration_project_rpc_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentConfigFile.ProtoReflect.Descriptor instead.
func (*AgentConfigFile) Descriptor() ([]byte, []int) {
	return file_internal_module_configuration_project_rpc_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *AgentConfigFile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AgentConfigFile) GetAgentName() string {
	if x != nil {
		return x.AgentName
	}
	return ""
}

var File_internal_module_configuration_project_rpc_rpc_proto protoreflect.FileDescriptor

var file_internal_module_configuration_project_rpc_rpc_proto_rawDesc = []byte{
	0x0a, 0x33, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x70, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x72, 0x70, 0x63, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x6d, 0x6f, 0x64, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xea, 0x01, 0x0a, 0x1b, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x4c, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x52,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01,
	0x02, 0x10, 0x01, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12,
	0x56, 0x0a, 0x0e, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x79, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62,
	0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x47, 0x69, 0x74, 0x61, 0x6c, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0d, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x79,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x5f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x22, 0x7b,
	0x0a, 0x1c, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b,
	0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x0c, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x45, 0x0a, 0x0f, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x32, 0xbc, 0x01, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0xa3, 0x01, 0x0a, 0x14,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46,
	0x69, 0x6c, 0x65, 0x73, 0x12, 0x43, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46, 0x69, 0x6c,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x44, 0x2e, 0x67, 0x69, 0x74, 0x6c,
	0x61, 0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x66, 0x5a, 0x64, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2d, 0x6f, 0x72, 0x67, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67,
	0x69, 0x74, 0x6c, 0x61, 0x62, 0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x35, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_internal_module_configuration_project_rpc_rpc_proto_rawDescOnce sync.Once
	file_internal_module_configuration_project_rpc_rpc_proto_rawDescData = file_internal_module_configuration_project_rpc_rpc_proto_rawDesc
)

func file_internal_module_configuration_project_rpc_rpc_proto_rawDescGZIP() []byte {
	file_internal_module_configuration_project_rpc_rpc_proto_rawDescOnce.Do(func() {
		file_internal_module_configuration_project_rpc_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_module_configuration_project_rpc_rpc_proto_rawDescData)
	})
	return file_internal_module_configuration_project_rpc_rpc_proto_rawDescData
}

var file_internal_module_configuration_project_rpc_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_internal_module_configuration_project_rpc_rpc_proto_goTypes = []interface{}{
	(*ListAgentConfigFilesRequest)(nil),  // 0: gitlab.agent.configuration_project.rpc.ListAgentConfigFilesRequest
	(*ListAgentConfigFilesResponse)(nil), // 1: gitlab.agent.configuration_project.rpc.ListAgentConfigFilesResponse
	(*AgentConfigFile)(nil),              // 2: gitlab.agent.configuration_project.rpc.AgentConfigFile
	(*modserver.Repository)(nil),         // 3: gitlab.agent.modserver.Repository
	(*modserver.GitalyAddress)(nil),      // 4: gitlab.agent.modserver.GitalyAddress
}
var file_internal_module_configuration_project_rpc_rpc_proto_depIdxs = []int32{
	3, // 0: gitlab.agent.configuration_project.rpc.ListAgentConfigFilesRequest.repository:type_name -> gitlab.agent.modserver.Repository
	4, // 1: gitlab.agent.configuration_project.rpc.ListAgentConfigFilesRequest.gitaly_address:type_name -> gitlab.agent.modserver.GitalyAddress
	2, // 2: gitlab.agent.configuration_project.rpc.ListAgentConfigFilesResponse.config_files:type_name -> gitlab.agent.configuration_project.rpc.AgentConfigFile
	0, // 3: gitlab.agent.configuration_project.rpc.ConfigurationProject.ListAgentConfigFiles:input_type -> gitlab.agent.configuration_project.rpc.ListAgentConfigFilesRequest
	1, // 4: gitlab.agent.configuration_project.rpc.ConfigurationProject.ListAgentConfigFiles:output_type -> gitlab.agent.configuration_project.rpc.ListAgentConfigFilesResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_internal_module_configuration_project_rpc_rpc_proto_init() }
func file_internal_module_configuration_project_rpc_rpc_proto_init() {
	if File_internal_module_configuration_project_rpc_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_module_configuration_project_rpc_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAgentConfigFilesRequest); i {
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
		file_internal_module_configuration_project_rpc_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAgentConfigFilesResponse); i {
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
		file_internal_module_configuration_project_rpc_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentConfigFile); i {
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
			RawDescriptor: file_internal_module_configuration_project_rpc_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_module_configuration_project_rpc_rpc_proto_goTypes,
		DependencyIndexes: file_internal_module_configuration_project_rpc_rpc_proto_depIdxs,
		MessageInfos:      file_internal_module_configuration_project_rpc_rpc_proto_msgTypes,
	}.Build()
	File_internal_module_configuration_project_rpc_rpc_proto = out.File
	file_internal_module_configuration_project_rpc_rpc_proto_rawDesc = nil
	file_internal_module_configuration_project_rpc_rpc_proto_goTypes = nil
	file_internal_module_configuration_project_rpc_rpc_proto_depIdxs = nil
}
