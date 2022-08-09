// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.0
// source: internal/module/agent_tracker/rpc/rpc.proto

package rpc

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	agent_tracker "gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/module/agent_tracker"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetConnectedAgentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Request:
	//	*GetConnectedAgentsRequest_ProjectId
	//	*GetConnectedAgentsRequest_AgentId
	Request isGetConnectedAgentsRequest_Request `protobuf_oneof:"request"`
}

func (x *GetConnectedAgentsRequest) Reset() {
	*x = GetConnectedAgentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_module_agent_tracker_rpc_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConnectedAgentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConnectedAgentsRequest) ProtoMessage() {}

func (x *GetConnectedAgentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_module_agent_tracker_rpc_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConnectedAgentsRequest.ProtoReflect.Descriptor instead.
func (*GetConnectedAgentsRequest) Descriptor() ([]byte, []int) {
	return file_internal_module_agent_tracker_rpc_rpc_proto_rawDescGZIP(), []int{0}
}

func (m *GetConnectedAgentsRequest) GetRequest() isGetConnectedAgentsRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *GetConnectedAgentsRequest) GetProjectId() int64 {
	if x, ok := x.GetRequest().(*GetConnectedAgentsRequest_ProjectId); ok {
		return x.ProjectId
	}
	return 0
}

func (x *GetConnectedAgentsRequest) GetAgentId() int64 {
	if x, ok := x.GetRequest().(*GetConnectedAgentsRequest_AgentId); ok {
		return x.AgentId
	}
	return 0
}

type isGetConnectedAgentsRequest_Request interface {
	isGetConnectedAgentsRequest_Request()
}

type GetConnectedAgentsRequest_ProjectId struct {
	ProjectId int64 `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3,oneof"`
}

type GetConnectedAgentsRequest_AgentId struct {
	AgentId int64 `protobuf:"varint,2,opt,name=agent_id,json=agentId,proto3,oneof"`
}

func (*GetConnectedAgentsRequest_ProjectId) isGetConnectedAgentsRequest_Request() {}

func (*GetConnectedAgentsRequest_AgentId) isGetConnectedAgentsRequest_Request() {}

type GetConnectedAgentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Agents []*agent_tracker.ConnectedAgentInfo `protobuf:"bytes,1,rep,name=agents,proto3" json:"agents,omitempty"`
}

func (x *GetConnectedAgentsResponse) Reset() {
	*x = GetConnectedAgentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_module_agent_tracker_rpc_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConnectedAgentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConnectedAgentsResponse) ProtoMessage() {}

func (x *GetConnectedAgentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_module_agent_tracker_rpc_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConnectedAgentsResponse.ProtoReflect.Descriptor instead.
func (*GetConnectedAgentsResponse) Descriptor() ([]byte, []int) {
	return file_internal_module_agent_tracker_rpc_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *GetConnectedAgentsResponse) GetAgents() []*agent_tracker.ConnectedAgentInfo {
	if x != nil {
		return x.Agents
	}
	return nil
}

var File_internal_module_agent_tracker_rpc_rpc_proto protoreflect.FileDescriptor

var file_internal_module_agent_tracker_rpc_rpc_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2f,
	0x72, 0x70, 0x63, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67,
	0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x72, 0x70, 0x63, 0x1a, 0x31, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2f, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x19, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x08, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x07, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x42, 0x0e, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x03, 0xf8, 0x42, 0x01, 0x22, 0x64, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x46, 0x0a, 0x06, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x06, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x9e, 0x01, 0x0a, 0x0c, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x8d, 0x01, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x39, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e,
	0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x5e, 0x5a, 0x5c, 0x67,
	0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62,
	0x2d, 0x6f, 0x72, 0x67, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2d, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2d,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_internal_module_agent_tracker_rpc_rpc_proto_rawDescOnce sync.Once
	file_internal_module_agent_tracker_rpc_rpc_proto_rawDescData = file_internal_module_agent_tracker_rpc_rpc_proto_rawDesc
)

func file_internal_module_agent_tracker_rpc_rpc_proto_rawDescGZIP() []byte {
	file_internal_module_agent_tracker_rpc_rpc_proto_rawDescOnce.Do(func() {
		file_internal_module_agent_tracker_rpc_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_module_agent_tracker_rpc_rpc_proto_rawDescData)
	})
	return file_internal_module_agent_tracker_rpc_rpc_proto_rawDescData
}

var file_internal_module_agent_tracker_rpc_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internal_module_agent_tracker_rpc_rpc_proto_goTypes = []interface{}{
	(*GetConnectedAgentsRequest)(nil),        // 0: gitlab.agent.agent_tracker.rpc.GetConnectedAgentsRequest
	(*GetConnectedAgentsResponse)(nil),       // 1: gitlab.agent.agent_tracker.rpc.GetConnectedAgentsResponse
	(*agent_tracker.ConnectedAgentInfo)(nil), // 2: gitlab.agent.agent_tracker.ConnectedAgentInfo
}
var file_internal_module_agent_tracker_rpc_rpc_proto_depIdxs = []int32{
	2, // 0: gitlab.agent.agent_tracker.rpc.GetConnectedAgentsResponse.agents:type_name -> gitlab.agent.agent_tracker.ConnectedAgentInfo
	0, // 1: gitlab.agent.agent_tracker.rpc.AgentTracker.GetConnectedAgents:input_type -> gitlab.agent.agent_tracker.rpc.GetConnectedAgentsRequest
	1, // 2: gitlab.agent.agent_tracker.rpc.AgentTracker.GetConnectedAgents:output_type -> gitlab.agent.agent_tracker.rpc.GetConnectedAgentsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_module_agent_tracker_rpc_rpc_proto_init() }
func file_internal_module_agent_tracker_rpc_rpc_proto_init() {
	if File_internal_module_agent_tracker_rpc_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_module_agent_tracker_rpc_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConnectedAgentsRequest); i {
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
		file_internal_module_agent_tracker_rpc_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConnectedAgentsResponse); i {
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
	file_internal_module_agent_tracker_rpc_rpc_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*GetConnectedAgentsRequest_ProjectId)(nil),
		(*GetConnectedAgentsRequest_AgentId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_module_agent_tracker_rpc_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_module_agent_tracker_rpc_rpc_proto_goTypes,
		DependencyIndexes: file_internal_module_agent_tracker_rpc_rpc_proto_depIdxs,
		MessageInfos:      file_internal_module_agent_tracker_rpc_rpc_proto_msgTypes,
	}.Build()
	File_internal_module_agent_tracker_rpc_rpc_proto = out.File
	file_internal_module_agent_tracker_rpc_rpc_proto_rawDesc = nil
	file_internal_module_agent_tracker_rpc_rpc_proto_goTypes = nil
	file_internal_module_agent_tracker_rpc_rpc_proto_depIdxs = nil
}
