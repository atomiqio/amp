// Code generated by protoc-gen-go.
// source: github.com/appcelerator/amp/api/rpc/stack/stack.proto
// DO NOT EDIT!

/*
Package stack is a generated protocol buffer package.

It is generated from these files:
	github.com/appcelerator/amp/api/rpc/stack/stack.proto

It has these top-level messages:
	DeployRequest
	DeployReply
	ListRequest
	ListReply
	StackListEntry
	RemoveRequest
	RemoveReply
*/
package stack

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import stacks "github.com/appcelerator/amp/data/stacks"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type DeployRequest struct {
	Name    string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Compose []byte `protobuf:"bytes,2,opt,name=compose,proto3" json:"compose,omitempty"`
}

func (m *DeployRequest) Reset()                    { *m = DeployRequest{} }
func (m *DeployRequest) String() string            { return proto.CompactTextString(m) }
func (*DeployRequest) ProtoMessage()               {}
func (*DeployRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DeployRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeployRequest) GetCompose() []byte {
	if m != nil {
		return m.Compose
	}
	return nil
}

type DeployReply struct {
	FullName string `protobuf:"bytes,1,opt,name=full_name,json=fullName" json:"full_name,omitempty"`
	Answer   string `protobuf:"bytes,2,opt,name=answer" json:"answer,omitempty"`
}

func (m *DeployReply) Reset()                    { *m = DeployReply{} }
func (m *DeployReply) String() string            { return proto.CompactTextString(m) }
func (*DeployReply) ProtoMessage()               {}
func (*DeployReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DeployReply) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *DeployReply) GetAnswer() string {
	if m != nil {
		return m.Answer
	}
	return ""
}

