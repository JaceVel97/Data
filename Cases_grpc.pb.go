// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package Cases

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// StudyCaseClient is the client API for StudyCase service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StudyCaseClient interface {
	ListCases(ctx context.Context, opts ...grpc.CallOption) (StudyCase_ListCasesClient, error)
}

type studyCaseClient struct {
	cc grpc.ClientConnInterface
}

func NewStudyCaseClient(cc grpc.ClientConnInterface) StudyCaseClient {
	return &studyCaseClient{cc}
}

func (c *studyCaseClient) ListCases(ctx context.Context, opts ...grpc.CallOption) (StudyCase_ListCasesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StudyCase_serviceDesc.Streams[0], "/StudyCase/ListCases", opts...)
	if err != nil {
		return nil, err
	}
	x := &studyCaseListCasesClient{stream}
	return x, nil
}

type StudyCase_ListCasesClient interface {
	Send(*Case) error
	CloseAndRecv() (*ServerResponse, error)
	grpc.ClientStream
}

type studyCaseListCasesClient struct {
	grpc.ClientStream
}

func (x *studyCaseListCasesClient) Send(m *Case) error {
	return x.ClientStream.SendMsg(m)
}

func (x *studyCaseListCasesClient) CloseAndRecv() (*ServerResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ServerResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StudyCaseServer is the server API for StudyCase service.
// All implementations must embed UnimplementedStudyCaseServer
// for forward compatibility
type StudyCaseServer interface {
	ListCases(StudyCase_ListCasesServer) error
	mustEmbedUnimplementedStudyCaseServer()
}

// UnimplementedStudyCaseServer must be embedded to have forward compatible implementations.
type UnimplementedStudyCaseServer struct {
}

func (UnimplementedStudyCaseServer) ListCases(StudyCase_ListCasesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListCases not implemented")
}
func (UnimplementedStudyCaseServer) mustEmbedUnimplementedStudyCaseServer() {}

// UnsafeStudyCaseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StudyCaseServer will
// result in compilation errors.
type UnsafeStudyCaseServer interface {
	mustEmbedUnimplementedStudyCaseServer()
}

func RegisterStudyCaseServer(s grpc.ServiceRegistrar, srv StudyCaseServer) {
	s.RegisterService(&_StudyCase_serviceDesc, srv)
}

func _StudyCase_ListCases_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StudyCaseServer).ListCases(&studyCaseListCasesServer{stream})
}

type StudyCase_ListCasesServer interface {
	SendAndClose(*ServerResponse) error
	Recv() (*Case, error)
	grpc.ServerStream
}

type studyCaseListCasesServer struct {
	grpc.ServerStream
}

func (x *studyCaseListCasesServer) SendAndClose(m *ServerResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *studyCaseListCasesServer) Recv() (*Case, error) {
	m := new(Case)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _StudyCase_serviceDesc = grpc.ServiceDesc{
	ServiceName: "StudyCase",
	HandlerType: (*StudyCaseServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListCases",
			Handler:       _StudyCase_ListCases_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "Cases.proto",
}
