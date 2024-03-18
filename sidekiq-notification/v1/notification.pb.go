// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: protobuf/v1/notification.proto

package v1

import (
	any1 "github.com/golang/protobuf/ptypes/any"
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

type MarkNotificationAsReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NotificationID string `protobuf:"bytes,1,opt,name=notificationID,proto3" json:"notificationID,omitempty"`
	ProfileID      string `protobuf:"bytes,2,opt,name=profileID,proto3" json:"profileID,omitempty"`
}

func (x *MarkNotificationAsReadRequest) Reset() {
	*x = MarkNotificationAsReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_v1_notification_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkNotificationAsReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkNotificationAsReadRequest) ProtoMessage() {}

func (x *MarkNotificationAsReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_v1_notification_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkNotificationAsReadRequest.ProtoReflect.Descriptor instead.
func (*MarkNotificationAsReadRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_v1_notification_proto_rawDescGZIP(), []int{0}
}

func (x *MarkNotificationAsReadRequest) GetNotificationID() string {
	if x != nil {
		return x.NotificationID
	}
	return ""
}

func (x *MarkNotificationAsReadRequest) GetProfileID() string {
	if x != nil {
		return x.ProfileID
	}
	return ""
}

type MarkAllNotificationAsReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProfileID string `protobuf:"bytes,1,opt,name=profileID,proto3" json:"profileID,omitempty"`
}

func (x *MarkAllNotificationAsReadRequest) Reset() {
	*x = MarkAllNotificationAsReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_v1_notification_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkAllNotificationAsReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkAllNotificationAsReadRequest) ProtoMessage() {}

func (x *MarkAllNotificationAsReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_v1_notification_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkAllNotificationAsReadRequest.ProtoReflect.Descriptor instead.
func (*MarkAllNotificationAsReadRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_v1_notification_proto_rawDescGZIP(), []int{1}
}

func (x *MarkAllNotificationAsReadRequest) GetProfileID() string {
	if x != nil {
		return x.ProfileID
	}
	return ""
}

type GetNotificationListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProfileID string `protobuf:"bytes,1,opt,name=profileID,proto3" json:"profileID,omitempty"`
}

func (x *GetNotificationListRequest) Reset() {
	*x = GetNotificationListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_v1_notification_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNotificationListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNotificationListRequest) ProtoMessage() {}

func (x *GetNotificationListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_v1_notification_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNotificationListRequest.ProtoReflect.Descriptor instead.
func (*GetNotificationListRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_v1_notification_proto_rawDescGZIP(), []int{2}
}

func (x *GetNotificationListRequest) GetProfileID() string {
	if x != nil {
		return x.ProfileID
	}
	return ""
}

type GetNotificationDisplayCountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProfileID string `protobuf:"bytes,1,opt,name=profileID,proto3" json:"profileID,omitempty"`
}

func (x *GetNotificationDisplayCountRequest) Reset() {
	*x = GetNotificationDisplayCountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_v1_notification_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNotificationDisplayCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNotificationDisplayCountRequest) ProtoMessage() {}

func (x *GetNotificationDisplayCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_v1_notification_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNotificationDisplayCountRequest.ProtoReflect.Descriptor instead.
func (*GetNotificationDisplayCountRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_v1_notification_proto_rawDescGZIP(), []int{3}
}

func (x *GetNotificationDisplayCountRequest) GetProfileID() string {
	if x != nil {
		return x.ProfileID
	}
	return ""
}

type GetNotificationDisplayCountReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetNotificationDisplayCountReply) Reset() {
	*x = GetNotificationDisplayCountReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_v1_notification_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNotificationDisplayCountReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNotificationDisplayCountReply) ProtoMessage() {}

func (x *GetNotificationDisplayCountReply) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_v1_notification_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNotificationDisplayCountReply.ProtoReflect.Descriptor instead.
func (*GetNotificationDisplayCountReply) Descriptor() ([]byte, []int) {
	return file_protobuf_v1_notification_proto_rawDescGZIP(), []int{4}
}

