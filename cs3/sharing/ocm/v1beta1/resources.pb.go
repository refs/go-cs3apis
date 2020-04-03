// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs3/sharing/ocm/v1beta1/resources.proto

package ocmv1beta1

import (
	fmt "fmt"
	v1beta11 "github.com/cs3org/go-cs3apis/cs3/identity/user/v1beta1"
	v1beta1 "github.com/cs3org/go-cs3apis/cs3/storage/provider/v1beta1"
	v1beta12 "github.com/cs3org/go-cs3apis/cs3/types/v1beta1"
	proto "github.com/golang/protobuf/proto"
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

// The state of the share.
type ShareState int32

const (
	// The share is no longer valid, for example, the share expired.
	ShareState_SHARE_STATE_INVALID ShareState = 0
	// New shares MUST be created in the SHARE_STATE_PENDING state.
	// This state means the share is pending to be accepted or rejected
	// by the recipient of the share.
	ShareState_SHARE_STATE_PENDING ShareState = 1
	// The recipient of the share has accepted the share.
	ShareState_SHARE_STATE_ACCEPTED ShareState = 2
	// The recipient of the share has rejected the share.
	// Do not means the share is removed, the recipient MAY
	// change the state to accepted or pending.
	ShareState_SHARE_STATE_REJECTED ShareState = 3
)

var ShareState_name = map[int32]string{
	0: "SHARE_STATE_INVALID",
	1: "SHARE_STATE_PENDING",
	2: "SHARE_STATE_ACCEPTED",
	3: "SHARE_STATE_REJECTED",
}

var ShareState_value = map[string]int32{
	"SHARE_STATE_INVALID":  0,
	"SHARE_STATE_PENDING":  1,
	"SHARE_STATE_ACCEPTED": 2,
	"SHARE_STATE_REJECTED": 3,
}

func (x ShareState) String() string {
	return proto.EnumName(ShareState_name, int32(x))
}

func (ShareState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b8b387acee7aaf43, []int{0}
}

// Shares are relationships between a resource owner
// (usually the authenticated user) who grants permissions to a recipient (grantee)
// on a specified resource (resource_id). UserShares represents both user and groups.
type Share struct {
	// REQUIRED.
	// Opaque unique identifier of the share.
	Id *ShareId `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// REQUIRED.
	// Unique identifier of the shared resource.
	ResourceId *v1beta1.ResourceId `protobuf:"bytes,2,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// REQUIRED.
	// Permissions for the grantee to use
	// the resource.
	Permissions *SharePermissions `protobuf:"bytes,3,opt,name=permissions,proto3" json:"permissions,omitempty"`
	// REQUIRED.
	// The receiver of the share, like a user, group ...
	Grantee *v1beta1.Grantee `protobuf:"bytes,4,opt,name=grantee,proto3" json:"grantee,omitempty"`
	// REQUIRED.
	// Uniquely identifies the owner of the share
	// (the resource owner at the time of creating the share).
	// In case the ownership of the underlying resource changes
	// the share owner field MAY change to reflect the change of ownsership.
	Owner *v1beta11.UserId `protobuf:"bytes,5,opt,name=owner,proto3" json:"owner,omitempty"`
	// REQUIRED.
	// Uniquely identifies a principal who initiates the share creation.
	// A creator can create shares on behalf of the owner (because of re-sharing,
	// because belonging to special groups, ...).
	// Creator and owner often result in being the same principal.
	Creator *v1beta11.UserId `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
	// REQUIRED.
	// Creation time of the share.
	Ctime *v1beta12.Timestamp `protobuf:"bytes,7,opt,name=ctime,proto3" json:"ctime,omitempty"`
	// REQUIRED.
	// Last modification time of the share.
	Mtime                *v1beta12.Timestamp `protobuf:"bytes,8,opt,name=mtime,proto3" json:"mtime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Share) Reset()         { *m = Share{} }
func (m *Share) String() string { return proto.CompactTextString(m) }
func (*Share) ProtoMessage()    {}
func (*Share) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b387acee7aaf43, []int{0}
}

func (m *Share) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Share.Unmarshal(m, b)
}
func (m *Share) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Share.Marshal(b, m, deterministic)
}
func (m *Share) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Share.Merge(m, src)
}
func (m *Share) XXX_Size() int {
	return xxx_messageInfo_Share.Size(m)
}
func (m *Share) XXX_DiscardUnknown() {
	xxx_messageInfo_Share.DiscardUnknown(m)
}

