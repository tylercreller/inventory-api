// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: kessel/inventory/v1beta1/resources/k8s_policies_service.proto

package resources

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateK8SPolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The policy to create in Kessel Asset Inventory
	K8SPolicy *K8SPolicy `protobuf:"bytes,1,opt,name=k8s_policy,json=k8sPolicy,proto3" json:"k8s_policy,omitempty"`
}

func (x *CreateK8SPolicyRequest) Reset() {
	*x = CreateK8SPolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_inventory_v1beta1_resources_k8s_policies_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateK8SPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateK8SPolicyRequest) ProtoMessage() {}

func (x *CreateK8SPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_inventory_v1beta1_resources_k8s_policies_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateK8SPolicyRequest.ProtoReflect.Descriptor instead.
func (*CreateK8SPolicyRequest) Descriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta1_resources_k8s_policies_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateK8SPolicyRequest) GetK8SPolicy() *K8SPolicy {
	if x != nil {
		return x.K8SPolicy
	}
	return nil
}

type CreateK8SPolicyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateK8SPolicyResponse) Reset() {
	*x = CreateK8SPolicyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_inventory_v1beta1_resources_k8s_policies_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateK8SPolicyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateK8SPolicyResponse) ProtoMessage() {}

func (x *CreateK8SPolicyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_inventory_v1beta1_resources_k8s_policies_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateK8SPolicyResponse.ProtoReflect.Descriptor instead.
func (*CreateK8SPolicyResponse) Descriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta1_resources_k8s_policies_service_proto_rawDescGZIP(), []int{1}
}

type UpdateK8SPolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource to be updated will be defined by
	// \"<reporter_data.reporter_type>:<reporter_instaance_id>:<reporter_data.local_resource_id>\"
	// from the request body.
	K8SPolicy *K8SPolicy `protobuf:"bytes,1,opt,name=k8s_policy,json=k8sPolicy,proto3" json:"k8s_policy,omitempty"`
}

func (x *UpdateK8SPolicyRequest) Reset() {
	*x = UpdateK8SPolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_inventory_v1beta1_resources_k8s_policies_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateK8SPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateK8SPolicyRequest) ProtoMessage() {}

func (x *UpdateK8SPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_inventory_v1beta1_resources_k8s_policies_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateK8SPolicyRequest.ProtoReflect.Descriptor instead.
func (*UpdateK8SPolicyRequest) Descriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta1_resources_k8s_policies_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateK8SPolicyRequest) GetK8SPolicy() *K8SPolicy {
	if x != nil {
		return x.K8SPolicy
	}
	return nil
}

type UpdateK8SPolicyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateK8SPolicyResponse) Reset() {
	*x = UpdateK8SPolicyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_inventory_v1beta1_resources_k8s_policies_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateK8SPolicyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateK8SPolicyResponse) ProtoMessage() {}

func (x *UpdateK8SPolicyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_inventory_v1beta1_resources_k8s_policies_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateK8SPolicyResponse.ProtoReflect.Descriptor instead.
func (*UpdateK8SPolicyResponse) Descriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta1_resources_k8s_policies_service_proto_rawDescGZIP(), []int{3}
}

type DeleteK8SPolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource to be deleted will be defined by
	// \"<reporter_data.reporter_type>:<reporter_instaance_id>:<reporter_data.local_resource_id>\"
	// from the request body.
	K8SPolicy *K8SPolicy `protobuf:"bytes,1,opt,name=k8s_policy,json=k8sPolicy,proto3" json:"k8s_policy,omitempty"`
}

func (x *DeleteK8SPolicyRequest) Reset() {
	*x = DeleteK8SPolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_inventory_v1beta1_resources_k8s_policies_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteK8SPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteK8SPolicyRequest) ProtoMessage() {}

func (x *DeleteK8SPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_inventory_v1beta1_resources_k8s_policies_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteK8SPolicyRequest.ProtoReflect.Descriptor instead.
func (*DeleteK8SPolicyRequest) Descriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta1_resources_k8s_policies_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteK8SPolicyRequest) GetK8SPolicy() *K8SPolicy {
	if x != nil {
		return x.K8SPolicy
	}
	return nil
}