func (x *GetNotificationDisplayCountReply) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type GenericReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data    *any1.Any `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Status  int32     `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Message string    `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GenericReply) Reset() {
	*x = GenericReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_v1_notification_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenericReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenericReply) ProtoMessage() {}

func (x *GenericReply) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_v1_notification_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenericReply.ProtoReflect.Descriptor instead.
func (*GenericReply) Descriptor() ([]byte, []int) {
	return file_protobuf_v1_notification_proto_rawDescGZIP(), []int{5}
}

func (x *GenericReply) GetData() *any1.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GenericReply) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GenericReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type NotificationHandlerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReceiverIDs []int32 `protobuf:"varint,1,rep,packed,name=receiverIDs,proto3" json:"receiverIDs,omitempty"`
	SenderID    int32   `protobuf:"varint,2,opt,name=senderID,proto3" json:"senderID,omitempty"`
	ThingType   string  `protobuf:"bytes,3,opt,name=thingType,proto3" json:"thingType,omitempty"`
	ThingID     string  `protobuf:"bytes,4,opt,name=thingID,proto3" json:"thingID,omitempty"`
	ActionType  string  `protobuf:"bytes,5,opt,name=actionType,proto3" json:"actionType,omitempty"`
	Message     string  `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *NotificationHandlerRequest) Reset() {
	*x = NotificationHandlerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_v1_notification_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationHandlerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationHandlerRequest) ProtoMessage() {}

func (x *NotificationHandlerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_v1_notification_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationHandlerRequest.ProtoReflect.Descriptor instead.
func (*NotificationHandlerRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_v1_notification_proto_rawDescGZIP(), []int{6}
}

func (x *NotificationHandlerRequest) GetReceiverIDs() []int32 {
	if x != nil {
		return x.ReceiverIDs
	}
	return nil
}

func (x *NotificationHandlerRequest) GetSenderID() int32 {
	if x != nil {
		return x.SenderID
	}
	return 0
}

func (x *NotificationHandlerRequest) GetThingType() string {
	if x != nil {
		return x.ThingType
	}
	return ""
}

func (x *NotificationHandlerRequest) GetThingID() string {
	if x != nil {
		return x.ThingID
	}
	return ""
}

func (x *NotificationHandlerRequest) GetActionType() string {
	if x != nil {
		return x.ActionType
	}
	return ""
}

func (x *NotificationHandlerRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_protobuf_v1_notification_proto protoreflect.FileDescriptor

var file_protobuf_v1_notification_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x65, 0x0a, 0x1d,
	0x4d, 0x61, 0x72, 0x6b, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x41, 0x73, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a,
	0x0e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x49, 0x44, 0x22, 0x40, 0x0a, 0x20, 0x4d, 0x61, 0x72, 0x6b, 0x41, 0x6c, 0x6c, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x73, 0x52, 0x65, 0x61, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x49, 0x44, 0x22, 0x3a, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49,
	0x44, 0x22, 0x42, 0x0a, 0x22, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x49, 0x44, 0x22, 0x38, 0x0a, 0x20, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x6a, 0x0a, 0x0c, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xcc, 0x01, 0x0a, 0x1a,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x49, 0x44, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x49, 0x44, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68, 0x69, 0x6e,
	0x67, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x68, 0x69,
	0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x49,
	0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x49, 0x44,
	0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xc5, 0x04, 0x0a, 0x13, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x69, 0x0a, 0x16, 0x4d, 0x61, 0x72, 0x6b, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x73, 0x52, 0x65, 0x61, 0x64, 0x12, 0x2e, 0x2e, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x61, 0x72, 0x6b, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41,
	0x73, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x6f, 0x0a,
	0x19, 0x4d, 0x61, 0x72, 0x6b, 0x41, 0x6c, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x41, 0x73, 0x52, 0x65, 0x61, 0x64, 0x12, 0x31, 0x2e, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x72,
	0x6b, 0x41, 0x6c, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x41, 0x73, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x63,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2b, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x87, 0x01, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x33, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x63, 0x0a,
	0x13, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x72, 0x12, 0x2b, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x42, 0x1a, 0x5a, 0x18, 0x2f, 0x73, 0x69, 0x64, 0x65, 0x6b, 0x69, 0x71, 0x2d, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_v1_notification_proto_rawDescOnce sync.Once
	file_protobuf_v1_notification_proto_rawDescData = file_protobuf_v1_notification_proto_rawDesc
)

func file_protobuf_v1_notification_proto_rawDescGZIP() []byte {
	file_protobuf_v1_notification_proto_rawDescOnce.Do(func() {
		file_protobuf_v1_notification_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_v1_notification_proto_rawDescData)
	})
	return file_protobuf_v1_notification_proto_rawDescData
}

