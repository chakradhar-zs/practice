package product

import (
	"api-product/internal/models"
	"api-product/internal/service"
	"strconv"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Handler struct {
	svc service.Product
}

func New(s service.Product) *Handler {
	return &Handler{svc: s}
}

func (h *Handler) Read(c *gofr.Context) (interface{}, error) {
	i := c.PathParam("id")
	brand := c.Param("brand")
	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}
	id, err := strconv.Atoi(i)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	resp, _ := h.svc.GetProduct(c, id, brand)

	return resp, nil
}

func (h *Handler) Create(c *gofr.Context) (interface{}, error) {
	var prod models.Product
	_ = c.Bind(&prod)

	resp, err := h.svc.CreateProduct(c, prod)
	if err != nil {
		return 0, err
	}
	return resp, nil
}

func (h *Handler) Update(c *gofr.Context) (interface{}, error) {
	var prod models.Product

	i := c.PathParam("id")
	if i == "" {
		return 0, errors.MissingParam{Param: []string{"id"}}
	}
	id, err := strconv.Atoi(i)
	if err != nil {
		return 0, errors.InvalidParam{Param: []string{"id"}}
	}
	_ = c.Bind(&prod)

	resp, err := h.svc.UpdateProduct(c, id, prod)
	if err != nil {
		return 0, err
	}
	return resp, nil
}

func (h *Handler) Delete(c *gofr.Context) (interface{}, error) {
	i := c.PathParam("id")
	if i == "" {
		return 0, errors.MissingParam{Param: []string{"id"}}
	}
	id, err := strconv.Atoi(i)
	if err != nil {
		return 0, errors.InvalidParam{Param: []string{"id"}}
	}
	resp, err := h.svc.DeleteProduct(c, id)
	if err != nil {
		return 0, errors.EntityNotFound{Entity: "id"}
	}
	return resp, nil
}

func (h *Handler) Index(ctx *gofr.Context) (interface{}, error) {
	brand := ctx.Param("brand")
	resp, err := h.svc.GetAllProducts(ctx, brand)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
