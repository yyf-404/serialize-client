// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/serialize.proto

package serialize

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

type MsgRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   int32    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	IsUser               bool     `protobuf:"varint,3,opt,name=isUser,proto3" json:"isUser,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgRequest) Reset()         { *m = MsgRequest{} }
func (m *MsgRequest) String() string { return proto.CompactTextString(m) }
func (*MsgRequest) ProtoMessage()    {}
func (*MsgRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ed1a15a203eeb68, []int{0}
}

func (m *MsgRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgRequest.Unmarshal(m, b)
}
func (m *MsgRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgRequest.Marshal(b, m, deterministic)
}
func (m *MsgRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRequest.Merge(m, src)
}
func (m *MsgRequest) XXX_Size() int {
	return xxx_messageInfo_MsgRequest.Size(m)
}
func (m *MsgRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRequest proto.InternalMessageInfo

func (m *MsgRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MsgRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MsgRequest) GetIsUser() bool {
	if m != nil {
		return m.IsUser
	}
	return false
}

type MsgReply struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgReply) Reset()         { *m = MsgReply{} }
func (m *MsgReply) String() string { return proto.CompactTextString(m) }
func (*MsgReply) ProtoMessage()    {}
func (*MsgReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ed1a15a203eeb68, []int{1}
}

func (m *MsgReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgReply.Unmarshal(m, b)
}
func (m *MsgReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgReply.Marshal(b, m, deterministic)
}
func (m *MsgReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgReply.Merge(m, src)
}
func (m *MsgReply) XXX_Size() int {
	return xxx_messageInfo_MsgReply.Size(m)
}
func (m *MsgReply) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgReply.DiscardUnknown(m)
}

var xxx_messageInfo_MsgReply proto.InternalMessageInfo

func (m *MsgReply) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *MsgReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgRequest)(nil), "serialize.MsgRequest")
	proto.RegisterType((*MsgReply)(nil), "serialize.MsgReply")
}

func init() { proto.RegisterFile("pb/serialize.proto", fileDescriptor_8ed1a15a203eeb68) }

var fileDescriptor_8ed1a15a203eeb68 = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x48, 0xd2, 0x2f,
	0x4e, 0x2d, 0xca, 0x4c, 0xcc, 0xc9, 0xac, 0x4a, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x84, 0x0b, 0x28, 0x79, 0x70, 0x71, 0xf9, 0x16, 0xa7, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97,
	0x08, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81,
	0xd9, 0x42, 0x7c, 0x5c, 0x4c, 0x99, 0x29, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x4c, 0x99,
	0x29, 0x42, 0x62, 0x5c, 0x6c, 0x99, 0xc5, 0xa1, 0xc5, 0xa9, 0x45, 0x12, 0xcc, 0x0a, 0x8c, 0x1a,
	0x1c, 0x41, 0x50, 0x9e, 0x92, 0x05, 0x17, 0x07, 0xd8, 0xa4, 0x82, 0x9c, 0x4a, 0x90, 0x39, 0xc9,
	0xf9, 0x29, 0x10, 0x73, 0x58, 0x83, 0xc0, 0x6c, 0x21, 0x09, 0x2e, 0xf6, 0xdc, 0xd4, 0xe2, 0xe2,
	0xc4, 0xf4, 0x54, 0xb0, 0x61, 0x9c, 0x41, 0x30, 0xae, 0x91, 0x37, 0x97, 0x40, 0x30, 0xcc, 0x41,
	0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0xe6, 0x5c, 0xec, 0xc1, 0xa9, 0x79, 0x29, 0xbe,
	0xc5, 0xe9, 0x42, 0xa2, 0x7a, 0x08, 0xf7, 0x23, 0xdc, 0x2a, 0x25, 0x8c, 0x2e, 0x5c, 0x90, 0x53,
	0xa9, 0xc4, 0x90, 0xc4, 0x06, 0xf6, 0xa2, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x05, 0xda, 0xa6,
	0x76, 0xf8, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SerializeServiceClient is the client API for SerializeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SerializeServiceClient interface {
	//  方法
	SendMsg(ctx context.Context, in *MsgRequest, opts ...grpc.CallOption) (*MsgReply, error)
}

type serializeServiceClient struct {
	cc *grpc.ClientConn
}

func NewSerializeServiceClient(cc *grpc.ClientConn) SerializeServiceClient {
	return &serializeServiceClient{cc}
}

func (c *serializeServiceClient) SendMsg(ctx context.Context, in *MsgRequest, opts ...grpc.CallOption) (*MsgReply, error) {
	out := new(MsgReply)
	err := c.cc.Invoke(ctx, "/serialize.SerializeService/SendMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SerializeServiceServer is the server API for SerializeService service.
type SerializeServiceServer interface {
	//  方法
	SendMsg(context.Context, *MsgRequest) (*MsgReply, error)
}

// UnimplementedSerializeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSerializeServiceServer struct {
}

func (*UnimplementedSerializeServiceServer) SendMsg(ctx context.Context, req *MsgRequest) (*MsgReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMsg not implemented")
}

func RegisterSerializeServiceServer(s *grpc.Server, srv SerializeServiceServer) {
	s.RegisterService(&_SerializeService_serviceDesc, srv)
}

func _SerializeService_SendMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SerializeServiceServer).SendMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serialize.SerializeService/SendMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SerializeServiceServer).SendMsg(ctx, req.(*MsgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SerializeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "serialize.SerializeService",
	HandlerType: (*SerializeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMsg",
			Handler:    _SerializeService_SendMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/serialize.proto",
}