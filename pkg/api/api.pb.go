// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.4
// source: pkg/api/api.proto

package api

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

type Codes int32

const (
	Codes_CODES_SUCCESS                 Codes = 0
	Codes_CODES_SERVER_INTERNAL_ERROR   Codes = 1
	Codes_CODES_INVALID_PARAMETER_ERROR Codes = 2
)

// Enum value maps for Codes.
var (
	Codes_name = map[int32]string{
		0: "CODES_SUCCESS",
		1: "CODES_SERVER_INTERNAL_ERROR",
		2: "CODES_INVALID_PARAMETER_ERROR",
	}
	Codes_value = map[string]int32{
		"CODES_SUCCESS":                 0,
		"CODES_SERVER_INTERNAL_ERROR":   1,
		"CODES_INVALID_PARAMETER_ERROR": 2,
	}
)

func (x Codes) Enum() *Codes {
	p := new(Codes)
	*p = x
	return p
}

func (x Codes) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Codes) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_api_api_proto_enumTypes[0].Descriptor()
}

func (Codes) Type() protoreflect.EnumType {
	return &file_pkg_api_api_proto_enumTypes[0]
}

func (x Codes) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Codes.Descriptor instead.
func (Codes) EnumDescriptor() ([]byte, []int) {
	return file_pkg_api_api_proto_rawDescGZIP(), []int{0}
}

type ImagineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Prompt    string `protobuf:"bytes,2,opt,name=prompt,proto3" json:"prompt,omitempty"`
	Webhook   string `protobuf:"bytes,3,opt,name=webhook,proto3" json:"webhook,omitempty"`
}

func (x *ImagineRequest) Reset() {
	*x = ImagineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImagineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImagineRequest) ProtoMessage() {}

func (x *ImagineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImagineRequest.ProtoReflect.Descriptor instead.
func (*ImagineRequest) Descriptor() ([]byte, []int) {
	return file_pkg_api_api_proto_rawDescGZIP(), []int{0}
}

func (x *ImagineRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *ImagineRequest) GetPrompt() string {
	if x != nil {
		return x.Prompt
	}
	return ""
}

func (x *ImagineRequest) GetWebhook() string {
	if x != nil {
		return x.Webhook
	}
	return ""
}

type ImagineResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string               `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Code      Codes                `protobuf:"varint,2,opt,name=code,proto3,enum=api.Codes" json:"code,omitempty"`
	Msg       string               `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	Data      *ImagineResponseData `protobuf:"bytes,4,opt,name=data,proto3,oneof" json:"data,omitempty"`
}

func (x *ImagineResponse) Reset() {
	*x = ImagineResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImagineResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImagineResponse) ProtoMessage() {}

func (x *ImagineResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImagineResponse.ProtoReflect.Descriptor instead.
func (*ImagineResponse) Descriptor() ([]byte, []int) {
	return file_pkg_api_api_proto_rawDescGZIP(), []int{1}
}

func (x *ImagineResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *ImagineResponse) GetCode() Codes {
	if x != nil {
		return x.Code
	}
	return Codes_CODES_SUCCESS
}

func (x *ImagineResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ImagineResponse) GetData() *ImagineResponseData {
	if x != nil {
		return x.Data
	}
	return nil
}

type ImagineResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId    string `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	StartTime int64  `protobuf:"varint,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
}

func (x *ImagineResponseData) Reset() {
	*x = ImagineResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImagineResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImagineResponseData) ProtoMessage() {}

func (x *ImagineResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImagineResponseData.ProtoReflect.Descriptor instead.
func (*ImagineResponseData) Descriptor() ([]byte, []int) {
	return file_pkg_api_api_proto_rawDescGZIP(), []int{2}
}

func (x *ImagineResponseData) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *ImagineResponseData) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

type UpscaleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Index     int32  `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	TaskId    string `protobuf:"bytes,3,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Webhook   string `protobuf:"bytes,4,opt,name=webhook,proto3" json:"webhook,omitempty"`
}

func (x *UpscaleRequest) Reset() {
	*x = UpscaleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpscaleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpscaleRequest) ProtoMessage() {}

func (x *UpscaleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpscaleRequest.ProtoReflect.Descriptor instead.
func (*UpscaleRequest) Descriptor() ([]byte, []int) {
	return file_pkg_api_api_proto_rawDescGZIP(), []int{3}
}

func (x *UpscaleRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *UpscaleRequest) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *UpscaleRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *UpscaleRequest) GetWebhook() string {
	if x != nil {
		return x.Webhook
	}
	return ""
}

type UpscaleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string               `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Code      Codes                `protobuf:"varint,2,opt,name=code,proto3,enum=api.Codes" json:"code,omitempty"`
	Msg       string               `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	Data      *UpscaleResponseData `protobuf:"bytes,4,opt,name=data,proto3,oneof" json:"data,omitempty"`
}

func (x *UpscaleResponse) Reset() {
	*x = UpscaleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpscaleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpscaleResponse) ProtoMessage() {}

func (x *UpscaleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpscaleResponse.ProtoReflect.Descriptor instead.
func (*UpscaleResponse) Descriptor() ([]byte, []int) {
	return file_pkg_api_api_proto_rawDescGZIP(), []int{4}
}

func (x *UpscaleResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *UpscaleResponse) GetCode() Codes {
	if x != nil {
		return x.Code
	}
	return Codes_CODES_SUCCESS
}

func (x *UpscaleResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *UpscaleResponse) GetData() *UpscaleResponseData {
	if x != nil {
		return x.Data
	}
	return nil
}

type UpscaleResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId    string `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	StartTime int64  `protobuf:"varint,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
}

