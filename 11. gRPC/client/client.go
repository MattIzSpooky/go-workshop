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
	"syscall"
	"time"
	"workshop/grpc/chat"
)

var username = flag.String("username", "default", "Username of person using the chat")
var chatServer = flag.String("server", ":8080", "Chat server address")
var debugEnabled = flag.Bool("debug", false, "Enabled/disable debug logging")

const defaultRoomName = "default"

func listenToMessages(ctx context.Context, client chat.ChatClient, room string) {
	joinRoomRequest := chat.JoinRoomRequest{Room: room, Username: *username}
	stream, err := client.ListenToRoom(ctx, &joinRoomRequest)

	if err != nil {
		log.Fatalf("listenToMessages() failed.")
	}

	for {
		select {
		case <-stream.Context().Done():
			return
		default:
			in, err := stream.Recv()
			if err == io.EOF {
				//wg.Done()
				return
			}
			if err != nil {
				//wg.Done()
				return
				//log.Fatalf("Failed to receive message from channel joining. \nErr: %v", err)
			}

			userString := color.GreenString(in.GetUsername())
			if *username == in.GetUsername() {
				userString = color.HiMagentaString(in.GetUsername())
			}

			fmt.Println(fmt.Sprintf("[%s >> %s]: (%s) -> %s",
				color.HiMagentaString(in.Time.AsTime().Format("2006-01-02T15:04:05 -07000")),
				color.MagentaString(in.GetRoom()),
				userString,
				color.CyanString(in.GetMessage())),
			)
		}
	}
}

func joinChannel(ctx context.Context, client chat.ChatClient, room string) error {
	joinRoomRequest := chat.JoinRoomRequest{Room: room, Username: *username}
	_, err := client.JoinRoom(ctx, &joinRoomRequest)

	return err
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

	err = joinChannel(ctx, client, room)
	exitOnError(err)
	pollRoomExists(client, ctx, room)

	go listenToMessages(ctx, client, room)
	go writeAndSendMessages(ctx, client, room)

	_, err = client.NotifyJoin(ctx, &chat.NotifyJoinMessage{
		Room:     room,
		Username: *username,
	})

	<-sigChan

	_, err = client.NotifyDisconnect(ctx, &chat.NotifyDisconnectRequest{
		Room:     room,
		Username: *username,
	})

	_, err = client.DisconnectFromRoom(ctx, &chat.DisconnectFromRoomMessage{
		Room:     room,
		Username: *username,
	})

	exitOnError(err)

	exitOnError(conn.Close())
}

func pollRoomExists(client chat.ChatClient, ctx context.Context, room string) {
	for {
		result, _ := client.CheckRoomExists(ctx, &chat.CheckRoomExistsMessage{Room: room})

		if result.Success {
			break
		}

		time.Sleep(250 * time.Millisecond)
	}
}

func writeAndSendMessages(ctx context.Context, client chat.ChatClient, room string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		go sendMessage(ctx, client, room, scanner.Text())
		fmt.Print("\033[1A\033[K")
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
