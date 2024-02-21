// ライセンスはいつか書いておく

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.2
// source: proto/k8s_controller_response_v1.proto

package v1

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

type CreateNamespaceReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateNamespaceReply) Reset() {
	*x = CreateNamespaceReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_k8s_controller_response_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNamespaceReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNamespaceReply) ProtoMessage() {}

func (x *CreateNamespaceReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_k8s_controller_response_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNamespaceReply.ProtoReflect.Descriptor instead.
func (*CreateNamespaceReply) Descriptor() ([]byte, []int) {
	return file_proto_k8s_controller_response_v1_proto_rawDescGZIP(), []int{0}
}

func (x *CreateNamespaceReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListNamespacesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Names []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
}

func (x *ListNamespacesReply) Reset() {
	*x = ListNamespacesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_k8s_controller_response_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNamespacesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNamespacesReply) ProtoMessage() {}

func (x *ListNamespacesReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_k8s_controller_response_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNamespacesReply.ProtoReflect.Descriptor instead.
func (*ListNamespacesReply) Descriptor() ([]byte, []int) {
	return file_proto_k8s_controller_response_v1_proto_rawDescGZIP(), []int{1}
}

func (x *ListNamespacesReply) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

type CreateDeploymentReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Version   string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *CreateDeploymentReply) Reset() {
	*x = CreateDeploymentReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_k8s_controller_response_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDeploymentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDeploymentReply) ProtoMessage() {}

func (x *CreateDeploymentReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_k8s_controller_response_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDeploymentReply.ProtoReflect.Descriptor instead.
func (*CreateDeploymentReply) Descriptor() ([]byte, []int) {
	return file_proto_k8s_controller_response_v1_proto_rawDescGZIP(), []int{2}
}

func (x *CreateDeploymentReply) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *CreateDeploymentReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateDeploymentReply) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type Repository struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	PullRequestId string `protobuf:"bytes,2,opt,name=pull_request_id,json=pullRequestId,proto3" json:"pull_request_id,omitempty"`
	Branch        string `protobuf:"bytes,3,opt,name=branch,proto3" json:"branch,omitempty"`
}

func (x *Repository) Reset() {
	*x = Repository{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_k8s_controller_response_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Repository) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Repository) ProtoMessage() {}

func (x *Repository) ProtoReflect() protoreflect.Message {
	mi := &file_proto_k8s_controller_response_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Repository.ProtoReflect.Descriptor instead.
func (*Repository) Descriptor() ([]byte, []int) {
	return file_proto_k8s_controller_response_v1_proto_rawDescGZIP(), []int{3}
}

func (x *Repository) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Repository) GetPullRequestId() string {
	if x != nil {
		return x.PullRequestId
	}
	return ""
}

func (x *Repository) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

type Application struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Status  string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Image   string `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	Age     string `protobuf:"bytes,5,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *Application) Reset() {
	*x = Application{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_k8s_controller_response_v1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Application) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Application) ProtoMessage() {}

func (x *Application) ProtoReflect() protoreflect.Message {
	mi := &file_proto_k8s_controller_response_v1_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Application.ProtoReflect.Descriptor instead.
func (*Application) Descriptor() ([]byte, []int) {
	return file_proto_k8s_controller_response_v1_proto_rawDescGZIP(), []int{4}
}

func (x *Application) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Application) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Application) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Application) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Application) GetAge() string {
	if x != nil {
		return x.Age
	}
	return ""
}

type GetOrgReposReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Organization string         `protobuf:"bytes,1,opt,name=organization,proto3" json:"organization,omitempty"`
	Repositories []*Repository  `protobuf:"bytes,2,rep,name=repositories,proto3" json:"repositories,omitempty"`
	Applications []*Application `protobuf:"bytes,3,rep,name=applications,proto3" json:"applications,omitempty"`
}

func (x *GetOrgReposReply) Reset() {
	*x = GetOrgReposReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_k8s_controller_response_v1_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrgReposReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrgReposReply) ProtoMessage() {}

func (x *GetOrgReposReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_k8s_controller_response_v1_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrgReposReply.ProtoReflect.Descriptor instead.
func (*GetOrgReposReply) Descriptor() ([]byte, []int) {
	return file_proto_k8s_controller_response_v1_proto_rawDescGZIP(), []int{5}
}

func (x *GetOrgReposReply) GetOrganization() string {
	if x != nil {
		return x.Organization
	}
	return ""
}

func (x *GetOrgReposReply) GetRepositories() []*Repository {
	if x != nil {
		return x.Repositories
	}
	return nil
}

func (x *GetOrgReposReply) GetApplications() []*Application {
	if x != nil {
		return x.Applications
	}
	return nil
}

type GetAllsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrganizationInfos []*GetOrgReposReply `protobuf:"bytes,1,rep,name=organization_infos,json=organizationInfos,proto3" json:"organization_infos,omitempty"`
}

