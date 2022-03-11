//
//Copyright 2021 The tKeel Authors.
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: api/subscribe/v1/error.proto

package v1

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

// @plugins=protoc-gen-go-errors
// 错误
type Error int32

const (
	// @msg=未知类型
	// @code=UNKNOWN
	Error_ERR_UNKNOWN Error = 0
	// @msg=未找到目标资源
	// @code=NOT_FOUND
	Error_ERR_NOT_FOUND Error = 1
	// @msg=获取列表信息失败
	// @code=INTERNAL
	Error_ERR_LIST Error = 2
	// @msg=请求参数无效
	// @code=INVALID_ARGUMENT
	Error_ERR_INVALID_ARGUMENT Error = 3
	// @msg=服务器内部错误
	// @code=INTERNAL
	Error_ERR_INTERNAL_ERROR Error = 4
	// @msg=订阅已存在
	// @code=INVALID_ARGUMENT
	Error_ERR_EXIST Error = 5
	// @msg=请确保您对资源拥有相应权限
	// @code=UNAUTHENTICATED
	Error_ERR_UNAUTHENTICATED Error = 6
	// @msg=内部资源获取失败
	// @code=INTERNAL
	Error_ERR_INTERNAL_QUERY Error = 7
	// @msg=未找到设备
	// @code=OK
	Error_ERR_DEVICE_NOT_FOUND Error = 8
	// @msg=必须字段未填充有效值
	// @code=INVALID_ARGUMENT
	Error_ERR_INVALID_ARGUMENT_SOME_FIELDS Error = 9
	// @msg=重复添加订阅
	// @code=INVALID_ARGUMENT
	Error_ERR_DUPLICATE_CREATE Error = 10
	// @msg=存在重复订阅
	// @code=OK
	Error_ERR_SOME__DUPLICATE_CREATE Error = 11
	// @msg=默认订阅不可删除
	// @code=INVALID_ARGUMENT
	Error_ERR_TRY_TO_DELETE_DEFAULT_SUBSCRIBE Error = 12
)

// Enum value maps for Error.
var (
	Error_name = map[int32]string{
		0:  "ERR_UNKNOWN",
		1:  "ERR_NOT_FOUND",
		2:  "ERR_LIST",
		3:  "ERR_INVALID_ARGUMENT",
		4:  "ERR_INTERNAL_ERROR",
		5:  "ERR_EXIST",
		6:  "ERR_UNAUTHENTICATED",
		7:  "ERR_INTERNAL_QUERY",
		8:  "ERR_DEVICE_NOT_FOUND",
		9:  "ERR_INVALID_ARGUMENT_SOME_FIELDS",
		10: "ERR_DUPLICATE_CREATE",
		11: "ERR_SOME__DUPLICATE_CREATE",
		12: "ERR_TRY_TO_DELETE_DEFAULT_SUBSCRIBE",
	}
	Error_value = map[string]int32{
		"ERR_UNKNOWN":                         0,
		"ERR_NOT_FOUND":                       1,
		"ERR_LIST":                            2,
		"ERR_INVALID_ARGUMENT":                3,
		"ERR_INTERNAL_ERROR":                  4,
		"ERR_EXIST":                           5,
		"ERR_UNAUTHENTICATED":                 6,
		"ERR_INTERNAL_QUERY":                  7,
		"ERR_DEVICE_NOT_FOUND":                8,
		"ERR_INVALID_ARGUMENT_SOME_FIELDS":    9,
		"ERR_DUPLICATE_CREATE":                10,
		"ERR_SOME__DUPLICATE_CREATE":          11,
		"ERR_TRY_TO_DELETE_DEFAULT_SUBSCRIBE": 12,
	}
)

func (x Error) Enum() *Error {
	p := new(Error)
	*p = x
	return p
}

func (x Error) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Error) Descriptor() protoreflect.EnumDescriptor {
	return file_api_subscribe_v1_error_proto_enumTypes[0].Descriptor()
}

func (Error) Type() protoreflect.EnumType {
	return &file_api_subscribe_v1_error_proto_enumTypes[0]
}

func (x Error) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Error.Descriptor instead.
func (Error) EnumDescriptor() ([]byte, []int) {
	return file_api_subscribe_v1_error_proto_rawDescGZIP(), []int{0}
}

