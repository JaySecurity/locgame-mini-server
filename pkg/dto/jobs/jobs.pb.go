// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.1
// source: jobs.proto

package jobs

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

type JobStatus int32

const (
	JobStatus_NotSet  JobStatus = 0
	JobStatus_Success JobStatus = 1
	JobStatus_Running JobStatus = 2
	JobStatus_Failed  JobStatus = 3
)

// Enum value maps for JobStatus.
var (
	JobStatus_name = map[int32]string{
		0: "NotSet",
		1: "Success",
		2: "Running",
		3: "Failed",
	}
	JobStatus_value = map[string]int32{
		"NotSet":  0,
		"Success": 1,
		"Running": 2,
		"Failed":  3,
	}
)

func (x JobStatus) Enum() *JobStatus {
	p := new(JobStatus)
	*p = x
	return p
}

func (x JobStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JobStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_jobs_proto_enumTypes[0].Descriptor()
}

func (JobStatus) Type() protoreflect.EnumType {
	return &file_jobs_proto_enumTypes[0]
}

func (x JobStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JobStatus.Descriptor instead.
func (JobStatus) EnumDescriptor() ([]byte, []int) {
	return file_jobs_proto_rawDescGZIP(), []int{0}
}

type JobData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          *base.ObjectID  `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty" yaml:"ID,omitempty" bson:"_id,omitempty"`
	ParentJobID *base.ObjectID  `protobuf:"bytes,2,opt,name=ParentJobID,proto3" json:"ParentJobID,omitempty" yaml:"ParentJobID,omitempty" bson:"parent_job_id,omitempty"`
	Name        string          `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty" yaml:"Name,omitempty" bson:"name,omitempty"`
	Attempt     uint32          `protobuf:"varint,4,opt,name=Attempt,proto3" json:"Attempt,omitempty" yaml:"Attempt,omitempty" bson:"attempt,omitempty"`
	Status      JobStatus       `protobuf:"varint,5,opt,name=Status,proto3,enum=proto.JobStatus" json:"Status,omitempty" yaml:"Status,omitempty" bson:"status,omitempty"`
	StartedAt   *base.Timestamp `protobuf:"bytes,6,opt,name=StartedAt,proto3" json:"StartedAt,omitempty" yaml:"StartedAt,omitempty" bson:"started_at,omitempty"`
	FinishedAt  *base.Timestamp `protobuf:"bytes,7,opt,name=FinishedAt,proto3" json:"FinishedAt,omitempty" yaml:"FinishedAt,omitempty" bson:"finished_at,omitempty"`
	Output      string          `protobuf:"bytes,8,opt,name=Output,proto3" json:"Output,omitempty" yaml:"Output,omitempty" bson:"output,omitempty"`
}

func (x *JobData) Reset() {
	*x = JobData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jobs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobData) ProtoMessage() {}

func (x *JobData) ProtoReflect() protoreflect.Message {
	mi := &file_jobs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobData.ProtoReflect.Descriptor instead.
func (*JobData) Descriptor() ([]byte, []int) {
	return file_jobs_proto_rawDescGZIP(), []int{0}
}

func (x *JobData) GetID() *base.ObjectID {
	if x != nil {
		return x.ID
	}
	return nil
}

func (x *JobData) GetParentJobID() *base.ObjectID {
	if x != nil {
		return x.ParentJobID
	}
	return nil
}

func (x *JobData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *JobData) GetAttempt() uint32 {
	if x != nil {
		return x.Attempt
	}
	return 0
}

func (x *JobData) GetStatus() JobStatus {
	if x != nil {
		return x.Status
	}
	return JobStatus_NotSet
}

func (x *JobData) GetStartedAt() *base.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *JobData) GetFinishedAt() *base.Timestamp {
	if x != nil {
		return x.FinishedAt
	}
	return nil
}

func (x *JobData) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

