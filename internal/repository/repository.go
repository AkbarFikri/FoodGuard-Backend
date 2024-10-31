package repository

import (
	"context"
	"github.com/AkbarFikri/FoodGuard-Backend/internal/entity"
	"gorm.io/gorm"
)

func New(db *gorm.DB) Repository {
	return &repository{
		DB: db,
	}
}

type repository struct {
	DB *gorm.DB
}

type Repository interface {
	NewClient(tx bool) (Client, error)
}

func (r *repository) NewClient(tx bool) (Client, error) {
	var db *gorm.DB
	var commitFunc, rollbackFunc func() error

	db = r.DB

	if tx {
		txx := r.DB.Begin()
		if txx.Error != nil {
			return Client{}, txx.Error
		}

		db = txx
		commitFunc = func() error { return txx.Commit().Error }
		rollbackFunc = func() error { return txx.Rollback().Error }
	} else {
		commitFunc = func() error { return nil }
		rollbackFunc = func() error { return nil }
	}

	return Client{
		User:     &userRepository{q: db},
		Commit:   commitFunc,
		Rollback: rollbackFunc,
	}, nil
}

type Client struct {
	User interface {
		GetByEmail(ctx context.Context, email string) (entity.User, error)
		Insert(ctx context.Context, user entity.User) error
	}
	Commit   func() error
	Rollback func() error
}

type userRepository struct {
	q *gorm.DB
}
