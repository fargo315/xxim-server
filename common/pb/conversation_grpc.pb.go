// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: conversation.proto

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

// GroupServiceClient is the client API for GroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GroupServiceClient interface {
	//GroupCreate 创建群组
	GroupCreate(ctx context.Context, in *GroupCreateReq, opts ...grpc.CallOption) (*GroupCreateResp, error)
	//CountJoinGroup 统计加入的群组数量
	CountJoinGroup(ctx context.Context, in *CountJoinGroupReq, opts ...grpc.CallOption) (*CountJoinGroupResp, error)
	//CountCreateGroup 统计创建的群组数量
	CountCreateGroup(ctx context.Context, in *CountCreateGroupReq, opts ...grpc.CallOption) (*CountCreateGroupResp, error)
	//GroupSubscribe 群组订阅
	GroupSubscribe(ctx context.Context, in *GroupSubscribeReq, opts ...grpc.CallOption) (*GroupSubscribeResp, error)
	//ListGroupSubscribers 列出群组订阅者
	ListGroupSubscribers(ctx context.Context, in *ListGroupSubscribersReq, opts ...grpc.CallOption) (*ListGroupSubscribersResp, error)
}

type groupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupServiceClient(cc grpc.ClientConnInterface) GroupServiceClient {
	return &groupServiceClient{cc}
}

