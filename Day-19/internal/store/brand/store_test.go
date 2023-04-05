package brand

import (
	"context"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/datastore"
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"

	"github.com/stretchr/testify/assert"

	"github.com/DATA-DOG/go-sqlmock"

	"Day-19/internal/models"
)

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
		output  models.Brand
		mockErr error
		expErr  error
	}{
		{desc: "Success",
			input:   6,
			output:  models.Brand{ID: 6, Name: "Bru"},
			mockErr: nil,
			expErr:  nil,
		},
		{desc: "Fail",
			input:   11,
			output:  models.Brand{},
			mockErr: errors.EntityNotFound{},
			expErr:  errors.EntityNotFound{},
		},
	}

	for i, val := range tests {
		st := New()
		row := mock.NewRows([]string{"id", "name"}).AddRow(val.output.ID, val.output.Name)
		mock.ExpectQuery("select id,name from brands").
			WithArgs(val.input).WillReturnRows(row).
			WillReturnError(val.mockErr)

		out, err := st.Get(ctx, val.input)
		assert.Equalf(t, val.output, out, "Test[%d] failed \n%s", i, val.desc)
		assert.Equalf(t, val.expErr, err, "Test[%d] failed \n%s", i, val.desc)
	}
}

func TestCreate(t *testing.T) {
	ctx := gofr.NewContext(nil, nil, gofr.New())
	db, mock, err := sqlmock.New()

	if err != nil {
		ctx.Logger.Error("Error while opening a mock db connection")
	}

	ctx.DataStore = datastore.DataStore{ORM: db}
	ctx.Context = context.Background()
	tests := []struct {
		desc    string
		input   models.Brand
		output  int64
		mockErr error
		expErr  error
	}{
		{desc: "Success",
			input:   models.Brand{ID: 3, Name: "Nike"},
			output:  1,
			mockErr: nil,
			expErr:  nil,
		},
		{desc: "Fail",
			input:   models.Brand{},
			output:  0,
			mockErr: errors.MissingParam{},
			expErr:  errors.MissingParam{},
		},
	}

	for i, val := range tests {
		st := New()

		mock.ExpectExec("insert into").WithArgs(val.input.ID, val.input.Name).
			WillReturnResult(sqlmock.NewResult(int64(val.input.ID), val.output)).
			WillReturnError(val.mockErr)

		out, err := st.Create(ctx, val.input)
		assert.Equalf(t, val.output, out, "Test[%d] failed \n%s", i, val.desc)
		assert.Equalf(t, val.expErr, err, "Test[%d] failed \n%s", i, val.desc)
	}
}

func TestUpdate(t *testing.T) {
	ctx := gofr.NewContext(nil, nil, gofr.New())
	db, mock, err := sqlmock.New()

	if err != nil {
		ctx.Logger.Error("Error while opening a mock db connection")
	}

	ctx.DataStore = datastore.DataStore{ORM: db}
	ctx.Context = context.Background()
	tests := []struct {
		desc    string
		input1  int
		input2  models.Brand
		output  int64
		mockErr error
		expErr  error
	}{
		{desc: "Success",
			input1:  6,
			input2:  models.Brand{ID: 6, Name: "bru"},
			output:  1,
			mockErr: nil,
			expErr:  nil,
		},
		{desc: "Fail",
			input1:  99,
			input2:  models.Brand{},
			output:  0,
			mockErr: errors.EntityNotFound{},
			expErr:  errors.EntityNotFound{},
		},
	}

	for i, val := range tests {
		st := New()

		mock.ExpectExec("update").
			WithArgs(val.input2.Name, val.input1).
			WillReturnResult(sqlmock.NewResult(int64(val.input1), val.output)).
			WillReturnError(val.mockErr)

		out, err := st.Update(ctx, val.input1, val.input2)
		assert.Equalf(t, val.output, out, "Test[%d] failed \n%s", i, val.desc)
		assert.Equalf(t, val.expErr, err, "Test[%d] failed \n%s", i, val.desc)
	}
}
