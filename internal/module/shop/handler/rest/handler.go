package handler

import (
	"codebase-app/internal/adapter"
	"codebase-app/internal/middleware"
	"codebase-app/internal/module/shop/entity"
	"codebase-app/internal/module/shop/ports"
	"codebase-app/internal/module/shop/repository"
	"codebase-app/internal/module/shop/service"
	"codebase-app/pkg/errmsg"
	"codebase-app/pkg/response"

	"github.com/gofiber/fiber/v2"
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
	router.Post("/products", middleware.UserIdHeader, middleware.UploadImageMiddleware, h.CreateProduct)
	router.Get("/products/all", middleware.UserIdHeader, h.GetAllProduct)
	router.Get("/products/:id", h.GetProductByid)
	router.Patch("/products/:id", middleware.UserIdHeader, middleware.UploadImageMiddleware, h.UpdateProduct)
	router.Delete("/products/:id", middleware.UserIdHeader, h.DeleteProduct)
	router.Post("/categories", middleware.UserIdHeader, h.CreateCategory)
	router.Get("/categories", middleware.UserIdHeader, h.GetCategory)
	router.Get("/categories/:id", middleware.UserIdHeader, h.GetCategoryId)
	router.Delete("/categories/:id", middleware.UserIdHeader, h.DeleteCategory)

}

func (h *shopHandler) CreateShop(c *fiber.Ctx) error {
	var (
		req = new(entity.CreateShopRequest)
		ctx = c.Context()
		v   = adapter.Adapters.Validator
		l   = middleware.GetLocals(c)
	)

	if err := c.BodyParser(req); err != nil {
		log.Warn().Err(err).Msg("handler::CreateShop - Parse request body")
		return c.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}

	req.UserId = l.UserId

	if err := v.Validate(req); err != nil {
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
		req = new(entity.GetShopRequest)
		ctx = c.Context()
		v   = adapter.Adapters.Validator
	)

	req.Id = c.Params("id")

	if err := v.Validate(req); err != nil {
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
		req = new(entity.DeleteShopRequest)
		ctx = c.Context()
		v   = adapter.Adapters.Validator
		l   = middleware.GetLocals(c)
	)
	req.UserId = l.UserId
	req.Id = c.Params("id")

	if err := v.Validate(req); err != nil {
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
		req = new(entity.UpdateShopRequest)
		ctx = c.Context()
		v   = adapter.Adapters.Validator
		l   = middleware.GetLocals(c)
	)

	if err := c.BodyParser(req); err != nil {
		log.Warn().Err(err).Msg("handler::UpdateShop - Parse request body")
		return c.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}

	req.UserId = l.UserId
	req.Id = c.Params("id")

	if err := v.Validate(req); err != nil {
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

func (h *shopHandler) CreateProduct(c *fiber.Ctx) error {
	var (
		req = new(entity.CreateProductRequest)
		ctx = c.Context()
		v   = adapter.Adapters.Validator
		l   = middleware.GetLocals(c)
	)

	if err := c.BodyParser(req); err != nil {
		log.Warn().Err(err).Msg("handler::CreateProduct - Parse request body")
		return c.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}

	req.UserId = l.UserId

	imageURL, ok := c.Locals("imageURL").(string)
	if ok && imageURL != "" {
		req.ImageURL = imageURL
	} else {
		req.ImageURL = ""
	}

	if err := v.Validate(req); err != nil {
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

func (h *shopHandler) GetAllProduct(c *fiber.Ctx) error {
	var (
		req = new(entity.GetProductRequest)
		ctx = c.Context()
		v   = adapter.Adapters.Validator
		//l   = middleware.GetLocals(c)
	)

	if err := c.QueryParser(req); err != nil {
		log.Warn().Err(err).Msg("handler::GetProduct - Parse request query")
		return c.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}

	req.SetDefault()

	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("payload", req).Msg("handler::GetShops - Validate request body")
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

func (h *shopHandler) GetProductByid(c *fiber.Ctx) error {
	var (
		req = new(entity.GetProductIdRequest)
		ctx = c.Context()
		v   = adapter.Adapters.Validator
	)

	req.Id = c.Params("id")

	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("payload", req).Msg("handler::GetProductByid - Validate request body")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	resp, err := h.service.GetProdctByid(ctx, req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.Status(fiber.StatusOK).JSON(response.Success(resp, ""))
}

func (h *shopHandler) UpdateProduct(c *fiber.Ctx) error {
	var (
		req = new(entity.UpdateProductRequest)
		ctx = c.Context()
		v   = adapter.Adapters.Validator
		l   = middleware.GetLocals(c)
	)

	if err := c.BodyParser(req); err != nil {
		log.Warn().Err(err).Msg("handler::UpdateProduct - Parse request body")
		return c.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}
	req.UserId = l.UserId
	req.Id = c.Params("id")

	if err := v.Validate(req); err != nil {
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

func (h *shopHandler) DeleteProduct(c *fiber.Ctx) error {
	var (
		req = new(entity.DeleteProductRequest)
		ctx = c.Context()
		v   = adapter.Adapters.Validator
		l   = middleware.GetLocals(c)
	)

	req.UserId = l.UserId
	req.Id = c.Params("id")

	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("payload", req).Msg("handler:DeleteProduct - Validate request body")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	err := h.service.DeleteProduct(ctx, req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.Status(fiber.StatusOK).JSON(response.Success(nil, "Product successfully deleted"))
}

func (h *shopHandler) CreateCategory(c *fiber.Ctx) error {
	var (
		req = new(entity.CreateCategoryRequest)
		ctx = c.Context()
		v   = adapter.Adapters.Validator
		l   = middleware.GetLocals(c)
	)
	if err := c.BodyParser(req); err != nil {
		log.Warn().Err(err).Msg("handler::CreateCategory - Parse request body")
		return c.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}
	req.UserId = l.UserId
	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("payload", req).Msg("handler::CreateCategory - Validate request body")
	}

	resp, err := h.service.CreateCategory(ctx, req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}
	return c.Status(fiber.StatusCreated).JSON(response.Success(resp, ""))
}

func (h *shopHandler) GetCategory(c *fiber.Ctx) error {
	var (
		req = new(entity.GetCategoryRequest)
		ctx = c.Context()
		l   = middleware.GetLocals(c)
	)

	req.UserId = l.UserId

	resp, err := h.service.GetCategory(ctx, req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}
	return c.Status(fiber.StatusOK).JSON(response.Success(resp, ""))
}

func (h *shopHandler) GetCategoryId(c *fiber.Ctx) error {
	var (
		req = new(entity.GetCategoryIdRequest)
		ctx = c.Context()
		v   = adapter.Adapters.Validator
		l   = middleware.GetLocals(c)
	)

	req.Id = c.Params("id")
	req.UserId = l.UserId

	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("payload", req).Msg("handler::GetProductByid - Validate request body")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	resp, err := h.service.GetCategoryId(ctx, req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.Status(fiber.StatusOK).JSON(response.Success(resp, ""))
}

func (h *shopHandler) DeleteCategory(c *fiber.Ctx) error {
	var (
		req = new(entity.DeleteCategoryRequest)
		ctx = c.Context()
		v   = adapter.Adapters.Validator
		l   = middleware.GetLocals(c)
	)

	req.Id = c.Params("id")
	req.UserId = l.UserId

	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("payload", req).Msg("handler::GetProductByid - Validate request body")
	}

	err := h.service.DeletCategory(ctx, req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}
	return c.Status(fiber.StatusOK).JSON(response.Success(nil, "Category successfully deleted"))
}
