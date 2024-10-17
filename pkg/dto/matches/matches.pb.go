// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.1
// source: matches.proto

package matches

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	accounts "locgame-mini-server/pkg/dto/accounts"
	base "locgame-mini-server/pkg/dto/base"
	game "locgame-mini-server/pkg/dto/game"
	resources "locgame-mini-server/pkg/dto/resources"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type QuickMatchDeclineReason int32

const (
	QuickMatchDeclineReason_Unknown       QuickMatchDeclineReason = 0
	QuickMatchDeclineReason_LeftTheGame   QuickMatchDeclineReason = 1
	QuickMatchDeclineReason_NotInterested QuickMatchDeclineReason = 2
)

// Enum value maps for QuickMatchDeclineReason.
var (
	QuickMatchDeclineReason_name = map[int32]string{
		0: "Unknown",
		1: "LeftTheGame",
		2: "NotInterested",
	}
	QuickMatchDeclineReason_value = map[string]int32{
		"Unknown":       0,
		"LeftTheGame":   1,
		"NotInterested": 2,
	}
)

func (x QuickMatchDeclineReason) Enum() *QuickMatchDeclineReason {
	p := new(QuickMatchDeclineReason)
	*p = x
	return p
}

func (x QuickMatchDeclineReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QuickMatchDeclineReason) Descriptor() protoreflect.EnumDescriptor {
	return file_matches_proto_enumTypes[0].Descriptor()
}

func (QuickMatchDeclineReason) Type() protoreflect.EnumType {
	return &file_matches_proto_enumTypes[0]
}

func (x QuickMatchDeclineReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use QuickMatchDeclineReason.Descriptor instead.
func (QuickMatchDeclineReason) EnumDescriptor() ([]byte, []int) {
	return file_matches_proto_rawDescGZIP(), []int{0}
}

type QuickMatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameType game.GameType `protobuf:"varint,1,opt,name=GameType,proto3,enum=game.GameType" json:"GameType,omitempty" yaml:"GameType,omitempty" bson:"game_type,omitempty"`
	Stake    int32         `protobuf:"varint,2,opt,name=Stake,proto3" json:"Stake,omitempty" yaml:"Stake,omitempty" bson:"stake,omitempty"`
}

func (x *QuickMatchRequest) Reset() {
	*x = QuickMatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_matches_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuickMatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuickMatchRequest) ProtoMessage() {}

func (x *QuickMatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_matches_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuickMatchRequest.ProtoReflect.Descriptor instead.
func (*QuickMatchRequest) Descriptor() ([]byte, []int) {
	return file_matches_proto_rawDescGZIP(), []int{0}
}

func (x *QuickMatchRequest) GetGameType() game.GameType {
	if x != nil {
		return x.GameType
	}
	return game.GameType(0)
}

func (x *QuickMatchRequest) GetStake() int32 {
	if x != nil {
		return x.Stake
	}
	return 0
}

type Stakes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LC []string `protobuf:"bytes,1,rep,name=LC,proto3" json:"LC,omitempty" yaml:"LC,omitempty" bson:"lc,omitempty"`
}

func (x *Stakes) Reset() {
	*x = Stakes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_matches_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stakes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stakes) ProtoMessage() {}

