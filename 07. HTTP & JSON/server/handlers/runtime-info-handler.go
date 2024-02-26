package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"strconv"
)

const RuntimeInfoRoute = "runtime-info"

type RuntimeInfo struct {
	// The `json:"whatever"` part tells Go how a field should look when marshalled to JSON.
	// This also works for some other formats such as yaml.
	Alloc      string `json:"alloc"`
	TotalAlloc string `json:"total-alloc"`
	Sys        string `json:"sys"`
	NumGC      string `json:"num-gc"`
	Frees      string `json:"frees"`
}

// A helper function to convert bytes to MiB.
func convertBytesToMiB(amountOfBytes uint64) uint64 {
	return amountOfBytes / 1024 / 1024
}

const allowedMethodForRuntimeInfoHandler = "GET"

func (h *Handler) RuntimeInfoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != allowedMethodForRuntimeInfoHandler {
		_, _ = h.writeMethodNotAllowed(w, allowedMethodForRuntimeInfoHandler)
		return
	}

	h.log.WriteInfo("Getting runtime info..")

	// Read some runtime stats.
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	// Write those values to our custom runtime information object
	runtimeInfo := RuntimeInfo{
		Alloc:      fmt.Sprintf("%v MiB", convertBytesToMiB(m.Alloc)),
		TotalAlloc: fmt.Sprintf("%v MiB", convertBytesToMiB(m.TotalAlloc)),
		Sys:        fmt.Sprintf("%v MiB", convertBytesToMiB(m.Sys)),
		NumGC:      strconv.Itoa(int(m.NumGC)),
		Frees:      strconv.FormatUint(m.Frees, 10),
	}

	// And send it back to the client.
	response, err := json.Marshal(runtimeInfo)

	if err != nil {
		h.writeInternalServerError(w, err)
		return
	}

	_, err = h.writeOk(w, response)

	if err != nil {
		h.writeInternalServerError(w, err)
	}
}
