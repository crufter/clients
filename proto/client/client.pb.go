// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/micro/clients/proto/client/client.proto

package go_micro_api_client

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

type Request struct {
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Endpoint             string   `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	ContentType          string   `protobuf:"bytes,3,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	Body                 []byte   `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_81c8aad97b6341e5, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *Request) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func (m *Request) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *Request) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type Response struct {
	Body                 []byte   `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_81c8aad97b6341e5, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type Message struct {
	Topic                string   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	ContentType          string   `protobuf:"bytes,2,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	Body                 []byte   `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_81c8aad97b6341e5, []int{2}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *Message) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *Message) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.api.client.Request")
	proto.RegisterType((*Response)(nil), "go.micro.api.client.Response")
	proto.RegisterType((*Message)(nil), "go.micro.api.client.Message")
}

func init() {
	proto.RegisterFile("github.com/micro/clients/proto/client/client.proto", fileDescriptor_81c8aad97b6341e5)
}

var fileDescriptor_81c8aad97b6341e5 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x40, 0x71, 0x1b, 0x92, 0x72, 0x74, 0x32, 0x0c, 0x56, 0x05, 0xa8, 0x64, 0xca, 0xe4, 0xa0,
	0xf2, 0x09, 0x15, 0x13, 0x62, 0x09, 0x88, 0x15, 0x25, 0xee, 0x29, 0x58, 0x4a, 0x6c, 0x13, 0x5f,
	0x2b, 0xe5, 0x5f, 0xf8, 0x58, 0x24, 0x27, 0xa5, 0x48, 0xd0, 0x89, 0xc9, 0x7e, 0xf7, 0xac, 0xbb,
	0xf3, 0x1d, 0xac, 0x6a, 0x4d, 0xef, 0xdb, 0x4a, 0x2a, 0xdb, 0xe6, 0xad, 0x56, 0x9d, 0xcd, 0x55,
	0xa3, 0xd1, 0x90, 0xcf, 0x5d, 0x67, 0x69, 0x4f, 0xe3, 0x21, 0x43, 0x8c, 0x5f, 0xd4, 0x56, 0x86,
	0xb7, 0xb2, 0x74, 0x5a, 0x0e, 0x2a, 0xdd, 0x41, 0x52, 0xe0, 0xc7, 0x16, 0x3d, 0x71, 0x01, 0x89,
	0xc7, 0x6e, 0xa7, 0x15, 0x0a, 0xb6, 0x64, 0xd9, 0x59, 0xb1, 0x47, 0xbe, 0x80, 0x19, 0x9a, 0x8d,
	0xb3, 0xda, 0x90, 0x98, 0x04, 0xf5, 0xcd, 0xfc, 0x16, 0xe6, 0xca, 0x1a, 0x42, 0x43, 0x6f, 0xd4,
	0x3b, 0x14, 0xd3, 0xe0, 0xcf, 0xc7, 0xd8, 0x4b, 0xef, 0x90, 0x73, 0x88, 0x2a, 0xbb, 0xe9, 0x45,
	0xb4, 0x64, 0xd9, 0xbc, 0x08, 0xf7, 0xf4, 0x06, 0x66, 0x05, 0x7a, 0x67, 0x8d, 0x3f, 0x78, 0xf6,
	0xc3, 0xbf, 0x42, 0xf2, 0x84, 0xde, 0x97, 0x35, 0xf2, 0x4b, 0x38, 0x25, 0xeb, 0xb4, 0x1a, 0xbb,
	0x1a, 0xe0, 0x57, 0xdd, 0xc9, 0xf1, 0xba, 0xd3, 0x43, 0xde, 0xd5, 0x27, 0x83, 0x78, 0x1d, 0xbe,
	0xce, 0x1f, 0x20, 0x5a, 0x97, 0x4d, 0xc3, 0xaf, 0xe4, 0x1f, 0x83, 0x91, 0xe3, 0x54, 0x16, 0xd7,
	0x47, 0xec, 0xd0, 0x7b, 0x7a, 0xc2, 0x1f, 0x21, 0x7e, 0xa6, 0x0e, 0xcb, 0xf6, 0x9f, 0x89, 0x32,
	0x76, 0xc7, 0xaa, 0x38, 0xac, 0xea, 0xfe, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x91, 0xbd, 0x35,
	0xe0, 0x01, 0x00, 0x00,
}
