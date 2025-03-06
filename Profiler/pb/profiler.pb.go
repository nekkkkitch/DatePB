// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: profiler.proto

package Profiler

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

type Name struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Name) Reset() {
	*x = Name{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profiler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Name) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Name) ProtoMessage() {}

func (x *Name) ProtoReflect() protoreflect.Message {
	mi := &file_profiler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Name.ProtoReflect.Descriptor instead.
func (*Name) Descriptor() ([]byte, []int) {
	return file_profiler_proto_rawDescGZIP(), []int{0}
}

func (x *Name) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Name) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Description struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Description) Reset() {
	*x = Description{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profiler_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Description) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Description) ProtoMessage() {}

func (x *Description) ProtoReflect() protoreflect.Message {
	mi := &file_profiler_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Description.ProtoReflect.Descriptor instead.
func (*Description) Descriptor() ([]byte, []int) {
	return file_profiler_proto_rawDescGZIP(), []int{1}
}

func (x *Description) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Description) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type Age struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Age int32  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *Age) Reset() {
	*x = Age{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profiler_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Age) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Age) ProtoMessage() {}

func (x *Age) ProtoReflect() protoreflect.Message {
	mi := &file_profiler_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Age.ProtoReflect.Descriptor instead.
func (*Age) Descriptor() ([]byte, []int) {
	return file_profiler_proto_rawDescGZIP(), []int{2}
}

func (x *Age) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Age) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

type Town struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Town string `protobuf:"bytes,2,opt,name=town,proto3" json:"town,omitempty"`
}

func (x *Town) Reset() {
	*x = Town{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profiler_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Town) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Town) ProtoMessage() {}

func (x *Town) ProtoReflect() protoreflect.Message {
	mi := &file_profiler_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Town.ProtoReflect.Descriptor instead.
func (*Town) Descriptor() ([]byte, []int) {
	return file_profiler_proto_rawDescGZIP(), []int{3}
}

func (x *Town) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Town) GetTown() string {
	if x != nil {
		return x.Town
	}
	return ""
}

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  []byte  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Lat float32 `protobuf:"fixed32,2,opt,name=lat,proto3" json:"lat,omitempty"`
	Lon float32 `protobuf:"fixed32,3,opt,name=lon,proto3" json:"lon,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profiler_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_profiler_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_profiler_proto_rawDescGZIP(), []int{4}
}

func (x *Location) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Location) GetLat() float32 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *Location) GetLon() float32 {
	if x != nil {
		return x.Lon
	}
	return 0
}

type Hobbies struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Hobbies []byte `protobuf:"bytes,2,opt,name=hobbies,proto3" json:"hobbies,omitempty"`
}

func (x *Hobbies) Reset() {
	*x = Hobbies{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profiler_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hobbies) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hobbies) ProtoMessage() {}

func (x *Hobbies) ProtoReflect() protoreflect.Message {
	mi := &file_profiler_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hobbies.ProtoReflect.Descriptor instead.
func (*Hobbies) Descriptor() ([]byte, []int) {
	return file_profiler_proto_rawDescGZIP(), []int{5}
}

func (x *Hobbies) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Hobbies) GetHobbies() []byte {
	if x != nil {
		return x.Hobbies
	}
	return nil
}

type Photo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Photo []byte `protobuf:"bytes,2,opt,name=photo,proto3" json:"photo,omitempty"`
}

func (x *Photo) Reset() {
	*x = Photo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profiler_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Photo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Photo) ProtoMessage() {}

func (x *Photo) ProtoReflect() protoreflect.Message {
	mi := &file_profiler_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Photo.ProtoReflect.Descriptor instead.
func (*Photo) Descriptor() ([]byte, []int) {
	return file_profiler_proto_rawDescGZIP(), []int{6}
}

func (x *Photo) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Photo) GetPhoto() []byte {
	if x != nil {
		return x.Photo
	}
	return nil
}

type PhotoID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PhotoID []byte `protobuf:"bytes,2,opt,name=photoID,proto3" json:"photoID,omitempty"`
}

func (x *PhotoID) Reset() {
	*x = PhotoID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profiler_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhotoID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhotoID) ProtoMessage() {}

func (x *PhotoID) ProtoReflect() protoreflect.Message {
	mi := &file_profiler_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhotoID.ProtoReflect.Descriptor instead.
func (*PhotoID) Descriptor() ([]byte, []int) {
	return file_profiler_proto_rawDescGZIP(), []int{7}
}

func (x *PhotoID) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *PhotoID) GetPhotoID() []byte {
	if x != nil {
		return x.PhotoID
	}
	return nil
}

type AgePrefs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Min int32  `protobuf:"varint,2,opt,name=min,proto3" json:"min,omitempty"`
	Max int32  `protobuf:"varint,3,opt,name=max,proto3" json:"max,omitempty"`
}

func (x *AgePrefs) Reset() {
	*x = AgePrefs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profiler_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgePrefs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgePrefs) ProtoMessage() {}

func (x *AgePrefs) ProtoReflect() protoreflect.Message {
	mi := &file_profiler_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgePrefs.ProtoReflect.Descriptor instead.
func (*AgePrefs) Descriptor() ([]byte, []int) {
	return file_profiler_proto_rawDescGZIP(), []int{8}
}

func (x *AgePrefs) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *AgePrefs) GetMin() int32 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *AgePrefs) GetMax() int32 {
	if x != nil {
		return x.Max
	}
	return 0
}

type Radius struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Radius int32  `protobuf:"varint,2,opt,name=radius,proto3" json:"radius,omitempty"`
}

func (x *Radius) Reset() {
	*x = Radius{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profiler_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Radius) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Radius) ProtoMessage() {}

func (x *Radius) ProtoReflect() protoreflect.Message {
	mi := &file_profiler_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Radius.ProtoReflect.Descriptor instead.
func (*Radius) Descriptor() ([]byte, []int) {
	return file_profiler_proto_rawDescGZIP(), []int{9}
}

func (x *Radius) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Radius) GetRadius() int32 {
	if x != nil {
		return x.Radius
	}
	return 0
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status            int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	AdditionalMessage string `protobuf:"bytes,2,opt,name=additionalMessage,proto3" json:"additionalMessage,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profiler_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_profiler_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_profiler_proto_rawDescGZIP(), []int{10}
}

