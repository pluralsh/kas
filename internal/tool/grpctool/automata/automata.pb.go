// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: internal/tool/grpctool/automata/automata.proto

package automata

import (
	reflect "reflect"

	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

var file_internal_tool_grpctool_automata_automata_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: ([]int32)(nil),
		Field:         92000,
		Name:          "gitlab.agent.grpctool.automata.next_allowed_field",
		Tag:           "varint,92000,rep,name=next_allowed_field",
		Filename:      "internal/tool/grpctool/automata/automata.proto",
	},
	{
		ExtendedType:  (*descriptor.OneofOptions)(nil),
		ExtensionType: ([]int32)(nil),
		Field:         92001,
		Name:          "gitlab.agent.grpctool.automata.first_allowed_field",
		Tag:           "varint,92001,rep,name=first_allowed_field",
		Filename:      "internal/tool/grpctool/automata/automata.proto",
	},
}

// Extension fields to descriptor.FieldOptions.
var (
	// repeated int32 next_allowed_field = 92000;
	E_NextAllowedField = &file_internal_tool_grpctool_automata_automata_proto_extTypes[0]
)

// Extension fields to descriptor.OneofOptions.
var (
	// repeated int32 first_allowed_field = 92001;
	E_FirstAllowedField = &file_internal_tool_grpctool_automata_automata_proto_extTypes[1]
)

var File_internal_tool_grpctool_automata_automata_proto protoreflect.FileDescriptor

var file_internal_tool_grpctool_automata_automata_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x6f, 0x6f, 0x6c, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x74, 0x6f, 0x6f, 0x6c, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x61, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x74, 0x6f, 0x6f, 0x6c, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x61,
	0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x3a, 0x4d, 0x0a, 0x12, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x64, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe0, 0xce, 0x05, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x10, 0x6e, 0x65, 0x78, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x3a, 0x4f, 0x0a, 0x13, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x64, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4f, 0x6e, 0x65, 0x6f, 0x66,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe1, 0xce, 0x05, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x11, 0x66, 0x69, 0x72, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x42, 0x58, 0x5a, 0x56, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2d, 0x6f, 0x72, 0x67, 0x2f, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x6f, 0x6f, 0x6c, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x74,
	0x6f, 0x6f, 0x6c, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_internal_tool_grpctool_automata_automata_proto_goTypes = []interface{}{
	(*descriptor.FieldOptions)(nil), // 0: google.protobuf.FieldOptions
	(*descriptor.OneofOptions)(nil), // 1: google.protobuf.OneofOptions
}
var file_internal_tool_grpctool_automata_automata_proto_depIdxs = []int32{
	0, // 0: gitlab.agent.grpctool.automata.next_allowed_field:extendee -> google.protobuf.FieldOptions
	1, // 1: gitlab.agent.grpctool.automata.first_allowed_field:extendee -> google.protobuf.OneofOptions
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	0, // [0:2] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_tool_grpctool_automata_automata_proto_init() }
func file_internal_tool_grpctool_automata_automata_proto_init() {
	if File_internal_tool_grpctool_automata_automata_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_tool_grpctool_automata_automata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_internal_tool_grpctool_automata_automata_proto_goTypes,
		DependencyIndexes: file_internal_tool_grpctool_automata_automata_proto_depIdxs,
		ExtensionInfos:    file_internal_tool_grpctool_automata_automata_proto_extTypes,
	}.Build()
	File_internal_tool_grpctool_automata_automata_proto = out.File
	file_internal_tool_grpctool_automata_automata_proto_rawDesc = nil
	file_internal_tool_grpctool_automata_automata_proto_goTypes = nil
	file_internal_tool_grpctool_automata_automata_proto_depIdxs = nil
}
