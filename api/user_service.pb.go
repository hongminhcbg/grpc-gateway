// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: user_service.proto

package api

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_user_service_proto protoreflect.FileDescriptor

var file_user_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x68, 0x6f, 0x6e, 0x67, 0x6d, 0x69, 0x6e, 0x68, 0x63, 0x62,
	0x67, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61,
	0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x10, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0x99, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x89, 0x01, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x2f, 0x2e, 0x68, 0x6f, 0x6e, 0x67, 0x6d, 0x69, 0x6e, 0x68, 0x63, 0x62, 0x67, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x30, 0x2e, 0x68, 0x6f, 0x6e, 0x67, 0x6d, 0x69, 0x6e, 0x68, 0x63, 0x62, 0x67,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x3a, 0x01, 0x2a, 0x42, 0xeb,
	0x01, 0x0a, 0x20, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x6f, 0x6e, 0x67, 0x6d, 0x69, 0x6e, 0x68, 0x63,
	0x62, 0x67, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x61, 0x70, 0x69, 0x42, 0x10, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x6f, 0x6e, 0x67, 0x6d, 0x69, 0x6e, 0x68, 0x63, 0x62, 0x67, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x61, 0x70, 0x69,
	0xa2, 0x02, 0x03, 0x48, 0x47, 0x41, 0xaa, 0x02, 0x1b, 0x48, 0x6f, 0x6e, 0x67, 0x6d, 0x69, 0x6e,
	0x68, 0x63, 0x62, 0x67, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x41, 0x70, 0x69, 0xca, 0x02, 0x1b, 0x48, 0x6f, 0x6e, 0x67, 0x6d, 0x69, 0x6e, 0x68, 0x63,
	0x62, 0x67, 0x5c, 0x47, 0x72, 0x70, 0x63, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5c, 0x41,
	0x70, 0x69, 0xe2, 0x02, 0x27, 0x48, 0x6f, 0x6e, 0x67, 0x6d, 0x69, 0x6e, 0x68, 0x63, 0x62, 0x67,
	0x5c, 0x47, 0x72, 0x70, 0x63, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5c, 0x41, 0x70, 0x69,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1d, 0x48,
	0x6f, 0x6e, 0x67, 0x6d, 0x69, 0x6e, 0x68, 0x63, 0x62, 0x67, 0x3a, 0x3a, 0x47, 0x72, 0x70, 0x63,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_user_service_proto_goTypes = []interface{}{
	(*CreateUserRequest)(nil),  // 0: hongminhcbg.grpc_gateway.api.CreateUserRequest
	(*CreateUserResponse)(nil), // 1: hongminhcbg.grpc_gateway.api.CreateUserResponse
}
var file_user_service_proto_depIdxs = []int32{
	0, // 0: hongminhcbg.grpc_gateway.api.UserService.CreateUser:input_type -> hongminhcbg.grpc_gateway.api.CreateUserRequest
	1, // 1: hongminhcbg.grpc_gateway.api.UserService.CreateUser:output_type -> hongminhcbg.grpc_gateway.api.CreateUserResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_service_proto_init() }
func file_user_service_proto_init() {
	if File_user_service_proto != nil {
		return
	}
	file_user_model_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_service_proto_goTypes,
		DependencyIndexes: file_user_service_proto_depIdxs,
	}.Build()
	File_user_service_proto = out.File
	file_user_service_proto_rawDesc = nil
	file_user_service_proto_goTypes = nil
	file_user_service_proto_depIdxs = nil
}
