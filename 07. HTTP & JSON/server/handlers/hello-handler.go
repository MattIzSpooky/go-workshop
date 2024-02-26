package handlers

import (
	"encoding/json"
	"net/http"
)

// The hello world route.
const HelloWorldRoute = "hello-world"

// The struct we want to send back.
type Message struct {
	// The `json:"message"` part tells Go how a field should look when marshalled to JSON.
	// This also works for some other formats such as yaml.
	Message string `json:"message"`
}

const allowedMethodForHelloHandler = "GET"

func (h *Handler) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	// Return method not allowed when a non-allowed method is used.
	if r.Method != allowedMethodForHelloHandler {
		_, _ = h.writeMethodNotAllowed(w, allowedMethodForHelloHandler)
		return
	}
	h.log.WriteInfo("Sending hello!")

	// Instantiate an object and marshall it to Json.
	response, err := json.Marshal(Message{Message: "hello world!"})
	// Alternatively, one can use json.NewEncoder(w).Encode(Message{Message: "hello world!"}) but this method uses more memory

	if err != nil {
		h.writeInternalServerError(w, err)
		return
	}

	// Write the result back to the client.
	_, err = h.writeOk(w, response)

	if err != nil {
		h.writeInternalServerError(w, err)
	}
}
