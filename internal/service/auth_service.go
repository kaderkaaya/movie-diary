package service

import (
	"context"
	model "moviediary/internal/model"
	repository "moviediary/internal/repository"
)

type AuthService struct {
	userRepository *repository.UserRepository
}

func NewAuthService(userRepository *repository.UserRepository) *AuthService {
	return &AuthService{userRepository: userRepository}
}

func (authService *AuthService) Register(ctx context.Context, username, email, password string) (*model.User, error) {
	user, err := authService.userRepository.CreateUser(ctx, username, email, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
