package handlers

import (
	"net/http"
)

func writeOk(w http.ResponseWriter, jsonAsBytes []byte) (int, error) {
	w.WriteHeader(http.StatusOK)
	return w.Write(jsonAsBytes)
}
