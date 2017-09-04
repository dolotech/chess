// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sts.proto

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type GameTableInfoArgs struct {
	RoomId  int32        `protobuf:"varint,1,opt,name=room_id,json=roomId" json:"room_id,omitempty"`
	TableId string       `protobuf:"bytes,2,opt,name=table_id,json=tableId" json:"table_id,omitempty"`
	Max     int32        `protobuf:"varint,3,opt,name=max" json:"max,omitempty"`
	Start   int32        `protobuf:"varint,4,opt,name=start" json:"start,omitempty"`
	End     int32        `protobuf:"varint,5,opt,name=end" json:"end,omitempty"`
	Cards   []*CardsInfo `protobuf:"bytes,6,rep,name=cards" json:"cards,omitempty"`
	Button  int32        `protobuf:"varint,7,opt,name=button" json:"button,omitempty"`
	Sb      int32        `protobuf:"varint,8,opt,name=sb" json:"sb,omitempty"`
	Bb      int32        `protobuf:"varint,9,opt,name=bb" json:"bb,omitempty"`
	SbPos   int32        `protobuf:"varint,10,opt,name=sb_pos,json=sbPos" json:"sb_pos,omitempty"`
	BbPos   int32        `protobuf:"varint,11,opt,name=bb_pos,json=bbPos" json:"bb_pos,omitempty"`
	Pot     []int32      `protobuf:"varint,12,rep,packed,name=pot" json:"pot,omitempty"`
	Player  []*Player    `protobuf:"bytes,13,rep,name=player" json:"player,omitempty"`
}

func (m *GameTableInfoArgs) Reset()                    { *m = GameTableInfoArgs{} }
func (m *GameTableInfoArgs) String() string            { return proto1.CompactTextString(m) }
func (*GameTableInfoArgs) ProtoMessage()               {}
func (*GameTableInfoArgs) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *GameTableInfoArgs) GetRoomId() int32 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *GameTableInfoArgs) GetTableId() string {
	if m != nil {
		return m.TableId
	}
	return ""
}

func (m *GameTableInfoArgs) GetMax() int32 {
	if m != nil {
		return m.Max
	}
	return 0
}

func (m *GameTableInfoArgs) GetStart() int32 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *GameTableInfoArgs) GetEnd() int32 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *GameTableInfoArgs) GetCards() []*CardsInfo {
	if m != nil {
		return m.Cards
	}
	return nil
}

func (m *GameTableInfoArgs) GetButton() int32 {
	if m != nil {
		return m.Button
	}
	return 0
}

func (m *GameTableInfoArgs) GetSb() int32 {
	if m != nil {
		return m.Sb
	}
	return 0
}

func (m *GameTableInfoArgs) GetBb() int32 {
	if m != nil {
		return m.Bb
	}
	return 0
}

func (m *GameTableInfoArgs) GetSbPos() int32 {
	if m != nil {
		return m.SbPos
	}
	return 0
}

func (m *GameTableInfoArgs) GetBbPos() int32 {
	if m != nil {
		return m.BbPos
	}
	return 0
}

func (m *GameTableInfoArgs) GetPot() []int32 {
	if m != nil {
		return m.Pot
	}
	return nil
}

func (m *GameTableInfoArgs) GetPlayer() []*Player {
	if m != nil {
		return m.Player
	}
	return nil
}

type CardsInfo struct {
	Suit  int32 `protobuf:"varint,1,opt,name=suit" json:"suit,omitempty"`
	Value int32 `protobuf:"varint,2,opt,name=value" json:"value,omitempty"`
}

func (m *CardsInfo) Reset()                    { *m = CardsInfo{} }
func (m *CardsInfo) String() string            { return proto1.CompactTextString(m) }
func (*CardsInfo) ProtoMessage()               {}
func (*CardsInfo) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *CardsInfo) GetSuit() int32 {
	if m != nil {
		return m.Suit
	}
	return 0
}

