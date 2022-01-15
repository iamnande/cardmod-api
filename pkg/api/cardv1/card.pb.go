// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: iamnande/cardmod/card/v1/card.proto

package cardv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// A collection of card resources.
type Cards struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// cards is a collection of card entities.
	Cards []*Card `protobuf:"bytes,1,rep,name=cards,proto3" json:"cards,omitempty"`
}

func (x *Cards) Reset() {
	*x = Cards{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iamnande_cardmod_card_v1_card_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cards) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cards) ProtoMessage() {}

func (x *Cards) ProtoReflect() protoreflect.Message {
	mi := &file_iamnande_cardmod_card_v1_card_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cards.ProtoReflect.Descriptor instead.
func (*Cards) Descriptor() ([]byte, []int) {
	return file_iamnande_cardmod_card_v1_card_proto_rawDescGZIP(), []int{0}
}

func (x *Cards) GetCards() []*Card {
	if x != nil {
		return x.Cards
	}
	return nil
}

// A card resource.
type Card struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the card.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name of the card.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Card) Reset() {
	*x = Card{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iamnande_cardmod_card_v1_card_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Card) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Card) ProtoMessage() {}

func (x *Card) ProtoReflect() protoreflect.Message {
	mi := &file_iamnande_cardmod_card_v1_card_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Card.ProtoReflect.Descriptor instead.
func (*Card) Descriptor() ([]byte, []int) {
	return file_iamnande_cardmod_card_v1_card_proto_rawDescGZIP(), []int{1}
}

func (x *Card) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Card) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_iamnande_cardmod_card_v1_card_proto protoreflect.FileDescriptor

var file_iamnande_cardmod_card_v1_card_proto_rawDesc = []byte{
	0x0a, 0x23, 0x69, 0x61, 0x6d, 0x6e, 0x61, 0x6e, 0x64, 0x65, 0x2f, 0x63, 0x61, 0x72, 0x64, 0x6d,
	0x6f, 0x64, 0x2f, 0x63, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x72, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x69, 0x61, 0x6d, 0x6e, 0x61, 0x6e, 0x64, 0x65, 0x2e,
	0x63, 0x61, 0x72, 0x64, 0x6d, 0x6f, 0x64, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3d, 0x0a, 0x05, 0x43, 0x61, 0x72, 0x64,
	0x73, 0x12, 0x34, 0x0a, 0x05, 0x63, 0x61, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x69, 0x61, 0x6d, 0x6e, 0x61, 0x6e, 0x64, 0x65, 0x2e, 0x63, 0x61, 0x72, 0x64,
	0x6d, 0x6f, 0x64, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x72, 0x64,
	0x52, 0x05, 0x63, 0x61, 0x72, 0x64, 0x73, 0x22, 0x2a, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x64, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x69, 0x61, 0x6d, 0x6e, 0x61, 0x6e, 0x64, 0x65, 0x2f, 0x63, 0x61, 0x72, 0x64, 0x6d,
	0x6f, 0x64, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x72, 0x64, 0x76,
	0x31, 0x3b, 0x63, 0x61, 0x72, 0x64, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_iamnande_cardmod_card_v1_card_proto_rawDescOnce sync.Once
	file_iamnande_cardmod_card_v1_card_proto_rawDescData = file_iamnande_cardmod_card_v1_card_proto_rawDesc
)

func file_iamnande_cardmod_card_v1_card_proto_rawDescGZIP() []byte {
	file_iamnande_cardmod_card_v1_card_proto_rawDescOnce.Do(func() {
		file_iamnande_cardmod_card_v1_card_proto_rawDescData = protoimpl.X.CompressGZIP(file_iamnande_cardmod_card_v1_card_proto_rawDescData)
	})
	return file_iamnande_cardmod_card_v1_card_proto_rawDescData
}

var file_iamnande_cardmod_card_v1_card_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_iamnande_cardmod_card_v1_card_proto_goTypes = []interface{}{
	(*Cards)(nil), // 0: iamnande.cardmod.card.v1.Cards
	(*Card)(nil),  // 1: iamnande.cardmod.card.v1.Card
}
var file_iamnande_cardmod_card_v1_card_proto_depIdxs = []int32{
	1, // 0: iamnande.cardmod.card.v1.Cards.cards:type_name -> iamnande.cardmod.card.v1.Card
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_iamnande_cardmod_card_v1_card_proto_init() }
func file_iamnande_cardmod_card_v1_card_proto_init() {
	if File_iamnande_cardmod_card_v1_card_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_iamnande_cardmod_card_v1_card_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cards); i {
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
		file_iamnande_cardmod_card_v1_card_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Card); i {
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
			RawDescriptor: file_iamnande_cardmod_card_v1_card_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_iamnande_cardmod_card_v1_card_proto_goTypes,
		DependencyIndexes: file_iamnande_cardmod_card_v1_card_proto_depIdxs,
		MessageInfos:      file_iamnande_cardmod_card_v1_card_proto_msgTypes,
	}.Build()
	File_iamnande_cardmod_card_v1_card_proto = out.File
	file_iamnande_cardmod_card_v1_card_proto_rawDesc = nil
	file_iamnande_cardmod_card_v1_card_proto_goTypes = nil
	file_iamnande_cardmod_card_v1_card_proto_depIdxs = nil
}
