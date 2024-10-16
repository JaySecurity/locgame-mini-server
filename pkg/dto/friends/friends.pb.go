// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: friends.proto

package friends

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	accounts "locgame-mini-server/pkg/dto/accounts"
	base "locgame-mini-server/pkg/dto/base"
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

type FriendRequestType int32

const (
	FriendRequestType_RequestReceived FriendRequestType = 0
	FriendRequestType_RequestAccepted FriendRequestType = 1
	FriendRequestType_RequestDeclined FriendRequestType = 2
	FriendRequestType_RequestCanceled FriendRequestType = 3
	FriendRequestType_FriendDeleted   FriendRequestType = 4
)

// Enum value maps for FriendRequestType.
var (
	FriendRequestType_name = map[int32]string{
		0: "RequestReceived",
		1: "RequestAccepted",
		2: "RequestDeclined",
		3: "RequestCanceled",
		4: "FriendDeleted",
	}
	FriendRequestType_value = map[string]int32{
		"RequestReceived": 0,
		"RequestAccepted": 1,
		"RequestDeclined": 2,
		"RequestCanceled": 3,
		"FriendDeleted":   4,
	}
)

func (x FriendRequestType) Enum() *FriendRequestType {
	p := new(FriendRequestType)
	*p = x
	return p
}

func (x FriendRequestType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FriendRequestType) Descriptor() protoreflect.EnumDescriptor {
	return file_friends_proto_enumTypes[0].Descriptor()
}

func (FriendRequestType) Type() protoreflect.EnumType {
	return &file_friends_proto_enumTypes[0]
}

