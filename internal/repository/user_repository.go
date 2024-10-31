package repository

import (
	"errors"
	"github.com/AkbarFikri/FoodGuard-Backend/internal/entity"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

func (r *userRepository) GetByEmail(ctx context.Context, email string) (entity.User, error) {
	var user entity.User
	if err := r.q.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.User{}, ErrRecordWithEmailNotFound
		}
		return entity.User{}, err
	}
	return user, nil
}

func (r *userRepository) Insert(ctx context.Context, user entity.User) error {
	if err := r.q.WithContext(ctx).Create(&user).Error; err != nil {
		return err
	}
	return nil
}
