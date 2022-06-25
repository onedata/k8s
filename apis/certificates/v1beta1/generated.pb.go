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
// source: k8s.io/api/certificates/v1beta1/generated.proto

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

// Describes a certificate signing request
type CertificateSigningRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// +optional
	Metadata *v1.ObjectMeta `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// The certificate request itself and any additional information.
	// +optional
	Spec *CertificateSigningRequestSpec `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	// Derived information about the request.
	// +optional
	Status *CertificateSigningRequestStatus `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
}

func (x *CertificateSigningRequest) Reset() {
	*x = CertificateSigningRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_io_api_certificates_v1beta1_generated_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertificateSigningRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertificateSigningRequest) ProtoMessage() {}

func (x *CertificateSigningRequest) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_io_api_certificates_v1beta1_generated_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertificateSigningRequest.ProtoReflect.Descriptor instead.
func (*CertificateSigningRequest) Descriptor() ([]byte, []int) {
	return file_k8s_io_api_certificates_v1beta1_generated_proto_rawDescGZIP(), []int{0}
}

func (x *CertificateSigningRequest) GetMetadata() *v1.ObjectMeta {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *CertificateSigningRequest) GetSpec() *CertificateSigningRequestSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *CertificateSigningRequest) GetStatus() *CertificateSigningRequestStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type CertificateSigningRequestCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// type of the condition. Known conditions include "Approved", "Denied", and "Failed".
	Type *string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	// Status of the condition, one of True, False, Unknown.
	// Approved, Denied, and Failed conditions may not be "False" or "Unknown".
	// Defaults to "True".
	// If unset, should be treated as "True".
	// +optional
	Status *string `protobuf:"bytes,6,opt,name=status" json:"status,omitempty"`
	// brief reason for the request state
	// +optional
	Reason *string `protobuf:"bytes,2,opt,name=reason" json:"reason,omitempty"`
	// human readable message with details about the request state
	// +optional
	Message *string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	// timestamp for the last update to this condition
	// +optional
	LastUpdateTime *v1.Time `protobuf:"bytes,4,opt,name=lastUpdateTime" json:"lastUpdateTime,omitempty"`
	// lastTransitionTime is the time the condition last transitioned from one status to another.
	// If unset, when a new condition type is added or an existing condition's status is changed,
	// the server defaults this to the current time.
	// +optional
	LastTransitionTime *v1.Time `protobuf:"bytes,5,opt,name=lastTransitionTime" json:"lastTransitionTime,omitempty"`
}

func (x *CertificateSigningRequestCondition) Reset() {
	*x = CertificateSigningRequestCondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_io_api_certificates_v1beta1_generated_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertificateSigningRequestCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertificateSigningRequestCondition) ProtoMessage() {}

func (x *CertificateSigningRequestCondition) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_io_api_certificates_v1beta1_generated_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertificateSigningRequestCondition.ProtoReflect.Descriptor instead.
func (*CertificateSigningRequestCondition) Descriptor() ([]byte, []int) {
	return file_k8s_io_api_certificates_v1beta1_generated_proto_rawDescGZIP(), []int{1}
}

func (x *CertificateSigningRequestCondition) GetType() string {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return ""
}

func (x *CertificateSigningRequestCondition) GetStatus() string {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ""
}

func (x *CertificateSigningRequestCondition) GetReason() string {
	if x != nil && x.Reason != nil {
		return *x.Reason
	}
	return ""
}

func (x *CertificateSigningRequestCondition) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

func (x *CertificateSigningRequestCondition) GetLastUpdateTime() *v1.Time {
	if x != nil {
		return x.LastUpdateTime
	}
	return nil
}

func (x *CertificateSigningRequestCondition) GetLastTransitionTime() *v1.Time {
	if x != nil {
		return x.LastTransitionTime
	}
	return nil
}

type CertificateSigningRequestList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// +optional
	Metadata *v1.ListMeta                 `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	Items    []*CertificateSigningRequest `protobuf:"bytes,2,rep,name=items" json:"items,omitempty"`
}

func (x *CertificateSigningRequestList) Reset() {
	*x = CertificateSigningRequestList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_io_api_certificates_v1beta1_generated_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertificateSigningRequestList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertificateSigningRequestList) ProtoMessage() {}

func (x *CertificateSigningRequestList) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_io_api_certificates_v1beta1_generated_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertificateSigningRequestList.ProtoReflect.Descriptor instead.
func (*CertificateSigningRequestList) Descriptor() ([]byte, []int) {
	return file_k8s_io_api_certificates_v1beta1_generated_proto_rawDescGZIP(), []int{2}
}

