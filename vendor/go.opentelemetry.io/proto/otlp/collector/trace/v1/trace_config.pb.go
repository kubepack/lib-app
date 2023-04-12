// Copyright 2019, OpenTelemetry Authors
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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: opentelemetry/proto/trace/v1/trace_config.proto

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

// How spans should be sampled:
// - Always off
// - Always on
// - Always follow the parent Span's decision (off if no parent).
type ConstantSampler_ConstantDecision int32

const (
	ConstantSampler_ALWAYS_OFF    ConstantSampler_ConstantDecision = 0
	ConstantSampler_ALWAYS_ON     ConstantSampler_ConstantDecision = 1
	ConstantSampler_ALWAYS_PARENT ConstantSampler_ConstantDecision = 2
)

// Enum value maps for ConstantSampler_ConstantDecision.
var (
	ConstantSampler_ConstantDecision_name = map[int32]string{
		0: "ALWAYS_OFF",
		1: "ALWAYS_ON",
		2: "ALWAYS_PARENT",
	}
	ConstantSampler_ConstantDecision_value = map[string]int32{
		"ALWAYS_OFF":    0,
		"ALWAYS_ON":     1,
		"ALWAYS_PARENT": 2,
	}
)

func (x ConstantSampler_ConstantDecision) Enum() *ConstantSampler_ConstantDecision {
	p := new(ConstantSampler_ConstantDecision)
	*p = x
	return p
}

func (x ConstantSampler_ConstantDecision) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConstantSampler_ConstantDecision) Descriptor() protoreflect.EnumDescriptor {
	return file_opentelemetry_proto_trace_v1_trace_config_proto_enumTypes[0].Descriptor()
}

func (ConstantSampler_ConstantDecision) Type() protoreflect.EnumType {
	return &file_opentelemetry_proto_trace_v1_trace_config_proto_enumTypes[0]
}

func (x ConstantSampler_ConstantDecision) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConstantSampler_ConstantDecision.Descriptor instead.
func (ConstantSampler_ConstantDecision) EnumDescriptor() ([]byte, []int) {
	return file_opentelemetry_proto_trace_v1_trace_config_proto_rawDescGZIP(), []int{1, 0}
}

// Global configuration of the trace service. All fields must be specified, or
// the default (zero) values will be used for each type.
type TraceConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The global default sampler used to make decisions on span sampling.
	//
	// Types that are assignable to Sampler:
	//	*TraceConfig_ConstantSampler
	//	*TraceConfig_TraceIdRatioBased
	//	*TraceConfig_RateLimitingSampler
	Sampler isTraceConfig_Sampler `protobuf_oneof:"sampler"`
	// The global default max number of attributes per span.
	MaxNumberOfAttributes int64 `protobuf:"varint,4,opt,name=max_number_of_attributes,json=maxNumberOfAttributes,proto3" json:"max_number_of_attributes,omitempty"`
	// The global default max number of annotation events per span.
	MaxNumberOfTimedEvents int64 `protobuf:"varint,5,opt,name=max_number_of_timed_events,json=maxNumberOfTimedEvents,proto3" json:"max_number_of_timed_events,omitempty"`
	// The global default max number of attributes per timed event.
	MaxNumberOfAttributesPerTimedEvent int64 `protobuf:"varint,6,opt,name=max_number_of_attributes_per_timed_event,json=maxNumberOfAttributesPerTimedEvent,proto3" json:"max_number_of_attributes_per_timed_event,omitempty"`
	// The global default max number of link entries per span.
	MaxNumberOfLinks int64 `protobuf:"varint,7,opt,name=max_number_of_links,json=maxNumberOfLinks,proto3" json:"max_number_of_links,omitempty"`
	// The global default max number of attributes per span.
	MaxNumberOfAttributesPerLink int64 `protobuf:"varint,8,opt,name=max_number_of_attributes_per_link,json=maxNumberOfAttributesPerLink,proto3" json:"max_number_of_attributes_per_link,omitempty"`
}

func (x *TraceConfig) Reset() {
	*x = TraceConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opentelemetry_proto_trace_v1_trace_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TraceConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TraceConfig) ProtoMessage() {}

