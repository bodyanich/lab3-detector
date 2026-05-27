// Package main starts the image metadata processor service.
package main

import (
	"encoding/json"
	"net/http"
)

type statusResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	writeStatusResponse(w, "ok")
}

func readyHandler(w http.ResponseWriter, _ *http.Request) {
	writeStatusResponse(w, "ready")
}

func writeStatusResponse(w http.ResponseWriter, status string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := statusResponse{
		Status:  status,
		Version: version,
	}

	_ = json.NewEncoder(w).Encode(response)
}
