// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.25.1
// source: node/v1/info.proto

package deepagent_node_v1

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

type CpuFeatureCacheKeyValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *CpuFeatureCacheKeyValue) Reset() {
	*x = CpuFeatureCacheKeyValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_v1_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CpuFeatureCacheKeyValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CpuFeatureCacheKeyValue) ProtoMessage() {}

func (x *CpuFeatureCacheKeyValue) ProtoReflect() protoreflect.Message {
	mi := &file_node_v1_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CpuFeatureCacheKeyValue.ProtoReflect.Descriptor instead.
func (*CpuFeatureCacheKeyValue) Descriptor() ([]byte, []int) {
	return file_node_v1_info_proto_rawDescGZIP(), []int{0}
}

func (x *CpuFeatureCacheKeyValue) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *CpuFeatureCacheKeyValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// "unknown service deepagent.node.v1.NodeCollectorService"
// CPU 特性
type CpuFeature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Arch              string                     `protobuf:"bytes,1,opt,name=arch,proto3" json:"arch,omitempty"`
	Brand             string                     `protobuf:"bytes,2,opt,name=brand,proto3" json:"brand,omitempty"`
	Family            string                     `protobuf:"bytes,3,opt,name=family,proto3" json:"family,omitempty"`
	Model             string                     `protobuf:"bytes,4,opt,name=model,proto3" json:"model,omitempty"`
	Stepping          string                     `protobuf:"bytes,5,opt,name=stepping,proto3" json:"stepping,omitempty"`
	Uarch             string                     `protobuf:"bytes,6,opt,name=uarch,proto3" json:"uarch,omitempty"`
	Implementer       string                     `protobuf:"bytes,7,opt,name=implementer,proto3" json:"implementer,omitempty"`
	Architecture      string                     `protobuf:"bytes,8,opt,name=architecture,proto3" json:"architecture,omitempty"`
	Variant           string                     `protobuf:"bytes,9,opt,name=variant,proto3" json:"variant,omitempty"`
	Part              string                     `protobuf:"bytes,10,opt,name=part,proto3" json:"part,omitempty"`
	Revision          string                     `protobuf:"bytes,11,opt,name=revision,proto3" json:"revision,omitempty"`
	Platform          string                     `protobuf:"bytes,12,opt,name=platform,proto3" json:"platform,omitempty"`
	Machine           string                     `protobuf:"bytes,13,opt,name=machine,proto3" json:"machine,omitempty"`
	Cpu               string                     `protobuf:"bytes,14,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Instruction       string                     `protobuf:"bytes,15,opt,name=instruction,proto3" json:"instruction,omitempty"`
	Microarchitecture string                     `protobuf:"bytes,16,opt,name=microarchitecture,proto3" json:"microarchitecture,omitempty"`
	Processors        string                     `protobuf:"bytes,17,opt,name=processors,proto3" json:"processors,omitempty"`
	Vendor            string                     `protobuf:"bytes,18,opt,name=vendor,proto3" json:"vendor,omitempty"`
	Flags             []string                   `protobuf:"bytes,19,rep,name=flags,proto3" json:"flags,omitempty"`
	Caches            []*CpuFeatureCacheKeyValue `protobuf:"bytes,20,rep,name=caches,proto3" json:"caches,omitempty"`
}

func (x *CpuFeature) Reset() {
	*x = CpuFeature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_v1_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CpuFeature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CpuFeature) ProtoMessage() {}

func (x *CpuFeature) ProtoReflect() protoreflect.Message {
	mi := &file_node_v1_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CpuFeature.ProtoReflect.Descriptor instead.
func (*CpuFeature) Descriptor() ([]byte, []int) {
	return file_node_v1_info_proto_rawDescGZIP(), []int{1}
}

func (x *CpuFeature) GetArch() string {
	if x != nil {
		return x.Arch
	}
	return ""
}

func (x *CpuFeature) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *CpuFeature) GetFamily() string {
	if x != nil {
		return x.Family
	}
	return ""
}

func (x *CpuFeature) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *CpuFeature) GetStepping() string {
	if x != nil {
		return x.Stepping
	}
	return ""
}

func (x *CpuFeature) GetUarch() string {
	if x != nil {
		return x.Uarch
	}
	return ""
}

func (x *CpuFeature) GetImplementer() string {
	if x != nil {
		return x.Implementer
	}
	return ""
}

func (x *CpuFeature) GetArchitecture() string {
	if x != nil {
		return x.Architecture
	}
	return ""
}

func (x *CpuFeature) GetVariant() string {
	if x != nil {
		return x.Variant
	}
	return ""
}

func (x *CpuFeature) GetPart() string {
	if x != nil {
		return x.Part
	}
	return ""
}

func (x *CpuFeature) GetRevision() string {
	if x != nil {
		return x.Revision
	}
	return ""
}

func (x *CpuFeature) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *CpuFeature) GetMachine() string {
	if x != nil {
		return x.Machine
	}
	return ""
}

func (x *CpuFeature) GetCpu() string {
	if x != nil {
		return x.Cpu
	}
	return ""
}

func (x *CpuFeature) GetInstruction() string {
	if x != nil {
		return x.Instruction
	}
	return ""
}

func (x *CpuFeature) GetMicroarchitecture() string {
	if x != nil {
		return x.Microarchitecture
	}
	return ""
}

func (x *CpuFeature) GetProcessors() string {
	if x != nil {
		return x.Processors
	}
	return ""
}

func (x *CpuFeature) GetVendor() string {
	if x != nil {
		return x.Vendor
	}
	return ""
}

func (x *CpuFeature) GetFlags() []string {
	if x != nil {
		return x.Flags
	}
	return nil
}

func (x *CpuFeature) GetCaches() []*CpuFeatureCacheKeyValue {
	if x != nil {
		return x.Caches
	}
	return nil
}

type NodeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyUuid string `protobuf:"bytes,1,opt,name=companyUuid,proto3" json:"companyUuid,omitempty"`
	// 唯一的数据id
	ObjectId uint64 `protobuf:"varint,2,opt,name=objectId,proto3" json:"objectId,omitempty"`
	// 主机名字
	HostName string `protobuf:"bytes,3,opt,name=hostName,proto3" json:"hostName,omitempty"`
	// CPU核心功能
	CpuFeature *CpuFeature `protobuf:"bytes,4,opt,name=cpu_feature,json=cpuFeature,proto3" json:"cpu_feature,omitempty"`
	// linux内核操作系统
	SystemOs string `protobuf:"bytes,5,opt,name=systemOs,proto3" json:"systemOs,omitempty"`
	// 内核版本
	SystemVersion string `protobuf:"bytes,6,opt,name=systemVersion,proto3" json:"systemVersion,omitempty"`
	// 主机类型是kvm还是虚拟机
	SystemType string `protobuf:"bytes,7,opt,name=systemType,proto3" json:"systemType,omitempty"`
	// 属于的集群
	Cluster string `protobuf:"bytes,8,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// 所属平台
	Platform string `protobuf:"bytes,9,opt,name=platform,proto3" json:"platform,omitempty"`
}

