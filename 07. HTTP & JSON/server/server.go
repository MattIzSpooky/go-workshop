package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"
	"workshop/http+json/server/log"
)

// Create our own http server object based on the standard library one.
type HTTPServer struct {
	http.Server
	log         *log.ServerLogger
	fullAddress string
}

// Initiates an HTTP server.
func NewServer() *HTTPServer {
	var fullAddress string
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	if strings.Contains(runtime.GOOS, "windows") || strings.Contains(runtime.GOOS, "darwin") {
		fullAddress = fmt.Sprintf("%s:%s", "localhost", port)
	} else {
		fullAddress = fmt.Sprintf(":%s", port)
	}

	// Initiate the http server
	server := &HTTPServer{
		Server: http.Server{
			Addr:         fullAddress,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		},
		log:         log.NewLogger(), // we use our own custom logger.
		fullAddress: fullAddress,
	}

	// Register the HTTP handlers.
	server.registerHandlers()

	return server
}

// Listen to requests.
func (s *HTTPServer) Serve() error {
	const runStringTemplate = "Running HTTP server on: http://%s"

	s.log.WriteInfo(fmt.Sprintf(runStringTemplate, s.fullAddress))

	// This here does the magic, this exposes the handlers so that they can be called through HTTP.
	if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		s.log.WriteErrorString(fmt.Sprintf("Could not listen on %s: %v\n", s.Addr, err.Error()))

		return err
	}

	return nil
}
func (s *HTTPServer) GracefulShutdown() error {
	s.log.WriteInfo("Server is shutting down")

	// Create a deadline for exiting the server. 15 seconds.
	// This is for shutting down the server.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	s.SetKeepAlivesEnabled(false)

	// s.Shutdown will try to gracefully clean up and handle the last requests.
	if err := s.Shutdown(ctx); err != nil {
		s.log.WriteErrorString(fmt.Sprintf("Could not gracefully shutdown the server: %v\n", err.Error()))
		return err
	}

	s.log.WriteInfo("Server has successfully closed down")

	return nil
}
