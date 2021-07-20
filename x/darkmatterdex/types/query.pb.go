// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: darkmatterdex/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
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

func init() { proto.RegisterFile("darkmatterdex/query.proto", fileDescriptor_e81081fb9321cc8e) }

var fileDescriptor_e81081fb9321cc8e = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0x49, 0x2c, 0xca,
	0xce, 0x4d, 0x2c, 0x29, 0x49, 0x2d, 0x4a, 0x49, 0xad, 0xd0, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x4f, 0x2a, 0xca, 0xcf, 0x48, 0xcc, 0x4d, 0x07, 0x11,
	0x7a, 0x28, 0xca, 0x50, 0x79, 0x52, 0x32, 0xe9, 0xf9, 0xf9, 0xe9, 0x39, 0xa9, 0xfa, 0x89, 0x05,
	0x99, 0xfa, 0x89, 0x79, 0x79, 0xf9, 0x25, 0x89, 0x25, 0x99, 0xf9, 0x79, 0xc5, 0x10, 0x63, 0xa4,
	0xb4, 0x92, 0xf3, 0x8b, 0x73, 0xf3, 0x8b, 0xf5, 0x93, 0x12, 0x8b, 0x53, 0x21, 0xe6, 0xeb, 0x97,
	0x19, 0x26, 0xa5, 0x96, 0x24, 0x1a, 0xea, 0x17, 0x24, 0xa6, 0x67, 0xe6, 0x81, 0x15, 0x43, 0xd4,
	0x1a, 0xb1, 0x73, 0xb1, 0x06, 0x82, 0x54, 0x38, 0x85, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91,
	0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3,
	0xb1, 0x1c, 0x43, 0x94, 0x55, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e,
	0x92, 0x03, 0xf5, 0x51, 0xfd, 0x51, 0x81, 0xc6, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03,
	0xdb, 0x62, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x34, 0x8b, 0x6c, 0xa6, 0xf5, 0x00, 0x00, 0x00,
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
	ServiceName: "brohamgoham.darkmatterdex.darkmatterdex.Query",
	HandlerType: (*QueryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "darkmatterdex/query.proto",
}
