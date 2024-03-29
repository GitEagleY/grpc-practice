// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: calculator.proto

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

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	Calculate(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorResponce, error)
	Avg(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_AvgClient, error)
	Max(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_MaxClient, error)
	Sqrt(ctx context.Context, in *SqrtRequest, opts ...grpc.CallOption) (*SqrtResponce, error)
}

type calculatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorServiceClient(cc grpc.ClientConnInterface) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Calculate(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorResponce, error) {
	out := new(CalculatorResponce)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Calculate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) Avg(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_AvgClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[0], "/calculator.CalculatorService/Avg", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceAvgClient{stream}
	return x, nil
}

type CalculatorService_AvgClient interface {
	Send(*AvgRequest) error
	CloseAndRecv() (*AvgResponce, error)
	grpc.ClientStream
}

type calculatorServiceAvgClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceAvgClient) Send(m *AvgRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceAvgClient) CloseAndRecv() (*AvgResponce, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AvgResponce)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) Max(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_MaxClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[1], "/calculator.CalculatorService/Max", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceMaxClient{stream}
	return x, nil
}

type CalculatorService_MaxClient interface {
	Send(*MaxRequest) error
	Recv() (*MaxResponce, error)
	grpc.ClientStream
}

type calculatorServiceMaxClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceMaxClient) Send(m *MaxRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceMaxClient) Recv() (*MaxResponce, error) {
	m := new(MaxResponce)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) Sqrt(ctx context.Context, in *SqrtRequest, opts ...grpc.CallOption) (*SqrtResponce, error) {
	out := new(SqrtResponce)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Sqrt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
// All implementations must embed UnimplementedCalculatorServiceServer
// for forward compatibility
type CalculatorServiceServer interface {
	Calculate(context.Context, *CalculatorRequest) (*CalculatorResponce, error)
	Avg(CalculatorService_AvgServer) error
	Max(CalculatorService_MaxServer) error
	Sqrt(context.Context, *SqrtRequest) (*SqrtResponce, error)
	mustEmbedUnimplementedCalculatorServiceServer()
}

// UnimplementedCalculatorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (UnimplementedCalculatorServiceServer) Calculate(context.Context, *CalculatorRequest) (*CalculatorResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Calculate not implemented")
}
func (UnimplementedCalculatorServiceServer) Avg(CalculatorService_AvgServer) error {
	return status.Errorf(codes.Unimplemented, "method Avg not implemented")
}
func (UnimplementedCalculatorServiceServer) Max(CalculatorService_MaxServer) error {
	return status.Errorf(codes.Unimplemented, "method Max not implemented")
}
func (UnimplementedCalculatorServiceServer) Sqrt(context.Context, *SqrtRequest) (*SqrtResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sqrt not implemented")
}
func (UnimplementedCalculatorServiceServer) mustEmbedUnimplementedCalculatorServiceServer() {}

// UnsafeCalculatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculatorServiceServer will
// result in compilation errors.
type UnsafeCalculatorServiceServer interface {
	mustEmbedUnimplementedCalculatorServiceServer()
}

func RegisterCalculatorServiceServer(s grpc.ServiceRegistrar, srv CalculatorServiceServer) {
	s.RegisterService(&CalculatorService_ServiceDesc, srv)
}

func _CalculatorService_Calculate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Calculate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Calculate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Calculate(ctx, req.(*CalculatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_Avg_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).Avg(&calculatorServiceAvgServer{stream})
}

type CalculatorService_AvgServer interface {
	SendAndClose(*AvgResponce) error
	Recv() (*AvgRequest, error)
	grpc.ServerStream
}

type calculatorServiceAvgServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceAvgServer) SendAndClose(m *AvgResponce) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceAvgServer) Recv() (*AvgRequest, error) {
	m := new(AvgRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculatorService_Max_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).Max(&calculatorServiceMaxServer{stream})
}

type CalculatorService_MaxServer interface {
	Send(*MaxResponce) error
	Recv() (*MaxRequest, error)
	grpc.ServerStream
}

type calculatorServiceMaxServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceMaxServer) Send(m *MaxResponce) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceMaxServer) Recv() (*MaxRequest, error) {
	m := new(MaxRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculatorService_Sqrt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SqrtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Sqrt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Sqrt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Sqrt(ctx, req.(*SqrtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CalculatorService_ServiceDesc is the grpc.ServiceDesc for CalculatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalculatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Calculate",
			Handler:    _CalculatorService_Calculate_Handler,
		},
		{
			MethodName: "Sqrt",
			Handler:    _CalculatorService_Sqrt_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Avg",
			Handler:       _CalculatorService_Avg_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Max",
			Handler:       _CalculatorService_Max_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "calculator.proto",
}
