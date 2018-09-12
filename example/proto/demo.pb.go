// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example/proto/demo.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	example/proto/demo.proto

It has these top-level messages:
	Demo
	Custom
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/infobloxopen/protoc-gen-preprocess/options"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type Demo struct {
	// Also it is possible to specify additional method on field level
	PreprocessedField string `protobuf:"bytes,1,opt,name=preprocessedField" json:"preprocessedField,omitempty"`
	// Preprocessor automatically checks if field is repeated and generates methods accordingly
	PreprocessedRepeatedField []string `protobuf:"bytes,2,rep,name=preprocessedRepeatedField" json:"preprocessedRepeatedField,omitempty"`
	// If a field does not fit preprocess method, it is just ignored
	Ignored int32 `protobuf:"varint,3,opt,name=ignored" json:"ignored,omitempty"`
}

func (m *Demo) Reset()                    { *m = Demo{} }
func (m *Demo) String() string            { return proto1.CompactTextString(m) }
func (*Demo) ProtoMessage()               {}
func (*Demo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Demo) GetPreprocessedField() string {
	if m != nil {
		return m.PreprocessedField
	}
	return ""
}

func (m *Demo) GetPreprocessedRepeatedField() []string {
	if m != nil {
		return m.PreprocessedRepeatedField
	}
	return nil
}

func (m *Demo) GetIgnored() int32 {
	if m != nil {
		return m.Ignored
	}
	return 0
}

// This message left as is to show that we can provide our own preprocessors
type Custom struct {
	DoItYourself string `protobuf:"bytes,1,opt,name=doItYourself" json:"doItYourself,omitempty"`
}

func (m *Custom) Reset()                    { *m = Custom{} }
func (m *Custom) String() string            { return proto1.CompactTextString(m) }
func (*Custom) ProtoMessage()               {}
func (*Custom) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Custom) GetDoItYourself() string {
	if m != nil {
		return m.DoItYourself
	}
	return ""
}

func init() {
	proto1.RegisterType((*Demo)(nil), "proto.Demo")
	proto1.RegisterType((*Custom)(nil), "proto.Custom")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DemoService service

type DemoServiceClient interface {
	Echo(ctx context.Context, in *Demo, opts ...grpc.CallOption) (*Demo, error)
}

type demoServiceClient struct {
	cc *grpc.ClientConn
}

func NewDemoServiceClient(cc *grpc.ClientConn) DemoServiceClient {
	return &demoServiceClient{cc}
}

func (c *demoServiceClient) Echo(ctx context.Context, in *Demo, opts ...grpc.CallOption) (*Demo, error) {
	out := new(Demo)
	err := grpc.Invoke(ctx, "/proto.DemoService/Echo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DemoService service

type DemoServiceServer interface {
	Echo(context.Context, *Demo) (*Demo, error)
}

func RegisterDemoServiceServer(s *grpc.Server, srv DemoServiceServer) {
	s.RegisterService(&_DemoService_serviceDesc, srv)
}

func _DemoService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Demo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DemoService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoServiceServer).Echo(ctx, req.(*Demo))
	}
	return interceptor(ctx, in, info, handler)
}

var _DemoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.DemoService",
	HandlerType: (*DemoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _DemoService_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example/proto/demo.proto",
}

func init() { proto1.RegisterFile("example/proto/demo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0x41, 0x4e, 0x02, 0x31,
	0x14, 0x86, 0x53, 0x86, 0x22, 0x3c, 0x5c, 0x68, 0x57, 0x48, 0x5c, 0x90, 0x71, 0x43, 0x8c, 0xd0,
	0x04, 0x13, 0x17, 0xc4, 0x8d, 0xa8, 0x24, 0x6e, 0xc7, 0x95, 0xcb, 0xa1, 0xf3, 0x18, 0x9a, 0xcc,
	0xf4, 0x35, 0x9d, 0x62, 0x58, 0x7b, 0x05, 0xaf, 0xe1, 0x9a, 0x05, 0xd7, 0xf0, 0x0a, 0x1e, 0xc4,
	0x30, 0x13, 0x04, 0x63, 0x5c, 0xb5, 0xfd, 0xbe, 0xff, 0x6f, 0xde, 0x83, 0x0e, 0xae, 0xe2, 0xdc,
	0x66, 0x28, 0xad, 0x23, 0x4f, 0x32, 0xc1, 0x9c, 0x86, 0xe5, 0x55, 0xf0, 0xf2, 0xe8, 0x4e, 0x53,
	0xed, 0x17, 0xcb, 0xd9, 0x50, 0x51, 0x2e, 0xb5, 0x99, 0xd3, 0x2c, 0xa3, 0x15, 0x59, 0x34, 0x55,
	0x41, 0x0d, 0x52, 0x34, 0x03, 0xeb, 0xd0, 0x3a, 0x52, 0x58, 0x14, 0x92, 0xac, 0xd7, 0x64, 0x0a,
	0xb9, 0x47, 0xd5, 0x77, 0xdd, 0xf3, 0x94, 0x28, 0xcd, 0x50, 0xc6, 0x56, 0xcb, 0xd8, 0x18, 0xf2,
	0x71, 0x19, 0xac, 0x6c, 0xf8, 0xc1, 0xa0, 0xfe, 0x80, 0x39, 0x89, 0x1b, 0x38, 0xdd, 0x57, 0x31,
	0x99, 0x6a, 0xcc, 0x92, 0x0e, 0xeb, 0xb1, 0x7e, 0x6b, 0xd2, 0xdc, 0xac, 0x79, 0x1d, 0x6a, 0xcd,
	0x20, 0xfa, 0x1b, 0x11, 0xb7, 0x70, 0x76, 0x08, 0x23, 0xb4, 0x18, 0xfb, 0x5d, 0xbf, 0xd6, 0x0b,
	0xfa, 0xad, 0xe8, 0xff, 0x80, 0xb8, 0x80, 0x23, 0x9d, 0x1a, 0x72, 0x98, 0x74, 0x82, 0x1e, 0xeb,
	0xf3, 0x49, 0x6b, 0xb3, 0xe6, 0x1c, 0x02, 0x60, 0x41, 0xb4, 0x33, 0xe3, 0x1f, 0xc6, 0xc2, 0x2b,
	0x68, 0xdc, 0x2f, 0x0b, 0x4f, 0xb9, 0x08, 0xe1, 0x38, 0xa1, 0x27, 0xff, 0x42, 0x4b, 0x57, 0x60,
	0x36, 0xaf, 0x46, 0x8d, 0x7e, 0xb1, 0xd1, 0x1d, 0xb4, 0xb7, 0xbb, 0x3d, 0xa3, 0x7b, 0xd5, 0x0a,
	0xc5, 0x08, 0xea, 0x8f, 0x6a, 0x41, 0xa2, 0x5d, 0xed, 0x3e, 0xdc, 0xba, 0xee, 0xe1, 0x23, 0x3c,
	0x79, 0xfb, 0xfc, 0x7a, 0xaf, 0x41, 0xc8, 0x25, 0xaa, 0x05, 0x8d, 0xd9, 0xe5, 0xac, 0x51, 0xda,
	0xeb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xca, 0x68, 0xa3, 0xec, 0xaf, 0x01, 0x00, 0x00,
}
