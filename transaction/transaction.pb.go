// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: transaction.proto

package transaction

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

type Listen struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// `gorm:"primaryKey;not null;autoIncrement"`
	Id          int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" gorm:"primaryKey;not null;autoIncrement"`
	MemberId    int64 `protobuf:"varint,2,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	SentenceId  int64 `protobuf:"varint,3,opt,name=sentence_id,json=sentenceId,proto3" json:"sentence_id,omitempty"`
	ParagraphId int64 `protobuf:"varint,4,opt,name=paragraph_id,json=paragraphId,proto3" json:"paragraph_id,omitempty"`
	DocumentId  int64 `protobuf:"varint,5,opt,name=document_id,json=documentId,proto3" json:"document_id,omitempty"`
	CreatedAt   int64 `protobuf:"varint,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Listen) Reset() {
	*x = Listen{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Listen) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Listen) ProtoMessage() {}

func (x *Listen) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Listen.ProtoReflect.Descriptor instead.
func (*Listen) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *Listen) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Listen) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *Listen) GetSentenceId() int64 {
	if x != nil {
		return x.SentenceId
	}
	return 0
}

func (x *Listen) GetParagraphId() int64 {
	if x != nil {
		return x.ParagraphId
	}
	return 0
}

func (x *Listen) GetDocumentId() int64 {
	if x != nil {
		return x.DocumentId
	}
	return 0
}

func (x *Listen) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type ListenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	MemberId    int64    `protobuf:"varint,2,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	SentenceId  int64    `protobuf:"varint,3,opt,name=sentence_id,json=sentenceId,proto3" json:"sentence_id,omitempty"`
	ParagraphId int64    `protobuf:"varint,4,opt,name=paragraph_id,json=paragraphId,proto3" json:"paragraph_id,omitempty"`
	DocumentId  int64    `protobuf:"varint,5,opt,name=document_id,json=documentId,proto3" json:"document_id,omitempty"`
	CreatedAt   int64    `protobuf:"varint,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Limit       int32    `protobuf:"varint,17,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset      int32    `protobuf:"varint,18,opt,name=offset,proto3" json:"offset,omitempty"`
	Cols        []string `protobuf:"bytes,16,rep,name=cols,proto3" json:"cols,omitempty"`
	Includes    []string `protobuf:"bytes,19,rep,name=includes,proto3" json:"includes,omitempty"`
}

func (x *ListenRequest) Reset() {
	*x = ListenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListenRequest) ProtoMessage() {}

func (x *ListenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListenRequest.ProtoReflect.Descriptor instead.
func (*ListenRequest) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{1}
}

func (x *ListenRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ListenRequest) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *ListenRequest) GetSentenceId() int64 {
	if x != nil {
		return x.SentenceId
	}
	return 0
}

func (x *ListenRequest) GetParagraphId() int64 {
	if x != nil {
		return x.ParagraphId
	}
	return 0
}

func (x *ListenRequest) GetDocumentId() int64 {
	if x != nil {
		return x.DocumentId
	}
	return 0
}

func (x *ListenRequest) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *ListenRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListenRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListenRequest) GetCols() []string {
	if x != nil {
		return x.Cols
	}
	return nil
}

func (x *ListenRequest) GetIncludes() []string {
	if x != nil {
		return x.Includes
	}
	return nil
}

