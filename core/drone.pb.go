// Code generated by protoc-gen-go. DO NOT EDIT.
// source: drone.proto

package core

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

type Status struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	StatusCode           int32    `protobuf:"varint,2,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9911fbfd24a7e00, []int{0}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Status) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

type FileFragment struct {
	FileName             string   `protobuf:"bytes,1,opt,name=fileName,proto3" json:"fileName,omitempty"`
	FragmentId           int32    `protobuf:"varint,2,opt,name=fragmentId,proto3" json:"fragmentId,omitempty"`
	FragmentContent      []byte   `protobuf:"bytes,3,opt,name=fragmentContent,proto3" json:"fragmentContent,omitempty"`
	TotalFragments       int32    `protobuf:"varint,4,opt,name=totalFragments,proto3" json:"totalFragments,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileFragment) Reset()         { *m = FileFragment{} }
func (m *FileFragment) String() string { return proto.CompactTextString(m) }
func (*FileFragment) ProtoMessage()    {}
func (*FileFragment) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9911fbfd24a7e00, []int{1}
}

func (m *FileFragment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileFragment.Unmarshal(m, b)
}
func (m *FileFragment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileFragment.Marshal(b, m, deterministic)
}
func (m *FileFragment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileFragment.Merge(m, src)
}
func (m *FileFragment) XXX_Size() int {
	return xxx_messageInfo_FileFragment.Size(m)
}
func (m *FileFragment) XXX_DiscardUnknown() {
	xxx_messageInfo_FileFragment.DiscardUnknown(m)
}

var xxx_messageInfo_FileFragment proto.InternalMessageInfo

func (m *FileFragment) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *FileFragment) GetFragmentId() int32 {
	if m != nil {
		return m.FragmentId
	}
	return 0
}

func (m *FileFragment) GetFragmentContent() []byte {
	if m != nil {
		return m.FragmentContent
	}
	return nil
}

func (m *FileFragment) GetTotalFragments() int32 {
	if m != nil {
		return m.TotalFragments
	}
	return 0
}

func init() {
	proto.RegisterType((*Status)(nil), "core.Status")
	proto.RegisterType((*FileFragment)(nil), "core.FileFragment")
}

func init() { proto.RegisterFile("drone.proto", fileDescriptor_a9911fbfd24a7e00) }

var fileDescriptor_a9911fbfd24a7e00 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x89, 0xee, 0xae, 0x3a, 0x5b, 0x14, 0xe6, 0x14, 0xf6, 0x20, 0x65, 0x0f, 0x92, 0x53,
	0x41, 0xbd, 0x79, 0xb4, 0x52, 0xf0, 0xe2, 0x21, 0x3e, 0x41, 0x6c, 0xa7, 0xa5, 0xd0, 0x26, 0x92,
	0x8c, 0x3e, 0x8d, 0x0f, 0x2b, 0x69, 0x8d, 0x5b, 0x7a, 0xfc, 0x3f, 0x86, 0x2f, 0x7f, 0x7e, 0xd8,
	0x37, 0xde, 0x59, 0x2a, 0x3e, 0xbd, 0x63, 0x87, 0x9b, 0xda, 0x79, 0x3a, 0x3e, 0xc3, 0xee, 0x9d,
	0x0d, 0x7f, 0x05, 0x94, 0x70, 0x31, 0x52, 0x08, 0xa6, 0x23, 0x29, 0x72, 0xa1, 0xae, 0x74, 0x8a,
	0x78, 0x0b, 0x10, 0xa6, 0x9b, 0xd2, 0x35, 0x24, 0xcf, 0x72, 0xa1, 0xb6, 0x7a, 0x41, 0x8e, 0x3f,
	0x02, 0xb2, 0xaa, 0x1f, 0xa8, 0xf2, 0xa6, 0x1b, 0xc9, 0x32, 0x1e, 0xe0, 0xb2, 0xed, 0x07, 0x7a,
	0x33, 0x63, 0x72, 0xfd, 0xe7, 0x28, 0x6b, 0xff, 0xee, 0x5e, 0x9b, 0x24, 0x3b, 0x11, 0x54, 0x70,
	0x93, 0x52, 0xe9, 0x2c, 0x93, 0x65, 0x79, 0x9e, 0x0b, 0x95, 0xe9, 0x35, 0xc6, 0x3b, 0xb8, 0x66,
	0xc7, 0x66, 0x48, 0xcf, 0x06, 0xb9, 0x99, 0x6c, 0x2b, 0xfa, 0xf0, 0x04, 0xdb, 0x97, 0xf8, 0x6f,
	0xbc, 0x87, 0xbd, 0xa6, 0x9a, 0xfa, 0x6f, 0x8a, 0x6d, 0x11, 0x8b, 0xb8, 0x40, 0xb1, 0x6c, 0x7e,
	0xc8, 0x66, 0x36, 0x4f, 0xa2, 0xc4, 0xc7, 0x6e, 0xda, 0xea, 0xf1, 0x37, 0x00, 0x00, 0xff, 0xff,
	0x28, 0x7d, 0xd9, 0x50, 0x3a, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DroneClient is the client API for Drone service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DroneClient interface {
	ReceiveFile(ctx context.Context, opts ...grpc.CallOption) (Drone_ReceiveFileClient, error)
}

type droneClient struct {
	cc grpc.ClientConnInterface
}

func NewDroneClient(cc grpc.ClientConnInterface) DroneClient {
	return &droneClient{cc}
}

func (c *droneClient) ReceiveFile(ctx context.Context, opts ...grpc.CallOption) (Drone_ReceiveFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Drone_serviceDesc.Streams[0], "/core.Drone/ReceiveFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &droneReceiveFileClient{stream}
	return x, nil
}

type Drone_ReceiveFileClient interface {
	Send(*FileFragment) error
	CloseAndRecv() (*Status, error)
	grpc.ClientStream
}

type droneReceiveFileClient struct {
	grpc.ClientStream
}

func (x *droneReceiveFileClient) Send(m *FileFragment) error {
	return x.ClientStream.SendMsg(m)
}

func (x *droneReceiveFileClient) CloseAndRecv() (*Status, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Status)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DroneServer is the server API for Drone service.
type DroneServer interface {
	ReceiveFile(Drone_ReceiveFileServer) error
}

// UnimplementedDroneServer can be embedded to have forward compatible implementations.
type UnimplementedDroneServer struct {
}

func (*UnimplementedDroneServer) ReceiveFile(srv Drone_ReceiveFileServer) error {
	return status.Errorf(codes.Unimplemented, "method ReceiveFile not implemented")
}

func RegisterDroneServer(s *grpc.Server, srv DroneServer) {
	s.RegisterService(&_Drone_serviceDesc, srv)
}

func _Drone_ReceiveFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DroneServer).ReceiveFile(&droneReceiveFileServer{stream})
}

type Drone_ReceiveFileServer interface {
	SendAndClose(*Status) error
	Recv() (*FileFragment, error)
	grpc.ServerStream
}

type droneReceiveFileServer struct {
	grpc.ServerStream
}

func (x *droneReceiveFileServer) SendAndClose(m *Status) error {
	return x.ServerStream.SendMsg(m)
}

func (x *droneReceiveFileServer) Recv() (*FileFragment, error) {
	m := new(FileFragment)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Drone_serviceDesc = grpc.ServiceDesc{
	ServiceName: "core.Drone",
	HandlerType: (*DroneServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReceiveFile",
			Handler:       _Drone_ReceiveFile_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "drone.proto",
}