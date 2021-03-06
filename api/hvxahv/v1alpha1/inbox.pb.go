// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1alpha1.25.0-devel
// 	protoc        v3.14.0
// source: api/hvxahv/v1alpha1/inbox.proto

package pb

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

type InboxData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	RequestId string `protobuf:"bytes,2,opt,name=RequestId,proto3" json:"RequestId,omitempty"`
	EventType string `protobuf:"bytes,3,opt,name=EventType,proto3" json:"EventType,omitempty"`
	Actor     string `protobuf:"bytes,4,opt,name=Actor,proto3" json:"Actor,omitempty"`
}

func (x *InboxData) Reset() {
	*x = InboxData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hvxahv_v1_inbox_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InboxData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InboxData) ProtoMessage() {}

func (x *InboxData) ProtoReflect() protoreflect.Message {
	mi := &file_api_hvxahv_v1_inbox_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InboxData.ProtoReflect.Descriptor instead.
func (*InboxData) Descriptor() ([]byte, []int) {
	return file_api_hvxahv_v1_inbox_proto_rawDescGZIP(), []int{0}
}

func (x *InboxData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InboxData) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *InboxData) GetEventType() string {
	if x != nil {
		return x.EventType
	}
	return ""
}

func (x *InboxData) GetActor() string {
	if x != nil {
		return x.Actor
	}
	return ""
}

type NewInboxReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reply string `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
}

func (x *NewInboxReply) Reset() {
	*x = NewInboxReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hvxahv_v1_inbox_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewInboxReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewInboxReply) ProtoMessage() {}

func (x *NewInboxReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_hvxahv_v1_inbox_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewInboxReply.ProtoReflect.Descriptor instead.
func (*NewInboxReply) Descriptor() ([]byte, []int) {
	return file_api_hvxahv_v1_inbox_proto_rawDescGZIP(), []int{1}
}

func (x *NewInboxReply) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

type Name struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Name) Reset() {
	*x = Name{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hvxahv_v1_inbox_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Name) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Name) ProtoMessage() {}

func (x *Name) ProtoReflect() protoreflect.Message {
	mi := &file_api_hvxahv_v1_inbox_proto_msgTypes[2]
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
	return file_api_hvxahv_v1_inbox_proto_rawDescGZIP(), []int{2}
}

func (x *Name) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetInboxReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inbox []*InboxData `protobuf:"bytes,1,rep,name=inbox,proto3" json:"inbox,omitempty"`
}

func (x *GetInboxReply) Reset() {
	*x = GetInboxReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hvxahv_v1_inbox_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInboxReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInboxReply) ProtoMessage() {}

func (x *GetInboxReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_hvxahv_v1_inbox_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInboxReply.ProtoReflect.Descriptor instead.
func (*GetInboxReply) Descriptor() ([]byte, []int) {
	return file_api_hvxahv_v1_inbox_proto_rawDescGZIP(), []int{3}
}

func (x *GetInboxReply) GetInbox() []*InboxData {
	if x != nil {
		return x.Inbox
	}
	return nil
}

var File_api_hvxahv_v1_inbox_proto protoreflect.FileDescriptor

var file_api_hvxahv_v1_inbox_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2f, 0x76, 0x31, 0x2f,
	0x69, 0x6e, 0x62, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x68, 0x76, 0x78,
	0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x71, 0x0a, 0x09,
	0x49, 0x6e, 0x62, 0x6f, 0x78, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x63, 0x74,
	0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x22,
	0x25, 0x0a, 0x0d, 0x4e, 0x65, 0x77, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1a, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x41, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x30, 0x0a, 0x05, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05,
	0x69, 0x6e, 0x62, 0x6f, 0x78, 0x32, 0x96, 0x01, 0x0a, 0x05, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x12,
	0x48, 0x0a, 0x08, 0x4e, 0x65, 0x77, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x12, 0x1a, 0x2e, 0x68, 0x76,
	0x78, 0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e,
	0x62, 0x6f, 0x78, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x1e, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76,
	0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x77, 0x49, 0x6e, 0x62,
	0x6f, 0x78, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x08, 0x47, 0x65, 0x74,
	0x49, 0x6e, 0x62, 0x6f, 0x78, 0x12, 0x15, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76,
	0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x1e, 0x2e, 0x68,
	0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47,
	0x65, 0x74, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x06,
	0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_hvxahv_v1_inbox_proto_rawDescOnce sync.Once
	file_api_hvxahv_v1_inbox_proto_rawDescData = file_api_hvxahv_v1_inbox_proto_rawDesc
)

func file_api_hvxahv_v1_inbox_proto_rawDescGZIP() []byte {
	file_api_hvxahv_v1_inbox_proto_rawDescOnce.Do(func() {
		file_api_hvxahv_v1_inbox_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_hvxahv_v1_inbox_proto_rawDescData)
	})
	return file_api_hvxahv_v1_inbox_proto_rawDescData
}

var file_api_hvxahv_v1_inbox_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_hvxahv_v1_inbox_proto_goTypes = []interface{}{
	(*InboxData)(nil),     // 0: hvxahv.v1alpha1.proto.InboxData
	(*NewInboxReply)(nil), // 1: hvxahv.v1alpha1.proto.NewInboxReply
	(*Name)(nil),          // 2: hvxahv.v1alpha1.proto.Name
	(*GetInboxReply)(nil), // 3: hvxahv.v1alpha1.proto.GetInboxReply
}
var file_api_hvxahv_v1_inbox_proto_depIdxs = []int32{
	0, // 0: hvxahv.v1alpha1.proto.GetInboxReply.inbox:type_name -> hvxahv.v1alpha1.proto.InboxData
	0, // 1: hvxahv.v1alpha1.proto.Inbox.NewInbox:input_type -> hvxahv.v1alpha1.proto.InboxData
	2, // 2: hvxahv.v1alpha1.proto.Inbox.GetInbox:input_type -> hvxahv.v1alpha1.proto.Name
	1, // 3: hvxahv.v1alpha1.proto.Inbox.NewInbox:output_type -> hvxahv.v1alpha1.proto.NewInboxReply
	3, // 4: hvxahv.v1alpha1.proto.Inbox.GetInbox:output_type -> hvxahv.v1alpha1.proto.GetInboxReply
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_hvxahv_v1_inbox_proto_init() }
func file_api_hvxahv_v1_inbox_proto_init() {
	if File_api_hvxahv_v1_inbox_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_hvxahv_v1_inbox_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InboxData); i {
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
		file_api_hvxahv_v1_inbox_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewInboxReply); i {
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
		file_api_hvxahv_v1_inbox_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_hvxahv_v1_inbox_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInboxReply); i {
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
			RawDescriptor: file_api_hvxahv_v1_inbox_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_hvxahv_v1_inbox_proto_goTypes,
		DependencyIndexes: file_api_hvxahv_v1_inbox_proto_depIdxs,
		MessageInfos:      file_api_hvxahv_v1_inbox_proto_msgTypes,
	}.Build()
	File_api_hvxahv_v1_inbox_proto = out.File
	file_api_hvxahv_v1_inbox_proto_rawDesc = nil
	file_api_hvxahv_v1_inbox_proto_goTypes = nil
	file_api_hvxahv_v1_inbox_proto_depIdxs = nil
}
