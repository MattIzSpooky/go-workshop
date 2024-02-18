package server

import (
	"github.com/rs/cors"
	"net/http"
	"workshop/http+json/server/handlers"
)

func (s *HTTPServer) registerHandlers() {
	mux := &http.ServeMux{}

	handler := handlers.NewHandler(s.log)
	handler.RegisterAll(mux)

	s.Handler = cors.Default().Handler(mux)
}
