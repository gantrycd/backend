// ライセンスはいつか書いておく

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.2
// source: proto/k8s_controller_service_v1.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_proto_k8s_controller_service_v1_proto protoreflect.FileDescriptor

var file_proto_k8s_controller_service_v1_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6b, 0x38, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x76,
	0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x63,
	0x64, 0x2e, 0x6b, 0x38, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x25, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6b, 0x38, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x76,
	0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6b,
	0x38, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0x91, 0x0a, 0x0a, 0x13, 0x4b, 0x38, 0x73, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x77, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x32, 0x2e, 0x67, 0x61, 0x6e,
	0x74, 0x72, 0x79, 0x63, 0x64, 0x2e, 0x6b, 0x38, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30,
	0x2e, 0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x63, 0x64, 0x2e, 0x6b, 0x38, 0x73, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x59, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2f, 0x2e, 0x67, 0x61, 0x6e,
	0x74, 0x72, 0x79, 0x63, 0x64, 0x2e, 0x6b, 0x38, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x5d, 0x0a, 0x0f, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x32,
	0x2e, 0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x63, 0x64, 0x2e, 0x6b, 0x38, 0x73, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x79, 0x0a, 0x0f, 0x41, 0x70,
	0x70, 0x6c, 0x79, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x33, 0x2e,
	0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x63, 0x64, 0x2e, 0x6b, 0x38, 0x73, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x31, 0x2e, 0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x63, 0x64, 0x2e, 0x6b, 0x38,
	0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x5f, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x33, 0x2e, 0x67, 0x61, 0x6e, 0x74,
	0x72, 0x79, 0x63, 0x64, 0x2e, 0x6b, 0x38, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4b, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x28, 0x2e, 0x67, 0x61, 0x6e, 0x74,
	0x72, 0x79, 0x63, 0x64, 0x2e, 0x6b, 0x38, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x6a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x67, 0x52, 0x65, 0x70,
	0x6f, 0x73, 0x12, 0x2d, 0x2e, 0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x63, 0x64, 0x2e, 0x6b, 0x38,
	0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x4f, 0x72, 0x67, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2c, 0x2e, 0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x63, 0x64, 0x2e, 0x6b, 0x38, 0x73,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x4f, 0x72, 0x67, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x71, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x12, 0x30, 0x2e, 0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x63, 0x64, 0x2e, 0x6b, 0x38, 0x73, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x63, 0x64, 0x2e, 0x6b, 0x38,
	0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x71, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x12, 0x30, 0x2e, 0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x63, 0x64, 0x2e, 0x6b,
	0x38, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x63, 0x64,
	0x2e, 0x6b, 0x38, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x59, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x30, 0x2e, 0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x63,
	0x64, 0x2e, 0x6b, 0x38, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x83, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x49, 0x6e,
	0x66, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x2e, 0x67, 0x61, 0x6e, 0x74, 0x72,
	0x79, 0x63, 0x64, 0x2e, 0x6b, 0x38, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x49,
	0x6e, 0x66, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x34, 0x2e, 0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x63, 0x64, 0x2e, 0x6b, 0x38, 0x73, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x6b, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x2e, 0x2e, 0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x63, 0x64,
	0x2e, 0x6b, 0x38, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x63, 0x64,
	0x2e, 0x6b, 0x38, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x63, 0x64, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x2f,
	0x6b, 0x38, 0x73, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_proto_k8s_controller_service_v1_proto_goTypes = []interface{}{
	(*CreateNamespaceRequest)(nil),     // 0: gantrycd.k8s_controller.v1.CreateNamespaceRequest
	(*emptypb.Empty)(nil),              // 1: google.protobuf.Empty
	(*DeleteNamespaceRequest)(nil),     // 2: gantrycd.k8s_controller.v1.DeleteNamespaceRequest
	(*CreateDeploymentRequest)(nil),    // 3: gantrycd.k8s_controller.v1.CreateDeploymentRequest
	(*DeleteDeploymentRequest)(nil),    // 4: gantrycd.k8s_controller.v1.DeleteDeploymentRequest
	(*GetOrgRepoRequest)(nil),          // 5: gantrycd.k8s_controller.v1.GetOrgRepoRequest
	(*CreatePreviewRequest)(nil),       // 6: gantrycd.k8s_controller.v1.CreatePreviewRequest
	(*DeletePreviewRequest)(nil),       // 7: gantrycd.k8s_controller.v1.DeletePreviewRequest
	(*GetDeployInfomationRequest)(nil), // 8: gantrycd.k8s_controller.v1.GetDeployInfomationRequest
	(*GetResourceRequest)(nil),         // 9: gantrycd.k8s_controller.v1.GetResourceRequest
	(*CreateNamespaceReply)(nil),       // 10: gantrycd.k8s_controller.v1.CreateNamespaceReply
	(*ListNamespacesReply)(nil),        // 11: gantrycd.k8s_controller.v1.ListNamespacesReply
	(*CreateDeploymentReply)(nil),      // 12: gantrycd.k8s_controller.v1.CreateDeploymentReply
	(*GetAllsReply)(nil),               // 13: gantrycd.k8s_controller.v1.GetAllsReply
	(*GetOrgReposReply)(nil),           // 14: gantrycd.k8s_controller.v1.GetOrgReposReply
	(*CreatePreviewReply)(nil),         // 15: gantrycd.k8s_controller.v1.CreatePreviewReply
	(*GetDeployInfomationReply)(nil),   // 16: gantrycd.k8s_controller.v1.GetDeployInfomationReply
	(*GetResourceReply)(nil),           // 17: gantrycd.k8s_controller.v1.GetResourceReply
}
var file_proto_k8s_controller_service_v1_proto_depIdxs = []int32{
	0,  // 0: gantrycd.k8s_controller.v1.K8sCustomController.CreateNamespace:input_type -> gantrycd.k8s_controller.v1.CreateNamespaceRequest
	1,  // 1: gantrycd.k8s_controller.v1.K8sCustomController.ListNamespaces:input_type -> google.protobuf.Empty
	2,  // 2: gantrycd.k8s_controller.v1.K8sCustomController.DeleteNamespace:input_type -> gantrycd.k8s_controller.v1.DeleteNamespaceRequest
	3,  // 3: gantrycd.k8s_controller.v1.K8sCustomController.ApplyDeployment:input_type -> gantrycd.k8s_controller.v1.CreateDeploymentRequest
	4,  // 4: gantrycd.k8s_controller.v1.K8sCustomController.DeleteDeployment:input_type -> gantrycd.k8s_controller.v1.DeleteDeploymentRequest
	1,  // 5: gantrycd.k8s_controller.v1.K8sCustomController.GetAlls:input_type -> google.protobuf.Empty
	5,  // 6: gantrycd.k8s_controller.v1.K8sCustomController.GetOrgRepos:input_type -> gantrycd.k8s_controller.v1.GetOrgRepoRequest
	6,  // 7: gantrycd.k8s_controller.v1.K8sCustomController.CreatePreview:input_type -> gantrycd.k8s_controller.v1.CreatePreviewRequest
	6,  // 8: gantrycd.k8s_controller.v1.K8sCustomController.UpdatePreview:input_type -> gantrycd.k8s_controller.v1.CreatePreviewRequest
	7,  // 9: gantrycd.k8s_controller.v1.K8sCustomController.DeletePreview:input_type -> gantrycd.k8s_controller.v1.DeletePreviewRequest
	8,  // 10: gantrycd.k8s_controller.v1.K8sCustomController.GetDeployInfomation:input_type -> gantrycd.k8s_controller.v1.GetDeployInfomationRequest
	9,  // 11: gantrycd.k8s_controller.v1.K8sCustomController.GetResource:input_type -> gantrycd.k8s_controller.v1.GetResourceRequest
	10, // 12: gantrycd.k8s_controller.v1.K8sCustomController.CreateNamespace:output_type -> gantrycd.k8s_controller.v1.CreateNamespaceReply
	11, // 13: gantrycd.k8s_controller.v1.K8sCustomController.ListNamespaces:output_type -> gantrycd.k8s_controller.v1.ListNamespacesReply
	1,  // 14: gantrycd.k8s_controller.v1.K8sCustomController.DeleteNamespace:output_type -> google.protobuf.Empty
	12, // 15: gantrycd.k8s_controller.v1.K8sCustomController.ApplyDeployment:output_type -> gantrycd.k8s_controller.v1.CreateDeploymentReply
	1,  // 16: gantrycd.k8s_controller.v1.K8sCustomController.DeleteDeployment:output_type -> google.protobuf.Empty
	13, // 17: gantrycd.k8s_controller.v1.K8sCustomController.GetAlls:output_type -> gantrycd.k8s_controller.v1.GetAllsReply
	14, // 18: gantrycd.k8s_controller.v1.K8sCustomController.GetOrgRepos:output_type -> gantrycd.k8s_controller.v1.GetOrgReposReply
	15, // 19: gantrycd.k8s_controller.v1.K8sCustomController.CreatePreview:output_type -> gantrycd.k8s_controller.v1.CreatePreviewReply
	15, // 20: gantrycd.k8s_controller.v1.K8sCustomController.UpdatePreview:output_type -> gantrycd.k8s_controller.v1.CreatePreviewReply
	1,  // 21: gantrycd.k8s_controller.v1.K8sCustomController.DeletePreview:output_type -> google.protobuf.Empty
	16, // 22: gantrycd.k8s_controller.v1.K8sCustomController.GetDeployInfomation:output_type -> gantrycd.k8s_controller.v1.GetDeployInfomationReply
	17, // 23: gantrycd.k8s_controller.v1.K8sCustomController.GetResource:output_type -> gantrycd.k8s_controller.v1.GetResourceReply
	12, // [12:24] is the sub-list for method output_type
	0,  // [0:12] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_proto_k8s_controller_service_v1_proto_init() }
func file_proto_k8s_controller_service_v1_proto_init() {
	if File_proto_k8s_controller_service_v1_proto != nil {
		return
	}
	file_proto_k8s_controller_request_v1_proto_init()
	file_proto_k8s_controller_response_v1_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_k8s_controller_service_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_k8s_controller_service_v1_proto_goTypes,
		DependencyIndexes: file_proto_k8s_controller_service_v1_proto_depIdxs,
	}.Build()
	File_proto_k8s_controller_service_v1_proto = out.File
	file_proto_k8s_controller_service_v1_proto_rawDesc = nil
	file_proto_k8s_controller_service_v1_proto_goTypes = nil
	file_proto_k8s_controller_service_v1_proto_depIdxs = nil
}
