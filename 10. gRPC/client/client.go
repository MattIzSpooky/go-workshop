package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"log"
	"os"
	"os/signal"
	"slices"
	"sync"
	"syscall"
	"workshop/grpc/chat"
)

var username = flag.String("username", "default", "Username of person using the chat")
var chatServer = flag.String("server", ":8080", "Chat server address")
var debugEnabled = flag.Bool("debug", false, "Enabled/disable debug logging")

const defaultRoomName = "default"

func joinChannel(ctx context.Context, client chat.ChatClient, room string) {

	joinRoomRequest := chat.JoinRoomRequest{Room: room, Username: *username}
	stream, err := client.JoinRoom(ctx, &joinRoomRequest)
	if err != nil {
		log.Fatalf("client.JoinChannel(ctx, &channel) throws: %v", err)
	}

	fmt.Printf("Joined room: %v \n", room)

	wg := new(sync.WaitGroup)
	wg.Add(1)

	// TODO: split this to another function
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				wg.Done()
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive message from channel joining. \nErr: %v", err)
			}

			if *username != in.GetUsername() {
				fmt.Println(fmt.Sprintf("[%s]: (%s) -> %s", color.MagentaString(in.GetRoom()), color.GreenString(in.GetUsername()), color.CyanString(in.GetMessage())))
			}
		}
	}()

	// Block until no more messages can be read.
	wg.Wait()
}

func sendMessage(ctx context.Context, client chat.ChatClient, room string, message string) {
	stream, err := client.SendMessage(ctx)
	if err != nil {
		log.Printf("Cannot send message: error: %v", err)
	}
	msg := chat.ChatMessage{
		Room:     room,
		Username: *username,
		Message:  message,
	}
	stream.Send(&msg)

	ack, err := stream.CloseAndRecv()

	if *debugEnabled {
		fmt.Printf("Message sent: %v \n", ack)
	}
}

func getExistingRoom(ctx context.Context, client chat.ChatClient, room string) (string, error) {
	var yesOrNoStr string
	fmt.Println("Do you still wish to join the room? [y/n]")
	fmt.Scanln(&yesOrNoStr)

	if yesOrNoStr != "y" {
		return "", fmt.Errorf("Not joining room [%s]", room)
	}

	return room, nil
}

func printRoomUsers(ctx context.Context, client chat.ChatClient, room string) {
	usersInChatReply, err := client.GetChatUsers(ctx, &chat.ChatUsersRequest{Room: room})

	if err != nil {
		panic(fmt.Sprintf("Error fetching users from room: [%s]. Error -> %s", room, err.Error()))
	}

	color.Cyan("Current users in chat: ")
	for _, user := range usersInChatReply.GetUsers() {
		color.Cyan(fmt.Sprintf("- %s", user))
	}
}

func main() {
	flag.Parse()

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(*chatServer, opts...)
	if err != nil {
		log.Fatalf("Fail to dail: %v", err)
	}

	ctx := context.Background()
	client := chat.NewChatClient(conn)

	listRoomsReply, err := client.ListRooms(ctx, &emptypb.Empty{})

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
			printRoomUsers(ctx, client, room)
			room, err = getExistingRoom(ctx, client, availableRooms[roomIdx])
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

	go joinChannel(ctx, client, room)
	go writeAndSendMessages(ctx, client, room)

	<-sigChan

	_, err = client.DisconnectFromRoom(ctx, &chat.DisconnectFromRoomMessage{
		Room:     room,
		Username: *username,
	})

	exitOnError(err)

	exitOnError(conn.Close())
}

func writeAndSendMessages(ctx context.Context, client chat.ChatClient, room string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		go sendMessage(ctx, client, room, scanner.Text())
	}
}

func exitOnError(err error) {
	if err != nil {
		color.Red(err.Error())
		os.Exit(-1)
	}
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

func printRooms(availableRooms []string) {
	color.Cyan("Current rooms: ")
	for i, room := range availableRooms {
		color.Cyan(fmt.Sprintf("%d: %s", i, room))
	}
}

func queryRoomNameFromUser(availableRooms []string) (string, error) {
	var roomNameFromUser string
	color.Green("What do you want the room to be called?")
	fmt.Scanln(&roomNameFromUser)

	newRoomIdx := slices.IndexFunc(availableRooms, func(r string) bool { return r == roomNameFromUser })

	if newRoomIdx != -1 {
		return "", fmt.Errorf("Room already exists [%s]. Closing client..", roomNameFromUser)
	}

	return roomNameFromUser, nil
}
