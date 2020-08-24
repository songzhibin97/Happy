// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package pbUser

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

// =================== (注册) ====================
type RegisterRequest struct {
	// 注册请求
	UserName             string   `protobuf:"bytes,1,opt,name=UserName,proto3" json:"UserName,omitempty" validate:"required,gte=4,lt=20"`
	Password             string   `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty" validate:"required,gte=6,lt=20"`
	ConfirmPassword      string   `protobuf:"bytes,3,opt,name=ConfirmPassword,proto3" json:"ConfirmPassword,omitempty" validate:"required,eqfield=Password,gte=6,lt=20"`
	UserInfo             string   `protobuf:"bytes,4,opt,name=UserInfo,proto3" json:"UserInfo,omitempty"`
	Email                string   `protobuf:"bytes,5,opt,name=Email,proto3" json:"Email,omitempty" validate:"required,email"`
	VerificationCode     string   `protobuf:"bytes,6,opt,name=VerificationCode,proto3" json:"VerificationCode,omitempty" validate:"required"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *RegisterRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *RegisterRequest) GetConfirmPassword() string {
	if m != nil {
		return m.ConfirmPassword
	}
	return ""
}

func (m *RegisterRequest) GetUserInfo() string {
	if m != nil {
		return m.UserInfo
	}
	return ""
}

func (m *RegisterRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *RegisterRequest) GetVerificationCode() string {
	if m != nil {
		return m.VerificationCode
	}
	return ""
}

// =================== (登录) ====================
type LoginRequest struct {
	// 登录请求
	UserName             string   `protobuf:"bytes,1,opt,name=UserName,proto3" json:"UserName,omitempty" validate:"required,gte=4,lt=20"`
	Password             string   `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty" validate:"required,gte=4,lt=20"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

// ==============(验证码)=============
type VerificationRequest struct {
	// 验证码
	Email                string   `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty" validate:"required,email"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerificationRequest) Reset()         { *m = VerificationRequest{} }
func (m *VerificationRequest) String() string { return proto.CompactTextString(m) }
func (*VerificationRequest) ProtoMessage()    {}
func (*VerificationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *VerificationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerificationRequest.Unmarshal(m, b)
}
func (m *VerificationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerificationRequest.Marshal(b, m, deterministic)
}
func (m *VerificationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerificationRequest.Merge(m, src)
}
func (m *VerificationRequest) XXX_Size() int {
	return xxx_messageInfo_VerificationRequest.Size(m)
}
func (m *VerificationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VerificationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VerificationRequest proto.InternalMessageInfo

func (m *VerificationRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

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
	return fileDescriptor_116e343673f7ffaf, []int{3}
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
	proto.RegisterType((*RegisterRequest)(nil), "pbUser.RegisterRequest")
	proto.RegisterType((*LoginRequest)(nil), "pbUser.LoginRequest")
	proto.RegisterType((*VerificationRequest)(nil), "pbUser.VerificationRequest")
	proto.RegisterType((*Response)(nil), "pbUser.Response")
	proto.RegisterMapType((map[string]string)(nil), "pbUser.Response.DataEntry")
}

func init() {
	proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf)
}

var fileDescriptor_116e343673f7ffaf = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x41, 0x4f, 0xf2, 0x40,
	0x10, 0xfd, 0x96, 0xb6, 0x04, 0xe6, 0x23, 0x81, 0x8c, 0x24, 0x36, 0xf5, 0x42, 0x7a, 0x22, 0x9a,
	0x34, 0x11, 0x0f, 0x18, 0x2f, 0x1e, 0x10, 0x13, 0x13, 0x35, 0xa6, 0x89, 0xde, 0x17, 0x19, 0xc8,
	0x46, 0xe8, 0xe2, 0x6e, 0xd1, 0xf0, 0x53, 0xfc, 0x1d, 0xfe, 0x12, 0xff, 0x91, 0xd9, 0x2d, 0x2d,
	0x15, 0x7b, 0xf3, 0x36, 0x33, 0xef, 0xf5, 0xbd, 0xd7, 0x99, 0x05, 0x58, 0x6b, 0x52, 0xd1, 0x4a,
	0xc9, 0x54, 0x62, 0x7d, 0x35, 0x79, 0xd4, 0xa4, 0xc2, 0x2f, 0x06, 0xed, 0x98, 0xe6, 0x42, 0xa7,
	0xa4, 0x62, 0x7a, 0x5d, 0x93, 0x4e, 0x31, 0x80, 0x86, 0xc1, 0xee, 0xf9, 0x92, 0x7c, 0xd6, 0x63,
	0xfd, 0x66, 0x5c, 0xf4, 0x06, 0x7b, 0xe0, 0x5a, 0xbf, 0x4b, 0x35, 0xf5, 0x6b, 0x19, 0x96, 0xf7,
	0xd8, 0x87, 0xf6, 0x48, 0x26, 0x33, 0xa1, 0x96, 0x05, 0xc5, 0xb1, 0x94, 0xfd, 0x71, 0xee, 0x70,
	0x93, 0xcc, 0xa4, 0xef, 0xee, 0x1c, 0x4c, 0x8f, 0x5d, 0xf0, 0xc6, 0x4b, 0x2e, 0x16, 0xbe, 0x67,
	0x81, 0xac, 0xc1, 0x63, 0xe8, 0x3c, 0x91, 0x12, 0x33, 0xf1, 0xcc, 0x53, 0x21, 0x93, 0x91, 0x9c,
	0x92, 0x5f, 0xb7, 0x84, 0x5f, 0xf3, 0xf0, 0x1a, 0x5a, 0xb7, 0x72, 0x2e, 0x92, 0x3f, 0xfe, 0x4f,
	0x78, 0x02, 0x07, 0x65, 0xed, 0x5c, 0xae, 0x08, 0xc8, 0x4a, 0x01, 0xc3, 0x0f, 0x06, 0x8d, 0x98,
	0xf4, 0x4a, 0x26, 0x9a, 0x10, 0xc1, 0xb5, 0x09, 0x0d, 0xc3, 0x8b, 0x6d, 0x8d, 0x1d, 0x70, 0xee,
	0xf4, 0x7c, 0x6b, 0x62, 0x4a, 0x8c, 0xc0, 0xbd, 0xe2, 0x29, 0xf7, 0x9d, 0x9e, 0xd3, 0xff, 0x3f,
	0x08, 0xa2, 0xec, 0x24, 0x51, 0xae, 0x12, 0x19, 0x70, 0x9c, 0xa4, 0x6a, 0x13, 0x5b, 0x5e, 0x30,
	0x84, 0x66, 0x31, 0x32, 0x72, 0x2f, 0xb4, 0xd9, 0x66, 0x30, 0xa5, 0xc9, 0xf5, 0xc6, 0x17, 0x6b,
	0xda, 0x5a, 0x64, 0xcd, 0x45, 0xed, 0x9c, 0x0d, 0x3e, 0x19, 0xb8, 0x46, 0x1a, 0x87, 0x26, 0x63,
	0x76, 0x6c, 0x3c, 0xdc, 0xf9, 0xfd, 0x38, 0x7f, 0xd0, 0xd9, 0x0f, 0x12, 0xfe, 0xc3, 0x53, 0xf0,
	0xec, 0x4a, 0xb1, 0x9b, 0x83, 0xe5, 0x0d, 0x57, 0x7e, 0x72, 0x09, 0xad, 0xf2, 0xf6, 0xf0, 0x28,
	0xe7, 0x54, 0xec, 0xb4, 0x4a, 0x60, 0x52, 0xb7, 0x2f, 0xf5, 0xec, 0x3b, 0x00, 0x00, 0xff, 0xff,
	0xd7, 0x62, 0x6b, 0x0c, 0xb7, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*Response, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*Response, error)
	Verification(ctx context.Context, in *VerificationRequest, opts ...grpc.CallOption) (*Response, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pbUser.User/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pbUser.User/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Verification(ctx context.Context, in *VerificationRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pbUser.User/Verification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	Register(context.Context, *RegisterRequest) (*Response, error)
	Login(context.Context, *LoginRequest) (*Response, error)
	Verification(context.Context, *VerificationRequest) (*Response, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) Register(ctx context.Context, req *RegisterRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedUserServer) Login(ctx context.Context, req *LoginRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedUserServer) Verification(ctx context.Context, req *VerificationRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verification not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbUser.User/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbUser.User/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Verification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Verification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbUser.User/Verification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Verification(ctx, req.(*VerificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pbUser.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _User_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _User_Login_Handler,
		},
		{
			MethodName: "Verification",
			Handler:    _User_Verification_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}