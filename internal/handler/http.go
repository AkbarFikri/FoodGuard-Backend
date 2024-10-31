package handler

import (
	"github.com/AkbarFikri/FoodGuard-Backend/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Handler struct {
	router   fiber.Router
	handlers []Handlers
}

type Handlers interface {
	Start(srv fiber.Router)
}

func New(client service.Client, router fiber.Router) *Handler {
	var handlers []Handlers

	authHandler := newAuthHandler(client.Auth)

	handlers = append(handlers, authHandler)
	return &Handler{
		router: router,
	}
}

func (h *Handler) RegisterHandler() error {
	h.router.Use(cors.New())
	h.router.Use(logger.New())

	for _, handler := range h.handlers {
		handler.Start(h.router)
	}

	return nil
}

type authHandler struct {
	service service.AuthService
}