type Speak struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// `gorm:"primaryKey;not null;autoIncrement"`
	Id          int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" gorm:"primaryKey;not null;autoIncrement"`
	MemberId    int64   `protobuf:"varint,2,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	SentenceId  int64   `protobuf:"varint,3,opt,name=sentence_id,json=sentenceId,proto3" json:"sentence_id,omitempty"`
	ParagraphId int64   `protobuf:"varint,4,opt,name=paragraph_id,json=paragraphId,proto3" json:"paragraph_id,omitempty"`
	CharacterId int64   `protobuf:"varint,5,opt,name=character_id,json=characterId,proto3" json:"character_id,omitempty"`
	CreatedAt   int64   `protobuf:"varint,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Score       float32 `protobuf:"fixed32,7,opt,name=score,proto3" json:"score,omitempty"`
	Hash        string  `protobuf:"bytes,8,opt,name=hash,proto3" json:"hash,omitempty"`
	Keyx        string  `protobuf:"bytes,9,opt,name=keyx,proto3" json:"keyx,omitempty"`
	Format      string  `protobuf:"bytes,10,opt,name=format,proto3" json:"format,omitempty"`
	// `gorm:"-"`
	Color   string `protobuf:"bytes,11,opt,name=color,proto3" json:"color,omitempty" gorm:"-"`
	Content string `protobuf:"bytes,12,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Speak) Reset() {
	*x = Speak{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Speak) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Speak) ProtoMessage() {}

func (x *Speak) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Speak.ProtoReflect.Descriptor instead.
func (*Speak) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{2}
}

func (x *Speak) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Speak) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *Speak) GetSentenceId() int64 {
	if x != nil {
		return x.SentenceId
	}
	return 0
}

func (x *Speak) GetParagraphId() int64 {
	if x != nil {
		return x.ParagraphId
	}
	return 0
}

func (x *Speak) GetCharacterId() int64 {
	if x != nil {
		return x.CharacterId
	}
	return 0
}

func (x *Speak) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Speak) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Speak) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *Speak) GetKeyx() string {
	if x != nil {
		return x.Keyx
	}
	return ""
}

func (x *Speak) GetFormat() string {
	if x != nil {
		return x.Format
	}
	return ""
}

func (x *Speak) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *Speak) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type Speaks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Speaks []*Speak `protobuf:"bytes,1,rep,name=speaks,proto3" json:"speaks,omitempty"`
	Total  int64    `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *Speaks) Reset() {
	*x = Speaks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Speaks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Speaks) ProtoMessage() {}

func (x *Speaks) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Speaks.ProtoReflect.Descriptor instead.
func (*Speaks) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{3}
}

func (x *Speaks) GetSpeaks() []*Speak {
	if x != nil {
		return x.Speaks
	}
	return nil
}

func (x *Speaks) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type SpeakRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	MemberId     int64    `protobuf:"varint,2,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	SentenceId   int64    `protobuf:"varint,3,opt,name=sentence_id,json=sentenceId,proto3" json:"sentence_id,omitempty"`
	ParagraphId  int64    `protobuf:"varint,4,opt,name=paragraph_id,json=paragraphId,proto3" json:"paragraph_id,omitempty"`
	CharacterId  int64    `protobuf:"varint,5,opt,name=character_id,json=characterId,proto3" json:"character_id,omitempty"`
	CreatedAt    int64    `protobuf:"varint,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Score        float32  `protobuf:"fixed32,7,opt,name=score,proto3" json:"score,omitempty"`
	Hash         string   `protobuf:"bytes,8,opt,name=hash,proto3" json:"hash,omitempty"`
	Keyx         string   `protobuf:"bytes,9,opt,name=keyx,proto3" json:"keyx,omitempty"`
	Format       string   `protobuf:"bytes,10,opt,name=format,proto3" json:"format,omitempty"`
	Limit        int32    `protobuf:"varint,17,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset       int32    `protobuf:"varint,18,opt,name=offset,proto3" json:"offset,omitempty"`
	Cols         []string `protobuf:"bytes,16,rep,name=cols,proto3" json:"cols,omitempty"`
	Includes     []string `protobuf:"bytes,19,rep,name=includes,proto3" json:"includes,omitempty"`
	SentenceIds  []int64  `protobuf:"varint,20,rep,packed,name=sentence_ids,json=sentenceIds,proto3" json:"sentence_ids,omitempty"`
	CharacterIds []int64  `protobuf:"varint,21,rep,packed,name=character_ids,json=characterIds,proto3" json:"character_ids,omitempty"`
	CurrentId    int64    `protobuf:"varint,22,opt,name=current_id,json=currentId,proto3" json:"current_id,omitempty"`
	Ids          []int64  `protobuf:"varint,23,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *SpeakRequest) Reset() {
	*x = SpeakRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpeakRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpeakRequest) ProtoMessage() {}

func (x *SpeakRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpeakRequest.ProtoReflect.Descriptor instead.
func (*SpeakRequest) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{4}
}

