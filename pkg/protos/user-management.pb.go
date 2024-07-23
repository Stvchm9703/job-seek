// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: user-management.proto

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

type GetUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_management_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_management_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_user_management_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	UserId  string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UserResponse) Reset() {
	*x = UserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_management_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponse) ProtoMessage() {}

func (x *UserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_management_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponse.ProtoReflect.Descriptor instead.
func (*UserResponse) Descriptor() ([]byte, []int) {
	return file_user_management_proto_rawDescGZIP(), []int{1}
}

func (x *UserResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *UserResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type UserAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName     string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	UserPassword string `protobuf:"bytes,3,opt,name=user_password,json=userPassword,proto3" json:"user_password,omitempty"`
	UserEmail    string `protobuf:"bytes,4,opt,name=user_email,json=userEmail,proto3" json:"user_email,omitempty"`
	UserPhone    string `protobuf:"bytes,5,opt,name=user_phone,json=userPhone,proto3" json:"user_phone,omitempty"`
	UserAddress  string `protobuf:"bytes,6,opt,name=user_address,json=userAddress,proto3" json:"user_address,omitempty"`
}

func (x *UserAccount) Reset() {
	*x = UserAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_management_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAccount) ProtoMessage() {}

func (x *UserAccount) ProtoReflect() protoreflect.Message {
	mi := &file_user_management_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAccount.ProtoReflect.Descriptor instead.
func (*UserAccount) Descriptor() ([]byte, []int) {
	return file_user_management_proto_rawDescGZIP(), []int{2}
}

func (x *UserAccount) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserAccount) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UserAccount) GetUserPassword() string {
	if x != nil {
		return x.UserPassword
	}
	return ""
}

func (x *UserAccount) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

func (x *UserAccount) GetUserPhone() string {
	if x != nil {
		return x.UserPhone
	}
	return ""
}

func (x *UserAccount) GetUserAddress() string {
	if x != nil {
		return x.UserAddress
	}
	return ""
}

type UserSearchPerfence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecordId string             `protobuf:"bytes,1,opt,name=record_id,json=recordId,proto3" json:"record_id,omitempty"`
	UserId   string             `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // seeker-id
	JobId    string             `protobuf:"bytes,3,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`    // job-id
	Keywords []*PerfenceKeyword `protobuf:"bytes,4,rep,name=keywords,proto3" json:"keywords,omitempty"`
}

func (x *UserSearchPerfence) Reset() {
	*x = UserSearchPerfence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_management_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSearchPerfence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSearchPerfence) ProtoMessage() {}

func (x *UserSearchPerfence) ProtoReflect() protoreflect.Message {
	mi := &file_user_management_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSearchPerfence.ProtoReflect.Descriptor instead.
func (*UserSearchPerfence) Descriptor() ([]byte, []int) {
	return file_user_management_proto_rawDescGZIP(), []int{3}
}

func (x *UserSearchPerfence) GetRecordId() string {
	if x != nil {
		return x.RecordId
	}
	return ""
}

func (x *UserSearchPerfence) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserSearchPerfence) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *UserSearchPerfence) GetKeywords() []*PerfenceKeyword {
	if x != nil {
		return x.Keywords
	}
	return nil
}

type PerfenceKeyword struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KwId       string `protobuf:"bytes,1,opt,name=kw_id,json=kwId,proto3" json:"kw_id,omitempty"`
	UserId     string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`              // seeker-id
	Keyword    string `protobuf:"bytes,3,opt,name=keyword,proto3" json:"keyword,omitempty"`                          //
	Value      string `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`                              //
	Type       string `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`                                //
	IsPositive bool   `protobuf:"varint,6,opt,name=is_positive,json=isPositive,proto3" json:"is_positive,omitempty"` // true: positive, false: negative
}

func (x *PerfenceKeyword) Reset() {
	*x = PerfenceKeyword{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_management_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PerfenceKeyword) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PerfenceKeyword) ProtoMessage() {}

func (x *PerfenceKeyword) ProtoReflect() protoreflect.Message {
	mi := &file_user_management_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PerfenceKeyword.ProtoReflect.Descriptor instead.
func (*PerfenceKeyword) Descriptor() ([]byte, []int) {
	return file_user_management_proto_rawDescGZIP(), []int{4}
}

func (x *PerfenceKeyword) GetKwId() string {
	if x != nil {
		return x.KwId
	}
	return ""
}

func (x *PerfenceKeyword) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *PerfenceKeyword) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *PerfenceKeyword) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *PerfenceKeyword) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PerfenceKeyword) GetIsPositive() bool {
	if x != nil {
		return x.IsPositive
	}
	return false
}

type GetPredictedPerfence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecordId string `protobuf:"bytes,1,opt,name=record_id,json=recordId,proto3" json:"record_id,omitempty"`
}

func (x *GetPredictedPerfence) Reset() {
	*x = GetPredictedPerfence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_management_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPredictedPerfence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPredictedPerfence) ProtoMessage() {}

func (x *GetPredictedPerfence) ProtoReflect() protoreflect.Message {
	mi := &file_user_management_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPredictedPerfence.ProtoReflect.Descriptor instead.
func (*GetPredictedPerfence) Descriptor() ([]byte, []int) {
	return file_user_management_proto_rawDescGZIP(), []int{5}
}

func (x *GetPredictedPerfence) GetRecordId() string {
	if x != nil {
		return x.RecordId
	}
	return ""
}

type PredictedPerfenceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	UserId   string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RecordId string `protobuf:"bytes,3,opt,name=record_id,json=recordId,proto3" json:"record_id,omitempty"`
	Message  string `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PredictedPerfenceResponse) Reset() {
	*x = PredictedPerfenceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_management_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredictedPerfenceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictedPerfenceResponse) ProtoMessage() {}

