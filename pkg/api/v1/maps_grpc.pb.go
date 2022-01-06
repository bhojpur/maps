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

// MapsServiceClient is the client API for MapsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MapsServiceClient interface {
	// StartLocalGeospatial starts a Geospatial on the Bhojpur.NET Platform directly.
	// The incoming requests are expected in the following order:
	//   1. metadata
	//   2. all bytes constituting the maps/config.yaml
	//   3. all bytes constituting the Geospatial YAML that will be executed (that the config.yaml points to)
	//   4. all bytes constituting the gzipped Bhojpur.NET Platform application tar stream
	//   5. the Bhojpur.NET Platform application tar stream done marker
	StartLocalGeospatial(ctx context.Context, opts ...grpc.CallOption) (MapsService_StartLocalGeospatialClient, error)
	// StartFromPreviousGeospatial starts a new Geospatial based on a previous one.
	// If the previous Geospatial does not have the can-replay condition set this call will result in an error.
	StartFromPreviousGeospatial(ctx context.Context, in *StartFromPreviousGeospatialRequest, opts ...grpc.CallOption) (*StartGeospatialResponse, error)
	// StartGeospatialRequest starts a new Geospatial based on its specification.
	StartGeospatial(ctx context.Context, in *StartGeospatialRequest, opts ...grpc.CallOption) (*StartGeospatialResponse, error)
	// Searches for Geospatial known to this instance
	ListGeospatial(ctx context.Context, in *ListGeospatialRequest, opts ...grpc.CallOption) (*ListGeospatialResponse, error)
	// Subscribe listens to new Geospatial updates
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (MapsService_SubscribeClient, error)
	// GetGeospatial retrieves details of a single Geospatial
	GetGeospatial(ctx context.Context, in *GetGeospatialRequest, opts ...grpc.CallOption) (*GetGeospatialResponse, error)
	// Listen listens to Geospatial updates and log output of a running Geospatial
	Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (MapsService_ListenClient, error)
	// StopGeospatial stops a currently running Geospatial
	StopGeospatial(ctx context.Context, in *StopGeospatialRequest, opts ...grpc.CallOption) (*StopGeospatialResponse, error)
}

type mapsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMapsServiceClient(cc grpc.ClientConnInterface) MapsServiceClient {
	return &mapsServiceClient{cc}
}

func (c *mapsServiceClient) StartLocalGeospatial(ctx context.Context, opts ...grpc.CallOption) (MapsService_StartLocalGeospatialClient, error) {
	stream, err := c.cc.NewStream(ctx, &MapsService_ServiceDesc.Streams[0], "/v1.MapsService/StartLocalGeospatial", opts...)
	if err != nil {
		return nil, err
	}
	x := &mapsServiceStartLocalGeospatialClient{stream}
	return x, nil
}

type MapsService_StartLocalGeospatialClient interface {
	Send(*StartLocalGeospatialRequest) error
	CloseAndRecv() (*StartGeospatialResponse, error)
	grpc.ClientStream
}

type mapsServiceStartLocalGeospatialClient struct {
	grpc.ClientStream
}

