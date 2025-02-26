// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: urlshortener.proto

package urlshortener_v3

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

type OriginalURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OriginalUrl string `protobuf:"bytes,1,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
}

func (x *OriginalURLRequest) Reset() {
	*x = OriginalURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_urlshortener_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OriginalURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OriginalURLRequest) ProtoMessage() {}

func (x *OriginalURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_urlshortener_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OriginalURLRequest.ProtoReflect.Descriptor instead.
func (*OriginalURLRequest) Descriptor() ([]byte, []int) {
	return file_urlshortener_proto_rawDescGZIP(), []int{0}
}

func (x *OriginalURLRequest) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

type ShortURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrl string `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
}

func (x *ShortURLResponse) Reset() {
	*x = ShortURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_urlshortener_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortURLResponse) ProtoMessage() {}

func (x *ShortURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_urlshortener_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortURLResponse.ProtoReflect.Descriptor instead.
func (*ShortURLResponse) Descriptor() ([]byte, []int) {
	return file_urlshortener_proto_rawDescGZIP(), []int{1}
}

func (x *ShortURLResponse) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

type ShortURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrl string `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
}

func (x *ShortURLRequest) Reset() {
	*x = ShortURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_urlshortener_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortURLRequest) ProtoMessage() {}

func (x *ShortURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_urlshortener_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortURLRequest.ProtoReflect.Descriptor instead.
func (*ShortURLRequest) Descriptor() ([]byte, []int) {
	return file_urlshortener_proto_rawDescGZIP(), []int{2}
}

func (x *ShortURLRequest) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

type OriginalURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OriginalUrl string `protobuf:"bytes,1,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
}

func (x *OriginalURLResponse) Reset() {
	*x = OriginalURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_urlshortener_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OriginalURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OriginalURLResponse) ProtoMessage() {}

func (x *OriginalURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_urlshortener_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OriginalURLResponse.ProtoReflect.Descriptor instead.
func (*OriginalURLResponse) Descriptor() ([]byte, []int) {
	return file_urlshortener_proto_rawDescGZIP(), []int{3}
}

func (x *OriginalURLResponse) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

var File_urlshortener_proto protoreflect.FileDescriptor

var file_urlshortener_proto_rawDesc = []byte{
	0x0a, 0x12, 0x75, 0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x75, 0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x5f, 0x76, 0x33, 0x22, 0x37, 0x0a, 0x12, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61,
	0x6c, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x22, 0x2f,
	0x0a, 0x10, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x22,
	0x2e, 0x0a, 0x0f, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x22,
	0x38, 0x0a, 0x13, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x52, 0x4c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x32, 0xc4, 0x01, 0x0a, 0x0c, 0x55, 0x52,
	0x4c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x12, 0x59, 0x0a, 0x0f, 0x4f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x54, 0x6f, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x12, 0x23, 0x2e,
	0x75, 0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x5f, 0x76, 0x33, 0x2e,
	0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x75, 0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x5f, 0x76, 0x33, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x0f, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x54, 0x6f,
	0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x12, 0x20, 0x2e, 0x75, 0x72, 0x6c, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x5f, 0x76, 0x33, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x75, 0x72, 0x6c,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x5f, 0x76, 0x33, 0x2e, 0x4f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x25, 0x5a, 0x23, 0x70, 0x6b, 0x67, 0x2f, 0x75, 0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x5f, 0x76, 0x33, 0x3b, 0x75, 0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x5f, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_urlshortener_proto_rawDescOnce sync.Once
	file_urlshortener_proto_rawDescData = file_urlshortener_proto_rawDesc
)

func file_urlshortener_proto_rawDescGZIP() []byte {
	file_urlshortener_proto_rawDescOnce.Do(func() {
		file_urlshortener_proto_rawDescData = protoimpl.X.CompressGZIP(file_urlshortener_proto_rawDescData)
	})
	return file_urlshortener_proto_rawDescData
}

var file_urlshortener_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_urlshortener_proto_goTypes = []interface{}{
	(*OriginalURLRequest)(nil),  // 0: urlshortener_v3.OriginalURLRequest
	(*ShortURLResponse)(nil),    // 1: urlshortener_v3.ShortURLResponse
	(*ShortURLRequest)(nil),     // 2: urlshortener_v3.ShortURLRequest
	(*OriginalURLResponse)(nil), // 3: urlshortener_v3.OriginalURLResponse
}
var file_urlshortener_proto_depIdxs = []int32{
	0, // 0: urlshortener_v3.URLShortener.OriginalToShort:input_type -> urlshortener_v3.OriginalURLRequest
	2, // 1: urlshortener_v3.URLShortener.ShortToOriginal:input_type -> urlshortener_v3.ShortURLRequest
	1, // 2: urlshortener_v3.URLShortener.OriginalToShort:output_type -> urlshortener_v3.ShortURLResponse
	3, // 3: urlshortener_v3.URLShortener.ShortToOriginal:output_type -> urlshortener_v3.OriginalURLResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_urlshortener_proto_init() }
func file_urlshortener_proto_init() {
	if File_urlshortener_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_urlshortener_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OriginalURLRequest); i {
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
		file_urlshortener_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortURLResponse); i {
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
		file_urlshortener_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortURLRequest); i {
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
		file_urlshortener_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OriginalURLResponse); i {
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
			RawDescriptor: file_urlshortener_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_urlshortener_proto_goTypes,
		DependencyIndexes: file_urlshortener_proto_depIdxs,
		MessageInfos:      file_urlshortener_proto_msgTypes,
	}.Build()
	File_urlshortener_proto = out.File
	file_urlshortener_proto_rawDesc = nil
	file_urlshortener_proto_goTypes = nil
	file_urlshortener_proto_depIdxs = nil
}