type DeleteK8SPolicyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteK8SPolicyResponse) Reset() {
	*x = DeleteK8SPolicyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_inventory_v1beta1_resources_k8s_policies_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteK8SPolicyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteK8SPolicyResponse) ProtoMessage() {}

func (x *DeleteK8SPolicyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_inventory_v1beta1_resources_k8s_policies_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteK8SPolicyResponse.ProtoReflect.Descriptor instead.
func (*DeleteK8SPolicyResponse) Descriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta1_resources_k8s_policies_service_proto_rawDescGZIP(), []int{5}
}

var File_kessel_inventory_v1beta1_resources_k8s_policies_service_proto protoreflect.FileDescriptor

var file_kessel_inventory_v1beta1_resources_k8s_policies_service_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x6b, 0x38, 0x73, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65,
	0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x22, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x33, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6b, 0x38, 0x73, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4b, 0x38, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x4c, 0x0a, 0x0a, 0x6b, 0x38, 0x73, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x4b, 0x38, 0x73, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x52, 0x09, 0x6b, 0x38, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0x19,
	0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x38, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x66, 0x0a, 0x16, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4b, 0x38, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x4c, 0x0a, 0x0a, 0x6b, 0x38, 0x73, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c,
	0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x4b, 0x38, 0x73,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x09, 0x6b, 0x38, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x22, 0x19, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4b, 0x38, 0x73, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x66, 0x0a, 0x16,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x38, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4c, 0x0a, 0x0a, 0x6b, 0x38, 0x73, 0x5f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x6b, 0x65, 0x73,
	0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e,
	0x4b, 0x38, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x09, 0x6b, 0x38, 0x73, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x22, 0x19, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x38,
	0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0xed, 0x04, 0x0a, 0x16, 0x4b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x4b, 0x38, 0x73, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xc4, 0x01, 0x0a, 0x0f, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x38, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x3a,
	0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x38, 0x73, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x6b, 0x65, 0x73,
	0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x38, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x38, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x32, 0x3a,
	0x01, 0x2a, 0x22, 0x2d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x6b, 0x38, 0x73, 0x2d, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65,
	0x73, 0x12, 0xc4, 0x01, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4b, 0x38, 0x73, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x3a, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4b, 0x38, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x3b, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4b, 0x38, 0x73,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x38,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x32, 0x3a, 0x01, 0x2a, 0x1a, 0x2d, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6b, 0x38, 0x73, 0x2d,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x12, 0xc4, 0x01, 0x0a, 0x0f, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4b, 0x38, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x3a, 0x2e, 0x6b,
	0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x38, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65,
	0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4b, 0x38, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x38, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x32, 0x3a, 0x01, 0x2a,
	0x2a, 0x2d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2f, 0x6b, 0x38, 0x73, 0x2d, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x42,
	0x86, 0x01, 0x0a, 0x32, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x50, 0x01, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x6b, 0x65, 0x73,
	0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2d, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kessel_inventory_v1beta1_resources_k8s_policies_service_proto_rawDescOnce sync.Once
	file_kessel_inventory_v1beta1_resources_k8s_policies_service_proto_rawDescData = file_kessel_inventory_v1beta1_resources_k8s_policies_service_proto_rawDesc
)

func file_kessel_inventory_v1beta1_resources_k8s_policies_service_proto_rawDescGZIP() []byte {
	file_kessel_inventory_v1beta1_resources_k8s_policies_service_proto_rawDescOnce.Do(func() {
		file_kessel_inventory_v1beta1_resources_k8s_policies_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_kessel_inventory_v1beta1_resources_k8s_policies_service_proto_rawDescData)
	})
	return file_kessel_inventory_v1beta1_resources_k8s_policies_service_proto_rawDescData
}