func (m *CardsInfo) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Player struct {
	Id             int32           `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Nickname       string          `protobuf:"bytes,2,opt,name=nickname" json:"nickname,omitempty"`
	Avatar         string          `protobuf:"bytes,3,opt,name=avatar" json:"avatar,omitempty"`
	Pos            int32           `protobuf:"varint,4,opt,name=pos" json:"pos,omitempty"`
	Bet            int32           `protobuf:"varint,5,opt,name=bet" json:"bet,omitempty"`
	Win            int32           `protobuf:"varint,6,opt,name=win" json:"win,omitempty"`
	FormerChips    int32           `protobuf:"varint,7,opt,name=former_chips,json=formerChips" json:"former_chips,omitempty"`
	CurrentChips   int32           `protobuf:"varint,8,opt,name=current_chips,json=currentChips" json:"current_chips,omitempty"`
	Action         string          `protobuf:"bytes,9,opt,name=action" json:"action,omitempty"`
	Cards          []*CardsInfo    `protobuf:"bytes,10,rep,name=cards" json:"cards,omitempty"`
	HandLevel      int32           `protobuf:"varint,11,opt,name=hand_level,json=handLevel" json:"hand_level,omitempty"`
	HandFinalValue int32           `protobuf:"varint,12,opt,name=hand_final_value,json=handFinalValue" json:"hand_final_value,omitempty"`
	Actions        []*PlayerAction `protobuf:"bytes,13,rep,name=actions" json:"actions,omitempty"`
}

func (m *Player) Reset()                    { *m = Player{} }
func (m *Player) String() string            { return proto1.CompactTextString(m) }
func (*Player) ProtoMessage()               {}
func (*Player) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *Player) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Player) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *Player) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *Player) GetPos() int32 {
	if m != nil {
		return m.Pos
	}
	return 0
}

func (m *Player) GetBet() int32 {
	if m != nil {
		return m.Bet
	}
	return 0
}

func (m *Player) GetWin() int32 {
	if m != nil {
		return m.Win
	}
	return 0
}

func (m *Player) GetFormerChips() int32 {
	if m != nil {
		return m.FormerChips
	}
	return 0
}

func (m *Player) GetCurrentChips() int32 {
	if m != nil {
		return m.CurrentChips
	}
	return 0
}

func (m *Player) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *Player) GetCards() []*CardsInfo {
	if m != nil {
		return m.Cards
	}
	return nil
}

func (m *Player) GetHandLevel() int32 {
	if m != nil {
		return m.HandLevel
	}
	return 0
}

func (m *Player) GetHandFinalValue() int32 {
	if m != nil {
		return m.HandFinalValue
	}
	return 0
}

func (m *Player) GetActions() []*PlayerAction {
	if m != nil {
		return m.Actions
	}
	return nil
}

type StsRes struct {
	Ret int32  `protobuf:"varint,1,opt,name=ret" json:"ret,omitempty"`
	Msg string `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
}

func (m *StsRes) Reset()                    { *m = StsRes{} }
func (m *StsRes) String() string            { return proto1.CompactTextString(m) }
func (*StsRes) ProtoMessage()               {}
func (*StsRes) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *StsRes) GetRet() int32 {
	if m != nil {
		return m.Ret
	}
	return 0
}

func (m *StsRes) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type PlayerAction struct {
	Action string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	Bet    int32  `protobuf:"varint,2,opt,name=bet" json:"bet,omitempty"`
}

func (m *PlayerAction) Reset()                    { *m = PlayerAction{} }
func (m *PlayerAction) String() string            { return proto1.CompactTextString(m) }
func (*PlayerAction) ProtoMessage()               {}
func (*PlayerAction) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

func (m *PlayerAction) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *PlayerAction) GetBet() int32 {
	if m != nil {
		return m.Bet
	}
	return 0
}

