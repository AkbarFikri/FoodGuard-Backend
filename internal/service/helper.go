package service

import (
	"github.com/AkbarFikri/FoodGuard-Backend/internal/dto"
	"github.com/AkbarFikri/FoodGuard-Backend/internal/entity"
	"strconv"
)

// helper for nutrition server
func formattedNutritionResponseToEntity(data dto.NutritionPredictResponse) entity.Nutrition {
	return entity.Nutrition{
		ID:             data.ID,
		Name:           data.Name,
		Score:          data.Score,
		Protein:        data.Protein,
		Carbohydrate:   data.Carbohydrates,
		Calorie:        data.Calories,
		Fat:            data.Fats,
		Sugar:          data.Sugar,
		Type:           data.Type,
		Recommendation: data.Recommendation,
	}
}

func formattedNutritions(data entity.Nutrition) dto.NutritionPredictResponse {
	return dto.NutritionPredictResponse{
		ID:             data.ID,
		Name:           data.Name,
		Score:          data.Score,
		Protein:        data.Protein,
		Sugar:          data.Sugar,
		Type:           data.Type,
		Recommendation: data.Recommendation,
		Carbohydrates:  data.Carbohydrate,
		Calories:       data.Calorie,
		Fats:           data.Calorie,
		CreatedAt:      strconv.FormatInt(data.CreatedAt.UnixMilli(), 10),
	}
}
