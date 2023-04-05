package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Test_pingHandler(t *testing.T) {

	//body, _ := json.Marshal(Message{Value: "pong"})

	testcases := []struct {
		method       string
		target       string
		body         io.Reader
		expectStatus int
		expbody      string
	}{
		{http.MethodGet, "/ping", nil, http.StatusOK, "pong"},

		{http.MethodPost, "/ping", nil, http.StatusMethodNotAllowed, ""},
		{http.MethodPost, "/vc", nil, http.StatusNotFound, ""},
		{http.MethodPut, "/ping", nil, http.StatusMethodNotAllowed, ""},
		{http.MethodDelete, "/ping", nil, http.StatusMethodNotAllowed, ""},
	}

	for _, val := range testcases {
		var msg Message
		w := httptest.NewRecorder()
		r := httptest.NewRequest(val.method, val.target, val.body)
		pingHandler(w, r)

		err := json.Unmarshal(w.Body.Bytes(), &msg)
		if err != nil {
			log.Print(err)
		}
		if w.Code != val.expectStatus {
			t.Error("Unexpected Response code")
		}
		if !reflect.DeepEqual(msg.Value, val.expbody) {
			t.Error("Unexpected Content")
		}
	}
}
