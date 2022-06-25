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
// source: k8s.io/api/batch/v1/generated.proto

package v1

import (
	v11 "github.com/onedata/k8s/apis/core/v1"
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

// Job represents the configuration of a single job.
type Job struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	Metadata *v1.ObjectMeta `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// Specification of the desired behavior of a job.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	// +optional
	Spec *JobSpec `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	// Current status of a job.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	// +optional
	Status *JobStatus `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
}

func (x *Job) Reset() {
	*x = Job{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_io_api_batch_v1_generated_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Job) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Job) ProtoMessage() {}

func (x *Job) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_io_api_batch_v1_generated_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Job.ProtoReflect.Descriptor instead.
func (*Job) Descriptor() ([]byte, []int) {
	return file_k8s_io_api_batch_v1_generated_proto_rawDescGZIP(), []int{0}
}

func (x *Job) GetMetadata() *v1.ObjectMeta {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Job) GetSpec() *JobSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Job) GetStatus() *JobStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

// JobCondition describes current state of a job.
type JobCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Type of job condition, Complete or Failed.
	Type *string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	// Status of the condition, one of True, False, Unknown.
	Status *string `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
	// Last time the condition was checked.
	// +optional
	LastProbeTime *v1.Time `protobuf:"bytes,3,opt,name=lastProbeTime" json:"lastProbeTime,omitempty"`
	// Last time the condition transit from one status to another.
	// +optional
	LastTransitionTime *v1.Time `protobuf:"bytes,4,opt,name=lastTransitionTime" json:"lastTransitionTime,omitempty"`
	// (brief) reason for the condition's last transition.
	// +optional
	Reason *string `protobuf:"bytes,5,opt,name=reason" json:"reason,omitempty"`
	// Human readable message indicating details about last transition.
	// +optional
	Message *string `protobuf:"bytes,6,opt,name=message" json:"message,omitempty"`
}

func (x *JobCondition) Reset() {
	*x = JobCondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_io_api_batch_v1_generated_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobCondition) ProtoMessage() {}

func (x *JobCondition) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_io_api_batch_v1_generated_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobCondition.ProtoReflect.Descriptor instead.
func (*JobCondition) Descriptor() ([]byte, []int) {
	return file_k8s_io_api_batch_v1_generated_proto_rawDescGZIP(), []int{1}
}

func (x *JobCondition) GetType() string {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return ""
}

func (x *JobCondition) GetStatus() string {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ""
}

func (x *JobCondition) GetLastProbeTime() *v1.Time {
	if x != nil {
		return x.LastProbeTime
	}
	return nil
}

func (x *JobCondition) GetLastTransitionTime() *v1.Time {
	if x != nil {
		return x.LastTransitionTime
	}
	return nil
}

func (x *JobCondition) GetReason() string {
	if x != nil && x.Reason != nil {
		return *x.Reason
	}
	return ""
}

func (x *JobCondition) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

