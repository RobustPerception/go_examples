// Code generated by protoc-gen-go.
// source: remote.proto
// DO NOT EDIT!

/*
Package remote is a generated protocol buffer package.

It is generated from these files:
	remote.proto

It has these top-level messages:
	Sample
	LabelPair
	TimeSeries
	WriteRequest
	ReadRequest
	ReadResponse
	Query
	LabelMatcher
	QueryResponse
*/
package remote

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MatchType int32

const (
	MatchType_EQUAL          MatchType = 0
	MatchType_NOT_EQUAL      MatchType = 1
	MatchType_REGEX_MATCH    MatchType = 2
	MatchType_REGEX_NO_MATCH MatchType = 3
)

var MatchType_name = map[int32]string{
	0: "EQUAL",
	1: "NOT_EQUAL",
	2: "REGEX_MATCH",
	3: "REGEX_NO_MATCH",
}
var MatchType_value = map[string]int32{
	"EQUAL":          0,
	"NOT_EQUAL":      1,
	"REGEX_MATCH":    2,
	"REGEX_NO_MATCH": 3,
}

func (x MatchType) String() string {
	return proto.EnumName(MatchType_name, int32(x))
}
func (MatchType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Sample struct {
	Value       float64 `protobuf:"fixed64,1,opt,name=value" json:"value,omitempty"`
	TimestampMs int64   `protobuf:"varint,2,opt,name=timestamp_ms,json=timestampMs" json:"timestamp_ms,omitempty"`
}

func (m *Sample) Reset()                    { *m = Sample{} }
func (m *Sample) String() string            { return proto.CompactTextString(m) }
func (*Sample) ProtoMessage()               {}
func (*Sample) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Sample) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Sample) GetTimestampMs() int64 {
	if m != nil {
		return m.TimestampMs
	}
	return 0
}

type LabelPair struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *LabelPair) Reset()                    { *m = LabelPair{} }
func (m *LabelPair) String() string            { return proto.CompactTextString(m) }
func (*LabelPair) ProtoMessage()               {}
func (*LabelPair) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LabelPair) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LabelPair) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type TimeSeries struct {
	Labels []*LabelPair `protobuf:"bytes,1,rep,name=labels" json:"labels,omitempty"`
	// Sorted by time, oldest sample first.
	Samples []*Sample `protobuf:"bytes,2,rep,name=samples" json:"samples,omitempty"`
}

func (m *TimeSeries) Reset()                    { *m = TimeSeries{} }
func (m *TimeSeries) String() string            { return proto.CompactTextString(m) }
func (*TimeSeries) ProtoMessage()               {}
func (*TimeSeries) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TimeSeries) GetLabels() []*LabelPair {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *TimeSeries) GetSamples() []*Sample {
	if m != nil {
		return m.Samples
	}
	return nil
}

type WriteRequest struct {
	Timeseries []*TimeSeries `protobuf:"bytes,1,rep,name=timeseries" json:"timeseries,omitempty"`
}

func (m *WriteRequest) Reset()                    { *m = WriteRequest{} }
func (m *WriteRequest) String() string            { return proto.CompactTextString(m) }
func (*WriteRequest) ProtoMessage()               {}
func (*WriteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *WriteRequest) GetTimeseries() []*TimeSeries {
	if m != nil {
		return m.Timeseries
	}
	return nil
}

type ReadRequest struct {
	Queries []*Query `protobuf:"bytes,1,rep,name=queries" json:"queries,omitempty"`
}

