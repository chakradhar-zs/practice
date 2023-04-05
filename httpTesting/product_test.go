package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProductHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	ProductHandler(w, r)

	response := w.Result()
	if response.StatusCode != http.StatusOK {
		t.Error("Unexpected Response code")
	}

	if response.Header.Get("content-type") != "application/json" {
		t.Error("Unexpected Content Type")
	}
}

func TestHttpClient(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	code, err := HttpClient(s.URL)
	if err != nil {
		t.Error(err)
	}
	if code != http.StatusOK {
		t.Error("Unexpected Status Code")
	}
}