func (x *Status) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Status) GetAdditionalMessage() string {
	if x != nil {
		return x.AdditionalMessage
	}
	return ""
}

var File_profiler_proto protoreflect.FileDescriptor

var file_profiler_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x22, 0x2a, 0x0a, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3f, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x27, 0x0a, 0x03, 0x41, 0x67, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65,
	0x22, 0x2a, 0x0a, 0x04, 0x54, 0x6f, 0x77, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x6f, 0x77, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x6f, 0x77, 0x6e, 0x22, 0x3e, 0x0a, 0x08,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6c, 0x6f, 0x6e, 0x22, 0x33, 0x0a, 0x07,
	0x48, 0x6f, 0x62, 0x62, 0x69, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x6f, 0x62, 0x62, 0x69,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x68, 0x6f, 0x62, 0x62, 0x69, 0x65,
	0x73, 0x22, 0x2d, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68,
	0x6f, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f,
	0x22, 0x33, 0x0a, 0x07, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x68, 0x6f, 0x74, 0x6f, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x68,
	0x6f, 0x74, 0x6f, 0x49, 0x44, 0x22, 0x3e, 0x0a, 0x08, 0x41, 0x67, 0x65, 0x50, 0x72, 0x65, 0x66,
	0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x6d, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x6d, 0x61, 0x78, 0x22, 0x30, 0x0a, 0x06, 0x52, 0x61, 0x64, 0x69, 0x75, 0x73, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x22, 0x4e, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2c, 0x0a, 0x11, 0x61, 0x64, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xb6, 0x04, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x0e, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x4e, 0x61,
	0x6d, 0x65, 0x1a, 0x10, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x10, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x44, 0x65, 0x73, 0x63, 0x74, 0x69, 0x70, 0x69, 0x6f, 0x6e, 0x12, 0x15, 0x2e, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x1a, 0x10, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x41,
	0x67, 0x65, 0x12, 0x0d, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x41, 0x67,
	0x65, 0x1a, 0x10, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54,
	0x6f, 0x77, 0x6e, 0x12, 0x0e, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x54,
	0x6f, 0x77, 0x6e, 0x1a, 0x10, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x2e, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x10, 0x2e,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x00, 0x12, 0x36, 0x0a, 0x0d, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x48, 0x6f, 0x62, 0x62, 0x69,
	0x65, 0x73, 0x12, 0x11, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x48, 0x6f,
	0x62, 0x62, 0x69, 0x65, 0x73, 0x1a, 0x10, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0b, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x72, 0x2e, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x34, 0x0a,
	0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x2e, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x49, 0x44, 0x1a,
	0x10, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x12, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x41, 0x67, 0x65, 0x12, 0x12, 0x2e, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x72, 0x2e, 0x41, 0x67, 0x65, 0x50, 0x72, 0x65, 0x66, 0x73, 0x1a, 0x10, 0x2e,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x00, 0x12, 0x3d, 0x0a, 0x15, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x72, 0x65, 0x64, 0x52, 0x61, 0x64, 0x69, 0x75, 0x73, 0x12, 0x10, 0x2e, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x52, 0x61, 0x64, 0x69, 0x75, 0x73, 0x1a, 0x10, 0x2e, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00,
	0x42, 0x02, 0x5a, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_profiler_proto_rawDescOnce sync.Once
	file_profiler_proto_rawDescData = file_profiler_proto_rawDesc
)

