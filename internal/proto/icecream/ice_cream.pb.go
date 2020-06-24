// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.11.4
// source: api/proto/ice_cream.proto

package icecream_pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type IceCream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId             string   `protobuf:"bytes,1,opt,name=ProductId,proto3" json:"ProductId,omitempty"`
	Name                  string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	ImageClosed           string   `protobuf:"bytes,3,opt,name=ImageClosed,proto3" json:"ImageClosed,omitempty"`
	ImageOpen             string   `protobuf:"bytes,4,opt,name=ImageOpen,proto3" json:"ImageOpen,omitempty"`
	Description           string   `protobuf:"bytes,5,opt,name=Description,proto3" json:"Description,omitempty"`
	Story                 string   `protobuf:"bytes,6,opt,name=Story,proto3" json:"Story,omitempty"`
	SourcingValues        []string `protobuf:"bytes,7,rep,name=SourcingValues,proto3" json:"SourcingValues,omitempty"`
	Ingredients           []string `protobuf:"bytes,8,rep,name=Ingredients,proto3" json:"Ingredients,omitempty"`
	AllergyInfo           string   `protobuf:"bytes,9,opt,name=AllergyInfo,proto3" json:"AllergyInfo,omitempty"`
	DietaryCertifications string   `protobuf:"bytes,10,opt,name=DietaryCertifications,proto3" json:"DietaryCertifications,omitempty"`
}

func (x *IceCream) Reset() {
	*x = IceCream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_ice_cream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IceCream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IceCream) ProtoMessage() {}

func (x *IceCream) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_ice_cream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IceCream.ProtoReflect.Descriptor instead.
func (*IceCream) Descriptor() ([]byte, []int) {
	return file_api_proto_ice_cream_proto_rawDescGZIP(), []int{0}
}

func (x *IceCream) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *IceCream) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *IceCream) GetImageClosed() string {
	if x != nil {
		return x.ImageClosed
	}
	return ""
}

func (x *IceCream) GetImageOpen() string {
	if x != nil {
		return x.ImageOpen
	}
	return ""
}

func (x *IceCream) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *IceCream) GetStory() string {
	if x != nil {
		return x.Story
	}
	return ""
}

func (x *IceCream) GetSourcingValues() []string {
	if x != nil {
		return x.SourcingValues
	}
	return nil
}

func (x *IceCream) GetIngredients() []string {
	if x != nil {
		return x.Ingredients
	}
	return nil
}

func (x *IceCream) GetAllergyInfo() string {
	if x != nil {
		return x.AllergyInfo
	}
	return ""
}

func (x *IceCream) GetDietaryCertifications() string {
	if x != nil {
		return x.DietaryCertifications
	}
	return ""
}

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IceCream *IceCream `protobuf:"bytes,1,opt,name=ice_cream,json=iceCream,proto3" json:"ice_cream,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_ice_cream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_ice_cream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_ice_cream_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRequest) GetIceCream() *IceCream {
	if x != nil {
		return x.IceCream
	}
	return nil
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IceCream *IceCream `protobuf:"bytes,1,opt,name=ice_cream,json=iceCream,proto3" json:"ice_cream,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_ice_cream_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_ice_cream_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_ice_cream_proto_rawDescGZIP(), []int{2}
}

