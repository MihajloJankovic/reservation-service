// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: appres.proto

package genfiles

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

type Emaill struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *Emaill) Reset() {
	*x = Emaill{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appres_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Emaill) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Emaill) ProtoMessage() {}

func (x *Emaill) ProtoReflect() protoreflect.Message {
	mi := &file_appres_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Emaill.ProtoReflect.Descriptor instead.
func (*Emaill) Descriptor() ([]byte, []int) {
	return file_appres_proto_rawDescGZIP(), []int{0}
}

func (x *Emaill) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type EmailCheck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DateFrom string `protobuf:"bytes,1,opt,name=dateFrom,proto3" json:"dateFrom,omitempty"`
	DateTo   string `protobuf:"bytes,2,opt,name=dateTo,proto3" json:"dateTo,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *EmailCheck) Reset() {
	*x = EmailCheck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appres_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailCheck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailCheck) ProtoMessage() {}

func (x *EmailCheck) ProtoReflect() protoreflect.Message {
	mi := &file_appres_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailCheck.ProtoReflect.Descriptor instead.
func (*EmailCheck) Descriptor() ([]byte, []int) {
	return file_appres_proto_rawDescGZIP(), []int{1}
}

func (x *EmailCheck) GetDateFrom() string {
	if x != nil {
		return x.DateFrom
	}
	return ""
}

func (x *EmailCheck) GetDateTo() string {
	if x != nil {
		return x.DateTo
	}
	return ""
}

func (x *EmailCheck) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type DateFromDateTo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DateFrom string `protobuf:"bytes,1,opt,name=dateFrom,proto3" json:"dateFrom,omitempty"`
	DateTo   string `protobuf:"bytes,2,opt,name=dateTo,proto3" json:"dateTo,omitempty"`
	Accid    string `protobuf:"bytes,3,opt,name=accid,proto3" json:"accid,omitempty"`
}

func (x *DateFromDateTo) Reset() {
	*x = DateFromDateTo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appres_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DateFromDateTo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DateFromDateTo) ProtoMessage() {}

func (x *DateFromDateTo) ProtoReflect() protoreflect.Message {
	mi := &file_appres_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DateFromDateTo.ProtoReflect.Descriptor instead.
func (*DateFromDateTo) Descriptor() ([]byte, []int) {
	return file_appres_proto_rawDescGZIP(), []int{2}
}

func (x *DateFromDateTo) GetDateFrom() string {
	if x != nil {
		return x.DateFrom
	}
	return ""
}

func (x *DateFromDateTo) GetDateTo() string {
	if x != nil {
		return x.DateTo
	}
	return ""
}

func (x *DateFromDateTo) GetAccid() string {
	if x != nil {
		return x.Accid
	}
	return ""
}

type DummyLista struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dummy []*ReservationResponse `protobuf:"bytes,1,rep,name=dummy,proto3" json:"dummy,omitempty"`
}

func (x *DummyLista) Reset() {
	*x = DummyLista{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appres_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DummyLista) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DummyLista) ProtoMessage() {}

func (x *DummyLista) ProtoReflect() protoreflect.Message {
	mi := &file_appres_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DummyLista.ProtoReflect.Descriptor instead.
func (*DummyLista) Descriptor() ([]byte, []int) {
	return file_appres_proto_rawDescGZIP(), []int{3}
}

func (x *DummyLista) GetDummy() []*ReservationResponse {
	if x != nil {
		return x.Dummy
	}
	return nil
}

type DeleteRequestaa struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *DeleteRequestaa) Reset() {
	*x = DeleteRequestaa{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appres_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequestaa) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequestaa) ProtoMessage() {}

func (x *DeleteRequestaa) ProtoReflect() protoreflect.Message {
	mi := &file_appres_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequestaa.ProtoReflect.Descriptor instead.
func (*DeleteRequestaa) Descriptor() ([]byte, []int) {
	return file_appres_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteRequestaa) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type ReservationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReservationRequest) Reset() {
	*x = ReservationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appres_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReservationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReservationRequest) ProtoMessage() {}

func (x *ReservationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_appres_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReservationRequest.ProtoReflect.Descriptor instead.
func (*ReservationRequest) Descriptor() ([]byte, []int) {
	return file_appres_proto_rawDescGZIP(), []int{5}
}

func (x *ReservationRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ReservationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	DateFrom string `protobuf:"bytes,3,opt,name=dateFrom,proto3" json:"dateFrom,omitempty"`
	DateTo   string `protobuf:"bytes,4,opt,name=dateTo,proto3" json:"dateTo,omitempty"`
	Accid    string `protobuf:"bytes,5,opt,name=accid,proto3" json:"accid,omitempty"`
}

func (x *ReservationResponse) Reset() {
	*x = ReservationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appres_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReservationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReservationResponse) ProtoMessage() {}

func (x *ReservationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_appres_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReservationResponse.ProtoReflect.Descriptor instead.
func (*ReservationResponse) Descriptor() ([]byte, []int) {
	return file_appres_proto_rawDescGZIP(), []int{6}
}

func (x *ReservationResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ReservationResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ReservationResponse) GetDateFrom() string {
	if x != nil {
		return x.DateFrom
	}
	return ""
}

func (x *ReservationResponse) GetDateTo() string {
	if x != nil {
		return x.DateTo
	}
	return ""
}

func (x *ReservationResponse) GetAccid() string {
	if x != nil {
		return x.Accid
	}
	return ""
}

type Emptyaa struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Emptyaa) Reset() {
	*x = Emptyaa{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appres_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Emptyaa) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Emptyaa) ProtoMessage() {}

func (x *Emptyaa) ProtoReflect() protoreflect.Message {
	mi := &file_appres_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Emptyaa.ProtoReflect.Descriptor instead.
func (*Emptyaa) Descriptor() ([]byte, []int) {
	return file_appres_proto_rawDescGZIP(), []int{7}
}

var File_appres_proto protoreflect.FileDescriptor

var file_appres_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x61, 0x70, 0x70, 0x72, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1e,
	0x0a, 0x06, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x56,
	0x0a, 0x0a, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x1a, 0x0a, 0x08,
	0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x5a, 0x0a, 0x0e, 0x44, 0x61, 0x74, 0x65, 0x46, 0x72,
	0x6f, 0x6d, 0x44, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x65,
	0x46, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x65,
	0x46, 0x72, 0x6f, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x12, 0x14, 0x0a, 0x05,
	0x61, 0x63, 0x63, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x63, 0x63,
	0x69, 0x64, 0x22, 0x38, 0x0a, 0x0a, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x61,
	0x12, 0x2a, 0x0a, 0x05, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x22, 0x23, 0x0a, 0x0f,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x61, 0x61, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69,
	0x64, 0x22, 0x24, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x85, 0x01, 0x0a, 0x13, 0x52, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f,
	0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f,
	0x6d, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x63, 0x63,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x63, 0x63, 0x69, 0x64, 0x22,
	0x09, 0x0a, 0x07, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x61, 0x61, 0x32, 0x8c, 0x04, 0x0a, 0x0b, 0x72,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x13, 0x2e, 0x52,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0b, 0x2e, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x61, 0x12, 0x2b,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x08, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x61, 0x61, 0x1a, 0x0b,
	0x2e, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x61, 0x12, 0x30, 0x0a, 0x0e, 0x53,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x2e,
	0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x1a, 0x08, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x61, 0x61, 0x12, 0x33, 0x0a,
	0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x14, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x08, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x61, 0x61, 0x12, 0x34, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x41, 0x63,
	0x63, 0x6f, 0x6d, 0x6e, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x61, 0x61, 0x1a, 0x08,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x61, 0x61, 0x12, 0x33, 0x0a, 0x16, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x0f, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x44, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x1a, 0x08, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x61, 0x61, 0x12, 0x34, 0x0a,
	0x1b, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x2e, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x1a, 0x08, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x61, 0x61, 0x12, 0x32, 0x0a, 0x1d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x07, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x6c, 0x1a, 0x08, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x61, 0x61, 0x12, 0x2d, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x07, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x6c, 0x1a, 0x08, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x61, 0x61, 0x12, 0x31, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x07, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x6c, 0x1a, 0x0b, 0x2e, 0x44,
	0x75, 0x6d, 0x6d, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x61, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2e, 0x2f,
	0x67, 0x65, 0x6e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_appres_proto_rawDescOnce sync.Once
	file_appres_proto_rawDescData = file_appres_proto_rawDesc
)

func file_appres_proto_rawDescGZIP() []byte {
	file_appres_proto_rawDescOnce.Do(func() {
		file_appres_proto_rawDescData = protoimpl.X.CompressGZIP(file_appres_proto_rawDescData)
	})
	return file_appres_proto_rawDescData
}

var file_appres_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_appres_proto_goTypes = []interface{}{
	(*Emaill)(nil),              // 0: Emaill
	(*EmailCheck)(nil),          // 1: EmailCheck
	(*DateFromDateTo)(nil),      // 2: DateFromDateTo
	(*DummyLista)(nil),          // 3: DummyLista
	(*DeleteRequestaa)(nil),     // 4: DeleteRequestaa
	(*ReservationRequest)(nil),  // 5: ReservationRequest
	(*ReservationResponse)(nil), // 6: ReservationResponse
	(*Emptyaa)(nil),             // 7: Emptyaa
}
var file_appres_proto_depIdxs = []int32{
	6,  // 0: DummyLista.dummy:type_name -> ReservationResponse
	5,  // 1: reservation.GetReservation:input_type -> ReservationRequest
	7,  // 2: reservation.GetAllReservations:input_type -> Emptyaa
	6,  // 3: reservation.SetReservation:input_type -> ReservationResponse
	6,  // 4: reservation.UpdateReservation:input_type -> ReservationResponse
	4,  // 5: reservation.DeleteByAccomnendation:input_type -> DeleteRequestaa
	2,  // 6: reservation.CheckActiveReservation:input_type -> DateFromDateTo
	1,  // 7: reservation.CheckUsersActiveReservation:input_type -> EmailCheck
	0,  // 8: reservation.CheckActiveReservationByEmail:input_type -> Emaill
	0,  // 9: reservation.DeleteReservationByEmail:input_type -> Emaill
	0,  // 10: reservation.GetAllReservationsByEmail:input_type -> Emaill
	3,  // 11: reservation.GetReservation:output_type -> DummyLista
	3,  // 12: reservation.GetAllReservations:output_type -> DummyLista
	7,  // 13: reservation.SetReservation:output_type -> Emptyaa
	7,  // 14: reservation.UpdateReservation:output_type -> Emptyaa
	7,  // 15: reservation.DeleteByAccomnendation:output_type -> Emptyaa
	7,  // 16: reservation.CheckActiveReservation:output_type -> Emptyaa
	7,  // 17: reservation.CheckUsersActiveReservation:output_type -> Emptyaa
	7,  // 18: reservation.CheckActiveReservationByEmail:output_type -> Emptyaa
	7,  // 19: reservation.DeleteReservationByEmail:output_type -> Emptyaa
	3,  // 20: reservation.GetAllReservationsByEmail:output_type -> DummyLista
	11, // [11:21] is the sub-list for method output_type
	1,  // [1:11] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_appres_proto_init() }
func file_appres_proto_init() {
	if File_appres_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_appres_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Emaill); i {
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
		file_appres_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailCheck); i {
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
		file_appres_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DateFromDateTo); i {
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
		file_appres_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DummyLista); i {
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
		file_appres_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequestaa); i {
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
		file_appres_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReservationRequest); i {
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
		file_appres_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReservationResponse); i {
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
		file_appres_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Emptyaa); i {
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
			RawDescriptor: file_appres_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_appres_proto_goTypes,
		DependencyIndexes: file_appres_proto_depIdxs,
		MessageInfos:      file_appres_proto_msgTypes,
	}.Build()
	File_appres_proto = out.File
	file_appres_proto_rawDesc = nil
	file_appres_proto_goTypes = nil
	file_appres_proto_depIdxs = nil
}
