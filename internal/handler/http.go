package handler

import (
	"github.com/AkbarFikri/FoodGuard-Backend/internal/middleware"
	"github.com/AkbarFikri/FoodGuard-Backend/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Handler struct {
	router     fiber.Router
	handlers   []Handlers
	middleware middleware.Middleware
}

type Handlers interface {
	Start(srv fiber.Router)
}

func New(client service.Client,
	router fiber.Router,
	validate *validator.Validate, middleware middleware.Middleware) *Handler {
	var handlers []Handlers

	authHandler := newAuthHandler(client.Auth, validate)
	nutritionHandler := newNutritionHandler(client.Nutrition, validate, middleware)

	handlers = append(handlers, authHandler, nutritionHandler)
	return &Handler{
		router:     router,
		handlers:   handlers,
		middleware: middleware,
	}
}

func (h *Handler) RegisterHandler() error {

	h.router.Use(h.middleware.NewRateLimitter)
	h.router.Use(cors.New())
	h.router.Use(logger.New())

	for _, handler := range h.handlers {
		handler.Start(h.router)
	}

	return nil
}

type authHandler struct {
	service  service.AuthService
	validate *validator.Validate
}

type nutritionHandler struct {
	service    service.NutritionService
	validate   *validator.Validate
	middleware middleware.Middleware
}
