// Copyright 2015 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: tinyurl.proto

package generated

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ShortURL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ShortURL) Reset() {
	*x = ShortURL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tinyurl_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortURL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortURL) ProtoMessage() {}

func (x *ShortURL) ProtoReflect() protoreflect.Message {
	mi := &file_tinyurl_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortURL.ProtoReflect.Descriptor instead.
func (*ShortURL) Descriptor() ([]byte, []int) {
	return file_tinyurl_proto_rawDescGZIP(), []int{0}
}

func (x *ShortURL) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type OriginalURL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *OriginalURL) Reset() {
	*x = OriginalURL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tinyurl_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OriginalURL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OriginalURL) ProtoMessage() {}

func (x *OriginalURL) ProtoReflect() protoreflect.Message {
	mi := &file_tinyurl_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OriginalURL.ProtoReflect.Descriptor instead.
func (*OriginalURL) Descriptor() ([]byte, []int) {
	return file_tinyurl_proto_rawDescGZIP(), []int{1}
}

func (x *OriginalURL) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type RespGetOriginalURLs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrl    string `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
	OriginalUrl string `protobuf:"bytes,2,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
}

func (x *RespGetOriginalURLs) Reset() {
	*x = RespGetOriginalURLs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tinyurl_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespGetOriginalURLs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespGetOriginalURLs) ProtoMessage() {}

func (x *RespGetOriginalURLs) ProtoReflect() protoreflect.Message {
	mi := &file_tinyurl_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespGetOriginalURLs.ProtoReflect.Descriptor instead.
func (*RespGetOriginalURLs) Descriptor() ([]byte, []int) {
	return file_tinyurl_proto_rawDescGZIP(), []int{2}
}

func (x *RespGetOriginalURLs) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

func (x *RespGetOriginalURLs) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

var File_tinyurl_proto protoreflect.FileDescriptor

var file_tinyurl_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x74, 0x69, 0x6e, 0x79, 0x75, 0x72, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x61, 0x70, 0x69, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x1a, 0x0a, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1f, 0x0a,
	0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x52, 0x4c, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x55,
	0x0a, 0x13, 0x72, 0x65, 0x73, 0x70, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61,
	0x6c, 0x55, 0x52, 0x4c, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55,
	0x72, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x61, 0x6c, 0x55, 0x72, 0x6c, 0x32, 0xbc, 0x01, 0x0a, 0x0a, 0x61, 0x70, 0x69, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x33,
	0x0a, 0x0e, 0x67, 0x65, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x52, 0x4c,
	0x12, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x1a,
	0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x52,
	0x4c, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x09, 0x66, 0x65, 0x74, 0x63, 0x68, 0x55, 0x52, 0x4c, 0x73,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72,
	0x65, 0x73, 0x70, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x52,
	0x4c, 0x73, 0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x73, 0x68, 0x61, 0x31, 0x30, 0x38, 0x62, 0x69, 0x6b, 0x2f, 0x74,
	0x69, 0x6e, 0x79, 0x5f, 0x75, 0x72, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tinyurl_proto_rawDescOnce sync.Once
	file_tinyurl_proto_rawDescData = file_tinyurl_proto_rawDesc
)

func file_tinyurl_proto_rawDescGZIP() []byte {
	file_tinyurl_proto_rawDescOnce.Do(func() {
		file_tinyurl_proto_rawDescData = protoimpl.X.CompressGZIP(file_tinyurl_proto_rawDescData)
	})
	return file_tinyurl_proto_rawDescData
}

var file_tinyurl_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tinyurl_proto_goTypes = []interface{}{
	(*ShortURL)(nil),            // 0: api.shortURL
	(*OriginalURL)(nil),         // 1: api.originalURL
	(*RespGetOriginalURLs)(nil), // 2: api.respGetOriginalURLs
	(*emptypb.Empty)(nil),       // 3: google.protobuf.Empty
}
var file_tinyurl_proto_depIdxs = []int32{
	3, // 0: api.apiService.ping:input_type -> google.protobuf.Empty
	0, // 1: api.apiService.getOriginalURL:input_type -> api.shortURL
	3, // 2: api.apiService.fetchURLs:input_type -> google.protobuf.Empty
	3, // 3: api.apiService.ping:output_type -> google.protobuf.Empty
	1, // 4: api.apiService.getOriginalURL:output_type -> api.originalURL
	2, // 5: api.apiService.fetchURLs:output_type -> api.respGetOriginalURLs
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tinyurl_proto_init() }
func file_tinyurl_proto_init() {
	if File_tinyurl_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tinyurl_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortURL); i {
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
		file_tinyurl_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OriginalURL); i {
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
		file_tinyurl_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespGetOriginalURLs); i {
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
			RawDescriptor: file_tinyurl_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tinyurl_proto_goTypes,
		DependencyIndexes: file_tinyurl_proto_depIdxs,
		MessageInfos:      file_tinyurl_proto_msgTypes,
	}.Build()
	File_tinyurl_proto = out.File
	file_tinyurl_proto_rawDesc = nil
	file_tinyurl_proto_goTypes = nil
	file_tinyurl_proto_depIdxs = nil
}
