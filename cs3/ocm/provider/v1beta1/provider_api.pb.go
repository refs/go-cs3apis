// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs3/ocm/provider/v1beta1/provider_api.proto

package providerv1beta1

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

type IsProviderAllowedRequest struct {
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The provider that we need to check against the list of verified mesh providers.
	Provider             *ProviderInfo `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *IsProviderAllowedRequest) Reset()         { *m = IsProviderAllowedRequest{} }
func (m *IsProviderAllowedRequest) String() string { return proto.CompactTextString(m) }
func (*IsProviderAllowedRequest) ProtoMessage()    {}
func (*IsProviderAllowedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9d7362468fbd522, []int{0}
}

func (m *IsProviderAllowedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsProviderAllowedRequest.Unmarshal(m, b)
}
func (m *IsProviderAllowedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsProviderAllowedRequest.Marshal(b, m, deterministic)
}
func (m *IsProviderAllowedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsProviderAllowedRequest.Merge(m, src)
}
func (m *IsProviderAllowedRequest) XXX_Size() int {
	return xxx_messageInfo_IsProviderAllowedRequest.Size(m)
}
func (m *IsProviderAllowedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IsProviderAllowedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IsProviderAllowedRequest proto.InternalMessageInfo

func (m *IsProviderAllowedRequest) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *IsProviderAllowedRequest) GetProvider() *ProviderInfo {
	if m != nil {
		return m.Provider
	}
	return nil
}

type IsProviderAllowedResponse struct {
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

func (m *IsProviderAllowedResponse) Reset()         { *m = IsProviderAllowedResponse{} }
func (m *IsProviderAllowedResponse) String() string { return proto.CompactTextString(m) }
func (*IsProviderAllowedResponse) ProtoMessage()    {}
func (*IsProviderAllowedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9d7362468fbd522, []int{1}
}

func (m *IsProviderAllowedResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsProviderAllowedResponse.Unmarshal(m, b)
}
func (m *IsProviderAllowedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsProviderAllowedResponse.Marshal(b, m, deterministic)
}
func (m *IsProviderAllowedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsProviderAllowedResponse.Merge(m, src)
}
func (m *IsProviderAllowedResponse) XXX_Size() int {
	return xxx_messageInfo_IsProviderAllowedResponse.Size(m)
}
func (m *IsProviderAllowedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IsProviderAllowedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IsProviderAllowedResponse proto.InternalMessageInfo

func (m *IsProviderAllowedResponse) GetStatus() *v1beta11.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *IsProviderAllowedResponse) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

type GetInfoByDomainRequest struct {
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The domain of the system provider.
	Domain               string   `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetInfoByDomainRequest) Reset()         { *m = GetInfoByDomainRequest{} }
func (m *GetInfoByDomainRequest) String() string { return proto.CompactTextString(m) }
func (*GetInfoByDomainRequest) ProtoMessage()    {}
func (*GetInfoByDomainRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9d7362468fbd522, []int{2}
}

func (m *GetInfoByDomainRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInfoByDomainRequest.Unmarshal(m, b)
}
func (m *GetInfoByDomainRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInfoByDomainRequest.Marshal(b, m, deterministic)
}
func (m *GetInfoByDomainRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInfoByDomainRequest.Merge(m, src)
}
func (m *GetInfoByDomainRequest) XXX_Size() int {
	return xxx_messageInfo_GetInfoByDomainRequest.Size(m)
}
func (m *GetInfoByDomainRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInfoByDomainRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetInfoByDomainRequest proto.InternalMessageInfo

func (m *GetInfoByDomainRequest) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *GetInfoByDomainRequest) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

