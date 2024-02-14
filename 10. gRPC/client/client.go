package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"log"
	"os"
	"workshop/grpc/chat"
)

const defaultRoomName = "default"

var roomName = flag.String("room", defaultRoomName, "Chatroom name")
var username = flag.String("sender", "default", "Username of person using the chat")
var chatServer = flag.String("server", ":8080", "Chat server address")

func joinChannel(ctx context.Context, client chat.ChatClient) {

	room := chat.JoinRoomRequest{Room: *roomName, Username: *username}
	stream, err := client.JoinRoom(ctx, &room)
	if err != nil {
		log.Fatalf("client.JoinChannel(ctx, &channel) throws: %v", err)
	}

	fmt.Printf("Joined channel: %v \n", *roomName)

	// TODO: refactor to wait group, this works for now.
	waitc := make(chan struct{})

	// TODO: split this to another function
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive message from channel joining. \nErr: %v", err)
			}

			// TODO: [print room]: username -> message
			if *username != in.GetUsername() {
				fmt.Printf("MESSAGE: (%v) -> %v \n", in.GetUsername(), in.Message)
			}
		}
	}()

	<-waitc
}

func sendMessage(ctx context.Context, client chat.ChatClient, message string) {
	stream, err := client.SendMessage(ctx)
	if err != nil {
		log.Printf("Cannot send message: error: %v", err)
	}
	msg := chat.ChatMessage{
		Room:     *roomName,
		Username: *username,
		Message:  message,
	}
	stream.Send(&msg)

	ack, err := stream.CloseAndRecv()
	fmt.Printf("Message sent: %v \n", ack)
}

func main() {
	flag.Parse()

	fmt.Println("--- CLIENT APP ---")
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(*chatServer, opts...)
	if err != nil {
		log.Fatalf("Fail to dail: %v", err)
	}

	defer conn.Close()

	ctx := context.Background()
	client := chat.NewChatClient(conn)

	listRoomsReply, err := client.ListRooms(ctx, &emptypb.Empty{})

	fmt.Println(listRoomsReply)

	// TODO: ask if you want to join the default room
	// TODO: if not, print rooms currently on the server
	// TODO: ask if the user wants to join a specific room
	// TODO: if not, create a new room
	// TODO: when joining an existing room, print the users and ask again if the user wants to join.

	//if *roomName == defaultRoomName {
	//
	//}

	go joinChannel(ctx, client)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		go sendMessage(ctx, client, scanner.Text())
	}

}