func (x *GetAllsReply) Reset() {
	*x = GetAllsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_k8s_controller_response_v1_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllsReply) ProtoMessage() {}

func (x *GetAllsReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_k8s_controller_response_v1_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllsReply.ProtoReflect.Descriptor instead.
func (*GetAllsReply) Descriptor() ([]byte, []int) {
	return file_proto_k8s_controller_response_v1_proto_rawDescGZIP(), []int{6}
}

func (x *GetAllsReply) GetOrganizationInfos() []*GetOrgReposReply {
	if x != nil {
		return x.OrganizationInfos
	}
	return nil
}

type Usage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContainerName string `protobuf:"bytes,1,opt,name=container_name,json=containerName,proto3" json:"container_name,omitempty"`
	Cpu           int64  `protobuf:"varint,2,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Mem           int64  `protobuf:"varint,3,opt,name=mem,proto3" json:"mem,omitempty"`
	Storage       int64  `protobuf:"varint,4,opt,name=storage,proto3" json:"storage,omitempty"`
}

func (x *Usage) Reset() {
	*x = Usage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_k8s_controller_response_v1_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Usage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Usage) ProtoMessage() {}

func (x *Usage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_k8s_controller_response_v1_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Usage.ProtoReflect.Descriptor instead.
func (*Usage) Descriptor() ([]byte, []int) {
	return file_proto_k8s_controller_response_v1_proto_rawDescGZIP(), []int{7}
}

func (x *Usage) GetContainerName() string {
	if x != nil {
		return x.ContainerName
	}
	return ""
}

func (x *Usage) GetCpu() int64 {
	if x != nil {
		return x.Cpu
	}
	return 0
}

func (x *Usage) GetMem() int64 {
	if x != nil {
		return x.Mem
	}
	return 0
}

func (x *Usage) GetStorage() int64 {
	if x != nil {
		return x.Storage
	}
	return 0
}

type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppName       string   `protobuf:"bytes,1,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`
	PodName       string   `protobuf:"bytes,2,opt,name=pod_name,json=podName,proto3" json:"pod_name,omitempty"`
	PullRequestId string   `protobuf:"bytes,3,opt,name=pull_request_id,json=pullRequestId,proto3" json:"pull_request_id,omitempty"`
	Branch        string   `protobuf:"bytes,4,opt,name=branch,proto3" json:"branch,omitempty"`
	Usages        []*Usage `protobuf:"bytes,5,rep,name=usages,proto3" json:"usages,omitempty"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_k8s_controller_response_v1_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_proto_k8s_controller_response_v1_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_proto_k8s_controller_response_v1_proto_rawDescGZIP(), []int{8}
}

func (x *Resource) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *Resource) GetPodName() string {
	if x != nil {
		return x.PodName
	}
	return ""
}

func (x *Resource) GetPullRequestId() string {
	if x != nil {
		return x.PullRequestId
	}
	return ""
}

func (x *Resource) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

func (x *Resource) GetUsages() []*Usage {
	if x != nil {
		return x.Usages
	}
	return nil
}

type GetResourceReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resources []*Resource `protobuf:"bytes,1,rep,name=resources,proto3" json:"resources,omitempty"`
	IsDisable bool        `protobuf:"varint,2,opt,name=is_disable,json=isDisable,proto3" json:"is_disable,omitempty"`
}

func (x *GetResourceReply) Reset() {
	*x = GetResourceReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_k8s_controller_response_v1_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResourceReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResourceReply) ProtoMessage() {}

func (x *GetResourceReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_k8s_controller_response_v1_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResourceReply.ProtoReflect.Descriptor instead.
func (*GetResourceReply) Descriptor() ([]byte, []int) {
	return file_proto_k8s_controller_response_v1_proto_rawDescGZIP(), []int{9}
}

func (x *GetResourceReply) GetResources() []*Resource {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *GetResourceReply) GetIsDisable() bool {
	if x != nil {
		return x.IsDisable
	}
	return false
}

type NodePort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target int32 `protobuf:"varint,1,opt,name=target,proto3" json:"target,omitempty"`
	Port   int32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *NodePort) Reset() {
	*x = NodePort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_k8s_controller_response_v1_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodePort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodePort) ProtoMessage() {}

func (x *NodePort) ProtoReflect() protoreflect.Message {
	mi := &file_proto_k8s_controller_response_v1_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodePort.ProtoReflect.Descriptor instead.
func (*NodePort) Descriptor() ([]byte, []int) {
	return file_proto_k8s_controller_response_v1_proto_rawDescGZIP(), []int{10}
}

func (x *NodePort) GetTarget() int32 {
	if x != nil {
		return x.Target
	}
	return 0
}

func (x *NodePort) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type CreatePreviewReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string      `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name      string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Version   string      `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	NodePorts []*NodePort `protobuf:"bytes,4,rep,name=node_ports,json=nodePorts,proto3" json:"node_ports,omitempty"`
}

func (x *CreatePreviewReply) Reset() {
	*x = CreatePreviewReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_k8s_controller_response_v1_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePreviewReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePreviewReply) ProtoMessage() {}

func (x *CreatePreviewReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_k8s_controller_response_v1_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePreviewReply.ProtoReflect.Descriptor instead.
func (*CreatePreviewReply) Descriptor() ([]byte, []int) {
	return file_proto_k8s_controller_response_v1_proto_rawDescGZIP(), []int{11}
}

func (x *CreatePreviewReply) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *CreatePreviewReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreatePreviewReply) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *CreatePreviewReply) GetNodePorts() []*NodePort {
	if x != nil {
		return x.NodePorts
	}
	return nil
}

type BuildImageReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image string `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *BuildImageReply) Reset() {
	*x = BuildImageReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_k8s_controller_response_v1_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildImageReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildImageReply) ProtoMessage() {}