var xxx_messageInfo_Share proto.InternalMessageInfo

func (m *Share) GetId() *ShareId {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Share) GetResourceId() *v1beta1.ResourceId {
	if m != nil {
		return m.ResourceId
	}
	return nil
}

func (m *Share) GetPermissions() *SharePermissions {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *Share) GetGrantee() *v1beta1.Grantee {
	if m != nil {
		return m.Grantee
	}
	return nil
}

func (m *Share) GetOwner() *v1beta11.UserId {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *Share) GetCreator() *v1beta11.UserId {
	if m != nil {
		return m.Creator
	}
	return nil
}

func (m *Share) GetCtime() *v1beta12.Timestamp {
	if m != nil {
		return m.Ctime
	}
	return nil
}

func (m *Share) GetMtime() *v1beta12.Timestamp {
	if m != nil {
		return m.Mtime
	}
	return nil
}

// The permissions for a share.
type SharePermissions struct {
	Permissions          *v1beta1.ResourcePermissions `protobuf:"bytes,1,opt,name=permissions,proto3" json:"permissions,omitempty"`
	Reshare              bool                         `protobuf:"varint,2,opt,name=reshare,proto3" json:"reshare,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *SharePermissions) Reset()         { *m = SharePermissions{} }
func (m *SharePermissions) String() string { return proto.CompactTextString(m) }
func (*SharePermissions) ProtoMessage()    {}
func (*SharePermissions) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b387acee7aaf43, []int{1}
}

func (m *SharePermissions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SharePermissions.Unmarshal(m, b)
}
func (m *SharePermissions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SharePermissions.Marshal(b, m, deterministic)
}
func (m *SharePermissions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SharePermissions.Merge(m, src)
}
func (m *SharePermissions) XXX_Size() int {
	return xxx_messageInfo_SharePermissions.Size(m)
}
func (m *SharePermissions) XXX_DiscardUnknown() {
	xxx_messageInfo_SharePermissions.DiscardUnknown(m)
}

var xxx_messageInfo_SharePermissions proto.InternalMessageInfo

func (m *SharePermissions) GetPermissions() *v1beta1.ResourcePermissions {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *SharePermissions) GetReshare() bool {
	if m != nil {
		return m.Reshare
	}
	return false
}

// A received share is the share that a grantee will receive.
// It expands the original share by adding state to the share,
// a display name from the perspective of the grantee and a
// resource mount path in case the share will be mounted
// in a path in a storage provider.
type ReceivedShare struct {
	// REQUIRED.
	Share *Share `protobuf:"bytes,1,opt,name=share,proto3" json:"share,omitempty"`
	// REQUIRED.
	// The state of the share.
	State                ShareState `protobuf:"varint,2,opt,name=state,proto3,enum=cs3.sharing.ocm.v1beta1.ShareState" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ReceivedShare) Reset()         { *m = ReceivedShare{} }
func (m *ReceivedShare) String() string { return proto.CompactTextString(m) }
func (*ReceivedShare) ProtoMessage()    {}
func (*ReceivedShare) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b387acee7aaf43, []int{2}
}

func (m *ReceivedShare) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceivedShare.Unmarshal(m, b)
}
func (m *ReceivedShare) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceivedShare.Marshal(b, m, deterministic)
}
func (m *ReceivedShare) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceivedShare.Merge(m, src)
}
func (m *ReceivedShare) XXX_Size() int {
	return xxx_messageInfo_ReceivedShare.Size(m)
}
func (m *ReceivedShare) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceivedShare.DiscardUnknown(m)
}

var xxx_messageInfo_ReceivedShare proto.InternalMessageInfo

func (m *ReceivedShare) GetShare() *Share {
	if m != nil {
		return m.Share
	}
	return nil
}

func (m *ReceivedShare) GetState() ShareState {
	if m != nil {
		return m.State
	}
	return ShareState_SHARE_STATE_INVALID
}

