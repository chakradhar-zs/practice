package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"developer.zopsmart.com/go/gofr/pkg/errors"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"developer.zopsmart.com/go/gofr/pkg/gofr/request"
)

func initTest(method, path string, body []byte) *gofr.Context {
	r := httptest.NewRequest(method, path, bytes.NewReader(body))
	req := request.NewHTTPRequest(r)
	ctx := gofr.NewContext(nil, req, gofr.New())

	return ctx
}

func Test_getByID(t *testing.T) {
	testcases := []struct {
		method string
		id     string
		desc   string
		output interface{}
		err    error
	}{
		{http.MethodGet, "3", "Success", Product{3, "sneaker shoes", "stylish", 1000, 3, "shoes", 4, "Available"}, nil},
		{http.MethodGet, "4", "Success", Product{4, "Rolex", "useful", 50000, 1, "wristwatch", 5, "Discontinued"}, nil},
		{http.MethodGet, "5", "Success", Product{5, "Bru", "tasty", 100, 3, "coffee", 6, "Available"}, nil},
		{http.MethodGet, "100", "Failure", Product{}, errors.EntityNotFound{}},
		{http.MethodGet, "abc", "Failure", Product{}, errors.InvalidParam{}},
	}

	for i, val := range testcases {
		ctx := initTest(val.method, "/product", nil)

		ctx.SetPathParams(map[string]string{"id": val.id})

		gotBody, gotError := getByID(ctx)

		assert.Equal(t, val.err, gotError, "Test[%d] failed. %s", i, val.desc)
		assert.Equal(t, val.output, gotBody, "Test[%d] failed. %s", i, val.desc)

	}
}

func TestCreateProduct(t *testing.T) {
	testcases := []struct {
		method string
		input  Product
		desc   string
		output int64
		err    error
	}{
		{http.MethodPost, Product{6, "maggi", "tasty", 50, 3, "noodles", 1, "Available"}, "Success", 1, nil},
		{http.MethodPost, Product{7, "Dairy Milk", "sweet", 100, 8, "dairymilk", 2, "Available"}, "Success", 1, nil},
		{http.MethodPost, Product{}, "Failure", 0, errors.MissingParam{}},
	}
	for i, val := range testcases {

		body, err := json.Marshal(val.input)
		if err != nil {
			log.Println("Error while marshalling")
		}
		ctx := initTest(val.method, "/product", body)

		row, gotError := CreateProduct(ctx)

		assert.Equal(t, val.err, gotError, "Test[%d] failed. %s", i, val.desc)
		assert.Equal(t, val.output, row, "Test[%d] failed. %s", i, val.desc)

	}
}
