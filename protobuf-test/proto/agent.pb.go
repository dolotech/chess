// Code generated by protoc-gen-go. DO NOT EDIT.
// source: agent.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	agent.proto

It has these top-level messages:
	User
	SearchResponse
	SomeOtherMessage
	ErrorStatus
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Id   int32  `protobuf:"varint,1,opt,name=Id" json:"Id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto1.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type SearchResponse struct {
	Results []*SearchResponse_Result `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
}

func (m *SearchResponse) Reset()                    { *m = SearchResponse{} }
func (m *SearchResponse) String() string            { return proto1.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()               {}
func (*SearchResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SearchResponse) GetResults() []*SearchResponse_Result {
	if m != nil {
		return m.Results
	}
	return nil
}

type SearchResponse_Result struct {
	Url      string   `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	Title    string   `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Snippets []string `protobuf:"bytes,3,rep,name=snippets" json:"snippets,omitempty"`
}

func (m *SearchResponse_Result) Reset()                    { *m = SearchResponse_Result{} }
func (m *SearchResponse_Result) String() string            { return proto1.CompactTextString(m) }
func (*SearchResponse_Result) ProtoMessage()               {}
func (*SearchResponse_Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *SearchResponse_Result) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *SearchResponse_Result) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SearchResponse_Result) GetSnippets() []string {
	if m != nil {
		return m.Snippets
	}
	return nil
}

type SomeOtherMessage struct {
	Result *SearchResponse_Result `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
}

func (m *SomeOtherMessage) Reset()                    { *m = SomeOtherMessage{} }
func (m *SomeOtherMessage) String() string            { return proto1.CompactTextString(m) }
func (*SomeOtherMessage) ProtoMessage()               {}
func (*SomeOtherMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SomeOtherMessage) GetResult() *SearchResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

type ErrorStatus struct {
	Message string                 `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	Details []*google_protobuf.Any `protobuf:"bytes,2,rep,name=details" json:"details,omitempty"`
}

func (m *ErrorStatus) Reset()                    { *m = ErrorStatus{} }
func (m *ErrorStatus) String() string            { return proto1.CompactTextString(m) }
func (*ErrorStatus) ProtoMessage()               {}
func (*ErrorStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ErrorStatus) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ErrorStatus) GetDetails() []*google_protobuf.Any {
	if m != nil {
		return m.Details
	}
	return nil
}

func init() {
	proto1.RegisterType((*User)(nil), "proto.User")
	proto1.RegisterType((*SearchResponse)(nil), "proto.SearchResponse")
	proto1.RegisterType((*SearchResponse_Result)(nil), "proto.SearchResponse.Result")
	proto1.RegisterType((*SomeOtherMessage)(nil), "proto.SomeOtherMessage")
	proto1.RegisterType((*ErrorStatus)(nil), "proto.ErrorStatus")
}

func init() { proto1.RegisterFile("agent.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xbf, 0x6b, 0xfb, 0x30,
	0x10, 0xc5, 0xb1, 0x9d, 0xc4, 0x5f, 0x9f, 0x21, 0x04, 0x91, 0x41, 0x5f, 0xd3, 0xc1, 0x78, 0x32,
	0x1d, 0x14, 0x68, 0x4b, 0xf7, 0x0e, 0x85, 0x06, 0xfa, 0x03, 0x64, 0x4a, 0x67, 0xa5, 0xbe, 0x3a,
	0x06, 0xdb, 0x32, 0x3a, 0x79, 0xc8, 0x3f, 0xd2, 0xbf, 0xb7, 0x54, 0x8a, 0x0b, 0x9d, 0x3a, 0xe9,
	0xde, 0xf1, 0xe1, 0xbd, 0xd3, 0x83, 0x54, 0x35, 0x38, 0x58, 0x31, 0x1a, 0x6d, 0x35, 0x5b, 0xba,
	0x27, 0xfb, 0xdf, 0x68, 0xdd, 0x74, 0xb8, 0x73, 0xea, 0x30, 0x7d, 0xec, 0xd4, 0x70, 0xf2, 0x44,
	0x71, 0x09, 0x8b, 0x57, 0x42, 0xc3, 0xd6, 0x10, 0xee, 0x6b, 0x1e, 0xe4, 0x41, 0xb9, 0x94, 0xe1,
	0xbe, 0x66, 0x0c, 0x16, 0xcf, 0xaa, 0x47, 0x1e, 0xe6, 0x41, 0x99, 0x48, 0x37, 0x17, 0x9f, 0x01,
	0xac, 0x2b, 0x54, 0xe6, 0xfd, 0x28, 0x91, 0x46, 0x3d, 0x10, 0xb2, 0x5b, 0x88, 0x0d, 0xd2, 0xd4,
	0x59, 0xe2, 0x41, 0x1e, 0x95, 0xe9, 0xd5, 0x85, 0xf7, 0x15, 0xbf, 0x39, 0x21, 0x1d, 0x24, 0x67,
	0x38, 0x7b, 0x84, 0x95, 0x5f, 0xb1, 0x0d, 0x44, 0x93, 0xe9, 0x5c, 0x72, 0x22, 0xbf, 0x47, 0xb6,
	0x85, 0xa5, 0x6d, 0x6d, 0x37, 0x67, 0x7b, 0xc1, 0x32, 0xf8, 0x47, 0x43, 0x3b, 0x8e, 0x68, 0x89,
	0x47, 0x79, 0x54, 0x26, 0xf2, 0x47, 0x17, 0x0f, 0xb0, 0xa9, 0x74, 0x8f, 0x2f, 0xf6, 0x88, 0xe6,
	0x09, 0x89, 0x54, 0x83, 0xec, 0x06, 0x56, 0x3e, 0xcc, 0x59, 0xff, 0x75, 0xd8, 0x99, 0x2d, 0xde,
	0x20, 0xbd, 0x37, 0x46, 0x9b, 0xca, 0x2a, 0x3b, 0x11, 0xe3, 0x10, 0xf7, 0xde, 0xef, 0x7c, 0xe0,
	0x2c, 0x99, 0x80, 0xb8, 0x46, 0xab, 0xda, 0x8e, 0x78, 0xe8, 0x3e, 0xbe, 0x15, 0xbe, 0x64, 0x31,
	0x97, 0x2c, 0xee, 0x86, 0x93, 0x9c, 0xa1, 0xc3, 0xca, 0xad, 0xaf, 0xbf, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xc6, 0x3f, 0x37, 0x10, 0x9f, 0x01, 0x00, 0x00,
}
