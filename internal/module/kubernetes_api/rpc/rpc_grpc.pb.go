// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.7
// source: internal/module/kubernetes_api/rpc/rpc.proto

package rpc

import (
	context "context"

	grpctool "gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/tool/grpctool"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	KubernetesApi_MakeRequest_FullMethodName = "/gitlab.agent.kubernetes_api.rpc.KubernetesApi/MakeRequest"
)

// KubernetesApiClient is the client API for KubernetesApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KubernetesApiClient interface {
	MakeRequest(ctx context.Context, opts ...grpc.CallOption) (KubernetesApi_MakeRequestClient, error)
}

type kubernetesApiClient struct {
	cc grpc.ClientConnInterface
}

func NewKubernetesApiClient(cc grpc.ClientConnInterface) KubernetesApiClient {
	return &kubernetesApiClient{cc}
}

func (c *kubernetesApiClient) MakeRequest(ctx context.Context, opts ...grpc.CallOption) (KubernetesApi_MakeRequestClient, error) {
	stream, err := c.cc.NewStream(ctx, &KubernetesApi_ServiceDesc.Streams[0], KubernetesApi_MakeRequest_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &kubernetesApiMakeRequestClient{stream}
	return x, nil
}

type KubernetesApi_MakeRequestClient interface {
	Send(*grpctool.HttpRequest) error
	Recv() (*grpctool.HttpResponse, error)
	grpc.ClientStream
}

type kubernetesApiMakeRequestClient struct {
	grpc.ClientStream
}

func (x *kubernetesApiMakeRequestClient) Send(m *grpctool.HttpRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *kubernetesApiMakeRequestClient) Recv() (*grpctool.HttpResponse, error) {
	m := new(grpctool.HttpResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// KubernetesApiServer is the server API for KubernetesApi service.
// All implementations must embed UnimplementedKubernetesApiServer
// for forward compatibility
type KubernetesApiServer interface {
	MakeRequest(KubernetesApi_MakeRequestServer) error
	mustEmbedUnimplementedKubernetesApiServer()
}

// UnimplementedKubernetesApiServer must be embedded to have forward compatible implementations.
type UnimplementedKubernetesApiServer struct {
}

func (UnimplementedKubernetesApiServer) MakeRequest(KubernetesApi_MakeRequestServer) error {
	return status.Errorf(codes.Unimplemented, "method MakeRequest not implemented")
}
func (UnimplementedKubernetesApiServer) mustEmbedUnimplementedKubernetesApiServer() {}

// UnsafeKubernetesApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KubernetesApiServer will
// result in compilation errors.
type UnsafeKubernetesApiServer interface {
	mustEmbedUnimplementedKubernetesApiServer()
}

func RegisterKubernetesApiServer(s grpc.ServiceRegistrar, srv KubernetesApiServer) {
	s.RegisterService(&KubernetesApi_ServiceDesc, srv)
}

func _KubernetesApi_MakeRequest_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(KubernetesApiServer).MakeRequest(&kubernetesApiMakeRequestServer{stream})
}

type KubernetesApi_MakeRequestServer interface {
	Send(*grpctool.HttpResponse) error
	Recv() (*grpctool.HttpRequest, error)
	grpc.ServerStream
}

type kubernetesApiMakeRequestServer struct {
	grpc.ServerStream
}

func (x *kubernetesApiMakeRequestServer) Send(m *grpctool.HttpResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *kubernetesApiMakeRequestServer) Recv() (*grpctool.HttpRequest, error) {
	m := new(grpctool.HttpRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// KubernetesApi_ServiceDesc is the grpc.ServiceDesc for KubernetesApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KubernetesApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gitlab.agent.kubernetes_api.rpc.KubernetesApi",
	HandlerType: (*KubernetesApiServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MakeRequest",
			Handler:       _KubernetesApi_MakeRequest_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "internal/module/kubernetes_api/rpc/rpc.proto",
}
