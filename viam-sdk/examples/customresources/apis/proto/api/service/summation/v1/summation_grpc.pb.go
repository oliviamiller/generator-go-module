// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: api/service/summation/v1/summation.proto

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

// SummationServiceClient is the client API for SummationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SummationServiceClient interface {
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
}

type summationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSummationServiceClient(cc grpc.ClientConnInterface) SummationServiceClient {
	return &summationServiceClient{cc}
}

func (c *summationServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/acme.service.summation.v1.SummationService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SummationServiceServer is the server API for SummationService service.
// All implementations must embed UnimplementedSummationServiceServer
// for forward compatibility
type SummationServiceServer interface {
	Sum(context.Context, *SumRequest) (*SumResponse, error)
	mustEmbedUnimplementedSummationServiceServer()
}

// UnimplementedSummationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSummationServiceServer struct {
}

func (UnimplementedSummationServiceServer) Sum(context.Context, *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (UnimplementedSummationServiceServer) mustEmbedUnimplementedSummationServiceServer() {}

// UnsafeSummationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SummationServiceServer will
// result in compilation errors.
type UnsafeSummationServiceServer interface {
	mustEmbedUnimplementedSummationServiceServer()
}

func RegisterSummationServiceServer(s grpc.ServiceRegistrar, srv SummationServiceServer) {
	s.RegisterService(&SummationService_ServiceDesc, srv)
}

func _SummationService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SummationServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/acme.service.summation.v1.SummationService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SummationServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SummationService_ServiceDesc is the grpc.ServiceDesc for SummationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SummationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "acme.service.summation.v1.SummationService",
	HandlerType: (*SummationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _SummationService_Sum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/service/summation/v1/summation.proto",
}
