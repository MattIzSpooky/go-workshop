package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"workshop/grpc/generated/chat"
)

var debugEnabled = flag.Bool("debug", true, "Enabled/disable debug logging")
var port = flag.Int("port", 8080, "The port for the chat server-app to start listening on")

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	chat.RegisterChatServer(grpcServer, &chatServer{rooms: nil})

	log.Printf("server-app listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
