package service

import (
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
}

type authService struct {
	repo repository.Repository
}
