// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/subspaces/v3/msgs_treasury.proto

package types

import (
	fmt "fmt"
	authz "github.com/cosmos/cosmos-sdk/x/authz"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgGrantTreasuryAuthorization grants an authorization on behalf of the treasury to a user
type MsgGrantTreasuryAuthorization struct {
	// Id of the subspace where the authorization should be granted
	SubspaceID uint64 `protobuf:"varint,1,opt,name=subspace_id,json=subspaceId,proto3" json:"subspace_id,omitempty" yaml:"subspace_id"`
	// Address of the user granting a treasury authorization
	Granter string `protobuf:"bytes,2,opt,name=granter,proto3" json:"granter,omitempty" yaml:"granter"`
	// Address of the user who is being granted a treasury authorization
	Grantee string `protobuf:"bytes,3,opt,name=grantee,proto3" json:"grantee,omitempty" yaml:"grantee"`
	// Grant represents the authorization to execute the provided methods
	Grant authz.Grant `protobuf:"bytes,4,opt,name=grant,proto3" json:"grant" yaml:"grant"`
}

func (m *MsgGrantTreasuryAuthorization) Reset()         { *m = MsgGrantTreasuryAuthorization{} }
func (m *MsgGrantTreasuryAuthorization) String() string { return proto.CompactTextString(m) }
func (*MsgGrantTreasuryAuthorization) ProtoMessage()    {}
func (*MsgGrantTreasuryAuthorization) Descriptor() ([]byte, []int) {
	return fileDescriptor_652ce4e8348c6ab3, []int{0}
}
func (m *MsgGrantTreasuryAuthorization) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgGrantTreasuryAuthorization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgGrantTreasuryAuthorization.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgGrantTreasuryAuthorization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgGrantTreasuryAuthorization.Merge(m, src)
}
func (m *MsgGrantTreasuryAuthorization) XXX_Size() int {
	return m.Size()
}
func (m *MsgGrantTreasuryAuthorization) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgGrantTreasuryAuthorization.DiscardUnknown(m)
}

var xxx_messageInfo_MsgGrantTreasuryAuthorization proto.InternalMessageInfo

func (m *MsgGrantTreasuryAuthorization) GetSubspaceID() uint64 {
	if m != nil {
		return m.SubspaceID
	}
	return 0
}

func (m *MsgGrantTreasuryAuthorization) GetGranter() string {
	if m != nil {
		return m.Granter
	}
	return ""
}

func (m *MsgGrantTreasuryAuthorization) GetGrantee() string {
	if m != nil {
		return m.Grantee
	}
	return ""
}

func (m *MsgGrantTreasuryAuthorization) GetGrant() authz.Grant {
	if m != nil {
		return m.Grant
	}
	return authz.Grant{}
}

// MsgGrantTreasuryAuthorizationResponse defines the Msg/MsgGrantTreasuryAuthorization response type
type MsgGrantTreasuryAuthorizationResponse struct {
}

