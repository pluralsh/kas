// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.0
// source: internal/module/agent_tracker/agent_tracker.proto

package agent_tracker

import (
	reflect "reflect"
	sync "sync"

	modshared "gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v14/internal/module/modshared"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ConnectedAgentInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentMeta    *modshared.AgentMeta   `protobuf:"bytes,1,opt,name=agent_meta,json=agentMeta,proto3" json:"agent_meta,omitempty"`
	ConnectedAt  *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=connected_at,json=connectedAt,proto3" json:"connected_at,omitempty"`
	ConnectionId int64                  `protobuf:"varint,3,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	AgentId      int64                  `protobuf:"varint,4,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	ProjectId    int64                  `protobuf:"varint,5,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *ConnectedAgentInfo) Reset() {
	*x = ConnectedAgentInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_module_agent_tracker_agent_tracker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectedAgentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectedAgentInfo) ProtoMessage() {}

func (x *ConnectedAgentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_internal_module_agent_tracker_agent_tracker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectedAgentInfo.ProtoReflect.Descriptor instead.
func (*ConnectedAgentInfo) Descriptor() ([]byte, []int) {
	return file_internal_module_agent_tracker_agent_tracker_proto_rawDescGZIP(), []int{0}
}

func (x *ConnectedAgentInfo) GetAgentMeta() *modshared.AgentMeta {
	if x != nil {
		return x.AgentMeta
	}
	return nil
}

func (x *ConnectedAgentInfo) GetConnectedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ConnectedAt
	}
	return nil
}

func (x *ConnectedAgentInfo) GetConnectionId() int64 {
	if x != nil {
		return x.ConnectionId
	}
	return 0
}

func (x *ConnectedAgentInfo) GetAgentId() int64 {
	if x != nil {
		return x.AgentId
	}
	return 0
}

func (x *ConnectedAgentInfo) GetProjectId() int64 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

var File_internal_module_agent_tracker_agent_tracker_proto protoreflect.FileDescriptor

var file_internal_module_agent_tracker_agent_tracker_proto_rawDesc = []byte{
	0x0a, 0x31, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2f,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x29, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x6d, 0x6f, 0x64, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf4, 0x01, 0x0a, 0x12,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x40, 0x0a, 0x0a, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x09, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x4d, 0x65, 0x74, 0x61, 0x12, 0x3d, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x49, 0x64, 0x42, 0x5a, 0x5a, 0x58, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2d, 0x6f, 0x72, 0x67, 0x2f, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x34,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_module_agent_tracker_agent_tracker_proto_rawDescOnce sync.Once
	file_internal_module_agent_tracker_agent_tracker_proto_rawDescData = file_internal_module_agent_tracker_agent_tracker_proto_rawDesc
)

func file_internal_module_agent_tracker_agent_tracker_proto_rawDescGZIP() []byte {
	file_internal_module_agent_tracker_agent_tracker_proto_rawDescOnce.Do(func() {
		file_internal_module_agent_tracker_agent_tracker_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_module_agent_tracker_agent_tracker_proto_rawDescData)
	})
	return file_internal_module_agent_tracker_agent_tracker_proto_rawDescData
}

var file_internal_module_agent_tracker_agent_tracker_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internal_module_agent_tracker_agent_tracker_proto_goTypes = []interface{}{
	(*ConnectedAgentInfo)(nil),    // 0: gitlab.agent.agent_tracker.ConnectedAgentInfo
	(*modshared.AgentMeta)(nil),   // 1: gitlab.agent.modshared.AgentMeta
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_internal_module_agent_tracker_agent_tracker_proto_depIdxs = []int32{
	1, // 0: gitlab.agent.agent_tracker.ConnectedAgentInfo.agent_meta:type_name -> gitlab.agent.modshared.AgentMeta
	2, // 1: gitlab.agent.agent_tracker.ConnectedAgentInfo.connected_at:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_internal_module_agent_tracker_agent_tracker_proto_init() }
func file_internal_module_agent_tracker_agent_tracker_proto_init() {
	if File_internal_module_agent_tracker_agent_tracker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_module_agent_tracker_agent_tracker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectedAgentInfo); i {
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
			RawDescriptor: file_internal_module_agent_tracker_agent_tracker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_module_agent_tracker_agent_tracker_proto_goTypes,
		DependencyIndexes: file_internal_module_agent_tracker_agent_tracker_proto_depIdxs,
		MessageInfos:      file_internal_module_agent_tracker_agent_tracker_proto_msgTypes,
	}.Build()
	File_internal_module_agent_tracker_agent_tracker_proto = out.File
	file_internal_module_agent_tracker_agent_tracker_proto_rawDesc = nil
	file_internal_module_agent_tracker_agent_tracker_proto_goTypes = nil
	file_internal_module_agent_tracker_agent_tracker_proto_depIdxs = nil
}
