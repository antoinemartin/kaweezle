// Generate with the following command:
// protoc --go-grpc_out=. --go_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative *.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.1
// source: elevated.proto

package config

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

type AddRouteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FixedAddress string `protobuf:"bytes,1,opt,name=fixedAddress,proto3" json:"fixedAddress,omitempty"`
	Netmask      string `protobuf:"bytes,2,opt,name=netmask,proto3" json:"netmask,omitempty"`
	Gateway      string `protobuf:"bytes,3,opt,name=gateway,proto3" json:"gateway,omitempty"`
}

func (x *AddRouteRequest) Reset() {
	*x = AddRouteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_elevated_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRouteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRouteRequest) ProtoMessage() {}

func (x *AddRouteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_elevated_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRouteRequest.ProtoReflect.Descriptor instead.
func (*AddRouteRequest) Descriptor() ([]byte, []int) {
	return file_elevated_proto_rawDescGZIP(), []int{0}
}

func (x *AddRouteRequest) GetFixedAddress() string {
	if x != nil {
		return x.FixedAddress
	}
	return ""
}

func (x *AddRouteRequest) GetNetmask() string {
	if x != nil {
		return x.Netmask
	}
	return ""
}

func (x *AddRouteRequest) GetGateway() string {
	if x != nil {
		return x.Gateway
	}
	return ""
}

type AddRouteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddRouteResponse) Reset() {
	*x = AddRouteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_elevated_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRouteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRouteResponse) ProtoMessage() {}

func (x *AddRouteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_elevated_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRouteResponse.ProtoReflect.Descriptor instead.
func (*AddRouteResponse) Descriptor() ([]byte, []int) {
	return file_elevated_proto_rawDescGZIP(), []int{1}
}

type RemoveRouteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FixedAddress string `protobuf:"bytes,1,opt,name=fixedAddress,proto3" json:"fixedAddress,omitempty"`
}

func (x *RemoveRouteRequest) Reset() {
	*x = RemoveRouteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_elevated_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveRouteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRouteRequest) ProtoMessage() {}

func (x *RemoveRouteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_elevated_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRouteRequest.ProtoReflect.Descriptor instead.
func (*RemoveRouteRequest) Descriptor() ([]byte, []int) {
	return file_elevated_proto_rawDescGZIP(), []int{2}
}

func (x *RemoveRouteRequest) GetFixedAddress() string {
	if x != nil {
		return x.FixedAddress
	}
	return ""
}

type RemoveRouteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveRouteResponse) Reset() {
	*x = RemoveRouteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_elevated_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveRouteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRouteResponse) ProtoMessage() {}

func (x *RemoveRouteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_elevated_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRouteResponse.ProtoReflect.Descriptor instead.
func (*RemoveRouteResponse) Descriptor() ([]byte, []int) {
	return file_elevated_proto_rawDescGZIP(), []int{3}
}

type ConfigureDomainsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DistributionName string   `protobuf:"bytes,1,opt,name=distributionName,proto3" json:"distributionName,omitempty"`
	IpAddress        string   `protobuf:"bytes,2,opt,name=ipAddress,proto3" json:"ipAddress,omitempty"`
	Domains          []string `protobuf:"bytes,3,rep,name=domains,proto3" json:"domains,omitempty"`
	Remove           bool     `protobuf:"varint,4,opt,name=remove,proto3" json:"remove,omitempty"`
}

func (x *ConfigureDomainsRequest) Reset() {
	*x = ConfigureDomainsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_elevated_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigureDomainsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigureDomainsRequest) ProtoMessage() {}

func (x *ConfigureDomainsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_elevated_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigureDomainsRequest.ProtoReflect.Descriptor instead.
func (*ConfigureDomainsRequest) Descriptor() ([]byte, []int) {
	return file_elevated_proto_rawDescGZIP(), []int{4}
}

func (x *ConfigureDomainsRequest) GetDistributionName() string {
	if x != nil {
		return x.DistributionName
	}
	return ""
}

func (x *ConfigureDomainsRequest) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

func (x *ConfigureDomainsRequest) GetDomains() []string {
	if x != nil {
		return x.Domains
	}
	return nil
}

func (x *ConfigureDomainsRequest) GetRemove() bool {
	if x != nil {
		return x.Remove
	}
	return false
}

type ConfigureDomainsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domains []string `protobuf:"bytes,1,rep,name=domains,proto3" json:"domains,omitempty"`
}

func (x *ConfigureDomainsResponse) Reset() {
	*x = ConfigureDomainsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_elevated_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigureDomainsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigureDomainsResponse) ProtoMessage() {}

