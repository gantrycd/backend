// ライセンスはいつか書いておく

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: proto/k8s_controller_request_v1.proto

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

type CreateNamespaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateNamespaceRequest) Reset() {
	*x = CreateNamespaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_k8s_controller_request_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNamespaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNamespaceRequest) ProtoMessage() {}

func (x *CreateNamespaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_k8s_controller_request_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNamespaceRequest.ProtoReflect.Descriptor instead.
func (*CreateNamespaceRequest) Descriptor() ([]byte, []int) {
	return file_proto_k8s_controller_request_v1_proto_rawDescGZIP(), []int{0}
}

func (x *CreateNamespaceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeleteNamespaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteNamespaceRequest) Reset() {
	*x = DeleteNamespaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_k8s_controller_request_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNamespaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNamespaceRequest) ProtoMessage() {}

func (x *DeleteNamespaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_k8s_controller_request_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNamespaceRequest.ProtoReflect.Descriptor instead.
func (*DeleteNamespaceRequest) Descriptor() ([]byte, []int) {
	return file_proto_k8s_controller_request_v1_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteNamespaceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateDeploymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace  string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Repository string `protobuf:"bytes,2,opt,name=repository,proto3" json:"repository,omitempty"`
	PrNumber   string `protobuf:"bytes,3,opt,name=pr_number,json=prNumber,proto3" json:"pr_number,omitempty"`
	Image      string `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	Branch     string `protobuf:"bytes,5,opt,name=branch,proto3" json:"branch,omitempty"`
	AppName    string `protobuf:"bytes,6,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`
	Replicas   string `protobuf:"bytes,7,opt,name=replicas,proto3" json:"replicas,omitempty"`
	CreatedBy  string `protobuf:"bytes,8,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
}

func (x *CreateDeploymentRequest) Reset() {
	*x = CreateDeploymentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_k8s_controller_request_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDeploymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDeploymentRequest) ProtoMessage() {}

func (x *CreateDeploymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_k8s_controller_request_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDeploymentRequest.ProtoReflect.Descriptor instead.
func (*CreateDeploymentRequest) Descriptor() ([]byte, []int) {
	return file_proto_k8s_controller_request_v1_proto_rawDescGZIP(), []int{2}
}

func (x *CreateDeploymentRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *CreateDeploymentRequest) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

func (x *CreateDeploymentRequest) GetPrNumber() string {
	if x != nil {
		return x.PrNumber
	}
	return ""
}

func (x *CreateDeploymentRequest) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *CreateDeploymentRequest) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

func (x *CreateDeploymentRequest) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *CreateDeploymentRequest) GetReplicas() string {
	if x != nil {
		return x.Replicas
	}
	return ""
}

func (x *CreateDeploymentRequest) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

type DeleteDeploymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace  string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Repository string `protobuf:"bytes,2,opt,name=repository,proto3" json:"repository,omitempty"`
	PrNumber   string `protobuf:"bytes,3,opt,name=pr_number,json=prNumber,proto3" json:"pr_number,omitempty"`
}

func (x *DeleteDeploymentRequest) Reset() {
	*x = DeleteDeploymentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_k8s_controller_request_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDeploymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDeploymentRequest) ProtoMessage() {}

func (x *DeleteDeploymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_k8s_controller_request_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDeploymentRequest.ProtoReflect.Descriptor instead.
func (*DeleteDeploymentRequest) Descriptor() ([]byte, []int) {
	return file_proto_k8s_controller_request_v1_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteDeploymentRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *DeleteDeploymentRequest) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

func (x *DeleteDeploymentRequest) GetPrNumber() string {
	if x != nil {
		return x.PrNumber
	}
	return ""
}

type GetOrgRepoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Organization string `protobuf:"bytes,1,opt,name=organization,proto3" json:"organization,omitempty"`
}

func (x *GetOrgRepoRequest) Reset() {
	*x = GetOrgRepoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_k8s_controller_request_v1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrgRepoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrgRepoRequest) ProtoMessage() {}

func (x *GetOrgRepoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_k8s_controller_request_v1_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrgRepoRequest.ProtoReflect.Descriptor instead.
func (*GetOrgRepoRequest) Descriptor() ([]byte, []int) {
	return file_proto_k8s_controller_request_v1_proto_rawDescGZIP(), []int{4}
}

func (x *GetOrgRepoRequest) GetOrganization() string {
	if x != nil {
		return x.Organization
	}
	return ""
}

type GetResourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Organization string `protobuf:"bytes,1,opt,name=organization,proto3" json:"organization,omitempty"`
	Repository   string `protobuf:"bytes,2,opt,name=repository,proto3" json:"repository,omitempty"`
}

func (x *GetResourceRequest) Reset() {
	*x = GetResourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_k8s_controller_request_v1_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResourceRequest) ProtoMessage() {}

func (x *GetResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_k8s_controller_request_v1_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResourceRequest.ProtoReflect.Descriptor instead.
func (*GetResourceRequest) Descriptor() ([]byte, []int) {
	return file_proto_k8s_controller_request_v1_proto_rawDescGZIP(), []int{5}
}

func (x *GetResourceRequest) GetOrganization() string {
	if x != nil {
		return x.Organization
	}
	return ""
}

func (x *GetResourceRequest) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

type GetRepoBranchesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Organization string `protobuf:"bytes,1,opt,name=organization,proto3" json:"organization,omitempty"`
	Repository   string `protobuf:"bytes,2,opt,name=repository,proto3" json:"repository,omitempty"`
}

func (x *GetRepoBranchesRequest) Reset() {
	*x = GetRepoBranchesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_k8s_controller_request_v1_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRepoBranchesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRepoBranchesRequest) ProtoMessage() {}

func (x *GetRepoBranchesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_k8s_controller_request_v1_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRepoBranchesRequest.ProtoReflect.Descriptor instead.
func (*GetRepoBranchesRequest) Descriptor() ([]byte, []int) {
	return file_proto_k8s_controller_request_v1_proto_rawDescGZIP(), []int{6}
}

func (x *GetRepoBranchesRequest) GetOrganization() string {
	if x != nil {
		return x.Organization
	}
	return ""
}

func (x *GetRepoBranchesRequest) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

type DeletePreviewRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Organization  string `protobuf:"bytes,1,opt,name=organization,proto3" json:"organization,omitempty"`
	Repository    string `protobuf:"bytes,2,opt,name=repository,proto3" json:"repository,omitempty"`
	PullRequestId string `protobuf:"bytes,3,opt,name=pull_request_id,json=pullRequestId,proto3" json:"pull_request_id,omitempty"`
	Branch        string `protobuf:"bytes,4,opt,name=branch,proto3" json:"branch,omitempty"`
}

func (x *DeletePreviewRequest) Reset() {
	*x = DeletePreviewRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_k8s_controller_request_v1_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePreviewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePreviewRequest) ProtoMessage() {}

func (x *DeletePreviewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_k8s_controller_request_v1_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePreviewRequest.ProtoReflect.Descriptor instead.
func (*DeletePreviewRequest) Descriptor() ([]byte, []int) {
	return file_proto_k8s_controller_request_v1_proto_rawDescGZIP(), []int{7}
}

func (x *DeletePreviewRequest) GetOrganization() string {
	if x != nil {
		return x.Organization
	}
	return ""
}

func (x *DeletePreviewRequest) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

func (x *DeletePreviewRequest) GetPullRequestId() string {
	if x != nil {
		return x.PullRequestId
	}
	return ""
}

func (x *DeletePreviewRequest) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_k8s_controller_request_v1_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_proto_k8s_controller_request_v1_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_proto_k8s_controller_request_v1_proto_rawDescGZIP(), []int{8}
}

func (x *Config) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Config) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type CreatePreviewRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Organization    string    `protobuf:"bytes,1,opt,name=organization,proto3" json:"organization,omitempty"`
	Repository      string    `protobuf:"bytes,2,opt,name=repository,proto3" json:"repository,omitempty"`
	PullRequestId   string    `protobuf:"bytes,3,opt,name=pull_request_id,json=pullRequestId,proto3" json:"pull_request_id,omitempty"`
	Branch          string    `protobuf:"bytes,4,opt,name=branch,proto3" json:"branch,omitempty"`
	Image           string    `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
	ImagePullPolicy string    `protobuf:"bytes,6,opt,name=image_pull_policy,json=imagePullPolicy,proto3" json:"image_pull_policy,omitempty"`
	Replicas        int32     `protobuf:"varint,7,opt,name=replicas,proto3" json:"replicas,omitempty"`
	Configs         []*Config `protobuf:"bytes,8,rep,name=configs,proto3" json:"configs,omitempty"`
	ExposePorts     []int32   `protobuf:"varint,9,rep,packed,name=expose_ports,json=exposePorts,proto3" json:"expose_ports,omitempty"`
}

func (x *CreatePreviewRequest) Reset() {
	*x = CreatePreviewRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_k8s_controller_request_v1_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePreviewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePreviewRequest) ProtoMessage() {}

func (x *CreatePreviewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_k8s_controller_request_v1_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePreviewRequest.ProtoReflect.Descriptor instead.
func (*CreatePreviewRequest) Descriptor() ([]byte, []int) {
	return file_proto_k8s_controller_request_v1_proto_rawDescGZIP(), []int{9}
}

func (x *CreatePreviewRequest) GetOrganization() string {
	if x != nil {
		return x.Organization
	}
	return ""
}

func (x *CreatePreviewRequest) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

func (x *CreatePreviewRequest) GetPullRequestId() string {
	if x != nil {
		return x.PullRequestId
	}
	return ""
}

func (x *CreatePreviewRequest) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

func (x *CreatePreviewRequest) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *CreatePreviewRequest) GetImagePullPolicy() string {
	if x != nil {
		return x.ImagePullPolicy
	}
	return ""
}

func (x *CreatePreviewRequest) GetReplicas() int32 {
	if x != nil {
		return x.Replicas
	}
	return 0
}

func (x *CreatePreviewRequest) GetConfigs() []*Config {
	if x != nil {
		return x.Configs
	}
	return nil
}

func (x *CreatePreviewRequest) GetExposePorts() []int32 {
	if x != nil {
		return x.ExposePorts
	}
	return nil
}

type BuildImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace      string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Repository     string `protobuf:"bytes,2,opt,name=repository,proto3" json:"repository,omitempty"`
	Branch         string `protobuf:"bytes,3,opt,name=branch,proto3" json:"branch,omitempty"`
	PullRequestId  string `protobuf:"bytes,4,opt,name=pull_request_id,json=pullRequestId,proto3" json:"pull_request_id,omitempty"`
	GitRepo        string `protobuf:"bytes,5,opt,name=git_repo,json=gitRepo,proto3" json:"git_repo,omitempty"`
	DockerfileDir  string `protobuf:"bytes,6,opt,name=dockerfile_dir,json=dockerfileDir,proto3" json:"dockerfile_dir,omitempty"`
	DockerfilePath string `protobuf:"bytes,7,opt,name=dockerfile_path,json=dockerfilePath,proto3" json:"dockerfile_path,omitempty"`
	ImageName      string `protobuf:"bytes,8,opt,name=image_name,json=imageName,proto3" json:"image_name,omitempty"`
	Token          string `protobuf:"bytes,9,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *BuildImageRequest) Reset() {
	*x = BuildImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_k8s_controller_request_v1_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildImageRequest) ProtoMessage() {}

func (x *BuildImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_k8s_controller_request_v1_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildImageRequest.ProtoReflect.Descriptor instead.
func (*BuildImageRequest) Descriptor() ([]byte, []int) {
	return file_proto_k8s_controller_request_v1_proto_rawDescGZIP(), []int{10}
}

func (x *BuildImageRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *BuildImageRequest) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

func (x *BuildImageRequest) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

func (x *BuildImageRequest) GetPullRequestId() string {
	if x != nil {
		return x.PullRequestId
	}
	return ""
}

func (x *BuildImageRequest) GetGitRepo() string {
	if x != nil {
		return x.GitRepo
	}
	return ""
}

func (x *BuildImageRequest) GetDockerfileDir() string {
	if x != nil {
		return x.DockerfileDir
	}
	return ""
}

func (x *BuildImageRequest) GetDockerfilePath() string {
	if x != nil {
		return x.DockerfilePath
	}
	return ""
}

func (x *BuildImageRequest) GetImageName() string {
	if x != nil {
		return x.ImageName
	}
	return ""
}

func (x *BuildImageRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_proto_k8s_controller_request_v1_proto protoreflect.FileDescriptor

var file_proto_k8s_controller_request_v1_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6b, 0x38, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x76,
	0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x63,
	0x64, 0x2e, 0x6b, 0x38, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x22, 0x2c, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x2c, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0xf8, 0x01, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x72, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x22, 0x74, 0x0a, 0x17, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x6f, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x22, 0x37, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x67, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x58, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x6f, 0x72, 0x79, 0x22, 0x5c, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x42, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a,
	0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72,
	0x79, 0x22, 0x9a, 0x01, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e,
	0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x26,
	0x0a, 0x0f, 0x70, 0x75, 0x6c, 0x6c, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x22, 0x32,
	0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0xd9, 0x02, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12,
	0x26, 0x0a, 0x0f, 0x70, 0x75, 0x6c, 0x6c, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x75, 0x6c, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x70,
	0x75, 0x6c, 0x6c, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x75, 0x6c, 0x6c, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x12, 0x3c, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x63, 0x64, 0x2e, 0x6b, 0x38, 0x73, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x65,
	0x78, 0x70, 0x6f, 0x73, 0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28,
	0x05, 0x52, 0x0b, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x22, 0xb1,
	0x02, 0x0a, 0x11, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f,
	0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x75,
	0x6c, 0x6c, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x69, 0x74, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x69, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x12, 0x25, 0x0a,
	0x0e, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x64, 0x69, 0x72, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x66, 0x69, 0x6c,
	0x65, 0x44, 0x69, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64,
	0x6f, 0x63, 0x6b, 0x65, 0x72, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1d, 0x0a,
	0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x63, 0x64, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x6b,
	0x38, 0x73, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_k8s_controller_request_v1_proto_rawDescOnce sync.Once
	file_proto_k8s_controller_request_v1_proto_rawDescData = file_proto_k8s_controller_request_v1_proto_rawDesc
)

func file_proto_k8s_controller_request_v1_proto_rawDescGZIP() []byte {
	file_proto_k8s_controller_request_v1_proto_rawDescOnce.Do(func() {
		file_proto_k8s_controller_request_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_k8s_controller_request_v1_proto_rawDescData)
	})
	return file_proto_k8s_controller_request_v1_proto_rawDescData
}

var file_proto_k8s_controller_request_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_k8s_controller_request_v1_proto_goTypes = []interface{}{
	(*CreateNamespaceRequest)(nil),  // 0: gantrycd.k8s_controller.v1.CreateNamespaceRequest
	(*DeleteNamespaceRequest)(nil),  // 1: gantrycd.k8s_controller.v1.DeleteNamespaceRequest
	(*CreateDeploymentRequest)(nil), // 2: gantrycd.k8s_controller.v1.CreateDeploymentRequest
	(*DeleteDeploymentRequest)(nil), // 3: gantrycd.k8s_controller.v1.DeleteDeploymentRequest
	(*GetOrgRepoRequest)(nil),       // 4: gantrycd.k8s_controller.v1.GetOrgRepoRequest
	(*GetResourceRequest)(nil),      // 5: gantrycd.k8s_controller.v1.GetResourceRequest
	(*GetRepoBranchesRequest)(nil),  // 6: gantrycd.k8s_controller.v1.GetRepoBranchesRequest
	(*DeletePreviewRequest)(nil),    // 7: gantrycd.k8s_controller.v1.DeletePreviewRequest
	(*Config)(nil),                  // 8: gantrycd.k8s_controller.v1.Config
	(*CreatePreviewRequest)(nil),    // 9: gantrycd.k8s_controller.v1.CreatePreviewRequest
	(*BuildImageRequest)(nil),       // 10: gantrycd.k8s_controller.v1.BuildImageRequest
}
var file_proto_k8s_controller_request_v1_proto_depIdxs = []int32{
	8, // 0: gantrycd.k8s_controller.v1.CreatePreviewRequest.configs:type_name -> gantrycd.k8s_controller.v1.Config
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_k8s_controller_request_v1_proto_init() }
func file_proto_k8s_controller_request_v1_proto_init() {
	if File_proto_k8s_controller_request_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_k8s_controller_request_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNamespaceRequest); i {
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
		file_proto_k8s_controller_request_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNamespaceRequest); i {
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
		file_proto_k8s_controller_request_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDeploymentRequest); i {
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
		file_proto_k8s_controller_request_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDeploymentRequest); i {
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
		file_proto_k8s_controller_request_v1_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrgRepoRequest); i {
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
		file_proto_k8s_controller_request_v1_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResourceRequest); i {
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
		file_proto_k8s_controller_request_v1_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRepoBranchesRequest); i {
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
		file_proto_k8s_controller_request_v1_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePreviewRequest); i {
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
		file_proto_k8s_controller_request_v1_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_proto_k8s_controller_request_v1_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePreviewRequest); i {
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
		file_proto_k8s_controller_request_v1_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildImageRequest); i {
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
			RawDescriptor: file_proto_k8s_controller_request_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_k8s_controller_request_v1_proto_goTypes,
		DependencyIndexes: file_proto_k8s_controller_request_v1_proto_depIdxs,
		MessageInfos:      file_proto_k8s_controller_request_v1_proto_msgTypes,
	}.Build()
	File_proto_k8s_controller_request_v1_proto = out.File
	file_proto_k8s_controller_request_v1_proto_rawDesc = nil
	file_proto_k8s_controller_request_v1_proto_goTypes = nil
	file_proto_k8s_controller_request_v1_proto_depIdxs = nil
}