func (x *CertificateSigningRequestList) GetMetadata() *v1.ListMeta {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *CertificateSigningRequestList) GetItems() []*CertificateSigningRequest {
	if x != nil {
		return x.Items
	}
	return nil
}

// This information is immutable after the request is created. Only the Request
// and Usages fields can be set on creation, other fields are derived by
// Kubernetes and cannot be modified by users.
type CertificateSigningRequestSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Base64-encoded PKCS#10 CSR data
	// +listType=atomic
	Request []byte `protobuf:"bytes,1,opt,name=request" json:"request,omitempty"`
	// Requested signer for the request. It is a qualified name in the form:
	// `scope-hostname.io/name`.
	// If empty, it will be defaulted:
	//  1. If it's a kubelet client certificate, it is assigned
	//     "kubernetes.io/kube-apiserver-client-kubelet".
	//  2. If it's a kubelet serving certificate, it is assigned
	//     "kubernetes.io/kubelet-serving".
	//  3. Otherwise, it is assigned "kubernetes.io/legacy-unknown".
	// Distribution of trust for signers happens out of band.
	// You can select on this field using `spec.signerName`.
	// +optional
	SignerName *string `protobuf:"bytes,7,opt,name=signerName" json:"signerName,omitempty"`
	// allowedUsages specifies a set of usage contexts the key will be
	// valid for.
	// See: https://tools.ietf.org/html/rfc5280#section-4.2.1.3
	//      https://tools.ietf.org/html/rfc5280#section-4.2.1.12
	// Valid values are:
	//  "signing",
	//  "digital signature",
	//  "content commitment",
	//  "key encipherment",
	//  "key agreement",
	//  "data encipherment",
	//  "cert sign",
	//  "crl sign",
	//  "encipher only",
	//  "decipher only",
	//  "any",
	//  "server auth",
	//  "client auth",
	//  "code signing",
	//  "email protection",
	//  "s/mime",
	//  "ipsec end system",
	//  "ipsec tunnel",
	//  "ipsec user",
	//  "timestamping",
	//  "ocsp signing",
	//  "microsoft sgc",
	//  "netscape sgc"
	// +listType=atomic
	Usages []string `protobuf:"bytes,5,rep,name=usages" json:"usages,omitempty"`
	// Information about the requesting user.
	// See user.Info interface for details.
	// +optional
	Username *string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	// UID information about the requesting user.
	// See user.Info interface for details.
	// +optional
	Uid *string `protobuf:"bytes,3,opt,name=uid" json:"uid,omitempty"`
	// Group information about the requesting user.
	// See user.Info interface for details.
	// +listType=atomic
	// +optional
	Groups []string `protobuf:"bytes,4,rep,name=groups" json:"groups,omitempty"`
	// Extra information about the requesting user.
	// See user.Info interface for details.
	// +optional
	Extra map[string]*ExtraValue `protobuf:"bytes,6,rep,name=extra" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (x *CertificateSigningRequestSpec) Reset() {
	*x = CertificateSigningRequestSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_io_api_certificates_v1beta1_generated_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertificateSigningRequestSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertificateSigningRequestSpec) ProtoMessage() {}

func (x *CertificateSigningRequestSpec) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_io_api_certificates_v1beta1_generated_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertificateSigningRequestSpec.ProtoReflect.Descriptor instead.
func (*CertificateSigningRequestSpec) Descriptor() ([]byte, []int) {
	return file_k8s_io_api_certificates_v1beta1_generated_proto_rawDescGZIP(), []int{3}
}

func (x *CertificateSigningRequestSpec) GetRequest() []byte {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *CertificateSigningRequestSpec) GetSignerName() string {
	if x != nil && x.SignerName != nil {
		return *x.SignerName
	}
	return ""
}

func (x *CertificateSigningRequestSpec) GetUsages() []string {
	if x != nil {
		return x.Usages
	}
	return nil
}

func (x *CertificateSigningRequestSpec) GetUsername() string {
	if x != nil && x.Username != nil {
		return *x.Username
	}
	return ""
}

func (x *CertificateSigningRequestSpec) GetUid() string {
	if x != nil && x.Uid != nil {
		return *x.Uid
	}
	return ""
}

func (x *CertificateSigningRequestSpec) GetGroups() []string {
	if x != nil {
		return x.Groups
	}
	return nil
}

func (x *CertificateSigningRequestSpec) GetExtra() map[string]*ExtraValue {
	if x != nil {
		return x.Extra
	}
	return nil
}

type CertificateSigningRequestStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Conditions applied to the request, such as approval or denial.
	// +listType=map
	// +listMapKey=type
	// +optional
	Conditions []*CertificateSigningRequestCondition `protobuf:"bytes,1,rep,name=conditions" json:"conditions,omitempty"`
	// If request was approved, the controller will place the issued certificate here.
	// +listType=atomic
	// +optional
	Certificate []byte `protobuf:"bytes,2,opt,name=certificate" json:"certificate,omitempty"`
}

func (x *CertificateSigningRequestStatus) Reset() {
	*x = CertificateSigningRequestStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_io_api_certificates_v1beta1_generated_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertificateSigningRequestStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertificateSigningRequestStatus) ProtoMessage() {}

func (x *CertificateSigningRequestStatus) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_io_api_certificates_v1beta1_generated_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertificateSigningRequestStatus.ProtoReflect.Descriptor instead.
func (*CertificateSigningRequestStatus) Descriptor() ([]byte, []int) {
	return file_k8s_io_api_certificates_v1beta1_generated_proto_rawDescGZIP(), []int{4}
}

func (x *CertificateSigningRequestStatus) GetConditions() []*CertificateSigningRequestCondition {
	if x != nil {
		return x.Conditions
	}
	return nil
}

func (x *CertificateSigningRequestStatus) GetCertificate() []byte {
	if x != nil {
		return x.Certificate
	}
	return nil
}

// ExtraValue masks the value so protobuf can generate
// +protobuf.nullable=true
// +protobuf.options.(gogoproto.goproto_stringer)=false
type ExtraValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []string `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (x *ExtraValue) Reset() {
	*x = ExtraValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_io_api_certificates_v1beta1_generated_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtraValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtraValue) ProtoMessage() {}