type GetInfoByDomainResponse struct {
	// REQUIRED.
	// The response status.
	Status *v1beta11.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The info of the provider
	ProviderInfo         *ProviderInfo `protobuf:"bytes,3,opt,name=provider_info,json=providerInfo,proto3" json:"provider_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetInfoByDomainResponse) Reset()         { *m = GetInfoByDomainResponse{} }
func (m *GetInfoByDomainResponse) String() string { return proto.CompactTextString(m) }
func (*GetInfoByDomainResponse) ProtoMessage()    {}
func (*GetInfoByDomainResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9d7362468fbd522, []int{3}
}

func (m *GetInfoByDomainResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInfoByDomainResponse.Unmarshal(m, b)
}
func (m *GetInfoByDomainResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInfoByDomainResponse.Marshal(b, m, deterministic)
}
func (m *GetInfoByDomainResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInfoByDomainResponse.Merge(m, src)
}
func (m *GetInfoByDomainResponse) XXX_Size() int {
	return xxx_messageInfo_GetInfoByDomainResponse.Size(m)
}
func (m *GetInfoByDomainResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInfoByDomainResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetInfoByDomainResponse proto.InternalMessageInfo

func (m *GetInfoByDomainResponse) GetStatus() *v1beta11.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *GetInfoByDomainResponse) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *GetInfoByDomainResponse) GetProviderInfo() *ProviderInfo {
	if m != nil {
		return m.ProviderInfo
	}
	return nil
}

type ListAllProvidersRequest struct {
	// OPTIONAL.
	// Opaque information.
	Opaque               *v1beta1.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ListAllProvidersRequest) Reset()         { *m = ListAllProvidersRequest{} }
func (m *ListAllProvidersRequest) String() string { return proto.CompactTextString(m) }
func (*ListAllProvidersRequest) ProtoMessage()    {}
func (*ListAllProvidersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9d7362468fbd522, []int{4}
}

func (m *ListAllProvidersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAllProvidersRequest.Unmarshal(m, b)
}
func (m *ListAllProvidersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAllProvidersRequest.Marshal(b, m, deterministic)
}
func (m *ListAllProvidersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAllProvidersRequest.Merge(m, src)
}
func (m *ListAllProvidersRequest) XXX_Size() int {
	return xxx_messageInfo_ListAllProvidersRequest.Size(m)
}
func (m *ListAllProvidersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAllProvidersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListAllProvidersRequest proto.InternalMessageInfo

func (m *ListAllProvidersRequest) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

type ListAllProvidersResponse struct {
	// REQUIRED.
	// The response status.
	Status *v1beta11.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The share.
	Providers            []*ProviderInfo `protobuf:"bytes,3,rep,name=providers,proto3" json:"providers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ListAllProvidersResponse) Reset()         { *m = ListAllProvidersResponse{} }
func (m *ListAllProvidersResponse) String() string { return proto.CompactTextString(m) }
func (*ListAllProvidersResponse) ProtoMessage()    {}
func (*ListAllProvidersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9d7362468fbd522, []int{5}
}

func (m *ListAllProvidersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAllProvidersResponse.Unmarshal(m, b)
}
func (m *ListAllProvidersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAllProvidersResponse.Marshal(b, m, deterministic)
}
func (m *ListAllProvidersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAllProvidersResponse.Merge(m, src)
}
func (m *ListAllProvidersResponse) XXX_Size() int {
	return xxx_messageInfo_ListAllProvidersResponse.Size(m)
}
func (m *ListAllProvidersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAllProvidersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListAllProvidersResponse proto.InternalMessageInfo

func (m *ListAllProvidersResponse) GetStatus() *v1beta11.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ListAllProvidersResponse) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *ListAllProvidersResponse) GetProviders() []*ProviderInfo {
	if m != nil {
		return m.Providers
	}
	return nil
}

func init() {
	proto.RegisterType((*IsProviderAllowedRequest)(nil), "cs3.ocm.provider.v1beta1.IsProviderAllowedRequest")
	proto.RegisterType((*IsProviderAllowedResponse)(nil), "cs3.ocm.provider.v1beta1.IsProviderAllowedResponse")
	proto.RegisterType((*GetInfoByDomainRequest)(nil), "cs3.ocm.provider.v1beta1.GetInfoByDomainRequest")
	proto.RegisterType((*GetInfoByDomainResponse)(nil), "cs3.ocm.provider.v1beta1.GetInfoByDomainResponse")
	proto.RegisterType((*ListAllProvidersRequest)(nil), "cs3.ocm.provider.v1beta1.ListAllProvidersRequest")
	proto.RegisterType((*ListAllProvidersResponse)(nil), "cs3.ocm.provider.v1beta1.ListAllProvidersResponse")
}

