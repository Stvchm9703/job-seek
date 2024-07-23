// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: prediction.proto

package protos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	PredictionService_UnaryServerUnaryClient_FullMethodName   = "/job_seek.prediction.PredictionService/UnaryServerUnaryClient"
	PredictionService_StreamServerUnaryClient_FullMethodName  = "/job_seek.prediction.PredictionService/StreamServerUnaryClient"
	PredictionService_UnaryServerStreamClient_FullMethodName  = "/job_seek.prediction.PredictionService/UnaryServerStreamClient"
	PredictionService_StreamServerStreamClient_FullMethodName = "/job_seek.prediction.PredictionService/StreamServerStreamClient"
)

// PredictionServiceClient is the client API for PredictionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PredictionServiceClient interface {
	UnaryServerUnaryClient(ctx context.Context, in *Simple, opts ...grpc.CallOption) (*Simple, error)
	StreamServerUnaryClient(ctx context.Context, in *Simple, opts ...grpc.CallOption) (PredictionService_StreamServerUnaryClientClient, error)
	UnaryServerStreamClient(ctx context.Context, opts ...grpc.CallOption) (PredictionService_UnaryServerStreamClientClient, error)
	StreamServerStreamClient(ctx context.Context, opts ...grpc.CallOption) (PredictionService_StreamServerStreamClientClient, error)
}

type predictionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPredictionServiceClient(cc grpc.ClientConnInterface) PredictionServiceClient {
	return &predictionServiceClient{cc}
}

func (c *predictionServiceClient) UnaryServerUnaryClient(ctx context.Context, in *Simple, opts ...grpc.CallOption) (*Simple, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Simple)
	err := c.cc.Invoke(ctx, PredictionService_UnaryServerUnaryClient_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *predictionServiceClient) StreamServerUnaryClient(ctx context.Context, in *Simple, opts ...grpc.CallOption) (PredictionService_StreamServerUnaryClientClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &PredictionService_ServiceDesc.Streams[0], PredictionService_StreamServerUnaryClient_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &predictionServiceStreamServerUnaryClientClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PredictionService_StreamServerUnaryClientClient interface {
	Recv() (*Simple, error)
	grpc.ClientStream
}

type predictionServiceStreamServerUnaryClientClient struct {
	grpc.ClientStream
}

func (x *predictionServiceStreamServerUnaryClientClient) Recv() (*Simple, error) {
	m := new(Simple)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *predictionServiceClient) UnaryServerStreamClient(ctx context.Context, opts ...grpc.CallOption) (PredictionService_UnaryServerStreamClientClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &PredictionService_ServiceDesc.Streams[1], PredictionService_UnaryServerStreamClient_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &predictionServiceUnaryServerStreamClientClient{ClientStream: stream}
	return x, nil
}

type PredictionService_UnaryServerStreamClientClient interface {
	Send(*Simple) error
	CloseAndRecv() (*Simple, error)
	grpc.ClientStream
}

type predictionServiceUnaryServerStreamClientClient struct {
	grpc.ClientStream
}

func (x *predictionServiceUnaryServerStreamClientClient) Send(m *Simple) error {
	return x.ClientStream.SendMsg(m)
}

func (x *predictionServiceUnaryServerStreamClientClient) CloseAndRecv() (*Simple, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Simple)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *predictionServiceClient) StreamServerStreamClient(ctx context.Context, opts ...grpc.CallOption) (PredictionService_StreamServerStreamClientClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &PredictionService_ServiceDesc.Streams[2], PredictionService_StreamServerStreamClient_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &predictionServiceStreamServerStreamClientClient{ClientStream: stream}
	return x, nil
}

type PredictionService_StreamServerStreamClientClient interface {
	Send(*Simple) error
	Recv() (*Simple, error)
	grpc.ClientStream
}

type predictionServiceStreamServerStreamClientClient struct {
	grpc.ClientStream
}

func (x *predictionServiceStreamServerStreamClientClient) Send(m *Simple) error {
	return x.ClientStream.SendMsg(m)
}

func (x *predictionServiceStreamServerStreamClientClient) Recv() (*Simple, error) {
	m := new(Simple)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PredictionServiceServer is the server API for PredictionService service.
// All implementations must embed UnimplementedPredictionServiceServer
// for forward compatibility
type PredictionServiceServer interface {
	UnaryServerUnaryClient(context.Context, *Simple) (*Simple, error)
	StreamServerUnaryClient(*Simple, PredictionService_StreamServerUnaryClientServer) error
	UnaryServerStreamClient(PredictionService_UnaryServerStreamClientServer) error
	StreamServerStreamClient(PredictionService_StreamServerStreamClientServer) error
	mustEmbedUnimplementedPredictionServiceServer()
}

// UnimplementedPredictionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPredictionServiceServer struct {
}