// Uniquely identifies a share in the share provider.
// A share MUST be uniquely identify by four (4) elements:
// 1) The share provider id
// 2) The owner of the share
// 3) The resource id
// 4) The grantee for the share
// This 4-tuple MUST be unique.
// For example, owner Alice shares the resource /home/docs with id
// home:1234 to an user named Bob. The 4-tuple will consist of
// 1) The share provider id = "user"
// 2) The owner of the share = "Alice"
// 3) The resource id = "home:1234"
// 4) The grantee for the share = Grantee("type" = "user", "" => "Bob")
type ShareKey struct {
	// REQUIRED.
	Owner *v1beta11.UserId `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	// REQUIRED.
	ResourceId *v1beta1.ResourceId `protobuf:"bytes,3,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// REQUIRED.
	Grantee              *v1beta1.Grantee `protobuf:"bytes,4,opt,name=grantee,proto3" json:"grantee,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ShareKey) Reset()         { *m = ShareKey{} }
func (m *ShareKey) String() string { return proto.CompactTextString(m) }
func (*ShareKey) ProtoMessage()    {}
func (*ShareKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b387acee7aaf43, []int{3}
}

func (m *ShareKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShareKey.Unmarshal(m, b)
}
func (m *ShareKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShareKey.Marshal(b, m, deterministic)
}
func (m *ShareKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShareKey.Merge(m, src)
}
func (m *ShareKey) XXX_Size() int {
	return xxx_messageInfo_ShareKey.Size(m)
}
func (m *ShareKey) XXX_DiscardUnknown() {
	xxx_messageInfo_ShareKey.DiscardUnknown(m)
}

var xxx_messageInfo_ShareKey proto.InternalMessageInfo

func (m *ShareKey) GetOwner() *v1beta11.UserId {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *ShareKey) GetResourceId() *v1beta1.ResourceId {
	if m != nil {
		return m.ResourceId
	}
	return nil
}

func (m *ShareKey) GetGrantee() *v1beta1.Grantee {
	if m != nil {
		return m.Grantee
	}
	return nil
}

// A share id identifies uniquely a // share in the share provider namespace.
// A ShareId MUST be unique inside the share provider.
type ShareId struct {
	// REQUIRED.
	// The internal id used by service implementor to
	// uniquely identity the share in the internal
	// implementation of the service.
	OpaqueId             string   `protobuf:"bytes,2,opt,name=opaque_id,json=opaqueId,proto3" json:"opaque_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShareId) Reset()         { *m = ShareId{} }
func (m *ShareId) String() string { return proto.CompactTextString(m) }
func (*ShareId) ProtoMessage()    {}
func (*ShareId) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b387acee7aaf43, []int{4}
}

func (m *ShareId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShareId.Unmarshal(m, b)
}
func (m *ShareId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShareId.Marshal(b, m, deterministic)
}
func (m *ShareId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShareId.Merge(m, src)
}
func (m *ShareId) XXX_Size() int {
	return xxx_messageInfo_ShareId.Size(m)
}
func (m *ShareId) XXX_DiscardUnknown() {
	xxx_messageInfo_ShareId.DiscardUnknown(m)
}

var xxx_messageInfo_ShareId proto.InternalMessageInfo

func (m *ShareId) GetOpaqueId() string {
	if m != nil {
		return m.OpaqueId
	}
	return ""
}

