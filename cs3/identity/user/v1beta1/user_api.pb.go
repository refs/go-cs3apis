// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs3/identity/user/v1beta1/user_api.proto

package userv1beta1

import (
	context "context"
	fmt "fmt"
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

type GetUserRequest struct {
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The id of the user.
	UserId               *UserId  `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb4e4164505d2418, []int{0}
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

func (m *GetUserRequest) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *GetUserRequest) GetUserId() *UserId {
	if m != nil {
		return m.UserId
	}
	return nil
}

type GetUserResponse struct {
	// REQUIRED.
	// The response status.
	Status *v1beta11.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The user information.
	User                 *User    `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserResponse) Reset()         { *m = GetUserResponse{} }
func (m *GetUserResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserResponse) ProtoMessage()    {}
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb4e4164505d2418, []int{1}
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

func (m *GetUserResponse) GetStatus() *v1beta11.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *GetUserResponse) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *GetUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type GetUserByClaimRequest struct {
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The claim on the basis of which users will be filtered.
	Claim string `protobuf:"bytes,2,opt,name=claim,proto3" json:"claim,omitempty"`
	// REQUIRED.
	// The value of the claim to find the specific user.
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserByClaimRequest) Reset()         { *m = GetUserByClaimRequest{} }
func (m *GetUserByClaimRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserByClaimRequest) ProtoMessage()    {}
func (*GetUserByClaimRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb4e4164505d2418, []int{2}
}

func (m *GetUserByClaimRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserByClaimRequest.Unmarshal(m, b)
}
func (m *GetUserByClaimRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserByClaimRequest.Marshal(b, m, deterministic)
}
func (m *GetUserByClaimRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserByClaimRequest.Merge(m, src)
}
func (m *GetUserByClaimRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserByClaimRequest.Size(m)
}
func (m *GetUserByClaimRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserByClaimRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserByClaimRequest proto.InternalMessageInfo

func (m *GetUserByClaimRequest) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *GetUserByClaimRequest) GetClaim() string {
	if m != nil {
		return m.Claim
	}
	return ""
}

