package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductHandler(t *testing.T) {
	tests := []struct {
		desc          string
		method        string
		expStatuscode int
	}{
		{"success case", http.MethodGet, http.StatusOK},
		{"error case: 405", http.MethodDelete, http.StatusMethodNotAllowed},
	}

	for _, tc := range tests {
		r, err := http.NewRequest(tc.method, "/products", nil)
		if err != nil {
			t.Errorf("could not create request: %v", err)
			continue
		}

		w := httptest.NewRecorder()

		ProductHandler(w, r)

		assert.Equalf(t, tc.expStatuscode, w.Code, "status code mismatch")

		if w.Code == http.StatusOK {
			var p []Product

			err := json.Unmarshal(w.Body.Bytes(), &p)
			if err != nil {
				t.Errorf("invalid format of body")
			}
		}
	}
}
