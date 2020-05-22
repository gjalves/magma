// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orc8r/protos/streamer.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// --------------------------------------------------------------------------
// Streamer provides a pipeline for the cloud to push the updates to the
// gateway as and when the update happens.
//
// The Streamer interface defines the semantics and consistency guarantees
// between the cloud and the gateway while abstracting the details of how
// its implemented in the cloud and what the gateway does with the updates.
//
// - The gateways call the GetUpdates() streaming API with a StreamRequest
//   indicating the stream name and the offset to continue streaming from.
// - The cloud sends a stream of DataUpdateBatch containing a batch of updates.
// - If resync is true, then the gateway can cleanup all its data and add
//   all the keys (the batch is guaranteed to contain only unique keys).
// - If resync is false, then the gateway can update the keys, or add new
//   ones if the key is not already present.
// - Key deletions are not yet supported (#15109350)
// --------------------------------------------------------------------------
type StreamRequest struct {
	GatewayId string `protobuf:"bytes,1,opt,name=gatewayId,proto3" json:"gatewayId,omitempty"`
	// Stream name to attach to. (Eg:) subscriberdb, config, etc.
	StreamName string `protobuf:"bytes,2,opt,name=stream_name,json=streamName,proto3" json:"stream_name,omitempty"`
	// Any extra data to send up with the stream request. This value will be
	// different per stream provider.
	ExtraArgs            *any.Any `protobuf:"bytes,3,opt,name=extra_args,json=extraArgs,proto3" json:"extra_args,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamRequest) Reset()         { *m = StreamRequest{} }
func (m *StreamRequest) String() string { return proto.CompactTextString(m) }
func (*StreamRequest) ProtoMessage()    {}
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_acdce76608ae0d01, []int{0}
}

func (m *StreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamRequest.Unmarshal(m, b)
}
func (m *StreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamRequest.Marshal(b, m, deterministic)
}
func (m *StreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamRequest.Merge(m, src)
}
func (m *StreamRequest) XXX_Size() int {
	return xxx_messageInfo_StreamRequest.Size(m)
}
func (m *StreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamRequest proto.InternalMessageInfo

func (m *StreamRequest) GetGatewayId() string {
	if m != nil {
		return m.GatewayId
	}
	return ""
}

func (m *StreamRequest) GetStreamName() string {
	if m != nil {
		return m.StreamName
	}
	return ""
}

func (m *StreamRequest) GetExtraArgs() *any.Any {
	if m != nil {
		return m.ExtraArgs
	}
	return nil
}

type DataUpdate struct {
	// Unique key for each item
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// value can be file contents, protobuf serialized message, etc.
	// For key deletions, the value field would be absent.
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataUpdate) Reset()         { *m = DataUpdate{} }
func (m *DataUpdate) String() string { return proto.CompactTextString(m) }
func (*DataUpdate) ProtoMessage()    {}
func (*DataUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_acdce76608ae0d01, []int{1}
}

func (m *DataUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataUpdate.Unmarshal(m, b)
}
func (m *DataUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataUpdate.Marshal(b, m, deterministic)
}
func (m *DataUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataUpdate.Merge(m, src)
}
func (m *DataUpdate) XXX_Size() int {
	return xxx_messageInfo_DataUpdate.Size(m)
}
func (m *DataUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_DataUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_DataUpdate proto.InternalMessageInfo

func (m *DataUpdate) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *DataUpdate) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type DataUpdateBatch struct {
	Updates []*DataUpdate `protobuf:"bytes,1,rep,name=updates,proto3" json:"updates,omitempty"`
	// If resync is true, the updates would be a snapshot of all the
	// contents in the cloud.
	Resync               bool     `protobuf:"varint,2,opt,name=resync,proto3" json:"resync,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataUpdateBatch) Reset()         { *m = DataUpdateBatch{} }
func (m *DataUpdateBatch) String() string { return proto.CompactTextString(m) }
func (*DataUpdateBatch) ProtoMessage()    {}
func (*DataUpdateBatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_acdce76608ae0d01, []int{2}
}

func (m *DataUpdateBatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataUpdateBatch.Unmarshal(m, b)
}
func (m *DataUpdateBatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataUpdateBatch.Marshal(b, m, deterministic)
}
func (m *DataUpdateBatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataUpdateBatch.Merge(m, src)
}
func (m *DataUpdateBatch) XXX_Size() int {
	return xxx_messageInfo_DataUpdateBatch.Size(m)
}
func (m *DataUpdateBatch) XXX_DiscardUnknown() {
	xxx_messageInfo_DataUpdateBatch.DiscardUnknown(m)
}

var xxx_messageInfo_DataUpdateBatch proto.InternalMessageInfo

func (m *DataUpdateBatch) GetUpdates() []*DataUpdate {
	if m != nil {
		return m.Updates
	}
	return nil
}

func (m *DataUpdateBatch) GetResync() bool {
	if m != nil {
		return m.Resync
	}
	return false
}

func init() {
	proto.RegisterType((*StreamRequest)(nil), "magma.orc8r.StreamRequest")
	proto.RegisterType((*DataUpdate)(nil), "magma.orc8r.DataUpdate")
	proto.RegisterType((*DataUpdateBatch)(nil), "magma.orc8r.DataUpdateBatch")
}

func init() { proto.RegisterFile("orc8r/protos/streamer.proto", fileDescriptor_acdce76608ae0d01) }

var fileDescriptor_acdce76608ae0d01 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xcd, 0x4f, 0xf2, 0x40,
	0x10, 0xc6, 0xdf, 0xbe, 0x44, 0x84, 0xa9, 0x5f, 0xd9, 0x10, 0x2d, 0x1f, 0x89, 0xa4, 0x27, 0x4e,
	0xad, 0x16, 0x0f, 0x5e, 0x21, 0x26, 0x7e, 0x1c, 0x8c, 0x29, 0xc1, 0x03, 0x31, 0x21, 0x03, 0x8c,
	0x2b, 0x91, 0x76, 0x71, 0x77, 0x8b, 0xf6, 0xec, 0x3f, 0x6e, 0xdc, 0x2d, 0x41, 0x0e, 0xde, 0x3c,
	0xb5, 0xcf, 0xec, 0xef, 0x99, 0x67, 0x32, 0x19, 0x68, 0x0a, 0x39, 0xbd, 0x94, 0xe1, 0x52, 0x0a,
	0x2d, 0x54, 0xa8, 0xb4, 0x24, 0x4c, 0x48, 0x06, 0x46, 0x33, 0x37, 0x41, 0x9e, 0x60, 0x60, 0x90,
	0x46, 0x9d, 0x0b, 0xc1, 0x17, 0x64, 0xd1, 0x49, 0xf6, 0x1c, 0x62, 0x9a, 0x5b, 0xce, 0xff, 0x74,
	0x60, 0x7f, 0x60, 0xac, 0x31, 0xbd, 0x65, 0xa4, 0x34, 0x6b, 0x41, 0x95, 0xa3, 0xa6, 0x77, 0xcc,
	0x6f, 0x67, 0x9e, 0xd3, 0x76, 0x3a, 0xd5, 0x78, 0x53, 0x60, 0xa7, 0xe0, 0xda, 0xa4, 0x71, 0x8a,
	0x09, 0x79, 0xff, 0xcd, 0x3b, 0xd8, 0xd2, 0x3d, 0x26, 0xc4, 0xba, 0x00, 0xf4, 0xa1, 0x25, 0x8e,
	0x51, 0x72, 0xe5, 0x95, 0xda, 0x4e, 0xc7, 0x8d, 0x6a, 0x81, 0x1d, 0x20, 0x58, 0x0f, 0x10, 0xf4,
	0xd2, 0x3c, 0xae, 0x1a, 0xae, 0x27, 0xb9, 0xf2, 0x2f, 0x00, 0xae, 0x50, 0xe3, 0x70, 0x39, 0x43,
	0x4d, 0xec, 0x08, 0x4a, 0xaf, 0x94, 0x17, 0xd9, 0xdf, 0xbf, 0xac, 0x06, 0x3b, 0x2b, 0x5c, 0x64,
	0x36, 0x6f, 0x2f, 0xb6, 0xc2, 0x7f, 0x82, 0xc3, 0x8d, 0xab, 0x8f, 0x7a, 0xfa, 0xc2, 0xce, 0x61,
	0x37, 0x33, 0x52, 0x79, 0x4e, 0xbb, 0xd4, 0x71, 0xa3, 0x93, 0xe0, 0xc7, 0x22, 0x82, 0x0d, 0x1e,
	0xaf, 0x39, 0x76, 0x0c, 0x65, 0x49, 0x2a, 0x4f, 0xa7, 0xa6, 0x79, 0x25, 0x2e, 0x54, 0xf4, 0x08,
	0x95, 0x41, 0xb1, 0x53, 0x76, 0x07, 0x70, 0x4d, 0x7a, 0x58, 0x38, 0x1a, 0x5b, 0x3d, 0xb7, 0xb6,
	0xd7, 0x68, 0xfd, 0x92, 0x67, 0xc6, 0xf3, 0xff, 0x9d, 0x39, 0xd1, 0x08, 0x0e, 0xac, 0xe5, 0x41,
	0x8a, 0xd5, 0x7c, 0x46, 0x92, 0xdd, 0xfc, 0x55, 0xf7, 0x7e, 0x73, 0x54, 0x37, 0x40, 0x68, 0x4f,
	0x63, 0x31, 0x9f, 0x84, 0x5c, 0x14, 0x17, 0x32, 0x29, 0x9b, 0x6f, 0xf7, 0x2b, 0x00, 0x00, 0xff,
	0xff, 0x5f, 0x6a, 0x9f, 0xf7, 0x38, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StreamerClient is the client API for Streamer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StreamerClient interface {
	// Get the stream of updates from the cloud.
	// The RPC call would be kept open to push new updates as they happen.
	GetUpdates(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (Streamer_GetUpdatesClient, error)
}

type streamerClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamerClient(cc grpc.ClientConnInterface) StreamerClient {
	return &streamerClient{cc}
}

func (c *streamerClient) GetUpdates(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (Streamer_GetUpdatesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Streamer_serviceDesc.Streams[0], "/magma.orc8r.Streamer/GetUpdates", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamerGetUpdatesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Streamer_GetUpdatesClient interface {
	Recv() (*DataUpdateBatch, error)
	grpc.ClientStream
}

type streamerGetUpdatesClient struct {
	grpc.ClientStream
}

func (x *streamerGetUpdatesClient) Recv() (*DataUpdateBatch, error) {
	m := new(DataUpdateBatch)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamerServer is the server API for Streamer service.
type StreamerServer interface {
	// Get the stream of updates from the cloud.
	// The RPC call would be kept open to push new updates as they happen.
	GetUpdates(*StreamRequest, Streamer_GetUpdatesServer) error
}

// UnimplementedStreamerServer can be embedded to have forward compatible implementations.
type UnimplementedStreamerServer struct {
}

func (*UnimplementedStreamerServer) GetUpdates(req *StreamRequest, srv Streamer_GetUpdatesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetUpdates not implemented")
}

func RegisterStreamerServer(s *grpc.Server, srv StreamerServer) {
	s.RegisterService(&_Streamer_serviceDesc, srv)
}

func _Streamer_GetUpdates_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamerServer).GetUpdates(m, &streamerGetUpdatesServer{stream})
}

type Streamer_GetUpdatesServer interface {
	Send(*DataUpdateBatch) error
	grpc.ServerStream
}

type streamerGetUpdatesServer struct {
	grpc.ServerStream
}

func (x *streamerGetUpdatesServer) Send(m *DataUpdateBatch) error {
	return x.ServerStream.SendMsg(m)
}

var _Streamer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.orc8r.Streamer",
	HandlerType: (*StreamerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetUpdates",
			Handler:       _Streamer_GetUpdates_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "orc8r/protos/streamer.proto",
}

// StreamProviderClient is the client API for StreamProvider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StreamProviderClient interface {
	GetUpdates(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (*DataUpdateBatch, error)
}

type streamProviderClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamProviderClient(cc grpc.ClientConnInterface) StreamProviderClient {
	return &streamProviderClient{cc}
}

func (c *streamProviderClient) GetUpdates(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (*DataUpdateBatch, error) {
	out := new(DataUpdateBatch)
	err := c.cc.Invoke(ctx, "/magma.orc8r.StreamProvider/GetUpdates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StreamProviderServer is the server API for StreamProvider service.
type StreamProviderServer interface {
	GetUpdates(context.Context, *StreamRequest) (*DataUpdateBatch, error)
}

// UnimplementedStreamProviderServer can be embedded to have forward compatible implementations.
type UnimplementedStreamProviderServer struct {
}

func (*UnimplementedStreamProviderServer) GetUpdates(ctx context.Context, req *StreamRequest) (*DataUpdateBatch, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUpdates not implemented")
}

func RegisterStreamProviderServer(s *grpc.Server, srv StreamProviderServer) {
	s.RegisterService(&_StreamProvider_serviceDesc, srv)
}

func _StreamProvider_GetUpdates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamProviderServer).GetUpdates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.StreamProvider/GetUpdates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamProviderServer).GetUpdates(ctx, req.(*StreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StreamProvider_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.orc8r.StreamProvider",
	HandlerType: (*StreamProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUpdates",
			Handler:    _StreamProvider_GetUpdates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orc8r/protos/streamer.proto",
}
