package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hilmiikhsan/shopeefun-product-service/internal/adapter"
	"github.com/hilmiikhsan/shopeefun-product-service/internal/middleware"
	"github.com/hilmiikhsan/shopeefun-product-service/internal/module/shop/entity"
	"github.com/hilmiikhsan/shopeefun-product-service/internal/module/shop/ports"
	"github.com/hilmiikhsan/shopeefun-product-service/internal/module/shop/repository"
	"github.com/hilmiikhsan/shopeefun-product-service/internal/module/shop/service"
	"github.com/hilmiikhsan/shopeefun-product-service/pkg/errmsg"
	"github.com/hilmiikhsan/shopeefun-product-service/pkg/response"
	"github.com/rs/zerolog/log"
)

type shopHandler struct {
	service ports.ShopService
}

func NewShopHandler() *shopHandler {
	var (
		handler = new(shopHandler)
		repo    = repository.NewShopRepository(adapter.Adapters.ShopeefunPostgres)
		service = service.NewShopService(repo)
	)
	handler.service = service

	return handler
}

func (h *shopHandler) Register(router fiber.Router) {
	router.Get("/shops", middleware.UserIdHeader, h.GetShops)
	router.Post("/shops", middleware.UserIdHeader, h.CreateShop)
	router.Get("/shops/:id", h.GetShop)
	router.Delete("/shops/:id", middleware.UserIdHeader, h.DeleteShop)
	router.Patch("/shops/:id", middleware.UserIdHeader, h.UpdateShop)
}

func (h *shopHandler) CreateShop(c *fiber.Ctx) error {
	var (
		req        = new(entity.CreateShopRequest)
		ctx        = c.Context()
		validators = adapter.Adapters.Validator
		locals     = middleware.GetLocals(c)
	)

	if err := c.BodyParser(req); err != nil {
		log.Warn().Err(err).Msg("handler::CreateShop - Parse request body")
		return c.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}

	req.UserId = locals.UserId

	if err := validators.Validate(req); err != nil {
		log.Warn().Err(err).Any("payload", req).Msg("handler::CreateShop - Validate request body")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	resp, err := h.service.CreateShop(ctx, req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.Status(fiber.StatusCreated).JSON(response.Success(resp, ""))

}

func (h *shopHandler) GetShop(c *fiber.Ctx) error {
	var (
		req        = new(entity.GetShopRequest)
		ctx        = c.Context()
		validators = adapter.Adapters.Validator
	)

	req.Id = c.Params("id")

	if err := validators.Validate(req); err != nil {
		log.Warn().Err(err).Any("payload", req).Msg("handler::GetShop - Validate request body")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	resp, err := h.service.GetShop(ctx, req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.Status(fiber.StatusOK).JSON(response.Success(resp, ""))
}

func (h *shopHandler) DeleteShop(c *fiber.Ctx) error {
	var (
		req        = new(entity.DeleteShopRequest)
		ctx        = c.Context()
		validators = adapter.Adapters.Validator
		locals     = middleware.GetLocals(c)
	)
	req.UserId = locals.UserId
	req.Id = c.Params("id")

	if err := validators.Validate(req); err != nil {
		log.Warn().Err(err).Any("payload", req).Msg("handler::DeleteShop - Validate request body")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	err := h.service.DeleteShop(ctx, req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.Status(fiber.StatusOK).JSON(response.Success(nil, ""))
}

func (h *shopHandler) UpdateShop(c *fiber.Ctx) error {
	var (
		req        = new(entity.UpdateShopRequest)
		ctx        = c.Context()
		validators = adapter.Adapters.Validator
		locals     = middleware.GetLocals(c)
	)

	if err := c.BodyParser(req); err != nil {
		log.Warn().Err(err).Msg("handler::UpdateShop - Parse request body")
		return c.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}

	req.UserId = locals.UserId
	req.Id = c.Params("id")

	if err := validators.Validate(req); err != nil {
		log.Warn().Err(err).Any("payload", req).Msg("handler::UpdateShop - Validate request body")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	resp, err := h.service.UpdateShop(ctx, req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.Status(fiber.StatusOK).JSON(response.Success(resp, ""))
}

func (h *shopHandler) GetShops(c *fiber.Ctx) error {
	var (
		req = new(entity.ShopsRequest)
		ctx = c.Context()
		v   = adapter.Adapters.Validator
		l   = middleware.GetLocals(c)
	)

	if err := c.QueryParser(req); err != nil {
		log.Warn().Err(err).Msg("handler::GetShops - Parse request query")
		return c.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}

	req.UserId = l.UserId
	req.SetDefault()

	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("payload", req).Msg("handler::GetShops - Validate request body")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	resp, err := h.service.GetShops(ctx, req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.Status(fiber.StatusOK).JSON(response.Success(resp, ""))
}