func init() {
	proto1.RegisterType((*GameTableInfoArgs)(nil), "proto.GameTableInfoArgs")
	proto1.RegisterType((*CardsInfo)(nil), "proto.CardsInfo")
	proto1.RegisterType((*Player)(nil), "proto.Player")
	proto1.RegisterType((*StsRes)(nil), "proto.StsRes")
	proto1.RegisterType((*PlayerAction)(nil), "proto.PlayerAction")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for StsService service

type StsServiceClient interface {
	// 获取游戏信息
	GameInfo(ctx context.Context, in *GameTableInfoArgs, opts ...grpc.CallOption) (*StsRes, error)
}

type stsServiceClient struct {
	cc *grpc.ClientConn
}

func NewStsServiceClient(cc *grpc.ClientConn) StsServiceClient {
	return &stsServiceClient{cc}
}

func (c *stsServiceClient) GameInfo(ctx context.Context, in *GameTableInfoArgs, opts ...grpc.CallOption) (*StsRes, error) {
	out := new(StsRes)
	err := grpc.Invoke(ctx, "/proto.StsService/GameInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for StsService service

type StsServiceServer interface {
	// 获取游戏信息
	GameInfo(context.Context, *GameTableInfoArgs) (*StsRes, error)
}

func RegisterStsServiceServer(s *grpc.Server, srv StsServiceServer) {
	s.RegisterService(&_StsService_serviceDesc, srv)
}

func _StsService_GameInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GameTableInfoArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StsServiceServer).GameInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.StsService/GameInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StsServiceServer).GameInfo(ctx, req.(*GameTableInfoArgs))
	}
	return interceptor(ctx, in, info, handler)
}

var _StsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.StsService",
	HandlerType: (*StsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GameInfo",
			Handler:    _StsService_GameInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sts.proto",
}

