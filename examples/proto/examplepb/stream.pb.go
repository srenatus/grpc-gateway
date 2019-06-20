// Code generated by protoc-gen-go. DO NOT EDIT.
// source: examples/proto/examplepb/stream.proto

package examplepb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import sub "github.com/grpc-ecosystem/grpc-gateway/examples/proto/sub"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import httpbody "google.golang.org/genproto/googleapis/api/httpbody"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StreamServiceClient is the client API for StreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StreamServiceClient interface {
	BulkCreate(ctx context.Context, opts ...grpc.CallOption) (StreamService_BulkCreateClient, error)
	List(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (StreamService_ListClient, error)
	HttpBodyStream(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (StreamService_HttpBodyStreamClient, error)
	BulkEcho(ctx context.Context, opts ...grpc.CallOption) (StreamService_BulkEchoClient, error)
}

type streamServiceClient struct {
	cc *grpc.ClientConn
}

func NewStreamServiceClient(cc *grpc.ClientConn) StreamServiceClient {
	return &streamServiceClient{cc}
}

func (c *streamServiceClient) BulkCreate(ctx context.Context, opts ...grpc.CallOption) (StreamService_BulkCreateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamService_serviceDesc.Streams[0], "/grpc.gateway.examples.examplepb.StreamService/BulkCreate", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceBulkCreateClient{stream}
	return x, nil
}

type StreamService_BulkCreateClient interface {
	Send(*ABitOfEverything) error
	CloseAndRecv() (*empty.Empty, error)
	grpc.ClientStream
}

type streamServiceBulkCreateClient struct {
	grpc.ClientStream
}

func (x *streamServiceBulkCreateClient) Send(m *ABitOfEverything) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamServiceBulkCreateClient) CloseAndRecv() (*empty.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(empty.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamServiceClient) List(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (StreamService_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamService_serviceDesc.Streams[1], "/grpc.gateway.examples.examplepb.StreamService/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamService_ListClient interface {
	Recv() (*ABitOfEverything, error)
	grpc.ClientStream
}

type streamServiceListClient struct {
	grpc.ClientStream
}

func (x *streamServiceListClient) Recv() (*ABitOfEverything, error) {
	m := new(ABitOfEverything)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamServiceClient) HttpBodyStream(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (StreamService_HttpBodyStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamService_serviceDesc.Streams[2], "/grpc.gateway.examples.examplepb.StreamService/HttpBodyStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceHttpBodyStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamService_HttpBodyStreamClient interface {
	Recv() (*httpbody.HttpBody, error)
	grpc.ClientStream
}

type streamServiceHttpBodyStreamClient struct {
	grpc.ClientStream
}

func (x *streamServiceHttpBodyStreamClient) Recv() (*httpbody.HttpBody, error) {
	m := new(httpbody.HttpBody)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamServiceClient) BulkEcho(ctx context.Context, opts ...grpc.CallOption) (StreamService_BulkEchoClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamService_serviceDesc.Streams[3], "/grpc.gateway.examples.examplepb.StreamService/BulkEcho", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceBulkEchoClient{stream}
	return x, nil
}

type StreamService_BulkEchoClient interface {
	Send(*sub.StringMessage) error
	Recv() (*sub.StringMessage, error)
	grpc.ClientStream
}

type streamServiceBulkEchoClient struct {
	grpc.ClientStream
}

func (x *streamServiceBulkEchoClient) Send(m *sub.StringMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamServiceBulkEchoClient) Recv() (*sub.StringMessage, error) {
	m := new(sub.StringMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamServiceServer is the server API for StreamService service.
type StreamServiceServer interface {
	BulkCreate(StreamService_BulkCreateServer) error
	List(*empty.Empty, StreamService_ListServer) error
	HttpBodyStream(*empty.Empty, StreamService_HttpBodyStreamServer) error
	BulkEcho(StreamService_BulkEchoServer) error
}

func RegisterStreamServiceServer(s *grpc.Server, srv StreamServiceServer) {
	s.RegisterService(&_StreamService_serviceDesc, srv)
}

func _StreamService_BulkCreate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServiceServer).BulkCreate(&streamServiceBulkCreateServer{stream})
}

type StreamService_BulkCreateServer interface {
	SendAndClose(*empty.Empty) error
	Recv() (*ABitOfEverything, error)
	grpc.ServerStream
}

type streamServiceBulkCreateServer struct {
	grpc.ServerStream
}

func (x *streamServiceBulkCreateServer) SendAndClose(m *empty.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamServiceBulkCreateServer) Recv() (*ABitOfEverything, error) {
	m := new(ABitOfEverything)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _StreamService_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamServiceServer).List(m, &streamServiceListServer{stream})
}

type StreamService_ListServer interface {
	Send(*ABitOfEverything) error
	grpc.ServerStream
}

type streamServiceListServer struct {
	grpc.ServerStream
}

func (x *streamServiceListServer) Send(m *ABitOfEverything) error {
	return x.ServerStream.SendMsg(m)
}

func _StreamService_HttpBodyStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamServiceServer).HttpBodyStream(m, &streamServiceHttpBodyStreamServer{stream})
}

type StreamService_HttpBodyStreamServer interface {
	Send(*httpbody.HttpBody) error
	grpc.ServerStream
}

type streamServiceHttpBodyStreamServer struct {
	grpc.ServerStream
}

func (x *streamServiceHttpBodyStreamServer) Send(m *httpbody.HttpBody) error {
	return x.ServerStream.SendMsg(m)
}

func _StreamService_BulkEcho_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServiceServer).BulkEcho(&streamServiceBulkEchoServer{stream})
}

type StreamService_BulkEchoServer interface {
	Send(*sub.StringMessage) error
	Recv() (*sub.StringMessage, error)
	grpc.ServerStream
}

