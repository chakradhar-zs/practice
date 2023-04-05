package product

import (
	"api-product/internal/models"
	"api-product/internal/store"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"developer.zopsmart.com/go/gofr/pkg/gofr/request"
	"developer.zopsmart.com/go/gofr/pkg/gofr/responder"
	"github.com/stretchr/testify/assert"
)

//type mockStore struct {
//}

//func (m mockStore) Get(ctx *gofr.Context, id int) (interface{}, error) {
//	if id == 3 {
//		return models.Product{3, "sneaker shoes", "stylish", 1000, 3, "shoes", 4, "Available"}, nil
//	}
//	if id == 333 {
//		return nil, errors.EntityNotFound{}
//	}
//	return nil, nil
//}

//func (m mockStore) Create(ctx *gofr.Context, prod models.Product) (interface{}, error) {
//	if reflect.DeepEqual(prod, models.Product{6, "maggi", "tasty", 50, 3, "noodles", 1, "Available"}) {
//		return 1, nil
//	}
//	if reflect.DeepEqual(prod, models.Product{}) {
//		return 0, errors.MissingParam{Param: []string{"body"}}
//	}
//	return nil, nil
//}
//
//func (m mockStore) Update(ctx *gofr.Context, id int, prod models.Product) (interface{}, error) {
//	if id == 333 {
//		return 0, errors.EntityNotFound{Entity: "id"}
//	}
//	if reflect.DeepEqual(prod, models.Product{6, "Maggi", "yummy", 50, 3, "noodles", 1, "Available"}) {
//		return 1, nil
//	}
//	if reflect.DeepEqual(prod, models.Product{}) {
//		return 0, errors.MissingParam{Param: []string{"body"}}
//	}
//	return nil, nil
//}
//
//func (m mockStore) Del(ctx *gofr.Context, id int) (interface{}, error) {
//	if id == 333 || id == 99 {
//		return 0, errors.EntityNotFound{Entity: "id"}
//	}
//	if id == 5 {
//		return 1, nil
//	}
//	return nil, nil
//}
func TestGetProduct(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	storeMock := store.NewMockProductStorer(ctrl)
	//svc:= New(getMock)
	tests := []struct {
		desc   string
		input  int
		input2 string
		output interface{}
		expErr error
		calls  []*gomock.Call
	}{
		{desc: "Success",
			input:  3,
			input2: "true",
			output: models.Product{Id: 3, Name: "sneaker shoes", Description: "stylish", Price: 1000, Quantity: 3, Category: "shoes", Brand: models.Brand{Id: 4, Name: "Nike"}, Status: "Available"},
			expErr: nil,
			calls: []*gomock.Call{
				storeMock.EXPECT().Get(gomock.AssignableToTypeOf(&gofr.Context{}), 3, "true").
					Return(models.Product{Id: 3, Name: "sneaker shoes", Description: "stylish", Price: 1000, Quantity: 3, Category: "shoes", Brand: models.Brand{Id: 4, Name: "Nike"}, Status: "Available"}, nil),
			}},
		{desc: "Fail",
			input:  333,
			input2: "true",
			output: nil,
			expErr: errors.EntityNotFound{},
			calls: []*gomock.Call{
				storeMock.EXPECT().Get(gomock.AssignableToTypeOf(&gofr.Context{}), 333, "true").
					Return(nil, errors.EntityNotFound{}),
			}},
		{desc: "Success",
			input:  4,
			input2: "false",
			output: models.Product{Id: 4, Name: "maggi", Description: "yum", Price: 100, Quantity: 3, Category: "noodles", Brand: models.Brand{Id: 5, Name: ""}, Status: "Available"},
			expErr: nil,
			calls: []*gomock.Call{
				storeMock.EXPECT().Get(gomock.AssignableToTypeOf(&gofr.Context{}), 4, "false").
					Return(models.Product{Id: 4, Name: "maggi", Description: "yum", Price: 100, Quantity: 3, Category: "noodles", Brand: models.Brand{Id: 5, Name: ""}, Status: "Available"}, nil),
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
		input  models.Product
		output interface{}
		expErr error
		Call   []*gomock.Call
	}{
		{desc: "Success",
			input:  models.Product{Id: 6, Name: "maggi", Description: "tasty", Price: 50, Quantity: 3, Category: "noodles", Brand: models.Brand{Id: 1, Name: ""}, Status: "Available"},
			output: 1,
			expErr: nil,
			Call: []*gomock.Call{
				storeMock.EXPECT().
					Create(gomock.AssignableToTypeOf(&gofr.Context{}), models.Product{Id: 6, Name: "maggi", Description: "tasty", Price: 50, Quantity: 3, Category: "noodles", Brand: models.Brand{Id: 1, Name: ""}, Status: "Available"}).
					Return(1, nil),
			}},
		{desc: "Fail",
			input:  models.Product{},
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
		input2 models.Product
		output interface{}
		expErr error
		Calls  []*gomock.Call
	}{
		{desc: "Success",
			input1: 6,
			input2: models.Product{Id: 6, Name: "Maggi", Description: "yummy", Price: 50, Quantity: 3, Category: "noodles", Brand: models.Brand{Id: 1, Name: ""}, Status: "Available"},
			output: 1,
			expErr: nil,
			Calls: []*gomock.Call{
				storeMock.EXPECT().
					Update(gomock.AssignableToTypeOf(&gofr.Context{}), 6, models.Product{Id: 6, Name: "Maggi", Description: "yummy", Price: 50, Quantity: 3, Category: "noodles", Brand: models.Brand{Id: 1, Name: ""}, Status: "Available"}).
					Return(1, nil),
			}},
		{desc: "Fail",
			input1: 2,
			input2: models.Product{},
			output: 0,
			expErr: errors.MissingParam{Param: []string{"body"}},
			Calls:  nil,
		},
		{desc: "Fail",
			input1: 333,
			input2: models.Product{Id: 333, Name: "Maggi", Description: "yummy", Price: 50, Quantity: 3, Category: "noodles", Brand: models.Brand{Id: 1, Name: ""}, Status: "Available"},
			output: 0,
			expErr: errors.EntityNotFound{Entity: "id"},
			Calls: []*gomock.Call{
				storeMock.EXPECT().
					Update(gomock.AssignableToTypeOf(&gofr.Context{}), 333, models.Product{Id: 333, Name: "Maggi", Description: "yummy", Price: 50, Quantity: 3, Category: "noodles", Brand: models.Brand{Id: 1, Name: ""}, Status: "Available"}).
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

//
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
			path:   "/product/?brand=true",
			input:  "true",
			output: []models.Product{{Id: 3, Name: "sneaker shoes", Description: "stylish", Price: 1000, Quantity: 3, Category: "shoes", Brand: models.Brand{Id: 4, Name: "Nike"}, Status: "Available"}},
			expErr: nil,
			calls: []*gomock.Call{
				storeMock.EXPECT().GetAll(gomock.AssignableToTypeOf(&gofr.Context{}), "true").
					Return([]models.Product{{Id: 3, Name: "sneaker shoes", Description: "stylish", Price: 1000, Quantity: 3, Category: "shoes", Brand: models.Brand{Id: 4, Name: "Nike"}, Status: "Available"}}, nil),
			}},
		{desc: "Fail",
			path:   "/product/?brand=true",
			input:  "true",
			output: nil,
			expErr: errors.EntityNotFound{},
			calls: []*gomock.Call{
				storeMock.EXPECT().GetAll(gomock.AssignableToTypeOf(&gofr.Context{}), "true").
					Return(nil, errors.EntityNotFound{}),
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
