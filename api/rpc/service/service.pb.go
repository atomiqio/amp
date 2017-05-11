// Code generated by protoc-gen-go.
// source: github.com/appcelerator/amp/api/rpc/service/service.proto
// DO NOT EDIT!

/*
Package service is a generated protocol buffer package.

It is generated from these files:
	github.com/appcelerator/amp/api/rpc/service/service.proto

It has these top-level messages:
	TasksRequest
	Task
	TasksReply
*/
package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
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

type TasksRequest struct {
	ServiceId string `protobuf:"bytes,1,opt,name=service_id,json=serviceId" json:"service_id,omitempty"`
}

func (m *TasksRequest) Reset()                    { *m = TasksRequest{} }
func (m *TasksRequest) String() string            { return proto.CompactTextString(m) }
func (*TasksRequest) ProtoMessage()               {}
func (*TasksRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TasksRequest) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

type Task struct {
	Id           string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Image        string `protobuf:"bytes,2,opt,name=image" json:"image,omitempty"`
	State        string `protobuf:"bytes,3,opt,name=state" json:"state,omitempty"`
	DesiredState string `protobuf:"bytes,4,opt,name=desired_state,json=desiredState" json:"desired_state,omitempty"`
	NodeId       string `protobuf:"bytes,5,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Task) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Task) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Task) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Task) GetDesiredState() string {
	if m != nil {
		return m.DesiredState
	}
	return ""
}

func (m *Task) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

type TasksReply struct {
	Tasks []*Task `protobuf:"bytes,1,rep,name=tasks" json:"tasks,omitempty"`
}

func (m *TasksReply) Reset()                    { *m = TasksReply{} }
func (m *TasksReply) String() string            { return proto.CompactTextString(m) }
func (*TasksReply) ProtoMessage()               {}
func (*TasksReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TasksReply) GetTasks() []*Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func init() {
	proto.RegisterType((*TasksRequest)(nil), "service.TasksRequest")
	proto.RegisterType((*Task)(nil), "service.Task")
	proto.RegisterType((*TasksReply)(nil), "service.TasksReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Service service

type ServiceClient interface {
	Tasks(ctx context.Context, in *TasksRequest, opts ...grpc.CallOption) (*TasksReply, error)
}

type serviceClient struct {
	cc *grpc.ClientConn
}

func NewServiceClient(cc *grpc.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Tasks(ctx context.Context, in *TasksRequest, opts ...grpc.CallOption) (*TasksReply, error) {
	out := new(TasksReply)
	err := grpc.Invoke(ctx, "/service.Service/Tasks", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Service service

type ServiceServer interface {
	Tasks(context.Context, *TasksRequest) (*TasksReply, error)
}

func RegisterServiceServer(s *grpc.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_Tasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Tasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Service/Tasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Tasks(ctx, req.(*TasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Tasks",
			Handler:    _Service_Tasks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/appcelerator/amp/api/rpc/service/service.proto",
}

func init() {
	proto.RegisterFile("github.com/appcelerator/amp/api/rpc/service/service.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x4d, 0x4e, 0xc3, 0x30,
	0x10, 0x85, 0x95, 0xb4, 0x69, 0xd5, 0xa1, 0x65, 0x61, 0xfe, 0xa2, 0x0a, 0x50, 0x95, 0x6e, 0xba,
	0xa1, 0x56, 0xcb, 0x8a, 0x23, 0x74, 0xdb, 0xb0, 0xa6, 0x72, 0xe3, 0x51, 0xb0, 0x48, 0x62, 0x63,
	0xbb, 0x95, 0x2a, 0x84, 0x84, 0xb8, 0x02, 0x47, 0xe3, 0x0a, 0x1c, 0x04, 0xc5, 0x4e, 0x40, 0xb0,
	0x8a, 0xde, 0xf7, 0xbe, 0x89, 0xc6, 0x03, 0x77, 0xb9, 0xb0, 0x8f, 0xbb, 0xed, 0x3c, 0x93, 0x25,
	0x65, 0x4a, 0x65, 0x58, 0xa0, 0x66, 0x56, 0x6a, 0xca, 0x4a, 0x45, 0x99, 0x12, 0x54, 0xab, 0x8c,
	0x1a, 0xd4, 0x7b, 0x91, 0x61, 0xfb, 0x9d, 0x2b, 0x2d, 0xad, 0x24, 0xfd, 0x26, 0x8e, 0x2f, 0x73,
	0x29, 0xf3, 0x02, 0x9d, 0xce, 0xaa, 0x4a, 0x5a, 0x66, 0x85, 0xac, 0x8c, 0xd7, 0x92, 0x1b, 0x18,
	0xde, 0x33, 0xf3, 0x64, 0xd6, 0xf8, 0xbc, 0x43, 0x63, 0xc9, 0x15, 0x40, 0x33, 0xb8, 0x11, 0x3c,
	0x0e, 0x26, 0xc1, 0x6c, 0xb0, 0x1e, 0x34, 0x64, 0xc5, 0x93, 0xb7, 0x00, 0xba, 0xb5, 0x4f, 0x8e,
	0x21, 0xfc, 0xe9, 0x43, 0xc1, 0xc9, 0x29, 0x44, 0xa2, 0x64, 0x39, 0xc6, 0xa1, 0x43, 0x3e, 0xd4,
	0xd4, 0x58, 0x66, 0x31, 0xee, 0x78, 0xea, 0x02, 0x99, 0xc2, 0x88, 0xa3, 0x11, 0x1a, 0xf9, 0xc6,
	0xb7, 0x5d, 0xd7, 0x0e, 0x1b, 0x98, 0x3a, 0xe9, 0x02, 0xfa, 0x95, 0xe4, 0x6e, 0x8b, 0xc8, 0xd5,
	0xbd, 0x3a, 0xae, 0x78, 0xb2, 0x00, 0x68, 0x36, 0x56, 0xc5, 0x81, 0x4c, 0x21, 0xb2, 0x75, 0x8a,
	0x83, 0x49, 0x67, 0x76, 0xb4, 0x1c, 0xcd, 0xdb, 0x2b, 0xd4, 0xce, 0xda, 0x77, 0xcb, 0x07, 0xe8,
	0xa7, 0x1e, 0x93, 0x14, 0x22, 0x37, 0x4d, 0xce, 0xfe, 0x98, 0xed, 0xfb, 0xc7, 0x27, 0xff, 0xb1,
	0x2a, 0x0e, 0xc9, 0xf5, 0xfb, 0xe7, 0xd7, 0x47, 0x18, 0x93, 0x73, 0xba, 0x5f, 0x50, 0xf7, 0x4b,
	0xfa, 0xf2, 0x7b, 0xa5, 0xd7, 0x6d, 0xcf, 0xdd, 0xf2, 0xf6, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x59,
	0x9d, 0x9c, 0x17, 0xaf, 0x01, 0x00, 0x00,
}