// The mechanism to identify a share
// in the share provider namespace.
type ShareReference struct {
	// REQUIRED.
	// One of the specifications MUST be specified.
	//
	// Types that are valid to be assigned to Spec:
	//	*ShareReference_Id
	//	*ShareReference_Key
	Spec                 isShareReference_Spec `protobuf_oneof:"spec"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ShareReference) Reset()         { *m = ShareReference{} }
func (m *ShareReference) String() string { return proto.CompactTextString(m) }
func (*ShareReference) ProtoMessage()    {}
func (*ShareReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b387acee7aaf43, []int{5}
}

func (m *ShareReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShareReference.Unmarshal(m, b)
}
func (m *ShareReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShareReference.Marshal(b, m, deterministic)
}
func (m *ShareReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShareReference.Merge(m, src)
}
func (m *ShareReference) XXX_Size() int {
	return xxx_messageInfo_ShareReference.Size(m)
}
func (m *ShareReference) XXX_DiscardUnknown() {
	xxx_messageInfo_ShareReference.DiscardUnknown(m)
}

var xxx_messageInfo_ShareReference proto.InternalMessageInfo

type isShareReference_Spec interface {
	isShareReference_Spec()
}

type ShareReference_Id struct {
	Id *ShareId `protobuf:"bytes,1,opt,name=id,proto3,oneof"`
}

type ShareReference_Key struct {
	Key *ShareKey `protobuf:"bytes,2,opt,name=key,proto3,oneof"`
}

func (*ShareReference_Id) isShareReference_Spec() {}

func (*ShareReference_Key) isShareReference_Spec() {}

func (m *ShareReference) GetSpec() isShareReference_Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *ShareReference) GetId() *ShareId {
	if x, ok := m.GetSpec().(*ShareReference_Id); ok {
		return x.Id
	}
	return nil
}

func (m *ShareReference) GetKey() *ShareKey {
	if x, ok := m.GetSpec().(*ShareReference_Key); ok {
		return x.Key
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ShareReference) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ShareReference_Id)(nil),
		(*ShareReference_Key)(nil),
	}
}

// A share grant specifies the share permissions
// for a grantee.
type ShareGrant struct {
	// REQUIRED.
	// The grantee of the grant.
	Grantee *v1beta1.Grantee `protobuf:"bytes,1,opt,name=grantee,proto3" json:"grantee,omitempty"`
	// REQUIRED.
	// The share permissions for the grant.
	Permissions          *SharePermissions `protobuf:"bytes,2,opt,name=permissions,proto3" json:"permissions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ShareGrant) Reset()         { *m = ShareGrant{} }
func (m *ShareGrant) String() string { return proto.CompactTextString(m) }
func (*ShareGrant) ProtoMessage()    {}
func (*ShareGrant) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b387acee7aaf43, []int{6}
}

func (m *ShareGrant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShareGrant.Unmarshal(m, b)
}
func (m *ShareGrant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShareGrant.Marshal(b, m, deterministic)
}
func (m *ShareGrant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShareGrant.Merge(m, src)
}
func (m *ShareGrant) XXX_Size() int {
	return xxx_messageInfo_ShareGrant.Size(m)
}
func (m *ShareGrant) XXX_DiscardUnknown() {
	xxx_messageInfo_ShareGrant.DiscardUnknown(m)
}

var xxx_messageInfo_ShareGrant proto.InternalMessageInfo

func (m *ShareGrant) GetGrantee() *v1beta1.Grantee {
	if m != nil {
		return m.Grantee
	}
	return nil
}

func (m *ShareGrant) GetPermissions() *SharePermissions {
	if m != nil {
		return m.Permissions
	}
	return nil
}

// Details of the sync'n'share system provider.
type ProviderInfo struct {
	// REQUIRED.
	// The domain of the sync'n'share provider.
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	// REQUIRED.
	// The API version supported by the provider.
	ApiVersion string `protobuf:"bytes,2,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// REQUIRED.
	// The endpoint on which the API is exposed.
	ApiEndpoint string `protobuf:"bytes,3,opt,name=api_endpoint,json=apiEndpoint,proto3" json:"api_endpoint,omitempty"`
	// REQUIRED.
	// The endpoint on which the webdav protocol is exposed.
	WebdavEndpoint       string   `protobuf:"bytes,4,opt,name=webdav_endpoint,json=webdavEndpoint,proto3" json:"webdav_endpoint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProviderInfo) Reset()         { *m = ProviderInfo{} }
func (m *ProviderInfo) String() string { return proto.CompactTextString(m) }
func (*ProviderInfo) ProtoMessage()    {}
func (*ProviderInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b387acee7aaf43, []int{7}
}

func (m *ProviderInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProviderInfo.Unmarshal(m, b)
}
func (m *ProviderInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProviderInfo.Marshal(b, m, deterministic)
}
func (m *ProviderInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProviderInfo.Merge(m, src)
}
func (m *ProviderInfo) XXX_Size() int {
	return xxx_messageInfo_ProviderInfo.Size(m)
}
func (m *ProviderInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ProviderInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ProviderInfo proto.InternalMessageInfo

func (m *ProviderInfo) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *ProviderInfo) GetApiVersion() string {
	if m != nil {
		return m.ApiVersion
	}
	return ""
}

func (m *ProviderInfo) GetApiEndpoint() string {
	if m != nil {
		return m.ApiEndpoint
	}
	return ""
}

func (m *ProviderInfo) GetWebdavEndpoint() string {
	if m != nil {
		return m.WebdavEndpoint
	}
	return ""
}

// List of users a user can share resources with.
type UserShareList struct {
	// REQUIRED.
	// The user who can share the resources.
	UserId *v1beta11.UserId `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// REQUIRED.
	// The list of users.
	Users                []*v1beta11.UserId `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *UserShareList) Reset()         { *m = UserShareList{} }
func (m *UserShareList) String() string { return proto.CompactTextString(m) }
func (*UserShareList) ProtoMessage()    {}
func (*UserShareList) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b387acee7aaf43, []int{8}
}

