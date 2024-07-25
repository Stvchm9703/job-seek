// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: data-store.proto

package protos

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

type DataStoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DataStoreRequest) Reset() {
	*x = DataStoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_store_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataStoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataStoreRequest) ProtoMessage() {}

func (x *DataStoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_data_store_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataStoreRequest.ProtoReflect.Descriptor instead.
func (*DataStoreRequest) Descriptor() ([]byte, []int) {
	return file_data_store_proto_rawDescGZIP(), []int{0}
}

type DataStoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DataStoreResponse) Reset() {
	*x = DataStoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_store_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataStoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataStoreResponse) ProtoMessage() {}

func (x *DataStoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_data_store_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataStoreResponse.ProtoReflect.Descriptor instead.
func (*DataStoreResponse) Descriptor() ([]byte, []int) {
	return file_data_store_proto_rawDescGZIP(), []int{1}
}

func (x *DataStoreResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CacheJobSearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobSearches []*JobSearchRequest `protobuf:"bytes,1,rep,name=job_searches,json=jobSearches,proto3" json:"job_searches,omitempty"`
}

func (x *CacheJobSearchResponse) Reset() {
	*x = CacheJobSearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_store_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CacheJobSearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CacheJobSearchResponse) ProtoMessage() {}

func (x *CacheJobSearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_data_store_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CacheJobSearchResponse.ProtoReflect.Descriptor instead.
func (*CacheJobSearchResponse) Descriptor() ([]byte, []int) {
	return file_data_store_proto_rawDescGZIP(), []int{2}
}

func (x *CacheJobSearchResponse) GetJobSearches() []*JobSearchRequest {
	if x != nil {
		return x.JobSearches
	}
	return nil
}

var File_data_store_proto protoreflect.FileDescriptor

var file_data_store_proto_rawDesc = []byte{
	0x0a, 0x10, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x1a, 0x10, 0x6a, 0x6f, 0x62, 0x2d, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x6a, 0x6f, 0x62, 0x2d, 0x61,
	0x70, 0x70, 0x6c, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x75, 0x73, 0x65, 0x72,
	0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x12, 0x0a, 0x10, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2d, 0x0a, 0x11, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x62, 0x0a, 0x16, 0x43, 0x61, 0x63, 0x68, 0x65, 0x4a, 0x6f, 0x62,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48,
	0x0a, 0x0c, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e,
	0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x6a, 0x6f, 0x62,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x73, 0x32, 0xfc, 0x13, 0x0a, 0x10, 0x44, 0x61, 0x74,
	0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x60, 0x0a,
	0x0f, 0x53, 0x61, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x25, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x26, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65,
	0x65, 0x6b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x61, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x28, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6a, 0x6f,
	0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x72, 0x0a, 0x18, 0x53, 0x61, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x2e,
	0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x1a, 0x26,
	0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x73, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x12, 0x28, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x6a, 0x6f,
	0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x99, 0x01, 0x0a, 0x24,
	0x53, 0x61, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x65, 0x64, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x12, 0x3a, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x65,
	0x64, 0x69, 0x63, 0x74, 0x65, 0x64, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x1a, 0x35, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x72, 0x65, 0x64,
	0x69, 0x63, 0x74, 0x65, 0x64, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x93, 0x01, 0x0a, 0x23, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x65, 0x64,
	0x69, 0x63, 0x74, 0x65, 0x64, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12,
	0x30, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x65, 0x64, 0x69, 0x63, 0x74, 0x65, 0x64, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x1a, 0x3a, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63,
	0x74, 0x65, 0x64, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x54, 0x0a,
	0x0c, 0x53, 0x61, 0x76, 0x65, 0x4a, 0x6f, 0x62, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x1c, 0x2e,
	0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x61, 0x70, 0x70,
	0x6c, 0x79, 0x2e, 0x4a, 0x6f, 0x62, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x1a, 0x26, 0x2e, 0x6a, 0x6f,
	0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x5e, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x41, 0x70, 0x70,
	0x6c, 0x79, 0x12, 0x26, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x6a, 0x6f,
	0x62, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x41, 0x70,
	0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6a, 0x6f, 0x62,
	0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x2e,
	0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x41, 0x70,
	0x70, 0x6c, 0x79, 0x12, 0x26, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x6a,
	0x6f, 0x62, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x41,
	0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6a, 0x6f,
	0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x79,
	0x2e, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f,
	0x62, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x1c, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65,
	0x6b, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x2e, 0x4a, 0x6f, 0x62, 0x41,
	0x70, 0x70, 0x6c, 0x79, 0x1a, 0x26, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x0f,
	0x53, 0x61, 0x76, 0x65, 0x4a, 0x6f, 0x62, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x12,
	0x1f, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x61,
	0x70, 0x70, 0x6c, 0x79, 0x2e, 0x4a, 0x6f, 0x62, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b,
	0x1a, 0x26, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4a,
	0x6f, 0x62, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x26, 0x2e, 0x6a, 0x6f, 0x62,
	0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x2e,
	0x4a, 0x6f, 0x62, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x6a, 0x6f,
	0x62, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x2e, 0x4a, 0x6f, 0x62, 0x42, 0x6f, 0x6f, 0x6b, 0x6d,
	0x61, 0x72, 0x6b, 0x12, 0x66, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x42, 0x6f,
	0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x26, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65,
	0x6b, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x2e, 0x4a, 0x6f, 0x62, 0x42,
	0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b,
	0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x61, 0x70,
	0x70, 0x6c, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x42, 0x6f, 0x6f, 0x6b, 0x6d,
	0x61, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x63, 0x0a, 0x11, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b,
	0x12, 0x26, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x6a, 0x6f, 0x62, 0x5f,
	0x61, 0x70, 0x70, 0x6c, 0x79, 0x2e, 0x4a, 0x6f, 0x62, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73,
	0x65, 0x65, 0x6b, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4b, 0x0a, 0x07, 0x53, 0x61, 0x76, 0x65, 0x4a, 0x6f, 0x62, 0x12, 0x18, 0x2e, 0x6a, 0x6f,
	0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x2e, 0x4a, 0x6f, 0x62, 0x1a, 0x26, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a,
	0x0c, 0x42, 0x61, 0x74, 0x63, 0x68, 0x53, 0x61, 0x76, 0x65, 0x4a, 0x6f, 0x62, 0x12, 0x18, 0x2e,
	0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x2e, 0x4a, 0x6f, 0x62, 0x1a, 0x26, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65,
	0x65, 0x6b, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x4a, 0x6f,
	0x62, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28,
	0x01, 0x30, 0x01, 0x12, 0x5a, 0x0a, 0x09, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4a, 0x6f, 0x62,
	0x12, 0x25, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x6a, 0x6f, 0x62, 0x5f,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65,
	0x65, 0x6b, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x4a, 0x6f,
	0x62, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4d, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x12, 0x18, 0x2e, 0x6a,
	0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x2e, 0x4a, 0x6f, 0x62, 0x1a, 0x26, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65,
	0x6b, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52,
	0x0a, 0x0e, 0x44, 0x72, 0x6f, 0x70, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x4a, 0x6f, 0x62,
	0x12, 0x18, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x6a, 0x6f, 0x62, 0x5f,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x4a, 0x6f, 0x62, 0x1a, 0x26, 0x2e, 0x6a, 0x6f, 0x62,
	0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x69, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x64,
	0x4a, 0x6f, 0x62, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x25, 0x2e, 0x6a, 0x6f, 0x62, 0x5f,
	0x73, 0x65, 0x65, 0x6b, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2b, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x4a, 0x6f, 0x62, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x69, 0x0a,
	0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x64, 0x4a, 0x6f, 0x62,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x25, 0x2e, 0x6a, 0x6f, 0x62,
	0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x25, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x6a, 0x6f, 0x62,
	0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x63, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43,
	0x61, 0x63, 0x68, 0x65, 0x64, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x25,
	0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b,
	0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x4a, 0x6f, 0x62, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a,
	0x11, 0x53, 0x61, 0x76, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x12, 0x22, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x6a, 0x6f,
	0x62, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x1a, 0x26, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65,
	0x6b, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6c,
	0x0a, 0x16, 0x42, 0x61, 0x74, 0x63, 0x68, 0x53, 0x61, 0x76, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x22, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73,
	0x65, 0x65, 0x6b, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x1a, 0x2a, 0x2e, 0x6a,
	0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x12, 0x69, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x12, 0x29, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x6a, 0x6f, 0x62, 0x5f,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x6a, 0x6f,
	0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x15, 0x5a, 0x13, 0x6a, 0x6f, 0x62, 0x2d, 0x73,
	0x65, 0x65, 0x6b, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_data_store_proto_rawDescOnce sync.Once
	file_data_store_proto_rawDescData = file_data_store_proto_rawDesc
)

func file_data_store_proto_rawDescGZIP() []byte {
	file_data_store_proto_rawDescOnce.Do(func() {
		file_data_store_proto_rawDescData = protoimpl.X.CompressGZIP(file_data_store_proto_rawDescData)
	})
	return file_data_store_proto_rawDescData
}

var file_data_store_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_data_store_proto_goTypes = []any{
	(*DataStoreRequest)(nil),                 // 0: job_seek.data_store.DataStoreRequest
	(*DataStoreResponse)(nil),                // 1: job_seek.data_store.DataStoreResponse
	(*CacheJobSearchResponse)(nil),           // 2: job_seek.data_store.CacheJobSearchResponse
	(*JobSearchRequest)(nil),                 // 3: job_seek.job_search.JobSearchRequest
	(*UserAccount)(nil),                      // 4: job_seek.user_management.UserAccount
	(*GetUserRequest)(nil),                   // 5: job_seek.user_management.GetUserRequest
	(*UserSearchPreference)(nil),             // 6: job_seek.user_management.UserSearchPreference
	(*UserJobSearchPredictedPreference)(nil), // 7: job_seek.user_management.UserJobSearchPredictedPreference
	(*GetPredictedPreference)(nil),           // 8: job_seek.user_management.GetPredictedPreference
	(*JobApply)(nil),                         // 9: job_seek.job_apply.JobApply
	(*GetJobApplyRequest)(nil),               // 10: job_seek.job_apply.GetJobApplyRequest
	(*JobBookmark)(nil),                      // 11: job_seek.job_apply.JobBookmark
	(*JobBookmarkRequest)(nil),               // 12: job_seek.job_apply.JobBookmarkRequest
	(*Job)(nil),                              // 13: job_seek.job_search.Job
	(*CompanyDetail)(nil),                    // 14: job_seek.job_search.CompanyDetail
	(*CompanyDetailRequest)(nil),             // 15: job_seek.job_search.CompanyDetailRequest
	(*UserResponse)(nil),                     // 16: job_seek.user_management.UserResponse
	(*PredictedPreferenceResponse)(nil),      // 17: job_seek.user_management.PredictedPreferenceResponse
	(*GetJobApplyResponse)(nil),              // 18: job_seek.job_apply.GetJobApplyResponse
	(*ListJobBookmarkResponse)(nil),          // 19: job_seek.job_apply.ListJobBookmarkResponse
	(*JobSearchResponse)(nil),                // 20: job_seek.job_search.JobSearchResponse
	(*CompanyDetailResponse)(nil),            // 21: job_seek.job_search.CompanyDetailResponse
}
var file_data_store_proto_depIdxs = []int32{
	3,  // 0: job_seek.data_store.CacheJobSearchResponse.job_searches:type_name -> job_seek.job_search.JobSearchRequest
	4,  // 1: job_seek.data_store.DataStoreService.SaveUserAccount:input_type -> job_seek.user_management.UserAccount
	5,  // 2: job_seek.data_store.DataStoreService.GetUserAccount:input_type -> job_seek.user_management.GetUserRequest
	6,  // 3: job_seek.data_store.DataStoreService.SaveUserSearchPreference:input_type -> job_seek.user_management.UserSearchPreference
	5,  // 4: job_seek.data_store.DataStoreService.GetUserSearchPreference:input_type -> job_seek.user_management.GetUserRequest
	7,  // 5: job_seek.data_store.DataStoreService.SaveUserJobSearchPredictedPreference:input_type -> job_seek.user_management.UserJobSearchPredictedPreference
	8,  // 6: job_seek.data_store.DataStoreService.GetUserJobSearchPredictedPreference:input_type -> job_seek.user_management.GetPredictedPreference
	9,  // 7: job_seek.data_store.DataStoreService.SaveJobApply:input_type -> job_seek.job_apply.JobApply
	10, // 8: job_seek.data_store.DataStoreService.GetJobApply:input_type -> job_seek.job_apply.GetJobApplyRequest
	10, // 9: job_seek.data_store.DataStoreService.ListJobApply:input_type -> job_seek.job_apply.GetJobApplyRequest
	9,  // 10: job_seek.data_store.DataStoreService.DeleteJobApply:input_type -> job_seek.job_apply.JobApply
	11, // 11: job_seek.data_store.DataStoreService.SaveJobBookmark:input_type -> job_seek.job_apply.JobBookmark
	12, // 12: job_seek.data_store.DataStoreService.GetJobBookmark:input_type -> job_seek.job_apply.JobBookmarkRequest
	12, // 13: job_seek.data_store.DataStoreService.ListJobBookmark:input_type -> job_seek.job_apply.JobBookmarkRequest
	12, // 14: job_seek.data_store.DataStoreService.DeleteJobBookmark:input_type -> job_seek.job_apply.JobBookmarkRequest
	13, // 15: job_seek.data_store.DataStoreService.SaveJob:input_type -> job_seek.job_search.Job
	13, // 16: job_seek.data_store.DataStoreService.BatchSaveJob:input_type -> job_seek.job_search.Job
	3,  // 17: job_seek.data_store.DataStoreService.SearchJob:input_type -> job_seek.job_search.JobSearchRequest
	13, // 18: job_seek.data_store.DataStoreService.DeleteJob:input_type -> job_seek.job_search.Job
	13, // 19: job_seek.data_store.DataStoreService.DropExpiredJob:input_type -> job_seek.job_search.Job
	0,  // 20: job_seek.data_store.DataStoreService.ListCachedJobSearch:input_type -> job_seek.data_store.DataStoreRequest
	3,  // 21: job_seek.data_store.DataStoreService.CreateCachedJobSearchList:input_type -> job_seek.job_search.JobSearchRequest
	3,  // 22: job_seek.data_store.DataStoreService.GetCachedJobSearch:input_type -> job_seek.job_search.JobSearchRequest
	14, // 23: job_seek.data_store.DataStoreService.SaveCompanyDetail:input_type -> job_seek.job_search.CompanyDetail
	14, // 24: job_seek.data_store.DataStoreService.BatchSaveCompanyDetail:input_type -> job_seek.job_search.CompanyDetail
	15, // 25: job_seek.data_store.DataStoreService.GetCompanyDetail:input_type -> job_seek.job_search.CompanyDetailRequest
	16, // 26: job_seek.data_store.DataStoreService.SaveUserAccount:output_type -> job_seek.user_management.UserResponse
	4,  // 27: job_seek.data_store.DataStoreService.GetUserAccount:output_type -> job_seek.user_management.UserAccount
	16, // 28: job_seek.data_store.DataStoreService.SaveUserSearchPreference:output_type -> job_seek.user_management.UserResponse
	6,  // 29: job_seek.data_store.DataStoreService.GetUserSearchPreference:output_type -> job_seek.user_management.UserSearchPreference
	17, // 30: job_seek.data_store.DataStoreService.SaveUserJobSearchPredictedPreference:output_type -> job_seek.user_management.PredictedPreferenceResponse
	7,  // 31: job_seek.data_store.DataStoreService.GetUserJobSearchPredictedPreference:output_type -> job_seek.user_management.UserJobSearchPredictedPreference
	1,  // 32: job_seek.data_store.DataStoreService.SaveJobApply:output_type -> job_seek.data_store.DataStoreResponse
	18, // 33: job_seek.data_store.DataStoreService.GetJobApply:output_type -> job_seek.job_apply.GetJobApplyResponse
	18, // 34: job_seek.data_store.DataStoreService.ListJobApply:output_type -> job_seek.job_apply.GetJobApplyResponse
	1,  // 35: job_seek.data_store.DataStoreService.DeleteJobApply:output_type -> job_seek.data_store.DataStoreResponse
	1,  // 36: job_seek.data_store.DataStoreService.SaveJobBookmark:output_type -> job_seek.data_store.DataStoreResponse
	11, // 37: job_seek.data_store.DataStoreService.GetJobBookmark:output_type -> job_seek.job_apply.JobBookmark
	19, // 38: job_seek.data_store.DataStoreService.ListJobBookmark:output_type -> job_seek.job_apply.ListJobBookmarkResponse
	1,  // 39: job_seek.data_store.DataStoreService.DeleteJobBookmark:output_type -> job_seek.data_store.DataStoreResponse
	1,  // 40: job_seek.data_store.DataStoreService.SaveJob:output_type -> job_seek.data_store.DataStoreResponse
	20, // 41: job_seek.data_store.DataStoreService.BatchSaveJob:output_type -> job_seek.job_search.JobSearchResponse
	20, // 42: job_seek.data_store.DataStoreService.SearchJob:output_type -> job_seek.job_search.JobSearchResponse
	1,  // 43: job_seek.data_store.DataStoreService.DeleteJob:output_type -> job_seek.data_store.DataStoreResponse
	1,  // 44: job_seek.data_store.DataStoreService.DropExpiredJob:output_type -> job_seek.data_store.DataStoreResponse
	2,  // 45: job_seek.data_store.DataStoreService.ListCachedJobSearch:output_type -> job_seek.data_store.CacheJobSearchResponse
	3,  // 46: job_seek.data_store.DataStoreService.CreateCachedJobSearchList:output_type -> job_seek.job_search.JobSearchRequest
	20, // 47: job_seek.data_store.DataStoreService.GetCachedJobSearch:output_type -> job_seek.job_search.JobSearchResponse
	1,  // 48: job_seek.data_store.DataStoreService.SaveCompanyDetail:output_type -> job_seek.data_store.DataStoreResponse
	21, // 49: job_seek.data_store.DataStoreService.BatchSaveCompanyDetail:output_type -> job_seek.job_search.CompanyDetailResponse
	21, // 50: job_seek.data_store.DataStoreService.GetCompanyDetail:output_type -> job_seek.job_search.CompanyDetailResponse
	26, // [26:51] is the sub-list for method output_type
	1,  // [1:26] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_data_store_proto_init() }
func file_data_store_proto_init() {
	if File_data_store_proto != nil {
		return
	}
	file_job_search_proto_init()
	file_job_apply_proto_init()
	file_user_management_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_data_store_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*DataStoreRequest); i {
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
		file_data_store_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*DataStoreResponse); i {
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
		file_data_store_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CacheJobSearchResponse); i {
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
			RawDescriptor: file_data_store_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_data_store_proto_goTypes,
		DependencyIndexes: file_data_store_proto_depIdxs,
		MessageInfos:      file_data_store_proto_msgTypes,
	}.Build()
	File_data_store_proto = out.File
	file_data_store_proto_rawDesc = nil
	file_data_store_proto_goTypes = nil
	file_data_store_proto_depIdxs = nil
}
