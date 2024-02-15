package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"log"
	"net"
	"slices"
	"sync"
	"workshop/grpc/chat"
)

var debugEnabled = flag.Bool("debug", false, "Enabled/disable debug logging")
var port = flag.Int("port", 8080, "The port for the chat server to start listening on")

type chatRoom struct {
	lock sync.Mutex
	name string
	//users       []string
	//msgChannels []chan *chat.ChatMessage
	users []user
}

type user struct {
	name       string
	msgChannel chan *chat.ChatMessage
}

type chatServer struct {
	chat.UnimplementedChatServer
	sync.Mutex

	rooms []chatRoom
}

func (s *chatServer) JoinRoom(joinRequest *chat.JoinRoomRequest, msgStream chat.Chat_JoinRoomServer) error {
	s.Lock()

	idx := slices.IndexFunc(s.rooms, func(c chatRoom) bool { return c.name == joinRequest.GetRoom() })

	// If room does not exist, create room.
	if idx == -1 {
		fmt.Println(fmt.Sprintf("Room [%s] does not exist. Creating room", joinRequest.GetRoom()))
		s.rooms = append(s.rooms, chatRoom{name: joinRequest.GetRoom(), lock: sync.Mutex{}})
	}

	// Check if the current username is already used.
	for _, r := range s.rooms {
		usersIdx := slices.IndexFunc(r.users, func(u user) bool { return u.name == joinRequest.GetUsername() })
		if usersIdx != -1 {
			s.Unlock()
			user := &r.users[usersIdx]
			return fmt.Errorf("username [%s] cannot be used. It is used in room [%s]", user.name, r.name)
		}
	}

	idx = slices.IndexFunc(s.rooms, func(c chatRoom) bool { return c.name == joinRequest.GetRoom() })

	msgChannel := make(chan *chat.ChatMessage)

	currRoom := &s.rooms[idx]
	currRoom.users = append(currRoom.users, user{
		name:       joinRequest.GetUsername(),
		msgChannel: msgChannel,
	})
	s.Unlock()

	for {
		select {
		case <-msgStream.Context().Done():
			return nil
		case msg := <-msgChannel:
			if *debugEnabled {
				fmt.Printf("Got message: %v \n", msg)
			}

			msgStream.SendMsg(msg)
		}
	}
}

func (s *chatServer) SendMessage(msgStream chat.Chat_SendMessageServer) error {
	msg, err := msgStream.Recv()

	if err == io.EOF {
		return nil
	}

	if err != nil {
		return err
	}

	ack := chat.MessageAck{Status: "SENT"}
	msgStream.SendAndClose(&ack)

	s.Lock()
	idx := slices.IndexFunc(s.rooms, func(c chatRoom) bool { return c.name == msg.GetRoom() })
	s.Unlock()

	currRoom := &s.rooms[idx]

	go func() {
		// When a message comes in, send it to all recipients in the room
		for _, user := range currRoom.users {
			user.msgChannel <- msg
		}
	}()

	return nil
}

func (s *chatServer) ListRooms(ctx context.Context, empty *emptypb.Empty) (*chat.ListRoomsReply, error) {
	var rooms []string

	s.Lock()
	for _, room := range s.rooms {
		rooms = append(rooms, room.name)
	}
	s.Unlock()

	return &chat.ListRoomsReply{Rooms: rooms}, nil
}

func (s *chatServer) DisconnectFromRoom(ctx context.Context, disconnectFromRoomMessage *chat.DisconnectFromRoomMessage) (*chat.DisconnectFromRoomReply, error) {
	s.Lock()
	idx := slices.IndexFunc(s.rooms, func(c chatRoom) bool { return c.name == disconnectFromRoomMessage.GetRoom() })
	currRoom := &s.rooms[idx]
	s.Unlock()

	currRoom.lock.Lock()
	usersIdx := slices.IndexFunc(currRoom.users, func(u user) bool { return u.name == disconnectFromRoomMessage.GetUsername() })

	if usersIdx == -1 {
		return &chat.DisconnectFromRoomReply{Success: false}, fmt.Errorf("User does not exist: %s", disconnectFromRoomMessage.GetUsername())
	}

	currentUser := currRoom.users[usersIdx]

	close(currentUser.msgChannel)
	currRoom.users = slices.Delete(currRoom.users, usersIdx, usersIdx+1)

	currRoom.lock.Unlock()

	return &chat.DisconnectFromRoomReply{Success: true}, nil
}

func (s *chatServer) GetChatUsers(ctx context.Context, chatUsersRequest *chat.ChatUsersRequest) (*chat.ChatUsersReply, error) {
	s.Lock()
	idx := slices.IndexFunc(s.rooms, func(c chatRoom) bool { return c.name == chatUsersRequest.GetRoom() })
	s.Unlock()

	if idx == -1 {
		return nil, errors.New(fmt.Sprintf("Could not find room [%s]", chatUsersRequest.GetRoom()))
	}

	currRoom := &s.rooms[idx]

	var usernames []string

	for _, user := range currRoom.users {
		usernames = append(usernames, user.name)
	}

	return &chat.ChatUsersReply{Users: usernames}, nil
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	chat.RegisterChatServer(grpcServer, &chatServer{rooms: nil})

	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
