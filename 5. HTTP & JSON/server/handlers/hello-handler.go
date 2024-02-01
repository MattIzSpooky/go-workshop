package handlers

import (
	"encoding/json"
	"net/http"
)

const HelloWorldRoute = "hello-world"

type Message struct {
	Message string `json:"message"`
}

func (h *Handler) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	h.log.WriteInfo("Sending hello!")

	response, err := json.Marshal(Message{Message: "hello world!"})
	// Alternatively, one can use json.NewEncoder(w).Encode(Message{Message: "hello world!"}) but this method uses more memory

	if err != nil {
		h.crash(w, err)
		return
	}

	_, err = writeOk(w, response)

	if err != nil {
		h.crash(w, err)
	}
}
