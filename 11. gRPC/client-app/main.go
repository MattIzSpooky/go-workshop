package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"os"
	"os/signal"
	"syscall"
	"workshop/grpc/client-app/client"
	"workshop/grpc/generated/chat"
)

var username = flag.String("username", "default", "Username of person using the chat")
var chatServer = flag.String("server", ":8080", "Chat server address")
var debugEnabled = flag.Bool("debug", false, "Enabled/disable debug logging")

const defaultRoomName = "default"

func main() {
	flag.Parse()

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(*chatServer, opts...)
	if err != nil {
		log.Fatalf("Fail to dail: %v", err)
	}

	ctx := context.Background()
	grpcClient := chat.NewChatClient(conn)

	listRoomsReply, err := grpcClient.ListRooms(ctx, &emptypb.Empty{})

	exitOnError(err)

	availableRooms := listRoomsReply.Rooms
	availableRoomsLength := len(availableRooms)

	room := defaultRoomName

	if availableRoomsLength > 0 {
		printRooms(availableRooms)
		roomIdx, err := queryUserForRoom(availableRoomsLength)

		exitOnError(err)

		// if room exists, ask user if it is okay to join it.
		if roomIdx != -1 {
			usersInChatReply, err := grpcClient.GetChatUsers(ctx, &chat.ChatUsersRequest{Room: room})
			if err != nil {
				panic(fmt.Sprintf("Error fetching users from room: [%s]. Error -> %s", room, err.Error()))
			}

			printRoomUsers(usersInChatReply.GetUsers())
			room, err = queryExistingRoom(availableRooms[roomIdx])
		} else {
			room, err = queryRoomNameFromUser(availableRooms)
		}

		exitOnError(err)

	} else {
		color.Magenta(fmt.Sprintf("There are currently no rooms. Attempting to create and join room: [%s]", room))
	}

	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGKILL,
		syscall.SIGHUP,
	)

	cClient := client.NewChatClient(room, *username, ctx, grpcClient)
	err = cClient.JoinAndListen()
	exitOnError(err)

	cClient.WriteAndSendMessages()

	<-sigChan

	err = cClient.Disconnect()

	exitOnError(err)
	exitOnError(conn.Close())
}
