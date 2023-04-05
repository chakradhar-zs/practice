package product

import (
	"errors"
	"mock-productservice/internal/models"
	"mock-productservice/internal/service"
	"mock-productservice/internal/store"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestProductCreateWithoutName(t *testing.T) {
	svc := New(nil, nil)

	_, err := svc.Create(&service.ProductCreate{
		Name:      "",
		BrandName: "",
	})

	if !reflect.DeepEqual(err, errors.New("Invalid product name")) {
		t.Errorf("unexpected error %v", err)
	}
}

func TestProductCreateWithName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	brandMock := service.NewMockBrandService(ctrl)
	brandMock.EXPECT().Get("Amul").Return(0, errors.New("unable to fetch brand"))

	svc := New(nil, brandMock)

	_, err := svc.Create(&service.ProductCreate{
		Name:      "Ghee",
		BrandName: "Amul",
	})

	if !reflect.DeepEqual(err, errors.New("unable to fetch brand")) {
		t.Errorf("unexpected error %v", err)
	}
}

func TestProductCreateWithMultipleNames(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	brandMock := service.NewMockBrandService(ctrl)
	brandMock.EXPECT().Get(gomock.Any()).Return(0, nil)
	brandMock.EXPECT().Create("Amul").Return(0, errors.New("error while creating"))

	svc := New(nil, brandMock)

	_, err := svc.Create(&service.ProductCreate{
		Name:      "Ghee 1 ltr",
		BrandName: "Amul",
	})

	if !reflect.DeepEqual(err, errors.New("error while creating")) {
		t.Errorf("unexpected error %v", err)
	}
}

func TestProductCreateSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	brandMock := service.NewMockBrandService(ctrl)
	brandMock.EXPECT().Get(gomock.Any()).Return(0, nil)
	brandMock.EXPECT().Create("Amul").Return(1, nil)

	storeMock := store.NewMockProductStorer(ctrl)
	storeMock.EXPECT().Create(&models.Product{
		Name:    "Ghee 1 ltr",
		BrandId: 1,
	}).Return(&models.Product{
		Name:    "Ghee 1 ltr",
		BrandId: 1,
		Id:      1,
	}, nil)
	svc := New(storeMock, brandMock)

	product, err := svc.Create(&service.ProductCreate{
		Name:      "Ghee 1 ltr",
		BrandName: "Amul",
	})

	if !reflect.DeepEqual(err, nil) {
		t.Errorf("unexpected error %v", err)
	}
	if product.Id != 1 {
		t.Errorf("unexpected id %v", product.Id)
	}
}

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	brandMock := service.NewMockBrandService(ctrl)
	storeMock := store.NewMockProductStorer(ctrl)
	svc := New(storeMock, brandMock)
	tcs := []struct {
		Input      *service.ProductCreate
		ExpErr     error
		ExpProduct *models.Product
		Call       []*gomock.Call
	}{
		{
			&service.ProductCreate{},
			errors.New("invalid product name"),
			nil,
			nil,
		},
		{
			&service.ProductCreate{
				Name:      "Ghee",
				BrandName: "Amul",
			},
			errors.New("failed to fetch"),
			nil,
			[]*gomock.Call{
				brandMock.EXPECT().Get("Amul").Return(0, errors.New("failed to fetch")),
			},
		},
		{
			&service.ProductCreate{
				Name:      "Ghee",
				BrandName: "Amul",
			},
			nil,
			&models.Product{
				Id:      1,
				Name:    "Ghee",
				BrandId: 1,
			},
			[]*gomock.Call{
				brandMock.EXPECT().Get(gomock.Any()).Return(1, nil),
				storeMock.EXPECT().Create(&models.Product{
					Name:    "Ghee",
					BrandId: 1,
				}).Return(&models.Product{
					Name:    "Ghee",
					BrandId: 1,
					Id:      1,
				}, nil),
			},
		},
	}

	for _, val := range tcs {
		out, err := svc.Create(val.Input)
		if !reflect.DeepEqual(err, val.ExpErr) {
			t.Errorf("unexpected error got %v, want %v", err, val.ExpErr)
		}

		if !reflect.DeepEqual(out, val.ExpProduct) {
			t.Errorf("unexpected product got %v, want %v", out, val.ExpProduct)
		}
	}
}
