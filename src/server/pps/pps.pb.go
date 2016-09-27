// Code generated by protoc-gen-go.
// source: server/pps/pps.proto
// DO NOT EDIT!

/*
Package pps is a generated protocol buffer package.

It is generated from these files:
	server/pps/pps.proto

It has these top-level messages:
	StartJobRequest
	StartJobResponse
	FinishJobRequest
*/
package pps

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "go.pedge.io/pb/go/google/protobuf"
import _ "github.com/pachyderm/pachyderm/src/server/pfs/fuse"
import pfs "github.com/pachyderm/pachyderm/src/client/pfs"
import pachyderm_pps "github.com/pachyderm/pachyderm/src/client/pps"

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

type StartJobRequest struct {
	Job *pachyderm_pps.Job `protobuf:"bytes,1,opt,name=job" json:"job,omitempty"`
}

func (m *StartJobRequest) Reset()                    { *m = StartJobRequest{} }
func (m *StartJobRequest) String() string            { return proto.CompactTextString(m) }
func (*StartJobRequest) ProtoMessage()               {}
func (*StartJobRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StartJobRequest) GetJob() *pachyderm_pps.Job {
	if m != nil {
		return m.Job
	}
	return nil
}

type StartJobResponse struct {
	Transform *pachyderm_pps.Transform `protobuf:"bytes,1,opt,name=transform" json:"transform,omitempty"`
	View      *pfs.View                `protobuf:"bytes,4,opt,name=view" json:"view,omitempty"`
	PodIndex  uint64                   `protobuf:"varint,3,opt,name=pod_index,json=podIndex" json:"pod_index,omitempty"`
}

func (m *StartJobResponse) Reset()                    { *m = StartJobResponse{} }
func (m *StartJobResponse) String() string            { return proto.CompactTextString(m) }
func (*StartJobResponse) ProtoMessage()               {}
func (*StartJobResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *StartJobResponse) GetTransform() *pachyderm_pps.Transform {
	if m != nil {
		return m.Transform
	}
	return nil
}

func (m *StartJobResponse) GetView() *pfs.View {
	if m != nil {
		return m.View
	}
	return nil
}

type FinishJobRequest struct {
	Job      *pachyderm_pps.Job `protobuf:"bytes,1,opt,name=job" json:"job,omitempty"`
	Success  bool               `protobuf:"varint,2,opt,name=success" json:"success,omitempty"`
	PodIndex uint64             `protobuf:"varint,3,opt,name=pod_index,json=podIndex" json:"pod_index,omitempty"`
}

func (m *FinishJobRequest) Reset()                    { *m = FinishJobRequest{} }
func (m *FinishJobRequest) String() string            { return proto.CompactTextString(m) }
func (*FinishJobRequest) ProtoMessage()               {}
func (*FinishJobRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *FinishJobRequest) GetJob() *pachyderm_pps.Job {
	if m != nil {
		return m.Job
	}
	return nil
}

