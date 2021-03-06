// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: iamnande/cardmod/calculate/v1/api.proto

package calculatev1

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

// Request message for Calculate.
type CalculateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Target is the desired refinement target.
	Target string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	// Quantity of the target to calculate against.
	Count int32 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *CalculateRequest) Reset() {
	*x = CalculateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iamnande_cardmod_calculate_v1_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculateRequest) ProtoMessage() {}

func (x *CalculateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iamnande_cardmod_calculate_v1_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculateRequest.ProtoReflect.Descriptor instead.
func (*CalculateRequest) Descriptor() ([]byte, []int) {
	return file_iamnande_cardmod_calculate_v1_api_proto_rawDescGZIP(), []int{0}
}

func (x *CalculateRequest) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *CalculateRequest) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

// Response output for Calculate.
type CalculateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A collection of refinement options.
	Refinements []*CalculateResponse_Refinement `protobuf:"bytes,1,rep,name=refinements,proto3" json:"refinements,omitempty"`
}

func (x *CalculateResponse) Reset() {
	*x = CalculateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iamnande_cardmod_calculate_v1_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculateResponse) ProtoMessage() {}

func (x *CalculateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_iamnande_cardmod_calculate_v1_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculateResponse.ProtoReflect.Descriptor instead.
func (*CalculateResponse) Descriptor() ([]byte, []int) {
	return file_iamnande_cardmod_calculate_v1_api_proto_rawDescGZIP(), []int{1}
}

func (x *CalculateResponse) GetRefinements() []*CalculateResponse_Refinement {
	if x != nil {
		return x.Refinements
	}
	return nil
}

// Refinement output of a given refinement.
type CalculateResponse_Refinement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Source option for the refinement.
	Source string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	// Quantity of the source needed for the desired amount.
	Count int32 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	// Complete origin calculation required.
	// NOTE: A maximum of two levels will ever be returned.
	Refinements []*CalculateResponse_Refinement `protobuf:"bytes,3,rep,name=refinements,proto3" json:"refinements,omitempty"`
}

func (x *CalculateResponse_Refinement) Reset() {
	*x = CalculateResponse_Refinement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iamnande_cardmod_calculate_v1_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculateResponse_Refinement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculateResponse_Refinement) ProtoMessage() {}

func (x *CalculateResponse_Refinement) ProtoReflect() protoreflect.Message {
	mi := &file_iamnande_cardmod_calculate_v1_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculateResponse_Refinement.ProtoReflect.Descriptor instead.
func (*CalculateResponse_Refinement) Descriptor() ([]byte, []int) {
	return file_iamnande_cardmod_calculate_v1_api_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CalculateResponse_Refinement) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *CalculateResponse_Refinement) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *CalculateResponse_Refinement) GetRefinements() []*CalculateResponse_Refinement {
	if x != nil {
		return x.Refinements
	}
	return nil
}

var File_iamnande_cardmod_calculate_v1_api_proto protoreflect.FileDescriptor

