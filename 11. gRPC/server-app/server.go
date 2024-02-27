package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"slices"
	"sync"
	"time"
	"workshop/grpc/generated/chat"
)

type chatServer struct {
	chat.UnimplementedChatServer
	sync.Mutex

	rooms []*chatRoom
}

func (s *chatServer) ListenToRoom(joinRequest *chat.JoinRoomRequest, msgStream chat.Chat_ListenToRoomServer) error {
	currRoom := s.findRoom(joinRequest.GetRoom())
	user := currRoom.findUser(joinRequest.GetUsername())

	if user == nil {
		return fmt.Errorf("user [%s] does not exist in room [%s]", joinRequest.GetUsername(), joinRequest.GetRoom())
	}

	for {
		select {
		case <-msgStream.Context().Done():
			return nil
		case msg := <-user.msgChannel:
			s.printDebug(fmt.Sprintf("Got message: %v \n", msg))
			err := msgStream.Send(msg)

			if msg == nil {
				s.printDebug(fmt.Sprintf("User [%s] disconnected from room [%s]", user.name, currRoom.name))
				return nil
			}

			if err != nil {
				return err
			}
		}
	}
}

func (s *chatServer) printDebug(message string) {
	if *debugEnabled {
		fmt.Println(message)
	}
}

func (s *chatServer) findRoom(roomName string) *chatRoom {
	s.Lock()
	defer s.Unlock()

	idx := slices.IndexFunc(s.rooms, func(c *chatRoom) bool { return c.name == roomName })

	if idx == -1 {
		return nil
	}

	return s.rooms[idx]
}

func (s *chatServer) JoinRoom(ctx context.Context, joinRequest *chat.JoinRoomRequest) (*chat.SuccessReply, error) {
	room := s.findRoom(joinRequest.GetRoom())

	// If room does not exist, create room.
	if room == nil {
		s.printDebug(fmt.Sprintf("Room [%s] does not exist. Creating room", joinRequest.GetRoom()))
		room = s.addRoom(room, joinRequest)
	}

	// Check if the current username is already used.
	reply, err := s.checkIfUserAlreadyExists(joinRequest)

	if err != nil {
		return reply, err
	}

	room.addUser(joinRequest.GetUsername())

	s.printDebug(fmt.Sprintf("User [%s] joined room [%s]", joinRequest.GetUsername(), joinRequest.GetRoom()))

	return &chat.SuccessReply{Success: true}, nil
}

func (s *chatServer) deleteRoom(room *chatRoom) {
	s.Lock()
	defer s.Unlock()

	idx := slices.IndexFunc(s.rooms, func(c *chatRoom) bool { return c.name == room.name })
	s.rooms = slices.Delete(s.rooms, idx, idx+1)
}

func (s *chatServer) addRoom(room *chatRoom, joinRequest *chat.JoinRoomRequest) *chatRoom {
	s.Lock()
	defer s.Unlock()

	room = &chatRoom{name: joinRequest.GetRoom()}
	s.rooms = append(s.rooms, room)

	return room
}

func (s *chatServer) checkIfUserAlreadyExists(joinRequest *chat.JoinRoomRequest) (*chat.SuccessReply, error) {
	for _, r := range s.rooms {
		user := r.findUser(joinRequest.GetUsername())

		if user != nil {
			errorMsg := fmt.Sprintf("username [%s] cannot be used. It is used in room [%s]", user.name, r.name)
			s.printDebug(errorMsg)

			return &chat.SuccessReply{Success: false}, errors.New(errorMsg)
		}
	}
	return nil, nil
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

	currRoom := s.findRoom(msg.GetRoom())

	msg.Time = timestamppb.New(time.Now())

	for _, user := range currRoom.users {
		user.msgChannel <- msg
	}

	return nil
}

func (s *chatServer) ListRooms(ctx context.Context, _ *emptypb.Empty) (*chat.ListRoomsReply, error) {
	var rooms []string

	s.Lock()
	for _, room := range s.rooms {
		rooms = append(rooms, room.name)
	}
	s.Unlock()

	return &chat.ListRoomsReply{Rooms: rooms}, nil
}

func (s *chatServer) NotifyDisconnect(ctx context.Context, notifyDisconnectMessage *chat.NotifyDisconnectRequest) (*chat.SuccessReply, error) {
	currRoom := s.findRoom(notifyDisconnectMessage.GetRoom())

	notifyDisconnectMsg := &chat.ChatMessage{
		Room:     currRoom.name,
		Username: "SYSTEM",
		Message:  fmt.Sprintf("User [%s] has disconnected.", notifyDisconnectMessage.GetUsername()),
		Time:     timestamppb.New(time.Now()),
	}

	for _, user := range currRoom.users {
		user.msgChannel <- notifyDisconnectMsg
	}

	return &chat.SuccessReply{Success: true}, nil
}

func (s *chatServer) CheckRoomExists(ctx context.Context, checkRoomExistsMessage *chat.CheckRoomExistsMessage) (*chat.SuccessReply, error) {
	idx := slices.IndexFunc(s.rooms, func(c *chatRoom) bool { return c.name == checkRoomExistsMessage.GetRoom() })

	if idx == -1 {
		return &chat.SuccessReply{Success: false}, nil
	}

	return &chat.SuccessReply{Success: true}, nil
}

func (s *chatServer) NotifyJoin(ctx context.Context, notifyJoinMessage *chat.NotifyJoinMessage) (*chat.SuccessReply, error) {
	currRoom := s.findRoom(notifyJoinMessage.GetRoom())

	notifyJoinMsg := &chat.ChatMessage{
		Room:     currRoom.name,
		Username: "SYSTEM",
		Message:  fmt.Sprintf("User [%s] has joined the room.", notifyJoinMessage.GetUsername()),
		Time:     timestamppb.New(time.Now()),
	}

	for _, user := range currRoom.users {
		user.msgChannel <- notifyJoinMsg
	}

	return &chat.SuccessReply{Success: true}, nil
}

func (s *chatServer) DisconnectFromRoom(ctx context.Context, disconnectFromRoomMessage *chat.DisconnectFromRoomMessage) (*chat.SuccessReply, error) {
	currRoom := s.findRoom(disconnectFromRoomMessage.GetRoom())

	err := currRoom.deleteUser(disconnectFromRoomMessage.GetUsername())

	if err != nil {
		return &chat.SuccessReply{Success: false}, errors.Join(fmt.Errorf("user does not exist: %s", disconnectFromRoomMessage.GetUsername()), err)
	}

	if currRoom.isEmpty() {
		s.printDebug(fmt.Sprintf("No more users in room [%s]. Deleting room...", currRoom.name))
		s.deleteRoom(currRoom)
		s.printDebug(fmt.Sprintf("Room [%s] deleted.", currRoom.name))
	}

	return &chat.SuccessReply{Success: true}, nil
}

func (s *chatServer) GetChatUsers(ctx context.Context, chatUsersRequest *chat.ChatUsersRequest) (*chat.ChatUsersReply, error) {
	currRoom := s.findRoom(chatUsersRequest.GetRoom())

	if currRoom == nil {
		return nil, errors.New(fmt.Sprintf("Could not find room [%s]", chatUsersRequest.GetRoom()))
	}

	var usernames []string

	for _, user := range currRoom.users {
		usernames = append(usernames, user.name)
	}

	return &chat.ChatUsersReply{Users: usernames}, nil
}
