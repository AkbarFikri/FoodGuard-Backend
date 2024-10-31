package entity

import "time"

type Nutrition struct {
	ID             string
	Name           string
	Type           string
	Carbohydrate   float64
	Sugar          float64
	Calorie        float64
	Fat            float64
	Protein        float64
	Recommendation string
	CreatedAt      time.Time
}