func (x *PredictedPerfenceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_management_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredictedPerfenceResponse.ProtoReflect.Descriptor instead.
func (*PredictedPerfenceResponse) Descriptor() ([]byte, []int) {
	return file_user_management_proto_rawDescGZIP(), []int{6}
}

func (x *PredictedPerfenceResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *PredictedPerfenceResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *PredictedPerfenceResponse) GetRecordId() string {
	if x != nil {
		return x.RecordId
	}
	return ""
}

func (x *PredictedPerfenceResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type UserJobSearchPredictedPerfence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecordId   string             `protobuf:"bytes,1,opt,name=record_id,json=recordId,proto3" json:"record_id,omitempty"`
	UserId     string             `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // seeker-id
	JobId      string             `protobuf:"bytes,3,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`    // job-id
	SiteKey    string             `protobuf:"bytes,4,opt,name=site_key,json=siteKey,proto3" json:"site_key,omitempty"`
	Locale     string             `protobuf:"bytes,5,opt,name=locale,proto3" json:"locale,omitempty"`
	JobKeyword []*PerfenceKeyword `protobuf:"bytes,6,rep,name=job_keyword,json=jobKeyword,proto3" json:"job_keyword,omitempty"` //
	Score      int32              `protobuf:"varint,7,opt,name=score,proto3" json:"score,omitempty"`                            // 0-100
	Count      int32              `protobuf:"varint,8,opt,name=count,proto3" json:"count,omitempty"`                            // 0-100
}

func (x *UserJobSearchPredictedPerfence) Reset() {
	*x = UserJobSearchPredictedPerfence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_management_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserJobSearchPredictedPerfence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserJobSearchPredictedPerfence) ProtoMessage() {}

func (x *UserJobSearchPredictedPerfence) ProtoReflect() protoreflect.Message {
	mi := &file_user_management_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserJobSearchPredictedPerfence.ProtoReflect.Descriptor instead.
func (*UserJobSearchPredictedPerfence) Descriptor() ([]byte, []int) {
	return file_user_management_proto_rawDescGZIP(), []int{7}
}

func (x *UserJobSearchPredictedPerfence) GetRecordId() string {
	if x != nil {
		return x.RecordId
	}
	return ""
}

func (x *UserJobSearchPredictedPerfence) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserJobSearchPredictedPerfence) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *UserJobSearchPredictedPerfence) GetSiteKey() string {
	if x != nil {
		return x.SiteKey
	}
	return ""
}

func (x *UserJobSearchPredictedPerfence) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *UserJobSearchPredictedPerfence) GetJobKeyword() []*PerfenceKeyword {
	if x != nil {
		return x.JobKeyword
	}
	return nil
}

func (x *UserJobSearchPredictedPerfence) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *UserJobSearchPredictedPerfence) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type PredictedPerfenceMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	MapId  string `protobuf:"bytes,2,opt,name=map_id,json=mapId,proto3" json:"map_id,omitempty"`
}

func (x *PredictedPerfenceMap) Reset() {
	*x = PredictedPerfenceMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_management_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredictedPerfenceMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictedPerfenceMap) ProtoMessage() {}

func (x *PredictedPerfenceMap) ProtoReflect() protoreflect.Message {
	mi := &file_user_management_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredictedPerfenceMap.ProtoReflect.Descriptor instead.
func (*PredictedPerfenceMap) Descriptor() ([]byte, []int) {
	return file_user_management_proto_rawDescGZIP(), []int{8}
}

func (x *PredictedPerfenceMap) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *PredictedPerfenceMap) GetMapId() string {
	if x != nil {
		return x.MapId
	}
	return ""
}