func (x *ConfigureDomainsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_elevated_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigureDomainsResponse.ProtoReflect.Descriptor instead.
func (*ConfigureDomainsResponse) Descriptor() ([]byte, []int) {
	return file_elevated_proto_rawDescGZIP(), []int{5}
}

func (x *ConfigureDomainsResponse) GetDomains() []string {
	if x != nil {
		return x.Domains
	}
	return nil
}

type StopRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timeout int32 `protobuf:"varint,1,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *StopRequest) Reset() {
	*x = StopRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_elevated_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopRequest) ProtoMessage() {}

func (x *StopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_elevated_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopRequest.ProtoReflect.Descriptor instead.
func (*StopRequest) Descriptor() ([]byte, []int) {
	return file_elevated_proto_rawDescGZIP(), []int{6}
}

func (x *StopRequest) GetTimeout() int32 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

type StopResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StopResponse) Reset() {
	*x = StopResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_elevated_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopResponse) ProtoMessage() {}

func (x *StopResponse) ProtoReflect() protoreflect.Message {
	mi := &file_elevated_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopResponse.ProtoReflect.Descriptor instead.
func (*StopResponse) Descriptor() ([]byte, []int) {
	return file_elevated_proto_rawDescGZIP(), []int{7}
}

var File_elevated_proto protoreflect.FileDescriptor

var file_elevated_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x65, 0x6c, 0x65, 0x76, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x69, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x66,
	0x69, 0x78, 0x65, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x66, 0x69, 0x78, 0x65, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6e, 0x65, 0x74, 0x6d, 0x61, 0x73, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x22, 0x12, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x38, 0x0a, 0x12, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a,
	0x0c, 0x66, 0x69, 0x78, 0x65, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x69, 0x78, 0x65, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x22, 0x15, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x95, 0x01, 0x0a, 0x17, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x07, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x22, 0x34, 0x0a, 0x18, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x44, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x22, 0x27, 0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22,
	0x0e, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0xb0, 0x02, 0x0a, 0x15, 0x45, 0x6c, 0x65, 0x76, 0x61, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x08, 0x41, 0x64, 0x64,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x41,
	0x64, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0b, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x12, 0x1f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a,
	0x04, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x13, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53,
	0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6b, 0x61, 0x77, 0x65, 0x65, 0x7a, 0x6c, 0x65, 0x2f, 0x6b, 0x61, 0x77, 0x65, 0x65, 0x7a,
	0x6c, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_elevated_proto_rawDescOnce sync.Once
	file_elevated_proto_rawDescData = file_elevated_proto_rawDesc
)

func file_elevated_proto_rawDescGZIP() []byte {
	file_elevated_proto_rawDescOnce.Do(func() {
		file_elevated_proto_rawDescData = protoimpl.X.CompressGZIP(file_elevated_proto_rawDescData)
	})
	return file_elevated_proto_rawDescData
}

var file_elevated_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_elevated_proto_goTypes = []interface{}{
	(*AddRouteRequest)(nil),          // 0: config.AddRouteRequest
	(*AddRouteResponse)(nil),         // 1: config.AddRouteResponse
	(*RemoveRouteRequest)(nil),       // 2: config.RemoveRouteRequest
	(*RemoveRouteResponse)(nil),      // 3: config.RemoveRouteResponse
	(*ConfigureDomainsRequest)(nil),  // 4: config.ConfigureDomainsRequest
	(*ConfigureDomainsResponse)(nil), // 5: config.ConfigureDomainsResponse
	(*StopRequest)(nil),              // 6: config.StopRequest
	(*StopResponse)(nil),             // 7: config.StopResponse
}
var file_elevated_proto_depIdxs = []int32{
	0, // 0: config.ElevatedConfiguration.AddRoute:input_type -> config.AddRouteRequest
	2, // 1: config.ElevatedConfiguration.RemoveRoute:input_type -> config.RemoveRouteRequest
	4, // 2: config.ElevatedConfiguration.ConfigureDomains:input_type -> config.ConfigureDomainsRequest
	6, // 3: config.ElevatedConfiguration.Stop:input_type -> config.StopRequest
	1, // 4: config.ElevatedConfiguration.AddRoute:output_type -> config.AddRouteResponse
	3, // 5: config.ElevatedConfiguration.RemoveRoute:output_type -> config.RemoveRouteResponse
	5, // 6: config.ElevatedConfiguration.ConfigureDomains:output_type -> config.ConfigureDomainsResponse
	7, // 7: config.ElevatedConfiguration.Stop:output_type -> config.StopResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_elevated_proto_init() }
func file_elevated_proto_init() {
	if File_elevated_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_elevated_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRouteRequest); i {
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
		file_elevated_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRouteResponse); i {
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
		file_elevated_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveRouteRequest); i {
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
		file_elevated_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveRouteResponse); i {
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
		file_elevated_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigureDomainsRequest); i {
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
		file_elevated_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigureDomainsResponse); i {
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
		file_elevated_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopRequest); i {
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
		file_elevated_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopResponse); i {
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
			RawDescriptor: file_elevated_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_elevated_proto_goTypes,
		DependencyIndexes: file_elevated_proto_depIdxs,
		MessageInfos:      file_elevated_proto_msgTypes,
	}.Build()
	File_elevated_proto = out.File
	file_elevated_proto_rawDesc = nil
	file_elevated_proto_goTypes = nil
	file_elevated_proto_depIdxs = nil
}
