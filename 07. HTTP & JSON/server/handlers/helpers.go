package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type errorMsg struct {
	Message string `json:"Message"`
}

func writeOk(w http.ResponseWriter, jsonAsBytes []byte) (int, error) {
	w.WriteHeader(http.StatusOK)
	return w.Write(jsonAsBytes)
}

func writeMethodNotAllowed(w http.ResponseWriter, allowedMethods ...string) (int, error) {
	methodNotAllowedMsg := errorMsg{
		Message: fmt.Sprintf("Method is not allowed. Allowed methods: %s", strings.Join(allowedMethods, ",")),
	}

	bytes, err := json.Marshal(methodNotAllowedMsg)

	if err != nil {
		return writeInternalServerError(w)
	}

	w.WriteHeader(http.StatusMethodNotAllowed)

	return w.Write(bytes)
}

func writeInternalServerError(w http.ResponseWriter) (int, error) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	return w.Write([]byte{})
}
