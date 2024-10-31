package postgres

import (
	"fmt"
	"github.com/AkbarFikri/FoodGuard-Backend/internal/pkg/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewInstance() (*gorm.DB, error) {
	DBName := env.GetString("DB_NAME", "")
	DBUser := env.GetString("DB_USER", "")
	DBPassword := env.GetString("DB_PASSWORD", "")
	DBHost := env.GetString("DB_HOST", "")
	DBPort := env.GetString("DB_PORT", "")
	DBSslMode := env.GetString("DB_SSLMODE", "")

	DBDSN := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		DBHost, DBUser, DBPassword, DBName, DBPort, DBSslMode,
	)

	db, err := gorm.Open(postgres.Open(DBDSN), &gorm.Config{})
	if err != nil {
		panic("Failed connect to database")
	}

	connection, err := db.DB()
	if err != nil {
		panic("Failed connect to database")
	}

	connection.SetMaxIdleConns(5)
	connection.SetMaxOpenConns(50)

	return db, nil
}
