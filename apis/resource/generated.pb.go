//
//Copyright The Kubernetes Authors.
//
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

// This file was autogenerated by go-to-protobuf. Do not edit it manually!

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: k8s.io/apimachinery/pkg/api/resource/generated.proto

package resource

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

// Quantity is a fixed-point representation of a number.
// It provides convenient marshaling/unmarshaling in JSON and YAML,
// in addition to String() and AsInt64() accessors.
//
// The serialization format is:
//
// <quantity>        ::= <signedNumber><suffix>
//   (Note that <suffix> may be empty, from the "" case in <decimalSI>.)
// <digit>           ::= 0 | 1 | ... | 9
// <digits>          ::= <digit> | <digit><digits>
// <number>          ::= <digits> | <digits>.<digits> | <digits>. | .<digits>
// <sign>            ::= "+" | "-"
// <signedNumber>    ::= <number> | <sign><number>
// <suffix>          ::= <binarySI> | <decimalExponent> | <decimalSI>
// <binarySI>        ::= Ki | Mi | Gi | Ti | Pi | Ei
//   (International System of units; See: http://physics.nist.gov/cuu/Units/binary.html)
// <decimalSI>       ::= m | "" | k | M | G | T | P | E
//   (Note that 1024 = 1Ki but 1000 = 1k; I didn't choose the capitalization.)
// <decimalExponent> ::= "e" <signedNumber> | "E" <signedNumber>
//
// No matter which of the three exponent forms is used, no quantity may represent
// a number greater than 2^63-1 in magnitude, nor may it have more than 3 decimal
// places. Numbers larger or more precise will be capped or rounded up.
// (E.g.: 0.1m will rounded up to 1m.)
// This may be extended in the future if we require larger or smaller quantities.
//
// When a Quantity is parsed from a string, it will remember the type of suffix
// it had, and will use the same type again when it is serialized.
//
// Before serializing, Quantity will be put in "canonical form".
// This means that Exponent/suffix will be adjusted up or down (with a
// corresponding increase or decrease in Mantissa) such that:
//   a. No precision is lost
//   b. No fractional digits will be emitted
//   c. The exponent (or suffix) is as large as possible.
// The sign will be omitted unless the number is negative.
//
// Examples:
//   1.5 will be serialized as "1500m"
//   1.5Gi will be serialized as "1536Mi"
//
// Note that the quantity will NEVER be internally represented by a
// floating point number. That is the whole point of this exercise.
//
// Non-canonical values will still parse as long as they are well formed,
// but will be re-emitted in their canonical form. (So always use canonical
// form, or don't diff.)
//
// This format is intended to make it difficult to use these numbers without
// writing some sort of special handling code in the hopes that that will
// cause implementors to also use a fixed point implementation.
//
// +protobuf=true
// +protobuf.embed=string
// +protobuf.options.marshal=false
// +protobuf.options.(gogoproto.goproto_stringer)=false
// +k8s:deepcopy-gen=true
// +k8s:openapi-gen=true
type Quantity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	String_ *string `protobuf:"bytes,1,opt,name=string" json:"string,omitempty"`
}

func (x *Quantity) Reset() {
	*x = Quantity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_io_apimachinery_pkg_api_resource_generated_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Quantity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Quantity) ProtoMessage() {}

func (x *Quantity) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_io_apimachinery_pkg_api_resource_generated_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Quantity.ProtoReflect.Descriptor instead.
func (*Quantity) Descriptor() ([]byte, []int) {
	return file_k8s_io_apimachinery_pkg_api_resource_generated_proto_rawDescGZIP(), []int{0}
}

func (x *Quantity) GetString_() string {
	if x != nil && x.String_ != nil {
		return *x.String_
	}
	return ""
}

var File_k8s_io_apimachinery_pkg_api_resource_generated_proto protoreflect.FileDescriptor

var file_k8s_io_apimachinery_pkg_api_resource_generated_proto_rawDesc = []byte{
	0x0a, 0x34, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61,
	0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x6b, 0x67, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x22, 0x0a, 0x08,
	0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x42, 0x26, 0x5a, 0x24, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x6d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
}

var (
	file_k8s_io_apimachinery_pkg_api_resource_generated_proto_rawDescOnce sync.Once
	file_k8s_io_apimachinery_pkg_api_resource_generated_proto_rawDescData = file_k8s_io_apimachinery_pkg_api_resource_generated_proto_rawDesc
)

func file_k8s_io_apimachinery_pkg_api_resource_generated_proto_rawDescGZIP() []byte {
	file_k8s_io_apimachinery_pkg_api_resource_generated_proto_rawDescOnce.Do(func() {
		file_k8s_io_apimachinery_pkg_api_resource_generated_proto_rawDescData = protoimpl.X.CompressGZIP(file_k8s_io_apimachinery_pkg_api_resource_generated_proto_rawDescData)
	})
	return file_k8s_io_apimachinery_pkg_api_resource_generated_proto_rawDescData
}

var file_k8s_io_apimachinery_pkg_api_resource_generated_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_k8s_io_apimachinery_pkg_api_resource_generated_proto_goTypes = []interface{}{
	(*Quantity)(nil), // 0: k8s.io.apimachinery.pkg.api.resource.Quantity
}
var file_k8s_io_apimachinery_pkg_api_resource_generated_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_k8s_io_apimachinery_pkg_api_resource_generated_proto_init() }
func file_k8s_io_apimachinery_pkg_api_resource_generated_proto_init() {
	if File_k8s_io_apimachinery_pkg_api_resource_generated_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_k8s_io_apimachinery_pkg_api_resource_generated_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Quantity); i {
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
			RawDescriptor: file_k8s_io_apimachinery_pkg_api_resource_generated_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_k8s_io_apimachinery_pkg_api_resource_generated_proto_goTypes,
		DependencyIndexes: file_k8s_io_apimachinery_pkg_api_resource_generated_proto_depIdxs,
		MessageInfos:      file_k8s_io_apimachinery_pkg_api_resource_generated_proto_msgTypes,
	}.Build()
	File_k8s_io_apimachinery_pkg_api_resource_generated_proto = out.File
	file_k8s_io_apimachinery_pkg_api_resource_generated_proto_rawDesc = nil
	file_k8s_io_apimachinery_pkg_api_resource_generated_proto_goTypes = nil
	file_k8s_io_apimachinery_pkg_api_resource_generated_proto_depIdxs = nil
}
