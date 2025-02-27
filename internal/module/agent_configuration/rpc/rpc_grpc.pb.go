// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.7
// source: internal/module/agent_configuration/rpc/rpc.proto

package rpc

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AgentConfiguration_GetConfiguration_FullMethodName = "/gitlab.agent.agent_configuration.rpc.AgentConfiguration/GetConfiguration"
)

// AgentConfigurationClient is the client API for AgentConfiguration service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentConfigurationClient interface {
	GetConfiguration(ctx context.Context, in *ConfigurationRequest, opts ...grpc.CallOption) (AgentConfiguration_GetConfigurationClient, error)
}

type agentConfigurationClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentConfigurationClient(cc grpc.ClientConnInterface) AgentConfigurationClient {
	return &agentConfigurationClient{cc}
}

func (c *agentConfigurationClient) GetConfiguration(ctx context.Context, in *ConfigurationRequest, opts ...grpc.CallOption) (AgentConfiguration_GetConfigurationClient, error) {
	stream, err := c.cc.NewStream(ctx, &AgentConfiguration_ServiceDesc.Streams[0], AgentConfiguration_GetConfiguration_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &agentConfigurationGetConfigurationClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AgentConfiguration_GetConfigurationClient interface {
	Recv() (*ConfigurationResponse, error)
	grpc.ClientStream
}

type agentConfigurationGetConfigurationClient struct {
	grpc.ClientStream
}

func (x *agentConfigurationGetConfigurationClient) Recv() (*ConfigurationResponse, error) {
	m := new(ConfigurationResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AgentConfigurationServer is the server API for AgentConfiguration service.
// All implementations must embed UnimplementedAgentConfigurationServer
// for forward compatibility
type AgentConfigurationServer interface {
	GetConfiguration(*ConfigurationRequest, AgentConfiguration_GetConfigurationServer) error
	mustEmbedUnimplementedAgentConfigurationServer()
}

// UnimplementedAgentConfigurationServer must be embedded to have forward compatible implementations.
type UnimplementedAgentConfigurationServer struct {
}

func (UnimplementedAgentConfigurationServer) GetConfiguration(*ConfigurationRequest, AgentConfiguration_GetConfigurationServer) error {
	return status.Errorf(codes.Unimplemented, "method GetConfiguration not implemented")
}
func (UnimplementedAgentConfigurationServer) mustEmbedUnimplementedAgentConfigurationServer() {}

// UnsafeAgentConfigurationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentConfigurationServer will
// result in compilation errors.
type UnsafeAgentConfigurationServer interface {
	mustEmbedUnimplementedAgentConfigurationServer()
}

func RegisterAgentConfigurationServer(s grpc.ServiceRegistrar, srv AgentConfigurationServer) {
	s.RegisterService(&AgentConfiguration_ServiceDesc, srv)
}

func _AgentConfiguration_GetConfiguration_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConfigurationRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AgentConfigurationServer).GetConfiguration(m, &agentConfigurationGetConfigurationServer{stream})
}

type AgentConfiguration_GetConfigurationServer interface {
	Send(*ConfigurationResponse) error
	grpc.ServerStream
}

type agentConfigurationGetConfigurationServer struct {
	grpc.ServerStream
}

func (x *agentConfigurationGetConfigurationServer) Send(m *ConfigurationResponse) error {
	return x.ServerStream.SendMsg(m)
}

// AgentConfiguration_ServiceDesc is the grpc.ServiceDesc for AgentConfiguration service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AgentConfiguration_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gitlab.agent.agent_configuration.rpc.AgentConfiguration",
	HandlerType: (*AgentConfigurationServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetConfiguration",
			Handler:       _AgentConfiguration_GetConfiguration_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "internal/module/agent_configuration/rpc/rpc.proto",
}
