// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.12.4
// source: appix/v1/clusters.proto

package v1

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

// gratos::model
type Cluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Cluster) Reset() {
	*x = Cluster{}
	mi := &file_appix_v1_clusters_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Cluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cluster) ProtoMessage() {}

func (x *Cluster) ProtoReflect() protoreflect.Message {
	mi := &file_appix_v1_clusters_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cluster.ProtoReflect.Descriptor instead.
func (*Cluster) Descriptor() ([]byte, []int) {
	return file_appix_v1_clusters_proto_rawDescGZIP(), []int{0}
}

func (x *Cluster) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Cluster) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Cluster) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CreateClustersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clusters []*Cluster `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
}

func (x *CreateClustersRequest) Reset() {
	*x = CreateClustersRequest{}
	mi := &file_appix_v1_clusters_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateClustersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClustersRequest) ProtoMessage() {}

func (x *CreateClustersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_appix_v1_clusters_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClustersRequest.ProtoReflect.Descriptor instead.
func (*CreateClustersRequest) Descriptor() ([]byte, []int) {
	return file_appix_v1_clusters_proto_rawDescGZIP(), []int{1}
}

func (x *CreateClustersRequest) GetClusters() []*Cluster {
	if x != nil {
		return x.Clusters
	}
	return nil
}

type CreateClustersReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Code    int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Action  string `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
}

func (x *CreateClustersReply) Reset() {
	*x = CreateClustersReply{}
	mi := &file_appix_v1_clusters_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateClustersReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClustersReply) ProtoMessage() {}

func (x *CreateClustersReply) ProtoReflect() protoreflect.Message {
	mi := &file_appix_v1_clusters_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClustersReply.ProtoReflect.Descriptor instead.
func (*CreateClustersReply) Descriptor() ([]byte, []int) {
	return file_appix_v1_clusters_proto_rawDescGZIP(), []int{2}
}

func (x *CreateClustersReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CreateClustersReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CreateClustersReply) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

type UpdateClustersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clusters []*Cluster `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
}

func (x *UpdateClustersRequest) Reset() {
	*x = UpdateClustersRequest{}
	mi := &file_appix_v1_clusters_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateClustersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateClustersRequest) ProtoMessage() {}

func (x *UpdateClustersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_appix_v1_clusters_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateClustersRequest.ProtoReflect.Descriptor instead.
func (*UpdateClustersRequest) Descriptor() ([]byte, []int) {
	return file_appix_v1_clusters_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateClustersRequest) GetClusters() []*Cluster {
	if x != nil {
		return x.Clusters
	}
	return nil
}

type UpdateClustersReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Code    int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Action  string `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
}

func (x *UpdateClustersReply) Reset() {
	*x = UpdateClustersReply{}
	mi := &file_appix_v1_clusters_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateClustersReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateClustersReply) ProtoMessage() {}

func (x *UpdateClustersReply) ProtoReflect() protoreflect.Message {
	mi := &file_appix_v1_clusters_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateClustersReply.ProtoReflect.Descriptor instead.
func (*UpdateClustersReply) Descriptor() ([]byte, []int) {
	return file_appix_v1_clusters_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateClustersReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UpdateClustersReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *UpdateClustersReply) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

type DeleteClustersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []uint32 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *DeleteClustersRequest) Reset() {
	*x = DeleteClustersRequest{}
	mi := &file_appix_v1_clusters_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteClustersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteClustersRequest) ProtoMessage() {}

