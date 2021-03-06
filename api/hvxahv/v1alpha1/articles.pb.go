// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1alpha1.25.0-devel
// 	protoc        v3.14.0
// source: api/hvxahv/v1alpha1/articles.proto

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

// ArticleData Article 文章模型
type ArticleData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Author    string `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	Content   string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Images    string `protobuf:"bytes,4,opt,name=images,proto3" json:"images,omitempty"`
	Type      string `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	Private   bool   `protobuf:"varint,6,opt,name=private,proto3" json:"private,omitempty"`
	IsComment bool   `protobuf:"varint,7,opt,name=isComment,proto3" json:"isComment,omitempty"`
}

func (x *ArticleData) Reset() {
	*x = ArticleData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hvxahv_v1_articles_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleData) ProtoMessage() {}

func (x *ArticleData) ProtoReflect() protoreflect.Message {
	mi := &file_api_hvxahv_v1_articles_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleData.ProtoReflect.Descriptor instead.
func (*ArticleData) Descriptor() ([]byte, []int) {
	return file_api_hvxahv_v1_articles_proto_rawDescGZIP(), []int{0}
}

func (x *ArticleData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ArticleData) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *ArticleData) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ArticleData) GetImages() string {
	if x != nil {
		return x.Images
	}
	return ""
}

func (x *ArticleData) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ArticleData) GetPrivate() bool {
	if x != nil {
		return x.Private
	}
	return false
}

func (x *ArticleData) GetIsComment() bool {
	if x != nil {
		return x.IsComment
	}
	return false
}

// NewArticleReply ...
type NewArticleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reply string `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
}

func (x *NewArticleReply) Reset() {
	*x = NewArticleReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hvxahv_v1_articles_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewArticleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewArticleReply) ProtoMessage() {}

func (x *NewArticleReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_hvxahv_v1_articles_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewArticleReply.ProtoReflect.Descriptor instead.
func (*NewArticleReply) Descriptor() ([]byte, []int) {
	return file_api_hvxahv_v1_articles_proto_rawDescGZIP(), []int{1}
}

func (x *NewArticleReply) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

type UpdateArticleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reply string `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
}

func (x *UpdateArticleReply) Reset() {
	*x = UpdateArticleReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hvxahv_v1_articles_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateArticleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateArticleReply) ProtoMessage() {}

func (x *UpdateArticleReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_hvxahv_v1_articles_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateArticleReply.ProtoReflect.Descriptor instead.
func (*UpdateArticleReply) Descriptor() ([]byte, []int) {
	return file_api_hvxahv_v1_articles_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateArticleReply) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

type DeleteArticleByID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteArticleByID) Reset() {
	*x = DeleteArticleByID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hvxahv_v1_articles_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteArticleByID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteArticleByID) ProtoMessage() {}