func (m *UserShareList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserShareList.Unmarshal(m, b)
}
func (m *UserShareList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserShareList.Marshal(b, m, deterministic)
}
func (m *UserShareList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserShareList.Merge(m, src)
}
func (m *UserShareList) XXX_Size() int {
	return xxx_messageInfo_UserShareList.Size(m)
}
func (m *UserShareList) XXX_DiscardUnknown() {
	xxx_messageInfo_UserShareList.DiscardUnknown(m)
}

var xxx_messageInfo_UserShareList proto.InternalMessageInfo

func (m *UserShareList) GetUserId() *v1beta11.UserId {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *UserShareList) GetUsers() []*v1beta11.UserId {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterEnum("cs3.sharing.ocm.v1beta1.ShareState", ShareState_name, ShareState_value)
	proto.RegisterType((*Share)(nil), "cs3.sharing.ocm.v1beta1.Share")
	proto.RegisterType((*SharePermissions)(nil), "cs3.sharing.ocm.v1beta1.SharePermissions")
	proto.RegisterType((*ReceivedShare)(nil), "cs3.sharing.ocm.v1beta1.ReceivedShare")
	proto.RegisterType((*ShareKey)(nil), "cs3.sharing.ocm.v1beta1.ShareKey")
	proto.RegisterType((*ShareId)(nil), "cs3.sharing.ocm.v1beta1.ShareId")
	proto.RegisterType((*ShareReference)(nil), "cs3.sharing.ocm.v1beta1.ShareReference")
	proto.RegisterType((*ShareGrant)(nil), "cs3.sharing.ocm.v1beta1.ShareGrant")
	proto.RegisterType((*ProviderInfo)(nil), "cs3.sharing.ocm.v1beta1.ProviderInfo")
	proto.RegisterType((*UserShareList)(nil), "cs3.sharing.ocm.v1beta1.UserShareList")
}

func init() {
	proto.RegisterFile("cs3/sharing/ocm/v1beta1/resources.proto", fileDescriptor_b8b387acee7aaf43)
}

