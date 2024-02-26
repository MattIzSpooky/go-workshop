package main

import (
	"fmt"
	"github.com/fatih/color"
	"slices"
)

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
		return "", fmt.Errorf("Room already exists [%s]. Closing client-app..", roomNameFromUser)
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
