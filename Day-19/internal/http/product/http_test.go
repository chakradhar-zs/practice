package product

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"developer.zopsmart.com/go/gofr/pkg/gofr/request"
	"developer.zopsmart.com/go/gofr/pkg/gofr/responder"

	"github.com/bmizerany/assert"
	"github.com/golang/mock/gomock"

	"Day-19/internal/constants"
	"Day-19/internal/models"
	"Day-19/internal/service"
)

func initTest(path string) *gofr.Context {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, path, nil)

	req := request.NewHTTPRequest(r)
	res := responder.NewContextualResponder(w, r)
	ctx := gofr.NewContext(res, req, gofr.New())

	return ctx
}

func TestRead(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	serviceMock := service.NewMockProduct(ctrl)

	tests := []struct {
		desc   string
		path   string
		input  string
		input1 string
		output interface{}
		expErr error
		calls  []*gomock.Call
	}{
		{desc: "Success",
			path:   "/product/?brand=false",
			input:  "3",
			input1: "false",
			output: models.Product{
				ID:          3,
				Name:        "sneaker shoes",
				Description: "stylish",
				Price:       1000,
				Quantity:    3,
				Category:    "shoes",
				Brand:       models.Brand{ID: 4, Name: ""},
				Status:      "Available",
			},
			expErr: nil,
			calls: []*gomock.Call{
				serviceMock.EXPECT().
					GetProduct(gomock.AssignableToTypeOf(&gofr.Context{}), 3, "false").
					Return(models.Product{
						ID: 3, Name: "sneaker shoes", Description: "stylish", Price: 1000, Quantity: 3, Category: "shoes",
						Brand: models.Brand{ID: 4, Name: ""}, Status: "Available"}, nil),
			}},
		{desc: "Fail",
			path:   "/product/?brand=true",
			input:  "abc",
			input1: "true",
			output: nil,
			expErr: errors.InvalidParam{Param: []string{"id"}},
			calls:  nil,
		},
		{desc: "Fail",
			path:   "/product/?brand=false",
			input:  "",
			input1: "",
			output: nil,
			expErr: errors.MissingParam{Param: []string{"id"}},
			calls:  nil,
		},
	}

	for i, val := range tests {
		ctx := initTest(val.path)
		ctx.SetPathParams(map[string]string{
			"id": val.input,
		})

		h := New(serviceMock)
		out, err := h.Read(ctx)
		assert.Equalf(t, val.output, out, "TEST[%d], failed.\n%s", i, val.desc)
		assert.Equalf(t, val.expErr, err, "TEST[%d], failed.\n%s", i, val.desc)
	}
}
func TestReadPro(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	serviceMock := service.NewMockProduct(ctrl)

	tests := []struct {
		desc   string
		path   string
		input  string
		input1 string
		output interface{}
		expErr error
		calls  []*gomock.Call
	}{
		{desc: "Success",
			path:   "/product/?brand=true",
			input:  "4",
			input1: "true",
			output: models.Product{
				ID:          4,
				Name:        "maggi",
				Description: "yum",
				Price:       100,
				Quantity:    3,
				Category:    "noodles",
				Brand:       models.Brand{ID: 5, Name: "Maggi"},
				Status:      "Available",
			},
			expErr: nil,
			calls: []*gomock.Call{
				serviceMock.EXPECT().GetProduct(gomock.AssignableToTypeOf(&gofr.Context{}), 4, "true").
					Return(models.Product{
						ID: 4, Name: "maggi", Description: "yum", Price: 100, Quantity: 3, Category: "noodles",
						Brand: models.Brand{ID: 5, Name: "Maggi"}, Status: "Available"}, nil),
			}},
	}

	for i, val := range tests {
		ctx := initTest(val.path)
		ctx.SetPathParams(map[string]string{
			"id": val.input,
		})

		h := New(serviceMock)
		out, err := h.Read(ctx)
		assert.Equalf(t, val.output, out, "TEST[%d], failed.\n%s", i, val.desc)
		assert.Equalf(t, val.expErr, err, "TEST[%d], failed.\n%s", i, val.desc)
	}
}

