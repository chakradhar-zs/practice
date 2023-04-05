package product

import (
	"context"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/datastore"
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"Day-19/internal/models"
)

const v = "true"

func TestGet(t *testing.T) {
	ctx := gofr.NewContext(nil, nil, gofr.New())
	db, mock, err := sqlmock.New()

	if err != nil {
		ctx.Logger.Error("Error while opening a mock db connection")
	}
	defer db.Close()

	ctx.DataStore = datastore.DataStore{ORM: db}
	ctx.Context = context.Background()

	tests := []struct {
		desc    string
		input   int
		input2  string
		output  models.Product
		mockErr error
		expErr  error
	}{
		{desc: "Success",
			input:  3,
			input2: "true",
			output: models.Product{
				ID: 3, Name: "sneaker shoes", Description: "stylish", Price: 1000, Quantity: 3, Category: "shoes",
				Brand: models.Brand{ID: 4, Name: "Nike"}, Status: "Available",
			},
			mockErr: nil,
			expErr:  nil,
		},
		{desc: "Success",
			input:  5,
			input2: "false",
			output: models.Product{
				ID: 5, Name: "Bru", Description: "tasty", Price: 100, Quantity: 3, Category: "coffee",
				Brand: models.Brand{ID: 6}, Status: "Available",
			},
			mockErr: nil,
			expErr:  nil,
		},
		{desc: "Fail",
			input:   333,
			input2:  "false",
			output:  models.Product{},
			mockErr: errors.EntityNotFound{},
			expErr:  errors.EntityNotFound{},
		},
		{desc: "Fail",
			input:   22,
			input2:  "true",
			output:  models.Product{},
			mockErr: errors.EntityNotFound{},
			expErr:  errors.EntityNotFound{},
		},
	}
	for _, val := range tests {
		st := New()
		row := mock.NewRows([]string{"id", "name", "description", "price", "quantity", "category", "brand_id", "status"}).
			AddRow(val.output.ID, val.output.Name, val.output.Description, val.output.Price,
				val.output.Quantity, val.output.Category, val.output.Brand.ID, val.output.Status)
		mock.ExpectQuery("select id,name,description,price,quantity,category,brand_id,status from products where id=?").
			WithArgs(val.input).
			WillReturnRows(row).
			WillReturnError(val.mockErr)

		rb := sqlmock.NewRows([]string{"name"}).AddRow(val.output.Brand.Name)
		mock.ExpectQuery("select name from brands").
			WithArgs(val.output.Brand.ID).
			WillReturnRows(rb).
			WillReturnError(val.expErr)

		out, err := st.Get(ctx, val.input, val.input2)
		assert.Equal(t, val.output, out, "TEST failed.")
		assert.Equal(t, val.expErr, err, "TEST failed.")
	}
}

func TestCreate(t *testing.T) {
	ctx := gofr.NewContext(nil, nil, gofr.New())
	db, mock, err := sqlmock.New()

	if err != nil {
		ctx.Logger.Error("Error while opening a mock db connection")
	}

	tests := []struct {
		desc    string
		input   *models.Product
		output  int64
		mockErr error
		expErr  error
	}{
		{desc: "Success",
			input: &models.Product{
				ID: 6, Name: "maggi", Description: "tasty", Price: 50, Quantity: 3, Category: "noodles",
				Brand: models.Brand{ID: 1, Name: ""}, Status: "Available",
			},
			output:  1,
			mockErr: nil,
			expErr:  nil,
		},
		{desc: "Fail",
			input:   &models.Product{},
			output:  0,
			mockErr: errors.MissingParam{},
			expErr:  errors.MissingParam{},
		},
	}

	for _, val := range tests {
		ctx.DataStore = datastore.DataStore{ORM: db}
		ctx.Context = context.Background()

		mock.ExpectExec("insert into").
			WithArgs(val.input.ID, val.input.Name, val.input.Description, val.input.Price, val.input.Quantity,
				val.input.Category, val.input.Brand.ID, val.input.Status).
			WillReturnResult(sqlmock.NewResult(6, val.output)).
			WillReturnError(val.mockErr)

		st := New()
		out, err := st.Create(ctx, val.input)
		assert.Equal(t, val.output, out, "TEST failed.")
		assert.Equal(t, val.expErr, err, "TEST failed.")
	}
}

