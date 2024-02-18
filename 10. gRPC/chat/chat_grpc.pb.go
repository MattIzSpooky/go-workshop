// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: chat/chat.proto

package chat

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatClient interface {
	GetChatUsers(ctx context.Context, in *ChatUsersRequest, opts ...grpc.CallOption) (*ChatUsersReply, error)
	ListRooms(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListRoomsReply, error)
	JoinRoom(ctx context.Context, in *JoinRoomRequest, opts ...grpc.CallOption) (*SuccessReply, error)
	ListenToRoom(ctx context.Context, in *JoinRoomRequest, opts ...grpc.CallOption) (Chat_ListenToRoomClient, error)
	SendMessage(ctx context.Context, opts ...grpc.CallOption) (Chat_SendMessageClient, error)
	DisconnectFromRoom(ctx context.Context, in *DisconnectFromRoomMessage, opts ...grpc.CallOption) (*SuccessReply, error)
	NotifyDisconnect(ctx context.Context, in *NotifyDisconnectRequest, opts ...grpc.CallOption) (*SuccessReply, error)
	NotifyJoin(ctx context.Context, in *NotifyJoinMessage, opts ...grpc.CallOption) (*SuccessReply, error)
	CheckRoomExists(ctx context.Context, in *CheckRoomExistsMessage, opts ...grpc.CallOption) (*SuccessReply, error)
}

type chatClient struct {
	cc grpc.ClientConnInterface
}