func (x *TraceConfig) ProtoReflect() protoreflect.Message {
	mi := &file_opentelemetry_proto_trace_v1_trace_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TraceConfig.ProtoReflect.Descriptor instead.
func (*TraceConfig) Descriptor() ([]byte, []int) {
	return file_opentelemetry_proto_trace_v1_trace_config_proto_rawDescGZIP(), []int{0}
}

func (m *TraceConfig) GetSampler() isTraceConfig_Sampler {
	if m != nil {
		return m.Sampler
	}
	return nil
}

func (x *TraceConfig) GetConstantSampler() *ConstantSampler {
	if x, ok := x.GetSampler().(*TraceConfig_ConstantSampler); ok {
		return x.ConstantSampler
	}
	return nil
}

func (x *TraceConfig) GetTraceIdRatioBased() *TraceIdRatioBased {
	if x, ok := x.GetSampler().(*TraceConfig_TraceIdRatioBased); ok {
		return x.TraceIdRatioBased
	}
	return nil
}

func (x *TraceConfig) GetRateLimitingSampler() *RateLimitingSampler {
	if x, ok := x.GetSampler().(*TraceConfig_RateLimitingSampler); ok {
		return x.RateLimitingSampler
	}
	return nil
}

func (x *TraceConfig) GetMaxNumberOfAttributes() int64 {
	if x != nil {
		return x.MaxNumberOfAttributes
	}
	return 0
}

func (x *TraceConfig) GetMaxNumberOfTimedEvents() int64 {
	if x != nil {
		return x.MaxNumberOfTimedEvents
	}
	return 0
}

func (x *TraceConfig) GetMaxNumberOfAttributesPerTimedEvent() int64 {
	if x != nil {
		return x.MaxNumberOfAttributesPerTimedEvent
	}
	return 0
}

func (x *TraceConfig) GetMaxNumberOfLinks() int64 {
	if x != nil {
		return x.MaxNumberOfLinks
	}
	return 0
}

func (x *TraceConfig) GetMaxNumberOfAttributesPerLink() int64 {
	if x != nil {
		return x.MaxNumberOfAttributesPerLink
	}
	return 0
}

type isTraceConfig_Sampler interface {
	isTraceConfig_Sampler()
}

type TraceConfig_ConstantSampler struct {
	ConstantSampler *ConstantSampler `protobuf:"bytes,1,opt,name=constant_sampler,json=constantSampler,proto3,oneof"`
}

type TraceConfig_TraceIdRatioBased struct {
	TraceIdRatioBased *TraceIdRatioBased `protobuf:"bytes,2,opt,name=trace_id_ratio_based,json=traceIdRatioBased,proto3,oneof"`
}

type TraceConfig_RateLimitingSampler struct {
	RateLimitingSampler *RateLimitingSampler `protobuf:"bytes,3,opt,name=rate_limiting_sampler,json=rateLimitingSampler,proto3,oneof"`
}

func (*TraceConfig_ConstantSampler) isTraceConfig_Sampler() {}

func (*TraceConfig_TraceIdRatioBased) isTraceConfig_Sampler() {}

func (*TraceConfig_RateLimitingSampler) isTraceConfig_Sampler() {}

// Sampler that always makes a constant decision on span sampling.
type ConstantSampler struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Decision ConstantSampler_ConstantDecision `protobuf:"varint,1,opt,name=decision,proto3,enum=opentelemetry.proto.trace.v1.ConstantSampler_ConstantDecision" json:"decision,omitempty"`
}

func (x *ConstantSampler) Reset() {
	*x = ConstantSampler{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opentelemetry_proto_trace_v1_trace_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConstantSampler) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConstantSampler) ProtoMessage() {}

func (x *ConstantSampler) ProtoReflect() protoreflect.Message {
	mi := &file_opentelemetry_proto_trace_v1_trace_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConstantSampler.ProtoReflect.Descriptor instead.
func (*ConstantSampler) Descriptor() ([]byte, []int) {
	return file_opentelemetry_proto_trace_v1_trace_config_proto_rawDescGZIP(), []int{1}
}

func (x *ConstantSampler) GetDecision() ConstantSampler_ConstantDecision {
	if x != nil {
		return x.Decision
	}
	return ConstantSampler_ALWAYS_OFF
}

// Sampler that tries to uniformly sample traces with a given ratio.
// The ratio of sampling a trace is equal to that of the specified ratio.
type TraceIdRatioBased struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The desired ratio of sampling. Must be within [0.0, 1.0].
	SamplingRatio float64 `protobuf:"fixed64,1,opt,name=samplingRatio,proto3" json:"samplingRatio,omitempty"`
}

func (x *TraceIdRatioBased) Reset() {
	*x = TraceIdRatioBased{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opentelemetry_proto_trace_v1_trace_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TraceIdRatioBased) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TraceIdRatioBased) ProtoMessage() {}