var File_user_management_proto protoreflect.FileDescriptor

var file_user_management_proto_rawDesc = []byte{
	0x0a, 0x15, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65,
	0x6b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x22, 0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x59, 0x0a, 0x0c,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xc9, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a,
	0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x22, 0xa8, 0x01, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x50, 0x65, 0x72, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x45, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6a, 0x6f, 0x62, 0x5f,
	0x73, 0x65, 0x65, 0x6b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x65, 0x72, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x4b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x22, 0xa4,
	0x01, 0x0a, 0x0f, 0x50, 0x65, 0x72, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x4b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x13, 0x0a, 0x05, 0x6b, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6b, 0x77, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x76, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x76, 0x65, 0x22, 0x33, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x50, 0x72, 0x65, 0x64,
	0x69, 0x63, 0x74, 0x65, 0x64, 0x50, 0x65, 0x72, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x64, 0x22, 0x83, 0x01, 0x0a, 0x19, 0x50,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x65, 0x64, 0x50, 0x65, 0x72, 0x66, 0x65, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x98, 0x02, 0x0a, 0x1e, 0x55, 0x73, 0x65, 0x72, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x65, 0x64, 0x50, 0x65, 0x72, 0x66, 0x65,
	0x6e, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x69, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x65, 0x12, 0x4a, 0x0a, 0x0b, 0x6a, 0x6f, 0x62, 0x5f, 0x6b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73,
	0x65, 0x65, 0x6b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x50, 0x65, 0x72, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x4b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x52, 0x0a, 0x6a, 0x6f, 0x62, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x46, 0x0a, 0x14, 0x50,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x65, 0x64, 0x50, 0x65, 0x72, 0x66, 0x65, 0x6e, 0x63, 0x65,
	0x4d, 0x61, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06,
	0x6d, 0x61, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x61,
	0x70, 0x49, 0x64, 0x32, 0xf1, 0x07, 0x0a, 0x15, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x60, 0x0a,
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
	0x6e, 0x74, 0x12, 0x6e, 0x0a, 0x16, 0x53, 0x61, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x50, 0x65, 0x72, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x2c, 0x2e, 0x6a,
	0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x50, 0x65, 0x72, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x1a, 0x26, 0x2e, 0x6a, 0x6f, 0x62,
	0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x70, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x65, 0x72, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x2c,
	0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x50, 0x65, 0x72, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x1a, 0x26, 0x2e, 0x6a,
	0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6f, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x65, 0x72, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x28, 0x2e,
	0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65,
	0x65, 0x6b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x65, 0x72,
	0x66, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x93, 0x01, 0x0a, 0x22, 0x53, 0x61, 0x76, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x65, 0x64, 0x69,
	0x63, 0x74, 0x65, 0x64, 0x50, 0x65, 0x72, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x38, 0x2e, 0x6a,
	0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4a, 0x6f, 0x62, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x65, 0x64, 0x50, 0x65,
	0x72, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x1a, 0x33, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65,
	0x6b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x65, 0x64, 0x50, 0x65, 0x72, 0x66, 0x65,
	0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x8d, 0x01, 0x0a, 0x21,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x65, 0x64, 0x50, 0x65, 0x72, 0x66, 0x65, 0x6e, 0x63,
	0x65, 0x12, 0x2e, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x65, 0x64, 0x50, 0x65, 0x72, 0x66, 0x65, 0x6e, 0x63,
	0x65, 0x1a, 0x38, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63,
	0x74, 0x65, 0x64, 0x50, 0x65, 0x72, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x99, 0x01, 0x0a, 0x24,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x65, 0x64, 0x50, 0x65, 0x72, 0x66,
	0x65, 0x6e, 0x63, 0x65, 0x12, 0x38, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x65,
	0x64, 0x69, 0x63, 0x74, 0x65, 0x64, 0x50, 0x65, 0x72, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x1a, 0x33,
	0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63,
	0x74, 0x65, 0x64, 0x50, 0x65, 0x72, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x15, 0x5a, 0x13, 0x6a, 0x6f, 0x62, 0x2d, 0x73,
	0x65, 0x65, 0x6b, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_management_proto_rawDescOnce sync.Once
	file_user_management_proto_rawDescData = file_user_management_proto_rawDesc
)

func file_user_management_proto_rawDescGZIP() []byte {
	file_user_management_proto_rawDescOnce.Do(func() {
		file_user_management_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_management_proto_rawDescData)
	})
	return file_user_management_proto_rawDescData
}

