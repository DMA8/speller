// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: messages.proto

package proto

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SearchEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventTimestamp *timestamp.Timestamp `protobuf:"bytes,1,opt,name=eventTimestamp,proto3" json:"eventTimestamp,omitempty"`
	Query          string               `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	ShardKey       string               `protobuf:"bytes,3,opt,name=shardKey,proto3" json:"shardKey,omitempty"`
	Resource       string               `protobuf:"bytes,4,opt,name=resource,proto3" json:"resource,omitempty"`
	Category       string               `protobuf:"bytes,5,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *SearchEvent) Reset() {
	*x = SearchEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchEvent) ProtoMessage() {}

func (x *SearchEvent) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchEvent.ProtoReflect.Descriptor instead.
func (*SearchEvent) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{0}
}

func (x *SearchEvent) GetEventTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.EventTimestamp
	}
	return nil
}

func (x *SearchEvent) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *SearchEvent) GetShardKey() string {
	if x != nil {
		return x.ShardKey
	}
	return ""
}

func (x *SearchEvent) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *SearchEvent) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

type BadSearchEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventTimestamp *timestamp.Timestamp `protobuf:"bytes,1,opt,name=eventTimestamp,proto3" json:"eventTimestamp,omitempty"`
	Query          string               `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	Error          string               `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *BadSearchEvent) Reset() {
	*x = BadSearchEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BadSearchEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BadSearchEvent) ProtoMessage() {}

func (x *BadSearchEvent) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BadSearchEvent.ProtoReflect.Descriptor instead.
func (*BadSearchEvent) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{1}
}

func (x *BadSearchEvent) GetEventTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.EventTimestamp
	}
	return nil
}

func (x *BadSearchEvent) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *BadSearchEvent) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type AddNewSearchRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventTimestamp *timestamp.Timestamp `protobuf:"bytes,1,opt,name=eventTimestamp,proto3" json:"eventTimestamp,omitempty"`
	Query          string               `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	Miner          string               `protobuf:"bytes,3,opt,name=miner,proto3" json:"miner,omitempty"`
	MinerArgs      string               `protobuf:"bytes,4,opt,name=minerArgs,proto3" json:"minerArgs,omitempty"`
	ShardKey       string               `protobuf:"bytes,5,opt,name=shardKey,proto3" json:"shardKey,omitempty"`
	Resource       string               `protobuf:"bytes,6,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *AddNewSearchRecord) Reset() {
	*x = AddNewSearchRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddNewSearchRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddNewSearchRecord) ProtoMessage() {}

func (x *AddNewSearchRecord) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddNewSearchRecord.ProtoReflect.Descriptor instead.
func (*AddNewSearchRecord) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{2}
}

func (x *AddNewSearchRecord) GetEventTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.EventTimestamp
	}
	return nil
}

func (x *AddNewSearchRecord) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *AddNewSearchRecord) GetMiner() string {
	if x != nil {
		return x.Miner
	}
	return ""
}

func (x *AddNewSearchRecord) GetMinerArgs() string {
	if x != nil {
		return x.MinerArgs
	}
	return ""
}

func (x *AddNewSearchRecord) GetShardKey() string {
	if x != nil {
		return x.ShardKey
	}
	return ""
}

func (x *AddNewSearchRecord) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

type AddNewSynonym struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventTimestamp *timestamp.Timestamp `protobuf:"bytes,1,opt,name=eventTimestamp,proto3" json:"eventTimestamp,omitempty"`
	Query          string               `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *AddNewSynonym) Reset() {
	*x = AddNewSynonym{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddNewSynonym) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddNewSynonym) ProtoMessage() {}

func (x *AddNewSynonym) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddNewSynonym.ProtoReflect.Descriptor instead.
func (*AddNewSynonym) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{3}
}

func (x *AddNewSynonym) GetEventTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.EventTimestamp
	}
	return nil
}

func (x *AddNewSynonym) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

var File_messages_proto protoreflect.FileDescriptor

var file_messages_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb, 0x01, 0x0a, 0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x42, 0x0a, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12,
	0x1a, 0x0a, 0x08, 0x73, 0x68, 0x61, 0x72, 0x64, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x68, 0x61, 0x72, 0x64, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x22, 0x80, 0x01, 0x0a, 0x0e, 0x42, 0x61, 0x64, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x42, 0x0a, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0xda, 0x01, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x4e, 0x65,
	0x77, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x42, 0x0a,
	0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x69, 0x6e, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x1c, 0x0a,
	0x09, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x41, 0x72, 0x67, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x41, 0x72, 0x67, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x68, 0x61, 0x72, 0x64, 0x4b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x68, 0x61, 0x72, 0x64, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x22, 0x69, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x4e, 0x65, 0x77, 0x53, 0x79, 0x6e,
	0x6f, 0x6e, 0x79, 0x6d, 0x12, 0x42, 0x0a, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x42, 0x3b,
	0x5a, 0x39, 0x67, 0x69, 0x74, 0x2e, 0x77, 0x69, 0x6c, 0x64, 0x62, 0x65, 0x72, 0x72, 0x69, 0x65,
	0x73, 0x2e, 0x72, 0x75, 0x2f, 0x69, 0x6e, 0x6e, 0x6f, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x77, 0x62, 0x78, 0x2d, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2f, 0x65, 0x78, 0x61, 0x63, 0x74,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_messages_proto_rawDescOnce sync.Once
	file_messages_proto_rawDescData = file_messages_proto_rawDesc
)

func file_messages_proto_rawDescGZIP() []byte {
	file_messages_proto_rawDescOnce.Do(func() {
		file_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_messages_proto_rawDescData)
	})
	return file_messages_proto_rawDescData
}

var file_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_messages_proto_goTypes = []interface{}{
	(*SearchEvent)(nil),         // 0: pb.SearchEvent
	(*BadSearchEvent)(nil),      // 1: pb.BadSearchEvent
	(*AddNewSearchRecord)(nil),  // 2: pb.AddNewSearchRecord
	(*AddNewSynonym)(nil),       // 3: pb.AddNewSynonym
	(*timestamp.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_messages_proto_depIdxs = []int32{
	4, // 0: pb.SearchEvent.eventTimestamp:type_name -> google.protobuf.Timestamp
	4, // 1: pb.BadSearchEvent.eventTimestamp:type_name -> google.protobuf.Timestamp
	4, // 2: pb.AddNewSearchRecord.eventTimestamp:type_name -> google.protobuf.Timestamp
	4, // 3: pb.AddNewSynonym.eventTimestamp:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_messages_proto_init() }
func file_messages_proto_init() {
	if File_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchEvent); i {
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
		file_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BadSearchEvent); i {
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
		file_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddNewSearchRecord); i {
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
		file_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddNewSynonym); i {
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
			RawDescriptor: file_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messages_proto_goTypes,
		DependencyIndexes: file_messages_proto_depIdxs,
		MessageInfos:      file_messages_proto_msgTypes,
	}.Build()
	File_messages_proto = out.File
	file_messages_proto_rawDesc = nil
	file_messages_proto_goTypes = nil
	file_messages_proto_depIdxs = nil
}