// JobList is a collection of jobs.
type JobList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Standard list metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	Metadata *v1.ListMeta `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// items is the list of Jobs.
	Items []*Job `protobuf:"bytes,2,rep,name=items" json:"items,omitempty"`
}

func (x *JobList) Reset() {
	*x = JobList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_io_api_batch_v1_generated_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobList) ProtoMessage() {}

func (x *JobList) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_io_api_batch_v1_generated_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobList.ProtoReflect.Descriptor instead.
func (*JobList) Descriptor() ([]byte, []int) {
	return file_k8s_io_api_batch_v1_generated_proto_rawDescGZIP(), []int{2}
}

func (x *JobList) GetMetadata() *v1.ListMeta {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *JobList) GetItems() []*Job {
	if x != nil {
		return x.Items
	}
	return nil
}

// JobSpec describes how the job execution will look like.
type JobSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies the maximum desired number of pods the job should
	// run at any given time. The actual number of pods running in steady state will
	// be less than this number when ((.spec.completions - .status.successful) < .spec.parallelism),
	// i.e. when the work left to do is less than max parallelism.
	// More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
	// +optional
	Parallelism *int32 `protobuf:"varint,1,opt,name=parallelism" json:"parallelism,omitempty"`
	// Specifies the desired number of successfully finished pods the
	// job should be run with.  Setting to nil means that the success of any
	// pod signals the success of all pods, and allows parallelism to have any positive
	// value.  Setting to 1 means that parallelism is limited to 1 and the success of that
	// pod signals the success of the job.
	// More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
	// +optional
	Completions *int32 `protobuf:"varint,2,opt,name=completions" json:"completions,omitempty"`
	// Specifies the duration in seconds relative to the startTime that the job may be active
	// before the system tries to terminate it; value must be positive integer
	// +optional
	ActiveDeadlineSeconds *int64 `protobuf:"varint,3,opt,name=activeDeadlineSeconds" json:"activeDeadlineSeconds,omitempty"`
	// Specifies the number of retries before marking this job failed.
	// Defaults to 6
	// +optional
	BackoffLimit *int32 `protobuf:"varint,7,opt,name=backoffLimit" json:"backoffLimit,omitempty"`
	// A label query over pods that should match the pod count.
	// Normally, the system sets this field for you.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
	// +optional
	Selector *v1.LabelSelector `protobuf:"bytes,4,opt,name=selector" json:"selector,omitempty"`
	// manualSelector controls generation of pod labels and pod selectors.
	// Leave `manualSelector` unset unless you are certain what you are doing.
	// When false or unset, the system pick labels unique to this job
	// and appends those labels to the pod template.  When true,
	// the user is responsible for picking unique labels and specifying
	// the selector.  Failure to pick a unique label may cause this
	// and other jobs to not function correctly.  However, You may see
	// `manualSelector=true` in jobs that were created with the old `extensions/v1beta1`
	// API.
	// More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/#specifying-your-own-pod-selector
	// +optional
	ManualSelector *bool `protobuf:"varint,5,opt,name=manualSelector" json:"manualSelector,omitempty"`
	// Describes the pod that will be created when executing a job.
	// More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
	Template *v11.PodTemplateSpec `protobuf:"bytes,6,opt,name=template" json:"template,omitempty"`
	// ttlSecondsAfterFinished limits the lifetime of a Job that has finished
	// execution (either Complete or Failed). If this field is set,
	// ttlSecondsAfterFinished after the Job finishes, it is eligible to be
	// automatically deleted. When the Job is being deleted, its lifecycle
	// guarantees (e.g. finalizers) will be honored. If this field is unset,
	// the Job won't be automatically deleted. If this field is set to zero,
	// the Job becomes eligible to be deleted immediately after it finishes.
	// This field is alpha-level and is only honored by servers that enable the
	// TTLAfterFinished feature.
	// +optional
	TtlSecondsAfterFinished *int32 `protobuf:"varint,8,opt,name=ttlSecondsAfterFinished" json:"ttlSecondsAfterFinished,omitempty"`
}

func (x *JobSpec) Reset() {
	*x = JobSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_io_api_batch_v1_generated_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobSpec) ProtoMessage() {}

func (x *JobSpec) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_io_api_batch_v1_generated_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobSpec.ProtoReflect.Descriptor instead.
func (*JobSpec) Descriptor() ([]byte, []int) {
	return file_k8s_io_api_batch_v1_generated_proto_rawDescGZIP(), []int{3}
}

func (x *JobSpec) GetParallelism() int32 {
	if x != nil && x.Parallelism != nil {
		return *x.Parallelism
	}
	return 0
}

func (x *JobSpec) GetCompletions() int32 {
	if x != nil && x.Completions != nil {
		return *x.Completions
	}
	return 0
}

func (x *JobSpec) GetActiveDeadlineSeconds() int64 {
	if x != nil && x.ActiveDeadlineSeconds != nil {
		return *x.ActiveDeadlineSeconds
	}
	return 0
}

func (x *JobSpec) GetBackoffLimit() int32 {
	if x != nil && x.BackoffLimit != nil {
		return *x.BackoffLimit
	}
	return 0
}

func (x *JobSpec) GetSelector() *v1.LabelSelector {
	if x != nil {
		return x.Selector
	}
	return nil
}

func (x *JobSpec) GetManualSelector() bool {
	if x != nil && x.ManualSelector != nil {
		return *x.ManualSelector
	}
	return false
}

func (x *JobSpec) GetTemplate() *v11.PodTemplateSpec {
	if x != nil {
		return x.Template
	}
	return nil
}

func (x *JobSpec) GetTtlSecondsAfterFinished() int32 {
	if x != nil && x.TtlSecondsAfterFinished != nil {
		return *x.TtlSecondsAfterFinished
	}
	return 0
}

// JobStatus represents the current state of a Job.
type JobStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The latest available observations of an object's current state.
	// When a job fails, one of the conditions will have type == "Failed".
	// More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	Conditions []*JobCondition `protobuf:"bytes,1,rep,name=conditions" json:"conditions,omitempty"`
	// Represents time when the job was acknowledged by the job controller.
	// It is not guaranteed to be set in happens-before order across separate operations.
	// It is represented in RFC3339 form and is in UTC.
	// +optional
	StartTime *v1.Time `protobuf:"bytes,2,opt,name=startTime" json:"startTime,omitempty"`
	// Represents time when the job was completed. It is not guaranteed to
	// be set in happens-before order across separate operations.
	// It is represented in RFC3339 form and is in UTC.
	// The completion time is only set when the job finishes successfully.
	// +optional
	CompletionTime *v1.Time `protobuf:"bytes,3,opt,name=completionTime" json:"completionTime,omitempty"`
	// The number of actively running pods.
	// +optional
	Active *int32 `protobuf:"varint,4,opt,name=active" json:"active,omitempty"`
	// The number of pods which reached phase Succeeded.
	// +optional
	Succeeded *int32 `protobuf:"varint,5,opt,name=succeeded" json:"succeeded,omitempty"`
	// The number of pods which reached phase Failed.
	// +optional
	Failed *int32 `protobuf:"varint,6,opt,name=failed" json:"failed,omitempty"`
}

func (x *JobStatus) Reset() {
	*x = JobStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_io_api_batch_v1_generated_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobStatus) ProtoMessage() {}

func (x *JobStatus) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_io_api_batch_v1_generated_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobStatus.ProtoReflect.Descriptor instead.
func (*JobStatus) Descriptor() ([]byte, []int) {
	return file_k8s_io_api_batch_v1_generated_proto_rawDescGZIP(), []int{4}
}

func (x *JobStatus) GetConditions() []*JobCondition {
	if x != nil {
		return x.Conditions
	}
	return nil
}

func (x *JobStatus) GetStartTime() *v1.Time {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *JobStatus) GetCompletionTime() *v1.Time {
	if x != nil {
		return x.CompletionTime
	}
	return nil
}

func (x *JobStatus) GetActive() int32 {
	if x != nil && x.Active != nil {
		return *x.Active
	}
	return 0
}

func (x *JobStatus) GetSucceeded() int32 {
	if x != nil && x.Succeeded != nil {
		return *x.Succeeded
	}
	return 0
}

func (x *JobStatus) GetFailed() int32 {
	if x != nil && x.Failed != nil {
		return *x.Failed
	}
	return 0
}

var File_k8s_io_api_batch_v1_generated_proto protoreflect.FileDescriptor

var file_k8s_io_api_batch_v1_generated_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x61, 0x74,
	0x63, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x22, 0x6b, 0x38, 0x73, 0x2e,
	0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34,
	0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x72, 0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6d, 0x65, 0x74,
	0x61, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69,
	0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x75,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70,
	0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x72,
	0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x01,
	0x0a, 0x03, 0x4a, 0x6f, 0x62, 0x12, 0x4c, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f,
	0x2e, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x6b,
	0x67, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x30, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62,
	0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x70, 0x65, 0x63, 0x52,
	0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x9a, 0x02,
	0x0a, 0x0c, 0x4a, 0x6f, 0x62, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x50, 0x0a, 0x0d, 0x6c, 0x61,
	0x73, 0x74, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x6d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x0d, 0x6c,
	0x61, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x5a, 0x0a, 0x12,
	0x6c, 0x61, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69,
	0x6f, 0x2e, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2e, 0x70,
	0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x52, 0x12, 0x6c, 0x61, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x85, 0x01, 0x0a, 0x07, 0x4a,
	0x6f, 0x62, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69,
	0x6f, 0x2e, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2e, 0x70,
	0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x2e, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62,
	0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x22, 0x9b, 0x03, 0x0a, 0x07, 0x4a, 0x6f, 0x62, 0x53, 0x70, 0x65, 0x63, 0x12, 0x20,
	0x0a, 0x0b, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x6c, 0x65, 0x6c, 0x69, 0x73, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x6c, 0x65, 0x6c, 0x69, 0x73, 0x6d,
	0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x34, 0x0a, 0x15, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x44, 0x65, 0x61, 0x64,
	0x6c, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x15, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e,
	0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x62, 0x61, 0x63, 0x6b,
	0x6f, 0x66, 0x66, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x62, 0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x4f, 0x0a, 0x08,
	0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33,
	0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6d, 0x65,
	0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x26, 0x0a,
	0x0e, 0x6d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x6d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x3f, 0x0a, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x64,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x70, 0x65, 0x63, 0x52, 0x08, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x38, 0x0a, 0x17, 0x74, 0x74, 0x6c, 0x53, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x73, 0x41, 0x66, 0x74, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x17, 0x74, 0x74, 0x6c, 0x53, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x73, 0x41, 0x66, 0x74, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x22, 0xba, 0x02, 0x0a, 0x09, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x41,
	0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x43, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x48, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70,
	0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x52, 0x0a, 0x0e, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69,
	0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52,
	0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x65, 0x64, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x65, 0x64, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x42, 0x15, 0x5a,
	0x13, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x61, 0x74, 0x63,
	0x68, 0x2f, 0x76, 0x31,
}

var (
	file_k8s_io_api_batch_v1_generated_proto_rawDescOnce sync.Once
	file_k8s_io_api_batch_v1_generated_proto_rawDescData = file_k8s_io_api_batch_v1_generated_proto_rawDesc
)

func file_k8s_io_api_batch_v1_generated_proto_rawDescGZIP() []byte {
	file_k8s_io_api_batch_v1_generated_proto_rawDescOnce.Do(func() {
		file_k8s_io_api_batch_v1_generated_proto_rawDescData = protoimpl.X.CompressGZIP(file_k8s_io_api_batch_v1_generated_proto_rawDescData)
	})
	return file_k8s_io_api_batch_v1_generated_proto_rawDescData
}

var file_k8s_io_api_batch_v1_generated_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_k8s_io_api_batch_v1_generated_proto_goTypes = []interface{}{
	(*Job)(nil),                 // 0: k8s.io.api.batch.v1.Job
	(*JobCondition)(nil),        // 1: k8s.io.api.batch.v1.JobCondition
	(*JobList)(nil),             // 2: k8s.io.api.batch.v1.JobList
	(*JobSpec)(nil),             // 3: k8s.io.api.batch.v1.JobSpec
	(*JobStatus)(nil),           // 4: k8s.io.api.batch.v1.JobStatus
	(*v1.ObjectMeta)(nil),       // 5: k8s.io.apimachinery.pkg.apis.meta.v1.ObjectMeta
	(*v1.Time)(nil),             // 6: k8s.io.apimachinery.pkg.apis.meta.v1.Time
	(*v1.ListMeta)(nil),         // 7: k8s.io.apimachinery.pkg.apis.meta.v1.ListMeta
	(*v1.LabelSelector)(nil),    // 8: k8s.io.apimachinery.pkg.apis.meta.v1.LabelSelector
	(*v11.PodTemplateSpec)(nil), // 9: k8s.io.api.core.v1.PodTemplateSpec
}
var file_k8s_io_api_batch_v1_generated_proto_depIdxs = []int32{
	5,  // 0: k8s.io.api.batch.v1.Job.metadata:type_name -> k8s.io.apimachinery.pkg.apis.meta.v1.ObjectMeta
	3,  // 1: k8s.io.api.batch.v1.Job.spec:type_name -> k8s.io.api.batch.v1.JobSpec
	4,  // 2: k8s.io.api.batch.v1.Job.status:type_name -> k8s.io.api.batch.v1.JobStatus
	6,  // 3: k8s.io.api.batch.v1.JobCondition.lastProbeTime:type_name -> k8s.io.apimachinery.pkg.apis.meta.v1.Time
	6,  // 4: k8s.io.api.batch.v1.JobCondition.lastTransitionTime:type_name -> k8s.io.apimachinery.pkg.apis.meta.v1.Time
	7,  // 5: k8s.io.api.batch.v1.JobList.metadata:type_name -> k8s.io.apimachinery.pkg.apis.meta.v1.ListMeta
	0,  // 6: k8s.io.api.batch.v1.JobList.items:type_name -> k8s.io.api.batch.v1.Job
	8,  // 7: k8s.io.api.batch.v1.JobSpec.selector:type_name -> k8s.io.apimachinery.pkg.apis.meta.v1.LabelSelector
	9,  // 8: k8s.io.api.batch.v1.JobSpec.template:type_name -> k8s.io.api.core.v1.PodTemplateSpec
	1,  // 9: k8s.io.api.batch.v1.JobStatus.conditions:type_name -> k8s.io.api.batch.v1.JobCondition
	6,  // 10: k8s.io.api.batch.v1.JobStatus.startTime:type_name -> k8s.io.apimachinery.pkg.apis.meta.v1.Time
	6,  // 11: k8s.io.api.batch.v1.JobStatus.completionTime:type_name -> k8s.io.apimachinery.pkg.apis.meta.v1.Time
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_k8s_io_api_batch_v1_generated_proto_init() }
func file_k8s_io_api_batch_v1_generated_proto_init() {
	if File_k8s_io_api_batch_v1_generated_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_k8s_io_api_batch_v1_generated_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Job); i {
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
		file_k8s_io_api_batch_v1_generated_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobCondition); i {
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
		file_k8s_io_api_batch_v1_generated_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobList); i {
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
		file_k8s_io_api_batch_v1_generated_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobSpec); i {
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
		file_k8s_io_api_batch_v1_generated_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobStatus); i {
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
			RawDescriptor: file_k8s_io_api_batch_v1_generated_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_k8s_io_api_batch_v1_generated_proto_goTypes,
		DependencyIndexes: file_k8s_io_api_batch_v1_generated_proto_depIdxs,
		MessageInfos:      file_k8s_io_api_batch_v1_generated_proto_msgTypes,
	}.Build()
	File_k8s_io_api_batch_v1_generated_proto = out.File
	file_k8s_io_api_batch_v1_generated_proto_rawDesc = nil
	file_k8s_io_api_batch_v1_generated_proto_goTypes = nil
	file_k8s_io_api_batch_v1_generated_proto_depIdxs = nil
}
