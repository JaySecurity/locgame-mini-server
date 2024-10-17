// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.1
// source: resources.proto

package resources

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	base "locgame-mini-server/pkg/dto/base"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CapacityType int32

const (
	CapacityType_SoftCapacity CapacityType = 0
	CapacityType_HardCapacity CapacityType = 1
)

// Enum value maps for CapacityType.
var (
	CapacityType_name = map[int32]string{
		0: "SoftCapacity",
		1: "HardCapacity",
	}
	CapacityType_value = map[string]int32{
		"SoftCapacity": 0,
		"HardCapacity": 1,
	}
)

func (x CapacityType) Enum() *CapacityType {
	p := new(CapacityType)
	*p = x
	return p
}

func (x CapacityType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CapacityType) Descriptor() protoreflect.EnumDescriptor {
	return file_resources_proto_enumTypes[0].Descriptor()
}

func (CapacityType) Type() protoreflect.EnumType {
	return &file_resources_proto_enumTypes[0]
}

func (x CapacityType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CapacityType.Descriptor instead.
func (CapacityType) EnumDescriptor() ([]byte, []int) {
	return file_resources_proto_rawDescGZIP(), []int{0}
}

type ResourceData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         int32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty" yaml:"ID,omitempty" bson:"id,omitempty"`
	Key        string `protobuf:"bytes,2,opt,name=Key,proto3" json:"Key,omitempty" yaml:"Key,omitempty" bson:"key,omitempty"`
	CategoryID int32  `protobuf:"varint,3,opt,name=CategoryID,proto3" json:"CategoryID,omitempty" yaml:"CategoryID,omitempty" bson:"category_id,omitempty"`
}

func (x *ResourceData) Reset() {
	*x = ResourceData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceData) ProtoMessage() {}

func (x *ResourceData) ProtoReflect() protoreflect.Message {
	mi := &file_resources_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceData.ProtoReflect.Descriptor instead.
func (*ResourceData) Descriptor() ([]byte, []int) {
	return file_resources_proto_rawDescGZIP(), []int{0}
}

func (x *ResourceData) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ResourceData) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ResourceData) GetCategoryID() int32 {
	if x != nil {
		return x.CategoryID
	}
	return 0
}

type ResourceCategory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID  int32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty" yaml:"ID,omitempty" bson:"id,omitempty"`
	Key string `protobuf:"bytes,2,opt,name=Key,proto3" json:"Key,omitempty" yaml:"Key,omitempty" bson:"key,omitempty"`
}

func (x *ResourceCategory) Reset() {
	*x = ResourceCategory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceCategory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceCategory) ProtoMessage() {}

func (x *ResourceCategory) ProtoReflect() protoreflect.Message {
	mi := &file_resources_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceCategory.ProtoReflect.Descriptor instead.
func (*ResourceCategory) Descriptor() ([]byte, []int) {
	return file_resources_proto_rawDescGZIP(), []int{1}
}

func (x *ResourceCategory) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ResourceCategory) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type ResourceAdjustment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceID int32  `protobuf:"varint,1,opt,name=ResourceID,proto3" json:"ResourceID,omitempty" yaml:"ResourceID,omitempty" bson:"resource_id,omitempty"`
	Quantity   int32  `protobuf:"varint,2,opt,name=Quantity,proto3" json:"Quantity,omitempty" yaml:"Quantity,omitempty" bson:"quantity,omitempty"`
	Reason     string `protobuf:"bytes,3,opt,name=Reason,proto3" json:"Reason,omitempty" yaml:"Reason,omitempty" bson:"reason,omitempty"`
}

func (x *ResourceAdjustment) Reset() {
	*x = ResourceAdjustment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceAdjustment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceAdjustment) ProtoMessage() {}

func (x *ResourceAdjustment) ProtoReflect() protoreflect.Message {
	mi := &file_resources_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceAdjustment.ProtoReflect.Descriptor instead.
func (*ResourceAdjustment) Descriptor() ([]byte, []int) {
	return file_resources_proto_rawDescGZIP(), []int{2}
}

func (x *ResourceAdjustment) GetResourceID() int32 {
	if x != nil {
		return x.ResourceID
	}
	return 0
}

func (x *ResourceAdjustment) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *ResourceAdjustment) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type ResourceAdjustments struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Adjustments []*ResourceAdjustment `protobuf:"bytes,1,rep,name=Adjustments,proto3" json:"Adjustments,omitempty" yaml:"Adjustments,omitempty" bson:"adjustments,omitempty"`
}

