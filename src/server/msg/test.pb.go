// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

/*
Package msg is a generated protocol buffer package.

It is generated from these files:
	test.proto

It has these top-level messages:
	LoginMsg
	ChatMsg
	RoomData
	PlayerInfo
*/
package msg

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

type LoginMsg struct {
	Name             *string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Key              *string `protobuf:"bytes,2,opt,name=Key" json:"Key,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *LoginMsg) Reset()                    { *m = LoginMsg{} }
func (m *LoginMsg) String() string            { return proto.CompactTextString(m) }
func (*LoginMsg) ProtoMessage()               {}
func (*LoginMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LoginMsg) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *LoginMsg) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

type ChatMsg struct {
	Key              *string `protobuf:"bytes,1,opt,name=Key" json:"Key,omitempty"`
	Name             *string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	Data             *string `protobuf:"bytes,3,opt,name=Data" json:"Data,omitempty"`
	Photo            []byte  `protobuf:"bytes,4,opt,name=photo" json:"photo,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ChatMsg) Reset()                    { *m = ChatMsg{} }
func (m *ChatMsg) String() string            { return proto.CompactTextString(m) }
func (*ChatMsg) ProtoMessage()               {}
func (*ChatMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ChatMsg) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *ChatMsg) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ChatMsg) GetData() string {
	if m != nil && m.Data != nil {
		return *m.Data
	}
	return ""
}

func (m *ChatMsg) GetPhoto() []byte {
	if m != nil {
		return m.Photo
	}
	return nil
}

type RoomData struct {
	OnlineCount      *int32        `protobuf:"varint,1,opt,name=OnlineCount" json:"OnlineCount,omitempty"`
	Players          []*PlayerInfo `protobuf:"bytes,2,rep,name=Players" json:"Players,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *RoomData) Reset()                    { *m = RoomData{} }
func (m *RoomData) String() string            { return proto.CompactTextString(m) }
func (*RoomData) ProtoMessage()               {}
func (*RoomData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RoomData) GetOnlineCount() int32 {
	if m != nil && m.OnlineCount != nil {
		return *m.OnlineCount
	}
	return 0
}

func (m *RoomData) GetPlayers() []*PlayerInfo {
	if m != nil {
		return m.Players
	}
	return nil
}

type PlayerInfo struct {
	Name             *string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Online           *bool   `protobuf:"varint,2,opt,name=online" json:"online,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PlayerInfo) Reset()                    { *m = PlayerInfo{} }
func (m *PlayerInfo) String() string            { return proto.CompactTextString(m) }
func (*PlayerInfo) ProtoMessage()               {}
func (*PlayerInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PlayerInfo) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *PlayerInfo) GetOnline() bool {
	if m != nil && m.Online != nil {
		return *m.Online
	}
	return false
}

func init() {
	proto.RegisterType((*LoginMsg)(nil), "msg.LoginMsg")
	proto.RegisterType((*ChatMsg)(nil), "msg.ChatMsg")
	proto.RegisterType((*RoomData)(nil), "msg.RoomData")
	proto.RegisterType((*PlayerInfo)(nil), "msg.PlayerInfo")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0x3f, 0x4f, 0x85, 0x30,
	0x14, 0xc5, 0x53, 0xca, 0xf3, 0xe1, 0xc5, 0x44, 0xd3, 0x18, 0xd3, 0xb1, 0xe9, 0x84, 0x0b, 0x31,
	0x4e, 0xee, 0xb8, 0x18, 0xff, 0xa6, 0x8b, 0x71, 0xec, 0x50, 0x0b, 0x09, 0xed, 0x25, 0xb4, 0x0e,
	0x7c, 0x7b, 0x43, 0x41, 0x64, 0x78, 0xdb, 0x39, 0xbf, 0xd3, 0x9e, 0x7b, 0x2f, 0x40, 0x34, 0x21,
	0xd6, 0xc3, 0x88, 0x11, 0x19, 0x75, 0xc1, 0xca, 0x3b, 0x28, 0x5e, 0xd0, 0x76, 0xfe, 0x35, 0x58,
	0xc6, 0x20, 0x7f, 0xd3, 0xce, 0x70, 0x22, 0x48, 0x75, 0xae, 0x92, 0x66, 0x57, 0x40, 0x9f, 0xcd,
	0xc4, 0xb3, 0x84, 0x66, 0x29, 0xbf, 0xe0, 0xd8, 0xb4, 0x3a, 0xce, 0x1f, 0xd6, 0x90, 0x6c, 0xe1,
	0x56, 0x91, 0xed, 0x2a, 0x18, 0xe4, 0x8f, 0x3a, 0x6a, 0x4e, 0x17, 0x36, 0x6b, 0x76, 0x0d, 0x87,
	0xa1, 0xc5, 0x88, 0x3c, 0x17, 0xa4, 0xba, 0x50, 0x8b, 0x91, 0x9f, 0x50, 0x28, 0x44, 0x97, 0x5e,
	0x08, 0x28, 0xdf, 0x7d, 0xdf, 0x79, 0xd3, 0xe0, 0x8f, 0x8f, 0x69, 0xc6, 0x41, 0xed, 0x11, 0xbb,
	0x85, 0xe3, 0x47, 0xaf, 0x27, 0x33, 0x06, 0x9e, 0x09, 0x5a, 0x95, 0xf7, 0x97, 0xb5, 0x0b, 0xb6,
	0x5e, 0xd8, 0x93, 0xff, 0x46, 0xf5, 0x97, 0xcb, 0x07, 0x80, 0x7f, 0x7c, 0xf2, 0xce, 0x1b, 0x38,
	0xc3, 0xd4, 0x9d, 0x56, 0x2f, 0xd4, 0xea, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x44, 0x72, 0x87,
	0x3b, 0x31, 0x01, 0x00, 0x00,
}
