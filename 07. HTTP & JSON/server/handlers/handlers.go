package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"workshop/http+json/server/log"
	"workshop/http+json/server/middleware"
)

// Our handler object.
type Handler struct {
	log *log.ServerLogger
}

func NewHandler(l *log.ServerLogger) *Handler {
	return &Handler{log: l}
}

// The root route.
const RootRoute = "/"

func (h *Handler) RegisterAll(mux *http.ServeMux) {
	// Instantiate some handlers with the logging & json middleware.
	helloWorld := applyMiddleware(middleware.LogMiddleware(http.HandlerFunc(h.HelloWorldHandler), h.log), middleware.JsonMiddleware)
	runtimeHandler := applyMiddleware(middleware.LogMiddleware(http.HandlerFunc(h.RuntimeInfoHandler), h.log), middleware.JsonMiddleware)

	// Register each handler
	h.registerHandler(h.createUrl(HelloWorldRoute), helloWorld, mux)
	h.registerHandler(h.createUrl(RuntimeInfoRoute), runtimeHandler, mux)
}

// Wraps a given handler with one or more middlewares.
func applyMiddleware(handler http.HandlerFunc, middlewares ...func(http.Handler) http.HandlerFunc) http.HandlerFunc {
	for _, m := range middlewares {
		handler = m(handler)
	}

	return handler
}

// Adds the handler to the serve mux.
func (h *Handler) registerHandler(route string, handler http.Handler, mux *http.ServeMux) {
	mux.Handle(route, handler)

	h.log.WriteInfo(fmt.Sprintf("[%s] has been registered.", route))
}

// A util function to create the url.
func (h *Handler) createUrl(parts ...string) string {
	url := RootRoute

	for _, part := range parts {
		url += part
	}

	return url
}

type errorMsg struct {
	Message string `json:"Message"`
}

// A helper function to return OK responses in the handlers.
func (h *Handler) writeOk(w http.ResponseWriter, jsonAsBytes []byte) (int, error) {
	w.WriteHeader(http.StatusOK)
	return w.Write(jsonAsBytes)
}

// A helper function to return method not allowed responses in the handlers.
func (h *Handler) writeMethodNotAllowed(w http.ResponseWriter, allowedMethods ...string) (int, error) {
	methodNotAllowedMsg := errorMsg{
		Message: fmt.Sprintf("Method is not allowed. Allowed methods: %s", strings.Join(allowedMethods, ",")),
	}

	bytes, err := json.Marshal(methodNotAllowedMsg)

	if err != nil {
		return h.writeInternalServerError(w, err)
	}

	w.WriteHeader(http.StatusMethodNotAllowed)

	return w.Write(bytes)
}

// A helper function to return internal server error responses in the handlers.
func (h *Handler) writeInternalServerError(w http.ResponseWriter, reason error) (int, error) {
	h.log.WriteErrorObject(reason)
	w.WriteHeader(http.StatusInternalServerError)
	return w.Write([]byte{})
}
