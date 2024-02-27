package main

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"google.golang.org/protobuf/types/known/emptypb"
	"slices"
	"workshop/grpc/generated/chat"
)

func queryRoomFromAvailableRooms(grpcClient chat.ChatClient, ctx context.Context) (string, error) {
	listRoomsReply, err := grpcClient.ListRooms(ctx, &emptypb.Empty{})

	if err != nil {
		return "", err
	}

	availableRoomsLength := len(listRoomsReply.Rooms)

	room := defaultRoomName

	if availableRoomsLength > 0 {
		printRooms(listRoomsReply.Rooms)
		roomIdx, err := queryUserForRoom(availableRoomsLength)

		exitOnError(err)

		// if room exists, ask user if it is okay to join it.
		if roomIdx != -1 {
			usersInChatReply, err := grpcClient.GetChatUsers(ctx, &chat.ChatUsersRequest{Room: room})
			if err != nil {
				panic(fmt.Sprintf("Error fetching users from room: [%s]. Error -> %s", room, err.Error()))
			}

			printRoomUsers(usersInChatReply.GetUsers())
			room, err = queryExistingRoom(listRoomsReply.Rooms[roomIdx])
		} else {
			room, err = queryRoomNameFromUser(listRoomsReply.Rooms)
		}

	} else {
		color.Magenta(fmt.Sprintf("There are currently no rooms. Attempting to create and join room: [%s]", room))
	}

	if err != nil {
		return "", err
	}

	return room, nil
}

func queryUserForRoom(availableRoomsLength int) (int, error) {
	var roomIdx int
	color.Green("Enter the number of the room you wish to join.")
	color.Cyan("Write -1 to create a new room.")
	fmt.Scanln(&roomIdx)

	if roomIdx > availableRoomsLength || roomIdx < -1 {
		return 0, fmt.Errorf("invalid room selected, selected #%d", roomIdx)
	}

	return roomIdx, nil
}

func queryRoomNameFromUser(availableRooms []string) (string, error) {
	var roomNameFromUser string
	color.Green("What do you want the room to be called?")
	fmt.Scanln(&roomNameFromUser)

	newRoomIdx := slices.IndexFunc(availableRooms, func(r string) bool { return r == roomNameFromUser })

	if newRoomIdx != -1 {
		return "", fmt.Errorf("Room already exists [%s]. Closing chatroom-app..", roomNameFromUser)
	}

	return roomNameFromUser, nil
}

func queryExistingRoom(room string) (string, error) {
	var yesOrNoStr string
	fmt.Println("Do you still wish to join the room? [y/n]")
	fmt.Scanln(&yesOrNoStr)

	if yesOrNoStr != "y" {
		return "", fmt.Errorf("Not joining room [%s]", room)
	}

	return room, nil
}
