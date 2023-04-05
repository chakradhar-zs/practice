package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Test_getAllBooks(t *testing.T) {
	var body []books
	testcases := []struct {
		method    string
		target    string
		expstatus int
		expbody   []books
	}{
		{http.MethodGet, "/", http.StatusOK, b},
		{http.MethodPost, "/", http.StatusMethodNotAllowed, body},
		{http.MethodPut, "/", http.StatusMethodNotAllowed, body},
		{http.MethodDelete, "/", http.StatusMethodNotAllowed, body},
	}

	for _, val := range testcases {
		var book books
		w := httptest.NewRecorder()
		r := httptest.NewRequest(val.method, val.target, nil)

		getAllBooks(w, r)
		err := json.Unmarshal(w.Body.Bytes(), &book)
		if err != nil {
			log.Println(err)
		}
		if w.Code != val.expstatus {
			t.Error("Unexpected Code")
		}

		if reflect.DeepEqual(book, val.expbody) {
			t.Error("Unexpected Content")
		}

	}
}