func (x *SpeakRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SpeakRequest) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *SpeakRequest) GetSentenceId() int64 {
	if x != nil {
		return x.SentenceId
	}
	return 0
}

func (x *SpeakRequest) GetParagraphId() int64 {
	if x != nil {
		return x.ParagraphId
	}
	return 0
}

func (x *SpeakRequest) GetCharacterId() int64 {
	if x != nil {
		return x.CharacterId
	}
	return 0
}

func (x *SpeakRequest) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *SpeakRequest) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *SpeakRequest) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *SpeakRequest) GetKeyx() string {
	if x != nil {
		return x.Keyx
	}
	return ""
}

func (x *SpeakRequest) GetFormat() string {
	if x != nil {
		return x.Format
	}
	return ""
}

func (x *SpeakRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *SpeakRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *SpeakRequest) GetCols() []string {
	if x != nil {
		return x.Cols
	}
	return nil
}

func (x *SpeakRequest) GetIncludes() []string {
	if x != nil {
		return x.Includes
	}
	return nil
}

func (x *SpeakRequest) GetSentenceIds() []int64 {
	if x != nil {
		return x.SentenceIds
	}
	return nil
}

func (x *SpeakRequest) GetCharacterIds() []int64 {
	if x != nil {
		return x.CharacterIds
	}
	return nil
}

func (x *SpeakRequest) GetCurrentId() int64 {
	if x != nil {
		return x.CurrentId
	}
	return 0
}

func (x *SpeakRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type Lookup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// `gorm:"primaryKey;not null;autoIncrement"`
	Id         int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" gorm:"primaryKey;not null;autoIncrement"`
	MemberId   int64 `protobuf:"varint,2,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	SentenceId int64 `protobuf:"varint,3,opt,name=sentence_id,json=sentenceId,proto3" json:"sentence_id,omitempty"`
	PosId      int64 `protobuf:"varint,4,opt,name=pos_id,json=posId,proto3" json:"pos_id,omitempty"`
	WordId     int64 `protobuf:"varint,5,opt,name=word_id,json=wordId,proto3" json:"word_id,omitempty"`
	CreatedAt  int64 `protobuf:"varint,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Lookup) Reset() {
	*x = Lookup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lookup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lookup) ProtoMessage() {}

func (x *Lookup) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lookup.ProtoReflect.Descriptor instead.
func (*Lookup) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{5}
}

func (x *Lookup) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Lookup) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *Lookup) GetSentenceId() int64 {
	if x != nil {
		return x.SentenceId
	}
	return 0
}

func (x *Lookup) GetPosId() int64 {
	if x != nil {
		return x.PosId
	}
	return 0
}

func (x *Lookup) GetWordId() int64 {
	if x != nil {
		return x.WordId
	}
	return 0
}

func (x *Lookup) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

var File_transaction_proto protoreflect.FileDescriptor

