package service

import (
	"errors"
	"github.com/AkbarFikri/FoodGuard-Backend/internal/entity"
	"github.com/AkbarFikri/FoodGuard-Backend/internal/pkg/helper"
	"github.com/AkbarFikri/FoodGuard-Backend/internal/pkg/token"
	"github.com/AkbarFikri/FoodGuard-Backend/internal/repository"
	"golang.org/x/net/context"
	"time"
)

func (s authService) Register(ctx context.Context, user entity.User) (string, error) {
	repo, err := s.repo.NewClient(false)
	if err != nil {
		return "", err
	}

	checkUser, err := repo.User.GetByEmail(ctx, user.Email)
	if err != nil && !errors.Is(err, repository.ErrRecordWithEmailNotFound) {
		return "", err
	}

	if checkUser.Email != "" {
		return "", ErrEmailAlreadyExists
	}

	ulidID, err := helper.NewUlidFromTimestamp(time.Now())
	if err != nil {
		return "", err
	}

	hashPass, err := helper.HashPassword(user.Password)
	if err != nil {
		return "", err
	}

	user.ID = ulidID
	user.Password = hashPass

	if err := repo.User.Insert(ctx, user); err != nil {
		return "", err
	}

	userClaims := map[string]interface{}{
		"email":    user.Email,
		"id":       user.ID,
		"username": user.Username,
	}
	accessToken, err := token.Sign(userClaims, 24*time.Hour)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (s authService) Login(ctx context.Context, user entity.User) (string, error) {
	repo, err := s.repo.NewClient(false)
	if err != nil {
		return "", err
	}

	checkUser, err := repo.User.GetByEmail(ctx, user.Email)
	if err != nil {
		if errors.Is(err, repository.ErrRecordWithEmailNotFound) {
			return "", ErrEmailOrPasswordWrong
		}
		return "", err
	}

	if err := helper.ComparePassword(checkUser.Password, user.Password); err != nil {
		return "", ErrEmailOrPasswordWrong
	}

	userClaims := map[string]interface{}{
		"email":    checkUser.Email,
		"id":       checkUser.ID,
		"username": checkUser.Username,
	}
	accessToken, err := token.Sign(userClaims, 24*time.Hour)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
