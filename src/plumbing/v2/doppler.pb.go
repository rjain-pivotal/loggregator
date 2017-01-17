// Code generated by protoc-gen-go.
// source: doppler.proto
// DO NOT EDIT!

/*
Package loggregator is a generated protocol buffer package.

It is generated from these files:
	doppler.proto
	envelope.proto
	metron.proto

It has these top-level messages:
	SenderResponse
	Envelope
	Value
	Log
	Counter
	Gauge
	GaugeValue
	Timer
	MetronResponse
*/
package loggregator

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type SenderResponse struct {
}

func (m *SenderResponse) Reset()                    { *m = SenderResponse{} }
func (m *SenderResponse) String() string            { return proto.CompactTextString(m) }
func (*SenderResponse) ProtoMessage()               {}
func (*SenderResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*SenderResponse)(nil), "loggregator.SenderResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for DopplerIngress service

type DopplerIngressClient interface {
	Sender(ctx context.Context, opts ...grpc.CallOption) (DopplerIngress_SenderClient, error)
}

type dopplerIngressClient struct {
	cc *grpc.ClientConn
}

func NewDopplerIngressClient(cc *grpc.ClientConn) DopplerIngressClient {
	return &dopplerIngressClient{cc}
}

func (c *dopplerIngressClient) Sender(ctx context.Context, opts ...grpc.CallOption) (DopplerIngress_SenderClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_DopplerIngress_serviceDesc.Streams[0], c.cc, "/loggregator.DopplerIngress/Sender", opts...)
	if err != nil {
		return nil, err
	}
	x := &dopplerIngressSenderClient{stream}
	return x, nil
}

type DopplerIngress_SenderClient interface {
	Send(*Envelope) error
	CloseAndRecv() (*SenderResponse, error)
	grpc.ClientStream
}

type dopplerIngressSenderClient struct {
	grpc.ClientStream
}

func (x *dopplerIngressSenderClient) Send(m *Envelope) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dopplerIngressSenderClient) CloseAndRecv() (*SenderResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SenderResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for DopplerIngress service

type DopplerIngressServer interface {
	Sender(DopplerIngress_SenderServer) error
}

func RegisterDopplerIngressServer(s *grpc.Server, srv DopplerIngressServer) {
	s.RegisterService(&_DopplerIngress_serviceDesc, srv)
}

func _DopplerIngress_Sender_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DopplerIngressServer).Sender(&dopplerIngressSenderServer{stream})
}

type DopplerIngress_SenderServer interface {
	SendAndClose(*SenderResponse) error
	Recv() (*Envelope, error)
	grpc.ServerStream
}

type dopplerIngressSenderServer struct {
	grpc.ServerStream
}

func (x *dopplerIngressSenderServer) SendAndClose(m *SenderResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dopplerIngressSenderServer) Recv() (*Envelope, error) {
	m := new(Envelope)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _DopplerIngress_serviceDesc = grpc.ServiceDesc{
	ServiceName: "loggregator.DopplerIngress",
	HandlerType: (*DopplerIngressServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Sender",
			Handler:       _DopplerIngress_Sender_Handler,
			ClientStreams: true,
		},
	},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("doppler.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 127 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xc9, 0x2f, 0x28,
	0xc8, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xce, 0xc9, 0x4f, 0x4f, 0x2f,
	0x4a, 0x4d, 0x4f, 0x2c, 0xc9, 0x2f, 0x92, 0xe2, 0x4b, 0xcd, 0x2b, 0x4b, 0xcd, 0xc9, 0x2f, 0x48,
	0x85, 0x48, 0x2a, 0x09, 0x70, 0xf1, 0x05, 0xa7, 0xe6, 0xa5, 0xa4, 0x16, 0x05, 0xa5, 0x16, 0x17,
	0xe4, 0xe7, 0x15, 0xa7, 0x1a, 0x05, 0x71, 0xf1, 0xb9, 0x40, 0xf4, 0x7b, 0xe6, 0xa5, 0x17, 0xa5,
	0x16, 0x17, 0x0b, 0x39, 0x70, 0xb1, 0x41, 0xd4, 0x08, 0x89, 0xea, 0x21, 0x99, 0xa5, 0xe7, 0x0a,
	0x35, 0x4a, 0x4a, 0x1a, 0x45, 0x18, 0xd5, 0x3c, 0x25, 0x06, 0x0d, 0xc6, 0x24, 0x36, 0xb0, 0x65,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xea, 0x24, 0x80, 0x9a, 0x00, 0x00, 0x00,
}