func (x *NodeInfo) Reset() {
	*x = NodeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_v1_info_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeInfo) ProtoMessage() {}

func (x *NodeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_node_v1_info_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeInfo.ProtoReflect.Descriptor instead.
func (*NodeInfo) Descriptor() ([]byte, []int) {
	return file_node_v1_info_proto_rawDescGZIP(), []int{2}
}

func (x *NodeInfo) GetCompanyUuid() string {
	if x != nil {
		return x.CompanyUuid
	}
	return ""
}

func (x *NodeInfo) GetObjectId() uint64 {
	if x != nil {
		return x.ObjectId
	}
	return 0
}

func (x *NodeInfo) GetHostName() string {
	if x != nil {
		return x.HostName
	}
	return ""
}

func (x *NodeInfo) GetCpuFeature() *CpuFeature {
	if x != nil {
		return x.CpuFeature
	}
	return nil
}

func (x *NodeInfo) GetSystemOs() string {
	if x != nil {
		return x.SystemOs
	}
	return ""
}

func (x *NodeInfo) GetSystemVersion() string {
	if x != nil {
		return x.SystemVersion
	}
	return ""
}

func (x *NodeInfo) GetSystemType() string {
	if x != nil {
		return x.SystemType
	}
	return ""
}

func (x *NodeInfo) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *NodeInfo) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