func (x *BuildImageReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_k8s_controller_response_v1_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildImageReply.ProtoReflect.Descriptor instead.
func (*BuildImageReply) Descriptor() ([]byte, []int) {
	return file_proto_k8s_controller_response_v1_proto_rawDescGZIP(), []int{12}
}

func (x *BuildImageReply) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

var File_proto_k8s_controller_response_v1_proto protoreflect.FileDescriptor

var file_proto_k8s_controller_response_v1_proto_rawDesc = []byte{
	0x0a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6b, 0x38, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f,
	0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x67, 0x61, 0x6e, 0x74, 0x72, 0x79,
	0x63, 0x64, 0x2e, 0x6b, 0x38, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x22, 0x2a, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x2b, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x63, 0x0a,
	0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x22, 0x60, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x75, 0x6c, 0x6c, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70,
	0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x22, 0x7b, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x67,
	0x65, 0x22, 0xcf, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x67, 0x52, 0x65, 0x70, 0x6f,
	0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4a, 0x0a, 0x0c, 0x72, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x63, 0x64, 0x2e, 0x6b, 0x38, 0x73, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x0c, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x4b, 0x0a, 0x0c, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x67,
	0x61, 0x6e, 0x74, 0x72, 0x79, 0x63, 0x64, 0x2e, 0x6b, 0x38, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0x6b, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x5b, 0x0a, 0x12, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2c, 0x2e, 0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x63, 0x64, 0x2e, 0x6b, 0x38, 0x73, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x4f, 0x72, 0x67, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x11, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x73,
	0x22, 0x6c, 0x0a, 0x05, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x63, 0x70, 0x75, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x63,
	0x70, 0x75, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x65, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x6d, 0x65, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x22, 0xbb,
	0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x61,
	0x70, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6f, 0x64, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6f, 0x64, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x75, 0x6c, 0x6c, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x75, 0x6c, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x12, 0x39, 0x0a, 0x06, 0x75, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x63, 0x64, 0x2e, 0x6b, 0x38, 0x73,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x06, 0x75, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x75, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x42, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x63, 0x64, 0x2e, 0x6b,
	0x38, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x64, 0x69, 0x73, 0x61, 0x62,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x44, 0x69, 0x73, 0x61,
	0x62, 0x6c, 0x65, 0x22, 0x36, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0xa5, 0x01, 0x0a, 0x12,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x43,
	0x0a, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x63, 0x64, 0x2e, 0x6b, 0x38,
	0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x4e, 0x6f, 0x64, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x50, 0x6f,
	0x72, 0x74, 0x73, 0x22, 0x27, 0x0a, 0x0f, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x3c, 0x5a, 0x3a,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x6e, 0x74, 0x72,
	0x79, 0x63, 0x64, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x6b, 0x38, 0x73, 0x2d, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_k8s_controller_response_v1_proto_rawDescOnce sync.Once
	file_proto_k8s_controller_response_v1_proto_rawDescData = file_proto_k8s_controller_response_v1_proto_rawDesc
)

func file_proto_k8s_controller_response_v1_proto_rawDescGZIP() []byte {
	file_proto_k8s_controller_response_v1_proto_rawDescOnce.Do(func() {
		file_proto_k8s_controller_response_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_k8s_controller_response_v1_proto_rawDescData)
	})
	return file_proto_k8s_controller_response_v1_proto_rawDescData
}

