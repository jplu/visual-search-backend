// Code generated by protoc-gen-go. DO NOT EDIT.
// source: faiss.proto

package faiss

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type HeartbeatResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeartbeatResponse) Reset()         { *m = HeartbeatResponse{} }
func (m *HeartbeatResponse) String() string { return proto.CompactTextString(m) }
func (*HeartbeatResponse) ProtoMessage()    {}
func (*HeartbeatResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_380cd395b6b16ebb, []int{0}
}

func (m *HeartbeatResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartbeatResponse.Unmarshal(m, b)
}
func (m *HeartbeatResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartbeatResponse.Marshal(b, m, deterministic)
}
func (m *HeartbeatResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartbeatResponse.Merge(m, src)
}
func (m *HeartbeatResponse) XXX_Size() int {
	return xxx_messageInfo_HeartbeatResponse.Size(m)
}
func (m *HeartbeatResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartbeatResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HeartbeatResponse proto.InternalMessageInfo

func (m *HeartbeatResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Vector struct {
	FloatVal             []float32 `protobuf:"fixed32,5,rep,packed,name=float_val,json=floatVal,proto3" json:"float_val,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Vector) Reset()         { *m = Vector{} }
func (m *Vector) String() string { return proto.CompactTextString(m) }
func (*Vector) ProtoMessage()    {}
func (*Vector) Descriptor() ([]byte, []int) {
	return fileDescriptor_380cd395b6b16ebb, []int{1}
}

func (m *Vector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vector.Unmarshal(m, b)
}
func (m *Vector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vector.Marshal(b, m, deterministic)
}
func (m *Vector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vector.Merge(m, src)
}
func (m *Vector) XXX_Size() int {
	return xxx_messageInfo_Vector.Size(m)
}
func (m *Vector) XXX_DiscardUnknown() {
	xxx_messageInfo_Vector.DiscardUnknown(m)
}

var xxx_messageInfo_Vector proto.InternalMessageInfo

func (m *Vector) GetFloatVal() []float32 {
	if m != nil {
		return m.FloatVal
	}
	return nil
}

type SearchRequest struct {
	Vector               *Vector  `protobuf:"bytes,1,opt,name=vector,proto3" json:"vector,omitempty"`
	TopK                 uint64   `protobuf:"varint,2,opt,name=top_k,json=topK,proto3" json:"top_k,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_380cd395b6b16ebb, []int{2}
}

func (m *SearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest.Unmarshal(m, b)
}
func (m *SearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest.Marshal(b, m, deterministic)
}
func (m *SearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest.Merge(m, src)
}
func (m *SearchRequest) XXX_Size() int {
	return xxx_messageInfo_SearchRequest.Size(m)
}
func (m *SearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest proto.InternalMessageInfo

func (m *SearchRequest) GetVector() *Vector {
	if m != nil {
		return m.Vector
	}
	return nil
}

func (m *SearchRequest) GetTopK() uint64 {
	if m != nil {
		return m.TopK
	}
	return 0
}

type Neighbor struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Score                float32  `protobuf:"fixed32,2,opt,name=score,proto3" json:"score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Neighbor) Reset()         { *m = Neighbor{} }
func (m *Neighbor) String() string { return proto.CompactTextString(m) }
func (*Neighbor) ProtoMessage()    {}
func (*Neighbor) Descriptor() ([]byte, []int) {
	return fileDescriptor_380cd395b6b16ebb, []int{3}
}

func (m *Neighbor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Neighbor.Unmarshal(m, b)
}
func (m *Neighbor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Neighbor.Marshal(b, m, deterministic)
}
func (m *Neighbor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Neighbor.Merge(m, src)
}
func (m *Neighbor) XXX_Size() int {
	return xxx_messageInfo_Neighbor.Size(m)
}
func (m *Neighbor) XXX_DiscardUnknown() {
	xxx_messageInfo_Neighbor.DiscardUnknown(m)
}

var xxx_messageInfo_Neighbor proto.InternalMessageInfo

func (m *Neighbor) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Neighbor) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

type SearchResponse struct {
	Neighbors            []*Neighbor `protobuf:"bytes,2,rep,name=neighbors,proto3" json:"neighbors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SearchResponse) Reset()         { *m = SearchResponse{} }
func (m *SearchResponse) String() string { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()    {}
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_380cd395b6b16ebb, []int{4}
}

func (m *SearchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchResponse.Unmarshal(m, b)
}
func (m *SearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchResponse.Marshal(b, m, deterministic)
}
func (m *SearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchResponse.Merge(m, src)
}
func (m *SearchResponse) XXX_Size() int {
	return xxx_messageInfo_SearchResponse.Size(m)
}
func (m *SearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchResponse proto.InternalMessageInfo

func (m *SearchResponse) GetNeighbors() []*Neighbor {
	if m != nil {
		return m.Neighbors
	}
	return nil
}

type SearchByIdRequest struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TopK                 uint64   `protobuf:"varint,2,opt,name=top_k,json=topK,proto3" json:"top_k,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchByIdRequest) Reset()         { *m = SearchByIdRequest{} }
func (m *SearchByIdRequest) String() string { return proto.CompactTextString(m) }
func (*SearchByIdRequest) ProtoMessage()    {}
func (*SearchByIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_380cd395b6b16ebb, []int{5}
}

func (m *SearchByIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchByIdRequest.Unmarshal(m, b)
}
func (m *SearchByIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchByIdRequest.Marshal(b, m, deterministic)
}
func (m *SearchByIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchByIdRequest.Merge(m, src)
}
func (m *SearchByIdRequest) XXX_Size() int {
	return xxx_messageInfo_SearchByIdRequest.Size(m)
}
func (m *SearchByIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchByIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchByIdRequest proto.InternalMessageInfo

func (m *SearchByIdRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SearchByIdRequest) GetTopK() uint64 {
	if m != nil {
		return m.TopK
	}
	return 0
}

type SearchByIdResponse struct {
	RequestId            uint64      `protobuf:"varint,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Neighbors            []*Neighbor `protobuf:"bytes,2,rep,name=neighbors,proto3" json:"neighbors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SearchByIdResponse) Reset()         { *m = SearchByIdResponse{} }
func (m *SearchByIdResponse) String() string { return proto.CompactTextString(m) }
func (*SearchByIdResponse) ProtoMessage()    {}
func (*SearchByIdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_380cd395b6b16ebb, []int{6}
}

func (m *SearchByIdResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchByIdResponse.Unmarshal(m, b)
}
func (m *SearchByIdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchByIdResponse.Marshal(b, m, deterministic)
}
func (m *SearchByIdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchByIdResponse.Merge(m, src)
}
func (m *SearchByIdResponse) XXX_Size() int {
	return xxx_messageInfo_SearchByIdResponse.Size(m)
}
func (m *SearchByIdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchByIdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchByIdResponse proto.InternalMessageInfo

func (m *SearchByIdResponse) GetRequestId() uint64 {
	if m != nil {
		return m.RequestId
	}
	return 0
}

func (m *SearchByIdResponse) GetNeighbors() []*Neighbor {
	if m != nil {
		return m.Neighbors
	}
	return nil
}

func init() {
	proto.RegisterType((*HeartbeatResponse)(nil), "faiss.HeartbeatResponse")
	proto.RegisterType((*Vector)(nil), "faiss.Vector")
	proto.RegisterType((*SearchRequest)(nil), "faiss.SearchRequest")
	proto.RegisterType((*Neighbor)(nil), "faiss.Neighbor")
	proto.RegisterType((*SearchResponse)(nil), "faiss.SearchResponse")
	proto.RegisterType((*SearchByIdRequest)(nil), "faiss.SearchByIdRequest")
	proto.RegisterType((*SearchByIdResponse)(nil), "faiss.SearchByIdResponse")
}

func init() { proto.RegisterFile("faiss.proto", fileDescriptor_380cd395b6b16ebb) }

var fileDescriptor_380cd395b6b16ebb = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x5f, 0x8b, 0xd3, 0x40,
	0x14, 0xc5, 0x49, 0xb6, 0x89, 0x9b, 0x5b, 0x77, 0x65, 0xc7, 0x55, 0x62, 0x16, 0x21, 0x04, 0x16,
	0xf2, 0xb2, 0xa9, 0x54, 0x04, 0x5f, 0x44, 0x14, 0x14, 0x4b, 0xc1, 0x87, 0x29, 0xf4, 0xb5, 0x4c,
	0x92, 0xdb, 0x34, 0x98, 0x76, 0xe2, 0xcc, 0xb4, 0xd0, 0x4f, 0xe8, 0xd7, 0x92, 0xce, 0x4c, 0xfa,
	0xcf, 0xbe, 0xf8, 0x38, 0xf7, 0xde, 0xf3, 0xbb, 0xe7, 0xce, 0x81, 0xfe, 0x9c, 0xd5, 0x52, 0x66,
	0xad, 0xe0, 0x8a, 0x13, 0x4f, 0x3f, 0xa2, 0x87, 0x8a, 0xf3, 0xaa, 0xc1, 0x81, 0x2e, 0xe6, 0xeb,
	0xf9, 0x00, 0x97, 0xad, 0xda, 0x9a, 0x99, 0xe4, 0x09, 0xee, 0x7e, 0x20, 0x13, 0x2a, 0x47, 0xa6,
	0x28, 0xca, 0x96, 0xaf, 0x24, 0x92, 0x10, 0x9e, 0x2d, 0x51, 0x4a, 0x56, 0x61, 0xe8, 0xc4, 0x4e,
	0x1a, 0xd0, 0xee, 0x99, 0x3c, 0x82, 0x3f, 0xc5, 0x42, 0x71, 0x41, 0x1e, 0x20, 0x98, 0x37, 0x9c,
	0xa9, 0xd9, 0x86, 0x35, 0xa1, 0x17, 0x5f, 0xa5, 0x2e, 0xbd, 0xd6, 0x85, 0x29, 0x6b, 0x92, 0x31,
	0xdc, 0x4c, 0x90, 0x89, 0x62, 0x41, 0xf1, 0xf7, 0x1a, 0xa5, 0x22, 0x8f, 0xe0, 0x6f, 0xb4, 0x4e,
	0x03, 0xfb, 0xc3, 0x9b, 0xcc, 0x18, 0x35, 0x30, 0x6a, 0x9b, 0xe4, 0x25, 0x78, 0x8a, 0xb7, 0xb3,
	0x5f, 0xa1, 0x1b, 0x3b, 0x69, 0x8f, 0xf6, 0x14, 0x6f, 0xc7, 0xc9, 0x3b, 0xb8, 0xfe, 0x89, 0x75,
	0xb5, 0xc8, 0xb9, 0x20, 0xb7, 0xe0, 0xd6, 0xa5, 0x66, 0xf4, 0xa8, 0x5b, 0x97, 0xe4, 0x1e, 0x3c,
	0x59, 0x70, 0x81, 0x5a, 0xe0, 0x52, 0xf3, 0x48, 0x3e, 0xc3, 0x6d, 0xb7, 0xde, 0x5e, 0xf4, 0x04,
	0xc1, 0xca, 0x32, 0x64, 0xe8, 0xc6, 0x57, 0x69, 0x7f, 0xf8, 0xc2, 0x5a, 0xe8, 0xd8, 0xf4, 0x30,
	0x91, 0x7c, 0x84, 0x3b, 0x03, 0xf8, 0xba, 0x1d, 0x95, 0xdd, 0x0d, 0xe7, 0xbb, 0x2f, 0x9a, 0xcd,
	0x81, 0x1c, 0x2b, 0xed, 0xfa, 0xb7, 0x00, 0xc2, 0x50, 0x66, 0x7b, 0x44, 0x60, 0x2b, 0xa3, 0xf2,
	0x3f, 0xdd, 0x0d, 0xff, 0x38, 0xf0, 0xfc, 0xfb, 0xae, 0x3b, 0x41, 0xb1, 0xa9, 0x0b, 0x24, 0x9f,
	0x20, 0xd8, 0x87, 0x48, 0x5e, 0x67, 0x26, 0xef, 0xac, 0xcb, 0x3b, 0xfb, 0xb6, 0xcb, 0x3b, 0x0a,
	0x2d, 0xf1, 0xdf, 0xb8, 0x3f, 0x80, 0x6f, 0x3c, 0x93, 0x7b, 0x3b, 0x73, 0x12, 0x5e, 0xf4, 0xea,
	0xac, 0x6a, 0x65, 0x5f, 0x00, 0x0e, 0xa7, 0x92, 0xf0, 0x64, 0xe8, 0xe8, 0xdf, 0xa2, 0x37, 0x17,
	0x3a, 0x06, 0x91, 0xfb, 0xda, 0xe3, 0xfb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe1, 0xde, 0x5c,
	0xed, 0xb7, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FaissServiceClient is the client API for FaissService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FaissServiceClient interface {
	Heartbeat(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*HeartbeatResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	SearchById(ctx context.Context, in *SearchByIdRequest, opts ...grpc.CallOption) (*SearchByIdResponse, error)
}

type faissServiceClient struct {
	cc *grpc.ClientConn
}

func NewFaissServiceClient(cc *grpc.ClientConn) FaissServiceClient {
	return &faissServiceClient{cc}
}

func (c *faissServiceClient) Heartbeat(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*HeartbeatResponse, error) {
	out := new(HeartbeatResponse)
	err := c.cc.Invoke(ctx, "/faiss.FaissService/Heartbeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *faissServiceClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/faiss.FaissService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *faissServiceClient) SearchById(ctx context.Context, in *SearchByIdRequest, opts ...grpc.CallOption) (*SearchByIdResponse, error) {
	out := new(SearchByIdResponse)
	err := c.cc.Invoke(ctx, "/faiss.FaissService/SearchById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FaissServiceServer is the server API for FaissService service.
type FaissServiceServer interface {
	Heartbeat(context.Context, *empty.Empty) (*HeartbeatResponse, error)
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
	SearchById(context.Context, *SearchByIdRequest) (*SearchByIdResponse, error)
}

// UnimplementedFaissServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFaissServiceServer struct {
}

func (*UnimplementedFaissServiceServer) Heartbeat(ctx context.Context, req *empty.Empty) (*HeartbeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Heartbeat not implemented")
}
func (*UnimplementedFaissServiceServer) Search(ctx context.Context, req *SearchRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (*UnimplementedFaissServiceServer) SearchById(ctx context.Context, req *SearchByIdRequest) (*SearchByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchById not implemented")
}

func RegisterFaissServiceServer(s *grpc.Server, srv FaissServiceServer) {
	s.RegisterService(&_FaissService_serviceDesc, srv)
}

func _FaissService_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FaissServiceServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/faiss.FaissService/Heartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FaissServiceServer).Heartbeat(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _FaissService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FaissServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/faiss.FaissService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FaissServiceServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FaissService_SearchById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FaissServiceServer).SearchById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/faiss.FaissService/SearchById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FaissServiceServer).SearchById(ctx, req.(*SearchByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FaissService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "faiss.FaissService",
	HandlerType: (*FaissServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Heartbeat",
			Handler:    _FaissService_Heartbeat_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _FaissService_Search_Handler,
		},
		{
			MethodName: "SearchById",
			Handler:    _FaissService_SearchById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "faiss.proto",
}