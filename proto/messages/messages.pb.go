// Code generated by protoc-gen-go.
// source: proto/messages/messages.proto
// DO NOT EDIT!

/*
Package messages is a generated protocol buffer package.

It is generated from these files:
	proto/messages/messages.proto

It has these top-level messages:
	MessageRequest
	MessageReply
*/
package messages

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

// The request message containing the new message's content
type MessageRequest struct {
	User []byte `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Text []byte `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (m *MessageRequest) Reset()                    { *m = MessageRequest{} }
func (m *MessageRequest) String() string            { return proto.CompactTextString(m) }
func (*MessageRequest) ProtoMessage()               {}
func (*MessageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// The response message containing the newely created message
type MessageReply struct {
	Id        []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	User      []byte `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Text      []byte `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	Timestamp int64  `protobuf:"varint,4,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *MessageReply) Reset()                    { *m = MessageReply{} }
func (m *MessageReply) String() string            { return proto.CompactTextString(m) }
func (*MessageReply) ProtoMessage()               {}
func (*MessageReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*MessageRequest)(nil), "messages.MessageRequest")
	proto.RegisterType((*MessageReply)(nil), "messages.MessageReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for Message service

type MessageClient interface {
	// Sends a greeting
	Create(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageReply, error)
}

type messageClient struct {
	cc *grpc.ClientConn
}

func NewMessageClient(cc *grpc.ClientConn) MessageClient {
	return &messageClient{cc}
}

func (c *messageClient) Create(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageReply, error) {
	out := new(MessageReply)
	err := grpc.Invoke(ctx, "/messages.Message/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Message service

type MessageServer interface {
	// Sends a greeting
	Create(context.Context, *MessageRequest) (*MessageReply, error)
}

func RegisterMessageServer(s *grpc.Server, srv MessageServer) {
	s.RegisterService(&_Message_serviceDesc, srv)
}

func _Message_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messages.Message/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).Create(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Message_serviceDesc = grpc.ServiceDesc{
	ServiceName: "messages.Message",
	HandlerType: (*MessageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Message_Create_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x2d, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x2d, 0x86, 0x33, 0xf4, 0xc0, 0xe2, 0x42,
	0x1c, 0x30, 0xbe, 0x92, 0x05, 0x17, 0x9f, 0x2f, 0x84, 0x1d, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c,
	0x22, 0x24, 0xc4, 0xc5, 0x52, 0x5a, 0x9c, 0x5a, 0x24, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x13, 0x04,
	0x66, 0x83, 0xc4, 0x4a, 0x52, 0x2b, 0x4a, 0x24, 0x98, 0x20, 0x62, 0x20, 0xb6, 0x52, 0x0a, 0x17,
	0x0f, 0x5c, 0x67, 0x41, 0x4e, 0xa5, 0x10, 0x1f, 0x17, 0x53, 0x66, 0x0a, 0x54, 0x17, 0x90, 0x05,
	0x37, 0x87, 0x09, 0x8b, 0x39, 0xcc, 0x08, 0x73, 0x84, 0x64, 0xb8, 0x38, 0x4b, 0x32, 0x81, 0xee,
	0x29, 0x49, 0xcc, 0x2d, 0x90, 0x60, 0x01, 0x4a, 0x30, 0x07, 0x21, 0x04, 0x8c, 0xdc, 0xb9, 0xd8,
	0xa1, 0xb6, 0x08, 0xd9, 0x70, 0xb1, 0x39, 0x17, 0xa5, 0x26, 0x96, 0xa4, 0x0a, 0x49, 0xe8, 0xc1,
	0xfd, 0x83, 0xea, 0x78, 0x29, 0x31, 0x2c, 0x32, 0x40, 0xc7, 0x29, 0x31, 0x24, 0xb1, 0x81, 0x7d,
	0x6e, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x10, 0x68, 0x61, 0x63, 0x1a, 0x01, 0x00, 0x00,
}
