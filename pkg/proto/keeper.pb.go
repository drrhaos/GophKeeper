// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.3
// source: keeper.proto

package proto

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login    string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keeper_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_keeper_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_keeper_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keeper_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_keeper_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_keeper_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *RegisterResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login    string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keeper_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_keeper_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_keeper_proto_rawDescGZIP(), []int{2}
}

func (x *LoginRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keeper_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_keeper_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_keeper_proto_rawDescGZIP(), []int{3}
}

func (x *LoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *LoginResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type FieldKeep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Login      string `protobuf:"bytes,2,opt,name=login,proto3" json:"login,omitempty"`
	Password   string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Data       string `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	TimeUpdate string `protobuf:"bytes,5,opt,name=timeUpdate,proto3" json:"timeUpdate,omitempty"`
	CardNumber string `protobuf:"bytes,6,opt,name=cardNumber,proto3" json:"cardNumber,omitempty"`
	CardCVC    string `protobuf:"bytes,7,opt,name=cardCVC,proto3" json:"cardCVC,omitempty"`
	CardDate   string `protobuf:"bytes,8,opt,name=cardDate,proto3" json:"cardDate,omitempty"`
	CardOwner  string `protobuf:"bytes,9,opt,name=cardOwner,proto3" json:"cardOwner,omitempty"`
}

func (x *FieldKeep) Reset() {
	*x = FieldKeep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keeper_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldKeep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldKeep) ProtoMessage() {}

func (x *FieldKeep) ProtoReflect() protoreflect.Message {
	mi := &file_keeper_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldKeep.ProtoReflect.Descriptor instead.
func (*FieldKeep) Descriptor() ([]byte, []int) {
	return file_keeper_proto_rawDescGZIP(), []int{4}
}

func (x *FieldKeep) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FieldKeep) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *FieldKeep) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *FieldKeep) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *FieldKeep) GetTimeUpdate() string {
	if x != nil {
		return x.TimeUpdate
	}
	return ""
}

func (x *FieldKeep) GetCardNumber() string {
	if x != nil {
		return x.CardNumber
	}
	return ""
}

func (x *FieldKeep) GetCardCVC() string {
	if x != nil {
		return x.CardCVC
	}
	return ""
}

func (x *FieldKeep) GetCardDate() string {
	if x != nil {
		return x.CardDate
	}
	return ""
}

func (x *FieldKeep) GetCardOwner() string {
	if x != nil {
		return x.CardOwner
	}
	return ""
}

type AddFieldKeepRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *FieldKeep `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *AddFieldKeepRequest) Reset() {
	*x = AddFieldKeepRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keeper_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFieldKeepRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFieldKeepRequest) ProtoMessage() {}

