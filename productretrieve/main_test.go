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

func Test_getByID(t *testing.T) {
	testcases := []struct {
		method    string
		target    string
		output    Product
		expStatus int
	}{
		{http.MethodGet, "/product/1", Product{1, "Amul Ghee", "pure", 200, 2, "ghee", 1, "Available"}, http.StatusOK},
		{http.MethodGet, "/product/2", Product{2, "Bag", "comfort", 3000, 1, "travel bags", 2, "Available"}, http.StatusOK},
		//{http.MethodGet, "/product/1000", Product{}, http.StatusNotFound},
		//{http.MethodPut, "/product/3", Product{}, http.StatusMethodNotAllowed},
		//{http.MethodGet, "/product/abc", Product{}, http.StatusBadRequest},
	}

	for _, v := range testcases {

		var p Product
		w := httptest.NewRecorder()
		r := httptest.NewRequest(v.method, v.target, nil)

		getByID(w, r)
		got, err := io.ReadAll(w.Body)
		if err != nil {
			t.Error("Invalid Body")
		}
		err = json.Unmarshal(got, &p)
		if err != nil {
			log.Fatal(err)
		}
		if w.Code != v.expStatus {
			t.Error("Unexpected Response Code")
		}
		if !reflect.DeepEqual(p, v.output) {
			t.Errorf("Retrieved Product Data is incorrect: %v, want %v", p, v.output)
		}

	}
}
