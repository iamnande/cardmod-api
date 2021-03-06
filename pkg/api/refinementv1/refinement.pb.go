// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: iamnande/cardmod/refinement/v1/refinement.proto

package refinementv1

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

// An refinement resource.
type Refinement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Source of the refinement.
	Source string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	// Target of the refinement.
	Target string `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	// Numerator of the refinement ratio.
	Numerator int32 `protobuf:"varint,3,opt,name=numerator,proto3" json:"numerator,omitempty"`
	// Denominator of the refinement ratio.
	Denominator int32 `protobuf:"varint,4,opt,name=denominator,proto3" json:"denominator,omitempty"`
}

func (x *Refinement) Reset() {
	*x = Refinement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iamnande_cardmod_refinement_v1_refinement_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Refinement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Refinement) ProtoMessage() {}

func (x *Refinement) ProtoReflect() protoreflect.Message {
	mi := &file_iamnande_cardmod_refinement_v1_refinement_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Refinement.ProtoReflect.Descriptor instead.
func (*Refinement) Descriptor() ([]byte, []int) {
	return file_iamnande_cardmod_refinement_v1_refinement_proto_rawDescGZIP(), []int{0}
}

func (x *Refinement) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Refinement) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *Refinement) GetNumerator() int32 {
	if x != nil {
		return x.Numerator
	}
	return 0
}

func (x *Refinement) GetDenominator() int32 {
	if x != nil {
		return x.Denominator
	}
	return 0
}

var File_iamnande_cardmod_refinement_v1_refinement_proto protoreflect.FileDescriptor

var file_iamnande_cardmod_refinement_v1_refinement_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x69, 0x61, 0x6d, 0x6e, 0x61, 0x6e, 0x64, 0x65, 0x2f, 0x63, 0x61, 0x72, 0x64, 0x6d,
	0x6f, 0x64, 0x2f, 0x72, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1e, 0x69, 0x61, 0x6d, 0x6e, 0x61, 0x6e, 0x64, 0x65, 0x2e, 0x63, 0x61, 0x72, 0x64,
	0x6d, 0x6f, 0x64, 0x2e, 0x72, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x22, 0x7c, 0x0a, 0x0a, 0x52, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x42,
	0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x61,
	0x6d, 0x6e, 0x61, 0x6e, 0x64, 0x65, 0x2f, 0x63, 0x61, 0x72, 0x64, 0x6d, 0x6f, 0x64, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x76, 0x31, 0x3b, 0x72, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_iamnande_cardmod_refinement_v1_refinement_proto_rawDescOnce sync.Once
	file_iamnande_cardmod_refinement_v1_refinement_proto_rawDescData = file_iamnande_cardmod_refinement_v1_refinement_proto_rawDesc
)

func file_iamnande_cardmod_refinement_v1_refinement_proto_rawDescGZIP() []byte {
	file_iamnande_cardmod_refinement_v1_refinement_proto_rawDescOnce.Do(func() {
		file_iamnande_cardmod_refinement_v1_refinement_proto_rawDescData = protoimpl.X.CompressGZIP(file_iamnande_cardmod_refinement_v1_refinement_proto_rawDescData)
	})
	return file_iamnande_cardmod_refinement_v1_refinement_proto_rawDescData
}

var file_iamnande_cardmod_refinement_v1_refinement_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_iamnande_cardmod_refinement_v1_refinement_proto_goTypes = []interface{}{
	(*Refinement)(nil), // 0: iamnande.cardmod.refinement.v1.Refinement
}
var file_iamnande_cardmod_refinement_v1_refinement_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_iamnande_cardmod_refinement_v1_refinement_proto_init() }
func file_iamnande_cardmod_refinement_v1_refinement_proto_init() {
	if File_iamnande_cardmod_refinement_v1_refinement_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_iamnande_cardmod_refinement_v1_refinement_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Refinement); i {
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
			RawDescriptor: file_iamnande_cardmod_refinement_v1_refinement_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_iamnande_cardmod_refinement_v1_refinement_proto_goTypes,
		DependencyIndexes: file_iamnande_cardmod_refinement_v1_refinement_proto_depIdxs,
		MessageInfos:      file_iamnande_cardmod_refinement_v1_refinement_proto_msgTypes,
	}.Build()
	File_iamnande_cardmod_refinement_v1_refinement_proto = out.File
	file_iamnande_cardmod_refinement_v1_refinement_proto_rawDesc = nil
	file_iamnande_cardmod_refinement_v1_refinement_proto_goTypes = nil
	file_iamnande_cardmod_refinement_v1_refinement_proto_depIdxs = nil
}
