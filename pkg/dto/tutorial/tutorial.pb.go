// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.1
// source: tutorial.proto

package tutorial

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CompleteTutorialStepRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty" yaml:"ID,omitempty" bson:"id,omitempty"`
}

func (x *CompleteTutorialStepRequest) Reset() {
	*x = CompleteTutorialStepRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tutorial_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteTutorialStepRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteTutorialStepRequest) ProtoMessage() {}

func (x *CompleteTutorialStepRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tutorial_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteTutorialStepRequest.ProtoReflect.Descriptor instead.
func (*CompleteTutorialStepRequest) Descriptor() ([]byte, []int) {
	return file_tutorial_proto_rawDescGZIP(), []int{0}
}

func (x *CompleteTutorialStepRequest) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

type TutorialData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentStep            int32   `protobuf:"varint,1,opt,name=CurrentStep,proto3" json:"CurrentStep,omitempty" yaml:"CurrentStep,omitempty" bson:"current_step,omitempty"`
	CompletedSoftTutorials []int32 `protobuf:"varint,2,rep,packed,name=CompletedSoftTutorials,proto3" json:"CompletedSoftTutorials,omitempty" yaml:"CompletedSoftTutorials,omitempty" bson:"completed_soft_tutorials,omitempty"`
}

func (x *TutorialData) Reset() {
	*x = TutorialData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tutorial_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TutorialData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TutorialData) ProtoMessage() {}

func (x *TutorialData) ProtoReflect() protoreflect.Message {
	mi := &file_tutorial_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TutorialData.ProtoReflect.Descriptor instead.
func (*TutorialData) Descriptor() ([]byte, []int) {
	return file_tutorial_proto_rawDescGZIP(), []int{1}
}

func (x *TutorialData) GetCurrentStep() int32 {
	if x != nil {
		return x.CurrentStep
	}
	return 0
}

func (x *TutorialData) GetCompletedSoftTutorials() []int32 {
	if x != nil {
		return x.CompletedSoftTutorials
	}
	return nil
}

var File_tutorial_proto protoreflect.FileDescriptor

var file_tutorial_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x22, 0x2d, 0x0a, 0x1b, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x74,
	0x65, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x22, 0x68, 0x0a, 0x0c, 0x54, 0x75, 0x74,
	0x6f, 0x72, 0x69, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x65, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x65, 0x70, 0x12, 0x36, 0x0a, 0x16, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x53, 0x6f, 0x66, 0x74, 0x54, 0x75, 0x74, 0x6f,
	0x72, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x16, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x53, 0x6f, 0x66, 0x74, 0x54, 0x75, 0x74, 0x6f, 0x72, 0x69,
	0x61, 0x6c, 0x73, 0x42, 0x3a, 0x5a, 0x24, 0x6c, 0x6f, 0x63, 0x67, 0x61, 0x6d, 0x65, 0x2d, 0x6d,
	0x69, 0x6e, 0x69, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x64,
	0x74, 0x6f, 0x2f, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0xaa, 0x02, 0x11, 0x4c, 0x6f,
	0x43, 0x2e, 0x44, 0x54, 0x4f, 0x73, 0x2e, 0x54, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tutorial_proto_rawDescOnce sync.Once
	file_tutorial_proto_rawDescData = file_tutorial_proto_rawDesc
)

func file_tutorial_proto_rawDescGZIP() []byte {
	file_tutorial_proto_rawDescOnce.Do(func() {
		file_tutorial_proto_rawDescData = protoimpl.X.CompressGZIP(file_tutorial_proto_rawDescData)
	})
	return file_tutorial_proto_rawDescData
}

var file_tutorial_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tutorial_proto_goTypes = []any{
	(*CompleteTutorialStepRequest)(nil), // 0: tutorial.CompleteTutorialStepRequest
	(*TutorialData)(nil),                // 1: tutorial.TutorialData
}
var file_tutorial_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tutorial_proto_init() }
func file_tutorial_proto_init() {
	if File_tutorial_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tutorial_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CompleteTutorialStepRequest); i {
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
		file_tutorial_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*TutorialData); i {
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
			RawDescriptor: file_tutorial_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tutorial_proto_goTypes,
		DependencyIndexes: file_tutorial_proto_depIdxs,
		MessageInfos:      file_tutorial_proto_msgTypes,
	}.Build()
	File_tutorial_proto = out.File
	file_tutorial_proto_rawDesc = nil
	file_tutorial_proto_goTypes = nil
	file_tutorial_proto_depIdxs = nil
}