func (x *DeleteClustersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_appix_v1_clusters_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteClustersRequest.ProtoReflect.Descriptor instead.
func (*DeleteClustersRequest) Descriptor() ([]byte, []int) {
	return file_appix_v1_clusters_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteClustersRequest) GetIds() []uint32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type DeleteClustersReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Code    int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Action  string `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
}

func (x *DeleteClustersReply) Reset() {
	*x = DeleteClustersReply{}
	mi := &file_appix_v1_clusters_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteClustersReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteClustersReply) ProtoMessage() {}

func (x *DeleteClustersReply) ProtoReflect() protoreflect.Message {
	mi := &file_appix_v1_clusters_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteClustersReply.ProtoReflect.Descriptor instead.
func (*DeleteClustersReply) Descriptor() ([]byte, []int) {
	return file_appix_v1_clusters_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteClustersReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *DeleteClustersReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *DeleteClustersReply) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

type GetClustersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetClustersRequest) Reset() {
	*x = GetClustersRequest{}
	mi := &file_appix_v1_clusters_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetClustersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClustersRequest) ProtoMessage() {}

func (x *GetClustersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_appix_v1_clusters_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClustersRequest.ProtoReflect.Descriptor instead.
func (*GetClustersRequest) Descriptor() ([]byte, []int) {
	return file_appix_v1_clusters_proto_rawDescGZIP(), []int{7}
}

func (x *GetClustersRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetClustersReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Code    int32    `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Action  string   `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
	Cluster *Cluster `protobuf:"bytes,4,opt,name=cluster,proto3" json:"cluster,omitempty"`
}

func (x *GetClustersReply) Reset() {
	*x = GetClustersReply{}
	mi := &file_appix_v1_clusters_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetClustersReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClustersReply) ProtoMessage() {}

func (x *GetClustersReply) ProtoReflect() protoreflect.Message {
	mi := &file_appix_v1_clusters_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClustersReply.ProtoReflect.Descriptor instead.
func (*GetClustersReply) Descriptor() ([]byte, []int) {
	return file_appix_v1_clusters_proto_rawDescGZIP(), []int{8}
}

func (x *GetClustersReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetClustersReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetClustersReply) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *GetClustersReply) GetCluster() *Cluster {
	if x != nil {
		return x.Cluster
	}
	return nil
}

// gratos::model
type ListClustersFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     uint32   `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize uint32   `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Names    []string `protobuf:"bytes,3,rep,name=names,proto3" json:"names,omitempty"`
	Ids      []uint32 `protobuf:"varint,4,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *ListClustersFilter) Reset() {
	*x = ListClustersFilter{}
	mi := &file_appix_v1_clusters_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListClustersFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListClustersFilter) ProtoMessage() {}

func (x *ListClustersFilter) ProtoReflect() protoreflect.Message {
	mi := &file_appix_v1_clusters_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListClustersFilter.ProtoReflect.Descriptor instead.
func (*ListClustersFilter) Descriptor() ([]byte, []int) {
	return file_appix_v1_clusters_proto_rawDescGZIP(), []int{9}
}

func (x *ListClustersFilter) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListClustersFilter) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListClustersFilter) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

func (x *ListClustersFilter) GetIds() []uint32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type ListClustersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter *ListClustersFilter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ListClustersRequest) Reset() {
	*x = ListClustersRequest{}
	mi := &file_appix_v1_clusters_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListClustersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListClustersRequest) ProtoMessage() {}

