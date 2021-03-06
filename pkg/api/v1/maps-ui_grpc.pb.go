// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// MapsUIClient is the client API for MapsUI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MapsUIClient interface {
	// ListGeospatialSpecs returns a list of Geospatial that can be started through the UI.
	ListGeospatialSpecs(ctx context.Context, in *ListGeospatialSpecsRequest, opts ...grpc.CallOption) (MapsUI_ListGeospatialSpecsClient, error)
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error)
}

type mapsUIClient struct {
	cc grpc.ClientConnInterface
}

func NewMapsUIClient(cc grpc.ClientConnInterface) MapsUIClient {
	return &mapsUIClient{cc}
}

func (c *mapsUIClient) ListGeospatialSpecs(ctx context.Context, in *ListGeospatialSpecsRequest, opts ...grpc.CallOption) (MapsUI_ListGeospatialSpecsClient, error) {
	stream, err := c.cc.NewStream(ctx, &MapsUI_ServiceDesc.Streams[0], "/v1.MapsUI/ListGeospatialSpecs", opts...)
	if err != nil {
		return nil, err
	}
	x := &mapsUIListGeospatialSpecsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MapsUI_ListGeospatialSpecsClient interface {
	Recv() (*ListGeospatialSpecsResponse, error)
	grpc.ClientStream
}

type mapsUIListGeospatialSpecsClient struct {
	grpc.ClientStream
}

func (x *mapsUIListGeospatialSpecsClient) Recv() (*ListGeospatialSpecsResponse, error) {
	m := new(ListGeospatialSpecsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *mapsUIClient) IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error) {
	out := new(IsReadOnlyResponse)
	err := c.cc.Invoke(ctx, "/v1.MapsUI/IsReadOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MapsUIServer is the server API for MapsUI service.
// All implementations must embed UnimplementedMapsUIServer
// for forward compatibility
type MapsUIServer interface {
	// ListGeospatialSpecs returns a list of Geospatial that can be started through the UI.
	ListGeospatialSpecs(*ListGeospatialSpecsRequest, MapsUI_ListGeospatialSpecsServer) error
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error)
	mustEmbedUnimplementedMapsUIServer()
}

// UnimplementedMapsUIServer must be embedded to have forward compatible implementations.
type UnimplementedMapsUIServer struct {
}

func (UnimplementedMapsUIServer) ListGeospatialSpecs(*ListGeospatialSpecsRequest, MapsUI_ListGeospatialSpecsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListGeospatialSpecs not implemented")
}
func (UnimplementedMapsUIServer) IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsReadOnly not implemented")
}
func (UnimplementedMapsUIServer) mustEmbedUnimplementedMapsUIServer() {}

// UnsafeMapsUIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MapsUIServer will
// result in compilation errors.
type UnsafeMapsUIServer interface {
	mustEmbedUnimplementedMapsUIServer()
}

func RegisterMapsUIServer(s grpc.ServiceRegistrar, srv MapsUIServer) {
	s.RegisterService(&MapsUI_ServiceDesc, srv)
}

func _MapsUI_ListGeospatialSpecs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListGeospatialSpecsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MapsUIServer).ListGeospatialSpecs(m, &mapsUIListGeospatialSpecsServer{stream})
}

type MapsUI_ListGeospatialSpecsServer interface {
	Send(*ListGeospatialSpecsResponse) error
	grpc.ServerStream
}

type mapsUIListGeospatialSpecsServer struct {
	grpc.ServerStream
}

func (x *mapsUIListGeospatialSpecsServer) Send(m *ListGeospatialSpecsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _MapsUI_IsReadOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsReadOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MapsUIServer).IsReadOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.MapsUI/IsReadOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MapsUIServer).IsReadOnly(ctx, req.(*IsReadOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MapsUI_ServiceDesc is the grpc.ServiceDesc for MapsUI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MapsUI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.MapsUI",
	HandlerType: (*MapsUIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsReadOnly",
			Handler:    _MapsUI_IsReadOnly_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListGeospatialSpecs",
			Handler:       _MapsUI_ListGeospatialSpecs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "maps-ui.proto",
}
