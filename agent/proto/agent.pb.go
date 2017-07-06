// Code generated by protoc-gen-go. DO NOT EDIT.
// source: agent.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	agent.proto
	auth.proto
	centre.proto
	game.proto
	room.proto

It has these top-level messages:
	AutoId
	SeedInfo
	BaseAck
	UserLoginReq
	TokenInfoArgs
	TokenInfoRes
	RefreshTokenArgs
	RefreshTokenRes
	DestroyTokenArgs
	DestroyTokenRes
	GameListArgs
	GameListRes
	GamePlaygroundArgs
	GamePlaygroundRes
	RoomInfoArgs
	RoomInfoRes
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

// 心跳包
type AutoId struct {
	Id int32 `protobuf:"varint,1,opt,name=Id" json:"Id,omitempty"`
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

// 通信加密种子
type SeedInfo struct {
	ClientSendSeed    int32 `protobuf:"varint,1,opt,name=ClientSendSeed" json:"ClientSendSeed,omitempty"`
	ClientReceiveSeed int32 `protobuf:"varint,2,opt,name=ClientReceiveSeed" json:"ClientReceiveSeed,omitempty"`
}

func (m *SeedInfo) Reset()                    { *m = SeedInfo{} }
func (m *SeedInfo) String() string            { return proto1.CompactTextString(m) }
func (*SeedInfo) ProtoMessage()               {}
func (*SeedInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SeedInfo) GetClientSendSeed() int32 {
	if m != nil {
		return m.ClientSendSeed
	}
	return 0
}

func (m *SeedInfo) GetClientReceiveSeed() int32 {
	if m != nil {
		return m.ClientReceiveSeed
	}
	return 0
}

// 一般性回复payload,1代表成功
type BaseAck struct {
	Ret int32  `protobuf:"varint,1,opt,name=Ret" json:"Ret,omitempty"`
	Msg string `protobuf:"bytes,2,opt,name=Msg" json:"Msg,omitempty"`
}

func (m *BaseAck) Reset()                    { *m = BaseAck{} }
func (m *BaseAck) String() string            { return proto1.CompactTextString(m) }
func (*BaseAck) ProtoMessage()               {}
func (*BaseAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *BaseAck) GetRet() int32 {
	if m != nil {
		return m.Ret
	}
	return 0
}

func (m *BaseAck) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type UserLoginReq struct {
	UserId     int32  `protobuf:"varint,1,opt,name=UserId" json:"UserId,omitempty"`
	AppFrom    string `protobuf:"bytes,2,opt,name=AppFrom" json:"AppFrom,omitempty"`
	AppChannel string `protobuf:"bytes,3,opt,name=AppChannel" json:"AppChannel,omitempty"`
	AppVer     int32  `protobuf:"varint,4,opt,name=AppVer" json:"AppVer,omitempty"`
	UniqueId   string `protobuf:"bytes,5,opt,name=UniqueId" json:"UniqueId,omitempty"`
	Token      string `protobuf:"bytes,6,opt,name=Token" json:"Token,omitempty"`
}

func (m *UserLoginReq) Reset()                    { *m = UserLoginReq{} }
func (m *UserLoginReq) String() string            { return proto1.CompactTextString(m) }
func (*UserLoginReq) ProtoMessage()               {}
func (*UserLoginReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UserLoginReq) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UserLoginReq) GetAppFrom() string {
	if m != nil {
		return m.AppFrom
	}
	return ""
}

func (m *UserLoginReq) GetAppChannel() string {
	if m != nil {
		return m.AppChannel
	}
	return ""
}

func (m *UserLoginReq) GetAppVer() int32 {
	if m != nil {
		return m.AppVer
	}
	return 0
}

func (m *UserLoginReq) GetUniqueId() string {
	if m != nil {
		return m.UniqueId
	}
	return ""
}

func (m *UserLoginReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto1.RegisterType((*AutoId)(nil), "proto.AutoId")
	proto1.RegisterType((*SeedInfo)(nil), "proto.SeedInfo")
	proto1.RegisterType((*BaseAck)(nil), "proto.BaseAck")
	proto1.RegisterType((*UserLoginReq)(nil), "proto.UserLoginReq")
}

func init() { proto1.RegisterFile("agent.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0xc7, 0x49, 0x6a, 0xd2, 0x3a, 0x4a, 0xd1, 0x41, 0x64, 0xf1, 0x20, 0x92, 0x83, 0x78, 0x50,
	0x2f, 0x3e, 0xc1, 0x5a, 0x10, 0x16, 0xf4, 0xb2, 0xb5, 0x9e, 0x8d, 0xdd, 0x31, 0x86, 0xd6, 0xd9,
	0x6d, 0xb2, 0xf5, 0xa1, 0x7c, 0x4a, 0xd9, 0x4d, 0x0c, 0x62, 0x4f, 0x99, 0xdf, 0xff, 0x8b, 0xb0,
	0x70, 0x50, 0x56, 0xc4, 0xfe, 0xd6, 0x35, 0xd6, 0x5b, 0xcc, 0xe2, 0xa7, 0x10, 0x90, 0xcb, 0xad,
	0xb7, 0xca, 0xe0, 0x14, 0x52, 0x65, 0x44, 0x72, 0x91, 0x5c, 0x65, 0x3a, 0x55, 0xa6, 0x78, 0x85,
	0xc9, 0x9c, 0xc8, 0x28, 0x7e, 0xb7, 0x78, 0x09, 0xd3, 0xd9, 0xba, 0x26, 0xf6, 0x73, 0x62, 0x13,
	0xd4, 0x3e, 0xf7, 0x4f, 0xc5, 0x6b, 0x38, 0xee, 0x14, 0x4d, 0x4b, 0xaa, 0xbf, 0x28, 0x46, 0xd3,
	0x18, 0xdd, 0x35, 0x8a, 0x1b, 0x18, 0xdf, 0x97, 0x2d, 0xc9, 0xe5, 0x0a, 0x8f, 0x60, 0xa4, 0xc9,
	0xf7, 0xab, 0xe1, 0x0c, 0xca, 0x53, 0x5b, 0xc5, 0xf2, 0xbe, 0x0e, 0x67, 0xf1, 0x9d, 0xc0, 0xe1,
	0xa2, 0xa5, 0xe6, 0xd1, 0x56, 0x35, 0x6b, 0xda, 0xe0, 0x29, 0xe4, 0x81, 0x87, 0xbf, 0xee, 0x09,
	0x05, 0x8c, 0xa5, 0x73, 0x0f, 0x8d, 0xfd, 0xec, 0xeb, 0xbf, 0x88, 0xe7, 0x00, 0xd2, 0xb9, 0xd9,
	0x47, 0xc9, 0x4c, 0x6b, 0x31, 0x8a, 0xe6, 0x1f, 0x25, 0x2c, 0x4a, 0xe7, 0x5e, 0xa8, 0x11, 0x7b,
	0xdd, 0x62, 0x47, 0x78, 0x06, 0x93, 0x05, 0xd7, 0x9b, 0x2d, 0x29, 0x23, 0xb2, 0xd8, 0x1a, 0x18,
	0x4f, 0x20, 0x7b, 0xb6, 0x2b, 0x62, 0x91, 0x47, 0xa3, 0x83, 0xb7, 0x3c, 0x3e, 0xef, 0xdd, 0x4f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x7c, 0x7c, 0x8f, 0x74, 0x01, 0x00, 0x00,
}
