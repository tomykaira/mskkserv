// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mozc/converter/segmenter_data.proto

/*
Package converter is a generated protocol buffer package.

It is generated from these files:
	mozc/converter/segmenter_data.proto

It has these top-level messages:
	SegmenterDataSizeInfo
*/
package converter

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

type SegmenterDataSizeInfo struct {
	CompressedLsize  *uint64 `protobuf:"varint,1,opt,name=compressed_lsize,json=compressedLsize" json:"compressed_lsize,omitempty"`
	CompressedRsize  *uint64 `protobuf:"varint,2,opt,name=compressed_rsize,json=compressedRsize" json:"compressed_rsize,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SegmenterDataSizeInfo) Reset()                    { *m = SegmenterDataSizeInfo{} }
func (m *SegmenterDataSizeInfo) String() string            { return proto.CompactTextString(m) }
func (*SegmenterDataSizeInfo) ProtoMessage()               {}
func (*SegmenterDataSizeInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SegmenterDataSizeInfo) GetCompressedLsize() uint64 {
	if m != nil && m.CompressedLsize != nil {
		return *m.CompressedLsize
	}
	return 0
}

func (m *SegmenterDataSizeInfo) GetCompressedRsize() uint64 {
	if m != nil && m.CompressedRsize != nil {
		return *m.CompressedRsize
	}
	return 0
}

func init() {
	proto.RegisterType((*SegmenterDataSizeInfo)(nil), "mozc.converter.SegmenterDataSizeInfo")
}

func init() { proto.RegisterFile("mozc/converter/segmenter_data.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0xcf, 0xb1, 0x0a, 0xc2, 0x40,
	0x0c, 0x80, 0x61, 0x2a, 0x4e, 0x1d, 0x54, 0x0a, 0x82, 0xa3, 0xe8, 0xa2, 0x20, 0x77, 0xef, 0x20,
	0x2e, 0x82, 0x53, 0xbb, 0xb9, 0x94, 0xf3, 0x1a, 0xeb, 0x51, 0xd3, 0x94, 0x24, 0x16, 0xec, 0xd3,
	0x4b, 0x05, 0x15, 0xbb, 0xfe, 0x7c, 0xcb, 0x1f, 0xaf, 0x91, 0x3a, 0x6f, 0x3d, 0xd5, 0x2d, 0xb0,
	0x02, 0x5b, 0x81, 0x12, 0xa1, 0x56, 0xe0, 0xbc, 0x70, 0xea, 0x4c, 0xc3, 0xa4, 0x94, 0x4c, 0x7a,
	0x64, 0xbe, 0x68, 0x85, 0xf1, 0x3c, 0xfb, 0xb8, 0x83, 0x53, 0x97, 0x85, 0x0e, 0x8e, 0xf5, 0x95,
	0x92, 0x6d, 0x3c, 0xf3, 0x84, 0x0d, 0x83, 0x08, 0x14, 0xf9, 0x5d, 0x42, 0x07, 0x8b, 0x68, 0x19,
	0x6d, 0xc6, 0xe9, 0xf4, 0xd7, 0x4f, 0x7d, 0x1e, 0x50, 0x7e, 0xd3, 0xd1, 0x90, 0xa6, 0x7d, 0xde,
	0x9b, 0xf3, 0xae, 0x0c, 0x7a, 0x7b, 0x5c, 0x8c, 0x27, 0xb4, 0x4a, 0xf8, 0xac, 0x5c, 0x60, 0x67,
	0x51, 0xaa, 0x4a, 0x80, 0x5b, 0xfb, 0xff, 0xf0, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x60, 0x91, 0x47,
	0x3d, 0xd4, 0x00, 0x00, 0x00,
}
