// Code generated by protoc-gen-go. DO NOT EDIT.
// source: greeter.proto

package greeter

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Request struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_e585294ab3f34af5, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Response struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_e585294ab3f34af5, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*Request)(nil), "greeter.Request")
	proto.RegisterType((*Response)(nil), "greeter.Response")
}

func init() { proto.RegisterFile("greeter.proto", fileDescriptor_e585294ab3f34af5) }

var fileDescriptor_e585294ab3f34af5 = []byte{
	// 142 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x2f, 0x4a, 0x4d,
	0x2d, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0xe4, 0xb9,
	0xd8, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x44, 0xb8, 0x58, 0x93, 0xf3, 0x4b, 0xf3,
	0x4a, 0x24, 0x18, 0x15, 0x18, 0x35, 0x98, 0x83, 0x20, 0x1c, 0x25, 0x05, 0x2e, 0x8e, 0xa0, 0xd4,
	0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0xec, 0x2a, 0x8c, 0xca, 0xb9, 0xd8, 0xdd, 0x21, 0xa6, 0x09,
	0x19, 0x73, 0xb1, 0x05, 0x97, 0x14, 0xa5, 0x26, 0xe6, 0x0a, 0x09, 0xe8, 0xc1, 0x2c, 0x84, 0x1a,
	0x2f, 0x25, 0x88, 0x24, 0x02, 0x31, 0x4f, 0x89, 0x41, 0x83, 0xd1, 0x80, 0x51, 0xc8, 0x9c, 0x8b,
	0x27, 0x38, 0xb5, 0xa8, 0x2c, 0xb5, 0x88, 0x24, 0xad, 0x06, 0x8c, 0x49, 0x6c, 0x60, 0xbf, 0x18,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb6, 0xc3, 0x57, 0x4a, 0xdc, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	Stream(ctx context.Context, opts ...grpc.CallOption) (Greeter_StreamClient, error)
	ServerStream(ctx context.Context, in *Request, opts ...grpc.CallOption) (Greeter_ServerStreamClient, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) Stream(ctx context.Context, opts ...grpc.CallOption) (Greeter_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Greeter_serviceDesc.Streams[0], "/greeter.Greeter/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterStreamClient{stream}
	return x, nil
}

type Greeter_StreamClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type greeterStreamClient struct {
	grpc.ClientStream
}

func (x *greeterStreamClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterStreamClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterClient) ServerStream(ctx context.Context, in *Request, opts ...grpc.CallOption) (Greeter_ServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Greeter_serviceDesc.Streams[1], "/greeter.Greeter/ServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Greeter_ServerStreamClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type greeterServerStreamClient struct {
	grpc.ClientStream
}

func (x *greeterServerStreamClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	Stream(Greeter_StreamServer) error
	ServerStream(*Request, Greeter_ServerStreamServer) error
}

// UnimplementedGreeterServer can be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (*UnimplementedGreeterServer) Stream(srv Greeter_StreamServer) error {
	return status.Errorf(codes.Unimplemented, "method Stream not implemented")
}
func (*UnimplementedGreeterServer) ServerStream(req *Request, srv Greeter_ServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerStream not implemented")
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServer).Stream(&greeterStreamServer{stream})
}

type Greeter_StreamServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type greeterStreamServer struct {
	grpc.ServerStream
}

func (x *greeterStreamServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterStreamServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Greeter_ServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreeterServer).ServerStream(m, &greeterServerStreamServer{stream})
}

type Greeter_ServerStreamServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type greeterServerStreamServer struct {
	grpc.ServerStream
}

func (x *greeterServerStreamServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greeter.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _Greeter_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ServerStream",
			Handler:       _Greeter_ServerStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "greeter.proto",
}