func init() {
	proto.RegisterFile("cs3/ocm/provider/v1beta1/provider_api.proto", fileDescriptor_f9d7362468fbd522)
}

var fileDescriptor_f9d7362468fbd522 = []byte{
	// 468 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xcd, 0x6a, 0xdb, 0x30,
	0x1c, 0xc7, 0x0e, 0x84, 0xf5, 0xdf, 0x8d, 0x76, 0x3a, 0x34, 0xaa, 0xe9, 0xa0, 0xf8, 0x30, 0x0a,
	0x03, 0x65, 0x8e, 0x9f, 0xa0, 0x6e, 0x61, 0x84, 0x15, 0x62, 0x3c, 0xd8, 0x61, 0x04, 0x86, 0xab,
	0xa8, 0x60, 0x88, 0x23, 0xd5, 0x92, 0x33, 0x02, 0x83, 0x3e, 0xc3, 0x5e, 0x61, 0xc7, 0x3d, 0xc7,
	0x4e, 0x7d, 0xa8, 0x31, 0x2c, 0x4b, 0x5a, 0x69, 0xea, 0x31, 0xaf, 0xd0, 0xa3, 0xfc, 0xfb, 0xf8,
	0x7f, 0xe8, 0x67, 0xc1, 0x1b, 0x2a, 0xe3, 0x31, 0xa7, 0xe5, 0x58, 0x54, 0x7c, 0x5d, 0x2c, 0x58,
	0x35, 0x5e, 0x47, 0x97, 0x4c, 0xe5, 0x91, 0xfb, 0xf0, 0x39, 0x17, 0x05, 0x11, 0x15, 0x57, 0x1c,
	0x61, 0x2a, 0x63, 0xc2, 0x69, 0x49, 0x2c, 0x46, 0x0c, 0x39, 0x38, 0xe9, 0xb4, 0xa9, 0x98, 0xe4,
	0x75, 0x45, 0x99, 0x6c, 0x3d, 0x82, 0xa3, 0x86, 0x59, 0x09, 0xea, 0x08, 0x52, 0xe5, 0xaa, 0xb6,
	0xe8, 0xab, 0x06, 0x55, 0x1b, 0xc1, 0xa4, 0xc3, 0xf5, 0xa9, 0x85, 0xc3, 0x6f, 0x1e, 0xe0, 0xa9,
	0x4c, 0x4d, 0x8d, 0xd3, 0xe5, 0x92, 0x7f, 0x61, 0x8b, 0x8c, 0x5d, 0xd7, 0x4c, 0x2a, 0x14, 0xc1,
	0x90, 0x8b, 0xfc, 0xba, 0x66, 0xd8, 0x3b, 0xf6, 0x4e, 0x76, 0x27, 0x87, 0xa4, 0x69, 0xb7, 0x95,
	0x1b, 0x33, 0x32, 0xd3, 0x84, 0xcc, 0x10, 0x51, 0x02, 0xcf, 0x6c, 0xc3, 0xd8, 0xd7, 0xa2, 0xd7,
	0xa4, 0x6b, 0x46, 0x62, 0xcb, 0x4e, 0x57, 0x57, 0x3c, 0x73, 0xba, 0xf0, 0x06, 0x0e, 0x1f, 0x68,
	0x49, 0x0a, 0xbe, 0x92, 0x0c, 0x8d, 0x61, 0xd8, 0xce, 0x67, 0x7a, 0x1a, 0x69, 0xfb, 0x4a, 0x50,
	0xe7, 0xfa, 0x41, 0xc3, 0x99, 0xa1, 0xdd, 0x19, 0xc2, 0xff, 0xc7, 0x21, 0x42, 0x0a, 0x07, 0xef,
	0x98, 0x6a, 0xba, 0x4a, 0x36, 0xe7, 0xbc, 0xcc, 0x8b, 0xd5, 0x23, 0x36, 0x72, 0x00, 0xc3, 0x85,
	0xf6, 0xd0, 0xf5, 0x77, 0x32, 0x73, 0x0a, 0x6f, 0x3d, 0x18, 0x6d, 0x55, 0x79, 0xba, 0x21, 0xd1,
	0x7b, 0x78, 0xe1, 0x02, 0x59, 0xac, 0xae, 0x38, 0x1e, 0xf4, 0xba, 0xae, 0xe7, 0xe2, 0xce, 0x29,
	0xbc, 0x80, 0xd1, 0x45, 0x21, 0xd5, 0xe9, 0x72, 0x69, 0x49, 0xf2, 0xff, 0x57, 0x16, 0xfe, 0xf4,
	0x00, 0x6f, 0xdb, 0x3d, 0xe1, 0x6e, 0xce, 0x61, 0xc7, 0x8e, 0x27, 0xf1, 0xe0, 0x78, 0xd0, 0x63,
	0x2f, 0x7f, 0x84, 0x93, 0x5f, 0x3e, 0xec, 0xba, 0x18, 0xa7, 0x53, 0xf4, 0x15, 0x5e, 0x6e, 0xe5,
	0x1a, 0x4d, 0xba, 0x7d, 0xbb, 0xfe, 0xcb, 0x20, 0xee, 0xa5, 0x31, 0x7b, 0x5b, 0xc3, 0xde, 0xbd,
	0xb8, 0xa1, 0xb7, 0xdd, 0x3e, 0x0f, 0xe7, 0x3f, 0x88, 0x7a, 0x28, 0x4c, 0xdd, 0x0d, 0xec, 0xdf,
	0xbf, 0x4b, 0xf4, 0x17, 0x9b, 0x8e, 0x18, 0x05, 0x93, 0x3e, 0x92, 0xb6, 0x74, 0x72, 0x03, 0x47,
	0x94, 0x97, 0x9d, 0xc2, 0x64, 0xdf, 0xed, 0x4a, 0x14, 0x69, 0xf3, 0x1c, 0xa6, 0xde, 0xa7, 0x3d,
	0xcb, 0x32, 0xa4, 0xef, 0xfe, 0xe0, 0x6c, 0x96, 0xfe, 0xf0, 0xf1, 0x99, 0x8c, 0xc9, 0x8c, 0x96,
	0xee, 0xb6, 0xc9, 0xc7, 0x28, 0x69, 0x08, 0xb7, 0x1a, 0x9a, 0xcf, 0x68, 0x39, 0xb7, 0xd0, 0xdc,
	0x40, 0x97, 0x43, 0xfd, 0xc8, 0xc6, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x28, 0xc2, 0x2b, 0xf7,
	0x14, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProviderAPIClient is the client API for ProviderAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProviderAPIClient interface {
	// Check if a given system provider is registered in the mesh or not.
	// MUST return CODE_UNAUTHENTICATED if the system is not registered
	IsProviderAllowed(ctx context.Context, in *IsProviderAllowedRequest, opts ...grpc.CallOption) (*IsProviderAllowedResponse, error)
	// Get the information of the provider identified by a specific domain.
	// MUST return CODE_NOT_FOUND if the sync'n'share system provider does not exist.
	GetInfoByDomain(ctx context.Context, in *GetInfoByDomainRequest, opts ...grpc.CallOption) (*GetInfoByDomainResponse, error)
	// Get the information of all the providers registered in the mesh.
	ListAllProviders(ctx context.Context, in *ListAllProvidersRequest, opts ...grpc.CallOption) (*ListAllProvidersResponse, error)
}

type providerAPIClient struct {
	cc *grpc.ClientConn
}

func NewProviderAPIClient(cc *grpc.ClientConn) ProviderAPIClient {
	return &providerAPIClient{cc}
}

func (c *providerAPIClient) IsProviderAllowed(ctx context.Context, in *IsProviderAllowedRequest, opts ...grpc.CallOption) (*IsProviderAllowedResponse, error) {
	out := new(IsProviderAllowedResponse)
	err := c.cc.Invoke(ctx, "/cs3.ocm.provider.v1beta1.ProviderAPI/IsProviderAllowed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerAPIClient) GetInfoByDomain(ctx context.Context, in *GetInfoByDomainRequest, opts ...grpc.CallOption) (*GetInfoByDomainResponse, error) {
	out := new(GetInfoByDomainResponse)
	err := c.cc.Invoke(ctx, "/cs3.ocm.provider.v1beta1.ProviderAPI/GetInfoByDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerAPIClient) ListAllProviders(ctx context.Context, in *ListAllProvidersRequest, opts ...grpc.CallOption) (*ListAllProvidersResponse, error) {
	out := new(ListAllProvidersResponse)
	err := c.cc.Invoke(ctx, "/cs3.ocm.provider.v1beta1.ProviderAPI/ListAllProviders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProviderAPIServer is the server API for ProviderAPI service.
type ProviderAPIServer interface {
	// Check if a given system provider is registered in the mesh or not.
	// MUST return CODE_UNAUTHENTICATED if the system is not registered
	IsProviderAllowed(context.Context, *IsProviderAllowedRequest) (*IsProviderAllowedResponse, error)
	// Get the information of the provider identified by a specific domain.
	// MUST return CODE_NOT_FOUND if the sync'n'share system provider does not exist.
	GetInfoByDomain(context.Context, *GetInfoByDomainRequest) (*GetInfoByDomainResponse, error)
	// Get the information of all the providers registered in the mesh.
	ListAllProviders(context.Context, *ListAllProvidersRequest) (*ListAllProvidersResponse, error)
}

// UnimplementedProviderAPIServer can be embedded to have forward compatible implementations.
type UnimplementedProviderAPIServer struct {
}

func (*UnimplementedProviderAPIServer) IsProviderAllowed(ctx context.Context, req *IsProviderAllowedRequest) (*IsProviderAllowedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsProviderAllowed not implemented")
}
func (*UnimplementedProviderAPIServer) GetInfoByDomain(ctx context.Context, req *GetInfoByDomainRequest) (*GetInfoByDomainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfoByDomain not implemented")
}
func (*UnimplementedProviderAPIServer) ListAllProviders(ctx context.Context, req *ListAllProvidersRequest) (*ListAllProvidersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAllProviders not implemented")
}

func RegisterProviderAPIServer(s *grpc.Server, srv ProviderAPIServer) {
	s.RegisterService(&_ProviderAPI_serviceDesc, srv)
}

func _ProviderAPI_IsProviderAllowed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsProviderAllowedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderAPIServer).IsProviderAllowed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cs3.ocm.provider.v1beta1.ProviderAPI/IsProviderAllowed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderAPIServer).IsProviderAllowed(ctx, req.(*IsProviderAllowedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProviderAPI_GetInfoByDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInfoByDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderAPIServer).GetInfoByDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cs3.ocm.provider.v1beta1.ProviderAPI/GetInfoByDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderAPIServer).GetInfoByDomain(ctx, req.(*GetInfoByDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProviderAPI_ListAllProviders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAllProvidersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderAPIServer).ListAllProviders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cs3.ocm.provider.v1beta1.ProviderAPI/ListAllProviders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderAPIServer).ListAllProviders(ctx, req.(*ListAllProvidersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProviderAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cs3.ocm.provider.v1beta1.ProviderAPI",
	HandlerType: (*ProviderAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsProviderAllowed",
			Handler:    _ProviderAPI_IsProviderAllowed_Handler,
		},
		{
			MethodName: "GetInfoByDomain",
			Handler:    _ProviderAPI_GetInfoByDomain_Handler,
		},
		{
			MethodName: "ListAllProviders",
			Handler:    _ProviderAPI_ListAllProviders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cs3/ocm/provider/v1beta1/provider_api.proto",
}
