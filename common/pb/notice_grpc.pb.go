// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: notice.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// NoticeServiceClient is the client API for NoticeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NoticeServiceClient interface {
	//NoticeInsert 插入公告
	NoticeInsert(ctx context.Context, in *NoticeInsertReq, opts ...grpc.CallOption) (*NoticeInsertResp, error)
}

type noticeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNoticeServiceClient(cc grpc.ClientConnInterface) NoticeServiceClient {
	return &noticeServiceClient{cc}
}

func (c *noticeServiceClient) NoticeInsert(ctx context.Context, in *NoticeInsertReq, opts ...grpc.CallOption) (*NoticeInsertResp, error) {
	out := new(NoticeInsertResp)
	err := c.cc.Invoke(ctx, "/pb.noticeService/NoticeInsert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NoticeServiceServer is the server API for NoticeService service.
// All implementations must embed UnimplementedNoticeServiceServer
// for forward compatibility
type NoticeServiceServer interface {
	//NoticeInsert 插入公告
	NoticeInsert(context.Context, *NoticeInsertReq) (*NoticeInsertResp, error)
	mustEmbedUnimplementedNoticeServiceServer()
}

// UnimplementedNoticeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNoticeServiceServer struct {
}

func (UnimplementedNoticeServiceServer) NoticeInsert(context.Context, *NoticeInsertReq) (*NoticeInsertResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NoticeInsert not implemented")
}
func (UnimplementedNoticeServiceServer) mustEmbedUnimplementedNoticeServiceServer() {}

// UnsafeNoticeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NoticeServiceServer will
// result in compilation errors.
type UnsafeNoticeServiceServer interface {
	mustEmbedUnimplementedNoticeServiceServer()
}

func RegisterNoticeServiceServer(s grpc.ServiceRegistrar, srv NoticeServiceServer) {
	s.RegisterService(&NoticeService_ServiceDesc, srv)
}

func _NoticeService_NoticeInsert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoticeInsertReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoticeServiceServer).NoticeInsert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.noticeService/NoticeInsert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoticeServiceServer).NoticeInsert(ctx, req.(*NoticeInsertReq))
	}
	return interceptor(ctx, in, info, handler)
}

// NoticeService_ServiceDesc is the grpc.ServiceDesc for NoticeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NoticeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.noticeService",
	HandlerType: (*NoticeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NoticeInsert",
			Handler:    _NoticeService_NoticeInsert_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notice.proto",
}
