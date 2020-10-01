// Code generated by protoc-gen-go. DO NOT EDIT.
// source: UniversaltestService/universaltestservice.proto

package UniversaltestService

import (
	universaltest "../universaltest"
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

type CreateUniversaltestServiceRequest struct {
	Data                 *universaltest.Universaltest `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *CreateUniversaltestServiceRequest) Reset()         { *m = CreateUniversaltestServiceRequest{} }
func (m *CreateUniversaltestServiceRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUniversaltestServiceRequest) ProtoMessage()    {}
func (*CreateUniversaltestServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c035c7f431c02372, []int{0}
}

func (m *CreateUniversaltestServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUniversaltestServiceRequest.Unmarshal(m, b)
}
func (m *CreateUniversaltestServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUniversaltestServiceRequest.Marshal(b, m, deterministic)
}
func (m *CreateUniversaltestServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUniversaltestServiceRequest.Merge(m, src)
}
func (m *CreateUniversaltestServiceRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUniversaltestServiceRequest.Size(m)
}
func (m *CreateUniversaltestServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUniversaltestServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUniversaltestServiceRequest proto.InternalMessageInfo

func (m *CreateUniversaltestServiceRequest) GetData() *universaltest.Universaltest {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateUniversaltestServiceRequest)(nil), "UniversaltestService.CreateUniversaltestServiceRequest")
}

func init() {
	proto.RegisterFile("UniversaltestService/universaltestservice.proto", fileDescriptor_c035c7f431c02372)
}

var fileDescriptor_c035c7f431c02372 = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x0f, 0xcd, 0xcb, 0x2c,
	0x4b, 0x2d, 0x2a, 0x4e, 0xcc, 0x29, 0x49, 0x2d, 0x2e, 0x09, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e,
	0xd5, 0x2f, 0x45, 0x16, 0x2c, 0x86, 0x08, 0xea, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x89, 0x60,
	0xd3, 0x20, 0x25, 0x93, 0x9e, 0x9f, 0x9f, 0x9e, 0x93, 0xaa, 0x9f, 0x58, 0x90, 0xa9, 0x9f, 0x98,
	0x97, 0x97, 0x5f, 0x92, 0x58, 0x92, 0x99, 0x9f, 0x57, 0x0c, 0xd1, 0x23, 0xa5, 0x88, 0x62, 0x1e,
	0xaa, 0xe9, 0x50, 0x25, 0xd2, 0x50, 0x03, 0xc0, 0xbc, 0xa4, 0xd2, 0x34, 0xfd, 0xd4, 0xdc, 0x82,
	0x92, 0x4a, 0x88, 0xa4, 0x52, 0x28, 0x97, 0xa2, 0x73, 0x51, 0x6a, 0x62, 0x49, 0x2a, 0x36, 0xbb,
	0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x0c, 0xb8, 0x58, 0x52, 0x12, 0x4b, 0x12, 0x25,
	0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0x64, 0xf4, 0x50, 0x6d, 0x41, 0xd1, 0x19, 0x04, 0x56, 0x69,
	0xb4, 0x99, 0x91, 0x0b, 0xab, 0x6f, 0x84, 0x66, 0x31, 0x72, 0x09, 0x43, 0x2c, 0x44, 0x31, 0x44,
	0xc8, 0x5c, 0x0f, 0x9b, 0x72, 0x3d, 0x82, 0x6e, 0x93, 0x52, 0xc2, 0xe7, 0x1a, 0xd7, 0xbc, 0x92,
	0xcc, 0x92, 0x4a, 0x25, 0xb5, 0xa6, 0xcb, 0x4f, 0x26, 0x33, 0x29, 0x28, 0x49, 0xe8, 0xe7, 0xe6,
	0x27, 0x67, 0x83, 0x1c, 0x88, 0x16, 0x0d, 0x56, 0x60, 0x57, 0x3b, 0x89, 0x47, 0x89, 0x62, 0x8d,
	0xb4, 0x24, 0x36, 0x70, 0x60, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xde, 0x5e, 0x3a, 0x59,
	0xd3, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UniversaltestServiceClient is the client API for UniversaltestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UniversaltestServiceClient interface {
	// Creates a new universaltest
	Createuniversaltest(ctx context.Context, in *CreateUniversaltestServiceRequest, opts ...grpc.CallOption) (*universaltest.UniversaltestEntity, error)
}

type universaltestServiceClient struct {
	cc *grpc.ClientConn
}

func NewUniversaltestServiceClient(cc *grpc.ClientConn) UniversaltestServiceClient {
	return &universaltestServiceClient{cc}
}

func (c *universaltestServiceClient) Createuniversaltest(ctx context.Context, in *CreateUniversaltestServiceRequest, opts ...grpc.CallOption) (*universaltest.UniversaltestEntity, error) {
	out := new(universaltest.UniversaltestEntity)
	err := c.cc.Invoke(ctx, "/UniversaltestService.UniversaltestService/Createuniversaltest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UniversaltestServiceServer is the server API for UniversaltestService service.
type UniversaltestServiceServer interface {
	// Creates a new universaltest
	Createuniversaltest(context.Context, *CreateUniversaltestServiceRequest) (*universaltest.UniversaltestEntity, error)
}

func RegisterUniversaltestServiceServer(s *grpc.Server, srv UniversaltestServiceServer) {
	s.RegisterService(&_UniversaltestService_serviceDesc, srv)
}

func _UniversaltestService_Createuniversaltest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUniversaltestServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniversaltestServiceServer).Createuniversaltest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UniversaltestService.UniversaltestService/Createuniversaltest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniversaltestServiceServer).Createuniversaltest(ctx, req.(*CreateUniversaltestServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UniversaltestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "UniversaltestService.UniversaltestService",
	HandlerType: (*UniversaltestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Createuniversaltest",
			Handler:    _UniversaltestService_Createuniversaltest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "UniversaltestService/universaltestservice.proto",
}
