// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/reactions/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

// GenesisState contains the data of the genesis state for the reactions module
type GenesisState struct {
	SubspacesData       []SubspaceDataEntry       `protobuf:"bytes,1,rep,name=subspaces_data,json=subspacesData,proto3" json:"subspaces_data"`
	RegisteredReactions []RegisteredReaction      `protobuf:"bytes,2,rep,name=registered_reactions,json=registeredReactions,proto3" json:"registered_reactions"`
	PostsData           []PostDataEntry           `protobuf:"bytes,3,rep,name=posts_data,json=postsData,proto3" json:"posts_data"`
	Reactions           []Reaction                `protobuf:"bytes,4,rep,name=reactions,proto3" json:"reactions"`
	SubspacesParams     []SubspaceReactionsParams `protobuf:"bytes,5,rep,name=subspaces_params,json=subspacesParams,proto3" json:"subspaces_params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_81b04d309ae6ac6f, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetSubspacesData() []SubspaceDataEntry {
	if m != nil {
		return m.SubspacesData
	}
	return nil
}

func (m *GenesisState) GetRegisteredReactions() []RegisteredReaction {
	if m != nil {
		return m.RegisteredReactions
	}
	return nil
}

func (m *GenesisState) GetPostsData() []PostDataEntry {
	if m != nil {
		return m.PostsData
	}
	return nil
}

func (m *GenesisState) GetReactions() []Reaction {
	if m != nil {
		return m.Reactions
	}
	return nil
}

func (m *GenesisState) GetSubspacesParams() []SubspaceReactionsParams {
	if m != nil {
		return m.SubspacesParams
	}
	return nil
}

// SubspaceDataEntry contains the data related to a single subspace
type SubspaceDataEntry struct {
	SubspaceID           uint64 `protobuf:"varint,1,opt,name=subspace_id,json=subspaceId,proto3" json:"subspace_id,omitempty"`
	RegisteredReactionID uint32 `protobuf:"varint,2,opt,name=registered_reaction_id,json=registeredReactionId,proto3" json:"registered_reaction_id,omitempty"`
}

func (m *SubspaceDataEntry) Reset()         { *m = SubspaceDataEntry{} }
func (m *SubspaceDataEntry) String() string { return proto.CompactTextString(m) }
func (*SubspaceDataEntry) ProtoMessage()    {}
func (*SubspaceDataEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_81b04d309ae6ac6f, []int{1}
}
func (m *SubspaceDataEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SubspaceDataEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SubspaceDataEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SubspaceDataEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubspaceDataEntry.Merge(m, src)
}
func (m *SubspaceDataEntry) XXX_Size() int {
	return m.Size()
}
func (m *SubspaceDataEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_SubspaceDataEntry.DiscardUnknown(m)
}

var xxx_messageInfo_SubspaceDataEntry proto.InternalMessageInfo

func (m *SubspaceDataEntry) GetSubspaceID() uint64 {
	if m != nil {
		return m.SubspaceID
	}
	return 0
}

func (m *SubspaceDataEntry) GetRegisteredReactionID() uint32 {
	if m != nil {
		return m.RegisteredReactionID
	}
	return 0
}

// PostDataEntry contains the data related to a single post
type PostDataEntry struct {
	SubspaceID uint64 `protobuf:"varint,1,opt,name=subspace_id,json=subspaceId,proto3" json:"subspace_id,omitempty"`
	PostID     uint64 `protobuf:"varint,2,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	ReactionID uint32 `protobuf:"varint,3,opt,name=reaction_id,json=reactionId,proto3" json:"reaction_id,omitempty"`
}

func (m *PostDataEntry) Reset()         { *m = PostDataEntry{} }
func (m *PostDataEntry) String() string { return proto.CompactTextString(m) }
func (*PostDataEntry) ProtoMessage()    {}
func (*PostDataEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_81b04d309ae6ac6f, []int{2}
}
func (m *PostDataEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PostDataEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PostDataEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PostDataEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostDataEntry.Merge(m, src)
}
func (m *PostDataEntry) XXX_Size() int {
	return m.Size()
}
func (m *PostDataEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_PostDataEntry.DiscardUnknown(m)
}

var xxx_messageInfo_PostDataEntry proto.InternalMessageInfo

