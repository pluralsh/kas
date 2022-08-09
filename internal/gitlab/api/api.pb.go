// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.0
// source: internal/gitlab/api/api.proto

package api

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type Configuration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DefaultNamespace string                 `protobuf:"bytes,1,opt,name=default_namespace,proto3" json:"default_namespace,omitempty"`
	AccessAs         *agentcfg.CiAccessAsCF `protobuf:"bytes,2,opt,name=access_as,proto3" json:"access_as,omitempty"`
}

func (x *Configuration) Reset() {
	*x = Configuration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_gitlab_api_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Configuration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Configuration) ProtoMessage() {}

func (x *Configuration) ProtoReflect() protoreflect.Message {
	mi := &file_internal_gitlab_api_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Configuration.ProtoReflect.Descriptor instead.
func (*Configuration) Descriptor() ([]byte, []int) {
	return file_internal_gitlab_api_api_proto_rawDescGZIP(), []int{0}
}

func (x *Configuration) GetDefaultNamespace() string {
	if x != nil {
		return x.DefaultNamespace
	}
	return ""
}

func (x *Configuration) GetAccessAs() *agentcfg.CiAccessAsCF {
	if x != nil {
		return x.AccessAs
	}
	return nil
}

type AllowedAgent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ConfigProject *ConfigProject `protobuf:"bytes,2,opt,name=config_project,proto3" json:"config_project,omitempty"`
	Configuration *Configuration `protobuf:"bytes,3,opt,name=configuration,proto3" json:"configuration,omitempty"`
}

func (x *AllowedAgent) Reset() {
	*x = AllowedAgent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_gitlab_api_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllowedAgent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllowedAgent) ProtoMessage() {}

func (x *AllowedAgent) ProtoReflect() protoreflect.Message {
	mi := &file_internal_gitlab_api_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllowedAgent.ProtoReflect.Descriptor instead.
func (*AllowedAgent) Descriptor() ([]byte, []int) {
	return file_internal_gitlab_api_api_proto_rawDescGZIP(), []int{1}
}

func (x *AllowedAgent) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AllowedAgent) GetConfigProject() *ConfigProject {
	if x != nil {
		return x.ConfigProject
	}
	return nil
}

func (x *AllowedAgent) GetConfiguration() *Configuration {
	if x != nil {
		return x.Configuration
	}
	return nil
}

type ConfigProject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ConfigProject) Reset() {
	*x = ConfigProject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_gitlab_api_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigProject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigProject) ProtoMessage() {}

func (x *ConfigProject) ProtoReflect() protoreflect.Message {
	mi := &file_internal_gitlab_api_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigProject.ProtoReflect.Descriptor instead.
func (*ConfigProject) Descriptor() ([]byte, []int) {
	return file_internal_gitlab_api_api_proto_rawDescGZIP(), []int{2}
}

func (x *ConfigProject) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Pipeline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Pipeline) Reset() {
	*x = Pipeline{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_gitlab_api_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pipeline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pipeline) ProtoMessage() {}

func (x *Pipeline) ProtoReflect() protoreflect.Message {
	mi := &file_internal_gitlab_api_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pipeline.ProtoReflect.Descriptor instead.
func (*Pipeline) Descriptor() ([]byte, []int) {
	return file_internal_gitlab_api_api_proto_rawDescGZIP(), []int{3}
}

func (x *Pipeline) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Groups []*Group `protobuf:"bytes,2,rep,name=groups,proto3" json:"groups,omitempty"`
}

func (x *Project) Reset() {
	*x = Project{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_gitlab_api_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project) ProtoMessage() {}

func (x *Project) ProtoReflect() protoreflect.Message {
	mi := &file_internal_gitlab_api_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project.ProtoReflect.Descriptor instead.
func (*Project) Descriptor() ([]byte, []int) {
	return file_internal_gitlab_api_api_proto_rawDescGZIP(), []int{4}
}

func (x *Project) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Project) GetGroups() []*Group {
	if x != nil {
		return x.Groups
	}
	return nil
}

type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_gitlab_api_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_internal_gitlab_api_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_internal_gitlab_api_api_proto_rawDescGZIP(), []int{5}
}

func (x *Group) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Job struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Job) Reset() {
	*x = Job{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_gitlab_api_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Job) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Job) ProtoMessage() {}

func (x *Job) ProtoReflect() protoreflect.Message {
	mi := &file_internal_gitlab_api_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Job.ProtoReflect.Descriptor instead.
func (*Job) Descriptor() ([]byte, []int) {
	return file_internal_gitlab_api_api_proto_rawDescGZIP(), []int{6}
}

func (x *Job) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_gitlab_api_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_internal_gitlab_api_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_internal_gitlab_api_api_proto_rawDescGZIP(), []int{7}
}

