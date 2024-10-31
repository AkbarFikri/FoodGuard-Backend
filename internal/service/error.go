package service

import (
	"github.com/AkbarFikri/FoodGuard-Backend/internal/pkg/response"
	"net/http"
)

// error from auth service
var (
	ErrEmailAlreadyExists   = response.NewError(http.StatusConflict, "email already exists")
	ErrEmailOrPasswordWrong = response.NewError(http.StatusBadRequest, "email or password wrong")
)