func (m *PostDataEntry) GetSubspaceID() uint64 {
	if m != nil {
		return m.SubspaceID
	}
	return 0
}

func (m *PostDataEntry) GetPostID() uint64 {
	if m != nil {
		return m.PostID
	}
	return 0
}

func (m *PostDataEntry) GetReactionID() uint32 {
	if m != nil {
		return m.ReactionID
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "desmos.reactions.v1.GenesisState")
	proto.RegisterType((*SubspaceDataEntry)(nil), "desmos.reactions.v1.SubspaceDataEntry")
	proto.RegisterType((*PostDataEntry)(nil), "desmos.reactions.v1.PostDataEntry")
}

func init() { proto.RegisterFile("desmos/reactions/v1/genesis.proto", fileDescriptor_81b04d309ae6ac6f) }

var fileDescriptor_81b04d309ae6ac6f = []byte{
	// 500 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xc1, 0x8a, 0xd3, 0x40,
	0x18, 0xc7, 0x3b, 0xdb, 0x5a, 0xdd, 0x6f, 0xed, 0xaa, 0xd9, 0x22, 0x71, 0xc1, 0xa4, 0x56, 0xd0,
	0x1e, 0x34, 0xc3, 0xae, 0x9e, 0xbc, 0x19, 0x2a, 0x4b, 0x40, 0x64, 0x49, 0x6f, 0x82, 0xd4, 0x49,
	0x33, 0xc6, 0x40, 0xd3, 0x09, 0x99, 0x69, 0xb1, 0x6f, 0xe1, 0xd1, 0x93, 0xf4, 0x39, 0x7c, 0x82,
	0x3d, 0xee, 0xd1, 0x53, 0x91, 0xf4, 0xe2, 0x63, 0x48, 0x26, 0x93, 0xb4, 0xbb, 0x89, 0x82, 0xb7,
	0xcc, 0x7c, 0xdf, 0xff, 0x97, 0xdf, 0xcc, 0xc7, 0xc0, 0x23, 0x9f, 0xf2, 0x88, 0x71, 0x9c, 0x50,
	0x32, 0x11, 0x21, 0x9b, 0x71, 0xbc, 0x38, 0xc1, 0x01, 0x9d, 0x51, 0x1e, 0x72, 0x2b, 0x4e, 0x98,
	0x60, 0xda, 0x51, 0xde, 0x62, 0x95, 0x2d, 0xd6, 0xe2, 0xe4, 0xb8, 0x1b, 0xb0, 0x80, 0xc9, 0x3a,
	0xce, 0xbe, 0xf2, 0xd6, 0xe3, 0x07, 0x01, 0x63, 0xc1, 0x94, 0x62, 0xb9, 0xf2, 0xe6, 0x9f, 0x30,
	0x99, 0x2d, 0x55, 0xc9, 0xbc, 0x5e, 0x12, 0x61, 0x44, 0xb9, 0x20, 0x51, 0x5c, 0x64, 0x27, 0x2c,
	0xfb, 0xcd, 0x38, 0x87, 0xe6, 0x0b, 0x55, 0xea, 0xd5, 0x49, 0x46, 0xcc, 0xa7, 0x53, 0xd5, 0xd1,
	0xff, 0xd1, 0x84, 0xdb, 0x67, 0xb9, 0xf5, 0x48, 0x10, 0x41, 0xb5, 0x11, 0x1c, 0xf2, 0xb9, 0xc7,
	0x63, 0x32, 0xa1, 0x7c, 0xec, 0x13, 0x41, 0x74, 0xd4, 0x6b, 0x0e, 0x0e, 0x4e, 0x9f, 0x58, 0x35,
	0xa7, 0xb1, 0x46, 0xaa, 0x75, 0x48, 0x04, 0x79, 0x33, 0x13, 0xc9, 0xd2, 0x6e, 0x5d, 0xac, 0xcd,
	0x86, 0xdb, 0x29, 0x19, 0x59, 0x45, 0xfb, 0x08, 0xdd, 0x84, 0x06, 0x21, 0x17, 0x34, 0xa1, 0xfe,
	0xb8, 0x24, 0xe8, 0x7b, 0x12, 0xfd, 0xb4, 0x16, 0xed, 0x96, 0x01, 0x57, 0x6d, 0x2b, 0xf6, 0x51,
	0x52, 0xa9, 0x70, 0xed, 0x0c, 0x20, 0x66, 0x5c, 0x28, 0xe5, 0xa6, 0xe4, 0xf6, 0x6b, 0xb9, 0xe7,
	0x8c, 0x8b, 0xeb, 0xba, 0xfb, 0x32, 0x2b, 0x55, 0x5f, 0xc3, 0xfe, 0xd6, 0xaf, 0x25, 0x39, 0x0f,
	0xff, 0xe2, 0x77, 0xc5, 0x6a, 0x9b, 0xd2, 0x3e, 0xc0, 0xdd, 0xed, 0x15, 0xc6, 0x24, 0x21, 0x11,
	0xd7, 0x6f, 0x48, 0xd2, 0xb3, 0x7f, 0x5e, 0x62, 0x79, 0x9a, 0x73, 0x99, 0x51, 0xe0, 0x3b, 0x25,
	0x2b, 0xdf, 0x7e, 0x75, 0xeb, 0xdb, 0xca, 0x44, 0xbf, 0x57, 0x26, 0xea, 0x7f, 0x47, 0x70, 0xaf,
	0x32, 0x01, 0x0d, 0xc3, 0x41, 0x11, 0x19, 0x87, 0xbe, 0x8e, 0x7a, 0x68, 0xd0, 0xb2, 0x0f, 0xd3,
	0xb5, 0x09, 0x45, 0xaf, 0x33, 0x74, 0xa1, 0x68, 0x71, 0x7c, 0xed, 0x1d, 0xdc, 0xaf, 0x99, 0x4e,
	0x96, 0xdd, 0xeb, 0xa1, 0x41, 0xc7, 0xd6, 0xd3, 0xb5, 0xd9, 0xad, 0x8e, 0xc3, 0x19, 0xba, 0xdd,
	0xea, 0x28, 0x1c, 0x7f, 0x47, 0x70, 0x85, 0xa0, 0x73, 0xe5, 0xbe, 0xff, 0x5f, 0xee, 0x31, 0xdc,
	0xcc, 0x86, 0x53, 0xd8, 0xb4, 0x6c, 0x48, 0xd7, 0x66, 0x3b, 0x83, 0x3a, 0x43, 0xb7, 0x9d, 0x95,
	0x1c, 0x3f, 0xa3, 0xee, 0x6a, 0x37, 0xa5, 0xb6, 0xa4, 0xee, 0xc8, 0x42, 0x52, 0xa3, 0x68, 0xbf,
	0xbd, 0x48, 0x0d, 0x74, 0x99, 0x1a, 0xe8, 0x57, 0x6a, 0xa0, 0xaf, 0x1b, 0xa3, 0x71, 0xb9, 0x31,
	0x1a, 0x3f, 0x37, 0x46, 0xe3, 0xfd, 0x69, 0x10, 0x8a, 0xcf, 0x73, 0xcf, 0x9a, 0xb0, 0x08, 0xe7,
	0x63, 0x7b, 0x3e, 0x25, 0x1e, 0x57, 0xdf, 0x78, 0xf1, 0x12, 0x7f, 0xd9, 0x79, 0x58, 0x62, 0x19,
	0x53, 0xee, 0xb5, 0xe5, 0xab, 0x7a, 0xf1, 0x27, 0x00, 0x00, 0xff, 0xff, 0x26, 0xc0, 0xec, 0x81,
	0x1e, 0x04, 0x00, 0x00,
}

