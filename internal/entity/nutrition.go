package entity

import "time"

type Nutrition struct {
	ID             string
	Name           string
	UserID         string
	Type           string
	Score          float32
	Carbohydrate   float64
	Sugar          float64
	Calorie        float64
	Fat            float64
	Protein        float64
	Recommendation string
	CreatedAt      time.Time
}