func (x *ListClustersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_appix_v1_clusters_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListClustersRequest.ProtoReflect.Descriptor instead.
func (*ListClustersRequest) Descriptor() ([]byte, []int) {
	return file_appix_v1_clusters_proto_rawDescGZIP(), []int{10}
}

func (x *ListClustersRequest) GetFilter() *ListClustersFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type ListClustersReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message  string     `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Code     int32      `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Action   string     `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
	Clusters []*Cluster `protobuf:"bytes,4,rep,name=clusters,proto3" json:"clusters,omitempty"`
}

func (x *ListClustersReply) Reset() {
	*x = ListClustersReply{}
	mi := &file_appix_v1_clusters_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListClustersReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListClustersReply) ProtoMessage() {}

func (x *ListClustersReply) ProtoReflect() protoreflect.Message {
	mi := &file_appix_v1_clusters_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListClustersReply.ProtoReflect.Descriptor instead.
func (*ListClustersReply) Descriptor() ([]byte, []int) {
	return file_appix_v1_clusters_proto_rawDescGZIP(), []int{11}
}

func (x *ListClustersReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ListClustersReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ListClustersReply) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *ListClustersReply) GetClusters() []*Cluster {
	if x != nil {
		return x.Clusters
	}
	return nil
}

var File_appix_v1_clusters_proto protoreflect.FileDescriptor

var file_appix_v1_clusters_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x70, 0x69, 0x78, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x70, 0x70, 0x69, 0x78, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x07, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4a, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x31, 0x0a, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x69, 0x78, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x73, 0x22, 0x5b, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x4a, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x08, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x61, 0x70, 0x70, 0x69, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x22, 0x5b, 0x0a, 0x13, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x29, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x03,
	0x69, 0x64, 0x73, 0x22, 0x5b, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x24, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x89, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2f, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x69, 0x78, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x22, 0x6d, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x03, 0x69, 0x64,
	0x73, 0x22, 0x4f, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x70, 0x70, 0x69, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x22, 0x8c, 0x01, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x31,
	0x0a, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x69, 0x78, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x73, 0x32, 0xea, 0x04, 0x0a, 0x08, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x7c,
	0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73,
	0x12, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x69, 0x78, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x69,
	0x78, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c,
	0x3a, 0x01, 0x2a, 0x22, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x7c, 0x0a, 0x0e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x23,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x69, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x69, 0x78, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01,
	0x2a, 0x22, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x73, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x7c, 0x0a, 0x0e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x23, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x69, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x69, 0x78, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x22,
	0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x73, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x6e, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70,
	0x70, 0x69, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x70, 0x70, 0x69, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x17, 0x12, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x74, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x70, 0x70, 0x69, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x70, 0x70, 0x69, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x20, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x22, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x27,
	0x0a, 0x0c, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x69, 0x78, 0x2e, 0x76, 0x31, 0x50, 0x01,
	0x5a, 0x15, 0x61, 0x70, 0x70, 0x69, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x69,
	0x78, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_appix_v1_clusters_proto_rawDescOnce sync.Once
	file_appix_v1_clusters_proto_rawDescData = file_appix_v1_clusters_proto_rawDesc
)

func file_appix_v1_clusters_proto_rawDescGZIP() []byte {
	file_appix_v1_clusters_proto_rawDescOnce.Do(func() {
		file_appix_v1_clusters_proto_rawDescData = protoimpl.X.CompressGZIP(file_appix_v1_clusters_proto_rawDescData)
	})
	return file_appix_v1_clusters_proto_rawDescData
}

var file_appix_v1_clusters_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_appix_v1_clusters_proto_goTypes = []any{
	(*Cluster)(nil),               // 0: api.appix.v1.Cluster
	(*CreateClustersRequest)(nil), // 1: api.appix.v1.CreateClustersRequest
	(*CreateClustersReply)(nil),   // 2: api.appix.v1.CreateClustersReply
	(*UpdateClustersRequest)(nil), // 3: api.appix.v1.UpdateClustersRequest
	(*UpdateClustersReply)(nil),   // 4: api.appix.v1.UpdateClustersReply
	(*DeleteClustersRequest)(nil), // 5: api.appix.v1.DeleteClustersRequest
	(*DeleteClustersReply)(nil),   // 6: api.appix.v1.DeleteClustersReply
	(*GetClustersRequest)(nil),    // 7: api.appix.v1.GetClustersRequest
	(*GetClustersReply)(nil),      // 8: api.appix.v1.GetClustersReply
	(*ListClustersFilter)(nil),    // 9: api.appix.v1.ListClustersFilter
	(*ListClustersRequest)(nil),   // 10: api.appix.v1.ListClustersRequest
	(*ListClustersReply)(nil),     // 11: api.appix.v1.ListClustersReply
}
var file_appix_v1_clusters_proto_depIdxs = []int32{
	0,  // 0: api.appix.v1.CreateClustersRequest.clusters:type_name -> api.appix.v1.Cluster
	0,  // 1: api.appix.v1.UpdateClustersRequest.clusters:type_name -> api.appix.v1.Cluster
	0,  // 2: api.appix.v1.GetClustersReply.cluster:type_name -> api.appix.v1.Cluster
	9,  // 3: api.appix.v1.ListClustersRequest.filter:type_name -> api.appix.v1.ListClustersFilter
	0,  // 4: api.appix.v1.ListClustersReply.clusters:type_name -> api.appix.v1.Cluster
	1,  // 5: api.appix.v1.Clusters.CreateClusters:input_type -> api.appix.v1.CreateClustersRequest
	3,  // 6: api.appix.v1.Clusters.UpdateClusters:input_type -> api.appix.v1.UpdateClustersRequest
	5,  // 7: api.appix.v1.Clusters.DeleteClusters:input_type -> api.appix.v1.DeleteClustersRequest
	7,  // 8: api.appix.v1.Clusters.GetClusters:input_type -> api.appix.v1.GetClustersRequest
	10, // 9: api.appix.v1.Clusters.ListClusters:input_type -> api.appix.v1.ListClustersRequest
	2,  // 10: api.appix.v1.Clusters.CreateClusters:output_type -> api.appix.v1.CreateClustersReply
	4,  // 11: api.appix.v1.Clusters.UpdateClusters:output_type -> api.appix.v1.UpdateClustersReply
	6,  // 12: api.appix.v1.Clusters.DeleteClusters:output_type -> api.appix.v1.DeleteClustersReply
	8,  // 13: api.appix.v1.Clusters.GetClusters:output_type -> api.appix.v1.GetClustersReply
	11, // 14: api.appix.v1.Clusters.ListClusters:output_type -> api.appix.v1.ListClustersReply
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_appix_v1_clusters_proto_init() }
func file_appix_v1_clusters_proto_init() {
	if File_appix_v1_clusters_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_appix_v1_clusters_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_appix_v1_clusters_proto_goTypes,
		DependencyIndexes: file_appix_v1_clusters_proto_depIdxs,
		MessageInfos:      file_appix_v1_clusters_proto_msgTypes,
	}.Build()
	File_appix_v1_clusters_proto = out.File
	file_appix_v1_clusters_proto_rawDesc = nil
	file_appix_v1_clusters_proto_goTypes = nil
	file_appix_v1_clusters_proto_depIdxs = nil
}