func (m *ReadRequest) Reset()                    { *m = ReadRequest{} }
func (m *ReadRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()               {}
func (*ReadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ReadRequest) GetQueries() []*Query {
	if m != nil {
		return m.Queries
	}
	return nil
}

type ReadResponse struct {
	// In same order as the request's queries.
	Responses []*QueryResponse `protobuf:"bytes,1,rep,name=responses" json:"responses,omitempty"`
}

func (m *ReadResponse) Reset()                    { *m = ReadResponse{} }
func (m *ReadResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()               {}
func (*ReadResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ReadResponse) GetResponses() []*QueryResponse {
	if m != nil {
		return m.Responses
	}
	return nil
}

type Query struct {
	StartTimestampMs int64           `protobuf:"varint,1,opt,name=start_timestamp_ms,json=startTimestampMs" json:"start_timestamp_ms,omitempty"`
	EndTimestampMs   int64           `protobuf:"varint,2,opt,name=end_timestamp_ms,json=endTimestampMs" json:"end_timestamp_ms,omitempty"`
	Matchers         []*LabelMatcher `protobuf:"bytes,3,rep,name=matchers" json:"matchers,omitempty"`
}

func (m *Query) Reset()                    { *m = Query{} }
func (m *Query) String() string            { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()               {}
func (*Query) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Query) GetStartTimestampMs() int64 {
	if m != nil {
		return m.StartTimestampMs
	}
	return 0
}

func (m *Query) GetEndTimestampMs() int64 {
	if m != nil {
		return m.EndTimestampMs
	}
	return 0
}

func (m *Query) GetMatchers() []*LabelMatcher {
	if m != nil {
		return m.Matchers
	}
	return nil
}

type LabelMatcher struct {
	Type  MatchType `protobuf:"varint,1,opt,name=type,enum=remote.MatchType" json:"type,omitempty"`
	Name  string    `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Value string    `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
}

func (m *LabelMatcher) Reset()                    { *m = LabelMatcher{} }
func (m *LabelMatcher) String() string            { return proto.CompactTextString(m) }
func (*LabelMatcher) ProtoMessage()               {}
func (*LabelMatcher) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *LabelMatcher) GetType() MatchType {
	if m != nil {
		return m.Type
	}
	return MatchType_EQUAL
}

func (m *LabelMatcher) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LabelMatcher) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type QueryResponse struct {
	Timeseries []*TimeSeries `protobuf:"bytes,1,rep,name=timeseries" json:"timeseries,omitempty"`
}