var fileDescriptor_b8b387acee7aaf43 = []byte{
	// 770 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0xdf, 0x6e, 0xfb, 0x34,
	0x14, 0xc7, 0x97, 0x74, 0xfd, 0x77, 0xba, 0x95, 0xca, 0x20, 0x16, 0x6d, 0xfc, 0xd9, 0x82, 0x60,
	0x1b, 0x42, 0x29, 0x5d, 0x41, 0x08, 0xb8, 0x40, 0x6d, 0x17, 0x6d, 0x61, 0xd3, 0x56, 0xb9, 0x65,
	0x17, 0x68, 0x52, 0x95, 0x25, 0xde, 0x66, 0xa1, 0xc4, 0xc1, 0x4e, 0x3b, 0xf5, 0x0e, 0x04, 0x6f,
	0xc0, 0x1d, 0x97, 0x5c, 0xf2, 0x28, 0xdc, 0xf0, 0x2a, 0x3c, 0x02, 0xb2, 0x9d, 0x74, 0x5b, 0xa5,
	0x56, 0xdd, 0x6f, 0x57, 0xad, 0x7d, 0xbe, 0x1f, 0xfb, 0xe4, 0x7c, 0xcf, 0x49, 0x60, 0x3f, 0x10,
	0xed, 0xa6, 0xb8, 0xf7, 0x39, 0x8d, 0xef, 0x9a, 0x2c, 0x88, 0x9a, 0x93, 0xd6, 0x0d, 0x49, 0xfd,
	0x56, 0x93, 0x13, 0xc1, 0xc6, 0x3c, 0x20, 0xc2, 0x49, 0x38, 0x4b, 0x19, 0xda, 0x0a, 0x44, 0xdb,
	0xc9, 0x84, 0x0e, 0x0b, 0x22, 0x27, 0x13, 0x6e, 0x1f, 0xca, 0x13, 0x68, 0x48, 0xe2, 0x94, 0xa6,
	0xd3, 0xe6, 0x58, 0x10, 0xbe, 0xe8, 0x8c, 0xed, 0xcf, 0xd4, 0x65, 0x29, 0xe3, 0xfe, 0x1d, 0x69,
	0x26, 0x9c, 0x4d, 0x68, 0xb8, 0x44, 0xfd, 0xbe, 0x54, 0xa7, 0xd3, 0x84, 0x88, 0x99, 0x44, 0xad,
	0x74, 0xd8, 0xfe, 0xaf, 0x00, 0xc5, 0xc1, 0xbd, 0xcf, 0x09, 0xfa, 0x1c, 0x4c, 0x1a, 0x5a, 0xc6,
	0xae, 0x71, 0x50, 0x3b, 0xda, 0x75, 0x16, 0xe4, 0xe9, 0x28, 0xad, 0x17, 0x62, 0x93, 0x86, 0xc8,
	0x83, 0x5a, 0x7e, 0xdb, 0x88, 0x86, 0x96, 0xa9, 0xd0, 0x03, 0x8d, 0xea, 0xf4, 0x9c, 0x3c, 0xbd,
	0x19, 0x8f, 0x33, 0xc0, 0x0b, 0x31, 0xf0, 0xd9, 0x7f, 0x74, 0x06, 0xb5, 0x84, 0xf0, 0x88, 0x0a,
	0x41, 0x59, 0x2c, 0xac, 0x82, 0x3a, 0xea, 0x70, 0x79, 0x16, 0xfd, 0x47, 0x00, 0x3f, 0xa5, 0xd1,
	0x77, 0x50, 0xbe, 0xe3, 0x7e, 0x9c, 0x12, 0x62, 0xad, 0xab, 0x83, 0x3e, 0x5e, 0x9e, 0xd3, 0x89,
	0x16, 0xe3, 0x9c, 0x42, 0x5f, 0x41, 0x91, 0x3d, 0xc4, 0x84, 0x5b, 0x45, 0x85, 0xef, 0x29, 0x3c,
	0x37, 0xc7, 0x91, 0xe6, 0xcc, 0xd8, 0x1f, 0x04, 0xe1, 0x5e, 0x88, 0xb5, 0x1e, 0x7d, 0x0b, 0xe5,
	0x80, 0x13, 0x3f, 0x65, 0xdc, 0x2a, 0xad, 0x8a, 0xe6, 0x04, 0x3a, 0x82, 0x62, 0x90, 0xd2, 0x88,
	0x58, 0x65, 0x85, 0xbe, 0xa7, 0x50, 0xed, 0x55, 0x8e, 0x0c, 0x69, 0x44, 0x44, 0xea, 0x47, 0x09,
	0xd6, 0x52, 0xc9, 0x44, 0x8a, 0xa9, 0xac, 0xc2, 0x28, 0xa9, 0xfd, 0xab, 0x01, 0x8d, 0xf9, 0x02,
	0xa2, 0xc1, 0x73, 0x03, 0x74, 0x1b, 0xb4, 0x56, 0xf3, 0x72, 0xa1, 0x11, 0x16, 0x94, 0x39, 0x91,
	0xfe, 0x11, 0xd5, 0x1c, 0x15, 0x9c, 0x2f, 0xed, 0x5f, 0x0c, 0xd8, 0xc4, 0x24, 0x20, 0x74, 0x42,
	0x42, 0xdd, 0x7e, 0x5f, 0x40, 0x51, 0x2b, 0xf5, 0xd5, 0x1f, 0x2c, 0xf7, 0x1e, 0x6b, 0x31, 0xfa,
	0x1a, 0x8a, 0x22, 0xf5, 0x53, 0x7d, 0x7e, 0xfd, 0xe8, 0xa3, 0xe5, 0xd4, 0x40, 0x4a, 0xb1, 0x26,
	0xec, 0x7f, 0x0d, 0xa8, 0xa8, 0xdd, 0x33, 0x32, 0x7d, 0x74, 0xdc, 0x7c, 0xa1, 0xe3, 0x73, 0x33,
	0x50, 0x78, 0xc5, 0x0c, 0xbc, 0xb6, 0x6d, 0xed, 0x4f, 0xa0, 0x9c, 0x8d, 0x27, 0xda, 0x81, 0x2a,
	0x4b, 0xfc, 0x9f, 0xc7, 0xb3, 0xc1, 0xac, 0xe2, 0x8a, 0xde, 0xf0, 0x42, 0xfb, 0x37, 0x03, 0xea,
	0xba, 0x8a, 0xe4, 0x96, 0x70, 0x12, 0x07, 0xb2, 0x8f, 0x5e, 0x30, 0xfc, 0xa7, 0x6b, 0x6a, 0xfc,
	0xbf, 0x84, 0xc2, 0x4f, 0x64, 0xfa, 0xac, 0x62, 0x0b, 0xa1, 0x33, 0x32, 0x3d, 0x5d, 0xc3, 0x52,
	0xdf, 0x2d, 0xc1, 0xba, 0x48, 0x48, 0x60, 0xff, 0x69, 0x00, 0xa8, 0x98, 0x7a, 0x8e, 0xa7, 0x4f,
	0x6f, 0xbc, 0xd1, 0xd0, 0xce, 0xbd, 0x42, 0xcc, 0xd7, 0xbc, 0x42, 0xec, 0x3f, 0x0c, 0xd8, 0xe8,
	0x67, 0x57, 0x7a, 0xf1, 0x2d, 0x43, 0xef, 0x42, 0x29, 0x64, 0x91, 0x4f, 0x63, 0x95, 0x5d, 0x15,
	0x67, 0x2b, 0xf4, 0x21, 0xd4, 0xfc, 0x84, 0x8e, 0x26, 0x84, 0x4b, 0x30, 0x2b, 0x35, 0xf8, 0x09,
	0xbd, 0xd2, 0x3b, 0x68, 0x0f, 0x36, 0xa4, 0x80, 0xc4, 0x61, 0xc2, 0x68, 0x9c, 0xaa, 0x0e, 0xa9,
	0x62, 0x09, 0xb9, 0xd9, 0x16, 0xda, 0x87, 0xb7, 0x1e, 0xc8, 0x4d, 0xe8, 0x4f, 0x1e, 0x55, 0xeb,
	0x4a, 0x55, 0xd7, 0xdb, 0xb9, 0xd0, 0xfe, 0xdd, 0x80, 0x4d, 0xd9, 0x7e, 0x2a, 0xf7, 0x73, 0x2a,
	0x52, 0xf4, 0x0d, 0x94, 0x65, 0x73, 0x8e, 0x66, 0xe6, 0xad, 0xd0, 0xb9, 0xa5, 0xb1, 0xfa, 0x95,
	0x3d, 0x2f, 0xff, 0xc9, 0x52, 0x15, 0x56, 0xec, 0x79, 0xa5, 0xff, 0x94, 0x67, 0xc6, 0xa9, 0x71,
	0x42, 0x5b, 0xf0, 0xf6, 0xe0, 0xb4, 0x83, 0xdd, 0xd1, 0x60, 0xd8, 0x19, 0xba, 0x23, 0xef, 0xe2,
	0xaa, 0x73, 0xee, 0x1d, 0x37, 0xd6, 0xe6, 0x03, 0x7d, 0xf7, 0xe2, 0xd8, 0xbb, 0x38, 0x69, 0x18,
	0xc8, 0x82, 0x77, 0x9e, 0x06, 0x3a, 0xbd, 0x9e, 0xdb, 0x1f, 0xba, 0xc7, 0x0d, 0x73, 0x3e, 0x82,
	0xdd, 0xef, 0xdd, 0x9e, 0x8c, 0x14, 0xba, 0x63, 0xd8, 0x09, 0x58, 0xb4, 0xc8, 0xcd, 0x6e, 0x3d,
	0x9f, 0x29, 0xd1, 0x97, 0x9f, 0xb5, 0xbe, 0xf1, 0x23, 0xb0, 0x20, 0xca, 0xa2, 0x7f, 0x99, 0x85,
	0xde, 0xe0, 0xf2, 0x6f, 0x73, 0xab, 0x27, 0xda, 0xca, 0x77, 0x49, 0x5f, 0x06, 0x91, 0x73, 0xd5,
	0xea, 0xca, 0xf8, 0x3f, 0x2a, 0x72, 0x9d, 0x45, 0xae, 0x2f, 0x83, 0xe8, 0x3a, 0x8b, 0xdc, 0x94,
	0xd4, 0x57, 0xb2, 0xfd, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x5b, 0x30, 0xac, 0xe1, 0x07,
	0x00, 0x00,
}
