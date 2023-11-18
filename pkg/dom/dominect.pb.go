// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dominect.proto

package dom

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

// The GameParameter packet defines the board of m x n fields (width x height).
type GameParameter struct {
	BoardWidth           uint32   `protobuf:"varint,1,opt,name=board_width,json=boardWidth,proto3" json:"board_width,omitempty"`
	BoardHeight          uint32   `protobuf:"varint,2,opt,name=board_height,json=boardHeight,proto3" json:"board_height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameParameter) Reset()         { *m = GameParameter{} }
func (m *GameParameter) String() string { return proto.CompactTextString(m) }
func (*GameParameter) ProtoMessage()    {}
func (*GameParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_1617187b855ec876, []int{0}
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

func (m *GameParameter) GetBoardWidth() uint32 {
	if m != nil {
		return m.BoardWidth
	}
	return 0
}

func (m *GameParameter) GetBoardHeight() uint32 {
	if m != nil {
		return m.BoardHeight
	}
	return 0
}

// The GameTurn packet defines the two occupied positions of the played domino piece.
// These positions should be adjacent (a valid domino piece), otherwise the server will respond with an invalid turn status.
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
	return fileDescriptor_1617187b855ec876, []int{1}
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

// The GameState packet defines the current board state of the match.
type GameState struct {
	BoardWidth  uint32 `protobuf:"varint,1,opt,name=board_width,json=boardWidth,proto3" json:"board_width,omitempty"`
	BoardHeight uint32 `protobuf:"varint,2,opt,name=board_height,json=boardHeight,proto3" json:"board_height,omitempty"`
	// Byte array containing the current board information.
	// Each byte/character corresponds to a field state.
	// Ascii '0' ... Field unoccupied.
	// Ascii '1' ... Field belongs to the first player.
	// Ascii '2' ... Field belongs to the second player.
	// The overlying GameState packet defines "beginning_player", which is set to true, if you are the first player.
	// Fields are ordered row-after-row inside the byte array "board_data".
	// The first field of "board_data" corresponds to the position (x|y)->(0|0), the second to (x|y)->(1|0).
	BoardData            []byte   `protobuf:"bytes,3,opt,name=board_data,json=boardData,proto3" json:"board_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameState) Reset()         { *m = GameState{} }
func (m *GameState) String() string { return proto.CompactTextString(m) }
func (*GameState) ProtoMessage()    {}
func (*GameState) Descriptor() ([]byte, []int) {
	return fileDescriptor_1617187b855ec876, []int{2}
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

func (m *GameState) GetBoardWidth() uint32 {
	if m != nil {
		return m.BoardWidth
	}
	return 0
}

func (m *GameState) GetBoardHeight() uint32 {
	if m != nil {
		return m.BoardHeight
	}
	return 0
}

func (m *GameState) GetBoardData() []byte {
	if m != nil {
		return m.BoardData
	}
	return nil
}

func init() {
	proto.RegisterType((*GameParameter)(nil), "dom.GameParameter")
	proto.RegisterType((*GameTurn)(nil), "dom.GameTurn")
	proto.RegisterType((*GameState)(nil), "dom.GameState")
}

func init() {
	proto.RegisterFile("dominect.proto", fileDescriptor_1617187b855ec876)
}

var fileDescriptor_1617187b855ec876 = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x90, 0xcf, 0x4a, 0xc4, 0x30,
	0x10, 0xc6, 0xd9, 0xac, 0x88, 0x3b, 0xee, 0x2e, 0x18, 0x2f, 0xb9, 0x88, 0xda, 0x93, 0x27, 0x97,
	0xc6, 0x37, 0x10, 0x41, 0xf1, 0x24, 0xbb, 0x82, 0xe0, 0xa5, 0xcc, 0x9a, 0xa1, 0x2d, 0x25, 0x4d,
	0x09, 0x23, 0xb6, 0x6f, 0x2f, 0x99, 0xe6, 0x11, 0xf6, 0xf8, 0xfd, 0xbe, 0x3f, 0x09, 0x03, 0x5b,
	0x17, 0x7c, 0xdb, 0xd3, 0x0f, 0x3f, 0x0e, 0x31, 0x70, 0xd0, 0x4b, 0x17, 0x7c, 0x71, 0x80, 0xcd,
	0x2b, 0x7a, 0xfa, 0xc0, 0x88, 0x9e, 0x98, 0xa2, 0xbe, 0x85, 0xcb, 0x63, 0xc0, 0xe8, 0xaa, 0xbf,
	0xd6, 0x71, 0x63, 0x16, 0x77, 0x8b, 0x87, 0xcd, 0x1e, 0x04, 0x7d, 0x25, 0xa2, 0xef, 0x61, 0x3d,
	0x07, 0x1a, 0x6a, 0xeb, 0x86, 0x8d, 0x92, 0xc4, 0x5c, 0x7a, 0x13, 0x54, 0xbc, 0xc3, 0x45, 0x1a,
	0xfd, 0xfc, 0x8d, 0xbd, 0xde, 0x82, 0x1a, 0xcb, 0x3c, 0xa3, 0xc6, 0x32, 0xe9, 0xa9, 0xcc, 0x25,
	0x35, 0x89, 0x1e, 0xad, 0x59, 0x66, 0xdf, 0x8a, 0x6f, 0xcd, 0x59, 0xf6, 0x6d, 0xd1, 0xc3, 0x2a,
	0x6d, 0x1d, 0x18, 0x99, 0x4e, 0xf1, 0x39, 0x7d, 0x03, 0x73, 0xa1, 0x72, 0xc8, 0x28, 0x0f, 0xaf,
	0xf7, 0x2b, 0x21, 0x2f, 0xc8, 0xf8, 0x7c, 0xfd, 0x7d, 0xc5, 0xd4, 0x85, 0xaa, 0x46, 0x4f, 0xbb,
	0xa1, 0xab, 0x77, 0x2e, 0xf8, 0xe3, 0xb9, 0x5c, 0xec, 0xe9, 0x3f, 0x00, 0x00, 0xff, 0xff, 0xa9,
	0x89, 0x5c, 0xc9, 0x43, 0x01, 0x00, 0x00,
}
