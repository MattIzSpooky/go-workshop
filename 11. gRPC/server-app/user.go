package main

import "workshop/grpc/generated/chat"

type user struct {
	name       string
	msgChannel chan *chat.ChatMessage
}
