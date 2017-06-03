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
	ServiceListRequest
	ServiceListReply
	ServiceListEntry
	ServiceEntity
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

type ServiceListRequest struct {
}

func (m *ServiceListRequest) Reset()                    { *m = ServiceListRequest{} }
func (m *ServiceListRequest) String() string            { return proto.CompactTextString(m) }
func (*ServiceListRequest) ProtoMessage()               {}
func (*ServiceListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type ServiceListReply struct {
	Entries []*ServiceListEntry `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty"`
}

func (m *ServiceListReply) Reset()                    { *m = ServiceListReply{} }
func (m *ServiceListReply) String() string            { return proto.CompactTextString(m) }
func (*ServiceListReply) ProtoMessage()               {}
func (*ServiceListReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ServiceListReply) GetEntries() []*ServiceListEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type ServiceListEntry struct {
	Service    *ServiceEntity `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
	ReadyTasks int32          `protobuf:"varint,2,opt,name=ready_tasks,json=readyTasks" json:"ready_tasks,omitempty"`
	TotalTasks int32          `protobuf:"varint,3,opt,name=total_tasks,json=totalTasks" json:"total_tasks,omitempty"`
	Status     string         `protobuf:"bytes,4,opt,name=status" json:"status,omitempty"`
}

func (m *ServiceListEntry) Reset()                    { *m = ServiceListEntry{} }
func (m *ServiceListEntry) String() string            { return proto.CompactTextString(m) }
func (*ServiceListEntry) ProtoMessage()               {}
func (*ServiceListEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ServiceListEntry) GetService() *ServiceEntity {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *ServiceListEntry) GetReadyTasks() int32 {
	if m != nil {
		return m.ReadyTasks
	}
	return 0
}

func (m *ServiceListEntry) GetTotalTasks() int32 {
	if m != nil {
		return m.TotalTasks
	}
	return 0
}

func (m *ServiceListEntry) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type ServiceEntity struct {
	Id    string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Mode  string `protobuf:"bytes,3,opt,name=mode" json:"mode,omitempty"`
	Image string `protobuf:"bytes,4,opt,name=image" json:"image,omitempty"`
	Tag   string `protobuf:"bytes,5,opt,name=tag" json:"tag,omitempty"`
}

func (m *ServiceEntity) Reset()                    { *m = ServiceEntity{} }
func (m *ServiceEntity) String() string            { return proto.CompactTextString(m) }
func (*ServiceEntity) ProtoMessage()               {}
func (*ServiceEntity) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ServiceEntity) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ServiceEntity) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ServiceEntity) GetMode() string {
	if m != nil {
		return m.Mode
	}
	return ""
}

func (m *ServiceEntity) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *ServiceEntity) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func init() {
	proto.RegisterType((*TasksRequest)(nil), "service.TasksRequest")
	proto.RegisterType((*Task)(nil), "service.Task")
	proto.RegisterType((*TasksReply)(nil), "service.TasksReply")
	proto.RegisterType((*ServiceListRequest)(nil), "service.ServiceListRequest")
	proto.RegisterType((*ServiceListReply)(nil), "service.ServiceListReply")
	proto.RegisterType((*ServiceListEntry)(nil), "service.ServiceListEntry")
	proto.RegisterType((*ServiceEntity)(nil), "service.ServiceEntity")
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
	ListService(ctx context.Context, in *ServiceListRequest, opts ...grpc.CallOption) (*ServiceListReply, error)
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

func (c *serviceClient) ListService(ctx context.Context, in *ServiceListRequest, opts ...grpc.CallOption) (*ServiceListReply, error) {
	out := new(ServiceListReply)
	err := grpc.Invoke(ctx, "/service.Service/ListService", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Service service

type ServiceServer interface {
	Tasks(context.Context, *TasksRequest) (*TasksReply, error)
	ListService(context.Context, *ServiceListRequest) (*ServiceListReply, error)
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

func _Service_ListService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ListService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Service/ListService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ListService(ctx, req.(*ServiceListRequest))
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
		{
			MethodName: "ListService",
			Handler:    _Service_ListService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/appcelerator/amp/api/rpc/service/service.proto",
}

func init() {
	proto.RegisterFile("github.com/appcelerator/amp/api/rpc/service/service.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0x4b, 0x6e, 0xdb, 0x30,
	0x10, 0x85, 0x6c, 0xcb, 0x46, 0xc6, 0x76, 0x10, 0xb0, 0xae, 0xab, 0xa4, 0xbf, 0x40, 0xd9, 0x64,
	0x53, 0xab, 0x71, 0x56, 0x3d, 0x40, 0x50, 0x04, 0xe8, 0x4a, 0xee, 0xae, 0x0b, 0x83, 0xb1, 0x08,
	0x95, 0xa8, 0x24, 0xaa, 0xe4, 0x38, 0x80, 0x50, 0x14, 0x28, 0x7a, 0x85, 0xee, 0x7b, 0x91, 0x1e,
	0xa3, 0x57, 0xe8, 0x41, 0x0a, 0x0e, 0xc9, 0x24, 0x4e, 0xb2, 0x32, 0xe7, 0xbd, 0xe7, 0x79, 0x9c,
	0x79, 0x14, 0xbc, 0x2b, 0x25, 0x7e, 0xde, 0x5e, 0x2d, 0x36, 0xaa, 0xce, 0x78, 0xdb, 0x6e, 0x44,
	0x25, 0x34, 0x47, 0xa5, 0x33, 0x5e, 0xb7, 0x19, 0x6f, 0x65, 0xa6, 0xdb, 0x4d, 0x66, 0x84, 0xbe,
	0x96, 0x1b, 0x11, 0x7e, 0x17, 0xad, 0x56, 0xa8, 0xd8, 0xc8, 0x97, 0x47, 0x2f, 0x4a, 0xa5, 0xca,
	0x4a, 0x90, 0x9c, 0x37, 0x8d, 0x42, 0x8e, 0x52, 0x35, 0xc6, 0xc9, 0xd2, 0x37, 0x30, 0xf9, 0xc8,
	0xcd, 0x17, 0x93, 0x8b, 0xaf, 0x5b, 0x61, 0x90, 0xbd, 0x04, 0xf0, 0x7f, 0x5c, 0xcb, 0x22, 0x89,
	0x8e, 0xa3, 0xd3, 0xbd, 0x7c, 0xcf, 0x23, 0x97, 0x45, 0xfa, 0x23, 0x82, 0x81, 0xd5, 0xb3, 0x7d,
	0xe8, 0xdd, 0xf0, 0x3d, 0x59, 0xb0, 0x19, 0xc4, 0xb2, 0xe6, 0xa5, 0x48, 0x7a, 0x04, 0xb9, 0xc2,
	0xa2, 0x06, 0x39, 0x8a, 0xa4, 0xef, 0x50, 0x2a, 0xd8, 0x09, 0x4c, 0x0b, 0x61, 0xa4, 0x16, 0xc5,
	0xda, 0xb1, 0x03, 0x62, 0x27, 0x1e, 0x5c, 0x91, 0xe8, 0x19, 0x8c, 0x1a, 0x55, 0xd0, 0x2d, 0x62,
	0xa2, 0x87, 0xb6, 0xbc, 0x2c, 0xd2, 0x33, 0x00, 0x7f, 0xe3, 0xb6, 0xea, 0xd8, 0x09, 0xc4, 0x68,
	0xab, 0x24, 0x3a, 0xee, 0x9f, 0x8e, 0x97, 0xd3, 0x45, 0xd8, 0x82, 0xd5, 0xe4, 0x8e, 0x4b, 0x67,
	0xc0, 0x56, 0x0e, 0xfe, 0x20, 0x0d, 0xfa, 0x51, 0xd3, 0xf7, 0x70, 0xb0, 0x83, 0xda, 0x76, 0xe7,
	0x30, 0x12, 0x0d, 0x6a, 0x29, 0x42, 0xc3, 0xc3, 0x9b, 0x86, 0x77, 0xb4, 0x17, 0x0d, 0xea, 0x2e,
	0x0f, 0xca, 0xf4, 0x77, 0xb4, 0xd3, 0x89, 0x58, 0xf6, 0x16, 0x42, 0x02, 0xb4, 0xa5, 0xf1, 0x72,
	0x7e, 0xbf, 0xd3, 0x45, 0x83, 0x12, 0xbb, 0x3c, 0xc8, 0xd8, 0x6b, 0x18, 0x6b, 0xc1, 0x8b, 0x6e,
	0xed, 0x06, 0xb2, 0x8b, 0x8c, 0x73, 0x20, 0x88, 0x06, 0xb6, 0x02, 0x54, 0xc8, 0x2b, 0x2f, 0xe8,
	0x3b, 0x01, 0x41, 0x4e, 0x30, 0x87, 0xa1, 0x5d, 0xe8, 0xd6, 0xf8, 0x8d, 0xfa, 0x2a, 0x55, 0x30,
	0xdd, 0xf1, 0x7c, 0x90, 0x1e, 0x83, 0x41, 0xc3, 0xeb, 0x10, 0x1e, 0x9d, 0x2d, 0x56, 0xab, 0x22,
	0x44, 0x47, 0xe7, 0xdb, 0x94, 0x07, 0x77, 0x53, 0x3e, 0x80, 0x3e, 0xf2, 0xd2, 0xc7, 0x64, 0x8f,
	0xcb, 0x3f, 0x11, 0x8c, 0xbc, 0x23, 0x5b, 0x41, 0xec, 0x6e, 0xf7, 0x74, 0x27, 0x9b, 0xf0, 0xe2,
	0x8e, 0x9e, 0xdc, 0x87, 0xdb, 0xaa, 0x4b, 0x5f, 0xfd, 0xfc, 0xfb, 0xef, 0x57, 0x2f, 0x61, 0xf3,
	0xec, 0xfa, 0x2c, 0xa3, 0x71, 0xb3, 0x6f, 0xb7, 0xef, 0xf2, 0x3b, 0xfb, 0x04, 0x63, 0xbb, 0xea,
	0xe0, 0xf1, 0xfc, 0xb1, 0x94, 0x82, 0xc1, 0xe1, 0xe3, 0xa4, 0xb5, 0x99, 0x91, 0xcd, 0x3e, 0x9b,
	0x58, 0x1b, 0xaf, 0x32, 0x57, 0x43, 0xfa, 0x34, 0xce, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0xef,
	0x67, 0x34, 0xd9, 0x7e, 0x03, 0x00, 0x00,
}
