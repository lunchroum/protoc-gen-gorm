// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: feature_demo/demo_multi_file.proto

package example

import (
	_ "github.com/lunchroum/protoc-gen-gorm/options"
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

type ExternalChild struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ExternalChild) Reset() {
	*x = ExternalChild{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feature_demo_demo_multi_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExternalChild) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExternalChild) ProtoMessage() {}

func (x *ExternalChild) ProtoReflect() protoreflect.Message {
	mi := &file_feature_demo_demo_multi_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExternalChild.ProtoReflect.Descriptor instead.
func (*ExternalChild) Descriptor() ([]byte, []int) {
	return file_feature_demo_demo_multi_file_proto_rawDescGZIP(), []int{0}
}

func (x *ExternalChild) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type BlogPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title  string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Author string `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
}

func (x *BlogPost) Reset() {
	*x = BlogPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feature_demo_demo_multi_file_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlogPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlogPost) ProtoMessage() {}

func (x *BlogPost) ProtoReflect() protoreflect.Message {
	mi := &file_feature_demo_demo_multi_file_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlogPost.ProtoReflect.Descriptor instead.
func (*BlogPost) Descriptor() ([]byte, []int) {
	return file_feature_demo_demo_multi_file_proto_rawDescGZIP(), []int{1}
}

func (x *BlogPost) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BlogPost) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *BlogPost) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

var File_feature_demo_demo_multi_file_proto protoreflect.FileDescriptor

var file_feature_demo_demo_multi_file_proto_rawDesc = []byte{
	0x0a, 0x22, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x64,
	0x65, 0x6d, 0x6f, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x1a, 0x12, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x27, 0x0a, 0x0d, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x68, 0x69,
	0x6c, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x3a, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x08, 0x01, 0x22, 0x50, 0x0a, 0x08, 0x42, 0x6c,
	0x6f, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x3a, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x08, 0x01, 0x42, 0x40, 0x5a, 0x3e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6c, 0x6f, 0x73, 0x6d,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x72,
	0x6d, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x5f, 0x64, 0x65, 0x6d, 0x6f, 0x3b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_feature_demo_demo_multi_file_proto_rawDescOnce sync.Once
	file_feature_demo_demo_multi_file_proto_rawDescData = file_feature_demo_demo_multi_file_proto_rawDesc
)

func file_feature_demo_demo_multi_file_proto_rawDescGZIP() []byte {
	file_feature_demo_demo_multi_file_proto_rawDescOnce.Do(func() {
		file_feature_demo_demo_multi_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_feature_demo_demo_multi_file_proto_rawDescData)
	})
	return file_feature_demo_demo_multi_file_proto_rawDescData
}

var file_feature_demo_demo_multi_file_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_feature_demo_demo_multi_file_proto_goTypes = []interface{}{
	(*ExternalChild)(nil), // 0: example.ExternalChild
	(*BlogPost)(nil),      // 1: example.BlogPost
}
var file_feature_demo_demo_multi_file_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_feature_demo_demo_multi_file_proto_init() }
func file_feature_demo_demo_multi_file_proto_init() {
	if File_feature_demo_demo_multi_file_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_feature_demo_demo_multi_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExternalChild); i {
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
		file_feature_demo_demo_multi_file_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlogPost); i {
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
			RawDescriptor: file_feature_demo_demo_multi_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_feature_demo_demo_multi_file_proto_goTypes,
		DependencyIndexes: file_feature_demo_demo_multi_file_proto_depIdxs,
		MessageInfos:      file_feature_demo_demo_multi_file_proto_msgTypes,
	}.Build()
	File_feature_demo_demo_multi_file_proto = out.File
	file_feature_demo_demo_multi_file_proto_rawDesc = nil
	file_feature_demo_demo_multi_file_proto_goTypes = nil
	file_feature_demo_demo_multi_file_proto_depIdxs = nil
}