func init() {
	proto.RegisterType((*StartJobRequest)(nil), "pachyderm.pps.StartJobRequest")
	proto.RegisterType((*StartJobResponse)(nil), "pachyderm.pps.StartJobResponse")
	proto.RegisterType((*FinishJobRequest)(nil), "pachyderm.pps.FinishJobRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for InternalJobAPI service

type InternalJobAPIClient interface {
	StartJob(ctx context.Context, in *StartJobRequest, opts ...grpc.CallOption) (*StartJobResponse, error)
	FinishJob(ctx context.Context, in *FinishJobRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}

type internalJobAPIClient struct {
	cc *grpc.ClientConn
}

func NewInternalJobAPIClient(cc *grpc.ClientConn) InternalJobAPIClient {
	return &internalJobAPIClient{cc}
}

func (c *internalJobAPIClient) StartJob(ctx context.Context, in *StartJobRequest, opts ...grpc.CallOption) (*StartJobResponse, error) {
	out := new(StartJobResponse)
	err := grpc.Invoke(ctx, "/pachyderm.pps.InternalJobAPI/StartJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalJobAPIClient) FinishJob(ctx context.Context, in *FinishJobRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/pachyderm.pps.InternalJobAPI/FinishJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for InternalJobAPI service

type InternalJobAPIServer interface {
	StartJob(context.Context, *StartJobRequest) (*StartJobResponse, error)
	FinishJob(context.Context, *FinishJobRequest) (*google_protobuf.Empty, error)
}

func RegisterInternalJobAPIServer(s *grpc.Server, srv InternalJobAPIServer) {
	s.RegisterService(&_InternalJobAPI_serviceDesc, srv)
}

func _InternalJobAPI_StartJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalJobAPIServer).StartJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pachyderm.pps.InternalJobAPI/StartJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalJobAPIServer).StartJob(ctx, req.(*StartJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalJobAPI_FinishJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinishJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalJobAPIServer).FinishJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pachyderm.pps.InternalJobAPI/FinishJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalJobAPIServer).FinishJob(ctx, req.(*FinishJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _InternalJobAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pachyderm.pps.InternalJobAPI",
	HandlerType: (*InternalJobAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartJob",
			Handler:    _InternalJobAPI_StartJob_Handler,
		},
		{
			MethodName: "FinishJob",
			Handler:    _InternalJobAPI_FinishJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("server/pps/pps.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x92, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0xc7, 0x1b, 0x5b, 0xb5, 0x59, 0x51, 0xcb, 0x52, 0x24, 0xa4, 0xa8, 0x25, 0x78, 0xe8, 0x69,
	0x03, 0x15, 0xf4, 0xac, 0x60, 0xa1, 0x05, 0x41, 0xa2, 0x78, 0xf0, 0x22, 0xf9, 0x98, 0xb4, 0x91,
	0x36, 0x3b, 0xee, 0x6c, 0x5a, 0xfb, 0x02, 0xbe, 0x8a, 0xaf, 0x29, 0x49, 0x1b, 0xab, 0x01, 0x05,
	0x0f, 0xbb, 0x30, 0xf3, 0xfb, 0xcf, 0x27, 0xc3, 0xda, 0x04, 0x6a, 0x0e, 0xca, 0x45, 0xa4, 0xfc,
	0x09, 0x54, 0x52, 0x4b, 0xbe, 0x8f, 0x7e, 0x38, 0x59, 0x46, 0xa0, 0x66, 0x02, 0x91, 0xec, 0xce,
	0x58, 0xca, 0xf1, 0x14, 0xdc, 0x02, 0x06, 0x59, 0xec, 0xc2, 0x0c, 0xf5, 0x72, 0xa5, 0xb5, 0xed,
	0x32, 0x43, 0x4c, 0x6e, 0x9c, 0x11, 0x14, 0xdf, 0x9a, 0xb5, 0xc3, 0x69, 0x02, 0xa9, 0x2e, 0x18,
	0xc6, 0x54, 0xf5, 0x7e, 0xaf, 0xe9, 0x5c, 0xb2, 0xc3, 0x7b, 0xed, 0x2b, 0x3d, 0x92, 0x81, 0x07,
	0xaf, 0x19, 0x90, 0xe6, 0x67, 0xac, 0xfe, 0x22, 0x03, 0xcb, 0xe8, 0x1a, 0xbd, 0xbd, 0x3e, 0x17,
	0x3f, 0x9a, 0x12, 0xb9, 0x2e, 0xc7, 0xce, 0xbb, 0xc1, 0x5a, 0x9b, 0x48, 0x42, 0x99, 0x12, 0xf0,
	0x0b, 0x66, 0x6a, 0xe5, 0xa7, 0x14, 0x4b, 0x35, 0x5b, 0x27, 0xb0, 0x2a, 0x09, 0x1e, 0x4a, 0xee,
	0x6d, 0xa4, 0xfc, 0x98, 0x35, 0xe6, 0x09, 0x2c, 0xac, 0x46, 0x11, 0x62, 0x8a, 0xbc, 0xeb, 0xc7,
	0x04, 0x16, 0x5e, 0xe1, 0xe6, 0x1d, 0x66, 0xa2, 0x8c, 0x9e, 0x93, 0x34, 0x82, 0x37, 0xab, 0xde,
	0x35, 0x7a, 0x0d, 0xaf, 0x89, 0x32, 0x1a, 0xe6, 0xb6, 0x23, 0x59, 0x6b, 0x90, 0xa4, 0x09, 0x4d,
	0xfe, 0x3b, 0x02, 0xb7, 0xd8, 0x2e, 0x65, 0x61, 0x08, 0x44, 0xd6, 0x56, 0xd7, 0xe8, 0x35, 0xbd,
	0xd2, 0xfc, 0xb3, 0x60, 0xff, 0xc3, 0x60, 0x07, 0xc3, 0x54, 0x83, 0x4a, 0xfd, 0xe9, 0x48, 0x06,
	0x57, 0x77, 0x43, 0x7e, 0xcb, 0x9a, 0xe5, 0x2e, 0xf8, 0x49, 0xa5, 0x5c, 0x65, 0xbd, 0xf6, 0xe9,
	0xaf, 0x7c, 0xb5, 0x44, 0xa7, 0xc6, 0x07, 0xcc, 0xfc, 0x1a, 0x89, 0x57, 0xf5, 0xd5, 0x61, 0xed,
	0x23, 0xb1, 0x3a, 0x14, 0x51, 0x1e, 0x8a, 0xb8, 0xc9, 0x0f, 0xc5, 0xa9, 0x5d, 0x6f, 0x3f, 0xd5,
	0x11, 0x29, 0xd8, 0x29, 0xc0, 0xf9, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x98, 0xca, 0xc4, 0xee,
	0x76, 0x02, 0x00, 0x00,
}
