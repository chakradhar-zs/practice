package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRetrieveProduct(t *testing.T) {
	testcases := []struct {
		method    string
		target    string
		expOutput Product
		expStatus int
	}{
		{http.MethodGet, "/product/3", Product{3, "sneaker shoes", "stylish", 1000, 3, "shoes", 4, "Available"}, http.StatusOK},
		{http.MethodGet, "/product/4", Product{4, "Rolex", "useful", 50000, 1, "wristwatch", 5, "Discontinued"}, http.StatusOK},
	}

	for _, val := range testcases {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		var r1 Result
		RetrieveProduct(w, r)
		got, err := io.ReadAll(w.Body)
		if err != nil {
			log.Println("Error while reading")
		}
		err = json.Unmarshal(got, &r1)
		if err != nil {
			log.Println("Error while unmarshalling")
		}
		var expError error
		if !assert.IsType(t, expError, r1.val2) {
			t.Error("Unexpected Error")
		}
		if reflect.DeepEqual(r1.val1, val.expOutput) {
			t.Error("Unexpected Body")
		}
		if w.Code != val.expStatus {
			t.Error("Unexpected Code")
		}
	}
}