func (x *ResourceAdjustments) Reset() {
	*x = ResourceAdjustments{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceAdjustments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceAdjustments) ProtoMessage() {}

func (x *ResourceAdjustments) ProtoReflect() protoreflect.Message {
	mi := &file_resources_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceAdjustments.ProtoReflect.Descriptor instead.
func (*ResourceAdjustments) Descriptor() ([]byte, []int) {
	return file_resources_proto_rawDescGZIP(), []int{3}
}

func (x *ResourceAdjustments) GetAdjustments() []*ResourceAdjustment {
	if x != nil {
		return x.Adjustments
	}
	return nil
}

type ResettableResources struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resources map[int32]*ResettableResource `protobuf:"bytes,1,rep,name=Resources,proto3" json:"Resources,omitempty" yaml:"Resources,omitempty" bson:"resources,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ResettableResources) Reset() {
	*x = ResettableResources{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResettableResources) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResettableResources) ProtoMessage() {}

func (x *ResettableResources) ProtoReflect() protoreflect.Message {
	mi := &file_resources_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResettableResources.ProtoReflect.Descriptor instead.
func (*ResettableResources) Descriptor() ([]byte, []int) {
	return file_resources_proto_rawDescGZIP(), []int{4}
}

func (x *ResettableResources) GetResources() map[int32]*ResettableResource {
	if x != nil {
		return x.Resources
	}
	return nil
}

type ResettableResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceID int32  `protobuf:"varint,1,opt,name=ResourceID,proto3" json:"ResourceID,omitempty" yaml:"ResourceID,omitempty" bson:"resource_id,omitempty"`
	ResetTime  string `protobuf:"bytes,2,opt,name=ResetTime,proto3" json:"ResetTime,omitempty" yaml:"ResetTime,omitempty" bson:"reset_time,omitempty"`
}

func (x *ResettableResource) Reset() {
	*x = ResettableResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResettableResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResettableResource) ProtoMessage() {}

func (x *ResettableResource) ProtoReflect() protoreflect.Message {
	mi := &file_resources_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResettableResource.ProtoReflect.Descriptor instead.
func (*ResettableResource) Descriptor() ([]byte, []int) {
	return file_resources_proto_rawDescGZIP(), []int{5}
}

func (x *ResettableResource) GetResourceID() int32 {
	if x != nil {
		return x.ResourceID
	}
	return 0
}

func (x *ResettableResource) GetResetTime() string {
	if x != nil {
		return x.ResetTime
	}
	return ""
}

type ResettableResourceData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NextResetTime *base.Timestamp `protobuf:"bytes,1,opt,name=NextResetTime,proto3" json:"NextResetTime,omitempty" yaml:"NextResetTime,omitempty" bson:"next_reset_time,omitempty"`
}

func (x *ResettableResourceData) Reset() {
	*x = ResettableResourceData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResettableResourceData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResettableResourceData) ProtoMessage() {}

func (x *ResettableResourceData) ProtoReflect() protoreflect.Message {
	mi := &file_resources_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResettableResourceData.ProtoReflect.Descriptor instead.
func (*ResettableResourceData) Descriptor() ([]byte, []int) {
	return file_resources_proto_rawDescGZIP(), []int{6}
}

func (x *ResettableResourceData) GetNextResetTime() *base.Timestamp {
	if x != nil {
		return x.NextResetTime
	}
	return nil
}

type CappedResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceID   int32           `protobuf:"varint,1,opt,name=ResourceID,proto3" json:"ResourceID,omitempty" yaml:"ResourceID,omitempty" bson:"resource_id,omitempty"`
	CapacityType CapacityType    `protobuf:"varint,2,opt,name=CapacityType,proto3,enum=resources.CapacityType" json:"CapacityType,omitempty" yaml:"CapacityType,omitempty" bson:"capacity_type,omitempty"`
	Capacities   map[int32]int32 `protobuf:"bytes,3,rep,name=Capacities,proto3" json:"Capacities,omitempty" yaml:"Capacities,omitempty" bson:"capacities,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *CappedResource) Reset() {
	*x = CappedResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CappedResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CappedResource) ProtoMessage() {}

func (x *CappedResource) ProtoReflect() protoreflect.Message {
	mi := &file_resources_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CappedResource.ProtoReflect.Descriptor instead.
func (*CappedResource) Descriptor() ([]byte, []int) {
	return file_resources_proto_rawDescGZIP(), []int{7}
}

func (x *CappedResource) GetResourceID() int32 {
	if x != nil {
		return x.ResourceID
	}
	return 0
}

func (x *CappedResource) GetCapacityType() CapacityType {
	if x != nil {
		return x.CapacityType
	}
	return CapacityType_SoftCapacity
}

func (x *CappedResource) GetCapacities() map[int32]int32 {
	if x != nil {
		return x.Capacities
	}
	return nil
}

type CappedResources struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resources map[int32]*CappedResource `protobuf:"bytes,1,rep,name=Resources,proto3" json:"Resources,omitempty" yaml:"Resources,omitempty" bson:"resources,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CappedResources) Reset() {
	*x = CappedResources{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CappedResources) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CappedResources) ProtoMessage() {}

func (x *CappedResources) ProtoReflect() protoreflect.Message {
	mi := &file_resources_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CappedResources.ProtoReflect.Descriptor instead.
func (*CappedResources) Descriptor() ([]byte, []int) {
	return file_resources_proto_rawDescGZIP(), []int{8}
}

func (x *CappedResources) GetResources() map[int32]*CappedResource {
	if x != nil {
		return x.Resources
	}
	return nil
}

type WithdrawRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LC int32 `protobuf:"varint,1,opt,name=LC,proto3" json:"LC,omitempty" yaml:"LC,omitempty" bson:"lc,omitempty"`
}

func (x *WithdrawRequest) Reset() {
	*x = WithdrawRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WithdrawRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WithdrawRequest) ProtoMessage() {}

func (x *WithdrawRequest) ProtoReflect() protoreflect.Message {
	mi := &file_resources_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WithdrawRequest.ProtoReflect.Descriptor instead.
func (*WithdrawRequest) Descriptor() ([]byte, []int) {
	return file_resources_proto_rawDescGZIP(), []int{9}
}

func (x *WithdrawRequest) GetLC() int32 {
	if x != nil {
		return x.LC
	}
	return 0
}

type WithdrawResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionHash string                `protobuf:"bytes,1,opt,name=TransactionHash,proto3" json:"TransactionHash,omitempty" yaml:"TransactionHash,omitempty" bson:"transaction_hash,omitempty"`
	LOCG            float64               `protobuf:"fixed64,2,opt,name=LOCG,proto3" json:"LOCG,omitempty" yaml:"LOCG,omitempty" bson:"locg,omitempty"`
	Adjustments     []*ResourceAdjustment `protobuf:"bytes,3,rep,name=Adjustments,proto3" json:"Adjustments,omitempty" yaml:"Adjustments,omitempty" bson:"adjustments,omitempty"`
}

func (x *WithdrawResponse) Reset() {
	*x = WithdrawResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WithdrawResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WithdrawResponse) ProtoMessage() {}

func (x *WithdrawResponse) ProtoReflect() protoreflect.Message {
	mi := &file_resources_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WithdrawResponse.ProtoReflect.Descriptor instead.
func (*WithdrawResponse) Descriptor() ([]byte, []int) {
	return file_resources_proto_rawDescGZIP(), []int{10}
}

func (x *WithdrawResponse) GetTransactionHash() string {
	if x != nil {
		return x.TransactionHash
	}
	return ""
}

func (x *WithdrawResponse) GetLOCG() float64 {
	if x != nil {
		return x.LOCG
	}
	return 0
}

func (x *WithdrawResponse) GetAdjustments() []*ResourceAdjustment {
	if x != nil {
		return x.Adjustments
	}
	return nil
}

var File_resources_proto protoreflect.FileDescriptor

var file_resources_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x0a, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x22, 0x34, 0x0a, 0x10, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x10,
	0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79,
	0x22, 0x68, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x64, 0x6a, 0x75,
	0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x56, 0x0a, 0x13, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x3f, 0x0a, 0x0b, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x64, 0x6a, 0x75, 0x73,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x22, 0xbf, 0x01, 0x0a, 0x13, 0x52, 0x65, 0x73, 0x65, 0x74, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x4b, 0x0a, 0x09, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x5b, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x33, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x52, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x65, 0x74, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x65,
	0x73, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x52,
	0x65, 0x73, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x4f, 0x0a, 0x16, 0x52, 0x65, 0x73, 0x65,
	0x74, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x35, 0x0a, 0x0d, 0x4e, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x65, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x4e, 0x65, 0x78, 0x74,
	0x52, 0x65, 0x73, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xf7, 0x01, 0x0a, 0x0e, 0x43, 0x61,
	0x70, 0x70, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x44, 0x12, 0x3b, 0x0a, 0x0c,
	0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x17, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x43,
	0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x43, 0x61, 0x70,
	0x61, 0x63, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x49, 0x0a, 0x0a, 0x43, 0x61, 0x70,
	0x61, 0x63, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x70, 0x70, 0x65, 0x64,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x1a, 0x3d, 0x0a, 0x0f, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0xb3, 0x01, 0x0a, 0x0f, 0x43, 0x61, 0x70, 0x70, 0x65, 0x64, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x47, 0x0a, 0x09, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x70, 0x70, 0x65, 0x64, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x1a, 0x57, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x2f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e,
	0x43, 0x61, 0x70, 0x70, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x21, 0x0a, 0x0f, 0x57, 0x69, 0x74,
	0x68, 0x64, 0x72, 0x61, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x4c, 0x43, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x4c, 0x43, 0x22, 0x91, 0x01, 0x0a,
	0x10, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x28, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x4c,
	0x4f, 0x43, 0x47, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x4c, 0x4f, 0x43, 0x47, 0x12,
	0x3f, 0x0a, 0x0b, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x0b, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x2a, 0x32, 0x0a, 0x0c, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x10, 0x0a, 0x0c, 0x53, 0x6f, 0x66, 0x74, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79,
	0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x48, 0x61, 0x72, 0x64, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69,
	0x74, 0x79, 0x10, 0x01, 0x42, 0x3c, 0x5a, 0x25, 0x6c, 0x6f, 0x63, 0x67, 0x61, 0x6d, 0x65, 0x2d,
	0x6d, 0x69, 0x6e, 0x69, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x64, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xaa, 0x02, 0x12,
	0x4c, 0x6f, 0x43, 0x2e, 0x44, 0x54, 0x4f, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resources_proto_rawDescOnce sync.Once
	file_resources_proto_rawDescData = file_resources_proto_rawDesc
)

func file_resources_proto_rawDescGZIP() []byte {
	file_resources_proto_rawDescOnce.Do(func() {
		file_resources_proto_rawDescData = protoimpl.X.CompressGZIP(file_resources_proto_rawDescData)
	})
	return file_resources_proto_rawDescData
}

var file_resources_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_resources_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_resources_proto_goTypes = []any{
	(CapacityType)(0),              // 0: resources.CapacityType
	(*ResourceData)(nil),           // 1: resources.ResourceData
	(*ResourceCategory)(nil),       // 2: resources.ResourceCategory
	(*ResourceAdjustment)(nil),     // 3: resources.ResourceAdjustment
	(*ResourceAdjustments)(nil),    // 4: resources.ResourceAdjustments
	(*ResettableResources)(nil),    // 5: resources.ResettableResources
	(*ResettableResource)(nil),     // 6: resources.ResettableResource
	(*ResettableResourceData)(nil), // 7: resources.ResettableResourceData
	(*CappedResource)(nil),         // 8: resources.CappedResource
	(*CappedResources)(nil),        // 9: resources.CappedResources
	(*WithdrawRequest)(nil),        // 10: resources.WithdrawRequest
	(*WithdrawResponse)(nil),       // 11: resources.WithdrawResponse
	nil,                            // 12: resources.ResettableResources.ResourcesEntry
	nil,                            // 13: resources.CappedResource.CapacitiesEntry
	nil,                            // 14: resources.CappedResources.ResourcesEntry
	(*base.Timestamp)(nil),         // 15: base.Timestamp
}
var file_resources_proto_depIdxs = []int32{
	3,  // 0: resources.ResourceAdjustments.Adjustments:type_name -> resources.ResourceAdjustment
	12, // 1: resources.ResettableResources.Resources:type_name -> resources.ResettableResources.ResourcesEntry
	15, // 2: resources.ResettableResourceData.NextResetTime:type_name -> base.Timestamp
	0,  // 3: resources.CappedResource.CapacityType:type_name -> resources.CapacityType
	13, // 4: resources.CappedResource.Capacities:type_name -> resources.CappedResource.CapacitiesEntry
	14, // 5: resources.CappedResources.Resources:type_name -> resources.CappedResources.ResourcesEntry
	3,  // 6: resources.WithdrawResponse.Adjustments:type_name -> resources.ResourceAdjustment
	6,  // 7: resources.ResettableResources.ResourcesEntry.value:type_name -> resources.ResettableResource
	8,  // 8: resources.CappedResources.ResourcesEntry.value:type_name -> resources.CappedResource
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_resources_proto_init() }
func file_resources_proto_init() {
	if File_resources_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resources_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ResourceData); i {
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
		file_resources_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ResourceCategory); i {
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
		file_resources_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ResourceAdjustment); i {
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
		file_resources_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ResourceAdjustments); i {
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
		file_resources_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ResettableResources); i {
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
		file_resources_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ResettableResource); i {
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
		file_resources_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ResettableResourceData); i {
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
		file_resources_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*CappedResource); i {
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
		file_resources_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*CappedResources); i {
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
		file_resources_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*WithdrawRequest); i {
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
		file_resources_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*WithdrawResponse); i {
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
			RawDescriptor: file_resources_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_proto_goTypes,
		DependencyIndexes: file_resources_proto_depIdxs,
		EnumInfos:         file_resources_proto_enumTypes,
		MessageInfos:      file_resources_proto_msgTypes,
	}.Build()
	File_resources_proto = out.File
	file_resources_proto_rawDesc = nil
	file_resources_proto_goTypes = nil
	file_resources_proto_depIdxs = nil
}