type ListRequest struct {
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type ListReply struct {
	Entries []*StackListEntry `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty"`
}

func (m *ListReply) Reset()                    { *m = ListReply{} }
func (m *ListReply) String() string            { return proto.CompactTextString(m) }
func (*ListReply) ProtoMessage()               {}
func (*ListReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ListReply) GetEntries() []*StackListEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type StackListEntry struct {
	Stack    *stacks.Stack `protobuf:"bytes,1,opt,name=stack" json:"stack,omitempty"`
	Services string        `protobuf:"bytes,2,opt,name=services" json:"services,omitempty"`
}

func (m *StackListEntry) Reset()                    { *m = StackListEntry{} }
func (m *StackListEntry) String() string            { return proto.CompactTextString(m) }
func (*StackListEntry) ProtoMessage()               {}
func (*StackListEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *StackListEntry) GetStack() *stacks.Stack {
	if m != nil {
		return m.Stack
	}
	return nil
}

func (m *StackListEntry) GetServices() string {
	if m != nil {
		return m.Services
	}
	return ""
}

type RemoveRequest struct {
	Stack string `protobuf:"bytes,1,opt,name=stack" json:"stack,omitempty"`
}

func (m *RemoveRequest) Reset()                    { *m = RemoveRequest{} }
func (m *RemoveRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveRequest) ProtoMessage()               {}
func (*RemoveRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RemoveRequest) GetStack() string {
	if m != nil {
		return m.Stack
	}
	return ""
}

type RemoveReply struct {
	Answer string `protobuf:"bytes,1,opt,name=answer" json:"answer,omitempty"`
}

func (m *RemoveReply) Reset()                    { *m = RemoveReply{} }
func (m *RemoveReply) String() string            { return proto.CompactTextString(m) }
func (*RemoveReply) ProtoMessage()               {}
func (*RemoveReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *RemoveReply) GetAnswer() string {
	if m != nil {
		return m.Answer
	}
	return ""
}

func init() {
	proto.RegisterType((*DeployRequest)(nil), "stack.DeployRequest")
	proto.RegisterType((*DeployReply)(nil), "stack.DeployReply")
	proto.RegisterType((*ListRequest)(nil), "stack.ListRequest")
	proto.RegisterType((*ListReply)(nil), "stack.ListReply")
	proto.RegisterType((*StackListEntry)(nil), "stack.StackListEntry")
	proto.RegisterType((*RemoveRequest)(nil), "stack.RemoveRequest")
	proto.RegisterType((*RemoveReply)(nil), "stack.RemoveReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Stack service

type StackClient interface {
	Deploy(ctx context.Context, in *DeployRequest, opts ...grpc.CallOption) (*DeployReply, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error)
	Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveReply, error)
}

type stackClient struct {
	cc *grpc.ClientConn
}

func NewStackClient(cc *grpc.ClientConn) StackClient {
	return &stackClient{cc}
}

func (c *stackClient) Deploy(ctx context.Context, in *DeployRequest, opts ...grpc.CallOption) (*DeployReply, error) {
	out := new(DeployReply)
	err := grpc.Invoke(ctx, "/stack.Stack/Deploy", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stackClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error) {
	out := new(ListReply)
	err := grpc.Invoke(ctx, "/stack.Stack/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stackClient) Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveReply, error) {
	out := new(RemoveReply)
	err := grpc.Invoke(ctx, "/stack.Stack/Remove", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Stack service

type StackServer interface {
	Deploy(context.Context, *DeployRequest) (*DeployReply, error)
	List(context.Context, *ListRequest) (*ListReply, error)
	Remove(context.Context, *RemoveRequest) (*RemoveReply, error)
}

func RegisterStackServer(s *grpc.Server, srv StackServer) {
	s.RegisterService(&_Stack_serviceDesc, srv)
}

func _Stack_Deploy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeployRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StackServer).Deploy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stack.Stack/Deploy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StackServer).Deploy(ctx, req.(*DeployRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stack_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StackServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stack.Stack/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StackServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stack_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StackServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stack.Stack/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StackServer).Remove(ctx, req.(*RemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Stack_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stack.Stack",
	HandlerType: (*StackServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Deploy",
			Handler:    _Stack_Deploy_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Stack_List_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _Stack_Remove_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/appcelerator/amp/api/rpc/stack/stack.proto",
}

func init() {
	proto.RegisterFile("github.com/appcelerator/amp/api/rpc/stack/stack.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xdf, 0x8a, 0xd4, 0x30,
	0x14, 0xc6, 0xe9, 0xba, 0xdb, 0xdd, 0x9e, 0x6e, 0x55, 0x0e, 0xb3, 0x52, 0xaa, 0x17, 0x43, 0x64,
	0x61, 0xd8, 0x8b, 0x06, 0x47, 0xbd, 0x11, 0x05, 0x11, 0xbd, 0x10, 0x64, 0xc1, 0xfa, 0x00, 0x92,
	0xad, 0x71, 0x2c, 0xb6, 0x4d, 0x4c, 0x32, 0x23, 0x45, 0xbc, 0xf1, 0x15, 0x7c, 0x34, 0x5f, 0x61,
	0x1f, 0x44, 0xf2, 0xa7, 0xb3, 0xad, 0x17, 0xde, 0x24, 0x73, 0x92, 0x73, 0x7e, 0xf3, 0x7d, 0x5f,
	0x0a, 0x4f, 0x37, 0x8d, 0xf9, 0xb2, 0xbd, 0x2a, 0x6b, 0xd1, 0x51, 0x26, 0x65, 0xcd, 0x5b, 0xae,
	0x98, 0x11, 0x8a, 0xb2, 0x4e, 0x52, 0x26, 0x1b, 0xaa, 0x64, 0x4d, 0xb5, 0x61, 0xf5, 0x57, 0xbf,
	0x96, 0x52, 0x09, 0x23, 0xf0, 0xc8, 0x15, 0xc5, 0x93, 0xff, 0x4d, 0x7f, 0x62, 0x86, 0xf9, 0x21,
	0x1d, 0x36, 0x3f, 0x5c, 0x3c, 0xd8, 0x08, 0xb1, 0x69, 0xb9, 0xc3, 0xb3, 0xbe, 0x17, 0x86, 0x99,
	0x46, 0xf4, 0xe1, 0x96, 0xbc, 0x80, 0xec, 0x35, 0x97, 0xad, 0x18, 0x2a, 0xfe, 0x6d, 0xcb, 0xb5,
	0x41, 0x84, 0xc3, 0x9e, 0x75, 0x3c, 0x8f, 0x96, 0xd1, 0x2a, 0xa9, 0xdc, 0x6f, 0xcc, 0xe1, 0xb8,
	0x16, 0x9d, 0x14, 0x9a, 0xe7, 0x07, 0xcb, 0x68, 0x75, 0x5a, 0x8d, 0x25, 0x79, 0x05, 0xe9, 0x38,
	0x2e, 0xdb, 0x01, 0xef, 0x43, 0xf2, 0x79, 0xdb, 0xb6, 0x1f, 0x27, 0x84, 0x13, 0x7b, 0x70, 0x69,
	0x29, 0xf7, 0x20, 0x66, 0xbd, 0xfe, 0xce, 0x95, 0x83, 0x24, 0x55, 0xa8, 0x48, 0x06, 0xe9, 0xbb,
	0x46, 0x9b, 0x20, 0x80, 0x3c, 0x87, 0xc4, 0x97, 0x16, 0x48, 0xe1, 0x98, 0xf7, 0x46, 0x35, 0x5c,
	0xe7, 0xd1, 0xf2, 0xd6, 0x2a, 0x5d, 0x9f, 0x95, 0x3e, 0x98, 0x0f, 0x76, 0xb5, 0x7d, 0x6f, 0x7a,
	0xa3, 0x86, 0x6a, 0xec, 0x22, 0xef, 0xe1, 0xf6, 0xfc, 0x0a, 0x1f, 0x82, 0x8f, 0xcf, 0xe9, 0x49,
	0xd7, 0x59, 0x19, 0xd2, 0x71, 0x6d, 0x95, 0xbf, 0xc3, 0x02, 0x4e, 0x34, 0x57, 0xbb, 0xa6, 0xe6,
	0x3a, 0xa8, 0xdb, 0xd7, 0xe4, 0x1c, 0xb2, 0x8a, 0x77, 0x62, 0xc7, 0xc7, 0x88, 0x16, 0x53, 0x62,
	0x12, 0x10, 0xe4, 0x1c, 0xd2, 0xb1, 0xcd, 0x2a, 0xbf, 0x71, 0x1b, 0x4d, 0xdd, 0xae, 0xaf, 0x23,
	0x38, 0x72, 0x7f, 0x8d, 0x6f, 0x21, 0xf6, 0xd9, 0xe1, 0x22, 0x98, 0x9a, 0xbd, 0x44, 0x81, 0xff,
	0x9c, 0xca, 0x76, 0x20, 0x67, 0xbf, 0xfe, 0x5c, 0xff, 0x3e, 0xb8, 0x43, 0x80, 0xee, 0x1e, 0x85,
	0x67, 0x7e, 0x16, 0x5d, 0xe0, 0x4b, 0x38, 0xb4, 0x86, 0x71, 0x1c, 0x99, 0xe4, 0x59, 0xdc, 0x9d,
	0x9d, 0x59, 0x08, 0x3a, 0xc8, 0x29, 0x4e, 0x20, 0x78, 0x09, 0xb1, 0x57, 0xbf, 0x17, 0x33, 0xf3,
	0xbc, 0x17, 0x33, 0xb1, 0x48, 0x0a, 0xc7, 0x59, 0x5c, 0xe0, 0x0d, 0x87, 0xfe, 0x70, 0xfb, 0xcf,
	0xab, 0xd8, 0x7d, 0x5e, 0x8f, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x3a, 0x64, 0xb2, 0x97, 0xf2,
	0x02, 0x00, 0x00,
}