func (x *Stakes) ProtoReflect() protoreflect.Message {
	mi := &file_matches_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stakes.ProtoReflect.Descriptor instead.
func (*Stakes) Descriptor() ([]byte, []int) {
	return file_matches_proto_rawDescGZIP(), []int{1}
}

func (x *Stakes) GetLC() []string {
	if x != nil {
		return x.LC
	}
	return nil
}

type QuickMatchList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Matches map[string]*Stakes `protobuf:"bytes,1,rep,name=Matches,proto3" json:"Matches,omitempty" yaml:"Matches,omitempty" bson:"matches,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *QuickMatchList) Reset() {
	*x = QuickMatchList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_matches_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuickMatchList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuickMatchList) ProtoMessage() {}

func (x *QuickMatchList) ProtoReflect() protoreflect.Message {
	mi := &file_matches_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuickMatchList.ProtoReflect.Descriptor instead.
func (*QuickMatchList) Descriptor() ([]byte, []int) {
	return file_matches_proto_rawDescGZIP(), []int{2}
}

func (x *QuickMatchList) GetMatches() map[string]*Stakes {
	if x != nil {
		return x.Matches
	}
	return nil
}

type QuickMatchDecline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reason QuickMatchDeclineReason `protobuf:"varint,1,opt,name=Reason,proto3,enum=matches.QuickMatchDeclineReason" json:"Reason,omitempty" yaml:"Reason,omitempty" bson:"reason,omitempty"`
}

func (x *QuickMatchDecline) Reset() {
	*x = QuickMatchDecline{}
	if protoimpl.UnsafeEnabled {
		mi := &file_matches_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuickMatchDecline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuickMatchDecline) ProtoMessage() {}

func (x *QuickMatchDecline) ProtoReflect() protoreflect.Message {
	mi := &file_matches_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuickMatchDecline.ProtoReflect.Descriptor instead.
func (*QuickMatchDecline) Descriptor() ([]byte, []int) {
	return file_matches_proto_rawDescGZIP(), []int{3}
}

func (x *QuickMatchDecline) GetReason() QuickMatchDeclineReason {
	if x != nil {
		return x.Reason
	}
	return QuickMatchDeclineReason_Unknown
}

type QuickMatchResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Adjustments []*resources.ResourceAdjustment `protobuf:"bytes,1,rep,name=Adjustments,proto3" json:"Adjustments,omitempty" yaml:"Adjustments,omitempty" bson:"adjustments,omitempty"`
}

func (x *QuickMatchResult) Reset() {
	*x = QuickMatchResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_matches_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuickMatchResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuickMatchResult) ProtoMessage() {}

func (x *QuickMatchResult) ProtoReflect() protoreflect.Message {
	mi := &file_matches_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuickMatchResult.ProtoReflect.Descriptor instead.
func (*QuickMatchResult) Descriptor() ([]byte, []int) {
	return file_matches_proto_rawDescGZIP(), []int{4}
}

func (x *QuickMatchResult) GetAdjustments() []*resources.ResourceAdjustment {
	if x != nil {
		return x.Adjustments
	}
	return nil
}

type QuickMatchCancel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *QuickMatchCancel) Reset() {
	*x = QuickMatchCancel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_matches_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuickMatchCancel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuickMatchCancel) ProtoMessage() {}

func (x *QuickMatchCancel) ProtoReflect() protoreflect.Message {
	mi := &file_matches_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuickMatchCancel.ProtoReflect.Descriptor instead.
func (*QuickMatchCancel) Descriptor() ([]byte, []int) {
	return file_matches_proto_rawDescGZIP(), []int{5}
}

type QuickMatchEnded struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Adjustments []*resources.ResourceAdjustment `protobuf:"bytes,1,rep,name=Adjustments,proto3" json:"Adjustments,omitempty" yaml:"Adjustments,omitempty" bson:"adjustments,omitempty"`
}

func (x *QuickMatchEnded) Reset() {
	*x = QuickMatchEnded{}
	if protoimpl.UnsafeEnabled {
		mi := &file_matches_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuickMatchEnded) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuickMatchEnded) ProtoMessage() {}

func (x *QuickMatchEnded) ProtoReflect() protoreflect.Message {
	mi := &file_matches_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuickMatchEnded.ProtoReflect.Descriptor instead.
func (*QuickMatchEnded) Descriptor() ([]byte, []int) {
	return file_matches_proto_rawDescGZIP(), []int{6}
}

func (x *QuickMatchEnded) GetAdjustments() []*resources.ResourceAdjustment {
	if x != nil {
		return x.Adjustments
	}
	return nil
}

type QuickMatchAccept struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OpponentID *base.ObjectID `protobuf:"bytes,1,opt,name=OpponentID,proto3" json:"OpponentID,omitempty" yaml:"OpponentID,omitempty" bson:"opponent_id,omitempty"`
}

func (x *QuickMatchAccept) Reset() {
	*x = QuickMatchAccept{}
	if protoimpl.UnsafeEnabled {
		mi := &file_matches_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuickMatchAccept) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuickMatchAccept) ProtoMessage() {}

func (x *QuickMatchAccept) ProtoReflect() protoreflect.Message {
	mi := &file_matches_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuickMatchAccept.ProtoReflect.Descriptor instead.
func (*QuickMatchAccept) Descriptor() ([]byte, []int) {
	return file_matches_proto_rawDescGZIP(), []int{7}
}

func (x *QuickMatchAccept) GetOpponentID() *base.ObjectID {
	if x != nil {
		return x.OpponentID
	}
	return nil
}

type QuickMatchData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerID   *base.ObjectID `protobuf:"bytes,1,opt,name=PlayerID,proto3" json:"PlayerID,omitempty" yaml:"PlayerID,omitempty" bson:"player_id,omitempty"`
	OpponentID *base.ObjectID `protobuf:"bytes,2,opt,name=OpponentID,proto3" json:"OpponentID,omitempty" yaml:"OpponentID,omitempty" bson:"opponent_id,omitempty"`
	Stake      int32          `protobuf:"varint,3,opt,name=Stake,proto3" json:"Stake,omitempty" yaml:"Stake,omitempty" bson:"stake,omitempty"`
}

func (x *QuickMatchData) Reset() {
	*x = QuickMatchData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_matches_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuickMatchData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuickMatchData) ProtoMessage() {}

func (x *QuickMatchData) ProtoReflect() protoreflect.Message {
	mi := &file_matches_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuickMatchData.ProtoReflect.Descriptor instead.
func (*QuickMatchData) Descriptor() ([]byte, []int) {
	return file_matches_proto_rawDescGZIP(), []int{8}
}

func (x *QuickMatchData) GetPlayerID() *base.ObjectID {
	if x != nil {
		return x.PlayerID
	}
	return nil
}

func (x *QuickMatchData) GetOpponentID() *base.ObjectID {
	if x != nil {
		return x.OpponentID
	}
	return nil
}

func (x *QuickMatchData) GetStake() int32 {
	if x != nil {
		return x.Stake
	}
	return 0
}

type QuickMatchPlayersListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query string `protobuf:"bytes,1,opt,name=Query,proto3" json:"Query,omitempty" yaml:"Query,omitempty" bson:"query,omitempty"`
}

func (x *QuickMatchPlayersListRequest) Reset() {
	*x = QuickMatchPlayersListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_matches_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuickMatchPlayersListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuickMatchPlayersListRequest) ProtoMessage() {}

func (x *QuickMatchPlayersListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_matches_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuickMatchPlayersListRequest.ProtoReflect.Descriptor instead.
func (*QuickMatchPlayersListRequest) Descriptor() ([]byte, []int) {
	return file_matches_proto_rawDescGZIP(), []int{9}
}

func (x *QuickMatchPlayersListRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

type QuickMatchPlayersListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*accounts.UserInfo `protobuf:"bytes,1,rep,name=Users,proto3" json:"Users,omitempty" yaml:"Users,omitempty" bson:"users,omitempty"`
}

func (x *QuickMatchPlayersListResponse) Reset() {
	*x = QuickMatchPlayersListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_matches_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuickMatchPlayersListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuickMatchPlayersListResponse) ProtoMessage() {}

func (x *QuickMatchPlayersListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_matches_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuickMatchPlayersListResponse.ProtoReflect.Descriptor instead.
func (*QuickMatchPlayersListResponse) Descriptor() ([]byte, []int) {
	return file_matches_proto_rawDescGZIP(), []int{10}
}

func (x *QuickMatchPlayersListResponse) GetUsers() []*accounts.UserInfo {
	if x != nil {
		return x.Users
	}
	return nil
}

type QuickMatchConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stakes []int32 `protobuf:"varint,1,rep,packed,name=Stakes,proto3" json:"Stakes,omitempty" yaml:"Stakes,omitempty" bson:"stakes,omitempty"`
	Fee    float32 `protobuf:"fixed32,2,opt,name=Fee,proto3" json:"Fee,omitempty" yaml:"Fee,omitempty" bson:"fee,omitempty"`
}

func (x *QuickMatchConfig) Reset() {
	*x = QuickMatchConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_matches_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuickMatchConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuickMatchConfig) ProtoMessage() {}

func (x *QuickMatchConfig) ProtoReflect() protoreflect.Message {
	mi := &file_matches_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuickMatchConfig.ProtoReflect.Descriptor instead.
func (*QuickMatchConfig) Descriptor() ([]byte, []int) {
	return file_matches_proto_rawDescGZIP(), []int{11}
}

func (x *QuickMatchConfig) GetStakes() []int32 {
	if x != nil {
		return x.Stakes
	}
	return nil
}

func (x *QuickMatchConfig) GetFee() float32 {
	if x != nil {
		return x.Fee
	}
	return 0
}

var File_matches_proto protoreflect.FileDescriptor

var file_matches_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x1a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x55, 0x0a, 0x11, 0x51, 0x75, 0x69, 0x63, 0x6b, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x08, 0x47, 0x61, 0x6d, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e,
	0x47, 0x61, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x47, 0x61, 0x6d, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x22, 0x18, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x6b,
	0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x4c, 0x43, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x02,
	0x4c, 0x43, 0x22, 0x9d, 0x01, 0x0a, 0x0e, 0x51, 0x75, 0x69, 0x63, 0x6b, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x07, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73,
	0x2e, 0x51, 0x75, 0x69, 0x63, 0x6b, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x2e,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x73, 0x1a, 0x4b, 0x0a, 0x0c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73,
	0x2e, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x4d, 0x0a, 0x11, 0x51, 0x75, 0x69, 0x63, 0x6b, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x44, 0x65, 0x63, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x73, 0x2e, 0x51, 0x75, 0x69, 0x63, 0x6b, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x63, 0x6c,
	0x69, 0x6e, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x22, 0x53, 0x0a, 0x10, 0x51, 0x75, 0x69, 0x63, 0x6b, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x3f, 0x0a, 0x0b, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41,
	0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x41, 0x64, 0x6a, 0x75, 0x73,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x12, 0x0a, 0x10, 0x51, 0x75, 0x69, 0x63, 0x6b, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x22, 0x52, 0x0a, 0x0f, 0x51, 0x75,
	0x69, 0x63, 0x6b, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x45, 0x6e, 0x64, 0x65, 0x64, 0x12, 0x3f, 0x0a,
	0x0b, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x0b, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x42,
	0x0a, 0x10, 0x51, 0x75, 0x69, 0x63, 0x6b, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x41, 0x63, 0x63, 0x65,
	0x70, 0x74, 0x12, 0x2e, 0x0a, 0x0a, 0x4f, 0x70, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x52, 0x0a, 0x4f, 0x70, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x49, 0x44, 0x22, 0x82, 0x01, 0x0a, 0x0e, 0x51, 0x75, 0x69, 0x63, 0x6b, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x2e, 0x0a, 0x0a, 0x4f, 0x70, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x44, 0x52, 0x0a, 0x4f, 0x70, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x49,
	0x44, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x22, 0x34, 0x0a, 0x1c, 0x51, 0x75, 0x69, 0x63, 0x6b,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x22, 0x49, 0x0a,
	0x1d, 0x51, 0x75, 0x69, 0x63, 0x6b, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28,
	0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x22, 0x3c, 0x0a, 0x10, 0x51, 0x75, 0x69, 0x63,
	0x6b, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x6b, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x06, 0x53, 0x74,
	0x61, 0x6b, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x46, 0x65, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x03, 0x46, 0x65, 0x65, 0x2a, 0x4a, 0x0a, 0x17, 0x51, 0x75, 0x69, 0x63, 0x6b, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x63, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x0f,
	0x0a, 0x0b, 0x4c, 0x65, 0x66, 0x74, 0x54, 0x68, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x10, 0x01, 0x12,
	0x11, 0x0a, 0x0d, 0x4e, 0x6f, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x10, 0x02, 0x42, 0x38, 0x5a, 0x23, 0x6c, 0x6f, 0x63, 0x67, 0x61, 0x6d, 0x65, 0x2d, 0x6d, 0x69,
	0x6e, 0x69, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x64, 0x74,
	0x6f, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0xaa, 0x02, 0x10, 0x4c, 0x6f, 0x43, 0x2e,
	0x44, 0x54, 0x4f, 0x73, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_matches_proto_rawDescOnce sync.Once
	file_matches_proto_rawDescData = file_matches_proto_rawDesc
)

func file_matches_proto_rawDescGZIP() []byte {
	file_matches_proto_rawDescOnce.Do(func() {
		file_matches_proto_rawDescData = protoimpl.X.CompressGZIP(file_matches_proto_rawDescData)
	})
	return file_matches_proto_rawDescData
}

var file_matches_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_matches_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_matches_proto_goTypes = []any{
	(QuickMatchDeclineReason)(0),          // 0: matches.QuickMatchDeclineReason
	(*QuickMatchRequest)(nil),             // 1: matches.QuickMatchRequest
	(*Stakes)(nil),                        // 2: matches.stakes
	(*QuickMatchList)(nil),                // 3: matches.QuickMatchList
	(*QuickMatchDecline)(nil),             // 4: matches.QuickMatchDecline
	(*QuickMatchResult)(nil),              // 5: matches.QuickMatchResult
	(*QuickMatchCancel)(nil),              // 6: matches.QuickMatchCancel
	(*QuickMatchEnded)(nil),               // 7: matches.QuickMatchEnded
	(*QuickMatchAccept)(nil),              // 8: matches.QuickMatchAccept
	(*QuickMatchData)(nil),                // 9: matches.QuickMatchData
	(*QuickMatchPlayersListRequest)(nil),  // 10: matches.QuickMatchPlayersListRequest
	(*QuickMatchPlayersListResponse)(nil), // 11: matches.QuickMatchPlayersListResponse
	(*QuickMatchConfig)(nil),              // 12: matches.QuickMatchConfig
	nil,                                   // 13: matches.QuickMatchList.MatchesEntry
	(game.GameType)(0),                    // 14: game.GameType
	(*resources.ResourceAdjustment)(nil),  // 15: resources.ResourceAdjustment
	(*base.ObjectID)(nil),                 // 16: base.ObjectID
	(*accounts.UserInfo)(nil),             // 17: accounts.UserInfo
}
var file_matches_proto_depIdxs = []int32{
	14, // 0: matches.QuickMatchRequest.GameType:type_name -> game.GameType
	13, // 1: matches.QuickMatchList.Matches:type_name -> matches.QuickMatchList.MatchesEntry
	0,  // 2: matches.QuickMatchDecline.Reason:type_name -> matches.QuickMatchDeclineReason
	15, // 3: matches.QuickMatchResult.Adjustments:type_name -> resources.ResourceAdjustment
	15, // 4: matches.QuickMatchEnded.Adjustments:type_name -> resources.ResourceAdjustment
	16, // 5: matches.QuickMatchAccept.OpponentID:type_name -> base.ObjectID
	16, // 6: matches.QuickMatchData.PlayerID:type_name -> base.ObjectID
	16, // 7: matches.QuickMatchData.OpponentID:type_name -> base.ObjectID
	17, // 8: matches.QuickMatchPlayersListResponse.Users:type_name -> accounts.UserInfo
	2,  // 9: matches.QuickMatchList.MatchesEntry.value:type_name -> matches.stakes
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_matches_proto_init() }
func file_matches_proto_init() {
	if File_matches_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_matches_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*QuickMatchRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_matches_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Stakes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_matches_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*QuickMatchList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_matches_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*QuickMatchDecline); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_matches_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*QuickMatchResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_matches_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*QuickMatchCancel); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_matches_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*QuickMatchEnded); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_matches_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*QuickMatchAccept); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_matches_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*QuickMatchData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_matches_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*QuickMatchPlayersListRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_matches_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*QuickMatchPlayersListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_matches_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*QuickMatchConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_matches_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_matches_proto_goTypes,
		DependencyIndexes: file_matches_proto_depIdxs,
		EnumInfos:         file_matches_proto_enumTypes,
		MessageInfos:      file_matches_proto_msgTypes,
	}.Build()
	File_matches_proto = out.File
	file_matches_proto_rawDesc = nil
	file_matches_proto_goTypes = nil
	file_matches_proto_depIdxs = nil
}
