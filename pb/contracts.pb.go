// Code generated by protoc-gen-go.
// source: contracts.proto
// DO NOT EDIT!

/*
Package deployment is a generated protocol buffer package.

It is generated from these files:
	contracts.proto

It has these top-level messages:
	Resource
	DeployRequest
	DeployResponse
*/
package deployment

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

type Resource struct {
	Bucket string `protobuf:"bytes,1,opt,name=bucket" json:"bucket,omitempty"`
	Key    string `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
}

func (m *Resource) Reset()                    { *m = Resource{} }
func (m *Resource) String() string            { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()               {}
func (*Resource) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Resource) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *Resource) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type DeployRequest struct {
	Resource *Resource `protobuf:"bytes,1,opt,name=resource" json:"resource,omitempty"`
}

func (m *DeployRequest) Reset()                    { *m = DeployRequest{} }
func (m *DeployRequest) String() string            { return proto.CompactTextString(m) }
func (*DeployRequest) ProtoMessage()               {}
func (*DeployRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DeployRequest) GetResource() *Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

type DeployResponse struct {
	Status   string    `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Resource *Resource `protobuf:"bytes,2,opt,name=resource" json:"resource,omitempty"`
}

func (m *DeployResponse) Reset()                    { *m = DeployResponse{} }
func (m *DeployResponse) String() string            { return proto.CompactTextString(m) }
func (*DeployResponse) ProtoMessage()               {}
func (*DeployResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DeployResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *DeployResponse) GetResource() *Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func init() {
	proto.RegisterType((*Resource)(nil), "deployment.Resource")
	proto.RegisterType((*DeployRequest)(nil), "deployment.DeployRequest")
	proto.RegisterType((*DeployResponse)(nil), "deployment.DeployResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Deployer service

type DeployerClient interface {
	Deploy(ctx context.Context, in *DeployRequest, opts ...grpc.CallOption) (Deployer_DeployClient, error)
}

type deployerClient struct {
	cc *grpc.ClientConn
}

func NewDeployerClient(cc *grpc.ClientConn) DeployerClient {
	return &deployerClient{cc}
}

func (c *deployerClient) Deploy(ctx context.Context, in *DeployRequest, opts ...grpc.CallOption) (Deployer_DeployClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Deployer_serviceDesc.Streams[0], c.cc, "/deployment.Deployer/Deploy", opts...)
	if err != nil {
		return nil, err
	}
	x := &deployerDeployClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Deployer_DeployClient interface {
	Recv() (*DeployResponse, error)
	grpc.ClientStream
}

type deployerDeployClient struct {
	grpc.ClientStream
}

func (x *deployerDeployClient) Recv() (*DeployResponse, error) {
	m := new(DeployResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Deployer service

type DeployerServer interface {
	Deploy(*DeployRequest, Deployer_DeployServer) error
}

func RegisterDeployerServer(s *grpc.Server, srv DeployerServer) {
	s.RegisterService(&_Deployer_serviceDesc, srv)
}

func _Deployer_Deploy_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DeployRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DeployerServer).Deploy(m, &deployerDeployServer{stream})
}

type Deployer_DeployServer interface {
	Send(*DeployResponse) error
	grpc.ServerStream
}

type deployerDeployServer struct {
	grpc.ServerStream
}

func (x *deployerDeployServer) Send(m *DeployResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Deployer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "deployment.Deployer",
	HandlerType: (*DeployerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Deploy",
			Handler:       _Deployer_Deploy_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "contracts.proto",
}

func init() { proto.RegisterFile("contracts.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0xce, 0xcf, 0x2b,
	0x29, 0x4a, 0x4c, 0x2e, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4a, 0x49, 0x2d,
	0xc8, 0xc9, 0xaf, 0xcc, 0x4d, 0xcd, 0x2b, 0x51, 0x32, 0xe1, 0xe2, 0x08, 0x4a, 0x2d, 0xce, 0x2f,
	0x2d, 0x4a, 0x4e, 0x15, 0x12, 0xe3, 0x62, 0x4b, 0x2a, 0x4d, 0xce, 0x4e, 0x2d, 0x91, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf2, 0x84, 0x04, 0xb8, 0x98, 0xb3, 0x53, 0x2b, 0x25, 0x98, 0xc0,
	0x82, 0x20, 0xa6, 0x92, 0x23, 0x17, 0xaf, 0x0b, 0xd8, 0x8c, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2,
	0x12, 0x21, 0x03, 0x2e, 0x8e, 0x22, 0xa8, 0x31, 0x60, 0xcd, 0xdc, 0x46, 0x22, 0x7a, 0x08, 0x5b,
	0xf4, 0x60, 0x56, 0x04, 0xc1, 0x55, 0x29, 0x45, 0x71, 0xf1, 0xc1, 0x8c, 0x28, 0x2e, 0xc8, 0xcf,
	0x2b, 0x06, 0x5b, 0x5f, 0x5c, 0x92, 0x58, 0x52, 0x5a, 0x0c, 0xb3, 0x1e, 0xc2, 0x43, 0x31, 0x9b,
	0x89, 0x18, 0xb3, 0x8d, 0x7c, 0xb9, 0x38, 0x20, 0x66, 0xa7, 0x16, 0x09, 0x39, 0x72, 0xb1, 0x41,
	0xd8, 0x42, 0x92, 0xc8, 0xba, 0x50, 0x9c, 0x2f, 0x25, 0x85, 0x4d, 0x0a, 0xe2, 0x2c, 0x03, 0xc6,
	0x24, 0x36, 0x70, 0xb0, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xda, 0x4b, 0xae, 0x65, 0x49,
	0x01, 0x00, 0x00,
}
