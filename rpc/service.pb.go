// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/service.proto

package play

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

type UserTrack struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TrackId              string   `protobuf:"bytes,2,opt,name=track_id,json=trackId,proto3" json:"track_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserTrack) Reset()         { *m = UserTrack{} }
func (m *UserTrack) String() string { return proto.CompactTextString(m) }
func (*UserTrack) ProtoMessage()    {}
func (*UserTrack) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_d916f37e90111965, []int{0}
}
func (m *UserTrack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserTrack.Unmarshal(m, b)
}
func (m *UserTrack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserTrack.Marshal(b, m, deterministic)
}
func (dst *UserTrack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserTrack.Merge(dst, src)
}
func (m *UserTrack) XXX_Size() int {
	return xxx_messageInfo_UserTrack.Size(m)
}
func (m *UserTrack) XXX_DiscardUnknown() {
	xxx_messageInfo_UserTrack.DiscardUnknown(m)
}

var xxx_messageInfo_UserTrack proto.InternalMessageInfo

func (m *UserTrack) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UserTrack) GetTrackId() string {
	if m != nil {
		return m.TrackId
	}
	return ""
}

type TrackData struct {
	TrackServerId        string   `protobuf:"bytes,1,opt,name=track_server_id,json=trackServerId,proto3" json:"track_server_id,omitempty"`
	StartPosition        int32    `protobuf:"varint,2,opt,name=start_position,json=startPosition,proto3" json:"start_position,omitempty"`
	NumBytes             int32    `protobuf:"varint,3,opt,name=num_bytes,json=numBytes,proto3" json:"num_bytes,omitempty"`
	Data                 []byte   `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrackData) Reset()         { *m = TrackData{} }
func (m *TrackData) String() string { return proto.CompactTextString(m) }
func (*TrackData) ProtoMessage()    {}
func (*TrackData) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_d916f37e90111965, []int{1}
}
func (m *TrackData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrackData.Unmarshal(m, b)
}
func (m *TrackData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrackData.Marshal(b, m, deterministic)
}
func (dst *TrackData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrackData.Merge(dst, src)
}
func (m *TrackData) XXX_Size() int {
	return xxx_messageInfo_TrackData.Size(m)
}
func (m *TrackData) XXX_DiscardUnknown() {
	xxx_messageInfo_TrackData.DiscardUnknown(m)
}

var xxx_messageInfo_TrackData proto.InternalMessageInfo

func (m *TrackData) GetTrackServerId() string {
	if m != nil {
		return m.TrackServerId
	}
	return ""
}

func (m *TrackData) GetStartPosition() int32 {
	if m != nil {
		return m.StartPosition
	}
	return 0
}

func (m *TrackData) GetNumBytes() int32 {
	if m != nil {
		return m.NumBytes
	}
	return 0
}

func (m *TrackData) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_d916f37e90111965, []int{2}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*UserTrack)(nil), "resonate.api.play.UserTrack")
	proto.RegisterType((*TrackData)(nil), "resonate.api.play.TrackData")
	proto.RegisterType((*Empty)(nil), "resonate.api.play.Empty")
}

func init() { proto.RegisterFile("rpc/service.proto", fileDescriptor_service_d916f37e90111965) }

var fileDescriptor_service_d916f37e90111965 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x89, 0xa6, 0x49, 0x33, 0x5a, 0xa5, 0x73, 0x31, 0xfe, 0x39, 0x94, 0x80, 0xd2, 0x53,
	0x14, 0xfd, 0x00, 0x42, 0xa9, 0x87, 0xdc, 0x4a, 0xa2, 0x17, 0x2f, 0x61, 0x9a, 0xec, 0x21, 0xd8,
	0x24, 0xcb, 0xee, 0x44, 0xc8, 0x67, 0xf0, 0x4b, 0xcb, 0x4e, 0xa5, 0x08, 0x7a, 0xdb, 0xf7, 0x7b,
	0x6f, 0x67, 0x78, 0x03, 0x73, 0xa3, 0xab, 0x7b, 0xab, 0xcc, 0x67, 0x53, 0xa9, 0x54, 0x9b, 0x9e,
	0x7b, 0x9c, 0x1b, 0x65, 0xfb, 0x8e, 0x58, 0xa5, 0xa4, 0x9b, 0x54, 0xef, 0x68, 0x4c, 0x9e, 0x21,
	0x7a, 0xb3, 0xca, 0xbc, 0x1a, 0xaa, 0x3e, 0xf0, 0x02, 0xc2, 0xc1, 0x2a, 0x53, 0x36, 0x75, 0xec,
	0x2d, 0xbc, 0x65, 0x94, 0x07, 0x4e, 0x66, 0x35, 0x5e, 0xc2, 0x94, 0x5d, 0xc2, 0x39, 0x47, 0xe2,
	0x84, 0xa2, 0xb3, 0x3a, 0xf9, 0xf2, 0x20, 0x92, 0xdf, 0x6b, 0x62, 0xc2, 0x3b, 0x38, 0xdf, 0x07,
	0xdd, 0xe2, 0xdf, 0x93, 0x66, 0x82, 0x0b, 0xa1, 0x59, 0x8d, 0xb7, 0x70, 0x66, 0x99, 0x0c, 0x97,
	0xba, 0xb7, 0x0d, 0x37, 0x7d, 0x27, 0x63, 0x27, 0xf9, 0x4c, 0xe8, 0xe6, 0x07, 0xe2, 0x35, 0x44,
	0xdd, 0xd0, 0x96, 0xdb, 0x91, 0x95, 0x8d, 0x8f, 0x25, 0x31, 0xed, 0x86, 0x76, 0xe5, 0x34, 0x22,
	0xf8, 0x35, 0x31, 0xc5, 0xfe, 0xc2, 0x5b, 0x9e, 0xe6, 0xf2, 0x4e, 0x42, 0x98, 0xbc, 0xb4, 0x9a,
	0xc7, 0xc7, 0x02, 0x4e, 0x36, 0x3b, 0x1a, 0x8b, 0x7d, 0x7f, 0x5c, 0x83, 0xef, 0x24, 0xde, 0xa4,
	0x7f, 0x4e, 0x90, 0x1e, 0xfa, 0x5f, 0xfd, 0xe7, 0x1e, 0xba, 0x3d, 0x78, 0xab, 0xe0, 0xdd, 0x77,
	0x6c, 0x1b, 0xc8, 0x39, 0x9f, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0x44, 0x3f, 0x28, 0x0e, 0x63,
	0x01, 0x00, 0x00,
}