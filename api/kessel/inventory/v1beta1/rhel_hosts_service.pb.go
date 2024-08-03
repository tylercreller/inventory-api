// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: kessel/inventory/v1beta1/rhel_hosts_service.proto

package v1beta1

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

type CreateRhelHostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Rhel Host to create in Kessel Asset Inventory
	Host *RhelHost `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
}

func (x *CreateRhelHostRequest) Reset() {
	*x = CreateRhelHostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_inventory_v1beta1_rhel_hosts_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRhelHostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRhelHostRequest) ProtoMessage() {}

func (x *CreateRhelHostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_inventory_v1beta1_rhel_hosts_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRhelHostRequest.ProtoReflect.Descriptor instead.
func (*CreateRhelHostRequest) Descriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta1_rhel_hosts_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRhelHostRequest) GetHost() *RhelHost {
	if x != nil {
		return x.Host
	}
	return nil
}

type CreateRhelHostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Rhel Host created in Kessel Asset Inventory
	Host *RhelHost `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
}

func (x *CreateRhelHostResponse) Reset() {
	*x = CreateRhelHostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_inventory_v1beta1_rhel_hosts_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRhelHostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRhelHostResponse) ProtoMessage() {}

func (x *CreateRhelHostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_inventory_v1beta1_rhel_hosts_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRhelHostResponse.ProtoReflect.Descriptor instead.
func (*CreateRhelHostResponse) Descriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta1_rhel_hosts_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRhelHostResponse) GetHost() *RhelHost {
	if x != nil {
		return x.Host
	}
	return nil
}

type UpdateRhelHostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A resource instance using the format:
	// \"<reporter_data.reporter_type>:<reporter_data.reporter_id>::<reporter_data.resourceId_alias>\"
	Resource string `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	// The Rhel host to update
	Host *RhelHost `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
}

func (x *UpdateRhelHostRequest) Reset() {
	*x = UpdateRhelHostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_inventory_v1beta1_rhel_hosts_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRhelHostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRhelHostRequest) ProtoMessage() {}

func (x *UpdateRhelHostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_inventory_v1beta1_rhel_hosts_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRhelHostRequest.ProtoReflect.Descriptor instead.
func (*UpdateRhelHostRequest) Descriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta1_rhel_hosts_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateRhelHostRequest) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *UpdateRhelHostRequest) GetHost() *RhelHost {
	if x != nil {
		return x.Host
	}
	return nil
}

type UpdateRhelHostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateRhelHostResponse) Reset() {
	*x = UpdateRhelHostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_inventory_v1beta1_rhel_hosts_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRhelHostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRhelHostResponse) ProtoMessage() {}

func (x *UpdateRhelHostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_inventory_v1beta1_rhel_hosts_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRhelHostResponse.ProtoReflect.Descriptor instead.
func (*UpdateRhelHostResponse) Descriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta1_rhel_hosts_service_proto_rawDescGZIP(), []int{3}
}

type DeleteRhelHostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A resource instance using the format:
	// \"<reporter_data.reporter_type>:<reporter_data.reporter_id>::<reporter_data.resourceId_alias>\"
	Resource string `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *DeleteRhelHostRequest) Reset() {
	*x = DeleteRhelHostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_inventory_v1beta1_rhel_hosts_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRhelHostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRhelHostRequest) ProtoMessage() {}

func (x *DeleteRhelHostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_inventory_v1beta1_rhel_hosts_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRhelHostRequest.ProtoReflect.Descriptor instead.
func (*DeleteRhelHostRequest) Descriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta1_rhel_hosts_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteRhelHostRequest) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

type DeleteRhelHostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteRhelHostResponse) Reset() {
	*x = DeleteRhelHostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_inventory_v1beta1_rhel_hosts_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRhelHostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRhelHostResponse) ProtoMessage() {}

func (x *DeleteRhelHostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_inventory_v1beta1_rhel_hosts_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRhelHostResponse.ProtoReflect.Descriptor instead.
func (*DeleteRhelHostResponse) Descriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta1_rhel_hosts_service_proto_rawDescGZIP(), []int{5}
}

var File_kessel_inventory_v1beta1_rhel_hosts_service_proto protoreflect.FileDescriptor

var file_kessel_inventory_v1beta1_rhel_hosts_service_proto_rawDesc = []byte{
	0x0a, 0x31, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x72, 0x68, 0x65, 0x6c, 0x5f,
	0x68, 0x6f, 0x73, 0x74, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e,
	0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x28, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x72, 0x68, 0x65, 0x6c, 0x5f,
	0x68, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x15, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x68, 0x65, 0x6c, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x52, 0x68, 0x65, 0x6c, 0x48, 0x6f, 0x73, 0x74, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52,
	0x04, 0x68, 0x6f, 0x73, 0x74, 0x22, 0x5a, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x68, 0x65, 0x6c, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x40, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x68, 0x65,
	0x6c, 0x48, 0x6f, 0x73, 0x74, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x04, 0x68, 0x6f, 0x73,
	0x74, 0x22, 0x7b, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x68, 0x65, 0x6c, 0x48,
	0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41,
	0x01, 0x02, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x04,
	0x68, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x68, 0x65, 0x6c, 0x48, 0x6f,
	0x73, 0x74, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x22, 0x18,
	0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x68, 0x65, 0x6c, 0x48, 0x6f, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x39, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x68, 0x65, 0x6c, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x20, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x22, 0x18, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x68, 0x65,
	0x6c, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x9c, 0x04,
	0x0a, 0x0c, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa7,
	0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x68, 0x65, 0x6c, 0x48, 0x6f, 0x73,
	0x74, 0x12, 0x33, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x68, 0x65, 0x6c, 0x48, 0x6f, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x73,
	0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x68, 0x65, 0x6c,
	0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x24, 0x3a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x22, 0x1c, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x12, 0xb2, 0x01, 0x0a, 0x0e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x68, 0x65, 0x6c, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x33, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x68, 0x65, 0x6c, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x34, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x68, 0x65, 0x6c, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x35, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2f, 0x3a, 0x04,
	0x68, 0x6f, 0x73, 0x74, 0x1a, 0x27, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x68, 0x6f, 0x73,
	0x74, 0x73, 0x2f, 0x7b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x7d, 0x12, 0xac, 0x01,
	0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x68, 0x65, 0x6c, 0x48, 0x6f, 0x73, 0x74,
	0x12, 0x33, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x68, 0x65, 0x6c, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x73, 0x73,
	0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x68, 0x65, 0x6c, 0x48,
	0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2f, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x29, 0x2a, 0x27, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x68, 0x6f, 0x73, 0x74,
	0x73, 0x2f, 0x7b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x7d, 0x42, 0x46, 0x0a, 0x1c,
	0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x50, 0x01, 0x5a, 0x24,
	0x61, 0x70, 0x69, 0x2f, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x50, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kessel_inventory_v1beta1_rhel_hosts_service_proto_rawDescOnce sync.Once
	file_kessel_inventory_v1beta1_rhel_hosts_service_proto_rawDescData = file_kessel_inventory_v1beta1_rhel_hosts_service_proto_rawDesc
)

func file_kessel_inventory_v1beta1_rhel_hosts_service_proto_rawDescGZIP() []byte {
	file_kessel_inventory_v1beta1_rhel_hosts_service_proto_rawDescOnce.Do(func() {
		file_kessel_inventory_v1beta1_rhel_hosts_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_kessel_inventory_v1beta1_rhel_hosts_service_proto_rawDescData)
	})
	return file_kessel_inventory_v1beta1_rhel_hosts_service_proto_rawDescData
}

var file_kessel_inventory_v1beta1_rhel_hosts_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_kessel_inventory_v1beta1_rhel_hosts_service_proto_goTypes = []any{
	(*CreateRhelHostRequest)(nil),  // 0: api.kessel.inventory.v1beta1.CreateRhelHostRequest
	(*CreateRhelHostResponse)(nil), // 1: api.kessel.inventory.v1beta1.CreateRhelHostResponse
	(*UpdateRhelHostRequest)(nil),  // 2: api.kessel.inventory.v1beta1.UpdateRhelHostRequest
	(*UpdateRhelHostResponse)(nil), // 3: api.kessel.inventory.v1beta1.UpdateRhelHostResponse
	(*DeleteRhelHostRequest)(nil),  // 4: api.kessel.inventory.v1beta1.DeleteRhelHostRequest
	(*DeleteRhelHostResponse)(nil), // 5: api.kessel.inventory.v1beta1.DeleteRhelHostResponse
	(*RhelHost)(nil),               // 6: api.kessel.inventory.v1beta1.RhelHost
}
var file_kessel_inventory_v1beta1_rhel_hosts_service_proto_depIdxs = []int32{
	6, // 0: api.kessel.inventory.v1beta1.CreateRhelHostRequest.host:type_name -> api.kessel.inventory.v1beta1.RhelHost
	6, // 1: api.kessel.inventory.v1beta1.CreateRhelHostResponse.host:type_name -> api.kessel.inventory.v1beta1.RhelHost
	6, // 2: api.kessel.inventory.v1beta1.UpdateRhelHostRequest.host:type_name -> api.kessel.inventory.v1beta1.RhelHost
	0, // 3: api.kessel.inventory.v1beta1.HostsService.CreateRhelHost:input_type -> api.kessel.inventory.v1beta1.CreateRhelHostRequest
	2, // 4: api.kessel.inventory.v1beta1.HostsService.UpdateRhelHost:input_type -> api.kessel.inventory.v1beta1.UpdateRhelHostRequest
	4, // 5: api.kessel.inventory.v1beta1.HostsService.DeleteRhelHost:input_type -> api.kessel.inventory.v1beta1.DeleteRhelHostRequest
	1, // 6: api.kessel.inventory.v1beta1.HostsService.CreateRhelHost:output_type -> api.kessel.inventory.v1beta1.CreateRhelHostResponse
	3, // 7: api.kessel.inventory.v1beta1.HostsService.UpdateRhelHost:output_type -> api.kessel.inventory.v1beta1.UpdateRhelHostResponse
	5, // 8: api.kessel.inventory.v1beta1.HostsService.DeleteRhelHost:output_type -> api.kessel.inventory.v1beta1.DeleteRhelHostResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_kessel_inventory_v1beta1_rhel_hosts_service_proto_init() }
func file_kessel_inventory_v1beta1_rhel_hosts_service_proto_init() {
	if File_kessel_inventory_v1beta1_rhel_hosts_service_proto != nil {
		return
	}
	file_kessel_inventory_v1beta1_rhel_host_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_kessel_inventory_v1beta1_rhel_hosts_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateRhelHostRequest); i {
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
		file_kessel_inventory_v1beta1_rhel_hosts_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateRhelHostResponse); i {
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
		file_kessel_inventory_v1beta1_rhel_hosts_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateRhelHostRequest); i {
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
		file_kessel_inventory_v1beta1_rhel_hosts_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateRhelHostResponse); i {
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
		file_kessel_inventory_v1beta1_rhel_hosts_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteRhelHostRequest); i {
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
		file_kessel_inventory_v1beta1_rhel_hosts_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteRhelHostResponse); i {
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
			RawDescriptor: file_kessel_inventory_v1beta1_rhel_hosts_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kessel_inventory_v1beta1_rhel_hosts_service_proto_goTypes,
		DependencyIndexes: file_kessel_inventory_v1beta1_rhel_hosts_service_proto_depIdxs,
		MessageInfos:      file_kessel_inventory_v1beta1_rhel_hosts_service_proto_msgTypes,
	}.Build()
	File_kessel_inventory_v1beta1_rhel_hosts_service_proto = out.File
	file_kessel_inventory_v1beta1_rhel_hosts_service_proto_rawDesc = nil
	file_kessel_inventory_v1beta1_rhel_hosts_service_proto_goTypes = nil
	file_kessel_inventory_v1beta1_rhel_hosts_service_proto_depIdxs = nil
}
