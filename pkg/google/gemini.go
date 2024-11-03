package google

import (
	"encoding/json"
	"github.com/AkbarFikri/FoodGuard-Backend/internal/dto"
	"github.com/google/generative-ai-go/genai"
	"golang.org/x/net/context"
	"strconv"
)

func (g googleGemini) GeneratePredictionFromFood(ctx context.Context, photo []byte) (dto.NutritionPredictResponse, error) {
	content, err := g.model.GenerateContent(ctx,
		genai.Text("I am providing an image of a food item. Please analyze the image and provide a JSON response containing the following details: the name of the food, the type of food, a confidence score for this prediction (as a float from 0 to 10, with one decimal place, where a higher score indicates a well-balanced and nutritionally favorable food item), the number of calories (in kcal), carbohydrates, sugar, fats, and protein (all in grams), and a health recommendation based on these values. If the item is not food, set 'name' to \"notfood\"."),
		genai.Text("Format the JSON response exactly like this example: {\"name\": \"Example Food Name\", \"type\": \"Example Food Type\", \"score\": 8.5, \"calorie\": 123.0, \"carbohydrates\": 45.0, \"sugar\": 12.0, \"fats\": 20.0, \"protein\": 15.0, \"recommendation\": \"Considering the high sugar content, it is advisable to choose foods with lower sugar levels for a balanced diet.\"}"),
		genai.Text("For the recommendation, provide a brief health summary based on the nutritional values. If there is a high value in sugar, fats, or calories, suggest alternatives or moderation for better health."),
		genai.Text("Give me only the JSON output in one-line, without anything else."),
		genai.Text(".Answer it in indonesian language"),
		genai.ImageData("jpeg", photo),
	)
	if err != nil {
		return dto.NutritionPredictResponse{}, err
	}

	part := content.Candidates[0].Content.Parts[0]
	byteJson, err := json.Marshal(part)
	if err != nil {
		return dto.NutritionPredictResponse{}, err
	}

	strJson, err := strconv.Unquote(string(byteJson))
	if err != nil {
		return dto.NutritionPredictResponse{}, err
	}

	var res dto.NutritionPredictResponse
	err = json.Unmarshal([]byte(strJson), &res)
	if err != nil {
		return dto.NutritionPredictResponse{}, err
	}

	return res, nil
}