func TestWrite(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	serviceMock := service.NewMockProduct(ctrl)

	product1 := models.Product{
		ID: 6, Name: "zs_maggi", Description: "tasty", Price: 50, Quantity: 3, Category: "noodles",
		Brand: models.Brand{ID: 1, Name: ""}, Status: "Available"}

	tests := []struct {
		desc   string
		org    string
		input  interface{}
		output interface{}
		expErr error
		calls  []*gomock.Call
	}{
		{desc: "Success",
			org: "zs",
			input: models.Product{
				ID:          6,
				Name:        "maggi",
				Description: "tasty",
				Price:       50,
				Quantity:    3,
				Category:    "noodles",
				Brand:       models.Brand{ID: 1, Name: ""},
				Status:      "Available",
			},
			output: 1,
			expErr: nil,
			calls: []*gomock.Call{
				serviceMock.EXPECT().
					CreateProduct(gomock.AssignableToTypeOf(&gofr.Context{}), &product1).
					Return(1, nil),
			}},
		{desc: "Fail",
			org:    "",
			input:  models.Product{},
			output: 0,
			expErr: errors.MissingParam{Param: []string{"body"}},
			calls: []*gomock.Call{
				serviceMock.EXPECT().CreateProduct(gomock.AssignableToTypeOf(&gofr.Context{}), &models.Product{}).
					Return(0, errors.MissingParam{Param: []string{"body"}}),
			}},
	}

	for i, val := range tests {
		body, _ := json.Marshal(val.input)
		ctx := context.WithValue(context.Background(), constants.CtxValue, val.org)
		r := httptest.NewRequest(http.MethodGet, "/", bytes.NewReader(body))
		r = r.WithContext(ctx)
		req := request.NewHTTPRequest(r)
		c := gofr.NewContext(nil, req, app)
		c.Context = r.Context()

		h := New(serviceMock)
		out, err := h.Create(c)
		fmt.Println(err)
		assert.Equalf(t, val.output, out, "TEST[%d], failed.\n%s", i, val.desc)
		assert.Equalf(t, val.expErr, err, "TEST[%d], failed.\n%s", i, val.desc)
	}
}

func TestUpdate(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	serviceMock := service.NewMockProduct(ctrl)
	tests := []struct {
		desc   string
		input1 string
		input2 interface{}
		output interface{}
		expErr error
		calls  []*gomock.Call
	}{
		{desc: "Success",
			input1: "6",
			input2: models.Product{
				ID:          6,
				Name:        "Maggi",
				Description: "yummy",
				Price:       50,
				Quantity:    3,
				Category:    "noodles",
				Brand:       models.Brand{ID: 1, Name: ""},
				Status:      "Available"},
			output: 1,
			expErr: nil,
			calls: []*gomock.Call{
				serviceMock.EXPECT().
					UpdateProduct(gomock.AssignableToTypeOf(&gofr.Context{}), 6, &models.Product{
						ID: 6, Name: "Maggi", Description: "yummy", Price: 50, Quantity: 3, Category: "noodles",
						Brand: models.Brand{ID: 1, Name: ""}, Status: "Available"}).
					Return(1, nil),
			}},
		{desc: "Fail",
			input1: "2",
			input2: models.Product{},
			output: 0,
			expErr: errors.MissingParam{Param: []string{"body"}},
			calls: []*gomock.Call{
				serviceMock.EXPECT().UpdateProduct(gomock.AssignableToTypeOf(&gofr.Context{}), 2, &models.Product{}).
					Return(0, errors.MissingParam{Param: []string{"body"}}),
			}},
		{desc: "Fail",
			input1: "abc",
			input2: models.Product{},
			output: 0,
			expErr: errors.InvalidParam{Param: []string{"id"}},
			calls:  nil,
		},
		{desc: "Fail",
			input1: "",
			input2: models.Product{},
			output: 0,
			expErr: errors.MissingParam{Param: []string{"id"}},
			calls:  nil,
		},
	}

	for i, val := range tests {
		body, _ := json.Marshal(val.input2)
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", bytes.NewReader(body))
		req := request.NewHTTPRequest(r)
		res := responder.NewContextualResponder(w, r)
		ctx := gofr.NewContext(res, req, app)
		ctx.SetPathParams(map[string]string{
			"id": val.input1,
		})

		h := New(serviceMock)
		out, err := h.Update(ctx)
		assert.Equalf(t, val.output, out, "TEST[%d], failed.\n%s", i, val.desc)
		assert.Equalf(t, val.expErr, err, "TEST[%d], failed.\n%s", i, val.desc)
	}
}

