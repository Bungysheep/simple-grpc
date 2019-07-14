// Code generated by protoc-gen-go. DO NOT EDIT.
// source: src/protofiles/simple.proto

package protofiles

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type SayHelloRequest struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName             string   `protobuf:"bytes,2,opt,name=lastName,proto3" json:"lastName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SayHelloRequest) Reset()         { *m = SayHelloRequest{} }
func (m *SayHelloRequest) String() string { return proto.CompactTextString(m) }
func (*SayHelloRequest) ProtoMessage()    {}
func (*SayHelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_266a9d6d68e82ed0, []int{0}
}

func (m *SayHelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SayHelloRequest.Unmarshal(m, b)
}
func (m *SayHelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SayHelloRequest.Marshal(b, m, deterministic)
}
func (m *SayHelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SayHelloRequest.Merge(m, src)
}
func (m *SayHelloRequest) XXX_Size() int {
	return xxx_messageInfo_SayHelloRequest.Size(m)
}
func (m *SayHelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SayHelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SayHelloRequest proto.InternalMessageInfo

func (m *SayHelloRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *SayHelloRequest) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

type SayHelloResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SayHelloResponse) Reset()         { *m = SayHelloResponse{} }
func (m *SayHelloResponse) String() string { return proto.CompactTextString(m) }
func (*SayHelloResponse) ProtoMessage()    {}
func (*SayHelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_266a9d6d68e82ed0, []int{1}
}

func (m *SayHelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SayHelloResponse.Unmarshal(m, b)
}
func (m *SayHelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SayHelloResponse.Marshal(b, m, deterministic)
}
func (m *SayHelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SayHelloResponse.Merge(m, src)
}
func (m *SayHelloResponse) XXX_Size() int {
	return xxx_messageInfo_SayHelloResponse.Size(m)
}
func (m *SayHelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SayHelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SayHelloResponse proto.InternalMessageInfo

func (m *SayHelloResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type FibonaciRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FibonaciRequest) Reset()         { *m = FibonaciRequest{} }
func (m *FibonaciRequest) String() string { return proto.CompactTextString(m) }
func (*FibonaciRequest) ProtoMessage()    {}
func (*FibonaciRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_266a9d6d68e82ed0, []int{2}
}

func (m *FibonaciRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibonaciRequest.Unmarshal(m, b)
}
func (m *FibonaciRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibonaciRequest.Marshal(b, m, deterministic)
}
func (m *FibonaciRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibonaciRequest.Merge(m, src)
}
func (m *FibonaciRequest) XXX_Size() int {
	return xxx_messageInfo_FibonaciRequest.Size(m)
}
func (m *FibonaciRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FibonaciRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FibonaciRequest proto.InternalMessageInfo

func (m *FibonaciRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type FibonaciResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FibonaciResponse) Reset()         { *m = FibonaciResponse{} }
func (m *FibonaciResponse) String() string { return proto.CompactTextString(m) }
func (*FibonaciResponse) ProtoMessage()    {}
func (*FibonaciResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_266a9d6d68e82ed0, []int{3}
}

func (m *FibonaciResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibonaciResponse.Unmarshal(m, b)
}
func (m *FibonaciResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibonaciResponse.Marshal(b, m, deterministic)
}
func (m *FibonaciResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibonaciResponse.Merge(m, src)
}
func (m *FibonaciResponse) XXX_Size() int {
	return xxx_messageInfo_FibonaciResponse.Size(m)
}
func (m *FibonaciResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FibonaciResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FibonaciResponse proto.InternalMessageInfo

func (m *FibonaciResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*SayHelloRequest)(nil), "protofiles.SayHelloRequest")
	proto.RegisterType((*SayHelloResponse)(nil), "protofiles.SayHelloResponse")
	proto.RegisterType((*FibonaciRequest)(nil), "protofiles.FibonaciRequest")
	proto.RegisterType((*FibonaciResponse)(nil), "protofiles.FibonaciResponse")
}

func init() { proto.RegisterFile("src/protofiles/simple.proto", fileDescriptor_266a9d6d68e82ed0) }

var fileDescriptor_266a9d6d68e82ed0 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x2e, 0x4a, 0xd6,
	0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xcb, 0xcc, 0x49, 0x2d, 0xd6, 0x2f, 0xce, 0xcc, 0x2d, 0xc8,
	0x49, 0xd5, 0x03, 0x8b, 0x08, 0x71, 0x21, 0x24, 0x94, 0xbc, 0xb9, 0xf8, 0x83, 0x13, 0x2b, 0x3d,
	0x52, 0x73, 0x72, 0xf2, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x64, 0xb8, 0x38, 0xd3,
	0x32, 0x8b, 0x8a, 0x4b, 0xfc, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x10,
	0x02, 0x42, 0x52, 0x5c, 0x1c, 0x39, 0x89, 0x50, 0x49, 0x26, 0xb0, 0x24, 0x9c, 0xaf, 0xa4, 0xc5,
	0x25, 0x80, 0x30, 0xac, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x8c, 0x8b, 0xad, 0x28, 0xb5,
	0xb8, 0x34, 0xa7, 0x04, 0x6a, 0x14, 0x94, 0xa7, 0xa4, 0xc9, 0xc5, 0xef, 0x96, 0x99, 0x94, 0x9f,
	0x97, 0x98, 0x9c, 0x09, 0xb3, 0x58, 0x8c, 0x8b, 0x2d, 0xaf, 0x34, 0x37, 0x29, 0xb5, 0x08, 0xac,
	0x94, 0x35, 0x08, 0xca, 0x03, 0x19, 0x8b, 0x50, 0x8a, 0xd5, 0x58, 0x56, 0x98, 0xb1, 0x46, 0xf3,
	0x19, 0xb9, 0x78, 0x83, 0xc1, 0x9e, 0x0d, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x15, 0x72, 0xe5,
	0xe2, 0x80, 0x39, 0x4a, 0x48, 0x5a, 0x0f, 0xe1, 0x75, 0x3d, 0x34, 0x7f, 0x4b, 0xc9, 0x60, 0x97,
	0x84, 0x5a, 0xe8, 0xce, 0xc5, 0x01, 0x73, 0x04, 0xaa, 0x31, 0x68, 0xbe, 0x40, 0x35, 0x06, 0xdd,
	0xdd, 0x06, 0x8c, 0x4e, 0x3c, 0x51, 0x48, 0xe1, 0x9f, 0xc4, 0x06, 0x66, 0x1b, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x41, 0x5e, 0x7e, 0x76, 0xb1, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SimpleServiceClient is the client API for SimpleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SimpleServiceClient interface {
	SayHello(ctx context.Context, in *SayHelloRequest, opts ...grpc.CallOption) (*SayHelloResponse, error)
	Fibonaci(ctx context.Context, in *FibonaciRequest, opts ...grpc.CallOption) (SimpleService_FibonaciClient, error)
}

type simpleServiceClient struct {
	cc *grpc.ClientConn
}

func NewSimpleServiceClient(cc *grpc.ClientConn) SimpleServiceClient {
	return &simpleServiceClient{cc}
}

func (c *simpleServiceClient) SayHello(ctx context.Context, in *SayHelloRequest, opts ...grpc.CallOption) (*SayHelloResponse, error) {
	out := new(SayHelloResponse)
	err := c.cc.Invoke(ctx, "/protofiles.SimpleService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simpleServiceClient) Fibonaci(ctx context.Context, in *FibonaciRequest, opts ...grpc.CallOption) (SimpleService_FibonaciClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SimpleService_serviceDesc.Streams[0], "/protofiles.SimpleService/Fibonaci", opts...)
	if err != nil {
		return nil, err
	}
	x := &simpleServiceFibonaciClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SimpleService_FibonaciClient interface {
	Recv() (*FibonaciResponse, error)
	grpc.ClientStream
}

type simpleServiceFibonaciClient struct {
	grpc.ClientStream
}

func (x *simpleServiceFibonaciClient) Recv() (*FibonaciResponse, error) {
	m := new(FibonaciResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SimpleServiceServer is the server API for SimpleService service.
type SimpleServiceServer interface {
	SayHello(context.Context, *SayHelloRequest) (*SayHelloResponse, error)
	Fibonaci(*FibonaciRequest, SimpleService_FibonaciServer) error
}

// UnimplementedSimpleServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSimpleServiceServer struct {
}

func (*UnimplementedSimpleServiceServer) SayHello(ctx context.Context, req *SayHelloRequest) (*SayHelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (*UnimplementedSimpleServiceServer) Fibonaci(req *FibonaciRequest, srv SimpleService_FibonaciServer) error {
	return status.Errorf(codes.Unimplemented, "method Fibonaci not implemented")
}

func RegisterSimpleServiceServer(s *grpc.Server, srv SimpleServiceServer) {
	s.RegisterService(&_SimpleService_serviceDesc, srv)
}

func _SimpleService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SayHelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protofiles.SimpleService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleServiceServer).SayHello(ctx, req.(*SayHelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimpleService_Fibonaci_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FibonaciRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SimpleServiceServer).Fibonaci(m, &simpleServiceFibonaciServer{stream})
}

type SimpleService_FibonaciServer interface {
	Send(*FibonaciResponse) error
	grpc.ServerStream
}

type simpleServiceFibonaciServer struct {
	grpc.ServerStream
}

func (x *simpleServiceFibonaciServer) Send(m *FibonaciResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _SimpleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protofiles.SimpleService",
	HandlerType: (*SimpleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _SimpleService_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Fibonaci",
			Handler:       _SimpleService_Fibonaci_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "src/protofiles/simple.proto",
}
