// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.7
// source: pkg/entity/entity.proto

package entity

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AgentMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version           string             `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	CommitId          string             `protobuf:"bytes,2,opt,name=commit_id,proto3" json:"commit_id,omitempty"`
	PodNamespace      string             `protobuf:"bytes,3,opt,name=pod_namespace,proto3" json:"pod_namespace,omitempty"`
	PodName           string             `protobuf:"bytes,4,opt,name=pod_name,proto3" json:"pod_name,omitempty"`
	KubernetesVersion *KubernetesVersion `protobuf:"bytes,5,opt,name=kubernetes_version,proto3" json:"kubernetes_version,omitempty"`
}

func (x *AgentMeta) Reset() {
	*x = AgentMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_entity_entity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentMeta) ProtoMessage() {}

func (x *AgentMeta) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_entity_entity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentMeta.ProtoReflect.Descriptor instead.
func (*AgentMeta) Descriptor() ([]byte, []int) {
	return file_pkg_entity_entity_proto_rawDescGZIP(), []int{0}
}

func (x *AgentMeta) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *AgentMeta) GetCommitId() string {
	if x != nil {
		return x.CommitId
	}
	return ""
}

func (x *AgentMeta) GetPodNamespace() string {
	if x != nil {
		return x.PodNamespace
	}
	return ""
}

func (x *AgentMeta) GetPodName() string {
	if x != nil {
		return x.PodName
	}
	return ""
}

func (x *AgentMeta) GetKubernetesVersion() *KubernetesVersion {
	if x != nil {
		return x.KubernetesVersion
	}
	return nil
}

type KubernetesVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Major      string `protobuf:"bytes,1,opt,name=major,proto3" json:"major,omitempty"`
	Minor      string `protobuf:"bytes,2,opt,name=minor,proto3" json:"minor,omitempty"`
	GitVersion string `protobuf:"bytes,3,opt,name=git_version,proto3" json:"git_version,omitempty"`
	Platform   string `protobuf:"bytes,4,opt,name=platform,proto3" json:"platform,omitempty"`
}

func (x *KubernetesVersion) Reset() {
	*x = KubernetesVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_entity_entity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubernetesVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesVersion) ProtoMessage() {}

func (x *KubernetesVersion) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_entity_entity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesVersion.ProtoReflect.Descriptor instead.
func (*KubernetesVersion) Descriptor() ([]byte, []int) {
	return file_pkg_entity_entity_proto_rawDescGZIP(), []int{1}
}

func (x *KubernetesVersion) GetMajor() string {
	if x != nil {
		return x.Major
	}
	return ""
}

func (x *KubernetesVersion) GetMinor() string {
	if x != nil {
		return x.Minor
	}
	return ""
}

func (x *KubernetesVersion) GetGitVersion() string {
	if x != nil {
		return x.GitVersion
	}
	return ""
}

func (x *KubernetesVersion) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

type GitalyInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address  string            `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Token    string            `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Features map[string]string `protobuf:"bytes,3,rep,name=features,proto3" json:"features,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GitalyInfo) Reset() {
	*x = GitalyInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_entity_entity_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GitalyInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GitalyInfo) ProtoMessage() {}

func (x *GitalyInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_entity_entity_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GitalyInfo.ProtoReflect.Descriptor instead.
func (*GitalyInfo) Descriptor() ([]byte, []int) {
	return file_pkg_entity_entity_proto_rawDescGZIP(), []int{2}
}

func (x *GitalyInfo) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *GitalyInfo) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GitalyInfo) GetFeatures() map[string]string {
	if x != nil {
		return x.Features
	}
	return nil
}

type GitalyRepository struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StorageName                   string   `protobuf:"bytes,2,opt,name=storage_name,proto3" json:"storage_name,omitempty"`
	RelativePath                  string   `protobuf:"bytes,3,opt,name=relative_path,proto3" json:"relative_path,omitempty"`
	GitObjectDirectory            string   `protobuf:"bytes,4,opt,name=git_object_directory,proto3" json:"git_object_directory,omitempty"`
	GitAlternateObjectDirectories []string `protobuf:"bytes,5,rep,name=git_alternate_object_directories,proto3" json:"git_alternate_object_directories,omitempty"`
	GlRepository                  string   `protobuf:"bytes,6,opt,name=gl_repository,proto3" json:"gl_repository,omitempty"`
	GlProjectPath                 string   `protobuf:"bytes,8,opt,name=gl_project_path,proto3" json:"gl_project_path,omitempty"`
}

func (x *GitalyRepository) Reset() {
	*x = GitalyRepository{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_entity_entity_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GitalyRepository) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GitalyRepository) ProtoMessage() {}

func (x *GitalyRepository) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_entity_entity_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GitalyRepository.ProtoReflect.Descriptor instead.
func (*GitalyRepository) Descriptor() ([]byte, []int) {
	return file_pkg_entity_entity_proto_rawDescGZIP(), []int{3}
}

