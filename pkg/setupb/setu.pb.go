// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: pkg/setupb/setu.proto

package setupb

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

type SetuReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SetuReq) Reset() {
	*x = SetuReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_setupb_setu_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetuReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetuReq) ProtoMessage() {}

func (x *SetuReq) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_setupb_setu_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetuReq.ProtoReflect.Descriptor instead.
func (*SetuReq) Descriptor() ([]byte, []int) {
	return file_pkg_setupb_setu_proto_rawDescGZIP(), []int{0}
}

func (x *SetuReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SetuResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chunk []byte `protobuf:"bytes,2,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (x *SetuResp) Reset() {
	*x = SetuResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_setupb_setu_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetuResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetuResp) ProtoMessage() {}

func (x *SetuResp) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_setupb_setu_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetuResp.ProtoReflect.Descriptor instead.
func (*SetuResp) Descriptor() ([]byte, []int) {
	return file_pkg_setupb_setu_proto_rawDescGZIP(), []int{1}
}

func (x *SetuResp) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

type FetchReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount uint64 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *FetchReq) Reset() {
	*x = FetchReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_setupb_setu_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchReq) ProtoMessage() {}

func (x *FetchReq) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_setupb_setu_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchReq.ProtoReflect.Descriptor instead.
func (*FetchReq) Descriptor() ([]byte, []int) {
	return file_pkg_setupb_setu_proto_rawDescGZIP(), []int{2}
}

func (x *FetchReq) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type InventoryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page      uint64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageLimit uint64 `protobuf:"varint,2,opt,name=pageLimit,proto3" json:"pageLimit,omitempty"`
}

func (x *InventoryReq) Reset() {
	*x = InventoryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_setupb_setu_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InventoryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InventoryReq) ProtoMessage() {}

func (x *InventoryReq) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_setupb_setu_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InventoryReq.ProtoReflect.Descriptor instead.
func (*InventoryReq) Descriptor() ([]byte, []int) {
	return file_pkg_setupb_setu_proto_rawDescGZIP(), []int{3}
}

func (x *InventoryReq) GetPage() uint64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *InventoryReq) GetPageLimit() uint64 {
	if x != nil {
		return x.PageLimit
	}
	return 0
}

type FetchResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrMsg string `protobuf:"bytes,1,opt,name=errMsg,proto3" json:"errMsg,omitempty"`
}

func (x *FetchResp) Reset() {
	*x = FetchResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_setupb_setu_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchResp) ProtoMessage() {}

func (x *FetchResp) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_setupb_setu_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchResp.ProtoReflect.Descriptor instead.
func (*FetchResp) Descriptor() ([]byte, []int) {
	return file_pkg_setupb_setu_proto_rawDescGZIP(), []int{4}
}

func (x *FetchResp) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

type InventoryResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Err  *FetchResp                `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
	Info []*InventoryResp_SetuInfo `protobuf:"bytes,2,rep,name=info,proto3" json:"info,omitempty"`
}

func (x *InventoryResp) Reset() {
	*x = InventoryResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_setupb_setu_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InventoryResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InventoryResp) ProtoMessage() {}

func (x *InventoryResp) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_setupb_setu_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InventoryResp.ProtoReflect.Descriptor instead.
func (*InventoryResp) Descriptor() ([]byte, []int) {
	return file_pkg_setupb_setu_proto_rawDescGZIP(), []int{5}
}

func (x *InventoryResp) GetErr() *FetchResp {
	if x != nil {
		return x.Err
	}
	return nil
}

