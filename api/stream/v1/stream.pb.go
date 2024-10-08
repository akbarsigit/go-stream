// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: stream/v1/stream.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListStreamsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*ListStreamsResponse_Stream `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *ListStreamsResponse) Reset() {
	*x = ListStreamsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_v1_stream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStreamsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStreamsResponse) ProtoMessage() {}

func (x *ListStreamsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stream_v1_stream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStreamsResponse.ProtoReflect.Descriptor instead.
func (*ListStreamsResponse) Descriptor() ([]byte, []int) {
	return file_stream_v1_stream_proto_rawDescGZIP(), []int{0}
}

func (x *ListStreamsResponse) GetResults() []*ListStreamsResponse_Stream {
	if x != nil {
		return x.Results
	}
	return nil
}

type ListStreamsResponse_Stream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *ListStreamsResponse_Stream) Reset() {
	*x = ListStreamsResponse_Stream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_v1_stream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStreamsResponse_Stream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStreamsResponse_Stream) ProtoMessage() {}

func (x *ListStreamsResponse_Stream) ProtoReflect() protoreflect.Message {
	mi := &file_stream_v1_stream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStreamsResponse_Stream.ProtoReflect.Descriptor instead.
func (*ListStreamsResponse_Stream) Descriptor() ([]byte, []int) {
	return file_stream_v1_stream_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ListStreamsResponse_Stream) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ListStreamsResponse_Stream) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListStreamsResponse_Stream) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

var File_stream_v1_stream_proto protoreflect.FileDescriptor

var file_stream_v1_stream_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa0, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x1a, 0x48, 0x0a, 0x06, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x32, 0x58, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1e, 0x2e, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0e, 0x5a,
	0x0c, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stream_v1_stream_proto_rawDescOnce sync.Once
	file_stream_v1_stream_proto_rawDescData = file_stream_v1_stream_proto_rawDesc
)

func file_stream_v1_stream_proto_rawDescGZIP() []byte {
	file_stream_v1_stream_proto_rawDescOnce.Do(func() {
		file_stream_v1_stream_proto_rawDescData = protoimpl.X.CompressGZIP(file_stream_v1_stream_proto_rawDescData)
	})
	return file_stream_v1_stream_proto_rawDescData
}

var file_stream_v1_stream_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_stream_v1_stream_proto_goTypes = []interface{}{
	(*ListStreamsResponse)(nil),        // 0: stream.v1.ListStreamsResponse
	(*ListStreamsResponse_Stream)(nil), // 1: stream.v1.ListStreamsResponse.Stream
	(*emptypb.Empty)(nil),              // 2: google.protobuf.Empty
}
var file_stream_v1_stream_proto_depIdxs = []int32{
	1, // 0: stream.v1.ListStreamsResponse.results:type_name -> stream.v1.ListStreamsResponse.Stream
	2, // 1: stream.v1.StreamService.ListStreams:input_type -> google.protobuf.Empty
	0, // 2: stream.v1.StreamService.ListStreams:output_type -> stream.v1.ListStreamsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_stream_v1_stream_proto_init() }
func file_stream_v1_stream_proto_init() {
	if File_stream_v1_stream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stream_v1_stream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStreamsResponse); i {
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
		file_stream_v1_stream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStreamsResponse_Stream); i {
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
			RawDescriptor: file_stream_v1_stream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stream_v1_stream_proto_goTypes,
		DependencyIndexes: file_stream_v1_stream_proto_depIdxs,
		MessageInfos:      file_stream_v1_stream_proto_msgTypes,
	}.Build()
	File_stream_v1_stream_proto = out.File
	file_stream_v1_stream_proto_rawDesc = nil
	file_stream_v1_stream_proto_goTypes = nil
	file_stream_v1_stream_proto_depIdxs = nil
}
