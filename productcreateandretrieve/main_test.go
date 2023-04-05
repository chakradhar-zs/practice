package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"productcreateandretrieve/db"
	"reflect"
	"testing"
)

func TestCreateProduct(t *testing.T) {
	testcases := []struct {
		method    string
		body      db.Product
		brandname string
		expStatus int
		expBody   int64
	}{
		{http.MethodPost, db.Product{3, "sneaker shoes", "stylish", 1000, 3, "shoes", 4, "Available"}, "Nike", http.StatusOK, 1},
		{http.MethodPost, db.Product{4, "Rolex", "useful", 50000, 1, "wristwatch", 5, "Discontinued"}, "Titan", http.StatusOK, 1},
		{http.MethodPost, db.Product{5, "Bru", "tasty", 100, 3, "coffee", 6, "Available"}, "Bru", http.StatusOK, 1},
		{http.MethodGet, db.Product{6, "juice", "cool", 20, 7, "softdrinks", 7, "Available"}, "Maaza", http.StatusMethodNotAllowed, 0},
		{http.MethodPut, db.Product{}, "", http.StatusMethodNotAllowed, 0},
	}

	for _, val := range testcases {

		var d db.Data
		d.Pro = val.body
		d.Brand = val.brandname
		b, err := json.Marshal(d)
		if err != nil {
			t.Error("Error while marshalling")
		}
		w := httptest.NewRecorder()
		r := httptest.NewRequest(val.method, "/", bytes.NewBuffer(b))

		op := CreateProductHandler(w, r)

		if w.Code != val.expStatus {
			t.Errorf("Unexpected Response Code %v", w.Code)
		}

		if op != val.expBody {
			t.Errorf("Unexpected Response Body %v , want 1", op)
		}
	}

}

func Test_getByID(t *testing.T) {
	testcases := []struct {
		method    string
		target    string
		output    db.Product
		expStatus int
	}{
		{http.MethodGet, "/product/3", db.Product{3, "sneaker shoes", "stylish", 1000, 3, "shoes", 4, "Available"}, http.StatusOK},
		{http.MethodGet, "/product/4", db.Product{4, "Rolex", "useful", 50000, 1, "wristwatch", 5, "Discontinued"}, http.StatusOK},
		{http.MethodGet, "/product/1000", db.Product{}, http.StatusNotFound},
		{http.MethodPut, "/product/89", db.Product{}, http.StatusMethodNotAllowed},
		{http.MethodGet, "/product/abc", db.Product{}, http.StatusBadRequest},
	}

	for _, v := range testcases {

		var p db.Product
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