func (x *TraceIdRatioBased) ProtoReflect() protoreflect.Message {
	mi := &file_opentelemetry_proto_trace_v1_trace_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TraceIdRatioBased.ProtoReflect.Descriptor instead.
func (*TraceIdRatioBased) Descriptor() ([]byte, []int) {
	return file_opentelemetry_proto_trace_v1_trace_config_proto_rawDescGZIP(), []int{2}
}

func (x *TraceIdRatioBased) GetSamplingRatio() float64 {
	if x != nil {
		return x.SamplingRatio
	}
	return 0
}

// Sampler that tries to sample with a rate per time window.
type RateLimitingSampler struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Rate per second.
	Qps int64 `protobuf:"varint,1,opt,name=qps,proto3" json:"qps,omitempty"`
}

func (x *RateLimitingSampler) Reset() {
	*x = RateLimitingSampler{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opentelemetry_proto_trace_v1_trace_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimitingSampler) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimitingSampler) ProtoMessage() {}

func (x *RateLimitingSampler) ProtoReflect() protoreflect.Message {
	mi := &file_opentelemetry_proto_trace_v1_trace_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimitingSampler.ProtoReflect.Descriptor instead.
func (*RateLimitingSampler) Descriptor() ([]byte, []int) {
	return file_opentelemetry_proto_trace_v1_trace_config_proto_rawDescGZIP(), []int{3}
}

func (x *RateLimitingSampler) GetQps() int64 {
	if x != nil {
		return x.Qps
	}
	return 0
}

var File_opentelemetry_proto_trace_v1_trace_config_proto protoreflect.FileDescriptor

var file_opentelemetry_proto_trace_v1_trace_config_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1c, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x22,
	0x84, 0x05, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x5a, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x5f, 0x73, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x6f, 0x70, 0x65, 0x6e,
	0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x74, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x72, 0x48, 0x00, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x74, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x72, 0x12, 0x62, 0x0a, 0x14, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x5f, 0x62, 0x61,
	0x73, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x6f, 0x70, 0x65, 0x6e,
	0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64,
	0x52, 0x61, 0x74, 0x69, 0x6f, 0x42, 0x61, 0x73, 0x65, 0x64, 0x48, 0x00, 0x52, 0x11, 0x74, 0x72,
	0x61, 0x63, 0x65, 0x49, 0x64, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x42, 0x61, 0x73, 0x65, 0x64, 0x12,
	0x67, 0x0a, 0x15, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x6e, 0x67,
	0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31,
	0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61,
	0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x72, 0x48, 0x00, 0x52, 0x13, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x6e,
	0x67, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x18, 0x6d, 0x61, 0x78, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x15, 0x6d, 0x61, 0x78, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x12, 0x3a, 0x0a, 0x1a, 0x6d, 0x61, 0x78, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f,
	0x6f, 0x66, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x64, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x16, 0x6d, 0x61, 0x78, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x4f, 0x66, 0x54, 0x69, 0x6d, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x54, 0x0a,
	0x28, 0x6d, 0x61, 0x78, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x61,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x64, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x22, 0x6d, 0x61, 0x78, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x50, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x64, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x2d, 0x0a, 0x13, 0x6d, 0x61, 0x78, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x10, 0x6d, 0x61, 0x78, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x4c, 0x69, 0x6e,
	0x6b, 0x73, 0x12, 0x47, 0x0a, 0x21, 0x6d, 0x61, 0x78, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x5f, 0x6f, 0x66, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x5f, 0x70,
	0x65, 0x72, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x1c, 0x6d,
	0x61, 0x78, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x50, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x42, 0x09, 0x0a, 0x07, 0x73,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x72, 0x22, 0xb3, 0x01, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x74, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x72, 0x12, 0x5a, 0x0a, 0x08, 0x64, 0x65,
	0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3e, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x74, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x74, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x64, 0x65,
	0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x44, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x4c,
	0x57, 0x41, 0x59, 0x53, 0x5f, 0x4f, 0x46, 0x46, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x4c,
	0x57, 0x41, 0x59, 0x53, 0x5f, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x4c, 0x57,
	0x41, 0x59, 0x53, 0x5f, 0x50, 0x41, 0x52, 0x45, 0x4e, 0x54, 0x10, 0x02, 0x22, 0x39, 0x0a, 0x11,
	0x54, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x42, 0x61, 0x73, 0x65,
	0x64, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x52, 0x61, 0x74,
	0x69, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69,
	0x6e, 0x67, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x22, 0x27, 0x0a, 0x13, 0x52, 0x61, 0x74, 0x65, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x72, 0x12, 0x10,
	0x0a, 0x03, 0x71, 0x70, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x71, 0x70, 0x73,
	0x42, 0x68, 0x0a, 0x1f, 0x69, 0x6f, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d,
	0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x42, 0x10, 0x54, 0x72, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x31, 0x67, 0x6f, 0x2e, 0x6f, 0x70, 0x65, 0x6e,
	0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x69, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x6f, 0x74, 0x6c, 0x70, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_opentelemetry_proto_trace_v1_trace_config_proto_rawDescOnce sync.Once
	file_opentelemetry_proto_trace_v1_trace_config_proto_rawDescData = file_opentelemetry_proto_trace_v1_trace_config_proto_rawDesc
)