func (x *ExtraValue) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_io_api_certificates_v1beta1_generated_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtraValue.ProtoReflect.Descriptor instead.
func (*ExtraValue) Descriptor() ([]byte, []int) {
	return file_k8s_io_api_certificates_v1beta1_generated_proto_rawDescGZIP(), []int{5}
}

func (x *ExtraValue) GetItems() []string {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_k8s_io_api_certificates_v1beta1_generated_proto protoreflect.FileDescriptor

var file_k8s_io_api_certificates_v1beta1_generated_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1f, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x1a, 0x22, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x6b, 0x38,
	0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72,
	0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x6b,
	0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x72, 0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x02, 0x0a, 0x19, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x4c, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61,
	0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x6b, 0x67, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x52, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x3e, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67,
	0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x70, 0x65, 0x63, 0x52,
	0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x58, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0xb2, 0x02, 0x0a, 0x22, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x53,
	0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x52, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6b,
	0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x72, 0x79, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x5a, 0x0a, 0x12, 0x6c, 0x61, 0x73, 0x74,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70,
	0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x52, 0x12, 0x6c, 0x61, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x22, 0xbd, 0x01, 0x0a, 0x1d, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69,
	0x6f, 0x2e, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2e, 0x70,
	0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x50, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x3a, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x53,
	0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x22, 0xff, 0x02, 0x0a, 0x1d, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x53, 0x70, 0x65, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x5f,
	0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x49, 0x2e,
	0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x45, 0x78,
	0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x1a,
	0x65, 0x0a, 0x0a, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x41, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b,
	0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xa8, 0x01, 0x0a, 0x1f, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x63, 0x0a, 0x0a, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x43,
	0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67, 0x6e,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x20, 0x0a, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x22, 0x22, 0x0a, 0x0a, 0x45, 0x78, 0x74, 0x72, 0x61, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x21, 0x5a, 0x1f, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
}

var (
	file_k8s_io_api_certificates_v1beta1_generated_proto_rawDescOnce sync.Once
	file_k8s_io_api_certificates_v1beta1_generated_proto_rawDescData = file_k8s_io_api_certificates_v1beta1_generated_proto_rawDesc
)

func file_k8s_io_api_certificates_v1beta1_generated_proto_rawDescGZIP() []byte {
	file_k8s_io_api_certificates_v1beta1_generated_proto_rawDescOnce.Do(func() {
		file_k8s_io_api_certificates_v1beta1_generated_proto_rawDescData = protoimpl.X.CompressGZIP(file_k8s_io_api_certificates_v1beta1_generated_proto_rawDescData)
	})
	return file_k8s_io_api_certificates_v1beta1_generated_proto_rawDescData
}