func (m *QueryResponse) Reset()                    { *m = QueryResponse{} }
func (m *QueryResponse) String() string            { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()               {}
func (*QueryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *QueryResponse) GetTimeseries() []*TimeSeries {
	if m != nil {
		return m.Timeseries
	}
	return nil
}

func init() {
	proto.RegisterType((*Sample)(nil), "remote.Sample")
	proto.RegisterType((*LabelPair)(nil), "remote.LabelPair")
	proto.RegisterType((*TimeSeries)(nil), "remote.TimeSeries")
	proto.RegisterType((*WriteRequest)(nil), "remote.WriteRequest")
	proto.RegisterType((*ReadRequest)(nil), "remote.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "remote.ReadResponse")
	proto.RegisterType((*Query)(nil), "remote.Query")
	proto.RegisterType((*LabelMatcher)(nil), "remote.LabelMatcher")
	proto.RegisterType((*QueryResponse)(nil), "remote.QueryResponse")
	proto.RegisterEnum("remote.MatchType", MatchType_name, MatchType_value)
}

func init() { proto.RegisterFile("remote.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x5d, 0x6b, 0x13, 0x41,
	0x14, 0x75, 0xb2, 0x4d, 0xea, 0xde, 0x6c, 0x62, 0xbc, 0x54, 0xc8, 0x63, 0x1c, 0x10, 0x57, 0x91,
	0x22, 0x2d, 0xfa, 0x1e, 0x43, 0x50, 0xa4, 0x69, 0xed, 0x74, 0x45, 0xdf, 0x96, 0xa9, 0xb9, 0xe0,
	0x42, 0x26, 0xbb, 0x9d, 0x99, 0x15, 0xf2, 0x33, 0xfc, 0xc7, 0x92, 0x99, 0xfd, 0x84, 0x3e, 0xf9,
	0x36, 0xf7, 0x9e, 0x73, 0xee, 0x9c, 0x9d, 0x73, 0x17, 0x22, 0x4d, 0x2a, 0xb7, 0x74, 0x5e, 0xe8,
	0xdc, 0xe6, 0x38, 0xf2, 0x15, 0x5f, 0xc2, 0xe8, 0x4e, 0xaa, 0x62, 0x47, 0x78, 0x06, 0xc3, 0x3f,
	0x72, 0x57, 0xd2, 0x9c, 0x2d, 0x58, 0xcc, 0x84, 0x2f, 0xf0, 0x25, 0x44, 0x36, 0x53, 0x64, 0xac,
	0x54, 0x45, 0xaa, 0xcc, 0x7c, 0xb0, 0x60, 0x71, 0x20, 0xc6, 0x4d, 0x6f, 0x63, 0xf8, 0x07, 0x08,
	0xaf, 0xe4, 0x3d, 0xed, 0xbe, 0xc9, 0x4c, 0x23, 0xc2, 0xc9, 0x5e, 0x2a, 0x3f, 0x24, 0x14, 0xee,
	0xdc, 0x4e, 0x1e, 0xb8, 0xa6, 0x2f, 0xb8, 0x04, 0x48, 0x32, 0x45, 0x77, 0xa4, 0x33, 0x32, 0xf8,
	0x06, 0x46, 0xbb, 0xe3, 0x10, 0x33, 0x67, 0x8b, 0x20, 0x1e, 0x5f, 0x3c, 0x3f, 0xaf, 0xec, 0x36,
	0xa3, 0x45, 0x45, 0xc0, 0x18, 0x4e, 0x8d, 0xb3, 0x7c, 0x74, 0x73, 0xe4, 0x4e, 0x6b, 0xae, 0xff,
	0x12, 0x51, 0xc3, 0xfc, 0x13, 0x44, 0x3f, 0x74, 0x66, 0x49, 0xd0, 0x43, 0x49, 0xc6, 0xe2, 0x05,
	0x80, 0x33, 0xee, 0xae, 0xac, 0x2e, 0xc2, 0x5a, 0xdc, 0x9a, 0x11, 0x1d, 0x16, 0xff, 0x08, 0x63,
	0x41, 0x72, 0x5b, 0x8f, 0x78, 0x0d, 0xa7, 0x0f, 0x65, 0x57, 0x3f, 0xa9, 0xf5, 0xb7, 0x25, 0xe9,
	0x83, 0xa8, 0x51, 0xbe, 0x82, 0xc8, 0xeb, 0x4c, 0x91, 0xef, 0x0d, 0xe1, 0x25, 0x84, 0xba, 0x3a,
	0xd7, 0xd2, 0x17, 0x7d, 0x69, 0x85, 0x8a, 0x96, 0xc7, 0xff, 0x32, 0x18, 0x3a, 0x10, 0xdf, 0x01,
	0x1a, 0x2b, 0xb5, 0x4d, 0x7b, 0x69, 0x30, 0x97, 0xc6, 0xcc, 0x21, 0x49, 0x1b, 0x09, 0xc6, 0x30,
	0xa3, 0xfd, 0x36, 0x7d, 0x24, 0xb9, 0x29, 0xed, 0xb7, 0x5d, 0xe6, 0x7b, 0x78, 0xaa, 0xa4, 0xfd,
	0xf5, 0x9b, 0xb4, 0x99, 0x07, 0xce, 0xd5, 0x59, 0xef, 0xe5, 0x37, 0x1e, 0x14, 0x0d, 0x8b, 0xa7,
	0x10, 0x75, 0x11, 0x7c, 0x05, 0x27, 0xf6, 0x50, 0xf8, 0xc4, 0xa7, 0x6d, 0x6e, 0x0e, 0x4e, 0x0e,
	0x05, 0x09, 0x07, 0x37, 0x8b, 0x31, 0x78, 0x6c, 0x31, 0x82, 0xee, 0x62, 0xac, 0x60, 0xd2, 0x7b,
	0x90, 0xff, 0x89, 0xed, 0xed, 0x57, 0x08, 0x1b, 0x07, 0x18, 0xc2, 0x70, 0x7d, 0xfb, 0x7d, 0x79,
	0x35, 0x7b, 0x82, 0x13, 0x08, 0xaf, 0x6f, 0x92, 0xd4, 0x97, 0x0c, 0x9f, 0xc1, 0x58, 0xac, 0x3f,
	0xaf, 0x7f, 0xa6, 0x9b, 0x65, 0xb2, 0xfa, 0x32, 0x1b, 0x20, 0xc2, 0xd4, 0x37, 0xae, 0x6f, 0xaa,
	0x5e, 0x70, 0x3f, 0x72, 0xbf, 0xcc, 0xe5, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x0c, 0x7b,
	0x5d, 0x42, 0x03, 0x00, 0x00,
}
