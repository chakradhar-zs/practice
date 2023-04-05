package brand

import (
	"strconv"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"

	"Day-19/internal/models"
	"Day-19/internal/service"
)

type Handler struct {
	svc service.Brand
}

func New(s service.Brand) *Handler {
	return &Handler{svc: s}
}

func (h *Handler) Read(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("id")

	if i == "" {
		return models.Brand{}, errors.MissingParam{Param: []string{"id"}}
	}

	id, err := strconv.Atoi(i)

	if err != nil {
		return models.Brand{}, errors.InvalidParam{Param: []string{"id"}}
	}

	resp, err := h.svc.GetBrand(ctx, id)

	if err != nil {
		return models.Brand{}, errors.EntityNotFound{}
	}

	return resp, nil
}

func (h *Handler) Create(ctx *gofr.Context) (interface{}, error) {
	var b models.Brand

	_ = ctx.Bind(&b)

	resp, err := h.svc.CreateBrand(ctx, b)

	if err != nil {
		return 0, err
	}

	return resp, nil
}

func (h *Handler) Update(ctx *gofr.Context) (interface{}, error) {
	var b models.Brand

	i := ctx.PathParam("id")

	if i == "" {
		return 0, errors.MissingParam{Param: []string{"id"}}
	}

	id, err := strconv.Atoi(i)

	if err != nil {
		return 0, errors.InvalidParam{Param: []string{"id"}}
	}

	_ = ctx.Bind(&b)

	resp, err := h.svc.UpdateBrand(ctx, id, b)

	if err != nil {
		return 0, err
	}

	return resp, nil
}
