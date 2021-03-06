// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/example/example.proto

package go_micro_srv_postRet

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// web -> srv
type Request struct {
	// 手机号
	Mobile string `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	// 密码
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	// 短信验证码
	SmsCode              string   `protobuf:"bytes,3,opt,name=sms_code,json=smsCode,proto3" json:"sms_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{0}
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

func (m *Request) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *Request) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Request) GetSmsCode() string {
	if m != nil {
		return m.SmsCode
	}
	return ""
}

// srv -> web
type Response struct {
	// 错误码
	Errno string `protobuf:"bytes,1,opt,name=errno,proto3" json:"errno,omitempty"`
	// 错误信息
	Errmsg string `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	// sessionId
	SessionId            string   `protobuf:"bytes,3,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{1}
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

func (m *Response) GetErrno() string {
	if m != nil {
		return m.Errno
	}
	return ""
}

func (m *Response) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func (m *Response) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.srv.postRet.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.postRet.Response")
}

func init() { proto.RegisterFile("proto/example/example.proto", fileDescriptor_097b3f5db5cf5789) }

var fileDescriptor_097b3f5db5cf5789 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xcb, 0x4a, 0x04, 0x31,
	0x10, 0x45, 0x1d, 0xc5, 0x49, 0x4f, 0x2d, 0x8b, 0x41, 0xda, 0xf1, 0x81, 0xf4, 0xca, 0x55, 0x04,
	0xfd, 0x04, 0x71, 0xa1, 0x2b, 0x09, 0x28, 0xee, 0x64, 0x66, 0x52, 0x34, 0x0d, 0x9d, 0xae, 0x58,
	0x15, 0x1f, 0x9f, 0x2f, 0x26, 0x51, 0x37, 0xba, 0x0a, 0xe7, 0x5e, 0xb8, 0xc9, 0x09, 0x1c, 0x45,
	0xe1, 0xc4, 0x17, 0xf4, 0xb1, 0x0e, 0x71, 0xa4, 0xef, 0xd3, 0xe6, 0x14, 0x97, 0x3d, 0xdb, 0x30,
	0x6c, 0x85, 0xad, 0xca, 0x9b, 0x8d, 0xac, 0xc9, 0x51, 0xea, 0x9e, 0xc0, 0x38, 0x7a, 0x79, 0x25,
	0x4d, 0x78, 0x00, 0xf3, 0xc0, 0x9b, 0x61, 0xa4, 0x76, 0x76, 0x36, 0x3b, 0x5f, 0xb8, 0x4a, 0xb8,
	0x82, 0x26, 0xae, 0x55, 0xdf, 0x59, 0x7c, 0xbb, 0x9b, 0x9b, 0x1f, 0xc6, 0x43, 0x68, 0x34, 0xe8,
	0xf3, 0x96, 0x3d, 0xb5, 0x7b, 0xb9, 0x33, 0x1a, 0xf4, 0x9a, 0x3d, 0x75, 0x8f, 0xd0, 0x38, 0xd2,
	0xc8, 0x93, 0x12, 0x2e, 0x61, 0x9f, 0x44, 0x26, 0xae, 0xcb, 0x05, 0xbe, 0x2e, 0x24, 0x91, 0xa0,
	0x7d, 0x9d, 0xad, 0x84, 0xc7, 0xb0, 0x50, 0x52, 0x1d, 0x78, 0xba, 0xf5, 0x75, 0xf5, 0x37, 0xb8,
	0x7c, 0x00, 0x73, 0x53, 0xc4, 0xf0, 0x0e, 0xcc, 0x7d, 0xf1, 0xc0, 0x13, 0xfb, 0x97, 0x9e, 0xad,
	0x6e, 0xab, 0xd3, 0xff, 0xea, 0xf2, 0xc0, 0x6e, 0x67, 0x33, 0xcf, 0xbf, 0x74, 0xf5, 0x19, 0x00,
	0x00, 0xff, 0xff, 0xd4, 0x5e, 0x7c, 0xfc, 0x44, 0x01, 0x00, 0x00,
}
