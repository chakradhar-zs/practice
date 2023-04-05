package product

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"developer.zopsmart.com/go/gofr/pkg/gofr/request"
	"developer.zopsmart.com/go/gofr/pkg/gofr/responder"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"Day-19/internal/models"
	"Day-19/internal/store"
)

func TestGetProduct(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	storeMock := store.NewMockProductStorer(ctrl)
	tests := []struct {
		desc   string
		input  int
		input2 string
		output interface{}
		expErr error
		calls  []*gomock.Call
	}{
		{
			desc:   "Success",
			input:  3,
			input2: "true",
			output: models.Product{
				ID: 3, Name: "sneaker shoes", Description: "stylish", Price: 1000, Quantity: 3, Category: "shoes",
				Brand: models.Brand{ID: 4, Name: "Nike"}, Status: "Available"},
			expErr: nil,
			calls: []*gomock.Call{
				storeMock.EXPECT().Get(gomock.AssignableToTypeOf(&gofr.Context{}), 3, "true").
					Return(models.Product{
						ID: 3, Name: "sneaker shoes", Description: "stylish", Price: 1000, Quantity: 3, Category: "shoes",
						Brand: models.Brand{ID: 4, Name: "Nike"}, Status: "Available"}, nil),
			}},
		{desc: "Fail",
			input:  333,
			input2: "true",
			output: models.Product{},
			expErr: errors.EntityNotFound{},
			calls: []*gomock.Call{
				storeMock.EXPECT().Get(gomock.AssignableToTypeOf(&gofr.Context{}), 333, "true").
					Return(models.Product{}, errors.EntityNotFound{}),
			}},
	}

	for i, val := range tests {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)

		req := request.NewHTTPRequest(r)
		res := responder.NewContextualResponder(w, r)
		ctx := gofr.NewContext(res, req, app)
		s := New(storeMock)
		out, err := s.GetProduct(ctx, val.input, val.input2)
		assert.Equalf(t, val.output, out, "TEST[%d], failed.\n%s", i, val.desc)
		assert.Equalf(t, val.expErr, err, "TEST[%d], failed.\n%s", i, val.desc)
	}
}

func TestGet(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	storeMock := store.NewMockProductStorer(ctrl)
	tests := []struct {
		desc   string
		input  int
		input2 string
		output interface{}
		expErr error
		calls  []*gomock.Call
	}{
		{desc: "Success",
			input:  4,
			input2: "false",
			output: models.Product{
				ID: 4, Name: "maggi", Description: "yum", Price: 100, Quantity: 3, Category: "noodles",
				Brand: models.Brand{ID: 5}, Status: "Available"},
			expErr: nil,
			calls: []*gomock.Call{
				storeMock.EXPECT().Get(gomock.AssignableToTypeOf(&gofr.Context{}), 4, "false").
					Return(models.Product{
						ID: 4, Name: "maggi", Description: "yum", Price: 100, Quantity: 3, Category: "noodles",
						Brand: models.Brand{ID: 5}, Status: "Available"}, nil),
			}},
	}

	for i, val := range tests {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)

		req := request.NewHTTPRequest(r)
		res := responder.NewContextualResponder(w, r)
		ctx := gofr.NewContext(res, req, app)
		s := New(storeMock)
		out, err := s.GetProduct(ctx, val.input, val.input2)
		assert.Equalf(t, val.output, out, "TEST[%d], failed.\n%s", i, val.desc)
		assert.Equalf(t, val.expErr, err, "TEST[%d], failed.\n%s", i, val.desc)
	}
}

func TestCreateProduct(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	storeMock := store.NewMockProductStorer(ctrl)
	tests := []struct {
		desc   string
		input  *models.Product
		output interface{}
		expErr error
		Call   []*gomock.Call
	}{
		{desc: "Success",
			input: &models.Product{
				ID: 6, Name: "maggi", Description: "tasty", Price: 50, Quantity: 3, Category: "noodles",
				Brand: models.Brand{ID: 1, Name: ""}, Status: "Available"},
			output: 1,
			expErr: nil,
			Call: []*gomock.Call{
				storeMock.EXPECT().
					Create(gomock.AssignableToTypeOf(&gofr.Context{}), &models.Product{
						ID: 6, Name: "maggi", Description: "tasty", Price: 50, Quantity: 3, Category: "noodles",
						Brand: models.Brand{ID: 1, Name: ""}, Status: "Available"}).
					Return(1, nil),
			}},
		{desc: "Fail",
			input:  &models.Product{},
			output: 0,
			expErr: errors.MissingParam{Param: []string{"body"}},
			Call:   nil,
		},
	}

	for i, val := range tests {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		req := request.NewHTTPRequest(r)
		res := responder.NewContextualResponder(w, r)
		ctx := gofr.NewContext(res, req, app)
		s := New(storeMock)
		out, err := s.CreateProduct(ctx, val.input)
		assert.Equalf(t, val.output, out, "TEST[%d], failed.\n%s", i, val.desc)
		assert.Equalf(t, val.expErr, err, "TEST[%d], failed.\n%s", i, val.desc)
	}
}