func (m *GetUserByClaimRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type GetUserByClaimResponse struct {
	// REQUIRED.
	// The response status.
	Status *v1beta11.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The user information.
	User                 *User    `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserByClaimResponse) Reset()         { *m = GetUserByClaimResponse{} }
func (m *GetUserByClaimResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserByClaimResponse) ProtoMessage()    {}
func (*GetUserByClaimResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb4e4164505d2418, []int{3}
}

func (m *GetUserByClaimResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserByClaimResponse.Unmarshal(m, b)
}
func (m *GetUserByClaimResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserByClaimResponse.Marshal(b, m, deterministic)
}
func (m *GetUserByClaimResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserByClaimResponse.Merge(m, src)
}
func (m *GetUserByClaimResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserByClaimResponse.Size(m)
}
func (m *GetUserByClaimResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserByClaimResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserByClaimResponse proto.InternalMessageInfo

func (m *GetUserByClaimResponse) GetStatus() *v1beta11.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *GetUserByClaimResponse) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *GetUserByClaimResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type GetUserGroupsRequest struct {
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The id of the user.
	UserId               *UserId  `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserGroupsRequest) Reset()         { *m = GetUserGroupsRequest{} }
func (m *GetUserGroupsRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserGroupsRequest) ProtoMessage()    {}
func (*GetUserGroupsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb4e4164505d2418, []int{4}
}

func (m *GetUserGroupsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserGroupsRequest.Unmarshal(m, b)
}
func (m *GetUserGroupsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserGroupsRequest.Marshal(b, m, deterministic)
}
func (m *GetUserGroupsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserGroupsRequest.Merge(m, src)
}
func (m *GetUserGroupsRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserGroupsRequest.Size(m)
}
func (m *GetUserGroupsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserGroupsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserGroupsRequest proto.InternalMessageInfo

func (m *GetUserGroupsRequest) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *GetUserGroupsRequest) GetUserId() *UserId {
	if m != nil {
		return m.UserId
	}
	return nil
}

type GetUserGroupsResponse struct {
	// REQUIRED.
	// The response status.
	Status *v1beta11.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The groups for the user.
	Groups               []string `protobuf:"bytes,3,rep,name=groups,proto3" json:"groups,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserGroupsResponse) Reset()         { *m = GetUserGroupsResponse{} }
func (m *GetUserGroupsResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserGroupsResponse) ProtoMessage()    {}
func (*GetUserGroupsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb4e4164505d2418, []int{5}
}

func (m *GetUserGroupsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserGroupsResponse.Unmarshal(m, b)
}
func (m *GetUserGroupsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserGroupsResponse.Marshal(b, m, deterministic)
}
func (m *GetUserGroupsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserGroupsResponse.Merge(m, src)
}
func (m *GetUserGroupsResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserGroupsResponse.Size(m)
}
func (m *GetUserGroupsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserGroupsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserGroupsResponse proto.InternalMessageInfo

func (m *GetUserGroupsResponse) GetStatus() *v1beta11.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *GetUserGroupsResponse) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *GetUserGroupsResponse) GetGroups() []string {
	if m != nil {
		return m.Groups
	}
	return nil
}

type IsInGroupRequest struct {
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The id of the user.
	UserId *UserId `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// REQUIRED.
	// The group to check.
	Group                string   `protobuf:"bytes,3,opt,name=group,proto3" json:"group,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsInGroupRequest) Reset()         { *m = IsInGroupRequest{} }
func (m *IsInGroupRequest) String() string { return proto.CompactTextString(m) }
func (*IsInGroupRequest) ProtoMessage()    {}
func (*IsInGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb4e4164505d2418, []int{6}
}

func (m *IsInGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsInGroupRequest.Unmarshal(m, b)
}
func (m *IsInGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsInGroupRequest.Marshal(b, m, deterministic)
}
func (m *IsInGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsInGroupRequest.Merge(m, src)
}
func (m *IsInGroupRequest) XXX_Size() int {
	return xxx_messageInfo_IsInGroupRequest.Size(m)
}
func (m *IsInGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IsInGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IsInGroupRequest proto.InternalMessageInfo

func (m *IsInGroupRequest) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *IsInGroupRequest) GetUserId() *UserId {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *IsInGroupRequest) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

type IsInGroupResponse struct {
	// REQUIRED.
	// The response status.
	Status *v1beta11.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// Tells if the user belongs to the group.
	Ok                   bool     `protobuf:"varint,3,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsInGroupResponse) Reset()         { *m = IsInGroupResponse{} }
func (m *IsInGroupResponse) String() string { return proto.CompactTextString(m) }
func (*IsInGroupResponse) ProtoMessage()    {}
func (*IsInGroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb4e4164505d2418, []int{7}
}

func (m *IsInGroupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsInGroupResponse.Unmarshal(m, b)
}
func (m *IsInGroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsInGroupResponse.Marshal(b, m, deterministic)
}
func (m *IsInGroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsInGroupResponse.Merge(m, src)
}
func (m *IsInGroupResponse) XXX_Size() int {
	return xxx_messageInfo_IsInGroupResponse.Size(m)
}
func (m *IsInGroupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IsInGroupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IsInGroupResponse proto.InternalMessageInfo

func (m *IsInGroupResponse) GetStatus() *v1beta11.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *IsInGroupResponse) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *IsInGroupResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

type FindUsersRequest struct {
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED. TODO(labkode): create proper filters for most common searches.
	// The filter to apply.
	Filter               string   `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindUsersRequest) Reset()         { *m = FindUsersRequest{} }
func (m *FindUsersRequest) String() string { return proto.CompactTextString(m) }
func (*FindUsersRequest) ProtoMessage()    {}
func (*FindUsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb4e4164505d2418, []int{8}
}

func (m *FindUsersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindUsersRequest.Unmarshal(m, b)
}
func (m *FindUsersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindUsersRequest.Marshal(b, m, deterministic)
}
func (m *FindUsersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindUsersRequest.Merge(m, src)
}
func (m *FindUsersRequest) XXX_Size() int {
	return xxx_messageInfo_FindUsersRequest.Size(m)
}
func (m *FindUsersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindUsersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindUsersRequest proto.InternalMessageInfo

func (m *FindUsersRequest) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *FindUsersRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

type FindUsersResponse struct {
	// REQUIRED.
	// The response status.
	Status *v1beta11.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The users matching the specified filter.
	Users                []*User  `protobuf:"bytes,3,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindUsersResponse) Reset()         { *m = FindUsersResponse{} }
func (m *FindUsersResponse) String() string { return proto.CompactTextString(m) }
func (*FindUsersResponse) ProtoMessage()    {}
func (*FindUsersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb4e4164505d2418, []int{9}
}

func (m *FindUsersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindUsersResponse.Unmarshal(m, b)
}
func (m *FindUsersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindUsersResponse.Marshal(b, m, deterministic)
}
func (m *FindUsersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindUsersResponse.Merge(m, src)
}
func (m *FindUsersResponse) XXX_Size() int {
	return xxx_messageInfo_FindUsersResponse.Size(m)
}
func (m *FindUsersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindUsersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindUsersResponse proto.InternalMessageInfo

func (m *FindUsersResponse) GetStatus() *v1beta11.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *FindUsersResponse) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *FindUsersResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type FindGroupsRequest struct {
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The filter to apply.
	Filter               string   `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindGroupsRequest) Reset()         { *m = FindGroupsRequest{} }
func (m *FindGroupsRequest) String() string { return proto.CompactTextString(m) }
func (*FindGroupsRequest) ProtoMessage()    {}
func (*FindGroupsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb4e4164505d2418, []int{10}
}

func (m *FindGroupsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindGroupsRequest.Unmarshal(m, b)
}
func (m *FindGroupsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindGroupsRequest.Marshal(b, m, deterministic)
}
func (m *FindGroupsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindGroupsRequest.Merge(m, src)
}
func (m *FindGroupsRequest) XXX_Size() int {
	return xxx_messageInfo_FindGroupsRequest.Size(m)
}
func (m *FindGroupsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindGroupsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindGroupsRequest proto.InternalMessageInfo

func (m *FindGroupsRequest) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *FindGroupsRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

type FindGroupsResponse struct {
	// REQUIRED.
	// The response status.
	Status *v1beta11.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The groups matching the specified filter.
	Groups               []string `protobuf:"bytes,3,rep,name=groups,proto3" json:"groups,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindGroupsResponse) Reset()         { *m = FindGroupsResponse{} }
func (m *FindGroupsResponse) String() string { return proto.CompactTextString(m) }
func (*FindGroupsResponse) ProtoMessage()    {}
func (*FindGroupsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb4e4164505d2418, []int{11}
}

func (m *FindGroupsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindGroupsResponse.Unmarshal(m, b)
}
func (m *FindGroupsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindGroupsResponse.Marshal(b, m, deterministic)
}
func (m *FindGroupsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindGroupsResponse.Merge(m, src)
}
func (m *FindGroupsResponse) XXX_Size() int {
	return xxx_messageInfo_FindGroupsResponse.Size(m)
}
func (m *FindGroupsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindGroupsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindGroupsResponse proto.InternalMessageInfo

func (m *FindGroupsResponse) GetStatus() *v1beta11.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *FindGroupsResponse) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *FindGroupsResponse) GetGroups() []string {
	if m != nil {
		return m.Groups
	}
	return nil
}

func init() {
	proto.RegisterType((*GetUserRequest)(nil), "cs3.identity.user.v1beta1.GetUserRequest")
	proto.RegisterType((*GetUserResponse)(nil), "cs3.identity.user.v1beta1.GetUserResponse")
	proto.RegisterType((*GetUserByClaimRequest)(nil), "cs3.identity.user.v1beta1.GetUserByClaimRequest")
	proto.RegisterType((*GetUserByClaimResponse)(nil), "cs3.identity.user.v1beta1.GetUserByClaimResponse")
	proto.RegisterType((*GetUserGroupsRequest)(nil), "cs3.identity.user.v1beta1.GetUserGroupsRequest")
	proto.RegisterType((*GetUserGroupsResponse)(nil), "cs3.identity.user.v1beta1.GetUserGroupsResponse")
	proto.RegisterType((*IsInGroupRequest)(nil), "cs3.identity.user.v1beta1.IsInGroupRequest")
	proto.RegisterType((*IsInGroupResponse)(nil), "cs3.identity.user.v1beta1.IsInGroupResponse")
	proto.RegisterType((*FindUsersRequest)(nil), "cs3.identity.user.v1beta1.FindUsersRequest")
	proto.RegisterType((*FindUsersResponse)(nil), "cs3.identity.user.v1beta1.FindUsersResponse")
	proto.RegisterType((*FindGroupsRequest)(nil), "cs3.identity.user.v1beta1.FindGroupsRequest")
	proto.RegisterType((*FindGroupsResponse)(nil), "cs3.identity.user.v1beta1.FindGroupsResponse")
}

func init() {
	proto.RegisterFile("cs3/identity/user/v1beta1/user_api.proto", fileDescriptor_fb4e4164505d2418)
}

var fileDescriptor_fb4e4164505d2418 = []byte{
	// 624 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0xd5, 0x3a, 0xbf, 0x38, 0xbf, 0x4c, 0xa1, 0xb4, 0xab, 0x12, 0x12, 0x8b, 0x8a, 0xe2, 0x53,
	0xca, 0x1f, 0xa7, 0x69, 0xc4, 0x85, 0x1b, 0x89, 0x44, 0xe5, 0x13, 0x91, 0x51, 0x39, 0xa0, 0xf0,
	0xc7, 0x75, 0xb6, 0xc8, 0x6a, 0x1a, 0xbb, 0x5e, 0x3b, 0x52, 0xb8, 0x70, 0x42, 0x5c, 0x41, 0x1c,
	0xb8, 0x73, 0xec, 0x01, 0x89, 0xaf, 0xc1, 0xa7, 0x42, 0x3b, 0x5e, 0xbb, 0x71, 0x50, 0x12, 0xa3,
	0x48, 0x09, 0xa7, 0x68, 0xf2, 0xde, 0x9b, 0x79, 0x3b, 0x9e, 0xf1, 0x1a, 0xea, 0x0e, 0x6f, 0x35,
	0xdc, 0x3e, 0x1b, 0x86, 0x6e, 0x38, 0x6e, 0x44, 0x9c, 0x05, 0x8d, 0x51, 0xf3, 0x84, 0x85, 0x76,
	0x13, 0x83, 0x37, 0xb6, 0xef, 0x1a, 0x7e, 0xe0, 0x85, 0x1e, 0xad, 0x39, 0xbc, 0x65, 0x24, 0x4c,
	0x43, 0x80, 0x86, 0x64, 0x6a, 0xfb, 0xb3, 0x93, 0x04, 0x8c, 0x7b, 0x51, 0xe0, 0x30, 0x1e, 0x67,
	0xd1, 0x6e, 0x0b, 0x6a, 0xe0, 0x3b, 0x29, 0x81, 0x87, 0x76, 0x18, 0x25, 0xe8, 0xae, 0x40, 0xc3,
	0xb1, 0xcf, 0x78, 0x8a, 0x63, 0x14, 0xc3, 0xfa, 0x07, 0xd8, 0x3c, 0x62, 0xe1, 0x31, 0x67, 0x81,
	0xc5, 0x2e, 0x22, 0xc6, 0x43, 0xda, 0x04, 0xd5, 0xf3, 0xed, 0x8b, 0x88, 0x55, 0xc9, 0x1e, 0xa9,
	0x6f, 0x1c, 0xd6, 0x0c, 0xe1, 0x32, 0xd6, 0xc8, 0x0c, 0xc6, 0x33, 0x24, 0x58, 0x92, 0x48, 0x1f,
	0x43, 0x09, 0x4f, 0xe6, 0xf6, 0xab, 0x0a, 0x6a, 0xee, 0x1a, 0x33, 0x4f, 0x66, 0x88, 0x5a, 0x66,
	0xdf, 0x52, 0x23, 0xfc, 0xd5, 0x2f, 0x09, 0xdc, 0x48, 0x1d, 0x70, 0xdf, 0x1b, 0x72, 0x46, 0x1b,
	0xa0, 0xc6, 0x67, 0x90, 0x16, 0x6e, 0x61, 0xba, 0xc0, 0x77, 0xd2, 0x24, 0xcf, 0x11, 0xb6, 0x24,
	0x6d, 0xc2, 0xb3, 0x92, 0xd7, 0x73, 0x0b, 0xfe, 0x13, 0x0e, 0xaa, 0x05, 0x14, 0xdc, 0x59, 0x60,
	0xd8, 0x42, 0xb2, 0x3e, 0x82, 0x9b, 0xd2, 0x6b, 0x7b, 0xdc, 0x19, 0xd8, 0xee, 0xf9, 0x12, 0x4d,
	0xdb, 0x81, 0xa2, 0x23, 0x52, 0xa0, 0xe5, 0xb2, 0x15, 0x07, 0xe2, 0xdf, 0x91, 0x3d, 0x88, 0x18,
	0xfa, 0x2a, 0x5b, 0x71, 0xa0, 0xff, 0x24, 0x50, 0x99, 0x2e, 0xfc, 0xaf, 0xf7, 0xea, 0x23, 0x81,
	0x1d, 0xe9, 0xf9, 0x28, 0xf0, 0x22, 0x9f, 0xaf, 0x69, 0xc0, 0xbe, 0x92, 0xf4, 0xa1, 0x25, 0x3e,
	0x56, 0xd8, 0xba, 0x0a, 0xa8, 0xef, 0xb0, 0x6a, 0xb5, 0xb0, 0x57, 0xa8, 0x97, 0x2d, 0x19, 0xe9,
	0xdf, 0x08, 0x6c, 0x99, 0xdc, 0x1c, 0xa2, 0xa5, 0xf5, 0x74, 0x46, 0xcc, 0x1a, 0xba, 0x49, 0x66,
	0x0d, 0x03, 0xfd, 0x13, 0x81, 0xed, 0x09, 0x67, 0x2b, 0xec, 0xd5, 0x26, 0x28, 0xde, 0x19, 0x9a,
	0xf9, 0xdf, 0x52, 0xbc, 0x33, 0xfd, 0x15, 0x6c, 0x3d, 0x75, 0x87, 0x7d, 0xe1, 0x7a, 0x99, 0xe1,
	0xa9, 0x80, 0x7a, 0xea, 0x0e, 0x42, 0x16, 0xc8, 0x4d, 0x93, 0x91, 0xfe, 0x83, 0xc0, 0xf6, 0x44,
	0xfe, 0x15, 0x1e, 0xf4, 0x11, 0x14, 0xc5, 0x23, 0x88, 0x67, 0x22, 0xc7, 0x42, 0xc5, 0x6c, 0xfd,
	0x75, 0xec, 0x77, 0xe9, 0x6d, 0x9a, 0xd5, 0x90, 0xcf, 0x04, 0xe8, 0x64, 0x81, 0xf5, 0xaf, 0xc9,
	0xe1, 0x97, 0x22, 0x94, 0x44, 0x0b, 0x9e, 0x74, 0x4d, 0xfa, 0x16, 0x4a, 0x72, 0x8f, 0xe9, 0xfe,
	0x9c, 0x8e, 0x65, 0xaf, 0x33, 0xed, 0x5e, 0x1e, 0xaa, 0x3c, 0x69, 0x94, 0x5e, 0x86, 0xf2, 0x2d,
	0x4b, 0x0f, 0x16, 0xab, 0xb3, 0x37, 0x81, 0xd6, 0xfc, 0x0b, 0x85, 0x2c, 0x1b, 0xc0, 0xf5, 0xcc,
	0x0b, 0x8a, 0x36, 0x16, 0xe7, 0xc8, 0x0c, 0x81, 0x76, 0x90, 0x5f, 0x20, 0x6b, 0x9e, 0x42, 0x39,
	0x5d, 0x72, 0x7a, 0x7f, 0x8e, 0x7c, 0xfa, 0x25, 0xa5, 0x3d, 0xc8, 0x47, 0xbe, 0xaa, 0x93, 0xee,
	0xd8, 0xdc, 0x3a, 0xd3, 0x9b, 0x3e, 0xb7, 0xce, 0x9f, 0x6b, 0xeb, 0x02, 0x5c, 0x8d, 0x2e, 0x5d,
	0xa4, 0xcd, 0x76, 0xef, 0x61, 0x4e, 0x76, 0x5c, 0xaa, 0xfd, 0x1e, 0x76, 0x1d, 0xef, 0x7c, 0xb6,
	0xa6, 0x7d, 0x0d, 0x27, 0xd6, 0x77, 0xbb, 0xe2, 0x0b, 0xab, 0x4b, 0x5e, 0x6e, 0x08, 0x54, 0x82,
	0xdf, 0x95, 0x42, 0xc7, 0x3c, 0xbe, 0x54, 0x6a, 0x1d, 0xde, 0x32, 0xcc, 0x44, 0x2e, 0x04, 0xc6,
	0x8b, 0x66, 0x5b, 0x30, 0x7e, 0x21, 0xd6, 0x4b, 0xb0, 0x9e, 0xc0, 0x7a, 0x12, 0x3b, 0x51, 0xf1,
	0xab, 0xad, 0xf5, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x8a, 0xf6, 0x86, 0xc9, 0x64, 0x0a, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserAPIClient is the client API for UserAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserAPIClient interface {
	// Gets the information about a user by the user id.
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	// Gets the information about a user based on a specified claim.
	GetUserByClaim(ctx context.Context, in *GetUserByClaimRequest, opts ...grpc.CallOption) (*GetUserByClaimResponse, error)
	// Gets the groups of a user.
	GetUserGroups(ctx context.Context, in *GetUserGroupsRequest, opts ...grpc.CallOption) (*GetUserGroupsResponse, error)
	// Tells if the user is in a certain group.
	IsInGroup(ctx context.Context, in *IsInGroupRequest, opts ...grpc.CallOption) (*IsInGroupResponse, error)
	// Finds users by any attribute of the user.
	// TODO(labkode): to define the filters that make more sense.
	FindUsers(ctx context.Context, in *FindUsersRequest, opts ...grpc.CallOption) (*FindUsersResponse, error)
	// Finds groups whose names match the specified filter.
	FindGroups(ctx context.Context, in *FindGroupsRequest, opts ...grpc.CallOption) (*FindGroupsResponse, error)
}

type userAPIClient struct {
	cc *grpc.ClientConn
}

func NewUserAPIClient(cc *grpc.ClientConn) UserAPIClient {
	return &userAPIClient{cc}
}

func (c *userAPIClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/cs3.identity.user.v1beta1.UserAPI/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAPIClient) GetUserByClaim(ctx context.Context, in *GetUserByClaimRequest, opts ...grpc.CallOption) (*GetUserByClaimResponse, error) {
	out := new(GetUserByClaimResponse)
	err := c.cc.Invoke(ctx, "/cs3.identity.user.v1beta1.UserAPI/GetUserByClaim", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAPIClient) GetUserGroups(ctx context.Context, in *GetUserGroupsRequest, opts ...grpc.CallOption) (*GetUserGroupsResponse, error) {
	out := new(GetUserGroupsResponse)
	err := c.cc.Invoke(ctx, "/cs3.identity.user.v1beta1.UserAPI/GetUserGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAPIClient) IsInGroup(ctx context.Context, in *IsInGroupRequest, opts ...grpc.CallOption) (*IsInGroupResponse, error) {
	out := new(IsInGroupResponse)
	err := c.cc.Invoke(ctx, "/cs3.identity.user.v1beta1.UserAPI/IsInGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAPIClient) FindUsers(ctx context.Context, in *FindUsersRequest, opts ...grpc.CallOption) (*FindUsersResponse, error) {
	out := new(FindUsersResponse)
	err := c.cc.Invoke(ctx, "/cs3.identity.user.v1beta1.UserAPI/FindUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAPIClient) FindGroups(ctx context.Context, in *FindGroupsRequest, opts ...grpc.CallOption) (*FindGroupsResponse, error) {
	out := new(FindGroupsResponse)
	err := c.cc.Invoke(ctx, "/cs3.identity.user.v1beta1.UserAPI/FindGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserAPIServer is the server API for UserAPI service.
type UserAPIServer interface {
	// Gets the information about a user by the user id.
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	// Gets the information about a user based on a specified claim.
	GetUserByClaim(context.Context, *GetUserByClaimRequest) (*GetUserByClaimResponse, error)
	// Gets the groups of a user.
	GetUserGroups(context.Context, *GetUserGroupsRequest) (*GetUserGroupsResponse, error)
	// Tells if the user is in a certain group.
	IsInGroup(context.Context, *IsInGroupRequest) (*IsInGroupResponse, error)
	// Finds users by any attribute of the user.
	// TODO(labkode): to define the filters that make more sense.
	FindUsers(context.Context, *FindUsersRequest) (*FindUsersResponse, error)
	// Finds groups whose names match the specified filter.
	FindGroups(context.Context, *FindGroupsRequest) (*FindGroupsResponse, error)
}

// UnimplementedUserAPIServer can be embedded to have forward compatible implementations.
type UnimplementedUserAPIServer struct {
}

func (*UnimplementedUserAPIServer) GetUser(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (*UnimplementedUserAPIServer) GetUserByClaim(ctx context.Context, req *GetUserByClaimRequest) (*GetUserByClaimResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByClaim not implemented")
}
func (*UnimplementedUserAPIServer) GetUserGroups(ctx context.Context, req *GetUserGroupsRequest) (*GetUserGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserGroups not implemented")
}
func (*UnimplementedUserAPIServer) IsInGroup(ctx context.Context, req *IsInGroupRequest) (*IsInGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsInGroup not implemented")
}
func (*UnimplementedUserAPIServer) FindUsers(ctx context.Context, req *FindUsersRequest) (*FindUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUsers not implemented")
}
func (*UnimplementedUserAPIServer) FindGroups(ctx context.Context, req *FindGroupsRequest) (*FindGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindGroups not implemented")
}

func RegisterUserAPIServer(s *grpc.Server, srv UserAPIServer) {
	s.RegisterService(&_UserAPI_serviceDesc, srv)
}

func _UserAPI_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAPIServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cs3.identity.user.v1beta1.UserAPI/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAPIServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAPI_GetUserByClaim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByClaimRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAPIServer).GetUserByClaim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cs3.identity.user.v1beta1.UserAPI/GetUserByClaim",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAPIServer).GetUserByClaim(ctx, req.(*GetUserByClaimRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAPI_GetUserGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAPIServer).GetUserGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cs3.identity.user.v1beta1.UserAPI/GetUserGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAPIServer).GetUserGroups(ctx, req.(*GetUserGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAPI_IsInGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsInGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAPIServer).IsInGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cs3.identity.user.v1beta1.UserAPI/IsInGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAPIServer).IsInGroup(ctx, req.(*IsInGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAPI_FindUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAPIServer).FindUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cs3.identity.user.v1beta1.UserAPI/FindUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAPIServer).FindUsers(ctx, req.(*FindUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAPI_FindGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAPIServer).FindGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cs3.identity.user.v1beta1.UserAPI/FindGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAPIServer).FindGroups(ctx, req.(*FindGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cs3.identity.user.v1beta1.UserAPI",
	HandlerType: (*UserAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _UserAPI_GetUser_Handler,
		},
		{
			MethodName: "GetUserByClaim",
			Handler:    _UserAPI_GetUserByClaim_Handler,
		},
		{
			MethodName: "GetUserGroups",
			Handler:    _UserAPI_GetUserGroups_Handler,
		},
		{
			MethodName: "IsInGroup",
			Handler:    _UserAPI_IsInGroup_Handler,
		},
		{
			MethodName: "FindUsers",
			Handler:    _UserAPI_FindUsers_Handler,
		},
		{
			MethodName: "FindGroups",
			Handler:    _UserAPI_FindGroups_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cs3/identity/user/v1beta1/user_api.proto",
}
