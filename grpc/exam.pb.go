// Code generated by protoc-gen-go. DO NOT EDIT.
// source: exam.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	exam.proto

It has these top-level messages:
	AutoId
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

// xi tiaobao
type AutoId struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *AutoId) Reset()                    { *m = AutoId{} }
func (m *AutoId) String() string            { return proto1.CompactTextString(m) }
func (*AutoId) ProtoMessage()               {}
func (*AutoId) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AutoId) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto1.RegisterType((*AutoId)(nil), "proto.AutoId")
}

func init() { proto1.RegisterFile("exam.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 71 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xad, 0x48, 0xcc,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x12, 0x5c, 0x6c, 0x8e, 0xa5,
	0x25, 0xf9, 0x9e, 0x29, 0x42, 0x7c, 0x5c, 0x4c, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac,
	0x41, 0x4c, 0x99, 0x29, 0x49, 0x6c, 0x60, 0x05, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5b,
	0x57, 0x0c, 0xbb, 0x35, 0x00, 0x00, 0x00,
}