var file_k8s_io_api_certificates_v1beta1_generated_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_k8s_io_api_certificates_v1beta1_generated_proto_goTypes = []interface{}{
	(*CertificateSigningRequest)(nil),          // 0: k8s.io.api.certificates.v1beta1.CertificateSigningRequest
	(*CertificateSigningRequestCondition)(nil), // 1: k8s.io.api.certificates.v1beta1.CertificateSigningRequestCondition
	(*CertificateSigningRequestList)(nil),      // 2: k8s.io.api.certificates.v1beta1.CertificateSigningRequestList
	(*CertificateSigningRequestSpec)(nil),      // 3: k8s.io.api.certificates.v1beta1.CertificateSigningRequestSpec
	(*CertificateSigningRequestStatus)(nil),    // 4: k8s.io.api.certificates.v1beta1.CertificateSigningRequestStatus
	(*ExtraValue)(nil),                         // 5: k8s.io.api.certificates.v1beta1.ExtraValue
	nil,                                        // 6: k8s.io.api.certificates.v1beta1.CertificateSigningRequestSpec.ExtraEntry
	(*v1.ObjectMeta)(nil),                      // 7: k8s.io.apimachinery.pkg.apis.meta.v1.ObjectMeta
	(*v1.Time)(nil),                            // 8: k8s.io.apimachinery.pkg.apis.meta.v1.Time
	(*v1.ListMeta)(nil),                        // 9: k8s.io.apimachinery.pkg.apis.meta.v1.ListMeta
}
var file_k8s_io_api_certificates_v1beta1_generated_proto_depIdxs = []int32{
	7,  // 0: k8s.io.api.certificates.v1beta1.CertificateSigningRequest.metadata:type_name -> k8s.io.apimachinery.pkg.apis.meta.v1.ObjectMeta
	3,  // 1: k8s.io.api.certificates.v1beta1.CertificateSigningRequest.spec:type_name -> k8s.io.api.certificates.v1beta1.CertificateSigningRequestSpec
	4,  // 2: k8s.io.api.certificates.v1beta1.CertificateSigningRequest.status:type_name -> k8s.io.api.certificates.v1beta1.CertificateSigningRequestStatus
	8,  // 3: k8s.io.api.certificates.v1beta1.CertificateSigningRequestCondition.lastUpdateTime:type_name -> k8s.io.apimachinery.pkg.apis.meta.v1.Time
	8,  // 4: k8s.io.api.certificates.v1beta1.CertificateSigningRequestCondition.lastTransitionTime:type_name -> k8s.io.apimachinery.pkg.apis.meta.v1.Time
	9,  // 5: k8s.io.api.certificates.v1beta1.CertificateSigningRequestList.metadata:type_name -> k8s.io.apimachinery.pkg.apis.meta.v1.ListMeta
	0,  // 6: k8s.io.api.certificates.v1beta1.CertificateSigningRequestList.items:type_name -> k8s.io.api.certificates.v1beta1.CertificateSigningRequest
	6,  // 7: k8s.io.api.certificates.v1beta1.CertificateSigningRequestSpec.extra:type_name -> k8s.io.api.certificates.v1beta1.CertificateSigningRequestSpec.ExtraEntry
	1,  // 8: k8s.io.api.certificates.v1beta1.CertificateSigningRequestStatus.conditions:type_name -> k8s.io.api.certificates.v1beta1.CertificateSigningRequestCondition
	5,  // 9: k8s.io.api.certificates.v1beta1.CertificateSigningRequestSpec.ExtraEntry.value:type_name -> k8s.io.api.certificates.v1beta1.ExtraValue
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_k8s_io_api_certificates_v1beta1_generated_proto_init() }
func file_k8s_io_api_certificates_v1beta1_generated_proto_init() {
	if File_k8s_io_api_certificates_v1beta1_generated_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_k8s_io_api_certificates_v1beta1_generated_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertificateSigningRequest); i {
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
		file_k8s_io_api_certificates_v1beta1_generated_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertificateSigningRequestCondition); i {
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
		file_k8s_io_api_certificates_v1beta1_generated_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertificateSigningRequestList); i {
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
		file_k8s_io_api_certificates_v1beta1_generated_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertificateSigningRequestSpec); i {
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
		file_k8s_io_api_certificates_v1beta1_generated_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertificateSigningRequestStatus); i {
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
		file_k8s_io_api_certificates_v1beta1_generated_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtraValue); i {
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
			RawDescriptor: file_k8s_io_api_certificates_v1beta1_generated_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_k8s_io_api_certificates_v1beta1_generated_proto_goTypes,
		DependencyIndexes: file_k8s_io_api_certificates_v1beta1_generated_proto_depIdxs,
		MessageInfos:      file_k8s_io_api_certificates_v1beta1_generated_proto_msgTypes,
	}.Build()
	File_k8s_io_api_certificates_v1beta1_generated_proto = out.File
	file_k8s_io_api_certificates_v1beta1_generated_proto_rawDesc = nil
	file_k8s_io_api_certificates_v1beta1_generated_proto_goTypes = nil
	file_k8s_io_api_certificates_v1beta1_generated_proto_depIdxs = nil
}
