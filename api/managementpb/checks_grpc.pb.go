// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: managementpb/checks.proto

package managementpb

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

// SecurityChecksClient is the client API for SecurityChecks service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SecurityChecksClient interface {
	// ListFailedServices returns a list of services with failed checks.
	ListFailedServices(ctx context.Context, in *ListFailedServicesRequest, opts ...grpc.CallOption) (*ListFailedServicesResponse, error)
	// GetFailedChecks returns the checks result for a given service.
	GetFailedChecks(ctx context.Context, in *GetFailedChecksRequest, opts ...grpc.CallOption) (*GetFailedChecksResponse, error)
	// ToggleCheckAlert allows to switch alerts state for a check result between "silenced" and "unsilenced".
	ToggleCheckAlert(ctx context.Context, in *ToggleCheckAlertRequest, opts ...grpc.CallOption) (*ToggleCheckAlertResponse, error)
	// Deprecated: Do not use.
	// GetSecurityCheckResults returns Security Thread Tool's latest checks results.
	GetSecurityCheckResults(ctx context.Context, in *GetSecurityCheckResultsRequest, opts ...grpc.CallOption) (*GetSecurityCheckResultsResponse, error)
	// StartSecurityChecks executes Security Thread Tool checks and returns when all checks are executed.
	StartSecurityChecks(ctx context.Context, in *StartSecurityChecksRequest, opts ...grpc.CallOption) (*StartSecurityChecksResponse, error)
	// StartChecksStream executes Checks and returns their results as a gRPC stream.
	// Current implementation expects to handle a websocket connection, and the initial websocket request must
	// have the Cookie header set with the proper `grafana_session` value (otherwise it fails with a 401 error).
	StartChecksStream(ctx context.Context, opts ...grpc.CallOption) (SecurityChecks_StartChecksStreamClient, error)
	// ListSecurityChecks returns a list of available Security Thread Tool checks.
	ListSecurityChecks(ctx context.Context, in *ListSecurityChecksRequest, opts ...grpc.CallOption) (*ListSecurityChecksResponse, error)
	// ChangeSecurityChecks enables/disables Security Thread Tool checks or changes their interval by names.
	ChangeSecurityChecks(ctx context.Context, in *ChangeSecurityChecksRequest, opts ...grpc.CallOption) (*ChangeSecurityChecksResponse, error)
}

type securityChecksClient struct {
	cc grpc.ClientConnInterface
}

func NewSecurityChecksClient(cc grpc.ClientConnInterface) SecurityChecksClient {
	return &securityChecksClient{cc}
}

