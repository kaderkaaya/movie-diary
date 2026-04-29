package repository

import (
	"context"
	model "moviediary/internal/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (userRepository *UserRepository) CreateUser(ctx context.Context, username, email, password string) (*model.User, error) {
	user := &model.User{
		Username:     username,
		Email:        email,
		PasswordHash: password,
	}
	if err := userRepository.db.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