func (x *GitalyRepository) GetStorageName() string {
	if x != nil {
		return x.StorageName
	}
	return ""
}

func (x *GitalyRepository) GetRelativePath() string {
	if x != nil {
		return x.RelativePath
	}
	return ""
}

func (x *GitalyRepository) GetGitObjectDirectory() string {
	if x != nil {
		return x.GitObjectDirectory
	}
	return ""
}

func (x *GitalyRepository) GetGitAlternateObjectDirectories() []string {
	if x != nil {
		return x.GitAlternateObjectDirectories
	}
	return nil
}

func (x *GitalyRepository) GetGlRepository() string {
	if x != nil {
		return x.GlRepository
	}
	return ""
}

func (x *GitalyRepository) GetGlProjectPath() string {
	if x != nil {
		return x.GlProjectPath
	}
	return ""
}

var File_pkg_entity_entity_proto protoreflect.FileDescriptor

var file_pkg_entity_entity_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x6b, 0x67, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x67, 0x69, 0x74, 0x6c, 0x61,
	0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdd, 0x01, 0x0a, 0x09, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x12, 0x24, 0x0a,
	0x0d, 0x70, 0x6f, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x6f, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x56, 0x0a, 0x12, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67, 0x69,
	0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x12, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x7d, 0x0a, 0x11, 0x4b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x61, 0x6a,
	0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x67, 0x69, 0x74, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x67,
	0x69, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x22, 0xcd, 0x01, 0x0a, 0x0a, 0x47, 0x69, 0x74, 0x61, 0x6c,
	0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x21, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x49,
	0x0a, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2d, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x47, 0x69, 0x74, 0x61, 0x6c, 0x79, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x1a, 0x3b, 0x0a, 0x0d, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xd0, 0x02, 0x0a, 0x10, 0x47, 0x69, 0x74, 0x61, 0x6c,
	0x79, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x2b, 0x0a, 0x0c, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x0c, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x0d, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x76, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x0d, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x76, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x12, 0x32, 0x0a, 0x14, 0x67, 0x69, 0x74, 0x5f, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x67, 0x69, 0x74, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x4a, 0x0a, 0x20, 0x67,
	0x69, 0x74, 0x5f, 0x61, 0x6c, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x20, 0x67, 0x69, 0x74, 0x5f, 0x61, 0x6c, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x2d, 0x0a, 0x0d, 0x67, 0x6c, 0x5f, 0x72, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x0d, 0x67, 0x6c, 0x5f, 0x72, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x31, 0x0a, 0x0f, 0x67, 0x6c, 0x5f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x0f, 0x67, 0x6c, 0x5f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74,
	0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2d, 0x6f,
	0x72, 0x67, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2d, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_entity_entity_proto_rawDescOnce sync.Once
	file_pkg_entity_entity_proto_rawDescData = file_pkg_entity_entity_proto_rawDesc
)

func file_pkg_entity_entity_proto_rawDescGZIP() []byte {
	file_pkg_entity_entity_proto_rawDescOnce.Do(func() {
		file_pkg_entity_entity_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_entity_entity_proto_rawDescData)
	})
	return file_pkg_entity_entity_proto_rawDescData
}

var file_pkg_entity_entity_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_entity_entity_proto_goTypes = []interface{}{
	(*AgentMeta)(nil),         // 0: gitlab.agent.entity.AgentMeta
	(*KubernetesVersion)(nil), // 1: gitlab.agent.entity.KubernetesVersion
	(*GitalyInfo)(nil),        // 2: gitlab.agent.entity.GitalyInfo
	(*GitalyRepository)(nil),  // 3: gitlab.agent.entity.GitalyRepository
	nil,                       // 4: gitlab.agent.entity.GitalyInfo.FeaturesEntry
}
var file_pkg_entity_entity_proto_depIdxs = []int32{
	1, // 0: gitlab.agent.entity.AgentMeta.kubernetes_version:type_name -> gitlab.agent.entity.KubernetesVersion
	4, // 1: gitlab.agent.entity.GitalyInfo.features:type_name -> gitlab.agent.entity.GitalyInfo.FeaturesEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_entity_entity_proto_init() }
func file_pkg_entity_entity_proto_init() {
	if File_pkg_entity_entity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_entity_entity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentMeta); i {
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
		file_pkg_entity_entity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubernetesVersion); i {
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
		file_pkg_entity_entity_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GitalyInfo); i {
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
		file_pkg_entity_entity_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GitalyRepository); i {
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
			RawDescriptor: file_pkg_entity_entity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_entity_entity_proto_goTypes,
		DependencyIndexes: file_pkg_entity_entity_proto_depIdxs,
		MessageInfos:      file_pkg_entity_entity_proto_msgTypes,
	}.Build()
	File_pkg_entity_entity_proto = out.File
	file_pkg_entity_entity_proto_rawDesc = nil
	file_pkg_entity_entity_proto_goTypes = nil
	file_pkg_entity_entity_proto_depIdxs = nil
}