func (c *securityChecksClient) ListFailedServices(ctx context.Context, in *ListFailedServicesRequest, opts ...grpc.CallOption) (*ListFailedServicesResponse, error) {
	out := new(ListFailedServicesResponse)
	err := c.cc.Invoke(ctx, "/management.SecurityChecks/ListFailedServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityChecksClient) GetFailedChecks(ctx context.Context, in *GetFailedChecksRequest, opts ...grpc.CallOption) (*GetFailedChecksResponse, error) {
	out := new(GetFailedChecksResponse)
	err := c.cc.Invoke(ctx, "/management.SecurityChecks/GetFailedChecks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityChecksClient) ToggleCheckAlert(ctx context.Context, in *ToggleCheckAlertRequest, opts ...grpc.CallOption) (*ToggleCheckAlertResponse, error) {
	out := new(ToggleCheckAlertResponse)
	err := c.cc.Invoke(ctx, "/management.SecurityChecks/ToggleCheckAlert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *securityChecksClient) GetSecurityCheckResults(ctx context.Context, in *GetSecurityCheckResultsRequest, opts ...grpc.CallOption) (*GetSecurityCheckResultsResponse, error) {
	out := new(GetSecurityCheckResultsResponse)
	err := c.cc.Invoke(ctx, "/management.SecurityChecks/GetSecurityCheckResults", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityChecksClient) StartSecurityChecks(ctx context.Context, in *StartSecurityChecksRequest, opts ...grpc.CallOption) (*StartSecurityChecksResponse, error) {
	out := new(StartSecurityChecksResponse)
	err := c.cc.Invoke(ctx, "/management.SecurityChecks/StartSecurityChecks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityChecksClient) StartChecksStream(ctx context.Context, opts ...grpc.CallOption) (SecurityChecks_StartChecksStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &SecurityChecks_ServiceDesc.Streams[0], "/management.SecurityChecks/StartChecksStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &securityChecksStartChecksStreamClient{stream}
	return x, nil
}

type SecurityChecks_StartChecksStreamClient interface {
	Send(*StartSecurityChecksRequest) error
	Recv() (*StartChecksStreamResponse, error)
	grpc.ClientStream
}

type securityChecksStartChecksStreamClient struct {
	grpc.ClientStream
}

func (x *securityChecksStartChecksStreamClient) Send(m *StartSecurityChecksRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *securityChecksStartChecksStreamClient) Recv() (*StartChecksStreamResponse, error) {
	m := new(StartChecksStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *securityChecksClient) ListSecurityChecks(ctx context.Context, in *ListSecurityChecksRequest, opts ...grpc.CallOption) (*ListSecurityChecksResponse, error) {
	out := new(ListSecurityChecksResponse)
	err := c.cc.Invoke(ctx, "/management.SecurityChecks/ListSecurityChecks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityChecksClient) ChangeSecurityChecks(ctx context.Context, in *ChangeSecurityChecksRequest, opts ...grpc.CallOption) (*ChangeSecurityChecksResponse, error) {
	out := new(ChangeSecurityChecksResponse)
	err := c.cc.Invoke(ctx, "/management.SecurityChecks/ChangeSecurityChecks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecurityChecksServer is the server API for SecurityChecks service.
// All implementations must embed UnimplementedSecurityChecksServer
// for forward compatibility
type SecurityChecksServer interface {
	// ListFailedServices returns a list of services with failed checks.
	ListFailedServices(context.Context, *ListFailedServicesRequest) (*ListFailedServicesResponse, error)
	// GetFailedChecks returns the checks result for a given service.
	GetFailedChecks(context.Context, *GetFailedChecksRequest) (*GetFailedChecksResponse, error)
	// ToggleCheckAlert allows to switch alerts state for a check result between "silenced" and "unsilenced".
	ToggleCheckAlert(context.Context, *ToggleCheckAlertRequest) (*ToggleCheckAlertResponse, error)
	// Deprecated: Do not use.
	// GetSecurityCheckResults returns Security Thread Tool's latest checks results.
	GetSecurityCheckResults(context.Context, *GetSecurityCheckResultsRequest) (*GetSecurityCheckResultsResponse, error)
	// StartSecurityChecks executes Security Thread Tool checks and returns when all checks are executed.
	StartSecurityChecks(context.Context, *StartSecurityChecksRequest) (*StartSecurityChecksResponse, error)
	// StartChecksStream executes Checks and returns their results as a gRPC stream.
	// Current implementation expects to handle a websocket connection, and the initial websocket request must
	// have the Cookie header set with the proper `grafana_session` value (otherwise it fails with a 401 error).
	StartChecksStream(SecurityChecks_StartChecksStreamServer) error
	// ListSecurityChecks returns a list of available Security Thread Tool checks.
	ListSecurityChecks(context.Context, *ListSecurityChecksRequest) (*ListSecurityChecksResponse, error)
	// ChangeSecurityChecks enables/disables Security Thread Tool checks or changes their interval by names.
	ChangeSecurityChecks(context.Context, *ChangeSecurityChecksRequest) (*ChangeSecurityChecksResponse, error)
	mustEmbedUnimplementedSecurityChecksServer()
}

// UnimplementedSecurityChecksServer must be embedded to have forward compatible implementations.
type UnimplementedSecurityChecksServer struct{}

func (UnimplementedSecurityChecksServer) ListFailedServices(context.Context, *ListFailedServicesRequest) (*ListFailedServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFailedServices not implemented")
}

func (UnimplementedSecurityChecksServer) GetFailedChecks(context.Context, *GetFailedChecksRequest) (*GetFailedChecksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFailedChecks not implemented")
}

func (UnimplementedSecurityChecksServer) ToggleCheckAlert(context.Context, *ToggleCheckAlertRequest) (*ToggleCheckAlertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ToggleCheckAlert not implemented")
}

func (UnimplementedSecurityChecksServer) GetSecurityCheckResults(context.Context, *GetSecurityCheckResultsRequest) (*GetSecurityCheckResultsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSecurityCheckResults not implemented")
}

func (UnimplementedSecurityChecksServer) StartSecurityChecks(context.Context, *StartSecurityChecksRequest) (*StartSecurityChecksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartSecurityChecks not implemented")
}

func (UnimplementedSecurityChecksServer) StartChecksStream(SecurityChecks_StartChecksStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method StartChecksStream not implemented")
}

func (UnimplementedSecurityChecksServer) ListSecurityChecks(context.Context, *ListSecurityChecksRequest) (*ListSecurityChecksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSecurityChecks not implemented")
}

func (UnimplementedSecurityChecksServer) ChangeSecurityChecks(context.Context, *ChangeSecurityChecksRequest) (*ChangeSecurityChecksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeSecurityChecks not implemented")
}
func (UnimplementedSecurityChecksServer) mustEmbedUnimplementedSecurityChecksServer() {}

// UnsafeSecurityChecksServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SecurityChecksServer will
// result in compilation errors.
type UnsafeSecurityChecksServer interface {
	mustEmbedUnimplementedSecurityChecksServer()
}

func RegisterSecurityChecksServer(s grpc.ServiceRegistrar, srv SecurityChecksServer) {
	s.RegisterService(&SecurityChecks_ServiceDesc, srv)
}

func _SecurityChecks_ListFailedServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFailedServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityChecksServer).ListFailedServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.SecurityChecks/ListFailedServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityChecksServer).ListFailedServices(ctx, req.(*ListFailedServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityChecks_GetFailedChecks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFailedChecksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityChecksServer).GetFailedChecks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.SecurityChecks/GetFailedChecks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityChecksServer).GetFailedChecks(ctx, req.(*GetFailedChecksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityChecks_ToggleCheckAlert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ToggleCheckAlertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityChecksServer).ToggleCheckAlert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.SecurityChecks/ToggleCheckAlert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityChecksServer).ToggleCheckAlert(ctx, req.(*ToggleCheckAlertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityChecks_GetSecurityCheckResults_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSecurityCheckResultsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityChecksServer).GetSecurityCheckResults(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.SecurityChecks/GetSecurityCheckResults",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityChecksServer).GetSecurityCheckResults(ctx, req.(*GetSecurityCheckResultsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityChecks_StartSecurityChecks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartSecurityChecksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityChecksServer).StartSecurityChecks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.SecurityChecks/StartSecurityChecks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityChecksServer).StartSecurityChecks(ctx, req.(*StartSecurityChecksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityChecks_StartChecksStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SecurityChecksServer).StartChecksStream(&securityChecksStartChecksStreamServer{stream})
}

type SecurityChecks_StartChecksStreamServer interface {
	Send(*StartChecksStreamResponse) error
	Recv() (*StartSecurityChecksRequest, error)
	grpc.ServerStream
}

type securityChecksStartChecksStreamServer struct {
	grpc.ServerStream
}

func (x *securityChecksStartChecksStreamServer) Send(m *StartChecksStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *securityChecksStartChecksStreamServer) Recv() (*StartSecurityChecksRequest, error) {
	m := new(StartSecurityChecksRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SecurityChecks_ListSecurityChecks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSecurityChecksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityChecksServer).ListSecurityChecks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.SecurityChecks/ListSecurityChecks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityChecksServer).ListSecurityChecks(ctx, req.(*ListSecurityChecksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityChecks_ChangeSecurityChecks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeSecurityChecksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityChecksServer).ChangeSecurityChecks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.SecurityChecks/ChangeSecurityChecks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityChecksServer).ChangeSecurityChecks(ctx, req.(*ChangeSecurityChecksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SecurityChecks_ServiceDesc is the grpc.ServiceDesc for SecurityChecks service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SecurityChecks_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "management.SecurityChecks",
	HandlerType: (*SecurityChecksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListFailedServices",
			Handler:    _SecurityChecks_ListFailedServices_Handler,
		},
		{
			MethodName: "GetFailedChecks",
			Handler:    _SecurityChecks_GetFailedChecks_Handler,
		},
		{
			MethodName: "ToggleCheckAlert",
			Handler:    _SecurityChecks_ToggleCheckAlert_Handler,
		},
		{
			MethodName: "GetSecurityCheckResults",
			Handler:    _SecurityChecks_GetSecurityCheckResults_Handler,
		},
		{
			MethodName: "StartSecurityChecks",
			Handler:    _SecurityChecks_StartSecurityChecks_Handler,
		},
		{
			MethodName: "ListSecurityChecks",
			Handler:    _SecurityChecks_ListSecurityChecks_Handler,
		},
		{
			MethodName: "ChangeSecurityChecks",
			Handler:    _SecurityChecks_ChangeSecurityChecks_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StartChecksStream",
			Handler:       _SecurityChecks_StartChecksStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "managementpb/checks.proto",
}
