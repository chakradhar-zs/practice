package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_hello(t *testing.T) {
	testcases := []struct {
		method    string
		target    string
		expstatus int
		expbody   string
	}{
		{http.MethodGet, "/", http.StatusOK, "hello-world"},
		{http.MethodPost, "/", http.StatusMethodNotAllowed, ""},
		{http.MethodPut, "/", http.StatusMethodNotAllowed, ""},
		{http.MethodDelete, "/", http.StatusMethodNotAllowed, ""},
	}

	for _, val := range testcases {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(val.method, val.target, nil)

		hello(w, r)

		if w.Code != val.expstatus {
			t.Error("Unexpected Status Code")
		}
		if w.Body.String() != val.expbody {
			t.Error("Unexpected Content")
		}

	}
}
