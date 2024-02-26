package server

import (
	"github.com/rs/cors"
	"net/http"
	"workshop/http+json/server/handlers"
)

func (s *HTTPServer) registerHandlers() {
	// ServeMux is an HTTP request multiplexer. See ServeMux for more information.
	mux := &http.ServeMux{}

	// register handlers onto the serve mux.
	handler := handlers.NewHandler(s.log)
	handler.RegisterAll(mux)

	// Add cors https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS
	s.Handler = cors.Default().Handler(mux)
}
