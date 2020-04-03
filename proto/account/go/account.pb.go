// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/account/account.proto

package _go

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
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

type User struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	PasswordHash         []byte   `protobuf:"bytes,3,opt,name=password_hash,json=passwordHash,proto3" json:"password_hash,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	CompanyId            uint64   `protobuf:"varint,5,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d492a0187472a3b, []int{0}
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

func (m *User) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPasswordHash() []byte {
	if m != nil {
		return m.PasswordHash
	}
	return nil
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetCompanyId() uint64 {
	if m != nil {
		return m.CompanyId
	}
	return 0
}

type CreateUserRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             []byte   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d492a0187472a3b, []int{1}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CreateUserRequest) GetPassword() []byte {
	if m != nil {
		return m.Password
	}
	return nil
}

type CreateUserResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserResponse) Reset()         { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()    {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d492a0187472a3b, []int{2}
}

func (m *CreateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResponse.Unmarshal(m, b)
}
func (m *CreateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResponse.Marshal(b, m, deterministic)
}
func (m *CreateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResponse.Merge(m, src)
}
func (m *CreateUserResponse) XXX_Size() int {
	return xxx_messageInfo_CreateUserResponse.Size(m)
}
func (m *CreateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResponse proto.InternalMessageInfo

func (m *CreateUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*CreateUserRequest)(nil), "user.CreateUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "user.CreateUserResponse")
}

func init() { proto.RegisterFile("proto/account/account.proto", fileDescriptor_5d492a0187472a3b) }

var fileDescriptor_5d492a0187472a3b = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xdf, 0x8b, 0xd3, 0x40,
	0x10, 0xc7, 0xd9, 0x9a, 0x56, 0x3b, 0xad, 0x82, 0x8b, 0x60, 0xa8, 0x68, 0x6b, 0x15, 0x89, 0x96,
	0x24, 0x1a, 0x7f, 0x40, 0x9f, 0x94, 0x88, 0xa0, 0x08, 0x3e, 0x44, 0x84, 0xa2, 0xd5, 0xb2, 0x4d,
	0x96, 0x64, 0xb1, 0xc9, 0xe6, 0x76, 0x37, 0xcd, 0xdd, 0xfb, 0x3d, 0xdc, 0x7f, 0x59, 0xe8, 0x5f,
	0x72, 0x64, 0xd3, 0xf4, 0x5a, 0xfa, 0x94, 0xf9, 0x7e, 0x67, 0x26, 0xf3, 0x99, 0xdd, 0x85, 0x47,
	0xb9, 0xe0, 0x8a, 0xbb, 0x24, 0x0c, 0x79, 0x91, 0xa9, 0xe6, 0xeb, 0x68, 0x17, 0x1b, 0x85, 0xa4,
	0x62, 0xf0, 0x21, 0x66, 0x2a, 0x29, 0x96, 0x4e, 0xc8, 0x53, 0x37, 0x2d, 0x99, 0xfa, 0xcf, 0x4b,
	0x37, 0xe6, 0xb6, 0x2e, 0xb1, 0xd7, 0x64, 0xc5, 0x22, 0xa2, 0xb8, 0x90, 0xee, 0x3e, 0xac, 0xbb,
	0xc7, 0x97, 0x08, 0x8c, 0x5f, 0x92, 0x0a, 0x7c, 0x0f, 0x5a, 0x2c, 0x32, 0xd1, 0x08, 0x59, 0x46,
	0xd0, 0x62, 0x11, 0x7e, 0x00, 0x6d, 0x9a, 0x12, 0xb6, 0x32, 0x5b, 0x23, 0x64, 0x75, 0x83, 0x5a,
	0xe0, 0x67, 0x70, 0x37, 0x27, 0x52, 0x96, 0x5c, 0x44, 0x8b, 0x84, 0xc8, 0xc4, 0xbc, 0x35, 0x42,
	0x56, 0x3f, 0xe8, 0x37, 0xe6, 0x57, 0x22, 0x13, 0x8c, 0xc1, 0xc8, 0x48, 0x4a, 0x4d, 0x43, 0x77,
	0xea, 0x18, 0x3f, 0x06, 0x08, 0x79, 0x9a, 0x93, 0xec, 0x62, 0xc1, 0x22, 0xb3, 0xad, 0xc7, 0x74,
	0x77, 0xce, 0xb7, 0x68, 0x7c, 0x85, 0xe0, 0xfe, 0x67, 0x41, 0x89, 0xa2, 0x15, 0x4c, 0x40, 0xcf,
	0x0a, 0x2a, 0x15, 0xfe, 0xde, 0x30, 0x54, 0x58, 0x5d, 0xff, 0xfd, 0x76, 0x33, 0x7c, 0x03, 0xf6,
	0xbf, 0x79, 0x39, 0xb1, 0xfe, 0xd8, 0x13, 0xe7, 0xef, 0xbc, 0x9c, 0xbc, 0x7c, 0xf5, 0xa9, 0x96,
	0x3b, 0x35, 0x77, 0x8e, 0xe4, 0xf3, 0x19, 0x6a, 0xd0, 0x5f, 0xc0, 0x9d, 0x86, 0x52, 0xef, 0xd4,
	0xf7, 0x61, 0xbb, 0x19, 0x76, 0x66, 0x28, 0xbf, 0x7d, 0xfe, 0x34, 0xd8, 0xe7, 0xc6, 0xef, 0x00,
	0x1f, 0x92, 0xc8, 0x9c, 0x67, 0x92, 0xe2, 0x27, 0xa0, 0xcf, 0x59, 0x93, 0xf4, 0x3c, 0x70, 0x2a,
	0xe1, 0xe8, 0x0a, 0xed, 0x7b, 0x3f, 0xa0, 0x57, 0xa9, 0x9f, 0x54, 0xac, 0x59, 0x48, 0xf1, 0x47,
	0x80, 0x9b, 0x9f, 0xe0, 0x87, 0x75, 0xf9, 0xc9, 0x82, 0x03, 0xf3, 0x34, 0x51, 0xcf, 0xf3, 0xbd,
	0xdf, 0xaf, 0x0f, 0x6e, 0x94, 0x0b, 0x16, 0xb3, 0x6c, 0x49, 0xb3, 0x4c, 0xf1, 0xc2, 0xf5, 0xa6,
	0xde, 0xd4, 0xff, 0xe2, 0x1e, 0x3f, 0x8a, 0x98, 0x2f, 0x3b, 0xda, 0x79, 0x7b, 0x1d, 0x00, 0x00,
	0xff, 0xff, 0x35, 0xc5, 0x19, 0xa2, 0x2f, 0x02, 0x00, 0x00,
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
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/account/account.proto",
}