func init() { proto1.RegisterFile("sts.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 527 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x4f, 0x6b, 0x1b, 0x3d,
	0x10, 0xc6, 0x5f, 0xef, 0x66, 0xd7, 0xde, 0x89, 0x13, 0xfc, 0xaa, 0xff, 0xd4, 0x40, 0xc1, 0xdd,
	0xd2, 0xe2, 0x43, 0x9b, 0x43, 0x4a, 0xa0, 0xd7, 0x60, 0x68, 0x31, 0xf4, 0x10, 0xd6, 0xa5, 0x57,
	0x23, 0xed, 0xca, 0xc9, 0xd2, 0xb5, 0xb4, 0x48, 0xb2, 0xd3, 0x5e, 0xfb, 0xe9, 0xfa, 0xb1, 0xca,
	0x8c, 0x64, 0x37, 0xa1, 0xd0, 0x93, 0xe7, 0xf9, 0xcd, 0xb0, 0x7e, 0xf4, 0x8c, 0x04, 0x85, 0xf3,
	0xee, 0xbc, 0xb7, 0xc6, 0x1b, 0x96, 0xd1, 0x4f, 0xf9, 0x2b, 0x81, 0xff, 0x3f, 0x89, 0x8d, 0xfa,
	0x22, 0x64, 0xa7, 0x16, 0x7a, 0x6d, 0xae, 0xec, 0x8d, 0x63, 0xcf, 0x60, 0x68, 0x8d, 0xd9, 0xac,
	0xda, 0x86, 0x0f, 0xa6, 0x83, 0x59, 0x56, 0xe5, 0x28, 0x17, 0x0d, 0x7b, 0x0e, 0x23, 0x8f, 0x93,
	0xd8, 0x49, 0xa6, 0x83, 0x59, 0x51, 0x0d, 0x49, 0x2f, 0x1a, 0x36, 0x81, 0x74, 0x23, 0xbe, 0xf3,
	0x94, 0xe6, 0xb1, 0x64, 0x8f, 0x21, 0x73, 0x5e, 0x58, 0xcf, 0x8f, 0x88, 0x05, 0x81, 0x73, 0x4a,
	0x37, 0x3c, 0x0b, 0x73, 0x4a, 0x37, 0xec, 0x0d, 0x64, 0xb5, 0xb0, 0x8d, 0xe3, 0xf9, 0x34, 0x9d,
	0x1d, 0x5f, 0x4c, 0x82, 0xc3, 0xf3, 0x39, 0x32, 0xb4, 0x54, 0x85, 0x36, 0x7b, 0x0a, 0xb9, 0xdc,
	0x7a, 0x6f, 0x34, 0x1f, 0x06, 0x53, 0x41, 0xb1, 0x53, 0x48, 0x9c, 0xe4, 0x23, 0x62, 0x89, 0x93,
	0xa8, 0xa5, 0xe4, 0x45, 0xd0, 0x52, 0xb2, 0x27, 0x90, 0x3b, 0xb9, 0xea, 0x8d, 0xe3, 0x10, 0x8d,
	0xc8, 0x6b, 0xe3, 0x10, 0xcb, 0x80, 0x8f, 0x03, 0x96, 0x84, 0x27, 0x90, 0xf6, 0xc6, 0xf3, 0xf1,
	0x34, 0x45, 0x7f, 0xbd, 0xf1, 0xec, 0x35, 0xe4, 0x7d, 0x27, 0x7e, 0x28, 0xcb, 0x4f, 0xc8, 0xe0,
	0x49, 0x34, 0x78, 0x4d, 0xb0, 0x8a, 0xcd, 0xf2, 0x12, 0x8a, 0x83, 0x65, 0xc6, 0xe0, 0xc8, 0x6d,
	0x5b, 0x1f, 0xe3, 0xa3, 0x1a, 0xf3, 0xd8, 0x89, 0x6e, 0xab, 0x28, 0xb9, 0xac, 0x0a, 0xa2, 0xfc,
	0x99, 0x42, 0x1e, 0xbe, 0x84, 0xc6, 0x0f, 0x89, 0x27, 0x6d, 0xc3, 0xce, 0x60, 0xa4, 0xdb, 0xfa,
	0x9b, 0x16, 0x1b, 0x15, 0xd3, 0x3e, 0x68, 0x0c, 0x43, 0xec, 0x84, 0x17, 0x96, 0x12, 0x2f, 0xaa,
	0xa8, 0x82, 0x7d, 0x17, 0x23, 0xc7, 0x12, 0x89, 0x54, 0x7e, 0x1f, 0xb8, 0x54, 0xb4, 0x82, 0xbb,
	0x56, 0xf3, 0x3c, 0x90, 0xbb, 0x56, 0xb3, 0x97, 0x30, 0x5e, 0x1b, 0xbb, 0x51, 0x76, 0x55, 0xdf,
	0xb6, 0xbd, 0x8b, 0x01, 0x1f, 0x07, 0x36, 0x47, 0xc4, 0x5e, 0xc1, 0x49, 0xbd, 0xb5, 0x56, 0x69,
	0x1f, 0x67, 0x42, 0xe0, 0xe3, 0x08, 0xc3, 0x10, 0xba, 0xaa, 0x7d, 0x6b, 0x34, 0xc5, 0x8f, 0xae,
	0x48, 0xfd, 0x59, 0x31, 0xfc, 0x7b, 0xc5, 0x2f, 0x00, 0x6e, 0x85, 0x6e, 0x56, 0x9d, 0xda, 0xa9,
	0x2e, 0xee, 0xa5, 0x40, 0xf2, 0x19, 0x01, 0x9b, 0xc1, 0x84, 0xda, 0xeb, 0x56, 0x8b, 0x6e, 0x15,
	0xc2, 0x1c, 0xd3, 0xd0, 0x29, 0xf2, 0x8f, 0x88, 0xbf, 0x22, 0x65, 0xef, 0x60, 0x18, 0xfe, 0xda,
	0xc5, 0xa5, 0x3d, 0x7a, 0xb0, 0xb4, 0x2b, 0xea, 0x55, 0xfb, 0x99, 0xf2, 0x2d, 0xe4, 0x4b, 0xef,
	0x2a, 0x45, 0x69, 0x59, 0xb5, 0xdf, 0x1b, 0x96, 0x74, 0xb1, 0xdd, 0x4d, 0x5c, 0x00, 0x96, 0xe5,
	0x07, 0x18, 0xdf, 0xff, 0xcc, 0xbd, 0x53, 0x0f, 0x1e, 0x9c, 0x3a, 0x26, 0x9f, 0x1c, 0x92, 0xbf,
	0x98, 0x03, 0x2c, 0xbd, 0x5b, 0x2a, 0xbb, 0x6b, 0x6b, 0xc5, 0x2e, 0x61, 0x84, 0x6f, 0x8f, 0x2e,
	0x0c, 0x8f, 0xfe, 0xfe, 0x7a, 0x8c, 0x67, 0xfb, 0xeb, 0x16, 0x0c, 0x96, 0xff, 0xc9, 0x9c, 0xf4,
	0xfb, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x21, 0xfc, 0xa7, 0xeb, 0xce, 0x03, 0x00, 0x00,
}