var file_protobuf_v1_notification_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_protobuf_v1_notification_proto_goTypes = []interface{}{
	(*MarkNotificationAsReadRequest)(nil),      // 0: notification.v1.MarkNotificationAsReadRequest
	(*MarkAllNotificationAsReadRequest)(nil),   // 1: notification.v1.MarkAllNotificationAsReadRequest
	(*GetNotificationListRequest)(nil),         // 2: notification.v1.GetNotificationListRequest
	(*GetNotificationDisplayCountRequest)(nil), // 3: notification.v1.GetNotificationDisplayCountRequest
	(*GetNotificationDisplayCountReply)(nil),   // 4: notification.v1.GetNotificationDisplayCountReply
	(*GenericReply)(nil),                       // 5: notification.v1.GenericReply
	(*NotificationHandlerRequest)(nil),         // 6: notification.v1.NotificationHandlerRequest
	(*any1.Any)(nil),                           // 7: google.protobuf.Any
}
var file_protobuf_v1_notification_proto_depIdxs = []int32{
	7, // 0: notification.v1.GenericReply.data:type_name -> google.protobuf.Any
	0, // 1: notification.v1.NotificationService.MarkNotificationAsRead:input_type -> notification.v1.MarkNotificationAsReadRequest
	1, // 2: notification.v1.NotificationService.MarkAllNotificationAsRead:input_type -> notification.v1.MarkAllNotificationAsReadRequest
	2, // 3: notification.v1.NotificationService.GetNotificationList:input_type -> notification.v1.GetNotificationListRequest
	3, // 4: notification.v1.NotificationService.GetNotificationDisplayCount:input_type -> notification.v1.GetNotificationDisplayCountRequest
	6, // 5: notification.v1.NotificationService.NotificationHandler:input_type -> notification.v1.NotificationHandlerRequest
	5, // 6: notification.v1.NotificationService.MarkNotificationAsRead:output_type -> notification.v1.GenericReply
	5, // 7: notification.v1.NotificationService.MarkAllNotificationAsRead:output_type -> notification.v1.GenericReply
	5, // 8: notification.v1.NotificationService.GetNotificationList:output_type -> notification.v1.GenericReply
	4, // 9: notification.v1.NotificationService.GetNotificationDisplayCount:output_type -> notification.v1.GetNotificationDisplayCountReply
	5, // 10: notification.v1.NotificationService.NotificationHandler:output_type -> notification.v1.GenericReply
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protobuf_v1_notification_proto_init() }
func file_protobuf_v1_notification_proto_init() {
	if File_protobuf_v1_notification_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_v1_notification_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarkNotificationAsReadRequest); i {
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
		file_protobuf_v1_notification_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarkAllNotificationAsReadRequest); i {
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
		file_protobuf_v1_notification_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNotificationListRequest); i {
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
		file_protobuf_v1_notification_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNotificationDisplayCountRequest); i {
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
		file_protobuf_v1_notification_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNotificationDisplayCountReply); i {
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
		file_protobuf_v1_notification_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenericReply); i {
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
		file_protobuf_v1_notification_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationHandlerRequest); i {
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
			RawDescriptor: file_protobuf_v1_notification_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_v1_notification_proto_goTypes,
		DependencyIndexes: file_protobuf_v1_notification_proto_depIdxs,
		MessageInfos:      file_protobuf_v1_notification_proto_msgTypes,
	}.Build()
	File_protobuf_v1_notification_proto = out.File
	file_protobuf_v1_notification_proto_rawDesc = nil
	file_protobuf_v1_notification_proto_goTypes = nil
	file_protobuf_v1_notification_proto_depIdxs = nil
}