var file_user_management_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_user_management_proto_goTypes = []any{
	(*GetUserRequest)(nil),                 // 0: job_seek.user_management.GetUserRequest
	(*UserResponse)(nil),                   // 1: job_seek.user_management.UserResponse
	(*UserAccount)(nil),                    // 2: job_seek.user_management.UserAccount
	(*UserSearchPerfence)(nil),             // 3: job_seek.user_management.UserSearchPerfence
	(*PerfenceKeyword)(nil),                // 4: job_seek.user_management.PerfenceKeyword
	(*GetPredictedPerfence)(nil),           // 5: job_seek.user_management.GetPredictedPerfence
	(*PredictedPerfenceResponse)(nil),      // 6: job_seek.user_management.PredictedPerfenceResponse
	(*UserJobSearchPredictedPerfence)(nil), // 7: job_seek.user_management.UserJobSearchPredictedPerfence
	(*PredictedPerfenceMap)(nil),           // 8: job_seek.user_management.PredictedPerfenceMap
}
var file_user_management_proto_depIdxs = []int32{
	4,  // 0: job_seek.user_management.UserSearchPerfence.keywords:type_name -> job_seek.user_management.PerfenceKeyword
	4,  // 1: job_seek.user_management.UserJobSearchPredictedPerfence.job_keyword:type_name -> job_seek.user_management.PerfenceKeyword
	2,  // 2: job_seek.user_management.UserManagementService.SaveUserAccount:input_type -> job_seek.user_management.UserAccount
	0,  // 3: job_seek.user_management.UserManagementService.GetUserAccount:input_type -> job_seek.user_management.GetUserRequest
	3,  // 4: job_seek.user_management.UserManagementService.SaveUserSearchPerfence:input_type -> job_seek.user_management.UserSearchPerfence
	3,  // 5: job_seek.user_management.UserManagementService.UpdateUserSearchPerfence:input_type -> job_seek.user_management.UserSearchPerfence
	0,  // 6: job_seek.user_management.UserManagementService.GetUserSearchPerfence:input_type -> job_seek.user_management.GetUserRequest
	7,  // 7: job_seek.user_management.UserManagementService.SaveUserJobSearchPredictedPerfence:input_type -> job_seek.user_management.UserJobSearchPredictedPerfence
	5,  // 8: job_seek.user_management.UserManagementService.GetUserJobSearchPredictedPerfence:input_type -> job_seek.user_management.GetPredictedPerfence
	7,  // 9: job_seek.user_management.UserManagementService.CreateUserJobSearchPredictedPerfence:input_type -> job_seek.user_management.UserJobSearchPredictedPerfence
	1,  // 10: job_seek.user_management.UserManagementService.SaveUserAccount:output_type -> job_seek.user_management.UserResponse
	2,  // 11: job_seek.user_management.UserManagementService.GetUserAccount:output_type -> job_seek.user_management.UserAccount
	1,  // 12: job_seek.user_management.UserManagementService.SaveUserSearchPerfence:output_type -> job_seek.user_management.UserResponse
	1,  // 13: job_seek.user_management.UserManagementService.UpdateUserSearchPerfence:output_type -> job_seek.user_management.UserResponse
	3,  // 14: job_seek.user_management.UserManagementService.GetUserSearchPerfence:output_type -> job_seek.user_management.UserSearchPerfence
	6,  // 15: job_seek.user_management.UserManagementService.SaveUserJobSearchPredictedPerfence:output_type -> job_seek.user_management.PredictedPerfenceResponse
	7,  // 16: job_seek.user_management.UserManagementService.GetUserJobSearchPredictedPerfence:output_type -> job_seek.user_management.UserJobSearchPredictedPerfence
	6,  // 17: job_seek.user_management.UserManagementService.CreateUserJobSearchPredictedPerfence:output_type -> job_seek.user_management.PredictedPerfenceResponse
	10, // [10:18] is the sub-list for method output_type
	2,  // [2:10] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_user_management_proto_init() }
func file_user_management_proto_init() {
	if File_user_management_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_management_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetUserRequest); i {
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
		file_user_management_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*UserResponse); i {
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
		file_user_management_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*UserAccount); i {
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
		file_user_management_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UserSearchPerfence); i {
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
		file_user_management_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*PerfenceKeyword); i {
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
		file_user_management_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetPredictedPerfence); i {
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
		file_user_management_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*PredictedPerfenceResponse); i {
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
		file_user_management_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*UserJobSearchPredictedPerfence); i {
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
		file_user_management_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*PredictedPerfenceMap); i {
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
			RawDescriptor: file_user_management_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_management_proto_goTypes,
		DependencyIndexes: file_user_management_proto_depIdxs,
		MessageInfos:      file_user_management_proto_msgTypes,
	}.Build()
	File_user_management_proto = out.File
	file_user_management_proto_rawDesc = nil
	file_user_management_proto_goTypes = nil
	file_user_management_proto_depIdxs = nil
}
