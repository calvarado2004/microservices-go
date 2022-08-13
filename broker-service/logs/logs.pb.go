// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: logs.proto

package logs

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

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Data string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_logs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_logs_proto_rawDescGZIP(), []int{0}
}

func (x *Log) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Log) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type LogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogEntry *Log `protobuf:"bytes,1,opt,name=logEntry,proto3" json:"logEntry,omitempty"`
}

func (x *LogRequest) Reset() {
	*x = LogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogRequest) ProtoMessage() {}

func (x *LogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_logs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogRequest.ProtoReflect.Descriptor instead.
func (*LogRequest) Descriptor() ([]byte, []int) {
	return file_logs_proto_rawDescGZIP(), []int{1}
}

func (x *LogRequest) GetLogEntry() *Log {
	if x != nil {
		return x.LogEntry
	}
	return nil
}

type LogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *LogResponse) Reset() {
	*x = LogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogResponse) ProtoMessage() {}

func (x *LogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_logs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogResponse.ProtoReflect.Descriptor instead.
func (*LogResponse) Descriptor() ([]byte, []int) {
	return file_logs_proto_rawDescGZIP(), []int{2}
}

func (x *LogResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_logs_proto protoreflect.FileDescriptor

var file_logs_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6c, 0x6f,
	0x67, 0x73, 0x22, 0x2d, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x33, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x25, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x09, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x08, 0x6c, 0x6f,
	0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x25, 0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x3d, 0x0a,
	0x0a, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x08, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x12, 0x10, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x4c,
	0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x6c, 0x6f, 0x67, 0x73,
	0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x07, 0x5a, 0x05,
	0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_logs_proto_rawDescOnce sync.Once
	file_logs_proto_rawDescData = file_logs_proto_rawDesc
)

func file_logs_proto_rawDescGZIP() []byte {
	file_logs_proto_rawDescOnce.Do(func() {
		file_logs_proto_rawDescData = protoimpl.X.CompressGZIP(file_logs_proto_rawDescData)
	})
	return file_logs_proto_rawDescData
}

var file_logs_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_logs_proto_goTypes = []interface{}{
	(*Log)(nil),         // 0: logs.Log
	(*LogRequest)(nil),  // 1: logs.LogRequest
	(*LogResponse)(nil), // 2: logs.LogResponse
}
var file_logs_proto_depIdxs = []int32{
	0, // 0: logs.LogRequest.logEntry:type_name -> logs.Log
	1, // 1: logs.LogService.WriteLog:input_type -> logs.LogRequest
	2, // 2: logs.LogService.WriteLog:output_type -> logs.LogResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_logs_proto_init() }
func file_logs_proto_init() {
	if File_logs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_logs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
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
		file_logs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogRequest); i {
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
		file_logs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogResponse); i {
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
			RawDescriptor: file_logs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_logs_proto_goTypes,
		DependencyIndexes: file_logs_proto_depIdxs,
		MessageInfos:      file_logs_proto_msgTypes,
	}.Build()
	File_logs_proto = out.File
	file_logs_proto_rawDesc = nil
	file_logs_proto_goTypes = nil
	file_logs_proto_depIdxs = nil
}