type RecurringJobData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID            *base.ObjectID  `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty" yaml:"ID,omitempty" bson:"_id,omitempty"`
	Name          string          `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty" yaml:"Name,omitempty" bson:"name,omitempty"`
	Schedule      string          `protobuf:"bytes,3,opt,name=Schedule,proto3" json:"Schedule,omitempty" yaml:"Schedule,omitempty" bson:"schedule,omitempty"`
	LastExecution *base.Timestamp `protobuf:"bytes,4,opt,name=LastExecution,proto3" json:"LastExecution,omitempty" yaml:"LastExecution,omitempty" bson:"last_execution,omitempty"`
	SuccessCount  uint32          `protobuf:"varint,5,opt,name=SuccessCount,proto3" json:"SuccessCount,omitempty" yaml:"SuccessCount,omitempty" bson:"success_count,omitempty"`
	ErrorCount    uint32          `protobuf:"varint,6,opt,name=ErrorCount,proto3" json:"ErrorCount,omitempty" yaml:"ErrorCount,omitempty" bson:"error_count,omitempty"`
	Disabled      bool            `protobuf:"varint,7,opt,name=Disabled,proto3" json:"Disabled,omitempty" yaml:"Disabled,omitempty" bson:"disabled"`
	Retries       uint32          `protobuf:"varint,8,opt,name=Retries,proto3" json:"Retries,omitempty" yaml:"Retries,omitempty" bson:"retries,omitempty"`
	Status        JobStatus       `protobuf:"varint,9,opt,name=Status,proto3,enum=proto.JobStatus" json:"Status,omitempty" yaml:"Status,omitempty" bson:"status,omitempty"`
	NextExecution *base.Timestamp `protobuf:"bytes,10,opt,name=NextExecution,proto3" json:"NextExecution,omitempty" yaml:"NextExecution,omitempty" bson:"next_execution,omitempty"`
}

func (x *RecurringJobData) Reset() {
	*x = RecurringJobData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jobs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecurringJobData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecurringJobData) ProtoMessage() {}

func (x *RecurringJobData) ProtoReflect() protoreflect.Message {
	mi := &file_jobs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecurringJobData.ProtoReflect.Descriptor instead.
func (*RecurringJobData) Descriptor() ([]byte, []int) {
	return file_jobs_proto_rawDescGZIP(), []int{1}
}

func (x *RecurringJobData) GetID() *base.ObjectID {
	if x != nil {
		return x.ID
	}
	return nil
}

func (x *RecurringJobData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RecurringJobData) GetSchedule() string {
	if x != nil {
		return x.Schedule
	}
	return ""
}

func (x *RecurringJobData) GetLastExecution() *base.Timestamp {
	if x != nil {
		return x.LastExecution
	}
	return nil
}

func (x *RecurringJobData) GetSuccessCount() uint32 {
	if x != nil {
		return x.SuccessCount
	}
	return 0
}

func (x *RecurringJobData) GetErrorCount() uint32 {
	if x != nil {
		return x.ErrorCount
	}
	return 0
}

func (x *RecurringJobData) GetDisabled() bool {
	if x != nil {
		return x.Disabled
	}
	return false
}

func (x *RecurringJobData) GetRetries() uint32 {
	if x != nil {
		return x.Retries
	}
	return 0
}

func (x *RecurringJobData) GetStatus() JobStatus {
	if x != nil {
		return x.Status
	}
	return JobStatus_NotSet
}

func (x *RecurringJobData) GetNextExecution() *base.Timestamp {
	if x != nil {
		return x.NextExecution
	}
	return nil
}

type TriggerNowMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobName string `protobuf:"bytes,1,opt,name=JobName,proto3" json:"JobName,omitempty" yaml:"JobName,omitempty" bson:"job_name,omitempty"`
}

func (x *TriggerNowMessage) Reset() {
	*x = TriggerNowMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jobs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TriggerNowMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TriggerNowMessage) ProtoMessage() {}

