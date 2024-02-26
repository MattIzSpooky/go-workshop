package client

import (
	"bufio"
	"context"
	"fmt"
	"github.com/fatih/color"
	"io"
	"log"
	"os"
	"time"
	"workshop/grpc/generated/chat"
)

type ChatClient struct {
	room           string
	username       string
	ctx            context.Context
	internalClient chat.ChatClient
}

func NewChatClient(room string, username string, ctx context.Context, internalClient chat.ChatClient) ChatClient {
	return ChatClient{
		room:           room,
		username:       username,
		ctx:            ctx,
		internalClient: internalClient,
	}
}

func (c ChatClient) JoinAndListen() error {
	err := c.joinRoom()

	if err != nil {
		return err
	}

	c.pollRoomExists()

	go c.listenToMessages()

	_, err = c.internalClient.NotifyJoin(c.ctx, &chat.NotifyJoinMessage{
		Room:     c.room,
		Username: c.username,
	})

	return err
}

func (c ChatClient) listenToMessages() {
	joinRoomRequest := chat.JoinRoomRequest{Room: c.room, Username: c.username}
	stream, err := c.internalClient.ListenToRoom(c.ctx, &joinRoomRequest)

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
			if c.username == in.GetUsername() {
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

func (c ChatClient) Disconnect() error {
	_, err := c.internalClient.NotifyDisconnect(c.ctx, &chat.NotifyDisconnectRequest{
		Room:     c.room,
		Username: c.username,
	})

	if err != nil {
		return err
	}

	_, err = c.internalClient.DisconnectFromRoom(c.ctx, &chat.DisconnectFromRoomMessage{
		Room:     c.room,
		Username: c.username,
	})

	return err
}

func (c ChatClient) WriteAndSendMessages() {
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			go c.sendMessage(scanner.Text())
			fmt.Print("\033[1A\033[K")
		}
	}()
}

func (c ChatClient) pollRoomExists() {
	for {
		result, _ := c.internalClient.CheckRoomExists(c.ctx, &chat.CheckRoomExistsMessage{Room: c.room})

		if result.Success {
			break
		}

		time.Sleep(250 * time.Millisecond)
	}
}

func (c ChatClient) joinRoom() error {
	joinRoomRequest := chat.JoinRoomRequest{Room: c.room, Username: c.username}
	_, err := c.internalClient.JoinRoom(c.ctx, &joinRoomRequest)

	return err
}

func (c ChatClient) sendMessage(message string) {
	stream, err := c.internalClient.SendMessage(c.ctx)
	if err != nil {
		log.Printf("Cannot send message: error: %v", err)
	}
	msg := chat.ChatMessage{
		Room:     c.room,
		Username: c.username,
		Message:  message,
	}
	stream.Send(&msg)

	_, err = stream.CloseAndRecv()
}