func (c *groupServiceClient) GroupCreate(ctx context.Context, in *GroupCreateReq, opts ...grpc.CallOption) (*GroupCreateResp, error) {
	out := new(GroupCreateResp)
	err := c.cc.Invoke(ctx, "/pb.groupService/GroupCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) CountJoinGroup(ctx context.Context, in *CountJoinGroupReq, opts ...grpc.CallOption) (*CountJoinGroupResp, error) {
	out := new(CountJoinGroupResp)
	err := c.cc.Invoke(ctx, "/pb.groupService/CountJoinGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) CountCreateGroup(ctx context.Context, in *CountCreateGroupReq, opts ...grpc.CallOption) (*CountCreateGroupResp, error) {
	out := new(CountCreateGroupResp)
	err := c.cc.Invoke(ctx, "/pb.groupService/CountCreateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) GroupSubscribe(ctx context.Context, in *GroupSubscribeReq, opts ...grpc.CallOption) (*GroupSubscribeResp, error) {
	out := new(GroupSubscribeResp)
	err := c.cc.Invoke(ctx, "/pb.groupService/GroupSubscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) ListGroupSubscribers(ctx context.Context, in *ListGroupSubscribersReq, opts ...grpc.CallOption) (*ListGroupSubscribersResp, error) {
	out := new(ListGroupSubscribersResp)
	err := c.cc.Invoke(ctx, "/pb.groupService/ListGroupSubscribers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupServiceServer is the server API for GroupService service.
// All implementations must embed UnimplementedGroupServiceServer
// for forward compatibility
type GroupServiceServer interface {
	//GroupCreate 创建群组
	GroupCreate(context.Context, *GroupCreateReq) (*GroupCreateResp, error)
	//CountJoinGroup 统计加入的群组数量
	CountJoinGroup(context.Context, *CountJoinGroupReq) (*CountJoinGroupResp, error)
	//CountCreateGroup 统计创建的群组数量
	CountCreateGroup(context.Context, *CountCreateGroupReq) (*CountCreateGroupResp, error)
	//GroupSubscribe 群组订阅
	GroupSubscribe(context.Context, *GroupSubscribeReq) (*GroupSubscribeResp, error)
	//ListGroupSubscribers 列出群组订阅者
	ListGroupSubscribers(context.Context, *ListGroupSubscribersReq) (*ListGroupSubscribersResp, error)
	mustEmbedUnimplementedGroupServiceServer()
}

// UnimplementedGroupServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGroupServiceServer struct {
}

func (UnimplementedGroupServiceServer) GroupCreate(context.Context, *GroupCreateReq) (*GroupCreateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupCreate not implemented")
}
func (UnimplementedGroupServiceServer) CountJoinGroup(context.Context, *CountJoinGroupReq) (*CountJoinGroupResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountJoinGroup not implemented")
}
func (UnimplementedGroupServiceServer) CountCreateGroup(context.Context, *CountCreateGroupReq) (*CountCreateGroupResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountCreateGroup not implemented")
}
func (UnimplementedGroupServiceServer) GroupSubscribe(context.Context, *GroupSubscribeReq) (*GroupSubscribeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupSubscribe not implemented")
}
func (UnimplementedGroupServiceServer) ListGroupSubscribers(context.Context, *ListGroupSubscribersReq) (*ListGroupSubscribersResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGroupSubscribers not implemented")
}
func (UnimplementedGroupServiceServer) mustEmbedUnimplementedGroupServiceServer() {}

// UnsafeGroupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GroupServiceServer will
// result in compilation errors.
type UnsafeGroupServiceServer interface {
	mustEmbedUnimplementedGroupServiceServer()
}

func RegisterGroupServiceServer(s grpc.ServiceRegistrar, srv GroupServiceServer) {
	s.RegisterService(&GroupService_ServiceDesc, srv)
}

func _GroupService_GroupCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).GroupCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.groupService/GroupCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).GroupCreate(ctx, req.(*GroupCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_CountJoinGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountJoinGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).CountJoinGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.groupService/CountJoinGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).CountJoinGroup(ctx, req.(*CountJoinGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_CountCreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountCreateGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).CountCreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.groupService/CountCreateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).CountCreateGroup(ctx, req.(*CountCreateGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_GroupSubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupSubscribeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).GroupSubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.groupService/GroupSubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).GroupSubscribe(ctx, req.(*GroupSubscribeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_ListGroupSubscribers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGroupSubscribersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).ListGroupSubscribers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.groupService/ListGroupSubscribers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).ListGroupSubscribers(ctx, req.(*ListGroupSubscribersReq))
	}
	return interceptor(ctx, in, info, handler)
}

// GroupService_ServiceDesc is the grpc.ServiceDesc for GroupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GroupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.groupService",
	HandlerType: (*GroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GroupCreate",
			Handler:    _GroupService_GroupCreate_Handler,
		},
		{
			MethodName: "CountJoinGroup",
			Handler:    _GroupService_CountJoinGroup_Handler,
		},
		{
			MethodName: "CountCreateGroup",
			Handler:    _GroupService_CountCreateGroup_Handler,
		},
		{
			MethodName: "GroupSubscribe",
			Handler:    _GroupService_GroupSubscribe_Handler,
		},
		{
			MethodName: "ListGroupSubscribers",
			Handler:    _GroupService_ListGroupSubscribers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "conversation.proto",
}

// FriendServiceClient is the client API for FriendService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FriendServiceClient interface {
	//FriendApply 添加好友
	FriendApply(ctx context.Context, in *FriendApplyReq, opts ...grpc.CallOption) (*FriendApplyResp, error)
	//FriendApplyHandle 处理好友申请
	FriendApplyHandle(ctx context.Context, in *FriendApplyHandleReq, opts ...grpc.CallOption) (*FriendApplyHandleResp, error)
	//ListFriendApply 列出好友申请
	ListFriendApply(ctx context.Context, in *ListFriendApplyReq, opts ...grpc.CallOption) (*ListFriendApplyResp, error)
	//CountFriend 统计好友数量
	CountFriend(ctx context.Context, in *CountFriendReq, opts ...grpc.CallOption) (*CountFriendResp, error)
}

type friendServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFriendServiceClient(cc grpc.ClientConnInterface) FriendServiceClient {
	return &friendServiceClient{cc}
}

