// Code generated by protoc-gen-go. DO NOT EDIT.
// source: community.proto

package pbCommunity

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

// ============ 获取社区列表 =============
type CommunityListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommunityListRequest) Reset()         { *m = CommunityListRequest{} }
func (m *CommunityListRequest) String() string { return proto.CompactTextString(m) }
func (*CommunityListRequest) ProtoMessage()    {}
func (*CommunityListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_857922b7acda88b9, []int{0}
}

func (m *CommunityListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommunityListRequest.Unmarshal(m, b)
}
func (m *CommunityListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommunityListRequest.Marshal(b, m, deterministic)
}
func (m *CommunityListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommunityListRequest.Merge(m, src)
}
func (m *CommunityListRequest) XXX_Size() int {
	return xxx_messageInfo_CommunityListRequest.Size(m)
}
func (m *CommunityListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CommunityListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CommunityListRequest proto.InternalMessageInfo

// =========== 社区分类详情 ============
type CommunityDetailRequest struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty" validate:"required"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommunityDetailRequest) Reset()         { *m = CommunityDetailRequest{} }
func (m *CommunityDetailRequest) String() string { return proto.CompactTextString(m) }
func (*CommunityDetailRequest) ProtoMessage()    {}
func (*CommunityDetailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_857922b7acda88b9, []int{1}
}

func (m *CommunityDetailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommunityDetailRequest.Unmarshal(m, b)
}
func (m *CommunityDetailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommunityDetailRequest.Marshal(b, m, deterministic)
}
func (m *CommunityDetailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommunityDetailRequest.Merge(m, src)
}
func (m *CommunityDetailRequest) XXX_Size() int {
	return xxx_messageInfo_CommunityDetailRequest.Size(m)
}
func (m *CommunityDetailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CommunityDetailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CommunityDetailRequest proto.InternalMessageInfo

func (m *CommunityDetailRequest) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

// Response
type Response struct {
	// 响应
	Code                 int32             `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg                  string            `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	Data                 map[string]string `protobuf:"bytes,3,rep,name=Data,proto3" json:"Data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_857922b7acda88b9, []int{2}
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

func (m *Response) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *Response) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*CommunityListRequest)(nil), "pbCommunity.CommunityListRequest")
	proto.RegisterType((*CommunityDetailRequest)(nil), "pbCommunity.CommunityDetailRequest")
	proto.RegisterType((*Response)(nil), "pbCommunity.Response")
	proto.RegisterMapType((map[string]string)(nil), "pbCommunity.Response.DataEntry")
}

func init() {
	proto.RegisterFile("community.proto", fileDescriptor_857922b7acda88b9)
}

var fileDescriptor_857922b7acda88b9 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0xce, 0xcf, 0xcd,
	0x2d, 0xcd, 0xcb, 0x2c, 0xa9, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2e, 0x48, 0x72,
	0x86, 0x09, 0x29, 0x89, 0x71, 0x89, 0xc0, 0x39, 0x3e, 0x99, 0xc5, 0x25, 0x41, 0xa9, 0x85, 0xa5,
	0xa9, 0xc5, 0x25, 0x4a, 0x1a, 0x5c, 0x62, 0x70, 0x71, 0x97, 0xd4, 0x92, 0xc4, 0xcc, 0x1c, 0xa8,
	0x8c, 0x10, 0x1f, 0x17, 0x93, 0xa7, 0x8b, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x73, 0x10, 0x93, 0xa7,
	0x8b, 0xd2, 0x3c, 0x46, 0x2e, 0x8e, 0xa0, 0xd4, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x21,
	0x2e, 0x16, 0xe7, 0xfc, 0x94, 0x54, 0xb0, 0x34, 0x6b, 0x10, 0x98, 0x2d, 0x24, 0xc0, 0xc5, 0xec,
	0x5b, 0x9c, 0x2e, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x62, 0x0a, 0x19, 0x73, 0xb1, 0xb8,
	0x24, 0x96, 0x24, 0x4a, 0x30, 0x2b, 0x30, 0x6b, 0x70, 0x1b, 0xc9, 0xeb, 0x21, 0x39, 0x48, 0x0f,
	0x66, 0x94, 0x1e, 0x48, 0x85, 0x6b, 0x5e, 0x49, 0x51, 0x65, 0x10, 0x58, 0xb1, 0x94, 0x39, 0x17,
	0x27, 0x5c, 0x08, 0x64, 0x66, 0x76, 0x6a, 0x25, 0xd8, 0x1a, 0xce, 0x20, 0x10, 0x53, 0x48, 0x84,
	0x8b, 0xb5, 0x2c, 0x31, 0xa7, 0x34, 0x15, 0x6a, 0x0f, 0x84, 0x63, 0xc5, 0x64, 0xc1, 0x68, 0xb4,
	0x92, 0x91, 0x8b, 0x13, 0x6e, 0xbe, 0x90, 0x37, 0x17, 0x2f, 0x8a, 0x87, 0x85, 0x14, 0x51, 0xac,
	0xc7, 0x16, 0x18, 0x52, 0xa2, 0x58, 0x5d, 0xa8, 0xc4, 0x20, 0xe4, 0xcf, 0xc5, 0x8f, 0x16, 0x4a,
	0x42, 0xca, 0xd8, 0x8d, 0x43, 0x09, 0x43, 0x9c, 0x06, 0x26, 0xb1, 0x81, 0xa3, 0xc8, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0xdf, 0xef, 0xe4, 0x5b, 0xb5, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CommunityClient is the client API for Community service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommunityClient interface {
	CommunityList(ctx context.Context, in *CommunityListRequest, opts ...grpc.CallOption) (*Response, error)
	CommunityDetail(ctx context.Context, in *CommunityDetailRequest, opts ...grpc.CallOption) (*Response, error)
}

type communityClient struct {
	cc grpc.ClientConnInterface
}

func NewCommunityClient(cc grpc.ClientConnInterface) CommunityClient {
	return &communityClient{cc}
}

func (c *communityClient) CommunityList(ctx context.Context, in *CommunityListRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pbCommunity.Community/CommunityList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communityClient) CommunityDetail(ctx context.Context, in *CommunityDetailRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pbCommunity.Community/CommunityDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommunityServer is the server API for Community service.
type CommunityServer interface {
	CommunityList(context.Context, *CommunityListRequest) (*Response, error)
	CommunityDetail(context.Context, *CommunityDetailRequest) (*Response, error)
}

// UnimplementedCommunityServer can be embedded to have forward compatible implementations.
type UnimplementedCommunityServer struct {
}

func (*UnimplementedCommunityServer) CommunityList(ctx context.Context, req *CommunityListRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommunityList not implemented")
}
func (*UnimplementedCommunityServer) CommunityDetail(ctx context.Context, req *CommunityDetailRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommunityDetail not implemented")
}

func RegisterCommunityServer(s *grpc.Server, srv CommunityServer) {
	s.RegisterService(&_Community_serviceDesc, srv)
}

func _Community_CommunityList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommunityListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServer).CommunityList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbCommunity.Community/CommunityList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServer).CommunityList(ctx, req.(*CommunityListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Community_CommunityDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommunityDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServer).CommunityDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbCommunity.Community/CommunityDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServer).CommunityDetail(ctx, req.(*CommunityDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Community_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pbCommunity.Community",
	HandlerType: (*CommunityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CommunityList",
			Handler:    _Community_CommunityList_Handler,
		},
		{
			MethodName: "CommunityDetail",
			Handler:    _Community_CommunityDetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "community.proto",
}
