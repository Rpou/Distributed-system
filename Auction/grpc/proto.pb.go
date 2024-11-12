// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.12.4
// source: grpc/proto.proto

package proto

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

type Bid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bid  int64 `protobuf:"varint,1,opt,name=Bid,proto3" json:"Bid,omitempty"`
	Time int64 `protobuf:"varint,2,opt,name=Time,proto3" json:"Time,omitempty"`
}

func (x *Bid) Reset() {
	*x = Bid{}
	mi := &file_grpc_proto_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Bid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bid) ProtoMessage() {}

func (x *Bid) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bid.ProtoReflect.Descriptor instead.
func (*Bid) Descriptor() ([]byte, []int) {
	return file_grpc_proto_proto_rawDescGZIP(), []int{0}
}

func (x *Bid) GetBid() int64 {
	if x != nil {
		return x.Bid
	}
	return 0
}

func (x *Bid) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

type ClientToNodeBid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bid        int64 `protobuf:"varint,1,opt,name=Bid,proto3" json:"Bid,omitempty"`
	NodeNumber int64 `protobuf:"varint,2,opt,name=NodeNumber,proto3" json:"NodeNumber,omitempty"`
}

func (x *ClientToNodeBid) Reset() {
	*x = ClientToNodeBid{}
	mi := &file_grpc_proto_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClientToNodeBid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientToNodeBid) ProtoMessage() {}

func (x *ClientToNodeBid) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientToNodeBid.ProtoReflect.Descriptor instead.
func (*ClientToNodeBid) Descriptor() ([]byte, []int) {
	return file_grpc_proto_proto_rawDescGZIP(), []int{1}
}

func (x *ClientToNodeBid) GetBid() int64 {
	if x != nil {
		return x.Bid
	}
	return 0
}

func (x *ClientToNodeBid) GetNodeNumber() int64 {
	if x != nil {
		return x.NodeNumber
	}
	return 0
}

type AcceptNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Giveacces bool   `protobuf:"varint,1,opt,name=giveacces,proto3" json:"giveacces,omitempty"`
	Status    string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *AcceptNodeRequest) Reset() {
	*x = AcceptNodeRequest{}
	mi := &file_grpc_proto_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AcceptNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptNodeRequest) ProtoMessage() {}

func (x *AcceptNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptNodeRequest.ProtoReflect.Descriptor instead.
func (*AcceptNodeRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_proto_rawDescGZIP(), []int{2}
}

func (x *AcceptNodeRequest) GetGiveacces() bool {
	if x != nil {
		return x.Giveacces
	}
	return false
}

func (x *AcceptNodeRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type AcceptClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuctionBid int64 `protobuf:"varint,1,opt,name=AuctionBid,proto3" json:"AuctionBid,omitempty"`
}

func (x *AcceptClientRequest) Reset() {
	*x = AcceptClientRequest{}
	mi := &file_grpc_proto_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AcceptClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptClientRequest) ProtoMessage() {}

func (x *AcceptClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptClientRequest.ProtoReflect.Descriptor instead.
func (*AcceptClientRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_proto_rawDescGZIP(), []int{3}
}

func (x *AcceptClientRequest) GetAuctionBid() int64 {
	if x != nil {
		return x.AuctionBid
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_grpc_proto_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_proto_msgTypes[4]
	if x != nil {
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
	return file_grpc_proto_proto_rawDescGZIP(), []int{4}
}

var File_grpc_proto_proto protoreflect.FileDescriptor

var file_grpc_proto_proto_rawDesc = []byte{
	0x0a, 0x10, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x03, 0x42, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x42, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x42, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x22,
	0x43, 0x0a, 0x0f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x4e, 0x6f, 0x64, 0x65, 0x42,
	0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x42, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x42, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x4e, 0x6f, 0x64, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x4e, 0x6f, 0x64, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x22, 0x49, 0x0a, 0x11, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x4e, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x69, 0x76,
	0x65, 0x61, 0x63, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x67, 0x69,
	0x76, 0x65, 0x61, 0x63, 0x63, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x35, 0x0a, 0x13, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x41, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x69, 0x64, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32,
	0x6e, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x25, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x04, 0x2e, 0x42, 0x69, 0x64,
	0x1a, 0x12, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x54, 0x6f, 0x4e, 0x6f, 0x64, 0x65, 0x42, 0x69, 0x64, 0x1a, 0x14, 0x2e, 0x41, 0x63, 0x63, 0x65,
	0x70, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42,
	0x16, 0x5a, 0x14, 0x49, 0x54, 0x55, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_proto_proto_rawDescOnce sync.Once
	file_grpc_proto_proto_rawDescData = file_grpc_proto_proto_rawDesc
)

func file_grpc_proto_proto_rawDescGZIP() []byte {
	file_grpc_proto_proto_rawDescOnce.Do(func() {
		file_grpc_proto_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_proto_proto_rawDescData)
	})
	return file_grpc_proto_proto_rawDescData
}

var file_grpc_proto_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_grpc_proto_proto_goTypes = []any{
	(*Bid)(nil),                 // 0: Bid
	(*ClientToNodeBid)(nil),     // 1: ClientToNodeBid
	(*AcceptNodeRequest)(nil),   // 2: AcceptNodeRequest
	(*AcceptClientRequest)(nil), // 3: AcceptClientRequest
	(*Empty)(nil),               // 4: Empty
}
var file_grpc_proto_proto_depIdxs = []int32{
	0, // 0: Communcation.Request:input_type -> Bid
	1, // 1: Communcation.ClientRequest:input_type -> ClientToNodeBid
	2, // 2: Communcation.Request:output_type -> AcceptNodeRequest
	3, // 3: Communcation.ClientRequest:output_type -> AcceptClientRequest
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_proto_proto_init() }
func file_grpc_proto_proto_init() {
	if File_grpc_proto_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpc_proto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_proto_proto_goTypes,
		DependencyIndexes: file_grpc_proto_proto_depIdxs,
		MessageInfos:      file_grpc_proto_proto_msgTypes,
	}.Build()
	File_grpc_proto_proto = out.File
	file_grpc_proto_proto_rawDesc = nil
	file_grpc_proto_proto_goTypes = nil
	file_grpc_proto_proto_depIdxs = nil
}