func (m *MsgGrantTreasuryAuthorizationResponse) Reset()         { *m = MsgGrantTreasuryAuthorizationResponse{} }
func (m *MsgGrantTreasuryAuthorizationResponse) String() string { return proto.CompactTextString(m) }
func (*MsgGrantTreasuryAuthorizationResponse) ProtoMessage()    {}
func (*MsgGrantTreasuryAuthorizationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_652ce4e8348c6ab3, []int{1}
}
func (m *MsgGrantTreasuryAuthorizationResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgGrantTreasuryAuthorizationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgGrantTreasuryAuthorizationResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgGrantTreasuryAuthorizationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgGrantTreasuryAuthorizationResponse.Merge(m, src)
}
func (m *MsgGrantTreasuryAuthorizationResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgGrantTreasuryAuthorizationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgGrantTreasuryAuthorizationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgGrantTreasuryAuthorizationResponse proto.InternalMessageInfo

// MsgRevokeTreasuryAuthorization revokes an existing treasury authorization from a user
type MsgRevokeTreasuryAuthorization struct {
	// Id of the subspace from which the authorization should be revoked
	SubspaceID uint64 `protobuf:"varint,1,opt,name=subspace_id,json=subspaceId,proto3" json:"subspace_id,omitempty" yaml:"subspace_id"`
	// Address of the user revoking the treasury authorization
	Granter string `protobuf:"bytes,2,opt,name=granter,proto3" json:"granter,omitempty" yaml:"granter"`
	// Address of the user who is being revoked the treasury authorization
	Grantee string `protobuf:"bytes,3,opt,name=grantee,proto3" json:"grantee,omitempty" yaml:"grant"`
	// Type url of the authorized message which is being revoked
	MsgTypeUrl string `protobuf:"bytes,4,opt,name=msg_type_url,json=msgTypeUrl,proto3" json:"msg_type_url,omitempty" yaml:"msg_type_url"`
}

func (m *MsgRevokeTreasuryAuthorization) Reset()         { *m = MsgRevokeTreasuryAuthorization{} }
func (m *MsgRevokeTreasuryAuthorization) String() string { return proto.CompactTextString(m) }
func (*MsgRevokeTreasuryAuthorization) ProtoMessage()    {}
func (*MsgRevokeTreasuryAuthorization) Descriptor() ([]byte, []int) {
	return fileDescriptor_652ce4e8348c6ab3, []int{2}
}
func (m *MsgRevokeTreasuryAuthorization) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRevokeTreasuryAuthorization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRevokeTreasuryAuthorization.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRevokeTreasuryAuthorization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRevokeTreasuryAuthorization.Merge(m, src)
}
func (m *MsgRevokeTreasuryAuthorization) XXX_Size() int {
	return m.Size()
}
func (m *MsgRevokeTreasuryAuthorization) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRevokeTreasuryAuthorization.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRevokeTreasuryAuthorization proto.InternalMessageInfo

func (m *MsgRevokeTreasuryAuthorization) GetSubspaceID() uint64 {
	if m != nil {
		return m.SubspaceID
	}
	return 0
}

func (m *MsgRevokeTreasuryAuthorization) GetGranter() string {
	if m != nil {
		return m.Granter
	}
	return ""
}

func (m *MsgRevokeTreasuryAuthorization) GetGrantee() string {
	if m != nil {
		return m.Grantee
	}
	return ""
}

func (m *MsgRevokeTreasuryAuthorization) GetMsgTypeUrl() string {
	if m != nil {
		return m.MsgTypeUrl
	}
	return ""
}

// MsgRevokeTreasuryAuthorizationResponse defines the Msg/MsgRevokeTreasuryAuthorization response type
type MsgRevokeTreasuryAuthorizationResponse struct {
}

func (m *MsgRevokeTreasuryAuthorizationResponse) Reset() {
	*m = MsgRevokeTreasuryAuthorizationResponse{}
}
func (m *MsgRevokeTreasuryAuthorizationResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRevokeTreasuryAuthorizationResponse) ProtoMessage()    {}
func (*MsgRevokeTreasuryAuthorizationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_652ce4e8348c6ab3, []int{3}
}
func (m *MsgRevokeTreasuryAuthorizationResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRevokeTreasuryAuthorizationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRevokeTreasuryAuthorizationResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRevokeTreasuryAuthorizationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRevokeTreasuryAuthorizationResponse.Merge(m, src)
}
func (m *MsgRevokeTreasuryAuthorizationResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRevokeTreasuryAuthorizationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRevokeTreasuryAuthorizationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRevokeTreasuryAuthorizationResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgGrantTreasuryAuthorization)(nil), "desmos.subspaces.v3.MsgGrantTreasuryAuthorization")
	proto.RegisterType((*MsgGrantTreasuryAuthorizationResponse)(nil), "desmos.subspaces.v3.MsgGrantTreasuryAuthorizationResponse")
	proto.RegisterType((*MsgRevokeTreasuryAuthorization)(nil), "desmos.subspaces.v3.MsgRevokeTreasuryAuthorization")
	proto.RegisterType((*MsgRevokeTreasuryAuthorizationResponse)(nil), "desmos.subspaces.v3.MsgRevokeTreasuryAuthorizationResponse")
}

