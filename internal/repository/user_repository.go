package repository

import (
	"context"
	model "moviediary/internal/model"

	"time"

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

func (userRepository *UserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	if err := userRepository.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (userRepository *UserRepository) CreateUserToken(ctx context.Context, UserID uint, token string) (*model.UserTokens, error) {
	userToken := &model.UserTokens{
		UserID:    UserID,
		Token:     token,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}
	if err := userRepository.db.WithContext(ctx).Create(userToken).Error; err != nil {
		return nil, err
	}
	return userToken, nil
}
