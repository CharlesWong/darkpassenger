// Code generated by protoc-gen-go.
// source: util.proto
// DO NOT EDIT!

package model

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type NullMessage struct {
}

func (m *NullMessage) Reset()                    { *m = NullMessage{} }
func (m *NullMessage) String() string            { return proto.CompactTextString(m) }
func (*NullMessage) ProtoMessage()               {}
func (*NullMessage) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func init() {
	proto.RegisterType((*NullMessage)(nil), "model.NullMessage")
}

var fileDescriptor2 = []byte{
	// 63 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0xc9, 0xcc,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcd, 0xcd, 0x4f, 0x49, 0xcd, 0x51, 0xe2, 0xe5,
	0xe2, 0xf6, 0x2b, 0xcd, 0xc9, 0xf1, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x4d, 0x62, 0x03, 0x4b,
	0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x26, 0x36, 0x68, 0x79, 0x2a, 0x00, 0x00, 0x00,
}