type NodeUsage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 唯一的数据id
	ObjectId uint64 `protobuf:"varint,1,opt,name=objectId,proto3" json:"objectId,omitempty"`
	// cpu使用率
	CpuValue uint64 `protobuf:"varint,2,opt,name=cpuValue,proto3" json:"cpuValue,omitempty"`
	// memory使用率
	MemoryValue uint64 `protobuf:"varint,3,opt,name=memoryValue,proto3" json:"memoryValue,omitempty"`
}

func (x *NodeUsage) Reset() {
	*x = NodeUsage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_v1_info_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeUsage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeUsage) ProtoMessage() {}

func (x *NodeUsage) ProtoReflect() protoreflect.Message {
	mi := &file_node_v1_info_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeUsage.ProtoReflect.Descriptor instead.
func (*NodeUsage) Descriptor() ([]byte, []int) {
	return file_node_v1_info_proto_rawDescGZIP(), []int{3}
}

func (x *NodeUsage) GetObjectId() uint64 {
	if x != nil {
		return x.ObjectId
	}
	return 0
}

func (x *NodeUsage) GetCpuValue() uint64 {
	if x != nil {
		return x.CpuValue
	}
	return 0
}

func (x *NodeUsage) GetMemoryValue() uint64 {
	if x != nil {
		return x.MemoryValue
	}
	return 0
}

type NodeRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *NodeRes) Reset() {
	*x = NodeRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_v1_info_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeRes) ProtoMessage() {}

func (x *NodeRes) ProtoReflect() protoreflect.Message {
	mi := &file_node_v1_info_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeRes.ProtoReflect.Descriptor instead.
func (*NodeRes) Descriptor() ([]byte, []int) {
	return file_node_v1_info_proto_rawDescGZIP(), []int{4}
}

func (x *NodeRes) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_node_v1_info_proto protoreflect.FileDescriptor

var file_node_v1_info_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x64, 0x65, 0x65, 0x70, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x41, 0x0a, 0x17, 0x43, 0x70, 0x75, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xd0, 0x04, 0x0a, 0x0a, 0x43,
	0x70, 0x75, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x63,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x63, 0x68, 0x12, 0x14, 0x0a,
	0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x72,
	0x61, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x65, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x65, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a,
	0x05, 0x75, 0x61, 0x72, 0x63, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x75, 0x61,
	0x72, 0x63, 0x68, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x72, 0x63,
	0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x61, 0x72,
	0x69, 0x61, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x61, 0x72, 0x69,
	0x61, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x72, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x61, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x70, 0x75,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x70, 0x75, 0x12, 0x20, 0x0a, 0x0b, 0x69,
	0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a,
	0x11, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x61,
	0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x73, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x76,
	0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x65, 0x6e,
	0x64, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x13, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x42, 0x0a, 0x06, 0x63, 0x61, 0x63,
	0x68, 0x65, 0x73, 0x18, 0x14, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x64, 0x65, 0x65, 0x70,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x70,
	0x75, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x4b, 0x65, 0x79,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x63, 0x61, 0x63, 0x68, 0x65, 0x73, 0x22, 0xbc, 0x02,
	0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x55, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x55, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x63, 0x70, 0x75, 0x5f, 0x66, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x64, 0x65, 0x65, 0x70,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x70,
	0x75, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x0a, 0x63, 0x70, 0x75, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x4f, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x4f, 0x73,
	0x12, 0x24, 0x0a, 0x0d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x22, 0x65, 0x0a, 0x09,
	0x4e, 0x6f, 0x64, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x70, 0x75, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x63, 0x70, 0x75, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x19, 0x0a, 0x07, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x32, 0x9e,
	0x01, 0x0a, 0x14, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x03, 0x72, 0x65, 0x67, 0x12, 0x1b,
	0x2e, 0x64, 0x65, 0x65, 0x70, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x1a, 0x2e, 0x64, 0x65,
	0x65, 0x70, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x06, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x64, 0x65, 0x65, 0x70, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x55, 0x73, 0x61, 0x67,
	0x65, 0x1a, 0x1a, 0x2e, 0x64, 0x65, 0x65, 0x70, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x6e, 0x6f,
	0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42,
	0x13, 0x5a, 0x11, 0x64, 0x65, 0x65, 0x70, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x6e, 0x6f, 0x64,
	0x65, 0x2e, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_node_v1_info_proto_rawDescOnce sync.Once
	file_node_v1_info_proto_rawDescData = file_node_v1_info_proto_rawDesc
)