var File_api_subscribe_v1_error_proto protoreflect.FileDescriptor

var file_api_subscribe_v1_error_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d,
	0x69, 0x6f, 0x2e, 0x74, 0x6b, 0x65, 0x65, 0x6c, 0x2e, 0x72, 0x75, 0x64, 0x64, 0x65, 0x72, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2a, 0xce, 0x02,
	0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x52, 0x52, 0x5f, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x45, 0x52, 0x52, 0x5f,
	0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x45,
	0x52, 0x52, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x52, 0x52,
	0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e,
	0x54, 0x10, 0x03, 0x12, 0x16, 0x0a, 0x12, 0x45, 0x52, 0x52, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52,
	0x4e, 0x41, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x45,
	0x52, 0x52, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x10, 0x05, 0x12, 0x17, 0x0a, 0x13, 0x45, 0x52,
	0x52, 0x5f, 0x55, 0x4e, 0x41, 0x55, 0x54, 0x48, 0x45, 0x4e, 0x54, 0x49, 0x43, 0x41, 0x54, 0x45,
	0x44, 0x10, 0x06, 0x12, 0x16, 0x0a, 0x12, 0x45, 0x52, 0x52, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52,
	0x4e, 0x41, 0x4c, 0x5f, 0x51, 0x55, 0x45, 0x52, 0x59, 0x10, 0x07, 0x12, 0x18, 0x0a, 0x14, 0x45,
	0x52, 0x52, 0x5f, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f,
	0x55, 0x4e, 0x44, 0x10, 0x08, 0x12, 0x24, 0x0a, 0x20, 0x45, 0x52, 0x52, 0x5f, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x4f,
	0x4d, 0x45, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x53, 0x10, 0x09, 0x12, 0x18, 0x0a, 0x14, 0x45,
	0x52, 0x52, 0x5f, 0x44, 0x55, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x52, 0x45,
	0x41, 0x54, 0x45, 0x10, 0x0a, 0x12, 0x1e, 0x0a, 0x1a, 0x45, 0x52, 0x52, 0x5f, 0x53, 0x4f, 0x4d,
	0x45, 0x5f, 0x5f, 0x44, 0x55, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x52, 0x45,
	0x41, 0x54, 0x45, 0x10, 0x0b, 0x12, 0x27, 0x0a, 0x23, 0x45, 0x52, 0x52, 0x5f, 0x54, 0x52, 0x59,
	0x5f, 0x54, 0x4f, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x44, 0x45, 0x46, 0x41, 0x55,
	0x4c, 0x54, 0x5f, 0x53, 0x55, 0x42, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x10, 0x0c, 0x42, 0x49,
	0x0a, 0x10, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x2e,
	0x76, 0x31, 0x50, 0x01, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x74, 0x6b, 0x65, 0x65, 0x6c, 0x2d, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2d, 0x62,
	0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_subscribe_v1_error_proto_rawDescOnce sync.Once
	file_api_subscribe_v1_error_proto_rawDescData = file_api_subscribe_v1_error_proto_rawDesc
)

func file_api_subscribe_v1_error_proto_rawDescGZIP() []byte {
	file_api_subscribe_v1_error_proto_rawDescOnce.Do(func() {
		file_api_subscribe_v1_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_subscribe_v1_error_proto_rawDescData)
	})
	return file_api_subscribe_v1_error_proto_rawDescData
}

var file_api_subscribe_v1_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_subscribe_v1_error_proto_goTypes = []interface{}{
	(Error)(0), // 0: io.tkeel.rudder.api.config.v1.Error
}
var file_api_subscribe_v1_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_subscribe_v1_error_proto_init() }
func file_api_subscribe_v1_error_proto_init() {
	if File_api_subscribe_v1_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_subscribe_v1_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_subscribe_v1_error_proto_goTypes,
		DependencyIndexes: file_api_subscribe_v1_error_proto_depIdxs,
		EnumInfos:         file_api_subscribe_v1_error_proto_enumTypes,
	}.Build()
	File_api_subscribe_v1_error_proto = out.File
	file_api_subscribe_v1_error_proto_rawDesc = nil
	file_api_subscribe_v1_error_proto_goTypes = nil
	file_api_subscribe_v1_error_proto_depIdxs = nil
}
