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

type HTTPServer struct {
	http.Server
	log         *log.ServerLogger
	fullAddress string
}

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

	server := &HTTPServer{
		Server: http.Server{
			Addr:         fullAddress,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		},
		log:         log.NewLogger(),
		fullAddress: fullAddress,
	}

	server.registerHandlers()

	return server
}

func (s *HTTPServer) Serve() error {
	const runStringTemplate = "Running HTTP server on: http://%s"

	s.log.WriteInfo(fmt.Sprintf(runStringTemplate, s.fullAddress))

	if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		s.log.WriteErrorString(fmt.Sprintf("Could not listen on %s: %v\n", s.Addr, err.Error()))

		return err
	}

	return nil
}
func (s *HTTPServer) GracefulShutdown() error {
	s.log.WriteInfo("Server is shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	s.SetKeepAlivesEnabled(false)

	if err := s.Shutdown(ctx); err != nil {
		s.log.WriteErrorString(fmt.Sprintf("Could not gracefully shutdown the server: %v\n", err.Error()))
		return err
	}

	s.log.WriteInfo("Server has successfully closed down")

	return nil
}

func (s *HTTPServer) writeError(w http.ResponseWriter, text string, statusCode int) {
	w.WriteHeader(statusCode)
	_, err := w.Write([]byte(fmt.Sprintf(`{ "error": "%s" }`, text)))

	if err != nil {
		s.log.WriteErrorString(fmt.Sprintf("Could not write error: %s", err.Error()))
	}
}

func (s *HTTPServer) Crash(w http.ResponseWriter, reason error) {
	s.log.WriteErrorObject(reason)
	w.WriteHeader(http.StatusInternalServerError)
}
