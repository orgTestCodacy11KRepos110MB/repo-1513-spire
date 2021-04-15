// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package nodeattestorv0

import (
	context "context"
	plugin "github.com/spiffe/spire/proto/spire/common/plugin"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// NodeAttestorClient is the client API for NodeAttestor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeAttestorClient interface {
	//* Returns the node attestation data for specific platform and the generated Base SPIFFE ID for CSR formation
	FetchAttestationData(ctx context.Context, opts ...grpc.CallOption) (NodeAttestor_FetchAttestationDataClient, error)
	//* Applies the plugin configuration and returns configuration errors
	Configure(ctx context.Context, in *plugin.ConfigureRequest, opts ...grpc.CallOption) (*plugin.ConfigureResponse, error)
	//* Returns the version and related metadata of the plugin
	GetPluginInfo(ctx context.Context, in *plugin.GetPluginInfoRequest, opts ...grpc.CallOption) (*plugin.GetPluginInfoResponse, error)
}

type nodeAttestorClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeAttestorClient(cc grpc.ClientConnInterface) NodeAttestorClient {
	return &nodeAttestorClient{cc}
}

func (c *nodeAttestorClient) FetchAttestationData(ctx context.Context, opts ...grpc.CallOption) (NodeAttestor_FetchAttestationDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &NodeAttestor_ServiceDesc.Streams[0], "/spire.agent.nodeattestor.NodeAttestor/FetchAttestationData", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeAttestorFetchAttestationDataClient{stream}
	return x, nil
}

type NodeAttestor_FetchAttestationDataClient interface {
	Send(*FetchAttestationDataRequest) error
	Recv() (*FetchAttestationDataResponse, error)
	grpc.ClientStream
}

type nodeAttestorFetchAttestationDataClient struct {
	grpc.ClientStream
}

func (x *nodeAttestorFetchAttestationDataClient) Send(m *FetchAttestationDataRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *nodeAttestorFetchAttestationDataClient) Recv() (*FetchAttestationDataResponse, error) {
	m := new(FetchAttestationDataResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeAttestorClient) Configure(ctx context.Context, in *plugin.ConfigureRequest, opts ...grpc.CallOption) (*plugin.ConfigureResponse, error) {
	out := new(plugin.ConfigureResponse)
	err := c.cc.Invoke(ctx, "/spire.agent.nodeattestor.NodeAttestor/Configure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeAttestorClient) GetPluginInfo(ctx context.Context, in *plugin.GetPluginInfoRequest, opts ...grpc.CallOption) (*plugin.GetPluginInfoResponse, error) {
	out := new(plugin.GetPluginInfoResponse)
	err := c.cc.Invoke(ctx, "/spire.agent.nodeattestor.NodeAttestor/GetPluginInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeAttestorServer is the server API for NodeAttestor service.
// All implementations must embed UnimplementedNodeAttestorServer
// for forward compatibility
type NodeAttestorServer interface {
	//* Returns the node attestation data for specific platform and the generated Base SPIFFE ID for CSR formation
	FetchAttestationData(NodeAttestor_FetchAttestationDataServer) error
	//* Applies the plugin configuration and returns configuration errors
	Configure(context.Context, *plugin.ConfigureRequest) (*plugin.ConfigureResponse, error)
	//* Returns the version and related metadata of the plugin
	GetPluginInfo(context.Context, *plugin.GetPluginInfoRequest) (*plugin.GetPluginInfoResponse, error)
	mustEmbedUnimplementedNodeAttestorServer()
}

// UnimplementedNodeAttestorServer must be embedded to have forward compatible implementations.
type UnimplementedNodeAttestorServer struct {
}

func (UnimplementedNodeAttestorServer) FetchAttestationData(NodeAttestor_FetchAttestationDataServer) error {
	return status.Errorf(codes.Unimplemented, "method FetchAttestationData not implemented")
}
func (UnimplementedNodeAttestorServer) Configure(context.Context, *plugin.ConfigureRequest) (*plugin.ConfigureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Configure not implemented")
}
func (UnimplementedNodeAttestorServer) GetPluginInfo(context.Context, *plugin.GetPluginInfoRequest) (*plugin.GetPluginInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPluginInfo not implemented")
}
func (UnimplementedNodeAttestorServer) mustEmbedUnimplementedNodeAttestorServer() {}

// UnsafeNodeAttestorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeAttestorServer will
// result in compilation errors.
type UnsafeNodeAttestorServer interface {
	mustEmbedUnimplementedNodeAttestorServer()
}

func RegisterNodeAttestorServer(s grpc.ServiceRegistrar, srv NodeAttestorServer) {
	s.RegisterService(&NodeAttestor_ServiceDesc, srv)
}

func _NodeAttestor_FetchAttestationData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NodeAttestorServer).FetchAttestationData(&nodeAttestorFetchAttestationDataServer{stream})
}

type NodeAttestor_FetchAttestationDataServer interface {
	Send(*FetchAttestationDataResponse) error
	Recv() (*FetchAttestationDataRequest, error)
	grpc.ServerStream
}

type nodeAttestorFetchAttestationDataServer struct {
	grpc.ServerStream
}

func (x *nodeAttestorFetchAttestationDataServer) Send(m *FetchAttestationDataResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *nodeAttestorFetchAttestationDataServer) Recv() (*FetchAttestationDataRequest, error) {
	m := new(FetchAttestationDataRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _NodeAttestor_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(plugin.ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeAttestorServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.agent.nodeattestor.NodeAttestor/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeAttestorServer).Configure(ctx, req.(*plugin.ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeAttestor_GetPluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(plugin.GetPluginInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeAttestorServer).GetPluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.agent.nodeattestor.NodeAttestor/GetPluginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeAttestorServer).GetPluginInfo(ctx, req.(*plugin.GetPluginInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NodeAttestor_ServiceDesc is the grpc.ServiceDesc for NodeAttestor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodeAttestor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spire.agent.nodeattestor.NodeAttestor",
	HandlerType: (*NodeAttestorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Configure",
			Handler:    _NodeAttestor_Configure_Handler,
		},
		{
			MethodName: "GetPluginInfo",
			Handler:    _NodeAttestor_GetPluginInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FetchAttestationData",
			Handler:       _NodeAttestor_FetchAttestationData_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "spire/plugin/agent/nodeattestor/v0/nodeattestor.proto",
}