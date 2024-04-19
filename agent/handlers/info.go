package handlers

import (
	"encoding/json"
	"github.com/pbnjay/memory"
	"github.com/uptrace/bunrouter"
	"net/http"
	"runtime"
)

type InfoResponse struct {
	Id          string `json:"id"`
	Os          string `json:"os"`
	Arch        string `json:"arch"`
	NumCores    int    `json:"num_cores"`
	MemoryBytes uint64 `json:"memory_bytes"`
}

func Info(w http.ResponseWriter, r bunrouter.Request) error {
	response := &InfoResponse{
		Os:          runtime.GOOS,
		Arch:        runtime.GOARCH,
		NumCores:    runtime.NumCPU(),
		MemoryBytes: memory.TotalMemory(),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	responseBytes, err := json.Marshal(response)
	if err != nil {
		return err
	}
	_, err = w.Write(responseBytes)
	return err
}
