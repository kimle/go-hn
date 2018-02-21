// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go-hackernews.proto

/*
Package gohackernews is a generated protocol buffer package.

It is generated from these files:
	go-hackernews.proto

It has these top-level messages:
	Story
	TopStories
	Stories
*/
package gohackernews

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

type Story struct {
	Id    int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Url   string `protobuf:"bytes,3,opt,name=url" json:"url,omitempty"`
}

func (m *Story) Reset()                    { *m = Story{} }
func (m *Story) String() string            { return proto.CompactTextString(m) }
func (*Story) ProtoMessage()               {}
func (*Story) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Story) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Story) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Story) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type TopStories struct {
	TopStories []*Story `protobuf:"bytes,1,rep,name=topStories" json:"topStories,omitempty"`
}

func (m *TopStories) Reset()                    { *m = TopStories{} }
func (m *TopStories) String() string            { return proto.CompactTextString(m) }
func (*TopStories) ProtoMessage()               {}
func (*TopStories) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TopStories) GetTopStories() []*Story {
	if m != nil {
		return m.TopStories
	}
	return nil
}

type Stories struct {
	Stories []*Story `protobuf:"bytes,1,rep,name=stories" json:"stories,omitempty"`
}

func (m *Stories) Reset()                    { *m = Stories{} }
func (m *Stories) String() string            { return proto.CompactTextString(m) }
func (*Stories) ProtoMessage()               {}
func (*Stories) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Stories) GetStories() []*Story {
	if m != nil {
		return m.Stories
	}
	return nil
}

func init() {
	proto.RegisterType((*Story)(nil), "gohackernews.Story")
	proto.RegisterType((*TopStories)(nil), "gohackernews.TopStories")
	proto.RegisterType((*Stories)(nil), "gohackernews.Stories")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Test service

type TestClient interface {
	GetStory(ctx context.Context, in *TopStories, opts ...grpc.CallOption) (*Story, error)
	GetStories(ctx context.Context, in *TopStories, opts ...grpc.CallOption) (*Stories, error)
}

type testClient struct {
	cc *grpc.ClientConn
}

func NewTestClient(cc *grpc.ClientConn) TestClient {
	return &testClient{cc}
}

func (c *testClient) GetStory(ctx context.Context, in *TopStories, opts ...grpc.CallOption) (*Story, error) {
	out := new(Story)
	err := grpc.Invoke(ctx, "/gohackernews.Test/GetStory", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testClient) GetStories(ctx context.Context, in *TopStories, opts ...grpc.CallOption) (*Stories, error) {
	out := new(Stories)
	err := grpc.Invoke(ctx, "/gohackernews.Test/GetStories", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Test service

type TestServer interface {
	GetStory(context.Context, *TopStories) (*Story, error)
	GetStories(context.Context, *TopStories) (*Stories, error)
}

func RegisterTestServer(s *grpc.Server, srv TestServer) {
	s.RegisterService(&_Test_serviceDesc, srv)
}

func _Test_GetStory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopStories)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).GetStory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gohackernews.Test/GetStory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).GetStory(ctx, req.(*TopStories))
	}
	return interceptor(ctx, in, info, handler)
}

func _Test_GetStories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopStories)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).GetStories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gohackernews.Test/GetStories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).GetStories(ctx, req.(*TopStories))
	}
	return interceptor(ctx, in, info, handler)
}

var _Test_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gohackernews.Test",
	HandlerType: (*TestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStory",
			Handler:    _Test_GetStory_Handler,
		},
		{
			MethodName: "GetStories",
			Handler:    _Test_GetStories_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go-hackernews.proto",
}

func init() { proto.RegisterFile("go-hackernews.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0xcf, 0xd7, 0xcd,
	0x48, 0x4c, 0xce, 0x4e, 0x2d, 0xca, 0x4b, 0x2d, 0x2f, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x49, 0xcf, 0x47, 0x88, 0x29, 0xd9, 0x73, 0xb1, 0x06, 0x97, 0xe4, 0x17, 0x55, 0x0a, 0xf1,
	0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06, 0x31, 0x65, 0xa6, 0x08, 0x89,
	0x70, 0xb1, 0x96, 0x64, 0x96, 0xe4, 0xa4, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x38,
	0x42, 0x02, 0x5c, 0xcc, 0xa5, 0x45, 0x39, 0x12, 0xcc, 0x60, 0x31, 0x10, 0x53, 0xc9, 0x91, 0x8b,
	0x2b, 0x24, 0xbf, 0x00, 0x64, 0x46, 0x66, 0x6a, 0xb1, 0x90, 0x31, 0x17, 0x57, 0x09, 0x9c, 0x27,
	0xc1, 0xa8, 0xc0, 0xac, 0xc1, 0x6d, 0x24, 0xac, 0x87, 0x6c, 0xa3, 0x1e, 0xd8, 0xba, 0x20, 0x24,
	0x65, 0x4a, 0x16, 0x5c, 0xec, 0x30, 0xfd, 0xba, 0x5c, 0xec, 0xc5, 0x84, 0x35, 0xc3, 0xd4, 0x18,
	0xb5, 0x30, 0x72, 0xb1, 0x84, 0xa4, 0x16, 0x97, 0x08, 0x59, 0x73, 0x71, 0xb8, 0xa7, 0x96, 0x40,
	0x7c, 0x22, 0x81, 0xaa, 0x05, 0xe1, 0x3a, 0x29, 0x6c, 0x86, 0x29, 0x31, 0x08, 0xd9, 0x73, 0x71,
	0x41, 0x35, 0x83, 0x9c, 0x80, 0x5b, 0xbb, 0x28, 0xa6, 0x76, 0x90, 0xf3, 0x19, 0x92, 0xd8, 0xc0,
	0x21, 0x6b, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x13, 0x25, 0x46, 0x83, 0x70, 0x01, 0x00, 0x00,
}
