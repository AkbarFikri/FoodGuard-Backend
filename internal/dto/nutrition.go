package dto

import "mime/multipart"

type NutritionPredictRequest struct {
	Picture *multipart.FileHeader `form:"picture" validate:"required"`
}

type NutritionPredictResponse struct {
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	UserID         string  `json:"userId"`
	Type           string  `json:"type"`
	Score          float32 `json:"score"`
	Calories       float64 `json:"calorie"`
	Carbohydrates  float64 `json:"carbohydrates"`
	Sugar          float64 `json:"sugar"`
	Fats           float64 `json:"fats"`
	Protein        float64 `json:"protein"`
	Recommendation string  `json:"recommendation"`
	CreatedAt      string  `json:"createdAt"`
}