func (x *DeleteArticleByID) ProtoReflect() protoreflect.Message {
	mi := &file_api_hvxahv_v1_articles_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteArticleByID.ProtoReflect.Descriptor instead.
func (*DeleteArticleByID) Descriptor() ([]byte, []int) {
	return file_api_hvxahv_v1_articles_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteArticleByID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteArticleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reply string `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
}

func (x *DeleteArticleReply) Reset() {
	*x = DeleteArticleReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hvxahv_v1_articles_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteArticleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteArticleReply) ProtoMessage() {}

func (x *DeleteArticleReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_hvxahv_v1_articles_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteArticleReply.ProtoReflect.Descriptor instead.
func (*DeleteArticleReply) Descriptor() ([]byte, []int) {
	return file_api_hvxahv_v1_articles_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteArticleReply) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

// GetArticle ...
type GetArticleData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetArticleData) Reset() {
	*x = GetArticleData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hvxahv_v1_articles_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArticleData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticleData) ProtoMessage() {}

func (x *GetArticleData) ProtoReflect() protoreflect.Message {
	mi := &file_api_hvxahv_v1_articles_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticleData.ProtoReflect.Descriptor instead.
func (*GetArticleData) Descriptor() ([]byte, []int) {
	return file_api_hvxahv_v1_articles_proto_rawDescGZIP(), []int{5}
}

func (x *GetArticleData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetArticleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Articles []*ArticleData `protobuf:"bytes,1,rep,name=articles,proto3" json:"articles,omitempty"`
}

func (x *GetArticleReply) Reset() {
	*x = GetArticleReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hvxahv_v1_articles_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArticleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticleReply) ProtoMessage() {}

func (x *GetArticleReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_hvxahv_v1_articles_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticleReply.ProtoReflect.Descriptor instead.
func (*GetArticleReply) Descriptor() ([]byte, []int) {
	return file_api_hvxahv_v1_articles_proto_rawDescGZIP(), []int{6}
}

func (x *GetArticleReply) GetArticles() []*ArticleData {
	if x != nil {
		return x.Articles
	}
	return nil
}

var File_api_hvxahv_v1_articles_proto protoreflect.FileDescriptor

var file_api_hvxahv_v1_articles_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xb3, 0x01, 0x0a, 0x0b, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x27, 0x0a, 0x0f, 0x4e, 0x65, 0x77, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2a,
	0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x42, 0x79, 0x49, 0x44, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x2a, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x24, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x4b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x38, 0x0a, 0x08, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e,
	0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x32, 0xe0,
	0x02, 0x0a, 0x08, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x12, 0x4e, 0x0a, 0x0a, 0x4e,
	0x65, 0x77, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x1c, 0x2e, 0x68, 0x76, 0x78, 0x61,
	0x68, 0x76, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x20, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76,
	0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x77, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x12, 0x1f, 0x2e, 0x68, 0x76, 0x78,
	0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x20, 0x2e, 0x68, 0x76,
	0x78, 0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12,
	0x54, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x12, 0x1c, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x23,
	0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x22, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e,
	0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x42, 0x79, 0x49, 0x44, 0x1a, 0x23, 0x2e, 0x68, 0x76, 0x78,
	0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_hvxahv_v1_articles_proto_rawDescOnce sync.Once
	file_api_hvxahv_v1_articles_proto_rawDescData = file_api_hvxahv_v1_articles_proto_rawDesc
)

func file_api_hvxahv_v1_articles_proto_rawDescGZIP() []byte {
	file_api_hvxahv_v1_articles_proto_rawDescOnce.Do(func() {
		file_api_hvxahv_v1_articles_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_hvxahv_v1_articles_proto_rawDescData)
	})
	return file_api_hvxahv_v1_articles_proto_rawDescData
}

var file_api_hvxahv_v1_articles_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_hvxahv_v1_articles_proto_goTypes = []interface{}{
	(*ArticleData)(nil),        // 0: hvxahv.v1alpha1.proto.ArticleData
	(*NewArticleReply)(nil),    // 1: hvxahv.v1alpha1.proto.NewArticleReply
	(*UpdateArticleReply)(nil), // 2: hvxahv.v1alpha1.proto.UpdateArticleReply
	(*DeleteArticleByID)(nil),  // 3: hvxahv.v1alpha1.proto.DeleteArticleByID
	(*DeleteArticleReply)(nil), // 4: hvxahv.v1alpha1.proto.DeleteArticleReply
	(*GetArticleData)(nil),     // 5: hvxahv.v1alpha1.proto.GetArticleData
	(*GetArticleReply)(nil),    // 6: hvxahv.v1alpha1.proto.GetArticleReply
}
var file_api_hvxahv_v1_articles_proto_depIdxs = []int32{
	0, // 0: hvxahv.v1alpha1.proto.GetArticleReply.articles:type_name -> hvxahv.v1alpha1.proto.ArticleData
	0, // 1: hvxahv.v1alpha1.proto.Articles.NewArticle:input_type -> hvxahv.v1alpha1.proto.ArticleData
	5, // 2: hvxahv.v1alpha1.proto.Articles.GetArticles:input_type -> hvxahv.v1alpha1.proto.GetArticleData
	0, // 3: hvxahv.v1alpha1.proto.Articles.UpdateArticle:input_type -> hvxahv.v1alpha1.proto.ArticleData
	3, // 4: hvxahv.v1alpha1.proto.Articles.DeleteArticle:input_type -> hvxahv.v1alpha1.proto.DeleteArticleByID
	1, // 5: hvxahv.v1alpha1.proto.Articles.NewArticle:output_type -> hvxahv.v1alpha1.proto.NewArticleReply
	6, // 6: hvxahv.v1alpha1.proto.Articles.GetArticles:output_type -> hvxahv.v1alpha1.proto.GetArticleReply
	2, // 7: hvxahv.v1alpha1.proto.Articles.UpdateArticle:output_type -> hvxahv.v1alpha1.proto.UpdateArticleReply
	4, // 8: hvxahv.v1alpha1.proto.Articles.DeleteArticle:output_type -> hvxahv.v1alpha1.proto.DeleteArticleReply
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_hvxahv_v1_articles_proto_init() }
func file_api_hvxahv_v1_articles_proto_init() {
	if File_api_hvxahv_v1_articles_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_hvxahv_v1_articles_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleData); i {
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
		file_api_hvxahv_v1_articles_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewArticleReply); i {
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
		file_api_hvxahv_v1_articles_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateArticleReply); i {
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
		file_api_hvxahv_v1_articles_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteArticleByID); i {
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
		file_api_hvxahv_v1_articles_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteArticleReply); i {
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
		file_api_hvxahv_v1_articles_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArticleData); i {
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
		file_api_hvxahv_v1_articles_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArticleReply); i {
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
			RawDescriptor: file_api_hvxahv_v1_articles_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_hvxahv_v1_articles_proto_goTypes,
		DependencyIndexes: file_api_hvxahv_v1_articles_proto_depIdxs,
		MessageInfos:      file_api_hvxahv_v1_articles_proto_msgTypes,
	}.Build()
	File_api_hvxahv_v1_articles_proto = out.File
	file_api_hvxahv_v1_articles_proto_rawDesc = nil
	file_api_hvxahv_v1_articles_proto_goTypes = nil
	file_api_hvxahv_v1_articles_proto_depIdxs = nil
}
