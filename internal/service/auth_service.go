package service

import (
	"context"
	model "moviediary/internal/model"
	repository "moviediary/internal/repository"
	utils "moviediary/pkg/utils"
)

type AuthService struct {
	userRepository *repository.UserRepository
}

func NewAuthService(userRepository *repository.UserRepository) *AuthService {
	return &AuthService{userRepository: userRepository}
}

func (authService *AuthService) Register(ctx context.Context, username, email, password string) (*model.User, error) {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}
	user, err := authService.userRepository.CreateUser(ctx, username, email, hashedPassword)
	if err != nil {
		return nil, err
	}
	return user, nil
}