func file_profiler_proto_rawDescGZIP() []byte {
	file_profiler_proto_rawDescOnce.Do(func() {
		file_profiler_proto_rawDescData = protoimpl.X.CompressGZIP(file_profiler_proto_rawDescData)
	})
	return file_profiler_proto_rawDescData
}

var file_profiler_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_profiler_proto_goTypes = []interface{}{
	(*Name)(nil),        // 0: Profiler.Name
	(*Description)(nil), // 1: Profiler.Description
	(*Age)(nil),         // 2: Profiler.Age
	(*Town)(nil),        // 3: Profiler.Town
	(*Location)(nil),    // 4: Profiler.Location
	(*Hobbies)(nil),     // 5: Profiler.Hobbies
	(*Photo)(nil),       // 6: Profiler.Photo
	(*PhotoID)(nil),     // 7: Profiler.PhotoID
	(*AgePrefs)(nil),    // 8: Profiler.AgePrefs
	(*Radius)(nil),      // 9: Profiler.Radius
	(*Status)(nil),      // 10: Profiler.Status
}
var file_profiler_proto_depIdxs = []int32{
	0,  // 0: Profiler.Profiler.ChangeName:input_type -> Profiler.Name
	1,  // 1: Profiler.Profiler.ChangeDesctipion:input_type -> Profiler.Description
	2,  // 2: Profiler.Profiler.ChangeAge:input_type -> Profiler.Age
	3,  // 3: Profiler.Profiler.ChangeTown:input_type -> Profiler.Town
	4,  // 4: Profiler.Profiler.UpdateLocation:input_type -> Profiler.Location
	5,  // 5: Profiler.Profiler.ChangeHobbies:input_type -> Profiler.Hobbies
	6,  // 6: Profiler.Profiler.UploadPhoto:input_type -> Profiler.Photo
	7,  // 7: Profiler.Profiler.DeletePhoto:input_type -> Profiler.PhotoID
	8,  // 8: Profiler.Profiler.ChangePreferredAge:input_type -> Profiler.AgePrefs
	9,  // 9: Profiler.Profiler.ChangePreferredRadius:input_type -> Profiler.Radius
	10, // 10: Profiler.Profiler.ChangeName:output_type -> Profiler.Status
	10, // 11: Profiler.Profiler.ChangeDesctipion:output_type -> Profiler.Status
	10, // 12: Profiler.Profiler.ChangeAge:output_type -> Profiler.Status
	10, // 13: Profiler.Profiler.ChangeTown:output_type -> Profiler.Status
	10, // 14: Profiler.Profiler.UpdateLocation:output_type -> Profiler.Status
	10, // 15: Profiler.Profiler.ChangeHobbies:output_type -> Profiler.Status
	10, // 16: Profiler.Profiler.UploadPhoto:output_type -> Profiler.Status
	10, // 17: Profiler.Profiler.DeletePhoto:output_type -> Profiler.Status
	10, // 18: Profiler.Profiler.ChangePreferredAge:output_type -> Profiler.Status
	10, // 19: Profiler.Profiler.ChangePreferredRadius:output_type -> Profiler.Status
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_profiler_proto_init() }
func file_profiler_proto_init() {
	if File_profiler_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_profiler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Name); i {
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
		file_profiler_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Description); i {
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
		file_profiler_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Age); i {
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
		file_profiler_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Town); i {
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
		file_profiler_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_profiler_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hobbies); i {
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
		file_profiler_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Photo); i {
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
		file_profiler_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhotoID); i {
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
		file_profiler_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgePrefs); i {
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
		file_profiler_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Radius); i {
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
		file_profiler_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
			RawDescriptor: file_profiler_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_profiler_proto_goTypes,
		DependencyIndexes: file_profiler_proto_depIdxs,
		MessageInfos:      file_profiler_proto_msgTypes,
	}.Build()
	File_profiler_proto = out.File
	file_profiler_proto_rawDesc = nil
	file_profiler_proto_goTypes = nil
	file_profiler_proto_depIdxs = nil
}
