// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messaging/data.proto

package messaging

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

type AppendRequest struct {
	Cid                  int32    `protobuf:"varint,1,opt,name=cid,proto3" json:"cid,omitempty"`
	Csn                  int32    `protobuf:"varint,2,opt,name=csn,proto3" json:"csn,omitempty"`
	Record               string   `protobuf:"bytes,3,opt,name=record,proto3" json:"record,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppendRequest) Reset()         { *m = AppendRequest{} }
func (m *AppendRequest) String() string { return proto.CompactTextString(m) }
func (*AppendRequest) ProtoMessage()    {}
func (*AppendRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_data_b161ddc57c3945d6, []int{0}
}
func (m *AppendRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppendRequest.Unmarshal(m, b)
}
func (m *AppendRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppendRequest.Marshal(b, m, deterministic)
}
func (dst *AppendRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppendRequest.Merge(dst, src)
}
func (m *AppendRequest) XXX_Size() int {
	return xxx_messageInfo_AppendRequest.Size(m)
}
func (m *AppendRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AppendRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AppendRequest proto.InternalMessageInfo

func (m *AppendRequest) GetCid() int32 {
	if m != nil {
		return m.Cid
	}
	return 0
}

func (m *AppendRequest) GetCsn() int32 {
	if m != nil {
		return m.Csn
	}
	return 0
}

func (m *AppendRequest) GetRecord() string {
	if m != nil {
		return m.Record
	}
	return ""
}

type AppendResponse struct {
	Csn                  int32    `protobuf:"varint,1,opt,name=csn,proto3" json:"csn,omitempty"`
	Gsn                  int32    `protobuf:"varint,2,opt,name=gsn,proto3" json:"gsn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppendResponse) Reset()         { *m = AppendResponse{} }
func (m *AppendResponse) String() string { return proto.CompactTextString(m) }
func (*AppendResponse) ProtoMessage()    {}
func (*AppendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_data_b161ddc57c3945d6, []int{1}
}
func (m *AppendResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppendResponse.Unmarshal(m, b)
}
func (m *AppendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppendResponse.Marshal(b, m, deterministic)
}
func (dst *AppendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppendResponse.Merge(dst, src)
}
func (m *AppendResponse) XXX_Size() int {
	return xxx_messageInfo_AppendResponse.Size(m)
}
func (m *AppendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AppendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AppendResponse proto.InternalMessageInfo

func (m *AppendResponse) GetCsn() int32 {
	if m != nil {
		return m.Csn
	}
	return 0
}

func (m *AppendResponse) GetGsn() int32 {
	if m != nil {
		return m.Gsn
	}
	return 0
}

type ReplicateRequest struct {
	ServerID             int32    `protobuf:"varint,1,opt,name=serverID,proto3" json:"serverID,omitempty"`
	Record               string   `protobuf:"bytes,2,opt,name=record,proto3" json:"record,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReplicateRequest) Reset()         { *m = ReplicateRequest{} }
func (m *ReplicateRequest) String() string { return proto.CompactTextString(m) }
func (*ReplicateRequest) ProtoMessage()    {}
func (*ReplicateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_data_b161ddc57c3945d6, []int{2}
}
func (m *ReplicateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplicateRequest.Unmarshal(m, b)
}
func (m *ReplicateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplicateRequest.Marshal(b, m, deterministic)
}
func (dst *ReplicateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplicateRequest.Merge(dst, src)
}
func (m *ReplicateRequest) XXX_Size() int {
	return xxx_messageInfo_ReplicateRequest.Size(m)
}
func (m *ReplicateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplicateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReplicateRequest proto.InternalMessageInfo

func (m *ReplicateRequest) GetServerID() int32 {
	if m != nil {
		return m.ServerID
	}
	return 0
}

func (m *ReplicateRequest) GetRecord() string {
	if m != nil {
		return m.Record
	}
	return ""
}

// No response needed. In event of failure, we should finalize everything
type ReplicateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReplicateResponse) Reset()         { *m = ReplicateResponse{} }
func (m *ReplicateResponse) String() string { return proto.CompactTextString(m) }
func (*ReplicateResponse) ProtoMessage()    {}
func (*ReplicateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_data_b161ddc57c3945d6, []int{3}
}
func (m *ReplicateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplicateResponse.Unmarshal(m, b)
}
func (m *ReplicateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplicateResponse.Marshal(b, m, deterministic)
}
func (dst *ReplicateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplicateResponse.Merge(dst, src)
}
func (m *ReplicateResponse) XXX_Size() int {
	return xxx_messageInfo_ReplicateResponse.Size(m)
}
func (m *ReplicateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplicateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReplicateResponse proto.InternalMessageInfo

type SubscribeRequest struct {
	SubscriptionGsn      int32    `protobuf:"varint,1,opt,name=subscription_gsn,json=subscriptionGsn,proto3" json:"subscription_gsn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeRequest) Reset()         { *m = SubscribeRequest{} }
func (m *SubscribeRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeRequest) ProtoMessage()    {}
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_data_b161ddc57c3945d6, []int{4}
}
func (m *SubscribeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeRequest.Unmarshal(m, b)
}
func (m *SubscribeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeRequest.Marshal(b, m, deterministic)
}
func (dst *SubscribeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeRequest.Merge(dst, src)
}
func (m *SubscribeRequest) XXX_Size() int {
	return xxx_messageInfo_SubscribeRequest.Size(m)
}
func (m *SubscribeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeRequest proto.InternalMessageInfo

func (m *SubscribeRequest) GetSubscriptionGsn() int32 {
	if m != nil {
		return m.SubscriptionGsn
	}
	return 0
}

type SubscribeResponse struct {
	Gsn                  int32    `protobuf:"varint,1,opt,name=gsn,proto3" json:"gsn,omitempty"`
	Record               string   `protobuf:"bytes,2,opt,name=record,proto3" json:"record,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeResponse) Reset()         { *m = SubscribeResponse{} }
func (m *SubscribeResponse) String() string { return proto.CompactTextString(m) }
func (*SubscribeResponse) ProtoMessage()    {}
func (*SubscribeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_data_b161ddc57c3945d6, []int{5}
}
func (m *SubscribeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeResponse.Unmarshal(m, b)
}
func (m *SubscribeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeResponse.Marshal(b, m, deterministic)
}
func (dst *SubscribeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeResponse.Merge(dst, src)
}
func (m *SubscribeResponse) XXX_Size() int {
	return xxx_messageInfo_SubscribeResponse.Size(m)
}
func (m *SubscribeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeResponse proto.InternalMessageInfo

func (m *SubscribeResponse) GetGsn() int32 {
	if m != nil {
		return m.Gsn
	}
	return 0
}

func (m *SubscribeResponse) GetRecord() string {
	if m != nil {
		return m.Record
	}
	return ""
}

type TrimRequest struct {
	Gsn                  int32    `protobuf:"varint,1,opt,name=gsn,proto3" json:"gsn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrimRequest) Reset()         { *m = TrimRequest{} }
func (m *TrimRequest) String() string { return proto.CompactTextString(m) }
func (*TrimRequest) ProtoMessage()    {}
func (*TrimRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_data_b161ddc57c3945d6, []int{6}
}
func (m *TrimRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrimRequest.Unmarshal(m, b)
}
func (m *TrimRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrimRequest.Marshal(b, m, deterministic)
}
func (dst *TrimRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrimRequest.Merge(dst, src)
}
func (m *TrimRequest) XXX_Size() int {
	return xxx_messageInfo_TrimRequest.Size(m)
}
func (m *TrimRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TrimRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TrimRequest proto.InternalMessageInfo

func (m *TrimRequest) GetGsn() int32 {
	if m != nil {
		return m.Gsn
	}
	return 0
}

type TrimResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrimResponse) Reset()         { *m = TrimResponse{} }
func (m *TrimResponse) String() string { return proto.CompactTextString(m) }
func (*TrimResponse) ProtoMessage()    {}
func (*TrimResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_data_b161ddc57c3945d6, []int{7}
}
func (m *TrimResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrimResponse.Unmarshal(m, b)
}
func (m *TrimResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrimResponse.Marshal(b, m, deterministic)
}
func (dst *TrimResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrimResponse.Merge(dst, src)
}
func (m *TrimResponse) XXX_Size() int {
	return xxx_messageInfo_TrimResponse.Size(m)
}
func (m *TrimResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TrimResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TrimResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AppendRequest)(nil), "messaging.AppendRequest")
	proto.RegisterType((*AppendResponse)(nil), "messaging.AppendResponse")
	proto.RegisterType((*ReplicateRequest)(nil), "messaging.ReplicateRequest")
	proto.RegisterType((*ReplicateResponse)(nil), "messaging.ReplicateResponse")
	proto.RegisterType((*SubscribeRequest)(nil), "messaging.SubscribeRequest")
	proto.RegisterType((*SubscribeResponse)(nil), "messaging.SubscribeResponse")
	proto.RegisterType((*TrimRequest)(nil), "messaging.TrimRequest")
	proto.RegisterType((*TrimResponse)(nil), "messaging.TrimResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DataClient is the client API for Data service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataClient interface {
	Append(ctx context.Context, in *AppendRequest, opts ...grpc.CallOption) (*AppendResponse, error)
	Replicate(ctx context.Context, opts ...grpc.CallOption) (Data_ReplicateClient, error)
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (Data_SubscribeClient, error)
	Trim(ctx context.Context, in *TrimRequest, opts ...grpc.CallOption) (*TrimResponse, error)
}

type dataClient struct {
	cc *grpc.ClientConn
}

func NewDataClient(cc *grpc.ClientConn) DataClient {
	return &dataClient{cc}
}

func (c *dataClient) Append(ctx context.Context, in *AppendRequest, opts ...grpc.CallOption) (*AppendResponse, error) {
	out := new(AppendResponse)
	err := c.cc.Invoke(ctx, "/messaging.Data/Append", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) Replicate(ctx context.Context, opts ...grpc.CallOption) (Data_ReplicateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Data_serviceDesc.Streams[0], "/messaging.Data/Replicate", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataReplicateClient{stream}
	return x, nil
}

type Data_ReplicateClient interface {
	Send(*ReplicateRequest) error
	CloseAndRecv() (*ReplicateResponse, error)
	grpc.ClientStream
}

type dataReplicateClient struct {
	grpc.ClientStream
}

func (x *dataReplicateClient) Send(m *ReplicateRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dataReplicateClient) CloseAndRecv() (*ReplicateResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ReplicateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dataClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (Data_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Data_serviceDesc.Streams[1], "/messaging.Data/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Data_SubscribeClient interface {
	Recv() (*SubscribeResponse, error)
	grpc.ClientStream
}

type dataSubscribeClient struct {
	grpc.ClientStream
}

func (x *dataSubscribeClient) Recv() (*SubscribeResponse, error) {
	m := new(SubscribeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dataClient) Trim(ctx context.Context, in *TrimRequest, opts ...grpc.CallOption) (*TrimResponse, error) {
	out := new(TrimResponse)
	err := c.cc.Invoke(ctx, "/messaging.Data/Trim", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataServer is the server API for Data service.
type DataServer interface {
	Append(context.Context, *AppendRequest) (*AppendResponse, error)
	Replicate(Data_ReplicateServer) error
	Subscribe(*SubscribeRequest, Data_SubscribeServer) error
	Trim(context.Context, *TrimRequest) (*TrimResponse, error)
}

func RegisterDataServer(s *grpc.Server, srv DataServer) {
	s.RegisterService(&_Data_serviceDesc, srv)
}

func _Data_Append_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).Append(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messaging.Data/Append",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).Append(ctx, req.(*AppendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Data_Replicate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DataServer).Replicate(&dataReplicateServer{stream})
}

type Data_ReplicateServer interface {
	SendAndClose(*ReplicateResponse) error
	Recv() (*ReplicateRequest, error)
	grpc.ServerStream
}

type dataReplicateServer struct {
	grpc.ServerStream
}

func (x *dataReplicateServer) SendAndClose(m *ReplicateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dataReplicateServer) Recv() (*ReplicateRequest, error) {
	m := new(ReplicateRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Data_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DataServer).Subscribe(m, &dataSubscribeServer{stream})
}

type Data_SubscribeServer interface {
	Send(*SubscribeResponse) error
	grpc.ServerStream
}

type dataSubscribeServer struct {
	grpc.ServerStream
}

func (x *dataSubscribeServer) Send(m *SubscribeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Data_Trim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrimRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).Trim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messaging.Data/Trim",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).Trim(ctx, req.(*TrimRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Data_serviceDesc = grpc.ServiceDesc{
	ServiceName: "messaging.Data",
	HandlerType: (*DataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Append",
			Handler:    _Data_Append_Handler,
		},
		{
			MethodName: "Trim",
			Handler:    _Data_Trim_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Replicate",
			Handler:       _Data_Replicate_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Subscribe",
			Handler:       _Data_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "messaging/data.proto",
}

func init() { proto.RegisterFile("messaging/data.proto", fileDescriptor_data_b161ddc57c3945d6) }

var fileDescriptor_data_b161ddc57c3945d6 = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x5f, 0x4f, 0xc2, 0x30,
	0x14, 0xc5, 0x19, 0x20, 0x91, 0xab, 0xe2, 0xa8, 0x06, 0xe7, 0x34, 0x91, 0xf4, 0x09, 0x5f, 0xd0,
	0xa8, 0x2f, 0x3e, 0x10, 0x63, 0x42, 0x34, 0xea, 0xdb, 0xf4, 0xdd, 0xec, 0x4f, 0xb3, 0x34, 0x91,
	0xae, 0xf6, 0x16, 0x3f, 0x8c, 0x9f, 0xd6, 0x8c, 0x6d, 0x5d, 0x07, 0xf8, 0xd6, 0x9e, 0xde, 0xf3,
	0xeb, 0xbd, 0xa7, 0x85, 0xe3, 0x05, 0x43, 0x0c, 0x53, 0x2e, 0xd2, 0xab, 0x24, 0xd4, 0xe1, 0x54,
	0xaa, 0x4c, 0x67, 0xa4, 0x6f, 0x54, 0xfa, 0x06, 0x07, 0x8f, 0x52, 0x32, 0x91, 0x04, 0xec, 0x7b,
	0xc9, 0x50, 0x13, 0x17, 0x3a, 0x31, 0x4f, 0x3c, 0x67, 0xec, 0x4c, 0x76, 0x82, 0x7c, 0xb9, 0x52,
	0x50, 0x78, 0xed, 0x52, 0x41, 0x41, 0x46, 0xd0, 0x53, 0x2c, 0xce, 0x54, 0xe2, 0x75, 0xc6, 0xce,
	0xa4, 0x1f, 0x94, 0x3b, 0x7a, 0x07, 0x83, 0x0a, 0x86, 0x32, 0x13, 0xc8, 0x2a, 0xaf, 0x53, 0x7b,
	0x5d, 0xe8, 0xa4, 0x35, 0x2d, 0x45, 0x41, 0x9f, 0xc0, 0x0d, 0x98, 0xfc, 0xe2, 0x71, 0xa8, 0x59,
	0xd5, 0x85, 0x0f, 0xbb, 0xc8, 0xd4, 0x0f, 0x53, 0x2f, 0xf3, 0xd2, 0x6c, 0xf6, 0xd6, 0xed, 0xed,
	0xc6, 0xed, 0x47, 0x30, 0xb4, 0x38, 0x45, 0x03, 0x74, 0x06, 0xee, 0xfb, 0x32, 0xc2, 0x58, 0xf1,
	0xc8, 0xc0, 0x2f, 0xc1, 0xc5, 0x42, 0x93, 0x9a, 0x67, 0xe2, 0x33, 0x35, 0x1d, 0x1e, 0xda, 0xfa,
	0x33, 0x0a, 0x3a, 0x83, 0xa1, 0x65, 0xaf, 0x87, 0xaa, 0x2d, 0xf9, 0xf2, 0xdf, 0x96, 0x2e, 0x60,
	0xef, 0x43, 0xf1, 0x85, 0x95, 0x6d, 0xd3, 0x48, 0x07, 0xb0, 0x5f, 0x14, 0x14, 0xe8, 0x9b, 0xdf,
	0x36, 0x74, 0xe7, 0xa1, 0x0e, 0xc9, 0x03, 0xf4, 0x8a, 0x28, 0x89, 0x37, 0x35, 0xaf, 0x35, 0x6d,
	0x3c, 0x95, 0x7f, 0xba, 0xe5, 0xa4, 0x1c, 0xbb, 0x45, 0x5e, 0xa1, 0x6f, 0xd2, 0x20, 0x67, 0x56,
	0xe5, 0x7a, 0xd6, 0xfe, 0xf9, 0xf6, 0xc3, 0x8a, 0x34, 0x71, 0x72, 0x96, 0x49, 0xa1, 0xc1, 0x5a,
	0x8f, 0xb6, 0xc1, 0xda, 0x08, 0x8e, 0xb6, 0xae, 0x1d, 0x72, 0x0f, 0xdd, 0x7c, 0x62, 0x32, 0xb2,
	0x2a, 0xad, 0x8c, 0xfc, 0x93, 0x0d, 0xbd, 0x32, 0x47, 0xbd, 0xd5, 0xef, 0xbd, 0xfd, 0x0b, 0x00,
	0x00, 0xff, 0xff, 0x2d, 0x54, 0x90, 0xb6, 0xd5, 0x02, 0x00, 0x00,
}