func (x *InventoryResp) GetInfo() []*InventoryResp_SetuInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type InventoryResp_SetuInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Uid   int64  `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`
	Url   string `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	IsR18 bool   `protobuf:"varint,5,opt,name=is_r18,json=isR18,proto3" json:"is_r18,omitempty"`
}

func (x *InventoryResp_SetuInfo) Reset() {
	*x = InventoryResp_SetuInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_setupb_setu_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InventoryResp_SetuInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InventoryResp_SetuInfo) ProtoMessage() {}

func (x *InventoryResp_SetuInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_setupb_setu_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InventoryResp_SetuInfo.ProtoReflect.Descriptor instead.
func (*InventoryResp_SetuInfo) Descriptor() ([]byte, []int) {
	return file_pkg_setupb_setu_proto_rawDescGZIP(), []int{5, 0}
}

func (x *InventoryResp_SetuInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *InventoryResp_SetuInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *InventoryResp_SetuInfo) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *InventoryResp_SetuInfo) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *InventoryResp_SetuInfo) GetIsR18() bool {
	if x != nil {
		return x.IsR18
	}
	return false
}

var File_pkg_setupb_setu_proto protoreflect.FileDescriptor

var file_pkg_setupb_setu_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x74, 0x75, 0x70, 0x62, 0x2f, 0x73, 0x65, 0x74,
	0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x73, 0x65, 0x74, 0x75, 0x22, 0x19, 0x0a,
	0x07, 0x53, 0x65, 0x74, 0x75, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x20, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x75,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x22, 0x22, 0x0a, 0x08, 0x46, 0x65,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x40,
	0x0a, 0x0c, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x22, 0x23, 0x0a, 0x09, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a,
	0x06, 0x65, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65,
	0x72, 0x72, 0x4d, 0x73, 0x67, 0x22, 0xd1, 0x01, 0x0a, 0x0d, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x12, 0x21, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x65, 0x74, 0x75, 0x2e, 0x46, 0x65, 0x74, 0x63,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x52, 0x03, 0x65, 0x72, 0x72, 0x12, 0x30, 0x0a, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x65, 0x74, 0x75, 0x2e,
	0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x2e, 0x53, 0x65,
	0x74, 0x75, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x1a, 0x6b, 0x0a, 0x08,
	0x53, 0x65, 0x74, 0x75, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x73, 0x5f, 0x72, 0x31, 0x38, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x52, 0x31, 0x38, 0x32, 0xa6, 0x01, 0x0a, 0x0b, 0x53, 0x65,
	0x74, 0x75, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x2e, 0x73, 0x65, 0x74, 0x75,
	0x2e, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e,
	0x73, 0x65, 0x74, 0x75, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x05, 0x46, 0x65, 0x74, 0x63, 0x68, 0x12, 0x0e, 0x2e,
	0x73, 0x65, 0x74, 0x75, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e,
	0x73, 0x65, 0x74, 0x75, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00,
	0x12, 0x30, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x65, 0x74, 0x75, 0x42, 0x79, 0x49, 0x64, 0x12,
	0x0d, 0x2e, 0x73, 0x65, 0x74, 0x75, 0x2e, 0x53, 0x65, 0x74, 0x75, 0x52, 0x65, 0x71, 0x1a, 0x0e,
	0x2e, 0x73, 0x65, 0x74, 0x75, 0x2e, 0x53, 0x65, 0x74, 0x75, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00,
	0x30, 0x01, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x30, 0x77, 0x30, 0x6d, 0x65, 0x77, 0x6f, 0x2f, 0x62, 0x75, 0x64, 0x6f, 0x6e, 0x67, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x74, 0x75, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pkg_setupb_setu_proto_rawDescOnce sync.Once
	file_pkg_setupb_setu_proto_rawDescData = file_pkg_setupb_setu_proto_rawDesc
)

func file_pkg_setupb_setu_proto_rawDescGZIP() []byte {
	file_pkg_setupb_setu_proto_rawDescOnce.Do(func() {
		file_pkg_setupb_setu_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_setupb_setu_proto_rawDescData)
	})
	return file_pkg_setupb_setu_proto_rawDescData
}

var file_pkg_setupb_setu_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pkg_setupb_setu_proto_goTypes = []interface{}{
	(*SetuReq)(nil),                // 0: setu.SetuReq
	(*SetuResp)(nil),               // 1: setu.SetuResp
	(*FetchReq)(nil),               // 2: setu.FetchReq
	(*InventoryReq)(nil),           // 3: setu.InventoryReq
	(*FetchResp)(nil),              // 4: setu.FetchResp
	(*InventoryResp)(nil),          // 5: setu.InventoryResp
	(*InventoryResp_SetuInfo)(nil), // 6: setu.InventoryResp.SetuInfo
}
var file_pkg_setupb_setu_proto_depIdxs = []int32{
	4, // 0: setu.InventoryResp.err:type_name -> setu.FetchResp
	6, // 1: setu.InventoryResp.info:type_name -> setu.InventoryResp.SetuInfo
	3, // 2: setu.SetuService.GetInventory:input_type -> setu.InventoryReq
	2, // 3: setu.SetuService.Fetch:input_type -> setu.FetchReq
	0, // 4: setu.SetuService.GetSetuById:input_type -> setu.SetuReq
	5, // 5: setu.SetuService.GetInventory:output_type -> setu.InventoryResp
	4, // 6: setu.SetuService.Fetch:output_type -> setu.FetchResp
	1, // 7: setu.SetuService.GetSetuById:output_type -> setu.SetuResp
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_setupb_setu_proto_init() }
func file_pkg_setupb_setu_proto_init() {
	if File_pkg_setupb_setu_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_setupb_setu_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetuReq); i {
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
		file_pkg_setupb_setu_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetuResp); i {
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
		file_pkg_setupb_setu_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchReq); i {
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
		file_pkg_setupb_setu_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InventoryReq); i {
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
		file_pkg_setupb_setu_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchResp); i {
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
		file_pkg_setupb_setu_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InventoryResp); i {
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
		file_pkg_setupb_setu_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InventoryResp_SetuInfo); i {
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
			RawDescriptor: file_pkg_setupb_setu_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_setupb_setu_proto_goTypes,
		DependencyIndexes: file_pkg_setupb_setu_proto_depIdxs,
		MessageInfos:      file_pkg_setupb_setu_proto_msgTypes,
	}.Build()
	File_pkg_setupb_setu_proto = out.File
	file_pkg_setupb_setu_proto_rawDesc = nil
	file_pkg_setupb_setu_proto_goTypes = nil
	file_pkg_setupb_setu_proto_depIdxs = nil
}
