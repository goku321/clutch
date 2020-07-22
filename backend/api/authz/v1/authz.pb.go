// Code generated by protoc-gen-go. DO NOT EDIT.
// source: authz/v1/authz.proto

package authzv1

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	v1 "github.com/lyft/clutch/backend/api/api/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Decision int32

const (
	Decision_UNSPECIFIED Decision = 0
	Decision_DENY        Decision = 1
	Decision_ALLOW       Decision = 2
)

var Decision_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "DENY",
	2: "ALLOW",
}

var Decision_value = map[string]int32{
	"UNSPECIFIED": 0,
	"DENY":        1,
	"ALLOW":       2,
}

func (x Decision) String() string {
	return proto.EnumName(Decision_name, int32(x))
}

func (Decision) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a9f76f3270a35635, []int{0}
}

type Subject struct {
	User                 string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Groups               []string `protobuf:"bytes,2,rep,name=groups,proto3" json:"groups,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Subject) Reset()         { *m = Subject{} }
func (m *Subject) String() string { return proto.CompactTextString(m) }
func (*Subject) ProtoMessage()    {}
func (*Subject) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9f76f3270a35635, []int{0}
}

func (m *Subject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Subject.Unmarshal(m, b)
}
func (m *Subject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Subject.Marshal(b, m, deterministic)
}
func (m *Subject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subject.Merge(m, src)
}
func (m *Subject) XXX_Size() int {
	return xxx_messageInfo_Subject.Size(m)
}
func (m *Subject) XXX_DiscardUnknown() {
	xxx_messageInfo_Subject.DiscardUnknown(m)
}

var xxx_messageInfo_Subject proto.InternalMessageInfo

func (m *Subject) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Subject) GetGroups() []string {
	if m != nil {
		return m.Groups
	}
	return nil
}

type CheckRequest struct {
	Subject              *Subject      `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	Method               string        `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	ActionType           v1.ActionType `protobuf:"varint,3,opt,name=action_type,json=actionType,proto3,enum=clutch.api.v1.ActionType" json:"action_type,omitempty"`
	Resource             string        `protobuf:"bytes,4,opt,name=resource,proto3" json:"resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CheckRequest) Reset()         { *m = CheckRequest{} }
func (m *CheckRequest) String() string { return proto.CompactTextString(m) }
func (*CheckRequest) ProtoMessage()    {}
func (*CheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9f76f3270a35635, []int{1}
}

func (m *CheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckRequest.Unmarshal(m, b)
}
func (m *CheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckRequest.Marshal(b, m, deterministic)
}
func (m *CheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckRequest.Merge(m, src)
}
func (m *CheckRequest) XXX_Size() int {
	return xxx_messageInfo_CheckRequest.Size(m)
}
func (m *CheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckRequest proto.InternalMessageInfo

func (m *CheckRequest) GetSubject() *Subject {
	if m != nil {
		return m.Subject
	}
	return nil
}

func (m *CheckRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *CheckRequest) GetActionType() v1.ActionType {
	if m != nil {
		return m.ActionType
	}
	return v1.ActionType_UNSPECIFIED
}

func (m *CheckRequest) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

type CheckResponse struct {
	Decision             Decision `protobuf:"varint,1,opt,name=decision,proto3,enum=clutch.authz.v1.Decision" json:"decision,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckResponse) Reset()         { *m = CheckResponse{} }
func (m *CheckResponse) String() string { return proto.CompactTextString(m) }
func (*CheckResponse) ProtoMessage()    {}
func (*CheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9f76f3270a35635, []int{2}
}

func (m *CheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckResponse.Unmarshal(m, b)
}
func (m *CheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckResponse.Marshal(b, m, deterministic)
}
func (m *CheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckResponse.Merge(m, src)
}
func (m *CheckResponse) XXX_Size() int {
	return xxx_messageInfo_CheckResponse.Size(m)
}
func (m *CheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckResponse proto.InternalMessageInfo

func (m *CheckResponse) GetDecision() Decision {
	if m != nil {
		return m.Decision
	}
	return Decision_UNSPECIFIED
}

func init() {
	proto.RegisterEnum("clutch.authz.v1.Decision", Decision_name, Decision_value)
	proto.RegisterType((*Subject)(nil), "clutch.authz.v1.Subject")
	proto.RegisterType((*CheckRequest)(nil), "clutch.authz.v1.CheckRequest")
	proto.RegisterType((*CheckResponse)(nil), "clutch.authz.v1.CheckResponse")
}

func init() {
	proto.RegisterFile("authz/v1/authz.proto", fileDescriptor_a9f76f3270a35635)
}

var fileDescriptor_a9f76f3270a35635 = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xcd, 0x6e, 0xd3, 0x40,
	0x18, 0x64, 0xdd, 0xa4, 0xb1, 0xbf, 0x00, 0xb1, 0x96, 0x0a, 0x8c, 0x55, 0x90, 0xe5, 0x53, 0xd4,
	0x83, 0x4d, 0x82, 0x7a, 0xa9, 0xb8, 0x38, 0x4d, 0x2a, 0x45, 0xaa, 0x4a, 0xe5, 0x80, 0x10, 0x5c,
	0xd0, 0x66, 0xb3, 0x8a, 0x0d, 0x89, 0xd7, 0x78, 0xd7, 0x96, 0xc2, 0x91, 0x23, 0x57, 0xde, 0x82,
	0x67, 0xe0, 0x2d, 0x78, 0x02, 0x24, 0x9e, 0x82, 0x13, 0xf2, 0xfa, 0x47, 0x11, 0x11, 0xb7, 0xef,
	0x9b, 0x99, 0x9d, 0x99, 0xb5, 0x17, 0x4e, 0x48, 0x2e, 0xa3, 0xcf, 0x7e, 0x31, 0xf2, 0xd5, 0xe0,
	0xa5, 0x19, 0x97, 0x1c, 0x0f, 0xe8, 0x26, 0x97, 0x34, 0xf2, 0x2a, 0xac, 0x18, 0xd9, 0xa7, 0x6b,
	0xce, 0xd7, 0x1b, 0xe6, 0x93, 0x34, 0xf6, 0x49, 0x92, 0x70, 0x49, 0x64, 0xcc, 0x13, 0x51, 0xc9,
	0x6d, 0xab, 0x84, 0x4b, 0x8b, 0x03, 0xe6, 0x41, 0xcd, 0x08, 0x1a, 0xb1, 0x2d, 0xa9, 0xc1, 0x47,
	0x05, 0xd9, 0xc4, 0x2b, 0x22, 0x99, 0xdf, 0x0c, 0x15, 0xe1, 0x9e, 0x43, 0x6f, 0x91, 0x2f, 0x3f,
	0x30, 0x2a, 0x31, 0x86, 0x4e, 0x2e, 0x58, 0x66, 0x21, 0x07, 0x0d, 0x8d, 0x50, 0xcd, 0xf8, 0x21,
	0x1c, 0xaf, 0x33, 0x9e, 0xa7, 0xc2, 0xd2, 0x9c, 0xa3, 0xa1, 0x11, 0xd6, 0x9b, 0xfb, 0x03, 0xc1,
	0xdd, 0xcb, 0x88, 0xd1, 0x8f, 0x21, 0xfb, 0x94, 0x33, 0x21, 0xf1, 0x0b, 0xe8, 0x89, 0xca, 0x47,
	0x9d, 0xef, 0x8f, 0x2d, 0xef, 0x9f, 0x0b, 0x79, 0x75, 0xce, 0x44, 0xff, 0x33, 0xe9, 0x7e, 0x45,
	0x9a, 0x89, 0xc2, 0xe6, 0x48, 0x19, 0xb3, 0x65, 0x32, 0xe2, 0x2b, 0x4b, 0x53, 0xe1, 0xf5, 0x86,
	0x2f, 0xa0, 0x4f, 0x68, 0x79, 0xb9, 0xf7, 0x72, 0x97, 0x32, 0xeb, 0xc8, 0x41, 0xc3, 0xfb, 0xe3,
	0xc7, 0xad, 0x73, 0x1a, 0x97, 0xbe, 0x81, 0x52, 0xbc, 0xda, 0xa5, 0x2c, 0x04, 0xd2, 0xce, 0xd8,
	0x06, 0x3d, 0x63, 0x82, 0xe7, 0x19, 0x65, 0x56, 0x47, 0xb9, 0xb6, 0xbb, 0x7b, 0x05, 0xf7, 0xea,
	0xf6, 0x22, 0xe5, 0x89, 0x60, 0xf8, 0x1c, 0xf4, 0x15, 0xa3, 0xb1, 0x88, 0x79, 0xa2, 0xfa, 0xef,
	0xa7, 0x34, 0xfd, 0xa7, 0xb5, 0x20, 0x6c, 0xa5, 0x67, 0xcf, 0x40, 0x6f, 0x50, 0x3c, 0x80, 0xfe,
	0xeb, 0x9b, 0xc5, 0xed, 0xec, 0x72, 0x7e, 0x35, 0x9f, 0x4d, 0xcd, 0x3b, 0x58, 0x87, 0xce, 0x74,
	0x76, 0xf3, 0xd6, 0x44, 0xd8, 0x80, 0x6e, 0x70, 0x7d, 0xfd, 0xf2, 0x8d, 0xa9, 0x8d, 0x25, 0xe8,
	0x41, 0x69, 0x18, 0xdc, 0xce, 0x71, 0x04, 0x5d, 0xd5, 0x02, 0x3f, 0x39, 0xc8, 0xda, 0xff, 0xb6,
	0xf6, 0xd3, 0xff, 0xd1, 0x55, 0x79, 0xd7, 0xf9, 0xfe, 0xeb, 0x54, 0xd3, 0xb5, 0x2f, 0x3f, 0x7f,
	0x7f, 0xd3, 0x4e, 0xdc, 0x41, 0xfb, 0xb2, 0x7c, 0x5a, 0xca, 0x2e, 0xd0, 0xd9, 0xc4, 0x78, 0xd7,
	0x53, 0x48, 0x31, 0x5a, 0x1e, 0xab, 0xff, 0xfe, 0xfc, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xab,
	0x5c, 0xae, 0xe3, 0x86, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthzAPIClient is the client API for AuthzAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthzAPIClient interface {
	Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error)
}

type authzAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthzAPIClient(cc grpc.ClientConnInterface) AuthzAPIClient {
	return &authzAPIClient{cc}
}

func (c *authzAPIClient) Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, "/clutch.authz.v1.AuthzAPI/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthzAPIServer is the server API for AuthzAPI service.
type AuthzAPIServer interface {
	Check(context.Context, *CheckRequest) (*CheckResponse, error)
}

// UnimplementedAuthzAPIServer can be embedded to have forward compatible implementations.
type UnimplementedAuthzAPIServer struct {
}

func (*UnimplementedAuthzAPIServer) Check(ctx context.Context, req *CheckRequest) (*CheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}

func RegisterAuthzAPIServer(s *grpc.Server, srv AuthzAPIServer) {
	s.RegisterService(&_AuthzAPI_serviceDesc, srv)
}

func _AuthzAPI_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthzAPIServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.authz.v1.AuthzAPI/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthzAPIServer).Check(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthzAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "clutch.authz.v1.AuthzAPI",
	HandlerType: (*AuthzAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _AuthzAPI_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authz/v1/authz.proto",
}
