package service

import (
	"github.com/AkbarFikri/FoodGuard-Backend/internal/dto"
	"github.com/AkbarFikri/FoodGuard-Backend/internal/pkg/helper"
	"golang.org/x/net/context"
	"io"
	"mime/multipart"
	"time"
)

func (s nutritionService) GeneratePrediction(ctx context.Context, request dto.NutritionPredictRequest, userID string) (dto.NutritionPredictResponse, error) {
	repo, err := s.repo.NewClient(false)
	if err != nil {
		return dto.NutritionPredictResponse{}, err
	}

	file, err := request.Picture.Open()
	if err != nil {
		return dto.NutritionPredictResponse{}, err
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return dto.NutritionPredictResponse{}, err
	}

	prediction, err := s.googleService.Gemini.GeneratePredictionFromFood(ctx, fileBytes)
	if err != nil {
		return dto.NutritionPredictResponse{}, err
	}

	if prediction.Name == "notfood" {
		return dto.NutritionPredictResponse{}, ErrPictureNotFood
	}

	entityNutrition := formattedNutritionResponseToEntity(prediction)

	ulid, err := helper.NewUlidFromTimestamp(time.Now())
	if err != nil {
		return dto.NutritionPredictResponse{}, err
	}

	entityNutrition.ID = ulid
	entityNutrition.UserID = userID
	entityNutrition.CreatedAt = time.Now().UTC()

	if err := repo.Nutrition.Insert(ctx, entityNutrition); err != nil {
		return dto.NutritionPredictResponse{}, err
	}

	prediction.UserID = entityNutrition.UserID
	prediction.CreatedAt = entityNutrition.CreatedAt.String()

	return prediction, nil
}

func (s nutritionService) GetAllUserNutrition(ctx context.Context, userID string) ([]dto.NutritionPredictResponse, error) {
	repo, err := s.repo.NewClient(false)
	if err != nil {
		return []dto.NutritionPredictResponse{}, err
	}

	nutritions, err := repo.Nutrition.GetAllByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	res := make([]dto.NutritionPredictResponse, len(nutritions))
	for i, nutrition := range nutritions {
		res[i] = formattedNutritions(nutrition)
	}

	return res, nil
}