var file_iamnande_cardmod_calculate_v1_api_proto_rawDesc = []byte{
	0x0a, 0x27, 0x69, 0x61, 0x6d, 0x6e, 0x61, 0x6e, 0x64, 0x65, 0x2f, 0x63, 0x61, 0x72, 0x64, 0x6d,
	0x6f, 0x64, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x69, 0x61, 0x6d, 0x6e, 0x61,
	0x6e, 0x64, 0x65, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x6d, 0x6f, 0x64, 0x2e, 0x63, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x63, 0x0a, 0x10, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x15, 0xfa, 0x42, 0x12, 0x72, 0x10, 0x10, 0x03, 0x18, 0x19,
	0x32, 0x0a, 0x5e, 0x5b, 0x2d, 0x2c, 0x20, 0x5c, 0x77, 0x5d, 0x2b, 0x24, 0x52, 0x06, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x12, 0x20, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x1a, 0x05, 0x18, 0xac, 0x02, 0x28, 0x01, 0x52,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x8e, 0x02, 0x0a, 0x11, 0x43, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x0b,
	0x72, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x3b, 0x2e, 0x69, 0x61, 0x6d, 0x6e, 0x61, 0x6e, 0x64, 0x65, 0x2e, 0x63, 0x61, 0x72,
	0x64, 0x6d, 0x6f, 0x64, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b,
	0x72, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x99, 0x01, 0x0a, 0x0a,
	0x52, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x5d, 0x0a, 0x0b, 0x72, 0x65, 0x66, 0x69,
	0x6e, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e,
	0x69, 0x61, 0x6d, 0x6e, 0x61, 0x6e, 0x64, 0x65, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x6d, 0x6f, 0x64,
	0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x52, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x72, 0x65, 0x66, 0x69,
	0x6e, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x7e, 0x0a, 0x0c, 0x43, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x65, 0x41, 0x50, 0x49, 0x12, 0x6e, 0x0a, 0x09, 0x43, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x65, 0x12, 0x2f, 0x2e, 0x69, 0x61, 0x6d, 0x6e, 0x61, 0x6e, 0x64, 0x65, 0x2e,
	0x63, 0x61, 0x72, 0x64, 0x6d, 0x6f, 0x64, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x69, 0x61, 0x6d, 0x6e, 0x61, 0x6e, 0x64, 0x65,
	0x2e, 0x63, 0x61, 0x72, 0x64, 0x6d, 0x6f, 0x64, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61,
	0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x61, 0x6d, 0x6e, 0x61, 0x6e, 0x64, 0x65, 0x2f, 0x63,
	0x61, 0x72, 0x64, 0x6d, 0x6f, 0x64, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x76, 0x31, 0x3b, 0x63, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_iamnande_cardmod_calculate_v1_api_proto_rawDescOnce sync.Once
	file_iamnande_cardmod_calculate_v1_api_proto_rawDescData = file_iamnande_cardmod_calculate_v1_api_proto_rawDesc
)

func file_iamnande_cardmod_calculate_v1_api_proto_rawDescGZIP() []byte {
	file_iamnande_cardmod_calculate_v1_api_proto_rawDescOnce.Do(func() {
		file_iamnande_cardmod_calculate_v1_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_iamnande_cardmod_calculate_v1_api_proto_rawDescData)
	})
	return file_iamnande_cardmod_calculate_v1_api_proto_rawDescData
}

var file_iamnande_cardmod_calculate_v1_api_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_iamnande_cardmod_calculate_v1_api_proto_goTypes = []interface{}{
	(*CalculateRequest)(nil),             // 0: iamnande.cardmod.calculate.v1.CalculateRequest
	(*CalculateResponse)(nil),            // 1: iamnande.cardmod.calculate.v1.CalculateResponse
	(*CalculateResponse_Refinement)(nil), // 2: iamnande.cardmod.calculate.v1.CalculateResponse.Refinement
}
var file_iamnande_cardmod_calculate_v1_api_proto_depIdxs = []int32{
	2, // 0: iamnande.cardmod.calculate.v1.CalculateResponse.refinements:type_name -> iamnande.cardmod.calculate.v1.CalculateResponse.Refinement
	2, // 1: iamnande.cardmod.calculate.v1.CalculateResponse.Refinement.refinements:type_name -> iamnande.cardmod.calculate.v1.CalculateResponse.Refinement
	0, // 2: iamnande.cardmod.calculate.v1.CalculateAPI.Calculate:input_type -> iamnande.cardmod.calculate.v1.CalculateRequest
	1, // 3: iamnande.cardmod.calculate.v1.CalculateAPI.Calculate:output_type -> iamnande.cardmod.calculate.v1.CalculateResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_iamnande_cardmod_calculate_v1_api_proto_init() }
func file_iamnande_cardmod_calculate_v1_api_proto_init() {
	if File_iamnande_cardmod_calculate_v1_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_iamnande_cardmod_calculate_v1_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculateRequest); i {
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
		file_iamnande_cardmod_calculate_v1_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculateResponse); i {
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
		file_iamnande_cardmod_calculate_v1_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculateResponse_Refinement); i {
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
			RawDescriptor: file_iamnande_cardmod_calculate_v1_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_iamnande_cardmod_calculate_v1_api_proto_goTypes,
		DependencyIndexes: file_iamnande_cardmod_calculate_v1_api_proto_depIdxs,
		MessageInfos:      file_iamnande_cardmod_calculate_v1_api_proto_msgTypes,
	}.Build()
	File_iamnande_cardmod_calculate_v1_api_proto = out.File
	file_iamnande_cardmod_calculate_v1_api_proto_rawDesc = nil
	file_iamnande_cardmod_calculate_v1_api_proto_goTypes = nil
	file_iamnande_cardmod_calculate_v1_api_proto_depIdxs = nil
}
