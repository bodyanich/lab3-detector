package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHealthHandler(t *testing.T) {
	version = "test-version"

	request := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	recorder := httptest.NewRecorder()

	healthHandler(recorder, request)

	require.Equal(t, http.StatusOK, recorder.Code)
	require.Equal(t, "application/json", recorder.Header().Get("Content-Type"))

	var response statusResponse
	err := json.NewDecoder(recorder.Body).Decode(&response)

	require.NoError(t, err)
	require.Equal(t, "ok", response.Status)
	require.Equal(t, "test-version", response.Version)
}