func TestUpdate(t *testing.T) {
	ctx := gofr.NewContext(nil, nil, gofr.New())
	db, mock, err := sqlmock.New()

	if err != nil {
		ctx.Logger.Error("Error while opening a mock db connection")
	}

	tests := []struct {
		desc    string
		input1  int
		input2  *models.Product
		output  int64
		mockErr error
		expErr  error
	}{
		{desc: "Success",
			input1: 6,
			input2: &models.Product{
				ID: 6, Name: "Maggi", Description: "yummy", Price: 50, Quantity: 3, Category: "noodles",
				Brand: models.Brand{ID: 1, Name: ""}, Status: "Available",
			},
			output:  1,
			mockErr: nil,
			expErr:  nil,
		},
		{desc: "Fail",
			input1:  2,
			input2:  &models.Product{},
			output:  0,
			mockErr: errors.EntityNotFound{},
			expErr:  errors.EntityNotFound{},
		},
		{desc: "Fail",
			input1: 333,
			input2: &models.Product{
				ID: 333, Name: "Maggi", Description: "yummy", Price: 50, Quantity: 3, Category: "noodles",
				Brand: models.Brand{ID: 1, Name: ""}, Status: "Available",
			},
			output:  0,
			mockErr: errors.EntityNotFound{},
			expErr:  errors.EntityNotFound{},
		},
	}

	for _, val := range tests {
		ctx.DataStore = datastore.DataStore{ORM: db}
		ctx.Context = context.Background()

		mock.ExpectExec("update ").
			WithArgs(val.input2.Name, val.input2.Description, val.input2.Price, val.input2.Quantity,
				val.input2.Category, val.input2.Brand.ID, val.input2.Status, val.input2.ID).
			WillReturnResult(sqlmock.NewResult(6, val.output)).
			WillReturnError(val.mockErr)

		st := New()
		out, err := st.Update(ctx, val.input1, val.input2)
		assert.Equal(t, val.output, out, "TEST failed.")
		assert.Equal(t, val.expErr, err, "TEST failed.")
	}
}

func TestDel(t *testing.T) {
	ctx := gofr.NewContext(nil, nil, gofr.New())
	db, mock, err := sqlmock.New()

	if err != nil {
		ctx.Logger.Error("Error while opening a mock db connection")
	}

	tests := []struct {
		desc    string
		input   int
		output  int64
		mockErr error
		expErr  error
	}{
		{desc: "Success",
			input:   5,
			output:  1,
			mockErr: nil,
			expErr:  nil,
		},
		{desc: "Fail",
			input:   99,
			output:  0,
			mockErr: errors.EntityNotFound{Entity: "id"},
			expErr:  errors.EntityNotFound{Entity: "id"},
		},
		{desc: "Fail",
			input:   333,
			output:  0,
			mockErr: errors.EntityNotFound{Entity: "id"},
			expErr:  errors.EntityNotFound{Entity: "id"},
		},
	}

	for _, val := range tests {
		ctx.DataStore = datastore.DataStore{ORM: db}
		ctx.Context = context.Background()

		mock.ExpectExec("delete ").
			WithArgs(val.input).
			WillReturnResult(sqlmock.NewResult(5, val.output)).
			WillReturnError(val.mockErr)

		st := New()
		out, err := st.Del(ctx, val.input)
		assert.Equal(t, val.output, out, "TEST failed.")
		assert.Equal(t, val.expErr, err, "TEST failed.")
	}
}

func TestGetAll(t *testing.T) {
	ctx := gofr.NewContext(nil, nil, gofr.New())
	db, mock, err := sqlmock.New()

	if err != nil {
		ctx.Logger.Error("Error while opening a mock db connection")
	}

	row := mock.NewRows([]string{"id", "name", "description", "price", "quantity", "category", "brand_id", "status"}).
		AddRow(3, "sneaker shoes", "stylish", 1000, 3, "shoes", 4, "Available")
	rb := sqlmock.NewRows([]string{"name"}).AddRow("Nike")

	tests := []struct {
		desc    string
		input   string
		output  []models.Product
		mockErr error
		expErr  error
	}{
		{desc: "Success",
			input: "true",
			output: []models.Product{{
				ID: 3, Name: "sneaker shoes", Description: "stylish", Price: 1000, Quantity: 3, Category: "shoes",
				Brand: models.Brand{ID: 4, Name: "Nike"}, Status: "Available",
			}},
			mockErr: nil,
			expErr:  nil,
		},
		{desc: "Fail",
			input:   "true",
			output:  nil,
			mockErr: errors.EntityNotFound{},
			expErr:  errors.EntityNotFound{},
		},
		{desc: "Fail",
			input:   "false",
			output:  nil,
			mockErr: errors.EntityNotFound{},
			expErr:  errors.EntityNotFound{},
		},
	}

	for _, val := range tests {
		ctx.DataStore = datastore.DataStore{ORM: db}
		ctx.Context = context.Background()

		mock.ExpectQuery("select").WillReturnRows(row).WillReturnError(val.mockErr)

		if val.input == v {
			mock.ExpectQuery("select name from brands").
				WithArgs(4).
				WillReturnRows(rb).
				WillReturnError(val.mockErr)
		}

		st := New()
		out, err := st.GetAll(ctx, val.input)
		assert.Equal(t, val.output, out, "TEST failed.")
		assert.Equal(t, val.expErr, err, "TEST failed.")
	}
}

