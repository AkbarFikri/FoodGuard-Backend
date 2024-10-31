package main

import (
	"github.com/AkbarFikri/FoodGuard-Backend/internal/config"
	"github.com/AkbarFikri/FoodGuard-Backend/internal/handler"
	"github.com/AkbarFikri/FoodGuard-Backend/internal/repository"
	"github.com/AkbarFikri/FoodGuard-Backend/internal/service"
	"github.com/AkbarFikri/FoodGuard-Backend/pkg/postgres"
	"github.com/joho/godotenv"
)

func main() {
	logger := config.NewLogger()

	if err := godotenv.Load(); err != nil {
		logger.Fatalf("Error loading .env file")
	}

	db, err := postgres.NewInstance()
	if err != nil {
		logger.Fatal(err)
	}

	app := config.NewFiber(logger)
	router := app.Group("/api/v1")

	repo := repository.New(db)
	srv := service.New(repo)
	handl := handler.New(srv, router)

	if err := handl.RegisterHandler(); err != nil {
		logger.Fatal(err)
	}

	logger.Fatal(app.Listen(":8080"))
}
