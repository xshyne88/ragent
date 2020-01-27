// Code generated by protoc-gen-go. DO NOT EDIT.
// source: communication.proto

package protobuf

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

type CommandRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CommandType          string   `protobuf:"bytes,2,opt,name=commandType,proto3" json:"commandType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommandRequest) Reset()         { *m = CommandRequest{} }
func (m *CommandRequest) String() string { return proto.CompactTextString(m) }
func (*CommandRequest) ProtoMessage()    {}
func (*CommandRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b64068f22c460ac1, []int{0}
}

func (m *CommandRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandRequest.Unmarshal(m, b)
}
func (m *CommandRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandRequest.Marshal(b, m, deterministic)
}
func (m *CommandRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandRequest.Merge(m, src)
}
func (m *CommandRequest) XXX_Size() int {
	return xxx_messageInfo_CommandRequest.Size(m)
}
func (m *CommandRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CommandRequest proto.InternalMessageInfo

func (m *CommandRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CommandRequest) GetCommandType() string {
	if m != nil {
		return m.CommandType
	}
	return ""
}

type CommandResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Acknowledged         bool     `protobuf:"varint,2,opt,name=acknowledged,proto3" json:"acknowledged,omitempty"`
	Payload              string   `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommandResponse) Reset()         { *m = CommandResponse{} }
func (m *CommandResponse) String() string { return proto.CompactTextString(m) }
func (*CommandResponse) ProtoMessage()    {}
func (*CommandResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b64068f22c460ac1, []int{1}
}

func (m *CommandResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandResponse.Unmarshal(m, b)
}
func (m *CommandResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandResponse.Marshal(b, m, deterministic)
}
func (m *CommandResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandResponse.Merge(m, src)
}
func (m *CommandResponse) XXX_Size() int {
	return xxx_messageInfo_CommandResponse.Size(m)
}
func (m *CommandResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CommandResponse proto.InternalMessageInfo

func (m *CommandResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CommandResponse) GetAcknowledged() bool {
	if m != nil {
		return m.Acknowledged
	}
	return false
}

func (m *CommandResponse) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func init() {
	proto.RegisterType((*CommandRequest)(nil), "communication.CommandRequest")
	proto.RegisterType((*CommandResponse)(nil), "communication.CommandResponse")
}

func init() { proto.RegisterFile("communication.proto", fileDescriptor_b64068f22c460ac1) }

var fileDescriptor_b64068f22c460ac1 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0xce, 0xcf, 0xcd,
	0x2d, 0xcd, 0xcb, 0x4c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x45, 0x11, 0x54, 0x72, 0xe2, 0xe2, 0x73, 0xce, 0xcf, 0xcd, 0x4d, 0xcc, 0x4b, 0x09, 0x4a,
	0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4,
	0xe0, 0x0c, 0x62, 0xca, 0x4c, 0x11, 0x52, 0xe0, 0xe2, 0x4e, 0x86, 0xa8, 0x08, 0xa9, 0x2c, 0x48,
	0x95, 0x60, 0x02, 0x4b, 0x20, 0x0b, 0x29, 0xc5, 0x73, 0xf1, 0xc3, 0xcd, 0x28, 0x2e, 0xc8, 0xcf,
	0x2b, 0x4e, 0xc5, 0x30, 0x44, 0x89, 0x8b, 0x27, 0x31, 0x39, 0x3b, 0x2f, 0xbf, 0x3c, 0x27, 0x35,
	0x25, 0x3d, 0x35, 0x05, 0x6c, 0x0a, 0x47, 0x10, 0x8a, 0x98, 0x90, 0x04, 0x17, 0x7b, 0x41, 0x62,
	0x65, 0x4e, 0x7e, 0x62, 0x8a, 0x04, 0x33, 0x58, 0x23, 0x8c, 0x6b, 0x94, 0xc2, 0xc5, 0x0b, 0xb5,
	0xc0, 0xb3, 0xb8, 0xb8, 0x34, 0xb5, 0x48, 0x28, 0x98, 0x8b, 0x27, 0x38, 0x35, 0x2f, 0x05, 0x2a,
	0x58, 0x2c, 0x24, 0xab, 0x87, 0xea, 0x55, 0x54, 0x2f, 0x49, 0xc9, 0xe1, 0x92, 0x86, 0xb8, 0x56,
	0x89, 0x41, 0x83, 0xd1, 0x80, 0xd1, 0x89, 0x2b, 0x8a, 0x03, 0x1c, 0x44, 0x49, 0xa5, 0x69, 0x49,
	0x6c, 0x60, 0x96, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xcc, 0x89, 0x5f, 0x00, 0x43, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CommandIssuerClient is the client API for CommandIssuer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommandIssuerClient interface {
	SendCommands(ctx context.Context, opts ...grpc.CallOption) (CommandIssuer_SendCommandsClient, error)
}

type commandIssuerClient struct {
	cc *grpc.ClientConn
}

func NewCommandIssuerClient(cc *grpc.ClientConn) CommandIssuerClient {
	return &commandIssuerClient{cc}
}

func (c *commandIssuerClient) SendCommands(ctx context.Context, opts ...grpc.CallOption) (CommandIssuer_SendCommandsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CommandIssuer_serviceDesc.Streams[0], "/communication.CommandIssuer/SendCommands", opts...)
	if err != nil {
		return nil, err
	}
	x := &commandIssuerSendCommandsClient{stream}
	return x, nil
}

type CommandIssuer_SendCommandsClient interface {
	Send(*CommandRequest) error
	Recv() (*CommandResponse, error)
	grpc.ClientStream
}

type commandIssuerSendCommandsClient struct {
	grpc.ClientStream
}

func (x *commandIssuerSendCommandsClient) Send(m *CommandRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *commandIssuerSendCommandsClient) Recv() (*CommandResponse, error) {
	m := new(CommandResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CommandIssuerServer is the server API for CommandIssuer service.
type CommandIssuerServer interface {
	SendCommands(CommandIssuer_SendCommandsServer) error
}

// UnimplementedCommandIssuerServer can be embedded to have forward compatible implementations.
type UnimplementedCommandIssuerServer struct {
}

func (*UnimplementedCommandIssuerServer) SendCommands(srv CommandIssuer_SendCommandsServer) error {
	return status.Errorf(codes.Unimplemented, "method SendCommands not implemented")
}

func RegisterCommandIssuerServer(s *grpc.Server, srv CommandIssuerServer) {
	s.RegisterService(&_CommandIssuer_serviceDesc, srv)
}

func _CommandIssuer_SendCommands_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CommandIssuerServer).SendCommands(&commandIssuerSendCommandsServer{stream})
}

type CommandIssuer_SendCommandsServer interface {
	Send(*CommandResponse) error
	Recv() (*CommandRequest, error)
	grpc.ServerStream
}

type commandIssuerSendCommandsServer struct {
	grpc.ServerStream
}

func (x *commandIssuerSendCommandsServer) Send(m *CommandResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *commandIssuerSendCommandsServer) Recv() (*CommandRequest, error) {
	m := new(CommandRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _CommandIssuer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "communication.CommandIssuer",
	HandlerType: (*CommandIssuerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendCommands",
			Handler:       _CommandIssuer_SendCommands_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "communication.proto",
}
