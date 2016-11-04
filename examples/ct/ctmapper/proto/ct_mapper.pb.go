// Code generated by protoc-gen-go.
// source: github.com/google/trillian/examples/ct/ctmapper/proto/ct_mapper.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	github.com/google/trillian/examples/ct/ctmapper/proto/ct_mapper.proto

It has these top-level messages:
	EntryList
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

type EntryList struct {
	Domain       string  `protobuf:"bytes,1,opt,name=domain" json:"domain,omitempty"`
	CertIndex    []int64 `protobuf:"varint,2,rep,packed,name=cert_index,json=certIndex" json:"cert_index,omitempty"`
	PrecertIndex []int64 `protobuf:"varint,3,rep,packed,name=precert_index,json=precertIndex" json:"precert_index,omitempty"`
}

func (m *EntryList) Reset()                    { *m = EntryList{} }
func (m *EntryList) String() string            { return proto1.CompactTextString(m) }
func (*EntryList) ProtoMessage()               {}
func (*EntryList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto1.RegisterType((*EntryList)(nil), "proto.EntryList")
}

func init() {
	proto1.RegisterFile("github.com/google/trillian/examples/ct/ctmapper/proto/ct_mapper.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x8d, 0xbf, 0xca, 0x83, 0x30,
	0x14, 0x47, 0xf1, 0x93, 0x4f, 0x48, 0x68, 0x97, 0x0c, 0xc5, 0xa5, 0x20, 0xed, 0xe2, 0xe4, 0x1d,
	0xfa, 0x0c, 0x0e, 0x85, 0x4e, 0xbe, 0x80, 0xc4, 0x78, 0x49, 0x03, 0xf9, 0x47, 0xbc, 0x05, 0xfb,
	0xf6, 0xc5, 0xd4, 0xa1, 0xd3, 0x8f, 0x73, 0x38, 0xf0, 0xe3, 0xbd, 0x36, 0xf4, 0x7c, 0x4d, 0x9d,
	0x0a, 0x0e, 0x74, 0x08, 0xda, 0x22, 0x50, 0x32, 0xd6, 0x1a, 0xe9, 0x01, 0x57, 0xe9, 0xa2, 0xc5,
	0x05, 0x14, 0x81, 0x22, 0x27, 0x63, 0xc4, 0x04, 0x31, 0x05, 0x0a, 0xa0, 0x68, 0xfc, 0x72, 0x97,
	0x59, 0xfc, 0xe7, 0xb9, 0x68, 0xce, 0x7a, 0x4f, 0xe9, 0xfd, 0x30, 0x0b, 0x89, 0x13, 0xaf, 0xe6,
	0xe0, 0xa4, 0xf1, 0x75, 0xd1, 0x14, 0x2d, 0x1b, 0x76, 0x12, 0x67, 0xce, 0x15, 0x26, 0x1a, 0x8d,
	0x9f, 0x71, 0xad, 0xff, 0x9a, 0xb2, 0x2d, 0x07, 0xb6, 0x99, 0xfb, 0x26, 0xc4, 0x95, 0x1f, 0x63,
	0xc2, 0x9f, 0xa2, 0xcc, 0xc5, 0x61, 0x97, 0x39, 0x9a, 0xaa, 0xfc, 0x77, 0xfb, 0x04, 0x00, 0x00,
	0xff, 0xff, 0xd5, 0xf4, 0x5d, 0xfe, 0xbf, 0x00, 0x00, 0x00,
}