package handler

import (
	"github.com/AkbarFikri/FoodGuard-Backend/internal/dto"
	"github.com/AkbarFikri/FoodGuard-Backend/internal/middleware"
	"github.com/AkbarFikri/FoodGuard-Backend/internal/pkg/helper"
	"github.com/AkbarFikri/FoodGuard-Backend/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func newNutritionHandler(service service.NutritionService,
	validate *validator.Validate, middleware middleware.Middleware) *nutritionHandler {
	return &nutritionHandler{
		service:    service,
		validate:   validate,
		middleware: middleware,
	}
}

func (h *nutritionHandler) Start(srv fiber.Router) {
	nutrition := srv.Group("/nutritions")

	nutrition.Use(h.middleware.NewtokenMiddleware)

	nutrition.Post("/predic", h.NutritionPicture)
	nutrition.Get("", h.HandleGetNutrition)
}

func (h *nutritionHandler) NutritionPicture(c *fiber.Ctx) error {
	var req dto.NutritionPredictRequest

	file, err := c.FormFile("picture")
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	req.Picture = file

	user, err := helper.GetUserFromContext(c)
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	res, err := h.service.GeneratePrediction(c.UserContext(), req, user.ID)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *nutritionHandler) HandleGetNutrition(c *fiber.Ctx) error {
	user, err := helper.GetUserFromContext(c)
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	res, err := h.service.GetAllUserNutrition(c.UserContext(), user.ID)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
