package main

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllProductsHandler(t *testing.T) {
	products = []Product{
		{Id: 1, Name: "Product 1", Description: "Description 1", Price: 1.99},
		{Id: 2, Name: "Product 2", Description: "Description 2", Price: 2.99},
		{Id: 3, Name: "Product 3", Description: "Description 3", Price: 3.99},
	}

	body, _ := json.Marshal(products)

	tests := []struct {
		desc           string
		expectedStatus int
		expectedBody   string
		method         string
	}{
		{
			desc:           "Success",
			expectedStatus: http.StatusOK,
			expectedBody:   string(body),
			method:         http.MethodGet,
		},
		{
			desc:           "Error case: invalid-method",
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "",
			method:         http.MethodPut,
		},
	}

	for i, tc := range tests {
		req := httptest.NewRequest(tc.method, "/products", nil)
		rec := httptest.NewRecorder()

		GetProducts(rec, req)

		assert.Equal(t, tc.expectedStatus, rec.Code, "TEST[%d], failed.\n%s", i, tc.desc)

		assert.Equal(t, tc.expectedBody, rec.Body.String(), "TEST[%d], failed.\n%s", i, tc.desc)
	}
}

func TestCreateProduct(t *testing.T) {
	successCase := Product{Id: 4, Name: "Product 1", Description: "This is product 1", Price: 10.99}
	invalidBody := map[string]interface{}{"id": "1"}

	tests := []struct {
		desc          string
		method        string
		reqBody       interface{}
		expStatuscode int
		expResp       Product
	}{
		{"success case", http.MethodPost, successCase, http.StatusCreated, successCase},
		{"error case: 405", http.MethodDelete, successCase, http.StatusMethodNotAllowed, Product{}},
		{"error case: 400", http.MethodPost, invalidBody, http.StatusBadRequest, Product{}},
	}

	for i, tc := range tests {
		body, _ := json.Marshal(tc.reqBody)

		r, err := http.NewRequest(tc.method, "/products", bytes.NewBuffer(body))
		if err != nil {
			t.Errorf("could not create request: %v", err)
			continue
		}

		w := httptest.NewRecorder()

		CreateProduct(w, r)

		assert.Equalf(t, tc.expStatuscode, w.Code, "status code mismatch")

		var p Product

		_ = json.Unmarshal(w.Body.Bytes(), &p)

		assert.Equalf(t, tc.expResp, p, "Test[%d]. Handler returned unexpected body: got %v want %v",
			i, p, tc.expResp)
	}
}
