package google

import (
	"github.com/AkbarFikri/FoodGuard-Backend/internal/dto"
	"github.com/AkbarFikri/FoodGuard-Backend/internal/pkg/env"
	"github.com/google/generative-ai-go/genai"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

type GoogleService struct {
	Gemini interface {
		GeneratePredictionFromFood(ctx context.Context, photo []byte) (dto.NutritionPredictResponse, error)
	}
}

func New() GoogleService {

	geminiKey := env.GetString("GEMINI_API_KEY", "")
	typeModel := env.GetString("GEMINI_TYPE_MODEL", "")
	client, err := genai.NewClient(context.Background(), option.WithAPIKey(geminiKey))
	if err != nil {
		panic(err)
	}

	model := client.GenerativeModel(typeModel)
	return GoogleService{
		Gemini: googleGemini{
			model: model,
		},
	}
}

type googleGemini struct {
	model *genai.GenerativeModel
}