func TestUpdateProduct(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	storeMock := store.NewMockProductStorer(ctrl)
	tests := []struct {
		desc   string
		input1 int
		input2 *models.Product
		output interface{}
		expErr error
		Calls  []*gomock.Call
	}{
		{desc: "Success",
			input1: 6,
			input2: &models.Product{
				ID: 6, Name: "Maggi", Description: "yummy", Price: 50, Quantity: 3, Category: "noodles",
				Brand: models.Brand{ID: 1, Name: ""}, Status: "Available"},
			output: 1,
			expErr: nil,
			Calls: []*gomock.Call{
				storeMock.EXPECT().
					Update(gomock.AssignableToTypeOf(&gofr.Context{}), 6, &models.Product{
						ID: 6, Name: "Maggi", Description: "yummy", Price: 50, Quantity: 3, Category: "noodles",
						Brand: models.Brand{ID: 1, Name: ""}, Status: "Available"}).
					Return(1, nil),
			}},
		{desc: "Fail",
			input1: 2,
			input2: &models.Product{},
			output: 0,
			expErr: errors.MissingParam{Param: []string{"body"}},
			Calls:  nil,
		},
		{desc: "Fail",
			input1: 333,
			input2: &models.Product{
				ID: 333, Name: "Maggi", Description: "yummy", Price: 50, Quantity: 3, Category: "noodles",
				Brand: models.Brand{ID: 1, Name: ""}, Status: "Available"},
			output: 0,
			expErr: errors.EntityNotFound{Entity: "id"},
			Calls: []*gomock.Call{
				storeMock.EXPECT().
					Update(gomock.AssignableToTypeOf(&gofr.Context{}), 333, &models.Product{
						ID: 333, Name: "Maggi", Description: "yummy", Price: 50, Quantity: 3, Category: "noodles",
						Brand: models.Brand{ID: 1, Name: ""}, Status: "Available"}).
					Return(0, errors.EntityNotFound{Entity: "id"}),
			}},
	}

	for i, val := range tests {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		req := request.NewHTTPRequest(r)
		res := responder.NewContextualResponder(w, r)
		ctx := gofr.NewContext(res, req, app)
		s := New(storeMock)
		out, err := s.UpdateProduct(ctx, val.input1, val.input2)
		assert.Equalf(t, val.output, out, "TEST[%d], failed.\n%s", i, val.desc)
		assert.Equalf(t, val.expErr, err, "TEST[%d], failed.\n%s", i, val.desc)
	}
}

func TestDeleteProduct(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	storeMock := store.NewMockProductStorer(ctrl)
	tests := []struct {
		desc   string
		input  int
		output interface{}
		expErr error
		calls  []*gomock.Call
	}{
		{desc: "Success",
			input:  5,
			output: 1,
			expErr: nil,
			calls: []*gomock.Call{
				storeMock.EXPECT().Del(gomock.AssignableToTypeOf(&gofr.Context{}), 5).
					Return(1, nil),
			}},
		{desc: "Fail",
			input:  99,
			output: 0,
			expErr: errors.EntityNotFound{Entity: "id"},
			calls: []*gomock.Call{
				storeMock.EXPECT().Del(gomock.AssignableToTypeOf(&gofr.Context{}), 99).
					Return(0, errors.EntityNotFound{Entity: "id"}),
			}},
		{desc: "Fail",
			input:  333,
			output: 0,
			expErr: errors.EntityNotFound{Entity: "id"},
			calls: []*gomock.Call{
				storeMock.EXPECT().Del(gomock.AssignableToTypeOf(&gofr.Context{}), 333).
					Return(0, errors.EntityNotFound{Entity: "id"}),
			}},
	}

	for i, val := range tests {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)

		req := request.NewHTTPRequest(r)
		res := responder.NewContextualResponder(w, r)
		ctx := gofr.NewContext(res, req, app)
		s := New(storeMock)
		out, err := s.DeleteProduct(ctx, val.input)
		assert.Equalf(t, val.output, out, "TEST[%d], failed.\n%s", i, val.desc)
		assert.Equalf(t, val.expErr, err, "TEST[%d], failed.\n%s", i, val.desc)
	}
}

