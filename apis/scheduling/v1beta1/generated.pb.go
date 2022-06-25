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
// source: k8s.io/api/scheduling/v1beta1/generated.proto

package v1beta1

import (
	_ "github.com/onedata/k8s/apis/core/v1"
	"github.com/onedata/k8s/apis/meta/v1"
	_ "github.com/onedata/k8s/runtime"
	_ "github.com/onedata/k8s/runtime/schema"
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

// DEPRECATED - This group version of PriorityClass is deprecated by scheduling.k8s.io/v1/PriorityClass.
// PriorityClass defines mapping from a priority class name to the priority
// integer value. The value can be any valid integer.
type PriorityClass struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	Metadata *v1.ObjectMeta `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// The value of this priority class. This is the actual priority that pods
	// receive when they have the name of this class in their pod spec.
	Value *int32 `protobuf:"varint,2,opt,name=value" json:"value,omitempty"`
	// globalDefault specifies whether this PriorityClass should be considered as
	// the default priority for pods that do not have any priority class.
	// Only one PriorityClass can be marked as `globalDefault`. However, if more than
	// one PriorityClasses exists with their `globalDefault` field set to true,
	// the smallest value of such global default PriorityClasses will be used as the default priority.
	// +optional
	GlobalDefault *bool `protobuf:"varint,3,opt,name=globalDefault" json:"globalDefault,omitempty"`
	// description is an arbitrary string that usually provides guidelines on
	// when this priority class should be used.
	// +optional
	Description *string `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	// PreemptionPolicy is the Policy for preempting pods with lower priority.
	// One of Never, PreemptLowerPriority.
	// Defaults to PreemptLowerPriority if unset.
	// This field is beta-level, gated by the NonPreemptingPriority feature-gate.
	// +optional
	PreemptionPolicy *string `protobuf:"bytes,5,opt,name=preemptionPolicy" json:"preemptionPolicy,omitempty"`
}

func (x *PriorityClass) Reset() {
	*x = PriorityClass{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_io_api_scheduling_v1beta1_generated_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PriorityClass) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PriorityClass) ProtoMessage() {}

func (x *PriorityClass) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_io_api_scheduling_v1beta1_generated_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PriorityClass.ProtoReflect.Descriptor instead.
func (*PriorityClass) Descriptor() ([]byte, []int) {
	return file_k8s_io_api_scheduling_v1beta1_generated_proto_rawDescGZIP(), []int{0}
}

func (x *PriorityClass) GetMetadata() *v1.ObjectMeta {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *PriorityClass) GetValue() int32 {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return 0
}

func (x *PriorityClass) GetGlobalDefault() bool {
	if x != nil && x.GlobalDefault != nil {
		return *x.GlobalDefault
	}
	return false
}

func (x *PriorityClass) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *PriorityClass) GetPreemptionPolicy() string {
	if x != nil && x.PreemptionPolicy != nil {
		return *x.PreemptionPolicy
	}
	return ""
}

