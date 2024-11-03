package main

import (
	"github.com/AkbarFikri/FoodGuard-Backend/internal/config"
	"github.com/AkbarFikri/FoodGuard-Backend/internal/handler"
	"github.com/AkbarFikri/FoodGuard-Backend/internal/middleware"
	"github.com/AkbarFikri/FoodGuard-Backend/internal/repository"
	"github.com/AkbarFikri/FoodGuard-Backend/internal/service"
	"github.com/AkbarFikri/FoodGuard-Backend/pkg/google"
	"github.com/AkbarFikri/FoodGuard-Backend/pkg/postgres"
	"github.com/joho/godotenv"
)

func main() {
	logger := config.NewLogger()
	validator := config.NewValidator()

	if err := godotenv.Load(); err != nil {
		logger.Fatalf("Error loading .env file")
	}

	db, err := postgres.NewInstance()
	if err != nil {
		logger.Fatal(err)
	}

	app := config.NewFiber(logger)
	router := app.Group("/api/v1")

	middlewares := middleware.New(logger)
	googles := google.New()

	repo := repository.New(db)
	srv := service.New(repo, googles)
	handl := handler.New(srv, router, validator, middlewares)

	if err := handl.RegisterHandler(); err != nil {
		logger.Fatal(err)
	}

	logger.Fatal(app.Listen("0.0.0.0:9090"))
}
