// Code generated by protoc-gen-go.
// source: msg.proto
// DO NOT EDIT!*/

package protos

import "github.com/golang/protobuf/proto"

type Msg struct {
	MsgType int32  `protobuf:"varint,1,opt,name=msgType" json:"msgType,omitempty"`
	MsgInfo string `protobuf:"bytes,2,opt,name=msgInfo" json:"msgInfo,omitempty"`
	MsgFrom string `protobuf:"bytes,3,opt,name=msgFrom" json:"msgFrom,omitempty"`
}

func (m *Msg) Reset()                      { *m = Msg{} }
func (m *Msg) String() string              { return proto.CompactTextString(m) }
func (m *Msg) ProtoMessage()               {}
func (m *Msg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*Msg)(nil), "protos.msg")
	proto.RegisterFile("msg.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 97 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xcc, 0x2d, 0x4e, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x53, 0xc5, 0x4a, 0xc1, 0x5c, 0xcc, 0xb9, 0xc5,
	0xe9, 0x42, 0x12, 0x5c, 0xec, 0xb9, 0xc5, 0xe9, 0x21, 0x95, 0x05, 0xa9, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0xac, 0x41, 0x30, 0x2e, 0x54, 0xc6, 0x33, 0x2f, 0x2d, 0x5f, 0x82, 0x49, 0x81, 0x51, 0x83,
	0x33, 0x08, 0xc6, 0x85, 0xca, 0xb8, 0x15, 0xe5, 0xe7, 0x4a, 0x30, 0xc3, 0x65, 0x40, 0xdc, 0x24,
	0x88, 0xe1, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfc, 0x57, 0x67, 0x81, 0x70, 0x00, 0x00,
	0x00,
}
