package repository

import (
	"github.com/AkbarFikri/FoodGuard-Backend/internal/pkg/response"
	"net/http"
)

// Error from user repository
var (
	ErrRecordWithEmailNotFound = response.NewError(http.StatusNotFound, "record with email not found")
)
