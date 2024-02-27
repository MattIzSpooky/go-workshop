package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"os/signal"
	"syscall"
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

	room, err := queryRoomFromAvailableRooms(grpcClient, ctx)
	exitOnError(err)

	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGKILL,
		syscall.SIGHUP,
	)

	chatroomObj := Chatroom{
		room:           room,
		username:       *username,
		ctx:            ctx,
		internalClient: grpcClient,
	}

	err = chatroomObj.JoinAndListen()
	exitOnError(err)

	chatroomObj.WriteAndSendMessages()

	<-sigChan

	err = chatroomObj.Disconnect()

	exitOnError(err)
	exitOnError(conn.Close())
}
