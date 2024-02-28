// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: api.proto

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

// ElectrocardiogramClient is the client API for Electrocardiogram service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ElectrocardiogramClient interface {
	ListRecords(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*ListRecordsResponse, error)
	StreamRecords(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Electrocardiogram_StreamRecordsClient, error)
	Analyze(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*Analysis, error)
}

type electrocardiogramClient struct {
	cc grpc.ClientConnInterface
}

func NewElectrocardiogramClient(cc grpc.ClientConnInterface) ElectrocardiogramClient {
	return &electrocardiogramClient{cc}
}

func (c *electrocardiogramClient) ListRecords(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*ListRecordsResponse, error) {
	out := new(ListRecordsResponse)
	err := c.cc.Invoke(ctx, "/Electrocardiogram/ListRecords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *electrocardiogramClient) StreamRecords(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Electrocardiogram_StreamRecordsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Electrocardiogram_ServiceDesc.Streams[0], "/Electrocardiogram/StreamRecords", opts...)
	if err != nil {
		return nil, err
	}
	x := &electrocardiogramStreamRecordsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Electrocardiogram_StreamRecordsClient interface {
	Recv() (*Record, error)
	grpc.ClientStream
}

type electrocardiogramStreamRecordsClient struct {
	grpc.ClientStream
}

func (x *electrocardiogramStreamRecordsClient) Recv() (*Record, error) {
	m := new(Record)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *electrocardiogramClient) Analyze(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*Analysis, error) {
	out := new(Analysis)
	err := c.cc.Invoke(ctx, "/Electrocardiogram/Analyze", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ElectrocardiogramServer is the server API for Electrocardiogram service.
// All implementations must embed UnimplementedElectrocardiogramServer
// for forward compatibility
type ElectrocardiogramServer interface {
	ListRecords(context.Context, *Filter) (*ListRecordsResponse, error)
	StreamRecords(*Empty, Electrocardiogram_StreamRecordsServer) error
	Analyze(context.Context, *Filter) (*Analysis, error)
	mustEmbedUnimplementedElectrocardiogramServer()
}

// UnimplementedElectrocardiogramServer must be embedded to have forward compatible implementations.
type UnimplementedElectrocardiogramServer struct {
}

func (UnimplementedElectrocardiogramServer) ListRecords(context.Context, *Filter) (*ListRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRecords not implemented")
}
func (UnimplementedElectrocardiogramServer) StreamRecords(*Empty, Electrocardiogram_StreamRecordsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamRecords not implemented")
}
func (UnimplementedElectrocardiogramServer) Analyze(context.Context, *Filter) (*Analysis, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Analyze not implemented")
}
func (UnimplementedElectrocardiogramServer) mustEmbedUnimplementedElectrocardiogramServer() {}

// UnsafeElectrocardiogramServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ElectrocardiogramServer will
// result in compilation errors.
type UnsafeElectrocardiogramServer interface {
	mustEmbedUnimplementedElectrocardiogramServer()
}

func RegisterElectrocardiogramServer(s grpc.ServiceRegistrar, srv ElectrocardiogramServer) {
	s.RegisterService(&Electrocardiogram_ServiceDesc, srv)
}

func _Electrocardiogram_ListRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElectrocardiogramServer).ListRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Electrocardiogram/ListRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElectrocardiogramServer).ListRecords(ctx, req.(*Filter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Electrocardiogram_StreamRecords_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ElectrocardiogramServer).StreamRecords(m, &electrocardiogramStreamRecordsServer{stream})
}

type Electrocardiogram_StreamRecordsServer interface {
	Send(*Record) error
	grpc.ServerStream
}

type electrocardiogramStreamRecordsServer struct {
	grpc.ServerStream
}

func (x *electrocardiogramStreamRecordsServer) Send(m *Record) error {
	return x.ServerStream.SendMsg(m)
}

func _Electrocardiogram_Analyze_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElectrocardiogramServer).Analyze(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Electrocardiogram/Analyze",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElectrocardiogramServer).Analyze(ctx, req.(*Filter))
	}
	return interceptor(ctx, in, info, handler)
}

// Electrocardiogram_ServiceDesc is the grpc.ServiceDesc for Electrocardiogram service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Electrocardiogram_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Electrocardiogram",
	HandlerType: (*ElectrocardiogramServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRecords",
			Handler:    _Electrocardiogram_ListRecords_Handler,
		},
		{
			MethodName: "Analyze",
			Handler:    _Electrocardiogram_Analyze_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamRecords",
			Handler:       _Electrocardiogram_StreamRecords_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api.proto",
}
