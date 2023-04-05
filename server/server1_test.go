package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingHandler(t *testing.T) {
	r, err := http.NewRequest(http.MethodGet, "/ping", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	w := httptest.NewRecorder()

	pingHandler(w, r)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("http returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "pong"
	got := w.Body.String()

	assert.Equal(t, expected, got, "Expected response body did not match the body returned from the server")
}
