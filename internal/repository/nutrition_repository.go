package repository

import (
	"github.com/AkbarFikri/FoodGuard-Backend/internal/entity"
	"golang.org/x/net/context"
)

func (r *nutritionRepository) Insert(ctx context.Context, nutrition entity.Nutrition) error {
	if err := r.q.WithContext(ctx).Create(&nutrition).Error; err != nil {
		return err
	}
	return nil
}

func (r *nutritionRepository) GetAllByUserID(ctx context.Context, userID string) ([]entity.Nutrition, error) {
	var nutritions []entity.Nutrition
	if err := r.q.WithContext(ctx).Find(&nutritions, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}
	return nutritions, nil
}