func (x *UpscaleResponseData) Reset() {
	*x = UpscaleResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpscaleResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpscaleResponseData) ProtoMessage() {}

func (x *UpscaleResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpscaleResponseData.ProtoReflect.Descriptor instead.
func (*UpscaleResponseData) Descriptor() ([]byte, []int) {
	return file_pkg_api_api_proto_rawDescGZIP(), []int{5}
}

func (x *UpscaleResponseData) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *UpscaleResponseData) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

var File_pkg_api_api_proto protoreflect.FileDescriptor

var file_pkg_api_api_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x61, 0x0a, 0x0e, 0x49, 0x6d, 0x61, 0x67,
	0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x6f,
	0x6d, 0x70, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x70,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x22, 0x9e, 0x01, 0x0a, 0x0f,
	0x49, 0x6d, 0x61, 0x67, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1e,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x12, 0x31, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x22, 0x4d, 0x0a, 0x13,
	0x49, 0x6d, 0x61, 0x67, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x78, 0x0a, 0x0e, 0x55,
	0x70, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x77,
	0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x65,
	0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x22, 0x9e, 0x01, 0x0a, 0x0f, 0x55, 0x70, 0x73, 0x63, 0x61, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x64,
	0x65, 0x73, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x31, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55,
	0x70, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a,
	0x05, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x22, 0x4d, 0x0a, 0x13, 0x55, 0x70, 0x73, 0x63, 0x61, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x17, 0x0a,
	0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x2a, 0x5e, 0x0a, 0x05, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x11,
	0x0a, 0x0d, 0x43, 0x4f, 0x44, 0x45, 0x53, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10,
	0x00, 0x12, 0x1f, 0x0a, 0x1b, 0x43, 0x4f, 0x44, 0x45, 0x53, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45,
	0x52, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0x01, 0x12, 0x21, 0x0a, 0x1d, 0x43, 0x4f, 0x44, 0x45, 0x53, 0x5f, 0x49, 0x4e, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x45, 0x54, 0x45, 0x52, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0x02, 0x32, 0x78, 0x0a, 0x0a, 0x41, 0x50, 0x49, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x49, 0x6d, 0x61, 0x67, 0x69, 0x6e, 0x65, 0x12, 0x13,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x69, 0x6e,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x55, 0x70, 0x73,
	0x63, 0x61, 0x6c, 0x65, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x73, 0x63, 0x61,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x55, 0x70, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_api_api_proto_rawDescOnce sync.Once
	file_pkg_api_api_proto_rawDescData = file_pkg_api_api_proto_rawDesc
)

func file_pkg_api_api_proto_rawDescGZIP() []byte {
	file_pkg_api_api_proto_rawDescOnce.Do(func() {
		file_pkg_api_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_api_api_proto_rawDescData)
	})
	return file_pkg_api_api_proto_rawDescData
}

var file_pkg_api_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_api_api_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pkg_api_api_proto_goTypes = []interface{}{
	(Codes)(0),                  // 0: api.Codes
	(*ImagineRequest)(nil),      // 1: api.ImagineRequest
	(*ImagineResponse)(nil),     // 2: api.ImagineResponse
	(*ImagineResponseData)(nil), // 3: api.ImagineResponseData
	(*UpscaleRequest)(nil),      // 4: api.UpscaleRequest
	(*UpscaleResponse)(nil),     // 5: api.UpscaleResponse
	(*UpscaleResponseData)(nil), // 6: api.UpscaleResponseData
}
var file_pkg_api_api_proto_depIdxs = []int32{
	0, // 0: api.ImagineResponse.code:type_name -> api.Codes
	3, // 1: api.ImagineResponse.data:type_name -> api.ImagineResponseData
	0, // 2: api.UpscaleResponse.code:type_name -> api.Codes
	6, // 3: api.UpscaleResponse.data:type_name -> api.UpscaleResponseData
	1, // 4: api.APIService.Imagine:input_type -> api.ImagineRequest
	4, // 5: api.APIService.Upscale:input_type -> api.UpscaleRequest
	2, // 6: api.APIService.Imagine:output_type -> api.ImagineResponse
	5, // 7: api.APIService.Upscale:output_type -> api.UpscaleResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pkg_api_api_proto_init() }
func file_pkg_api_api_proto_init() {
	if File_pkg_api_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_api_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImagineRequest); i {
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
		file_pkg_api_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImagineResponse); i {
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
		file_pkg_api_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImagineResponseData); i {
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
		file_pkg_api_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpscaleRequest); i {
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
		file_pkg_api_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpscaleResponse); i {
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
		file_pkg_api_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpscaleResponseData); i {
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
	file_pkg_api_api_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_pkg_api_api_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_api_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_api_api_proto_goTypes,
		DependencyIndexes: file_pkg_api_api_proto_depIdxs,
		EnumInfos:         file_pkg_api_api_proto_enumTypes,
		MessageInfos:      file_pkg_api_api_proto_msgTypes,
	}.Build()
	File_pkg_api_api_proto = out.File
	file_pkg_api_api_proto_rawDesc = nil
	file_pkg_api_api_proto_goTypes = nil
	file_pkg_api_api_proto_depIdxs = nil
}
