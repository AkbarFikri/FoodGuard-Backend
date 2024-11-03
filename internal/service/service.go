package service

import (
	"context"
	"github.com/AkbarFikri/FoodGuard-Backend/internal/dto"
	"github.com/AkbarFikri/FoodGuard-Backend/internal/entity"
	"github.com/AkbarFikri/FoodGuard-Backend/internal/repository"
	"github.com/AkbarFikri/FoodGuard-Backend/pkg/google"
)

type Client struct {
	Auth      AuthService
	Nutrition NutritionService
}

func New(repo repository.Repository, googles google.GoogleService) Client {
	return Client{
		Auth:      authService{repo: repo},
		Nutrition: nutritionService{repo: repo, googleService: googles},
	}
}

type AuthService interface {
	Register(ctx context.Context, user entity.User) (string, error)
	Login(ctx context.Context, user entity.User) (string, error)
}

type authService struct {
	repo repository.Repository
}

type NutritionService interface {
	GeneratePrediction(ctx context.Context, request dto.NutritionPredictRequest, userID string) (dto.NutritionPredictResponse, error)
	GetAllUserNutrition(ctx context.Context, userID string) ([]dto.NutritionPredictResponse, error)
}

type nutritionService struct {
	repo          repository.Repository
	googleService google.GoogleService
}