func (c *friendServiceClient) FriendApply(ctx context.Context, in *FriendApplyReq, opts ...grpc.CallOption) (*FriendApplyResp, error) {
	out := new(FriendApplyResp)
	err := c.cc.Invoke(ctx, "/pb.friendService/FriendApply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendServiceClient) FriendApplyHandle(ctx context.Context, in *FriendApplyHandleReq, opts ...grpc.CallOption) (*FriendApplyHandleResp, error) {
	out := new(FriendApplyHandleResp)
	err := c.cc.Invoke(ctx, "/pb.friendService/FriendApplyHandle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendServiceClient) ListFriendApply(ctx context.Context, in *ListFriendApplyReq, opts ...grpc.CallOption) (*ListFriendApplyResp, error) {
	out := new(ListFriendApplyResp)
	err := c.cc.Invoke(ctx, "/pb.friendService/ListFriendApply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendServiceClient) CountFriend(ctx context.Context, in *CountFriendReq, opts ...grpc.CallOption) (*CountFriendResp, error) {
	out := new(CountFriendResp)
	err := c.cc.Invoke(ctx, "/pb.friendService/CountFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FriendServiceServer is the server API for FriendService service.
// All implementations must embed UnimplementedFriendServiceServer
// for forward compatibility
type FriendServiceServer interface {
	//FriendApply 添加好友
	FriendApply(context.Context, *FriendApplyReq) (*FriendApplyResp, error)
	//FriendApplyHandle 处理好友申请
	FriendApplyHandle(context.Context, *FriendApplyHandleReq) (*FriendApplyHandleResp, error)
	//ListFriendApply 列出好友申请
	ListFriendApply(context.Context, *ListFriendApplyReq) (*ListFriendApplyResp, error)
	//CountFriend 统计好友数量
	CountFriend(context.Context, *CountFriendReq) (*CountFriendResp, error)
	mustEmbedUnimplementedFriendServiceServer()
}

// UnimplementedFriendServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFriendServiceServer struct {
}

func (UnimplementedFriendServiceServer) FriendApply(context.Context, *FriendApplyReq) (*FriendApplyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FriendApply not implemented")
}
func (UnimplementedFriendServiceServer) FriendApplyHandle(context.Context, *FriendApplyHandleReq) (*FriendApplyHandleResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FriendApplyHandle not implemented")
}
func (UnimplementedFriendServiceServer) ListFriendApply(context.Context, *ListFriendApplyReq) (*ListFriendApplyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFriendApply not implemented")
}
func (UnimplementedFriendServiceServer) CountFriend(context.Context, *CountFriendReq) (*CountFriendResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountFriend not implemented")
}
func (UnimplementedFriendServiceServer) mustEmbedUnimplementedFriendServiceServer() {}

// UnsafeFriendServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FriendServiceServer will
// result in compilation errors.
type UnsafeFriendServiceServer interface {
	mustEmbedUnimplementedFriendServiceServer()
}

func RegisterFriendServiceServer(s grpc.ServiceRegistrar, srv FriendServiceServer) {
	s.RegisterService(&FriendService_ServiceDesc, srv)
}

func _FriendService_FriendApply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FriendApplyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServiceServer).FriendApply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.friendService/FriendApply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServiceServer).FriendApply(ctx, req.(*FriendApplyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendService_FriendApplyHandle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FriendApplyHandleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServiceServer).FriendApplyHandle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.friendService/FriendApplyHandle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServiceServer).FriendApplyHandle(ctx, req.(*FriendApplyHandleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendService_ListFriendApply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFriendApplyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServiceServer).ListFriendApply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.friendService/ListFriendApply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServiceServer).ListFriendApply(ctx, req.(*ListFriendApplyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendService_CountFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountFriendReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServiceServer).CountFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.friendService/CountFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServiceServer).CountFriend(ctx, req.(*CountFriendReq))
	}
	return interceptor(ctx, in, info, handler)
}

// FriendService_ServiceDesc is the grpc.ServiceDesc for FriendService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FriendService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.friendService",
	HandlerType: (*FriendServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FriendApply",
			Handler:    _FriendService_FriendApply_Handler,
		},
		{
			MethodName: "FriendApplyHandle",
			Handler:    _FriendService_FriendApplyHandle_Handler,
		},
		{
			MethodName: "ListFriendApply",
			Handler:    _FriendService_ListFriendApply_Handler,
		},
		{
			MethodName: "CountFriend",
			Handler:    _FriendService_CountFriend_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "conversation.proto",
}