func (x *CreateResponse) GetIceCream() *IceCream {
	if x != nil {
		return x.IceCream
	}
	return nil
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IceCream *IceCream `protobuf:"bytes,1,opt,name=ice_cream,json=iceCream,proto3" json:"ice_cream,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_ice_cream_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_ice_cream_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_ice_cream_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateRequest) GetIceCream() *IceCream {
	if x != nil {
		return x.IceCream
	}
	return nil
}

type UpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IceCream *IceCream `protobuf:"bytes,1,opt,name=ice_cream,json=iceCream,proto3" json:"ice_cream,omitempty"`
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_ice_cream_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_ice_cream_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResponse.ProtoReflect.Descriptor instead.
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_ice_cream_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateResponse) GetIceCream() *IceCream {
	if x != nil {
		return x.IceCream
	}
	return nil
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_ice_cream_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_ice_cream_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_ice_cream_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IceCream *IceCream `protobuf:"bytes,1,opt,name=ice_cream,json=iceCream,proto3" json:"ice_cream,omitempty"`
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_ice_cream_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_ice_cream_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_ice_cream_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteResponse) GetIceCream() *IceCream {
	if x != nil {
		return x.IceCream
	}
	return nil
}

type GetByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *GetByIdRequest) Reset() {
	*x = GetByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_ice_cream_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByIdRequest) ProtoMessage() {}

func (x *GetByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_ice_cream_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByIdRequest.ProtoReflect.Descriptor instead.
func (*GetByIdRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_ice_cream_proto_rawDescGZIP(), []int{7}
}

func (x *GetByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IceCream *IceCream `protobuf:"bytes,1,opt,name=ice_cream,json=iceCream,proto3" json:"ice_cream,omitempty"`
}

func (x *GetByIdResponse) Reset() {
	*x = GetByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_ice_cream_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByIdResponse) ProtoMessage() {}

func (x *GetByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_ice_cream_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByIdResponse.ProtoReflect.Descriptor instead.
func (*GetByIdResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_ice_cream_proto_rawDescGZIP(), []int{8}
}

func (x *GetByIdResponse) GetIceCream() *IceCream {
	if x != nil {
		return x.IceCream
	}
	return nil
}

var File_api_proto_ice_cream_proto protoreflect.FileDescriptor

var file_api_proto_ice_cream_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x63, 0x65, 0x5f,
	0x63, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x69, 0x63, 0x65,
	0x63, 0x72, 0x65, 0x61, 0x6d, 0x22, 0xd6, 0x02, 0x0a, 0x08, 0x49, 0x63, 0x65, 0x43, 0x72, 0x65,
	0x61, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x6c, 0x6f,
	0x73, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x43, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4f,
	0x70, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x4f, 0x70, 0x65, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x26, 0x0a, 0x0e,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x49, 0x6e, 0x67, 0x72, 0x65,
	0x64, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x6c, 0x6c, 0x65, 0x72, 0x67,
	0x79, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x6c, 0x6c,
	0x65, 0x72, 0x67, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x34, 0x0a, 0x15, 0x44, 0x69, 0x65, 0x74,
	0x61, 0x72, 0x79, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x44, 0x69, 0x65, 0x74, 0x61, 0x72, 0x79,
	0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x40,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2f, 0x0a, 0x09, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x69, 0x63, 0x65, 0x63, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x49, 0x63,
	0x65, 0x43, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x08, 0x69, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x6d,
	0x22, 0x41, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2f, 0x0a, 0x09, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x69, 0x63, 0x65, 0x63, 0x72, 0x65, 0x61, 0x6d,
	0x2e, 0x49, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x08, 0x69, 0x63, 0x65, 0x43, 0x72,
	0x65, 0x61, 0x6d, 0x22, 0x40, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x09, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x69, 0x63, 0x65, 0x63, 0x72, 0x65,
	0x61, 0x6d, 0x2e, 0x49, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x08, 0x69, 0x63, 0x65,
	0x43, 0x72, 0x65, 0x61, 0x6d, 0x22, 0x41, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x09, 0x69, 0x63, 0x65, 0x5f, 0x63,
	0x72, 0x65, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x69, 0x63, 0x65,
	0x63, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x49, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x08,
	0x69, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x6d, 0x22, 0x1f, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x0e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x09, 0x69,
	0x63, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x69, 0x63, 0x65, 0x63, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x49, 0x63, 0x65, 0x43, 0x72, 0x65,
	0x61, 0x6d, 0x52, 0x08, 0x69, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x6d, 0x22, 0x20, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x22, 0x42,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2f, 0x0a, 0x09, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x69, 0x63, 0x65, 0x63, 0x72, 0x65, 0x61, 0x6d, 0x2e,
	0x49, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x08, 0x69, 0x63, 0x65, 0x43, 0x72, 0x65,
	0x61, 0x6d, 0x32, 0x8c, 0x02, 0x0a, 0x0b, 0x49, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x6d, 0x41,
	0x70, 0x69, 0x12, 0x3d, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x69,
	0x63, 0x65, 0x63, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x69, 0x63, 0x65, 0x63, 0x72, 0x65, 0x61, 0x6d,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3d, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x69, 0x63,
	0x65, 0x63, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x69, 0x63, 0x65, 0x63, 0x72, 0x65, 0x61, 0x6d, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3d, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x69, 0x63, 0x65,
	0x63, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x69, 0x63, 0x65, 0x63, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x40, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x18, 0x2e, 0x69, 0x63, 0x65,
	0x63, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x69, 0x63, 0x65, 0x63, 0x72, 0x65, 0x61, 0x6d, 0x2e,
	0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x16, 0x5a, 0x14, 0x69, 0x63, 0x65, 0x63, 0x72, 0x65, 0x61, 0x6d, 0x3b, 0x69, 0x63,
	0x65, 0x63, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_proto_ice_cream_proto_rawDescOnce sync.Once
	file_api_proto_ice_cream_proto_rawDescData = file_api_proto_ice_cream_proto_rawDesc
)

func file_api_proto_ice_cream_proto_rawDescGZIP() []byte {
	file_api_proto_ice_cream_proto_rawDescOnce.Do(func() {
		file_api_proto_ice_cream_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_ice_cream_proto_rawDescData)
	})
	return file_api_proto_ice_cream_proto_rawDescData
}

var file_api_proto_ice_cream_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_proto_ice_cream_proto_goTypes = []interface{}{
	(*IceCream)(nil),        // 0: icecream.IceCream
	(*CreateRequest)(nil),   // 1: icecream.CreateRequest
	(*CreateResponse)(nil),  // 2: icecream.CreateResponse
	(*UpdateRequest)(nil),   // 3: icecream.UpdateRequest
	(*UpdateResponse)(nil),  // 4: icecream.UpdateResponse
	(*DeleteRequest)(nil),   // 5: icecream.DeleteRequest
	(*DeleteResponse)(nil),  // 6: icecream.DeleteResponse
	(*GetByIdRequest)(nil),  // 7: icecream.GetByIdRequest
	(*GetByIdResponse)(nil), // 8: icecream.GetByIdResponse
}
var file_api_proto_ice_cream_proto_depIdxs = []int32{
	0,  // 0: icecream.CreateRequest.ice_cream:type_name -> icecream.IceCream
	0,  // 1: icecream.CreateResponse.ice_cream:type_name -> icecream.IceCream
	0,  // 2: icecream.UpdateRequest.ice_cream:type_name -> icecream.IceCream
	0,  // 3: icecream.UpdateResponse.ice_cream:type_name -> icecream.IceCream
	0,  // 4: icecream.DeleteResponse.ice_cream:type_name -> icecream.IceCream
	0,  // 5: icecream.GetByIdResponse.ice_cream:type_name -> icecream.IceCream
	1,  // 6: icecream.IceCreamApi.Create:input_type -> icecream.CreateRequest
	3,  // 7: icecream.IceCreamApi.Update:input_type -> icecream.UpdateRequest
	5,  // 8: icecream.IceCreamApi.Delete:input_type -> icecream.DeleteRequest
	7,  // 9: icecream.IceCreamApi.GetById:input_type -> icecream.GetByIdRequest
	2,  // 10: icecream.IceCreamApi.Create:output_type -> icecream.CreateResponse
	4,  // 11: icecream.IceCreamApi.Update:output_type -> icecream.UpdateResponse
	6,  // 12: icecream.IceCreamApi.Delete:output_type -> icecream.DeleteResponse
	8,  // 13: icecream.IceCreamApi.GetById:output_type -> icecream.GetByIdResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_api_proto_ice_cream_proto_init() }
func file_api_proto_ice_cream_proto_init() {
	if File_api_proto_ice_cream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_ice_cream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IceCream); i {
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
		file_api_proto_ice_cream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_api_proto_ice_cream_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
		file_api_proto_ice_cream_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_api_proto_ice_cream_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResponse); i {
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
		file_api_proto_ice_cream_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_api_proto_ice_cream_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
		file_api_proto_ice_cream_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByIdRequest); i {
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
		file_api_proto_ice_cream_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByIdResponse); i {
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
			RawDescriptor: file_api_proto_ice_cream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_ice_cream_proto_goTypes,
		DependencyIndexes: file_api_proto_ice_cream_proto_depIdxs,
		MessageInfos:      file_api_proto_ice_cream_proto_msgTypes,
	}.Build()
	File_api_proto_ice_cream_proto = out.File
	file_api_proto_ice_cream_proto_rawDesc = nil
	file_api_proto_ice_cream_proto_goTypes = nil
	file_api_proto_ice_cream_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// IceCreamApiClient is the client API for IceCreamApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IceCreamApiClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	GetById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*GetByIdResponse, error)
}

type iceCreamApiClient struct {
	cc grpc.ClientConnInterface
}

func NewIceCreamApiClient(cc grpc.ClientConnInterface) IceCreamApiClient {
	return &iceCreamApiClient{cc}
}

func (c *iceCreamApiClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/icecream.IceCreamApi/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iceCreamApiClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/icecream.IceCreamApi/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iceCreamApiClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/icecream.IceCreamApi/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iceCreamApiClient) GetById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*GetByIdResponse, error) {
	out := new(GetByIdResponse)
	err := c.cc.Invoke(ctx, "/icecream.IceCreamApi/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IceCreamApiServer is the server API for IceCreamApi service.
type IceCreamApiServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	GetById(context.Context, *GetByIdRequest) (*GetByIdResponse, error)
}

// UnimplementedIceCreamApiServer can be embedded to have forward compatible implementations.
type UnimplementedIceCreamApiServer struct {
}

func (*UnimplementedIceCreamApiServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedIceCreamApiServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedIceCreamApiServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedIceCreamApiServer) GetById(context.Context, *GetByIdRequest) (*GetByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}

func RegisterIceCreamApiServer(s *grpc.Server, srv IceCreamApiServer) {
	s.RegisterService(&_IceCreamApi_serviceDesc, srv)
}

func _IceCreamApi_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IceCreamApiServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/icecream.IceCreamApi/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IceCreamApiServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IceCreamApi_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IceCreamApiServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/icecream.IceCreamApi/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IceCreamApiServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IceCreamApi_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IceCreamApiServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/icecream.IceCreamApi/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IceCreamApiServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IceCreamApi_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IceCreamApiServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/icecream.IceCreamApi/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IceCreamApiServer).GetById(ctx, req.(*GetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IceCreamApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "icecream.IceCreamApi",
	HandlerType: (*IceCreamApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _IceCreamApi_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _IceCreamApi_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _IceCreamApi_Delete_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _IceCreamApi_GetById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/ice_cream.proto",
}
