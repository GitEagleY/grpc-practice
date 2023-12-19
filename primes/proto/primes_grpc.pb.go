// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: primes.proto

package proto

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

// PrimeServiceClient is the client API for PrimeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PrimeServiceClient interface {
	Calculate(ctx context.Context, in *PrimeRequest, opts ...grpc.CallOption) (PrimeService_CalculateClient, error)
}

type primeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPrimeServiceClient(cc grpc.ClientConnInterface) PrimeServiceClient {
	return &primeServiceClient{cc}
}

func (c *primeServiceClient) Calculate(ctx context.Context, in *PrimeRequest, opts ...grpc.CallOption) (PrimeService_CalculateClient, error) {
	stream, err := c.cc.NewStream(ctx, &PrimeService_ServiceDesc.Streams[0], "/primes.PrimeService/Calculate", opts...)
	if err != nil {
		return nil, err
	}
	x := &primeServiceCalculateClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PrimeService_CalculateClient interface {
	Recv() (*PrimeResponce, error)
	grpc.ClientStream
}

type primeServiceCalculateClient struct {
	grpc.ClientStream
}

func (x *primeServiceCalculateClient) Recv() (*PrimeResponce, error) {
	m := new(PrimeResponce)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PrimeServiceServer is the server API for PrimeService service.
// All implementations must embed UnimplementedPrimeServiceServer
// for forward compatibility
type PrimeServiceServer interface {
	Calculate(*PrimeRequest, PrimeService_CalculateServer) error
	mustEmbedUnimplementedPrimeServiceServer()
}

// UnimplementedPrimeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPrimeServiceServer struct {
}

func (UnimplementedPrimeServiceServer) Calculate(*PrimeRequest, PrimeService_CalculateServer) error {
	return status.Errorf(codes.Unimplemented, "method Calculate not implemented")
}
func (UnimplementedPrimeServiceServer) mustEmbedUnimplementedPrimeServiceServer() {}

// UnsafePrimeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PrimeServiceServer will
// result in compilation errors.
type UnsafePrimeServiceServer interface {
	mustEmbedUnimplementedPrimeServiceServer()
}

func RegisterPrimeServiceServer(s grpc.ServiceRegistrar, srv PrimeServiceServer) {
	s.RegisterService(&PrimeService_ServiceDesc, srv)
}

func _PrimeService_Calculate_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PrimeServiceServer).Calculate(m, &primeServiceCalculateServer{stream})
}

type PrimeService_CalculateServer interface {
	Send(*PrimeResponce) error
	grpc.ServerStream
}

type primeServiceCalculateServer struct {
	grpc.ServerStream
}

func (x *primeServiceCalculateServer) Send(m *PrimeResponce) error {
	return x.ServerStream.SendMsg(m)
}

// PrimeService_ServiceDesc is the grpc.ServiceDesc for PrimeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PrimeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "primes.PrimeService",
	HandlerType: (*PrimeServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Calculate",
			Handler:       _PrimeService_Calculate_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "primes.proto",
}