func init() {
	proto.RegisterFile("desmos/subspaces/v3/msgs_treasury.proto", fileDescriptor_652ce4e8348c6ab3)
}

var fileDescriptor_652ce4e8348c6ab3 = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x92, 0xcf, 0x6b, 0xd4, 0x40,
	0x14, 0xc7, 0x77, 0x6a, 0x55, 0x3a, 0x2d, 0x22, 0x69, 0xc1, 0xa5, 0x62, 0xb2, 0x0c, 0x6a, 0x17,
	0xd1, 0x0c, 0xed, 0x7a, 0xd1, 0x9b, 0x41, 0x29, 0x05, 0x7b, 0x89, 0xf5, 0xe2, 0x65, 0x99, 0x6c,
	0x86, 0xd9, 0x60, 0xb2, 0x13, 0xe6, 0x4d, 0x82, 0xe9, 0xbf, 0xe0, 0xc5, 0x3f, 0xab, 0xc7, 0x1e,
	0x3d, 0x05, 0xc9, 0xfe, 0x07, 0x39, 0x7b, 0x90, 0xcd, 0x24, 0x6b, 0xd0, 0xb2, 0xe7, 0xde, 0xde,
	0xcc, 0xfb, 0xbc, 0x1f, 0xdf, 0x2f, 0x0f, 0x1f, 0x85, 0x1c, 0x12, 0x09, 0x14, 0xb2, 0x00, 0x52,
	0x36, 0xe3, 0x40, 0xf3, 0x09, 0x4d, 0x40, 0xc0, 0x54, 0x2b, 0xce, 0x20, 0x53, 0x85, 0x9b, 0x2a,
	0xa9, 0xa5, 0xb5, 0x6f, 0x40, 0x77, 0x0d, 0xba, 0xf9, 0xe4, 0xf0, 0x40, 0x48, 0x21, 0x9b, 0x3c,
	0x5d, 0x45, 0x06, 0x3d, 0x1c, 0xdd, 0xd8, 0x53, 0x86, 0x3c, 0x86, 0x8e, 0x98, 0xc9, 0x86, 0x60,
	0x99, 0x9e, 0x5f, 0xd2, 0xfc, 0x38, 0xe0, 0x9a, 0x1d, 0x9b, 0x97, 0x21, 0xc8, 0xf7, 0x2d, 0xfc,
	0xe4, 0x1c, 0xc4, 0xa9, 0x62, 0x0b, 0x7d, 0xd1, 0x6e, 0xf2, 0x2e, 0xd3, 0x73, 0xa9, 0xa2, 0x4b,
	0xa6, 0x23, 0xb9, 0xb0, 0x3e, 0xe0, 0xdd, 0x6e, 0xc0, 0x34, 0x0a, 0x87, 0x68, 0x84, 0xc6, 0xdb,
	0xde, 0xd3, 0xaa, 0x74, 0xf0, 0xa7, 0xf6, 0xfb, 0xec, 0x7d, 0x5d, 0x3a, 0x56, 0xc1, 0x92, 0xf8,
	0x2d, 0xe9, 0xa1, 0xc4, 0xc7, 0xdd, 0xeb, 0x2c, 0xb4, 0x5e, 0xe2, 0xfb, 0x62, 0x35, 0x84, 0xab,
	0xe1, 0xd6, 0x08, 0x8d, 0x77, 0x3c, 0xab, 0x2e, 0x9d, 0x07, 0xa6, 0xa8, 0x4d, 0x10, 0xbf, 0x43,
	0xfe, 0xd2, 0x7c, 0x78, 0xe7, 0x66, 0x9a, 0xaf, 0x69, 0x6e, 0x9d, 0xe2, 0xbb, 0x4d, 0x38, 0xdc,
	0x1e, 0xa1, 0xf1, 0xee, 0xc9, 0x63, 0xd7, 0xc8, 0x76, 0x8d, 0xd0, 0x56, 0xb6, 0xdb, 0x68, 0xf4,
	0x0e, 0xae, 0x4a, 0x67, 0x50, 0x97, 0xce, 0x5e, 0xaf, 0x19, 0xf1, 0x4d, 0x3d, 0x39, 0xc2, 0xcf,
	0x36, 0x9a, 0xe1, 0x73, 0x48, 0xe5, 0x02, 0x38, 0xf9, 0x8d, 0xb0, 0x7d, 0x0e, 0xc2, 0xe7, 0xb9,
	0xfc, 0xca, 0x6f, 0x91, 0x6f, 0x2f, 0xfe, 0xf5, 0xed, 0xe1, 0x7f, 0x52, 0xd7, 0xae, 0xbd, 0xc1,
	0x7b, 0x09, 0x88, 0xa9, 0x2e, 0x52, 0x3e, 0xcd, 0x54, 0xdc, 0x98, 0xb7, 0xe3, 0x3d, 0xaa, 0x4b,
	0x67, 0xdf, 0x14, 0xf4, 0xb3, 0xc4, 0xc7, 0x09, 0x88, 0x8b, 0x22, 0xe5, 0x9f, 0x55, 0x4c, 0xc6,
	0xf8, 0xf9, 0x66, 0xf5, 0x9d, 0x51, 0xde, 0xc7, 0xab, 0xca, 0x46, 0xd7, 0x95, 0x8d, 0x7e, 0x55,
	0x36, 0xfa, 0xb1, 0xb4, 0x07, 0xd7, 0x4b, 0x7b, 0xf0, 0x73, 0x69, 0x0f, 0xbe, 0x9c, 0x88, 0x48,
	0xcf, 0xb3, 0xc0, 0x9d, 0xc9, 0x84, 0x9a, 0x43, 0x7e, 0x15, 0xb3, 0x00, 0xda, 0x98, 0xe6, 0xaf,
	0xe9, 0xb7, 0xde, 0x65, 0xaf, 0xf6, 0x80, 0xe0, 0x5e, 0x73, 0xb4, 0x93, 0x3f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x13, 0xf7, 0x89, 0xc5, 0x4e, 0x03, 0x00, 0x00,
}

