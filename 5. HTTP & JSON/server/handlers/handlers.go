package handlers

import (
	"fmt"
	"net/http"
	"workshop/http+json/server/log"
	"workshop/http+json/server/middleware"
)

type Handler struct {
	log *log.ServerLogger
}

func NewHandler(l *log.ServerLogger) *Handler {
	return &Handler{log: l}
}

func (h *Handler) crash(w http.ResponseWriter, reason error) {
	h.log.WriteErrorObject(reason)
	w.WriteHeader(http.StatusInternalServerError)
}

const RootRoute = "/"

func (h *Handler) RegisterAll(mux *http.ServeMux) {
	helloWorld := applyMiddleware(middleware.LogMiddleware(http.HandlerFunc(h.HelloWorldHandler), h.log), middleware.JsonMiddleware)
	runtimeHandler := applyMiddleware(middleware.LogMiddleware(http.HandlerFunc(h.RuntimeInfoHandler), h.log), middleware.JsonMiddleware)

	h.registerHandler(h.createUrl(HelloWorldRoute), helloWorld, mux)
	h.registerHandler(h.createUrl(RuntimeInfoRoute), runtimeHandler, mux)
}

func applyMiddleware(handler http.HandlerFunc, middlewares ...func(http.Handler) http.HandlerFunc) http.HandlerFunc {
	for _, m := range middlewares {
		handler = m(handler)
	}

	return handler
}

func (h *Handler) registerHandler(route string, handler http.Handler, mux *http.ServeMux) {
	mux.Handle(route, handler)

	h.log.WriteInfo(fmt.Sprintf("[%s] has been registered.", route))
}

func (h *Handler) createUrl(parts ...string) string {
	url := RootRoute

	for _, part := range parts {
		url += part
	}

	return url
}
