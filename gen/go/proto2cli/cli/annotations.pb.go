// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: proto2cli/cli/annotations.proto

package annotations

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_proto2cli_cli_annotations_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         1071,
		Name:          "proto2cli.cli.SvcDesc",
		Tag:           "bytes,1071,opt,name=SvcDesc",
		Filename:      "proto2cli/cli/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         1071,
		Name:          "proto2cli.cli.MsgDesc",
		Tag:           "bytes,1071,opt,name=MsgDesc",
		Filename:      "proto2cli/cli/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         1071,
		Name:          "proto2cli.cli.FlagName",
		Tag:           "bytes,1071,opt,name=FlagName",
		Filename:      "proto2cli/cli/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         1072,
		Name:          "proto2cli.cli.ShortFlagName",
		Tag:           "bytes,1072,opt,name=ShortFlagName",
		Filename:      "proto2cli/cli/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         1073,
		Name:          "proto2cli.cli.FieldDesc",
		Tag:           "bytes,1073,opt,name=FieldDesc",
		Filename:      "proto2cli/cli/annotations.proto",
	},
}

// Extension fields to descriptorpb.ServiceOptions.
var (
	// Override any CLI help text for a particular message
	//
	// optional string SvcDesc = 1071;
	E_SvcDesc = &file_proto2cli_cli_annotations_proto_extTypes[0]
)

// Extension fields to descriptorpb.MessageOptions.
var (
	// Override any CLI help text for a particular message
	//
	// optional string MsgDesc = 1071;
	E_MsgDesc = &file_proto2cli_cli_annotations_proto_extTypes[1]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// An optional flag to be used for this field.  Eg "--my-flag-name"
	// If a flag is not provided than the fqn of the field is used
	//
	// optional string FlagName = 1071;
	E_FlagName = &file_proto2cli_cli_annotations_proto_extTypes[2]
	// Optional short flag to be used for this command - eg "-m"
	//
	// optional string ShortFlagName = 1072;
	E_ShortFlagName = &file_proto2cli_cli_annotations_proto_extTypes[3]
	// Override any CLI help text for a particular field.
	// Otherwise the documentation comment is used.
	//
	// optional string FieldDesc = 1073;
	E_FieldDesc = &file_proto2cli_cli_annotations_proto_extTypes[4]
)

var File_proto2cli_cli_annotations_proto protoreflect.FileDescriptor

var file_proto2cli_cli_annotations_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x63, 0x6c, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x63, 0x6c, 0x69, 0x2e, 0x63, 0x6c, 0x69,
	0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x3a, 0x3a, 0x0a, 0x07, 0x53, 0x76, 0x63, 0x44, 0x65, 0x73, 0x63, 0x12, 0x1f, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xaf,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x76, 0x63, 0x44, 0x65, 0x73, 0x63, 0x3a, 0x3a,
	0x0a, 0x07, 0x4d, 0x73, 0x67, 0x44, 0x65, 0x73, 0x63, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xaf, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x4d, 0x73, 0x67, 0x44, 0x65, 0x73, 0x63, 0x3a, 0x3a, 0x0a, 0x08, 0x46, 0x6c,
	0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xaf, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x6c,
	0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x3a, 0x44, 0x0a, 0x0d, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x46,
	0x6c, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb0, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x46, 0x6c, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x3a, 0x3c, 0x0a, 0x09,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x44, 0x65, 0x73, 0x63, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb1, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x44, 0x65, 0x73, 0x63, 0x42, 0xc4, 0x01, 0x0a, 0x11, 0x63,
	0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x63, 0x6c, 0x69, 0x2e, 0x63, 0x6c, 0x69,
	0x42, 0x10, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x61, 0x6e, 0x79, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x63, 0x6c,
	0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x63,
	0x6c, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x3b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0xa2, 0x02,
	0x03, 0x50, 0x43, 0x58, 0xaa, 0x02, 0x0d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x63, 0x6c, 0x69,
	0x2e, 0x43, 0x6c, 0x69, 0xca, 0x02, 0x0d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x63, 0x6c, 0x69,
	0x5c, 0x43, 0x6c, 0x69, 0xe2, 0x02, 0x19, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x63, 0x6c, 0x69,
	0x5c, 0x43, 0x6c, 0x69, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x0e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x63, 0x6c, 0x69, 0x3a, 0x3a, 0x43, 0x6c,
	0x69,
}

var file_proto2cli_cli_annotations_proto_goTypes = []interface{}{
	(*descriptorpb.ServiceOptions)(nil), // 0: google.protobuf.ServiceOptions
	(*descriptorpb.MessageOptions)(nil), // 1: google.protobuf.MessageOptions
	(*descriptorpb.FieldOptions)(nil),   // 2: google.protobuf.FieldOptions
}
var file_proto2cli_cli_annotations_proto_depIdxs = []int32{
	0, // 0: proto2cli.cli.SvcDesc:extendee -> google.protobuf.ServiceOptions
	1, // 1: proto2cli.cli.MsgDesc:extendee -> google.protobuf.MessageOptions
	2, // 2: proto2cli.cli.FlagName:extendee -> google.protobuf.FieldOptions
	2, // 3: proto2cli.cli.ShortFlagName:extendee -> google.protobuf.FieldOptions
	2, // 4: proto2cli.cli.FieldDesc:extendee -> google.protobuf.FieldOptions
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	0, // [0:5] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto2cli_cli_annotations_proto_init() }
func file_proto2cli_cli_annotations_proto_init() {
	if File_proto2cli_cli_annotations_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto2cli_cli_annotations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 5,
			NumServices:   0,
		},
		GoTypes:           file_proto2cli_cli_annotations_proto_goTypes,
		DependencyIndexes: file_proto2cli_cli_annotations_proto_depIdxs,
		ExtensionInfos:    file_proto2cli_cli_annotations_proto_extTypes,
	}.Build()
	File_proto2cli_cli_annotations_proto = out.File
	file_proto2cli_cli_annotations_proto_rawDesc = nil
	file_proto2cli_cli_annotations_proto_goTypes = nil
	file_proto2cli_cli_annotations_proto_depIdxs = nil
}