func file_opentelemetry_proto_trace_v1_trace_config_proto_rawDescGZIP() []byte {
	file_opentelemetry_proto_trace_v1_trace_config_proto_rawDescOnce.Do(func() {
		file_opentelemetry_proto_trace_v1_trace_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_opentelemetry_proto_trace_v1_trace_config_proto_rawDescData)
	})
	return file_opentelemetry_proto_trace_v1_trace_config_proto_rawDescData
}

var file_opentelemetry_proto_trace_v1_trace_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_opentelemetry_proto_trace_v1_trace_config_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_opentelemetry_proto_trace_v1_trace_config_proto_goTypes = []interface{}{
	(ConstantSampler_ConstantDecision)(0), // 0: opentelemetry.proto.trace.v1.ConstantSampler.ConstantDecision
	(*TraceConfig)(nil),                   // 1: opentelemetry.proto.trace.v1.TraceConfig
	(*ConstantSampler)(nil),               // 2: opentelemetry.proto.trace.v1.ConstantSampler
	(*TraceIdRatioBased)(nil),             // 3: opentelemetry.proto.trace.v1.TraceIdRatioBased
	(*RateLimitingSampler)(nil),           // 4: opentelemetry.proto.trace.v1.RateLimitingSampler
}
var file_opentelemetry_proto_trace_v1_trace_config_proto_depIdxs = []int32{
	2, // 0: opentelemetry.proto.trace.v1.TraceConfig.constant_sampler:type_name -> opentelemetry.proto.trace.v1.ConstantSampler
	3, // 1: opentelemetry.proto.trace.v1.TraceConfig.trace_id_ratio_based:type_name -> opentelemetry.proto.trace.v1.TraceIdRatioBased
	4, // 2: opentelemetry.proto.trace.v1.TraceConfig.rate_limiting_sampler:type_name -> opentelemetry.proto.trace.v1.RateLimitingSampler
	0, // 3: opentelemetry.proto.trace.v1.ConstantSampler.decision:type_name -> opentelemetry.proto.trace.v1.ConstantSampler.ConstantDecision
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_opentelemetry_proto_trace_v1_trace_config_proto_init() }
func file_opentelemetry_proto_trace_v1_trace_config_proto_init() {
	if File_opentelemetry_proto_trace_v1_trace_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_opentelemetry_proto_trace_v1_trace_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TraceConfig); i {
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
		file_opentelemetry_proto_trace_v1_trace_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConstantSampler); i {
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
		file_opentelemetry_proto_trace_v1_trace_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TraceIdRatioBased); i {
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
		file_opentelemetry_proto_trace_v1_trace_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLimitingSampler); i {
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
	file_opentelemetry_proto_trace_v1_trace_config_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*TraceConfig_ConstantSampler)(nil),
		(*TraceConfig_TraceIdRatioBased)(nil),
		(*TraceConfig_RateLimitingSampler)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_opentelemetry_proto_trace_v1_trace_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_opentelemetry_proto_trace_v1_trace_config_proto_goTypes,
		DependencyIndexes: file_opentelemetry_proto_trace_v1_trace_config_proto_depIdxs,
		EnumInfos:         file_opentelemetry_proto_trace_v1_trace_config_proto_enumTypes,
		MessageInfos:      file_opentelemetry_proto_trace_v1_trace_config_proto_msgTypes,
	}.Build()
	File_opentelemetry_proto_trace_v1_trace_config_proto = out.File
	file_opentelemetry_proto_trace_v1_trace_config_proto_rawDesc = nil
	file_opentelemetry_proto_trace_v1_trace_config_proto_goTypes = nil
	file_opentelemetry_proto_trace_v1_trace_config_proto_depIdxs = nil
}