// ConversationServiceClient is the client API for ConversationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConversationServiceClient interface {
	//ConversationSettingUpdate 更新会话设置
	ConversationSettingUpdate(ctx context.Context, in *ConversationSettingUpdateReq, opts ...grpc.CallOption) (*ConversationSettingUpdateResp, error)
	//ListJoinedConversations 列出加入的会话
	ListJoinedConversations(ctx context.Context, in *ListJoinedConversationsReq, opts ...grpc.CallOption) (*ListJoinedConversationsResp, error)
}

type conversationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConversationServiceClient(cc grpc.ClientConnInterface) ConversationServiceClient {
	return &conversationServiceClient{cc}
}

func (c *conversationServiceClient) ConversationSettingUpdate(ctx context.Context, in *ConversationSettingUpdateReq, opts ...grpc.CallOption) (*ConversationSettingUpdateResp, error) {
	out := new(ConversationSettingUpdateResp)
	err := c.cc.Invoke(ctx, "/pb.conversationService/ConversationSettingUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conversationServiceClient) ListJoinedConversations(ctx context.Context, in *ListJoinedConversationsReq, opts ...grpc.CallOption) (*ListJoinedConversationsResp, error) {
	out := new(ListJoinedConversationsResp)
	err := c.cc.Invoke(ctx, "/pb.conversationService/ListJoinedConversations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConversationServiceServer is the server API for ConversationService service.
// All implementations must embed UnimplementedConversationServiceServer
// for forward compatibility
type ConversationServiceServer interface {
	//ConversationSettingUpdate 更新会话设置
	ConversationSettingUpdate(context.Context, *ConversationSettingUpdateReq) (*ConversationSettingUpdateResp, error)
	//ListJoinedConversations 列出加入的会话
	ListJoinedConversations(context.Context, *ListJoinedConversationsReq) (*ListJoinedConversationsResp, error)
	mustEmbedUnimplementedConversationServiceServer()
}

// UnimplementedConversationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedConversationServiceServer struct {
}

func (UnimplementedConversationServiceServer) ConversationSettingUpdate(context.Context, *ConversationSettingUpdateReq) (*ConversationSettingUpdateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConversationSettingUpdate not implemented")
}
func (UnimplementedConversationServiceServer) ListJoinedConversations(context.Context, *ListJoinedConversationsReq) (*ListJoinedConversationsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListJoinedConversations not implemented")
}
func (UnimplementedConversationServiceServer) mustEmbedUnimplementedConversationServiceServer() {}

// UnsafeConversationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConversationServiceServer will
// result in compilation errors.
type UnsafeConversationServiceServer interface {
	mustEmbedUnimplementedConversationServiceServer()
}

func RegisterConversationServiceServer(s grpc.ServiceRegistrar, srv ConversationServiceServer) {
	s.RegisterService(&ConversationService_ServiceDesc, srv)
}

func _ConversationService_ConversationSettingUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConversationSettingUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversationServiceServer).ConversationSettingUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.conversationService/ConversationSettingUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversationServiceServer).ConversationSettingUpdate(ctx, req.(*ConversationSettingUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConversationService_ListJoinedConversations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListJoinedConversationsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversationServiceServer).ListJoinedConversations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.conversationService/ListJoinedConversations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversationServiceServer).ListJoinedConversations(ctx, req.(*ListJoinedConversationsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ConversationService_ServiceDesc is the grpc.ServiceDesc for ConversationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConversationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.conversationService",
	HandlerType: (*ConversationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConversationSettingUpdate",
			Handler:    _ConversationService_ConversationSettingUpdate_Handler,
		},
		{
			MethodName: "ListJoinedConversations",
			Handler:    _ConversationService_ListJoinedConversations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "conversation.proto",
}

// SubscriptionServiceClient is the client API for SubscriptionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubscriptionServiceClient interface {
	//SubscriptionSubscribe 订阅号订阅
	SubscriptionSubscribe(ctx context.Context, in *SubscriptionSubscribeReq, opts ...grpc.CallOption) (*SubscriptionSubscribeResp, error)
	//SubscriptionAfterOnline 订阅号在做用户上线后的操作
	SubscriptionAfterOnline(ctx context.Context, in *SubscriptionAfterOnlineReq, opts ...grpc.CallOption) (*SubscriptionAfterOnlineResp, error)
	//UpsertUserSubscription 更新用户订阅的订阅号
	UpsertUserSubscription(ctx context.Context, in *UpsertUserSubscriptionReq, opts ...grpc.CallOption) (*UpsertUserSubscriptionResp, error)
	//DeleteUserSubscription 删除用户订阅的订阅号
	DeleteUserSubscription(ctx context.Context, in *DeleteUserSubscriptionReq, opts ...grpc.CallOption) (*DeleteUserSubscriptionResp, error)
	//ListSubscriptionSubscribers 列出订阅号订阅者
	ListSubscriptionSubscribers(ctx context.Context, in *ListSubscriptionSubscribersReq, opts ...grpc.CallOption) (*ListSubscriptionSubscribersResp, error)
}

type subscriptionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSubscriptionServiceClient(cc grpc.ClientConnInterface) SubscriptionServiceClient {
	return &subscriptionServiceClient{cc}
}

func (c *subscriptionServiceClient) SubscriptionSubscribe(ctx context.Context, in *SubscriptionSubscribeReq, opts ...grpc.CallOption) (*SubscriptionSubscribeResp, error) {
	out := new(SubscriptionSubscribeResp)
	err := c.cc.Invoke(ctx, "/pb.subscriptionService/SubscriptionSubscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) SubscriptionAfterOnline(ctx context.Context, in *SubscriptionAfterOnlineReq, opts ...grpc.CallOption) (*SubscriptionAfterOnlineResp, error) {
	out := new(SubscriptionAfterOnlineResp)
	err := c.cc.Invoke(ctx, "/pb.subscriptionService/SubscriptionAfterOnline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) UpsertUserSubscription(ctx context.Context, in *UpsertUserSubscriptionReq, opts ...grpc.CallOption) (*UpsertUserSubscriptionResp, error) {
	out := new(UpsertUserSubscriptionResp)
	err := c.cc.Invoke(ctx, "/pb.subscriptionService/UpsertUserSubscription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) DeleteUserSubscription(ctx context.Context, in *DeleteUserSubscriptionReq, opts ...grpc.CallOption) (*DeleteUserSubscriptionResp, error) {
	out := new(DeleteUserSubscriptionResp)
	err := c.cc.Invoke(ctx, "/pb.subscriptionService/DeleteUserSubscription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) ListSubscriptionSubscribers(ctx context.Context, in *ListSubscriptionSubscribersReq, opts ...grpc.CallOption) (*ListSubscriptionSubscribersResp, error) {
	out := new(ListSubscriptionSubscribersResp)
	err := c.cc.Invoke(ctx, "/pb.subscriptionService/ListSubscriptionSubscribers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubscriptionServiceServer is the server API for SubscriptionService service.
// All implementations must embed UnimplementedSubscriptionServiceServer
// for forward compatibility
type SubscriptionServiceServer interface {
	//SubscriptionSubscribe 订阅号订阅
	SubscriptionSubscribe(context.Context, *SubscriptionSubscribeReq) (*SubscriptionSubscribeResp, error)
	//SubscriptionAfterOnline 订阅号在做用户上线后的操作
	SubscriptionAfterOnline(context.Context, *SubscriptionAfterOnlineReq) (*SubscriptionAfterOnlineResp, error)
	//UpsertUserSubscription 更新用户订阅的订阅号
	UpsertUserSubscription(context.Context, *UpsertUserSubscriptionReq) (*UpsertUserSubscriptionResp, error)
	//DeleteUserSubscription 删除用户订阅的订阅号
	DeleteUserSubscription(context.Context, *DeleteUserSubscriptionReq) (*DeleteUserSubscriptionResp, error)
	//ListSubscriptionSubscribers 列出订阅号订阅者
	ListSubscriptionSubscribers(context.Context, *ListSubscriptionSubscribersReq) (*ListSubscriptionSubscribersResp, error)
	mustEmbedUnimplementedSubscriptionServiceServer()
}

// UnimplementedSubscriptionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSubscriptionServiceServer struct {
}

func (UnimplementedSubscriptionServiceServer) SubscriptionSubscribe(context.Context, *SubscriptionSubscribeReq) (*SubscriptionSubscribeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubscriptionSubscribe not implemented")
}
func (UnimplementedSubscriptionServiceServer) SubscriptionAfterOnline(context.Context, *SubscriptionAfterOnlineReq) (*SubscriptionAfterOnlineResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubscriptionAfterOnline not implemented")
}
func (UnimplementedSubscriptionServiceServer) UpsertUserSubscription(context.Context, *UpsertUserSubscriptionReq) (*UpsertUserSubscriptionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertUserSubscription not implemented")
}
func (UnimplementedSubscriptionServiceServer) DeleteUserSubscription(context.Context, *DeleteUserSubscriptionReq) (*DeleteUserSubscriptionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserSubscription not implemented")
}
func (UnimplementedSubscriptionServiceServer) ListSubscriptionSubscribers(context.Context, *ListSubscriptionSubscribersReq) (*ListSubscriptionSubscribersResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSubscriptionSubscribers not implemented")
}
func (UnimplementedSubscriptionServiceServer) mustEmbedUnimplementedSubscriptionServiceServer() {}

// UnsafeSubscriptionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubscriptionServiceServer will
// result in compilation errors.
type UnsafeSubscriptionServiceServer interface {
	mustEmbedUnimplementedSubscriptionServiceServer()
}

func RegisterSubscriptionServiceServer(s grpc.ServiceRegistrar, srv SubscriptionServiceServer) {
	s.RegisterService(&SubscriptionService_ServiceDesc, srv)
}

func _SubscriptionService_SubscriptionSubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscriptionSubscribeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).SubscriptionSubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.subscriptionService/SubscriptionSubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).SubscriptionSubscribe(ctx, req.(*SubscriptionSubscribeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_SubscriptionAfterOnline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscriptionAfterOnlineReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).SubscriptionAfterOnline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.subscriptionService/SubscriptionAfterOnline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).SubscriptionAfterOnline(ctx, req.(*SubscriptionAfterOnlineReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_UpsertUserSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertUserSubscriptionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).UpsertUserSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.subscriptionService/UpsertUserSubscription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).UpsertUserSubscription(ctx, req.(*UpsertUserSubscriptionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_DeleteUserSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserSubscriptionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).DeleteUserSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.subscriptionService/DeleteUserSubscription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).DeleteUserSubscription(ctx, req.(*DeleteUserSubscriptionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_ListSubscriptionSubscribers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSubscriptionSubscribersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).ListSubscriptionSubscribers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.subscriptionService/ListSubscriptionSubscribers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).ListSubscriptionSubscribers(ctx, req.(*ListSubscriptionSubscribersReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SubscriptionService_ServiceDesc is the grpc.ServiceDesc for SubscriptionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SubscriptionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.subscriptionService",
	HandlerType: (*SubscriptionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubscriptionSubscribe",
			Handler:    _SubscriptionService_SubscriptionSubscribe_Handler,
		},
		{
			MethodName: "SubscriptionAfterOnline",
			Handler:    _SubscriptionService_SubscriptionAfterOnline_Handler,
		},
		{
			MethodName: "UpsertUserSubscription",
			Handler:    _SubscriptionService_UpsertUserSubscription_Handler,
		},
		{
			MethodName: "DeleteUserSubscription",
			Handler:    _SubscriptionService_DeleteUserSubscription_Handler,
		},
		{
			MethodName: "ListSubscriptionSubscribers",
			Handler:    _SubscriptionService_ListSubscriptionSubscribers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "conversation.proto",
}