func (x *User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type Environment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slug string `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug,omitempty"`
}

func (x *Environment) Reset() {
	*x = Environment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_gitlab_api_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Environment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Environment) ProtoMessage() {}

func (x *Environment) ProtoReflect() protoreflect.Message {
	mi := &file_internal_gitlab_api_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Environment.ProtoReflect.Descriptor instead.
func (*Environment) Descriptor() ([]byte, []int) {
	return file_internal_gitlab_api_api_proto_rawDescGZIP(), []int{8}
}

func (x *Environment) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

type AllowedAgentsForJob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AllowedAgents []*AllowedAgent `protobuf:"bytes,1,rep,name=allowed_agents,proto3" json:"allowed_agents,omitempty"`
	Job           *Job            `protobuf:"bytes,2,opt,name=job,proto3" json:"job,omitempty"`
	Pipeline      *Pipeline       `protobuf:"bytes,3,opt,name=pipeline,proto3" json:"pipeline,omitempty"`
	Project       *Project        `protobuf:"bytes,4,opt,name=project,proto3" json:"project,omitempty"`
	User          *User           `protobuf:"bytes,5,opt,name=user,proto3" json:"user,omitempty"`
	Environment   *Environment    `protobuf:"bytes,6,opt,name=environment,proto3" json:"environment,omitempty"`
}

func (x *AllowedAgentsForJob) Reset() {
	*x = AllowedAgentsForJob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_gitlab_api_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllowedAgentsForJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllowedAgentsForJob) ProtoMessage() {}

func (x *AllowedAgentsForJob) ProtoReflect() protoreflect.Message {
	mi := &file_internal_gitlab_api_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllowedAgentsForJob.ProtoReflect.Descriptor instead.
func (*AllowedAgentsForJob) Descriptor() ([]byte, []int) {
	return file_internal_gitlab_api_api_proto_rawDescGZIP(), []int{9}
}

func (x *AllowedAgentsForJob) GetAllowedAgents() []*AllowedAgent {
	if x != nil {
		return x.AllowedAgents
	}
	return nil
}

func (x *AllowedAgentsForJob) GetJob() *Job {
	if x != nil {
		return x.Job
	}
	return nil
}

func (x *AllowedAgentsForJob) GetPipeline() *Pipeline {
	if x != nil {
		return x.Pipeline
	}
	return nil
}

func (x *AllowedAgentsForJob) GetProject() *Project {
	if x != nil {
		return x.Project
	}
	return nil
}

func (x *AllowedAgentsForJob) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *AllowedAgentsForJob) GetEnvironment() *Environment {
	if x != nil {
		return x.Environment
	}
	return nil
}

var File_internal_gitlab_api_api_proto protoreflect.FileDescriptor

var file_internal_gitlab_api_api_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x69, 0x74, 0x6c, 0x61,
	0x62, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x17, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x69,
	0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x1b, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x63, 0x66, 0x67, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x63, 0x66, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80,
	0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2c, 0x0a, 0x11, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x41,
	0x0a, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x61, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x63, 0x66, 0x67, 0x2e, 0x43, 0x69, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x41, 0x73, 0x43, 0x46, 0x52, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x61,
	0x73, 0x22, 0xc6, 0x01, 0x0a, 0x0c, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x58, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67, 0x69, 0x74,
	0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x4c, 0x0a, 0x0d,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x1f, 0x0a, 0x0d, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1a, 0x0a, 0x08, 0x50,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x51, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x36, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0x17, 0x0a, 0x05, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x15, 0x0a, 0x03, 0x4a, 0x6f, 0x62, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3b, 0x0a, 0x04, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x23, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2a, 0x0a, 0x0b, 0x45, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x04, 0x73,
	0x6c, 0x75, 0x67, 0x22, 0xb2, 0x03, 0x0a, 0x13, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x73, 0x46, 0x6f, 0x72, 0x4a, 0x6f, 0x62, 0x12, 0x4d, 0x0a, 0x0e, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x64, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x0e, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x64, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x38, 0x0a, 0x03, 0x6a, 0x6f,
	0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62,
	0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x4a, 0x6f, 0x62, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52,
	0x03, 0x6a, 0x6f, 0x62, 0x12, 0x47, 0x0a, 0x08, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01,
	0x02, 0x10, 0x01, 0x52, 0x08, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x44, 0x0a,
	0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x69,
	0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x3b, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x12, 0x46, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x65, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x50, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x6c,
	0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2d, 0x6f, 0x72,
	0x67, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2d, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_internal_gitlab_api_api_proto_rawDescOnce sync.Once
	file_internal_gitlab_api_api_proto_rawDescData = file_internal_gitlab_api_api_proto_rawDesc
)

