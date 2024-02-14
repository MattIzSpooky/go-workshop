package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"log"
	"net"
	"slices"
	"workshop/grpc/chat"
)

type chatRoom struct {
	name        string
	users       []string
	msgChannels []chan *chat.ChatMessage
}

type chatServer struct {
	chat.UnimplementedChatServer

	rooms []chatRoom
}

func (s *chatServer) JoinRoom(joinRequest *chat.JoinRoomRequest, msgStream chat.Chat_JoinRoomServer) error {
	idx := slices.IndexFunc(s.rooms, func(c chatRoom) bool { return c.name == joinRequest.GetRoom() })

	// If room does not exist, create room.
	if idx == -1 {
		fmt.Println(fmt.Sprintf("Room [%s] does not exist. Creating room", joinRequest.GetRoom()))
		s.rooms = append(s.rooms, chatRoom{name: joinRequest.GetRoom()})
	}

	// Check if the current username is already used.
	for _, r := range s.rooms {
		usersIdx := slices.IndexFunc(r.users, func(u string) bool { return u == joinRequest.GetUsername() })
		if usersIdx != -1 {
			username := &r.users[usersIdx]
			return fmt.Errorf("username [%s] cannot be used. It is used in room [%s]", *username, r.name)
		}
	}

	idx = slices.IndexFunc(s.rooms, func(c chatRoom) bool { return c.name == joinRequest.GetRoom() })

	msgChannel := make(chan *chat.ChatMessage)

	currRoom := &s.rooms[idx]
	currRoom.users = append(currRoom.users, joinRequest.GetUsername())
	currRoom.msgChannels = append(currRoom.msgChannels, msgChannel)

	for {
		select {
		case <-msgStream.Context().Done():
			return nil
		case msg := <-msgChannel:
			fmt.Printf("GO ROUTINE (got message): %v \n", msg)
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

	idx := slices.IndexFunc(s.rooms, func(c chatRoom) bool { return c.name == msg.GetRoom() })

	currRoom := &s.rooms[idx]

	go func() {
		// When a message comes in, send it to all recipients in the room
		streams := currRoom.msgChannels
		for _, msgChan := range streams {
			msgChan <- msg
		}
	}()

	return nil
}

func (s *chatServer) ListRooms(ctx context.Context, empty *emptypb.Empty) (*chat.ListRoomsReply, error) {
	var rooms []string

	for _, room := range s.rooms {
		rooms = append(rooms, room.name)
	}

	return &chat.ListRoomsReply{Rooms: rooms}, nil
}

func (s *chatServer) GetChatUsers(ctx context.Context, chatUsersRequest *chat.ChatUsersRequest) (*chat.ChatUsersReply, error) {
	idx := slices.IndexFunc(s.rooms, func(c chatRoom) bool { return c.name == chatUsersRequest.GetRoom() })

	if idx == -1 {
		return nil, errors.New(fmt.Sprintf("Could not find room [%s]", chatUsersRequest.GetRoom()))
	}

	currRoom := s.rooms[idx]

	return &chat.ChatUsersReply{Users: currRoom.users}, nil
}

func main() {
	port := 8080
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

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