// PriorityClassList is a collection of priority classes.
type PriorityClassList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Standard list metadata
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	Metadata *v1.ListMeta `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// items is the list of PriorityClasses
	Items []*PriorityClass `protobuf:"bytes,2,rep,name=items" json:"items,omitempty"`
}

func (x *PriorityClassList) Reset() {
	*x = PriorityClassList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_io_api_scheduling_v1beta1_generated_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PriorityClassList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PriorityClassList) ProtoMessage() {}

func (x *PriorityClassList) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_io_api_scheduling_v1beta1_generated_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PriorityClassList.ProtoReflect.Descriptor instead.
func (*PriorityClassList) Descriptor() ([]byte, []int) {
	return file_k8s_io_api_scheduling_v1beta1_generated_proto_rawDescGZIP(), []int{1}
}

func (x *PriorityClassList) GetMetadata() *v1.ListMeta {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *PriorityClassList) GetItems() []*PriorityClass {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_k8s_io_api_scheduling_v1beta1_generated_proto protoreflect.FileDescriptor

var file_k8s_io_api_scheduling_v1beta1_generated_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1d, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x22,
	0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x34, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x6d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f,
	0x2f, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x6b, 0x38, 0x73, 0x2e, 0x69,
	0x6f, 0x2f, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xe7, 0x01, 0x0a, 0x0d, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x12, 0x4c, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61,
	0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x6b, 0x67, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x67, 0x6c, 0x6f, 0x62, 0x61,
	0x6c, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d,
	0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2a, 0x0a, 0x10, 0x70, 0x72, 0x65, 0x65, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x72, 0x65, 0x65, 0x6d,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0xa3, 0x01, 0x0a, 0x11,
	0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x4a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69,
	0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d,
	0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x42, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6b,
	0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x69,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x42, 0x1f, 0x5a, 0x1d, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31,
}

var (
	file_k8s_io_api_scheduling_v1beta1_generated_proto_rawDescOnce sync.Once
	file_k8s_io_api_scheduling_v1beta1_generated_proto_rawDescData = file_k8s_io_api_scheduling_v1beta1_generated_proto_rawDesc
)

func file_k8s_io_api_scheduling_v1beta1_generated_proto_rawDescGZIP() []byte {
	file_k8s_io_api_scheduling_v1beta1_generated_proto_rawDescOnce.Do(func() {
		file_k8s_io_api_scheduling_v1beta1_generated_proto_rawDescData = protoimpl.X.CompressGZIP(file_k8s_io_api_scheduling_v1beta1_generated_proto_rawDescData)
	})
	return file_k8s_io_api_scheduling_v1beta1_generated_proto_rawDescData
}

var file_k8s_io_api_scheduling_v1beta1_generated_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_k8s_io_api_scheduling_v1beta1_generated_proto_goTypes = []interface{}{
	(*PriorityClass)(nil),     // 0: k8s.io.api.scheduling.v1beta1.PriorityClass
	(*PriorityClassList)(nil), // 1: k8s.io.api.scheduling.v1beta1.PriorityClassList
	(*v1.ObjectMeta)(nil),     // 2: k8s.io.apimachinery.pkg.apis.meta.v1.ObjectMeta
	(*v1.ListMeta)(nil),       // 3: k8s.io.apimachinery.pkg.apis.meta.v1.ListMeta
}
var file_k8s_io_api_scheduling_v1beta1_generated_proto_depIdxs = []int32{
	2, // 0: k8s.io.api.scheduling.v1beta1.PriorityClass.metadata:type_name -> k8s.io.apimachinery.pkg.apis.meta.v1.ObjectMeta
	3, // 1: k8s.io.api.scheduling.v1beta1.PriorityClassList.metadata:type_name -> k8s.io.apimachinery.pkg.apis.meta.v1.ListMeta
	0, // 2: k8s.io.api.scheduling.v1beta1.PriorityClassList.items:type_name -> k8s.io.api.scheduling.v1beta1.PriorityClass
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_k8s_io_api_scheduling_v1beta1_generated_proto_init() }
func file_k8s_io_api_scheduling_v1beta1_generated_proto_init() {
	if File_k8s_io_api_scheduling_v1beta1_generated_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_k8s_io_api_scheduling_v1beta1_generated_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PriorityClass); i {
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
		file_k8s_io_api_scheduling_v1beta1_generated_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PriorityClassList); i {
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
			RawDescriptor: file_k8s_io_api_scheduling_v1beta1_generated_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_k8s_io_api_scheduling_v1beta1_generated_proto_goTypes,
		DependencyIndexes: file_k8s_io_api_scheduling_v1beta1_generated_proto_depIdxs,
		MessageInfos:      file_k8s_io_api_scheduling_v1beta1_generated_proto_msgTypes,
	}.Build()
	File_k8s_io_api_scheduling_v1beta1_generated_proto = out.File
	file_k8s_io_api_scheduling_v1beta1_generated_proto_rawDesc = nil
	file_k8s_io_api_scheduling_v1beta1_generated_proto_goTypes = nil
	file_k8s_io_api_scheduling_v1beta1_generated_proto_depIdxs = nil
}