func TestGetByName(t *testing.T) {
	ctx := gofr.NewContext(nil, nil, gofr.New())
	db, mock, err := sqlmock.New()

	if err != nil {
		ctx.Logger.Error("Error while opening a mock db connection")
	}

	tests := []struct {
		desc    string
		input   string
		input2  string
		output  []models.Product
		mockErr error
		expErr  error
	}{
		{desc: "Success",
			input:  "zs_sneaker shoes",
			input2: "true",
			output: []models.Product{{
				ID: 3, Name: "zs_sneaker shoes", Description: "stylish", Price: 1000, Quantity: 3, Category: "shoes",
				Brand: models.Brand{ID: 4, Name: "Nike"}, Status: "Available",
			}},
			mockErr: nil,
			expErr:  nil,
		},
		{desc: "Fail",
			input:  "bag",
			input2: "true",
			output: []models.Product{{
				ID: 0, Name: "", Description: "", Price: 0, Quantity: 0, Category: "",
				Brand: models.Brand{ID: 0, Name: ""}, Status: ""}},
			mockErr: errors.EntityNotFound{},
			expErr:  errors.EntityNotFound{},
		},
		{desc: "Fail",
			input:  "chair",
			input2: "false",
			output: []models.Product{{
				ID: 0, Name: "", Description: "", Price: 0, Quantity: 0, Category: "",
				Brand: models.Brand{ID: 0, Name: ""}, Status: ""}},
			mockErr: errors.EntityNotFound{},
			expErr:  errors.EntityNotFound{},
		},
	}

	for _, val := range tests {
		ctx.DataStore = datastore.DataStore{ORM: db}
		ctx.Context = context.Background()

		row := mock.NewRows([]string{"id", "name", "description", "price", "quantity", "category", "brand_id", "status"}).
			AddRow(val.output[0].ID, val.output[0].Name, val.output[0].Description,
				val.output[0].Price, val.output[0].Quantity, val.output[0].Category, val.output[0].Brand.ID, val.output[0].Status)
		rb := sqlmock.NewRows([]string{"name"}).AddRow(val.output[0].Brand.Name)

		mock.ExpectQuery("select id,name,description,price,quantity,category,brand_id,status from products where name=?").
			WithArgs(val.input).
			WillReturnRows(row).WillReturnError(val.mockErr)

		if val.input2 == "true" {
			mock.ExpectQuery("select name from brands where id=?").
				WithArgs(val.output[0].Brand.ID).
				WillReturnRows(rb).
				WillReturnError(val.mockErr)
		}

		st := New()
		out, err := st.GetByName(ctx, val.input, val.input2)
		assert.Equal(t, val.output, out, "TEST failed.")
		assert.Equal(t, val.expErr, err, "TEST failed.")
	}
}

func TestByName(t *testing.T) {
	ctx := gofr.NewContext(nil, nil, gofr.New())
	db, mock, err := sqlmock.New()

	if err != nil {
		ctx.Logger.Error("Error while opening a mock db connection")
	}

	tests := []struct {
		desc    string
		input   string
		input2  string
		output  []models.Product
		mockErr error
		expErr  error
	}{
		{desc: "Fail",
			input:  "zs_nike",
			input2: "true",
			output: []models.Product{{
				ID: 0, Name: "", Description: "", Price: 0, Quantity: 0, Category: "",
				Brand: models.Brand{ID: 0, Name: ""}, Status: ""}},
			mockErr: errors.EntityNotFound{},
			expErr:  errors.EntityNotFound{},
		},
	}

	for _, val := range tests {
		ctx.DataStore = datastore.DataStore{ORM: db}
		ctx.Context = context.Background()

		row := mock.NewRows([]string{"id", "name", "description", "price", "quantity", "category", "brand_id", "status"}).
			AddRow(3, "zs_nike", 99, 100, 1, "shoe", 4, "Available")
		rb := sqlmock.NewRows([]string{"name"}).AddRow("Nike")

		mock.ExpectQuery("select id,name,description,price,quantity,category,brand_id,status from products where name=?").
			WithArgs(val.input).
			WillReturnRows(row).WillReturnError(val.mockErr)

		if val.input2 == "true" {
			mock.ExpectQuery("select name from brands where id=?").
				WithArgs(4).
				WillReturnRows(rb).
				WillReturnError(val.mockErr)
		}

		st := New()
		out, err := st.GetByName(ctx, val.input, val.input2)
		assert.Equal(t, val.output, out, "TEST failed.")
		assert.Equal(t, val.expErr, err, "TEST failed.")
	}
}