func (x *mapsServiceStartLocalGeospatialClient) Send(m *StartLocalGeospatialRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *mapsServiceStartLocalGeospatialClient) CloseAndRecv() (*StartGeospatialResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StartGeospatialResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *mapsServiceClient) StartFromPreviousGeospatial(ctx context.Context, in *StartFromPreviousGeospatialRequest, opts ...grpc.CallOption) (*StartGeospatialResponse, error) {
	out := new(StartGeospatialResponse)
	err := c.cc.Invoke(ctx, "/v1.MapsService/StartFromPreviousGeospatial", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mapsServiceClient) StartGeospatial(ctx context.Context, in *StartGeospatialRequest, opts ...grpc.CallOption) (*StartGeospatialResponse, error) {
	out := new(StartGeospatialResponse)
	err := c.cc.Invoke(ctx, "/v1.MapsService/StartGeospatial", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mapsServiceClient) ListGeospatial(ctx context.Context, in *ListGeospatialRequest, opts ...grpc.CallOption) (*ListGeospatialResponse, error) {
	out := new(ListGeospatialResponse)
	err := c.cc.Invoke(ctx, "/v1.MapsService/ListGeospatial", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mapsServiceClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (MapsService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &MapsService_ServiceDesc.Streams[1], "/v1.MapsService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &mapsServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MapsService_SubscribeClient interface {
	Recv() (*SubscribeResponse, error)
	grpc.ClientStream
}

type mapsServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *mapsServiceSubscribeClient) Recv() (*SubscribeResponse, error) {
	m := new(SubscribeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *mapsServiceClient) GetGeospatial(ctx context.Context, in *GetGeospatialRequest, opts ...grpc.CallOption) (*GetGeospatialResponse, error) {
	out := new(GetGeospatialResponse)
	err := c.cc.Invoke(ctx, "/v1.MapsService/GetGeospatial", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mapsServiceClient) Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (MapsService_ListenClient, error) {
	stream, err := c.cc.NewStream(ctx, &MapsService_ServiceDesc.Streams[2], "/v1.MapsService/Listen", opts...)
	if err != nil {
		return nil, err
	}
	x := &mapsServiceListenClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MapsService_ListenClient interface {
	Recv() (*ListenResponse, error)
	grpc.ClientStream
}

type mapsServiceListenClient struct {
	grpc.ClientStream
}

func (x *mapsServiceListenClient) Recv() (*ListenResponse, error) {
	m := new(ListenResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *mapsServiceClient) StopGeospatial(ctx context.Context, in *StopGeospatialRequest, opts ...grpc.CallOption) (*StopGeospatialResponse, error) {
	out := new(StopGeospatialResponse)
	err := c.cc.Invoke(ctx, "/v1.MapsService/StopGeospatial", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MapsServiceServer is the server API for MapsService service.
// All implementations must embed UnimplementedMapsServiceServer
// for forward compatibility
type MapsServiceServer interface {
	// StartLocalGeospatial starts a Geospatial on the Bhojpur.NET Platform directly.
	// The incoming requests are expected in the following order:
	//   1. metadata
	//   2. all bytes constituting the maps/config.yaml
	//   3. all bytes constituting the Geospatial YAML that will be executed (that the config.yaml points to)
	//   4. all bytes constituting the gzipped Bhojpur.NET Platform application tar stream
	//   5. the Bhojpur.NET Platform application tar stream done marker
	StartLocalGeospatial(MapsService_StartLocalGeospatialServer) error
	// StartFromPreviousGeospatial starts a new Geospatial based on a previous one.
	// If the previous Geospatial does not have the can-replay condition set this call will result in an error.
	StartFromPreviousGeospatial(context.Context, *StartFromPreviousGeospatialRequest) (*StartGeospatialResponse, error)
	// StartGeospatialRequest starts a new Geospatial based on its specification.
	StartGeospatial(context.Context, *StartGeospatialRequest) (*StartGeospatialResponse, error)
	// Searches for Geospatial known to this instance
	ListGeospatial(context.Context, *ListGeospatialRequest) (*ListGeospatialResponse, error)
	// Subscribe listens to new Geospatial updates
	Subscribe(*SubscribeRequest, MapsService_SubscribeServer) error
	// GetGeospatial retrieves details of a single Geospatial
	GetGeospatial(context.Context, *GetGeospatialRequest) (*GetGeospatialResponse, error)
	// Listen listens to Geospatial updates and log output of a running Geospatial
	Listen(*ListenRequest, MapsService_ListenServer) error
	// StopGeospatial stops a currently running Geospatial
	StopGeospatial(context.Context, *StopGeospatialRequest) (*StopGeospatialResponse, error)
	mustEmbedUnimplementedMapsServiceServer()
}

// UnimplementedMapsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMapsServiceServer struct {
}

func (UnimplementedMapsServiceServer) StartLocalGeospatial(MapsService_StartLocalGeospatialServer) error {
	return status.Errorf(codes.Unimplemented, "method StartLocalGeospatial not implemented")
}
func (UnimplementedMapsServiceServer) StartFromPreviousGeospatial(context.Context, *StartFromPreviousGeospatialRequest) (*StartGeospatialResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartFromPreviousGeospatial not implemented")
}
func (UnimplementedMapsServiceServer) StartGeospatial(context.Context, *StartGeospatialRequest) (*StartGeospatialResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartGeospatial not implemented")
}
func (UnimplementedMapsServiceServer) ListGeospatial(context.Context, *ListGeospatialRequest) (*ListGeospatialResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGeospatial not implemented")
}
func (UnimplementedMapsServiceServer) Subscribe(*SubscribeRequest, MapsService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedMapsServiceServer) GetGeospatial(context.Context, *GetGeospatialRequest) (*GetGeospatialResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGeospatial not implemented")
}
func (UnimplementedMapsServiceServer) Listen(*ListenRequest, MapsService_ListenServer) error {
	return status.Errorf(codes.Unimplemented, "method Listen not implemented")
}
func (UnimplementedMapsServiceServer) StopGeospatial(context.Context, *StopGeospatialRequest) (*StopGeospatialResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopGeospatial not implemented")
}
func (UnimplementedMapsServiceServer) mustEmbedUnimplementedMapsServiceServer() {}

// UnsafeMapsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MapsServiceServer will
// result in compilation errors.
type UnsafeMapsServiceServer interface {
	mustEmbedUnimplementedMapsServiceServer()
}

func RegisterMapsServiceServer(s grpc.ServiceRegistrar, srv MapsServiceServer) {
	s.RegisterService(&MapsService_ServiceDesc, srv)
}

func _MapsService_StartLocalGeospatial_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MapsServiceServer).StartLocalGeospatial(&mapsServiceStartLocalGeospatialServer{stream})
}

type MapsService_StartLocalGeospatialServer interface {
	SendAndClose(*StartGeospatialResponse) error
	Recv() (*StartLocalGeospatialRequest, error)
	grpc.ServerStream
}

type mapsServiceStartLocalGeospatialServer struct {
	grpc.ServerStream
}

func (x *mapsServiceStartLocalGeospatialServer) SendAndClose(m *StartGeospatialResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *mapsServiceStartLocalGeospatialServer) Recv() (*StartLocalGeospatialRequest, error) {
	m := new(StartLocalGeospatialRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _MapsService_StartFromPreviousGeospatial_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartFromPreviousGeospatialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MapsServiceServer).StartFromPreviousGeospatial(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.MapsService/StartFromPreviousGeospatial",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MapsServiceServer).StartFromPreviousGeospatial(ctx, req.(*StartFromPreviousGeospatialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MapsService_StartGeospatial_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartGeospatialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MapsServiceServer).StartGeospatial(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.MapsService/StartGeospatial",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MapsServiceServer).StartGeospatial(ctx, req.(*StartGeospatialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MapsService_ListGeospatial_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGeospatialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MapsServiceServer).ListGeospatial(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.MapsService/ListGeospatial",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MapsServiceServer).ListGeospatial(ctx, req.(*ListGeospatialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MapsService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MapsServiceServer).Subscribe(m, &mapsServiceSubscribeServer{stream})
}

type MapsService_SubscribeServer interface {
	Send(*SubscribeResponse) error
	grpc.ServerStream
}

type mapsServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *mapsServiceSubscribeServer) Send(m *SubscribeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _MapsService_GetGeospatial_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGeospatialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MapsServiceServer).GetGeospatial(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.MapsService/GetGeospatial",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MapsServiceServer).GetGeospatial(ctx, req.(*GetGeospatialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MapsService_Listen_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListenRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MapsServiceServer).Listen(m, &mapsServiceListenServer{stream})
}

type MapsService_ListenServer interface {
	Send(*ListenResponse) error
	grpc.ServerStream
}

type mapsServiceListenServer struct {
	grpc.ServerStream
}

func (x *mapsServiceListenServer) Send(m *ListenResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _MapsService_StopGeospatial_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopGeospatialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MapsServiceServer).StopGeospatial(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.MapsService/StopGeospatial",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MapsServiceServer).StopGeospatial(ctx, req.(*StopGeospatialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MapsService_ServiceDesc is the grpc.ServiceDesc for MapsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MapsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.MapsService",
	HandlerType: (*MapsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartFromPreviousGeospatial",
			Handler:    _MapsService_StartFromPreviousGeospatial_Handler,
		},
		{
			MethodName: "StartGeospatial",
			Handler:    _MapsService_StartGeospatial_Handler,
		},
		{
			MethodName: "ListGeospatial",
			Handler:    _MapsService_ListGeospatial_Handler,
		},
		{
			MethodName: "GetGeospatial",
			Handler:    _MapsService_GetGeospatial_Handler,
		},
		{
			MethodName: "StopGeospatial",
			Handler:    _MapsService_StopGeospatial_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StartLocalGeospatial",
			Handler:       _MapsService_StartLocalGeospatial_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Subscribe",
			Handler:       _MapsService_Subscribe_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Listen",
			Handler:       _MapsService_Listen_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "maps.proto",
}
