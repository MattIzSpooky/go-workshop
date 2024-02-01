package main

import (
	"os"
	"os/signal"
	"syscall"
	"workshop/http+json/server"
)

func main() {
	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGKILL,
		syscall.SIGHUP,
	)

	httpServer := server.NewServer()

	go func() {
		if err := httpServer.Serve(); err != nil {
			panic(err)
		}
	}()

	<-sigChan

	if err := httpServer.GracefulShutdown(); err != nil {
		panic(err)
	}
}