func (x *AddFieldKeepRequest) ProtoReflect() protoreflect.Message {
	mi := &file_keeper_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFieldKeepRequest.ProtoReflect.Descriptor instead.
func (*AddFieldKeepRequest) Descriptor() ([]byte, []int) {
	return file_keeper_proto_rawDescGZIP(), []int{5}
}

func (x *AddFieldKeepRequest) GetData() *FieldKeep {
	if x != nil {
		return x.Data
	}
	return nil
}

type AddFieldKeepResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid  string     `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Data  *FieldKeep `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Error string     `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *AddFieldKeepResponse) Reset() {
	*x = AddFieldKeepResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keeper_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFieldKeepResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFieldKeepResponse) ProtoMessage() {}

func (x *AddFieldKeepResponse) ProtoReflect() protoreflect.Message {
	mi := &file_keeper_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFieldKeepResponse.ProtoReflect.Descriptor instead.
func (*AddFieldKeepResponse) Descriptor() ([]byte, []int) {
	return file_keeper_proto_rawDescGZIP(), []int{6}
}

func (x *AddFieldKeepResponse) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *AddFieldKeepResponse) GetData() *FieldKeep {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *AddFieldKeepResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type EditFieldKeepRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string     `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Data *FieldKeep `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *EditFieldKeepRequest) Reset() {
	*x = EditFieldKeepRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keeper_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditFieldKeepRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditFieldKeepRequest) ProtoMessage() {}

func (x *EditFieldKeepRequest) ProtoReflect() protoreflect.Message {
	mi := &file_keeper_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditFieldKeepRequest.ProtoReflect.Descriptor instead.
func (*EditFieldKeepRequest) Descriptor() ([]byte, []int) {
	return file_keeper_proto_rawDescGZIP(), []int{7}
}

func (x *EditFieldKeepRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *EditFieldKeepRequest) GetData() *FieldKeep {
	if x != nil {
		return x.Data
	}
	return nil
}

type EditFieldKeepResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *FieldKeep `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Error string     `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *EditFieldKeepResponse) Reset() {
	*x = EditFieldKeepResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keeper_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditFieldKeepResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditFieldKeepResponse) ProtoMessage() {}

func (x *EditFieldKeepResponse) ProtoReflect() protoreflect.Message {
	mi := &file_keeper_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditFieldKeepResponse.ProtoReflect.Descriptor instead.
func (*EditFieldKeepResponse) Descriptor() ([]byte, []int) {
	return file_keeper_proto_rawDescGZIP(), []int{8}
}

func (x *EditFieldKeepResponse) GetData() *FieldKeep {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *EditFieldKeepResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type DeleteFieldKeepRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *DeleteFieldKeepRequest) Reset() {
	*x = DeleteFieldKeepRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keeper_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFieldKeepRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFieldKeepRequest) ProtoMessage() {}

func (x *DeleteFieldKeepRequest) ProtoReflect() protoreflect.Message {
	mi := &file_keeper_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFieldKeepRequest.ProtoReflect.Descriptor instead.
func (*DeleteFieldKeepRequest) Descriptor() ([]byte, []int) {
	return file_keeper_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteFieldKeepRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type DeleteFieldKeepResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid  string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DeleteFieldKeepResponse) Reset() {
	*x = DeleteFieldKeepResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keeper_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFieldKeepResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFieldKeepResponse) ProtoMessage() {}

func (x *DeleteFieldKeepResponse) ProtoReflect() protoreflect.Message {
	mi := &file_keeper_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFieldKeepResponse.ProtoReflect.Descriptor instead.
func (*DeleteFieldKeepResponse) Descriptor() ([]byte, []int) {
	return file_keeper_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteFieldKeepResponse) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *DeleteFieldKeepResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type ListFieldsKeepRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListFieldsKeepRequest) Reset() {
	*x = ListFieldsKeepRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keeper_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFieldsKeepRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFieldsKeepRequest) ProtoMessage() {}

func (x *ListFieldsKeepRequest) ProtoReflect() protoreflect.Message {
	mi := &file_keeper_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFieldsKeepRequest.ProtoReflect.Descriptor instead.
func (*ListFieldsKeepRequest) Descriptor() ([]byte, []int) {
	return file_keeper_proto_rawDescGZIP(), []int{11}
}

type ListFielsdKeepResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  map[string]*FieldKeep `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Error string                `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ListFielsdKeepResponse) Reset() {
	*x = ListFielsdKeepResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keeper_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFielsdKeepResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFielsdKeepResponse) ProtoMessage() {}

func (x *ListFielsdKeepResponse) ProtoReflect() protoreflect.Message {
	mi := &file_keeper_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFielsdKeepResponse.ProtoReflect.Descriptor instead.
func (*ListFielsdKeepResponse) Descriptor() ([]byte, []int) {
	return file_keeper_proto_rawDescGZIP(), []int{12}
}

func (x *ListFielsdKeepResponse) GetData() map[string]*FieldKeep {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ListFielsdKeepResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_keeper_proto protoreflect.FileDescriptor

var file_keeper_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x3e, 0x0a,
	0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x40, 0x0a,
	0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22,
	0x3b, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0xf9, 0x01, 0x0a,
	0x09, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4b, 0x65, 0x65, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x72, 0x64, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x72, 0x64, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61, 0x72, 0x64, 0x43, 0x56, 0x43, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x61, 0x72, 0x64, 0x43, 0x56, 0x43, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x61, 0x72, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x61, 0x72, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x61,
	0x72, 0x64, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x61, 0x72, 0x64, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x22, 0x40, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4b, 0x65, 0x65, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x29, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4b, 0x65, 0x65, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x6b, 0x0a, 0x14, 0x41, 0x64,
	0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4b, 0x65, 0x65, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65,
	0x72, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4b, 0x65, 0x65, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x55, 0x0a, 0x14, 0x45, 0x64, 0x69, 0x74, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4b, 0x65, 0x65, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4b, 0x65, 0x65, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x58,
	0x0a, 0x15, 0x45, 0x64, 0x69, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4b, 0x65, 0x65, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70,
	0x65, 0x72, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4b, 0x65, 0x65, 0x70, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x2c, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4b, 0x65, 0x65, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x43, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x4b, 0x65, 0x65, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x17, 0x0a, 0x15, 0x4c,
	0x69, 0x73, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x4b, 0x65, 0x65, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0xc0, 0x01, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x65,
	0x6c, 0x73, 0x64, 0x4b, 0x65, 0x65, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x40, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e,
	0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46,
	0x69, 0x65, 0x6c, 0x73, 0x64, 0x4b, 0x65, 0x65, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0x4e, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70,
	0x65, 0x72, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4b, 0x65, 0x65, 0x70, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xac, 0x05, 0x0a, 0x0a, 0x47, 0x6f, 0x70, 0x68,
	0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x12, 0x5f, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x1b, 0x2e, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x12, 0x3a, 0x01, 0x2a, 0x22, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x53, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x12, 0x18, 0x2e, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x67, 0x6f, 0x70,
	0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x3a, 0x01, 0x2a,
	0x22, 0x0a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x73, 0x0a, 0x08,
	0x41, 0x64, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x70, 0x68, 0x6b,
	0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4b, 0x65,
	0x65, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x67, 0x6f, 0x70, 0x68,
	0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4b,
	0x65, 0x65, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x92, 0x41, 0x0e,
	0x62, 0x0c, 0x0a, 0x0a, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0d, 0x3a, 0x01, 0x2a, 0x22, 0x08, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x64,
	0x64, 0x12, 0x77, 0x0a, 0x09, 0x45, 0x64, 0x69, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x20,
	0x2e, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x45, 0x64, 0x69, 0x74,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x4b, 0x65, 0x65, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x45, 0x64,
	0x69, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4b, 0x65, 0x65, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x25, 0x92, 0x41, 0x0e, 0x62, 0x0c, 0x0a, 0x0a, 0x0a, 0x06, 0x62, 0x65,
	0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a, 0x01, 0x2a, 0x1a,
	0x09, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x64, 0x69, 0x74, 0x12, 0x80, 0x01, 0x0a, 0x08, 0x44,
	0x65, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x22, 0x2e, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65,
	0x65, 0x70, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4b, 0x65, 0x65, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x67, 0x6f,
	0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4b, 0x65, 0x65, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x2b, 0x92, 0x41, 0x0e, 0x62, 0x0c, 0x0a, 0x0a, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65,
	0x72, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x2a, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2f, 0x7b, 0x75, 0x75, 0x69, 0x64, 0x7d, 0x12, 0x77, 0x0a,
	0x0a, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x21, 0x2e, 0x67, 0x6f,
	0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x4b, 0x65, 0x65, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22,
	0x2e, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x46, 0x69, 0x65, 0x6c, 0x73, 0x64, 0x4b, 0x65, 0x65, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x22, 0x92, 0x41, 0x0e, 0x62, 0x0c, 0x0a, 0x0a, 0x0a, 0x06, 0x62, 0x65, 0x61,
	0x72, 0x65, 0x72, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x42, 0xb3, 0x01, 0x92, 0x41, 0x99, 0x01, 0x12, 0x46, 0x0a,
	0x0a, 0x47, 0x6f, 0x70, 0x68, 0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x22, 0x33, 0x0a, 0x0a, 0x47,
	0x6f, 0x70, 0x68, 0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x12, 0x25, 0x68, 0x74, 0x74, 0x70, 0x73,
	0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x72,
	0x72, 0x68, 0x61, 0x6f, 0x73, 0x2f, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72,
	0x32, 0x03, 0x31, 0x2e, 0x30, 0x5a, 0x4f, 0x0a, 0x4d, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65,
	0x72, 0x12, 0x43, 0x08, 0x02, 0x12, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x20, 0x42, 0x65, 0x61, 0x72,
	0x65, 0x72, 0x20, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x61,
	0x20, 0x73, 0x70, 0x61, 0x63, 0x65, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x4a, 0x57, 0x54, 0x20, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x5a, 0x14, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70,
	0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_keeper_proto_rawDescOnce sync.Once
	file_keeper_proto_rawDescData = file_keeper_proto_rawDesc
)

func file_keeper_proto_rawDescGZIP() []byte {
	file_keeper_proto_rawDescOnce.Do(func() {
		file_keeper_proto_rawDescData = protoimpl.X.CompressGZIP(file_keeper_proto_rawDescData)
	})
	return file_keeper_proto_rawDescData
}

var file_keeper_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_keeper_proto_goTypes = []any{
	(*RegisterRequest)(nil),         // 0: gophkeeper.RegisterRequest
	(*RegisterResponse)(nil),        // 1: gophkeeper.RegisterResponse
	(*LoginRequest)(nil),            // 2: gophkeeper.LoginRequest
	(*LoginResponse)(nil),           // 3: gophkeeper.LoginResponse
	(*FieldKeep)(nil),               // 4: gophkeeper.FieldKeep
	(*AddFieldKeepRequest)(nil),     // 5: gophkeeper.AddFieldKeepRequest
	(*AddFieldKeepResponse)(nil),    // 6: gophkeeper.AddFieldKeepResponse
	(*EditFieldKeepRequest)(nil),    // 7: gophkeeper.EditFieldKeepRequest
	(*EditFieldKeepResponse)(nil),   // 8: gophkeeper.EditFieldKeepResponse
	(*DeleteFieldKeepRequest)(nil),  // 9: gophkeeper.DeleteFieldKeepRequest
	(*DeleteFieldKeepResponse)(nil), // 10: gophkeeper.DeleteFieldKeepResponse
	(*ListFieldsKeepRequest)(nil),   // 11: gophkeeper.ListFieldsKeepRequest
	(*ListFielsdKeepResponse)(nil),  // 12: gophkeeper.ListFielsdKeepResponse
	nil,                             // 13: gophkeeper.ListFielsdKeepResponse.DataEntry
}
var file_keeper_proto_depIdxs = []int32{
	4,  // 0: gophkeeper.AddFieldKeepRequest.data:type_name -> gophkeeper.FieldKeep
	4,  // 1: gophkeeper.AddFieldKeepResponse.data:type_name -> gophkeeper.FieldKeep
	4,  // 2: gophkeeper.EditFieldKeepRequest.data:type_name -> gophkeeper.FieldKeep
	4,  // 3: gophkeeper.EditFieldKeepResponse.data:type_name -> gophkeeper.FieldKeep
	13, // 4: gophkeeper.ListFielsdKeepResponse.data:type_name -> gophkeeper.ListFielsdKeepResponse.DataEntry
	4,  // 5: gophkeeper.ListFielsdKeepResponse.DataEntry.value:type_name -> gophkeeper.FieldKeep
	0,  // 6: gophkeeper.GophKeeper.Register:input_type -> gophkeeper.RegisterRequest
	2,  // 7: gophkeeper.GophKeeper.Login:input_type -> gophkeeper.LoginRequest
	5,  // 8: gophkeeper.GophKeeper.AddField:input_type -> gophkeeper.AddFieldKeepRequest
	7,  // 9: gophkeeper.GophKeeper.EditField:input_type -> gophkeeper.EditFieldKeepRequest
	9,  // 10: gophkeeper.GophKeeper.DelField:input_type -> gophkeeper.DeleteFieldKeepRequest
	11, // 11: gophkeeper.GophKeeper.ListFields:input_type -> gophkeeper.ListFieldsKeepRequest
	1,  // 12: gophkeeper.GophKeeper.Register:output_type -> gophkeeper.RegisterResponse
	3,  // 13: gophkeeper.GophKeeper.Login:output_type -> gophkeeper.LoginResponse
	6,  // 14: gophkeeper.GophKeeper.AddField:output_type -> gophkeeper.AddFieldKeepResponse
	8,  // 15: gophkeeper.GophKeeper.EditField:output_type -> gophkeeper.EditFieldKeepResponse
	10, // 16: gophkeeper.GophKeeper.DelField:output_type -> gophkeeper.DeleteFieldKeepResponse
	12, // 17: gophkeeper.GophKeeper.ListFields:output_type -> gophkeeper.ListFielsdKeepResponse
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_keeper_proto_init() }
func file_keeper_proto_init() {
	if File_keeper_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_keeper_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*RegisterRequest); i {
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
		file_keeper_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*RegisterResponse); i {
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
		file_keeper_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*LoginRequest); i {
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
		file_keeper_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*LoginResponse); i {
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
		file_keeper_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*FieldKeep); i {
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
		file_keeper_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*AddFieldKeepRequest); i {
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
		file_keeper_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*AddFieldKeepResponse); i {
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
		file_keeper_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*EditFieldKeepRequest); i {
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
		file_keeper_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*EditFieldKeepResponse); i {
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
		file_keeper_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteFieldKeepRequest); i {
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
		file_keeper_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteFieldKeepResponse); i {
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
		file_keeper_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*ListFieldsKeepRequest); i {
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
		file_keeper_proto_msgTypes[12].Exporter = func(v any, i int) any {
			switch v := v.(*ListFielsdKeepResponse); i {
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
			RawDescriptor: file_keeper_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_keeper_proto_goTypes,
		DependencyIndexes: file_keeper_proto_depIdxs,
		MessageInfos:      file_keeper_proto_msgTypes,
	}.Build()
	File_keeper_proto = out.File
	file_keeper_proto_rawDesc = nil
	file_keeper_proto_goTypes = nil
	file_keeper_proto_depIdxs = nil
}
