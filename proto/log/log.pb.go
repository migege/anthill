// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/migege/anthill/proto/log/log.proto

/*
Package com_mayibot_ah_log is a generated protocol buffer package.

It is generated from these files:
	github.com/migege/anthill/proto/log/log.proto

It has these top-level messages:
	Response
	Info
	Profit
*/
package com_mayibot_ah_log

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

type Response struct {
	Code    int64  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Response) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Info struct {
	Info string `protobuf:"bytes,1,opt,name=info" json:"info,omitempty"`
	Ts   int64  `protobuf:"varint,2,opt,name=ts" json:"ts,omitempty"`
}

func (m *Info) Reset()                    { *m = Info{} }
func (m *Info) String() string            { return proto.CompactTextString(m) }
func (*Info) ProtoMessage()               {}
func (*Info) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Info) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func (m *Info) GetTs() int64 {
	if m != nil {
		return m.Ts
	}
	return 0
}

type Profit struct {
	Profit float64 `protobuf:"fixed64,1,opt,name=profit" json:"profit,omitempty"`
	Info   string  `protobuf:"bytes,2,opt,name=info" json:"info,omitempty"`
	Ts     int64   `protobuf:"varint,3,opt,name=ts" json:"ts,omitempty"`
}

func (m *Profit) Reset()                    { *m = Profit{} }
func (m *Profit) String() string            { return proto.CompactTextString(m) }
func (*Profit) ProtoMessage()               {}
func (*Profit) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Profit) GetProfit() float64 {
	if m != nil {
		return m.Profit
	}
	return 0
}

func (m *Profit) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func (m *Profit) GetTs() int64 {
	if m != nil {
		return m.Ts
	}
	return 0
}

func init() {
	proto.RegisterType((*Response)(nil), "com.mayibot.ah.log.Response")
	proto.RegisterType((*Info)(nil), "com.mayibot.ah.log.Info")
	proto.RegisterType((*Profit)(nil), "com.mayibot.ah.log.Profit")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Logger service

type LoggerClient interface {
	Log(ctx context.Context, in *Info, opts ...client.CallOption) (*Response, error)
	LogStatus(ctx context.Context, in *Info, opts ...client.CallOption) (*Response, error)
	LogProfit(ctx context.Context, in *Profit, opts ...client.CallOption) (*Response, error)
	Status(ctx context.Context, in *Info, opts ...client.CallOption) (Logger_StatusClient, error)
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
		serviceName = "com.mayibot.ah.log"
	}
	return &loggerClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *loggerClient) Log(ctx context.Context, in *Info, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Logger.Log", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loggerClient) LogStatus(ctx context.Context, in *Info, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Logger.LogStatus", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loggerClient) LogProfit(ctx context.Context, in *Profit, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Logger.LogProfit", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loggerClient) Status(ctx context.Context, in *Info, opts ...client.CallOption) (Logger_StatusClient, error) {
	req := c.c.NewRequest(c.serviceName, "Logger.Status", &Info{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &loggerStatusClient{stream}, nil
}

type Logger_StatusClient interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*Info, error)
}

type loggerStatusClient struct {
	stream client.Streamer
}

func (x *loggerStatusClient) Close() error {
	return x.stream.Close()
}

func (x *loggerStatusClient) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *loggerStatusClient) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *loggerStatusClient) Recv() (*Info, error) {
	m := new(Info)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Logger service

type LoggerHandler interface {
	Log(context.Context, *Info, *Response) error
	LogStatus(context.Context, *Info, *Response) error
	LogProfit(context.Context, *Profit, *Response) error
	Status(context.Context, *Info, Logger_StatusStream) error
}

func RegisterLoggerHandler(s server.Server, hdlr LoggerHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Logger{hdlr}, opts...))
}

type Logger struct {
	LoggerHandler
}

func (h *Logger) Log(ctx context.Context, in *Info, out *Response) error {
	return h.LoggerHandler.Log(ctx, in, out)
}

func (h *Logger) LogStatus(ctx context.Context, in *Info, out *Response) error {
	return h.LoggerHandler.LogStatus(ctx, in, out)
}

func (h *Logger) LogProfit(ctx context.Context, in *Profit, out *Response) error {
	return h.LoggerHandler.LogProfit(ctx, in, out)
}

