// Code generated by protoc-gen-go. DO NOT EDIT.
// source: authentication.proto

package pb

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

type AuthenticationRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticationRequest) Reset()         { *m = AuthenticationRequest{} }
func (m *AuthenticationRequest) String() string { return proto.CompactTextString(m) }
func (*AuthenticationRequest) ProtoMessage()    {}
func (*AuthenticationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0dbc99083440df2, []int{0}
}

func (m *AuthenticationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticationRequest.Unmarshal(m, b)
}
func (m *AuthenticationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticationRequest.Marshal(b, m, deterministic)
}
func (m *AuthenticationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticationRequest.Merge(m, src)
}
func (m *AuthenticationRequest) XXX_Size() int {
	return xxx_messageInfo_AuthenticationRequest.Size(m)
}
func (m *AuthenticationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticationRequest proto.InternalMessageInfo

func (m *AuthenticationRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AuthenticationRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type AuthenticationResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticationResponse) Reset()         { *m = AuthenticationResponse{} }
func (m *AuthenticationResponse) String() string { return proto.CompactTextString(m) }
func (*AuthenticationResponse) ProtoMessage()    {}
func (*AuthenticationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0dbc99083440df2, []int{1}
}

func (m *AuthenticationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticationResponse.Unmarshal(m, b)
}
func (m *AuthenticationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticationResponse.Marshal(b, m, deterministic)
}
func (m *AuthenticationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticationResponse.Merge(m, src)
}
func (m *AuthenticationResponse) XXX_Size() int {
	return xxx_messageInfo_AuthenticationResponse.Size(m)
}
func (m *AuthenticationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticationResponse proto.InternalMessageInfo

func (m *AuthenticationResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*AuthenticationRequest)(nil), "pb.AuthenticationRequest")
	proto.RegisterType((*AuthenticationResponse)(nil), "pb.AuthenticationResponse")
}

func init() { proto.RegisterFile("authentication.proto", fileDescriptor_d0dbc99083440df2) }

var fileDescriptor_d0dbc99083440df2 = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x49, 0x2c, 0x2d, 0xc9,
	0x48, 0xcd, 0x2b, 0xc9, 0x4c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x62, 0x2a, 0x48, 0x52, 0xf2, 0xe7, 0x12, 0x75, 0x44, 0x91, 0x0b, 0x4a, 0x2d, 0x2c, 0x4d,
	0x2d, 0x2e, 0x11, 0x92, 0xe2, 0xe2, 0x28, 0x2d, 0x4e, 0x2d, 0xca, 0x4b, 0xcc, 0x4d, 0x95, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf3, 0x41, 0x72, 0x05, 0x89, 0xc5, 0xc5, 0xe5, 0xf9, 0x45,
	0x29, 0x12, 0x4c, 0x10, 0x39, 0x18, 0x5f, 0xc9, 0x80, 0x4b, 0x0c, 0xdd, 0xc0, 0xe2, 0x82, 0xfc,
	0xbc, 0xe2, 0x54, 0x21, 0x31, 0x2e, 0xb6, 0xa2, 0xd4, 0xe2, 0xd2, 0x9c, 0x12, 0xa8, 0x79, 0x50,
	0x9e, 0x51, 0x22, 0xba, 0x13, 0x82, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x85, 0x3c, 0xb8, 0xf8,
	0x9d, 0x8b, 0x52, 0x53, 0x40, 0x12, 0x89, 0x39, 0xce, 0x19, 0xa9, 0xc9, 0xd9, 0x42, 0x92, 0x7a,
	0x05, 0x49, 0x7a, 0x58, 0x1d, 0x2c, 0x25, 0x85, 0x4d, 0x0a, 0x62, 0x75, 0x12, 0x1b, 0xd8, 0xc3,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x60, 0x66, 0x3c, 0x99, 0x08, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthenticationServiceClient is the client API for AuthenticationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthenticationServiceClient interface {
	CredentialCheck(ctx context.Context, in *AuthenticationRequest, opts ...grpc.CallOption) (*AuthenticationResponse, error)
}

type authenticationServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthenticationServiceClient(cc *grpc.ClientConn) AuthenticationServiceClient {
	return &authenticationServiceClient{cc}
}

func (c *authenticationServiceClient) CredentialCheck(ctx context.Context, in *AuthenticationRequest, opts ...grpc.CallOption) (*AuthenticationResponse, error) {
	out := new(AuthenticationResponse)
	err := c.cc.Invoke(ctx, "/pb.AuthenticationService/CredentialCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationServiceServer is the server API for AuthenticationService service.
type AuthenticationServiceServer interface {
	CredentialCheck(context.Context, *AuthenticationRequest) (*AuthenticationResponse, error)
}

// UnimplementedAuthenticationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthenticationServiceServer struct {
}

func (*UnimplementedAuthenticationServiceServer) CredentialCheck(ctx context.Context, req *AuthenticationRequest) (*AuthenticationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CredentialCheck not implemented")
}

func RegisterAuthenticationServiceServer(s *grpc.Server, srv AuthenticationServiceServer) {
	s.RegisterService(&_AuthenticationService_serviceDesc, srv)
}

func _AuthenticationService_CredentialCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).CredentialCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AuthenticationService/CredentialCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).CredentialCheck(ctx, req.(*AuthenticationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthenticationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.AuthenticationService",
	HandlerType: (*AuthenticationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CredentialCheck",
			Handler:    _AuthenticationService_CredentialCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authentication.proto",
}