var file_transaction_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0xb9, 0x01, 0x0a, 0x06, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x6e, 0x74,
	0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73,
	0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x72,
	0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x70, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x9e, 0x02, 0x0a,
	0x0d, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x70, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x70, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x49, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x12, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x6c, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x6c,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x73, 0x18, 0x13, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x73, 0x22, 0xc0, 0x02,
	0x0a, 0x05, 0x53, 0x70, 0x65, 0x61, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x65, 0x6e, 0x74, 0x65,
	0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61,
	0x70, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x70, 0x61, 0x72,
	0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x72,
	0x61, 0x63, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x68, 0x61, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x78, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x22, 0x4a, 0x0a, 0x06, 0x53, 0x70, 0x65, 0x61, 0x6b, 0x73, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x70,
	0x65, 0x61, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x70, 0x65, 0x61, 0x6b, 0x52, 0x06,
	0x73, 0x70, 0x65, 0x61, 0x6b, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0xee, 0x03, 0x0a,
	0x0c, 0x53, 0x70, 0x65, 0x61, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65,
	0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70,
	0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x70, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x49, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65,
	0x79, 0x78, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x78, 0x12, 0x16,
	0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x11, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x12, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x6c, 0x73, 0x18, 0x10, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x6c, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x63, 0x6c,
	0x75, 0x64, 0x65, 0x73, 0x18, 0x13, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x63, 0x6c,
	0x75, 0x64, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x14, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0b, 0x73, 0x65, 0x6e, 0x74,
	0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x15, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0c,
	0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x49, 0x64, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x69,
	0x64, 0x73, 0x18, 0x17, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0xa5, 0x01,
	0x0a, 0x06, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x65, 0x6e, 0x74,
	0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x6f, 0x73, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x77, 0x6f, 0x72, 0x64, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0x85, 0x02, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x0a,
	0x4c, 0x69, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x12, 0x1a, 0x2e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x12, 0x3a, 0x0a, 0x0e, 0x45,
	0x6e, 0x64, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x12, 0x13, 0x2e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x1a, 0x13, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x12, 0x3b, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x70, 0x65, 0x61, 0x6b, 0x12, 0x19, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x53, 0x70, 0x65, 0x61, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x13, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x70,
	0x65, 0x61, 0x6b, 0x73, 0x12, 0x37, 0x0a, 0x0d, 0x45, 0x6e, 0x64, 0x4c, 0x65, 0x61, 0x72, 0x6e,
	0x53, 0x70, 0x65, 0x61, 0x6b, 0x12, 0x12, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x70, 0x65, 0x61, 0x6b, 0x1a, 0x12, 0x2e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x70, 0x65, 0x61, 0x6b, 0x42, 0x2f, 0x5a,
	0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x51, 0x75, 0x61, 0x6e,
	0x67, 0x41, 0x6e, 0x68, 0x4d, 0x54, 0x41, 0x2f, 0x65, 0x6e, 0x74, 0x2d, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transaction_proto_rawDescOnce sync.Once
	file_transaction_proto_rawDescData = file_transaction_proto_rawDesc
)

func file_transaction_proto_rawDescGZIP() []byte {
	file_transaction_proto_rawDescOnce.Do(func() {
		file_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_transaction_proto_rawDescData)
	})
	return file_transaction_proto_rawDescData
}

var file_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_transaction_proto_goTypes = []interface{}{
	(*Listen)(nil),        // 0: transaction.Listen
	(*ListenRequest)(nil), // 1: transaction.ListenRequest
	(*Speak)(nil),         // 2: transaction.Speak
	(*Speaks)(nil),        // 3: transaction.Speaks
	(*SpeakRequest)(nil),  // 4: transaction.SpeakRequest
	(*Lookup)(nil),        // 5: transaction.Lookup
}
var file_transaction_proto_depIdxs = []int32{
	2, // 0: transaction.Speaks.speaks:type_name -> transaction.Speak
	1, // 1: transaction.TransactionService.ListListen:input_type -> transaction.ListenRequest
	0, // 2: transaction.TransactionService.EndLearnListen:input_type -> transaction.Listen
	4, // 3: transaction.TransactionService.ListSpeak:input_type -> transaction.SpeakRequest
	2, // 4: transaction.TransactionService.EndLearnSpeak:input_type -> transaction.Speak
	0, // 5: transaction.TransactionService.ListListen:output_type -> transaction.Listen
	0, // 6: transaction.TransactionService.EndLearnListen:output_type -> transaction.Listen
	3, // 7: transaction.TransactionService.ListSpeak:output_type -> transaction.Speaks
	2, // 8: transaction.TransactionService.EndLearnSpeak:output_type -> transaction.Speak
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_transaction_proto_init() }
func file_transaction_proto_init() {
	if File_transaction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Listen); i {
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
		file_transaction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListenRequest); i {
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
		file_transaction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Speak); i {
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
		file_transaction_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Speaks); i {
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
		file_transaction_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpeakRequest); i {
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
		file_transaction_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Lookup); i {
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
			RawDescriptor: file_transaction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_transaction_proto_goTypes,
		DependencyIndexes: file_transaction_proto_depIdxs,
		MessageInfos:      file_transaction_proto_msgTypes,
	}.Build()
	File_transaction_proto = out.File
	file_transaction_proto_rawDesc = nil
	file_transaction_proto_goTypes = nil
	file_transaction_proto_depIdxs = nil
}
