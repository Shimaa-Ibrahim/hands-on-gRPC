// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: bidirectional-stream-api/maximumpb/maximum.proto

package maximumpb

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

type MaximumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int64 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *MaximumRequest) Reset() {
	*x = MaximumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bidirectional_stream_api_maximumpb_maximum_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaximumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaximumRequest) ProtoMessage() {}

func (x *MaximumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bidirectional_stream_api_maximumpb_maximum_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaximumRequest.ProtoReflect.Descriptor instead.
func (*MaximumRequest) Descriptor() ([]byte, []int) {
	return file_bidirectional_stream_api_maximumpb_maximum_proto_rawDescGZIP(), []int{0}
}

func (x *MaximumRequest) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

type MaximumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaximumNumber int64 `protobuf:"varint,1,opt,name=maximum_number,json=maximumNumber,proto3" json:"maximum_number,omitempty"`
}

func (x *MaximumResponse) Reset() {
	*x = MaximumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bidirectional_stream_api_maximumpb_maximum_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaximumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaximumResponse) ProtoMessage() {}

func (x *MaximumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bidirectional_stream_api_maximumpb_maximum_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaximumResponse.ProtoReflect.Descriptor instead.
func (*MaximumResponse) Descriptor() ([]byte, []int) {
	return file_bidirectional_stream_api_maximumpb_maximum_proto_rawDescGZIP(), []int{1}
}

func (x *MaximumResponse) GetMaximumNumber() int64 {
	if x != nil {
		return x.MaximumNumber
	}
	return 0
}

var File_bidirectional_stream_api_maximumpb_maximum_proto protoreflect.FileDescriptor

var file_bidirectional_stream_api_maximumpb_maximum_proto_rawDesc = []byte{
	0x0a, 0x30, 0x62, 0x69, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2d,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x78, 0x69, 0x6d,
	0x75, 0x6d, 0x70, 0x62, 0x2f, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x22, 0x28, 0x0a, 0x0e, 0x4d,
	0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x38, 0x0a, 0x0f, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x61, 0x78, 0x69,
	0x6d, 0x75, 0x6d, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0d, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x32,
	0x54, 0x0a, 0x0e, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x42, 0x0a, 0x07, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x12, 0x17, 0x2e, 0x6d,
	0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x2e, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x2e,
	0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0b, 0x5a, 0x09, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bidirectional_stream_api_maximumpb_maximum_proto_rawDescOnce sync.Once
	file_bidirectional_stream_api_maximumpb_maximum_proto_rawDescData = file_bidirectional_stream_api_maximumpb_maximum_proto_rawDesc
)

func file_bidirectional_stream_api_maximumpb_maximum_proto_rawDescGZIP() []byte {
	file_bidirectional_stream_api_maximumpb_maximum_proto_rawDescOnce.Do(func() {
		file_bidirectional_stream_api_maximumpb_maximum_proto_rawDescData = protoimpl.X.CompressGZIP(file_bidirectional_stream_api_maximumpb_maximum_proto_rawDescData)
	})
	return file_bidirectional_stream_api_maximumpb_maximum_proto_rawDescData
}

var file_bidirectional_stream_api_maximumpb_maximum_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_bidirectional_stream_api_maximumpb_maximum_proto_goTypes = []interface{}{
	(*MaximumRequest)(nil),  // 0: maximum.MaximumRequest
	(*MaximumResponse)(nil), // 1: maximum.MaximumResponse
}
var file_bidirectional_stream_api_maximumpb_maximum_proto_depIdxs = []int32{
	0, // 0: maximum.MaximumService.Maximum:input_type -> maximum.MaximumRequest
	1, // 1: maximum.MaximumService.Maximum:output_type -> maximum.MaximumResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bidirectional_stream_api_maximumpb_maximum_proto_init() }
func file_bidirectional_stream_api_maximumpb_maximum_proto_init() {
	if File_bidirectional_stream_api_maximumpb_maximum_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bidirectional_stream_api_maximumpb_maximum_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaximumRequest); i {
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
		file_bidirectional_stream_api_maximumpb_maximum_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaximumResponse); i {
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
			RawDescriptor: file_bidirectional_stream_api_maximumpb_maximum_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bidirectional_stream_api_maximumpb_maximum_proto_goTypes,
		DependencyIndexes: file_bidirectional_stream_api_maximumpb_maximum_proto_depIdxs,
		MessageInfos:      file_bidirectional_stream_api_maximumpb_maximum_proto_msgTypes,
	}.Build()
	File_bidirectional_stream_api_maximumpb_maximum_proto = out.File
	file_bidirectional_stream_api_maximumpb_maximum_proto_rawDesc = nil
	file_bidirectional_stream_api_maximumpb_maximum_proto_goTypes = nil
	file_bidirectional_stream_api_maximumpb_maximum_proto_depIdxs = nil
}
