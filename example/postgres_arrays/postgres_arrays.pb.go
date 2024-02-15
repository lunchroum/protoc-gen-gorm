// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: postgres_arrays/postgres_arrays.proto

package postgres_arrays

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

type Example struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id for example
	Id             string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description    string    `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	ArrayOfBools   []bool    `protobuf:"varint,20,rep,packed,name=array_of_bools,json=arrayOfBools,proto3" json:"array_of_bools,omitempty"`
	ArrayOfFloat64 []float64 `protobuf:"fixed64,30,rep,packed,name=array_of_float64,json=arrayOfFloat64,proto3" json:"array_of_float64,omitempty"`
	ArrayOfInt64   []int64   `protobuf:"varint,40,rep,packed,name=array_of_int64,json=arrayOfInt64,proto3" json:"array_of_int64,omitempty"`
	ArrayOfString  []string  `protobuf:"bytes,50,rep,name=array_of_string,json=arrayOfString,proto3" json:"array_of_string,omitempty"`
}

func (x *Example) Reset() {
	*x = Example{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postgres_arrays_postgres_arrays_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Example) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Example) ProtoMessage() {}

func (x *Example) ProtoReflect() protoreflect.Message {
	mi := &file_postgres_arrays_postgres_arrays_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Example.ProtoReflect.Descriptor instead.
func (*Example) Descriptor() ([]byte, []int) {
	return file_postgres_arrays_postgres_arrays_proto_rawDescGZIP(), []int{0}
}

func (x *Example) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Example) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Example) GetArrayOfBools() []bool {
	if x != nil {
		return x.ArrayOfBools
	}
	return nil
}

func (x *Example) GetArrayOfFloat64() []float64 {
	if x != nil {
		return x.ArrayOfFloat64
	}
	return nil
}

func (x *Example) GetArrayOfInt64() []int64 {
	if x != nil {
		return x.ArrayOfInt64
	}
	return nil
}

func (x *Example) GetArrayOfString() []string {
	if x != nil {
		return x.ArrayOfString
	}
	return nil
}

var File_postgres_arrays_postgres_arrays_proto protoreflect.FileDescriptor

var file_postgres_arrays_postgres_arrays_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x72, 0x72, 0x61, 0x79,
	0x73, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x72, 0x72, 0x61, 0x79,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65,
	0x73, 0x2e, 0x61, 0x72, 0x72, 0x61, 0x79, 0x73, 0x1a, 0x12, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x67, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf1, 0x01, 0x0a,
	0x07, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0xba, 0xb9, 0x19, 0x0a, 0x0a, 0x08, 0x12, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x28, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0e, 0x61, 0x72,
	0x72, 0x61, 0x79, 0x5f, 0x6f, 0x66, 0x5f, 0x62, 0x6f, 0x6f, 0x6c, 0x73, 0x18, 0x14, 0x20, 0x03,
	0x28, 0x08, 0x52, 0x0c, 0x61, 0x72, 0x72, 0x61, 0x79, 0x4f, 0x66, 0x42, 0x6f, 0x6f, 0x6c, 0x73,
	0x12, 0x28, 0x0a, 0x10, 0x61, 0x72, 0x72, 0x61, 0x79, 0x5f, 0x6f, 0x66, 0x5f, 0x66, 0x6c, 0x6f,
	0x61, 0x74, 0x36, 0x34, 0x18, 0x1e, 0x20, 0x03, 0x28, 0x01, 0x52, 0x0e, 0x61, 0x72, 0x72, 0x61,
	0x79, 0x4f, 0x66, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x36, 0x34, 0x12, 0x24, 0x0a, 0x0e, 0x61, 0x72,
	0x72, 0x61, 0x79, 0x5f, 0x6f, 0x66, 0x5f, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x28, 0x20, 0x03,
	0x28, 0x03, 0x52, 0x0c, 0x61, 0x72, 0x72, 0x61, 0x79, 0x4f, 0x66, 0x49, 0x6e, 0x74, 0x36, 0x34,
	0x12, 0x26, 0x0a, 0x0f, 0x61, 0x72, 0x72, 0x61, 0x79, 0x5f, 0x6f, 0x66, 0x5f, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x18, 0x32, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x72, 0x72, 0x61, 0x79,
	0x4f, 0x66, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x3a, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x08, 0x01,
	0x42, 0x4b, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b,
	0x6c, 0x6f, 0x73, 0x6d, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x67, 0x6f, 0x72, 0x6d, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x70, 0x6f,
	0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x72, 0x72, 0x61, 0x79, 0x73, 0x3b, 0x70, 0x6f,
	0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x72, 0x72, 0x61, 0x79, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_postgres_arrays_postgres_arrays_proto_rawDescOnce sync.Once
	file_postgres_arrays_postgres_arrays_proto_rawDescData = file_postgres_arrays_postgres_arrays_proto_rawDesc
)

func file_postgres_arrays_postgres_arrays_proto_rawDescGZIP() []byte {
	file_postgres_arrays_postgres_arrays_proto_rawDescOnce.Do(func() {
		file_postgres_arrays_postgres_arrays_proto_rawDescData = protoimpl.X.CompressGZIP(file_postgres_arrays_postgres_arrays_proto_rawDescData)
	})
	return file_postgres_arrays_postgres_arrays_proto_rawDescData
}

var file_postgres_arrays_postgres_arrays_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_postgres_arrays_postgres_arrays_proto_goTypes = []interface{}{
	(*Example)(nil), // 0: postgres.arrays.Example
}
var file_postgres_arrays_postgres_arrays_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_postgres_arrays_postgres_arrays_proto_init() }
func file_postgres_arrays_postgres_arrays_proto_init() {
	if File_postgres_arrays_postgres_arrays_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_postgres_arrays_postgres_arrays_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Example); i {
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
			RawDescriptor: file_postgres_arrays_postgres_arrays_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_postgres_arrays_postgres_arrays_proto_goTypes,
		DependencyIndexes: file_postgres_arrays_postgres_arrays_proto_depIdxs,
		MessageInfos:      file_postgres_arrays_postgres_arrays_proto_msgTypes,
	}.Build()
	File_postgres_arrays_postgres_arrays_proto = out.File
	file_postgres_arrays_postgres_arrays_proto_rawDesc = nil
	file_postgres_arrays_postgres_arrays_proto_goTypes = nil
	file_postgres_arrays_postgres_arrays_proto_depIdxs = nil
}