func (this *GenesisState) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GenesisState)
	if !ok {
		that2, ok := that.(GenesisState)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.SubspacesData) != len(that1.SubspacesData) {
		return false
	}
	for i := range this.SubspacesData {
		if !this.SubspacesData[i].Equal(&that1.SubspacesData[i]) {
			return false
		}
	}
	if len(this.RegisteredReactions) != len(that1.RegisteredReactions) {
		return false
	}
	for i := range this.RegisteredReactions {
		if !this.RegisteredReactions[i].Equal(&that1.RegisteredReactions[i]) {
			return false
		}
	}
	if len(this.PostsData) != len(that1.PostsData) {
		return false
	}
	for i := range this.PostsData {
		if !this.PostsData[i].Equal(&that1.PostsData[i]) {
			return false
		}
	}
	if len(this.Reactions) != len(that1.Reactions) {
		return false
	}
	for i := range this.Reactions {
		if !this.Reactions[i].Equal(&that1.Reactions[i]) {
			return false
		}
	}
	if len(this.SubspacesParams) != len(that1.SubspacesParams) {
		return false
	}
	for i := range this.SubspacesParams {
		if !this.SubspacesParams[i].Equal(&that1.SubspacesParams[i]) {
			return false
		}
	}
	return true
}
func (this *SubspaceDataEntry) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SubspaceDataEntry)
	if !ok {
		that2, ok := that.(SubspaceDataEntry)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.SubspaceID != that1.SubspaceID {
		return false
	}
	if this.RegisteredReactionID != that1.RegisteredReactionID {
		return false
	}
	return true
}
func (this *PostDataEntry) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PostDataEntry)
	if !ok {
		that2, ok := that.(PostDataEntry)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.SubspaceID != that1.SubspaceID {
		return false
	}
	if this.PostID != that1.PostID {
		return false
	}
	if this.ReactionID != that1.ReactionID {
		return false
	}
	return true
}
func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SubspacesParams) > 0 {
		for iNdEx := len(m.SubspacesParams) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SubspacesParams[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Reactions) > 0 {
		for iNdEx := len(m.Reactions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Reactions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.PostsData) > 0 {
		for iNdEx := len(m.PostsData) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PostsData[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.RegisteredReactions) > 0 {
		for iNdEx := len(m.RegisteredReactions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RegisteredReactions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.SubspacesData) > 0 {
		for iNdEx := len(m.SubspacesData) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SubspacesData[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *SubspaceDataEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SubspaceDataEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SubspaceDataEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RegisteredReactionID != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.RegisteredReactionID))
		i--
		dAtA[i] = 0x10
	}
	if m.SubspaceID != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.SubspaceID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *PostDataEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PostDataEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PostDataEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ReactionID != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.ReactionID))
		i--
		dAtA[i] = 0x18
	}
	if m.PostID != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.PostID))
		i--
		dAtA[i] = 0x10
	}
	if m.SubspaceID != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.SubspaceID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.SubspacesData) > 0 {
		for _, e := range m.SubspacesData {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.RegisteredReactions) > 0 {
		for _, e := range m.RegisteredReactions {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.PostsData) > 0 {
		for _, e := range m.PostsData {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Reactions) > 0 {
		for _, e := range m.Reactions {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.SubspacesParams) > 0 {
		for _, e := range m.SubspacesParams {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *SubspaceDataEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SubspaceID != 0 {
		n += 1 + sovGenesis(uint64(m.SubspaceID))
	}
	if m.RegisteredReactionID != 0 {
		n += 1 + sovGenesis(uint64(m.RegisteredReactionID))
	}
	return n
}

func (m *PostDataEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SubspaceID != 0 {
		n += 1 + sovGenesis(uint64(m.SubspaceID))
	}
	if m.PostID != 0 {
		n += 1 + sovGenesis(uint64(m.PostID))
	}
	if m.ReactionID != 0 {
		n += 1 + sovGenesis(uint64(m.ReactionID))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubspacesData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubspacesData = append(m.SubspacesData, SubspaceDataEntry{})
			if err := m.SubspacesData[len(m.SubspacesData)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegisteredReactions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RegisteredReactions = append(m.RegisteredReactions, RegisteredReaction{})
			if err := m.RegisteredReactions[len(m.RegisteredReactions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PostsData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PostsData = append(m.PostsData, PostDataEntry{})
			if err := m.PostsData[len(m.PostsData)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reactions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reactions = append(m.Reactions, Reaction{})
			if err := m.Reactions[len(m.Reactions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubspacesParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubspacesParams = append(m.SubspacesParams, SubspaceReactionsParams{})
			if err := m.SubspacesParams[len(m.SubspacesParams)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *SubspaceDataEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: SubspaceDataEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SubspaceDataEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubspaceID", wireType)
			}
			m.SubspaceID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegisteredReactionID", wireType)
			}
			m.RegisteredReactionID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RegisteredReactionID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *PostDataEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: PostDataEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PostDataEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubspaceID", wireType)
			}
			m.SubspaceID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PostID", wireType)
			}
			m.PostID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PostID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReactionID", wireType)
			}
			m.ReactionID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReactionID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)