func TestGetAllProducts(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	storeMock := store.NewMockProductStorer(ctrl)
	tests := []struct {
		desc   string
		path   string
		input  string
		output interface{}
		expErr error
		calls  []*gomock.Call
	}{
		{desc: "Success",
			path:  "/product/?brand=true",
			input: "true",
			output: []models.Product{{
				ID: 3, Name: "sneaker shoes", Description: "stylish", Price: 1000, Quantity: 3, Category: "shoes",
				Brand: models.Brand{ID: 4, Name: "Nike"}, Status: "Available"}},
			expErr: nil,
			calls: []*gomock.Call{
				storeMock.EXPECT().GetAll(gomock.AssignableToTypeOf(&gofr.Context{}), "true").
					Return([]models.Product{{
						ID: 3, Name: "sneaker shoes", Description: "stylish", Price: 1000, Quantity: 3, Category: "shoes",
						Brand: models.Brand{ID: 4, Name: "Nike"}, Status: "Available"}}, nil),
			}},
		{desc: "Fail",
			path:   "/product/?brand=true",
			input:  "true",
			output: []models.Product{},
			expErr: errors.EntityNotFound{},
			calls: []*gomock.Call{
				storeMock.EXPECT().GetAll(gomock.AssignableToTypeOf(&gofr.Context{}), "true").
					Return([]models.Product{}, errors.EntityNotFound{}),
			}},
	}

	for i, val := range tests {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, val.path, nil)

		req := request.NewHTTPRequest(r)
		res := responder.NewContextualResponder(w, r)
		ctx := gofr.NewContext(res, req, app)
		s := New(storeMock)
		out, err := s.GetAllProducts(ctx, val.input)
		assert.Equalf(t, val.output, out, "TEST[%d], failed.\n%s", i, val.desc)
		assert.Equalf(t, val.expErr, err, "TEST[%d], failed.\n%s", i, val.desc)
	}
}

func TestGetProductByNAme(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	storeMock := store.NewMockProductStorer(ctrl)
	tests := []struct {
		desc   string
		path   string
		input1 string
		input2 string
		output interface{}
		expErr error
		calls  []*gomock.Call
	}{
		{desc: "Success",
			path:   "/product/?brand=true",
			input1: "zs_sneaker shoes",
			input2: "true",
			output: []models.Product{{
				ID: 3, Name: "zs_sneaker shoes", Description: "stylish", Price: 1000, Quantity: 3, Category: "shoes",
				Brand: models.Brand{ID: 4, Name: "Nike"}, Status: "Available"}},
			expErr: nil,
			calls: []*gomock.Call{
				storeMock.EXPECT().GetByName(gomock.AssignableToTypeOf(&gofr.Context{}), "zs_sneaker shoes", "true").
					Return([]models.Product{{
						ID: 3, Name: "zs_sneaker shoes", Description: "stylish", Price: 1000, Quantity: 3, Category: "shoes",
						Brand: models.Brand{ID: 4, Name: "Nike"}, Status: "Available"}}, nil),
			}},
		{desc: "Fail",
			path:   "/product/?brand=true",
			input1: "",
			input2: "true",
			output: []models.Product{},
			expErr: errors.EntityNotFound{},
			calls: []*gomock.Call{
				storeMock.EXPECT().GetByName(gomock.AssignableToTypeOf(&gofr.Context{}), "", "true").
					Return([]models.Product{}, errors.EntityNotFound{}),
			}},
	}

	for i, val := range tests {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, val.path, nil)

		req := request.NewHTTPRequest(r)
		res := responder.NewContextualResponder(w, r)
		ctx := gofr.NewContext(res, req, app)
		s := New(storeMock)
		out, err := s.GetProductByNAme(ctx, val.input1, val.input2)
		assert.Equalf(t, val.output, out, "TEST[%d], failed.\n%s", i, val.desc)
		assert.Equalf(t, val.expErr, err, "TEST[%d], failed.\n%s", i, val.desc)
	}
}
