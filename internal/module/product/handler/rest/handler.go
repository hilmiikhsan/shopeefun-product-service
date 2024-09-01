package rest

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/hilmiikhsan/shopeefun-product-service/internal/adapter"
	"github.com/hilmiikhsan/shopeefun-product-service/internal/middleware"
	"github.com/hilmiikhsan/shopeefun-product-service/internal/module/product/entity"
	"github.com/hilmiikhsan/shopeefun-product-service/internal/module/product/ports"
	"github.com/hilmiikhsan/shopeefun-product-service/internal/module/product/repository"
	"github.com/hilmiikhsan/shopeefun-product-service/internal/module/product/service"
	"github.com/hilmiikhsan/shopeefun-product-service/pkg/errmsg"
	"github.com/hilmiikhsan/shopeefun-product-service/pkg/response"
	"github.com/rs/zerolog/log"
)

type productHandler struct {
	service ports.ProductService
}

func NewProductHandler() *productHandler {
	var (
		handler = new(productHandler)
		repo    = repository.NewProductRepository(adapter.Adapters.ShopeefunPostgres)
		service = service.NewProductService(repo)
	)
	handler.service = service

	return handler
}

func (h *productHandler) Register(router fiber.Router) {
	router.Post("/products", middleware.UserIdHeader, h.CreateProduct)
	router.Get("/products/:id", middleware.UserIdHeader, h.GetProduct)
	router.Get("/products", middleware.UserIdHeader, h.GetProducts)
	router.Patch("/products/:id", middleware.UserIdHeader, h.UpdateProduct)
	router.Delete("/products/:id", middleware.UserIdHeader, h.DeleteProduct)
}

func (h *productHandler) CreateProduct(c *fiber.Ctx) error {
	var (
		req        = new(entity.CreateProductRequest)
		ctx        = c.Context()
		validators = adapter.Adapters.Validator
	)

	if err := c.BodyParser(req); err != nil {
		log.Warn().Err(err).Msg("handler::CreateProduct - Parse request body")
		return c.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}

	if err := validators.Validate(req); err != nil {
		log.Warn().Err(err).Any("payload", req).Msg("handler::CreateProduct - Validate request body")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	resp, err := h.service.CreateProduct(ctx, req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.Status(fiber.StatusCreated).JSON(response.Success(resp, ""))
}

func (h *productHandler) GetProduct(c *fiber.Ctx) error {
	var (
		req        = new(entity.GetProductRequest)
		ctx        = c.Context()
		validators = adapter.Adapters.Validator
	)

	req.Id = c.Params("id")

	fmt.Println("ID: ", req.Id)

	if err := validators.Validate(req); err != nil {
		log.Warn().Err(err).Any("payload", req).Msg("handler::GetShop - Validate request body")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	resp, err := h.service.GetProduct(ctx, req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.Status(fiber.StatusOK).JSON(response.Success(resp, ""))
}

func (h *productHandler) GetProducts(c *fiber.Ctx) error {
	var (
		req        = new(entity.ProductRequest)
		ctx        = c.Context()
		validators = adapter.Adapters.Validator
	)

	if err := c.QueryParser(req); err != nil {
		log.Warn().Err(err).Msg("handler::GetProducts - Parse request query")
		return c.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}

	req.SetDefault()

	if err := validators.Validate(req); err != nil {
		log.Warn().Err(err).Any("payload", req).Msg("handler::GetProducts - Validate query params")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	resp, err := h.service.GetProducts(ctx, req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.Status(fiber.StatusOK).JSON(response.Success(resp, ""))
}

func (h *productHandler) UpdateProduct(c *fiber.Ctx) error {
	var (
		req        = new(entity.UpdateProductRequest)
		ctx        = c.Context()
		validators = adapter.Adapters.Validator
	)

	if err := c.BodyParser(req); err != nil {
		log.Warn().Err(err).Msg("handler::UpdateProduct - Parse request body")
		return c.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}

	req.Id = c.Params("id")

	if err := validators.Validate(req); err != nil {
		log.Warn().Err(err).Any("payload", req).Msg("handler::UpdateProduct - Validate request body")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	resp, err := h.service.UpdateProduct(ctx, req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.Status(fiber.StatusOK).JSON(response.Success(resp, ""))
}

func (h *productHandler) DeleteProduct(c *fiber.Ctx) error {
	var (
		req        = new(entity.DeleteProductRequest)
		ctx        = c.Context()
		validators = adapter.Adapters.Validator
	)

	req.Id = c.Params("id")

	if err := validators.Validate(req); err != nil {
		log.Warn().Err(err).Any("payload", req).Msg("handler::DeleteProduct - Validate request body")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	err := h.service.DeleteProduct(ctx, req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.Status(fiber.StatusOK).JSON(response.Success(nil, ""))
}