func (x *TriggerNowMessage) ProtoReflect() protoreflect.Message {
	mi := &file_jobs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TriggerNowMessage.ProtoReflect.Descriptor instead.
func (*TriggerNowMessage) Descriptor() ([]byte, []int) {
	return file_jobs_proto_rawDescGZIP(), []int{2}
}

func (x *TriggerNowMessage) GetJobName() string {
	if x != nil {
		return x.JobName
	}
	return ""
}

var File_jobs_proto protoreflect.FileDescriptor

var file_jobs_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xab, 0x02, 0x0a, 0x07, 0x4a, 0x6f, 0x62, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x52, 0x02, 0x49, 0x44, 0x12, 0x30, 0x0a, 0x0b, 0x50,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x4a, 0x6f, 0x62, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44,
	0x52, 0x0b, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x4a, 0x6f, 0x62, 0x49, 0x44, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x41, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x12, 0x28, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2d, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x2f, 0x0a, 0x0a, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x46, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x65, 0x64, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0xf4, 0x02,
	0x0a, 0x10, 0x52, 0x65, 0x63, 0x75, 0x72, 0x72, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x1e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x12, 0x35, 0x0a, 0x0d, 0x4c, 0x61, 0x73, 0x74, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x4c, 0x61, 0x73, 0x74,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x53, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0c, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x65, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x52, 0x65, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4a, 0x6f, 0x62, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x35, 0x0a,
	0x0d, 0x4e, 0x65, 0x78, 0x74, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x4e, 0x65, 0x78, 0x74, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2d, 0x0a, 0x11, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x4e,
	0x6f, 0x77, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4a, 0x6f, 0x62,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4a, 0x6f, 0x62, 0x4e,
	0x61, 0x6d, 0x65, 0x2a, 0x3d, 0x0a, 0x09, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x6f, 0x74, 0x53, 0x65, 0x74, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x75, 0x6e,
	0x6e, 0x69, 0x6e, 0x67, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64,
	0x10, 0x03, 0x42, 0x34, 0x48, 0x01, 0x5a, 0x20, 0x6c, 0x6f, 0x63, 0x67, 0x61, 0x6d, 0x65, 0x2d,
	0x6d, 0x69, 0x6e, 0x69, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x64, 0x74, 0x6f, 0x2f, 0x6a, 0x6f, 0x62, 0x73, 0xaa, 0x02, 0x0d, 0x4c, 0x6f, 0x43, 0x2e, 0x44,
	0x54, 0x4f, 0x73, 0x2e, 0x4a, 0x6f, 0x62, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_jobs_proto_rawDescOnce sync.Once
	file_jobs_proto_rawDescData = file_jobs_proto_rawDesc
)

func file_jobs_proto_rawDescGZIP() []byte {
	file_jobs_proto_rawDescOnce.Do(func() {
		file_jobs_proto_rawDescData = protoimpl.X.CompressGZIP(file_jobs_proto_rawDescData)
	})
	return file_jobs_proto_rawDescData
}

var file_jobs_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_jobs_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_jobs_proto_goTypes = []any{
	(JobStatus)(0),            // 0: proto.JobStatus
	(*JobData)(nil),           // 1: proto.JobData
	(*RecurringJobData)(nil),  // 2: proto.RecurringJobData
	(*TriggerNowMessage)(nil), // 3: proto.TriggerNowMessage
	(*base.ObjectID)(nil),     // 4: base.ObjectID
	(*base.Timestamp)(nil),    // 5: base.Timestamp
}
var file_jobs_proto_depIdxs = []int32{
	4, // 0: proto.JobData.ID:type_name -> base.ObjectID
	4, // 1: proto.JobData.ParentJobID:type_name -> base.ObjectID
	0, // 2: proto.JobData.Status:type_name -> proto.JobStatus
	5, // 3: proto.JobData.StartedAt:type_name -> base.Timestamp
	5, // 4: proto.JobData.FinishedAt:type_name -> base.Timestamp
	4, // 5: proto.RecurringJobData.ID:type_name -> base.ObjectID
	5, // 6: proto.RecurringJobData.LastExecution:type_name -> base.Timestamp
	0, // 7: proto.RecurringJobData.Status:type_name -> proto.JobStatus
	5, // 8: proto.RecurringJobData.NextExecution:type_name -> base.Timestamp
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_jobs_proto_init() }
func file_jobs_proto_init() {
	if File_jobs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_jobs_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*JobData); i {
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
		file_jobs_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*RecurringJobData); i {
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
		file_jobs_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*TriggerNowMessage); i {
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
			RawDescriptor: file_jobs_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_jobs_proto_goTypes,
		DependencyIndexes: file_jobs_proto_depIdxs,
		EnumInfos:         file_jobs_proto_enumTypes,
		MessageInfos:      file_jobs_proto_msgTypes,
	}.Build()
	File_jobs_proto = out.File
	file_jobs_proto_rawDesc = nil
	file_jobs_proto_goTypes = nil
	file_jobs_proto_depIdxs = nil
}
