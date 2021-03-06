// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cilium/api/nphds.proto

package cilium

import (
	context "context"
	fmt "fmt"
	_ "github.com/cilium/proxy/go/envoy/annotations"
	v3 "github.com/cilium/proxy/go/envoy/service/discovery/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// The mapping of a network policy identifier to the IP addresses of all the
// hosts on which the network policy is enforced.
// A host may be associated only with one network policy.
type NetworkPolicyHosts struct {
	// The unique identifier of the network policy enforced on the hosts.
	Policy uint64 `protobuf:"varint,1,opt,name=policy,proto3" json:"policy,omitempty"`
	// The set of IP addresses of the hosts on which the network policy is enforced.
	// Optional. May be empty.
	HostAddresses        []string `protobuf:"bytes,2,rep,name=host_addresses,json=hostAddresses,proto3" json:"host_addresses,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkPolicyHosts) Reset()         { *m = NetworkPolicyHosts{} }
func (m *NetworkPolicyHosts) String() string { return proto.CompactTextString(m) }
func (*NetworkPolicyHosts) ProtoMessage()    {}
func (*NetworkPolicyHosts) Descriptor() ([]byte, []int) {
	return fileDescriptor_b09d37d0b8c67603, []int{0}
}

func (m *NetworkPolicyHosts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkPolicyHosts.Unmarshal(m, b)
}
func (m *NetworkPolicyHosts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkPolicyHosts.Marshal(b, m, deterministic)
}
func (m *NetworkPolicyHosts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkPolicyHosts.Merge(m, src)
}
func (m *NetworkPolicyHosts) XXX_Size() int {
	return xxx_messageInfo_NetworkPolicyHosts.Size(m)
}
func (m *NetworkPolicyHosts) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkPolicyHosts.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkPolicyHosts proto.InternalMessageInfo

func (m *NetworkPolicyHosts) GetPolicy() uint64 {
	if m != nil {
		return m.Policy
	}
	return 0
}

func (m *NetworkPolicyHosts) GetHostAddresses() []string {
	if m != nil {
		return m.HostAddresses
	}
	return nil
}

func init() {
	proto.RegisterType((*NetworkPolicyHosts)(nil), "cilium.NetworkPolicyHosts")
}

func init() { proto.RegisterFile("cilium/api/nphds.proto", fileDescriptor_b09d37d0b8c67603) }

var fileDescriptor_b09d37d0b8c67603 = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x52, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0xc5, 0xa1, 0x04, 0xd5, 0x12, 0x0c, 0x11, 0x6a, 0x43, 0x60, 0x08, 0x59, 0xa8, 0x2a, 0x1a,
	0xa3, 0x76, 0x2b, 0x13, 0x15, 0x42, 0x4c, 0x08, 0xb5, 0x1f, 0x50, 0x85, 0xe4, 0xd4, 0x5a, 0xa4,
	0x71, 0xf0, 0xb9, 0x41, 0x1d, 0x58, 0x18, 0x19, 0x8b, 0xc4, 0x17, 0xc0, 0x17, 0xf1, 0x0b, 0xcc,
	0x2c, 0x6c, 0x9d, 0x50, 0xe3, 0x94, 0x22, 0x95, 0x81, 0x89, 0xcd, 0xe7, 0x7b, 0x7e, 0xf7, 0xee,
	0x3d, 0xd3, 0x4a, 0xc8, 0x63, 0x3e, 0x1e, 0xb1, 0x20, 0xe5, 0x2c, 0x49, 0x87, 0x11, 0xfa, 0xa9,
	0x14, 0x4a, 0x58, 0xa6, 0xbe, 0x77, 0xea, 0x90, 0x64, 0x62, 0xc2, 0x10, 0x64, 0xc6, 0x43, 0x60,
	0x11, 0xc7, 0x50, 0x64, 0x20, 0x27, 0x2c, 0x6b, 0x2d, 0x0b, 0xfd, 0xc6, 0xd9, 0x1f, 0x08, 0x31,
	0x88, 0x21, 0xe7, 0x0a, 0x92, 0x44, 0xa8, 0x40, 0x71, 0x91, 0x14, 0x8c, 0x8e, 0xab, 0x99, 0x7e,
	0x34, 0x98, 0x04, 0x14, 0x63, 0x19, 0x42, 0x81, 0xa8, 0x66, 0x41, 0xcc, 0xa3, 0x40, 0x01, 0x5b,
	0x1c, 0x74, 0xc3, 0xe3, 0xd4, 0xba, 0x04, 0x75, 0x27, 0xe4, 0xcd, 0x95, 0x88, 0x79, 0x38, 0xb9,
	0x10, 0xa8, 0xd0, 0xaa, 0x50, 0x33, 0xcd, 0x4b, 0x9b, 0xb8, 0xa4, 0x56, 0xea, 0x16, 0x95, 0x75,
	0x42, 0xb7, 0x87, 0x02, 0x55, 0x3f, 0x88, 0x22, 0x09, 0x88, 0x80, 0xb6, 0xe1, 0xae, 0xd7, 0xca,
	0x9d, 0x9d, 0x59, 0x67, 0x63, 0x4a, 0x0c, 0x9b, 0xcc, 0x3a, 0xe5, 0x29, 0x31, 0xbd, 0x92, 0x34,
	0x5c, 0xd2, 0xdd, 0x9a, 0x63, 0x4f, 0x17, 0xd0, 0xe6, 0x87, 0x41, 0xbd, 0xd5, 0x59, 0x67, 0x8b,
	0x4d, 0x7b, 0xda, 0x07, 0xeb, 0x9e, 0xda, 0x3d, 0x25, 0x21, 0x18, 0xfd, 0xa2, 0xeb, 0xc8, 0xcf,
	0x37, 0xf5, 0x0b, 0xcf, 0xfc, 0xa5, 0x4d, 0x59, 0xcb, 0xff, 0x66, 0xea, 0xc2, 0xed, 0x18, 0x50,
	0x39, 0x8d, 0x3f, 0xa2, 0x31, 0x15, 0x09, 0x82, 0xb7, 0x56, 0x23, 0xc7, 0xc4, 0x7a, 0x25, 0xb4,
	0x7a, 0x0e, 0x2a, 0x1c, 0xfe, 0xf7, 0xf8, 0xc6, 0xc3, 0xdb, 0xfb, 0x93, 0x71, 0xe8, 0x79, 0x2c,
	0x6b, 0x2e, 0xf3, 0x6f, 0x27, 0x5a, 0x46, 0x5f, 0x27, 0xd0, 0x9f, 0x5b, 0x8a, 0x6d, 0x52, 0x77,
	0x0e, 0x1e, 0x5f, 0x9e, 0x3f, 0x37, 0xf7, 0xe8, 0xae, 0xfe, 0x4d, 0xfe, 0xaa, 0xdc, 0x6b, 0x33,
	0x8f, 0xb8, 0xf5, 0x15, 0x00, 0x00, 0xff, 0xff, 0xed, 0x91, 0x9e, 0x9d, 0x89, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NetworkPolicyHostsDiscoveryServiceClient is the client API for NetworkPolicyHostsDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NetworkPolicyHostsDiscoveryServiceClient interface {
	StreamNetworkPolicyHosts(ctx context.Context, opts ...grpc.CallOption) (NetworkPolicyHostsDiscoveryService_StreamNetworkPolicyHostsClient, error)
	FetchNetworkPolicyHosts(ctx context.Context, in *v3.DiscoveryRequest, opts ...grpc.CallOption) (*v3.DiscoveryResponse, error)
}

type networkPolicyHostsDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewNetworkPolicyHostsDiscoveryServiceClient(cc *grpc.ClientConn) NetworkPolicyHostsDiscoveryServiceClient {
	return &networkPolicyHostsDiscoveryServiceClient{cc}
}

func (c *networkPolicyHostsDiscoveryServiceClient) StreamNetworkPolicyHosts(ctx context.Context, opts ...grpc.CallOption) (NetworkPolicyHostsDiscoveryService_StreamNetworkPolicyHostsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NetworkPolicyHostsDiscoveryService_serviceDesc.Streams[0], "/cilium.NetworkPolicyHostsDiscoveryService/StreamNetworkPolicyHosts", opts...)
	if err != nil {
		return nil, err
	}
	x := &networkPolicyHostsDiscoveryServiceStreamNetworkPolicyHostsClient{stream}
	return x, nil
}

type NetworkPolicyHostsDiscoveryService_StreamNetworkPolicyHostsClient interface {
	Send(*v3.DiscoveryRequest) error
	Recv() (*v3.DiscoveryResponse, error)
	grpc.ClientStream
}

type networkPolicyHostsDiscoveryServiceStreamNetworkPolicyHostsClient struct {
	grpc.ClientStream
}

func (x *networkPolicyHostsDiscoveryServiceStreamNetworkPolicyHostsClient) Send(m *v3.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *networkPolicyHostsDiscoveryServiceStreamNetworkPolicyHostsClient) Recv() (*v3.DiscoveryResponse, error) {
	m := new(v3.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *networkPolicyHostsDiscoveryServiceClient) FetchNetworkPolicyHosts(ctx context.Context, in *v3.DiscoveryRequest, opts ...grpc.CallOption) (*v3.DiscoveryResponse, error) {
	out := new(v3.DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/cilium.NetworkPolicyHostsDiscoveryService/FetchNetworkPolicyHosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkPolicyHostsDiscoveryServiceServer is the server API for NetworkPolicyHostsDiscoveryService service.
type NetworkPolicyHostsDiscoveryServiceServer interface {
	StreamNetworkPolicyHosts(NetworkPolicyHostsDiscoveryService_StreamNetworkPolicyHostsServer) error
	FetchNetworkPolicyHosts(context.Context, *v3.DiscoveryRequest) (*v3.DiscoveryResponse, error)
}

// UnimplementedNetworkPolicyHostsDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNetworkPolicyHostsDiscoveryServiceServer struct {
}

func (*UnimplementedNetworkPolicyHostsDiscoveryServiceServer) StreamNetworkPolicyHosts(srv NetworkPolicyHostsDiscoveryService_StreamNetworkPolicyHostsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamNetworkPolicyHosts not implemented")
}
func (*UnimplementedNetworkPolicyHostsDiscoveryServiceServer) FetchNetworkPolicyHosts(ctx context.Context, req *v3.DiscoveryRequest) (*v3.DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchNetworkPolicyHosts not implemented")
}

func RegisterNetworkPolicyHostsDiscoveryServiceServer(s *grpc.Server, srv NetworkPolicyHostsDiscoveryServiceServer) {
	s.RegisterService(&_NetworkPolicyHostsDiscoveryService_serviceDesc, srv)
}

func _NetworkPolicyHostsDiscoveryService_StreamNetworkPolicyHosts_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NetworkPolicyHostsDiscoveryServiceServer).StreamNetworkPolicyHosts(&networkPolicyHostsDiscoveryServiceStreamNetworkPolicyHostsServer{stream})
}

type NetworkPolicyHostsDiscoveryService_StreamNetworkPolicyHostsServer interface {
	Send(*v3.DiscoveryResponse) error
	Recv() (*v3.DiscoveryRequest, error)
	grpc.ServerStream
}

type networkPolicyHostsDiscoveryServiceStreamNetworkPolicyHostsServer struct {
	grpc.ServerStream
}

func (x *networkPolicyHostsDiscoveryServiceStreamNetworkPolicyHostsServer) Send(m *v3.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *networkPolicyHostsDiscoveryServiceStreamNetworkPolicyHostsServer) Recv() (*v3.DiscoveryRequest, error) {
	m := new(v3.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _NetworkPolicyHostsDiscoveryService_FetchNetworkPolicyHosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkPolicyHostsDiscoveryServiceServer).FetchNetworkPolicyHosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cilium.NetworkPolicyHostsDiscoveryService/FetchNetworkPolicyHosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkPolicyHostsDiscoveryServiceServer).FetchNetworkPolicyHosts(ctx, req.(*v3.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NetworkPolicyHostsDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cilium.NetworkPolicyHostsDiscoveryService",
	HandlerType: (*NetworkPolicyHostsDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchNetworkPolicyHosts",
			Handler:    _NetworkPolicyHostsDiscoveryService_FetchNetworkPolicyHosts_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamNetworkPolicyHosts",
			Handler:       _NetworkPolicyHostsDiscoveryService_StreamNetworkPolicyHosts_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "cilium/api/nphds.proto",
}
