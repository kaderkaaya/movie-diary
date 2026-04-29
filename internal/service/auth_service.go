package service

import (
	"context"
	config "moviediary/internal/config"
	model "moviediary/internal/model"
	repository "moviediary/internal/repository"
	apperror "moviediary/pkg/apperror"
	utils "moviediary/pkg/utils"
	"strings"
	"time"
)

type AuthService struct {
	userRepository *repository.UserRepository
}

func NewAuthService(userRepository *repository.UserRepository) *AuthService {
	return &AuthService{userRepository: userRepository}
}

func (authService *AuthService) Register(ctx context.Context, username, email, password string) (*model.User, error) {
	username = strings.TrimSpace(username)
	email = strings.ToLower(strings.TrimSpace(email))

	if username == "" {
		return nil, apperror.ErrUserEmpty
	}

	if email == "" {
		return nil, apperror.ErrEmailEmpty
	}

	if password == "" {
		return nil, apperror.ErrPasswordEmpty
	}

	existingUser, err := authService.userRepository.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		return nil, apperror.ErrEmailAlreadyExists
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, apperror.ErrPasswordHashError
	}

	user, err := authService.userRepository.CreateUser(ctx, username, email, hashedPassword)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (authService *AuthService) Login(ctx context.Context, email, password string) (*model.AuthResponse, error) {
	if email == "" {
		return nil, apperror.ErrEmailEmpty
	}
	if password == "" {
		return nil, apperror.ErrPasswordEmpty
	}
	user, err := authService.userRepository.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, apperror.ErrUserNotFound
	}
	userPassword := utils.VerifyPassword(user.PasswordHash, password)
	if !userPassword {
		return nil, apperror.ErrInvalidPassword
	}
	token, err := utils.GenerateJWT(user.ID, config.Load().JwtSecret, 24*time.Hour)
	if err != nil {
		return nil, err
	}

	authService.userRepository.CreateUserToken(ctx, user.ID, token)
	return &model.AuthResponse{
		User:  user,
		Token: token,
	}, nil
}
