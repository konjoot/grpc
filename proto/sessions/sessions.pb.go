// Code generated by protoc-gen-go.
// source: proto/sessions/sessions.proto
// DO NOT EDIT!

/*
Package sessions is a generated protocol buffer package.

It is generated from these files:
	proto/sessions/sessions.proto

It has these top-level messages:
	SessionRequest
	SessionReply
*/
package sessions

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

// The request message containing the user's credentials.
type SessionRequest struct {
	Login []byte `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Pass  []byte `protobuf:"bytes,2,opt,name=pass,proto3" json:"pass,omitempty"`
}

func (m *SessionRequest) Reset()                    { *m = SessionRequest{} }
func (m *SessionRequest) String() string            { return proto.CompactTextString(m) }
func (*SessionRequest) ProtoMessage()               {}
func (*SessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// The response message containing the token
type SessionReply struct {
	Token []byte `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (m *SessionReply) Reset()                    { *m = SessionReply{} }
func (m *SessionReply) String() string            { return proto.CompactTextString(m) }
func (*SessionReply) ProtoMessage()               {}
func (*SessionReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*SessionRequest)(nil), "sessions.SessionRequest")
	proto.RegisterType((*SessionReply)(nil), "sessions.SessionReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for Session service

type SessionClient interface {
	// Sends a greeting
	Create(ctx context.Context, in *SessionRequest, opts ...grpc.CallOption) (*SessionReply, error)
}

type sessionClient struct {
	cc *grpc.ClientConn
}

func NewSessionClient(cc *grpc.ClientConn) SessionClient {
	return &sessionClient{cc}
}

func (c *sessionClient) Create(ctx context.Context, in *SessionRequest, opts ...grpc.CallOption) (*SessionReply, error) {
	out := new(SessionReply)
	err := grpc.Invoke(ctx, "/sessions.Session/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Session service

type SessionServer interface {
	// Sends a greeting
	Create(context.Context, *SessionRequest) (*SessionReply, error)
}

func RegisterSessionServer(s *grpc.Server, srv SessionServer) {
	s.RegisterService(&_Session_serviceDesc, srv)
}

func _Session_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sessions.Session/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServer).Create(ctx, req.(*SessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Session_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sessions.Session",
	HandlerType: (*SessionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Session_Create_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 149 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x2d, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x4e, 0x2d, 0x2e, 0xce, 0xcc, 0xcf, 0x2b, 0x86, 0x33, 0xf4, 0xc0, 0xe2, 0x42,
	0x1c, 0x30, 0xbe, 0x92, 0x15, 0x17, 0x5f, 0x30, 0x84, 0x1d, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c,
	0x22, 0x24, 0xc2, 0xc5, 0x9a, 0x93, 0x9f, 0x9e, 0x99, 0x27, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x13,
	0x04, 0xe1, 0x08, 0x09, 0x71, 0xb1, 0x14, 0x24, 0x16, 0x17, 0x4b, 0x30, 0x81, 0x05, 0xc1, 0x6c,
	0x25, 0x15, 0x2e, 0x1e, 0xb8, 0xde, 0x82, 0x9c, 0x4a, 0x90, 0xce, 0x92, 0xfc, 0xec, 0x54, 0xb8,
	0x4e, 0x30, 0xc7, 0xc8, 0x9d, 0x8b, 0x1d, 0xaa, 0x4a, 0xc8, 0x86, 0x8b, 0xcd, 0xb9, 0x28, 0x35,
	0xb1, 0x24, 0x55, 0x48, 0x42, 0x0f, 0xee, 0x22, 0x54, 0xeb, 0xa5, 0xc4, 0xb0, 0xc8, 0x00, 0x0d,
	0x57, 0x62, 0x48, 0x62, 0x03, 0xbb, 0xdd, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x5a, 0xc9, 0x62,
	0x68, 0xdc, 0x00, 0x00, 0x00,
}
