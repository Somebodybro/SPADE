// [START declaration]

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: spade.proto

package __

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Define empty message for functions returning nil
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spade_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_spade_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_spade_proto_rawDescGZIP(), []int{0}
}

type PublicParamsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PublicParamsReq) Reset() {
	*x = PublicParamsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spade_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicParamsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicParamsReq) ProtoMessage() {}

func (x *PublicParamsReq) ProtoReflect() protoreflect.Message {
	mi := &file_spade_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicParamsReq.ProtoReflect.Descriptor instead.
func (*PublicParamsReq) Descriptor() ([]byte, []int) {
	return file_spade_proto_rawDescGZIP(), []int{1}
}

type PublicParamsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Q   []byte   `protobuf:"bytes,1,opt,name=q,proto3" json:"q,omitempty"`     // Store q as bytes
	G   []byte   `protobuf:"bytes,2,opt,name=g,proto3" json:"g,omitempty"`     // Store g as bytes
	Mpk [][]byte `protobuf:"bytes,3,rep,name=mpk,proto3" json:"mpk,omitempty"` // Store mpk as bytes
}

func (x *PublicParamsResp) Reset() {
	*x = PublicParamsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spade_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicParamsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicParamsResp) ProtoMessage() {}

func (x *PublicParamsResp) ProtoReflect() protoreflect.Message {
	mi := &file_spade_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicParamsResp.ProtoReflect.Descriptor instead.
func (*PublicParamsResp) Descriptor() ([]byte, []int) {
	return file_spade_proto_rawDescGZIP(), []int{2}
}

func (x *PublicParamsResp) GetQ() []byte {
	if x != nil {
		return x.Q
	}
	return nil
}

func (x *PublicParamsResp) GetG() []byte {
	if x != nil {
		return x.G
	}
	return nil
}

func (x *PublicParamsResp) GetMpk() [][]byte {
	if x != nil {
		return x.Mpk
	}
	return nil
}

type UserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                // user's identifier
	RegKey     []byte   `protobuf:"bytes,2,opt,name=regKey,proto3" json:"regKey,omitempty"`         // user's registration key
	Ciphertext [][]byte `protobuf:"bytes,3,rep,name=ciphertext,proto3" json:"ciphertext,omitempty"` // user's encrypted data
}

func (x *UserReq) Reset() {
	*x = UserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spade_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserReq) ProtoMessage() {}

func (x *UserReq) ProtoReflect() protoreflect.Message {
	mi := &file_spade_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserReq.ProtoReflect.Descriptor instead.
func (*UserReq) Descriptor() ([]byte, []int) {
	return file_spade_proto_rawDescGZIP(), []int{3}
}

func (x *UserReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserReq) GetRegKey() []byte {
	if x != nil {
		return x.RegKey
	}
	return nil
}

func (x *UserReq) GetCiphertext() [][]byte {
	if x != nil {
		return x.Ciphertext
	}
	return nil
}

type UserResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flag bool `protobuf:"varint,1,opt,name=flag,proto3" json:"flag,omitempty"` // OK or Not
}

func (x *UserResp) Reset() {
	*x = UserResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spade_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResp) ProtoMessage() {}

func (x *UserResp) ProtoReflect() protoreflect.Message {
	mi := &file_spade_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResp.ProtoReflect.Descriptor instead.
func (*UserResp) Descriptor() ([]byte, []int) {
	return file_spade_proto_rawDescGZIP(), []int{4}
}

func (x *UserResp) GetFlag() bool {
	if x != nil {
		return x.Flag
	}
	return false
}

type AnalystReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`       // the corresponding user id for accessing his/her ciphertext
	Value int64 `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"` // the value for search query, which Analyst is looking for
}

func (x *AnalystReq) Reset() {
	*x = AnalystReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spade_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalystReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalystReq) ProtoMessage() {}

func (x *AnalystReq) ProtoReflect() protoreflect.Message {
	mi := &file_spade_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalystReq.ProtoReflect.Descriptor instead.
func (*AnalystReq) Descriptor() ([]byte, []int) {
	return file_spade_proto_rawDescGZIP(), []int{5}
}

func (x *AnalystReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AnalystReq) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type AnalystResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dkv        [][]byte `protobuf:"bytes,1,rep,name=dkv,proto3" json:"dkv,omitempty"`               // corresponding decryption key for query value
	Ciphertext [][]byte `protobuf:"bytes,2,rep,name=ciphertext,proto3" json:"ciphertext,omitempty"` // corresponding user's ciphertext
}

func (x *AnalystResp) Reset() {
	*x = AnalystResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spade_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalystResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalystResp) ProtoMessage() {}

