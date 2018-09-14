// Code generated by protoc-gen-go. DO NOT EDIT.
// source: screenpb/screen.proto

package screenpb

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

type SampleObject struct {
	FirstInt             int32    `protobuf:"varint,1,opt,name=first_int,json=firstInt,proto3" json:"first_int,omitempty"`
	SecondInt            int32    `protobuf:"varint,2,opt,name=second_int,json=secondInt,proto3" json:"second_int,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SampleObject) Reset()         { *m = SampleObject{} }
func (m *SampleObject) String() string { return proto.CompactTextString(m) }
func (*SampleObject) ProtoMessage()    {}
func (*SampleObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_screen_36886b82cc8bedcc, []int{0}
}
func (m *SampleObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SampleObject.Unmarshal(m, b)
}
func (m *SampleObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SampleObject.Marshal(b, m, deterministic)
}
func (dst *SampleObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SampleObject.Merge(dst, src)
}
func (m *SampleObject) XXX_Size() int {
	return xxx_messageInfo_SampleObject.Size(m)
}
func (m *SampleObject) XXX_DiscardUnknown() {
	xxx_messageInfo_SampleObject.DiscardUnknown(m)
}

var xxx_messageInfo_SampleObject proto.InternalMessageInfo

func (m *SampleObject) GetFirstInt() int32 {
	if m != nil {
		return m.FirstInt
	}
	return 0
}

func (m *SampleObject) GetSecondInt() int32 {
	if m != nil {
		return m.SecondInt
	}
	return 0
}

type SampleRequest struct {
	Calculation          *SampleObject `protobuf:"bytes,1,opt,name=calculation,proto3" json:"calculation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SampleRequest) Reset()         { *m = SampleRequest{} }
func (m *SampleRequest) String() string { return proto.CompactTextString(m) }
func (*SampleRequest) ProtoMessage()    {}
func (*SampleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_screen_36886b82cc8bedcc, []int{1}
}
func (m *SampleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SampleRequest.Unmarshal(m, b)
}
func (m *SampleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SampleRequest.Marshal(b, m, deterministic)
}
func (dst *SampleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SampleRequest.Merge(dst, src)
}
func (m *SampleRequest) XXX_Size() int {
	return xxx_messageInfo_SampleRequest.Size(m)
}
func (m *SampleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SampleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SampleRequest proto.InternalMessageInfo

func (m *SampleRequest) GetCalculation() *SampleObject {
	if m != nil {
		return m.Calculation
	}
	return nil
}

type SampleResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SampleResponse) Reset()         { *m = SampleResponse{} }
func (m *SampleResponse) String() string { return proto.CompactTextString(m) }
func (*SampleResponse) ProtoMessage()    {}
func (*SampleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_screen_36886b82cc8bedcc, []int{2}
}
func (m *SampleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SampleResponse.Unmarshal(m, b)
}
func (m *SampleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SampleResponse.Marshal(b, m, deterministic)
}
func (dst *SampleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SampleResponse.Merge(dst, src)
}
func (m *SampleResponse) XXX_Size() int {
	return xxx_messageInfo_SampleResponse.Size(m)
}
func (m *SampleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SampleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SampleResponse proto.InternalMessageInfo

func (m *SampleResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*SampleObject)(nil), "screen.SampleObject")
	proto.RegisterType((*SampleRequest)(nil), "screen.SampleRequest")
	proto.RegisterType((*SampleResponse)(nil), "screen.SampleResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ScreenServiceClient is the client API for ScreenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ScreenServiceClient interface {
	Sample(ctx context.Context, in *SampleRequest, opts ...grpc.CallOption) (*SampleResponse, error)
}

type screenServiceClient struct {
	cc *grpc.ClientConn
}

func NewScreenServiceClient(cc *grpc.ClientConn) ScreenServiceClient {
	return &screenServiceClient{cc}
}

func (c *screenServiceClient) Sample(ctx context.Context, in *SampleRequest, opts ...grpc.CallOption) (*SampleResponse, error) {
	out := new(SampleResponse)
	err := c.cc.Invoke(ctx, "/screen.ScreenService/Sample", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScreenServiceServer is the server API for ScreenService service.
type ScreenServiceServer interface {
	Sample(context.Context, *SampleRequest) (*SampleResponse, error)
}

func RegisterScreenServiceServer(s *grpc.Server, srv ScreenServiceServer) {
	s.RegisterService(&_ScreenService_serviceDesc, srv)
}

func _ScreenService_Sample_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SampleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScreenServiceServer).Sample(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/screen.ScreenService/Sample",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScreenServiceServer).Sample(ctx, req.(*SampleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ScreenService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "screen.ScreenService",
	HandlerType: (*ScreenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sample",
			Handler:    _ScreenService_Sample_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "screenpb/screen.proto",
}

func init() { proto.RegisterFile("screenpb/screen.proto", fileDescriptor_screen_36886b82cc8bedcc) }

var fileDescriptor_screen_36886b82cc8bedcc = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0x5d, 0xc1, 0xb0, 0x3b, 0xb5, 0x1e, 0x82, 0x2d, 0xa2, 0x08, 0x92, 0x53, 0x4f, 0x15,
	0x2a, 0x08, 0x5e, 0xbd, 0x48, 0x7b, 0x11, 0xda, 0x9b, 0x17, 0x69, 0xe3, 0x08, 0x91, 0x98, 0xc4,
	0x64, 0xea, 0xef, 0x17, 0x92, 0x16, 0xea, 0xde, 0x32, 0x6f, 0x32, 0xdf, 0x7b, 0x3c, 0x28, 0x82,
	0xf4, 0x88, 0xc6, 0x4d, 0xf7, 0xe9, 0x51, 0x3b, 0x6f, 0xc9, 0x72, 0x96, 0x26, 0xd1, 0xc1, 0xf9,
	0x30, 0x7e, 0x3b, 0x8d, 0xaf, 0xd3, 0x17, 0x4a, 0xe2, 0x37, 0x70, 0xf8, 0x54, 0x3e, 0xd0, 0xbb,
	0x32, 0x74, 0xb5, 0xbb, 0xdb, 0x55, 0x67, 0xfd, 0x3e, 0x0a, 0xad, 0x21, 0x7e, 0x0b, 0x10, 0x50,
	0x5a, 0xf3, 0x11, 0xb7, 0xa7, 0x71, 0x7b, 0x48, 0x4a, 0x6b, 0x48, 0xbc, 0x40, 0x9e, 0x58, 0x3d,
	0xfe, 0xcc, 0x18, 0x88, 0x3f, 0x42, 0x26, 0x47, 0x2d, 0x67, 0x3d, 0x92, 0xb2, 0x26, 0xe2, 0xb2,
	0xe6, 0xb2, 0x5e, 0x82, 0x6c, 0x7d, 0xfb, 0xed, 0x47, 0x51, 0xc1, 0xc5, 0x0a, 0x0a, 0xce, 0x9a,
	0x80, 0xbc, 0x04, 0xe6, 0x31, 0xcc, 0x7a, 0xcd, 0xb4, 0x4c, 0x4d, 0x07, 0xf9, 0x10, 0x69, 0x03,
	0xfa, 0x5f, 0x25, 0x91, 0x3f, 0x01, 0x4b, 0xa7, 0xbc, 0xf8, 0xef, 0xb3, 0x64, 0xba, 0x2e, 0x8f,
	0xe5, 0xe4, 0x20, 0x4e, 0x9e, 0xe1, 0x6d, 0xbf, 0x76, 0x35, 0xb1, 0xd8, 0xd2, 0xc3, 0x5f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xf6, 0x6a, 0x11, 0xe3, 0x3e, 0x01, 0x00, 0x00,
}
