package handler

import (
	"github.com/AkbarFikri/FoodGuard-Backend/internal/dto"
	"github.com/AkbarFikri/FoodGuard-Backend/internal/entity"
	"github.com/AkbarFikri/FoodGuard-Backend/internal/middleware"
	"github.com/AkbarFikri/FoodGuard-Backend/internal/pkg/helper"
	"github.com/AkbarFikri/FoodGuard-Backend/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func newAuthHandler(service service.AuthService, validate *validator.Validate, middleware middleware.Middleware) *authHandler {
	return &authHandler{
		service:    service,
		validate:   validate,
		middleware: middleware,
	}
}

func (h *authHandler) Start(srv fiber.Router) {
	auth := srv.Group("/auth")

	auth.Post("/register", h.HandleRegister)
	auth.Post("/login", h.HandleLogin)

	user := srv.Group("/user")

	user.Get("/current", h.middleware.NewtokenMiddleware, h.HandleCurrentUser)
}

func (h *authHandler) HandleRegister(c *fiber.Ctx) error {
	var req dto.RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := h.validate.Struct(&req); err != nil {
		return err
	}

	userReq := entity.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}

	res, err := h.service.Register(c.Context(), userReq)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"accessToken": res,
	})
}

func (h *authHandler) HandleLogin(c *fiber.Ctx) error {
	var req dto.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := h.validate.Struct(&req); err != nil {
		return err
	}

	userReq := entity.User{
		Email:    req.Email,
		Password: req.Password,
	}

	res, err := h.service.Login(c.Context(), userReq)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"accessToken": res,
	})
}

func (h *authHandler) HandleCurrentUser(c *fiber.Ctx) error {
	user, err := helper.GetUserFromContext(c)
	if err != nil {
		return fiber.ErrUnauthorized
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"username": user.Username,
		"email":    user.Email,
		"id":       user.ID,
	})
}
