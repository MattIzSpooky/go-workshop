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
	Alloc      string `json:"alloc"`
	TotalAlloc string `json:"total-alloc"`
	Sys        string `json:"sys"`
	NumGC      string `json:"num-gc"`
	Frees      string `json:"frees"`
}

func convertBytesToMiB(amountOfBytes uint64) uint64 {
	return amountOfBytes / 1024 / 1024
}

const allowedMethodForRuntimeInfoHandler = "GET"

func (h *Handler) RuntimeInfoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != allowedMethodForRuntimeInfoHandler {
		_, _ = writeMethodNotAllowed(w, allowedMethodForRuntimeInfoHandler)
		return
	}

	h.log.WriteInfo("Getting runtime info..")

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	runtimeInfo := RuntimeInfo{
		Alloc:      fmt.Sprintf("%v MiB", convertBytesToMiB(m.Alloc)),
		TotalAlloc: fmt.Sprintf("%v MiB", convertBytesToMiB(m.TotalAlloc)),
		Sys:        fmt.Sprintf("%v MiB", convertBytesToMiB(m.Sys)),
		NumGC:      strconv.Itoa(int(m.NumGC)),
		Frees:      strconv.FormatUint(m.Frees, 10),
	}

	response, err := json.Marshal(runtimeInfo)

	if err != nil {
		h.crash(w, err)
		return
	}

	_, err = writeOk(w, response)

	if err != nil {
		h.crash(w, err)
	}
}