func TestDelete(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	serviceMock := service.NewMockProduct(ctrl)
	tests := []struct {
		desc   string
		input  string
		output interface{}
		expErr error
		calls  []*gomock.Call
	}{
		{desc: "Success",
			input:  "3",
			output: 1,
			expErr: nil,
			calls: []*gomock.Call{
				serviceMock.EXPECT().DeleteProduct(gomock.AssignableToTypeOf(&gofr.Context{}), 3).
					Return(1, nil),
			}},
		{desc: "Fail",
			input:  "333",
			output: 0,
			expErr: errors.EntityNotFound{Entity: "id"},
			calls: []*gomock.Call{
				serviceMock.EXPECT().DeleteProduct(gomock.AssignableToTypeOf(&gofr.Context{}), 333).
					Return(0, errors.EntityNotFound{Entity: "id"}),
			}},
		{desc: "Fail",
			input:  "abc",
			output: 0,
			expErr: errors.InvalidParam{Param: []string{"id"}},
			calls:  nil,
		},
		{desc: "Fail",
			input:  "",
			output: 0,
			expErr: errors.MissingParam{Param: []string{"id"}},
			calls:  nil,
		},
	}

	for i, val := range tests {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		req := request.NewHTTPRequest(r)
		res := responder.NewContextualResponder(w, r)
		ctx := gofr.NewContext(res, req, app)
		ctx.SetPathParams(map[string]string{
			"id": val.input,
		})

		h := New(serviceMock)
		out, err := h.Delete(ctx)
		assert.Equalf(t, val.output, out, "TEST[%d], failed.\n%s", i, val.desc)
		assert.Equalf(t, val.expErr, err, "TEST[%d], failed.\n%s", i, val.desc)
	}
}

func TestIndex(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	serviceMock := service.NewMockProduct(ctrl)
	tests := []struct {
		desc   string
		org    string
		name   string
		path   string
		output interface{}
		expErr error
		calls  []*gomock.Call
	}{
		{desc: "Success",
			org:  "zs",
			name: "sneaker shoes",
			path: "/product/?brand=true",
			output: []models.Product{{
				ID:          3,
				Name:        "sneaker shoes",
				Description: "stylish",
				Price:       1000,
				Quantity:    3,
				Category:    "shoes",
				Brand:       models.Brand{ID: 4, Name: "Nike"},
				Status:      "Available"}},
			expErr: nil,
			calls: []*gomock.Call{
				serviceMock.EXPECT().GetProductByNAme(gomock.AssignableToTypeOf(&gofr.Context{}), "zs_sneaker shoes", "true").
					Return([]models.Product{{
						ID: 3, Name: "zs_sneaker shoes", Description: "stylish", Price: 1000, Quantity: 3, Category: "shoes",
						Brand: models.Brand{ID: 4, Name: "Nike"}, Status: "Available"}}, nil),
			}},
		{desc: "Fail",
			org:    "",
			name:   "",
			path:   "/product/?brand=true",
			output: nil,
			expErr: errors.EntityNotFound{},
			calls: []*gomock.Call{
				serviceMock.EXPECT().GetAllProducts(gomock.AssignableToTypeOf(&gofr.Context{}), "true").
					Return(nil, errors.EntityNotFound{}),
			}},
		{desc: "Success",
			org:  "",
			name: "",
			path: "/product/?brand=true",
			output: []models.Product{{
				ID: 6, Name: "zs_coffee", Description: "tasty", Price: 100, Quantity: 5, Category: "bru",
				Brand: models.Brand{ID: 5, Name: "bru"}, Status: "Available",
			}},
			expErr: nil,
			calls: []*gomock.Call{
				serviceMock.EXPECT().GetAllProducts(gomock.AssignableToTypeOf(&gofr.Context{}), "true").
					Return([]models.Product{{
						ID: 6, Name: "zs_coffee", Description: "tasty", Price: 100, Quantity: 5, Category: "bru",
						Brand: models.Brand{ID: 5, Name: "bru"}, Status: "Available",
					}}, nil),
			}},
	}

	for i, val := range tests {
		path := val.path + "&organization=" + val.org + "&name=" + val.name
		r, _ := http.NewRequest(http.MethodGet, path, http.NoBody)
		ctx := context.WithValue(context.Background(), constants.CtxValue, val.org)
		r = r.WithContext(ctx)
		req := request.NewHTTPRequest(r)
		c := gofr.NewContext(nil, req, app)

		h := New(serviceMock)
		out, err := h.Index(c)
		assert.Equalf(t, val.output, out, "TEST[%d], failed.\n%s", i, val.desc)
		assert.Equalf(t, val.expErr, err, "TEST[%d], failed.\n%s", i, val.desc)
	}
}
