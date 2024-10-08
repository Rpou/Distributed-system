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

type Posts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posts       []string `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
	LamportTime int64    `protobuf:"varint,2,opt,name=lamport_time,json=lamportTime,proto3" json:"lamport_time,omitempty"`
}

func (x *Posts) Reset() {
	*x = Posts{}
	mi := &file_grpc_proto_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Posts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Posts) ProtoMessage() {}

func (x *Posts) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Posts.ProtoReflect.Descriptor instead.
func (*Posts) Descriptor() ([]byte, []int) {
	return file_grpc_proto_proto_rawDescGZIP(), []int{0}
}

func (x *Posts) GetPosts() []string {
	if x != nil {
		return x.Posts
	}
	return nil
}

func (x *Posts) GetLamportTime() int64 {
	if x != nil {
		return x.LamportTime
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
	mi := &file_grpc_proto_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_grpc_proto_proto_rawDescGZIP(), []int{1}
}

type ClientNumber struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cn int64 `protobuf:"varint,1,opt,name=cn,proto3" json:"cn,omitempty"`
}

func (x *ClientNumber) Reset() {
	*x = ClientNumber{}
	mi := &file_grpc_proto_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClientNumber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientNumber) ProtoMessage() {}

func (x *ClientNumber) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ClientNumber.ProtoReflect.Descriptor instead.
func (*ClientNumber) Descriptor() ([]byte, []int) {
	return file_grpc_proto_proto_rawDescGZIP(), []int{2}
}

func (x *ClientNumber) GetCn() int64 {
	if x != nil {
		return x.Cn
	}
	return 0
}

type Connected struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Con bool `protobuf:"varint,1,opt,name=con,proto3" json:"con,omitempty"`
}

func (x *Connected) Reset() {
	*x = Connected{}
	mi := &file_grpc_proto_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Connected) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Connected) ProtoMessage() {}

func (x *Connected) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Connected.ProtoReflect.Descriptor instead.
func (*Connected) Descriptor() ([]byte, []int) {
	return file_grpc_proto_proto_rawDescGZIP(), []int{3}
}

func (x *Connected) GetCon() bool {
	if x != nil {
		return x.Con
	}
	return false
}

type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post        string `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	LamportTime int64  `protobuf:"varint,2,opt,name=lamport_time,json=lamportTime,proto3" json:"lamport_time,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	mi := &file_grpc_proto_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_grpc_proto_proto_rawDescGZIP(), []int{4}
}

func (x *Post) GetPost() string {
	if x != nil {
		return x.Post
	}
	return ""
}

func (x *Post) GetLamportTime() int64 {
	if x != nil {
		return x.LamportTime
	}
	return 0
}

type Posted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posted bool `protobuf:"varint,1,opt,name=posted,proto3" json:"posted,omitempty"`
}

func (x *Posted) Reset() {
	*x = Posted{}
	mi := &file_grpc_proto_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Posted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Posted) ProtoMessage() {}

func (x *Posted) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Posted.ProtoReflect.Descriptor instead.
func (*Posted) Descriptor() ([]byte, []int) {
	return file_grpc_proto_proto_rawDescGZIP(), []int{5}
}

func (x *Posted) GetPosted() bool {
	if x != nil {
		return x.Posted
	}
	return false
}

var File_grpc_proto_proto protoreflect.FileDescriptor

var file_grpc_proto_proto_rawDesc = []byte{
	0x0a, 0x10, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x40, 0x0a, 0x05, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x70, 0x6f, 0x73, 0x74,
	0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1e, 0x0a,
	0x0c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x63, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x63, 0x6e, 0x22, 0x1d, 0x0a,
	0x09, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x63, 0x6f, 0x6e, 0x22, 0x3d, 0x0a, 0x04,
	0x50, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x61, 0x6d, 0x70,
	0x6f, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x20, 0x0a, 0x06, 0x50,
	0x6f, 0x73, 0x74, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x64, 0x32, 0xa0, 0x01,
	0x0a, 0x0c, 0x43, 0x68, 0x69, 0x74, 0x74, 0x79, 0x63, 0x68, 0x61, 0x74, 0x44, 0x42, 0x12, 0x1c,
	0x0a, 0x08, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x06, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x22, 0x00, 0x12, 0x1f, 0x0a, 0x0b,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x05, 0x2e, 0x50, 0x6f,
	0x73, 0x74, 0x1a, 0x07, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x64, 0x22, 0x00, 0x12, 0x26, 0x0a,
	0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x0d, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x1a, 0x0a, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x0a, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x12, 0x0d, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x1a, 0x0a, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x22, 0x00,
	0x42, 0x17, 0x5a, 0x15, 0x43, 0x68, 0x69, 0x74, 0x74, 0x79, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
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

var file_grpc_proto_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_grpc_proto_proto_goTypes = []any{
	(*Posts)(nil),        // 0: Posts
	(*Empty)(nil),        // 1: Empty
	(*ClientNumber)(nil), // 2: ClientNumber
	(*Connected)(nil),    // 3: Connected
	(*Post)(nil),         // 4: Post
	(*Posted)(nil),       // 5: Posted
}
var file_grpc_proto_proto_depIdxs = []int32{
	1, // 0: ChittychatDB.GetPosts:input_type -> Empty
	4, // 1: ChittychatDB.PublishPost:input_type -> Post
	2, // 2: ChittychatDB.Connect:input_type -> ClientNumber
	2, // 3: ChittychatDB.Disconnect:input_type -> ClientNumber
	0, // 4: ChittychatDB.GetPosts:output_type -> Posts
	5, // 5: ChittychatDB.PublishPost:output_type -> Posted
	3, // 6: ChittychatDB.Connect:output_type -> Connected
	3, // 7: ChittychatDB.Disconnect:output_type -> Connected
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
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
			NumMessages:   6,
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
