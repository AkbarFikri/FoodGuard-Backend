package handler

import (
	"github.com/AkbarFikri/FoodGuard-Backend/internal/service"
	"github.com/gofiber/fiber/v2"
)

func newAuthHandler(service service.AuthService) *authHandler {
	return &authHandler{
		service: service,
	}
}

func (h *authHandler) Start(srv fiber.Router) {
	auth := srv.Group("/auth")

	auth.Post("/register")
	auth.Post("/login")
}
