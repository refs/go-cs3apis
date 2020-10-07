// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs3/ocm/invite/v1beta1/invite_api.proto

package invitev1beta1

import (
	context "context"
	fmt "fmt"
	v1beta13 "github.com/cs3org/go-cs3apis/cs3/identity/user/v1beta1"
	v1beta12 "github.com/cs3org/go-cs3apis/cs3/ocm/provider/v1beta1"
	v1beta11 "github.com/cs3org/go-cs3apis/cs3/rpc/v1beta1"
	v1beta1 "github.com/cs3org/go-cs3apis/cs3/types/v1beta1"
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

type GenerateInviteTokenRequest struct {
	// OPTIONAL.
	// Opaque information.
	Opaque               *v1beta1.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GenerateInviteTokenRequest) Reset()         { *m = GenerateInviteTokenRequest{} }
func (m *GenerateInviteTokenRequest) String() string { return proto.CompactTextString(m) }
func (*GenerateInviteTokenRequest) ProtoMessage()    {}
func (*GenerateInviteTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0e70953130e3677, []int{0}
}

func (m *GenerateInviteTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateInviteTokenRequest.Unmarshal(m, b)
}
func (m *GenerateInviteTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateInviteTokenRequest.Marshal(b, m, deterministic)
}
func (m *GenerateInviteTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateInviteTokenRequest.Merge(m, src)
}
func (m *GenerateInviteTokenRequest) XXX_Size() int {
	return xxx_messageInfo_GenerateInviteTokenRequest.Size(m)
}
func (m *GenerateInviteTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateInviteTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateInviteTokenRequest proto.InternalMessageInfo

func (m *GenerateInviteTokenRequest) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

type GenerateInviteTokenResponse struct {
	// REQUIRED.
	// The response status.
	Status *v1beta11.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The generated token.
	InviteToken          *InviteToken `protobuf:"bytes,3,opt,name=invite_token,json=inviteToken,proto3" json:"invite_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GenerateInviteTokenResponse) Reset()         { *m = GenerateInviteTokenResponse{} }
func (m *GenerateInviteTokenResponse) String() string { return proto.CompactTextString(m) }
func (*GenerateInviteTokenResponse) ProtoMessage()    {}
func (*GenerateInviteTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0e70953130e3677, []int{1}
}

func (m *GenerateInviteTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateInviteTokenResponse.Unmarshal(m, b)
}
func (m *GenerateInviteTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateInviteTokenResponse.Marshal(b, m, deterministic)
}
func (m *GenerateInviteTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateInviteTokenResponse.Merge(m, src)
}
func (m *GenerateInviteTokenResponse) XXX_Size() int {
	return xxx_messageInfo_GenerateInviteTokenResponse.Size(m)
}
func (m *GenerateInviteTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateInviteTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateInviteTokenResponse proto.InternalMessageInfo

func (m *GenerateInviteTokenResponse) GetStatus() *v1beta11.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *GenerateInviteTokenResponse) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *GenerateInviteTokenResponse) GetInviteToken() *InviteToken {
	if m != nil {
		return m.InviteToken
	}
	return nil
}

type ForwardInviteRequest struct {
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The token to authenticate with.
	InviteToken *InviteToken `protobuf:"bytes,2,opt,name=invite_token,json=inviteToken,proto3" json:"invite_token,omitempty"`
	// REQUIRED.
	// The details of the sync'n'share system provider of the user who sent the invite.
	OriginSystemProvider *v1beta12.ProviderInfo `protobuf:"bytes,3,opt,name=origin_system_provider,json=originSystemProvider,proto3" json:"origin_system_provider,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ForwardInviteRequest) Reset()         { *m = ForwardInviteRequest{} }
func (m *ForwardInviteRequest) String() string { return proto.CompactTextString(m) }
func (*ForwardInviteRequest) ProtoMessage()    {}
func (*ForwardInviteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0e70953130e3677, []int{2}
}

func (m *ForwardInviteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ForwardInviteRequest.Unmarshal(m, b)
}
func (m *ForwardInviteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ForwardInviteRequest.Marshal(b, m, deterministic)
}
func (m *ForwardInviteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForwardInviteRequest.Merge(m, src)
}
func (m *ForwardInviteRequest) XXX_Size() int {
	return xxx_messageInfo_ForwardInviteRequest.Size(m)
}
func (m *ForwardInviteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ForwardInviteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ForwardInviteRequest proto.InternalMessageInfo

func (m *ForwardInviteRequest) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *ForwardInviteRequest) GetInviteToken() *InviteToken {
	if m != nil {
		return m.InviteToken
	}
	return nil
}

func (m *ForwardInviteRequest) GetOriginSystemProvider() *v1beta12.ProviderInfo {
	if m != nil {
		return m.OriginSystemProvider
	}
	return nil
}

type ForwardInviteResponse struct {
	// REQUIRED.
	// The response status.
	Status *v1beta11.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque               *v1beta1.Opaque `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ForwardInviteResponse) Reset()         { *m = ForwardInviteResponse{} }
func (m *ForwardInviteResponse) String() string { return proto.CompactTextString(m) }
func (*ForwardInviteResponse) ProtoMessage()    {}
func (*ForwardInviteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0e70953130e3677, []int{3}
}

func (m *ForwardInviteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ForwardInviteResponse.Unmarshal(m, b)
}
func (m *ForwardInviteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ForwardInviteResponse.Marshal(b, m, deterministic)
}
func (m *ForwardInviteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForwardInviteResponse.Merge(m, src)
}
func (m *ForwardInviteResponse) XXX_Size() int {
	return xxx_messageInfo_ForwardInviteResponse.Size(m)
}
func (m *ForwardInviteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ForwardInviteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ForwardInviteResponse proto.InternalMessageInfo

func (m *ForwardInviteResponse) GetStatus() *v1beta11.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ForwardInviteResponse) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

type AcceptInviteRequest struct {
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The token to authenticate with.
	InviteToken *InviteToken `protobuf:"bytes,2,opt,name=invite_token,json=inviteToken,proto3" json:"invite_token,omitempty"`
	// REQUIRED.
	// The user who accepted the invite.
	RemoteUser           *v1beta13.User `protobuf:"bytes,3,opt,name=remote_user,json=remoteUser,proto3" json:"remote_user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AcceptInviteRequest) Reset()         { *m = AcceptInviteRequest{} }
func (m *AcceptInviteRequest) String() string { return proto.CompactTextString(m) }
func (*AcceptInviteRequest) ProtoMessage()    {}
func (*AcceptInviteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0e70953130e3677, []int{4}
}

func (m *AcceptInviteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AcceptInviteRequest.Unmarshal(m, b)
}
func (m *AcceptInviteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AcceptInviteRequest.Marshal(b, m, deterministic)
}
func (m *AcceptInviteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AcceptInviteRequest.Merge(m, src)
}
func (m *AcceptInviteRequest) XXX_Size() int {
	return xxx_messageInfo_AcceptInviteRequest.Size(m)
}
func (m *AcceptInviteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AcceptInviteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AcceptInviteRequest proto.InternalMessageInfo

func (m *AcceptInviteRequest) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *AcceptInviteRequest) GetInviteToken() *InviteToken {
	if m != nil {
		return m.InviteToken
	}
	return nil
}

func (m *AcceptInviteRequest) GetRemoteUser() *v1beta13.User {
	if m != nil {
		return m.RemoteUser
	}
	return nil
}

type AcceptInviteResponse struct {
	// REQUIRED.
	// The response status.
	Status *v1beta11.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque               *v1beta1.Opaque `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AcceptInviteResponse) Reset()         { *m = AcceptInviteResponse{} }
func (m *AcceptInviteResponse) String() string { return proto.CompactTextString(m) }
func (*AcceptInviteResponse) ProtoMessage()    {}
func (*AcceptInviteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0e70953130e3677, []int{5}
}

func (m *AcceptInviteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AcceptInviteResponse.Unmarshal(m, b)
}
func (m *AcceptInviteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AcceptInviteResponse.Marshal(b, m, deterministic)
}
func (m *AcceptInviteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AcceptInviteResponse.Merge(m, src)
}
func (m *AcceptInviteResponse) XXX_Size() int {
	return xxx_messageInfo_AcceptInviteResponse.Size(m)
}
func (m *AcceptInviteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AcceptInviteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AcceptInviteResponse proto.InternalMessageInfo

func (m *AcceptInviteResponse) GetStatus() *v1beta11.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *AcceptInviteResponse) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

type GetRemoteUserRequest struct {
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The id of the user.
	RemoteUserId         *v1beta13.UserId `protobuf:"bytes,2,opt,name=remote_user_id,json=remoteUserId,proto3" json:"remote_user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetRemoteUserRequest) Reset()         { *m = GetRemoteUserRequest{} }
func (m *GetRemoteUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetRemoteUserRequest) ProtoMessage()    {}
func (*GetRemoteUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0e70953130e3677, []int{6}
}

func (m *GetRemoteUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRemoteUserRequest.Unmarshal(m, b)
}
func (m *GetRemoteUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRemoteUserRequest.Marshal(b, m, deterministic)
}
func (m *GetRemoteUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRemoteUserRequest.Merge(m, src)
}
func (m *GetRemoteUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetRemoteUserRequest.Size(m)
}
func (m *GetRemoteUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRemoteUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRemoteUserRequest proto.InternalMessageInfo

func (m *GetRemoteUserRequest) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *GetRemoteUserRequest) GetRemoteUserId() *v1beta13.UserId {
	if m != nil {
		return m.RemoteUserId
	}
	return nil
}

type GetRemoteUserResponse struct {
	// REQUIRED.
	// The response status.
	Status *v1beta11.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The user information.
	RemoteUser           *v1beta13.User `protobuf:"bytes,3,opt,name=remote_user,json=remoteUser,proto3" json:"remote_user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetRemoteUserResponse) Reset()         { *m = GetRemoteUserResponse{} }
func (m *GetRemoteUserResponse) String() string { return proto.CompactTextString(m) }
func (*GetRemoteUserResponse) ProtoMessage()    {}
func (*GetRemoteUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0e70953130e3677, []int{7}
}

func (m *GetRemoteUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRemoteUserResponse.Unmarshal(m, b)
}
func (m *GetRemoteUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRemoteUserResponse.Marshal(b, m, deterministic)
}
func (m *GetRemoteUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRemoteUserResponse.Merge(m, src)
}
func (m *GetRemoteUserResponse) XXX_Size() int {
	return xxx_messageInfo_GetRemoteUserResponse.Size(m)
}
func (m *GetRemoteUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRemoteUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetRemoteUserResponse proto.InternalMessageInfo

func (m *GetRemoteUserResponse) GetStatus() *v1beta11.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *GetRemoteUserResponse) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *GetRemoteUserResponse) GetRemoteUser() *v1beta13.User {
	if m != nil {
		return m.RemoteUser
	}
	return nil
}

func init() {
	proto.RegisterType((*GenerateInviteTokenRequest)(nil), "cs3.ocm.invite.v1beta1.GenerateInviteTokenRequest")
	proto.RegisterType((*GenerateInviteTokenResponse)(nil), "cs3.ocm.invite.v1beta1.GenerateInviteTokenResponse")
	proto.RegisterType((*ForwardInviteRequest)(nil), "cs3.ocm.invite.v1beta1.ForwardInviteRequest")
	proto.RegisterType((*ForwardInviteResponse)(nil), "cs3.ocm.invite.v1beta1.ForwardInviteResponse")
	proto.RegisterType((*AcceptInviteRequest)(nil), "cs3.ocm.invite.v1beta1.AcceptInviteRequest")
	proto.RegisterType((*AcceptInviteResponse)(nil), "cs3.ocm.invite.v1beta1.AcceptInviteResponse")
	proto.RegisterType((*GetRemoteUserRequest)(nil), "cs3.ocm.invite.v1beta1.GetRemoteUserRequest")
	proto.RegisterType((*GetRemoteUserResponse)(nil), "cs3.ocm.invite.v1beta1.GetRemoteUserResponse")
}

func init() {
	proto.RegisterFile("cs3/ocm/invite/v1beta1/invite_api.proto", fileDescriptor_b0e70953130e3677)
}

var fileDescriptor_b0e70953130e3677 = []byte{
	// 589 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0x5d, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0x53, 0x29, 0x12, 0x9b, 0xb4, 0x0f, 0x6e, 0x1a, 0x8a, 0x01, 0x01, 0x46, 0x2a, 0x45,
	0x94, 0x8d, 0x92, 0x5c, 0x80, 0xa6, 0x52, 0x23, 0x3f, 0x25, 0x72, 0x81, 0x07, 0x14, 0xc9, 0x72,
	0xd7, 0x03, 0x5a, 0x81, 0xbd, 0xee, 0xee, 0x3a, 0x28, 0x20, 0x71, 0x08, 0x8e, 0xc0, 0x23, 0x77,
	0xe0, 0x00, 0xf0, 0xd8, 0x8b, 0x70, 0x05, 0x64, 0xef, 0xda, 0xf9, 0xc1, 0x86, 0xd0, 0x4a, 0x11,
	0x6f, 0xd9, 0x7c, 0xdf, 0x7c, 0x3b, 0xf3, 0xcd, 0xec, 0x18, 0x3d, 0x22, 0xa2, 0xdf, 0x61, 0x24,
	0xec, 0xd0, 0x68, 0x4a, 0x25, 0x74, 0xa6, 0xdd, 0x73, 0x90, 0x7e, 0x57, 0x1f, 0x3d, 0x3f, 0xa6,
	0x38, 0xe6, 0x4c, 0x32, 0xb3, 0x4d, 0x44, 0x1f, 0x33, 0x12, 0x62, 0x85, 0x60, 0x4d, 0xb4, 0x1e,
	0xa7, 0x02, 0x34, 0x80, 0x48, 0x52, 0x39, 0xeb, 0x24, 0x02, 0x78, 0xa1, 0xc1, 0x41, 0xb0, 0x84,
	0x13, 0x10, 0x4a, 0xc2, 0x3a, 0xa8, 0xb8, 0x6b, 0x95, 0x77, 0x98, 0xf3, 0x62, 0xce, 0xa6, 0x34,
	0xf8, 0x83, 0xe2, 0x9d, 0x94, 0xc9, 0x63, 0x52, 0x10, 0x84, 0xf4, 0x65, 0x92, 0xa3, 0x77, 0x53,
	0x54, 0xce, 0x62, 0x10, 0x05, 0x9e, 0x9d, 0x14, 0x6c, 0x8f, 0x90, 0x35, 0x84, 0x08, 0xb8, 0x2f,
	0xc1, 0xc9, 0x12, 0x7a, 0xce, 0xde, 0x42, 0xe4, 0xc2, 0x45, 0x02, 0x42, 0x9a, 0x5d, 0x54, 0x67,
	0xb1, 0x7f, 0x91, 0xc0, 0xbe, 0x71, 0xdf, 0x38, 0x6c, 0xf4, 0x6e, 0xe1, 0xd4, 0x00, 0x15, 0xaf,
	0xd5, 0xf0, 0x28, 0x23, 0xb8, 0x9a, 0x68, 0x7f, 0x37, 0xd0, 0xed, 0x52, 0x45, 0x11, 0xb3, 0x48,
	0x80, 0xd9, 0x41, 0x75, 0x95, 0x9f, 0x96, 0xbc, 0x99, 0x49, 0xf2, 0x98, 0x14, 0x82, 0x67, 0x19,
	0xec, 0x6a, 0xda, 0x42, 0x0e, 0xb5, 0x35, 0x73, 0x30, 0x4f, 0x51, 0x53, 0xb7, 0x4e, 0xa6, 0x77,
	0xef, 0x6f, 0x65, 0x81, 0x0f, 0x71, 0x79, 0xf7, 0xf0, 0x62, 0x9a, 0x0d, 0x3a, 0x3f, 0xd8, 0x3f,
	0x0d, 0xd4, 0x3a, 0x65, 0xfc, 0xbd, 0xcf, 0x03, 0xc5, 0xb9, 0xba, 0x2f, 0xbf, 0xe5, 0x54, 0xbb,
	0x5a, 0x4e, 0xe6, 0x04, 0xb5, 0x19, 0xa7, 0x6f, 0x68, 0xe4, 0x89, 0x99, 0x90, 0x10, 0x7a, 0xf9,
	0x7c, 0xe8, 0x2a, 0x0f, 0x0a, 0xc5, 0x1c, 0x28, 0x34, 0xc7, 0xfa, 0x0f, 0x27, 0x7a, 0xcd, 0xdc,
	0x96, 0x52, 0x39, 0xcb, 0x44, 0x72, 0xc4, 0xfe, 0x88, 0xf6, 0x56, 0x0a, 0xde, 0x5c, 0xdb, 0xec,
	0x4b, 0x03, 0xed, 0x1e, 0x13, 0x02, 0xb1, 0xfc, 0x6f, 0xdc, 0x7e, 0x86, 0x1a, 0x1c, 0x42, 0x26,
	0xc1, 0x4b, 0x1f, 0xb5, 0xb6, 0xf8, 0x5e, 0x26, 0x93, 0x3f, 0x77, 0x9c, 0x22, 0x85, 0xd2, 0x0b,
	0x01, 0xdc, 0x45, 0x2a, 0x26, 0xfd, 0x6d, 0x7f, 0x40, 0xad, 0xe5, 0x9a, 0x36, 0x68, 0xe8, 0x67,
	0x03, 0xb5, 0x86, 0x20, 0xdd, 0x22, 0x9b, 0x6b, 0x38, 0x3a, 0x44, 0x3b, 0x0b, 0x4e, 0x78, 0x34,
	0xd0, 0x69, 0x3c, 0xf8, 0x8b, 0x19, 0x4e, 0xe0, 0x36, 0xe7, 0x76, 0x38, 0x81, 0xfd, 0xcd, 0x40,
	0x7b, 0x2b, 0x49, 0x6d, 0x70, 0x35, 0x5c, 0xbb, 0xa1, 0xbd, 0xcb, 0x2d, 0x74, 0x43, 0xf5, 0xf2,
	0x78, 0xec, 0x98, 0x9f, 0xd0, 0x6e, 0xc9, 0xb6, 0x33, 0x7b, 0x55, 0x93, 0x56, 0xbd, 0x6c, 0xad,
	0xfe, 0x3f, 0xc5, 0x68, 0xcf, 0xde, 0xa1, 0xed, 0xa5, 0x07, 0x6b, 0x1e, 0x55, 0xa9, 0x94, 0x2d,
	0x32, 0xeb, 0xe9, 0x9a, 0x6c, 0x7d, 0x1b, 0x45, 0xcd, 0xc5, 0x61, 0x36, 0x9f, 0x54, 0x85, 0x97,
	0x3c, 0x63, 0xeb, 0x68, 0x3d, 0xf2, 0xbc, 0xb0, 0xa5, 0x29, 0xa9, 0x2e, 0xac, 0x6c, 0xc2, 0xab,
	0x0b, 0x2b, 0x1d, 0xbd, 0x41, 0x82, 0x2c, 0xc2, 0xc2, 0x8a, 0x98, 0xc1, 0x8e, 0xee, 0x77, 0x4c,
	0xc7, 0xe9, 0x47, 0x73, 0x6c, 0xbc, 0xda, 0x56, 0x0c, 0x4d, 0xf8, 0x52, 0xdb, 0x3a, 0x19, 0x39,
	0x5f, 0x6b, 0xed, 0x13, 0xd1, 0xc7, 0x23, 0x12, 0xea, 0x7d, 0x82, 0x5f, 0x76, 0x07, 0x29, 0xfc,
	0x23, 0x03, 0x26, 0x23, 0x12, 0x4e, 0x14, 0x30, 0xd1, 0xc0, 0x79, 0x3d, 0xfb, 0x08, 0xf7, 0x7f,
	0x05, 0x00, 0x00, 0xff, 0xff, 0xd0, 0x5f, 0x6a, 0x47, 0x81, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InviteAPIClient is the client API for InviteAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InviteAPIClient interface {
	// Generates a new token for the user with a validity of 24 hours.
	GenerateInviteToken(ctx context.Context, in *GenerateInviteTokenRequest, opts ...grpc.CallOption) (*GenerateInviteTokenResponse, error)
	// Forwards a received invite to the sync'n'share system provider.
	ForwardInvite(ctx context.Context, in *ForwardInviteRequest, opts ...grpc.CallOption) (*ForwardInviteResponse, error)
	// Completes an invitation acceptance.
	AcceptInvite(ctx context.Context, in *AcceptInviteRequest, opts ...grpc.CallOption) (*AcceptInviteResponse, error)
	// Retrieves details about a remote user who has accepted an invite to share.
	GetRemoteUser(ctx context.Context, in *GetRemoteUserRequest, opts ...grpc.CallOption) (*GetRemoteUserResponse, error)
}

type inviteAPIClient struct {
	cc *grpc.ClientConn
}

func NewInviteAPIClient(cc *grpc.ClientConn) InviteAPIClient {
	return &inviteAPIClient{cc}
}

func (c *inviteAPIClient) GenerateInviteToken(ctx context.Context, in *GenerateInviteTokenRequest, opts ...grpc.CallOption) (*GenerateInviteTokenResponse, error) {
	out := new(GenerateInviteTokenResponse)
	err := c.cc.Invoke(ctx, "/cs3.ocm.invite.v1beta1.InviteAPI/GenerateInviteToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inviteAPIClient) ForwardInvite(ctx context.Context, in *ForwardInviteRequest, opts ...grpc.CallOption) (*ForwardInviteResponse, error) {
	out := new(ForwardInviteResponse)
	err := c.cc.Invoke(ctx, "/cs3.ocm.invite.v1beta1.InviteAPI/ForwardInvite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inviteAPIClient) AcceptInvite(ctx context.Context, in *AcceptInviteRequest, opts ...grpc.CallOption) (*AcceptInviteResponse, error) {
	out := new(AcceptInviteResponse)
	err := c.cc.Invoke(ctx, "/cs3.ocm.invite.v1beta1.InviteAPI/AcceptInvite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inviteAPIClient) GetRemoteUser(ctx context.Context, in *GetRemoteUserRequest, opts ...grpc.CallOption) (*GetRemoteUserResponse, error) {
	out := new(GetRemoteUserResponse)
	err := c.cc.Invoke(ctx, "/cs3.ocm.invite.v1beta1.InviteAPI/GetRemoteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InviteAPIServer is the server API for InviteAPI service.
type InviteAPIServer interface {
	// Generates a new token for the user with a validity of 24 hours.
	GenerateInviteToken(context.Context, *GenerateInviteTokenRequest) (*GenerateInviteTokenResponse, error)
	// Forwards a received invite to the sync'n'share system provider.
	ForwardInvite(context.Context, *ForwardInviteRequest) (*ForwardInviteResponse, error)
	// Completes an invitation acceptance.
	AcceptInvite(context.Context, *AcceptInviteRequest) (*AcceptInviteResponse, error)
	// Retrieves details about a remote user who has accepted an invite to share.
	GetRemoteUser(context.Context, *GetRemoteUserRequest) (*GetRemoteUserResponse, error)
}

// UnimplementedInviteAPIServer can be embedded to have forward compatible implementations.
type UnimplementedInviteAPIServer struct {
}

func (*UnimplementedInviteAPIServer) GenerateInviteToken(ctx context.Context, req *GenerateInviteTokenRequest) (*GenerateInviteTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateInviteToken not implemented")
}
func (*UnimplementedInviteAPIServer) ForwardInvite(ctx context.Context, req *ForwardInviteRequest) (*ForwardInviteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForwardInvite not implemented")
}
func (*UnimplementedInviteAPIServer) AcceptInvite(ctx context.Context, req *AcceptInviteRequest) (*AcceptInviteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptInvite not implemented")
}
func (*UnimplementedInviteAPIServer) GetRemoteUser(ctx context.Context, req *GetRemoteUserRequest) (*GetRemoteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRemoteUser not implemented")
}

func RegisterInviteAPIServer(s *grpc.Server, srv InviteAPIServer) {
	s.RegisterService(&_InviteAPI_serviceDesc, srv)
}

func _InviteAPI_GenerateInviteToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateInviteTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InviteAPIServer).GenerateInviteToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cs3.ocm.invite.v1beta1.InviteAPI/GenerateInviteToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InviteAPIServer).GenerateInviteToken(ctx, req.(*GenerateInviteTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InviteAPI_ForwardInvite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForwardInviteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InviteAPIServer).ForwardInvite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cs3.ocm.invite.v1beta1.InviteAPI/ForwardInvite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InviteAPIServer).ForwardInvite(ctx, req.(*ForwardInviteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InviteAPI_AcceptInvite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptInviteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InviteAPIServer).AcceptInvite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cs3.ocm.invite.v1beta1.InviteAPI/AcceptInvite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InviteAPIServer).AcceptInvite(ctx, req.(*AcceptInviteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InviteAPI_GetRemoteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRemoteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InviteAPIServer).GetRemoteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cs3.ocm.invite.v1beta1.InviteAPI/GetRemoteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InviteAPIServer).GetRemoteUser(ctx, req.(*GetRemoteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _InviteAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cs3.ocm.invite.v1beta1.InviteAPI",
	HandlerType: (*InviteAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateInviteToken",
			Handler:    _InviteAPI_GenerateInviteToken_Handler,
		},
		{
			MethodName: "ForwardInvite",
			Handler:    _InviteAPI_ForwardInvite_Handler,
		},
		{
			MethodName: "AcceptInvite",
			Handler:    _InviteAPI_AcceptInvite_Handler,
		},
		{
			MethodName: "GetRemoteUser",
			Handler:    _InviteAPI_GetRemoteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cs3/ocm/invite/v1beta1/invite_api.proto",
}
