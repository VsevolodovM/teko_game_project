// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tko.proto

package tko

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GameParameter struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameParameter) Reset()         { *m = GameParameter{} }
func (m *GameParameter) String() string { return proto.CompactTextString(m) }
func (*GameParameter) ProtoMessage()    {}
func (*GameParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17e2fb55a57754, []int{0}
}

func (m *GameParameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameParameter.Unmarshal(m, b)
}
func (m *GameParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameParameter.Marshal(b, m, deterministic)
}
func (m *GameParameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameParameter.Merge(m, src)
}
func (m *GameParameter) XXX_Size() int {
	return xxx_messageInfo_GameParameter.Size(m)
}
func (m *GameParameter) XXX_DiscardUnknown() {
	xxx_messageInfo_GameParameter.DiscardUnknown(m)
}

var xxx_messageInfo_GameParameter proto.InternalMessageInfo

// x1, y1: x- and y-coordinates of the stone that is about to be moved (is not used for the first 8 half-moves)
// x2, y2: x- and y-coordinates of the field, where a stone is about to be placed
type GameTurn struct {
	X1                   uint32   `protobuf:"varint,1,opt,name=x1,proto3" json:"x1,omitempty"`
	Y1                   uint32   `protobuf:"varint,2,opt,name=y1,proto3" json:"y1,omitempty"`
	X2                   uint32   `protobuf:"varint,3,opt,name=x2,proto3" json:"x2,omitempty"`
	Y2                   uint32   `protobuf:"varint,4,opt,name=y2,proto3" json:"y2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameTurn) Reset()         { *m = GameTurn{} }
func (m *GameTurn) String() string { return proto.CompactTextString(m) }
func (*GameTurn) ProtoMessage()    {}
func (*GameTurn) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17e2fb55a57754, []int{1}
}

func (m *GameTurn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameTurn.Unmarshal(m, b)
}
func (m *GameTurn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameTurn.Marshal(b, m, deterministic)
}
func (m *GameTurn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameTurn.Merge(m, src)
}
func (m *GameTurn) XXX_Size() int {
	return xxx_messageInfo_GameTurn.Size(m)
}
func (m *GameTurn) XXX_DiscardUnknown() {
	xxx_messageInfo_GameTurn.DiscardUnknown(m)
}

var xxx_messageInfo_GameTurn proto.InternalMessageInfo

func (m *GameTurn) GetX1() uint32 {
	if m != nil {
		return m.X1
	}
	return 0
}

func (m *GameTurn) GetY1() uint32 {
	if m != nil {
		return m.Y1
	}
	return 0
}

func (m *GameTurn) GetX2() uint32 {
	if m != nil {
		return m.X2
	}
	return 0
}

func (m *GameTurn) GetY2() uint32 {
	if m != nil {
		return m.Y2
	}
	return 0
}

// 1d array representing the board
// the board should be reconstructed in the following way (using the listed indices of the array):
//   -------------> x
// |  0  1  2  3  4
// |  5  6  7  8  9
// | 10 11 12 13 14
// | 15 16 17 18 19
// V 20 21 22 23 24
// y
type GameState struct {
	Board                []int32  `protobuf:"varint,1,rep,packed,name=board,proto3" json:"board,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameState) Reset()         { *m = GameState{} }
func (m *GameState) String() string { return proto.CompactTextString(m) }
func (*GameState) ProtoMessage()    {}
func (*GameState) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17e2fb55a57754, []int{2}
}

func (m *GameState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameState.Unmarshal(m, b)
}
func (m *GameState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameState.Marshal(b, m, deterministic)
}
func (m *GameState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameState.Merge(m, src)
}
func (m *GameState) XXX_Size() int {
	return xxx_messageInfo_GameState.Size(m)
}
func (m *GameState) XXX_DiscardUnknown() {
	xxx_messageInfo_GameState.DiscardUnknown(m)
}

var xxx_messageInfo_GameState proto.InternalMessageInfo

func (m *GameState) GetBoard() []int32 {
	if m != nil {
		return m.Board
	}
	return nil
}

func init() {
	proto.RegisterType((*GameParameter)(nil), "tko.GameParameter")
	proto.RegisterType((*GameTurn)(nil), "tko.GameTurn")
	proto.RegisterType((*GameState)(nil), "tko.GameState")
}

func init() {
	proto.RegisterFile("tko.proto", fileDescriptor_2d17e2fb55a57754)
}

var fileDescriptor_2d17e2fb55a57754 = []byte{
	// 163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0xc9, 0xce, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0xc9, 0xce, 0x57, 0xe2, 0xe7, 0xe2, 0x75, 0x4f,
	0xcc, 0x4d, 0x0d, 0x48, 0x2c, 0x4a, 0xcc, 0x4d, 0x2d, 0x49, 0x2d, 0x52, 0xf2, 0xe2, 0xe2, 0x00,
	0x09, 0x84, 0x94, 0x16, 0xe5, 0x09, 0xf1, 0x71, 0x31, 0x55, 0x18, 0x4a, 0x30, 0x2a, 0x30, 0x6a,
	0xf0, 0x06, 0x31, 0x55, 0x18, 0x82, 0xf8, 0x95, 0x86, 0x12, 0x4c, 0x10, 0x7e, 0x25, 0x98, 0x5f,
	0x61, 0x24, 0xc1, 0x0c, 0x95, 0x37, 0x02, 0xcb, 0x1b, 0x49, 0xb0, 0x40, 0xe5, 0x8d, 0x94, 0x14,
	0xb9, 0x38, 0x41, 0x66, 0x05, 0x97, 0x24, 0x96, 0xa4, 0x0a, 0x89, 0x70, 0xb1, 0x26, 0xe5, 0x27,
	0x16, 0xa5, 0x48, 0x30, 0x2a, 0x30, 0x6b, 0xb0, 0x06, 0x41, 0x38, 0x4e, 0xc2, 0x51, 0x82, 0x25,
	0xa9, 0xd9, 0xf9, 0xf1, 0xe9, 0x89, 0xb9, 0xa9, 0xfa, 0x05, 0xd9, 0xe9, 0xfa, 0x25, 0xd9, 0xf9,
	0x49, 0x6c, 0x60, 0x07, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x42, 0xe2, 0xe3, 0xf0, 0xad,
	0x00, 0x00, 0x00,
}
