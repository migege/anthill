// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/migege/anthill/proto/log/log.proto

/*
Package migege_anthill_log is a generated protocol buffer package.

It is generated from these files:
	github.com/migege/anthill/proto/log/log.proto

It has these top-level messages:
	LogRequest
	LogResponse
*/
package migege_anthill_log

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type LogRequest struct {
	Info string `protobuf:"bytes,1,opt,name=info" json:"info,omitempty"`
}

func (m *LogRequest) Reset()                    { *m = LogRequest{} }
func (m *LogRequest) String() string            { return proto.CompactTextString(m) }
func (*LogRequest) ProtoMessage()               {}
func (*LogRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LogRequest) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

type LogResponse struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *LogResponse) Reset()                    { *m = LogResponse{} }
func (m *LogResponse) String() string            { return proto.CompactTextString(m) }
func (*LogResponse) ProtoMessage()               {}
func (*LogResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LogResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*LogRequest)(nil), "migege.anthill.log.LogRequest")
	proto.RegisterType((*LogResponse)(nil), "migege.anthill.log.LogResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Logger service

type LoggerClient interface {
	Log(ctx context.Context, in *LogRequest, opts ...client.CallOption) (*LogResponse, error)
}

type loggerClient struct {
	c           client.Client
	serviceName string
}

func NewLoggerClient(serviceName string, c client.Client) LoggerClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "migege.anthill.log"
	}
	return &loggerClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *loggerClient) Log(ctx context.Context, in *LogRequest, opts ...client.CallOption) (*LogResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Logger.Log", in)
	out := new(LogResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Logger service

type LoggerHandler interface {
	Log(context.Context, *LogRequest, *LogResponse) error
}

func RegisterLoggerHandler(s server.Server, hdlr LoggerHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Logger{hdlr}, opts...))
}

type Logger struct {
	LoggerHandler
}

func (h *Logger) Log(ctx context.Context, in *LogRequest, out *LogResponse) error {
	return h.LoggerHandler.Log(ctx, in, out)
}

func init() { proto.RegisterFile("github.com/migege/anthill/proto/log/log.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8e, 0xb1, 0xca, 0xc2, 0x40,
	0x0c, 0x80, 0xff, 0xd2, 0x9f, 0x82, 0x71, 0x91, 0x4c, 0xe2, 0x60, 0x4b, 0x27, 0x17, 0x73, 0xa0,
	0x2f, 0xe1, 0xd0, 0xa9, 0x6f, 0x60, 0x25, 0xa6, 0x07, 0xd7, 0xa6, 0xf6, 0xae, 0xef, 0x2f, 0xbd,
	0x16, 0x1c, 0xc4, 0x21, 0xf0, 0x85, 0x7c, 0xf0, 0x05, 0xce, 0x62, 0x43, 0x3b, 0x35, 0xf4, 0xd0,
	0xce, 0x74, 0x56, 0x58, 0xd8, 0xdc, 0xfb, 0xd0, 0x5a, 0xe7, 0xcc, 0x30, 0x6a, 0x50, 0xe3, 0x54,
	0xe6, 0xa1, 0xb8, 0x21, 0x2e, 0x0e, 0xad, 0x0e, 0x39, 0x95, 0xb2, 0x00, 0xa8, 0x54, 0x6a, 0x7e,
	0x4d, 0xec, 0x03, 0x22, 0xfc, 0xdb, 0xfe, 0xa9, 0xfb, 0xa4, 0x48, 0x4e, 0x9b, 0x3a, 0x72, 0x99,
	0xc3, 0x36, 0x1a, 0x7e, 0xd0, 0xde, 0x33, 0xee, 0x20, 0xed, 0xbc, 0xac, 0xc6, 0x8c, 0x97, 0x1a,
	0xb2, 0x4a, 0x45, 0x78, 0xc4, 0x1b, 0xa4, 0x95, 0x0a, 0x1e, 0xe9, 0x3b, 0x44, 0x9f, 0xca, 0x21,
	0xff, 0x79, 0x5f, 0x1a, 0xe5, 0x5f, 0x93, 0xc5, 0x8f, 0xaf, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xef, 0xfc, 0x34, 0xdc, 0xe2, 0x00, 0x00, 0x00,
}