func NewChatClient(cc grpc.ClientConnInterface) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) GetChatUsers(ctx context.Context, in *ChatUsersRequest, opts ...grpc.CallOption) (*ChatUsersReply, error) {
	out := new(ChatUsersReply)
	err := c.cc.Invoke(ctx, "/chat.Chat/GetChatUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) ListRooms(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListRoomsReply, error) {
	out := new(ListRoomsReply)
	err := c.cc.Invoke(ctx, "/chat.Chat/ListRooms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) JoinRoom(ctx context.Context, in *JoinRoomRequest, opts ...grpc.CallOption) (*SuccessReply, error) {
	out := new(SuccessReply)
	err := c.cc.Invoke(ctx, "/chat.Chat/JoinRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) ListenToRoom(ctx context.Context, in *JoinRoomRequest, opts ...grpc.CallOption) (Chat_ListenToRoomClient, error) {
	stream, err := c.cc.NewStream(ctx, &Chat_ServiceDesc.Streams[0], "/chat.Chat/ListenToRoom", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatListenToRoomClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Chat_ListenToRoomClient interface {
	Recv() (*ChatMessage, error)
	grpc.ClientStream
}

type chatListenToRoomClient struct {
	grpc.ClientStream
}

func (x *chatListenToRoomClient) Recv() (*ChatMessage, error) {
	m := new(ChatMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatClient) SendMessage(ctx context.Context, opts ...grpc.CallOption) (Chat_SendMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &Chat_ServiceDesc.Streams[1], "/chat.Chat/SendMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatSendMessageClient{stream}
	return x, nil
}

type Chat_SendMessageClient interface {
	Send(*ChatMessage) error
	CloseAndRecv() (*MessageAck, error)
	grpc.ClientStream
}

type chatSendMessageClient struct {
	grpc.ClientStream
}

func (x *chatSendMessageClient) Send(m *ChatMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatSendMessageClient) CloseAndRecv() (*MessageAck, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(MessageAck)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatClient) DisconnectFromRoom(ctx context.Context, in *DisconnectFromRoomMessage, opts ...grpc.CallOption) (*SuccessReply, error) {
	out := new(SuccessReply)
	err := c.cc.Invoke(ctx, "/chat.Chat/DisconnectFromRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) NotifyDisconnect(ctx context.Context, in *NotifyDisconnectRequest, opts ...grpc.CallOption) (*SuccessReply, error) {
	out := new(SuccessReply)
	err := c.cc.Invoke(ctx, "/chat.Chat/NotifyDisconnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) NotifyJoin(ctx context.Context, in *NotifyJoinMessage, opts ...grpc.CallOption) (*SuccessReply, error) {
	out := new(SuccessReply)
	err := c.cc.Invoke(ctx, "/chat.Chat/NotifyJoin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) CheckRoomExists(ctx context.Context, in *CheckRoomExistsMessage, opts ...grpc.CallOption) (*SuccessReply, error) {
	out := new(SuccessReply)
	err := c.cc.Invoke(ctx, "/chat.Chat/CheckRoomExists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServer is the server API for Chat service.
// All implementations must embed UnimplementedChatServer
// for forward compatibility
type ChatServer interface {
	GetChatUsers(context.Context, *ChatUsersRequest) (*ChatUsersReply, error)
	ListRooms(context.Context, *emptypb.Empty) (*ListRoomsReply, error)
	JoinRoom(context.Context, *JoinRoomRequest) (*SuccessReply, error)
	ListenToRoom(*JoinRoomRequest, Chat_ListenToRoomServer) error
	SendMessage(Chat_SendMessageServer) error
	DisconnectFromRoom(context.Context, *DisconnectFromRoomMessage) (*SuccessReply, error)
	NotifyDisconnect(context.Context, *NotifyDisconnectRequest) (*SuccessReply, error)
	NotifyJoin(context.Context, *NotifyJoinMessage) (*SuccessReply, error)
	CheckRoomExists(context.Context, *CheckRoomExistsMessage) (*SuccessReply, error)
	mustEmbedUnimplementedChatServer()
}

// UnimplementedChatServer must be embedded to have forward compatible implementations.
type UnimplementedChatServer struct {
}

func (UnimplementedChatServer) GetChatUsers(context.Context, *ChatUsersRequest) (*ChatUsersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatUsers not implemented")
}
func (UnimplementedChatServer) ListRooms(context.Context, *emptypb.Empty) (*ListRoomsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRooms not implemented")
}
func (UnimplementedChatServer) JoinRoom(context.Context, *JoinRoomRequest) (*SuccessReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinRoom not implemented")
}
func (UnimplementedChatServer) ListenToRoom(*JoinRoomRequest, Chat_ListenToRoomServer) error {
	return status.Errorf(codes.Unimplemented, "method ListenToRoom not implemented")
}
func (UnimplementedChatServer) SendMessage(Chat_SendMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedChatServer) DisconnectFromRoom(context.Context, *DisconnectFromRoomMessage) (*SuccessReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisconnectFromRoom not implemented")
}
func (UnimplementedChatServer) NotifyDisconnect(context.Context, *NotifyDisconnectRequest) (*SuccessReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyDisconnect not implemented")
}
func (UnimplementedChatServer) NotifyJoin(context.Context, *NotifyJoinMessage) (*SuccessReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyJoin not implemented")
}
func (UnimplementedChatServer) CheckRoomExists(context.Context, *CheckRoomExistsMessage) (*SuccessReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckRoomExists not implemented")
}
func (UnimplementedChatServer) mustEmbedUnimplementedChatServer() {}

// UnsafeChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServer will
// result in compilation errors.
type UnsafeChatServer interface {
	mustEmbedUnimplementedChatServer()
}

func RegisterChatServer(s grpc.ServiceRegistrar, srv ChatServer) {
	s.RegisterService(&Chat_ServiceDesc, srv)
}

func _Chat_GetChatUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).GetChatUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.Chat/GetChatUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).GetChatUsers(ctx, req.(*ChatUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_ListRooms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).ListRooms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.Chat/ListRooms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).ListRooms(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_JoinRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).JoinRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.Chat/JoinRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).JoinRoom(ctx, req.(*JoinRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_ListenToRoom_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(JoinRoomRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServer).ListenToRoom(m, &chatListenToRoomServer{stream})
}

type Chat_ListenToRoomServer interface {
	Send(*ChatMessage) error
	grpc.ServerStream
}

type chatListenToRoomServer struct {
	grpc.ServerStream
}

func (x *chatListenToRoomServer) Send(m *ChatMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _Chat_SendMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServer).SendMessage(&chatSendMessageServer{stream})
}

type Chat_SendMessageServer interface {
	SendAndClose(*MessageAck) error
	Recv() (*ChatMessage, error)
	grpc.ServerStream
}

type chatSendMessageServer struct {
	grpc.ServerStream
}

func (x *chatSendMessageServer) SendAndClose(m *MessageAck) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatSendMessageServer) Recv() (*ChatMessage, error) {
	m := new(ChatMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Chat_DisconnectFromRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisconnectFromRoomMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).DisconnectFromRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.Chat/DisconnectFromRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).DisconnectFromRoom(ctx, req.(*DisconnectFromRoomMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_NotifyDisconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyDisconnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).NotifyDisconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.Chat/NotifyDisconnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).NotifyDisconnect(ctx, req.(*NotifyDisconnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_NotifyJoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyJoinMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).NotifyJoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.Chat/NotifyJoin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).NotifyJoin(ctx, req.(*NotifyJoinMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_CheckRoomExists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRoomExistsMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).CheckRoomExists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.Chat/CheckRoomExists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).CheckRoomExists(ctx, req.(*CheckRoomExistsMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// Chat_ServiceDesc is the grpc.ServiceDesc for Chat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetChatUsers",
			Handler:    _Chat_GetChatUsers_Handler,
		},
		{
			MethodName: "ListRooms",
			Handler:    _Chat_ListRooms_Handler,
		},
		{
			MethodName: "JoinRoom",
			Handler:    _Chat_JoinRoom_Handler,
		},
		{
			MethodName: "DisconnectFromRoom",
			Handler:    _Chat_DisconnectFromRoom_Handler,
		},
		{
			MethodName: "NotifyDisconnect",
			Handler:    _Chat_NotifyDisconnect_Handler,
		},
		{
			MethodName: "NotifyJoin",
			Handler:    _Chat_NotifyJoin_Handler,
		},
		{
			MethodName: "CheckRoomExists",
			Handler:    _Chat_CheckRoomExists_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListenToRoom",
			Handler:       _Chat_ListenToRoom_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SendMessage",
			Handler:       _Chat_SendMessage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "chat/chat.proto",
}