var file_proto_k8s_controller_response_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_proto_k8s_controller_response_v1_proto_goTypes = []interface{}{
	(*CreateNamespaceReply)(nil),  // 0: gantrycd.k8s_controller.v1.CreateNamespaceReply
	(*ListNamespacesReply)(nil),   // 1: gantrycd.k8s_controller.v1.ListNamespacesReply
	(*CreateDeploymentReply)(nil), // 2: gantrycd.k8s_controller.v1.CreateDeploymentReply
	(*Repository)(nil),            // 3: gantrycd.k8s_controller.v1.repository
	(*Application)(nil),           // 4: gantrycd.k8s_controller.v1.application
	(*GetOrgReposReply)(nil),      // 5: gantrycd.k8s_controller.v1.GetOrgReposReply
	(*GetAllsReply)(nil),          // 6: gantrycd.k8s_controller.v1.GetAllsReply
	(*Usage)(nil),                 // 7: gantrycd.k8s_controller.v1.Usage
	(*Resource)(nil),              // 8: gantrycd.k8s_controller.v1.Resource
	(*GetResourceReply)(nil),      // 9: gantrycd.k8s_controller.v1.GetResourceReply
	(*NodePort)(nil),              // 10: gantrycd.k8s_controller.v1.NodePort
	(*CreatePreviewReply)(nil),    // 11: gantrycd.k8s_controller.v1.CreatePreviewReply
	(*BuildImageReply)(nil),       // 12: gantrycd.k8s_controller.v1.BuildImageReply
}
var file_proto_k8s_controller_response_v1_proto_depIdxs = []int32{
	3,  // 0: gantrycd.k8s_controller.v1.GetOrgReposReply.repositories:type_name -> gantrycd.k8s_controller.v1.repository
	4,  // 1: gantrycd.k8s_controller.v1.GetOrgReposReply.applications:type_name -> gantrycd.k8s_controller.v1.application
	5,  // 2: gantrycd.k8s_controller.v1.GetAllsReply.organization_infos:type_name -> gantrycd.k8s_controller.v1.GetOrgReposReply
	7,  // 3: gantrycd.k8s_controller.v1.Resource.usages:type_name -> gantrycd.k8s_controller.v1.Usage
	8,  // 4: gantrycd.k8s_controller.v1.GetResourceReply.resources:type_name -> gantrycd.k8s_controller.v1.Resource
	10, // 5: gantrycd.k8s_controller.v1.CreatePreviewReply.node_ports:type_name -> gantrycd.k8s_controller.v1.NodePort
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_k8s_controller_response_v1_proto_init() }
func file_proto_k8s_controller_response_v1_proto_init() {
	if File_proto_k8s_controller_response_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_k8s_controller_response_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNamespaceReply); i {
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
		file_proto_k8s_controller_response_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNamespacesReply); i {
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
		file_proto_k8s_controller_response_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDeploymentReply); i {
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
		file_proto_k8s_controller_response_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Repository); i {
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
		file_proto_k8s_controller_response_v1_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Application); i {
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
		file_proto_k8s_controller_response_v1_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrgReposReply); i {
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
		file_proto_k8s_controller_response_v1_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllsReply); i {
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
		file_proto_k8s_controller_response_v1_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Usage); i {
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
		file_proto_k8s_controller_response_v1_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
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
		file_proto_k8s_controller_response_v1_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResourceReply); i {
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
		file_proto_k8s_controller_response_v1_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodePort); i {
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
		file_proto_k8s_controller_response_v1_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePreviewReply); i {
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
		file_proto_k8s_controller_response_v1_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildImageReply); i {
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
			RawDescriptor: file_proto_k8s_controller_response_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_k8s_controller_response_v1_proto_goTypes,
		DependencyIndexes: file_proto_k8s_controller_response_v1_proto_depIdxs,
		MessageInfos:      file_proto_k8s_controller_response_v1_proto_msgTypes,
	}.Build()
	File_proto_k8s_controller_response_v1_proto = out.File
	file_proto_k8s_controller_response_v1_proto_rawDesc = nil
	file_proto_k8s_controller_response_v1_proto_goTypes = nil
	file_proto_k8s_controller_response_v1_proto_depIdxs = nil
}
