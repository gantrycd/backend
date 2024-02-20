// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.2
// source: proto/bff_v1.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetOrgResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrgRepository []*OrgRepository `protobuf:"bytes,1,rep,name=orgRepository,proto3" json:"orgRepository,omitempty"`
}

func (x *GetOrgResponse) Reset() {
	*x = GetOrgResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bff_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrgResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrgResponse) ProtoMessage() {}

func (x *GetOrgResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bff_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrgResponse.ProtoReflect.Descriptor instead.
func (*GetOrgResponse) Descriptor() ([]byte, []int) {
	return file_proto_bff_v1_proto_rawDescGZIP(), []int{0}
}

func (x *GetOrgResponse) GetOrgRepository() []*OrgRepository {
	if x != nil {
		return x.OrgRepository
	}
	return nil
}

type OrgRepository struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Organization string   `protobuf:"bytes,1,opt,name=organization,proto3" json:"organization,omitempty"`
	Repository   []string `protobuf:"bytes,2,rep,name=repository,proto3" json:"repository,omitempty"`
}

func (x *OrgRepository) Reset() {
	*x = OrgRepository{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bff_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrgRepository) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrgRepository) ProtoMessage() {}

func (x *OrgRepository) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bff_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrgRepository.ProtoReflect.Descriptor instead.
func (*OrgRepository) Descriptor() ([]byte, []int) {
	return file_proto_bff_v1_proto_rawDescGZIP(), []int{1}
}

func (x *OrgRepository) GetOrganization() string {
	if x != nil {
		return x.Organization
	}
	return ""
}

func (x *OrgRepository) GetRepository() []string {
	if x != nil {
		return x.Repository
	}
	return nil
}

var File_proto_bff_v1_proto protoreflect.FileDescriptor

var file_proto_bff_v1_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x66, 0x66, 0x5f, 0x76, 0x31, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x63, 0x64, 0x2e, 0x62,
	0x66, 0x66, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x56, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x44, 0x0a, 0x0d, 0x6f, 0x72, 0x67, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x61, 0x6e, 0x74,
	0x72, 0x79, 0x63, 0x64, 0x2e, 0x62, 0x66, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x67, 0x52,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x0d, 0x6f, 0x72, 0x67, 0x52, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x22, 0x53, 0x0a, 0x0d, 0x4f, 0x72, 0x67, 0x52,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a,
	0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x32, 0x5e, 0x0a,
	0x0a, 0x42, 0x66, 0x66, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x06, 0x47,
	0x65, 0x74, 0x4f, 0x72, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1f, 0x2e,
	0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x63, 0x64, 0x2e, 0x62, 0x66, 0x66, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x4f, 0x72, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0d,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x07, 0x12, 0x05, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x42, 0x2e, 0x5a,
	0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x6e, 0x74,
	0x72, 0x79, 0x63, 0x64, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x66, 0x66, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_bff_v1_proto_rawDescOnce sync.Once
	file_proto_bff_v1_proto_rawDescData = file_proto_bff_v1_proto_rawDesc
)

func file_proto_bff_v1_proto_rawDescGZIP() []byte {
	file_proto_bff_v1_proto_rawDescOnce.Do(func() {
		file_proto_bff_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_bff_v1_proto_rawDescData)
	})
	return file_proto_bff_v1_proto_rawDescData
}

var file_proto_bff_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_bff_v1_proto_goTypes = []interface{}{
	(*GetOrgResponse)(nil), // 0: gantrycd.bff.v1.GetOrgResponse
	(*OrgRepository)(nil),  // 1: gantrycd.bff.v1.OrgRepository
	(*emptypb.Empty)(nil),  // 2: google.protobuf.Empty
}
var file_proto_bff_v1_proto_depIdxs = []int32{
	1, // 0: gantrycd.bff.v1.GetOrgResponse.orgRepository:type_name -> gantrycd.bff.v1.OrgRepository
	2, // 1: gantrycd.bff.v1.BffService.GetOrg:input_type -> google.protobuf.Empty
	0, // 2: gantrycd.bff.v1.BffService.GetOrg:output_type -> gantrycd.bff.v1.GetOrgResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_bff_v1_proto_init() }
func file_proto_bff_v1_proto_init() {
	if File_proto_bff_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_bff_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrgResponse); i {
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
		file_proto_bff_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrgRepository); i {
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
			RawDescriptor: file_proto_bff_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_bff_v1_proto_goTypes,
		DependencyIndexes: file_proto_bff_v1_proto_depIdxs,
		MessageInfos:      file_proto_bff_v1_proto_msgTypes,
	}.Build()
	File_proto_bff_v1_proto = out.File
	file_proto_bff_v1_proto_rawDesc = nil
	file_proto_bff_v1_proto_goTypes = nil
	file_proto_bff_v1_proto_depIdxs = nil
}
