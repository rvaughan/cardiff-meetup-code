// Code generated by protoc-gen-go.
// source: speak.proto
// DO NOT EDIT!

/*
Package speak is a generated protocol buffer package.

It is generated from these files:
	speak.proto

It has these top-level messages:
	SpeakEvent
	SpeakResponse
	VoiceResponse
	Empty
*/
package speak

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type SpeakEvent struct {
	Speech     string `protobuf:"bytes,1,opt,name=speech" json:"speech,omitempty"`
	Voice      string `protobuf:"bytes,2,opt,name=voice" json:"voice,omitempty"`
	SpeechRate int64  `protobuf:"varint,3,opt,name=speechRate" json:"speechRate,omitempty"`
}

func (m *SpeakEvent) Reset()                    { *m = SpeakEvent{} }
func (m *SpeakEvent) String() string            { return proto.CompactTextString(m) }
func (*SpeakEvent) ProtoMessage()               {}
func (*SpeakEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SpeakEvent) GetSpeech() string {
	if m != nil {
		return m.Speech
	}
	return ""
}

func (m *SpeakEvent) GetVoice() string {
	if m != nil {
		return m.Voice
	}
	return ""
}

func (m *SpeakEvent) GetSpeechRate() int64 {
	if m != nil {
		return m.SpeechRate
	}
	return 0
}

type SpeakResponse struct {
	Result  int64  `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *SpeakResponse) Reset()                    { *m = SpeakResponse{} }
func (m *SpeakResponse) String() string            { return proto.CompactTextString(m) }
func (*SpeakResponse) ProtoMessage()               {}
func (*SpeakResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SpeakResponse) GetResult() int64 {
	if m != nil {
		return m.Result
	}
	return 0
}

func (m *SpeakResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type VoiceResponse struct {
	Voices []string `protobuf:"bytes,1,rep,name=voices" json:"voices,omitempty"`
}

func (m *VoiceResponse) Reset()                    { *m = VoiceResponse{} }
func (m *VoiceResponse) String() string            { return proto.CompactTextString(m) }
func (*VoiceResponse) ProtoMessage()               {}
func (*VoiceResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *VoiceResponse) GetVoices() []string {
	if m != nil {
		return m.Voices
	}
	return nil
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*SpeakEvent)(nil), "speak.SpeakEvent")
	proto.RegisterType((*SpeakResponse)(nil), "speak.SpeakResponse")
	proto.RegisterType((*VoiceResponse)(nil), "speak.VoiceResponse")
	proto.RegisterType((*Empty)(nil), "speak.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SpeakService service

type SpeakServiceClient interface {
	SaySomething(ctx context.Context, in *SpeakEvent, opts ...grpc.CallOption) (*Empty, error)
	GetVoices(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*VoiceResponse, error)
}

type speakServiceClient struct {
	cc *grpc.ClientConn
}

func NewSpeakServiceClient(cc *grpc.ClientConn) SpeakServiceClient {
	return &speakServiceClient{cc}
}

func (c *speakServiceClient) SaySomething(ctx context.Context, in *SpeakEvent, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/speak.SpeakService/SaySomething", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *speakServiceClient) GetVoices(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*VoiceResponse, error) {
	out := new(VoiceResponse)
	err := grpc.Invoke(ctx, "/speak.SpeakService/GetVoices", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SpeakService service

type SpeakServiceServer interface {
	SaySomething(context.Context, *SpeakEvent) (*Empty, error)
	GetVoices(context.Context, *Empty) (*VoiceResponse, error)
}

func RegisterSpeakServiceServer(s *grpc.Server, srv SpeakServiceServer) {
	s.RegisterService(&_SpeakService_serviceDesc, srv)
}

func _SpeakService_SaySomething_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpeakEvent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpeakServiceServer).SaySomething(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/speak.SpeakService/SaySomething",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpeakServiceServer).SaySomething(ctx, req.(*SpeakEvent))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpeakService_GetVoices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpeakServiceServer).GetVoices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/speak.SpeakService/GetVoices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpeakServiceServer).GetVoices(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _SpeakService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "speak.SpeakService",
	HandlerType: (*SpeakServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaySomething",
			Handler:    _SpeakService_SaySomething_Handler,
		},
		{
			MethodName: "GetVoices",
			Handler:    _SpeakService_GetVoices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "speak.proto",
}

func init() { proto.RegisterFile("speak.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xc1, 0x6a, 0xc2, 0x40,
	0x10, 0x86, 0x4d, 0x43, 0x14, 0xa7, 0xf1, 0xd0, 0x41, 0xca, 0xe2, 0xa1, 0x84, 0xbd, 0xd4, 0x93,
	0x60, 0xfb, 0x04, 0x3d, 0x48, 0xef, 0x1b, 0xe8, 0xa1, 0xb7, 0x54, 0x06, 0x95, 0x36, 0xd9, 0x25,
	0x33, 0x0d, 0xf8, 0xf6, 0x65, 0x27, 0x2b, 0xd5, 0xdb, 0x7e, 0x33, 0xf0, 0x7f, 0xff, 0x0e, 0xdc,
	0x73, 0xa0, 0xe6, 0x7b, 0x13, 0x7a, 0x2f, 0x1e, 0x0b, 0x05, 0xfb, 0x09, 0x50, 0xc7, 0xc7, 0x6e,
	0xa0, 0x4e, 0xf0, 0x11, 0xa6, 0x1c, 0x88, 0xf6, 0x47, 0x93, 0x55, 0xd9, 0x7a, 0xee, 0x12, 0xe1,
	0x12, 0x8a, 0xc1, 0x9f, 0xf6, 0x64, 0xee, 0x74, 0x3c, 0x02, 0x3e, 0x01, 0x8c, 0x7b, 0xd7, 0x08,
	0x99, 0xbc, 0xca, 0xd6, 0xb9, 0xbb, 0x9a, 0xd8, 0x37, 0x58, 0x68, 0xb6, 0x23, 0x0e, 0xbe, 0x63,
	0x8a, 0xf1, 0x3d, 0xf1, 0xef, 0x8f, 0x68, 0x7c, 0xee, 0x12, 0xa1, 0x81, 0x59, 0x4b, 0xcc, 0xcd,
	0xe1, 0x22, 0xb8, 0xa0, 0x7d, 0x86, 0xc5, 0x47, 0x74, 0x5d, 0x47, 0xa8, 0x9c, 0x4d, 0x56, 0xe5,
	0xb1, 0xe1, 0x48, 0x76, 0x06, 0xc5, 0xae, 0x0d, 0x72, 0x7e, 0x11, 0x28, 0x55, 0x5a, 0x53, 0x3f,
	0xc4, 0x92, 0x5b, 0x28, 0xeb, 0xe6, 0x5c, 0xfb, 0x96, 0xe4, 0x78, 0xea, 0x0e, 0xf8, 0xb0, 0x19,
	0xaf, 0xf0, 0xff, 0xeb, 0x55, 0x99, 0x46, 0x1a, 0x60, 0x27, 0xb8, 0x85, 0xf9, 0x3b, 0x89, 0x7a,
	0x19, 0x6f, 0x96, 0xab, 0x65, 0xa2, 0x9b, 0x52, 0x76, 0xf2, 0x35, 0xd5, 0xa3, 0xbe, 0xfe, 0x05,
	0x00, 0x00, 0xff, 0xff, 0x86, 0x3a, 0x13, 0x2e, 0x63, 0x01, 0x00, 0x00,
}