func file_node_v1_info_proto_rawDescGZIP() []byte {
	file_node_v1_info_proto_rawDescOnce.Do(func() {
		file_node_v1_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_node_v1_info_proto_rawDescData)
	})
	return file_node_v1_info_proto_rawDescData
}

var file_node_v1_info_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_node_v1_info_proto_goTypes = []interface{}{
	(*CpuFeatureCacheKeyValue)(nil), // 0: deepagent.node.v1.CpuFeatureCacheKeyValue
	(*CpuFeature)(nil),              // 1: deepagent.node.v1.CpuFeature
	(*NodeInfo)(nil),                // 2: deepagent.node.v1.NodeInfo
	(*NodeUsage)(nil),               // 3: deepagent.node.v1.NodeUsage
	(*NodeRes)(nil),                 // 4: deepagent.node.v1.NodeRes
}
var file_node_v1_info_proto_depIdxs = []int32{
	0, // 0: deepagent.node.v1.CpuFeature.caches:type_name -> deepagent.node.v1.CpuFeatureCacheKeyValue
	1, // 1: deepagent.node.v1.NodeInfo.cpu_feature:type_name -> deepagent.node.v1.CpuFeature
	2, // 2: deepagent.node.v1.NodeCollectorService.reg:input_type -> deepagent.node.v1.NodeInfo
	3, // 3: deepagent.node.v1.NodeCollectorService.update:input_type -> deepagent.node.v1.NodeUsage
	4, // 4: deepagent.node.v1.NodeCollectorService.reg:output_type -> deepagent.node.v1.NodeRes
	4, // 5: deepagent.node.v1.NodeCollectorService.update:output_type -> deepagent.node.v1.NodeRes
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_node_v1_info_proto_init() }
func file_node_v1_info_proto_init() {
	if File_node_v1_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_node_v1_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CpuFeatureCacheKeyValue); i {
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
		file_node_v1_info_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CpuFeature); i {
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
		file_node_v1_info_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeInfo); i {
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
		file_node_v1_info_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeUsage); i {
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
		file_node_v1_info_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeRes); i {
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
			RawDescriptor: file_node_v1_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_node_v1_info_proto_goTypes,
		DependencyIndexes: file_node_v1_info_proto_depIdxs,
		MessageInfos:      file_node_v1_info_proto_msgTypes,
	}.Build()
	File_node_v1_info_proto = out.File
	file_node_v1_info_proto_rawDesc = nil
	file_node_v1_info_proto_goTypes = nil
	file_node_v1_info_proto_depIdxs = nil
}