func (h *Logger) Status(ctx context.Context, stream server.Streamer) error {
	m := new(Info)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.LoggerHandler.Status(ctx, m, &loggerStatusStream{stream})
}

type Logger_StatusStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Info) error
}

type loggerStatusStream struct {
	stream server.Streamer
}

func (x *loggerStatusStream) Close() error {
	return x.stream.Close()
}

func (x *loggerStatusStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *loggerStatusStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *loggerStatusStream) Send(m *Info) error {
	return x.stream.Send(m)
}

func init() { proto.RegisterFile("github.com/migege/anthill/proto/log/log.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0x31, 0x4b, 0xc4, 0x40,
	0x10, 0x85, 0x2f, 0xc9, 0x11, 0xcd, 0x14, 0x16, 0x53, 0x48, 0x38, 0x2c, 0x8e, 0x54, 0x87, 0xe0,
	0x46, 0xb4, 0xb1, 0xd3, 0x42, 0x11, 0x21, 0x85, 0xac, 0xbf, 0x20, 0x89, 0x9b, 0xc9, 0x42, 0x36,
	0x13, 0xb2, 0x7b, 0x85, 0xbf, 0xc5, 0x3f, 0x2b, 0x59, 0x2f, 0x36, 0x77, 0x72, 0xc5, 0x15, 0x0b,
	0xf3, 0x78, 0xf3, 0x3e, 0xe6, 0xb1, 0x70, 0x43, 0xda, 0xb5, 0xdb, 0x4a, 0xd4, 0x6c, 0x72, 0xa3,
	0x49, 0x91, 0xca, 0xcb, 0xde, 0xb5, 0xba, 0xeb, 0xf2, 0x61, 0x64, 0xc7, 0x79, 0xc7, 0x34, 0x3d,
	0xe1, 0x15, 0x62, 0xcd, 0x46, 0x98, 0xf2, 0x4b, 0x57, 0xec, 0x44, 0xd9, 0x8a, 0x8e, 0x29, 0x7b,
	0x80, 0x73, 0xa9, 0xec, 0xc0, 0xbd, 0x55, 0x88, 0xb0, 0xac, 0xf9, 0x53, 0xa5, 0xc1, 0x3a, 0xd8,
	0x44, 0xd2, 0xcf, 0x98, 0xc2, 0x99, 0x51, 0xd6, 0x96, 0xa4, 0xd2, 0x70, 0x1d, 0x6c, 0x12, 0x39,
	0xcb, 0xec, 0x1a, 0x96, 0x6f, 0x7d, 0xc3, 0x53, 0x4a, 0xf7, 0x0d, 0xfb, 0x54, 0x22, 0xfd, 0x8c,
	0x17, 0x10, 0x3a, 0xeb, 0x03, 0x91, 0x0c, 0x9d, 0xcd, 0x9e, 0x21, 0x7e, 0x1f, 0xb9, 0xd1, 0x0e,
	0x2f, 0x21, 0x1e, 0xfc, 0xe4, 0xf7, 0x03, 0xb9, 0x53, 0x7f, 0x94, 0x70, 0x8f, 0x12, 0xcd, 0x94,
	0xbb, 0xef, 0x10, 0xe2, 0x82, 0x89, 0xd4, 0x88, 0x8f, 0x10, 0x15, 0x4c, 0x98, 0x8a, 0xfd, 0x4a,
	0x62, 0xba, 0x6a, 0x75, 0x75, 0xc8, 0x99, 0x9b, 0x66, 0x0b, 0x7c, 0x81, 0xa4, 0x60, 0xfa, 0x70,
	0xa5, 0xdb, 0xda, 0x13, 0x30, 0xaf, 0x1e, 0xb3, 0xeb, 0xb6, 0x3a, 0xb4, 0xfc, 0xeb, 0x1d, 0x05,
	0x3d, 0x41, 0x7c, 0xf4, 0x98, 0x7f, 0x9d, 0x6c, 0x71, 0x1b, 0x54, 0xb1, 0xff, 0xe4, 0xfb, 0x9f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xe0, 0xd5, 0x5f, 0x54, 0x15, 0x02, 0x00, 0x00,
}
