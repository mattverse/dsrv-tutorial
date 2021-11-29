// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: arbiter/stake/v1beta/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("arbiter/stake/v1beta/query.proto", fileDescriptor_aa5558554306a625) }

var fileDescriptor_aa5558554306a625 = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0xcf, 0x31, 0x4a, 0x05, 0x31,
	0x10, 0xc6, 0xf1, 0xdd, 0x42, 0x85, 0x57, 0x8a, 0x36, 0x0f, 0x09, 0xd6, 0x82, 0x19, 0x56, 0x6f,
	0x60, 0x67, 0x69, 0x6b, 0x37, 0x79, 0x0c, 0x31, 0xe8, 0xcb, 0xc4, 0xcc, 0xac, 0xb8, 0xb7, 0xf0,
	0x58, 0x96, 0x5b, 0x5a, 0xca, 0xee, 0x45, 0x64, 0x93, 0xd8, 0x05, 0xf2, 0x63, 0xfe, 0x7c, 0xbb,
	0x6b, 0xcc, 0x2e, 0x28, 0x65, 0x10, 0xc5, 0x57, 0x82, 0x8f, 0xc1, 0x91, 0x22, 0xbc, 0x8f, 0x94,
	0x27, 0x9b, 0x32, 0x2b, 0x9f, 0x5f, 0x36, 0x61, 0x8b, 0xb0, 0x55, 0x0c, 0xfb, 0x0b, 0xcf, 0x9e,
	0x8b, 0x80, 0xed, 0x55, 0xf1, 0xfe, 0xca, 0x33, 0xfb, 0x37, 0x02, 0x4c, 0x01, 0x30, 0x46, 0x56,
	0xd4, 0xc0, 0x51, 0xda, 0xef, 0xcd, 0x81, 0xe5, 0xc8, 0x02, 0x0e, 0x85, 0x6a, 0xa3, 0x05, 0x07,
	0x48, 0xe8, 0x43, 0x2c, 0xb8, 0xda, 0xbb, 0xb3, 0xdd, 0xc9, 0xd3, 0x26, 0x1e, 0x1e, 0xbf, 0x17,
	0xd3, 0xcf, 0x8b, 0xe9, 0x7f, 0x17, 0xd3, 0x7f, 0xad, 0xa6, 0x9b, 0x57, 0xd3, 0xfd, 0xac, 0xa6,
	0x7b, 0x06, 0x1f, 0xf4, 0x65, 0x74, 0xf6, 0xc0, 0x47, 0x10, 0x4c, 0x81, 0xa2, 0xdc, 0xb6, 0xc2,
	0xff, 0xaa, 0xcf, 0xb6, 0x4b, 0xa7, 0x44, 0xe2, 0x4e, 0xcb, 0xe9, 0xfb, 0xbf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x08, 0x52, 0xd8, 0x7e, 0xf5, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

// QueryServer is the server API for Query service.
type QueryServer interface {
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "arbiter.stake.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "arbiter/stake/v1beta/query.proto",
}