func (x *AnalystResp) ProtoReflect() protoreflect.Message {
	mi := &file_spade_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalystResp.ProtoReflect.Descriptor instead.
func (*AnalystResp) Descriptor() ([]byte, []int) {
	return file_spade_proto_rawDescGZIP(), []int{6}
}

func (x *AnalystResp) GetDkv() [][]byte {
	if x != nil {
		return x.Dkv
	}
	return nil
}

func (x *AnalystResp) GetCiphertext() [][]byte {
	if x != nil {
		return x.Ciphertext
	}
	return nil
}

var File_spade_proto protoreflect.FileDescriptor

var file_spade_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x70, 0x61, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73,
	0x70, 0x61, 0x64, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x11, 0x0a, 0x0f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x52, 0x65, 0x71, 0x22, 0x40, 0x0a, 0x10, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0c, 0x0a, 0x01, 0x71, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x01, 0x71, 0x12, 0x0c, 0x0a, 0x01, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x01, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x70, 0x6b, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0c, 0x52, 0x03, 0x6d, 0x70, 0x6b, 0x22, 0x51, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x72, 0x65, 0x67, 0x4b, 0x65, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x69,
	0x70, 0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0a,
	0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x22, 0x1e, 0x0a, 0x08, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x22, 0x32, 0x0a, 0x0a, 0x41, 0x6e,
	0x61, 0x6c, 0x79, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x3f,
	0x0a, 0x0b, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a,
	0x03, 0x64, 0x6b, 0x76, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x03, 0x64, 0x6b, 0x76, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0c, 0x52, 0x0a, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x32,
	0xd1, 0x01, 0x0a, 0x07, 0x43, 0x75, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x4e, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1b,
	0x2e, 0x73, 0x70, 0x61, 0x64, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x73, 0x70,
	0x61, 0x64, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0b, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x13, 0x2e, 0x73, 0x70, 0x61,
	0x64, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a,
	0x14, 0x2e, 0x73, 0x70, 0x61, 0x64, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x16, 0x2e, 0x73, 0x70, 0x61, 0x64, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x6e,
	0x61, 0x6c, 0x79, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x73, 0x70, 0x61, 0x64, 0x65,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x00, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_spade_proto_rawDescOnce sync.Once
	file_spade_proto_rawDescData = file_spade_proto_rawDesc
)

func file_spade_proto_rawDescGZIP() []byte {
	file_spade_proto_rawDescOnce.Do(func() {
		file_spade_proto_rawDescData = protoimpl.X.CompressGZIP(file_spade_proto_rawDescData)
	})
	return file_spade_proto_rawDescData
}

var file_spade_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_spade_proto_goTypes = []interface{}{
	(*Empty)(nil),            // 0: spadeproto.Empty
	(*PublicParamsReq)(nil),  // 1: spadeproto.PublicParamsReq
	(*PublicParamsResp)(nil), // 2: spadeproto.PublicParamsResp
	(*UserReq)(nil),          // 3: spadeproto.UserReq
	(*UserResp)(nil),         // 4: spadeproto.UserResp
	(*AnalystReq)(nil),       // 5: spadeproto.AnalystReq
	(*AnalystResp)(nil),      // 6: spadeproto.AnalystResp
}
var file_spade_proto_depIdxs = []int32{
	1, // 0: spadeproto.Curator.GetPublicParams:input_type -> spadeproto.PublicParamsReq
	3, // 1: spadeproto.Curator.UserRequest:input_type -> spadeproto.UserReq
	5, // 2: spadeproto.Curator.Query:input_type -> spadeproto.AnalystReq
	2, // 3: spadeproto.Curator.GetPublicParams:output_type -> spadeproto.PublicParamsResp
	4, // 4: spadeproto.Curator.UserRequest:output_type -> spadeproto.UserResp
	6, // 5: spadeproto.Curator.Query:output_type -> spadeproto.AnalystResp
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_spade_proto_init() }
func file_spade_proto_init() {
	if File_spade_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spade_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_spade_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicParamsReq); i {
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
		file_spade_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicParamsResp); i {
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
		file_spade_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserReq); i {
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
		file_spade_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserResp); i {
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
		file_spade_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalystReq); i {
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
		file_spade_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalystResp); i {
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
			RawDescriptor: file_spade_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spade_proto_goTypes,
		DependencyIndexes: file_spade_proto_depIdxs,
		MessageInfos:      file_spade_proto_msgTypes,
	}.Build()
	File_spade_proto = out.File
	file_spade_proto_rawDesc = nil
	file_spade_proto_goTypes = nil
	file_spade_proto_depIdxs = nil
}