func (x FriendRequestType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FriendRequestType.Descriptor instead.
func (FriendRequestType) EnumDescriptor() ([]byte, []int) {
	return file_friends_proto_rawDescGZIP(), []int{0}
}

type FriendlyMatchDeclineReason int32

const (
	FriendlyMatchDeclineReason_Unknown     FriendlyMatchDeclineReason = 0
	FriendlyMatchDeclineReason_LeftTheGame FriendlyMatchDeclineReason = 1 // TODO
)

// Enum value maps for FriendlyMatchDeclineReason.
var (
	FriendlyMatchDeclineReason_name = map[int32]string{
		0: "Unknown",
		1: "LeftTheGame",
	}
	FriendlyMatchDeclineReason_value = map[string]int32{
		"Unknown":     0,
		"LeftTheGame": 1,
	}
)

func (x FriendlyMatchDeclineReason) Enum() *FriendlyMatchDeclineReason {
	p := new(FriendlyMatchDeclineReason)
	*p = x
	return p
}

func (x FriendlyMatchDeclineReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FriendlyMatchDeclineReason) Descriptor() protoreflect.EnumDescriptor {
	return file_friends_proto_enumTypes[1].Descriptor()
}

func (FriendlyMatchDeclineReason) Type() protoreflect.EnumType {
	return &file_friends_proto_enumTypes[1]
}

func (x FriendlyMatchDeclineReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FriendlyMatchDeclineReason.Descriptor instead.
func (FriendlyMatchDeclineReason) EnumDescriptor() ([]byte, []int) {
	return file_friends_proto_rawDescGZIP(), []int{1}
}

type FriendsData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Friends         []*accounts.UserInfo `protobuf:"bytes,1,rep,name=Friends,proto3" json:"Friends,omitempty" yaml:"Friends,omitempty" bson:"friends,omitempty"`
	IncomingInvites []*accounts.UserInfo `protobuf:"bytes,2,rep,name=IncomingInvites,proto3" json:"IncomingInvites,omitempty" yaml:"IncomingInvites,omitempty" bson:"incoming_invites,omitempty"`
	OutgoingInvites []*accounts.UserInfo `protobuf:"bytes,3,rep,name=OutgoingInvites,proto3" json:"OutgoingInvites,omitempty" yaml:"OutgoingInvites,omitempty" bson:"outgoing_invites,omitempty"`
}

func (x *FriendsData) Reset() {
	*x = FriendsData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friends_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendsData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendsData) ProtoMessage() {}

func (x *FriendsData) ProtoReflect() protoreflect.Message {
	mi := &file_friends_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendsData.ProtoReflect.Descriptor instead.
func (*FriendsData) Descriptor() ([]byte, []int) {
	return file_friends_proto_rawDescGZIP(), []int{0}
}

func (x *FriendsData) GetFriends() []*accounts.UserInfo {
	if x != nil {
		return x.Friends
	}
	return nil
}

func (x *FriendsData) GetIncomingInvites() []*accounts.UserInfo {
	if x != nil {
		return x.IncomingInvites
	}
	return nil
}

func (x *FriendsData) GetOutgoingInvites() []*accounts.UserInfo {
	if x != nil {
		return x.OutgoingInvites
	}
	return nil
}

type FriendChangeData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     FriendRequestType  `protobuf:"varint,1,opt,name=Type,proto3,enum=friends.FriendRequestType" json:"Type,omitempty" yaml:"Type,omitempty" bson:"type,omitempty"`
	PlayerID *base.ObjectID     `protobuf:"bytes,2,opt,name=PlayerID,proto3" json:"PlayerID,omitempty" yaml:"PlayerID,omitempty" bson:"player_id,omitempty"`
	UserInfo *accounts.UserInfo `protobuf:"bytes,3,opt,name=UserInfo,proto3" json:"UserInfo,omitempty" yaml:"UserInfo,omitempty" bson:"user_info,omitempty"`
}

func (x *FriendChangeData) Reset() {
	*x = FriendChangeData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friends_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendChangeData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendChangeData) ProtoMessage() {}

func (x *FriendChangeData) ProtoReflect() protoreflect.Message {
	mi := &file_friends_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendChangeData.ProtoReflect.Descriptor instead.
func (*FriendChangeData) Descriptor() ([]byte, []int) {
	return file_friends_proto_rawDescGZIP(), []int{1}
}

func (x *FriendChangeData) GetType() FriendRequestType {
	if x != nil {
		return x.Type
	}
	return FriendRequestType_RequestReceived
}

func (x *FriendChangeData) GetPlayerID() *base.ObjectID {
	if x != nil {
		return x.PlayerID
	}
	return nil
}

func (x *FriendChangeData) GetUserInfo() *accounts.UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

type FriendlyMatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OpponentID *base.ObjectID `protobuf:"bytes,1,opt,name=OpponentID,proto3" json:"OpponentID,omitempty" yaml:"OpponentID,omitempty" bson:"opponent_id,omitempty"`
	Stake      int32          `protobuf:"varint,2,opt,name=Stake,proto3" json:"Stake,omitempty" yaml:"Stake,omitempty" bson:"stake,omitempty"`
}

func (x *FriendlyMatchRequest) Reset() {
	*x = FriendlyMatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friends_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendlyMatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendlyMatchRequest) ProtoMessage() {}

func (x *FriendlyMatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_friends_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendlyMatchRequest.ProtoReflect.Descriptor instead.
func (*FriendlyMatchRequest) Descriptor() ([]byte, []int) {
	return file_friends_proto_rawDescGZIP(), []int{2}
}

func (x *FriendlyMatchRequest) GetOpponentID() *base.ObjectID {
	if x != nil {
		return x.OpponentID
	}
	return nil
}

func (x *FriendlyMatchRequest) GetStake() int32 {
	if x != nil {
		return x.Stake
	}
	return 0
}

type FriendlyMatchDecline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reason FriendlyMatchDeclineReason `protobuf:"varint,1,opt,name=Reason,proto3,enum=friends.FriendlyMatchDeclineReason" json:"Reason,omitempty" yaml:"Reason,omitempty" bson:"reason,omitempty"`
}

func (x *FriendlyMatchDecline) Reset() {
	*x = FriendlyMatchDecline{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friends_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendlyMatchDecline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendlyMatchDecline) ProtoMessage() {}

func (x *FriendlyMatchDecline) ProtoReflect() protoreflect.Message {
	mi := &file_friends_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendlyMatchDecline.ProtoReflect.Descriptor instead.
func (*FriendlyMatchDecline) Descriptor() ([]byte, []int) {
	return file_friends_proto_rawDescGZIP(), []int{3}
}

func (x *FriendlyMatchDecline) GetReason() FriendlyMatchDeclineReason {
	if x != nil {
		return x.Reason
	}
	return FriendlyMatchDeclineReason_Unknown
}

type FriendlyMatchResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Adjustments []*resources.ResourceAdjustment `protobuf:"bytes,1,rep,name=Adjustments,proto3" json:"Adjustments,omitempty" yaml:"Adjustments,omitempty" bson:"adjustments,omitempty"`
}

func (x *FriendlyMatchResult) Reset() {
	*x = FriendlyMatchResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friends_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendlyMatchResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendlyMatchResult) ProtoMessage() {}

func (x *FriendlyMatchResult) ProtoReflect() protoreflect.Message {
	mi := &file_friends_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendlyMatchResult.ProtoReflect.Descriptor instead.
func (*FriendlyMatchResult) Descriptor() ([]byte, []int) {
	return file_friends_proto_rawDescGZIP(), []int{4}
}

func (x *FriendlyMatchResult) GetAdjustments() []*resources.ResourceAdjustment {
	if x != nil {
		return x.Adjustments
	}
	return nil
}

type FriendlyMatchCancel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FriendlyMatchCancel) Reset() {
	*x = FriendlyMatchCancel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friends_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendlyMatchCancel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendlyMatchCancel) ProtoMessage() {}

func (x *FriendlyMatchCancel) ProtoReflect() protoreflect.Message {
	mi := &file_friends_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendlyMatchCancel.ProtoReflect.Descriptor instead.
func (*FriendlyMatchCancel) Descriptor() ([]byte, []int) {
	return file_friends_proto_rawDescGZIP(), []int{5}
}

type FriendlyMatchEnded struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Adjustments []*resources.ResourceAdjustment `protobuf:"bytes,1,rep,name=Adjustments,proto3" json:"Adjustments,omitempty" yaml:"Adjustments,omitempty" bson:"adjustments,omitempty"`
}

func (x *FriendlyMatchEnded) Reset() {
	*x = FriendlyMatchEnded{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friends_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendlyMatchEnded) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendlyMatchEnded) ProtoMessage() {}

func (x *FriendlyMatchEnded) ProtoReflect() protoreflect.Message {
	mi := &file_friends_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendlyMatchEnded.ProtoReflect.Descriptor instead.
func (*FriendlyMatchEnded) Descriptor() ([]byte, []int) {
	return file_friends_proto_rawDescGZIP(), []int{6}
}

func (x *FriendlyMatchEnded) GetAdjustments() []*resources.ResourceAdjustment {
	if x != nil {
		return x.Adjustments
	}
	return nil
}

type FriendlyMatchAccept struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FriendlyMatchAccept) Reset() {
	*x = FriendlyMatchAccept{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friends_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendlyMatchAccept) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendlyMatchAccept) ProtoMessage() {}

func (x *FriendlyMatchAccept) ProtoReflect() protoreflect.Message {
	mi := &file_friends_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendlyMatchAccept.ProtoReflect.Descriptor instead.
func (*FriendlyMatchAccept) Descriptor() ([]byte, []int) {
	return file_friends_proto_rawDescGZIP(), []int{7}
}

type FriendlyMatchData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerID   *base.ObjectID `protobuf:"bytes,1,opt,name=PlayerID,proto3" json:"PlayerID,omitempty" yaml:"PlayerID,omitempty" bson:"player_id,omitempty"`
	OpponentID *base.ObjectID `protobuf:"bytes,2,opt,name=OpponentID,proto3" json:"OpponentID,omitempty" yaml:"OpponentID,omitempty" bson:"opponent_id,omitempty"`
	Stake      int32          `protobuf:"varint,3,opt,name=Stake,proto3" json:"Stake,omitempty" yaml:"Stake,omitempty" bson:"stake,omitempty"`
}

func (x *FriendlyMatchData) Reset() {
	*x = FriendlyMatchData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friends_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendlyMatchData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendlyMatchData) ProtoMessage() {}

func (x *FriendlyMatchData) ProtoReflect() protoreflect.Message {
	mi := &file_friends_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendlyMatchData.ProtoReflect.Descriptor instead.
func (*FriendlyMatchData) Descriptor() ([]byte, []int) {
	return file_friends_proto_rawDescGZIP(), []int{8}
}

func (x *FriendlyMatchData) GetPlayerID() *base.ObjectID {
	if x != nil {
		return x.PlayerID
	}
	return nil
}

func (x *FriendlyMatchData) GetOpponentID() *base.ObjectID {
	if x != nil {
		return x.OpponentID
	}
	return nil
}

func (x *FriendlyMatchData) GetStake() int32 {
	if x != nil {
		return x.Stake
	}
	return 0
}

type FindRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query string `protobuf:"bytes,1,opt,name=Query,proto3" json:"Query,omitempty" yaml:"Query,omitempty" bson:"query,omitempty"`
}

func (x *FindRequest) Reset() {
	*x = FindRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friends_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindRequest) ProtoMessage() {}

func (x *FindRequest) ProtoReflect() protoreflect.Message {
	mi := &file_friends_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindRequest.ProtoReflect.Descriptor instead.
func (*FindRequest) Descriptor() ([]byte, []int) {
	return file_friends_proto_rawDescGZIP(), []int{9}
}

func (x *FindRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

type FindResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*accounts.UserInfo `protobuf:"bytes,1,rep,name=Users,proto3" json:"Users,omitempty" yaml:"Users,omitempty" bson:"users,omitempty"`
}

func (x *FindResponse) Reset() {
	*x = FindResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friends_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindResponse) ProtoMessage() {}

func (x *FindResponse) ProtoReflect() protoreflect.Message {
	mi := &file_friends_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindResponse.ProtoReflect.Descriptor instead.
func (*FindResponse) Descriptor() ([]byte, []int) {
	return file_friends_proto_rawDescGZIP(), []int{10}
}

func (x *FindResponse) GetUsers() []*accounts.UserInfo {
	if x != nil {
		return x.Users
	}
	return nil
}

type FriendStatusData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsOnline bool           `protobuf:"varint,1,opt,name=IsOnline,proto3" json:"IsOnline,omitempty" yaml:"IsOnline,omitempty" bson:"is_online"`
	FriendID *base.ObjectID `protobuf:"bytes,2,opt,name=FriendID,proto3" json:"FriendID,omitempty" yaml:"FriendID,omitempty" bson:"friend_id,omitempty"`
}

func (x *FriendStatusData) Reset() {
	*x = FriendStatusData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friends_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendStatusData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendStatusData) ProtoMessage() {}

func (x *FriendStatusData) ProtoReflect() protoreflect.Message {
	mi := &file_friends_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendStatusData.ProtoReflect.Descriptor instead.
func (*FriendStatusData) Descriptor() ([]byte, []int) {
	return file_friends_proto_rawDescGZIP(), []int{11}
}

func (x *FriendStatusData) GetIsOnline() bool {
	if x != nil {
		return x.IsOnline
	}
	return false
}

func (x *FriendStatusData) GetFriendID() *base.ObjectID {
	if x != nil {
		return x.FriendID
	}
	return nil
}

type FriendlyMatchConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stakes []int32 `protobuf:"varint,1,rep,packed,name=Stakes,proto3" json:"Stakes,omitempty" yaml:"Stakes,omitempty" bson:"stakes,omitempty"`
	Fee    float32 `protobuf:"fixed32,2,opt,name=Fee,proto3" json:"Fee,omitempty" yaml:"Fee,omitempty" bson:"fee,omitempty"`
}

func (x *FriendlyMatchConfig) Reset() {
	*x = FriendlyMatchConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friends_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendlyMatchConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendlyMatchConfig) ProtoMessage() {}

func (x *FriendlyMatchConfig) ProtoReflect() protoreflect.Message {
	mi := &file_friends_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendlyMatchConfig.ProtoReflect.Descriptor instead.
func (*FriendlyMatchConfig) Descriptor() ([]byte, []int) {
	return file_friends_proto_rawDescGZIP(), []int{12}
}

func (x *FriendlyMatchConfig) GetStakes() []int32 {
	if x != nil {
		return x.Stakes
	}
	return nil
}

func (x *FriendlyMatchConfig) GetFee() float32 {
	if x != nil {
		return x.Fee
	}
	return 0
}

var File_friends_proto protoreflect.FileDescriptor

var file_friends_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x1a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb7, 0x01, 0x0a, 0x0b, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2c, 0x0a, 0x07, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x73, 0x12, 0x3c, 0x0a, 0x0f, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x49,
	0x6e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x0f, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65,
	0x73, 0x12, 0x3c, 0x0a, 0x0f, 0x4f, 0x75, 0x74, 0x67, 0x6f, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x76,
	0x69, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f,
	0x4f, 0x75, 0x74, 0x67, 0x6f, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x22,
	0x9e, 0x01, 0x0a, 0x10, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x2e, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x2e, 0x46, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x2e, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0x5c, 0x0a, 0x14, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x6c, 0x79, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x0a, 0x4f, 0x70, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x52, 0x0a, 0x4f, 0x70,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x6b,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x22, 0x53,
	0x0a, 0x14, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x6c, 0x79, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x44,
	0x65, 0x63, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73,
	0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x6c, 0x79, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65,
	0x63, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x52, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x22, 0x56, 0x0a, 0x13, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x6c, 0x79, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x3f, 0x0a, 0x0b, 0x41, 0x64,
	0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b,
	0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x15, 0x0a, 0x13, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x6c, 0x79, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x43, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x22, 0x55, 0x0a, 0x12, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x6c, 0x79, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x45, 0x6e, 0x64, 0x65, 0x64, 0x12, 0x3f, 0x0a, 0x0b, 0x41, 0x64, 0x6a, 0x75,
	0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x41, 0x64,
	0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x15, 0x0a, 0x13, 0x46, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x6c, 0x79, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74,
	0x22, 0x85, 0x01, 0x0a, 0x11, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x6c, 0x79, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x2e, 0x0a, 0x0a, 0x4f, 0x70, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x52, 0x0a, 0x4f, 0x70, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x22, 0x23, 0x0a, 0x0b, 0x46, 0x69, 0x6e, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x22, 0x38, 0x0a,
	0x0c, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a,
	0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x22, 0x5a, 0x0a, 0x10, 0x46, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x49,
	0x73, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x49,
	0x73, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x2a, 0x0a, 0x08, 0x46, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x52, 0x08, 0x46, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x49, 0x44, 0x22, 0x3f, 0x0a, 0x13, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x6c, 0x79, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x6b, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x06, 0x53, 0x74, 0x61, 0x6b,
	0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x46, 0x65, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x03, 0x46, 0x65, 0x65, 0x2a, 0x7a, 0x0a, 0x11, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x10, 0x00, 0x12, 0x13,
	0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65,
	0x64, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x65,
	0x63, 0x6c, 0x69, 0x6e, 0x65, 0x64, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x65, 0x64, 0x10, 0x03, 0x12, 0x11, 0x0a,
	0x0d, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x10, 0x04,
	0x2a, 0x3a, 0x0a, 0x1a, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x6c, 0x79, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x44, 0x65, 0x63, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x0b,
	0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x4c,
	0x65, 0x66, 0x74, 0x54, 0x68, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x10, 0x01, 0x42, 0x34, 0x5a, 0x1f,
	0x6c, 0x6f, 0x63, 0x67, 0x61, 0x6d, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x64, 0x74, 0x6f, 0x2f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0xaa,
	0x02, 0x10, 0x4c, 0x6f, 0x43, 0x2e, 0x44, 0x54, 0x4f, 0x73, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_friends_proto_rawDescOnce sync.Once
	file_friends_proto_rawDescData = file_friends_proto_rawDesc
)

func file_friends_proto_rawDescGZIP() []byte {
	file_friends_proto_rawDescOnce.Do(func() {
		file_friends_proto_rawDescData = protoimpl.X.CompressGZIP(file_friends_proto_rawDescData)
	})
	return file_friends_proto_rawDescData
}

var file_friends_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_friends_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_friends_proto_goTypes = []any{
	(FriendRequestType)(0),               // 0: friends.FriendRequestType
	(FriendlyMatchDeclineReason)(0),      // 1: friends.FriendlyMatchDeclineReason
	(*FriendsData)(nil),                  // 2: friends.FriendsData
	(*FriendChangeData)(nil),             // 3: friends.FriendChangeData
	(*FriendlyMatchRequest)(nil),         // 4: friends.FriendlyMatchRequest
	(*FriendlyMatchDecline)(nil),         // 5: friends.FriendlyMatchDecline
	(*FriendlyMatchResult)(nil),          // 6: friends.FriendlyMatchResult
	(*FriendlyMatchCancel)(nil),          // 7: friends.FriendlyMatchCancel
	(*FriendlyMatchEnded)(nil),           // 8: friends.FriendlyMatchEnded
	(*FriendlyMatchAccept)(nil),          // 9: friends.FriendlyMatchAccept
	(*FriendlyMatchData)(nil),            // 10: friends.FriendlyMatchData
	(*FindRequest)(nil),                  // 11: friends.FindRequest
	(*FindResponse)(nil),                 // 12: friends.FindResponse
	(*FriendStatusData)(nil),             // 13: friends.FriendStatusData
	(*FriendlyMatchConfig)(nil),          // 14: friends.FriendlyMatchConfig
	(*accounts.UserInfo)(nil),            // 15: accounts.UserInfo
	(*base.ObjectID)(nil),                // 16: base.ObjectID
	(*resources.ResourceAdjustment)(nil), // 17: resources.ResourceAdjustment
}
var file_friends_proto_depIdxs = []int32{
	15, // 0: friends.FriendsData.Friends:type_name -> accounts.UserInfo
	15, // 1: friends.FriendsData.IncomingInvites:type_name -> accounts.UserInfo
	15, // 2: friends.FriendsData.OutgoingInvites:type_name -> accounts.UserInfo
	0,  // 3: friends.FriendChangeData.Type:type_name -> friends.FriendRequestType
	16, // 4: friends.FriendChangeData.PlayerID:type_name -> base.ObjectID
	15, // 5: friends.FriendChangeData.UserInfo:type_name -> accounts.UserInfo
	16, // 6: friends.FriendlyMatchRequest.OpponentID:type_name -> base.ObjectID
	1,  // 7: friends.FriendlyMatchDecline.Reason:type_name -> friends.FriendlyMatchDeclineReason
	17, // 8: friends.FriendlyMatchResult.Adjustments:type_name -> resources.ResourceAdjustment
	17, // 9: friends.FriendlyMatchEnded.Adjustments:type_name -> resources.ResourceAdjustment
	16, // 10: friends.FriendlyMatchData.PlayerID:type_name -> base.ObjectID
	16, // 11: friends.FriendlyMatchData.OpponentID:type_name -> base.ObjectID
	15, // 12: friends.FindResponse.Users:type_name -> accounts.UserInfo
	16, // 13: friends.FriendStatusData.FriendID:type_name -> base.ObjectID
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_friends_proto_init() }
func file_friends_proto_init() {
	if File_friends_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_friends_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*FriendsData); i {
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
		file_friends_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*FriendChangeData); i {
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
		file_friends_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*FriendlyMatchRequest); i {
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
		file_friends_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*FriendlyMatchDecline); i {
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
		file_friends_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*FriendlyMatchResult); i {
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
		file_friends_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*FriendlyMatchCancel); i {
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
		file_friends_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*FriendlyMatchEnded); i {
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
		file_friends_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*FriendlyMatchAccept); i {
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
		file_friends_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*FriendlyMatchData); i {
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
		file_friends_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*FindRequest); i {
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
		file_friends_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*FindResponse); i {
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
		file_friends_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*FriendStatusData); i {
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
		file_friends_proto_msgTypes[12].Exporter = func(v any, i int) any {
			switch v := v.(*FriendlyMatchConfig); i {
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
			RawDescriptor: file_friends_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_friends_proto_goTypes,
		DependencyIndexes: file_friends_proto_depIdxs,
		EnumInfos:         file_friends_proto_enumTypes,
		MessageInfos:      file_friends_proto_msgTypes,
	}.Build()
	File_friends_proto = out.File
	file_friends_proto_rawDesc = nil
	file_friends_proto_goTypes = nil
	file_friends_proto_depIdxs = nil
}
