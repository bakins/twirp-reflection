// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: bakins/twirp/reflection/v0/reflection.proto

package twirp_reflection

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

type ListServicesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListServicesRequest) Reset() {
	*x = ListServicesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bakins_twirp_reflection_v0_reflection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListServicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListServicesRequest) ProtoMessage() {}

func (x *ListServicesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bakins_twirp_reflection_v0_reflection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListServicesRequest.ProtoReflect.Descriptor instead.
func (*ListServicesRequest) Descriptor() ([]byte, []int) {
	return file_bakins_twirp_reflection_v0_reflection_proto_rawDescGZIP(), []int{0}
}

type ListServicesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Services []*ListServiceResponse `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
}

func (x *ListServicesResponse) Reset() {
	*x = ListServicesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bakins_twirp_reflection_v0_reflection_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListServicesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListServicesResponse) ProtoMessage() {}

func (x *ListServicesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bakins_twirp_reflection_v0_reflection_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListServicesResponse.ProtoReflect.Descriptor instead.
func (*ListServicesResponse) Descriptor() ([]byte, []int) {
	return file_bakins_twirp_reflection_v0_reflection_proto_rawDescGZIP(), []int{1}
}

func (x *ListServicesResponse) GetServices() []*ListServiceResponse {
	if x != nil {
		return x.Services
	}
	return nil
}

type ListServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Full name of a registered service, including its package name. The format
	// is <package>.<service>
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ListServiceResponse) Reset() {
	*x = ListServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bakins_twirp_reflection_v0_reflection_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListServiceResponse) ProtoMessage() {}

func (x *ListServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bakins_twirp_reflection_v0_reflection_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListServiceResponse.ProtoReflect.Descriptor instead.
func (*ListServiceResponse) Descriptor() ([]byte, []int) {
	return file_bakins_twirp_reflection_v0_reflection_proto_rawDescGZIP(), []int{2}
}

func (x *ListServiceResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetServiceDescriptorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Full name of a registered service, including its package name. The format
	// is <package>.<service>
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetServiceDescriptorRequest) Reset() {
	*x = GetServiceDescriptorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bakins_twirp_reflection_v0_reflection_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServiceDescriptorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServiceDescriptorRequest) ProtoMessage() {}

func (x *GetServiceDescriptorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bakins_twirp_reflection_v0_reflection_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServiceDescriptorRequest.ProtoReflect.Descriptor instead.
func (*GetServiceDescriptorRequest) Descriptor() ([]byte, []int) {
	return file_bakins_twirp_reflection_v0_reflection_proto_rawDescGZIP(), []int{3}
}

func (x *GetServiceDescriptorRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetServiceDescriptorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Serialized FileDescriptorProto message. We avoid taking a dependency on
	// descriptor.proto, which uses proto2 only features, by making them opaque
	// bytes instead.
	FileDescriptor []byte `protobuf:"bytes,1,opt,name=file_descriptor,json=fileDescriptor,proto3" json:"file_descriptor,omitempty"`
}

func (x *GetServiceDescriptorResponse) Reset() {
	*x = GetServiceDescriptorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bakins_twirp_reflection_v0_reflection_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServiceDescriptorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServiceDescriptorResponse) ProtoMessage() {}

func (x *GetServiceDescriptorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bakins_twirp_reflection_v0_reflection_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServiceDescriptorResponse.ProtoReflect.Descriptor instead.
func (*GetServiceDescriptorResponse) Descriptor() ([]byte, []int) {
	return file_bakins_twirp_reflection_v0_reflection_proto_rawDescGZIP(), []int{4}
}

func (x *GetServiceDescriptorResponse) GetFileDescriptor() []byte {
	if x != nil {
		return x.FileDescriptor
	}
	return nil
}

type GetSymbolDescriptorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
}

func (x *GetSymbolDescriptorRequest) Reset() {
	*x = GetSymbolDescriptorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bakins_twirp_reflection_v0_reflection_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSymbolDescriptorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSymbolDescriptorRequest) ProtoMessage() {}

func (x *GetSymbolDescriptorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bakins_twirp_reflection_v0_reflection_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSymbolDescriptorRequest.ProtoReflect.Descriptor instead.
func (*GetSymbolDescriptorRequest) Descriptor() ([]byte, []int) {
	return file_bakins_twirp_reflection_v0_reflection_proto_rawDescGZIP(), []int{5}
}

func (x *GetSymbolDescriptorRequest) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

type GetSymbolDescriptorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileDescriptors [][]byte `protobuf:"bytes,1,rep,name=file_descriptors,json=fileDescriptors,proto3" json:"file_descriptors,omitempty"`
}

func (x *GetSymbolDescriptorResponse) Reset() {
	*x = GetSymbolDescriptorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bakins_twirp_reflection_v0_reflection_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSymbolDescriptorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSymbolDescriptorResponse) ProtoMessage() {}

func (x *GetSymbolDescriptorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bakins_twirp_reflection_v0_reflection_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSymbolDescriptorResponse.ProtoReflect.Descriptor instead.
func (*GetSymbolDescriptorResponse) Descriptor() ([]byte, []int) {
	return file_bakins_twirp_reflection_v0_reflection_proto_rawDescGZIP(), []int{6}
}

func (x *GetSymbolDescriptorResponse) GetFileDescriptors() [][]byte {
	if x != nil {
		return x.FileDescriptors
	}
	return nil
}

var File_bakins_twirp_reflection_v0_reflection_proto protoreflect.FileDescriptor

var file_bakins_twirp_reflection_v0_reflection_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x62, 0x61, 0x6b, 0x69, 0x6e, 0x73, 0x2f, 0x74, 0x77, 0x69, 0x72, 0x70, 0x2f, 0x72,
	0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x30, 0x2f, 0x72, 0x65, 0x66,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x62,
	0x61, 0x6b, 0x69, 0x6e, 0x73, 0x2e, 0x74, 0x77, 0x69, 0x72, 0x70, 0x2e, 0x72, 0x65, 0x66, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x30, 0x22, 0x15, 0x0a, 0x13, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x63, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x62, 0x61, 0x6b,
	0x69, 0x6e, 0x73, 0x2e, 0x74, 0x77, 0x69, 0x72, 0x70, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x30, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x29, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x31, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x47, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x66, 0x69,
	0x6c, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x22, 0x34, 0x0a, 0x1a,
	0x47, 0x65, 0x74, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62,
	0x6f, 0x6c, 0x22, 0x48, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x29, 0x0a, 0x10, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0f, 0x66, 0x69, 0x6c,
	0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x73, 0x32, 0xa1, 0x03, 0x0a,
	0x17, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x71, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x2f, 0x2e, 0x62, 0x61, 0x6b, 0x69, 0x6e,
	0x73, 0x2e, 0x74, 0x77, 0x69, 0x72, 0x70, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x30, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x62, 0x61, 0x6b, 0x69,
	0x6e, 0x73, 0x2e, 0x74, 0x77, 0x69, 0x72, 0x70, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x30, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x89, 0x01, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x12, 0x37, 0x2e, 0x62, 0x61, 0x6b, 0x69, 0x6e, 0x73, 0x2e, 0x74, 0x77,
	0x69, 0x72, 0x70, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x30, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e,
	0x62, 0x61, 0x6b, 0x69, 0x6e, 0x73, 0x2e, 0x74, 0x77, 0x69, 0x72, 0x70, 0x2e, 0x72, 0x65, 0x66,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x30, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x86, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53,
	0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12,
	0x36, 0x2e, 0x62, 0x61, 0x6b, 0x69, 0x6e, 0x73, 0x2e, 0x74, 0x77, 0x69, 0x72, 0x70, 0x2e, 0x72,
	0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x30, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x62, 0x61, 0x6b, 0x69, 0x6e, 0x73,
	0x2e, 0x74, 0x77, 0x69, 0x72, 0x70, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x30, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62,
	0x61, 0x6b, 0x69, 0x6e, 0x73, 0x2f, 0x74, 0x77, 0x69, 0x72, 0x70, 0x2d, 0x72, 0x65, 0x66, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bakins_twirp_reflection_v0_reflection_proto_rawDescOnce sync.Once
	file_bakins_twirp_reflection_v0_reflection_proto_rawDescData = file_bakins_twirp_reflection_v0_reflection_proto_rawDesc
)

func file_bakins_twirp_reflection_v0_reflection_proto_rawDescGZIP() []byte {
	file_bakins_twirp_reflection_v0_reflection_proto_rawDescOnce.Do(func() {
		file_bakins_twirp_reflection_v0_reflection_proto_rawDescData = protoimpl.X.CompressGZIP(file_bakins_twirp_reflection_v0_reflection_proto_rawDescData)
	})
	return file_bakins_twirp_reflection_v0_reflection_proto_rawDescData
}

var file_bakins_twirp_reflection_v0_reflection_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_bakins_twirp_reflection_v0_reflection_proto_goTypes = []interface{}{
	(*ListServicesRequest)(nil),          // 0: bakins.twirp.reflection.v0.ListServicesRequest
	(*ListServicesResponse)(nil),         // 1: bakins.twirp.reflection.v0.ListServicesResponse
	(*ListServiceResponse)(nil),          // 2: bakins.twirp.reflection.v0.ListServiceResponse
	(*GetServiceDescriptorRequest)(nil),  // 3: bakins.twirp.reflection.v0.GetServiceDescriptorRequest
	(*GetServiceDescriptorResponse)(nil), // 4: bakins.twirp.reflection.v0.GetServiceDescriptorResponse
	(*GetSymbolDescriptorRequest)(nil),   // 5: bakins.twirp.reflection.v0.GetSymbolDescriptorRequest
	(*GetSymbolDescriptorResponse)(nil),  // 6: bakins.twirp.reflection.v0.GetSymbolDescriptorResponse
}
var file_bakins_twirp_reflection_v0_reflection_proto_depIdxs = []int32{
	2, // 0: bakins.twirp.reflection.v0.ListServicesResponse.services:type_name -> bakins.twirp.reflection.v0.ListServiceResponse
	0, // 1: bakins.twirp.reflection.v0.ServerReflectionService.ListServices:input_type -> bakins.twirp.reflection.v0.ListServicesRequest
	3, // 2: bakins.twirp.reflection.v0.ServerReflectionService.GetServiceDescriptor:input_type -> bakins.twirp.reflection.v0.GetServiceDescriptorRequest
	5, // 3: bakins.twirp.reflection.v0.ServerReflectionService.GetSymbolDescriptor:input_type -> bakins.twirp.reflection.v0.GetSymbolDescriptorRequest
	1, // 4: bakins.twirp.reflection.v0.ServerReflectionService.ListServices:output_type -> bakins.twirp.reflection.v0.ListServicesResponse
	4, // 5: bakins.twirp.reflection.v0.ServerReflectionService.GetServiceDescriptor:output_type -> bakins.twirp.reflection.v0.GetServiceDescriptorResponse
	6, // 6: bakins.twirp.reflection.v0.ServerReflectionService.GetSymbolDescriptor:output_type -> bakins.twirp.reflection.v0.GetSymbolDescriptorResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_bakins_twirp_reflection_v0_reflection_proto_init() }
func file_bakins_twirp_reflection_v0_reflection_proto_init() {
	if File_bakins_twirp_reflection_v0_reflection_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bakins_twirp_reflection_v0_reflection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListServicesRequest); i {
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
		file_bakins_twirp_reflection_v0_reflection_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListServicesResponse); i {
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
		file_bakins_twirp_reflection_v0_reflection_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListServiceResponse); i {
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
		file_bakins_twirp_reflection_v0_reflection_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServiceDescriptorRequest); i {
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
		file_bakins_twirp_reflection_v0_reflection_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServiceDescriptorResponse); i {
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
		file_bakins_twirp_reflection_v0_reflection_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSymbolDescriptorRequest); i {
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
		file_bakins_twirp_reflection_v0_reflection_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSymbolDescriptorResponse); i {
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
			RawDescriptor: file_bakins_twirp_reflection_v0_reflection_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bakins_twirp_reflection_v0_reflection_proto_goTypes,
		DependencyIndexes: file_bakins_twirp_reflection_v0_reflection_proto_depIdxs,
		MessageInfos:      file_bakins_twirp_reflection_v0_reflection_proto_msgTypes,
	}.Build()
	File_bakins_twirp_reflection_v0_reflection_proto = out.File
	file_bakins_twirp_reflection_v0_reflection_proto_rawDesc = nil
	file_bakins_twirp_reflection_v0_reflection_proto_goTypes = nil
	file_bakins_twirp_reflection_v0_reflection_proto_depIdxs = nil
}