func (m *MsgGrantTreasuryAuthorization) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgGrantTreasuryAuthorization) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgGrantTreasuryAuthorization) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Grant.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMsgsTreasury(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Grantee) > 0 {
		i -= len(m.Grantee)
		copy(dAtA[i:], m.Grantee)
		i = encodeVarintMsgsTreasury(dAtA, i, uint64(len(m.Grantee)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Granter) > 0 {
		i -= len(m.Granter)
		copy(dAtA[i:], m.Granter)
		i = encodeVarintMsgsTreasury(dAtA, i, uint64(len(m.Granter)))
		i--
		dAtA[i] = 0x12
	}
	if m.SubspaceID != 0 {
		i = encodeVarintMsgsTreasury(dAtA, i, uint64(m.SubspaceID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgGrantTreasuryAuthorizationResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgGrantTreasuryAuthorizationResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgGrantTreasuryAuthorizationResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgRevokeTreasuryAuthorization) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRevokeTreasuryAuthorization) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRevokeTreasuryAuthorization) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MsgTypeUrl) > 0 {
		i -= len(m.MsgTypeUrl)
		copy(dAtA[i:], m.MsgTypeUrl)
		i = encodeVarintMsgsTreasury(dAtA, i, uint64(len(m.MsgTypeUrl)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Grantee) > 0 {
		i -= len(m.Grantee)
		copy(dAtA[i:], m.Grantee)
		i = encodeVarintMsgsTreasury(dAtA, i, uint64(len(m.Grantee)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Granter) > 0 {
		i -= len(m.Granter)
		copy(dAtA[i:], m.Granter)
		i = encodeVarintMsgsTreasury(dAtA, i, uint64(len(m.Granter)))
		i--
		dAtA[i] = 0x12
	}
	if m.SubspaceID != 0 {
		i = encodeVarintMsgsTreasury(dAtA, i, uint64(m.SubspaceID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgRevokeTreasuryAuthorizationResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRevokeTreasuryAuthorizationResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRevokeTreasuryAuthorizationResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintMsgsTreasury(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsgsTreasury(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgGrantTreasuryAuthorization) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SubspaceID != 0 {
		n += 1 + sovMsgsTreasury(uint64(m.SubspaceID))
	}
	l = len(m.Granter)
	if l > 0 {
		n += 1 + l + sovMsgsTreasury(uint64(l))
	}
	l = len(m.Grantee)
	if l > 0 {
		n += 1 + l + sovMsgsTreasury(uint64(l))
	}
	l = m.Grant.Size()
	n += 1 + l + sovMsgsTreasury(uint64(l))
	return n
}

func (m *MsgGrantTreasuryAuthorizationResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgRevokeTreasuryAuthorization) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SubspaceID != 0 {
		n += 1 + sovMsgsTreasury(uint64(m.SubspaceID))
	}
	l = len(m.Granter)
	if l > 0 {
		n += 1 + l + sovMsgsTreasury(uint64(l))
	}
	l = len(m.Grantee)
	if l > 0 {
		n += 1 + l + sovMsgsTreasury(uint64(l))
	}
	l = len(m.MsgTypeUrl)
	if l > 0 {
		n += 1 + l + sovMsgsTreasury(uint64(l))
	}
	return n
}

func (m *MsgRevokeTreasuryAuthorizationResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovMsgsTreasury(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsgsTreasury(x uint64) (n int) {
	return sovMsgsTreasury(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgGrantTreasuryAuthorization) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgsTreasury
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgGrantTreasuryAuthorization: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgGrantTreasuryAuthorization: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubspaceID", wireType)
			}
			m.SubspaceID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsTreasury
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SubspaceID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Granter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsTreasury
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsgsTreasury
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsTreasury
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Granter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Grantee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsTreasury
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsgsTreasury
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsTreasury
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Grantee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Grant", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsTreasury
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMsgsTreasury
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsTreasury
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Grant.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgsTreasury(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgsTreasury
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgGrantTreasuryAuthorizationResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgsTreasury
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgGrantTreasuryAuthorizationResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgGrantTreasuryAuthorizationResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsgsTreasury(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgsTreasury
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgRevokeTreasuryAuthorization) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgsTreasury
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgRevokeTreasuryAuthorization: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRevokeTreasuryAuthorization: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubspaceID", wireType)
			}
			m.SubspaceID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsTreasury
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SubspaceID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Granter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsTreasury
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsgsTreasury
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsTreasury
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Granter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Grantee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsTreasury
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsgsTreasury
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsTreasury
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Grantee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgTypeUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsTreasury
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsgsTreasury
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsTreasury
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MsgTypeUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgsTreasury(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgsTreasury
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgRevokeTreasuryAuthorizationResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgsTreasury
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgRevokeTreasuryAuthorizationResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRevokeTreasuryAuthorizationResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsgsTreasury(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgsTreasury
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMsgsTreasury(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsgsTreasury
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMsgsTreasury
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMsgsTreasury
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthMsgsTreasury
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsgsTreasury
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsgsTreasury
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsgsTreasury        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsgsTreasury          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsgsTreasury = fmt.Errorf("proto: unexpected end of group")
)