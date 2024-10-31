package service

import (
	"context"
	"github.com/AkbarFikri/FoodGuard-Backend/internal/entity"
	"github.com/AkbarFikri/FoodGuard-Backend/internal/repository"
)

type Client struct {
	Auth AuthService
}

func New(repo repository.Repository) Client {
	return Client{
		Auth: authService{repo: repo},
	}
}

type AuthService interface {
	Register(ctx context.Context, user entity.User) (string, error)
	Login(ctx context.Context, user entity.User) (string, error)
}

type authService struct {
	repo repository.Repository
}
