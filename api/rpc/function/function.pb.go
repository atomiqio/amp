// Code generated by protoc-gen-go.
// source: github.com/appcelerator/amp/api/rpc/function/function.proto
// DO NOT EDIT!

/*
Package function is a generated protocol buffer package.

It is generated from these files:
	github.com/appcelerator/amp/api/rpc/function/function.proto

It has these top-level messages:
	CreateRequest
	CreateReply
	ListRequest
	ListReply
	DeleteRequest
*/
package function

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"
import functions "github.com/appcelerator/amp/data/functions"

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

type CreateRequest struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Image string `protobuf:"bytes,2,opt,name=image" json:"image,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRequest) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

type CreateReply struct {
	Function *functions.Function `protobuf:"bytes,1,opt,name=function" json:"function,omitempty"`
}

func (m *CreateReply) Reset()                    { *m = CreateReply{} }
func (m *CreateReply) String() string            { return proto.CompactTextString(m) }
func (*CreateReply) ProtoMessage()               {}
func (*CreateReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateReply) GetFunction() *functions.Function {
	if m != nil {
		return m.Function
	}
	return nil
}

type ListRequest struct {
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type ListReply struct {
	Functions []*functions.Function `protobuf:"bytes,1,rep,name=functions" json:"functions,omitempty"`
}

func (m *ListReply) Reset()                    { *m = ListReply{} }
func (m *ListReply) String() string            { return proto.CompactTextString(m) }
func (*ListReply) ProtoMessage()               {}
func (*ListReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ListReply) GetFunctions() []*functions.Function {
	if m != nil {
		return m.Functions
	}
	return nil
}

type DeleteRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *DeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "function.CreateRequest")
	proto.RegisterType((*CreateReply)(nil), "function.CreateReply")
	proto.RegisterType((*ListRequest)(nil), "function.ListRequest")
	proto.RegisterType((*ListReply)(nil), "function.ListReply")
	proto.RegisterType((*DeleteRequest)(nil), "function.DeleteRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Function service

type FunctionClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateReply, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}

type functionClient struct {
	cc *grpc.ClientConn
}

func NewFunctionClient(cc *grpc.ClientConn) FunctionClient {
	return &functionClient{cc}
}

func (c *functionClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateReply, error) {
	out := new(CreateReply)
	err := grpc.Invoke(ctx, "/function.Function/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *functionClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error) {
	out := new(ListReply)
	err := grpc.Invoke(ctx, "/function.Function/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *functionClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/function.Function/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Function service

type FunctionServer interface {
	Create(context.Context, *CreateRequest) (*CreateReply, error)
	List(context.Context, *ListRequest) (*ListReply, error)
	Delete(context.Context, *DeleteRequest) (*google_protobuf.Empty, error)
}

func RegisterFunctionServer(s *grpc.Server, srv FunctionServer) {
	s.RegisterService(&_Function_serviceDesc, srv)
}

func _Function_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FunctionServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/function.Function/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FunctionServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Function_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FunctionServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/function.Function/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FunctionServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Function_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FunctionServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/function.Function/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FunctionServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Function_serviceDesc = grpc.ServiceDesc{
	ServiceName: "function.Function",
	HandlerType: (*FunctionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Function_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Function_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Function_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/appcelerator/amp/api/rpc/function/function.proto",
}

func init() {
	proto.RegisterFile("github.com/appcelerator/amp/api/rpc/function/function.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xbb, 0x6e, 0x83, 0x30,
	0x14, 0x15, 0x69, 0x1a, 0x25, 0x37, 0x4a, 0x07, 0xa7, 0x8f, 0x88, 0xa5, 0x11, 0x53, 0x27, 0x5b,
	0x4d, 0x96, 0xa6, 0x95, 0xb2, 0xf4, 0x31, 0x75, 0xca, 0x1f, 0x18, 0xb8, 0x71, 0x2d, 0x01, 0x76,
	0xc1, 0x0c, 0x7c, 0x57, 0x7f, 0xb0, 0x02, 0x83, 0x21, 0xad, 0xda, 0x05, 0xdd, 0x63, 0xdf, 0xf3,
	0xc2, 0xf0, 0x24, 0xa4, 0xf9, 0x28, 0x43, 0x1a, 0xa9, 0x94, 0x71, 0xad, 0x23, 0x4c, 0x30, 0xe7,
	0x46, 0xe5, 0x8c, 0xa7, 0x9a, 0x71, 0x2d, 0x59, 0xae, 0x23, 0x76, 0x2c, 0xb3, 0xc8, 0x48, 0x95,
	0xb9, 0x81, 0xea, 0x5c, 0x19, 0x45, 0xa6, 0x1d, 0xf6, 0xb7, 0x03, 0x19, 0xa1, 0x12, 0x9e, 0x09,
	0xd6, 0xac, 0x84, 0xe5, 0x91, 0x69, 0x53, 0x69, 0x2c, 0x18, 0xa6, 0xda, 0x54, 0xf6, 0x6b, 0xe9,
	0xfe, 0xe3, 0x7f, 0xde, 0x31, 0x37, 0xdc, 0xf9, 0x15, 0xfd, 0x64, 0xb9, 0xc1, 0x0e, 0x16, 0xcf,
	0x39, 0x72, 0x83, 0x07, 0xfc, 0x2c, 0xb1, 0x30, 0x84, 0xc0, 0x38, 0xe3, 0x29, 0xae, 0xbc, 0xb5,
	0x77, 0x37, 0x3b, 0x34, 0x33, 0xb9, 0x84, 0x73, 0x99, 0x72, 0x81, 0xab, 0x51, 0x73, 0x68, 0x41,
	0xb0, 0x87, 0x79, 0x47, 0xd5, 0x49, 0x45, 0x18, 0xb8, 0x1a, 0x0d, 0x79, 0xbe, 0x59, 0xd2, 0xde,
	0xed, 0xad, 0x9d, 0x0e, 0x6e, 0x29, 0x58, 0xc0, 0xfc, 0x5d, 0x16, 0xa6, 0x35, 0x0e, 0xf6, 0x30,
	0xb3, 0xb0, 0x16, 0xbb, 0x87, 0x99, 0xe3, 0xae, 0xbc, 0xf5, 0xd9, 0x5f, 0x6a, 0xfd, 0x56, 0x70,
	0x0b, 0x8b, 0x17, 0x4c, 0xb0, 0x6f, 0x72, 0x01, 0x23, 0x19, 0xb7, 0x3d, 0x46, 0x32, 0xde, 0x7c,
	0x79, 0x30, 0xed, 0x88, 0xe4, 0x01, 0x26, 0x36, 0x3c, 0xb9, 0x71, 0xba, 0xf4, 0xe4, 0x4f, 0xf8,
	0x57, 0xbf, 0x2f, 0xea, 0x68, 0x1b, 0x18, 0xd7, 0x39, 0xc9, 0xe0, 0x7a, 0x50, 0xc3, 0x5f, 0xfe,
	0x3c, 0xae, 0x39, 0x3b, 0x98, 0xd8, 0x6c, 0x43, 0xb7, 0x93, 0xb4, 0xfe, 0x35, 0x15, 0x4a, 0x89,
	0x04, 0x69, 0xf7, 0xde, 0xf4, 0xb5, 0x7e, 0xe2, 0x70, 0xd2, 0xe0, 0xed, 0x77, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xda, 0x26, 0x6a, 0xb6, 0x61, 0x02, 0x00, 0x00,
}
