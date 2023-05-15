// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: notice.proto

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

type NoticeInsertReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonReq *CommonReq `protobuf:"bytes,1,opt,name=commonReq,proto3" json:"commonReq"`
}

func (x *NoticeInsertReq) Reset() {
	*x = NoticeInsertReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoticeInsertReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoticeInsertReq) ProtoMessage() {}

func (x *NoticeInsertReq) ProtoReflect() protoreflect.Message {
	mi := &file_notice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoticeInsertReq.ProtoReflect.Descriptor instead.
func (*NoticeInsertReq) Descriptor() ([]byte, []int) {
	return file_notice_proto_rawDescGZIP(), []int{0}
}

func (x *NoticeInsertReq) GetCommonReq() *CommonReq {
	if x != nil {
		return x.CommonReq
	}
	return nil
}

type NoticeInsertResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonResp *CommonResp `protobuf:"bytes,1,opt,name=commonResp,proto3" json:"commonResp"`
}

func (x *NoticeInsertResp) Reset() {
	*x = NoticeInsertResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoticeInsertResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoticeInsertResp) ProtoMessage() {}

func (x *NoticeInsertResp) ProtoReflect() protoreflect.Message {
	mi := &file_notice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoticeInsertResp.ProtoReflect.Descriptor instead.
func (*NoticeInsertResp) Descriptor() ([]byte, []int) {
	return file_notice_proto_rawDescGZIP(), []int{1}
}

func (x *NoticeInsertResp) GetCommonResp() *CommonResp {
	if x != nil {
		return x.CommonResp
	}
	return nil
}

var File_notice_proto protoreflect.FileDescriptor

var file_notice_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x3e, 0x0a, 0x0f, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x12, 0x2b, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x22, 0x42, 0x0a, 0x10, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x2e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x32, 0x4a, 0x0a, 0x0d, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x49,
	0x6e, 0x73, 0x65, 0x72, 0x74, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x63,
	0x65, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e,
	0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notice_proto_rawDescOnce sync.Once
	file_notice_proto_rawDescData = file_notice_proto_rawDesc
)

func file_notice_proto_rawDescGZIP() []byte {
	file_notice_proto_rawDescOnce.Do(func() {
		file_notice_proto_rawDescData = protoimpl.X.CompressGZIP(file_notice_proto_rawDescData)
	})
	return file_notice_proto_rawDescData
}

var file_notice_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_notice_proto_goTypes = []interface{}{
	(*NoticeInsertReq)(nil),  // 0: pb.NoticeInsertReq
	(*NoticeInsertResp)(nil), // 1: pb.NoticeInsertResp
	(*CommonReq)(nil),        // 2: pb.CommonReq
	(*CommonResp)(nil),       // 3: pb.CommonResp
}
var file_notice_proto_depIdxs = []int32{
	2, // 0: pb.NoticeInsertReq.commonReq:type_name -> pb.CommonReq
	3, // 1: pb.NoticeInsertResp.commonResp:type_name -> pb.CommonResp
	0, // 2: pb.noticeService.NoticeInsert:input_type -> pb.NoticeInsertReq
	1, // 3: pb.noticeService.NoticeInsert:output_type -> pb.NoticeInsertResp
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_notice_proto_init() }
func file_notice_proto_init() {
	if File_notice_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_notice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoticeInsertReq); i {
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
		file_notice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoticeInsertResp); i {
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
			RawDescriptor: file_notice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notice_proto_goTypes,
		DependencyIndexes: file_notice_proto_depIdxs,
		MessageInfos:      file_notice_proto_msgTypes,
	}.Build()
	File_notice_proto = out.File
	file_notice_proto_rawDesc = nil
	file_notice_proto_goTypes = nil
	file_notice_proto_depIdxs = nil
}
