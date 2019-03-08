// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api-pb/types.proto

package apipb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

type User struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_710d12a8008818d7, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type AddUserRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddUserRequest) Reset()         { *m = AddUserRequest{} }
func (m *AddUserRequest) String() string { return proto.CompactTextString(m) }
func (*AddUserRequest) ProtoMessage()    {}
func (*AddUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_710d12a8008818d7, []int{1}
}

func (m *AddUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddUserRequest.Unmarshal(m, b)
}
func (m *AddUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddUserRequest.Marshal(b, m, deterministic)
}
func (m *AddUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddUserRequest.Merge(m, src)
}
func (m *AddUserRequest) XXX_Size() int {
	return xxx_messageInfo_AddUserRequest.Size(m)
}
func (m *AddUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddUserRequest proto.InternalMessageInfo

func (m *AddUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type AddUserResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddUserResponse) Reset()         { *m = AddUserResponse{} }
func (m *AddUserResponse) String() string { return proto.CompactTextString(m) }
func (*AddUserResponse) ProtoMessage()    {}
func (*AddUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_710d12a8008818d7, []int{2}
}

func (m *AddUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddUserResponse.Unmarshal(m, b)
}
func (m *AddUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddUserResponse.Marshal(b, m, deterministic)
}
func (m *AddUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddUserResponse.Merge(m, src)
}
func (m *AddUserResponse) XXX_Size() int {
	return xxx_messageInfo_AddUserResponse.Size(m)
}
func (m *AddUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddUserResponse proto.InternalMessageInfo

type GetUserRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_710d12a8008818d7, []int{3}
}

func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (m *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(m, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type GetUserResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserResponse) Reset()         { *m = GetUserResponse{} }
func (m *GetUserResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserResponse) ProtoMessage()    {}
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_710d12a8008818d7, []int{4}
}

func (m *GetUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserResponse.Unmarshal(m, b)
}
func (m *GetUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserResponse.Marshal(b, m, deterministic)
}
func (m *GetUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserResponse.Merge(m, src)
}
func (m *GetUserResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserResponse.Size(m)
}
func (m *GetUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserResponse proto.InternalMessageInfo

func (m *GetUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "apipb.User")
	proto.RegisterType((*AddUserRequest)(nil), "apipb.AddUserRequest")
	proto.RegisterType((*AddUserResponse)(nil), "apipb.AddUserResponse")
	proto.RegisterType((*GetUserRequest)(nil), "apipb.GetUserRequest")
	proto.RegisterType((*GetUserResponse)(nil), "apipb.GetUserResponse")
}

func init() { proto.RegisterFile("api-pb/types.proto", fileDescriptor_710d12a8008818d7) }

var fileDescriptor_710d12a8008818d7 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xe9, 0xb2, 0x6a, 0x9d, 0xe2, 0x2e, 0x46, 0x94, 0xa5, 0x08, 0x2e, 0x39, 0x89, 0x68,
	0x8b, 0x15, 0x3c, 0x78, 0xf3, 0xe4, 0xc9, 0xcb, 0x8a, 0x0f, 0x90, 0x6e, 0xc6, 0x12, 0xd0, 0x66,
	0x4c, 0x52, 0xc5, 0xab, 0xaf, 0xe0, 0x63, 0xf8, 0x38, 0xbe, 0x82, 0x0f, 0x22, 0xcd, 0x86, 0xd0,
	0xdd, 0x8b, 0xb7, 0xfc, 0x33, 0xff, 0x7c, 0xf9, 0x87, 0x01, 0x26, 0x48, 0x5d, 0x50, 0x5d, 0xba,
	0x0f, 0x42, 0x5b, 0x90, 0xd1, 0x4e, 0xb3, 0x2d, 0x41, 0x8a, 0xea, 0xfc, 0xb8, 0xd1, 0xba, 0x79,
	0xc6, 0x52, 0x90, 0x2a, 0x45, 0xdb, 0x6a, 0x27, 0x9c, 0xd2, 0x6d, 0x30, 0xe5, 0xf3, 0xd0, 0xf5,
	0xaa, 0xee, 0x9e, 0x4a, 0x89, 0x76, 0x69, 0x14, 0x39, 0x6d, 0x56, 0x0e, 0x7e, 0x0d, 0xe3, 0x47,
	0x8b, 0x86, 0x31, 0x18, 0xb7, 0xe2, 0x05, 0x67, 0xc9, 0x3c, 0x39, 0xdd, 0x5d, 0xf8, 0x37, 0xcb,
	0x21, 0x25, 0x61, 0xed, 0xbb, 0x36, 0x72, 0x36, 0xf2, 0xf5, 0xa8, 0xf9, 0x25, 0x4c, 0x6e, 0xa5,
	0xec, 0x47, 0x17, 0xf8, 0xda, 0xa1, 0x75, 0xec, 0x04, 0xc6, 0x9d, 0x45, 0xe3, 0x09, 0x59, 0x95,
	0x15, 0x3e, 0x5f, 0xe1, 0x1d, 0xbe, 0xc1, 0xf7, 0x61, 0x1a, 0x47, 0x2c, 0xe9, 0xd6, 0x22, 0x3f,
	0x87, 0xc9, 0x1d, 0xba, 0x21, 0x25, 0x87, 0xb4, 0x37, 0x0f, 0xb2, 0x44, 0xcd, 0x2b, 0x98, 0x46,
	0xf7, 0x0a, 0xf0, 0xef, 0xa7, 0xd5, 0x77, 0x02, 0x59, 0x2f, 0x1f, 0xd0, 0xbc, 0xa9, 0x25, 0xb2,
	0x7b, 0xd8, 0x09, 0x21, 0xd8, 0x61, 0x70, 0xaf, 0xef, 0x91, 0x1f, 0x6d, 0x96, 0x43, 0xd6, 0x83,
	0xcf, 0x9f, 0xdf, 0xaf, 0xd1, 0x1e, 0x4f, 0x4b, 0x21, 0x65, 0xcf, 0xbe, 0x49, 0xce, 0x7a, 0x5c,
	0x88, 0x14, 0x71, 0xeb, 0x0b, 0x45, 0xdc, 0x46, 0xf2, 0x01, 0xae, 0x41, 0x17, 0x70, 0xf5, 0xb6,
	0x3f, 0xca, 0xd5, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x73, 0xcc, 0x28, 0xa4, 0xf1, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*AddUserResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*AddUserResponse, error) {
	out := new(AddUserResponse)
	err := c.cc.Invoke(ctx, "/apipb.UserService/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/apipb.UserService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	AddUser(context.Context, *AddUserRequest) (*AddUserResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apipb.UserService/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddUser(ctx, req.(*AddUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apipb.UserService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "apipb.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUser",
			Handler:    _UserService_AddUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api-pb/types.proto",
}