func (UnimplementedPredictionServiceServer) UnaryServerUnaryClient(context.Context, *Simple) (*Simple, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnaryServerUnaryClient not implemented")
}
func (UnimplementedPredictionServiceServer) StreamServerUnaryClient(*Simple, PredictionService_StreamServerUnaryClientServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamServerUnaryClient not implemented")
}
func (UnimplementedPredictionServiceServer) UnaryServerStreamClient(PredictionService_UnaryServerStreamClientServer) error {
	return status.Errorf(codes.Unimplemented, "method UnaryServerStreamClient not implemented")
}
func (UnimplementedPredictionServiceServer) StreamServerStreamClient(PredictionService_StreamServerStreamClientServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamServerStreamClient not implemented")
}
func (UnimplementedPredictionServiceServer) mustEmbedUnimplementedPredictionServiceServer() {}

// UnsafePredictionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PredictionServiceServer will
// result in compilation errors.
type UnsafePredictionServiceServer interface {
	mustEmbedUnimplementedPredictionServiceServer()
}

func RegisterPredictionServiceServer(s grpc.ServiceRegistrar, srv PredictionServiceServer) {
	s.RegisterService(&PredictionService_ServiceDesc, srv)
}

func _PredictionService_UnaryServerUnaryClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Simple)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictionServiceServer).UnaryServerUnaryClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PredictionService_UnaryServerUnaryClient_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictionServiceServer).UnaryServerUnaryClient(ctx, req.(*Simple))
	}
	return interceptor(ctx, in, info, handler)
}

func _PredictionService_StreamServerUnaryClient_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Simple)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PredictionServiceServer).StreamServerUnaryClient(m, &predictionServiceStreamServerUnaryClientServer{ServerStream: stream})
}

type PredictionService_StreamServerUnaryClientServer interface {
	Send(*Simple) error
	grpc.ServerStream
}

type predictionServiceStreamServerUnaryClientServer struct {
	grpc.ServerStream
}

func (x *predictionServiceStreamServerUnaryClientServer) Send(m *Simple) error {
	return x.ServerStream.SendMsg(m)
}

func _PredictionService_UnaryServerStreamClient_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PredictionServiceServer).UnaryServerStreamClient(&predictionServiceUnaryServerStreamClientServer{ServerStream: stream})
}

type PredictionService_UnaryServerStreamClientServer interface {
	SendAndClose(*Simple) error
	Recv() (*Simple, error)
	grpc.ServerStream
}

type predictionServiceUnaryServerStreamClientServer struct {
	grpc.ServerStream
}

func (x *predictionServiceUnaryServerStreamClientServer) SendAndClose(m *Simple) error {
	return x.ServerStream.SendMsg(m)
}

func (x *predictionServiceUnaryServerStreamClientServer) Recv() (*Simple, error) {
	m := new(Simple)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PredictionService_StreamServerStreamClient_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PredictionServiceServer).StreamServerStreamClient(&predictionServiceStreamServerStreamClientServer{ServerStream: stream})
}

type PredictionService_StreamServerStreamClientServer interface {
	Send(*Simple) error
	Recv() (*Simple, error)
	grpc.ServerStream
}

type predictionServiceStreamServerStreamClientServer struct {
	grpc.ServerStream
}

func (x *predictionServiceStreamServerStreamClientServer) Send(m *Simple) error {
	return x.ServerStream.SendMsg(m)
}

func (x *predictionServiceStreamServerStreamClientServer) Recv() (*Simple, error) {
	m := new(Simple)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PredictionService_ServiceDesc is the grpc.ServiceDesc for PredictionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PredictionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "job_seek.prediction.PredictionService",
	HandlerType: (*PredictionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnaryServerUnaryClient",
			Handler:    _PredictionService_UnaryServerUnaryClient_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamServerUnaryClient",
			Handler:       _PredictionService_StreamServerUnaryClient_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "UnaryServerStreamClient",
			Handler:       _PredictionService_UnaryServerStreamClient_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamServerStreamClient",
			Handler:       _PredictionService_StreamServerStreamClient_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "prediction.proto",
}
