package main

import (
	"os"
	"os/signal"
	"syscall"
	"workshop/http+json/server"
)

func main() {
	// Create a channel that listens to 1 os signal.
	sigChan := make(chan os.Signal, 1)

	// Listen to various os signals related to exiting an application (clean or unclean exit).
	signal.Notify(sigChan,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGKILL,
		syscall.SIGHUP,
	)

	httpServer := server.NewServer()

	// Run the HTTP server in a separate go routine and have the main thread wait for the application to exit
	// So that it can clean up etc.

	go func() {
		if err := httpServer.Serve(); err != nil {
			panic(err)
		}
	}()

	// Block until an os signal has been received. This keeps the server to run until an os signal is received.
	<-sigChan

	// Clean up before exiting. Handling last requests, etc.
	if err := httpServer.GracefulShutdown(); err != nil {
		panic(err)
	}
}