type streamServiceBulkEchoServer struct {
	grpc.ServerStream
}

func (x *streamServiceBulkEchoServer) Send(m *sub.StringMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamServiceBulkEchoServer) Recv() (*sub.StringMessage, error) {
	m := new(sub.StringMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _StreamService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.gateway.examples.examplepb.StreamService",
	HandlerType: (*StreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BulkCreate",
			Handler:       _StreamService_BulkCreate_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "List",
			Handler:       _StreamService_List_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "HttpBodyStream",
			Handler:       _StreamService_HttpBodyStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "BulkEcho",
			Handler:       _StreamService_BulkEcho_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "examples/proto/examplepb/stream.proto",
}

func init() {
	proto.RegisterFile("examples/proto/examplepb/stream.proto", fileDescriptor_stream_2e65de8b98ed6ca3)
}

var fileDescriptor_stream_2e65de8b98ed6ca3 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x89, 0x14, 0xd1, 0x88, 0x1e, 0x16, 0x11, 0x8c, 0x42, 0x35, 0x28, 0xd6, 0x1e, 0x76,
	0xd3, 0x7a, 0xf3, 0x66, 0xa4, 0xe0, 0x41, 0xf1, 0xd0, 0x9b, 0x97, 0xb2, 0x9b, 0x6e, 0x93, 0xa5,
	0x49, 0x76, 0xd9, 0x9d, 0x54, 0x03, 0x9e, 0x3c, 0x7a, 0xed, 0x2b, 0xf8, 0x46, 0xbe, 0x82, 0x0f,
	0x22, 0xcd, 0x9f, 0x22, 0x62, 0x68, 0x3d, 0xee, 0x7c, 0xf3, 0x65, 0xbe, 0xdf, 0x4c, 0xec, 0x73,
	0xfe, 0x42, 0x13, 0x15, 0x73, 0x43, 0x94, 0x96, 0x20, 0x49, 0xf5, 0x54, 0x8c, 0x18, 0xd0, 0x9c,
	0x26, 0xb8, 0x28, 0xa3, 0x76, 0xa8, 0x55, 0x80, 0x43, 0x0a, 0xfc, 0x99, 0xe6, 0xb8, 0xf6, 0xe0,
	0x65, 0xb7, 0x73, 0x1c, 0x4a, 0x19, 0xc6, 0x9c, 0x50, 0x25, 0x08, 0x4d, 0x53, 0x09, 0x14, 0x84,
	0x4c, 0x4d, 0x69, 0x77, 0x0e, 0x7f, 0xa8, 0x11, 0x80, 0x62, 0x72, 0x9c, 0x57, 0xd2, 0x51, 0x25,
	0x15, 0x2f, 0x96, 0x4d, 0x08, 0x4f, 0x14, 0xd4, 0x62, 0xbf, 0x31, 0x1d, 0x1d, 0x31, 0x01, 0x23,
	0x39, 0x19, 0xf1, 0x19, 0xd7, 0x39, 0x44, 0x22, 0x0d, 0x2b, 0xcf, 0xc9, 0x2f, 0x8f, 0xc9, 0x18,
	0x49, 0xb8, 0x31, 0x34, 0xe4, 0x65, 0x47, 0xff, 0xa3, 0x65, 0xef, 0x0e, 0x0b, 0xba, 0x21, 0xd7,
	0x33, 0x11, 0x70, 0xf4, 0x6e, 0xd9, 0xb6, 0x9f, 0xc5, 0xd3, 0x5b, 0xcd, 0x29, 0x70, 0xd4, 0xc3,
	0x2b, 0x70, 0xf1, 0x8d, 0x2f, 0xe0, 0x71, 0x32, 0x58, 0xce, 0x76, 0x0e, 0x70, 0xc9, 0x81, 0x6b,
	0x0e, 0x3c, 0x58, 0x70, 0xb8, 0xe4, 0xed, 0xf3, 0x6b, 0xbe, 0x71, 0xe9, 0x9e, 0x91, 0x59, 0xaf,
	0x8e, 0xff, 0x57, 0x78, 0xc2, 0xb2, 0x78, 0x7a, 0x6d, 0x75, 0x3b, 0x16, 0x7a, 0xb5, 0x5b, 0xf7,
	0xc2, 0x00, 0x6a, 0xf8, 0xa4, 0xf3, 0xff, 0x74, 0xee, 0x45, 0x91, 0xe2, 0x14, 0xb5, 0x57, 0xa4,
	0xf0, 0x2c, 0xa4, 0xed, 0xbd, 0x3b, 0x00, 0xe5, 0xcb, 0x71, 0x5e, 0xee, 0xa8, 0x31, 0xc7, 0x7e,
	0x5d, 0xa7, 0x4a, 0xe0, 0xda, 0xe3, 0x7a, 0xc5, 0xa8, 0x2e, 0xea, 0xac, 0x02, 0xae, 0xff, 0x03,
	0xcf, 0x42, 0x73, 0xcb, 0xde, 0x5a, 0xac, 0x7f, 0x10, 0x44, 0x12, 0x75, 0x1a, 0xf0, 0x4c, 0xc6,
	0xf0, 0x10, 0xb4, 0x48, 0xc3, 0x87, 0xf2, 0x9a, 0xce, 0xda, 0x9d, 0xeb, 0x5f, 0x81, 0x07, 0x91,
	0x2c, 0xae, 0xe0, 0x59, 0xfe, 0xce, 0xd3, 0xf6, 0x72, 0xa5, 0x6c, 0xb3, 0x80, 0xbf, 0xfa, 0x0e,
	0x00, 0x00, 0xff, 0xff, 0x5e, 0xc0, 0x30, 0xea, 0x30, 0x03, 0x00, 0x00,
}
