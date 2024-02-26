package main

import (
	"fmt"
	"github.com/fatih/color"
)

func printRooms(availableRooms []string) {
	color.Cyan("Current rooms: ")
	for i, room := range availableRooms {
		color.Cyan(fmt.Sprintf("%d: %s", i, room))
	}
}

func printRoomUsers(users []string) {
	color.Cyan("Current users in chat: ")
	for _, user := range users {
		color.Cyan(fmt.Sprintf("- %s", user))
	}
}