var file_kessel_inventory_v1beta1_resources_k8s_policies_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_kessel_inventory_v1beta1_resources_k8s_policies_service_proto_goTypes = []any{
	(*CreateK8SPolicyRequest)(nil),  // 0: kessel.inventory.v1beta1.resources.CreateK8sPolicyRequest
	(*CreateK8SPolicyResponse)(nil), // 1: kessel.inventory.v1beta1.resources.CreateK8sPolicyResponse
	(*UpdateK8SPolicyRequest)(nil),  // 2: kessel.inventory.v1beta1.resources.UpdateK8sPolicyRequest
	(*UpdateK8SPolicyResponse)(nil), // 3: kessel.inventory.v1beta1.resources.UpdateK8sPolicyResponse
	(*DeleteK8SPolicyRequest)(nil),  // 4: kessel.inventory.v1beta1.resources.DeleteK8sPolicyRequest
	(*DeleteK8SPolicyResponse)(nil), // 5: kessel.inventory.v1beta1.resources.DeleteK8sPolicyResponse
	(*K8SPolicy)(nil),               // 6: kessel.inventory.v1beta1.resources.K8sPolicy
}
var file_kessel_inventory_v1beta1_resources_k8s_policies_service_proto_depIdxs = []int32{
	6, // 0: kessel.inventory.v1beta1.resources.CreateK8sPolicyRequest.k8s_policy:type_name -> kessel.inventory.v1beta1.resources.K8sPolicy
	6, // 1: kessel.inventory.v1beta1.resources.UpdateK8sPolicyRequest.k8s_policy:type_name -> kessel.inventory.v1beta1.resources.K8sPolicy
	6, // 2: kessel.inventory.v1beta1.resources.DeleteK8sPolicyRequest.k8s_policy:type_name -> kessel.inventory.v1beta1.resources.K8sPolicy
	0, // 3: kessel.inventory.v1beta1.resources.KesselK8sPolicyService.CreateK8sPolicy:input_type -> kessel.inventory.v1beta1.resources.CreateK8sPolicyRequest
	2, // 4: kessel.inventory.v1beta1.resources.KesselK8sPolicyService.UpdateK8sPolicy:input_type -> kessel.inventory.v1beta1.resources.UpdateK8sPolicyRequest
	4, // 5: kessel.inventory.v1beta1.resources.KesselK8sPolicyService.DeleteK8sPolicy:input_type -> kessel.inventory.v1beta1.resources.DeleteK8sPolicyRequest
	1, // 6: kessel.inventory.v1beta1.resources.KesselK8sPolicyService.CreateK8sPolicy:output_type -> kessel.inventory.v1beta1.resources.CreateK8sPolicyResponse
	3, // 7: kessel.inventory.v1beta1.resources.KesselK8sPolicyService.UpdateK8sPolicy:output_type -> kessel.inventory.v1beta1.resources.UpdateK8sPolicyResponse
	5, // 8: kessel.inventory.v1beta1.resources.KesselK8sPolicyService.DeleteK8sPolicy:output_type -> kessel.inventory.v1beta1.resources.DeleteK8sPolicyResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_kessel_inventory_v1beta1_resources_k8s_policies_service_proto_init() }
func file_kessel_inventory_v1beta1_resources_k8s_policies_service_proto_init() {
	if File_kessel_inventory_v1beta1_resources_k8s_policies_service_proto != nil {
		return
	}
	file_kessel_inventory_v1beta1_resources_k8s_policy_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_kessel_inventory_v1beta1_resources_k8s_policies_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateK8SPolicyRequest); i {
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
		file_kessel_inventory_v1beta1_resources_k8s_policies_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateK8SPolicyResponse); i {
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
		file_kessel_inventory_v1beta1_resources_k8s_policies_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateK8SPolicyRequest); i {
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
		file_kessel_inventory_v1beta1_resources_k8s_policies_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateK8SPolicyResponse); i {
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
		file_kessel_inventory_v1beta1_resources_k8s_policies_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteK8SPolicyRequest); i {
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
		file_kessel_inventory_v1beta1_resources_k8s_policies_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteK8SPolicyResponse); i {
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
			RawDescriptor: file_kessel_inventory_v1beta1_resources_k8s_policies_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kessel_inventory_v1beta1_resources_k8s_policies_service_proto_goTypes,
		DependencyIndexes: file_kessel_inventory_v1beta1_resources_k8s_policies_service_proto_depIdxs,
		MessageInfos:      file_kessel_inventory_v1beta1_resources_k8s_policies_service_proto_msgTypes,
	}.Build()
	File_kessel_inventory_v1beta1_resources_k8s_policies_service_proto = out.File
	file_kessel_inventory_v1beta1_resources_k8s_policies_service_proto_rawDesc = nil
	file_kessel_inventory_v1beta1_resources_k8s_policies_service_proto_goTypes = nil
	file_kessel_inventory_v1beta1_resources_k8s_policies_service_proto_depIdxs = nil
}
