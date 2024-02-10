package main

import (
	"context"
	"fmt"
	"github.com/Azure/go-amqp"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx := context.TODO()
	addr := "amqp://artemis:artemis@localhost:61616/"

	// create connection
	conn, err := amqp.Dial(context.TODO(), addr, nil)
	if err != nil {
		log.Fatal("Dialing AMQP server:", err)
	}
	defer conn.Close()

	// open a session
	session, err := conn.NewSession(context.TODO(), nil)
	if err != nil {
		log.Fatal("Creating AMQP session:", err)
	}

	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGKILL,
		syscall.SIGHUP,
	)

	go runQueueExample(ctx, session, addr)
	go runTopicExample(ctx, session, addr)

	<-sigChan

	session.Close(ctx)
}

func runQueueExample(ctx context.Context, session *amqp.Session, addr string) {
	// send a message

	// create a sender
	sender, err := session.NewSender(context.TODO(), "/queue-name", &amqp.SenderOptions{
		TargetCapabilities: []string{"queue"},
	})
	if err != nil {
		log.Fatal("Creating sender link:", err)
	}

	go func() {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

		for i := 0; i < 1000; i++ {
			str := fmt.Sprintf("Hello, world %d!", i)
			// send message
			err = sender.Send(context.TODO(), &amqp.Message{
				Data: [][]byte{[]byte(str)},
				Properties: &amqp.MessageProperties{
					To: &addr,
				},
			}, nil)
			if err != nil {
				log.Fatal("Sending message:", err)
			}
		}

		sender.Close(ctx)
		cancel()
	}()

	// continuously read messages
	// create a receiver
	receiver, err := session.NewReceiver(ctx, "/queue-name", &amqp.ReceiverOptions{
		SourceCapabilities: []string{"queue"},
	})
	if err != nil {
		log.Fatal("Creating receiver link:", err)
	}
	defer func() {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		receiver.Close(ctx)
		cancel()
	}()

	for {
		msg, err := receiver.Receive(context.TODO(), nil)
		if msg != nil {
			fmt.Println("Received from queue:", string(msg.GetData()))
			receiver.AcceptMessage(context.TODO(), msg)
		} else {
			fmt.Println("No More Messages")
			break
		}
		if err != nil {
			fmt.Println("Error")
		}
	}
}

func runTopicExample(ctx context.Context, session *amqp.Session, addr string) {
	// send a message
	// create a sender
	sender, err := session.NewSender(context.TODO(), "/topic-name", &amqp.SenderOptions{
		TargetCapabilities: []string{"topic"},
	})
	if err != nil {
		log.Fatal("Creating sender link:", err)
	}

	go func() {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

		for i := 0; i < 1000; i++ {
			str := fmt.Sprintf("Hello, world %d!", i)
			// send message
			err = sender.Send(context.TODO(), &amqp.Message{
				Data: [][]byte{[]byte(str)},
				Properties: &amqp.MessageProperties{
					To: &addr,
				},
			}, nil)
			if err != nil {
				log.Fatal("Sending message:", err)
			}
		}

		sender.Close(ctx)
		cancel()
	}()

	// continuously read messages
	// create a receiver
	receiver, err := session.NewReceiver(ctx, "/topic-name", &amqp.ReceiverOptions{
		SourceCapabilities: []string{"topic"},
	})
	if err != nil {
		log.Fatal("Creating receiver link:", err)
	}
	defer func() {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		receiver.Close(ctx)
		cancel()
	}()

	for {
		msg, err := receiver.Receive(context.TODO(), nil)
		if msg != nil {
			fmt.Println("Received from topic:", string(msg.GetData()))
			receiver.AcceptMessage(context.TODO(), msg)
		} else {
			fmt.Println("No More Messages")
			break
		}
		if err != nil {
			fmt.Println("Error")
		}
	}
}
