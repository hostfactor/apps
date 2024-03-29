// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: manifest/app.proto

package manifest

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

//	type Manifest struct {
//		Name        string    `yaml:"name,omitempty" json:"name,omitempty"`
//		Author      string    `yaml:"author,omitempty" json:"author,omitempty"`
//		Description string    `yaml:"description,omitempty" json:"description,omitempty"`
//		Features    []Feature `yaml:"features,omitempty" json:"features,omitempty"`
//	}
//
//	type Feature struct {
//		Name        string            `yaml:"name,omitempty" json:"name,omitempty"`
//		Tags        []string          `yaml:"tags,omitempty" json:"tags,omitempty"`
//		Context     string            `yaml:"context,omitempty" json:"context,omitempty"`
//		Watch       *FeatureWatch     `yaml:"watch,omitempty" json:"watch,omitempty"`
//		BuildArgs   map[string]string `yaml:"build_args,omitempty" json:"build_args,omitempty"`
//		Description string            `yaml:"description" json:"description"`
//	}
//
//	type FeatureWatch struct {
//		GithubRelease *GithubReleaseFeatureWatch `yaml:"github_release,omitempty" json:"github_release,omitempty"`
//	}
//
//	type GithubReleaseFeatureWatch struct {
//		Repo string `yaml:"repo,omitempty" json:"repo,omitempty"`
//	}
type AppManifest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the container.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The author of the container
	Author string `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	// A short description for the container
	Description string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Builds      []*Build `protobuf:"bytes,4,rep,name=builds,proto3" json:"builds,omitempty"`
	// Images that are built prior to building the app. Can be used for multi-stage builds.
	Dependencies []*BuildSpec `protobuf:"bytes,5,rep,name=dependencies,proto3" json:"dependencies,omitempty"`
}

func (x *AppManifest) Reset() {
	*x = AppManifest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manifest_app_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppManifest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppManifest) ProtoMessage() {}

func (x *AppManifest) ProtoReflect() protoreflect.Message {
	mi := &file_manifest_app_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppManifest.ProtoReflect.Descriptor instead.
func (*AppManifest) Descriptor() ([]byte, []int) {
	return file_manifest_app_proto_rawDescGZIP(), []int{0}
}

func (x *AppManifest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AppManifest) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *AppManifest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AppManifest) GetBuilds() []*Build {
	if x != nil {
		return x.Builds
	}
	return nil
}

func (x *AppManifest) GetDependencies() []*BuildSpec {
	if x != nil {
		return x.Dependencies
	}
	return nil
}

type BuildSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tags []string `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
	// The docker context.
	Context   string            `protobuf:"bytes,2,opt,name=context,proto3" json:"context,omitempty"`
	BuildArgs map[string]string `protobuf:"bytes,3,rep,name=build_args,proto3" json:"build_args,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *BuildSpec) Reset() {
	*x = BuildSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manifest_app_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildSpec) ProtoMessage() {}

func (x *BuildSpec) ProtoReflect() protoreflect.Message {
	mi := &file_manifest_app_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildSpec.ProtoReflect.Descriptor instead.
func (*BuildSpec) Descriptor() ([]byte, []int) {
	return file_manifest_app_proto_rawDescGZIP(), []int{1}
}

func (x *BuildSpec) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *BuildSpec) GetContext() string {
	if x != nil {
		return x.Context
	}
	return ""
}

func (x *BuildSpec) GetBuildArgs() map[string]string {
	if x != nil {
		return x.BuildArgs
	}
	return nil
}

// Generates an image under the scope of the app.
type Build struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the feature. Used as a post-fix for identifying this particular app.
	Name        string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Spec        *BuildSpec    `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	Watch       *FeatureWatch `protobuf:"bytes,3,opt,name=watch,proto3" json:"watch,omitempty"`
	Description string        `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Build) Reset() {
	*x = Build{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manifest_app_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Build) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Build) ProtoMessage() {}

func (x *Build) ProtoReflect() protoreflect.Message {
	mi := &file_manifest_app_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Build.ProtoReflect.Descriptor instead.
func (*Build) Descriptor() ([]byte, []int) {
	return file_manifest_app_proto_rawDescGZIP(), []int{2}
}

func (x *Build) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Build) GetSpec() *BuildSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Build) GetWatch() *FeatureWatch {
	if x != nil {
		return x.Watch
	}
	return nil
}

func (x *Build) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type FeatureWatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GithubRelease *FeatureWatch_GithubRelease `protobuf:"bytes,1,opt,name=github_release,proto3,oneof" json:"github_release,omitempty"`
}

func (x *FeatureWatch) Reset() {
	*x = FeatureWatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manifest_app_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeatureWatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureWatch) ProtoMessage() {}

func (x *FeatureWatch) ProtoReflect() protoreflect.Message {
	mi := &file_manifest_app_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureWatch.ProtoReflect.Descriptor instead.
func (*FeatureWatch) Descriptor() ([]byte, []int) {
	return file_manifest_app_proto_rawDescGZIP(), []int{3}
}

func (x *FeatureWatch) GetGithubRelease() *FeatureWatch_GithubRelease {
	if x != nil {
		return x.GithubRelease
	}
	return nil
}

type FeatureWatch_GithubRelease struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Repo string `protobuf:"bytes,1,opt,name=repo,proto3" json:"repo,omitempty"`
}

func (x *FeatureWatch_GithubRelease) Reset() {
	*x = FeatureWatch_GithubRelease{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manifest_app_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeatureWatch_GithubRelease) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureWatch_GithubRelease) ProtoMessage() {}

func (x *FeatureWatch_GithubRelease) ProtoReflect() protoreflect.Message {
	mi := &file_manifest_app_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureWatch_GithubRelease.ProtoReflect.Descriptor instead.
func (*FeatureWatch_GithubRelease) Descriptor() ([]byte, []int) {
	return file_manifest_app_proto_rawDescGZIP(), []int{3, 0}
}

func (x *FeatureWatch_GithubRelease) GetRepo() string {
	if x != nil {
		return x.Repo
	}
	return ""
}

var File_manifest_app_proto protoreflect.FileDescriptor

var file_manifest_app_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x2f, 0x61, 0x70, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x22, 0xbd,
	0x01, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x06,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d,
	0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x06, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x73, 0x12, 0x37, 0x0a, 0x0c, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65,
	0x6e, 0x63, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x61,
	0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x53, 0x70, 0x65, 0x63,
	0x52, 0x0c, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x22, 0xbb,
	0x01, 0x0a, 0x09, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x53, 0x70, 0x65, 0x63, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x42, 0x0a, 0x0a, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x5f, 0x61, 0x72, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x53,
	0x70, 0x65, 0x63, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x41, 0x72, 0x67, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0a, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x61, 0x72, 0x67, 0x73, 0x1a, 0x3c,
	0x0a, 0x0e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x41, 0x72, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x94, 0x01, 0x0a,
	0x05, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x73, 0x70,
	0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x61, 0x6e, 0x69, 0x66,
	0x65, 0x73, 0x74, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73,
	0x70, 0x65, 0x63, 0x12, 0x2c, 0x0a, 0x05, 0x77, 0x61, 0x74, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x05, 0x77, 0x61, 0x74, 0x63,
	0x68, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x99, 0x01, 0x0a, 0x0c, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x57,
	0x61, 0x74, 0x63, 0x68, 0x12, 0x51, 0x0a, 0x0e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x5f, 0x72,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6d,
	0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x57,
	0x61, 0x74, 0x63, 0x68, 0x2e, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x52, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x48, 0x00, 0x52, 0x0e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x5f, 0x72, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x88, 0x01, 0x01, 0x1a, 0x23, 0x0a, 0x0d, 0x47, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x70, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x42, 0x11, 0x0a, 0x0f,
	0x5f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x5f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x42,
	0x84, 0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74,
	0x42, 0x08, 0x41, 0x70, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x66, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f,
	0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0xa2, 0x02, 0x03, 0x4d, 0x58, 0x58, 0xaa, 0x02,
	0x08, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0xca, 0x02, 0x08, 0x4d, 0x61, 0x6e, 0x69,
	0x66, 0x65, 0x73, 0x74, 0xe2, 0x02, 0x14, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x4d, 0x61,
	0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_manifest_app_proto_rawDescOnce sync.Once
	file_manifest_app_proto_rawDescData = file_manifest_app_proto_rawDesc
)

func file_manifest_app_proto_rawDescGZIP() []byte {
	file_manifest_app_proto_rawDescOnce.Do(func() {
		file_manifest_app_proto_rawDescData = protoimpl.X.CompressGZIP(file_manifest_app_proto_rawDescData)
	})
	return file_manifest_app_proto_rawDescData
}

var file_manifest_app_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_manifest_app_proto_goTypes = []interface{}{
	(*AppManifest)(nil),                // 0: manifest.AppManifest
	(*BuildSpec)(nil),                  // 1: manifest.BuildSpec
	(*Build)(nil),                      // 2: manifest.Build
	(*FeatureWatch)(nil),               // 3: manifest.FeatureWatch
	nil,                                // 4: manifest.BuildSpec.BuildArgsEntry
	(*FeatureWatch_GithubRelease)(nil), // 5: manifest.FeatureWatch.GithubRelease
}
var file_manifest_app_proto_depIdxs = []int32{
	2, // 0: manifest.AppManifest.builds:type_name -> manifest.Build
	1, // 1: manifest.AppManifest.dependencies:type_name -> manifest.BuildSpec
	4, // 2: manifest.BuildSpec.build_args:type_name -> manifest.BuildSpec.BuildArgsEntry
	1, // 3: manifest.Build.spec:type_name -> manifest.BuildSpec
	3, // 4: manifest.Build.watch:type_name -> manifest.FeatureWatch
	5, // 5: manifest.FeatureWatch.github_release:type_name -> manifest.FeatureWatch.GithubRelease
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_manifest_app_proto_init() }
func file_manifest_app_proto_init() {
	if File_manifest_app_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_manifest_app_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppManifest); i {
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
		file_manifest_app_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildSpec); i {
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
		file_manifest_app_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Build); i {
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
		file_manifest_app_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeatureWatch); i {
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
		file_manifest_app_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeatureWatch_GithubRelease); i {
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
	file_manifest_app_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_manifest_app_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_manifest_app_proto_goTypes,
		DependencyIndexes: file_manifest_app_proto_depIdxs,
		MessageInfos:      file_manifest_app_proto_msgTypes,
	}.Build()
	File_manifest_app_proto = out.File
	file_manifest_app_proto_rawDesc = nil
	file_manifest_app_proto_goTypes = nil
	file_manifest_app_proto_depIdxs = nil
}