func file_internal_gitlab_api_api_proto_rawDescGZIP() []byte {
	file_internal_gitlab_api_api_proto_rawDescOnce.Do(func() {
		file_internal_gitlab_api_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_gitlab_api_api_proto_rawDescData)
	})
	return file_internal_gitlab_api_api_proto_rawDescData
}

var file_internal_gitlab_api_api_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_internal_gitlab_api_api_proto_goTypes = []interface{}{
	(*Configuration)(nil),         // 0: gitlab.agent.gitlab.api.Configuration
	(*AllowedAgent)(nil),          // 1: gitlab.agent.gitlab.api.AllowedAgent
	(*ConfigProject)(nil),         // 2: gitlab.agent.gitlab.api.ConfigProject
	(*Pipeline)(nil),              // 3: gitlab.agent.gitlab.api.Pipeline
	(*Project)(nil),               // 4: gitlab.agent.gitlab.api.Project
	(*Group)(nil),                 // 5: gitlab.agent.gitlab.api.Group
	(*Job)(nil),                   // 6: gitlab.agent.gitlab.api.Job
	(*User)(nil),                  // 7: gitlab.agent.gitlab.api.User
	(*Environment)(nil),           // 8: gitlab.agent.gitlab.api.Environment
	(*AllowedAgentsForJob)(nil),   // 9: gitlab.agent.gitlab.api.AllowedAgentsForJob
	(*agentcfg.CiAccessAsCF)(nil), // 10: gitlab.agent.agentcfg.CiAccessAsCF
}
var file_internal_gitlab_api_api_proto_depIdxs = []int32{
	10, // 0: gitlab.agent.gitlab.api.Configuration.access_as:type_name -> gitlab.agent.agentcfg.CiAccessAsCF
	2,  // 1: gitlab.agent.gitlab.api.AllowedAgent.config_project:type_name -> gitlab.agent.gitlab.api.ConfigProject
	0,  // 2: gitlab.agent.gitlab.api.AllowedAgent.configuration:type_name -> gitlab.agent.gitlab.api.Configuration
	5,  // 3: gitlab.agent.gitlab.api.Project.groups:type_name -> gitlab.agent.gitlab.api.Group
	1,  // 4: gitlab.agent.gitlab.api.AllowedAgentsForJob.allowed_agents:type_name -> gitlab.agent.gitlab.api.AllowedAgent
	6,  // 5: gitlab.agent.gitlab.api.AllowedAgentsForJob.job:type_name -> gitlab.agent.gitlab.api.Job
	3,  // 6: gitlab.agent.gitlab.api.AllowedAgentsForJob.pipeline:type_name -> gitlab.agent.gitlab.api.Pipeline
	4,  // 7: gitlab.agent.gitlab.api.AllowedAgentsForJob.project:type_name -> gitlab.agent.gitlab.api.Project
	7,  // 8: gitlab.agent.gitlab.api.AllowedAgentsForJob.user:type_name -> gitlab.agent.gitlab.api.User
	8,  // 9: gitlab.agent.gitlab.api.AllowedAgentsForJob.environment:type_name -> gitlab.agent.gitlab.api.Environment
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_internal_gitlab_api_api_proto_init() }
func file_internal_gitlab_api_api_proto_init() {
	if File_internal_gitlab_api_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_gitlab_api_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Configuration); i {
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
		file_internal_gitlab_api_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllowedAgent); i {
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
		file_internal_gitlab_api_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigProject); i {
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
		file_internal_gitlab_api_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pipeline); i {
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
		file_internal_gitlab_api_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Project); i {
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
		file_internal_gitlab_api_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group); i {
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
		file_internal_gitlab_api_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Job); i {
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
		file_internal_gitlab_api_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_internal_gitlab_api_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Environment); i {
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
		file_internal_gitlab_api_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllowedAgentsForJob); i {
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
			RawDescriptor: file_internal_gitlab_api_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_gitlab_api_api_proto_goTypes,
		DependencyIndexes: file_internal_gitlab_api_api_proto_depIdxs,
		MessageInfos:      file_internal_gitlab_api_api_proto_msgTypes,
	}.Build()
	File_internal_gitlab_api_api_proto = out.File
	file_internal_gitlab_api_api_proto_rawDesc = nil
	file_internal_gitlab_api_api_proto_goTypes = nil
	file_internal_gitlab_api_api_proto_depIdxs = nil
}
